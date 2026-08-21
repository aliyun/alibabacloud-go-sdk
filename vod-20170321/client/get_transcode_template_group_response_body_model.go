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
	// The request ID.
	//
	// example:
	//
	// 6730AC93-7B12-4B*****7F-49EE1FE8BC49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The transcoding template group data.
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
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the template group was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-12-12T10:20:51Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the template group is the default one. Valid values:
	//
	// - **Default**: The template group is the default one.
	//
	// - **NotDefault**: The template group is not the default one.
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Indicates whether the template group is locked. Valid values:
	//
	// - **Disabled**: Not locked.
	//
	// - **Enabled**: Locked.
	//
	// example:
	//
	// Enabled
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The time when the template group was last modified. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
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
	// The transcoding template group ID.
	//
	// example:
	//
	// a59b11f697c716*****6ae1502142d0
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The list of transcoding template configurations.
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
	// The audio stream transcoding configuration parameters (JSON string).
	//
	// example:
	//
	// {"Codec":"AAC","Remove":"false","Bitrate":"44","Samplerate":"32000","Channels":"2","Profile":"aac_low"}
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The video clipping configuration (JSON string). For example, set this parameter if you want to extract 5 seconds of content from a video to generate a new video.
	//
	// example:
	//
	// {"TimeSpan":{"Seek":"1","Duration":"5"}
	Clip *string `json:"Clip,omitempty" xml:"Clip,omitempty"`
	// The container format for encapsulating audio and video streams (JSON string).
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
	// "Content": "Test copyright watermark text"
	//
	// }
	CopyrightMark *string `json:"CopyrightMark,omitempty" xml:"CopyrightMark,omitempty"`
	// The definition mark for normal transcoding templates:
	//
	// - **LD*	- (low definition)
	//
	// - **SD*	- (standard definition)
	//
	// - **HD*	- (high definition)
	//
	// - **FHD*	- (full high definition)
	//
	// - **OD*	- (original definition, container format conversion)
	//
	// - **2K**
	//
	// - **4K**
	//
	// - **SQ*	- (standard audio quality)
	//
	// - **HQ*	- (high audio quality)
	//
	// The definition mark for Narrowband HD 1.0 built-in transcoding templates:
	//
	// - **LD-NBV1*	- (low definition)
	//
	// - **SD-NBV1*	- (standard definition)
	//
	// - **HD-NBV1*	- (high definition)
	//
	// - **FHD-NBV1*	- (full high definition)
	//
	// - **2K-NBV1**
	//
	// - **4K-NBV1**
	//
	// > - The definition mark of transcoding templates cannot be modified.
	//
	// > - The audio and video resolution, bitrate, and other parameters of Narrowband HD 1.0 transcoding templates are built into the system and cannot be modified.
	//
	// > - Narrowband HD 1.0 transcoding templates can only be created in FLV, M3U8 (HLS), or MP4 format.
	//
	// example:
	//
	// SD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The transcoding encryption configuration.
	//
	// example:
	//
	// "EncryptType":"Private"
	EncryptSetting *string `json:"EncryptSetting,omitempty" xml:"EncryptSetting,omitempty"`
	// The segment setting parameters for transcoding. Required for HLS (JSON string).
	//
	// example:
	//
	// "Segment": { "Duration":"6" }
	MuxConfig *string `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty"`
	// The packaging configuration. Only HLS adaptive bitrate streaming packaging and DASH packaging are supported (JSON string).
	//
	// example:
	//
	// "PackageType":"HLSPackage","PackageConfig":{   "BandWidth":"900000"  }
	PackageSetting *string `json:"PackageSetting,omitempty" xml:"PackageSetting,omitempty"`
	// The video rotation parameter. Controls the rotation angle of the video. For example, if set to 180, the video is flipped upside down. Value range: `[0,360]`.
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The subtitle configuration (JSON string).
	//
	// example:
	//
	// [{"SubtitleUrl":"http://outin-test.oss-cn-shanghai.aliyuncs.com/subtitles/c737fece-14f1-4364-b107-d5f7f8edde0e.ass","CharEncode":"utf-8"}]
	SubtitleList *string `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty"`
	// The transcoding template name.
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
	// The conditional transcoding parameters. Set this parameter if you want to perform basic logic checks based on the bitrate or resolution of the source file before outputting the transcoded video (JSON string).
	//
	// example:
	//
	// {"IsCheckReso":"true","IsCheckResoFail":"false","IsCheckVideoBitrate":"false","IsCheckVideoBitrateFail":"false","IsCheckAudioBitrate":"false","IsCheckAudioBitrateFail":"false"}
	TransConfig *string `json:"TransConfig,omitempty" xml:"TransConfig,omitempty"`
	// The custom transcoding output path.
	//
	// example:
	//
	// {MediaId}/transcoce_1
	TranscodeFileRegular *string `json:"TranscodeFileRegular,omitempty" xml:"TranscodeFileRegular,omitempty"`
	// The transcoding template ID.
	//
	// example:
	//
	// 696d29a11erc057*****a3acc398d02f4
	TranscodeTemplateId *string `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
	// The templatetype. Valid values:
	//
	// - **Normal*	- (default): a normal transcoding template. The PackageSetting parameter cannot be configured in Settings for this type of template.
	//
	// - **VideoPackage**: a video stream packaging template. This type of template first transcodes and then builds adaptive bitrate streaming. The PackageSetting parameter must be configured in Settings for this type of template.
	//
	// - **SubtitlePackage**: a subtitle packaging template. This type of template does not transcode but only builds the corresponding subtitle information into the adaptive bitrate streaming output file. The PackageSetting parameter must be configured in Settings for this type of template. This type of template cannot exist alone in a template group and must be configured together with a VideoPackage type template. Only one SubtitlePackage template can be configured in a template group.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The video stream transcoding configuration parameters (JSON string).
	//
	// example:
	//
	// {"Codec":"H.264","Bitrate":"900","Width":"960","Remove":"false","Fps":"30"}
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
	// The IDs of associated image and text watermark templates.
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
