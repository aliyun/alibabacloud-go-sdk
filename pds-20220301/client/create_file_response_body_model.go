// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *CreateFileResponseBody
	GetDomainId() *string
	SetDriveId(v string) *CreateFileResponseBody
	GetDriveId() *string
	SetExist(v bool) *CreateFileResponseBody
	GetExist() *bool
	SetFileId(v string) *CreateFileResponseBody
	GetFileId() *string
	SetFileName(v string) *CreateFileResponseBody
	GetFileName() *string
	SetParentFileId(v string) *CreateFileResponseBody
	GetParentFileId() *string
	SetPartInfoList(v []*UploadPartInfo) *CreateFileResponseBody
	GetPartInfoList() []*UploadPartInfo
	SetRapidUpload(v bool) *CreateFileResponseBody
	GetRapidUpload() *bool
	SetStatus(v string) *CreateFileResponseBody
	GetStatus() *string
	SetType(v string) *CreateFileResponseBody
	GetType() *string
	SetUploadId(v string) *CreateFileResponseBody
	GetUploadId() *string
}

type CreateFileResponseBody struct {
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
	// Indicates whether the file exists.
	//
	// example:
	//
	// false
	Exist *bool `json:"exist,omitempty" xml:"exist,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The file name.
	//
	// example:
	//
	// a.txt
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// The ID of the parent directory.
	//
	// example:
	//
	// fileid5
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The information about the file parts.
	PartInfoList []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// Indicates whether the file is instantly uploaded.
	//
	// example:
	//
	// true
	RapidUpload *bool `json:"rapid_upload,omitempty" xml:"rapid_upload,omitempty"`
	// The state of the file.
	//
	// example:
	//
	// uploading
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the file.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the upload task.
	//
	// example:
	//
	// uploadid1
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateFileResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateFileResponseBody) GetExist() *bool {
	return s.Exist
}

func (s *CreateFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *CreateFileResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *CreateFileResponseBody) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *CreateFileResponseBody) GetPartInfoList() []*UploadPartInfo {
	return s.PartInfoList
}

func (s *CreateFileResponseBody) GetRapidUpload() *bool {
	return s.RapidUpload
}

func (s *CreateFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateFileResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateFileResponseBody) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateFileResponseBody) SetDomainId(v string) *CreateFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateFileResponseBody) SetDriveId(v string) *CreateFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateFileResponseBody) SetExist(v bool) *CreateFileResponseBody {
	s.Exist = &v
	return s
}

func (s *CreateFileResponseBody) SetFileId(v string) *CreateFileResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateFileResponseBody) SetFileName(v string) *CreateFileResponseBody {
	s.FileName = &v
	return s
}

func (s *CreateFileResponseBody) SetParentFileId(v string) *CreateFileResponseBody {
	s.ParentFileId = &v
	return s
}

func (s *CreateFileResponseBody) SetPartInfoList(v []*UploadPartInfo) *CreateFileResponseBody {
	s.PartInfoList = v
	return s
}

func (s *CreateFileResponseBody) SetRapidUpload(v bool) *CreateFileResponseBody {
	s.RapidUpload = &v
	return s
}

func (s *CreateFileResponseBody) SetStatus(v string) *CreateFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFileResponseBody) SetType(v string) *CreateFileResponseBody {
	s.Type = &v
	return s
}

func (s *CreateFileResponseBody) SetUploadId(v string) *CreateFileResponseBody {
	s.UploadId = &v
	return s
}

func (s *CreateFileResponseBody) Validate() error {
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
