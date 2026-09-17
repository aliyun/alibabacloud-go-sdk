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
	SetPkgLabels(v []*string) *UpdateCloudAppInfoRequest
	GetPkgLabels() []*string
	SetStablePatchId(v string) *UpdateCloudAppInfoRequest
	GetStablePatchId() *string
}

type UpdateCloudAppInfoRequest struct {
	// The cloud application ID, which corresponds to a unique application package.
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
	// For testing purposes
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the patch package to upload.
	//
	// 1. Not supported when PkgType is set to android.
	//
	// 2. Only one patch can be in the uploading state at a time for the same AppId (only one patch in a non-final state is allowed per AppId).
	Patch *UpdateCloudAppInfoRequestPatch `json:"Patch,omitempty" xml:"Patch,omitempty" type:"Struct"`
	// The cloud application labels. You can select multiple labels. This operation resets the cloud application labels.
	//
	// 1. Valid values:
	//
	//   a. hot
	//
	//   b. game
	//
	//   c. app
	//
	// 2. Special cases:
	//
	//   a. To delete all labels, set this parameter to ["NULL"].
	PkgLabels []*string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty" type:"Repeated"`
	// The stable PatchId. When a PatchId is not specified during business operations (such as session startup), this PatchId is used by default. Not supported when PkgType is set to android.
	//
	// Special values:
	//
	// 1. origin: cancels the patch version and uses the initial version by default.
	//
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

func (s *UpdateCloudAppInfoRequest) GetPkgLabels() []*string {
	return s.PkgLabels
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

func (s *UpdateCloudAppInfoRequest) SetPkgLabels(v []*string) *UpdateCloudAppInfoRequest {
	s.PkgLabels = v
	return s
}

func (s *UpdateCloudAppInfoRequest) SetStablePatchId(v string) *UpdateCloudAppInfoRequest {
	s.StablePatchId = &v
	return s
}

func (s *UpdateCloudAppInfoRequest) Validate() error {
	if s.Patch != nil {
		if err := s.Patch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCloudAppInfoRequestPatch struct {
	// Specifies whether to automatically set the patch as the stable patch after a successful upload. Default value: false.
	//
	// example:
	//
	// false
	AsStablePatch *bool `json:"AsStablePatch,omitempty" xml:"AsStablePatch,omitempty"`
	// The download URL of the patch package.
	//
	// Either RenderingInstanceId or DownloadURL is required. DownloadURL takes priority.
	//
	// example:
	//
	// https://test_host/app/test-tar-pkg.tar
	DownloadURL *string `json:"DownloadURL,omitempty" xml:"DownloadURL,omitempty"`
	// The MD5 hash of the patch package, used for integrity verification. Valid only when DownloadURL is not empty. Required when DownloadURL is not empty.
	//
	// example:
	//
	// 346f6404395adfg5bae1e45g4e943bf7
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name or description of the patch package, which serves as a unique identifier under the AppId.
	//
	// Naming conventions:
	//
	// 1. Cannot be set to origin or all.
	//
	// 2. Must be 1 to 50 characters in length.
	//
	// 3. Can contain lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 4. Must start and end with a letter or digit.
	//
	// example:
	//
	// p1
	PatchName *string `json:"PatchName,omitempty" xml:"PatchName,omitempty"`
	// The format of the installation package. The default value is the file extension of the download URL. Valid only when DownloadURL is not empty. Valid values:
	//
	// 1. tar.gz
	//
	// 2. tar
	//
	// 3. zip
	//
	// 4. rar
	//
	// example:
	//
	// tar
	PkgFormat *string `json:"PkgFormat,omitempty" xml:"PkgFormat,omitempty"`
	// The relative path of the post-command within the application package. Only supported for Windows applications.
	//
	// example:
	//
	// install.ps1
	PostCommandPath *string `json:"PostCommandPath,omitempty" xml:"PostCommandPath,omitempty"`
	// The timeout period for the post-command execution, in seconds. Only supported for Windows applications.
	//
	// example:
	//
	// 10
	PostCommandTimeoutSec *int32 `json:"PostCommandTimeoutSec,omitempty" xml:"PostCommandTimeoutSec,omitempty"`
	// The instance ID of the instance used to create the patch package. Valid only for Android application marketplace scenarios (PkgType=andrpid_appmarket). Either RenderingInstanceId or DownloadURL is required. DownloadURL takes priority.
	//
	// example:
	//
	// render-d7ec79fe47ce47aca2d8d7500d25a28a
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpdateCloudAppInfoRequestPatch) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAppInfoRequestPatch) GoString() string {
	return s.String()
}

func (s *UpdateCloudAppInfoRequestPatch) GetAsStablePatch() *bool {
	return s.AsStablePatch
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

func (s *UpdateCloudAppInfoRequestPatch) GetPkgFormat() *string {
	return s.PkgFormat
}

func (s *UpdateCloudAppInfoRequestPatch) GetPostCommandPath() *string {
	return s.PostCommandPath
}

func (s *UpdateCloudAppInfoRequestPatch) GetPostCommandTimeoutSec() *int32 {
	return s.PostCommandTimeoutSec
}

func (s *UpdateCloudAppInfoRequestPatch) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpdateCloudAppInfoRequestPatch) SetAsStablePatch(v bool) *UpdateCloudAppInfoRequestPatch {
	s.AsStablePatch = &v
	return s
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

func (s *UpdateCloudAppInfoRequestPatch) SetPkgFormat(v string) *UpdateCloudAppInfoRequestPatch {
	s.PkgFormat = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) SetPostCommandPath(v string) *UpdateCloudAppInfoRequestPatch {
	s.PostCommandPath = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) SetPostCommandTimeoutSec(v int32) *UpdateCloudAppInfoRequestPatch {
	s.PostCommandTimeoutSec = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) SetRenderingInstanceId(v string) *UpdateCloudAppInfoRequestPatch {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) Validate() error {
	return dara.Validate(s)
}
