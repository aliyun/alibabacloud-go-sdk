// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDelegateAccountResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDelegateAccountResponseBody
  GetRequestId() *string 
}

type EnableDelegateAccountResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 768F908D-A66A-5A5D-816C-20C93CBBFEE3
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDelegateAccountResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDelegateAccountResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDelegateAccountResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDelegateAccountResponseBody) SetRequestId(v string) *EnableDelegateAccountResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDelegateAccountResponseBody) Validate() error {
  return dara.Validate(s)
}

