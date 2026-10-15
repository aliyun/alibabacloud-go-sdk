// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateAppPluginRequest
	GetCategory() *string
	SetDescription(v string) *CreateAppPluginRequest
	GetDescription() *string
	SetDownloadUrl(v string) *CreateAppPluginRequest
	GetDownloadUrl() *string
	SetExtend(v string) *CreateAppPluginRequest
	GetExtend() *string
	SetIcon(v string) *CreateAppPluginRequest
	GetIcon() *string
	SetOssKey(v string) *CreateAppPluginRequest
	GetOssKey() *string
	SetPluginId(v string) *CreateAppPluginRequest
	GetPluginId() *string
	SetPluginName(v string) *CreateAppPluginRequest
	GetPluginName() *string
	SetPluginVersion(v string) *CreateAppPluginRequest
	GetPluginVersion() *string
	SetTags(v string) *CreateAppPluginRequest
	GetTags() *string
	SetVisibility(v string) *CreateAppPluginRequest
	GetVisibility() *string
}

type CreateAppPluginRequest struct {
	// The category of the plug-in.
	//
	// example:
	//
	// {\\"Depth\\": 4, \\"ParentCategoryId\\": 5281, \\"Name\\": \\"1\\", \\"CategoryId\\": 5282, \\"CreateTime\\": \\"2025-08-26 10:20:37\\"}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the plug-in. The value must be 10 to 512 characters in length.
	//
	// example:
	//
	// 123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The downloadable ZIP URL. This parameter is mutually exclusive with OssKey.
	//
	// example:
	//
	// http://geoip-sdk-user.oss-cn-zhangjiakou.aliyuncs.com/product/v1/ipv6-inner/common/v1.20230608104522.dex?Expires=1753324652&OSSAccessKeyId=x6R87i97yKeKRggz&Signature=IwS6V56TlgjcX0AgikXLvf3CVtw%3D
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The extended JSON. This parameter is optional and uses the same format as the extend field in sandbox synchronization. If this parameter is not empty, the value is written to the database.
	//
	// example:
	//
	// {\\"user_device_id\\":\\"6bef45cb0c76de284d24de074c088b73\\"}\\n
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The URL of the icon.
	//
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The OSS key after frontend direct upload. This parameter is mutually exclusive with DownloadUrl.
	//
	// example:
	//
	// backend/detection/objects/r-0008ujvfksltf5j45n81/task-000hckiuwyau0gwn17vq.jpg
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The unique identifier of the plug-in. The value can contain lowercase letters, digits, and hyphens (-), and must be 3 to 64 characters in length.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The name of the plug-in. The value must be 2 to 64 characters in length.
	//
	// example:
	//
	// alisecguard
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The version number in semver format. Default value: 0.1.0.
	//
	// example:
	//
	// 1.1.6
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
	// The tags, separated by commas (,).
	//
	// example:
	//
	// []
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The visibility of the plug-in. Default value: private.
	//
	// example:
	//
	// private
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateAppPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPluginRequest) GoString() string {
	return s.String()
}

func (s *CreateAppPluginRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateAppPluginRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppPluginRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateAppPluginRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateAppPluginRequest) GetIcon() *string {
	return s.Icon
}

func (s *CreateAppPluginRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateAppPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *CreateAppPluginRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *CreateAppPluginRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *CreateAppPluginRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateAppPluginRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateAppPluginRequest) SetCategory(v string) *CreateAppPluginRequest {
	s.Category = &v
	return s
}

func (s *CreateAppPluginRequest) SetDescription(v string) *CreateAppPluginRequest {
	s.Description = &v
	return s
}

func (s *CreateAppPluginRequest) SetDownloadUrl(v string) *CreateAppPluginRequest {
	s.DownloadUrl = &v
	return s
}

func (s *CreateAppPluginRequest) SetExtend(v string) *CreateAppPluginRequest {
	s.Extend = &v
	return s
}

func (s *CreateAppPluginRequest) SetIcon(v string) *CreateAppPluginRequest {
	s.Icon = &v
	return s
}

func (s *CreateAppPluginRequest) SetOssKey(v string) *CreateAppPluginRequest {
	s.OssKey = &v
	return s
}

func (s *CreateAppPluginRequest) SetPluginId(v string) *CreateAppPluginRequest {
	s.PluginId = &v
	return s
}

func (s *CreateAppPluginRequest) SetPluginName(v string) *CreateAppPluginRequest {
	s.PluginName = &v
	return s
}

func (s *CreateAppPluginRequest) SetPluginVersion(v string) *CreateAppPluginRequest {
	s.PluginVersion = &v
	return s
}

func (s *CreateAppPluginRequest) SetTags(v string) *CreateAppPluginRequest {
	s.Tags = &v
	return s
}

func (s *CreateAppPluginRequest) SetVisibility(v string) *CreateAppPluginRequest {
	s.Visibility = &v
	return s
}

func (s *CreateAppPluginRequest) Validate() error {
	return dara.Validate(s)
}
