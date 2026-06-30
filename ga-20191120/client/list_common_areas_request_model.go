// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonAreasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *ListCommonAreasRequest
	GetIpVersion() *string
	SetIsEpg(v bool) *ListCommonAreasRequest
	GetIsEpg() *bool
	SetIsIpSet(v bool) *ListCommonAreasRequest
	GetIsIpSet() *bool
}

type ListCommonAreasRequest struct {
	// The IP address protocol used to connect to Global Accelerator (GA). Valid values:
	//
	// - **IPv4*	- (default): IPv4 address protocol. Queries regions that support IPv4.
	//
	// - **IPv6**: IPv6 address protocol. Queries regions that support IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Specifies whether the region is an endpoint group region supported by Global Accelerator.
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// true
	IsEpg *bool `json:"IsEpg,omitempty" xml:"IsEpg,omitempty"`
	// Specifies whether the region is an acceleration area supported by Global Accelerator.
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// true
	IsIpSet *bool `json:"IsIpSet,omitempty" xml:"IsIpSet,omitempty"`
}

func (s ListCommonAreasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommonAreasRequest) GoString() string {
	return s.String()
}

func (s *ListCommonAreasRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListCommonAreasRequest) GetIsEpg() *bool {
	return s.IsEpg
}

func (s *ListCommonAreasRequest) GetIsIpSet() *bool {
	return s.IsIpSet
}

func (s *ListCommonAreasRequest) SetIpVersion(v string) *ListCommonAreasRequest {
	s.IpVersion = &v
	return s
}

func (s *ListCommonAreasRequest) SetIsEpg(v bool) *ListCommonAreasRequest {
	s.IsEpg = &v
	return s
}

func (s *ListCommonAreasRequest) SetIsIpSet(v bool) *ListCommonAreasRequest {
	s.IsIpSet = &v
	return s
}

func (s *ListCommonAreasRequest) Validate() error {
	return dara.Validate(s)
}
