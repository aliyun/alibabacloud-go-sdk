// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextDatabaseApiKeysRequest
	GetMaxResults() *int32
	SetMemberId(v string) *ListContextDatabaseApiKeysRequest
	GetMemberId() *string
	SetNextToken(v string) *ListContextDatabaseApiKeysRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListContextDatabaseApiKeysRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseApiKeysRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// (null)
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseApiKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextDatabaseApiKeysRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *ListContextDatabaseApiKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextDatabaseApiKeysRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseApiKeysRequest) SetMaxResults(v int32) *ListContextDatabaseApiKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContextDatabaseApiKeysRequest) SetMemberId(v string) *ListContextDatabaseApiKeysRequest {
	s.MemberId = &v
	return s
}

func (s *ListContextDatabaseApiKeysRequest) SetNextToken(v string) *ListContextDatabaseApiKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListContextDatabaseApiKeysRequest) SetWorkspaceId(v string) *ListContextDatabaseApiKeysRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
