// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetDocumentListResponseBody
	GetCost() *int64
	SetData(v *GetDocumentListResponseBodyData) *GetDocumentListResponseBody
	GetData() *GetDocumentListResponseBodyData
	SetDataType(v string) *GetDocumentListResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetDocumentListResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetDocumentListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentListResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetDocumentListResponseBody
	GetTime() *string
}

type GetDocumentListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDocumentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDocumentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetDocumentListResponseBody) GetData() *GetDocumentListResponseBodyData {
	return s.Data
}

func (s *GetDocumentListResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDocumentListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDocumentListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentListResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetDocumentListResponseBody) SetCost(v int64) *GetDocumentListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentListResponseBody) SetData(v *GetDocumentListResponseBodyData) *GetDocumentListResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentListResponseBody) SetDataType(v string) *GetDocumentListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentListResponseBody) SetErrCode(v string) *GetDocumentListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentListResponseBody) SetMessage(v string) *GetDocumentListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentListResponseBody) SetRequestId(v string) *GetDocumentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentListResponseBody) SetSuccess(v bool) *GetDocumentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentListResponseBody) SetTime(v string) *GetDocumentListResponseBody {
	s.Time = &v
	return s
}

func (s *GetDocumentListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetDocumentListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetDocumentListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetDocumentListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDocumentListResponseBodyData) GetRecords() []*GetDocumentListResponseBodyDataRecords {
	return s.Records
}

func (s *GetDocumentListResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *GetDocumentListResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *GetDocumentListResponseBodyData) SetCurrentPage(v int64) *GetDocumentListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetPageSize(v int64) *GetDocumentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetRecords(v []*GetDocumentListResponseBodyDataRecords) *GetDocumentListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetDocumentListResponseBodyData) SetTotalPages(v int64) *GetDocumentListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetTotalRecords(v int64) *GetDocumentListResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *GetDocumentListResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDocumentListResponseBodyDataRecords struct {
	// example:
	//
	// 8326748346
	DocId        *string                `json:"docId,omitempty" xml:"docId,omitempty"`
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// skjdhshbv
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// WaitRefresh
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// null
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDocumentListResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBodyDataRecords) GetDocId() *string {
	return s.DocId
}

func (s *GetDocumentListResponseBodyDataRecords) GetDocumentMeta() map[string]interface{} {
	return s.DocumentMeta
}

func (s *GetDocumentListResponseBodyDataRecords) GetFileType() *string {
	return s.FileType
}

func (s *GetDocumentListResponseBodyDataRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDocumentListResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDocumentListResponseBodyDataRecords) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetDocumentListResponseBodyDataRecords) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetDocumentListResponseBodyDataRecords) GetTitle() *string {
	return s.Title
}

func (s *GetDocumentListResponseBodyDataRecords) GetUrl() *string {
	return s.Url
}

func (s *GetDocumentListResponseBodyDataRecords) SetDocId(v string) *GetDocumentListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetDocumentMeta(v map[string]interface{}) *GetDocumentListResponseBodyDataRecords {
	s.DocumentMeta = v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetFileType(v string) *GetDocumentListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetGmtCreate(v string) *GetDocumentListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetGmtModified(v string) *GetDocumentListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetLibraryId(v string) *GetDocumentListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetStatusCode(v string) *GetDocumentListResponseBodyDataRecords {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetTitle(v string) *GetDocumentListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetUrl(v string) *GetDocumentListResponseBodyDataRecords {
	s.Url = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
