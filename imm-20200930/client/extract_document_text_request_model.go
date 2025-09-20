// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractDocumentTextRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCredentialConfig(v *CredentialConfig) *ExtractDocumentTextRequest
  GetCredentialConfig() *CredentialConfig 
  SetProjectName(v string) *ExtractDocumentTextRequest
  GetProjectName() *string 
  SetSourceType(v string) *ExtractDocumentTextRequest
  GetSourceType() *string 
  SetSourceURI(v string) *ExtractDocumentTextRequest
  GetSourceURI() *string 
}

type ExtractDocumentTextRequest struct {
  // **If you have no special requirements, leave this parameter empty.**
  // 
  // The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
  CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s ExtractDocumentTextRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtractDocumentTextRequest) GoString() string {
  return s.String()
}

func (s *ExtractDocumentTextRequest) GetCredentialConfig() *CredentialConfig  {
  return s.CredentialConfig
}

func (s *ExtractDocumentTextRequest) GetProjectName() *string  {
  return s.ProjectName
}

func (s *ExtractDocumentTextRequest) GetSourceType() *string  {
  return s.SourceType
}

func (s *ExtractDocumentTextRequest) GetSourceURI() *string  {
  return s.SourceURI
}

func (s *ExtractDocumentTextRequest) SetCredentialConfig(v *CredentialConfig) *ExtractDocumentTextRequest {
  s.CredentialConfig = v
  return s
}

func (s *ExtractDocumentTextRequest) SetProjectName(v string) *ExtractDocumentTextRequest {
  s.ProjectName = &v
  return s
}

func (s *ExtractDocumentTextRequest) SetSourceType(v string) *ExtractDocumentTextRequest {
  s.SourceType = &v
  return s
}

func (s *ExtractDocumentTextRequest) SetSourceURI(v string) *ExtractDocumentTextRequest {
  s.SourceURI = &v
  return s
}

func (s *ExtractDocumentTextRequest) Validate() error {
  return dara.Validate(s)
}

