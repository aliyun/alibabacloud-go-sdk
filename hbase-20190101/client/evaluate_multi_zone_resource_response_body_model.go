// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateMultiZoneResourceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EvaluateMultiZoneResourceResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EvaluateMultiZoneResourceResponseBody
  GetSuccess() *bool 
}

type EvaluateMultiZoneResourceResponseBody struct {
  // example:
  // 
  // FB703B69-D4D4-4879-B9FE-6A37F67C46FD
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EvaluateMultiZoneResourceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateMultiZoneResourceResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateMultiZoneResourceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateMultiZoneResourceResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EvaluateMultiZoneResourceResponseBody) SetRequestId(v string) *EvaluateMultiZoneResourceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateMultiZoneResourceResponseBody) SetSuccess(v bool) *EvaluateMultiZoneResourceResponseBody {
  s.Success = &v
  return s
}

func (s *EvaluateMultiZoneResourceResponseBody) Validate() error {
  return dara.Validate(s)
}

