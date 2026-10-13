// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncSourcesResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncSourcesResponseBodyData) *ListContextDatabaseSyncSourcesResponseBody
	GetData() *ListContextDatabaseSyncSourcesResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncSourcesResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncSourcesResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncSourcesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncSourcesResponseBody) GetData() *ListContextDatabaseSyncSourcesResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncSourcesResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncSourcesResponseBody) SetCode(v string) *ListContextDatabaseSyncSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBody) SetData(v *ListContextDatabaseSyncSourcesResponseBodyData) *ListContextDatabaseSyncSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncSourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBody) SetMessage(v string) *ListContextDatabaseSyncSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBody) SetRequestId(v string) *ListContextDatabaseSyncSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncSourcesResponseBodyData struct {
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of data selection scopes.
	//
	// example:
	//
	// [{"Id":"src-as1llqmkol****", "ConnectionId":"conn-as1llqmkol****", "Name":"R&D team document scope", "SourceConfig":"{\\"folderIds\\":[\\"https://alidocs.dingtalk.com/i/p/p2N9kLmQxYz****\\"],\\"pageSize\\":50,\\"maxDepth\\":5}"}]
	Sources []*ListContextDatabaseSyncSourcesResponseBodyDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContextDatabaseSyncSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) GetSources() []*ListContextDatabaseSyncSourcesResponseBodyDataSources {
	return s.Sources
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncSourcesResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncSourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) SetSources(v []*ListContextDatabaseSyncSourcesResponseBodyDataSources) *ListContextDatabaseSyncSourcesResponseBodyData {
	s.Sources = v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncSourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyData) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncSourcesResponseBodyDataSources struct {
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
	// The list of selected node identifiers.
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
	// The datasource config. The structure varies based on the plugin.
	//
	// example:
	//
	// {"folderIds":["https://alidocs.dingtalk.com/i/p/p2N9kLmQxYz****"],"pageSize":50,"maxDepth":5}
	SourceConfig *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
}

func (s ListContextDatabaseSyncSourcesResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetConfirmedAt() *string {
	return s.ConfirmedAt
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetId() *string {
	return s.Id
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetName() *string {
	return s.Name
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetSelectedNodeKeys() []*string {
	return s.SelectedNodeKeys
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetSelectionPolicy() *string {
	return s.SelectionPolicy
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) GetSourceConfig() *string {
	return s.SourceConfig
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetConfirmedAt(v string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.ConfirmedAt = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetConnectionId(v string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.ConnectionId = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetId(v string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.Id = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetName(v string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.Name = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetSelectedNodeKeys(v []*string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.SelectedNodeKeys = v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetSelectionPolicy(v string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.SelectionPolicy = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) SetSourceConfig(v string) *ListContextDatabaseSyncSourcesResponseBodyDataSources {
	s.SourceConfig = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponseBodyDataSources) Validate() error {
	return dara.Validate(s)
}
