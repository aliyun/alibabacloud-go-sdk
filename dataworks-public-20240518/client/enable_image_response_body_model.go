// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableImageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableImageResponseBody
  GetData() *bool 
  SetRequestId(v string) *EnableImageResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableImageResponseBody
  GetSuccess() *bool 
}

type EnableImageResponseBody struct {
  // The result of the API request.
  // 
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID, which is used to locate logs and troubleshoot issues.
  // 
  // example:
  // 
  // 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableImageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableImageResponseBody) GoString() string {
  return s.String()
}

func (s *EnableImageResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableImageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableImageResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableImageResponseBody) SetData(v bool) *EnableImageResponseBody {
  s.Data = &v
  return s
}

func (s *EnableImageResponseBody) SetRequestId(v string) *EnableImageResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableImageResponseBody) SetSuccess(v bool) *EnableImageResponseBody {
  s.Success = &v
  return s
}

func (s *EnableImageResponseBody) Validate() error {
  return dara.Validate(s)
}

