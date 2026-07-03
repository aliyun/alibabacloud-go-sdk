// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DeleteBindAccountRequest
	GetAccessId() *string
	SetAccountId(v string) *DeleteBindAccountRequest
	GetAccountId() *string
	SetBindId(v int64) *DeleteBindAccountRequest
	GetBindId() *int64
	SetCloudCode(v string) *DeleteBindAccountRequest
	GetCloudCode() *string
	SetRegionId(v string) *DeleteBindAccountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteBindAccountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DeleteBindAccountRequest
	GetRoleType() *int32
}

type DeleteBindAccountRequest struct {
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
	// The binding ID. Call the [ListBindAccount](https://api.aliyun-inc.com/#/publishment/document/cloud-siem/863fdf54478f4cc5877e27c2a5fe9e44?tenantUuid=f382fccd88b94c5c8c864def6815b854\\&activeTabKey=api%7CListBindAccount) operation to obtain the binding ID.
	//
	// example:
	//
	// 10
	BindId *int64 `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The code for the cloud service provider. Valid values:
	//
	// - qcloud: Tencent Cloud
	//
	// - aliyun: Alibaba Cloud
	//
	// - hcloud: Huawei Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region where the threat analysis data center is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in mainland China or the China (Hong Kong) region.
	//
	// - ap-southeast-1: Your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can specify this parameter to perform the operation from the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that are managed by your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DeleteBindAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountRequest) GetAccessId() *string {
	return s.AccessId
}

func (s *DeleteBindAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteBindAccountRequest) GetBindId() *int64 {
	return s.BindId
}

func (s *DeleteBindAccountRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DeleteBindAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBindAccountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteBindAccountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DeleteBindAccountRequest) SetAccessId(v string) *DeleteBindAccountRequest {
	s.AccessId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetAccountId(v string) *DeleteBindAccountRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetBindId(v int64) *DeleteBindAccountRequest {
	s.BindId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetCloudCode(v string) *DeleteBindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *DeleteBindAccountRequest) SetRegionId(v string) *DeleteBindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBindAccountRequest) SetRoleFor(v int64) *DeleteBindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteBindAccountRequest) SetRoleType(v int32) *DeleteBindAccountRequest {
	s.RoleType = &v
	return s
}

func (s *DeleteBindAccountRequest) Validate() error {
	return dara.Validate(s)
}
