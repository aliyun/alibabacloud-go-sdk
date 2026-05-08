// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAvatarProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SaveAvatarProjectRequest
	GetAgentId() *string
	SetBitRate(v string) *SaveAvatarProjectRequest
	GetBitRate() *string
	SetFrameRate(v string) *SaveAvatarProjectRequest
	GetFrameRate() *string
	SetFrames(v []*SaveAvatarProjectRequestFrames) *SaveAvatarProjectRequest
	GetFrames() []*SaveAvatarProjectRequestFrames
	SetOperateType(v string) *SaveAvatarProjectRequest
	GetOperateType() *string
	SetProjectId(v string) *SaveAvatarProjectRequest
	GetProjectId() *string
	SetProjectName(v string) *SaveAvatarProjectRequest
	GetProjectName() *string
	SetResSpecType(v string) *SaveAvatarProjectRequest
	GetResSpecType() *string
	SetResolution(v string) *SaveAvatarProjectRequest
	GetResolution() *string
	SetScaleType(v string) *SaveAvatarProjectRequest
	GetScaleType() *string
	SetScriptModelTag(v string) *SaveAvatarProjectRequest
	GetScriptModelTag() *string
	SetSynchronizedDisplay(v string) *SaveAvatarProjectRequest
	GetSynchronizedDisplay() *string
}

type SaveAvatarProjectRequest struct {
	// example:
	//
	// 1000196
	AgentId   *string                           `json:"agentId,omitempty" xml:"agentId,omitempty"`
	BitRate   *string                           `json:"bitRate,omitempty" xml:"bitRate,omitempty"`
	FrameRate *string                           `json:"frameRate,omitempty" xml:"frameRate,omitempty"`
	Frames    []*SaveAvatarProjectRequestFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// example:
	//
	// CREATE
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// 787594567117586432
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// df_cs_471437
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// STANDARD
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	Resolution  *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// example:
	//
	// 9:16
	ScaleType           *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	ScriptModelTag      *string `json:"scriptModelTag,omitempty" xml:"scriptModelTag,omitempty"`
	SynchronizedDisplay *string `json:"synchronizedDisplay,omitempty" xml:"synchronizedDisplay,omitempty"`
}

func (s SaveAvatarProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SaveAvatarProjectRequest) GetBitRate() *string {
	return s.BitRate
}

func (s *SaveAvatarProjectRequest) GetFrameRate() *string {
	return s.FrameRate
}

func (s *SaveAvatarProjectRequest) GetFrames() []*SaveAvatarProjectRequestFrames {
	return s.Frames
}

func (s *SaveAvatarProjectRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *SaveAvatarProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SaveAvatarProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SaveAvatarProjectRequest) GetResSpecType() *string {
	return s.ResSpecType
}

func (s *SaveAvatarProjectRequest) GetResolution() *string {
	return s.Resolution
}

func (s *SaveAvatarProjectRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *SaveAvatarProjectRequest) GetScriptModelTag() *string {
	return s.ScriptModelTag
}

func (s *SaveAvatarProjectRequest) GetSynchronizedDisplay() *string {
	return s.SynchronizedDisplay
}

func (s *SaveAvatarProjectRequest) SetAgentId(v string) *SaveAvatarProjectRequest {
	s.AgentId = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetBitRate(v string) *SaveAvatarProjectRequest {
	s.BitRate = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetFrameRate(v string) *SaveAvatarProjectRequest {
	s.FrameRate = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetFrames(v []*SaveAvatarProjectRequestFrames) *SaveAvatarProjectRequest {
	s.Frames = v
	return s
}

func (s *SaveAvatarProjectRequest) SetOperateType(v string) *SaveAvatarProjectRequest {
	s.OperateType = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetProjectId(v string) *SaveAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetProjectName(v string) *SaveAvatarProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetResSpecType(v string) *SaveAvatarProjectRequest {
	s.ResSpecType = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetResolution(v string) *SaveAvatarProjectRequest {
	s.Resolution = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetScaleType(v string) *SaveAvatarProjectRequest {
	s.ScaleType = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetScriptModelTag(v string) *SaveAvatarProjectRequest {
	s.ScriptModelTag = &v
	return s
}

func (s *SaveAvatarProjectRequest) SetSynchronizedDisplay(v string) *SaveAvatarProjectRequest {
	s.SynchronizedDisplay = &v
	return s
}

func (s *SaveAvatarProjectRequest) Validate() error {
	if s.Frames != nil {
		for _, item := range s.Frames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveAvatarProjectRequestFrames struct {
	Index       *int32                                     `json:"index,omitempty" xml:"index,omitempty"`
	Layers      []*SaveAvatarProjectRequestFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	VideoScript *SaveAvatarProjectRequestFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s SaveAvatarProjectRequestFrames) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectRequestFrames) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFrames) GetIndex() *int32 {
	return s.Index
}

func (s *SaveAvatarProjectRequestFrames) GetLayers() []*SaveAvatarProjectRequestFramesLayers {
	return s.Layers
}

func (s *SaveAvatarProjectRequestFrames) GetVideoScript() *SaveAvatarProjectRequestFramesVideoScript {
	return s.VideoScript
}

func (s *SaveAvatarProjectRequestFrames) SetIndex(v int32) *SaveAvatarProjectRequestFrames {
	s.Index = &v
	return s
}

func (s *SaveAvatarProjectRequestFrames) SetLayers(v []*SaveAvatarProjectRequestFramesLayers) *SaveAvatarProjectRequestFrames {
	s.Layers = v
	return s
}

func (s *SaveAvatarProjectRequestFrames) SetVideoScript(v *SaveAvatarProjectRequestFramesVideoScript) *SaveAvatarProjectRequestFrames {
	s.VideoScript = v
	return s
}

func (s *SaveAvatarProjectRequestFrames) Validate() error {
	if s.Layers != nil {
		for _, item := range s.Layers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoScript != nil {
		if err := s.VideoScript.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveAvatarProjectRequestFramesLayers struct {
	// example:
	//
	// 100
	Height   *int32                                        `json:"height,omitempty" xml:"height,omitempty"`
	Index    *int32                                        `json:"index,omitempty" xml:"index,omitempty"`
	Material *SaveAvatarProjectRequestFramesLayersMaterial `json:"material,omitempty" xml:"material,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PositionX *int32 `json:"positionX,omitempty" xml:"positionX,omitempty"`
	// example:
	//
	// 1
	PositionY *int32 `json:"positionY,omitempty" xml:"positionY,omitempty"`
	// example:
	//
	// ANCHOR
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s SaveAvatarProjectRequestFramesLayers) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectRequestFramesLayers) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFramesLayers) GetHeight() *int32 {
	return s.Height
}

func (s *SaveAvatarProjectRequestFramesLayers) GetIndex() *int32 {
	return s.Index
}

func (s *SaveAvatarProjectRequestFramesLayers) GetMaterial() *SaveAvatarProjectRequestFramesLayersMaterial {
	return s.Material
}

func (s *SaveAvatarProjectRequestFramesLayers) GetPositionX() *int32 {
	return s.PositionX
}

func (s *SaveAvatarProjectRequestFramesLayers) GetPositionY() *int32 {
	return s.PositionY
}

func (s *SaveAvatarProjectRequestFramesLayers) GetType() *string {
	return s.Type
}

func (s *SaveAvatarProjectRequestFramesLayers) GetWidth() *int32 {
	return s.Width
}

func (s *SaveAvatarProjectRequestFramesLayers) SetHeight(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.Height = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetIndex(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.Index = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetMaterial(v *SaveAvatarProjectRequestFramesLayersMaterial) *SaveAvatarProjectRequestFramesLayers {
	s.Material = v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetPositionX(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.PositionX = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetPositionY(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.PositionY = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetType(v string) *SaveAvatarProjectRequestFramesLayers {
	s.Type = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) SetWidth(v int32) *SaveAvatarProjectRequestFramesLayers {
	s.Width = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayers) Validate() error {
	if s.Material != nil {
		if err := s.Material.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveAvatarProjectRequestFramesLayersMaterial struct {
	// example:
	//
	// image/png
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 434508
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// https://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SaveAvatarProjectRequestFramesLayersMaterial) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectRequestFramesLayersMaterial) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) GetFormat() *string {
	return s.Format
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) GetId() *string {
	return s.Id
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) GetUrl() *string {
	return s.Url
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) SetFormat(v string) *SaveAvatarProjectRequestFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) SetId(v string) *SaveAvatarProjectRequestFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) SetUrl(v string) *SaveAvatarProjectRequestFramesLayersMaterial {
	s.Url = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesLayersMaterial) Validate() error {
	return dara.Validate(s)
}

type SaveAvatarProjectRequestFramesVideoScript struct {
	Emotion   *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
	PitchRate *string `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	// example:
	//
	// 1.0
	SpeedRate     *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent   *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	// example:
	//
	// 1
	VoiceTemplateId *string `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
	// example:
	//
	// 50
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s SaveAvatarProjectRequestFramesVideoScript) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectRequestFramesVideoScript) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetEmotion() *string {
	return s.Emotion
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetPitchRate() *string {
	return s.PitchRate
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetSpeedRate() *string {
	return s.SpeedRate
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetTextContent() *string {
	return s.TextContent
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetVoiceTemplateId() *string {
	return s.VoiceTemplateId
}

func (s *SaveAvatarProjectRequestFramesVideoScript) GetVolume() *string {
	return s.Volume
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetEmotion(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.Emotion = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetPitchRate(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.PitchRate = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetSpeedRate(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetTextContent(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.TextContent = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetVoiceLanguage(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.VoiceLanguage = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetVoiceTemplateId(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) SetVolume(v string) *SaveAvatarProjectRequestFramesVideoScript {
	s.Volume = &v
	return s
}

func (s *SaveAvatarProjectRequestFramesVideoScript) Validate() error {
	return dara.Validate(s)
}
