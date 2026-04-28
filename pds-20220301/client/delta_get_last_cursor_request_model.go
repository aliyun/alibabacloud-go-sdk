// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeltaGetLastCursorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeltaGetLastCursorRequest
	GetDriveId() *string
	SetSyncRootId(v string) *DeltaGetLastCursorRequest
	GetSyncRootId() *string
}

type DeltaGetLastCursorRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the root file of the synced folder.
	//
	// example:
	//
	// 622fb09598ae66777c7040109a16f49381f6abe1
	SyncRootId *string `json:"sync_root_id,omitempty" xml:"sync_root_id,omitempty"`
}

func (s DeltaGetLastCursorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeltaGetLastCursorRequest) GoString() string {
	return s.String()
}

func (s *DeltaGetLastCursorRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeltaGetLastCursorRequest) GetSyncRootId() *string {
	return s.SyncRootId
}

func (s *DeltaGetLastCursorRequest) SetDriveId(v string) *DeltaGetLastCursorRequest {
	s.DriveId = &v
	return s
}

func (s *DeltaGetLastCursorRequest) SetSyncRootId(v string) *DeltaGetLastCursorRequest {
	s.SyncRootId = &v
	return s
}

func (s *DeltaGetLastCursorRequest) Validate() error {
	return dara.Validate(s)
}
