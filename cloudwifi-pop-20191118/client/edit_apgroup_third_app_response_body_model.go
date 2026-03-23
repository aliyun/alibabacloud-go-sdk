// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditApgroupThirdAppResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v map[string]interface{}) *EditApgroupThirdAppResponseBody
  GetData() map[string]interface{} 
  SetErrorCode(v int32) *EditApgroupThirdAppResponseBody
  GetErrorCode() *int32 
  SetErrorMessage(v string) *EditApgroupThirdAppResponseBody
  GetErrorMessage() *string 
  SetIsSuccess(v bool) *EditApgroupThirdAppResponseBody
  GetIsSuccess() *bool 
  SetRequestId(v string) *EditApgroupThirdAppResponseBody
  GetRequestId() *string 
}

type EditApgroupThirdAppResponseBody struct {
  Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
  ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditApgroupThirdAppResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditApgroupThirdAppResponseBody) GoString() string {
  return s.String()
}

func (s *EditApgroupThirdAppResponseBody) GetData() map[string]interface{}  {
  return s.Data
}

func (s *EditApgroupThirdAppResponseBody) GetErrorCode() *int32  {
  return s.ErrorCode
}

func (s *EditApgroupThirdAppResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EditApgroupThirdAppResponseBody) GetIsSuccess() *bool  {
  return s.IsSuccess
}

func (s *EditApgroupThirdAppResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditApgroupThirdAppResponseBody) SetData(v map[string]interface{}) *EditApgroupThirdAppResponseBody {
  s.Data = v
  return s
}

func (s *EditApgroupThirdAppResponseBody) SetErrorCode(v int32) *EditApgroupThirdAppResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EditApgroupThirdAppResponseBody) SetErrorMessage(v string) *EditApgroupThirdAppResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EditApgroupThirdAppResponseBody) SetIsSuccess(v bool) *EditApgroupThirdAppResponseBody {
  s.IsSuccess = &v
  return s
}

func (s *EditApgroupThirdAppResponseBody) SetRequestId(v string) *EditApgroupThirdAppResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditApgroupThirdAppResponseBody) Validate() error {
  return dara.Validate(s)
}

