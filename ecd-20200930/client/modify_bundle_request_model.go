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
	// The cloud computer template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-2g65ljy4291vl****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The new cloud computer template name.
	//
	// example:
	//
	// newName
	BundleName *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	// The new cloud computer template description.
	//
	// example:
	//
	// newDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new image ID. The new image must meet the following conditions:
	//
	// - The new image must be in the Available state.
	//
	// - The new image must have the same operating system as the original image.
	//
	// - The disk size required by the new image cannot be larger than that of the original image.
	//
	// - The GPU type of the new image must be the same as that of the original image.
	//
	// example:
	//
	// m-aea3oaww001np****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The operating system language. Currently, only system images are supported.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
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
