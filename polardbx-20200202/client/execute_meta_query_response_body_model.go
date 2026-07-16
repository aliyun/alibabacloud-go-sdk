// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMetaQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v *ExecuteMetaQueryResponseBodyAccessDeniedDetail) *ExecuteMetaQueryResponseBody
  GetAccessDeniedDetail() *ExecuteMetaQueryResponseBodyAccessDeniedDetail 
  SetData(v *ExecuteMetaQueryResponseBodyData) *ExecuteMetaQueryResponseBody
  GetData() *ExecuteMetaQueryResponseBodyData 
  SetMessage(v string) *ExecuteMetaQueryResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteMetaQueryResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteMetaQueryResponseBody
  GetSuccess() *bool 
}

type ExecuteMetaQueryResponseBody struct {
  // The details about the access denial.
  AccessDeniedDetail *ExecuteMetaQueryResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
  // The instance details.
  Data *ExecuteMetaQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The additional information returned. If the request is successful, success is returned. If the request fails, the corresponding error code is returned.
  // 
  // example:
  // 
  // ""
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // A501A191-BD70-5E50-98A9-C2A486A82****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteMetaQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMetaQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteMetaQueryResponseBody) GetAccessDeniedDetail() *ExecuteMetaQueryResponseBodyAccessDeniedDetail  {
  return s.AccessDeniedDetail
}

func (s *ExecuteMetaQueryResponseBody) GetData() *ExecuteMetaQueryResponseBodyData  {
  return s.Data
}

func (s *ExecuteMetaQueryResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteMetaQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteMetaQueryResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteMetaQueryResponseBody) SetAccessDeniedDetail(v *ExecuteMetaQueryResponseBodyAccessDeniedDetail) *ExecuteMetaQueryResponseBody {
  s.AccessDeniedDetail = v
  return s
}

func (s *ExecuteMetaQueryResponseBody) SetData(v *ExecuteMetaQueryResponseBodyData) *ExecuteMetaQueryResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteMetaQueryResponseBody) SetMessage(v string) *ExecuteMetaQueryResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteMetaQueryResponseBody) SetRequestId(v string) *ExecuteMetaQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteMetaQueryResponseBody) SetSuccess(v bool) *ExecuteMetaQueryResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteMetaQueryResponseBody) Validate() error {
  if s.AccessDeniedDetail != nil {
    if err := s.AccessDeniedDetail.Validate(); err != nil {
      return err
    }
  }
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteMetaQueryResponseBodyAccessDeniedDetail struct {
  // The description is the same as above.
  // 
  // example:
  // 
  // xxx
  AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
  // The description is the same as above.
  // 
  // example:
  // 
  // 222
  AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
  // The diagnostic information.
  // 
  // example:
  // 
  // AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
  EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
  // NoPermissionType
  // 
  // example:
  // 
  // ImplicitDeny
  NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
  // PolicyType
  // 
  // example:
  // 
  // PRIORITY
  PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ExecuteMetaQueryResponseBodyAccessDeniedDetail) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMetaQueryResponseBodyAccessDeniedDetail) GoString() string {
  return s.String()
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) GetAuthAction() *string  {
  return s.AuthAction
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string  {
  return s.AuthPrincipalType
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string  {
  return s.EncodedDiagnosticMessage
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) GetNoPermissionType() *string  {
  return s.NoPermissionType
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) GetPolicyType() *string  {
  return s.PolicyType
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ExecuteMetaQueryResponseBodyAccessDeniedDetail {
  s.AuthAction = &v
  return s
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ExecuteMetaQueryResponseBodyAccessDeniedDetail {
  s.AuthPrincipalType = &v
  return s
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ExecuteMetaQueryResponseBodyAccessDeniedDetail {
  s.EncodedDiagnosticMessage = &v
  return s
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ExecuteMetaQueryResponseBodyAccessDeniedDetail {
  s.NoPermissionType = &v
  return s
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ExecuteMetaQueryResponseBodyAccessDeniedDetail {
  s.PolicyType = &v
  return s
}

func (s *ExecuteMetaQueryResponseBodyAccessDeniedDetail) Validate() error {
  return dara.Validate(s)
}

type ExecuteMetaQueryResponseBodyData struct {
  // The column names.
  Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
  // The total number of data rows.
  // 
  // example:
  // 
  // 97901
  RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
  // The number of affected or returned rows. This field is available only for compute nodes (CNs).
  Rows []map[string]interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteMetaQueryResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMetaQueryResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteMetaQueryResponseBodyData) GetColumns() []*string  {
  return s.Columns
}

func (s *ExecuteMetaQueryResponseBodyData) GetRowCount() *int32  {
  return s.RowCount
}

func (s *ExecuteMetaQueryResponseBodyData) GetRows() []map[string]interface{}  {
  return s.Rows
}

func (s *ExecuteMetaQueryResponseBodyData) SetColumns(v []*string) *ExecuteMetaQueryResponseBodyData {
  s.Columns = v
  return s
}

func (s *ExecuteMetaQueryResponseBodyData) SetRowCount(v int32) *ExecuteMetaQueryResponseBodyData {
  s.RowCount = &v
  return s
}

func (s *ExecuteMetaQueryResponseBodyData) SetRows(v []map[string]interface{}) *ExecuteMetaQueryResponseBodyData {
  s.Rows = v
  return s
}

func (s *ExecuteMetaQueryResponseBodyData) Validate() error {
  return dara.Validate(s)
}

