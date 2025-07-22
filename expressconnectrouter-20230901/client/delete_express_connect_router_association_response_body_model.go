// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRouterAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteExpressConnectRouterAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DeleteExpressConnectRouterAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExpressConnectRouterAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExpressConnectRouterAssociationResponseBody
	GetSuccess() *bool
}

type DeleteExpressConnectRouterAssociationResponseBody struct {
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

func (s DeleteExpressConnectRouterAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetCode(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
