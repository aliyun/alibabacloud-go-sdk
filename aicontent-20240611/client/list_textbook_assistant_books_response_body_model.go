// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantBooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTextbookAssistantBooksResponseBodyData) *ListTextbookAssistantBooksResponseBody
	GetData() *ListTextbookAssistantBooksResponseBodyData
	SetErrCode(v string) *ListTextbookAssistantBooksResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListTextbookAssistantBooksResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListTextbookAssistantBooksResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTextbookAssistantBooksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTextbookAssistantBooksResponseBody
	GetSuccess() *bool
}

type ListTextbookAssistantBooksResponseBody struct {
	Data *ListTextbookAssistantBooksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B_USER_NOT_FOUND_EXCEPTION
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B695B377-7029-5805-9DE2-1AAE06C1BF6B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantBooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBody) GetData() *ListTextbookAssistantBooksResponseBodyData {
	return s.Data
}

func (s *ListTextbookAssistantBooksResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListTextbookAssistantBooksResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListTextbookAssistantBooksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTextbookAssistantBooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTextbookAssistantBooksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTextbookAssistantBooksResponseBody) SetData(v *ListTextbookAssistantBooksResponseBodyData) *ListTextbookAssistantBooksResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetErrCode(v string) *ListTextbookAssistantBooksResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetErrMessage(v string) *ListTextbookAssistantBooksResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantBooksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetRequestId(v string) *ListTextbookAssistantBooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetSuccess(v bool) *ListTextbookAssistantBooksResponseBody {
	s.Success = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTextbookAssistantBooksResponseBodyData struct {
	BookList       []*ListTextbookAssistantBooksResponseBodyDataBookList     `json:"bookList,omitempty" xml:"bookList,omitempty" type:"Repeated"`
	PaginationData *ListTextbookAssistantBooksResponseBodyDataPaginationData `json:"paginationData,omitempty" xml:"paginationData,omitempty" type:"Struct"`
}

func (s ListTextbookAssistantBooksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBodyData) GetBookList() []*ListTextbookAssistantBooksResponseBodyDataBookList {
	return s.BookList
}

func (s *ListTextbookAssistantBooksResponseBodyData) GetPaginationData() *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	return s.PaginationData
}

func (s *ListTextbookAssistantBooksResponseBodyData) SetBookList(v []*ListTextbookAssistantBooksResponseBodyDataBookList) *ListTextbookAssistantBooksResponseBodyData {
	s.BookList = v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyData) SetPaginationData(v *ListTextbookAssistantBooksResponseBodyDataPaginationData) *ListTextbookAssistantBooksResponseBodyData {
	s.PaginationData = v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyData) Validate() error {
	if s.BookList != nil {
		for _, item := range s.BookList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PaginationData != nil {
		if err := s.PaginationData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTextbookAssistantBooksResponseBodyDataBookList struct {
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// example:
	//
	// 231698
	BookId   *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	BookName *string `json:"bookName,omitempty" xml:"bookName,omitempty"`
	// example:
	//
	// null
	CoverImage *string `json:"coverImage,omitempty" xml:"coverImage,omitempty"`
	// example:
	//
	// 2024-7（1）
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// example:
	//
	// 3
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 2024-7（1）
	Impression *string `json:"impression,omitempty" xml:"impression,omitempty"`
	// example:
	//
	// 9787107382505
	Isbn      *string `json:"isbn,omitempty" xml:"isbn,omitempty"`
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
	// example:
	//
	// ENGLISH
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 0
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s ListTextbookAssistantBooksResponseBodyDataBookList) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBodyDataBookList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetAuthor() *string {
	return s.Author
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetBookId() *string {
	return s.BookId
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetBookName() *string {
	return s.BookName
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetCoverImage() *string {
	return s.CoverImage
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetEdition() *string {
	return s.Edition
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetGrade() *string {
	return s.Grade
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetImpression() *string {
	return s.Impression
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetIsbn() *string {
	return s.Isbn
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetPublisher() *string {
	return s.Publisher
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetSubject() *string {
	return s.Subject
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetVersion() *string {
	return s.Version
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) GetVolume() *string {
	return s.Volume
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetAuthor(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Author = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetBookId(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetBookName(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.BookName = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetCoverImage(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.CoverImage = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetEdition(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Edition = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetGrade(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetImpression(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Impression = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetIsbn(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Isbn = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetPublisher(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Publisher = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetSubject(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Subject = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetVersion(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Version = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetVolume(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Volume = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) Validate() error {
	return dara.Validate(s)
}

type ListTextbookAssistantBooksResponseBodyDataPaginationData struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTextbookAssistantBooksResponseBodyDataPaginationData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBodyDataPaginationData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) SetCurrentPage(v int32) *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	s.CurrentPage = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) SetMaxResults(v int32) *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	s.MaxResults = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) SetTotalCount(v int32) *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	s.TotalCount = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) Validate() error {
	return dara.Validate(s)
}
