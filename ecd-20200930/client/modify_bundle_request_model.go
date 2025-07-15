// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBundleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v string) *ModifyBundleRequest
	GetBundleId() *string
	SetBundleName(v string) *ModifyBundleRequest
	GetBundleName() *string
	SetDescription(v string) *ModifyBundleRequest
	GetDescription() *string
	SetImageId(v string) *ModifyBundleRequest
	GetImageId() *string
	SetLanguage(v string) *ModifyBundleRequest
	GetLanguage() *string
	SetRegionId(v string) *ModifyBundleRequest
	GetRegionId() *string
}

type ModifyBundleRequest struct {
	// The ID of the cloud computer template that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-2g65ljy4291vl****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The name of the new cloud computer template.
	//
	// example:
	//
	// newName
	BundleName *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	// The description of the new cloud computer template.
	//
	// example:
	//
	// newDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new image ID. The new image must meet the following conditions:
	//
	// 	- The new image must be in the Available state.
	//
	// 	- The operating system of the new image must be the same as that of the original image.
	//
	// 	- The required disk size for the new image cannot be greater than that for the original image.
	//
	// 	- The GPU type of the new image must be the same as that of the original image.
	//
	// example:
	//
	// m-aea3oaww001np****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The OS language. This parameter is available only for system images.
	//
	// Valid values:
	//
	// 	- en-US: American English
	//
	// 	- zh-HK: Traditional Chinese (Hong Kong)
	//
	// 	- zh-CN: Simplified Chinese.
	//
	// 	- ja-JP: Japanese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyBundleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBundleRequest) GoString() string {
	return s.String()
}

func (s *ModifyBundleRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *ModifyBundleRequest) GetBundleName() *string {
	return s.BundleName
}

func (s *ModifyBundleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyBundleRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyBundleRequest) GetLanguage() *string {
	return s.Language
}

func (s *ModifyBundleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBundleRequest) SetBundleId(v string) *ModifyBundleRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleName(v string) *ModifyBundleRequest {
	s.BundleName = &v
	return s
}

func (s *ModifyBundleRequest) SetDescription(v string) *ModifyBundleRequest {
	s.Description = &v
	return s
}

func (s *ModifyBundleRequest) SetImageId(v string) *ModifyBundleRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyBundleRequest) SetLanguage(v string) *ModifyBundleRequest {
	s.Language = &v
	return s
}

func (s *ModifyBundleRequest) SetRegionId(v string) *ModifyBundleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBundleRequest) Validate() error {
	return dara.Validate(s)
}
