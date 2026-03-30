// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScriptResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportScriptResponseBody
  GetCode() *string 
  SetData(v string) *ExportScriptResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportScriptResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportScriptResponseBody
  GetMessage() *string 
  SetParams(v []*string) *ExportScriptResponseBody
  GetParams() []*string 
  SetRequestId(v string) *ExportScriptResponseBody
  GetRequestId() *string 
}

type ExportScriptResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 64241e64-190c-45d1-af66-06f51c07b090
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // Instance af81a389-91f0-4157-8d82-720edd02b66a
  // 
  //  does not exist.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // example:
  // 
  // 83C6DA11-798F-5524-AB1D-555263E7A400
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportScriptResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportScriptResponseBody) GoString() string {
  return s.String()
}

func (s *ExportScriptResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportScriptResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportScriptResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportScriptResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportScriptResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *ExportScriptResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportScriptResponseBody) SetCode(v string) *ExportScriptResponseBody {
  s.Code = &v
  return s
}

func (s *ExportScriptResponseBody) SetData(v string) *ExportScriptResponseBody {
  s.Data = &v
  return s
}

func (s *ExportScriptResponseBody) SetHttpStatusCode(v int32) *ExportScriptResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportScriptResponseBody) SetMessage(v string) *ExportScriptResponseBody {
  s.Message = &v
  return s
}

func (s *ExportScriptResponseBody) SetParams(v []*string) *ExportScriptResponseBody {
  s.Params = v
  return s
}

func (s *ExportScriptResponseBody) SetRequestId(v string) *ExportScriptResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportScriptResponseBody) Validate() error {
  return dara.Validate(s)
}

