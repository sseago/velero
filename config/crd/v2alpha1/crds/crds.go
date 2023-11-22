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

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcY_s\xe3\xb8\r\x7fϧ\xc0l\x1f\xf2\xb2Rn\xaf\x9dN\xc7o\xbbN;\x93\xe9m\xea\xb9\xec\xe4\x9d\x12a\x99\x17\x8adI\xc8i\xda\xe9w\uf014d\xfd\xa1\xe3d\xafwz3\t\x82?\xfc\x00\x02 ]\x14ŕp\xea\x11}P\xd6l@8\x85\xff\"4\xfc+\x94O\x7f\t\xa5\xb27\xc7OWO\xca\xc8\rl\xbb@\xb6\xfd\x19\x83\xed|\x8d\xb7\xb8WF\x91\xb2\xe6\xaaE\x12R\x90\xd8\\\x01\bc,\t\x1e\x0e\xfc\x13\xa0\xb6\x86\xbc\xd5\x1a}Ѡ)\x9f\xba\n\xabNi\x89>*\x1f\xb6>\xfeP~\xfa\xb1\xfc\xe1\n\xc0\x88\x167\xc0\xfa\xa4}6\xda\n\x19\xca#j\xf4\xb6T\xf6*8\xacYq\xe3m\xe76p\x9aH\v\xfbM\x13\xe0[A\xe2\xb6\xd7\x11\x87\xb5\n\xf4\xf7\xd5\xd4O*P\x9cv\xba\xf3B/\xf6\x8e3A\x99\xa6\xd3\xc2\xcf\xe7\xae\x00Bm\x1dn\xe0\x9e\xb7v\xa2F\x1e\xebm\x8aP\n\x10RF\x96\x84\xdeye\b\xfd\xd6\xea\xae\x1d\xd8)@b\xa8\xbdr\x14Y\x98\u0082@\x82\xba\x00\xa1\xab\x0f \x02\xdc\xe3\xf3͝\xd9y\xdbx\f\t\x16\xc0/\xc1\x9a\x9d\xa0\xc3\x06\xca$^\xba\x83\b\xd8\xcf&*\x1f\xe2D?D/\x8c7\x90W\xa6\xc9!\xf8\xa6Z\x04\xd9\xf9\xe8B\xb6\xbbF\xa0\x83\nsh\xcf\"0<O(\xcf\x02\x89\xf3\xac.\x90h\xdd\x12\xd1di\x82$\x05a\x0e\xd0ֶN#\xa1\x84\xea\x85p0co}+h\x03\xcaП\xfft\x9e\x8b\x9e\xac2.\xbd\xb5fN\xcc\x17\x1e\x85\xc9pB\xc2^j\xd0gٱ$\xf4\xaf\x01B\xac\xe0\xcbd}B\x92\xf4N\xc7/B\xe1\x90\x03\xbb\a: |\x11\xf5S\xe7\xe0\x81\xac\x17\r\xc2O\xb6N\xee{>\xa0\xc7(Q%\t\x8e^P\xec;볮sX\x97I\xb6W6\xe8Z\xf8o\xbe\xd1\xff=\xb6j\x8f\"\x1b[C\xaa)\xa3\x84\xb2&\x1f`\x9f\x1b|SpMI4Vℱ\x19&\x15\xc0y[c\b\xaf\x04<+\x98\xa1\xb8?\r\xac\xa8I\x12\xc7\x1f\x85v\a\xf1)%\x99\xfa\x80\xad\xd8\xf4+\xacC\xf3yw\xf7\xf8Ǉ\xd90\xbc\x920DM\x813\x05\xc3wޒ\xad\xad\x86\n\xe9\x19\xd1$\u05f7\xf6\x88\x9e\xf3\\\xa3L\x185r֖S\x81S\xce\xe6\xf8\x8e\xfax6Mz\x8c\xd1\xc3\x00\xfd\xd4\xfb\xc0{:\xf4\xa4\x86,\xdc\xeb>\x15\x98\xc9\xe8\u008ek65I\x81\xe4ʂɌ>\x97\xa2\xec\xd9I\xceR\x01<:\x8f\x01\r\xcd!\xf4\xdc\xedA\x18\xb0\xd5/XS\t\x0f\xe8Y\r\x84\x83\xed\xb4d\xe3\x8e\xe8\t<ֶ1\xeaߣ\xee\x00d\xe3\xa6Z\x10\xf6%\xe1\xf4\xc5\xdcm\x84\x86\xa3\xd0\x1d~\x8c\x94\xb5\xe2\x05<\xf2.Й\x89\xbe(\x12J\xf8\xca<)\xb3\xb7\x1b8\x10\xb9\xb0\xb9\xb9i\x14\r\x85\xb5\xb6m\xdb\x19E/7\x91oUud}\xb8\x91xD}\x13TS\b_\x1f\x14aM\x9d\xc7\x1b\xe1T\x11\xa1\x9bX\\\xcbV\xfe\xc1\xf7\xa58\\ϰ\xaeb-}\xb1&\xbe\xe2\x01.\x8c\x1c\xe8\xa2_\x9a\xac8\x11\xcdC\xcc\xce\xcf\x7f}\xf8\x06\xc3\xd6\xd1\x19K\xf6#暈\xe1\xe4\x02&L\x99=\xfa\xe4Ľ\xb7mԉF:\xab\f\xc5\x1f\xb5Vh\x96\xf4\x87\xaej\x15\xb1\xdf\xff\xd9a \xf6U\t\xdb\xd8m@\x85\xd09>Ⲅ;\x03[Ѣފ\x80\xbf\xb9\x03\x98\xe9P0\xb1os\xc1\xb4QZ\n'\xd6&\x13C\xa7s\xc6_ӓ\xff\xe0\xb0f\xd71{\xbcL\xedU_\x01\xf8\xf8\x8a\x99l9S\x99?\xb2\xfce\xab\xc0Rh\x81\xe9Kn\xcd\x00\xccLrm_\x8eB\x92\\)\x05\xd0gK\x98Gg\x83\"\xeb_N\x85\xac\\i8\xe3\x00\xfejaj\xd4\x17,\xd9F!PF2\x938\xc6\x1d\xa7\x88\xa4 b\xb2\xa6\xb1|.\xce\x13\x9c\xbe;\xe2U\x1c\xa8\x01\x89m2\xd9\x1a\xa3\f\x9c:<\x98vrK\xcb*k5\x8ae\xde\xe3\xd8\xfa\xcaIzk\xcd^5k\x1b\xa7\xcd\xe89\xc7_\xa0/\x13\x86\x93-\xd9\n\x8e9FR\xc4zQ\f\x01ɉw\xaf\x9a\xce\xe726\x7f{\x85Z\x86s\xbe\\\x9d\x8f\xc1\xe0\xb8\xcb\x05w\x8e(\x87\xe3ї\x97I\xcd#\x1b\xf3H\x88\x8d&Of\x10\xa6\x10,\xe1n?Ѩ\x02|\xf8\x00\xd6Çt\x19\xf9\xf01\x85k\xa74\x15jZx3\x1a\x9f\x95\xd6þ\xef\x8a\xe2\xb1\xfar\x03d;\xba@\xc0?\x16\xe2\v\x1e\x88;\xb3h;Yx\x16\x8a\xc6r\x97\xc1<n\x1d>B\x85{N\xb1\x1e\xa9\xf3\x86O\x02z\xcf)'D\x95\xb6\xa3w\x19\x15\x8cp\xe1`\xe9\xee\xf6\x829\x0f\xa3\xe0\x90]\xeen\x87\xdc\xf2\x18\xbd0\xa6\x98^\x12\xc8\xe6\x1c\x8aC;#c1z\x1f\xdaX\x01ǫ\xdf%\xc8s\xe9\x01\xb7\xf5\xaaQ\xdcV\x98q\xe6\x94\xf2\x8e|U\xcc\x05\xa2\n\xd1>\x94й\x04\x9cS\fW\xd7\nA\xaa\xfd\x1e=\x1aJ\xf55m\xbc{\xdc^\x87\xd3&9\x9d\xfb\t\x86\xd8a\xb5\xc29\x94ܱ\xb3g{\xa2\xdeE\x11\t\xdf =F3.\xf0\xf3m\":\x90Õ\x9b\xafW\\\bz\xef&\x8d\xb0{\xdcr\a\x961c\xf7\xb8Fx\xbe\xcaAߊ\x9f\xf1\xe0\n\xe5\xca\x7f=\x9e\u05c8\xbd\x90N\x01\xdc\xf1\r;\xef\x1es\x85t\xa4\x03\xe8 \x88%\xfa\xab\x13T/Y\x9d0\x9c\x8fޝ߇\xb7~\x13\xe0\xed\xab\x88\xb7K\xc8g\xf0V/\xbf\x1a2\x17o\xe5Q\xaeQ\x17\xafx\xae\x00w\xcc\x0e\xd6o/Q\xf9\x9d\x8b|w\xb5\x90Y\xa6\xf8\xc5\xf4)Y.'\xe6\x99f1;=\x92ojC\xe3\xe5\xf6\xad\x8dhz\xb2\xea\xdd^w>\xa6\xa1\xfe!\x8boe\xdfՊ\xd6\xe9\thzۿԾ\xadW\xc4\xfb\x9e\x97\x93z'ƫl|r\x18ޙr\xfd\xdbI_Z\x1a\xd3#\xabC\txD\x03\xdcj\v\xa5Q\x0e:C\t߸\x1b\x8f\x17\x9f\xeb\xe5\x15)\xf2\xdd+\x8ae\x97{\xa6\f\xe8\xf5\xbaὉ\xaf;\x05\xabXI\x98NkQi\xdc\x00\xf9\xee\\\xff\x98=(-\x86 \x9aK\x89\xfak\x92JW\xc5~\t\x88\x8a\x9b\x8aeO{\x1dz߿\xabh\x18+/a\xb8\xb72\x020\xdf\xf1~\xf3.,\xb1\a\xbf\x00f\xc72\xb9\x98\x1f\xa1\xbd~=@ӵ\xb9\xcct\x8fϙ\xd1\xcfu\x8d.\x97-\v\xd8yt\xc2g\xa7V\x0f\xc7\xd3\xc9t\xd9\xc9%\xcea.\xabs|\x99\xcd\xcc\xfd-\x1e\x86w1\xdd\xe3\xbbD\xf6p;:X=\x1c\xe6\xf8xj\xba\xb6Bό\xc7\xe7ف\xfa!Kf\x0e\xa00r沓\x86\xb1'\x8c\xaa\xf8$s\x95J\x17\xb8\xa1K\x96*8-rEv\xb0d־\x9c\x0e\xc8\xea\xfd\xec\xbd\xfd\xca\xf8\x98\x9d/¹\x17\xe9\x9c\x17\xa6oˋ\xf9\xf1\x91\xfa\xb7\xd9\xe1\x95\v\xdd\xfcO\x83K-\xf5L\xf8R\x82\xef\xff\xafȥ\xf7i\xa6^\xe7\xe5\xf96\xbfgJ\xce\x12\xb5\x1a\x8c\xc8\xe5Dw\xff\xac2\x1d\xe9\xaa\xf1\xb1p\x03\xff\xf9\xef\xd5\xff\x02\x00\x00\xff\xffG\x0e\xcf\xec\xfa\x1b\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcYK\x93۸\x11\xbeϯ\xe8r\x0e\xbe\x8c8\xebM*\x95\xd2\xcd\xd6d\xabTYOT\x96w\xee\x10\xd9\"\xb1\x03\x02\f\x1e\x9aLR\xf9\xef\xa9\x06\b\n$\xa1\x97c\x87\x87\xa9\x11\x80nt7\xd0_?\xb0X,\xeeXǟQ\x1b\xae\xe4\x12X\xc7\xf1\x9f\x16%\xfd2\xc5\xcb_L\xc1\xd5\xc3\xe1\xc3\xdd\v\x97\xd5\x12V\xceX\xd5~A\xa3\x9c.\xf1\x11\xf7\\r˕\xbckѲ\x8aY\xb6\xbc\x03`R*\xcbh\xd8\xd0O\x80RI\xab\x95\x10\xa8\x175\xca\xe2\xc5\xedp縨P{\xe6q\xeb\xc3OŇ\x9f\x8b\x9f\xee\x00$kq\t\xc4\xcfuB\xb1\xca\x14\a\x14\xa8U\xc1՝\xe9\xb0$\xb6\xb5V\xae[\xc2q\"\x90\xf5[\x06q\x1f\x99e\xbfy\x0e~Ppc\xff6\x99\xf8\x95\x1b\xeb';\xe14\x13\xa3]\xfd\xb8\xe1\xb2v\x82\xe9t\xe6\x0e\xc0\x94\xaa\xc3%<і\x1d+\x91\xc6zM\xbc\b\v`U\xe5m\xc3\xc4FsiQ\xaf\x94pm\xb4\xc9\x02*4\xa5\xe6\x9d\xf5\xba\x1f\x05\x02c\x99u\x06\x8c+\x1b`\x06\x9e\xf0\xf5a-7Z\xd5\x1aM\x10\t\xe0w\xa3\xe4\x86\xd9f\tEX^t\r3\xd8\xcf\x06\xf3m\xfdD?d\xdfHZc5\x97un\xff\xaf\xbcE\xa8\x9c\xf6\xc7F:\x97\b\xb6\xe1&\x15\xec\x95\x19\x12N[\xacN\x8a\xe1牙\xb1\xac\xed\xa6\xf2$\xa4A\xa0\x8aỶ\xb3Rm'\xd0b\x05\xbb7\x8bQ\x89\xbd\xd2-\xb3K\xe0\xd2\xfe\xf9O\xa7-ћ\xaa\xf0\xa4\x8fJ\x8e\xcd\xf2\x89F!\x19\x0e\x92\xd0\tը\xb3\xb6Q\x96\x89\xffE\x10K\f>%\xf4A\x92\xc07\x1d\xbf(\n]7P{\xb0\r\xc2'V\xbe\xb8\x0e\xb6ViV#\xfc\xaa\xcapx\xaf\r\xea\xfe\xf0va\x89i\x94\x13\x15\xec\xa2\xc6\x00\xc6*\x9d=\xc5\x0e\xcb\"P\xf5|#\xdb\xc9Q\x8e\xf7\xfcΗ\xac\xd4Ȳ\x97,\xa2L\xe1Wp%\xf37\xedc\x8dWݲԚRU8\x98\x0eS\x89\xb8\x81N\xab\x12\x8d9s\xef\x89|$\xc3\xd3q`f\x96\xb0\xe2\xf03\x13]\xc3>\x04\x94)\x1blٲ\xa7P\x1dʏ\x9b\xf5\xf3\x1f\xb7\xa3a8\x89\x19\xac\xb4\x86\xc0\x82Dﴲ\xaaT\x02vh_\x11\xa5\xc7-h\xd5\x015\x81\\ͥ\x01&\xab\x81'\xa4\v\x8ePM\x97\xdc\xf3\xa3\xd90\xd9_'աN\x8f\x1dh\xcb\x0e\xb5\xe5\x11}×\x84\x95dt\xa2\xc4{\xd23\xac\x82\x8a\xe2\t\x06-z,Ū7M8'n@c\xa7Ѡ\xb4c\x11z\xc3\xed\x81IP\xbb߱\xb4\x05lQ\x13\x9bx\xffK%\x0f\xa8-h,U-\xf9\xbf\x06\xde\x06\xac\xf2\x9b\nf\xb1\x0f\a\xc7\xcfc\xb7d\x02\x0eL8\xbc'\xdbA\xcb\xde@#\xed\x02N&\xfc\xfc\x12S\xc0g\xa5\x11\xb8ܫ%4\xd6vf\xf9\xf0Ps\x1b\xc3i\xa9\xda\xd6In\xdf\x1e\xbc\xb9\xf9\xceY\xa5\xcdC\x85\a\x14\x0f\x86\xd7\v\xa6ˆ[,\xad\xd3\xf8\xc0:\xbe\xf0\xa2K\x1fR\x8b\xb6\xfa\x83\xee\x03\xb0y?\x92uv\xd1\xc2\xe7c\xe1\x99\x13\xa0\x90H\xb7\x9c\xf5\xa4A\x8b\xa3\xa1i\x88\xac\xf3\xe5\xafۯ\x10\xb7\xf6\x871\xb5\xbe\xb7\xfb\x91\xd0\x1c\x8f\x80\f\xc6\xe5\x1eu8ĽV\xad牲\xea\x14\x97\xd6\xff(\x05G95\xbfq\xbb\x96[:\xf7\x7f84\x96Ϊ\x80\x95\xcf1`\x87\xe0:\xf2\uea80\xb5\x84\x15kQ\xac\x98\xc1\x1f~\x00di\xb3 \xc3^w\x04iz4]\x1c\xac\x96L\xc4\f\xe7\xc4y\x1d\xdd~\xdbaI\aG\xb6#\"\xbe\xe7}\f \xdfe\xc9\xcab\xc4.\xef\xae\xf4e\xa1\x7f\xbah\"ϧ\x1cM\x14K&\x10\x1b\xa3QX9c\n \xa6!l\xa0\xd1\xd8)í\xd2o\xc48D\xafb\xc6\xe1\x84\xf1\xe9+\x99,Q\\\xd0d\xe5\x17\x01\x97\x15\xd9\x11\x87;G\xf0\x10\x18x\x99\x94\xac\x15\xf9\xc4)\xf3\x86om\x89\x86\xae\xa8AK\x1a\xc9L`\xe1\x12\x8e\xb9\x1d\xa49\xdcT\xab\x9dR\x02\xd9\x14\xefJ÷\x92u\xa6Q\xf6\x82n\xeb=ĕ_\xdf:\xa4\xcdW\xdb\xf5=\xfd\x89\xe3t/\x0e\xbc\xea\x01\x98\x9c\x87\xb2\x9c9\xc8B\x00ZZ\xb4ڮ\xc1\xf4\xe4s#H'\x04\xdb\t\\\x82\xd5n\xae\xd8\xe9kH_d\xbb\x12\xccd\x17L\x14ܦ\xebs\xd7/2\x84ү\xb0\r\x9bB\xcd`q\x8a?\x94\xac'D|HK\xe0\x95\xdb&Ky\xe6\xfeA\x9ft\xb1\x1a\xafV(Y\x9eէO\u0082:j\x7fF\x99\xcd\xf3\xca\xeb{I3\x82\xe5o\xd1,\xb0<}\x13g\xba=\x8f\br\xdaM\xa4<\xa5\x9c\"\a#\x90\xc0\n\\w\xbb\xec\xe4\xe1\\c5\x97y1:\xaf\xcc\xf4X\xe9\x13n;\x03w\xe8\xf3\xadϔQ\xad\x94\xdc\xf3z\xbewZ:\x9e\U000d1cea͂F\xb2%Y\x9cb\x04I\xb2\xf0\xc9\xdd\"\x06\x10J\x93\xf6\xbcv\xfa\x94\xeb\xef9\x8a\xca\xdc\xec\xed\x17\xecᅸ\x80a\x83\x121\xda\xf5P\x95\xe4\xaf\xe1B8\xe3+G\x9a\xcc(\x10bJA\x90x\xe4\xc8\r\xbc{\aJû\xd0Qxw\x1f\xe2\x8f\xe3\xc2.x\x9aDg8\xber!\xe2\xbe7\x85\xa5!\x95\xa6BF\xb9K \xfe\xf7\xc9\xf2\x89\x1d,\xd5W^w\xab\xe0\x95q;\xe4\xae9\x04\x8f\xbc\xcc=\xecpO\xf9\x92F봤ІZS\x06a<K\xe52\xd8~F)\x93ę\v\nMC\x92ׂ\xfe\x9fbv\xea\xe8\x19e\\w\x9b\x84>\x85\x1dz7\x97\x84\x1c\xaf\x8er*\xcdkNu\x81\x1cf\x8eyK\x00\x87\x8c\xa4}U\xee\xe1\xca\xe3mA\x99BL\xe2\b\x00\x8f\xec\xc8C\xc3\xe6\x04\xe0Tw\xac\xb6\xeb\fρ\xa2\xea\xfd+\xe3\x9d\x17\xad\xb1y^]e\a\x12%\x83\xd74\xfc\xda\xf0\xb2\x19\x9f۬F\xf0\xb2\xb0\x17\xf49\xeaMb\x86\xd6\xdbi\xd0\x1c\xc9\xfa\xdbh\xf1\xc4OF\b7Թ\x91\xff\\\xa8\xf3\x89J\xc74\x13\x02\xc5/\\\xa0\t\xfb^\x11\x017s\xaa\xc1\xac\xaeݡ&\xc3\xeeir\xd8\xe0D\f웕\xe4\"\x1djJ\xdaB\xb6\xe9L,\xdcNk\x06٦\xd3|6\x03\xdb\xf9ȹȗ\x10\x935S؛L\xa7\x002\x9d\x1a\xfbbvv\U000fceaa\xcc\xf2\x9d\x9b\xeb\n\xadВ\xedϧtZ\xa3\xb4\xb1Q\xab\xf6\xdfTj\x95\xa1ř6\xb1.\x95's\n\xdf\xcb\xd0U\x02\xff,\x96L\xbe\x91\x16ۨ\xb9\x93?\xb2\v\x94\xbe\xb7Bܰ\x02<\xa0\x04\xf2\r\xc6\x05\x85R\xcf\xd2\x14S\x9a<\xbe\r\\\xfa\xb0\x12\xae_\xbc\x8c\xb1\x89\xd4\xf7h\xbe\x12Z\xf8&\xc1{s\x86\xa7\x8fj\xe4\xa9\x19#\xcc!&\xf6g+fq\x91ezU\xb2\x92\x85\xa1!y\xfb\x82ƉL\xc4\xfe\x81\xc9[\xd82Կ&\x9b\xbc\x9d\xafژ\x01\x06:0\xe9q\xfb\\\x11\xfb\xed\x19]\x8bư\xfaR`\xfd\x1cV\x85\xdeSO\x02lG\x89\xcdX\xb4\xf7\xa6w\xb6\x9b\x02\x86T\xd5%\t\x9eT巗7w\x82o\x92\xa4c\xb6\xb9 Ɇ\xd9&\x02\xcc\xde\t\xe1ifiP_\x15퐼\xe9{eC\xbe\xedpI<Z\x93\x03@\xbc\xe6\"\xa1tm\xae\xc2z\xc2\xd7\xcc\xe8ǲ\xc4\xcef4[\xc0Fc\xc7tvj\xf6F\x96N\x86\xceN\xce\xd1\xe2\\\x96\xe7\xf0\f\x95\x99\xfb\xc5C\xe3Mv\xee\xe5\xbbd\xea\xd8\x0ej\x94\x88\xc8\xee_\x8a\x8ey\x81\x7f\x8b\x1aW\xca9o\x97\xd5\xe8\xb8\x12\xfa!\x03\xf5\x9c\b\x85\xb9\x89\xed\xaaX@T\xdct\x82\xbd\xe52\x9e(a\x826\x89\xdfN\x9f\tnM\xa8\x86w\xbb|\x1a\x95{|˝\xc1\xa9\x8c\x06\xe0\xf8\x1e\xf7cv8\x03\x8cѓ\u05cfWVF\xeb\xc7\xe8u\xbcBi)\x89=>\xcd\x1c\xd3ly\xb6\xd6M\xfa\xa7\xb7U\x06\xa3\xd7\xdcK\x12\x8f\x16_\xc8L\xfaw\xe4\\^\xb2%\x17'`\xf1\x8f\x04\xab\xe9K\xdf\xfd\xf0p\xc8l\xffRQ6L\xd6\xe4\x10\x92\x82\x9b\x0f\x8e9ƳTc\x94X\x8c\xc5\xff\x7f\xe6\x14\xd9\xeb2\x1b\xf4\x92W\t\xef\xbe=\x95\x8e\xb8\xdd\xf02\xb4\x84\x7f\xff\xe7\xee\xbf\x01\x00\x00\xff\xff\xf8I\x957\xdd!\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := io.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1.CustomResourceDefinition))
	}
	return objs
}
