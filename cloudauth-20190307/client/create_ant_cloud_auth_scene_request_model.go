// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntCloudAuthSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindMiniProgram(v string) *CreateAntCloudAuthSceneRequest
	GetBindMiniProgram() *string
	SetCheckFileBody(v string) *CreateAntCloudAuthSceneRequest
	GetCheckFileBody() *string
	SetCheckFileName(v string) *CreateAntCloudAuthSceneRequest
	GetCheckFileName() *string
	SetMiniProgramName(v string) *CreateAntCloudAuthSceneRequest
	GetMiniProgramName() *string
	SetPlatform(v string) *CreateAntCloudAuthSceneRequest
	GetPlatform() *string
	SetReturnPicCount(v int64) *CreateAntCloudAuthSceneRequest
	GetReturnPicCount() *int64
	SetReturnVideoLength(v int64) *CreateAntCloudAuthSceneRequest
	GetReturnVideoLength() *int64
	SetSceneName(v string) *CreateAntCloudAuthSceneRequest
	GetSceneName() *string
	SetStoreImage(v string) *CreateAntCloudAuthSceneRequest
	GetStoreImage() *string
}

type CreateAntCloudAuthSceneRequest struct {
	// Whether to enable binding of the mini program:
	//
	// - **Y**: Enable
	//
	// - **N (default)**: Not enabled
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
	// Binding platform for the mini program:
	//
	// - **WECHAT**: WeChat
	//
	// - **ALIPAY**: Alipay
	//
	// - **TIKTOK**: TikTok
	//
	// example:
	//
	// WECHAT
	Platform          *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ReturnPicCount    *int64  `json:"ReturnPicCount,omitempty" xml:"ReturnPicCount,omitempty"`
	ReturnVideoLength *int64  `json:"ReturnVideoLength,omitempty" xml:"ReturnVideoLength,omitempty"`
	// Scene name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试场景
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Whether to deliver the files generated from the authentication to the customer\\"s OSS:
	//
	// - **Y**: Yes
	//
	// - **N**: No
	//
	// example:
	//
	// Y
	StoreImage *string `json:"StoreImage,omitempty" xml:"StoreImage,omitempty"`
}

func (s CreateAntCloudAuthSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAntCloudAuthSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateAntCloudAuthSceneRequest) GetBindMiniProgram() *string {
	return s.BindMiniProgram
}

func (s *CreateAntCloudAuthSceneRequest) GetCheckFileBody() *string {
	return s.CheckFileBody
}

func (s *CreateAntCloudAuthSceneRequest) GetCheckFileName() *string {
	return s.CheckFileName
}

func (s *CreateAntCloudAuthSceneRequest) GetMiniProgramName() *string {
	return s.MiniProgramName
}

func (s *CreateAntCloudAuthSceneRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateAntCloudAuthSceneRequest) GetReturnPicCount() *int64 {
	return s.ReturnPicCount
}

func (s *CreateAntCloudAuthSceneRequest) GetReturnVideoLength() *int64 {
	return s.ReturnVideoLength
}

func (s *CreateAntCloudAuthSceneRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *CreateAntCloudAuthSceneRequest) GetStoreImage() *string {
	return s.StoreImage
}

func (s *CreateAntCloudAuthSceneRequest) SetBindMiniProgram(v string) *CreateAntCloudAuthSceneRequest {
	s.BindMiniProgram = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetCheckFileBody(v string) *CreateAntCloudAuthSceneRequest {
	s.CheckFileBody = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetCheckFileName(v string) *CreateAntCloudAuthSceneRequest {
	s.CheckFileName = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetMiniProgramName(v string) *CreateAntCloudAuthSceneRequest {
	s.MiniProgramName = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetPlatform(v string) *CreateAntCloudAuthSceneRequest {
	s.Platform = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetReturnPicCount(v int64) *CreateAntCloudAuthSceneRequest {
	s.ReturnPicCount = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetReturnVideoLength(v int64) *CreateAntCloudAuthSceneRequest {
	s.ReturnVideoLength = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetSceneName(v string) *CreateAntCloudAuthSceneRequest {
	s.SceneName = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetStoreImage(v string) *CreateAntCloudAuthSceneRequest {
	s.StoreImage = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) Validate() error {
	return dara.Validate(s)
}
