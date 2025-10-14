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
  SetSuccess(v bool) *EnableSqlAuditResponseBody
  GetSuccess() *bool 
}

type EnableSqlAuditResponseBody struct {
  // example:
  // 
  // DC3ABA3E-0F8A-4596-9104-F5167C******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *EnableSqlAuditResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSqlAuditResponseBody) SetRequestId(v string) *EnableSqlAuditResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSqlAuditResponseBody) SetSuccess(v bool) *EnableSqlAuditResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSqlAuditResponseBody) Validate() error {
  return dara.Validate(s)
}

