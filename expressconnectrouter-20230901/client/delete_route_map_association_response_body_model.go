// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteRouteMapAssociationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteRouteMapAssociationResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DeleteRouteMapAssociationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteRouteMapAssociationResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteRouteMapAssociationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRouteMapAssociationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRouteMapAssociationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRouteMapAssociationResponseBody
	GetSuccess() *bool
}

type DeleteRouteMapAssociationResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether the call was successful. A value of 200 indicates success. Other values indicate failure. For more information about the causes of failure, see error codes.
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
	// 05130E79-588D-5C40-A718-C4863A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **True**: successful.
	//
	// - **False**: failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRouteMapAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapAssociationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteRouteMapAssociationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRouteMapAssociationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteRouteMapAssociationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteRouteMapAssociationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRouteMapAssociationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRouteMapAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteMapAssociationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRouteMapAssociationResponseBody) SetAccessDeniedDetail(v string) *DeleteRouteMapAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetCode(v string) *DeleteRouteMapAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetDynamicCode(v string) *DeleteRouteMapAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetDynamicMessage(v string) *DeleteRouteMapAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetHttpStatusCode(v int32) *DeleteRouteMapAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetMessage(v string) *DeleteRouteMapAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetRequestId(v string) *DeleteRouteMapAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) SetSuccess(v bool) *DeleteRouteMapAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRouteMapAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
