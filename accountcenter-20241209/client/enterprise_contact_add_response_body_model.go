// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactAddResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseContactAddResponseBody
  GetCode() *string 
  SetData(v *EnterpriseContactAddResponseBodyData) *EnterpriseContactAddResponseBody
  GetData() *EnterpriseContactAddResponseBodyData 
  SetMessage(v string) *EnterpriseContactAddResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseContactAddResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseContactAddResponseBody
  GetSuccess() *bool 
}

type EnterpriseContactAddResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseContactAddResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // Successful!
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 40EA3ECB-2167-5A8E-9327-F7E59E508FA8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseContactAddResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactAddResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseContactAddResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseContactAddResponseBody) GetData() *EnterpriseContactAddResponseBodyData  {
  return s.Data
}

func (s *EnterpriseContactAddResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseContactAddResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseContactAddResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseContactAddResponseBody) SetCode(v string) *EnterpriseContactAddResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseContactAddResponseBody) SetData(v *EnterpriseContactAddResponseBodyData) *EnterpriseContactAddResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseContactAddResponseBody) SetMessage(v string) *EnterpriseContactAddResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseContactAddResponseBody) SetRequestId(v string) *EnterpriseContactAddResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseContactAddResponseBody) SetSuccess(v bool) *EnterpriseContactAddResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseContactAddResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseContactAddResponseBodyData struct {
  // example:
  // 
  // xxx
  ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
  ErrorList []*EnterpriseContactAddResponseBodyDataErrorList `json:"ErrorList,omitempty" xml:"ErrorList,omitempty" type:"Repeated"`
  // example:
  // 
  // xxx
  Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EnterpriseContactAddResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactAddResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseContactAddResponseBodyData) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactAddResponseBodyData) GetErrorList() []*EnterpriseContactAddResponseBodyDataErrorList  {
  return s.ErrorList
}

func (s *EnterpriseContactAddResponseBodyData) GetResult() *bool  {
  return s.Result
}

func (s *EnterpriseContactAddResponseBodyData) SetContactId(v int64) *EnterpriseContactAddResponseBodyData {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactAddResponseBodyData) SetErrorList(v []*EnterpriseContactAddResponseBodyDataErrorList) *EnterpriseContactAddResponseBodyData {
  s.ErrorList = v
  return s
}

func (s *EnterpriseContactAddResponseBodyData) SetResult(v bool) *EnterpriseContactAddResponseBodyData {
  s.Result = &v
  return s
}

func (s *EnterpriseContactAddResponseBodyData) Validate() error {
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

type EnterpriseContactAddResponseBodyDataErrorList struct {
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

func (s EnterpriseContactAddResponseBodyDataErrorList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactAddResponseBodyDataErrorList) GoString() string {
  return s.String()
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) GetErrorDesc() *string  {
  return s.ErrorDesc
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) GetItem() *string  {
  return s.Item
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) SetErrorCode(v string) *EnterpriseContactAddResponseBodyDataErrorList {
  s.ErrorCode = &v
  return s
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) SetErrorDesc(v string) *EnterpriseContactAddResponseBodyDataErrorList {
  s.ErrorDesc = &v
  return s
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) SetItem(v string) *EnterpriseContactAddResponseBodyDataErrorList {
  s.Item = &v
  return s
}

func (s *EnterpriseContactAddResponseBodyDataErrorList) Validate() error {
  return dara.Validate(s)
}

