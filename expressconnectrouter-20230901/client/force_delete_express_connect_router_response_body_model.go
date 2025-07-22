// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForceDeleteExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ForceDeleteExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ForceDeleteExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ForceDeleteExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ForceDeleteExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ForceDeleteExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ForceDeleteExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ForceDeleteExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ForceDeleteExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type ForceDeleteExpressConnectRouterResponseBody struct {
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

func (s ForceDeleteExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForceDeleteExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetCode(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetDynamicCode(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetDynamicMessage(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *ForceDeleteExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetMessage(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetRequestId(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetSuccess(v bool) *ForceDeleteExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
