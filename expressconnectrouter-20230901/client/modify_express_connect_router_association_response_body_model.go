// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyExpressConnectRouterAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyExpressConnectRouterAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyExpressConnectRouterAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyExpressConnectRouterAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyExpressConnectRouterAssociationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyExpressConnectRouterAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyExpressConnectRouterAssociationResponseBody
	GetSuccess() *bool
}

type ModifyExpressConnectRouterAssociationResponseBody struct {
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
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DFDA79A0-D135-5193-B0AC-3B5608FDB1D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetCode(v string) *ModifyExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetMessage(v string) *ModifyExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
