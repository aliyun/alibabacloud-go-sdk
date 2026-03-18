// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantArticlesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *ListTextbookAssistantArticlesRequest
	GetAuthToken() *string
	SetDirectoryId(v string) *ListTextbookAssistantArticlesRequest
	GetDirectoryId() *string
}

type ListTextbookAssistantArticlesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_a893b8492c4be046cbc906c566aeb8c9
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 90aa861b4d9311efbe6e0c42a106bb02
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
}

func (s ListTextbookAssistantArticlesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticlesRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListTextbookAssistantArticlesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTextbookAssistantArticlesRequest) SetAuthToken(v string) *ListTextbookAssistantArticlesRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantArticlesRequest) SetDirectoryId(v string) *ListTextbookAssistantArticlesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantArticlesRequest) Validate() error {
	return dara.Validate(s)
}
