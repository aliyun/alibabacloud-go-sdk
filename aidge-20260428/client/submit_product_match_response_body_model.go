// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProductMatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitProductMatchResponseBody
	GetCode() *string
	SetData(v *SubmitProductMatchResponseBodyData) *SubmitProductMatchResponseBody
	GetData() *SubmitProductMatchResponseBodyData
	SetMessage(v string) *SubmitProductMatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitProductMatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitProductMatchResponseBody
	GetSuccess() *bool
}

type SubmitProductMatchResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submit result of the matching product identification asynchronous task.
	Data *SubmitProductMatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitProductMatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitProductMatchResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitProductMatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitProductMatchResponseBody) GetData() *SubmitProductMatchResponseBodyData {
	return s.Data
}

func (s *SubmitProductMatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitProductMatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitProductMatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitProductMatchResponseBody) SetCode(v string) *SubmitProductMatchResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitProductMatchResponseBody) SetData(v *SubmitProductMatchResponseBodyData) *SubmitProductMatchResponseBody {
	s.Data = v
	return s
}

func (s *SubmitProductMatchResponseBody) SetMessage(v string) *SubmitProductMatchResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitProductMatchResponseBody) SetRequestId(v string) *SubmitProductMatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitProductMatchResponseBody) SetSuccess(v bool) *SubmitProductMatchResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitProductMatchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitProductMatchResponseBodyData struct {
	// The task acceptance time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-20T09:30:00Z
	SubmittedAt *string `json:"SubmittedAt,omitempty" xml:"SubmittedAt,omitempty"`
	// The asynchronous task ID used for QueryAsyncTaskResult queries.
	//
	// example:
	//
	// b7ea15cb609f47b7999d2d68dfbf3c90
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitProductMatchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitProductMatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitProductMatchResponseBodyData) GetSubmittedAt() *string {
	return s.SubmittedAt
}

func (s *SubmitProductMatchResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitProductMatchResponseBodyData) SetSubmittedAt(v string) *SubmitProductMatchResponseBodyData {
	s.SubmittedAt = &v
	return s
}

func (s *SubmitProductMatchResponseBodyData) SetTaskId(v string) *SubmitProductMatchResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitProductMatchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
