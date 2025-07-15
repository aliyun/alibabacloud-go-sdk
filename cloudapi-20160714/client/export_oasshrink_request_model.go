// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOASShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApiIdListShrink(v string) *ExportOASShrinkRequest
  GetApiIdListShrink() *string 
  SetDataFormat(v string) *ExportOASShrinkRequest
  GetDataFormat() *string 
  SetGroupId(v string) *ExportOASShrinkRequest
  GetGroupId() *string 
  SetOasVersion(v string) *ExportOASShrinkRequest
  GetOasVersion() *string 
  SetPageNumber(v int32) *ExportOASShrinkRequest
  GetPageNumber() *int32 
  SetSecurityToken(v string) *ExportOASShrinkRequest
  GetSecurityToken() *string 
  SetStageName(v string) *ExportOASShrinkRequest
  GetStageName() *string 
  SetWithXExtensions(v bool) *ExportOASShrinkRequest
  GetWithXExtensions() *bool 
}

type ExportOASShrinkRequest struct {
  // The APIs that you want to export.
  ApiIdListShrink *string `json:"ApiIdList,omitempty" xml:"ApiIdList,omitempty"`
  // The exported format:
  // 
  // 	- json
  // 
  // 	- yaml
  // 
  // example:
  // 
  // yaml
  DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
  // The API group ID.
  // 
  // example:
  // 
  // 42925e7f5209438186d5560239af5xxx
  GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // The OAS version. Valid values:
  // 
  // 	- **oas2**
  // 
  // 	- **oas3**
  // 
  // example:
  // 
  // oas2
  OasVersion *string `json:"OasVersion,omitempty" xml:"OasVersion,omitempty"`
  // The number of pages in which you want to export the APIs.
  // 
  // example:
  // 
  // 1
  PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
  SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
  // The environment to which the API is published. Valid values:
  // 
  // 	- **RELEASE**: the production environment
  // 
  // 	- **PRE**: the pre-release environment
  // 
  // 	- **TEST**: the test environment
  // 
  // example:
  // 
  // RELEASE
  StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
  // Specifies whether to export API Gateway extensions at the same time.
  // 
  // example:
  // 
  // true
  WithXExtensions *bool `json:"WithXExtensions,omitempty" xml:"WithXExtensions,omitempty"`
}

func (s ExportOASShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportOASShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportOASShrinkRequest) GetApiIdListShrink() *string  {
  return s.ApiIdListShrink
}

func (s *ExportOASShrinkRequest) GetDataFormat() *string  {
  return s.DataFormat
}

func (s *ExportOASShrinkRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *ExportOASShrinkRequest) GetOasVersion() *string  {
  return s.OasVersion
}

func (s *ExportOASShrinkRequest) GetPageNumber() *int32  {
  return s.PageNumber
}

func (s *ExportOASShrinkRequest) GetSecurityToken() *string  {
  return s.SecurityToken
}

func (s *ExportOASShrinkRequest) GetStageName() *string  {
  return s.StageName
}

func (s *ExportOASShrinkRequest) GetWithXExtensions() *bool  {
  return s.WithXExtensions
}

func (s *ExportOASShrinkRequest) SetApiIdListShrink(v string) *ExportOASShrinkRequest {
  s.ApiIdListShrink = &v
  return s
}

func (s *ExportOASShrinkRequest) SetDataFormat(v string) *ExportOASShrinkRequest {
  s.DataFormat = &v
  return s
}

func (s *ExportOASShrinkRequest) SetGroupId(v string) *ExportOASShrinkRequest {
  s.GroupId = &v
  return s
}

func (s *ExportOASShrinkRequest) SetOasVersion(v string) *ExportOASShrinkRequest {
  s.OasVersion = &v
  return s
}

func (s *ExportOASShrinkRequest) SetPageNumber(v int32) *ExportOASShrinkRequest {
  s.PageNumber = &v
  return s
}

func (s *ExportOASShrinkRequest) SetSecurityToken(v string) *ExportOASShrinkRequest {
  s.SecurityToken = &v
  return s
}

func (s *ExportOASShrinkRequest) SetStageName(v string) *ExportOASShrinkRequest {
  s.StageName = &v
  return s
}

func (s *ExportOASShrinkRequest) SetWithXExtensions(v bool) *ExportOASShrinkRequest {
  s.WithXExtensions = &v
  return s
}

func (s *ExportOASShrinkRequest) Validate() error {
  return dara.Validate(s)
}

