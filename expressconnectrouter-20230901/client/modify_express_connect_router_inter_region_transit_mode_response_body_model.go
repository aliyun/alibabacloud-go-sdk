// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterInterRegionTransitModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
	GetSuccess() *bool
}

type ModifyExpressConnectRouterInterRegionTransitModeResponseBody struct {
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

func (s ModifyExpressConnectRouterInterRegionTransitModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) Validate() error {
	return dara.Validate(s)
}
