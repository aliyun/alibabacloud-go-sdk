// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndSessionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EndSessionResponseBody
  GetCode() *string 
  SetData(v string) *EndSessionResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *EndSessionResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EndSessionResponseBody
  GetMessage() *string 
  SetParams(v []*string) *EndSessionResponseBody
  GetParams() []*string 
  SetRequestId(v string) *EndSessionResponseBody
  GetRequestId() *string 
}

type EndSessionResponseBody struct {
  // The business processing status code. A value of OK indicates success. If an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration. No error codes are added or modified for this operation.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The session ID of the terminated text session, which corresponds to the session ID in the request. The service still returns this ID even if the session context has expired.
  // 
  // example:
  // 
  // text-11111111-2222-4333-8444-555555555555
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The HTTP status code recorded in the business response body. A value of 200 indicates a normal response. To determine whether the business operation is successful, you must also check the Code field. You cannot determine the success of the call based solely on the presence of this field.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The business response or error message. This field is usually empty or omitted for successful requests. For failed requests, it provides an error description to assist in troubleshooting. Use this field in conjunction with the Code and Params fields to identify the cause of the error. The example shows the error message returned when an instance does not exist, which corresponds to the NotExists.InstanceId error code. This example is for format reference only.
  // 
  // example:
  // 
  // Instance instance-example-001 does not exist.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The list of parameters used to localize error messages. These are not scenario variables or business results. A successful response usually returns an empty array. If an error occurs, the parameters corresponding to the error template may be returned.
  // 
  // example:
  // 
  // []
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // The ID used to trace and troubleshoot the current POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway, so you do not need to specify system fields.
  // 
  // example:
  // 
  // AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndSessionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndSessionResponseBody) GoString() string {
  return s.String()
}

func (s *EndSessionResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EndSessionResponseBody) GetData() *string  {
  return s.Data
}

func (s *EndSessionResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EndSessionResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EndSessionResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *EndSessionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndSessionResponseBody) SetCode(v string) *EndSessionResponseBody {
  s.Code = &v
  return s
}

func (s *EndSessionResponseBody) SetData(v string) *EndSessionResponseBody {
  s.Data = &v
  return s
}

func (s *EndSessionResponseBody) SetHttpStatusCode(v int32) *EndSessionResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EndSessionResponseBody) SetMessage(v string) *EndSessionResponseBody {
  s.Message = &v
  return s
}

func (s *EndSessionResponseBody) SetParams(v []*string) *EndSessionResponseBody {
  s.Params = v
  return s
}

func (s *EndSessionResponseBody) SetRequestId(v string) *EndSessionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndSessionResponseBody) Validate() error {
  return dara.Validate(s)
}

