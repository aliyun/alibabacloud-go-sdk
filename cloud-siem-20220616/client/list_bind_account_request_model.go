// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudCode(v string) *ListBindAccountRequest
	GetCloudCode() *string
	SetRegionId(v string) *ListBindAccountRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListBindAccountRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListBindAccountRequest
	GetRoleType() *int32
}

type ListBindAccountRequest struct {
	// The code of the cloud service provider. Valid values:
	//
	// - qcloud: Tencent Cloud.
	//
	// - aliyun: Alibaba Cloud.
	//
	// - hcloud: Huawei Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The region where the data management center of Threat Analysis is deployed. Select the region of the data management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose perspective the administrator wants to switch to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that are managed by the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListBindAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindAccountRequest) GoString() string {
	return s.String()
}

func (s *ListBindAccountRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListBindAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBindAccountRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListBindAccountRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListBindAccountRequest) SetCloudCode(v string) *ListBindAccountRequest {
	s.CloudCode = &v
	return s
}

func (s *ListBindAccountRequest) SetRegionId(v string) *ListBindAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ListBindAccountRequest) SetRoleFor(v int64) *ListBindAccountRequest {
	s.RoleFor = &v
	return s
}

func (s *ListBindAccountRequest) SetRoleType(v int32) *ListBindAccountRequest {
	s.RoleType = &v
	return s
}

func (s *ListBindAccountRequest) Validate() error {
	return dara.Validate(s)
}
