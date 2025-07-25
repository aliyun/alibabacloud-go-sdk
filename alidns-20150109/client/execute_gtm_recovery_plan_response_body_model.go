// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteGtmRecoveryPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteGtmRecoveryPlanResponseBody
  GetRequestId() *string 
}

type ExecuteGtmRecoveryPlanResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 6856BCF6-11D6-4D7E-AC53-FD579933522B
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteGtmRecoveryPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteGtmRecoveryPlanResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteGtmRecoveryPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteGtmRecoveryPlanResponseBody) SetRequestId(v string) *ExecuteGtmRecoveryPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteGtmRecoveryPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

