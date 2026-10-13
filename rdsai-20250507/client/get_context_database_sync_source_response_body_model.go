// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncSourceResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncSourceResponseBodyData) *GetContextDatabaseSyncSourceResponseBody
	GetData() *GetContextDatabaseSyncSourceResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncSourceResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncSourceResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncSourceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Example value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetContextDatabaseSyncSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncSourceResponseBody) GetData() *GetContextDatabaseSyncSourceResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncSourceResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncSourceResponseBody) SetCode(v string) *GetContextDatabaseSyncSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBody) SetData(v *GetContextDatabaseSyncSourceResponseBodyData) *GetContextDatabaseSyncSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBody) SetMessage(v string) *GetContextDatabaseSyncSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBody) SetRequestId(v string) *GetContextDatabaseSyncSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncSourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncSourceResponseBodyData struct {
	// The confirmation time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	ConfirmedAt *string `json:"ConfirmedAt,omitempty" xml:"ConfirmedAt,omitempty"`
	// The connection ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The data source ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// R&D team document scope
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of selected node keys.
	//
	// example:
	//
	// ["doc:1234567890"]
	SelectedNodeKeys []*string `json:"SelectedNodeKeys,omitempty" xml:"SelectedNodeKeys,omitempty" type:"Repeated"`
	// The selection policy.
	//
	// example:
	//
	// subtree
	SelectionPolicy *string `json:"SelectionPolicy,omitempty" xml:"SelectionPolicy,omitempty"`
	// The datasource config, which varies by plug-in.
	//
	// example:
	//
	// {"folderIds":["https://alidocs.dingtalk.com/i/p/p2N9kLmQxYz****"],"pageSize":50,"maxDepth":5}
	SourceConfig *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
}

func (s GetContextDatabaseSyncSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetConfirmedAt() *string {
	return s.ConfirmedAt
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetSelectedNodeKeys() []*string {
	return s.SelectedNodeKeys
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetSelectionPolicy() *string {
	return s.SelectionPolicy
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) GetSourceConfig() *string {
	return s.SourceConfig
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetConfirmedAt(v string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.ConfirmedAt = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetConnectionId(v string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetId(v string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetName(v string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetSelectedNodeKeys(v []*string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.SelectedNodeKeys = v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetSelectionPolicy(v string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.SelectionPolicy = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) SetSourceConfig(v string) *GetContextDatabaseSyncSourceResponseBodyData {
	s.SourceConfig = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
