// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFileResponseBody
	GetCode() *string
	SetData(v *ListFileResponseBodyData) *ListFileResponseBody
	GetData() *ListFileResponseBodyData
	SetMessage(v string) *ListFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFileResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListFileResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListFileResponseBody
	GetSuccess() *bool
}

type ListFileResponseBody struct {
	// example:
	//
	// success
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8F97A63B-55F1-527F-9D6E-467B6A7E8CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFileResponseBody) GetData() *ListFileResponseBodyData {
	return s.Data
}

func (s *ListFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFileResponseBody) SetCode(v string) *ListFileResponseBody {
	s.Code = &v
	return s
}

func (s *ListFileResponseBody) SetData(v *ListFileResponseBodyData) *ListFileResponseBody {
	s.Data = v
	return s
}

func (s *ListFileResponseBody) SetMessage(v string) *ListFileResponseBody {
	s.Message = &v
	return s
}

func (s *ListFileResponseBody) SetRequestId(v string) *ListFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileResponseBody) SetStatus(v string) *ListFileResponseBody {
	s.Status = &v
	return s
}

func (s *ListFileResponseBody) SetSuccess(v bool) *ListFileResponseBody {
	s.Success = &v
	return s
}

func (s *ListFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFileResponseBodyData struct {
	FileList []*ListFileResponseBodyDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4jzbJk9J6lNeuXD9hP0viA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 48
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyData) GetFileList() []*ListFileResponseBodyDataFileList {
	return s.FileList
}

func (s *ListFileResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListFileResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFileResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFileResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileResponseBodyData) SetFileList(v []*ListFileResponseBodyDataFileList) *ListFileResponseBodyData {
	s.FileList = v
	return s
}

func (s *ListFileResponseBodyData) SetHasNext(v bool) *ListFileResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *ListFileResponseBodyData) SetMaxResults(v int32) *ListFileResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListFileResponseBodyData) SetNextToken(v string) *ListFileResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListFileResponseBodyData) SetTotalCount(v int32) *ListFileResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFileResponseBodyDataFileList struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2023-08-18 11:03:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// file_5ff599b3455a45db8c41b0054b361518_10098576
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// auto-test-1721096109278.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// example:
	//
	// 512
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// example:
	//
	// 200
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFileResponseBodyDataFileList) String() string {
	return dara.Prettify(s)
}

func (s ListFileResponseBodyDataFileList) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyDataFileList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListFileResponseBodyDataFileList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFileResponseBodyDataFileList) GetFileId() *string {
	return s.FileId
}

func (s *ListFileResponseBodyDataFileList) GetFileName() *string {
	return s.FileName
}

func (s *ListFileResponseBodyDataFileList) GetFileType() *string {
	return s.FileType
}

func (s *ListFileResponseBodyDataFileList) GetParser() *string {
	return s.Parser
}

func (s *ListFileResponseBodyDataFileList) GetSizeInBytes() *int64 {
	return s.SizeInBytes
}

func (s *ListFileResponseBodyDataFileList) GetStatus() *string {
	return s.Status
}

func (s *ListFileResponseBodyDataFileList) GetTags() []*string {
	return s.Tags
}

func (s *ListFileResponseBodyDataFileList) SetCategoryId(v string) *ListFileResponseBodyDataFileList {
	s.CategoryId = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetCreateTime(v string) *ListFileResponseBodyDataFileList {
	s.CreateTime = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetFileId(v string) *ListFileResponseBodyDataFileList {
	s.FileId = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetFileName(v string) *ListFileResponseBodyDataFileList {
	s.FileName = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetFileType(v string) *ListFileResponseBodyDataFileList {
	s.FileType = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetParser(v string) *ListFileResponseBodyDataFileList {
	s.Parser = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetSizeInBytes(v int64) *ListFileResponseBodyDataFileList {
	s.SizeInBytes = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetStatus(v string) *ListFileResponseBodyDataFileList {
	s.Status = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetTags(v []*string) *ListFileResponseBodyDataFileList {
	s.Tags = v
	return s
}

func (s *ListFileResponseBodyDataFileList) Validate() error {
	return dara.Validate(s)
}
