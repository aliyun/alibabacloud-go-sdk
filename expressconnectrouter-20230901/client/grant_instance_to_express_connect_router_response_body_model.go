// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GrantInstanceToExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GrantInstanceToExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *GrantInstanceToExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantInstanceToExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrantInstanceToExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type GrantInstanceToExpressConnectRouterResponseBody struct {
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

func (s GrantInstanceToExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetCode(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetDynamicCode(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetRequestId(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetSuccess(v bool) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
