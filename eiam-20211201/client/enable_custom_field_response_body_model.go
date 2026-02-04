// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomFieldResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCustomFieldResponseBody
  GetRequestId() *string 
}

type EnableCustomFieldResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCustomFieldResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomFieldResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCustomFieldResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCustomFieldResponseBody) SetRequestId(v string) *EnableCustomFieldResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCustomFieldResponseBody) Validate() error {
  return dara.Validate(s)
}

