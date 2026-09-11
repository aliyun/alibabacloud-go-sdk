// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteToMasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *PromoteToMasterResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PromoteToMasterResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *PromoteToMasterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PromoteToMasterResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *PromoteToMasterResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *PromoteToMasterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PromoteToMasterResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *PromoteToMasterResponseBody
	GetTaskId() *string
}

type PromoteToMasterResponseBody struct {
	// Dynamic error code. This parameter will be unpublished soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the **%s*	- placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Error code returned when the invocation fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message returned when the invocation fails.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 210ec2e116055198849072222d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Job ID. This parameter will be unpublished soon.
	//
	// example:
	//
	// z2v12jfo309****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PromoteToMasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PromoteToMasterResponseBody) GoString() string {
	return s.String()
}

func (s *PromoteToMasterResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PromoteToMasterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PromoteToMasterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PromoteToMasterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PromoteToMasterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PromoteToMasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PromoteToMasterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PromoteToMasterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PromoteToMasterResponseBody) SetDynamicCode(v string) *PromoteToMasterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetDynamicMessage(v string) *PromoteToMasterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetErrCode(v string) *PromoteToMasterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetErrMessage(v string) *PromoteToMasterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetHttpStatusCode(v int32) *PromoteToMasterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetRequestId(v string) *PromoteToMasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetSuccess(v bool) *PromoteToMasterResponseBody {
	s.Success = &v
	return s
}

func (s *PromoteToMasterResponseBody) SetTaskId(v string) *PromoteToMasterResponseBody {
	s.TaskId = &v
	return s
}

func (s *PromoteToMasterResponseBody) Validate() error {
	return dara.Validate(s)
}
