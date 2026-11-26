// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAppPluginVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangelog(v string) *UploadAppPluginVersionRequest
	GetChangelog() *string
	SetDescription(v string) *UploadAppPluginVersionRequest
	GetDescription() *string
	SetDownloadUrl(v string) *UploadAppPluginVersionRequest
	GetDownloadUrl() *string
	SetExtend(v string) *UploadAppPluginVersionRequest
	GetExtend() *string
	SetOssKey(v string) *UploadAppPluginVersionRequest
	GetOssKey() *string
	SetPluginId(v string) *UploadAppPluginVersionRequest
	GetPluginId() *string
	SetPluginVersion(v string) *UploadAppPluginVersionRequest
	GetPluginVersion() *string
}

type UploadAppPluginVersionRequest struct {
	// The changelog.
	//
	// example:
	//
	// Fixed an issue with reading plug-in configurations
	Changelog *string `json:"Changelog,omitempty" xml:"Changelog,omitempty"`
	// The description. This parameter is optional. If this parameter is not empty, the description in the main table is synchronously updated.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The downloadable ZIP link. This parameter is mutually exclusive with OssKey.
	//
	// example:
	//
	// https://sfile.chatglm.cn/chatglm4/file/a8/a87340a5ba.pdf
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The extended JSON. This parameter is optional and uses the same format as the extend synchronized to the sandbox. If this parameter is not empty, the extend in the main table is synchronously updated.
	//
	// example:
	//
	// {"deliveryNodeName":"Delivery quality inspection","deliveryNodeStatus":"Finish","deliveryOperatorRole":"Provider"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The Object Storage Service (OSS) key after frontend direct upload. This parameter is mutually exclusive with DownloadUrl.
	//
	// example:
	//
	// backend/detection/objects/r-0008ujvfksltf5j45n81/task-000hckiuwyau0gwn17vq.jpg
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The ID of the gateway plug-in.
	//
	// example:
	//
	// aliplayer-react
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The new version number in SemVer format. The value must be greater than the current version.
	//
	// example:
	//
	// v2
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
}

func (s UploadAppPluginVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAppPluginVersionRequest) GoString() string {
	return s.String()
}

func (s *UploadAppPluginVersionRequest) GetChangelog() *string {
	return s.Changelog
}

func (s *UploadAppPluginVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadAppPluginVersionRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *UploadAppPluginVersionRequest) GetExtend() *string {
	return s.Extend
}

func (s *UploadAppPluginVersionRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *UploadAppPluginVersionRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UploadAppPluginVersionRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *UploadAppPluginVersionRequest) SetChangelog(v string) *UploadAppPluginVersionRequest {
	s.Changelog = &v
	return s
}

func (s *UploadAppPluginVersionRequest) SetDescription(v string) *UploadAppPluginVersionRequest {
	s.Description = &v
	return s
}

func (s *UploadAppPluginVersionRequest) SetDownloadUrl(v string) *UploadAppPluginVersionRequest {
	s.DownloadUrl = &v
	return s
}

func (s *UploadAppPluginVersionRequest) SetExtend(v string) *UploadAppPluginVersionRequest {
	s.Extend = &v
	return s
}

func (s *UploadAppPluginVersionRequest) SetOssKey(v string) *UploadAppPluginVersionRequest {
	s.OssKey = &v
	return s
}

func (s *UploadAppPluginVersionRequest) SetPluginId(v string) *UploadAppPluginVersionRequest {
	s.PluginId = &v
	return s
}

func (s *UploadAppPluginVersionRequest) SetPluginVersion(v string) *UploadAppPluginVersionRequest {
	s.PluginVersion = &v
	return s
}

func (s *UploadAppPluginVersionRequest) Validate() error {
	return dara.Validate(s)
}
