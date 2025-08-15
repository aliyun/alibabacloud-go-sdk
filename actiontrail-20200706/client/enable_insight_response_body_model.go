// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInsightResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableInsightResponseBody
  GetRequestId() *string 
}

type EnableInsightResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 45AA79B7-0240-52AB-B158-3F9A512228ED
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInsightResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInsightResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInsightResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInsightResponseBody) SetRequestId(v string) *EnableInsightResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInsightResponseBody) Validate() error {
  return dara.Validate(s)
}

