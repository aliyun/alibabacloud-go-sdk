// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *GetUploadUrlRequest
	GetDriveId() *string
	SetFileId(v string) *GetUploadUrlRequest
	GetFileId() *string
	SetPartInfoList(v []*GetUploadUrlRequestPartInfoList) *GetUploadUrlRequest
	GetPartInfoList() []*GetUploadUrlRequestPartInfoList
	SetShareId(v string) *GetUploadUrlRequest
	GetShareId() *string
	SetUploadId(v string) *GetUploadUrlRequest
	GetUploadId() *string
}

type GetUploadUrlRequest struct {
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
	// 5d5b846942cf94fa72324c14a4bda34e81da635d
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The information about the file parts.
	//
	// This parameter is required.
	PartInfoList []*GetUploadUrlRequestPartInfoList `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// The share ID.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10166D06127B413BA1EC8ABB1144D111
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s GetUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetUploadUrlRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetUploadUrlRequest) GetPartInfoList() []*GetUploadUrlRequestPartInfoList {
	return s.PartInfoList
}

func (s *GetUploadUrlRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetUploadUrlRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *GetUploadUrlRequest) SetDriveId(v string) *GetUploadUrlRequest {
	s.DriveId = &v
	return s
}

func (s *GetUploadUrlRequest) SetFileId(v string) *GetUploadUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetUploadUrlRequest) SetPartInfoList(v []*GetUploadUrlRequestPartInfoList) *GetUploadUrlRequest {
	s.PartInfoList = v
	return s
}

func (s *GetUploadUrlRequest) SetShareId(v string) *GetUploadUrlRequest {
	s.ShareId = &v
	return s
}

func (s *GetUploadUrlRequest) SetUploadId(v string) *GetUploadUrlRequest {
	s.UploadId = &v
	return s
}

func (s *GetUploadUrlRequest) Validate() error {
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

type GetUploadUrlRequestPartInfoList struct {
	ContentMd5  *string `json:"content_md5,omitempty" xml:"content_md5,omitempty"`
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	ParallelSha1Ctx   *GetUploadUrlRequestPartInfoListParallelSha1Ctx   `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	ParallelSha256Ctx *GetUploadUrlRequestPartInfoListParallelSha256Ctx `json:"parallel_sha256_ctx,omitempty" xml:"parallel_sha256_ctx,omitempty" type:"Struct"`
	// The serial number of a part.
	//
	// example:
	//
	// 1
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
}

func (s GetUploadUrlRequestPartInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoList) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoList) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *GetUploadUrlRequestPartInfoList) GetContentType() *string {
	return s.ContentType
}

func (s *GetUploadUrlRequestPartInfoList) GetParallelSha1Ctx() *GetUploadUrlRequestPartInfoListParallelSha1Ctx {
	return s.ParallelSha1Ctx
}

func (s *GetUploadUrlRequestPartInfoList) GetParallelSha256Ctx() *GetUploadUrlRequestPartInfoListParallelSha256Ctx {
	return s.ParallelSha256Ctx
}

func (s *GetUploadUrlRequestPartInfoList) GetPartNumber() *int32 {
	return s.PartNumber
}

func (s *GetUploadUrlRequestPartInfoList) SetContentMd5(v string) *GetUploadUrlRequestPartInfoList {
	s.ContentMd5 = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetContentType(v string) *GetUploadUrlRequestPartInfoList {
	s.ContentType = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetParallelSha1Ctx(v *GetUploadUrlRequestPartInfoListParallelSha1Ctx) *GetUploadUrlRequestPartInfoList {
	s.ParallelSha1Ctx = v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetParallelSha256Ctx(v *GetUploadUrlRequestPartInfoListParallelSha256Ctx) *GetUploadUrlRequestPartInfoList {
	s.ParallelSha256Ctx = v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetPartNumber(v int32) *GetUploadUrlRequestPartInfoList {
	s.PartNumber = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) Validate() error {
	if s.ParallelSha1Ctx != nil {
		if err := s.ParallelSha1Ctx.Validate(); err != nil {
			return err
		}
	}
	if s.ParallelSha256Ctx != nil {
		if err := s.ParallelSha256Ctx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUploadUrlRequestPartInfoListParallelSha1Ctx struct {
	// The first to fifth 32-bit variables of the SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	H []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	// The size of the file part. Unit: bytes. The value must be a multiple of 64. This parameter takes effect only if the parallel upload feature is enabled.
	//
	// example:
	//
	// 10240
	PartOffset *int64 `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s GetUploadUrlRequestPartInfoListParallelSha1Ctx) String() string {
	return dara.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoListParallelSha1Ctx) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) GetH() []*int64 {
	return s.H
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) GetPartOffset() *int64 {
	return s.PartOffset
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) SetH(v []*int64) *GetUploadUrlRequestPartInfoListParallelSha1Ctx {
	s.H = v
	return s
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) SetPartOffset(v int64) *GetUploadUrlRequestPartInfoListParallelSha1Ctx {
	s.PartOffset = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) Validate() error {
	return dara.Validate(s)
}

type GetUploadUrlRequestPartInfoListParallelSha256Ctx struct {
	H          []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	PartOffset *int64   `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s GetUploadUrlRequestPartInfoListParallelSha256Ctx) String() string {
	return dara.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoListParallelSha256Ctx) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) GetH() []*int64 {
	return s.H
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) GetPartOffset() *int64 {
	return s.PartOffset
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) SetH(v []*int64) *GetUploadUrlRequestPartInfoListParallelSha256Ctx {
	s.H = v
	return s
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) SetPartOffset(v int64) *GetUploadUrlRequestPartInfoListParallelSha256Ctx {
	s.PartOffset = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) Validate() error {
	return dara.Validate(s)
}
