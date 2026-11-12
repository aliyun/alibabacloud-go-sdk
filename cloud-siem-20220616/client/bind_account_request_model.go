// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *BindAccountRequest
	GetAccessId() *string
	SetAccountId(v string) *BindAccountRequest
	GetAccountId() *string
	SetAccountName(v string) *BindAccountRequest
	GetAccountName() *string
	SetCloudCode(v string) *BindAccountRequest
	GetCloudCode() *string
	SetRegionId(v string) *BindAccountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *BindAccountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *BindAccountRequest
	GetRoleType() *int32
}

type BindAccountRequest struct {
	// The AccessKey ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the multi-cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample multi-cloud account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The code of the multi-cloud provider. Valid values:
	//
	// - hcloud: Huawei Cloud
	//
	// - qcloud: Tencent Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region of the data management center for Threat Analysis. Select the region of the management center based on the region of your assets. Valid values:
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
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s BindAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAccountRequest) GoString() string {
	return s.String()
}

func (s *BindAccountRequest) GetAccessId() *string {
	return s.AccessId
}

func (s *BindAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *BindAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *BindAccountRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *BindAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindAccountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *BindAccountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *BindAccountRequest) SetAccessId(v string) *BindAccountRequest {
	s.AccessId = &v
	return s
}

func (s *BindAccountRequest) SetAccountId(v string) *BindAccountRequest {
	s.AccountId = &v
	return s
}

func (s *BindAccountRequest) SetAccountName(v string) *BindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *BindAccountRequest) SetCloudCode(v string) *BindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *BindAccountRequest) SetRegionId(v string) *BindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *BindAccountRequest) SetRoleFor(v int64) *BindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *BindAccountRequest) SetRoleType(v int32) *BindAccountRequest {
	s.RoleType = &v
	return s
}

func (s *BindAccountRequest) Validate() error {
	return dara.Validate(s)
}
