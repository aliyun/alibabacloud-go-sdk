// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogCode(v string) *OpenDeliveryRequest
	GetLogCode() *string
	SetProductCode(v string) *OpenDeliveryRequest
	GetProductCode() *string
	SetRegionId(v string) *OpenDeliveryRequest
	GetRegionId() *string
	SetRoleFor(v int64) *OpenDeliveryRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *OpenDeliveryRequest
	GetRoleType() *int32
}

type OpenDeliveryRequest struct {
	// The log code of the cloud service, such as the process log in Security Center. This parameter is optional. If you do not specify this parameter, the operation applies to all logs of the cloud service.
	//
	// example:
	//
	// cloud_siem_cfw_flow
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
	// cfw
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region where the data management center of Threat Analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are located in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Select this value if your assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s OpenDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDeliveryRequest) GoString() string {
	return s.String()
}

func (s *OpenDeliveryRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *OpenDeliveryRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *OpenDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenDeliveryRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *OpenDeliveryRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *OpenDeliveryRequest) SetLogCode(v string) *OpenDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *OpenDeliveryRequest) SetProductCode(v string) *OpenDeliveryRequest {
	s.ProductCode = &v
	return s
}

func (s *OpenDeliveryRequest) SetRegionId(v string) *OpenDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *OpenDeliveryRequest) SetRoleFor(v int64) *OpenDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *OpenDeliveryRequest) SetRoleType(v int32) *OpenDeliveryRequest {
	s.RoleType = &v
	return s
}

func (s *OpenDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
