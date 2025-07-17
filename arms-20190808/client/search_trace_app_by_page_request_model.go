// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *SearchTraceAppByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchTraceAppByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *SearchTraceAppByPageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SearchTraceAppByPageRequest
	GetResourceGroupId() *string
	SetTags(v []*SearchTraceAppByPageRequestTags) *SearchTraceAppByPageRequest
	GetTags() []*SearchTraceAppByPageRequestTags
	SetTraceAppName(v string) *SearchTraceAppByPageRequest
	GetTraceAppName() *string
}

type SearchTraceAppByPageRequest struct {
	// The number of the page to return. Default value: `1`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: `10`.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A list of tags.
	Tags []*SearchTraceAppByPageRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTraceAppByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTraceAppByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByPageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchTraceAppByPageRequest) GetTags() []*SearchTraceAppByPageRequestTags {
	return s.Tags
}

func (s *SearchTraceAppByPageRequest) GetTraceAppName() *string {
	return s.TraceAppName
}

func (s *SearchTraceAppByPageRequest) SetPageNumber(v int32) *SearchTraceAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetPageSize(v int32) *SearchTraceAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetRegionId(v string) *SearchTraceAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetResourceGroupId(v string) *SearchTraceAppByPageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetTags(v []*SearchTraceAppByPageRequestTags) *SearchTraceAppByPageRequest {
	s.Tags = v
	return s
}

func (s *SearchTraceAppByPageRequest) SetTraceAppName(v string) *SearchTraceAppByPageRequest {
	s.TraceAppName = &v
	return s
}

func (s *SearchTraceAppByPageRequest) Validate() error {
	return dara.Validate(s)
}

type SearchTraceAppByPageRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTraceAppByPageRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageRequestTags) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageRequestTags) GetKey() *string {
	return s.Key
}

func (s *SearchTraceAppByPageRequestTags) GetValue() *string {
	return s.Value
}

func (s *SearchTraceAppByPageRequestTags) SetKey(v string) *SearchTraceAppByPageRequestTags {
	s.Key = &v
	return s
}

func (s *SearchTraceAppByPageRequestTags) SetValue(v string) *SearchTraceAppByPageRequestTags {
	s.Value = &v
	return s
}

func (s *SearchTraceAppByPageRequestTags) Validate() error {
	return dara.Validate(s)
}
