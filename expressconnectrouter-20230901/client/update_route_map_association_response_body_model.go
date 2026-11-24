// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateRouteMapAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateRouteMapAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *UpdateRouteMapAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateRouteMapAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateRouteMapAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateRouteMapAssociationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRouteMapAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRouteMapAssociationResponseBody
	GetSuccess() *bool
}

type UpdateRouteMapAssociationResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the call was successful. A return value of 200 indicates success. Other values indicate failure. For the cause of failure, see error codes.
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
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
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

func (s UpdateRouteMapAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateRouteMapAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRouteMapAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateRouteMapAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateRouteMapAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRouteMapAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRouteMapAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRouteMapAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRouteMapAssociationResponseBody) SetAccessDeniedDetail(v string) *UpdateRouteMapAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetCode(v string) *UpdateRouteMapAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetDynamicCode(v string) *UpdateRouteMapAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetDynamicMessage(v string) *UpdateRouteMapAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetHttpStatusCode(v int32) *UpdateRouteMapAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetMessage(v string) *UpdateRouteMapAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetRequestId(v string) *UpdateRouteMapAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) SetSuccess(v bool) *UpdateRouteMapAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRouteMapAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
