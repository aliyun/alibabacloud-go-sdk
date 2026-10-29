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
	// Used for testing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the patch package to be uploaded. 1. This parameter is not supported when PkgType is set to android. 2. Only one patch package can be in the uploading state (non-final state) for the same AppId at a time.
	PatchShrink *string `json:"Patch,omitempty" xml:"Patch,omitempty"`
	// The tags of the cloud application. You can specify multiple tags to reset the cloud application tags. Valid values:
	//
	// - hot
	//
	// - game
	//
	// - app
	//
	// To delete all tags, set the value to ["NULL"].
	PkgLabelsShrink *string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty"`
	// The ID of the stable patch. This is the default patch ID used in customer business scenarios, such as session startup, when no patch ID is specified. This parameter is not supported when PkgType is set to android. Special value: origin: cancels the patch version and uses the initial version by default.
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
