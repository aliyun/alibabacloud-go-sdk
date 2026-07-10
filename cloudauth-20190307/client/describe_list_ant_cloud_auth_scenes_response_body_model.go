// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListAntCloudAuthScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeListAntCloudAuthScenesResponseBody
	GetRequestId() *string
	SetScenes(v []*DescribeListAntCloudAuthScenesResponseBodyScenes) *DescribeListAntCloudAuthScenesResponseBody
	GetScenes() []*DescribeListAntCloudAuthScenesResponseBodyScenes
}

type DescribeListAntCloudAuthScenesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CC1AB3F5-22A2-589F-ABDD-B766694AA671
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scenes.
	Scenes []*DescribeListAntCloudAuthScenesResponseBodyScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
}

func (s DescribeListAntCloudAuthScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeListAntCloudAuthScenesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListAntCloudAuthScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeListAntCloudAuthScenesResponseBody) GetScenes() []*DescribeListAntCloudAuthScenesResponseBodyScenes {
	return s.Scenes
}

func (s *DescribeListAntCloudAuthScenesResponseBody) SetRequestId(v string) *DescribeListAntCloudAuthScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBody) SetScenes(v []*DescribeListAntCloudAuthScenesResponseBodyScenes) *DescribeListAntCloudAuthScenesResponseBody {
	s.Scenes = v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBody) Validate() error {
	if s.Scenes != nil {
		for _, item := range s.Scenes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeListAntCloudAuthScenesResponseBodyScenes struct {
	// The application ID.
	//
	// example:
	//
	// 2a3a13b6-ee85-457e-bd15-b48115cb396e
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// The creation time.
	//
	// example:
	//
	// 1260051251634779
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 189884094677xxxx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
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
	// The bound domain name.
	//
	// example:
	//
	// www.ddos.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The mini program name.
	//
	// example:
	//
	// 测试APP
	MiniProgramName *string `json:"MiniProgramName,omitempty" xml:"MiniProgramName,omitempty"`
	// The modifier.
	//
	// example:
	//
	// 189884094677xxxx
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The mini program platform. Valid values:
	//
	// - **WECHAT**: WeChat
	//
	// - **ALIPAY**: Alipay
	//
	// - **TIKTOK**: TikTok.
	//
	// example:
	//
	// WECHAT
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The number of evidence face photos (1 to 5).
	//
	// example:
	//
	// 1
	ReturnPicCount *int64 `json:"ReturnPicCount,omitempty" xml:"ReturnPicCount,omitempty"`
	// The duration of the evidence video, in seconds.
	//
	// example:
	//
	// 1
	ReturnVideoLength *int64 `json:"ReturnVideoLength,omitempty" xml:"ReturnVideoLength,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 100001xxxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// 测试场景
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// Indicates whether the scene is enabled. The value 1 indicates enabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to deliver files generated during authentication to the customer\\"s OSS. Valid values:
	//
	// - **Y**: Enabled.
	//
	// - **N**: Disabled.
	//
	// example:
	//
	// Y
	StoreImage *string `json:"StoreImage,omitempty" xml:"StoreImage,omitempty"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 1260051251634779
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeListAntCloudAuthScenesResponseBodyScenes) String() string {
	return dara.Prettify(s)
}

func (s DescribeListAntCloudAuthScenesResponseBodyScenes) GoString() string {
	return s.String()
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetBindMiniProgram() *string {
	return s.BindMiniProgram
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetCreator() *string {
	return s.Creator
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetDeviceRiskPlus() *string {
	return s.DeviceRiskPlus
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetDomain() *string {
	return s.Domain
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetMiniProgramName() *string {
	return s.MiniProgramName
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetReturnPicCount() *int64 {
	return s.ReturnPicCount
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetReturnVideoLength() *int64 {
	return s.ReturnVideoLength
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetStoreImage() *string {
	return s.StoreImage
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetAppId(v int64) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.AppId = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetBindMiniProgram(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.BindMiniProgram = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetCreateTime(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.CreateTime = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetCreator(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.Creator = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetDeviceRiskPlus(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.DeviceRiskPlus = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetDomain(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.Domain = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetMiniProgramName(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.MiniProgramName = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetModifier(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.Modifier = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetPlatform(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.Platform = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetReturnPicCount(v int64) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.ReturnPicCount = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetReturnVideoLength(v int64) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.ReturnVideoLength = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetSceneId(v int64) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.SceneId = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetSceneName(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.SceneName = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetStatus(v int32) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.Status = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetStoreImage(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.StoreImage = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) SetUpdateTime(v string) *DescribeListAntCloudAuthScenesResponseBodyScenes {
	s.UpdateTime = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesResponseBodyScenes) Validate() error {
	return dara.Validate(s)
}
