// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableInstanceResponseBody
  GetRequestId() *string 
  SetStatus(v string) *EnableInstanceResponseBody
  GetStatus() *string 
}

type EnableInstanceResponseBody struct {
  // example:
  // 
  // 3a530dc0-7cfa-48f6-9539-bf9001e77b16
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInstanceResponseBody) GetStatus() *string  {
  return s.Status
}

func (s *EnableInstanceResponseBody) SetRequestId(v string) *EnableInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInstanceResponseBody) SetStatus(v string) *EnableInstanceResponseBody {
  s.Status = &v
  return s
}

func (s *EnableInstanceResponseBody) Validate() error {
  return dara.Validate(s)
}

