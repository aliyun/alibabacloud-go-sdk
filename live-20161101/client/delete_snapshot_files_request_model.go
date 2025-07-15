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
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The timestamps when the snapshots that you want to delete were captured.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1653641526637
	CreateTimestampList []*int64 `json:"CreateTimestampList,omitempty" xml:"CreateTimestampList,omitempty" type:"Repeated"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to also delete the corresponding file in Object Storage Service (OSS) at the same time. Value values:
	//
	// 	- **true**: deletes the corresponding file in OSS.
	//
	// 	- **false**: does not delete the corresponding file in OSS.
	//
	// >  To delete the corresponding file in OSS, you must have the permissions on the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	RemoveFile *bool `json:"RemoveFile,omitempty" xml:"RemoveFile,omitempty"`
	// The name of the live stream.
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
