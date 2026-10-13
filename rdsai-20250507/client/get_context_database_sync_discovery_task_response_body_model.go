// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncDiscoveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) *GetContextDatabaseSyncDiscoveryTaskResponseBody
	GetData() *GetContextDatabaseSyncDiscoveryTaskResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncDiscoveryTaskResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncDiscoveryTaskResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncDiscoveryTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncDiscoveryTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
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

func (s GetContextDatabaseSyncDiscoveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncDiscoveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) GetData() *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) SetCode(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) SetData(v *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) SetMessage(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) SetRequestId(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncDiscoveryTaskResponseBodyData struct {
	// The time when the cancellation was requested.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	CancelRequestedAt *string `json:"CancelRequestedAt,omitempty" xml:"CancelRequestedAt,omitempty"`
	// The connection ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The number of containers.
	//
	// example:
	//
	// 1000000000000000
	Containers *int64 `json:"Containers,omitempty" xml:"Containers,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The error message.
	//
	// example:
	//
	// File size exceeds the limit
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The plug-in reference.
	//
	// example:
	//
	// dingtalk@1.0.0
	PluginRef *string `json:"PluginRef,omitempty" xml:"PluginRef,omitempty"`
	// The number of scanned items.
	//
	// example:
	//
	// 1000000000000000
	Scanned *int64 `json:"Scanned,omitempty" xml:"Scanned,omitempty"`
	// The datasource config. The structure varies depending on the plug-in.
	//
	// example:
	//
	// {"folderIds":["https://alidocs.dingtalk.com/i/p/p2N9kLmQxYz****"],"pageSize":50,"maxDepth":5}
	SourceConfig *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
	// The data source ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The discovery task status. Valid values: queued, running, succeeded, cancelled, and failed.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The discovery task ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The time when the task was last updated.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetContextDatabaseSyncDiscoveryTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetCancelRequestedAt() *string {
	return s.CancelRequestedAt
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetContainers() *int64 {
	return s.Containers
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetPluginRef() *string {
	return s.PluginRef
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetScanned() *int64 {
	return s.Scanned
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetSourceConfig() *string {
	return s.SourceConfig
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetCancelRequestedAt(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.CancelRequestedAt = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetConnectionId(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetContainers(v int64) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.Containers = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetCreatedAt(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetError(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetPluginRef(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.PluginRef = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetScanned(v int64) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.Scanned = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetSourceConfig(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.SourceConfig = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetSourceId(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetState(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.State = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetTaskId(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) SetUpdatedAt(v string) *GetContextDatabaseSyncDiscoveryTaskResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
