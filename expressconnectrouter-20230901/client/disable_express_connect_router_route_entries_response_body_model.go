// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableExpressConnectRouterRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableExpressConnectRouterRouteEntriesResponseBody
	GetSuccess() *bool
}

type DisableExpressConnectRouterRouteEntriesResponseBody struct {
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

func (s DisableExpressConnectRouterRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
