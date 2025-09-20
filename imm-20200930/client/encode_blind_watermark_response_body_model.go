// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncodeBlindWatermarkResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EncodeBlindWatermarkResponseBody
  GetRequestId() *string 
}

type EncodeBlindWatermarkResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 8E0DD64B-28C6-4653-8FF7-93E4C234BCF0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncodeBlindWatermarkResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EncodeBlindWatermarkResponseBody) GoString() string {
  return s.String()
}

func (s *EncodeBlindWatermarkResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EncodeBlindWatermarkResponseBody) SetRequestId(v string) *EncodeBlindWatermarkResponseBody {
  s.RequestId = &v
  return s
}

func (s *EncodeBlindWatermarkResponseBody) Validate() error {
  return dara.Validate(s)
}

