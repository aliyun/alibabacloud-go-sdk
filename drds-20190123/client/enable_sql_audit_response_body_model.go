// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlAuditResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSqlAuditResponseBody
  GetRequestId() *string 
  SetResult(v bool) *EnableSqlAuditResponseBody
  GetResult() *bool 
  SetSuccess(v bool) *EnableSqlAuditResponseBody
  GetSuccess() *bool 
}

type EnableSqlAuditResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 463A5F0F-12AD-4544-A902-B2B983******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indices whether the SQL audit feature is enabled.
  // 
  // example:
  // 
  // true
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlAuditResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlAuditResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSqlAuditResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSqlAuditResponseBody) GetResult() *bool  {
  return s.Result
}

func (s *EnableSqlAuditResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSqlAuditResponseBody) SetRequestId(v string) *EnableSqlAuditResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSqlAuditResponseBody) SetResult(v bool) *EnableSqlAuditResponseBody {
  s.Result = &v
  return s
}

func (s *EnableSqlAuditResponseBody) SetSuccess(v bool) *EnableSqlAuditResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSqlAuditResponseBody) Validate() error {
  return dara.Validate(s)
}

