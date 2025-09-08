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
	// The log code of the cloud service, such as the code of the process log for Security Center. You can obtain the log code from the response of the ListDelivery operation.
	//
	// example:
	//
	// cloud_siem_aegis_proc
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
	// sas
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
