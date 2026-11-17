// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndTrialSessionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EndTrialSessionResponseBody
  GetCode() *string 
  SetData(v string) *EndTrialSessionResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *EndTrialSessionResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EndTrialSessionResponseBody
  GetMessage() *string 
  SetParams(v []*string) *EndTrialSessionResponseBody
  GetParams() []*string 
  SetRequestId(v string) *EndTrialSessionResponseBody
  GetRequestId() *string 
}

type EndTrialSessionResponseBody struct {
  // The business processing status code. OK indicates success. If an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration. No error codes are added or modified in this operation.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The SessionId of the text session ended in this request, which corresponds to the SessionId in the request. The termination service still returns this ID even if the context has expired.
  // 
  // example:
  // 
  // text-11111111-2222-4333-8444-555555555555
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine the success of the call solely based on the presence of this field.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The business response or error message. It is typically empty or not returned upon success. Do not expect a fixed text string upon success. Upon failure, it provides an error description to assist with troubleshooting. Use it in conjunction with Code and Params to identify the cause of the error. The example shows an error message for an expired visitor token in an experience link (corresponding to the Expired.VisitorToken code) and serves as a format reference only.
  // 
  // example:
  // 
  // VisitorToken: example-visitor-token is expired.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The list of parameters used for the localized display of error messages. These are not scenario variables or business results. A normal successful response usually returns an empty array. When an error occurs, parameters corresponding to the error template may be returned.
  // 
  // example:
  // 
  // []
  Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
  // The ID used to associate and troubleshoot this POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway, and the caller does not need to fill in system fields.
  // 
  // example:
  // 
  // AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndTrialSessionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndTrialSessionResponseBody) GoString() string {
  return s.String()
}

func (s *EndTrialSessionResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EndTrialSessionResponseBody) GetData() *string  {
  return s.Data
}

func (s *EndTrialSessionResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EndTrialSessionResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EndTrialSessionResponseBody) GetParams() []*string  {
  return s.Params
}

func (s *EndTrialSessionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndTrialSessionResponseBody) SetCode(v string) *EndTrialSessionResponseBody {
  s.Code = &v
  return s
}

func (s *EndTrialSessionResponseBody) SetData(v string) *EndTrialSessionResponseBody {
  s.Data = &v
  return s
}

func (s *EndTrialSessionResponseBody) SetHttpStatusCode(v int32) *EndTrialSessionResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EndTrialSessionResponseBody) SetMessage(v string) *EndTrialSessionResponseBody {
  s.Message = &v
  return s
}

func (s *EndTrialSessionResponseBody) SetParams(v []*string) *EndTrialSessionResponseBody {
  s.Params = v
  return s
}

func (s *EndTrialSessionResponseBody) SetRequestId(v string) *EndTrialSessionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndTrialSessionResponseBody) Validate() error {
  return dara.Validate(s)
}

