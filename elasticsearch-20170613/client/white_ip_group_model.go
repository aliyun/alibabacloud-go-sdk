// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhiteIpGroup interface {
	dara.Model
	String() string
	GoString() string
	SetWhiteIpType(v string) *WhiteIpGroup
	GetWhiteIpType() *string
	SetGroupName(v string) *WhiteIpGroup
	GetGroupName() *string
	SetIps(v []*string) *WhiteIpGroup
	GetIps() []*string
}

type WhiteIpGroup struct {
	WhiteIpType *string   `json:"WhiteIpType,omitempty" xml:"WhiteIpType,omitempty"`
	GroupName   *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips         []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s WhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s WhiteIpGroup) GoString() string {
	return s.String()
}

func (s *WhiteIpGroup) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *WhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *WhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *WhiteIpGroup) SetWhiteIpType(v string) *WhiteIpGroup {
	s.WhiteIpType = &v
	return s
}

func (s *WhiteIpGroup) SetGroupName(v string) *WhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *WhiteIpGroup) SetIps(v []*string) *WhiteIpGroup {
	s.Ips = v
	return s
}

func (s *WhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
