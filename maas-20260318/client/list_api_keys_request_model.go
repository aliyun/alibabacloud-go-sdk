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
	// API Key ID。
	//
	// example:
	//
	// 3685841
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// VHayKbYhhm4=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// llm-ep8ba0dr6seiddri
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
