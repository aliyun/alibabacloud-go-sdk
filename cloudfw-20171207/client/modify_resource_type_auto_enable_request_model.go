// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceTypeAutoEnableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyResourceTypeAutoEnableRequest
	GetLang() *string
	SetRegionNo(v string) *ModifyResourceTypeAutoEnableRequest
	GetRegionNo() *string
	SetResourceTypeAutoEnable(v string) *ModifyResourceTypeAutoEnableRequest
	GetResourceTypeAutoEnable() *string
}

type ModifyResourceTypeAutoEnableRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// {"SlbEIP":true,"GaEIP":false,"EcsIPv6":true,"NatPublicIP":true,"SlbIPv6":false,"BastionHostIngressIP":false,"EIP":true,"NatEIP":true,"SlbPublicIP":true,"EcsEIP":true,"EniEIP":true,"HAVIP":true,"NlbEIP":true,"NlbIPv6":false,"EniEIPv6":false,"EcsPublicIP":true,"AlbIPv6":true,"BastionHostIP":false,"BastionHostEgressIP":true,"GaEIPV6":false,"AlbEIP":false}
	ResourceTypeAutoEnable *string `json:"ResourceTypeAutoEnable,omitempty" xml:"ResourceTypeAutoEnable,omitempty"`
}

func (s ModifyResourceTypeAutoEnableRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceTypeAutoEnableRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceTypeAutoEnableRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyResourceTypeAutoEnableRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *ModifyResourceTypeAutoEnableRequest) GetResourceTypeAutoEnable() *string {
	return s.ResourceTypeAutoEnable
}

func (s *ModifyResourceTypeAutoEnableRequest) SetLang(v string) *ModifyResourceTypeAutoEnableRequest {
	s.Lang = &v
	return s
}

func (s *ModifyResourceTypeAutoEnableRequest) SetRegionNo(v string) *ModifyResourceTypeAutoEnableRequest {
	s.RegionNo = &v
	return s
}

func (s *ModifyResourceTypeAutoEnableRequest) SetResourceTypeAutoEnable(v string) *ModifyResourceTypeAutoEnableRequest {
	s.ResourceTypeAutoEnable = &v
	return s
}

func (s *ModifyResourceTypeAutoEnableRequest) Validate() error {
	return dara.Validate(s)
}
