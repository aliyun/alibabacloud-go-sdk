// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RecallDocumentResponseBody
	GetCost() *int64
	SetData(v *RecallDocumentResponseBodyData) *RecallDocumentResponseBody
	GetData() *RecallDocumentResponseBodyData
	SetDataType(v string) *RecallDocumentResponseBody
	GetDataType() *string
	SetErrCode(v string) *RecallDocumentResponseBody
	GetErrCode() *string
	SetMessage(v string) *RecallDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecallDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecallDocumentResponseBody
	GetSuccess() *bool
	SetTime(v string) *RecallDocumentResponseBody
	GetTime() *string
}

type RecallDocumentResponseBody struct {
	// example:
	//
	// 0
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RecallDocumentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0bc13a9517168617617186457e401f
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

func (s RecallDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RecallDocumentResponseBody) GetData() *RecallDocumentResponseBodyData {
	return s.Data
}

func (s *RecallDocumentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RecallDocumentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RecallDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecallDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecallDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecallDocumentResponseBody) GetTime() *string {
	return s.Time
}

func (s *RecallDocumentResponseBody) SetCost(v int64) *RecallDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *RecallDocumentResponseBody) SetData(v *RecallDocumentResponseBodyData) *RecallDocumentResponseBody {
	s.Data = v
	return s
}

func (s *RecallDocumentResponseBody) SetDataType(v string) *RecallDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *RecallDocumentResponseBody) SetErrCode(v string) *RecallDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RecallDocumentResponseBody) SetMessage(v string) *RecallDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *RecallDocumentResponseBody) SetRequestId(v string) *RecallDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallDocumentResponseBody) SetSuccess(v bool) *RecallDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *RecallDocumentResponseBody) SetTime(v string) *RecallDocumentResponseBody {
	s.Time = &v
	return s
}

func (s *RecallDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyData struct {
	ChunkList     []*RecallDocumentResponseBodyDataChunkList     `json:"chunkList,omitempty" xml:"chunkList,omitempty" type:"Repeated"`
	ChunkPartList []*RecallDocumentResponseBodyDataChunkPartList `json:"chunkPartList,omitempty" xml:"chunkPartList,omitempty" type:"Repeated"`
	ChunkTextList []*string                                      `json:"chunkTextList,omitempty" xml:"chunkTextList,omitempty" type:"Repeated"`
	Documents     []*RecallDocumentResponseBodyDataDocuments     `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	EmbeddingElapsedMs *int64                                         `json:"embeddingElapsedMs,omitempty" xml:"embeddingElapsedMs,omitempty"`
	TextChunkList      []*RecallDocumentResponseBodyDataTextChunkList `json:"textChunkList,omitempty" xml:"textChunkList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TextSearchElapsedMs *int64 `json:"textSearchElapsedMs,omitempty" xml:"textSearchElapsedMs,omitempty"`
	// example:
	//
	// 400
	TotalElapsedMs  *int64                                           `json:"totalElapsedMs,omitempty" xml:"totalElapsedMs,omitempty"`
	VectorChunkList []*RecallDocumentResponseBodyDataVectorChunkList `json:"vectorChunkList,omitempty" xml:"vectorChunkList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	VectorSearchElapsedMs *int64 `json:"vectorSearchElapsedMs,omitempty" xml:"vectorSearchElapsedMs,omitempty"`
}

func (s RecallDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyData) GetChunkList() []*RecallDocumentResponseBodyDataChunkList {
	return s.ChunkList
}

func (s *RecallDocumentResponseBodyData) GetChunkPartList() []*RecallDocumentResponseBodyDataChunkPartList {
	return s.ChunkPartList
}

func (s *RecallDocumentResponseBodyData) GetChunkTextList() []*string {
	return s.ChunkTextList
}

func (s *RecallDocumentResponseBodyData) GetDocuments() []*RecallDocumentResponseBodyDataDocuments {
	return s.Documents
}

func (s *RecallDocumentResponseBodyData) GetEmbeddingElapsedMs() *int64 {
	return s.EmbeddingElapsedMs
}

func (s *RecallDocumentResponseBodyData) GetTextChunkList() []*RecallDocumentResponseBodyDataTextChunkList {
	return s.TextChunkList
}

func (s *RecallDocumentResponseBodyData) GetTextSearchElapsedMs() *int64 {
	return s.TextSearchElapsedMs
}

func (s *RecallDocumentResponseBodyData) GetTotalElapsedMs() *int64 {
	return s.TotalElapsedMs
}

func (s *RecallDocumentResponseBodyData) GetVectorChunkList() []*RecallDocumentResponseBodyDataVectorChunkList {
	return s.VectorChunkList
}

func (s *RecallDocumentResponseBodyData) GetVectorSearchElapsedMs() *int64 {
	return s.VectorSearchElapsedMs
}

func (s *RecallDocumentResponseBodyData) SetChunkList(v []*RecallDocumentResponseBodyDataChunkList) *RecallDocumentResponseBodyData {
	s.ChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetChunkPartList(v []*RecallDocumentResponseBodyDataChunkPartList) *RecallDocumentResponseBodyData {
	s.ChunkPartList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetChunkTextList(v []*string) *RecallDocumentResponseBodyData {
	s.ChunkTextList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetDocuments(v []*RecallDocumentResponseBodyDataDocuments) *RecallDocumentResponseBodyData {
	s.Documents = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetEmbeddingElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.EmbeddingElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTextChunkList(v []*RecallDocumentResponseBodyDataTextChunkList) *RecallDocumentResponseBodyData {
	s.TextChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTextSearchElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.TextSearchElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTotalElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.TotalElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetVectorChunkList(v []*RecallDocumentResponseBodyDataVectorChunkList) *RecallDocumentResponseBodyData {
	s.VectorChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetVectorSearchElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.VectorSearchElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataChunkList struct {
	// example:
	//
	// 823746762354
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 839468263472
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// dscsbdsk
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 982374872364
	NextChunkId *string                                       `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 827364827364832
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

func (s RecallDocumentResponseBodyDataChunkList) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkList) GetChunkId() *string {
	return s.ChunkId
}

func (s *RecallDocumentResponseBodyDataChunkList) GetChunkMeta() map[string]interface{} {
	return s.ChunkMeta
}

func (s *RecallDocumentResponseBodyDataChunkList) GetChunkOssUrl() *string {
	return s.ChunkOssUrl
}

func (s *RecallDocumentResponseBodyDataChunkList) GetChunkText() *string {
	return s.ChunkText
}

func (s *RecallDocumentResponseBodyDataChunkList) GetChunkType() *string {
	return s.ChunkType
}

func (s *RecallDocumentResponseBodyDataChunkList) GetDocId() *string {
	return s.DocId
}

func (s *RecallDocumentResponseBodyDataChunkList) GetFileType() *string {
	return s.FileType
}

func (s *RecallDocumentResponseBodyDataChunkList) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RecallDocumentResponseBodyDataChunkList) GetLibraryName() *string {
	return s.LibraryName
}

func (s *RecallDocumentResponseBodyDataChunkList) GetNextChunkId() *string {
	return s.NextChunkId
}

func (s *RecallDocumentResponseBodyDataChunkList) GetPos() []*RecallDocumentResponseBodyDataChunkListPos {
	return s.Pos
}

func (s *RecallDocumentResponseBodyDataChunkList) GetPreChunkId() *string {
	return s.PreChunkId
}

func (s *RecallDocumentResponseBodyDataChunkList) GetScore() *float32 {
	return s.Score
}

func (s *RecallDocumentResponseBodyDataChunkList) GetTitle() *string {
	return s.Title
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetPos(v []*RecallDocumentResponseBodyDataChunkListPos) *RecallDocumentResponseBodyDataChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataChunkList {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataChunkListPos) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkListPos) GetAxisArray() []*float64 {
	return s.AxisArray
}

func (s *RecallDocumentResponseBodyDataChunkListPos) GetPage() *int32 {
	return s.Page
}

func (s *RecallDocumentResponseBodyDataChunkListPos) GetTextHighlightArea() []*int32 {
	return s.TextHighlightArea
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataChunkListPos {
	s.TextHighlightArea = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataChunkPartList struct {
	// example:
	//
	// 98327482364
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 92837482364
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// sjdhgjsd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 2387648263542
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataChunkPartListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 32874682764
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	Title *string  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataChunkPartList) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkPartList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetChunkId() *string {
	return s.ChunkId
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetChunkMeta() map[string]interface{} {
	return s.ChunkMeta
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetChunkOssUrl() *string {
	return s.ChunkOssUrl
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetChunkText() *string {
	return s.ChunkText
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetChunkType() *string {
	return s.ChunkType
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetDocId() *string {
	return s.DocId
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetFileType() *string {
	return s.FileType
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetLibraryName() *string {
	return s.LibraryName
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetNextChunkId() *string {
	return s.NextChunkId
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetPos() []*RecallDocumentResponseBodyDataChunkPartListPos {
	return s.Pos
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetPreChunkId() *string {
	return s.PreChunkId
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetScore() *float32 {
	return s.Score
}

func (s *RecallDocumentResponseBodyDataChunkPartList) GetTitle() *string {
	return s.Title
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkText(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkType(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetDocId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetFileType(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetLibraryId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetLibraryName(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetPos(v []*RecallDocumentResponseBodyDataChunkPartListPos) *RecallDocumentResponseBodyDataChunkPartList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetScore(v float32) *RecallDocumentResponseBodyDataChunkPartList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetTitle(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataChunkPartListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataChunkPartListPos) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkPartListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) GetAxisArray() []*float64 {
	return s.AxisArray
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) GetPage() *int32 {
	return s.Page
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) GetTextHighlightArea() []*int32 {
	return s.TextHighlightArea
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetPage(v int32) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.TextHighlightArea = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataDocuments struct {
	// example:
	//
	// 92837482364
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// {"a":"1"}
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
	// sjdhgjsd
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/test.pdf
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RecallDocumentResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataDocuments) GetDocId() *string {
	return s.DocId
}

func (s *RecallDocumentResponseBodyDataDocuments) GetDocumentMeta() map[string]interface{} {
	return s.DocumentMeta
}

func (s *RecallDocumentResponseBodyDataDocuments) GetFileType() *string {
	return s.FileType
}

func (s *RecallDocumentResponseBodyDataDocuments) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RecallDocumentResponseBodyDataDocuments) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RecallDocumentResponseBodyDataDocuments) GetTitle() *string {
	return s.Title
}

func (s *RecallDocumentResponseBodyDataDocuments) GetUrl() *string {
	return s.Url
}

func (s *RecallDocumentResponseBodyDataDocuments) SetDocId(v string) *RecallDocumentResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetDocumentMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataDocuments {
	s.DocumentMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetFileType(v string) *RecallDocumentResponseBodyDataDocuments {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetGmtCreate(v string) *RecallDocumentResponseBodyDataDocuments {
	s.GmtCreate = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetLibraryId(v string) *RecallDocumentResponseBodyDataDocuments {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetTitle(v string) *RecallDocumentResponseBodyDataDocuments {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetUrl(v string) *RecallDocumentResponseBodyDataDocuments {
	s.Url = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataTextChunkList struct {
	// example:
	//
	// 32874682364
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 8372467263542
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// djsgfsjd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 23874682432
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataTextChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 89473868346
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	Title *string  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataTextChunkList) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataTextChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetChunkId() *string {
	return s.ChunkId
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetChunkMeta() map[string]interface{} {
	return s.ChunkMeta
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetChunkOssUrl() *string {
	return s.ChunkOssUrl
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetChunkText() *string {
	return s.ChunkText
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetChunkType() *string {
	return s.ChunkType
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetDocId() *string {
	return s.DocId
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetFileType() *string {
	return s.FileType
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetLibraryName() *string {
	return s.LibraryName
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetNextChunkId() *string {
	return s.NextChunkId
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetPos() []*RecallDocumentResponseBodyDataTextChunkListPos {
	return s.Pos
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetPreChunkId() *string {
	return s.PreChunkId
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetScore() *float32 {
	return s.Score
}

func (s *RecallDocumentResponseBodyDataTextChunkList) GetTitle() *string {
	return s.Title
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetPos(v []*RecallDocumentResponseBodyDataTextChunkListPos) *RecallDocumentResponseBodyDataTextChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataTextChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataTextChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataTextChunkListPos) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataTextChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) GetAxisArray() []*float64 {
	return s.AxisArray
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) GetPage() *int32 {
	return s.Page
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) GetTextHighlightArea() []*int32 {
	return s.TextHighlightArea
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.TextHighlightArea = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataVectorChunkList struct {
	// example:
	//
	// 8723642345276
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// https://oss-xxxx-hangzhou.com/test.pdf
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 78326476235675372
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// djsgfsjd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 293846872343
	NextChunkId *string                                             `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataVectorChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 873647326542
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

func (s RecallDocumentResponseBodyDataVectorChunkList) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataVectorChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetChunkId() *string {
	return s.ChunkId
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetChunkMeta() map[string]interface{} {
	return s.ChunkMeta
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetChunkOssUrl() *string {
	return s.ChunkOssUrl
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetChunkText() *string {
	return s.ChunkText
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetChunkType() *string {
	return s.ChunkType
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetDocId() *string {
	return s.DocId
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetFileType() *string {
	return s.FileType
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetLibraryName() *string {
	return s.LibraryName
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetNextChunkId() *string {
	return s.NextChunkId
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetPos() []*RecallDocumentResponseBodyDataVectorChunkListPos {
	return s.Pos
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetPreChunkId() *string {
	return s.PreChunkId
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetScore() *float32 {
	return s.Score
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) GetTitle() *string {
	return s.Title
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetPos(v []*RecallDocumentResponseBodyDataVectorChunkListPos) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentResponseBodyDataVectorChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataVectorChunkListPos) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentResponseBodyDataVectorChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) GetAxisArray() []*float64 {
	return s.AxisArray
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) GetPage() *int32 {
	return s.Page
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) GetTextHighlightArea() []*int32 {
	return s.TextHighlightArea
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.TextHighlightArea = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) Validate() error {
	return dara.Validate(s)
}
