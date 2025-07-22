// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableExpressConnectRouterRouteEntriesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetAccessDeniedDetail() *string 
  SetCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetCode() *string 
  SetDynamicCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetDynamicMessage() *string 
  SetHttpStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableExpressConnectRouterRouteEntriesResponseBody
  GetSuccess() *bool 
}

type EnableExpressConnectRouterRouteEntriesResponseBody struct {
  // The details about the access denial.
  // 
  // example:
  // 
  // Authentication is failed for ****
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  // The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The dynamic error code.
  // 
  // example:
  // 
  // IllegalParamFormat.EcrId
  DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
  // The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
  // 
  // >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
  // 
  // example:
  // 
  // The param format of EcrId ***	- is illegal.
  DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 6FABF516-FED3-5697-BDA2-B18C5D9A****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesResponseBody) GoString() string {
  return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.Code = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.Message = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *EnableExpressConnectRouterRouteEntriesResponseBody {
  s.Success = &v
  return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) Validate() error {
  return dara.Validate(s)
}

