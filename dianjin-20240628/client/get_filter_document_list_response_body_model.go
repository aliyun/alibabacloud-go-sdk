// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFilterDocumentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetFilterDocumentListResponseBody
	GetCost() *int64
	SetData(v *GetFilterDocumentListResponseBodyData) *GetFilterDocumentListResponseBody
	GetData() *GetFilterDocumentListResponseBodyData
	SetDataType(v string) *GetFilterDocumentListResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetFilterDocumentListResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetFilterDocumentListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFilterDocumentListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFilterDocumentListResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetFilterDocumentListResponseBody
	GetTime() *string
}

type GetFilterDocumentListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                 `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetFilterDocumentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 7ADF010C-FD89-569D-A079-2D4D5247E943
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetFilterDocumentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetFilterDocumentListResponseBody) GetData() *GetFilterDocumentListResponseBodyData {
	return s.Data
}

func (s *GetFilterDocumentListResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetFilterDocumentListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetFilterDocumentListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFilterDocumentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFilterDocumentListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFilterDocumentListResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetFilterDocumentListResponseBody) SetCost(v int64) *GetFilterDocumentListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetData(v *GetFilterDocumentListResponseBodyData) *GetFilterDocumentListResponseBody {
	s.Data = v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetDataType(v string) *GetFilterDocumentListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetErrCode(v string) *GetFilterDocumentListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetMessage(v string) *GetFilterDocumentListResponseBody {
	s.Message = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetRequestId(v string) *GetFilterDocumentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetSuccess(v bool) *GetFilterDocumentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetTime(v string) *GetFilterDocumentListResponseBody {
	s.Time = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFilterDocumentListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetFilterDocumentListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetFilterDocumentListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetFilterDocumentListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetFilterDocumentListResponseBodyData) GetRecords() []*GetFilterDocumentListResponseBodyDataRecords {
	return s.Records
}

func (s *GetFilterDocumentListResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *GetFilterDocumentListResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *GetFilterDocumentListResponseBodyData) SetCurrentPage(v int64) *GetFilterDocumentListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetPageSize(v int64) *GetFilterDocumentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetRecords(v []*GetFilterDocumentListResponseBodyDataRecords) *GetFilterDocumentListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetTotalPages(v int64) *GetFilterDocumentListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetTotalRecords(v int64) *GetFilterDocumentListResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetFilterDocumentListResponseBodyDataRecords struct {
	// example:
	//
	// 29368126816
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// {"a": "1"}
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
	// sdfgsjdfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// WaitRefresh
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// null
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetFilterDocumentListResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetDocId() *string {
	return s.DocId
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetDocumentMeta() map[string]interface{} {
	return s.DocumentMeta
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetFileType() *string {
	return s.FileType
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetTitle() *string {
	return s.Title
}

func (s *GetFilterDocumentListResponseBodyDataRecords) GetUrl() *string {
	return s.Url
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetDocId(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetDocumentMeta(v map[string]interface{}) *GetFilterDocumentListResponseBodyDataRecords {
	s.DocumentMeta = v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetFileType(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetGmtCreate(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetGmtModified(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetLibraryId(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetStatusCode(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.StatusCode = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetTitle(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetUrl(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.Url = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
