// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoolAppCode(v string) *BatchQueryGroupMemberRequest
	GetCoolAppCode() *string
	SetMaxResults(v int64) *BatchQueryGroupMemberRequest
	GetMaxResults() *int64
	SetNextToken(v string) *BatchQueryGroupMemberRequest
	GetNextToken() *string
	SetOpenConversationId(v string) *BatchQueryGroupMemberRequest
	GetOpenConversationId() *string
	SetTenantContext(v *BatchQueryGroupMemberRequestTenantContext) *BatchQueryGroupMemberRequest
	GetTenantContext() *BatchQueryGroupMemberRequestTenantContext
}

type BatchQueryGroupMemberRequest struct {
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
	OpenConversationId *string                                    `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	TenantContext      *BatchQueryGroupMemberRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s BatchQueryGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberRequest) GetCoolAppCode() *string {
	return s.CoolAppCode
}

func (s *BatchQueryGroupMemberRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *BatchQueryGroupMemberRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *BatchQueryGroupMemberRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *BatchQueryGroupMemberRequest) GetTenantContext() *BatchQueryGroupMemberRequestTenantContext {
	return s.TenantContext
}

func (s *BatchQueryGroupMemberRequest) SetCoolAppCode(v string) *BatchQueryGroupMemberRequest {
	s.CoolAppCode = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetMaxResults(v int64) *BatchQueryGroupMemberRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetNextToken(v string) *BatchQueryGroupMemberRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetOpenConversationId(v string) *BatchQueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetTenantContext(v *BatchQueryGroupMemberRequestTenantContext) *BatchQueryGroupMemberRequest {
	s.TenantContext = v
	return s
}

func (s *BatchQueryGroupMemberRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchQueryGroupMemberRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s BatchQueryGroupMemberRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberRequestTenantContext) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchQueryGroupMemberRequestTenantContext) SetTenantId(v string) *BatchQueryGroupMemberRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *BatchQueryGroupMemberRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
