// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantBookDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *ListTextbookAssistantBookDirectoriesRequest
	GetAuthToken() *string
	SetBookId(v string) *ListTextbookAssistantBookDirectoriesRequest
	GetBookId() *string
	SetScenario(v string) *ListTextbookAssistantBookDirectoriesRequest
	GetScenario() *string
}

type ListTextbookAssistantBookDirectoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 231698
	BookId *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListTextbookAssistantBookDirectoriesRequest) GetBookId() *string {
	return s.BookId
}

func (s *ListTextbookAssistantBookDirectoriesRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ListTextbookAssistantBookDirectoriesRequest) SetAuthToken(v string) *ListTextbookAssistantBookDirectoriesRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesRequest) SetBookId(v string) *ListTextbookAssistantBookDirectoriesRequest {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesRequest) SetScenario(v string) *ListTextbookAssistantBookDirectoriesRequest {
	s.Scenario = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
