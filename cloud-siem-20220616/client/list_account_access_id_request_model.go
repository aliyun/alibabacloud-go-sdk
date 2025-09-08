// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountAccessIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudCode(v string) *ListAccountAccessIdRequest
	GetCloudCode() *string
	SetRegionId(v string) *ListAccountAccessIdRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListAccountAccessIdRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListAccountAccessIdRequest
	GetRoleType() *int32
}

type ListAccountAccessIdRequest struct {
	// The code of the cloud service provider.
	//
	// Valid values:
	//
	// 	- qcloud
	//
	// 	- hcloud
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

func (s ListAccountAccessIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountAccessIdRequest) GoString() string {
	return s.String()
}

func (s *ListAccountAccessIdRequest) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListAccountAccessIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccountAccessIdRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListAccountAccessIdRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListAccountAccessIdRequest) SetCloudCode(v string) *ListAccountAccessIdRequest {
	s.CloudCode = &v
	return s
}

func (s *ListAccountAccessIdRequest) SetRegionId(v string) *ListAccountAccessIdRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountAccessIdRequest) SetRoleFor(v int64) *ListAccountAccessIdRequest {
	s.RoleFor = &v
	return s
}

func (s *ListAccountAccessIdRequest) SetRoleType(v int32) *ListAccountAccessIdRequest {
	s.RoleType = &v
	return s
}

func (s *ListAccountAccessIdRequest) Validate() error {
	return dara.Validate(s)
}
