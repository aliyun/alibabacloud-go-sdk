// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRightsSeparationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMessage(v string) *EnableRightsSeparationResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableRightsSeparationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableRightsSeparationResponseBody
  GetSuccess() *bool 
}

type EnableRightsSeparationResponseBody struct {
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 73559800-3c8c-11ec-bd40-99cfcff3fe1e
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRightsSeparationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableRightsSeparationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableRightsSeparationResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableRightsSeparationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableRightsSeparationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableRightsSeparationResponseBody) SetMessage(v string) *EnableRightsSeparationResponseBody {
  s.Message = &v
  return s
}

func (s *EnableRightsSeparationResponseBody) SetRequestId(v string) *EnableRightsSeparationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableRightsSeparationResponseBody) SetSuccess(v bool) *EnableRightsSeparationResponseBody {
  s.Success = &v
  return s
}

func (s *EnableRightsSeparationResponseBody) Validate() error {
  return dara.Validate(s)
}

