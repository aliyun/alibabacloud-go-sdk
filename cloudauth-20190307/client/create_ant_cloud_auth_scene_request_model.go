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
	SetDegradeAppScheme(v string) *CreateAntCloudAuthSceneRequest
	GetDegradeAppScheme() *string
	SetDegradeSubCodes(v string) *CreateAntCloudAuthSceneRequest
	GetDegradeSubCodes() *string
	SetDegradeType(v string) *CreateAntCloudAuthSceneRequest
	GetDegradeType() *string
	SetDeviceRiskPlus(v string) *CreateAntCloudAuthSceneRequest
	GetDeviceRiskPlus() *string
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
	SetUseDegrade(v string) *CreateAntCloudAuthSceneRequest
	GetUseDegrade() *string
}

type CreateAntCloudAuthSceneRequest struct {
	// Specifies whether to enable mini program binding. Valid values:
	//
	// - **Y**: Enabled.
	//
	// - **N (default)**: Disabled.
	//
	// example:
	//
	// Y
	BindMiniProgram *string `json:"BindMiniProgram,omitempty" xml:"BindMiniProgram,omitempty"`
	// The content of the uploaded verification file.
	//
	// example:
	//
	// 774c4aab45981ff4a86cde9255a11xxx
	CheckFileBody *string `json:"CheckFileBody,omitempty" xml:"CheckFileBody,omitempty"`
	// The name of the uploaded verification file.
	//
	// example:
	//
	// test.txt
	CheckFileName *string `json:"CheckFileName,omitempty" xml:"CheckFileName,omitempty"`
	// The iOS app scheme for degradation redirect.
	//
	// example:
	//
	// cloudauth://callback
	DegradeAppScheme *string `json:"DegradeAppScheme,omitempty" xml:"DegradeAppScheme,omitempty"`
	// The SubCode that triggers degradation.
	//
	// example:
	//
	// 201,202
	DegradeSubCodes *string `json:"DegradeSubCodes,omitempty" xml:"DegradeSubCodes,omitempty"`
	// The degradation verification type.
	//
	// example:
	//
	// ALIPAY
	DegradeType *string `json:"DegradeType,omitempty" xml:"DegradeType,omitempty"`
	// Specifies whether to enable enhanced device risk detection. Valid values:
	//
	// - **Y**: Enabled.
	//
	// - **N**: Disabled.
	//
	// example:
	//
	// N
	DeviceRiskPlus *string `json:"DeviceRiskPlus,omitempty" xml:"DeviceRiskPlus,omitempty"`
	// The mini program name.
	//
	// example:
	//
	// TestApp
	MiniProgramName *string `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	// The mini program platform to bind. Valid values:
	//
	// - **WECHAT**: WeChat.
	//
	// - **ALIPAY**: Alipay.
	//
	// - **TIKTOK**: TikTok.
	//
	// example:
	//
	// WECHAT
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The number of face photos for evidence storage (1-5).
	//
	// example:
	//
	// 1
	ReturnPicCount *int64 `json:"ReturnPicCount,omitempty" xml:"ReturnPicCount,omitempty"`
	// The duration of the evidence storage video, in seconds.
	//
	// example:
	//
	// 1
	ReturnVideoLength *int64 `json:"ReturnVideoLength,omitempty" xml:"ReturnVideoLength,omitempty"`
	// The scenario name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestScenario
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Specifies whether to deliver files generated during verification to the customer\\"s OSS. Valid values:
	//
	// - **Y**: Yes.
	//
	// - **N**: No.
	//
	// example:
	//
	// Y
	StoreImage *string `json:"StoreImage,omitempty" xml:"StoreImage,omitempty"`
	// Specifies whether to enable degradation.
	//
	// example:
	//
	// Y
	UseDegrade *string `json:"UseDegrade,omitempty" xml:"UseDegrade,omitempty"`
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

func (s *CreateAntCloudAuthSceneRequest) GetDegradeAppScheme() *string {
	return s.DegradeAppScheme
}

func (s *CreateAntCloudAuthSceneRequest) GetDegradeSubCodes() *string {
	return s.DegradeSubCodes
}

func (s *CreateAntCloudAuthSceneRequest) GetDegradeType() *string {
	return s.DegradeType
}

func (s *CreateAntCloudAuthSceneRequest) GetDeviceRiskPlus() *string {
	return s.DeviceRiskPlus
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

func (s *CreateAntCloudAuthSceneRequest) GetUseDegrade() *string {
	return s.UseDegrade
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

func (s *CreateAntCloudAuthSceneRequest) SetDegradeAppScheme(v string) *CreateAntCloudAuthSceneRequest {
	s.DegradeAppScheme = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetDegradeSubCodes(v string) *CreateAntCloudAuthSceneRequest {
	s.DegradeSubCodes = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetDegradeType(v string) *CreateAntCloudAuthSceneRequest {
	s.DegradeType = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) SetDeviceRiskPlus(v string) *CreateAntCloudAuthSceneRequest {
	s.DeviceRiskPlus = &v
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

func (s *CreateAntCloudAuthSceneRequest) SetUseDegrade(v string) *CreateAntCloudAuthSceneRequest {
	s.UseDegrade = &v
	return s
}

func (s *CreateAntCloudAuthSceneRequest) Validate() error {
	return dara.Validate(s)
}
