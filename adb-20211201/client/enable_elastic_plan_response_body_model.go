// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableElasticPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableElasticPlanResponseBody
  GetRequestId() *string 
}

type EnableElasticPlanResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // A5C433C2-001F-58E3-99F5-3274C14DF8BD
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableElasticPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableElasticPlanResponseBody) GoString() string {
  return s.String()
}

func (s *EnableElasticPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableElasticPlanResponseBody) SetRequestId(v string) *EnableElasticPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableElasticPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

