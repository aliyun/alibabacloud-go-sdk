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
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data field returned by the operation.
	Data *ListFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F97A63B-xxxx-527F-9D6E-467B6A7E8CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFileResponseBodyData struct {
	// The list of files in the category.
	FileList []*ListFileResponseBodyDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// Indicates whether there is a next page of category data that matches the query conditions. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// The number of entries per page for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// 4jzbJk9J6lNeuXD9hP0viA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries in the returned results.
	//
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
	if s.FileList != nil {
		for _, item := range s.FileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFileResponseBodyDataFileList struct {
	// The ID of the category to which the file belongs.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The timestamp when the file was added to Alibaba Cloud Model Studio. Format: yyyy-MM-dd HH:mm:ss. Time zone: UTC+8.
	//
	// example:
	//
	// 2024-09-09 11:03:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The file ID, which is the `FileId` returned by the **AddFile*	- operation. You can also obtain it on the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) page by clicking the icon next to the file name.
	//
	// example:
	//
	// file_5ff599b3455a45db8c41b0054b361518_xxxxxxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file name.
	//
	// example:
	//
	// product-overview.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file format type. Valid values: pdf, docx, doc, txt, md, pptx, ppt, xlsx, xls, html, png, jpg, jpeg, bmp, and gif.
	//
	// example:
	//
	// docx
	FileType          *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	ParseErrorMessage *string `json:"ParseErrorMessage,omitempty" xml:"ParseErrorMessage,omitempty"`
	// The document parser. Valid values:
	//
	// - DASHSCOPE_DOCMIND: Alibaba Cloud intelligent document parsing.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The file size in bytes.
	//
	// example:
	//
	// 512
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// The file parsing status. Valid values:
	//
	// - INIT: Initialization state, waiting to be scheduled.
	//
	// - PARSING: Parsing in progress.
	//
	// - PARSE_SUCCESS: Parsing completed.
	//
	// - PARSE_FAILED: Parsing failed.
	//
	// example:
	//
	// PARSE_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags associated with the file. A document can be associated with multiple tags.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListFileResponseBodyDataFileList) GetParseErrorMessage() *string {
	return s.ParseErrorMessage
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

func (s *ListFileResponseBodyDataFileList) SetParseErrorMessage(v string) *ListFileResponseBodyDataFileList {
	s.ParseErrorMessage = &v
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
