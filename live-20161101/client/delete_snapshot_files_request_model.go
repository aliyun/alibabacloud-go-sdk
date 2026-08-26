// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteSnapshotFilesRequest
	GetAppName() *string
	SetCreateTimestampList(v []*int64) *DeleteSnapshotFilesRequest
	GetCreateTimestampList() []*int64
	SetDomainName(v string) *DeleteSnapshotFilesRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteSnapshotFilesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSnapshotFilesRequest
	GetRegionId() *string
	SetRemoveFile(v bool) *DeleteSnapshotFilesRequest
	GetRemoveFile() *bool
	SetStreamName(v string) *DeleteSnapshotFilesRequest
	GetStreamName() *string
}

type DeleteSnapshotFilesRequest struct {
	// The AppName of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// A list of timestamps of the snapshots to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1653641526637
	CreateTimestampList []*int64 `json:"CreateTimestampList,omitempty" xml:"CreateTimestampList,omitempty" type:"Repeated"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to also delete the snapshot files from OSS. Valid values:
	//
	// - **true**: Deletes.
	//
	// - **false**: Does not delete.
	//
	// > To delete files from OSS, you must have the required permissions for OSS file operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	RemoveFile *bool `json:"RemoveFile,omitempty" xml:"RemoveFile,omitempty"`
	// The stream name.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteSnapshotFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotFilesRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteSnapshotFilesRequest) GetCreateTimestampList() []*int64 {
	return s.CreateTimestampList
}

func (s *DeleteSnapshotFilesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteSnapshotFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSnapshotFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotFilesRequest) GetRemoveFile() *bool {
	return s.RemoveFile
}

func (s *DeleteSnapshotFilesRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteSnapshotFilesRequest) SetAppName(v string) *DeleteSnapshotFilesRequest {
	s.AppName = &v
	return s
}

func (s *DeleteSnapshotFilesRequest) SetCreateTimestampList(v []*int64) *DeleteSnapshotFilesRequest {
	s.CreateTimestampList = v
	return s
}

func (s *DeleteSnapshotFilesRequest) SetDomainName(v string) *DeleteSnapshotFilesRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSnapshotFilesRequest) SetOwnerId(v int64) *DeleteSnapshotFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSnapshotFilesRequest) SetRegionId(v string) *DeleteSnapshotFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotFilesRequest) SetRemoveFile(v bool) *DeleteSnapshotFilesRequest {
	s.RemoveFile = &v
	return s
}

func (s *DeleteSnapshotFilesRequest) SetStreamName(v string) *DeleteSnapshotFilesRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteSnapshotFilesRequest) Validate() error {
	return dara.Validate(s)
}
