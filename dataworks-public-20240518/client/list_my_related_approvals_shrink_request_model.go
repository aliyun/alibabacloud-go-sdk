// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyRelatedApprovalsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTypesShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetAccessTypesShrink() *string
	SetDefSchema(v string) *ListMyRelatedApprovalsShrinkRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListMyRelatedApprovalsShrinkRequest
	GetEndTime() *int64
	SetGranteeShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetGranteeShrink() *string
	SetNextToken(v string) *ListMyRelatedApprovalsShrinkRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListMyRelatedApprovalsShrinkRequest
	GetPageSize() *int32
	SetResourceShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetResourceShrink() *string
	SetResourceTypeShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetResourceTypeShrink() *string
	SetStartTime(v int64) *ListMyRelatedApprovalsShrinkRequest
	GetStartTime() *int64
	SetStatusesShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetStatusesShrink() *string
}

type ListMyRelatedApprovalsShrinkRequest struct {
	AccessTypesShrink *string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// example:
	//
	// 1779724799999
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GranteeShrink *string `json:"Grantee,omitempty" xml:"Grantee,omitempty"`
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
	//
	// example:
	//
	// ["table", "column"]
	ResourceTypeShrink *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1771948800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// WAIT_APPROVAL
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListMyRelatedApprovalsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetAccessTypesShrink() *string {
	return s.AccessTypesShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetGranteeShrink() *string {
	return s.GranteeShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetResourceShrink() *string {
	return s.ResourceShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetResourceTypeShrink() *string {
	return s.ResourceTypeShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetAccessTypesShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.AccessTypesShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetDefSchema(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetEndTime(v int64) *ListMyRelatedApprovalsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetGranteeShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.GranteeShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetNextToken(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetPageSize(v int32) *ListMyRelatedApprovalsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetResourceShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.ResourceShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetResourceTypeShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.ResourceTypeShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetStartTime(v int64) *ListMyRelatedApprovalsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetStatusesShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
