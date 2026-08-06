// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateDocParserJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileFormat(v string) *CreateDocParserJobAdvanceRequest
	GetFileFormat() *string
	SetFileName(v string) *CreateDocParserJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetImageMode(v string) *CreateDocParserJobAdvanceRequest
	GetImageMode() *string
	SetOssFileUrl(v string) *CreateDocParserJobAdvanceRequest
	GetOssFileUrl() *string
	SetOutputFormat(v string) *CreateDocParserJobAdvanceRequest
	GetOutputFormat() *string
	SetRegionId(v string) *CreateDocParserJobAdvanceRequest
	GetRegionId() *string
	SetResultType(v string) *CreateDocParserJobAdvanceRequest
	GetResultType() *string
	SetTableFormat(v string) *CreateDocParserJobAdvanceRequest
	GetTableFormat() *string
}

type CreateDocParserJobAdvanceRequest struct {
	// The format of the input file. Valid values:
	//
	// - **pdf**: PDF file.
	//
	// - **docx**: Word file in docx format.
	//
	// - **doc**: Word file in doc format.
	//
	// - **pptx**: PPT file in pptx format.
	//
	// - **ppt**: PPT file in ppt format.
	//
	// - **txt**: plain text file.
	//
	// - **md**: Markdown file.
	//
	// - **png**: PNG image.
	//
	// - **jpg**: JPG image.
	//
	// - **jpeg**: JPEG image.
	//
	// This parameter is required.
	//
	// example:
	//
	// pdf
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The file name, which must include the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// document.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The HTTP or HTTPS URL of the file to be parsed.
	//
	// >SDKs for various languages provide an additional `CreateDocParserJobAdvance` method that supports passing a local file stream directly (such as InputStream in Java), without the need to upload the file to OSS and construct a FileUrl in advance. When using the Advance method, replace the `FileUrl` parameter (URL string) with the `FileUrlObject` parameter (file stream). All other request parameters remain unchanged. The SDK automatically performs the following operations:
	//
	// >1. Obtains temporary OSS upload credentials.
	//
	// >2. Uploads the file stream directly to OSS.
	//
	// >3. Calls the CreateDocParserJob operation with the generated OSS URL.
	//
	// example:
	//
	// https://xxx.oss-cn-beijing.aliyuncs.com/document.pdf?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImageMode     *string   `json:"ImageMode,omitempty" xml:"ImageMode,omitempty"`
	// The OSS file URL.
	OssFileUrl *string `json:"OssFileUrl,omitempty" xml:"OssFileUrl,omitempty"`
	// The output format of the parsing result. Valid values:
	//
	// - **markdown**: Markdown format.
	//
	// This parameter is required.
	//
	// example:
	//
	// markdown
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResultType  *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	TableFormat *string `json:"TableFormat,omitempty" xml:"TableFormat,omitempty"`
}

func (s CreateDocParserJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobAdvanceRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateDocParserJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *CreateDocParserJobAdvanceRequest) GetImageMode() *string {
	return s.ImageMode
}

func (s *CreateDocParserJobAdvanceRequest) GetOssFileUrl() *string {
	return s.OssFileUrl
}

func (s *CreateDocParserJobAdvanceRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateDocParserJobAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobAdvanceRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobAdvanceRequest) GetTableFormat() *string {
	return s.TableFormat
}

func (s *CreateDocParserJobAdvanceRequest) SetFileFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileName(v string) *CreateDocParserJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetImageMode(v string) *CreateDocParserJobAdvanceRequest {
	s.ImageMode = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetOssFileUrl(v string) *CreateDocParserJobAdvanceRequest {
	s.OssFileUrl = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetOutputFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetRegionId(v string) *CreateDocParserJobAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetResultType(v string) *CreateDocParserJobAdvanceRequest {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetTableFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.TableFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
