// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryFailedContextDatabaseSyncRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryFailedContextDatabaseSyncRunResponseBody
	GetCode() *string
	SetData(v *RetryFailedContextDatabaseSyncRunResponseBodyData) *RetryFailedContextDatabaseSyncRunResponseBody
	GetData() *RetryFailedContextDatabaseSyncRunResponseBodyData
	SetHttpStatusCode(v int64) *RetryFailedContextDatabaseSyncRunResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *RetryFailedContextDatabaseSyncRunResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryFailedContextDatabaseSyncRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryFailedContextDatabaseSyncRunResponseBody
	GetSuccess() *bool
}

type RetryFailedContextDatabaseSyncRunResponseBody struct {
	// example:
	//
	// 示例值
	Code *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RetryFailedContextDatabaseSyncRunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryFailedContextDatabaseSyncRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryFailedContextDatabaseSyncRunResponseBody) GoString() string {
	return s.String()
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) GetData() *RetryFailedContextDatabaseSyncRunResponseBodyData {
	return s.Data
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) SetCode(v string) *RetryFailedContextDatabaseSyncRunResponseBody {
	s.Code = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) SetData(v *RetryFailedContextDatabaseSyncRunResponseBodyData) *RetryFailedContextDatabaseSyncRunResponseBody {
	s.Data = v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) SetHttpStatusCode(v int64) *RetryFailedContextDatabaseSyncRunResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) SetMessage(v string) *RetryFailedContextDatabaseSyncRunResponseBody {
	s.Message = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) SetRequestId(v string) *RetryFailedContextDatabaseSyncRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) SetSuccess(v bool) *RetryFailedContextDatabaseSyncRunResponseBody {
	s.Success = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryFailedContextDatabaseSyncRunResponseBodyData struct {
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RetryOfRunId *string `json:"RetryOfRunId,omitempty" xml:"RetryOfRunId,omitempty"`
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
}

func (s RetryFailedContextDatabaseSyncRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetryFailedContextDatabaseSyncRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetryFailedContextDatabaseSyncRunResponseBodyData) GetRetryOfRunId() *string {
	return s.RetryOfRunId
}

func (s *RetryFailedContextDatabaseSyncRunResponseBodyData) GetRunId() *string {
	return s.RunId
}

func (s *RetryFailedContextDatabaseSyncRunResponseBodyData) SetRetryOfRunId(v string) *RetryFailedContextDatabaseSyncRunResponseBodyData {
	s.RetryOfRunId = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBodyData) SetRunId(v string) *RetryFailedContextDatabaseSyncRunResponseBodyData {
	s.RunId = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
