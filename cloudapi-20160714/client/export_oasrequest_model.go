// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOASRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApiIdList(v []*string) *ExportOASRequest
  GetApiIdList() []*string 
  SetDataFormat(v string) *ExportOASRequest
  GetDataFormat() *string 
  SetGroupId(v string) *ExportOASRequest
  GetGroupId() *string 
  SetOasVersion(v string) *ExportOASRequest
  GetOasVersion() *string 
  SetPageNumber(v int32) *ExportOASRequest
  GetPageNumber() *int32 
  SetSecurityToken(v string) *ExportOASRequest
  GetSecurityToken() *string 
  SetStageName(v string) *ExportOASRequest
  GetStageName() *string 
  SetWithXExtensions(v bool) *ExportOASRequest
  GetWithXExtensions() *bool 
}

type ExportOASRequest struct {
  // The APIs that you want to export.
  ApiIdList []*string `json:"ApiIdList,omitempty" xml:"ApiIdList,omitempty" type:"Repeated"`
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

func (s ExportOASRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportOASRequest) GoString() string {
  return s.String()
}

func (s *ExportOASRequest) GetApiIdList() []*string  {
  return s.ApiIdList
}

func (s *ExportOASRequest) GetDataFormat() *string  {
  return s.DataFormat
}

func (s *ExportOASRequest) GetGroupId() *string  {
  return s.GroupId
}

func (s *ExportOASRequest) GetOasVersion() *string  {
  return s.OasVersion
}

func (s *ExportOASRequest) GetPageNumber() *int32  {
  return s.PageNumber
}

func (s *ExportOASRequest) GetSecurityToken() *string  {
  return s.SecurityToken
}

func (s *ExportOASRequest) GetStageName() *string  {
  return s.StageName
}

func (s *ExportOASRequest) GetWithXExtensions() *bool  {
  return s.WithXExtensions
}

func (s *ExportOASRequest) SetApiIdList(v []*string) *ExportOASRequest {
  s.ApiIdList = v
  return s
}

func (s *ExportOASRequest) SetDataFormat(v string) *ExportOASRequest {
  s.DataFormat = &v
  return s
}

func (s *ExportOASRequest) SetGroupId(v string) *ExportOASRequest {
  s.GroupId = &v
  return s
}

func (s *ExportOASRequest) SetOasVersion(v string) *ExportOASRequest {
  s.OasVersion = &v
  return s
}

func (s *ExportOASRequest) SetPageNumber(v int32) *ExportOASRequest {
  s.PageNumber = &v
  return s
}

func (s *ExportOASRequest) SetSecurityToken(v string) *ExportOASRequest {
  s.SecurityToken = &v
  return s
}

func (s *ExportOASRequest) SetStageName(v string) *ExportOASRequest {
  s.StageName = &v
  return s
}

func (s *ExportOASRequest) SetWithXExtensions(v bool) *ExportOASRequest {
  s.WithXExtensions = &v
  return s
}

func (s *ExportOASRequest) Validate() error {
  return dara.Validate(s)
}

