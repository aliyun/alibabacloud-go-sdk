// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowlogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteFlowlogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteFlowlogResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DeleteFlowlogResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteFlowlogResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteFlowlogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteFlowlogResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFlowlogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFlowlogResponseBody
	GetSuccess() *bool
}

type DeleteFlowlogResponseBody struct {
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
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **True**
	//
	// - **False**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowlogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteFlowlogResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFlowlogResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteFlowlogResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteFlowlogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFlowlogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFlowlogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowlogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFlowlogResponseBody) SetAccessDeniedDetail(v string) *DeleteFlowlogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetCode(v string) *DeleteFlowlogResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetDynamicCode(v string) *DeleteFlowlogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetDynamicMessage(v string) *DeleteFlowlogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetHttpStatusCode(v int32) *DeleteFlowlogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetMessage(v string) *DeleteFlowlogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetRequestId(v string) *DeleteFlowlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetSuccess(v bool) *DeleteFlowlogResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowlogResponseBody) Validate() error {
	return dara.Validate(s)
}
