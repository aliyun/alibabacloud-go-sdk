// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsByLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudCode(v string) *ListAccountsByLogRequest
	GetCloudCode() *string
	SetLogCodes(v []*string) *ListAccountsByLogRequest
	GetLogCodes() []*string
	SetProdCode(v string) *ListAccountsByLogRequest
	GetProdCode() *string
	SetRegionId(v string) *ListAccountsByLogRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListAccountsByLogRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListAccountsByLogRequest
	GetRoleType() *int32
}

type ListAccountsByLogRequest struct {
	// The code of the cloud provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The list of log codes. This parameter is a repeatList array parameter that must be passed in dot-separated expanded form, for example, LogCodes.1=cloud_siem_hcloud_waf_alert_log. Passing a JSON array string directly is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// LogCodes.1=cloud_siem_hcloud_waf_alert_log
	LogCodes []*string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty" type:"Repeated"`
	// The code of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region where the CTDR data management center is located. Select the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: the assets belong to the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: the assets belong to regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
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

func (s ListAccountsByLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsByLogRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListAccountsByLogRequest) GetLogCodes() []*string {
	return s.LogCodes
}

func (s *ListAccountsByLogRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListAccountsByLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccountsByLogRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListAccountsByLogRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListAccountsByLogRequest) SetCloudCode(v string) *ListAccountsByLogRequest {
	s.CloudCode = &v
	return s
}

func (s *ListAccountsByLogRequest) SetLogCodes(v []*string) *ListAccountsByLogRequest {
	s.LogCodes = v
	return s
}

func (s *ListAccountsByLogRequest) SetProdCode(v string) *ListAccountsByLogRequest {
	s.ProdCode = &v
	return s
}

func (s *ListAccountsByLogRequest) SetRegionId(v string) *ListAccountsByLogRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountsByLogRequest) SetRoleFor(v int64) *ListAccountsByLogRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAccountsByLogRequest) SetRoleType(v int32) *ListAccountsByLogRequest {
	s.RoleType = &v
	return s
}

func (s *ListAccountsByLogRequest) Validate() error {
	return dara.Validate(s)
}
