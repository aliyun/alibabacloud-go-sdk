// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDataIngestionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDataIngestionResponseBody
  GetRequestId() *string 
}

type EnableDataIngestionResponseBody struct {
  // example:
  // 
  // 6276D891-*****-55B2-87B9-74D413F7****。
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDataIngestionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDataIngestionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDataIngestionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDataIngestionResponseBody) SetRequestId(v string) *EnableDataIngestionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDataIngestionResponseBody) Validate() error {
  return dara.Validate(s)
}

