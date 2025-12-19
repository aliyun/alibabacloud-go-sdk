// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteTaskResponseBody
  GetCode() *string 
  SetData(v *ExecuteTaskResponseBodyData) *ExecuteTaskResponseBody
  GetData() *ExecuteTaskResponseBodyData 
  SetMessage(v string) *ExecuteTaskResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteTaskResponseBody
  GetRequestId() *string 
}

type ExecuteTaskResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *ExecuteTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // Success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // A07FFDF2-78FA-1B48-9E38-88E833A93187
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTaskResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteTaskResponseBody) GetData() *ExecuteTaskResponseBodyData  {
  return s.Data
}

func (s *ExecuteTaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTaskResponseBody) SetCode(v string) *ExecuteTaskResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteTaskResponseBody) SetData(v *ExecuteTaskResponseBodyData) *ExecuteTaskResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteTaskResponseBody) SetMessage(v string) *ExecuteTaskResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteTaskResponseBody) SetRequestId(v string) *ExecuteTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTaskResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTaskResponseBodyData struct {
  // example:
  // 
  // SyntaxError
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // example:
  // 
  // 2615
  Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
  // example:
  // 
  // test_name
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // example:
  // 
  // FINISH
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ExecuteTaskResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteTaskResponseBodyData) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteTaskResponseBodyData) GetId() *int32  {
  return s.Id
}

func (s *ExecuteTaskResponseBodyData) GetName() *string  {
  return s.Name
}

func (s *ExecuteTaskResponseBodyData) GetStatus() *string  {
  return s.Status
}

func (s *ExecuteTaskResponseBodyData) SetErrorMessage(v string) *ExecuteTaskResponseBodyData {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteTaskResponseBodyData) SetId(v int32) *ExecuteTaskResponseBodyData {
  s.Id = &v
  return s
}

func (s *ExecuteTaskResponseBodyData) SetName(v string) *ExecuteTaskResponseBodyData {
  s.Name = &v
  return s
}

func (s *ExecuteTaskResponseBodyData) SetStatus(v string) *ExecuteTaskResponseBodyData {
  s.Status = &v
  return s
}

func (s *ExecuteTaskResponseBodyData) Validate() error {
  return dara.Validate(s)
}

