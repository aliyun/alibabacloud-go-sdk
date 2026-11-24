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
	// The status code that indicates whether the ECR instance was deleted. A value of 200 indicates that the deletion was successful. Any other value indicates that the deletion failed. For the cause of the failure, see the error codes.
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
	// The dynamic error message that is used to replace `%s` in the error message of the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
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
	// The ID of the request.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the ECR instance was deleted. Valid values:
	//
	// - **True**: succeeded.
	//
	// - **False**: failed.
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
