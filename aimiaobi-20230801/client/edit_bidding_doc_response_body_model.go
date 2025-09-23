// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditBiddingDocResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditBiddingDocResponseBody
  GetCode() *string 
  SetData(v *EditBiddingDocResponseBodyData) *EditBiddingDocResponseBody
  GetData() *EditBiddingDocResponseBodyData 
  SetHttpStatusCode(v int32) *EditBiddingDocResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EditBiddingDocResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EditBiddingDocResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditBiddingDocResponseBody
  GetSuccess() *bool 
}

type EditBiddingDocResponseBody struct {
  // example:
  // 
  // successful
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EditBiddingDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 1813ceee-7fe5-41b4-87e5-982a4d18cca5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditBiddingDocResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditBiddingDocResponseBody) GoString() string {
  return s.String()
}

func (s *EditBiddingDocResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditBiddingDocResponseBody) GetData() *EditBiddingDocResponseBodyData  {
  return s.Data
}

func (s *EditBiddingDocResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EditBiddingDocResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditBiddingDocResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditBiddingDocResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditBiddingDocResponseBody) SetCode(v string) *EditBiddingDocResponseBody {
  s.Code = &v
  return s
}

func (s *EditBiddingDocResponseBody) SetData(v *EditBiddingDocResponseBodyData) *EditBiddingDocResponseBody {
  s.Data = v
  return s
}

func (s *EditBiddingDocResponseBody) SetHttpStatusCode(v int32) *EditBiddingDocResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EditBiddingDocResponseBody) SetMessage(v string) *EditBiddingDocResponseBody {
  s.Message = &v
  return s
}

func (s *EditBiddingDocResponseBody) SetRequestId(v string) *EditBiddingDocResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditBiddingDocResponseBody) SetSuccess(v bool) *EditBiddingDocResponseBody {
  s.Success = &v
  return s
}

func (s *EditBiddingDocResponseBody) Validate() error {
  return dara.Validate(s)
}

type EditBiddingDocResponseBodyData struct {
  // example:
  // 
  // 3f7045e099474ba28ceca1b4eb6d6e21
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EditBiddingDocResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EditBiddingDocResponseBodyData) GoString() string {
  return s.String()
}

func (s *EditBiddingDocResponseBodyData) GetTaskId() *string  {
  return s.TaskId
}

func (s *EditBiddingDocResponseBodyData) SetTaskId(v string) *EditBiddingDocResponseBodyData {
  s.TaskId = &v
  return s
}

func (s *EditBiddingDocResponseBodyData) Validate() error {
  return dara.Validate(s)
}

