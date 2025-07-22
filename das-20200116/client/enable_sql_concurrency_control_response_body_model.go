// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlConcurrencyControlResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableSqlConcurrencyControlResponseBody
  GetCode() *string 
  SetData(v string) *EnableSqlConcurrencyControlResponseBody
  GetData() *string 
  SetMessage(v string) *EnableSqlConcurrencyControlResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableSqlConcurrencyControlResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *EnableSqlConcurrencyControlResponseBody
  GetSuccess() *string 
}

type EnableSqlConcurrencyControlResponseBody struct {
  // The HTTP status code returned.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The detailed information, including the error codes and the number of entries that are returned.
  // 
  // example:
  // 
  // Null
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The returned message.
  // 
  // >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
  // 
  // example:
  // 
  // Successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // B6D17591-B48B-4D31-9CD6-9B9796B2****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // example:
  // 
  // true
  Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSqlConcurrencyControlResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlConcurrencyControlResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSqlConcurrencyControlResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableSqlConcurrencyControlResponseBody) GetData() *string  {
  return s.Data
}

func (s *EnableSqlConcurrencyControlResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableSqlConcurrencyControlResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSqlConcurrencyControlResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *EnableSqlConcurrencyControlResponseBody) SetCode(v string) *EnableSqlConcurrencyControlResponseBody {
  s.Code = &v
  return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetData(v string) *EnableSqlConcurrencyControlResponseBody {
  s.Data = &v
  return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetMessage(v string) *EnableSqlConcurrencyControlResponseBody {
  s.Message = &v
  return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetRequestId(v string) *EnableSqlConcurrencyControlResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSqlConcurrencyControlResponseBody) SetSuccess(v string) *EnableSqlConcurrencyControlResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSqlConcurrencyControlResponseBody) Validate() error {
  return dara.Validate(s)
}

