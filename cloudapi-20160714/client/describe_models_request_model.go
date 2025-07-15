// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeModelsRequest
	GetGroupId() *string
	SetModelId(v string) *DescribeModelsRequest
	GetModelId() *string
	SetModelName(v string) *DescribeModelsRequest
	GetModelName() *string
	SetPageNumber(v int32) *DescribeModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeModelsRequest
	GetPageSize() *int32
	SetTag(v []*DescribeModelsRequestTag) *DescribeModelsRequest
	GetTag() []*DescribeModelsRequestTag
}

type DescribeModelsRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30e792398d6c4569b04c0e53a3494381
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the model.
	//
	// example:
	//
	// 123
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	Tag []*DescribeModelsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeModelsRequest) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeModelsRequest) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelsRequest) GetTag() []*DescribeModelsRequestTag {
	return s.Tag
}

func (s *DescribeModelsRequest) SetGroupId(v string) *DescribeModelsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeModelsRequest) SetModelId(v string) *DescribeModelsRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeModelsRequest) SetModelName(v string) *DescribeModelsRequest {
	s.ModelName = &v
	return s
}

func (s *DescribeModelsRequest) SetPageNumber(v int32) *DescribeModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelsRequest) SetPageSize(v int32) *DescribeModelsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeModelsRequest) SetTag(v []*DescribeModelsRequestTag) *DescribeModelsRequest {
	s.Tag = v
	return s
}

func (s *DescribeModelsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeModelsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeModelsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeModelsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeModelsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeModelsRequestTag) SetKey(v string) *DescribeModelsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeModelsRequestTag) SetValue(v string) *DescribeModelsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeModelsRequestTag) Validate() error {
	return dara.Validate(s)
}
