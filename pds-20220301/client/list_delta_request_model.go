// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeltaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *ListDeltaRequest
	GetCursor() *string
	SetDriveId(v string) *ListDeltaRequest
	GetDriveId() *string
	SetLimit(v int32) *ListDeltaRequest
	GetLimit() *int32
	SetSyncRootId(v string) *ListDeltaRequest
	GetSyncRootId() *string
}

type ListDeltaRequest struct {
	// The cursor of the incremental information.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The maximum number of results to return. Valid values: 0 to 100. Default value: 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The ID of the root file of the synced folder.
	//
	// example:
	//
	// 622fb09598ae66777c7040109a16f49381f6abe1
	SyncRootId *string `json:"sync_root_id,omitempty" xml:"sync_root_id,omitempty"`
}

func (s ListDeltaRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeltaRequest) GoString() string {
	return s.String()
}

func (s *ListDeltaRequest) GetCursor() *string {
	return s.Cursor
}

func (s *ListDeltaRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListDeltaRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListDeltaRequest) GetSyncRootId() *string {
	return s.SyncRootId
}

func (s *ListDeltaRequest) SetCursor(v string) *ListDeltaRequest {
	s.Cursor = &v
	return s
}

func (s *ListDeltaRequest) SetDriveId(v string) *ListDeltaRequest {
	s.DriveId = &v
	return s
}

func (s *ListDeltaRequest) SetLimit(v int32) *ListDeltaRequest {
	s.Limit = &v
	return s
}

func (s *ListDeltaRequest) SetSyncRootId(v string) *ListDeltaRequest {
	s.SyncRootId = &v
	return s
}

func (s *ListDeltaRequest) Validate() error {
	return dara.Validate(s)
}
