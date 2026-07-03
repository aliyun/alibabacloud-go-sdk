// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogCode(v string) *CloseDeliveryRequest
	GetLogCode() *string
	SetProductCode(v string) *CloseDeliveryRequest
	GetProductCode() *string
	SetRegionId(v string) *CloseDeliveryRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CloseDeliveryRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *CloseDeliveryRequest
	GetRoleType() *int32
}

type CloseDeliveryRequest struct {
	// The code of the log within the cloud service. For example, the process log of Security Center. For valid values, see the return value of the ListDelivery operation.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The code of the cloud service. Valid values:
	//
	// - qcloud_waf
	//
	// - qlcoud_cfw
	//
	// - hcloud_waf
	//
	// - hcloud_cfw
	//
	// - ddos
	//
	// - sas
	//
	// - cfw
	//
	// - config
	//
	// - csk
	//
	// - fc
	//
	// - rds
	//
	// - nas
	//
	// - apigateway
	//
	// - cdn
	//
	// - mongodb
	//
	// - eip
	//
	// - slb
	//
	// - vpc
	//
	// - actiontrail
	//
	// - waf
	//
	// - bastionhost
	//
	// - oss
	//
	// - polardb
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region where the data management center of threat analysis is located. Select a region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this region if your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Select this region if your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can use this parameter to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CloseDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CloseDeliveryRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *CloseDeliveryRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CloseDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseDeliveryRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CloseDeliveryRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *CloseDeliveryRequest) SetLogCode(v string) *CloseDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *CloseDeliveryRequest) SetProductCode(v string) *CloseDeliveryRequest {
	s.ProductCode = &v
	return s
}

func (s *CloseDeliveryRequest) SetRegionId(v string) *CloseDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *CloseDeliveryRequest) SetRoleFor(v int64) *CloseDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *CloseDeliveryRequest) SetRoleType(v int32) *CloseDeliveryRequest {
	s.RoleType = &v
	return s
}

func (s *CloseDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
