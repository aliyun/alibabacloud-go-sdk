// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactEditResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseContactEditResponseBody
  GetCode() *string 
  SetData(v *EnterpriseContactEditResponseBodyData) *EnterpriseContactEditResponseBody
  GetData() *EnterpriseContactEditResponseBodyData 
  SetMessage(v string) *EnterpriseContactEditResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseContactEditResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseContactEditResponseBody
  GetSuccess() *bool 
}

type EnterpriseContactEditResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseContactEditResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // SUCCESS
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 516C2364-18B7-5BAC-9288-AAEA85EEA351
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseContactEditResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactEditResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseContactEditResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseContactEditResponseBody) GetData() *EnterpriseContactEditResponseBodyData  {
  return s.Data
}

func (s *EnterpriseContactEditResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseContactEditResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseContactEditResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseContactEditResponseBody) SetCode(v string) *EnterpriseContactEditResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseContactEditResponseBody) SetData(v *EnterpriseContactEditResponseBodyData) *EnterpriseContactEditResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseContactEditResponseBody) SetMessage(v string) *EnterpriseContactEditResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseContactEditResponseBody) SetRequestId(v string) *EnterpriseContactEditResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseContactEditResponseBody) SetSuccess(v bool) *EnterpriseContactEditResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseContactEditResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseContactEditResponseBodyData struct {
  // example:
  // 
  // xxx
  ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
  ErrorList []*EnterpriseContactEditResponseBodyDataErrorList `json:"ErrorList,omitempty" xml:"ErrorList,omitempty" type:"Repeated"`
  // example:
  // 
  // true
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EnterpriseContactEditResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactEditResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseContactEditResponseBodyData) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactEditResponseBodyData) GetErrorList() []*EnterpriseContactEditResponseBodyDataErrorList  {
  return s.ErrorList
}

func (s *EnterpriseContactEditResponseBodyData) GetResult() *bool  {
  return s.Result
}

func (s *EnterpriseContactEditResponseBodyData) SetContactId(v int64) *EnterpriseContactEditResponseBodyData {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactEditResponseBodyData) SetErrorList(v []*EnterpriseContactEditResponseBodyDataErrorList) *EnterpriseContactEditResponseBodyData {
  s.ErrorList = v
  return s
}

func (s *EnterpriseContactEditResponseBodyData) SetResult(v bool) *EnterpriseContactEditResponseBodyData {
  s.Result = &v
  return s
}

func (s *EnterpriseContactEditResponseBodyData) Validate() error {
  if s.ErrorList != nil {
    for _, item := range s.ErrorList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseContactEditResponseBodyDataErrorList struct {
  // example:
  // 
  // MOBILE_CODE_ILLEGAL
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // example:
  // 
  // Mobile Code Illegal
  ErrorDesc *string `json:"ErrorDesc,omitempty" xml:"ErrorDesc,omitempty"`
  // example:
  // 
  // MOBILE_VERIFY_CODE
  Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
}

func (s EnterpriseContactEditResponseBodyDataErrorList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactEditResponseBodyDataErrorList) GoString() string {
  return s.String()
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) GetErrorDesc() *string  {
  return s.ErrorDesc
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) GetItem() *string  {
  return s.Item
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) SetErrorCode(v string) *EnterpriseContactEditResponseBodyDataErrorList {
  s.ErrorCode = &v
  return s
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) SetErrorDesc(v string) *EnterpriseContactEditResponseBodyDataErrorList {
  s.ErrorDesc = &v
  return s
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) SetItem(v string) *EnterpriseContactEditResponseBodyDataErrorList {
  s.Item = &v
  return s
}

func (s *EnterpriseContactEditResponseBodyDataErrorList) Validate() error {
  return dara.Validate(s)
}

