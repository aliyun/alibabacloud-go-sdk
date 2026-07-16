// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFileResponseBody
	GetCode() *string
	SetData(v *DescribeFileResponseBodyData) *DescribeFileResponseBody
	GetData() *DescribeFileResponseBodyData
	SetMessage(v string) *DescribeFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeFileResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeFileResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribeFileResponseBody
	GetSuccess() *bool
}

type DescribeFileResponseBody struct {
	// The error status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data field returned by the operation.
	Data *DescribeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the call was successful. Valid values:
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

func (s DescribeFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFileResponseBody) GetData() *DescribeFileResponseBodyData {
	return s.Data
}

func (s *DescribeFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFileResponseBody) SetCode(v string) *DescribeFileResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFileResponseBody) SetData(v *DescribeFileResponseBodyData) *DescribeFileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFileResponseBody) SetMessage(v string) *DescribeFileResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFileResponseBody) SetRequestId(v string) *DescribeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileResponseBody) SetStatus(v string) *DescribeFileResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFileResponseBody) SetSuccess(v bool) *DescribeFileResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileResponseBodyData struct {
	// The ID of the category to which the file belongs.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The timestamp when the file was added to Model Studio. Format: yyyy-MM-dd HH:mm:ss. Time zone: UTC+8.
	//
	// example:
	//
	// 2024-09-09 12:45:43
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The file ID.
	//
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_xxxxxxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file name.
	//
	// example:
	//
	// XXX产品介绍.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type (extension). Possible values: pdf, docx, doc, txt, md, pptx, ppt, xlsx, xls, html, png, jpg, jpeg, bmp, and gif.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The reason for the parsing failure.
	//
	// example:
	//
	// Error Message
	ParseErrorMessage      *string `json:"ParseErrorMessage,omitempty" xml:"ParseErrorMessage,omitempty"`
	ParseResultDownloadUrl *string `json:"ParseResultDownloadUrl,omitempty" xml:"ParseResultDownloadUrl,omitempty"`
	// The parser type used to parse the file. Possible values:
	//
	// - DASHSCOPE_DOCMIND: the default document parser.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 1234
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// <props="china">
	//
	// For files used in document-based knowledge bases (type: UNSTRUCTURED), possible values:
	//
	//
	//
	// <props="intl">
	//
	// For files used in unstructured knowledge bases (type: UNSTRUCTURED), possible values:
	//
	//
	//
	//
	// - INIT: Pending parsing.
	//
	// - IN_PARSE_QUEUE: Queued for parsing.
	//
	// - PARSING: Being parsed.
	//
	// - PARSE_SUCCESS: Parsing completed.
	//
	// <note>The document can be imported into a knowledge base only after the status changes to PARSE_SUCCESS.</note>
	//
	// - PARSE_FAILED: Parsing failed.
	//
	// <props="china">
	//
	// For files used in agent application [session interaction](https://www.alibabacloud.com/help/en/model-studio/user-guide/file-interaction) (type: SESSION_FILE), possible values:
	//
	// - INIT: Pending parsing.
	//
	// - IN_PARSE_QUEUE: Queued for parsing.
	//
	// - PARSING: Being parsed.
	//
	// - PARSE_SUCCESS: Parsing completed.
	//
	// - PARSE_FAILED: Parsing failed.
	//
	// - SAFE_CHECKING: Safety check in progress.
	//
	// - SAFE_CHECK_FAILED: Safety check failed.
	//
	// - INDEX_BUILDING: Index being built.
	//
	// - INDEX_BUILD_SUCCESS: Index built.
	//
	// - INDEX_BUILDING_FAILED: Index building failed.
	//
	// - INDEX_DELETED: File index deleted.
	//
	// - FILE_IS_READY: File is ready.
	//
	// <note>Q&A is available only after the status changes to FILE_IS_READY.</note>
	//
	// - FILE_EXPIRED: File expired.
	//
	// <note>The file is valid only for the current user session. After the user closes the session, the file expires (maximum validity period: 7 days). Long-term retention is not supported.</note>
	//
	// .
	//
	// example:
	//
	// PARSE_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags associated with the file. A file can be associated with multiple tags.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeFileResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFileResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *DescribeFileResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *DescribeFileResponseBodyData) GetFileType() *string {
	return s.FileType
}

func (s *DescribeFileResponseBodyData) GetParseErrorMessage() *string {
	return s.ParseErrorMessage
}

func (s *DescribeFileResponseBodyData) GetParseResultDownloadUrl() *string {
	return s.ParseResultDownloadUrl
}

func (s *DescribeFileResponseBodyData) GetParser() *string {
	return s.Parser
}

func (s *DescribeFileResponseBodyData) GetSizeInBytes() *int64 {
	return s.SizeInBytes
}

func (s *DescribeFileResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileResponseBodyData) GetTags() []*string {
	return s.Tags
}

func (s *DescribeFileResponseBodyData) SetCategoryId(v string) *DescribeFileResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetCreateTime(v string) *DescribeFileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileId(v string) *DescribeFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileName(v string) *DescribeFileResponseBodyData {
	s.FileName = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileType(v string) *DescribeFileResponseBodyData {
	s.FileType = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetParseErrorMessage(v string) *DescribeFileResponseBodyData {
	s.ParseErrorMessage = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetParseResultDownloadUrl(v string) *DescribeFileResponseBodyData {
	s.ParseResultDownloadUrl = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetParser(v string) *DescribeFileResponseBodyData {
	s.Parser = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetSizeInBytes(v int64) *DescribeFileResponseBodyData {
	s.SizeInBytes = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetStatus(v string) *DescribeFileResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetTags(v []*string) *DescribeFileResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
