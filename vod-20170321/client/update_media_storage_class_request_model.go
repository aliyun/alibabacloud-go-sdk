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
	// Specifies whether to allow modifying the storage class of media assets whose storage duration does not meet the minimum requirement. Valid values:
	//
	// - **true**: allows the modification.
	//
	// - **false*	- (default): does not allow the modification.
	//
	// > If the storage duration of a media asset does not meet the minimum requirement and you force the storage class modification, additional data retrieval fees are incurred.
	//
	// example:
	//
	// false
	AllowUpdateWithoutTimeLimit *bool `json:"AllowUpdateWithoutTimeLimit,omitempty" xml:"AllowUpdateWithoutTimeLimit,omitempty"`
	// The media asset IDs, which are the audio or video IDs (VideoId). Separate multiple IDs with commas (,). You can specify up to 20 IDs. You can obtain the IDs by using one of the following methods:
	//
	// - For audio or video files uploaded in the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Assets*	- > **Audio/Video*	- to view the audio or video IDs.
	//
	// - When you call the [Obtain the upload URL and credential for audio and video files](https://help.aliyun.com/document_detail/55407.html) operation, the value of the VideoId response parameter is the media asset ID.
	//
	// - After an audio or video file is uploaded, you can call the [Search for media asset information](https://help.aliyun.com/document_detail/86044.html) operation to query the media asset ID. The value of the VideoId response parameter is the media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d56c2ac0cee271ed80004710b5ba****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The restoration priority. This parameter is required only for Cold Archive media assets. If you do not specify this parameter, the default value **Standard*	- is used. Valid values:
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
	// The scope of the modification. If you do not specify this parameter, the default value **All*	- is used. Valid values:
	//
	// - **All*	- (default): applies tiered storage to all resources of the media asset, including the source file and transcoded streams.
	//
	// - **SourceFile**: applies tiered storage only to the video source file of the media asset ID. Resources other than the source file are stored in Standard storage.
	//
	// example:
	//
	// All
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The storage class. Valid values:
	//
	// - **Standard**: Standard
	//
	// - **IA**: Infrequent Access (IA)
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
