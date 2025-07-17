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
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data fields.
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	return dara.Validate(s)
}

type DescribeFileResponseBodyData struct {
	// The ID of the category to which the document belongs.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The timestamp when the document was uploaded to Model Studio. Format: yyyy-MM-dd HH:mm:ss. Time zone: UTC + 8.
	//
	// example:
	//
	// 2024-05-26 12:45:43
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The primary key ID of the document.
	//
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type of the document. The value is an extension. Valid values: pdf, docx, doc, txt, md, pptx, and ppt.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The parser that is used to parse the document. Valid value:
	//
	// 	- DASHSCOPE_DOCMIND: The default document parser.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The size of the document. Unit: bytes.
	//
	// example:
	//
	// 1234
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// The status of the document. Valid values:
	//
	// 	- INIT: pending parsing.
	//
	// 	- PARSING
	//
	// 	- PARSE_SUCCESS
	//
	// 	- PARSE_FAILED
	//
	// example:
	//
	// PARSE_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are associated with the document. A document can be associated with multiple tags.
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
