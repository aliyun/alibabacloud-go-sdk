// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateAt(v string) *GetUploadUrlResponseBody
	GetCreateAt() *string
	SetDomainId(v string) *GetUploadUrlResponseBody
	GetDomainId() *string
	SetDriveId(v string) *GetUploadUrlResponseBody
	GetDriveId() *string
	SetFileId(v string) *GetUploadUrlResponseBody
	GetFileId() *string
	SetPartInfoList(v []*UploadPartInfo) *GetUploadUrlResponseBody
	GetPartInfoList() []*UploadPartInfo
	SetUploadId(v string) *GetUploadUrlResponseBody
	GetUploadId() *string
}

type GetUploadUrlResponseBody struct {
	// The time when the upload task was created.
	//
	// example:
	//
	// 2019-09-11T16:34:36.977Z
	CreateAt *string `json:"create_at,omitempty" xml:"create_at,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 5d5b846942cf94fa72324c14a4bda34e81da635d
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The information about the file parts.
	PartInfoList []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// The ID of the upload task.
	//
	// example:
	//
	// 10166D06127B413BA1EC8ABB1144D111
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s GetUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponseBody) GetCreateAt() *string {
	return s.CreateAt
}

func (s *GetUploadUrlResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *GetUploadUrlResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *GetUploadUrlResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *GetUploadUrlResponseBody) GetPartInfoList() []*UploadPartInfo {
	return s.PartInfoList
}

func (s *GetUploadUrlResponseBody) GetUploadId() *string {
	return s.UploadId
}

func (s *GetUploadUrlResponseBody) SetCreateAt(v string) *GetUploadUrlResponseBody {
	s.CreateAt = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetDomainId(v string) *GetUploadUrlResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetDriveId(v string) *GetUploadUrlResponseBody {
	s.DriveId = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetFileId(v string) *GetUploadUrlResponseBody {
	s.FileId = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetPartInfoList(v []*UploadPartInfo) *GetUploadUrlResponseBody {
	s.PartInfoList = v
	return s
}

func (s *GetUploadUrlResponseBody) SetUploadId(v string) *GetUploadUrlResponseBody {
	s.UploadId = &v
	return s
}

func (s *GetUploadUrlResponseBody) Validate() error {
	if s.PartInfoList != nil {
		for _, item := range s.PartInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
