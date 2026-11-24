// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachExpressConnectRouterChildInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DetachExpressConnectRouterChildInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachExpressConnectRouterChildInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetachExpressConnectRouterChildInstanceResponseBody
	GetSuccess() *bool
}

type DetachExpressConnectRouterChildInstanceResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the call is successful. A return value of 200 indicates success, and other values indicate failure. For the cause of failure, see error codes.
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
	// The dynamic error message, which is used to replace `%s` in the error message of the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the specified **DtsJobId*	- request parameter is invalid.
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
	// The message returned for the operation.
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
	// Indicates whether the call is successful. Valid values:
	//
	// - **True**: successful.
	//
	// - **False**: failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
