// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateExpressConnectRouterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateExpressConnectRouterResponseBody
	GetCode() *string
	SetDynamicCode(v string) *CreateExpressConnectRouterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateExpressConnectRouterResponseBody
	GetDynamicMessage() *string
	SetEcrId(v string) *CreateExpressConnectRouterResponseBody
	GetEcrId() *string
	SetHttpStatusCode(v int32) *CreateExpressConnectRouterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateExpressConnectRouterResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExpressConnectRouterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExpressConnectRouterResponseBody
	GetSuccess() *bool
}

type CreateExpressConnectRouterResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The code that indicates whether the ECR instance was created. A value of 200 indicates that the instance was created. Other values indicate that the creation failed. For more information about the failure causes, see the error codes.
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
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The param format of Name ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The Express Connect Router (ECR) instance ID.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
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
	// Indicates whether the ECR instance was created. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressConnectRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateExpressConnectRouterResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateExpressConnectRouterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateExpressConnectRouterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateExpressConnectRouterResponseBody) GetEcrId() *string {
	return s.EcrId
}

func (s *CreateExpressConnectRouterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateExpressConnectRouterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExpressConnectRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressConnectRouterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *CreateExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetCode(v string) *CreateExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetDynamicCode(v string) *CreateExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetDynamicMessage(v string) *CreateExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetEcrId(v string) *CreateExpressConnectRouterResponseBody {
	s.EcrId = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *CreateExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetMessage(v string) *CreateExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetRequestId(v string) *CreateExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetSuccess(v bool) *CreateExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
