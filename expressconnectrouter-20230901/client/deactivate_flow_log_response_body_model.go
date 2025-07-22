// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeactivateFlowLogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeactivateFlowLogResponseBody
	GetCode() *string
	SetDynamicCode(v string) *DeactivateFlowLogResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeactivateFlowLogResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeactivateFlowLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeactivateFlowLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeactivateFlowLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeactivateFlowLogResponseBody
	GetSuccess() *bool
}

type DeactivateFlowLogResponseBody struct {
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

func (s DeactivateFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateFlowLogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeactivateFlowLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeactivateFlowLogResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeactivateFlowLogResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeactivateFlowLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeactivateFlowLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeactivateFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateFlowLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeactivateFlowLogResponseBody) SetAccessDeniedDetail(v string) *DeactivateFlowLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetCode(v string) *DeactivateFlowLogResponseBody {
	s.Code = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetDynamicCode(v string) *DeactivateFlowLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetDynamicMessage(v string) *DeactivateFlowLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetHttpStatusCode(v int32) *DeactivateFlowLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetMessage(v string) *DeactivateFlowLogResponseBody {
	s.Message = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetRequestId(v string) *DeactivateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) SetSuccess(v bool) *DeactivateFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *DeactivateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
