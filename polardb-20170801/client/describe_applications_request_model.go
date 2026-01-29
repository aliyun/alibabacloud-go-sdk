// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v string) *DescribeApplicationsRequest
	GetApplicationIds() *string
	SetApplicationTypes(v string) *DescribeApplicationsRequest
	GetApplicationTypes() *string
	SetDBClusterId(v string) *DescribeApplicationsRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribeApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApplicationsRequest
	GetRegionId() *string
	SetTag(v []*DescribeApplicationsRequestTag) *DescribeApplicationsRequest
	GetTag() []*DescribeApplicationsRequestTag
}

type DescribeApplicationsRequest struct {
	// example:
	//
	// pa-**************
	ApplicationIds   *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	ApplicationTypes *string `json:"ApplicationTypes,omitempty" xml:"ApplicationTypes,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag      []*DescribeApplicationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsRequest) GetApplicationIds() *string {
	return s.ApplicationIds
}

func (s *DescribeApplicationsRequest) GetApplicationTypes() *string {
	return s.ApplicationTypes
}

func (s *DescribeApplicationsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationsRequest) GetTag() []*DescribeApplicationsRequestTag {
	return s.Tag
}

func (s *DescribeApplicationsRequest) SetApplicationIds(v string) *DescribeApplicationsRequest {
	s.ApplicationIds = &v
	return s
}

func (s *DescribeApplicationsRequest) SetApplicationTypes(v string) *DescribeApplicationsRequest {
	s.ApplicationTypes = &v
	return s
}

func (s *DescribeApplicationsRequest) SetDBClusterId(v string) *DescribeApplicationsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApplicationsRequest) SetPageNumber(v int32) *DescribeApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationsRequest) SetPageSize(v int32) *DescribeApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationsRequest) SetRegionId(v string) *DescribeApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationsRequest) SetTag(v []*DescribeApplicationsRequestTag) *DescribeApplicationsRequest {
	s.Tag = v
	return s
}

func (s *DescribeApplicationsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationsRequestTag struct {
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationsRequestTag) SetKey(v string) *DescribeApplicationsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeApplicationsRequestTag) SetValue(v string) *DescribeApplicationsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeApplicationsRequestTag) Validate() error {
	return dara.Validate(s)
}
