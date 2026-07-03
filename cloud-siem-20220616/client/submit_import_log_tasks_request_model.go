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
	// The list of accounts for log ingestion. The value must be a JSON array. Valid values:
	//
	// - AccountId: The ID of the account.
	//
	// - Imported: Specifies whether to enable or disable log ingestion for the account. Valid values:
	//
	//   - 0: Disable ingestion.
	//
	//   - 1: Enable ingestion.
	//
	// example:
	//
	// [{"AccountId":"123123","Imported":1}]
	Accounts *string `json:"Accounts,omitempty" xml:"Accounts,omitempty"`
	// Specifies whether to automatically enable log ingestion for accounts that are configured with the specified log. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// example:
	//
	// 1
	AutoImported *int32 `json:"AutoImported,omitempty" xml:"AutoImported,omitempty"`
	// The code of the cloud service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The list of logs to be ingested. The value must be a JSON array.
	//
	// example:
	//
	// ["cloud_siem_qcloud_cfw_alert_log"]
	LogCodes *string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty"`
	// The code of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region where the data management center for Threat Analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member that the administrator wants to access.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts within the enterprise.
	//
	// example:
	//
	// 1
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
