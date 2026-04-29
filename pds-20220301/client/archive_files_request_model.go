// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ArchiveFilesRequest
	GetDriveId() *string
	SetFileIds(v []*string) *ArchiveFilesRequest
	GetFileIds() []*string
	SetName(v string) *ArchiveFilesRequest
	GetName() *string
	SetShareId(v string) *ArchiveFilesRequest
	GetShareId() *string
}

type ArchiveFilesRequest struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	FileIds []*string `json:"file_ids,omitempty" xml:"file_ids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_archive_files.zip
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s ArchiveFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ArchiveFilesRequest) GoString() string {
	return s.String()
}

func (s *ArchiveFilesRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ArchiveFilesRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *ArchiveFilesRequest) GetName() *string {
	return s.Name
}

func (s *ArchiveFilesRequest) GetShareId() *string {
	return s.ShareId
}

func (s *ArchiveFilesRequest) SetDriveId(v string) *ArchiveFilesRequest {
	s.DriveId = &v
	return s
}

func (s *ArchiveFilesRequest) SetFileIds(v []*string) *ArchiveFilesRequest {
	s.FileIds = v
	return s
}

func (s *ArchiveFilesRequest) SetName(v string) *ArchiveFilesRequest {
	s.Name = &v
	return s
}

func (s *ArchiveFilesRequest) SetShareId(v string) *ArchiveFilesRequest {
	s.ShareId = &v
	return s
}

func (s *ArchiveFilesRequest) Validate() error {
	return dara.Validate(s)
}
