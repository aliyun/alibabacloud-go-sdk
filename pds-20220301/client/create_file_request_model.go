// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckNameMode(v string) *CreateFileRequest
	GetCheckNameMode() *string
	SetContentHash(v string) *CreateFileRequest
	GetContentHash() *string
	SetContentHashName(v string) *CreateFileRequest
	GetContentHashName() *string
	SetContentType(v string) *CreateFileRequest
	GetContentType() *string
	SetDescription(v string) *CreateFileRequest
	GetDescription() *string
	SetDriveId(v string) *CreateFileRequest
	GetDriveId() *string
	SetFileId(v string) *CreateFileRequest
	GetFileId() *string
	SetHidden(v bool) *CreateFileRequest
	GetHidden() *bool
	SetLocalCreatedAt(v string) *CreateFileRequest
	GetLocalCreatedAt() *string
	SetLocalModifiedAt(v string) *CreateFileRequest
	GetLocalModifiedAt() *string
	SetName(v string) *CreateFileRequest
	GetName() *string
	SetParallelUpload(v bool) *CreateFileRequest
	GetParallelUpload() *bool
	SetParentFileId(v string) *CreateFileRequest
	GetParentFileId() *string
	SetPartInfoList(v []*CreateFileRequestPartInfoList) *CreateFileRequest
	GetPartInfoList() []*CreateFileRequestPartInfoList
	SetPreHash(v string) *CreateFileRequest
	GetPreHash() *string
	SetShareId(v string) *CreateFileRequest
	GetShareId() *string
	SetSize(v int64) *CreateFileRequest
	GetSize() *int64
	SetType(v string) *CreateFileRequest
	GetType() *string
	SetUserTags(v []*UserTag) *CreateFileRequest
	GetUserTags() []*UserTag
}

type CreateFileRequest struct {
	// The processing method that is used if the file that you want to create has the same name as an existing file in the cloud. Valid values:
	//
	// ignore: allows you to create the file by using the same name as an existing file in the cloud.
	//
	// auto_rename: automatically renames the file that you want to create. By default, the current point in time is added to the end of the file name. Example: xxx_20060102_150405.
	//
	// refuse: does not create the file that you want to create but returns the information about the file that has the same name in the cloud.
	//
	// Default value: ignore.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ignore
	CheckNameMode *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	// The hash value of the file content. The value is calculated based on the algorithm specified by content_hash_name.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 7C4A8D09CA3762AF61E59520943DC26494F8941B
	ContentHash *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	// The name of the algorithm that is used to calculate the hash value of the file content. Only SHA1 is supported.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// sha1
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	// The type of the file content. Default value: application/oct-stream.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// application/json
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The description of the file. The description can be up to 1,024 characters in length. By default, this parameter is left empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 重要文件
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The drive ID. This parameter is required if the file is not uploaded by using the share URL of the file.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID. This parameter is required if check_name_mode is set to ignore.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to hide the file or folder. By default, the file or folder is not hidden.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// The time when the local file was created. By default, this parameter is left empty. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format based on the UTC+0 time zone.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalCreatedAt *string `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	// The time when the local file was modified. By default, this parameter is left empty. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format based on the UTC+0 time zone.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalModifiedAt *string `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	// The name of the file. The name can be up to 1,024 bytes in length based on the UTF-8 encoding rule and cannot contain forward slash (/).
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// a.txt
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to enable the parallel upload feature.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	ParallelUpload *bool `json:"parallel_upload,omitempty" xml:"parallel_upload,omitempty"`
	// The ID of the parent directory. If you want to create a file or folder in the root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// fileid1
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The information about the file parts. You can specify up to 10,000 parts. By default, if you do not specify this parameter, only one part is returned.
	PartInfoList []*CreateFileRequestPartInfoList `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// The SHA-1 hash value of the first 1 KB data of the file. This parameter is required if you perform instant file upload by using the pre-hashing feature. If the SHA-1 hash value is not matched on the cloud, the client does not need to calculate the SHA-1 hash value of the entire file.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 7C4A8D09CA3762AF61E59520943DC26494F89411
	PreHash *string `json:"pre_hash,omitempty" xml:"pre_hash,omitempty"`
	// The share ID. This parameter is required if the file is uploaded by using the share URL of the file.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The type of the file. Valid values:
	//
	// file folder
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The custom tags. You can specify up to 1,000 tags.
	UserTags []*UserTag `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
}

func (s CreateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) GetCheckNameMode() *string {
	return s.CheckNameMode
}

func (s *CreateFileRequest) GetContentHash() *string {
	return s.ContentHash
}

func (s *CreateFileRequest) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *CreateFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateFileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateFileRequest) GetHidden() *bool {
	return s.Hidden
}

func (s *CreateFileRequest) GetLocalCreatedAt() *string {
	return s.LocalCreatedAt
}

func (s *CreateFileRequest) GetLocalModifiedAt() *string {
	return s.LocalModifiedAt
}

func (s *CreateFileRequest) GetName() *string {
	return s.Name
}

func (s *CreateFileRequest) GetParallelUpload() *bool {
	return s.ParallelUpload
}

func (s *CreateFileRequest) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *CreateFileRequest) GetPartInfoList() []*CreateFileRequestPartInfoList {
	return s.PartInfoList
}

func (s *CreateFileRequest) GetPreHash() *string {
	return s.PreHash
}

func (s *CreateFileRequest) GetShareId() *string {
	return s.ShareId
}

func (s *CreateFileRequest) GetSize() *int64 {
	return s.Size
}

func (s *CreateFileRequest) GetType() *string {
	return s.Type
}

func (s *CreateFileRequest) GetUserTags() []*UserTag {
	return s.UserTags
}

func (s *CreateFileRequest) SetCheckNameMode(v string) *CreateFileRequest {
	s.CheckNameMode = &v
	return s
}

func (s *CreateFileRequest) SetContentHash(v string) *CreateFileRequest {
	s.ContentHash = &v
	return s
}

func (s *CreateFileRequest) SetContentHashName(v string) *CreateFileRequest {
	s.ContentHashName = &v
	return s
}

func (s *CreateFileRequest) SetContentType(v string) *CreateFileRequest {
	s.ContentType = &v
	return s
}

func (s *CreateFileRequest) SetDescription(v string) *CreateFileRequest {
	s.Description = &v
	return s
}

func (s *CreateFileRequest) SetDriveId(v string) *CreateFileRequest {
	s.DriveId = &v
	return s
}

func (s *CreateFileRequest) SetFileId(v string) *CreateFileRequest {
	s.FileId = &v
	return s
}

func (s *CreateFileRequest) SetHidden(v bool) *CreateFileRequest {
	s.Hidden = &v
	return s
}

func (s *CreateFileRequest) SetLocalCreatedAt(v string) *CreateFileRequest {
	s.LocalCreatedAt = &v
	return s
}

func (s *CreateFileRequest) SetLocalModifiedAt(v string) *CreateFileRequest {
	s.LocalModifiedAt = &v
	return s
}

func (s *CreateFileRequest) SetName(v string) *CreateFileRequest {
	s.Name = &v
	return s
}

func (s *CreateFileRequest) SetParallelUpload(v bool) *CreateFileRequest {
	s.ParallelUpload = &v
	return s
}

func (s *CreateFileRequest) SetParentFileId(v string) *CreateFileRequest {
	s.ParentFileId = &v
	return s
}

func (s *CreateFileRequest) SetPartInfoList(v []*CreateFileRequestPartInfoList) *CreateFileRequest {
	s.PartInfoList = v
	return s
}

func (s *CreateFileRequest) SetPreHash(v string) *CreateFileRequest {
	s.PreHash = &v
	return s
}

func (s *CreateFileRequest) SetShareId(v string) *CreateFileRequest {
	s.ShareId = &v
	return s
}

func (s *CreateFileRequest) SetSize(v int64) *CreateFileRequest {
	s.Size = &v
	return s
}

func (s *CreateFileRequest) SetType(v string) *CreateFileRequest {
	s.Type = &v
	return s
}

func (s *CreateFileRequest) SetUserTags(v []*UserTag) *CreateFileRequest {
	s.UserTags = v
	return s
}

func (s *CreateFileRequest) Validate() error {
	if s.PartInfoList != nil {
		for _, item := range s.PartInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserTags != nil {
		for _, item := range s.UserTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFileRequestPartInfoList struct {
	// The MD5 hash value of the file part. This parameter is required when the MD5 hash value of the file part needs to be verified during part upload.
	//
	// example:
	//
	// ASKJDJSKDJJSJDJS
	ContentMd5 *string `json:"content_md5,omitempty" xml:"content_md5,omitempty"`
	// example:
	//
	// application/json
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	ParallelSha1Ctx *CreateFileRequestPartInfoListParallelSha1Ctx `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	// The serial number of a file part. The number starts from 1.
	//
	// example:
	//
	// 1
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
}

func (s CreateFileRequestPartInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequestPartInfoList) GoString() string {
	return s.String()
}

func (s *CreateFileRequestPartInfoList) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *CreateFileRequestPartInfoList) GetContentType() *string {
	return s.ContentType
}

func (s *CreateFileRequestPartInfoList) GetParallelSha1Ctx() *CreateFileRequestPartInfoListParallelSha1Ctx {
	return s.ParallelSha1Ctx
}

func (s *CreateFileRequestPartInfoList) GetPartNumber() *int32 {
	return s.PartNumber
}

func (s *CreateFileRequestPartInfoList) SetContentMd5(v string) *CreateFileRequestPartInfoList {
	s.ContentMd5 = &v
	return s
}

func (s *CreateFileRequestPartInfoList) SetContentType(v string) *CreateFileRequestPartInfoList {
	s.ContentType = &v
	return s
}

func (s *CreateFileRequestPartInfoList) SetParallelSha1Ctx(v *CreateFileRequestPartInfoListParallelSha1Ctx) *CreateFileRequestPartInfoList {
	s.ParallelSha1Ctx = v
	return s
}

func (s *CreateFileRequestPartInfoList) SetPartNumber(v int32) *CreateFileRequestPartInfoList {
	s.PartNumber = &v
	return s
}

func (s *CreateFileRequestPartInfoList) Validate() error {
	if s.ParallelSha1Ctx != nil {
		if err := s.ParallelSha1Ctx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFileRequestPartInfoListParallelSha1Ctx struct {
	// The first to fifth 32-bit variables of the SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	H []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	// The size of the file content before the file part. Unit: bytes. The value must be a multiple of 64. This parameter takes effect only if the parallel upload feature is enabled.
	//
	// example:
	//
	// 10240
	PartOffset *int64 `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s CreateFileRequestPartInfoListParallelSha1Ctx) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequestPartInfoListParallelSha1Ctx) GoString() string {
	return s.String()
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) GetH() []*int64 {
	return s.H
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) GetPartOffset() *int64 {
	return s.PartOffset
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) SetH(v []*int64) *CreateFileRequestPartInfoListParallelSha1Ctx {
	s.H = v
	return s
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) SetPartOffset(v int64) *CreateFileRequestPartInfoListParallelSha1Ctx {
	s.PartOffset = &v
	return s
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) Validate() error {
	return dara.Validate(s)
}
