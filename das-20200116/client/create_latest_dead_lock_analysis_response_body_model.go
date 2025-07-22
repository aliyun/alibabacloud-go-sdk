// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLatestDeadLockAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateLatestDeadLockAnalysisResponseBody
	GetCode() *int64
	SetData(v bool) *CreateLatestDeadLockAnalysisResponseBody
	GetData() *bool
	SetMessage(v string) *CreateLatestDeadLockAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLatestDeadLockAnalysisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLatestDeadLockAnalysisResponseBody
	GetSuccess() *bool
}

type CreateLatestDeadLockAnalysisResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLatestDeadLockAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLatestDeadLockAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLatestDeadLockAnalysisResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateLatestDeadLockAnalysisResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateLatestDeadLockAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLatestDeadLockAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLatestDeadLockAnalysisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLatestDeadLockAnalysisResponseBody) SetCode(v int64) *CreateLatestDeadLockAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponseBody) SetData(v bool) *CreateLatestDeadLockAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponseBody) SetMessage(v string) *CreateLatestDeadLockAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponseBody) SetRequestId(v string) *CreateLatestDeadLockAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponseBody) SetSuccess(v bool) *CreateLatestDeadLockAnalysisResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
