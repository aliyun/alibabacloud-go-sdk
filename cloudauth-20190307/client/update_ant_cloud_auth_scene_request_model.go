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
	SetDegradeAppScheme(v string) *UpdateAntCloudAuthSceneRequest
	GetDegradeAppScheme() *string
	SetDegradeSubCodes(v string) *UpdateAntCloudAuthSceneRequest
	GetDegradeSubCodes() *string
	SetDegradeType(v string) *UpdateAntCloudAuthSceneRequest
	GetDegradeType() *string
	SetDeviceRiskPlus(v string) *UpdateAntCloudAuthSceneRequest
	GetDeviceRiskPlus() *string
	SetMiniProgramName(v string) *UpdateAntCloudAuthSceneRequest
	GetMiniProgramName() *string
	SetPlatform(v string) *UpdateAntCloudAuthSceneRequest
	GetPlatform() *string
	SetReturnPicCount(v int64) *UpdateAntCloudAuthSceneRequest
	GetReturnPicCount() *int64
	SetReturnVideoLength(v int64) *UpdateAntCloudAuthSceneRequest
	GetReturnVideoLength() *int64
	SetSceneId(v int64) *UpdateAntCloudAuthSceneRequest
	GetSceneId() *int64
	SetSceneName(v string) *UpdateAntCloudAuthSceneRequest
	GetSceneName() *string
	SetStatus(v int32) *UpdateAntCloudAuthSceneRequest
	GetStatus() *int32
	SetStoreImage(v string) *UpdateAntCloudAuthSceneRequest
	GetStoreImage() *string
	SetUseDegrade(v string) *UpdateAntCloudAuthSceneRequest
	GetUseDegrade() *string
}

type UpdateAntCloudAuthSceneRequest struct {
	// Specifies whether to bind a mini program. Valid values:
	//
	// - **Y**: enabled.
	//
	// - **N (default)**: disabled.
	//
	// 	Notice: If you enable mini program binding, make sure that you specify all parameters related to the mini program binding.
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
	// The iOS app scheme for degradation.
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
	// Specifies whether to enable degraded authentication.
	//
	// example:
	//
	// ALIPAY
	DegradeType *string `json:"DegradeType,omitempty" xml:"DegradeType,omitempty"`
	// Specifies whether to enable enhanced device risk detection. Valid values:
	//
	// - **Y**: enabled.
	//
	// - **N**: disabled.
	//
	// example:
	//
	// Y
	DeviceRiskPlus *string `json:"DeviceRiskPlus,omitempty" xml:"DeviceRiskPlus,omitempty"`
	// The name of the mini program.
	//
	// example:
	//
	// TestApp
	MiniProgramName *string `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	// The mini program platform. Valid values:
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
	// The number of returned photos (1 to 5). This parameter takes effect after StoreImage is enabled for authentication file retention.
	//
	// example:
	//
	// 1
	ReturnPicCount *int64 `json:"ReturnPicCount,omitempty" xml:"ReturnPicCount,omitempty"`
	// The duration of the returned video (1 to 2 seconds). This parameter takes effect after StoreImage is enabled.
	//
	// example:
	//
	// 2
	ReturnVideoLength *int64 `json:"ReturnVideoLength,omitempty" xml:"ReturnVideoLength,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000013372
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// This parameter has no effect. You do not need to specify this parameter.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to deliver authentication files generated during the authentication process to your OSS bucket. Valid values:
	//
	// - **Y**: enabled.
	//
	// - **N (default)**: disabled.
	//
	// example:
	//
	// Y
	StoreImage *string `json:"StoreImage,omitempty" xml:"StoreImage,omitempty"`
	// Specifies whether to enable degraded authentication.
	//
	// example:
	//
	// Y
	UseDegrade *string `json:"UseDegrade,omitempty" xml:"UseDegrade,omitempty"`
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

func (s *UpdateAntCloudAuthSceneRequest) GetDegradeAppScheme() *string {
	return s.DegradeAppScheme
}

func (s *UpdateAntCloudAuthSceneRequest) GetDegradeSubCodes() *string {
	return s.DegradeSubCodes
}

func (s *UpdateAntCloudAuthSceneRequest) GetDegradeType() *string {
	return s.DegradeType
}

func (s *UpdateAntCloudAuthSceneRequest) GetDeviceRiskPlus() *string {
	return s.DeviceRiskPlus
}

func (s *UpdateAntCloudAuthSceneRequest) GetMiniProgramName() *string {
	return s.MiniProgramName
}

func (s *UpdateAntCloudAuthSceneRequest) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateAntCloudAuthSceneRequest) GetReturnPicCount() *int64 {
	return s.ReturnPicCount
}

func (s *UpdateAntCloudAuthSceneRequest) GetReturnVideoLength() *int64 {
	return s.ReturnVideoLength
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

func (s *UpdateAntCloudAuthSceneRequest) GetUseDegrade() *string {
	return s.UseDegrade
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

func (s *UpdateAntCloudAuthSceneRequest) SetDegradeAppScheme(v string) *UpdateAntCloudAuthSceneRequest {
	s.DegradeAppScheme = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetDegradeSubCodes(v string) *UpdateAntCloudAuthSceneRequest {
	s.DegradeSubCodes = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetDegradeType(v string) *UpdateAntCloudAuthSceneRequest {
	s.DegradeType = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetDeviceRiskPlus(v string) *UpdateAntCloudAuthSceneRequest {
	s.DeviceRiskPlus = &v
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

func (s *UpdateAntCloudAuthSceneRequest) SetReturnPicCount(v int64) *UpdateAntCloudAuthSceneRequest {
	s.ReturnPicCount = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) SetReturnVideoLength(v int64) *UpdateAntCloudAuthSceneRequest {
	s.ReturnVideoLength = &v
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

func (s *UpdateAntCloudAuthSceneRequest) SetUseDegrade(v string) *UpdateAntCloudAuthSceneRequest {
	s.UseDegrade = &v
	return s
}

func (s *UpdateAntCloudAuthSceneRequest) Validate() error {
	return dara.Validate(s)
}
