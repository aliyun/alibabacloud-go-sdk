// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentChunkListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetDocumentChunkListResponseBody
	GetCost() *int64
	SetData(v *GetDocumentChunkListResponseBodyData) *GetDocumentChunkListResponseBody
	GetData() *GetDocumentChunkListResponseBodyData
	SetDataType(v string) *GetDocumentChunkListResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetDocumentChunkListResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetDocumentChunkListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentChunkListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentChunkListResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetDocumentChunkListResponseBody
	GetTime() *string
}

type GetDocumentChunkListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDocumentChunkListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 2B8F6DC9-6FAF-576F-9095-CCD90FB2BDDF
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

func (s GetDocumentChunkListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentChunkListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetDocumentChunkListResponseBody) GetData() *GetDocumentChunkListResponseBodyData {
	return s.Data
}

func (s *GetDocumentChunkListResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDocumentChunkListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDocumentChunkListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentChunkListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentChunkListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentChunkListResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetDocumentChunkListResponseBody) SetCost(v int64) *GetDocumentChunkListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetData(v *GetDocumentChunkListResponseBodyData) *GetDocumentChunkListResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetDataType(v string) *GetDocumentChunkListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetErrCode(v string) *GetDocumentChunkListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetMessage(v string) *GetDocumentChunkListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetRequestId(v string) *GetDocumentChunkListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetSuccess(v bool) *GetDocumentChunkListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetTime(v string) *GetDocumentChunkListResponseBody {
	s.Time = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDocumentChunkListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetDocumentChunkListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetDocumentChunkListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetDocumentChunkListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDocumentChunkListResponseBodyData) GetRecords() []*GetDocumentChunkListResponseBodyDataRecords {
	return s.Records
}

func (s *GetDocumentChunkListResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *GetDocumentChunkListResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *GetDocumentChunkListResponseBodyData) SetCurrentPage(v int64) *GetDocumentChunkListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetPageSize(v int64) *GetDocumentChunkListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetRecords(v []*GetDocumentChunkListResponseBodyDataRecords) *GetDocumentChunkListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetTotalPages(v int64) *GetDocumentChunkListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetTotalRecords(v int64) *GetDocumentChunkListResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDocumentChunkListResponseBodyDataRecords struct {
	// example:
	//
	// 28377468263482764
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// oss-xxxx-hangzhou.com/test.pdf
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 8947387648356
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// jhsdvne
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 947538465
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*GetDocumentChunkListResponseBodyDataRecordsPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 9848346548365
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetDocumentChunkListResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetChunkId() *string {
	return s.ChunkId
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetChunkMeta() map[string]interface{} {
	return s.ChunkMeta
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetChunkOssUrl() *string {
	return s.ChunkOssUrl
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetChunkText() *string {
	return s.ChunkText
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetChunkType() *string {
	return s.ChunkType
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetDocId() *string {
	return s.DocId
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetFileType() *string {
	return s.FileType
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetLibraryName() *string {
	return s.LibraryName
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetNextChunkId() *string {
	return s.NextChunkId
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetPos() []*GetDocumentChunkListResponseBodyDataRecordsPos {
	return s.Pos
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetPreChunkId() *string {
	return s.PreChunkId
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetScore() *float32 {
	return s.Score
}

func (s *GetDocumentChunkListResponseBodyDataRecords) GetTitle() *string {
	return s.Title
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkMeta(v map[string]interface{}) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkMeta = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkOssUrl(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkOssUrl = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkText(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkText = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkType(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkType = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetDocId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetFileType(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetLibraryId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetLibraryName(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.LibraryName = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetNextChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.NextChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetPos(v []*GetDocumentChunkListResponseBodyDataRecordsPos) *GetDocumentChunkListResponseBodyDataRecords {
	s.Pos = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetPreChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.PreChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetScore(v float32) *GetDocumentChunkListResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetTitle(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}

type GetDocumentChunkListResponseBodyDataRecordsPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s GetDocumentChunkListResponseBodyDataRecordsPos) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyDataRecordsPos) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) GetAxisArray() []*float64 {
	return s.AxisArray
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) GetPage() *int32 {
	return s.Page
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) GetTextHighlightArea() []*int32 {
	return s.TextHighlightArea
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetAxisArray(v []*float64) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.AxisArray = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetPage(v int32) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.Page = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetTextHighlightArea(v []*int32) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.TextHighlightArea = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) Validate() error {
	return dara.Validate(s)
}
