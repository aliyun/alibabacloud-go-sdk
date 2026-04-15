// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryGroupMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoolAppCode(v string) *BatchQueryGroupMemberShrinkRequest
	GetCoolAppCode() *string
	SetMaxResults(v int64) *BatchQueryGroupMemberShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *BatchQueryGroupMemberShrinkRequest
	GetNextToken() *string
	SetOpenConversationId(v string) *BatchQueryGroupMemberShrinkRequest
	GetOpenConversationId() *string
	SetTenantContextShrink(v string) *BatchQueryGroupMemberShrinkRequest
	GetTenantContextShrink() *string
}

type BatchQueryGroupMemberShrinkRequest struct {
	// example:
	//
	// COOLAPP_XXXXX
	CoolAppCode *string `json:"CoolAppCode,omitempty" xml:"CoolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// weqrwereqsadqaadfafa
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId  *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s BatchQueryGroupMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberShrinkRequest) GetCoolAppCode() *string {
	return s.CoolAppCode
}

func (s *BatchQueryGroupMemberShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *BatchQueryGroupMemberShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *BatchQueryGroupMemberShrinkRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *BatchQueryGroupMemberShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *BatchQueryGroupMemberShrinkRequest) SetCoolAppCode(v string) *BatchQueryGroupMemberShrinkRequest {
	s.CoolAppCode = &v
	return s
}

func (s *BatchQueryGroupMemberShrinkRequest) SetMaxResults(v int64) *BatchQueryGroupMemberShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryGroupMemberShrinkRequest) SetNextToken(v string) *BatchQueryGroupMemberShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberShrinkRequest) SetOpenConversationId(v string) *BatchQueryGroupMemberShrinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BatchQueryGroupMemberShrinkRequest) SetTenantContextShrink(v string) *BatchQueryGroupMemberShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *BatchQueryGroupMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
