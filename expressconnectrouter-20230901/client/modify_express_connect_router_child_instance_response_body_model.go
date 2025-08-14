// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterChildInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyExpressConnectRouterChildInstanceResponseBody
	GetSuccess() *bool
}

type ModifyExpressConnectRouterChildInstanceResponseBody struct {
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IllegalParamFormat.EcrId
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DEB383C4-32C9-53DC-9B8B-8DBA26038074
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterChildInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyExpressConnectRouterChildInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
