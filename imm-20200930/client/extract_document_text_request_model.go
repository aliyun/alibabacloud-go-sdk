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
  // **If there are no special requirements, leave it blank.**
  // 
  // Chain authorization configuration, optional. For more information, see [Using Chain Authorization to Access Other Entity Resources](https://help.aliyun.com/document_detail/465340.html).
  CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
  if s.CredentialConfig != nil {
    if err := s.CredentialConfig.Validate(); err != nil {
      return err
    }
  }
  return nil
}

