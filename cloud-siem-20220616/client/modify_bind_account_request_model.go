// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBindAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *ModifyBindAccountRequest
	GetAccessId() *string
	SetAccountId(v string) *ModifyBindAccountRequest
	GetAccountId() *string
	SetAccountName(v string) *ModifyBindAccountRequest
	GetAccountName() *string
	SetBindId(v int64) *ModifyBindAccountRequest
	GetBindId() *int64
	SetCloudCode(v string) *ModifyBindAccountRequest
	GetCloudCode() *string
	SetRegionId(v string) *ModifyBindAccountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ModifyBindAccountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ModifyBindAccountRequest
	GetRoleType() *int32
}

type ModifyBindAccountRequest struct {
	// The AccessKey ID of the Alibaba Cloud account.
	//
	// example:
	//
	// ABCXXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the multicloud account.
	//
	// example:
	//
	// sas_account_xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the binding record. This is the BindId value returned by the ListBindAccount operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	BindId *int64 `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The code of the multicloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region where the Data Management center for threat analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can specify this parameter to switch to the member\\"s view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
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

func (s ModifyBindAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBindAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountRequest) GetAccessId() *string {
	return s.AccessId
}

func (s *ModifyBindAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifyBindAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyBindAccountRequest) GetBindId() *int64 {
	return s.BindId
}

func (s *ModifyBindAccountRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ModifyBindAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBindAccountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ModifyBindAccountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ModifyBindAccountRequest) SetAccessId(v string) *ModifyBindAccountRequest {
	s.AccessId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetAccountId(v string) *ModifyBindAccountRequest {
	s.AccountId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetAccountName(v string) *ModifyBindAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyBindAccountRequest) SetBindId(v int64) *ModifyBindAccountRequest {
	s.BindId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetCloudCode(v string) *ModifyBindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *ModifyBindAccountRequest) SetRegionId(v string) *ModifyBindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBindAccountRequest) SetRoleFor(v int64) *ModifyBindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *ModifyBindAccountRequest) SetRoleType(v int32) *ModifyBindAccountRequest {
	s.RoleType = &v
	return s
}

func (s *ModifyBindAccountRequest) Validate() error {
	return dara.Validate(s)
}
