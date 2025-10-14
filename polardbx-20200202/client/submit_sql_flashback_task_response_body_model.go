// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlFlashbackTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitSqlFlashbackTaskResponseBodyData) *SubmitSqlFlashbackTaskResponseBody
	GetData() *SubmitSqlFlashbackTaskResponseBodyData
	SetMessage(v string) *SubmitSqlFlashbackTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitSqlFlashbackTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSqlFlashbackTaskResponseBody
	GetSuccess() *bool
}

type SubmitSqlFlashbackTaskResponseBody struct {
	Data *SubmitSqlFlashbackTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetData() *SubmitSqlFlashbackTaskResponseBodyData {
	return s.Data
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSqlFlashbackTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetData(v *SubmitSqlFlashbackTaskResponseBodyData) *SubmitSqlFlashbackTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetMessage(v string) *SubmitSqlFlashbackTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetRequestId(v string) *SubmitSqlFlashbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) SetSuccess(v bool) *SubmitSqlFlashbackTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSqlFlashbackTaskResponseBodyData struct {
	// example:
	//
	// 1111
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSqlFlashbackTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlFlashbackTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSqlFlashbackTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitSqlFlashbackTaskResponseBodyData) SetTaskId(v string) *SubmitSqlFlashbackTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitSqlFlashbackTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
