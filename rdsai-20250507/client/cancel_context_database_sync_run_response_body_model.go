// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelContextDatabaseSyncRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelContextDatabaseSyncRunResponseBody
	GetCode() *string
	SetData(v *CancelContextDatabaseSyncRunResponseBodyData) *CancelContextDatabaseSyncRunResponseBody
	GetData() *CancelContextDatabaseSyncRunResponseBodyData
	SetHttpStatusCode(v int64) *CancelContextDatabaseSyncRunResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CancelContextDatabaseSyncRunResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelContextDatabaseSyncRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelContextDatabaseSyncRunResponseBody
	GetSuccess() *bool
}

type CancelContextDatabaseSyncRunResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CancelContextDatabaseSyncRunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Example value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelContextDatabaseSyncRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelContextDatabaseSyncRunResponseBody) GoString() string {
	return s.String()
}

func (s *CancelContextDatabaseSyncRunResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelContextDatabaseSyncRunResponseBody) GetData() *CancelContextDatabaseSyncRunResponseBodyData {
	return s.Data
}

func (s *CancelContextDatabaseSyncRunResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CancelContextDatabaseSyncRunResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelContextDatabaseSyncRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelContextDatabaseSyncRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelContextDatabaseSyncRunResponseBody) SetCode(v string) *CancelContextDatabaseSyncRunResponseBody {
	s.Code = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBody) SetData(v *CancelContextDatabaseSyncRunResponseBodyData) *CancelContextDatabaseSyncRunResponseBody {
	s.Data = v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBody) SetHttpStatusCode(v int64) *CancelContextDatabaseSyncRunResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBody) SetMessage(v string) *CancelContextDatabaseSyncRunResponseBody {
	s.Message = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBody) SetRequestId(v string) *CancelContextDatabaseSyncRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBody) SetSuccess(v bool) *CancelContextDatabaseSyncRunResponseBody {
	s.Success = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelContextDatabaseSyncRunResponseBodyData struct {
	// The run ID, a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The run status after the cancellation request is submitted. For runs in the queued or running state, the operation immediately returns cancelling.
	//
	// example:
	//
	// cancelling
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelContextDatabaseSyncRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelContextDatabaseSyncRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelContextDatabaseSyncRunResponseBodyData) GetRunId() *string {
	return s.RunId
}

func (s *CancelContextDatabaseSyncRunResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CancelContextDatabaseSyncRunResponseBodyData) SetRunId(v string) *CancelContextDatabaseSyncRunResponseBodyData {
	s.RunId = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBodyData) SetStatus(v string) *CancelContextDatabaseSyncRunResponseBodyData {
	s.Status = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
