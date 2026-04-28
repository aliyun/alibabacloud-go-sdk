// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RefreshAdvisorCheckResponseBodyData) *RefreshAdvisorCheckResponseBody
	GetData() *RefreshAdvisorCheckResponseBodyData
	SetRequestId(v string) *RefreshAdvisorCheckResponseBody
	GetRequestId() *string
}

type RefreshAdvisorCheckResponseBody struct {
	Data *RefreshAdvisorCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshAdvisorCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCheckResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponseBody) GetData() *RefreshAdvisorCheckResponseBodyData {
	return s.Data
}

func (s *RefreshAdvisorCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshAdvisorCheckResponseBody) SetData(v *RefreshAdvisorCheckResponseBodyData) *RefreshAdvisorCheckResponseBody {
	s.Data = v
	return s
}

func (s *RefreshAdvisorCheckResponseBody) SetRequestId(v string) *RefreshAdvisorCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefreshAdvisorCheckResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 12345678
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// manual-1be17af1121b4974822e69daee4f2481
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RefreshAdvisorCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RefreshAdvisorCheckResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RefreshAdvisorCheckResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *RefreshAdvisorCheckResponseBodyData) GetTraceId() *string {
	return s.TraceId
}

func (s *RefreshAdvisorCheckResponseBodyData) SetMessage(v string) *RefreshAdvisorCheckResponseBodyData {
	s.Message = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBodyData) SetSuccess(v bool) *RefreshAdvisorCheckResponseBodyData {
	s.Success = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBodyData) SetTaskId(v int64) *RefreshAdvisorCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBodyData) SetTraceId(v string) *RefreshAdvisorCheckResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *RefreshAdvisorCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
