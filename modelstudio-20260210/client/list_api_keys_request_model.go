// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ListApiKeysRequest
	GetApiKeyId() *int64
	SetDescription(v string) *ListApiKeysRequest
	GetDescription() *string
	SetMaxResults(v int32) *ListApiKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiKeysRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListApiKeysRequest
	GetWorkspaceId() *string
}

type ListApiKeysRequest struct {
	// Exact search by API Key ID.
	//
	// example:
	//
	// 3076140
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// Fuzzy search by description keyword.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Page size.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Used to return more results. This parameter is not required for the first query. The token required for subsequent queries can be obtained from the returned results.
	//
	// example:
	//
	// w9Z+S5+TZyw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Exact search by workspace ID.
	//
	// example:
	//
	// ws-ac3ef438bec22dc5
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListApiKeysRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ListApiKeysRequest) GetDescription() *string {
	return s.Description
}

func (s *ListApiKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiKeysRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiKeysRequest) SetApiKeyId(v int64) *ListApiKeysRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ListApiKeysRequest) SetDescription(v string) *ListApiKeysRequest {
	s.Description = &v
	return s
}

func (s *ListApiKeysRequest) SetMaxResults(v int32) *ListApiKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiKeysRequest) SetNextToken(v string) *ListApiKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiKeysRequest) SetWorkspaceId(v string) *ListApiKeysRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
