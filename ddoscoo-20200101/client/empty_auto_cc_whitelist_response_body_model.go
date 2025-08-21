// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptyAutoCcWhitelistResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EmptyAutoCcWhitelistResponseBody
  GetRequestId() *string 
}

type EmptyAutoCcWhitelistResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptyAutoCcWhitelistResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EmptyAutoCcWhitelistResponseBody) GoString() string {
  return s.String()
}

func (s *EmptyAutoCcWhitelistResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EmptyAutoCcWhitelistResponseBody) SetRequestId(v string) *EmptyAutoCcWhitelistResponseBody {
  s.RequestId = &v
  return s
}

func (s *EmptyAutoCcWhitelistResponseBody) Validate() error {
  return dara.Validate(s)
}

