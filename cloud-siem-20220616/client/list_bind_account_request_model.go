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
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleFor  *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	RoleType *int32  `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
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
