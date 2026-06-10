// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *CreateCdsFileRequest
	GetCdsId() *string
	SetConflictPolicy(v string) *CreateCdsFileRequest
	GetConflictPolicy() *string
	SetEndUserId(v string) *CreateCdsFileRequest
	GetEndUserId() *string
	SetFileHash(v string) *CreateCdsFileRequest
	GetFileHash() *string
	SetFileLength(v int64) *CreateCdsFileRequest
	GetFileLength() *int64
	SetFileName(v string) *CreateCdsFileRequest
	GetFileName() *string
	SetFileType(v string) *CreateCdsFileRequest
	GetFileType() *string
	SetGroupId(v string) *CreateCdsFileRequest
	GetGroupId() *string
	SetParentFileId(v string) *CreateCdsFileRequest
	GetParentFileId() *string
	SetRegionId(v string) *CreateCdsFileRequest
	GetRegionId() *string
}

type CreateCdsFileRequest struct {
	// Enterprise cloud disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-82414*****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// How to handle files with the same name.
	//
	// example:
	//
	// ignore
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// User ID.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// SHA-1 hash value of the file.
	//
	// example:
	//
	// 7C4A8D09CA3762AF61E59520943DC26494F8****
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// File size. Unit: Byte.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1048576
	FileLength *int64 `json:"FileLength,omitempty" xml:"FileLength,omitempty"`
	// File name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testFile.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// File type.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Team space ID.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Parent file ID. Get this from the `FileId` parameter returned by the [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) API.
	//
	// example:
	//
	// 637c9163b453b1a384874264ba79f3f9eab9****
	ParentFileId *string `json:"ParentFileId,omitempty" xml:"ParentFileId,omitempty"`
	// Region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to list regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileRequest) GoString() string {
	return s.String()
}

func (s *CreateCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCdsFileRequest) GetConflictPolicy() *string {
	return s.ConflictPolicy
}

func (s *CreateCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateCdsFileRequest) GetFileHash() *string {
	return s.FileHash
}

func (s *CreateCdsFileRequest) GetFileLength() *int64 {
	return s.FileLength
}

func (s *CreateCdsFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateCdsFileRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateCdsFileRequest) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *CreateCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCdsFileRequest) SetCdsId(v string) *CreateCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *CreateCdsFileRequest) SetConflictPolicy(v string) *CreateCdsFileRequest {
	s.ConflictPolicy = &v
	return s
}

func (s *CreateCdsFileRequest) SetEndUserId(v string) *CreateCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *CreateCdsFileRequest) SetFileHash(v string) *CreateCdsFileRequest {
	s.FileHash = &v
	return s
}

func (s *CreateCdsFileRequest) SetFileLength(v int64) *CreateCdsFileRequest {
	s.FileLength = &v
	return s
}

func (s *CreateCdsFileRequest) SetFileName(v string) *CreateCdsFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateCdsFileRequest) SetFileType(v string) *CreateCdsFileRequest {
	s.FileType = &v
	return s
}

func (s *CreateCdsFileRequest) SetGroupId(v string) *CreateCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *CreateCdsFileRequest) SetParentFileId(v string) *CreateCdsFileRequest {
	s.ParentFileId = &v
	return s
}

func (s *CreateCdsFileRequest) SetRegionId(v string) *CreateCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
