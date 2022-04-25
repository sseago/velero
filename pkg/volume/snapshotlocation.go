/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package volume

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/vmware-tanzu/velero/internal/credentials"
	velerov1api "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
)

// If the VSL specifies a credential, add credential file path to config
func UpdateVolumeSnapshotLocationWithCredentialConfig(location *velerov1api.VolumeSnapshotLocation, credentialStore credentials.FileStore, logger logrus.FieldLogger) error {
	// add the bucket name and prefix to the config map so that object stores
	// can use them when initializing. The AWS object store uses the bucket
	// name to determine the bucket's region when setting up its client.
	if location.Spec.Config == nil {
		location.Spec.Config = make(map[string]string)
	}
	// If the VSL specifies a credential, fetch its path on disk and pass to
	// plugin via the config.
	if location.Spec.Credential != nil && credentialStore != nil {
		credsFile, err := credentialStore.Path(location.Spec.Credential)
		if err != nil {
			return errors.Wrap(err, "unable to get credentials")
		}

		location.Spec.Config["credentialsFile"] = credsFile
	}
	return nil
}
