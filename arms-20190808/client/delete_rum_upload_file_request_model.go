// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumUploadFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchItems(v string) *DeleteRumUploadFileRequest
	GetBatchItems() *string
	SetFileName(v string) *DeleteRumUploadFileRequest
	GetFileName() *string
	SetPid(v string) *DeleteRumUploadFileRequest
	GetPid() *string
	SetRegionId(v string) *DeleteRumUploadFileRequest
	GetRegionId() *string
	SetUuid(v string) *DeleteRumUploadFileRequest
	GetUuid() *string
	SetVersionId(v string) *DeleteRumUploadFileRequest
	GetVersionId() *string
}

type DeleteRumUploadFileRequest struct {
	// Information of files to be deleted in JSON array format. If a single file needs to be deleted, this field should be left empty. If multiple files need to be deleted, just fill in this field.
	//
	// example:
	//
	// [{
	//
	//     "fileName" : "test.js.map",
	//
	//     "version" : "1.0.0"
	//
	//   },
	//
	//   {
	//
	//     "fileName" : "test.dSYM",
	//
	//     "version" : "1.20.1",
	//
	//     "uuid" : "xxxx-xxxx-xxxx-xxxx"
	//
	//   }]
	BatchItems *string `json:"BatchItems,omitempty" xml:"BatchItems,omitempty"`
	// The file name, with the extension.
	//
	// example:
	//
	// test.js.map
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The file ID.
	//
	// example:
	//
	// MS4wLjAtbWFpbi4wZjM0NzRlOSxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The version number of the file.
	//
	// example:
	//
	// 1.0.0
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteRumUploadFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumUploadFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteRumUploadFileRequest) GetBatchItems() *string {
	return s.BatchItems
}

func (s *DeleteRumUploadFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *DeleteRumUploadFileRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteRumUploadFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRumUploadFileRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteRumUploadFileRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DeleteRumUploadFileRequest) SetBatchItems(v string) *DeleteRumUploadFileRequest {
	s.BatchItems = &v
	return s
}

func (s *DeleteRumUploadFileRequest) SetFileName(v string) *DeleteRumUploadFileRequest {
	s.FileName = &v
	return s
}

func (s *DeleteRumUploadFileRequest) SetPid(v string) *DeleteRumUploadFileRequest {
	s.Pid = &v
	return s
}

func (s *DeleteRumUploadFileRequest) SetRegionId(v string) *DeleteRumUploadFileRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRumUploadFileRequest) SetUuid(v string) *DeleteRumUploadFileRequest {
	s.Uuid = &v
	return s
}

func (s *DeleteRumUploadFileRequest) SetVersionId(v string) *DeleteRumUploadFileRequest {
	s.VersionId = &v
	return s
}

func (s *DeleteRumUploadFileRequest) Validate() error {
	return dara.Validate(s)
}
