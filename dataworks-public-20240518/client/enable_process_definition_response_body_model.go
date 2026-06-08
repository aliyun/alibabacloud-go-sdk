// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableProcessDefinitionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableProcessDefinitionResponseBody
  GetRequestId() *string 
}

type EnableProcessDefinitionResponseBody struct {
  // example:
  // 
  // 0bc5df3a17***903790e8e8a
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableProcessDefinitionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableProcessDefinitionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableProcessDefinitionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableProcessDefinitionResponseBody) SetRequestId(v string) *EnableProcessDefinitionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableProcessDefinitionResponseBody) Validate() error {
  return dara.Validate(s)
}

