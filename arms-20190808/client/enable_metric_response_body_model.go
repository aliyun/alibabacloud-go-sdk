// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMetricResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int64) *EnableMetricResponseBody
  GetCode() *int64 
  SetData(v string) *EnableMetricResponseBody
  GetData() *string 
  SetMessage(v string) *EnableMetricResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableMetricResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableMetricResponseBody
  GetSuccess() *bool 
}

type EnableMetricResponseBody struct {
  // The HTTP status code. The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned struct.
  // 
  // example:
  // 
  // success
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 0231DA4B-3D11-5433-9376-3B5B46C7228D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableMetricResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableMetricResponseBody) GoString() string {
  return s.String()
}

func (s *EnableMetricResponseBody) GetCode() *int64  {
  return s.Code
}

func (s *EnableMetricResponseBody) GetData() *string  {
  return s.Data
}

func (s *EnableMetricResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableMetricResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableMetricResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableMetricResponseBody) SetCode(v int64) *EnableMetricResponseBody {
  s.Code = &v
  return s
}

func (s *EnableMetricResponseBody) SetData(v string) *EnableMetricResponseBody {
  s.Data = &v
  return s
}

func (s *EnableMetricResponseBody) SetMessage(v string) *EnableMetricResponseBody {
  s.Message = &v
  return s
}

func (s *EnableMetricResponseBody) SetRequestId(v string) *EnableMetricResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableMetricResponseBody) SetSuccess(v bool) *EnableMetricResponseBody {
  s.Success = &v
  return s
}

func (s *EnableMetricResponseBody) Validate() error {
  return dara.Validate(s)
}

