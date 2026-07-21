// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantBooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *ListTextbookAssistantBooksRequest
	GetAuthToken() *string
	SetBookId(v string) *ListTextbookAssistantBooksRequest
	GetBookId() *string
	SetGrade(v string) *ListTextbookAssistantBooksRequest
	GetGrade() *string
	SetMaxResults(v string) *ListTextbookAssistantBooksRequest
	GetMaxResults() *string
	SetPage(v string) *ListTextbookAssistantBooksRequest
	GetPage() *string
	SetVersion(v string) *ListTextbookAssistantBooksRequest
	GetVersion() *string
	SetVolume(v string) *ListTextbookAssistantBooksRequest
	GetVolume() *string
}

type ListTextbookAssistantBooksRequest struct {
	// The authorization token for the API call. You can obtain this token by calling the authorization API for the AI textbook assistant feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// tc_197bf5bb81889cc79eb51ae9b8c0cea3
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// The book ID.
	//
	// example:
	//
	// 231698
	BookId *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	// The grade level. The value is a string from "1" to "9".
	//
	// example:
	//
	// 1
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// The maximum number of results to return per page. The value cannot exceed 20.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *string `json:"page,omitempty" xml:"page,omitempty"`
	// The textbook version.
	//
	// example:
	//
	// 人教版
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The volume. Valid values: 0 (all-in-one volume), 1 (first volume), and 2 (second volume).
	//
	// example:
	//
	// 1
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s ListTextbookAssistantBooksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBooksRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListTextbookAssistantBooksRequest) GetBookId() *string {
	return s.BookId
}

func (s *ListTextbookAssistantBooksRequest) GetGrade() *string {
	return s.Grade
}

func (s *ListTextbookAssistantBooksRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListTextbookAssistantBooksRequest) GetPage() *string {
	return s.Page
}

func (s *ListTextbookAssistantBooksRequest) GetVersion() *string {
	return s.Version
}

func (s *ListTextbookAssistantBooksRequest) GetVolume() *string {
	return s.Volume
}

func (s *ListTextbookAssistantBooksRequest) SetAuthToken(v string) *ListTextbookAssistantBooksRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetBookId(v string) *ListTextbookAssistantBooksRequest {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetGrade(v string) *ListTextbookAssistantBooksRequest {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetMaxResults(v string) *ListTextbookAssistantBooksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetPage(v string) *ListTextbookAssistantBooksRequest {
	s.Page = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetVersion(v string) *ListTextbookAssistantBooksRequest {
	s.Version = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetVolume(v string) *ListTextbookAssistantBooksRequest {
	s.Volume = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) Validate() error {
	return dara.Validate(s)
}
