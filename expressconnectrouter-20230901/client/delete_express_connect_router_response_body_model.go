// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DeleteExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type DeleteExpressConnectRouterResponseBody struct {
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

func (s DeleteExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetCode(v string) *DeleteExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetDynamicCode(v string) *DeleteExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DeleteExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DeleteExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetMessage(v string) *DeleteExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetRequestId(v string) *DeleteExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetSuccess(v bool) *DeleteExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
