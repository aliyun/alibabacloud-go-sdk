// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *SetStorageRequest
	GetRegion() *string
	SetRegionId(v string) *SetStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *SetStorageRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *SetStorageRequest
	GetRoleType() *int32
	SetTtl(v int32) *SetStorageRequest
	GetTtl() *int32
}

type SetStorageRequest struct {
	// The log storage region.
	//
	// If the Data Management center is in cn-hangzhou, the default value of Region is **cn-shanghai**. If the Data Management center is in **ap-southeast-1**, the default value of **Region*	- is **ap-southeast-1**.
	//
	// The log storage region cannot be changed. To change the region, contact the Threat Analysis operations team.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region of the Data Management center for Threat Analysis. Select the region for the Data Management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Select this value if your assets are in a region outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the view of a specific member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The log storage duration in days. The default value is 180. The minimum value is 30 and the maximum value is 3000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SetStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s SetStorageRequest) GoString() string {
	return s.String()
}

func (s *SetStorageRequest) GetRegion() *string {
	return s.Region
}

func (s *SetStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *SetStorageRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *SetStorageRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SetStorageRequest) SetRegion(v string) *SetStorageRequest {
	s.Region = &v
	return s
}

func (s *SetStorageRequest) SetRegionId(v string) *SetStorageRequest {
	s.RegionId = &v
	return s
}

func (s *SetStorageRequest) SetRoleFor(v int64) *SetStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *SetStorageRequest) SetRoleType(v int32) *SetStorageRequest {
	s.RoleType = &v
	return s
}

func (s *SetStorageRequest) SetTtl(v int32) *SetStorageRequest {
	s.Ttl = &v
	return s
}

func (s *SetStorageRequest) Validate() error {
	return dara.Validate(s)
}
