// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubtaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetSubtaskResponseBody
	GetCode() *int32
	SetDetails(v string) *GetSubtaskResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetSubtaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetSubtaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubtaskResponseBody
	GetRequestId() *string
	SetSubtask(v *SimpleSubtask) *GetSubtaskResponseBody
	GetSubtask() *SimpleSubtask
	SetSuccess(v bool) *GetSubtaskResponseBody
	GetSuccess() *bool
}

type GetSubtaskResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Subtask.
	Subtask *SimpleSubtask `json:"Subtask,omitempty" xml:"Subtask,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubtaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubtaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubtaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSubtaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetSubtaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSubtaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubtaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubtaskResponseBody) GetSubtask() *SimpleSubtask {
	return s.Subtask
}

func (s *GetSubtaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubtaskResponseBody) SetCode(v int32) *GetSubtaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubtaskResponseBody) SetDetails(v string) *GetSubtaskResponseBody {
	s.Details = &v
	return s
}

func (s *GetSubtaskResponseBody) SetErrorCode(v string) *GetSubtaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSubtaskResponseBody) SetMessage(v string) *GetSubtaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubtaskResponseBody) SetRequestId(v string) *GetSubtaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubtaskResponseBody) SetSubtask(v *SimpleSubtask) *GetSubtaskResponseBody {
	s.Subtask = v
	return s
}

func (s *GetSubtaskResponseBody) SetSuccess(v bool) *GetSubtaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubtaskResponseBody) Validate() error {
	if s.Subtask != nil {
		if err := s.Subtask.Validate(); err != nil {
			return err
		}
	}
	return nil
}
