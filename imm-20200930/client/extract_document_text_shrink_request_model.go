// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractDocumentTextShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCredentialConfigShrink(v string) *ExtractDocumentTextShrinkRequest
  GetCredentialConfigShrink() *string 
  SetProjectName(v string) *ExtractDocumentTextShrinkRequest
  GetProjectName() *string 
  SetSourceType(v string) *ExtractDocumentTextShrinkRequest
  GetSourceType() *string 
  SetSourceURI(v string) *ExtractDocumentTextShrinkRequest
  GetSourceURI() *string 
}

type ExtractDocumentTextShrinkRequest struct {
  // **If you have no special requirements, leave this parameter empty.**
  // 
  // The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
  CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
  // The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/477051.html) operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // immtest
  ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
  // The type of the filename extension of the source data. By default, the filename extension of the source data is the same as the filename extension of the input document. If the input document has no extension, you can specify this parameter. Valid values:
  // 
  // 	- Text documents: doc, docx, wps, wpss, docm, dotm, dot, dotx, and html
  // 
  // 	- Presentation documents: pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, and dpss
  // 
  // 	- Table documents: xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, and ets
  // 
  // 	- PDF documents: pdf.
  // 
  // example:
  // 
  // docx
  SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
  // The URI of the Object Storage Service (OSS) bucket in which the document is stored.
  // 
  // Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the file that has an extension.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // oss://test-bucket/test-object
  SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s ExtractDocumentTextShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtractDocumentTextShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExtractDocumentTextShrinkRequest) GetCredentialConfigShrink() *string  {
  return s.CredentialConfigShrink
}

func (s *ExtractDocumentTextShrinkRequest) GetProjectName() *string  {
  return s.ProjectName
}

func (s *ExtractDocumentTextShrinkRequest) GetSourceType() *string  {
  return s.SourceType
}

func (s *ExtractDocumentTextShrinkRequest) GetSourceURI() *string  {
  return s.SourceURI
}

func (s *ExtractDocumentTextShrinkRequest) SetCredentialConfigShrink(v string) *ExtractDocumentTextShrinkRequest {
  s.CredentialConfigShrink = &v
  return s
}

func (s *ExtractDocumentTextShrinkRequest) SetProjectName(v string) *ExtractDocumentTextShrinkRequest {
  s.ProjectName = &v
  return s
}

func (s *ExtractDocumentTextShrinkRequest) SetSourceType(v string) *ExtractDocumentTextShrinkRequest {
  s.SourceType = &v
  return s
}

func (s *ExtractDocumentTextShrinkRequest) SetSourceURI(v string) *ExtractDocumentTextShrinkRequest {
  s.SourceURI = &v
  return s
}

func (s *ExtractDocumentTextShrinkRequest) Validate() error {
  return dara.Validate(s)
}

