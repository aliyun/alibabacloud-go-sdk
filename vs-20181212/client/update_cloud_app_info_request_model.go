// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateCloudAppInfoRequest
	GetAppId() *string
	SetDescription(v string) *UpdateCloudAppInfoRequest
	GetDescription() *string
	SetPatch(v *UpdateCloudAppInfoRequestPatch) *UpdateCloudAppInfoRequest
	GetPatch() *UpdateCloudAppInfoRequestPatch
	SetStablePatchId(v string) *UpdateCloudAppInfoRequest
	GetStablePatchId() *string
}

type UpdateCloudAppInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId       *string                         `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Patch       *UpdateCloudAppInfoRequestPatch `json:"Patch,omitempty" xml:"Patch,omitempty" type:"Struct"`
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	StablePatchId *string `json:"StablePatchId,omitempty" xml:"StablePatchId,omitempty"`
}

func (s UpdateCloudAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAppInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudAppInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCloudAppInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCloudAppInfoRequest) GetPatch() *UpdateCloudAppInfoRequestPatch {
	return s.Patch
}

func (s *UpdateCloudAppInfoRequest) GetStablePatchId() *string {
	return s.StablePatchId
}

func (s *UpdateCloudAppInfoRequest) SetAppId(v string) *UpdateCloudAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCloudAppInfoRequest) SetDescription(v string) *UpdateCloudAppInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateCloudAppInfoRequest) SetPatch(v *UpdateCloudAppInfoRequestPatch) *UpdateCloudAppInfoRequest {
	s.Patch = v
	return s
}

func (s *UpdateCloudAppInfoRequest) SetStablePatchId(v string) *UpdateCloudAppInfoRequest {
	s.StablePatchId = &v
	return s
}

func (s *UpdateCloudAppInfoRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudAppInfoRequestPatch struct {
	// example:
	//
	// https://test_host/app/test-tar-pkg.tar
	DownloadURL *string `json:"DownloadURL,omitempty" xml:"DownloadURL,omitempty"`
	// example:
	//
	// 346f6404395adfg5bae1e45g4e943bf7
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// p1
	PatchName *string `json:"PatchName,omitempty" xml:"PatchName,omitempty"`
}

func (s UpdateCloudAppInfoRequestPatch) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAppInfoRequestPatch) GoString() string {
	return s.String()
}

func (s *UpdateCloudAppInfoRequestPatch) GetDownloadURL() *string {
	return s.DownloadURL
}

func (s *UpdateCloudAppInfoRequestPatch) GetMd5() *string {
	return s.Md5
}

func (s *UpdateCloudAppInfoRequestPatch) GetPatchName() *string {
	return s.PatchName
}

func (s *UpdateCloudAppInfoRequestPatch) SetDownloadURL(v string) *UpdateCloudAppInfoRequestPatch {
	s.DownloadURL = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) SetMd5(v string) *UpdateCloudAppInfoRequestPatch {
	s.Md5 = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) SetPatchName(v string) *UpdateCloudAppInfoRequestPatch {
	s.PatchName = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) Validate() error {
	return dara.Validate(s)
}
