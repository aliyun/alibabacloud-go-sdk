// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectRouterAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetAssociationId(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetAssociationId() *string
	SetCode(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateExpressConnectRouterAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExpressConnectRouterAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExpressConnectRouterAssociationResponseBody
	GetSuccess() *bool
}

type CreateExpressConnectRouterAssociationResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID used to associate the Express Connect Router (ECR) with a VPC or TR.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// Indicates whether the call was successful. A return value of 200 indicates that the attachment was successful. Any other value indicates that the attachment failed. For more information about the failure causes, see Error codes.
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
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the **DtsJobId*	- request parameter is invalid.
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
	// The message returned by the operation.
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
	// Indicates whether the call was successful. Valid values:
	//
	// - **True**: The call was successful.
	//
	// - **False**: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressConnectRouterAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetAssociationId() *string {
	return s.AssociationId
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressConnectRouterAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetAssociationId(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.AssociationId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetCode(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *CreateExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetMessage(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *CreateExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
