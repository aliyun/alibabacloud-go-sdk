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
	// Specifies whether to change the storage class of a media asset that is stored for less than the minimum storage duration. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If you forcibly change the storage class of a media asset that is stored for less than the minimum storage duration, additional data retrieval fees are incurred.
	//
	// example:
	//
	// false
	AllowUpdateWithoutTimeLimit *bool `json:"AllowUpdateWithoutTimeLimit,omitempty" xml:"AllowUpdateWithoutTimeLimit,omitempty"`
	// The media asset ID. You can specify a maximum of 20 IDs. Separate multiple IDs with commas (,). You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, you can view the ID of the media asset. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of the VideoId parameter from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to upload media assets.
	//
	// 	- Obtain the value of the VideoId parameter from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you call to query the media ID after the media asset is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// d56c2ac0cee271ed80004710b5ba****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The restoration priority. This parameter is required only when you restore a Cold Archive media asset. Valid values:
	//
	// 	- **Expedited**
	//
	// 	- **Standard**
	//
	// 	- **Bulk**
	//
	// example:
	//
	// Standard
	RestoreTier *string `json:"RestoreTier,omitempty" xml:"RestoreTier,omitempty"`
	// The modification range. Valid values:
	//
	// 	- **All**: modifies the storage classes of all resources including the source files and transcoded streams.
	//
	// 	- **SourceFile**: modifies the storage classes of only the source files. The storage class of other resources is Standard.
	//
	// example:
	//
	// All
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The storage class. Valid values:
	//
	// 	- **Standard**
	//
	// 	- **IA**
	//
	// 	- **Archive**
	//
	// 	- **ColdArchive**
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
