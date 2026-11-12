// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialType(v string) *ListCredentialsRequest
	GetCredentialType() *string
	SetMaxResults(v int32) *ListCredentialsRequest
	GetMaxResults() *int32
	SetName(v string) *ListCredentialsRequest
	GetName() *string
	SetNameLike(v string) *ListCredentialsRequest
	GetNameLike() *string
	SetNextToken(v string) *ListCredentialsRequest
	GetNextToken() *string
}

type ListCredentialsRequest struct {
	// The credential type used for filtering. Valid values: apiKey and mcpHeader.
	//
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// The maximum number of records to return per page. Valid values: 1 to 100. If this parameter is not specified, 10 records are returned by default.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The credential name used for filtering.
	//
	// example:
	//
	// credentialxxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The keyword used to perform fuzzy matching on credential names.
	//
	// example:
	//
	// model
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The cursor used to query the next page. Leave this parameter empty in the first call. In subsequent calls, specify the nextToken value returned in the previous response.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ListCredentialsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialsRequest) GetName() *string {
	return s.Name
}

func (s *ListCredentialsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListCredentialsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialsRequest) SetCredentialType(v string) *ListCredentialsRequest {
	s.CredentialType = &v
	return s
}

func (s *ListCredentialsRequest) SetMaxResults(v int32) *ListCredentialsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialsRequest) SetName(v string) *ListCredentialsRequest {
	s.Name = &v
	return s
}

func (s *ListCredentialsRequest) SetNameLike(v string) *ListCredentialsRequest {
	s.NameLike = &v
	return s
}

func (s *ListCredentialsRequest) SetNextToken(v string) *ListCredentialsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
