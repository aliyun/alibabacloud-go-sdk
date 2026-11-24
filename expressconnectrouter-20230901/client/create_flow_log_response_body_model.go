// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateFlowLogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateFlowLogResponseBody
	GetCode() *string
	SetDynamicCode(v string) *CreateFlowLogResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateFlowLogResponseBody
	GetDynamicMessage() *string
	SetFlowLogId(v string) *CreateFlowLogResponseBody
	GetFlowLogId() *string
	SetHttpStatusCode(v int32) *CreateFlowLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateFlowLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFlowLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFlowLogResponseBody
	GetSuccess() *bool
}

type CreateFlowLogResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Authentication is failed for ****
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code that indicates whether the call was successful. A value of 200 indicates that the request was successful. Other values indicate that the request failed. For more information about the causes of failure, see the error codes.
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
	// The dynamic error message, which is used to replace the **%s*	- placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the specified **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The param format of EcrId ***	- is illegal.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The ID of the flow log instance.
	//
	// example:
	//
	// ecr-flog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
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

func (s CreateFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowLogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateFlowLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFlowLogResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateFlowLogResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateFlowLogResponseBody) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *CreateFlowLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateFlowLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlowLogResponseBody) SetAccessDeniedDetail(v string) *CreateFlowLogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetCode(v string) *CreateFlowLogResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetDynamicCode(v string) *CreateFlowLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetDynamicMessage(v string) *CreateFlowLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetFlowLogId(v string) *CreateFlowLogResponseBody {
	s.FlowLogId = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetHttpStatusCode(v int32) *CreateFlowLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetMessage(v string) *CreateFlowLogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetRequestId(v string) *CreateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetSuccess(v bool) *CreateFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
