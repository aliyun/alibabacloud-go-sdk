// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeConfigGroupRequest
	GetGroupId() *string
	SetGroupIds(v []*string) *DescribeConfigGroupRequest
	GetGroupIds() []*string
	SetName(v string) *DescribeConfigGroupRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeConfigGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeConfigGroupRequest
	GetPageSize() *int32
	SetProductType(v string) *DescribeConfigGroupRequest
	GetProductType() *string
	SetRegionId(v string) *DescribeConfigGroupRequest
	GetRegionId() *string
	SetStatuses(v []*string) *DescribeConfigGroupRequest
	GetStatuses() []*string
	SetType(v string) *DescribeConfigGroupRequest
	GetType() *string
}

type DescribeConfigGroupRequest struct {
	// The configuration group ID.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of configuration group IDs.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The configuration group name.
	//
	// example:
	//
	// Scheduled task configuration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type used by the configuration group.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID. This feature is not region-specific. Set this parameter to `cn-shanghai`.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of configuration group statuses.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// The configuration group type.
	//
	// example:
	//
	// Timer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeConfigGroupRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *DescribeConfigGroupRequest) GetName() *string {
	return s.Name
}

func (s *DescribeConfigGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConfigGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConfigGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeConfigGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConfigGroupRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *DescribeConfigGroupRequest) GetType() *string {
	return s.Type
}

func (s *DescribeConfigGroupRequest) SetGroupId(v string) *DescribeConfigGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeConfigGroupRequest) SetGroupIds(v []*string) *DescribeConfigGroupRequest {
	s.GroupIds = v
	return s
}

func (s *DescribeConfigGroupRequest) SetName(v string) *DescribeConfigGroupRequest {
	s.Name = &v
	return s
}

func (s *DescribeConfigGroupRequest) SetPageNumber(v int32) *DescribeConfigGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeConfigGroupRequest) SetPageSize(v int32) *DescribeConfigGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConfigGroupRequest) SetProductType(v string) *DescribeConfigGroupRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeConfigGroupRequest) SetRegionId(v string) *DescribeConfigGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConfigGroupRequest) SetStatuses(v []*string) *DescribeConfigGroupRequest {
	s.Statuses = v
	return s
}

func (s *DescribeConfigGroupRequest) SetType(v string) *DescribeConfigGroupRequest {
	s.Type = &v
	return s
}

func (s *DescribeConfigGroupRequest) Validate() error {
	return dara.Validate(s)
}
