// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptyAutoCcBlacklistResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EmptyAutoCcBlacklistResponseBody
  GetRequestId() *string 
}

type EmptyAutoCcBlacklistResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptyAutoCcBlacklistResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EmptyAutoCcBlacklistResponseBody) GoString() string {
  return s.String()
}

func (s *EmptyAutoCcBlacklistResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EmptyAutoCcBlacklistResponseBody) SetRequestId(v string) *EmptyAutoCcBlacklistResponseBody {
  s.RequestId = &v
  return s
}

func (s *EmptyAutoCcBlacklistResponseBody) Validate() error {
  return dara.Validate(s)
}

