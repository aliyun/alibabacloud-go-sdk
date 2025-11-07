// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAntCloudAuthSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindMiniProgram(v string) *UpdateAntCloudAuthSceneRequest
	GetBindMiniProgram() *string
	SetCheckFileBody(v string) *UpdateAntCloudAuthSceneRequest
	GetCheckFileBody() *string
	SetCheckFileName(v string) *UpdateAntCloudAuthSceneRequest
	GetCheckFileName() *string
	SetMiniProgramName(v string) *UpdateAntCloudAuthSceneRequest
	GetMiniProgramName() *string
	SetPlatform(v string) *UpdateAntCloudAuthSceneRequest
	GetPlatform() *string
	SetSceneId(v int64) *UpdateAntCloudAuthSceneRequest
	GetSceneId() *int64
	SetSceneName(v string) *UpdateAntCloudAuthSceneRequest
	GetSceneName() *string
	SetStatus(v int32) *UpdateAntCloudAuthSceneRequest
	GetStatus() *int32
	SetStoreImage(v string) *UpdateAntCloudAuthSceneRequest
	GetStoreImage() *string
}

type UpdateAntCloudAuthSceneRequest struct {
	// Whether to enable binding to a mini program:
	//
	// - **Y**: Enable
	//
	// - **N (default)**: Do not enable
	//
	// 	Notice: If enabling the binding of a mini program, please ensure all parameters for the mini program are passed.
	//
	// example:
	//
	// Y
	BindMiniProgram *string `json:"BindMiniProgram,omitempty" xml:"BindMiniProgram,omitempty"`
	// Content of the uploaded verification file.
	//
	// example:
	//
	// 774c4aab45981ff4a86cde9255a11xxx
	CheckFileBody *string `json:"CheckFileBody,omitempty" xml:"CheckFileBody,omitempty"`
	// Name of the uploaded verification file.
	//
	// example:
	//
	// 测试.txt
	CheckFileName *string `json:"CheckFileName,omitempty" xml:"CheckFileName,omitempty"`
	// Mini program name.
	//
	// example:
	//
	// 测试APP
	MiniProgramName *string `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	// Binding mini program platform:
	//
	// - **WECHAT**: WeChat
	//
	// - **ALIPAY**: Alipay
	//
	// - **TIKTOK**: TikTok
	//
	// example:
	//
	// IOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// Scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000013372
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Scenario name.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Currently meaningless, can be omitted.
	//
	// example:
	//
	// -
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Whether to enable delivering the authentication files generated during the authentication process to the user\\"s OSS:
	//
	// - **Y**: Enable
	//
	// - **N (default)**: Disable
	//
	// example:
	//
	// Y
	StoreImage *string `json:"StoreImage,omitempty" xml:"StoreImage,omitempty"`
}

func (s UpdateAntCloudAuthSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAntCloudAuthSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateAntCloudAuthSceneRequest) GetBindMiniProgram() *string {
	return s.BindMiniProgram
}

func (s *UpdateAntCloudAuthSceneRequest) GetCheckFileBody() *string {
	return s.CheckFileBody
}

func (s *UpdateAntCloudAuthSceneRequest) GetCheckFileName() *string {
	return s.CheckFileName
}

func (s *UpdateAntCloudAuthSceneRequest) GetMiniProgramName() *string {
	return s.MiniProgramName
}

func (s *UpdateAntCloudAuthSceneRequest) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateAntCloudAuthSceneRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *UpdateAntCloudAuthSceneRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *UpdateAntCloudAuthSceneRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateAntCloudAuthSceneRequest) GetStoreImage() *string {
	return s.StoreImage
}

func (s *UpdateAntCloudAuthSceneRequest) SetBindMiniProgram(v string) *UpdateAntCloudAuthSceneRequest {
	s.BindMiniProgram = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetCheckFileBody(v string) *UpdateAntCloudAuthSceneRequest {
	s.CheckFileBody = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetCheckFileName(v string) *UpdateAntCloudAuthSceneRequest {
	s.CheckFileName = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetMiniProgramName(v string) *UpdateAntCloudAuthSceneRequest {
	s.MiniProgramName = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetPlatform(v string) *UpdateAntCloudAuthSceneRequest {
	s.Platform = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetSceneId(v int64) *UpdateAntCloudAuthSceneRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetSceneName(v string) *UpdateAntCloudAuthSceneRequest {
	s.SceneName = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetStatus(v int32) *UpdateAntCloudAuthSceneRequest {
	s.Status = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetStoreImage(v string) *UpdateAntCloudAuthSceneRequest {
	s.StoreImage = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) Validate() error {
	return dara.Validate(s)
}
