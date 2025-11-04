// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *CreateVodPackagingAssetRequest
	GetAssetName() *string
	SetContentId(v string) *CreateVodPackagingAssetRequest
	GetContentId() *string
	SetDescription(v string) *CreateVodPackagingAssetRequest
	GetDescription() *string
	SetGroupName(v string) *CreateVodPackagingAssetRequest
	GetGroupName() *string
	SetInput(v *CreateVodPackagingAssetRequestInput) *CreateVodPackagingAssetRequest
	GetInput() *CreateVodPackagingAssetRequestInput
}

type CreateVodPackagingAssetRequest struct {
	// The name of the asset. The name must be unique and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// hls_3s
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The content ID in the digital rights management (DRM) system. The maximum length is 256 characters. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// movie
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// The asset description.
	//
	// example:
	//
	// HLS 3 second packaging
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The asset input configurations.
	Input *CreateVodPackagingAssetRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s CreateVodPackagingAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingAssetRequest) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingAssetRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *CreateVodPackagingAssetRequest) GetContentId() *string {
	return s.ContentId
}

func (s *CreateVodPackagingAssetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVodPackagingAssetRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateVodPackagingAssetRequest) GetInput() *CreateVodPackagingAssetRequestInput {
	return s.Input
}

func (s *CreateVodPackagingAssetRequest) SetAssetName(v string) *CreateVodPackagingAssetRequest {
	s.AssetName = &v
	return s
}

func (s *CreateVodPackagingAssetRequest) SetContentId(v string) *CreateVodPackagingAssetRequest {
	s.ContentId = &v
	return s
}

func (s *CreateVodPackagingAssetRequest) SetDescription(v string) *CreateVodPackagingAssetRequest {
	s.Description = &v
	return s
}

func (s *CreateVodPackagingAssetRequest) SetGroupName(v string) *CreateVodPackagingAssetRequest {
	s.GroupName = &v
	return s
}

func (s *CreateVodPackagingAssetRequest) SetInput(v *CreateVodPackagingAssetRequestInput) *CreateVodPackagingAssetRequest {
	s.Input = v
	return s
}

func (s *CreateVodPackagingAssetRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVodPackagingAssetRequestInput struct {
	// The URL of the media file. Only M3U8 files stored in OSS are supported.
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The input type. Only Object Storage Service (OSS) is supported.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateVodPackagingAssetRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingAssetRequestInput) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingAssetRequestInput) GetMedia() *string {
	return s.Media
}

func (s *CreateVodPackagingAssetRequestInput) GetType() *string {
	return s.Type
}

func (s *CreateVodPackagingAssetRequestInput) SetMedia(v string) *CreateVodPackagingAssetRequestInput {
	s.Media = &v
	return s
}

func (s *CreateVodPackagingAssetRequestInput) SetType(v string) *CreateVodPackagingAssetRequestInput {
	s.Type = &v
	return s
}

func (s *CreateVodPackagingAssetRequestInput) Validate() error {
	return dara.Validate(s)
}
