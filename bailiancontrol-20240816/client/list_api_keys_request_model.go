// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *ListApiKeysRequest
	GetApiKeyId() *string
	SetDescription(v string) *ListApiKeysRequest
	GetDescription() *string
	SetMaxResults(v int32) *ListApiKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiKeysRequest
	GetNextToken() *string
	SetSkip(v int32) *ListApiKeysRequest
	GetSkip() *int32
	SetUid(v string) *ListApiKeysRequest
	GetUid() *string
	SetWorkspaceId(v string) *ListApiKeysRequest
	GetWorkspaceId() *string
}

type ListApiKeysRequest struct {
	ApiKeyId    *string `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Skip        *int32  `json:"skip,omitempty" xml:"skip,omitempty"`
	Uid         *string `json:"uid,omitempty" xml:"uid,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListApiKeysRequest) GetApiKeyId() *string {
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

func (s *ListApiKeysRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListApiKeysRequest) GetUid() *string {
	return s.Uid
}

func (s *ListApiKeysRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiKeysRequest) SetApiKeyId(v string) *ListApiKeysRequest {
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

func (s *ListApiKeysRequest) SetSkip(v int32) *ListApiKeysRequest {
	s.Skip = &v
	return s
}

func (s *ListApiKeysRequest) SetUid(v string) *ListApiKeysRequest {
	s.Uid = &v
	return s
}

func (s *ListApiKeysRequest) SetWorkspaceId(v string) *ListApiKeysRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
