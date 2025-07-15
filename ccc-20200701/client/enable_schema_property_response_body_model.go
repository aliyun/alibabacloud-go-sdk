// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSchemaPropertyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableSchemaPropertyResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *EnableSchemaPropertyResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableSchemaPropertyResponseBody
  GetMessage() *string 
  SetParams(v []*string) *EnableSchemaPropertyResponseBody
  GetParams() []*string 
  SetRequestId(v string) *EnableSchemaPropertyResponseBody
  GetRequestId() *string 
}

type EnableSchemaPropertyResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // The operation is not allowed. User state (READY) does not meet expectations (OFFLINE).
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // 2778FA12-EDD6-42AA-9B15-AF855072E5E5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSchemaPropertyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSchemaPropertyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSchemaPropertyResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableSchemaPropertyResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableSchemaPropertyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableSchemaPropertyResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *EnableSchemaPropertyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSchemaPropertyResponseBody) SetCode(v string) *EnableSchemaPropertyResponseBody {
  s.Code = &v
  return s
}

func (s *EnableSchemaPropertyResponseBody) SetHttpStatusCode(v int32) *EnableSchemaPropertyResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableSchemaPropertyResponseBody) SetMessage(v string) *EnableSchemaPropertyResponseBody {
  s.Message = &v
  return s
}

func (s *EnableSchemaPropertyResponseBody) SetParams(v []*string) *EnableSchemaPropertyResponseBody {
  s.Params = v
  return s
}

func (s *EnableSchemaPropertyResponseBody) SetRequestId(v string) *EnableSchemaPropertyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSchemaPropertyResponseBody) Validate() error {
  return dara.Validate(s)
}

