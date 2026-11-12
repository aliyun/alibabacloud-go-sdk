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
	// If the data management center is located in **cn-hangzhou**, the default value of **Region*	- is China (Shanghai) (cn-shanghai). If the data management center is located in **ap-southeast-1**, the default value of **Region*	- is Singapore (ap-southeast-1).
	//
	// The log storage region cannot be modified. To modify the region, contact the Threat Analysis operations personnel.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region of the data management center for Threat Analysis. Select the region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Select this value if your assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose view the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account
	//
	// - 1: the view of all accounts in the enterprise
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The log retention period in days. Default value: 180. Valid values: 30 to 3000.
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
