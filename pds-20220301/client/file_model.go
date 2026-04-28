// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFile interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*string) *File
	GetActionList() []*string
	SetAutoDeleteLeftSec(v int64) *File
	GetAutoDeleteLeftSec() *int64
	SetCategory(v string) *File
	GetCategory() *string
	SetContentHash(v string) *File
	GetContentHash() *string
	SetContentHashName(v string) *File
	GetContentHashName() *string
	SetContentType(v string) *File
	GetContentType() *string
	SetCrc64Hash(v string) *File
	GetCrc64Hash() *string
	SetCreatedAt(v string) *File
	GetCreatedAt() *string
	SetDescription(v string) *File
	GetDescription() *string
	SetDirSizeInfo(v *FileDirSizeInfo) *File
	GetDirSizeInfo() *FileDirSizeInfo
	SetDomainId(v string) *File
	GetDomainId() *string
	SetDownloadUrl(v string) *File
	GetDownloadUrl() *string
	SetDriveId(v string) *File
	GetDriveId() *string
	SetFileExtension(v string) *File
	GetFileExtension() *string
	SetFileId(v string) *File
	GetFileId() *string
	SetHidden(v bool) *File
	GetHidden() *bool
	SetIdPath(v string) *File
	GetIdPath() *string
	SetImageMediaMetadata(v *ImageMediaMetadata) *File
	GetImageMediaMetadata() *ImageMediaMetadata
	SetLabels(v []*string) *File
	GetLabels() []*string
	SetLocalCreatedAt(v string) *File
	GetLocalCreatedAt() *string
	SetLocalModifiedAt(v string) *File
	GetLocalModifiedAt() *string
	SetName(v string) *File
	GetName() *string
	SetNamePath(v string) *File
	GetNamePath() *string
	SetParentFileId(v string) *File
	GetParentFileId() *string
	SetRevisionId(v string) *File
	GetRevisionId() *string
	SetSize(v int64) *File
	GetSize() *int64
	SetStarred(v bool) *File
	GetStarred() *bool
	SetStatus(v string) *File
	GetStatus() *string
	SetThumbnail(v string) *File
	GetThumbnail() *string
	SetThumbnailUrls(v map[string]*string) *File
	GetThumbnailUrls() map[string]*string
	SetTrashedAt(v string) *File
	GetTrashedAt() *string
	SetType(v string) *File
	GetType() *string
	SetUpdatedAt(v string) *File
	GetUpdatedAt() *string
	SetUploadId(v string) *File
	GetUploadId() *string
	SetUserTags(v map[string]*string) *File
	GetUserTags() map[string]*string
	SetVideoMediaMetadata(v *VideoMediaMetadata) *File
	GetVideoMediaMetadata() *VideoMediaMetadata
}

type File struct {
	// The permissions.
	ActionList []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// The remaining time until the file is automatically deleted from the recycle bin (if in it).
	AutoDeleteLeftSec *int64 `json:"auto_delete_left_sec,omitempty" xml:"auto_delete_left_sec,omitempty"`
	// The category. Drive and Photo Service (PDS) classifies files based on their extensions and mime-type. The supported categories include doc, image, audio, and video.
	//
	// example:
	//
	// image
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The hash value of the content.
	//
	// example:
	//
	// EA4942AA8761213890A5C386F88E6464D2C31CA3
	ContentHash *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	// The name of the hash algorithm. Set the value to sha1.
	//
	// example:
	//
	// sha1
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	// The type of the content.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// crc64
	//
	// example:
	//
	// 3574582125365864471
	Crc64Hash *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	// The time when the file was created.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// image file
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the folder structure. This parameter is returned only if you include the dir_size field in the fields parameter by calling the ListFile or GetFile operation.
	DirSizeInfo *FileDirSizeInfo `json:"dir_size_info,omitempty" xml:"dir_size_info,omitempty" type:"Struct"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The download URL. The default validity period of the download URL is 15 minutes. If the URL expires, you can obtain the URL by calling the GetFile operation.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	DownloadUrl *string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file name extension.
	//
	// example:
	//
	// txt
	FileExtension *string `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 5d79206586bb5dd69fb34c349282718146c55da7
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to hide the file.
	//
	// example:
	//
	// false
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// The file ID path.
	//
	// example:
	//
	// id1/id2
	IdPath *string `json:"id_path,omitempty" xml:"id_path,omitempty"`
	// The image metadata. This parameter takes effect only if the value-added image processing feature is enabled.
	ImageMediaMetadata *ImageMediaMetadata `json:"image_media_metadata,omitempty" xml:"image_media_metadata,omitempty"`
	// The labels of the file.
	//
	// example:
	//
	// ["label1:1", "label2:2"]
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The time when the local file was created. The time refers to the local time when the file was uploaded. This parameter helps identify the local upload time.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalCreatedAt *string `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	// The time when the local file was modified. The time refers to the local time when the modified file was uploaded. This parameter helps identify the local update time.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalModifiedAt *string `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	// The file name.
	//
	// example:
	//
	// 1.mov
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The file path.
	//
	// example:
	//
	// root/f1/f2
	NamePath *string `json:"name_path,omitempty" xml:"name_path,omitempty"`
	// The parent folder ID.
	//
	// example:
	//
	// 3d5b846942cf94fa72324c8a4bda34e81da635a
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The version ID. If a file that has the same file ID with an existing one is uploaded, a new version ID is generated for the file.
	//
	// example:
	//
	// 5d5b846942cf94fa72324c14a4bda34e81da635d
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	// The file size
	//
	// or folder size. The folder size is calculated based on all descendant files and folders in the folder. Note: The folder size can be returned only when you call the ListFile or GetFile operation and include the dir_size field in the fields parameter.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to add the file to favorites.
	//
	// example:
	//
	// false
	Starred *bool `json:"starred,omitempty" xml:"starred,omitempty"`
	// The status of the file. Only files and directories in the available state can be accessed. If you call the GetFile operation to obtain a file that is in the uploading state, a response indicating that the file does not exist is returned. If you call the ListFile operation to query files, files in the uploading state are not returned.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The URL of the thumbnail. This parameter is deprecated and we recommend that you use thumbnail_urls.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Thumbnail *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	// The information about the returned thumbnail. The value corresponds to the key that is specified by thumbnail_processes.
	ThumbnailUrls map[string]*string `json:"thumbnail_urls,omitempty" xml:"thumbnail_urls,omitempty"`
	// The time when the file was put into the recycle bin.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	TrashedAt *string `json:"trashed_at,omitempty" xml:"trashed_at,omitempty"`
	// The file type.
	//
	// Valid values:
	//
	// 	- file
	//
	// 	- folder
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the file was modified.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The upload ID.
	//
	// example:
	//
	// C9DCFE5A82644AC7A02DB74C30C934A6
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	// The custom tags.
	UserTags map[string]*string `json:"user_tags,omitempty" xml:"user_tags,omitempty"`
	// The audio and video information.
	VideoMediaMetadata *VideoMediaMetadata `json:"video_media_metadata,omitempty" xml:"video_media_metadata,omitempty"`
}

func (s File) String() string {
	return dara.Prettify(s)
}

func (s File) GoString() string {
	return s.String()
}

func (s *File) GetActionList() []*string {
	return s.ActionList
}

func (s *File) GetAutoDeleteLeftSec() *int64 {
	return s.AutoDeleteLeftSec
}

func (s *File) GetCategory() *string {
	return s.Category
}

func (s *File) GetContentHash() *string {
	return s.ContentHash
}

func (s *File) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *File) GetContentType() *string {
	return s.ContentType
}

func (s *File) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *File) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *File) GetDescription() *string {
	return s.Description
}

func (s *File) GetDirSizeInfo() *FileDirSizeInfo {
	return s.DirSizeInfo
}

func (s *File) GetDomainId() *string {
	return s.DomainId
}

func (s *File) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *File) GetDriveId() *string {
	return s.DriveId
}

func (s *File) GetFileExtension() *string {
	return s.FileExtension
}

func (s *File) GetFileId() *string {
	return s.FileId
}

func (s *File) GetHidden() *bool {
	return s.Hidden
}

func (s *File) GetIdPath() *string {
	return s.IdPath
}

func (s *File) GetImageMediaMetadata() *ImageMediaMetadata {
	return s.ImageMediaMetadata
}

func (s *File) GetLabels() []*string {
	return s.Labels
}

func (s *File) GetLocalCreatedAt() *string {
	return s.LocalCreatedAt
}

func (s *File) GetLocalModifiedAt() *string {
	return s.LocalModifiedAt
}

func (s *File) GetName() *string {
	return s.Name
}

func (s *File) GetNamePath() *string {
	return s.NamePath
}

func (s *File) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *File) GetRevisionId() *string {
	return s.RevisionId
}

func (s *File) GetSize() *int64 {
	return s.Size
}

func (s *File) GetStarred() *bool {
	return s.Starred
}

func (s *File) GetStatus() *string {
	return s.Status
}

func (s *File) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *File) GetThumbnailUrls() map[string]*string {
	return s.ThumbnailUrls
}

func (s *File) GetTrashedAt() *string {
	return s.TrashedAt
}

func (s *File) GetType() *string {
	return s.Type
}

func (s *File) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *File) GetUploadId() *string {
	return s.UploadId
}

func (s *File) GetUserTags() map[string]*string {
	return s.UserTags
}

func (s *File) GetVideoMediaMetadata() *VideoMediaMetadata {
	return s.VideoMediaMetadata
}

func (s *File) SetActionList(v []*string) *File {
	s.ActionList = v
	return s
}

func (s *File) SetAutoDeleteLeftSec(v int64) *File {
	s.AutoDeleteLeftSec = &v
	return s
}

func (s *File) SetCategory(v string) *File {
	s.Category = &v
	return s
}

func (s *File) SetContentHash(v string) *File {
	s.ContentHash = &v
	return s
}

func (s *File) SetContentHashName(v string) *File {
	s.ContentHashName = &v
	return s
}

func (s *File) SetContentType(v string) *File {
	s.ContentType = &v
	return s
}

func (s *File) SetCrc64Hash(v string) *File {
	s.Crc64Hash = &v
	return s
}

func (s *File) SetCreatedAt(v string) *File {
	s.CreatedAt = &v
	return s
}

func (s *File) SetDescription(v string) *File {
	s.Description = &v
	return s
}

func (s *File) SetDirSizeInfo(v *FileDirSizeInfo) *File {
	s.DirSizeInfo = v
	return s
}

func (s *File) SetDomainId(v string) *File {
	s.DomainId = &v
	return s
}

func (s *File) SetDownloadUrl(v string) *File {
	s.DownloadUrl = &v
	return s
}

func (s *File) SetDriveId(v string) *File {
	s.DriveId = &v
	return s
}

func (s *File) SetFileExtension(v string) *File {
	s.FileExtension = &v
	return s
}

func (s *File) SetFileId(v string) *File {
	s.FileId = &v
	return s
}

func (s *File) SetHidden(v bool) *File {
	s.Hidden = &v
	return s
}

func (s *File) SetIdPath(v string) *File {
	s.IdPath = &v
	return s
}

func (s *File) SetImageMediaMetadata(v *ImageMediaMetadata) *File {
	s.ImageMediaMetadata = v
	return s
}

func (s *File) SetLabels(v []*string) *File {
	s.Labels = v
	return s
}

func (s *File) SetLocalCreatedAt(v string) *File {
	s.LocalCreatedAt = &v
	return s
}

func (s *File) SetLocalModifiedAt(v string) *File {
	s.LocalModifiedAt = &v
	return s
}

func (s *File) SetName(v string) *File {
	s.Name = &v
	return s
}

func (s *File) SetNamePath(v string) *File {
	s.NamePath = &v
	return s
}

func (s *File) SetParentFileId(v string) *File {
	s.ParentFileId = &v
	return s
}

func (s *File) SetRevisionId(v string) *File {
	s.RevisionId = &v
	return s
}

func (s *File) SetSize(v int64) *File {
	s.Size = &v
	return s
}

func (s *File) SetStarred(v bool) *File {
	s.Starred = &v
	return s
}

func (s *File) SetStatus(v string) *File {
	s.Status = &v
	return s
}

func (s *File) SetThumbnail(v string) *File {
	s.Thumbnail = &v
	return s
}

func (s *File) SetThumbnailUrls(v map[string]*string) *File {
	s.ThumbnailUrls = v
	return s
}

func (s *File) SetTrashedAt(v string) *File {
	s.TrashedAt = &v
	return s
}

func (s *File) SetType(v string) *File {
	s.Type = &v
	return s
}

func (s *File) SetUpdatedAt(v string) *File {
	s.UpdatedAt = &v
	return s
}

func (s *File) SetUploadId(v string) *File {
	s.UploadId = &v
	return s
}

func (s *File) SetUserTags(v map[string]*string) *File {
	s.UserTags = v
	return s
}

func (s *File) SetVideoMediaMetadata(v *VideoMediaMetadata) *File {
	s.VideoMediaMetadata = v
	return s
}

func (s *File) Validate() error {
	if s.DirSizeInfo != nil {
		if err := s.DirSizeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ImageMediaMetadata != nil {
		if err := s.ImageMediaMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.VideoMediaMetadata != nil {
		if err := s.VideoMediaMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FileDirSizeInfo struct {
	// The number of all descendant folders in the folder, which is calculated recursively.
	//
	// example:
	//
	// 1
	DirCount *int64 `json:"dir_count,omitempty" xml:"dir_count,omitempty"`
	// The number of all descendant files in the folder, which is calculated recursively.
	//
	// example:
	//
	// 10
	FileCount *int64 `json:"file_count,omitempty" xml:"file_count,omitempty"`
}

func (s FileDirSizeInfo) String() string {
	return dara.Prettify(s)
}

func (s FileDirSizeInfo) GoString() string {
	return s.String()
}

func (s *FileDirSizeInfo) GetDirCount() *int64 {
	return s.DirCount
}

func (s *FileDirSizeInfo) GetFileCount() *int64 {
	return s.FileCount
}

func (s *FileDirSizeInfo) SetDirCount(v int64) *FileDirSizeInfo {
	s.DirCount = &v
	return s
}

func (s *FileDirSizeInfo) SetFileCount(v int64) *FileDirSizeInfo {
	s.FileCount = &v
	return s
}

func (s *FileDirSizeInfo) Validate() error {
	return dara.Validate(s)
}
