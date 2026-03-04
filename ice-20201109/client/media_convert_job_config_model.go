// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertJobConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInputs(v []*MediaConvertJobConfigInputs) *MediaConvertJobConfig
	GetInputs() []*MediaConvertJobConfigInputs
	SetJobName(v string) *MediaConvertJobConfig
	GetJobName() *string
	SetOutputGroups(v []*MediaConvertJobConfigOutputGroups) *MediaConvertJobConfig
	GetOutputGroups() []*MediaConvertJobConfigOutputGroups
}

type MediaConvertJobConfig struct {
	// Inputs.
	Inputs []*MediaConvertJobConfigInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// The task name.
	//
	// 	- Maximum length: 64 bytes.
	//
	// example:
	//
	// Name
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// An array of output group configurations.
	OutputGroups []*MediaConvertJobConfigOutputGroups `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty" type:"Repeated"`
}

func (s MediaConvertJobConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfig) GetInputs() []*MediaConvertJobConfigInputs {
	return s.Inputs
}

func (s *MediaConvertJobConfig) GetJobName() *string {
	return s.JobName
}

func (s *MediaConvertJobConfig) GetOutputGroups() []*MediaConvertJobConfigOutputGroups {
	return s.OutputGroups
}

func (s *MediaConvertJobConfig) SetInputs(v []*MediaConvertJobConfigInputs) *MediaConvertJobConfig {
	s.Inputs = v
	return s
}

func (s *MediaConvertJobConfig) SetJobName(v string) *MediaConvertJobConfig {
	s.JobName = &v
	return s
}

func (s *MediaConvertJobConfig) SetOutputGroups(v []*MediaConvertJobConfigOutputGroups) *MediaConvertJobConfig {
	s.OutputGroups = v
	return s
}

func (s *MediaConvertJobConfig) Validate() error {
	if s.Inputs != nil {
		for _, item := range s.Inputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputGroups != nil {
		for _, item := range s.OutputGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MediaConvertJobConfigInputs struct {
	// The input file.
	InputFile *MediaConvertJobConfigInputsInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The name of the input file. InputRef can reference it in the output configuration.
	//
	// example:
	//
	// InputVideo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MediaConvertJobConfigInputs) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigInputs) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigInputs) GetInputFile() *MediaConvertJobConfigInputsInputFile {
	return s.InputFile
}

func (s *MediaConvertJobConfigInputs) GetName() *string {
	return s.Name
}

func (s *MediaConvertJobConfigInputs) SetInputFile(v *MediaConvertJobConfigInputsInputFile) *MediaConvertJobConfigInputs {
	s.InputFile = v
	return s
}

func (s *MediaConvertJobConfigInputs) SetName(v string) *MediaConvertJobConfigInputs {
	s.Name = &v
	return s
}

func (s *MediaConvertJobConfigInputs) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertJobConfigInputsInputFile struct {
	// The URL or ID of the media file.
	//
	// 	- If Type is OSS, this is the file URL (OSS or HTTP). Do not use a signed URL, as it may cause authentication failure.
	//
	// 	- If Type is Media, this is the media asset ID. The source stream is used by default.
	//
	// example:
	//
	// http://bucket.loction.aliyuncs.com/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media file. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaConvertJobConfigInputsInputFile) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigInputsInputFile) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigInputsInputFile) GetMedia() *string {
	return s.Media
}

func (s *MediaConvertJobConfigInputsInputFile) GetType() *string {
	return s.Type
}

func (s *MediaConvertJobConfigInputsInputFile) SetMedia(v string) *MediaConvertJobConfigInputsInputFile {
	s.Media = &v
	return s
}

func (s *MediaConvertJobConfigInputsInputFile) SetType(v string) *MediaConvertJobConfigInputsInputFile {
	s.Type = &v
	return s
}

func (s *MediaConvertJobConfigInputsInputFile) Validate() error {
	return dara.Validate(s)
}

type MediaConvertJobConfigOutputGroups struct {
	// The configuration for this output group.
	GroupConfig *MediaConvertJobConfigOutputGroupsGroupConfig `json:"GroupConfig,omitempty" xml:"GroupConfig,omitempty" type:"Struct"`
	// The name of the output group.
	//
	// example:
	//
	// hls-group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of individual output stream configurations. Each object in this array defines a specific rendition.
	Outputs []*MediaConvertJobConfigOutputGroupsOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s MediaConvertJobConfigOutputGroups) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroups) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroups) GetGroupConfig() *MediaConvertJobConfigOutputGroupsGroupConfig {
	return s.GroupConfig
}

func (s *MediaConvertJobConfigOutputGroups) GetName() *string {
	return s.Name
}

func (s *MediaConvertJobConfigOutputGroups) GetOutputs() []*MediaConvertJobConfigOutputGroupsOutputs {
	return s.Outputs
}

func (s *MediaConvertJobConfigOutputGroups) SetGroupConfig(v *MediaConvertJobConfigOutputGroupsGroupConfig) *MediaConvertJobConfigOutputGroups {
	s.GroupConfig = v
	return s
}

func (s *MediaConvertJobConfigOutputGroups) SetName(v string) *MediaConvertJobConfigOutputGroups {
	s.Name = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroups) SetOutputs(v []*MediaConvertJobConfigOutputGroupsOutputs) *MediaConvertJobConfigOutputGroups {
	s.Outputs = v
	return s
}

func (s *MediaConvertJobConfigOutputGroups) Validate() error {
	if s.GroupConfig != nil {
		if err := s.GroupConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MediaConvertJobConfigOutputGroupsGroupConfig struct {
	// Configures an extension to an existing manifest file. This allows you to reference an existing manifest and combine it with the results of the current output group to generate a new manifest.
	ManifestExtend *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend `json:"ManifestExtend,omitempty" xml:"ManifestExtend,omitempty" type:"Struct"`
	// The name of the manifest file. This parameter is only applicable when Type is set to Hls or Dash.
	//
	// example:
	//
	// manifest
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The base output directory. All files generated by this OutputGroup are placed in this directory.
	OutputFileBase *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase `json:"OutputFileBase,omitempty" xml:"OutputFileBase,omitempty" type:"Struct"`
	// The type of the output group. Valid values:
	//
	// 	- File: An individual file.
	//
	// 	- Hls: HLS files for adaptive bitrate streaming.
	//
	// example:
	//
	// Hls
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaConvertJobConfigOutputGroupsGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroupsGroupConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) GetManifestExtend() *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend {
	return s.ManifestExtend
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) GetManifestName() *string {
	return s.ManifestName
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) GetOutputFileBase() *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase {
	return s.OutputFileBase
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) GetType() *string {
	return s.Type
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) SetManifestExtend(v *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) *MediaConvertJobConfigOutputGroupsGroupConfig {
	s.ManifestExtend = v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) SetManifestName(v string) *MediaConvertJobConfigOutputGroupsGroupConfig {
	s.ManifestName = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) SetOutputFileBase(v *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) *MediaConvertJobConfigOutputGroupsGroupConfig {
	s.OutputFileBase = v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) SetType(v string) *MediaConvertJobConfigOutputGroupsGroupConfig {
	s.Type = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfig) Validate() error {
	if s.ManifestExtend != nil {
		if err := s.ManifestExtend.Validate(); err != nil {
			return err
		}
	}
	if s.OutputFileBase != nil {
		if err := s.OutputFileBase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend struct {
	// Specifies streams to exclude from the referenced manifest. Multiple conditions within a single exclusion object are combined using AND logic. Multiple exclusion objects in the array are combined using OR logic.
	Excludes []*MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Repeated"`
	// References the Name of an input that contains the manifest to be extended.
	//
	// example:
	//
	// Input-Manifest
	InputRef *string `json:"InputRef,omitempty" xml:"InputRef,omitempty"`
}

func (s MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) GetExcludes() []*MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes {
	return s.Excludes
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) GetInputRef() *string {
	return s.InputRef
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) SetExcludes(v []*MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend {
	s.Excludes = v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) SetInputRef(v string) *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend {
	s.InputRef = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtend) Validate() error {
	if s.Excludes != nil {
		for _, item := range s.Excludes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes struct {
	// Excludes streams based on their Language field. It must conform to RFC 5646.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Excludes streams based on their Name field.
	//
	// example:
	//
	// audio-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Excludes streams based on their Type field.
	//
	// Valid values:
	//
	// 	- Audio
	//
	// 	- Subtitle
	//
	// example:
	//
	// Audio
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) GetLanguage() *string {
	return s.Language
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) GetName() *string {
	return s.Name
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) GetType() *string {
	return s.Type
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) SetLanguage(v string) *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes {
	s.Language = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) SetName(v string) *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes {
	s.Name = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) SetType(v string) *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes {
	s.Type = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigManifestExtendExcludes) Validate() error {
	return dara.Validate(s)
}

type MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase struct {
	// The media asset location.
	//
	// 	- If Type is OSS, this is the destination URL (OSS or HTTP).
	//
	// 	- If Type is Media, this is the destination media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/dir
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) GetMedia() *string {
	return s.Media
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) GetType() *string {
	return s.Type
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) SetMedia(v string) *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase {
	s.Media = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) SetType(v string) *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase {
	s.Type = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsGroupConfigOutputFileBase) Validate() error {
	return dara.Validate(s)
}

type MediaConvertJobConfigOutputGroupsOutputs struct {
	// Additional feature parameters. See [MediaConvertJobFeature](https://help.aliyun.com/document_detail/2979993.html) for details.
	//
	// example:
	//
	// {}
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// HLS-specific settings for this stream\\"s behavior in the manifest. Effective only when the group Type is Hls.
	HlsGroupConfig *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig `json:"HlsGroupConfig,omitempty" xml:"HlsGroupConfig,omitempty" type:"Struct"`
	// A name to label this output. This is for identification purposes only and does not affect the filename.
	//
	// example:
	//
	// group-output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filename for this output. It is used in conjunction with OutputFileBase from the GroupConfig.
	//
	// example:
	//
	// 720p.mp4
	OutputFileName *string `json:"OutputFileName,omitempty" xml:"OutputFileName,omitempty"`
	// A JSON string of parameters to override the settings in the specified template.
	//
	// example:
	//
	// {}
	OverrideParams *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	// The priority of the task, from 1 to 10. A higher value indicates a higher priority. Default: 6.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the transcoding template.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MediaConvertJobConfigOutputGroupsOutputs) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroupsOutputs) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetFeatures() *string {
	return s.Features
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetHlsGroupConfig() *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	return s.HlsGroupConfig
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetName() *string {
	return s.Name
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetOutputFileName() *string {
	return s.OutputFileName
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetOverrideParams() *string {
	return s.OverrideParams
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetPriority() *int32 {
	return s.Priority
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetFeatures(v string) *MediaConvertJobConfigOutputGroupsOutputs {
	s.Features = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetHlsGroupConfig(v *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) *MediaConvertJobConfigOutputGroupsOutputs {
	s.HlsGroupConfig = v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetName(v string) *MediaConvertJobConfigOutputGroupsOutputs {
	s.Name = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetOutputFileName(v string) *MediaConvertJobConfigOutputGroupsOutputs {
	s.OutputFileName = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetOverrideParams(v string) *MediaConvertJobConfigOutputGroupsOutputs {
	s.OverrideParams = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetPriority(v int32) *MediaConvertJobConfigOutputGroupsOutputs {
	s.Priority = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) SetTemplateId(v string) *MediaConvertJobConfigOutputGroupsOutputs {
	s.TemplateId = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputs) Validate() error {
	if s.HlsGroupConfig != nil {
		if err := s.HlsGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig struct {
	// The audio group this video stream references. Effective when Type is video.
	//
	// Default value: audio.
	//
	// example:
	//
	// audio
	AudioGroup *string `json:"AudioGroup,omitempty" xml:"AudioGroup,omitempty"`
	// Whether this stream should be automatically selected by the player. Effective when Type is audio or subtitle.
	//
	// example:
	//
	// true
	AutoSelect *string `json:"AutoSelect,omitempty" xml:"AutoSelect,omitempty"`
	// Whether this stream must be played. Effective when Type is audio or subtitle.
	//
	// example:
	//
	// true
	Forced *string `json:"Forced,omitempty" xml:"Forced,omitempty"`
	// The GROUP-ID attribute for this stream in the HLS manifest. Effective when Type is audio or subtitle.
	//
	// Default value: Same as the Type value.
	//
	// example:
	//
	// audio
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// Whether this is the default stream within its group. Effective when Type is audio or subtitle. Only one stream per group can be the default.
	//
	// example:
	//
	// false
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The LANGUAGE attribute for this stream in the HLS manifest (must conform to RFC 5646). Effective when Type is audio or subtitle.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The NAME attribute for this stream in the HLS manifest. **Required*	- when Type is audio or subtitle.
	//
	// example:
	//
	// audio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subtitle group this video or hybrid stream references. Effective when Type is video or hybrid.
	//
	// Default value: subtitle.
	//
	// example:
	//
	// subtitle
	SubtitleGroup *string `json:"SubtitleGroup,omitempty" xml:"SubtitleGroup,omitempty"`
	// Specifies the stream type.
	//
	// Valid values:
	//
	// 	- video: Video stream.
	//
	// 	- audio: Audio stream.
	//
	// 	- subtitle: Subtitle stream.
	//
	// 	- hybrid: A stream containing both audio and video.
	//
	// Default value: hybrid.
	//
	// example:
	//
	// hybrid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetAudioGroup() *string {
	return s.AudioGroup
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetAutoSelect() *string {
	return s.AutoSelect
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetForced() *string {
	return s.Forced
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetGroup() *string {
	return s.Group
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetIsDefault() *string {
	return s.IsDefault
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetLanguage() *string {
	return s.Language
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetName() *string {
	return s.Name
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetSubtitleGroup() *string {
	return s.SubtitleGroup
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) GetType() *string {
	return s.Type
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetAudioGroup(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.AudioGroup = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetAutoSelect(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.AutoSelect = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetForced(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.Forced = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetGroup(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.Group = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetIsDefault(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.IsDefault = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetLanguage(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.Language = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetName(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.Name = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetSubtitleGroup(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.SubtitleGroup = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) SetType(v string) *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig {
	s.Type = &v
	return s
}

func (s *MediaConvertJobConfigOutputGroupsOutputsHlsGroupConfig) Validate() error {
	return dara.Validate(s)
}
