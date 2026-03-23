// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectApConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *EffectApConfigResponseBody
  GetData() *string 
  SetErrorCode(v int32) *EffectApConfigResponseBody
  GetErrorCode() *int32 
  SetErrorMessage(v string) *EffectApConfigResponseBody
  GetErrorMessage() *string 
  SetIsSuccess(v bool) *EffectApConfigResponseBody
  GetIsSuccess() *bool 
}

type EffectApConfigResponseBody struct {
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s EffectApConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EffectApConfigResponseBody) GoString() string {
  return s.String()
}

func (s *EffectApConfigResponseBody) GetData() *string  {
  return s.Data
}

func (s *EffectApConfigResponseBody) GetErrorCode() *int32  {
  return s.ErrorCode
}

func (s *EffectApConfigResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EffectApConfigResponseBody) GetIsSuccess() *bool  {
  return s.IsSuccess
}

func (s *EffectApConfigResponseBody) SetData(v string) *EffectApConfigResponseBody {
  s.Data = &v
  return s
}

func (s *EffectApConfigResponseBody) SetErrorCode(v int32) *EffectApConfigResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EffectApConfigResponseBody) SetErrorMessage(v string) *EffectApConfigResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EffectApConfigResponseBody) SetIsSuccess(v bool) *EffectApConfigResponseBody {
  s.IsSuccess = &v
  return s
}

func (s *EffectApConfigResponseBody) Validate() error {
  return dara.Validate(s)
}

