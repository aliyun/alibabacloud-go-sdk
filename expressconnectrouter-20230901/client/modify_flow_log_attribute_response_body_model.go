// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyFlowLogAttributeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ModifyFlowLogAttributeResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ModifyFlowLogAttributeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyFlowLogAttributeResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ModifyFlowLogAttributeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyFlowLogAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyFlowLogAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyFlowLogAttributeResponseBody
	GetSuccess() *bool
}

type ModifyFlowLogAttributeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The code that indicates whether the call was successful. A value of 200 indicates that the request was successful. Any other value indicates that the request failed. For the causes of failure, see the error codes.
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
	// The dynamic error message, which is used to replace `%s` in the error message of the **ErrMessage*	- response parameter.
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
	// The message returned for the instance operation.
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

func (s ModifyFlowLogAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyFlowLogAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyFlowLogAttributeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyFlowLogAttributeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyFlowLogAttributeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyFlowLogAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyFlowLogAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFlowLogAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyFlowLogAttributeResponseBody) SetAccessDeniedDetail(v string) *ModifyFlowLogAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetCode(v string) *ModifyFlowLogAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetDynamicCode(v string) *ModifyFlowLogAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetDynamicMessage(v string) *ModifyFlowLogAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetHttpStatusCode(v int32) *ModifyFlowLogAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetMessage(v string) *ModifyFlowLogAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetRequestId(v string) *ModifyFlowLogAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetSuccess(v bool) *ModifyFlowLogAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
