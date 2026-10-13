// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartContextDatabaseSyncSourceDiscoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBody
	GetCode() *string
	SetData(v *StartContextDatabaseSyncSourceDiscoveryResponseBodyData) *StartContextDatabaseSyncSourceDiscoveryResponseBody
	GetData() *StartContextDatabaseSyncSourceDiscoveryResponseBodyData
	SetHttpStatusCode(v int64) *StartContextDatabaseSyncSourceDiscoveryResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartContextDatabaseSyncSourceDiscoveryResponseBody
	GetSuccess() *bool
}

type StartContextDatabaseSyncSourceDiscoveryResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *StartContextDatabaseSyncSourceDiscoveryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s StartContextDatabaseSyncSourceDiscoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartContextDatabaseSyncSourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) GetData() *StartContextDatabaseSyncSourceDiscoveryResponseBodyData {
	return s.Data
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) SetCode(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	s.Code = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) SetData(v *StartContextDatabaseSyncSourceDiscoveryResponseBodyData) *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	s.Data = v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) SetHttpStatusCode(v int64) *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) SetMessage(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	s.Message = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) SetRequestId(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) SetSuccess(v bool) *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	s.Success = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartContextDatabaseSyncSourceDiscoveryResponseBodyData struct {
	// The discovery task ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartContextDatabaseSyncSourceDiscoveryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartContextDatabaseSyncSourceDiscoveryResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBodyData) SetTaskId(v string) *StartContextDatabaseSyncSourceDiscoveryResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
