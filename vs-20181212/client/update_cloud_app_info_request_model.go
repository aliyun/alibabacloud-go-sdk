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
	Patch *UpdateCloudAppInfoRequestPatch `json:"Patch,omitempty" xml:"Patch,omitempty" type:"Struct"`
	// The tags for the cloud application. You can select multiple tags. This action resets all existing tags for the cloud application.
	//
	// 1. Valid values:
	//
	//    hot, game, and app.
	//
	// 2. Special case:
	//
	//    To delete all tags, enter ["NULL"].
	PkgLabels []*string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty" type:"Repeated"`
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
	// Specifies whether to automatically set the patch as the stable version after it is successfully uploaded. The default value is false.
	//
	// example:
	//
	// false
	AsStablePatch *bool `json:"AsStablePatch,omitempty" xml:"AsStablePatch,omitempty"`
	// The download URL for the patch package.
	//
	// You must specify either RenderingInstanceId or DownloadURL.
	//
	// DownloadURL takes precedence.
	//
	// example:
	//
	// https://test_host/app/test-tar-pkg.tar
	DownloadURL *string `json:"DownloadURL,omitempty" xml:"DownloadURL,omitempty"`
	// The MD5 hash of the patch package, used to verify integrity. This parameter is valid only if DownloadURL is not empty. It is required if DownloadURL is not empty.
	//
	// example:
	//
	// 346f6404395adfg5bae1e45g4e943bf7
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name or description of the patch package. This is a unique identifier under the AppId.
	//
	// Default naming conventions:
	//
	// 1. Cannot be origin or all.
	//
	// 2. Must be 1 to 50 characters in length.
	//
	// 3. Can contain lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 4. The first and last characters must be a letter or a digit.
	//
	// example:
	//
	// p1
	PatchName *string `json:"PatchName,omitempty" xml:"PatchName,omitempty"`
	// The format of the installation package. By default, the system uses the file extension from the download URL. This parameter is valid only if DownloadURL is not empty. Valid values:
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
	// The instance ID required to create the patch package. This parameter is valid only in the Android application marketplace scenario (PkgType=andrpid_appmarket). Specify either RenderingInstanceId or DownloadURL. DownloadURL takes precedence.
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

func (s *UpdateCloudAppInfoRequestPatch) SetRenderingInstanceId(v string) *UpdateCloudAppInfoRequestPatch {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpdateCloudAppInfoRequestPatch) Validate() error {
	return dara.Validate(s)
}
