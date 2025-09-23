// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSharingWithResourceDirectoryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSharingWithResourceDirectoryResponseBody
  GetRequestId() *string 
}

type EnableSharingWithResourceDirectoryResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 2F23CFB6-A721-4E90-AC1E-0E30FA8B45DA
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSharingWithResourceDirectoryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSharingWithResourceDirectoryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSharingWithResourceDirectoryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSharingWithResourceDirectoryResponseBody) SetRequestId(v string) *EnableSharingWithResourceDirectoryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSharingWithResourceDirectoryResponseBody) Validate() error {
  return dara.Validate(s)
}

