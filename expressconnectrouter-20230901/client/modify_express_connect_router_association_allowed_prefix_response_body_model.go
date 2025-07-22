// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterAssociationAllowedPrefixResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
	GetSuccess() *bool
}

type ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody struct {
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
	// 05130E79-588D-5C40-A718-C4863A59****
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

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) Validate() error {
	return dara.Validate(s)
}
