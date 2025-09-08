// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImportLogTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v string) *SubmitImportLogTasksRequest
	GetAccounts() *string
	SetAutoImported(v int32) *SubmitImportLogTasksRequest
	GetAutoImported() *int32
	SetCloudCode(v string) *SubmitImportLogTasksRequest
	GetCloudCode() *string
	SetLogCodes(v string) *SubmitImportLogTasksRequest
	GetLogCodes() *string
	SetProdCode(v string) *SubmitImportLogTasksRequest
	GetProdCode() *string
	SetRegionId(v string) *SubmitImportLogTasksRequest
	GetRegionId() *string
	SetRoleFor(v int64) *SubmitImportLogTasksRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *SubmitImportLogTasksRequest
	GetRoleType() *int32
}

type SubmitImportLogTasksRequest struct {
	// The accounts that you want to add. The value is a JSON array. Valid values:
	//
	// 	- AccountId: the IDs of the accounts.
	//
	// 	- Imported: specifies whether to add the accounts. Valid values:
	//
	//     	- 0: no
	//
	//     	- 1: yes
	//
	// example:
	//
	// [{"AccountId":"123123","Imported":1}]
	Accounts *string `json:"Accounts,omitempty" xml:"Accounts,omitempty"`
	// Specifies whether to automatically add the account for which the logging feature is configured. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// ["cloud_siem_qcloud_cfw_alert_log"]
	AutoImported *int32 `json:"AutoImported,omitempty" xml:"AutoImported,omitempty"`
	// The code that is used for multi-cloud environments. Valid values:
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The logs that you want to collect. The value is a JSON array.
	//
	// example:
	//
	// ["cloud_siem_qcloud_cfw_alert_log"]
	LogCodes *string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty"`
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
	// 0
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s SubmitImportLogTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportLogTasksRequest) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksRequest) GetAccounts() *string {
	return s.Accounts
}

func (s *SubmitImportLogTasksRequest) GetAutoImported() *int32 {
	return s.AutoImported
}

func (s *SubmitImportLogTasksRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *SubmitImportLogTasksRequest) GetLogCodes() *string {
	return s.LogCodes
}

func (s *SubmitImportLogTasksRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *SubmitImportLogTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SubmitImportLogTasksRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *SubmitImportLogTasksRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *SubmitImportLogTasksRequest) SetAccounts(v string) *SubmitImportLogTasksRequest {
	s.Accounts = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetAutoImported(v int32) *SubmitImportLogTasksRequest {
	s.AutoImported = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetCloudCode(v string) *SubmitImportLogTasksRequest {
	s.CloudCode = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetLogCodes(v string) *SubmitImportLogTasksRequest {
	s.LogCodes = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetProdCode(v string) *SubmitImportLogTasksRequest {
	s.ProdCode = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetRegionId(v string) *SubmitImportLogTasksRequest {
	s.RegionId = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetRoleFor(v int64) *SubmitImportLogTasksRequest {
	s.RoleFor = &v
	return s
}

func (s *SubmitImportLogTasksRequest) SetRoleType(v int32) *SubmitImportLogTasksRequest {
	s.RoleType = &v
	return s
}

func (s *SubmitImportLogTasksRequest) Validate() error {
	return dara.Validate(s)
}
