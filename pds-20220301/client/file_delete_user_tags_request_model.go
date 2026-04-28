// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileDeleteUserTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *FileDeleteUserTagsRequest
	GetDriveId() *string
	SetFileId(v string) *FileDeleteUserTagsRequest
	GetFileId() *string
	SetKeyList(v []*string) *FileDeleteUserTagsRequest
	GetKeyList() []*string
}

type FileDeleteUserTagsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The tags that you want to remove from a file. You cannot leave this parameter empty. You can specify up to 1,000 tags.
	//
	// This parameter is required.
	KeyList []*string `json:"key_list,omitempty" xml:"key_list,omitempty" type:"Repeated"`
}

func (s FileDeleteUserTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s FileDeleteUserTagsRequest) GoString() string {
	return s.String()
}

func (s *FileDeleteUserTagsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *FileDeleteUserTagsRequest) GetFileId() *string {
	return s.FileId
}

func (s *FileDeleteUserTagsRequest) GetKeyList() []*string {
	return s.KeyList
}

func (s *FileDeleteUserTagsRequest) SetDriveId(v string) *FileDeleteUserTagsRequest {
	s.DriveId = &v
	return s
}

func (s *FileDeleteUserTagsRequest) SetFileId(v string) *FileDeleteUserTagsRequest {
	s.FileId = &v
	return s
}

func (s *FileDeleteUserTagsRequest) SetKeyList(v []*string) *FileDeleteUserTagsRequest {
	s.KeyList = v
	return s
}

func (s *FileDeleteUserTagsRequest) Validate() error {
	return dara.Validate(s)
}
