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
	// The log code of the cloud service, such as the code of the process log for Security Center. This parameter is optional. If you leave this parameter empty, operations are performed on all logs of the cloud service.
	//
	// example:
	//
	// cloud_siem_cfw_flow
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The code of the cloud service. Valid values:
	//
	// 	- qcloud_waf
	//
	// 	- qlcoud_cfw
	//
	// 	- hcloud_waf
	//
	// 	- hcloud_cfw
	//
	// 	- ddos
	//
	// 	- sas
	//
	// 	- cfw
	//
	// 	- config
	//
	// 	- csk
	//
	// 	- fc
	//
	// 	- rds
	//
	// 	- nas
	//
	// 	- apigateway
	//
	// 	- cdn
	//
	// 	- mongodb
	//
	// 	- eip
	//
	// 	- slb
	//
	// 	- vpc
	//
	// 	- actiontrail
	//
	// 	- waf
	//
	// 	- bastionhost
	//
	// 	- oss
	//
	// 	- polardb
	//
	// This parameter is required.
	//
	// example:
	//
	// cfw
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
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
