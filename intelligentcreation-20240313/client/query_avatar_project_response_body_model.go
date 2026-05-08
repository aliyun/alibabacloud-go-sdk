// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvatarProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *QueryAvatarProjectResponseBody
	GetAgentId() *string
	SetErrorMsg(v string) *QueryAvatarProjectResponseBody
	GetErrorMsg() *string
	SetFrames(v []*QueryAvatarProjectResponseBodyFrames) *QueryAvatarProjectResponseBody
	GetFrames() []*QueryAvatarProjectResponseBodyFrames
	SetProjectName(v string) *QueryAvatarProjectResponseBody
	GetProjectName() *string
	SetRequestId(v string) *QueryAvatarProjectResponseBody
	GetRequestId() *string
	SetResSpecType(v string) *QueryAvatarProjectResponseBody
	GetResSpecType() *string
	SetScaleType(v string) *QueryAvatarProjectResponseBody
	GetScaleType() *string
	SetScriptModelTag(v string) *QueryAvatarProjectResponseBody
	GetScriptModelTag() *string
	SetStatus(v string) *QueryAvatarProjectResponseBody
	GetStatus() *string
}

type QueryAvatarProjectResponseBody struct {
	// example:
	//
	// 1000222
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                 `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Frames   []*QueryAvatarProjectResponseBodyFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// example:
	//
	// doc_test_3
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 2C331582-7390-5949-8D9A-AC8239185B37
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResSpecType    *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	ScaleType      *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	ScriptModelTag *string `json:"scriptModelTag,omitempty" xml:"scriptModelTag,omitempty"`
	// example:
	//
	// DEPLOYING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAvatarProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *QueryAvatarProjectResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryAvatarProjectResponseBody) GetFrames() []*QueryAvatarProjectResponseBodyFrames {
	return s.Frames
}

func (s *QueryAvatarProjectResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryAvatarProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvatarProjectResponseBody) GetResSpecType() *string {
	return s.ResSpecType
}

func (s *QueryAvatarProjectResponseBody) GetScaleType() *string {
	return s.ScaleType
}

func (s *QueryAvatarProjectResponseBody) GetScriptModelTag() *string {
	return s.ScriptModelTag
}

func (s *QueryAvatarProjectResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryAvatarProjectResponseBody) SetAgentId(v string) *QueryAvatarProjectResponseBody {
	s.AgentId = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetErrorMsg(v string) *QueryAvatarProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetFrames(v []*QueryAvatarProjectResponseBodyFrames) *QueryAvatarProjectResponseBody {
	s.Frames = v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetProjectName(v string) *QueryAvatarProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetRequestId(v string) *QueryAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetResSpecType(v string) *QueryAvatarProjectResponseBody {
	s.ResSpecType = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetScaleType(v string) *QueryAvatarProjectResponseBody {
	s.ScaleType = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetScriptModelTag(v string) *QueryAvatarProjectResponseBody {
	s.ScriptModelTag = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) SetStatus(v string) *QueryAvatarProjectResponseBody {
	s.Status = &v
	return s
}

func (s *QueryAvatarProjectResponseBody) Validate() error {
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

type QueryAvatarProjectResponseBodyFrames struct {
	Index       *int32                                           `json:"index,omitempty" xml:"index,omitempty"`
	Layers      []*QueryAvatarProjectResponseBodyFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	VideoScript *QueryAvatarProjectResponseBodyFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s QueryAvatarProjectResponseBodyFrames) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFrames) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFrames) GetIndex() *int32 {
	return s.Index
}

func (s *QueryAvatarProjectResponseBodyFrames) GetLayers() []*QueryAvatarProjectResponseBodyFramesLayers {
	return s.Layers
}

func (s *QueryAvatarProjectResponseBodyFrames) GetVideoScript() *QueryAvatarProjectResponseBodyFramesVideoScript {
	return s.VideoScript
}

func (s *QueryAvatarProjectResponseBodyFrames) SetIndex(v int32) *QueryAvatarProjectResponseBodyFrames {
	s.Index = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFrames) SetLayers(v []*QueryAvatarProjectResponseBodyFramesLayers) *QueryAvatarProjectResponseBodyFrames {
	s.Layers = v
	return s
}

func (s *QueryAvatarProjectResponseBodyFrames) SetVideoScript(v *QueryAvatarProjectResponseBodyFramesVideoScript) *QueryAvatarProjectResponseBodyFrames {
	s.VideoScript = v
	return s
}

func (s *QueryAvatarProjectResponseBodyFrames) Validate() error {
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

type QueryAvatarProjectResponseBodyFramesLayers struct {
	Height    *int32                                              `json:"height,omitempty" xml:"height,omitempty"`
	Index     *int32                                              `json:"index,omitempty" xml:"index,omitempty"`
	Material  *QueryAvatarProjectResponseBodyFramesLayersMaterial `json:"material,omitempty" xml:"material,omitempty" type:"Struct"`
	PositionX *int32                                              `json:"positionX,omitempty" xml:"positionX,omitempty"`
	PositionY *int32                                              `json:"positionY,omitempty" xml:"positionY,omitempty"`
	Type      *string                                             `json:"type,omitempty" xml:"type,omitempty"`
	Width     *int32                                              `json:"width,omitempty" xml:"width,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesLayers) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesLayers) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetHeight() *int32 {
	return s.Height
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetIndex() *int32 {
	return s.Index
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetMaterial() *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	return s.Material
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetPositionX() *int32 {
	return s.PositionX
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetPositionY() *int32 {
	return s.PositionY
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetType() *string {
	return s.Type
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) GetWidth() *int32 {
	return s.Width
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetHeight(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Height = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetIndex(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Index = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetMaterial(v *QueryAvatarProjectResponseBodyFramesLayersMaterial) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Material = v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetPositionX(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.PositionX = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetPositionY(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.PositionY = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetType(v string) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Type = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) SetWidth(v int32) *QueryAvatarProjectResponseBodyFramesLayers {
	s.Width = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayers) Validate() error {
	if s.Material != nil {
		if err := s.Material.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAvatarProjectResponseBodyFramesLayersMaterial struct {
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesLayersMaterial) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesLayersMaterial) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) GetFormat() *string {
	return s.Format
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) GetId() *string {
	return s.Id
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) GetUrl() *string {
	return s.Url
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) SetFormat(v string) *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) SetId(v string) *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) SetUrl(v string) *QueryAvatarProjectResponseBodyFramesLayersMaterial {
	s.Url = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesLayersMaterial) Validate() error {
	return dara.Validate(s)
}

type QueryAvatarProjectResponseBodyFramesVideoScript struct {
	Emotion         *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
	PitchRate       *string `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	SpeedRate       *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent     *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	VoiceLanguage   *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	VoiceTemplateId *string `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
	Volume          *int32  `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s QueryAvatarProjectResponseBodyFramesVideoScript) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectResponseBodyFramesVideoScript) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetEmotion() *string {
	return s.Emotion
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetPitchRate() *string {
	return s.PitchRate
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetSpeedRate() *string {
	return s.SpeedRate
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetTextContent() *string {
	return s.TextContent
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetVoiceTemplateId() *string {
	return s.VoiceTemplateId
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) GetVolume() *int32 {
	return s.Volume
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetEmotion(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.Emotion = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetPitchRate(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.PitchRate = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetSpeedRate(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetTextContent(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.TextContent = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVoiceLanguage(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.VoiceLanguage = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVoiceTemplateId(v string) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) SetVolume(v int32) *QueryAvatarProjectResponseBodyFramesVideoScript {
	s.Volume = &v
	return s
}

func (s *QueryAvatarProjectResponseBodyFramesVideoScript) Validate() error {
	return dara.Validate(s)
}
