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
	// The error code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data field.
	Data *DescribeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
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
	// The ID of the file.
	//
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_xxxxxxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// XXX Product Introduction.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type, which is the file extension. Valid values: pdf, docx, doc, txt, md, pptx, ppt, xlsx, xls, html, png, jpg, jpeg, bmp, and gif.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The download URL of the source file.
	//
	// example:
	//
	// https:xxxx.com/xxx.pdf?sign=xxxx&expiredAt=xxxx
	OriginFileDownloadUrl *string `json:"OriginFileDownloadUrl,omitempty" xml:"OriginFileDownloadUrl,omitempty"`
	// The cause of the parsing failure.
	//
	// example:
	//
	// Error Message
	ParseErrorMessage *string `json:"ParseErrorMessage,omitempty" xml:"ParseErrorMessage,omitempty"`
	// The download URL of the file parsing result.
	//
	// example:
	//
	// https://example.com/parse-result
	ParseResultDownloadUrl *string `json:"ParseResultDownloadUrl,omitempty" xml:"ParseResultDownloadUrl,omitempty"`
	// The type of the parser used to parse the file. Valid values:
	//
	// - DASHSCOPE_DOCMIND: the default document parser.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The size of the file, in bytes.
	//
	// example:
	//
	// 1234
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// <props="china">
	//
	// The status of a file used for a document knowledge base (type: UNSTRUCTURED). Valid values:
	//
	//
	//
	// <props="intl">
	//
	// The status of a file used for an unstructured knowledge base (type: UNSTRUCTURED). Valid values:
	//
	//
	//
	//
	// - INIT: to be parsed.
	//
	// - IN_PARSE_QUEUE: queuing for parsing.
	//
	// - PARSING: parsing.
	//
	// - PARSE_SUCCESS: parsed.
	//
	// <note>You can import a document into a knowledge base only after its status changes to PARSE_SUCCESS.</note>
	//
	// - PARSE_FAILED: parsing failed.
	//
	// <props="china">
	//
	// The status of a file used for [session interaction](https://www.alibabacloud.com/help/en/model-studio/user-guide/file-interaction) in an agent application (type: SESSION_FILE). Valid values:
	//
	// - INIT: to be parsed.
	//
	// - IN_PARSE_QUEUE: queuing for parsing.
	//
	// - PARSING: parsing.
	//
	// - PARSE_SUCCESS: parsed.
	//
	// - PARSE_FAILED: parsing failed.
	//
	// - SAFE_CHECKING: undergoing security check.
	//
	// - SAFE_CHECK_FAILED: security check failed.
	//
	// - INDEX_BUILDING: building index.
	//
	// - INDEX_BUILD_SUCCESS: index built.
	//
	// - INDEX_BUILDING_FAILED: index building failed.
	//
	// - INDEX_DELETED: file index deleted.
	//
	// - FILE_IS_READY: file ready.
	//
	// <note>You can perform Q&A only after the status changes to FILE_IS_READY.</note>
	//
	// - FILE_EXPIRED: file expired.
	//
	// <note>The file is valid only for the current session of the user. After the user closes the session, the file expires. The maximum validity period is 7 days. Long-term storage is not supported.</note>
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

func (s *DescribeFileResponseBodyData) GetOriginFileDownloadUrl() *string {
	return s.OriginFileDownloadUrl
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

func (s *DescribeFileResponseBodyData) SetOriginFileDownloadUrl(v string) *DescribeFileResponseBodyData {
	s.OriginFileDownloadUrl = &v
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
