// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListContextDatabaseSyncConnectionsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListContextDatabaseSyncConnectionsRequest
	GetPageSize() *int32
	SetPluginId(v string) *ListContextDatabaseSyncConnectionsRequest
	GetPluginId() *string
	SetWorkspaceId(v string) *ListContextDatabaseSyncConnectionsRequest
	GetWorkspaceId() *string
}

type ListContextDatabaseSyncConnectionsRequest struct {
	// The page number. Pages start from 1.
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
	// The plug-in ID.
	//
	// example:
	//
	// dingtalk
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListContextDatabaseSyncConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncConnectionsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncConnectionsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListContextDatabaseSyncConnectionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListContextDatabaseSyncConnectionsRequest) SetPage(v int32) *ListContextDatabaseSyncConnectionsRequest {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsRequest) SetPageSize(v int32) *ListContextDatabaseSyncConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsRequest) SetPluginId(v string) *ListContextDatabaseSyncConnectionsRequest {
	s.PluginId = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsRequest) SetWorkspaceId(v string) *ListContextDatabaseSyncConnectionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
