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
	// The code that is used for multi-cloud environments.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The codes of logs. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cloud_siem_hcloud_waf_alert_log"]
	LogCodes []*string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty" type:"Repeated"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions inside China.
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
