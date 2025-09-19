// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceAccessResourceDirectoryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableServiceAccessResourceDirectoryResponseBody
  GetRequestId() *string 
}

type EnableServiceAccessResourceDirectoryResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 52870893-48A7-5A9E-9E05-6253E5B6****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServiceAccessResourceDirectoryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceAccessResourceDirectoryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableServiceAccessResourceDirectoryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableServiceAccessResourceDirectoryResponseBody) SetRequestId(v string) *EnableServiceAccessResourceDirectoryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableServiceAccessResourceDirectoryResponseBody) Validate() error {
  return dara.Validate(s)
}

