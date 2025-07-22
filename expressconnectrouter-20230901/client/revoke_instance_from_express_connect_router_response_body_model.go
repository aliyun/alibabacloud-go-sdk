// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeInstanceFromExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type RevokeInstanceFromExpressConnectRouterResponseBody struct {
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetDynamicCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetDynamicMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetRequestId(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetSuccess(v bool) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
