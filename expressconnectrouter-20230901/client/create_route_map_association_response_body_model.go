// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateRouteMapAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetAssociationId(v string) *CreateRouteMapAssociationResponseBody
	GetAssociationId() *string
	SetCode(v string) *CreateRouteMapAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *CreateRouteMapAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateRouteMapAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateRouteMapAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateRouteMapAssociationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRouteMapAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRouteMapAssociationResponseBody
	GetSuccess() *bool
}

type CreateRouteMapAssociationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the route map association.
	//
	// example:
	//
	// ecr-routemap-assoc-qdoeneegffq4yv****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The status code that indicates whether the call was successful. A value of 200 indicates success. Other values indicate failure. For more information about the causes of failure, see error codes.
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
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- response parameter.
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
	// The message returned for the operation.
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRouteMapAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteMapAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateRouteMapAssociationResponseBody) GetAssociationId() *string {
	return s.AssociationId
}

func (s *CreateRouteMapAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRouteMapAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateRouteMapAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateRouteMapAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRouteMapAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRouteMapAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteMapAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRouteMapAssociationResponseBody) SetAccessDeniedDetail(v string) *CreateRouteMapAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetAssociationId(v string) *CreateRouteMapAssociationResponseBody {
	s.AssociationId = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetCode(v string) *CreateRouteMapAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetDynamicCode(v string) *CreateRouteMapAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetDynamicMessage(v string) *CreateRouteMapAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetHttpStatusCode(v int32) *CreateRouteMapAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetMessage(v string) *CreateRouteMapAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetRequestId(v string) *CreateRouteMapAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) SetSuccess(v bool) *CreateRouteMapAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRouteMapAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
