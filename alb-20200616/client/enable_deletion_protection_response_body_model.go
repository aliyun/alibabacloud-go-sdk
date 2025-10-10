// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeletionProtectionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDeletionProtectionResponseBody
  GetRequestId() *string 
}

type EnableDeletionProtectionResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // CEF72CEB-54B6-4AE8-B225-F876FF7BA984
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDeletionProtectionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDeletionProtectionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDeletionProtectionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDeletionProtectionResponseBody) SetRequestId(v string) *EnableDeletionProtectionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDeletionProtectionResponseBody) Validate() error {
  return dara.Validate(s)
}

