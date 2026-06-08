// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPendingApprovalsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTypesShrink(v string) *ListPendingApprovalsShrinkRequest
	GetAccessTypesShrink() *string
	SetDefSchema(v string) *ListPendingApprovalsShrinkRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListPendingApprovalsShrinkRequest
	GetEndTime() *int64
	SetGranteeShrink(v string) *ListPendingApprovalsShrinkRequest
	GetGranteeShrink() *string
	SetNextToken(v string) *ListPendingApprovalsShrinkRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListPendingApprovalsShrinkRequest
	GetPageSize() *int32
	SetResourceShrink(v string) *ListPendingApprovalsShrinkRequest
	GetResourceShrink() *string
	SetResourceTypeShrink(v string) *ListPendingApprovalsShrinkRequest
	GetResourceTypeShrink() *string
	SetStartTime(v int64) *ListPendingApprovalsShrinkRequest
	GetStartTime() *int64
}

type ListPendingApprovalsShrinkRequest struct {
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
}

func (s ListPendingApprovalsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsShrinkRequest) GetAccessTypesShrink() *string {
	return s.AccessTypesShrink
}

func (s *ListPendingApprovalsShrinkRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListPendingApprovalsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListPendingApprovalsShrinkRequest) GetGranteeShrink() *string {
	return s.GranteeShrink
}

func (s *ListPendingApprovalsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPendingApprovalsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPendingApprovalsShrinkRequest) GetResourceShrink() *string {
	return s.ResourceShrink
}

func (s *ListPendingApprovalsShrinkRequest) GetResourceTypeShrink() *string {
	return s.ResourceTypeShrink
}

func (s *ListPendingApprovalsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListPendingApprovalsShrinkRequest) SetAccessTypesShrink(v string) *ListPendingApprovalsShrinkRequest {
	s.AccessTypesShrink = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetDefSchema(v string) *ListPendingApprovalsShrinkRequest {
	s.DefSchema = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetEndTime(v int64) *ListPendingApprovalsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetGranteeShrink(v string) *ListPendingApprovalsShrinkRequest {
	s.GranteeShrink = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetNextToken(v string) *ListPendingApprovalsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetPageSize(v int32) *ListPendingApprovalsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetResourceShrink(v string) *ListPendingApprovalsShrinkRequest {
	s.ResourceShrink = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetResourceTypeShrink(v string) *ListPendingApprovalsShrinkRequest {
	s.ResourceTypeShrink = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) SetStartTime(v int64) *ListPendingApprovalsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListPendingApprovalsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
