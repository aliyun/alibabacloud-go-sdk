// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListFileRequest
	GetCategory() *string
	SetDriveId(v string) *ListFileRequest
	GetDriveId() *string
	SetFields(v string) *ListFileRequest
	GetFields() *string
	SetLimit(v int32) *ListFileRequest
	GetLimit() *int32
	SetMarker(v string) *ListFileRequest
	GetMarker() *string
	SetOrderBy(v string) *ListFileRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListFileRequest
	GetOrderDirection() *string
	SetParentFileId(v string) *ListFileRequest
	GetParentFileId() *string
	SetShareId(v string) *ListFileRequest
	GetShareId() *string
	SetStatus(v string) *ListFileRequest
	GetStatus() *string
	SetThumbnailProcesses(v map[string]*ImageProcess) *ListFileRequest
	GetThumbnailProcesses() map[string]*ImageProcess
	SetType(v string) *ListFileRequest
	GetType() *string
}

type ListFileRequest struct {
	// The file category. Valid values:
	//
	// app: installation package zip: compressed package image doc: document video audio others
	//
	// By default, files of all categories are returned.
	//
	// example:
	//
	// image
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The field that is used to return additional information about a child subject. Valid values:
	//
	// 	- url: returns the URL of the thumbnail of a file in the response.
	//
	// 	- exif: returns the Exchangeable Image File Format (EXIF) data of a file in the response.
	//
	// 	- cropping_suggestion: returns the cropping suggestion on a file in the response.
	//
	// 	- characteristic_hash: returns the characteristic hash value of a file in the response.
	//
	// 	- video_metadata: returns the metadata of a video file, such as the video duration, bitrate, height, and width, in the response.
	//
	// 	- video_preview_metadata: returns the transcoding information of a video file, such as the transcoding specification for each definition, in the response.
	//
	// 	- investigation_info: returns the investigation information in the response.
	//
	// 	- dir_size: returns the statistics on each subfolder in the response.
	//
	// 	- user_tags: returns the user tags of each child subject in the response.
	//
	// You can specify multiple fields by separating them with commas (,). Example: "url,dir_size,user_tags".
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned entries must be less than or equal to the value of this parameter.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The name of the entry after which the list begins. Entries whose names are alphabetically after the value of this parameter are returned. If you do not specify this parameter, all entries are returned.\\
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The sorting field. Valid values:
	//
	// created_at: sorts the entries by creation time. updated_at: sorts the entries by update time. size: sorts the entries by file size. name: sorts the entries by file name.
	//
	// Default value: created_at.
	//
	// Enumeration:
	//
	// 	- updated_at: update time
	//
	// 	- size: file size
	//
	// 	- name: file name
	//
	// 	- created_at: creation time
	//
	// example:
	//
	// updated_at
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The sorting direction. Valid values:
	//
	// ASC: ascending order DESC: descending order
	//
	// Default value: ASC.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	// The ID of the parent folder. If the parent folder is a root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The share ID. If you want to share a file, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The state of the files to return. Valid values:
	//
	// available: returns only normal files. uploading: returns only files that are being uploaded.
	//
	// By default, only files in the available state are returned.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The thumbnail configurations. Up to five thumbnails can be returned at a time. The value contains key-value pairs. You can customize the keys. The URL of a thumbnail is returned based on the key.
	ThumbnailProcesses map[string]*ImageProcess `json:"thumbnail_processes,omitempty" xml:"thumbnail_processes,omitempty"`
	// The file type. Valid values:
	//
	// file: returns only files. folder: returns only folders.
	//
	// By default, files of all types are returned.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileRequest) GoString() string {
	return s.String()
}

func (s *ListFileRequest) GetCategory() *string {
	return s.Category
}

func (s *ListFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListFileRequest) GetFields() *string {
	return s.Fields
}

func (s *ListFileRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFileRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListFileRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListFileRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListFileRequest) GetParentFileId() *string {
	return s.ParentFileId
}

func (s *ListFileRequest) GetShareId() *string {
	return s.ShareId
}

func (s *ListFileRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFileRequest) GetThumbnailProcesses() map[string]*ImageProcess {
	return s.ThumbnailProcesses
}

func (s *ListFileRequest) GetType() *string {
	return s.Type
}

func (s *ListFileRequest) SetCategory(v string) *ListFileRequest {
	s.Category = &v
	return s
}

func (s *ListFileRequest) SetDriveId(v string) *ListFileRequest {
	s.DriveId = &v
	return s
}

func (s *ListFileRequest) SetFields(v string) *ListFileRequest {
	s.Fields = &v
	return s
}

func (s *ListFileRequest) SetLimit(v int32) *ListFileRequest {
	s.Limit = &v
	return s
}

func (s *ListFileRequest) SetMarker(v string) *ListFileRequest {
	s.Marker = &v
	return s
}

func (s *ListFileRequest) SetOrderBy(v string) *ListFileRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFileRequest) SetOrderDirection(v string) *ListFileRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListFileRequest) SetParentFileId(v string) *ListFileRequest {
	s.ParentFileId = &v
	return s
}

func (s *ListFileRequest) SetShareId(v string) *ListFileRequest {
	s.ShareId = &v
	return s
}

func (s *ListFileRequest) SetStatus(v string) *ListFileRequest {
	s.Status = &v
	return s
}

func (s *ListFileRequest) SetThumbnailProcesses(v map[string]*ImageProcess) *ListFileRequest {
	s.ThumbnailProcesses = v
	return s
}

func (s *ListFileRequest) SetType(v string) *ListFileRequest {
	s.Type = &v
	return s
}

func (s *ListFileRequest) Validate() error {
	return dara.Validate(s)
}
