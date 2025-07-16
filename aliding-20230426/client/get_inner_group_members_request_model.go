// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetInnerGroupMembersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetInnerGroupMembersRequest
	GetNextToken() *string
	SetOpenConversationId(v string) *GetInnerGroupMembersRequest
	GetOpenConversationId() *string
}

type GetInnerGroupMembersRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// UZr*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cidB8Pzg****FIWPv2PMA==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
}

func (s GetInnerGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInnerGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetInnerGroupMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInnerGroupMembersRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetInnerGroupMembersRequest) SetMaxResults(v int32) *GetInnerGroupMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *GetInnerGroupMembersRequest) SetNextToken(v string) *GetInnerGroupMembersRequest {
	s.NextToken = &v
	return s
}

func (s *GetInnerGroupMembersRequest) SetOpenConversationId(v string) *GetInnerGroupMembersRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetInnerGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
