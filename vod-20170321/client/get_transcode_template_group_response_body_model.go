// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTranscodeTemplateGroupResponseBody
	GetRequestId() *string
	SetTranscodeTemplateGroup(v *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) *GetTranscodeTemplateGroupResponseBody
	GetTranscodeTemplateGroup() *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup
}

type GetTranscodeTemplateGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6730AC93-7B12-4B*****7F-49EE1FE8BC49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data of the transcoding template group.
	TranscodeTemplateGroup *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup `json:"TranscodeTemplateGroup,omitempty" xml:"TranscodeTemplateGroup,omitempty" type:"Struct"`
}

func (s GetTranscodeTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeTemplateGroupResponseBody) GetTranscodeTemplateGroup() *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	return s.TranscodeTemplateGroup
}

func (s *GetTranscodeTemplateGroupResponseBody) SetRequestId(v string) *GetTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroup(v *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) *GetTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroup = v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBody) Validate() error {
	if s.TranscodeTemplateGroup != nil {
		if err := s.TranscodeTemplateGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the template group was created. The time follows the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T10:20:51Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Specifies whether the template group is the default template group. Valid values:
	//
	// - **Default**: The default template group.
	//
	// - **NotDefault**: A non-default template group.
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Specifies whether the template group is locked. Valid values:
	//
	// - **Disabled**: Not locked.
	//
	// - **Enabled**: Locked.
	//
	// example:
	//
	// Enabled
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The time when the template group was modified. The time follows the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T11:20:51Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the template group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// a59b11f697c716*****6ae1502142d0
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The list of transcoding template configuration information.
	TranscodeTemplateList []*GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList `json:"TranscodeTemplateList,omitempty" xml:"TranscodeTemplateList,omitempty" type:"Repeated"`
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetAppId() *string {
	return s.AppId
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetLocked() *string {
	return s.Locked
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetName() *string {
	return s.Name
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GetTranscodeTemplateList() []*GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	return s.TranscodeTemplateList
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetAppId(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.AppId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetCreationTime(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetIsDefault(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.IsDefault = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetLocked(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.Locked = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetModifyTime(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.ModifyTime = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetName(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.Name = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetTranscodeTemplateGroupId(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetTranscodeTemplateList(v []*GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.TranscodeTemplateList = v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) Validate() error {
	if s.TranscodeTemplateList != nil {
		for _, item := range s.TranscodeTemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList struct {
	// The transcoding configuration parameters for the audio stream. The value is a JSON string.
	//
	// example:
	//
	// {"Codec":"AAC","Remove":"false","Bitrate":"44","Samplerate":"32000","Channels":"2","Profile":"aac_low"}
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The video clipping configuration. The value is a JSON string. For example, if you need to extract 5 seconds of content from a video to generate a new video, you can set this parameter.
	//
	// example:
	//
	// {"TimeSpan":{"Seek":"1","Duration":"5"}
	Clip *string `json:"Clip,omitempty" xml:"Clip,omitempty"`
	// The container format that encapsulates the audio and video streams. The value is a JSON string.
	//
	// example:
	//
	// "Format":"m3u8"
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The copyright watermark information.
	//
	// example:
	//
	// {
	//
	// 	"Content": "Test copyright watermark text"
	//
	// }
	CopyrightMark *string `json:"CopyrightMark,omitempty" xml:"CopyrightMark,omitempty"`
	// The definition tags for standard transcoding templates:
	//
	// - **LD*	- (Smooth)
	//
	// - **SD*	- (Standard definition)
	//
	// - **HD*	- (High definition)
	//
	// - **FHD*	- (Full high definition)
	//
	// - **OD*	- (Original definition, which means container format conversion)
	//
	// - **2K**
	//
	// - **4K**
	//
	// - **SQ*	- (Standard sound quality)
	//
	// - **HQ*	- (High sound quality)
	//
	// The definition tags for built-in Narrowband HD™ 1.0 transcoding templates:
	//
	// - **LD-NBV1*	- (Smooth)
	//
	// - **SD-NBV1*	- (Standard definition)
	//
	// - **HD-NBV1*	- (High definition)
	//
	// - **FHD-NBV1*	- (Full high definition)
	//
	// - **2K-NBV1**
	//
	// - **4K-NBV1**
	//
	// > - You cannot modify the definition tags of transcoding templates.
	//
	// > - The audio and video resolutions, bitrates, and other parameters of Narrowband HD™ 1.0 transcoding templates are built-in and cannot be modified.
	//
	// > - Narrowband HD™ 1.0 transcoding templates support only the FLV, M3U8 (HLS), and MP4 formats.
	//
	// example:
	//
	// SD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The encryption configuration information for transcoding.
	//
	// example:
	//
	// "EncryptType":"Private"
	EncryptSetting *string `json:"EncryptSetting,omitempty" xml:"EncryptSetting,omitempty"`
	// The segment settings for transcoding. This parameter is required for HLS. The value is a JSON string.
	//
	// example:
	//
	// "Segment": { "Duration":"6" }
	MuxConfig *string `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty"`
	// The packaging configuration. Only HLS adaptive bitrate streaming packaging and DASH packaging are supported. The value is a JSON string.
	//
	// example:
	//
	// "PackageType":"HLSPackage","PackageConfig":{   "BandWidth":"900000"  }
	PackageSetting *string `json:"PackageSetting,omitempty" xml:"PackageSetting,omitempty"`
	// The video rotation parameter. This parameter controls the rotation angle of the video image. For example, if you set this parameter to 180, the video image is upside down. Valid values: `[0,360]`.
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The subtitle configuration. The value is a JSON string.
	//
	// example:
	//
	// [{"SubtitleUrl":"http://outin-test.oss-cn-shanghai.aliyuncs.com/subtitles/c737fece-14f1-4364-b107-d5f7f8edde0e.ass","CharEncode":"utf-8"}]
	SubtitleList *string `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty"`
	// The name of the transcoding template.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The tracing watermark information.
	//
	// example:
	//
	// {
	//
	// 	"Enable": true
	//
	// }
	TraceMark *string `json:"TraceMark,omitempty" xml:"TraceMark,omitempty"`
	// The conditional transcoding parameters. If you need to make basic logical judgments based on the bitrate and resolution of the source video before outputting the transcoded video, you can set this parameter. The value is a JSON string.
	//
	// example:
	//
	// {"IsCheckReso":"true","IsCheckResoFail":"false","IsCheckVideoBitrate":"false","IsCheckVideoBitrateFail":"false","IsCheckAudioBitrate":"false","IsCheckAudioBitrateFail":"false"}
	TransConfig *string `json:"TransConfig,omitempty" xml:"TransConfig,omitempty"`
	// The custom output path for transcoding.
	//
	// example:
	//
	// {MediaId}/transcoce_1
	TranscodeFileRegular *string `json:"TranscodeFileRegular,omitempty" xml:"TranscodeFileRegular,omitempty"`
	// The ID of the transcoding template.
	//
	// example:
	//
	// 696d29a11erc057*****a3acc398d02f4
	TranscodeTemplateId *string `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
	// The templatetype. Valid values:
	//
	// - **Normal*	- (default): A standard transcoding template. You cannot configure the Settings for the PackageSetting parameter for this type of template.
	//
	// - **VideoPackage**: A video stream packaging template. This type of template transcodes the video stream first and then builds an adaptive bitrate stream. You must configure the Settings for the PackageSetting parameter.
	//
	// - **SubtitlePackage**: A subtitle packaging template. This type of template does not transcode the video stream. It only builds the corresponding subtitle information into the adaptive bitrate output file. You must configure the Settings for the PackageSetting parameter. This type of template cannot exist independently in a template group and must be configured together with a VideoPackage template. You can configure only one SubtitlePackage template in a template group.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The transcoding configuration parameters for the video stream. The value is a JSON string.
	//
	// example:
	//
	// {"Codec":"H.264","Bitrate":"900","Width":"960","Remove":"false","Fps":"30"}
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
	// The IDs of the associated image and text watermark templates.
	WatermarkIds []*string `json:"WatermarkIds,omitempty" xml:"WatermarkIds,omitempty" type:"Repeated"`
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetAudio() *string {
	return s.Audio
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetClip() *string {
	return s.Clip
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetContainer() *string {
	return s.Container
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetCopyrightMark() *string {
	return s.CopyrightMark
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetDefinition() *string {
	return s.Definition
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetEncryptSetting() *string {
	return s.EncryptSetting
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetMuxConfig() *string {
	return s.MuxConfig
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetPackageSetting() *string {
	return s.PackageSetting
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetRotate() *string {
	return s.Rotate
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetSubtitleList() *string {
	return s.SubtitleList
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetTraceMark() *string {
	return s.TraceMark
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetTransConfig() *string {
	return s.TransConfig
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetTranscodeFileRegular() *string {
	return s.TranscodeFileRegular
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetTranscodeTemplateId() *string {
	return s.TranscodeTemplateId
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetType() *string {
	return s.Type
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetVideo() *string {
	return s.Video
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GetWatermarkIds() []*string {
	return s.WatermarkIds
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetAudio(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Audio = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetClip(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Clip = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetContainer(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Container = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetCopyrightMark(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.CopyrightMark = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetDefinition(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Definition = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetEncryptSetting(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.EncryptSetting = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetMuxConfig(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.MuxConfig = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetPackageSetting(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.PackageSetting = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetRotate(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Rotate = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetSubtitleList(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.SubtitleList = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTemplateName(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TemplateName = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTraceMark(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TraceMark = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTransConfig(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TransConfig = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTranscodeFileRegular(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TranscodeFileRegular = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTranscodeTemplateId(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TranscodeTemplateId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetType(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Type = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetVideo(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Video = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetWatermarkIds(v []*string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.WatermarkIds = v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) Validate() error {
	return dara.Validate(s)
}
