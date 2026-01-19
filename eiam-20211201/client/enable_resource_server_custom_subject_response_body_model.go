// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceServerCustomSubjectResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableResourceServerCustomSubjectResponseBody
  GetRequestId() *string 
}

type EnableResourceServerCustomSubjectResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableResourceServerCustomSubjectResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceServerCustomSubjectResponseBody) GoString() string {
  return s.String()
}

func (s *EnableResourceServerCustomSubjectResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableResourceServerCustomSubjectResponseBody) SetRequestId(v string) *EnableResourceServerCustomSubjectResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableResourceServerCustomSubjectResponseBody) Validate() error {
  return dara.Validate(s)
}

