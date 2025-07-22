// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ActivateFlowLogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ActivateFlowLogResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ActivateFlowLogResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ActivateFlowLogResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ActivateFlowLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ActivateFlowLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActivateFlowLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ActivateFlowLogResponseBody
	GetSuccess() *bool
}

type ActivateFlowLogResponseBody struct {
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

func (s ActivateFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateFlowLogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ActivateFlowLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActivateFlowLogResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ActivateFlowLogResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ActivateFlowLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ActivateFlowLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActivateFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateFlowLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ActivateFlowLogResponseBody) SetAccessDeniedDetail(v string) *ActivateFlowLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetCode(v string) *ActivateFlowLogResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetDynamicCode(v string) *ActivateFlowLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetDynamicMessage(v string) *ActivateFlowLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetHttpStatusCode(v int32) *ActivateFlowLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetMessage(v string) *ActivateFlowLogResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetRequestId(v string) *ActivateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateFlowLogResponseBody) SetSuccess(v bool) *ActivateFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *ActivateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
