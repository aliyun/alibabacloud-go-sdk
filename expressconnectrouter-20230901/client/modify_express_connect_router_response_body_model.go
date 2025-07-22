// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type ModifyExpressConnectRouterResponseBody struct {
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
	// IllegalParamFormat.Name
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in **ErrMessage**.
	//
	// >  For example, if the value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
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
	// Indicates whether the request was successful Valid values:
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

func (s ModifyExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetCode(v string) *ModifyExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetMessage(v string) *ModifyExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
