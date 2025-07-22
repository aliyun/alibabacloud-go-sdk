// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachExpressConnectRouterChildInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AttachExpressConnectRouterChildInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody
	GetCode() *string
	SetDynamicCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachExpressConnectRouterChildInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AttachExpressConnectRouterChildInstanceResponseBody
	GetSuccess() *bool
}

type AttachExpressConnectRouterChildInstanceResponseBody struct {
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

func (s AttachExpressConnectRouterChildInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
