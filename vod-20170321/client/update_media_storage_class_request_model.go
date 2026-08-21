// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaStorageClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUpdateWithoutTimeLimit(v bool) *UpdateMediaStorageClassRequest
	GetAllowUpdateWithoutTimeLimit() *bool
	SetMediaIds(v string) *UpdateMediaStorageClassRequest
	GetMediaIds() *string
	SetRestoreTier(v string) *UpdateMediaStorageClassRequest
	GetRestoreTier() *string
	SetScope(v string) *UpdateMediaStorageClassRequest
	GetScope() *string
	SetStorageClass(v string) *UpdateMediaStorageClassRequest
	GetStorageClass() *string
}

type UpdateMediaStorageClassRequest struct {
	// Specifies whether to allow storage class modification for media assets that have not met the minimum storage duration requirement. Valid values:
	//
	// - **true**: Allowed.
	//
	// - **false (default)**: Not allowed.
	//
	// >If the storage duration of a media asset is insufficient and you force a storage class modification, additional retrieval fees are incurred.
	//
	// example:
	//
	// false
	AllowUpdateWithoutTimeLimit *bool `json:"AllowUpdateWithoutTimeLimit,omitempty" xml:"AllowUpdateWithoutTimeLimit,omitempty"`
	// The media IDs, which are audio or video IDs (VideoId). Separate multiple IDs with commas (,). A maximum of 20 IDs are supported. You can obtain the IDs by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// d56c2ac0cee271ed80004710b5ba****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The restore priority (required only for ColdArchive media assets). If this parameter is not specified, the default value **Standard*	- is used. Valid values:
	//
	// - **Expedited**: Expedited
	//
	// - **Standard*	- (default): Standard
	//
	// - **Bulk**: Bulk
	//
	// example:
	//
	// Standard
	RestoreTier *string `json:"RestoreTier,omitempty" xml:"RestoreTier,omitempty"`
	// The scope of the modification. If this parameter is not specified, the default value **All*	- is used. Valid values:
	//
	// - **All*	- (default): Applies tiered storage to all resources (source files and transcoded streams) of the media asset.
	//
	// - **SourceFile**: Applies tiered storage only to the source file of the media asset. Resources other than the source file use Standard storage.
	//
	// example:
	//
	// All
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The storage class. Valid values:
	//
	// - **Standard**: Standard
	//
	// - **IA**: Infrequent Access
	//
	// - **Archive**: Archive
	//
	// - **ColdArchive**: Cold Archive
	//
	// This parameter is required.
	//
	// example:
	//
	// Archive
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s UpdateMediaStorageClassRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaStorageClassRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaStorageClassRequest) GetAllowUpdateWithoutTimeLimit() *bool {
	return s.AllowUpdateWithoutTimeLimit
}

func (s *UpdateMediaStorageClassRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *UpdateMediaStorageClassRequest) GetRestoreTier() *string {
	return s.RestoreTier
}

func (s *UpdateMediaStorageClassRequest) GetScope() *string {
	return s.Scope
}

func (s *UpdateMediaStorageClassRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *UpdateMediaStorageClassRequest) SetAllowUpdateWithoutTimeLimit(v bool) *UpdateMediaStorageClassRequest {
	s.AllowUpdateWithoutTimeLimit = &v
	return s
}

func (s *UpdateMediaStorageClassRequest) SetMediaIds(v string) *UpdateMediaStorageClassRequest {
	s.MediaIds = &v
	return s
}

func (s *UpdateMediaStorageClassRequest) SetRestoreTier(v string) *UpdateMediaStorageClassRequest {
	s.RestoreTier = &v
	return s
}

func (s *UpdateMediaStorageClassRequest) SetScope(v string) *UpdateMediaStorageClassRequest {
	s.Scope = &v
	return s
}

func (s *UpdateMediaStorageClassRequest) SetStorageClass(v string) *UpdateMediaStorageClassRequest {
	s.StorageClass = &v
	return s
}

func (s *UpdateMediaStorageClassRequest) Validate() error {
	return dara.Validate(s)
}
