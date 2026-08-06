// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileFormat(v string) *CreateDocParserJobRequest
	GetFileFormat() *string
	SetFileName(v string) *CreateDocParserJobRequest
	GetFileName() *string
	SetFileUrl(v string) *CreateDocParserJobRequest
	GetFileUrl() *string
	SetImageMode(v string) *CreateDocParserJobRequest
	GetImageMode() *string
	SetOssFileUrl(v string) *CreateDocParserJobRequest
	GetOssFileUrl() *string
	SetOutputFormat(v string) *CreateDocParserJobRequest
	GetOutputFormat() *string
	SetRegionId(v string) *CreateDocParserJobRequest
	GetRegionId() *string
	SetResultType(v string) *CreateDocParserJobRequest
	GetResultType() *string
	SetTableFormat(v string) *CreateDocParserJobRequest
	GetTableFormat() *string
}

type CreateDocParserJobRequest struct {
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
	FileUrl   *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImageMode *string `json:"ImageMode,omitempty" xml:"ImageMode,omitempty"`
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

func (s CreateDocParserJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateDocParserJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateDocParserJobRequest) GetImageMode() *string {
	return s.ImageMode
}

func (s *CreateDocParserJobRequest) GetOssFileUrl() *string {
	return s.OssFileUrl
}

func (s *CreateDocParserJobRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateDocParserJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobRequest) GetTableFormat() *string {
	return s.TableFormat
}

func (s *CreateDocParserJobRequest) SetFileFormat(v string) *CreateDocParserJobRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileName(v string) *CreateDocParserJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileUrl(v string) *CreateDocParserJobRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateDocParserJobRequest) SetImageMode(v string) *CreateDocParserJobRequest {
	s.ImageMode = &v
	return s
}

func (s *CreateDocParserJobRequest) SetOssFileUrl(v string) *CreateDocParserJobRequest {
	s.OssFileUrl = &v
	return s
}

func (s *CreateDocParserJobRequest) SetOutputFormat(v string) *CreateDocParserJobRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) SetRegionId(v string) *CreateDocParserJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobRequest) SetResultType(v string) *CreateDocParserJobRequest {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobRequest) SetTableFormat(v string) *CreateDocParserJobRequest {
	s.TableFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) Validate() error {
	return dara.Validate(s)
}
