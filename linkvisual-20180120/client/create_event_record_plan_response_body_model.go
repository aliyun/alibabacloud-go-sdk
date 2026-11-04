// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRecordPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEventRecordPlanResponseBody
	GetCode() *string
	SetData(v string) *CreateEventRecordPlanResponseBody
	GetData() *string
	SetErrorMessage(v string) *CreateEventRecordPlanResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateEventRecordPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEventRecordPlanResponseBody
	GetSuccess() *bool
}

type CreateEventRecordPlanResponseBody struct {
	// The return code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: an error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the recording plan.
	//
	// example:
	//
	// fc82774fa749485bad7d719f9670****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// product is not active.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a195f615-8c64-4cc6-af57-2f42c6cca15c
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

func (s CreateEventRecordPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRecordPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEventRecordPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateEventRecordPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateEventRecordPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventRecordPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEventRecordPlanResponseBody) SetCode(v string) *CreateEventRecordPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetData(v string) *CreateEventRecordPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetErrorMessage(v string) *CreateEventRecordPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetRequestId(v string) *CreateEventRecordPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) SetSuccess(v bool) *CreateEventRecordPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEventRecordPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
