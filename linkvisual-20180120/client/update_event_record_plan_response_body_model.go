// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventRecordPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEventRecordPlanResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateEventRecordPlanResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateEventRecordPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventRecordPlanResponseBody
	GetSuccess() *bool
}

type UpdateEventRecordPlanResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: An error occurred. For more information about error codes, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventRecordPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventRecordPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateEventRecordPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventRecordPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventRecordPlanResponseBody) SetCode(v string) *UpdateEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetErrorMessage(v string) *UpdateEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetRequestId(v string) *UpdateEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) SetSuccess(v bool) *UpdateEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventRecordPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
