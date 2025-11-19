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
	// The information about the transcoding template group.
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
	// The time when the transcoding template group was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T10:20:51Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the template group is the default one. Valid values:
	//
	// 	- **Default**
	//
	// 	- **NotDefault**
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Indicates whether the transcoding template group is locked. Valid values:
	//
	// 	- **Disabled**: The template group is not locked.
	//
	// 	- **Enabled**: The template group is locked.
	//
	// example:
	//
	// Enabled
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The time when the transcoding template group was last modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T11:20:51Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the transcoding template group.
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
	// The information about the transcoding templates.
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
	// The transcoding configurations of the audio stream. The value is a JSON string.
	//
	// example:
	//
	// {\\"Codec\\":\\"AAC\\",\\"Remove\\":\\"false\\",\\"Bitrate\\":\\"44\\",\\"Samplerate\\":\\"32000\\",\\"Channels\\":\\"2\\",\\"Profile\\":\\"aac_low\\"}
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The clipping configurations of the video. The value is a JSON string. For example, this parameter is returned if you extract 5 seconds of content from a video to generate a new video.
	//
	// example:
	//
	// {\\"TimeSpan\\":{\\"Seek\\":\\"1\\",\\"Duration\\":\\"5\\"}
	Clip *string `json:"Clip,omitempty" xml:"Clip,omitempty"`
	// The format of the container used to encapsulate audio and video streams. The value is a JSON string.
	//
	// example:
	//
	// "Format":"m3u8"
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The content of the copyright watermark.
	//
	// example:
	//
	// {
	//
	// 	"Content": "test"
	//
	// }
	CopyrightMark *string `json:"CopyrightMark,omitempty" xml:"CopyrightMark,omitempty"`
	// Valid values for the definition of a common transcoding template:
	//
	// 	- **LD**: low definition.
	//
	// 	- **SD**: standard definition.
	//
	// 	- **HD**: high definition.
	//
	// 	- **FHD**: ultra high definition.
	//
	// 	- **OD**: original quality.
	//
	// 	- **2K**
	//
	// 	- **4K**
	//
	// 	- **SQ**: standard sound quality.
	//
	// 	- **HQ**: high sound quality.
	//
	// Valid values for the definition of a Narrowband HD™ 1.0 transcoding template:
	//
	// 	- **LD-NBV1**: low definition.
	//
	// 	- **SD-NBV1**: standard definition.
	//
	// 	- **HD-NBV1**: high definition.
	//
	// 	- **FHD-NBV1**: ultra high definition.
	//
	// 	- **2K-NBV1**
	//
	// 	- **4K-NBV1**
	//
	// > 	- You cannot change the definition of a transcoding template.
	//
	// >	- You cannot modify the system parameters, such as the video resolution, audio resolution, and bitrate, of Narrowband HD™ 1.0 transcoding templates.
	//
	// >	- You can create only Narrowband HD™ 1.0 transcoding templates that support the FLV, M3U8 (HLS), and MP4 output formats.
	//
	// example:
	//
	// SD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The encryption configuration for transcoding.
	//
	// example:
	//
	// "EncryptType":"Private"
	EncryptSetting *string `json:"EncryptSetting,omitempty" xml:"EncryptSetting,omitempty"`
	// The transcoding segment configurations. This parameter must be returned if HTTP-Live-Streaming (HLS) encryption is used. The value is a JSON string.
	//
	// example:
	//
	// "Segment": { "Duration":"6" }
	MuxConfig *string `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty"`
	// The packaging configuration. Only HLS packaging and DASH packaging are supported. The value is a JSON string.
	//
	// example:
	//
	// "PackageType":"HLSPackage","PackageConfig":{   "BandWidth":"900000"  }
	PackageSetting *string `json:"PackageSetting,omitempty" xml:"PackageSetting,omitempty"`
	// The video rotation identifier. It is used to control the image rotation angle. For example, if you set this parameter to 180, the video image is turned upside down. Valid values: `[0,360]`.
	//
	// example:
	//
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The subtitle configurations. The value is a JSON string.
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
	// The content of the tracing watermark.
	//
	// example:
	//
	// {
	//
	// 	"Enable": true
	//
	// }
	TraceMark *string `json:"TraceMark,omitempty" xml:"TraceMark,omitempty"`
	// The conditional transcoding configurations. This parameter can be used if you want to determine the basic logic based on the bitrate and resolution of the source file before the video is transcoded. The value is a JSON-formatted string.
	//
	// example:
	//
	// {"IsCheckReso":"true","IsCheckResoFail":"false","IsCheckVideoBitrate":"false","IsCheckVideoBitrateFail":"false","IsCheckAudioBitrate":"false","IsCheckAudioBitrateFail":"false"}
	TransConfig *string `json:"TransConfig,omitempty" xml:"TransConfig,omitempty"`
	// The custom path used to store the output files.
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
	// The type of the transcoding template. Valid values:
	//
	// 	- **Normal*	- (default): a common transcoding template. The PackageSetting parameter cannot be set for this type of template.
	//
	// 	- **VideoPackage**: a video stream package template. If this type of template is used, ApsaraVideo VOD transcodes a video into video streams in different bitrates and packages these video streams with a file. The PackageSetting parameter must be set for this type of template.
	//
	// 	- **SubtitlePackage**: a subtitle package template. If this type of template is used, ApsaraVideo VOD adds the subtitle information to the output file generated by packaging the multi-bitrate video streams of the corresponding video without transcoding. You must set the PackageSetting parameter for a subtitle package template and associate the subtitle package template with a video stream package template. A template group can contain only one subtitle package template.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The transcoding configurations of the video stream. The value is a JSON string.
	//
	// example:
	//
	// {"Codec":"H.264","Bitrate":"900","Width":"960","Remove":"false","Fps":"30"}
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
	// The IDs of the associated watermarks.
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
