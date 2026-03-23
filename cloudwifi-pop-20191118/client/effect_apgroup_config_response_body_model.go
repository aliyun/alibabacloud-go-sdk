// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectApgroupConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *EffectApgroupConfigResponseBody
  GetData() *string 
  SetErrorCode(v int32) *EffectApgroupConfigResponseBody
  GetErrorCode() *int32 
  SetErrorMessage(v string) *EffectApgroupConfigResponseBody
  GetErrorMessage() *string 
  SetIsSuccess(v bool) *EffectApgroupConfigResponseBody
  GetIsSuccess() *bool 
}

type EffectApgroupConfigResponseBody struct {
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s EffectApgroupConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EffectApgroupConfigResponseBody) GoString() string {
  return s.String()
}

func (s *EffectApgroupConfigResponseBody) GetData() *string  {
  return s.Data
}

func (s *EffectApgroupConfigResponseBody) GetErrorCode() *int32  {
  return s.ErrorCode
}

func (s *EffectApgroupConfigResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EffectApgroupConfigResponseBody) GetIsSuccess() *bool  {
  return s.IsSuccess
}

func (s *EffectApgroupConfigResponseBody) SetData(v string) *EffectApgroupConfigResponseBody {
  s.Data = &v
  return s
}

func (s *EffectApgroupConfigResponseBody) SetErrorCode(v int32) *EffectApgroupConfigResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EffectApgroupConfigResponseBody) SetErrorMessage(v string) *EffectApgroupConfigResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EffectApgroupConfigResponseBody) SetIsSuccess(v bool) *EffectApgroupConfigResponseBody {
  s.IsSuccess = &v
  return s
}

func (s *EffectApgroupConfigResponseBody) Validate() error {
  return dara.Validate(s)
}

