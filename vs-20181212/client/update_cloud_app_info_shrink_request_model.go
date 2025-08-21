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
	SetStablePatchId(v string) *UpdateCloudAppInfoShrinkRequest
	GetStablePatchId() *string
}

type UpdateCloudAppInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PatchShrink *string `json:"Patch,omitempty" xml:"Patch,omitempty"`
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

func (s *UpdateCloudAppInfoShrinkRequest) SetStablePatchId(v string) *UpdateCloudAppInfoShrinkRequest {
	s.StablePatchId = &v
	return s
}

func (s *UpdateCloudAppInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
