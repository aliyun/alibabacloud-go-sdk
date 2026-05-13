// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceActionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *WorkspaceActionStatusResponseBodyData) *WorkspaceActionStatusResponseBody
	GetData() *WorkspaceActionStatusResponseBodyData
	SetErrorCode(v string) *WorkspaceActionStatusResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *WorkspaceActionStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *WorkspaceActionStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *WorkspaceActionStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WorkspaceActionStatusResponseBody
	GetSuccess() *bool
}

type WorkspaceActionStatusResponseBody struct {
	Data *WorkspaceActionStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// This record is being collected, please wait for a moment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WorkspaceActionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *WorkspaceActionStatusResponseBody) GetData() *WorkspaceActionStatusResponseBodyData {
	return s.Data
}

func (s *WorkspaceActionStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WorkspaceActionStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WorkspaceActionStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WorkspaceActionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WorkspaceActionStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WorkspaceActionStatusResponseBody) SetData(v *WorkspaceActionStatusResponseBodyData) *WorkspaceActionStatusResponseBody {
	s.Data = v
	return s
}

func (s *WorkspaceActionStatusResponseBody) SetErrorCode(v string) *WorkspaceActionStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WorkspaceActionStatusResponseBody) SetHttpStatusCode(v int32) *WorkspaceActionStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WorkspaceActionStatusResponseBody) SetMessage(v string) *WorkspaceActionStatusResponseBody {
	s.Message = &v
	return s
}

func (s *WorkspaceActionStatusResponseBody) SetRequestId(v string) *WorkspaceActionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *WorkspaceActionStatusResponseBody) SetSuccess(v bool) *WorkspaceActionStatusResponseBody {
	s.Success = &v
	return s
}

func (s *WorkspaceActionStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WorkspaceActionStatusResponseBodyData struct {
	// example:
	//
	// action failed
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 2026-01-13T14:30:20.582182728+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2026-01-13T14:30:20.582182728+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"phase\\": \\"Created\\", \\"latestExecError\\": {\\"message\\": \\"\\", \\"code\\": \\"\\", \\"requestId\\": \\"\\", \\"extraInfo\\": \\"\\", \\"title\\": \\"\\"}}
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WorkspaceActionStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceActionStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *WorkspaceActionStatusResponseBodyData) GetData() *string {
	return s.Data
}

func (s *WorkspaceActionStatusResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *WorkspaceActionStatusResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *WorkspaceActionStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *WorkspaceActionStatusResponseBodyData) SetData(v string) *WorkspaceActionStatusResponseBodyData {
	s.Data = &v
	return s
}

func (s *WorkspaceActionStatusResponseBodyData) SetEndTime(v string) *WorkspaceActionStatusResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *WorkspaceActionStatusResponseBodyData) SetStartTime(v string) *WorkspaceActionStatusResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *WorkspaceActionStatusResponseBodyData) SetStatus(v string) *WorkspaceActionStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *WorkspaceActionStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
