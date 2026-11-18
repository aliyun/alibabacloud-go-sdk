// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialType(v string) *CreateUploadMediaRequest
	GetCredentialType() *string
	SetEntityId(v string) *CreateUploadMediaRequest
	GetEntityId() *string
	SetFileInfo(v string) *CreateUploadMediaRequest
	GetFileInfo() *string
	SetMediaId(v string) *CreateUploadMediaRequest
	GetMediaId() *string
	SetMediaMetaData(v string) *CreateUploadMediaRequest
	GetMediaMetaData() *string
	SetOverwrite(v bool) *CreateUploadMediaRequest
	GetOverwrite() *bool
	SetProjectId(v string) *CreateUploadMediaRequest
	GetProjectId() *string
	SetStorageLocation(v string) *CreateUploadMediaRequest
	GetStorageLocation() *string
	SetUserData(v string) *CreateUploadMediaRequest
	GetUserData() *string
}

type CreateUploadMediaRequest struct {
	// The type of the upload credential. Valid values:
	//
	// - `STS_TOKEN`: the default value. Direct upload by using the OSS SDK.
	//
	// - `POST_POLICY`: form-based upload.
	//
	// - `PRESIGNED_URL`: pre-signed PUT URL.
	//
	// - `UPLOAD_TICKET`: ticket mode. A real credential is issued only when the ticket is exchanged.
	//
	// example:
	//
	// STS_TOKEN
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The entity ID. The default value is AiSaasMediaAsset. You can ignore this parameter in common calls.
	//
	// example:
	//
	// AiSaasMediaAsset
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The file information, which is a JSON string. This parameter is required when you create a media asset. The **Type*	- field is required and can be set to video, audio, or image. The **Name*	- (file name), **Ext*	- (file name extension without the period), and **Size*	- (file size in bytes) fields are optional. This parameter can be omitted when you overwrite an existing media asset (Overwrite=true). If this parameter is omitted, the original media asset type and source file name extension are used.
	//
	// example:
	//
	// {"Type":"video","Name":"demo.mp4","Ext":"mp4","Size":10485760}
	FileInfo *string `json:"FileInfo,omitempty" xml:"FileInfo,omitempty"`
	// The ID of the existing media asset. This parameter is required when Overwrite is set to true (overwrite mode). If CredentialType is set to UPLOAD_TICKET, this parameter indicates a ticket exchange. This parameter is not required in other credential modes.
	//
	// example:
	//
	// b7e94fc0e*****6e7c5486601
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The metadata of the media asset, which is a JSON string. This parameter is required when you create a media asset. It can contain the following fields: **Title*	- (the title; if not specified, FileInfo.Name is used), **Description**, **CateId*	- (the category ID), **Tags*	- (comma-separated tags), **BusinessType*	- (falls back to general if an invalid value is provided), **CoverURL**, **DynamicMetaData*	- (a JSON string, such as ProductionId), **ReferenceId*	- (the external reference ID), and **RegisterConfig*	- (an object that contains SearchLibName to specify the search library). This parameter can be omitted when you overwrite an existing media asset.
	//
	// example:
	//
	// {"Title":"demo","Description":"test","Tags":"test,demo","BusinessType":"general","DynamicMetaData":"{\\"ProductionId\\":\\"pd_123456\\"}"}
	MediaMetaData *string `json:"MediaMetaData,omitempty" xml:"MediaMetaData,omitempty"`
	// Specifies whether to overwrite the existing media asset. Default value: false. If this parameter is set to true, you must specify MediaId. The original media asset is reset (old files, snapshots, and analysis results are cleared, and the status is set to uploading), and a new credential is issued as if it were the first upload.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The project ID. This parameter has the highest priority and is written to dynamicMetaData.ProductionId, overwriting the value with the same name in DynamicMetaData. It is automatically paired with the default folder `fd__default`. This parameter is supported only by the default entity. If you specify this parameter for a non-default entity, an `InvalidParameter` error is returned. If this parameter is not specified, `ProductionId` is obtained from the existing fallback chain.
	//
	// example:
	//
	// pd_123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The specified storage location. The format is `bucket.endpoint`. If this parameter is not specified, the default bucket selection policy is used.
	//
	// example:
	//
	// yike-shared-ap-southeast-1-1.oss-ap-southeast-1.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The custom user data, which is a JSON string.
	//
	// example:
	//
	// {"bizId":"order-123"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateUploadMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadMediaRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadMediaRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateUploadMediaRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateUploadMediaRequest) GetFileInfo() *string {
	return s.FileInfo
}

func (s *CreateUploadMediaRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadMediaRequest) GetMediaMetaData() *string {
	return s.MediaMetaData
}

func (s *CreateUploadMediaRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *CreateUploadMediaRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateUploadMediaRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *CreateUploadMediaRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateUploadMediaRequest) SetCredentialType(v string) *CreateUploadMediaRequest {
	s.CredentialType = &v
	return s
}

func (s *CreateUploadMediaRequest) SetEntityId(v string) *CreateUploadMediaRequest {
	s.EntityId = &v
	return s
}

func (s *CreateUploadMediaRequest) SetFileInfo(v string) *CreateUploadMediaRequest {
	s.FileInfo = &v
	return s
}

func (s *CreateUploadMediaRequest) SetMediaId(v string) *CreateUploadMediaRequest {
	s.MediaId = &v
	return s
}

func (s *CreateUploadMediaRequest) SetMediaMetaData(v string) *CreateUploadMediaRequest {
	s.MediaMetaData = &v
	return s
}

func (s *CreateUploadMediaRequest) SetOverwrite(v bool) *CreateUploadMediaRequest {
	s.Overwrite = &v
	return s
}

func (s *CreateUploadMediaRequest) SetProjectId(v string) *CreateUploadMediaRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateUploadMediaRequest) SetStorageLocation(v string) *CreateUploadMediaRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadMediaRequest) SetUserData(v string) *CreateUploadMediaRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadMediaRequest) Validate() error {
	return dara.Validate(s)
}
