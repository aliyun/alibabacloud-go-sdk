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
  // **If there are no special requirements, leave it blank.**
  // 
  // Chain authorization configuration, optional. For more information, see [Using Chain Authorization to Access Other Entity Resources](https://help.aliyun.com/document_detail/465340.html).
  CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
  // Project name. For how to obtain it, see [Creating a Project](https://help.aliyun.com/document_detail/477051.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // immtest
  ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
  // Suffix type of the source data. By default, the type of the source data is determined based on the suffix of the input object. When the input object does not have a suffix, you can set this parameter. The available values are as follows:
  // 
  // - Word Documents: doc, docx, wps, wpss, docm, dotm, dot, dotx, html
  // 
  // - Presentation Documents (PPT): pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, dpss
  // 
  // - Spreadsheet Documents (Excel): xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, ets
  // 
  // - PDF Documents: pdf
  // 
  // example:
  // 
  // docx
  SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
  // Storage address of the source data.
  // 
  // The OSS address rule is oss://${Bucket}/${Object}, where `${Bucket}` is the name of the OSS Bucket in the same region (Region) as the current project, and `${Object}` is the complete path of the file including the file extension.
  // 
  // 	Notice: Currently, only HTTP protocol addresses are supported.
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

