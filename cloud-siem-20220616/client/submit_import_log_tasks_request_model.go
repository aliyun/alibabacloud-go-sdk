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
	// The list of accounts to be onboarded, in JSON array format. Valid values:
	//
	// - AccountId: the ID of the account to be onboarded.
	//
	// - Imported: specifies whether to onboard or disconnect the account. Valid values:
	//
	//   - 0: Disconnect the account.
	//
	//   - 1: Onboard the account.
	//
	// example:
	//
	// [{"AccountId":"123123","Imported":1}]
	Accounts *string `json:"Accounts,omitempty" xml:"Accounts,omitempty"`
	// Specifies whether to automatically onboard accounts that are configured with the log. Valid values:
	//
	// - 1: Automatically onboard.
	//
	// - 0: Do not automatically onboard.
	//
	// example:
	//
	// 1
	AutoImported *int32 `json:"AutoImported,omitempty" xml:"AutoImported,omitempty"`
	// The multi-cloud code.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The list of logs to be onboarded, in JSON array format.
	//
	// example:
	//
	// ["cloud_siem_qcloud_cfw_alert_log"]
	LogCodes *string `json:"LogCodes,omitempty" xml:"LogCodes,omitempty"`
	// The product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region of the data management center for Threat Analysis. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: Current Alibaba Cloud account view.
	//
	// - 1: View of all accounts in the enterprise.
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
