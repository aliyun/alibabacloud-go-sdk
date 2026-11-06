// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptorIdList(v []*string) *UpdateImageRequest
	GetAdaptorIdList() []*string
	SetDescription(v string) *UpdateImageRequest
	GetDescription() *string
	SetDistro(v string) *UpdateImageRequest
	GetDistro() *string
	SetImageId(v string) *UpdateImageRequest
	GetImageId() *string
	SetImageName(v string) *UpdateImageRequest
	GetImageName() *string
	SetLanguage(v string) *UpdateImageRequest
	GetLanguage() *string
	SetVersionId(v string) *UpdateImageRequest
	GetVersionId() *string
	SetVersionName(v string) *UpdateImageRequest
	GetVersionName() *string
}

type UpdateImageRequest struct {
	// The complete list of application adapter IDs for the selected image version. A maximum of 100 items are supported. **When a non-empty list is passed, it must include the adapters that you still want to retain. Omitting this parameter or passing an empty list does not clear the existing associations.**
	AdaptorIdList []*string `json:"AdaptorIdList,omitempty" xml:"AdaptorIdList,omitempty" type:"Repeated"`
	// The updated description of the image.
	//
	// example:
	//
	// Update office software and basic configurations
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The operating system distribution information of the image. This must be consistent with the actual operating system distribution of the image.
	//
	// example:
	//
	// Ubuntu 22.04 64-bit
	Distro *string `json:"Distro,omitempty" xml:"Distro,omitempty"`
	// The ID of the image to be modified. You must provide at least one of ImageId and VersionId. If both are provided, they must belong to the same image.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The new name of the image. Specify this parameter only when you need to update the name.
	//
	// example:
	//
	// OfficeImage_202611
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The language of the image. Valid values:
	//
	// - **zh-CN**: Simplified Chinese.
	//
	// - **zh-HK**: Traditional Chinese (Hong Kong).
	//
	// - **en-US**: English.
	//
	// - **ja-JP**: Japanese.
	//
	// - **CN**: A compatible value for Simplified Chinese.
	//
	// - **EN**: A compatible value for English.
	//
	// - **JP**: A compatible value for Japanese.
	//
	// - **CN-HK**: A compatible value for Traditional Chinese (Hong Kong).
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the image version to be modified. If this parameter is omitted, the current effective version of the image specified by ImageId is used. You must provide at least one of ImageId and VersionId.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The new display name of the image version.
	//
	// example:
	//
	// Office environment-202611
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) GetAdaptorIdList() []*string {
	return s.AdaptorIdList
}

func (s *UpdateImageRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateImageRequest) GetDistro() *string {
	return s.Distro
}

func (s *UpdateImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *UpdateImageRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateImageRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateImageRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *UpdateImageRequest) SetAdaptorIdList(v []*string) *UpdateImageRequest {
	s.AdaptorIdList = v
	return s
}

func (s *UpdateImageRequest) SetDescription(v string) *UpdateImageRequest {
	s.Description = &v
	return s
}

func (s *UpdateImageRequest) SetDistro(v string) *UpdateImageRequest {
	s.Distro = &v
	return s
}

func (s *UpdateImageRequest) SetImageId(v string) *UpdateImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateImageRequest) SetImageName(v string) *UpdateImageRequest {
	s.ImageName = &v
	return s
}

func (s *UpdateImageRequest) SetLanguage(v string) *UpdateImageRequest {
	s.Language = &v
	return s
}

func (s *UpdateImageRequest) SetVersionId(v string) *UpdateImageRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateImageRequest) SetVersionName(v string) *UpdateImageRequest {
	s.VersionName = &v
	return s
}

func (s *UpdateImageRequest) Validate() error {
	return dara.Validate(s)
}
