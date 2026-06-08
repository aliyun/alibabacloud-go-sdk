// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyApplicationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefSchema(v string) *ListMyApplicationsShrinkRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListMyApplicationsShrinkRequest
	GetEndTime() *int64
	SetNextToken(v string) *ListMyApplicationsShrinkRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListMyApplicationsShrinkRequest
	GetPageSize() *int32
	SetResourceShrink(v string) *ListMyApplicationsShrinkRequest
	GetResourceShrink() *string
	SetResourceTypeShrink(v string) *ListMyApplicationsShrinkRequest
	GetResourceTypeShrink() *string
	SetStartTime(v int64) *ListMyApplicationsShrinkRequest
	GetStartTime() *int64
	SetStatusesShrink(v string) *ListMyApplicationsShrinkRequest
	GetStatusesShrink() *string
}

type ListMyApplicationsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1779724799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// eyJpZCI6MTIzfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceShrink *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// This parameter is required.
	ResourceTypeShrink *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1771948800000
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListMyApplicationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyApplicationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMyApplicationsShrinkRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyApplicationsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMyApplicationsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyApplicationsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyApplicationsShrinkRequest) GetResourceShrink() *string {
	return s.ResourceShrink
}

func (s *ListMyApplicationsShrinkRequest) GetResourceTypeShrink() *string {
	return s.ResourceTypeShrink
}

func (s *ListMyApplicationsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMyApplicationsShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListMyApplicationsShrinkRequest) SetDefSchema(v string) *ListMyApplicationsShrinkRequest {
	s.DefSchema = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetEndTime(v int64) *ListMyApplicationsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetNextToken(v string) *ListMyApplicationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetPageSize(v int32) *ListMyApplicationsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetResourceShrink(v string) *ListMyApplicationsShrinkRequest {
	s.ResourceShrink = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetResourceTypeShrink(v string) *ListMyApplicationsShrinkRequest {
	s.ResourceTypeShrink = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetStartTime(v int64) *ListMyApplicationsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) SetStatusesShrink(v string) *ListMyApplicationsShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListMyApplicationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
