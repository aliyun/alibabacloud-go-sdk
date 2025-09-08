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
	// The storage region of logs.
	//
	// If the data management center is **cn-hangzhou**, the default value of **Region*	- is cn-shanghai, which specifies the China (Shanghai) region. If the data management center is **ap-southeast-1**, the default value of **Region*	- is ap-southeast-1, which specifies the Singapore region.
	//
	// The region for log storage cannot be changed. To change the region, contact the technical support of threat analysis.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
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
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The storage duration of logs. Default value: 180. Minimum value: 30. Maximum value: 3000. Unit: days.
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
