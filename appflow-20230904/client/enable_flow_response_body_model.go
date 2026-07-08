// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFlowResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *EnableFlowResponseBody
  GetData() *string 
  SetRequestId(v string) *EnableFlowResponseBody
  GetRequestId() *string 
}

type EnableFlowResponseBody struct {
  // The response data.
  // 
  // example:
  // 
  // None
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 17CADBF7-B0F4-5FE6-87EB-76B1A69AC422
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableFlowResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableFlowResponseBody) GoString() string {
  return s.String()
}

func (s *EnableFlowResponseBody) GetData() *string  {
  return s.Data
}

func (s *EnableFlowResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableFlowResponseBody) SetData(v string) *EnableFlowResponseBody {
  s.Data = &v
  return s
}

func (s *EnableFlowResponseBody) SetRequestId(v string) *EnableFlowResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableFlowResponseBody) Validate() error {
  return dara.Validate(s)
}

