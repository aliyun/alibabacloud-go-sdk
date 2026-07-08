// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAppInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateCloudAppInfoShrinkRequest
	GetAppId() *string
	SetDescription(v string) *UpdateCloudAppInfoShrinkRequest
	GetDescription() *string
	SetPatchShrink(v string) *UpdateCloudAppInfoShrinkRequest
	GetPatchShrink() *string
	SetPkgLabelsShrink(v string) *UpdateCloudAppInfoShrinkRequest
	GetPkgLabelsShrink() *string
	SetStablePatchId(v string) *UpdateCloudAppInfoShrinkRequest
	GetStablePatchId() *string
}

type UpdateCloudAppInfoShrinkRequest struct {
	// The ID of the cloud application, which corresponds to a unique application package.
	//
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// 用于测试使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Information about the patch package to upload.
	//
	// 1. This parameter is not supported when PkgType is android.
	//
	// 2. For the same AppId, only one patch can be in the process of uploading at a time. This means only one patch can be in a state other than its desired state.
	PatchShrink *string `json:"Patch,omitempty" xml:"Patch,omitempty"`
	// The tags for the cloud application. You can select multiple tags. This action resets all existing tags for the cloud application.
	//
	// 1. Valid values:
	//
	//    hot, game, and app.
	//
	// 2. Special case:
	//
	//    To delete all tags, enter ["NULL"].
	PkgLabelsShrink *string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty"`
	// The ID of the stable patch. This patch is used by default if you do not specify a PatchId when the application is in use, such as during a session startup. This parameter is not supported when PkgType is android.
	//
	// Special value:
	//
	// 1. If you set this parameter to origin, the patch version is removed and the initial version is used.
	//
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	StablePatchId *string `json:"StablePatchId,omitempty" xml:"StablePatchId,omitempty"`
}

func (s UpdateCloudAppInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAppInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudAppInfoShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCloudAppInfoShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCloudAppInfoShrinkRequest) GetPatchShrink() *string {
	return s.PatchShrink
}

func (s *UpdateCloudAppInfoShrinkRequest) GetPkgLabelsShrink() *string {
	return s.PkgLabelsShrink
}

func (s *UpdateCloudAppInfoShrinkRequest) GetStablePatchId() *string {
	return s.StablePatchId
}

func (s *UpdateCloudAppInfoShrinkRequest) SetAppId(v string) *UpdateCloudAppInfoShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCloudAppInfoShrinkRequest) SetDescription(v string) *UpdateCloudAppInfoShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateCloudAppInfoShrinkRequest) SetPatchShrink(v string) *UpdateCloudAppInfoShrinkRequest {
	s.PatchShrink = &v
	return s
}

func (s *UpdateCloudAppInfoShrinkRequest) SetPkgLabelsShrink(v string) *UpdateCloudAppInfoShrinkRequest {
	s.PkgLabelsShrink = &v
	return s
}

func (s *UpdateCloudAppInfoShrinkRequest) SetStablePatchId(v string) *UpdateCloudAppInfoShrinkRequest {
	s.StablePatchId = &v
	return s
}

func (s *UpdateCloudAppInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
