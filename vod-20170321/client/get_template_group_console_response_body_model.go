// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateGroupConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTemplateGroupConsoleResponseBody
	GetRequestId() *string
	SetTemplateGroup(v *GetTemplateGroupConsoleResponseBodyTemplateGroup) *GetTemplateGroupConsoleResponseBody
	GetTemplateGroup() *GetTemplateGroupConsoleResponseBodyTemplateGroup
}

type GetTemplateGroupConsoleResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateGroup *GetTemplateGroupConsoleResponseBodyTemplateGroup `json:"TemplateGroup,omitempty" xml:"TemplateGroup,omitempty" type:"Struct"`
}

func (s GetTemplateGroupConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateGroupConsoleResponseBody) GetTemplateGroup() *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	return s.TemplateGroup
}

func (s *GetTemplateGroupConsoleResponseBody) SetRequestId(v string) *GetTemplateGroupConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBody) SetTemplateGroup(v *GetTemplateGroupConsoleResponseBodyTemplateGroup) *GetTemplateGroupConsoleResponseBody {
	s.TemplateGroup = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroup struct {
	DefaultGroup  *string                                                    `json:"DefaultGroup,omitempty" xml:"DefaultGroup,omitempty"`
	GroupId       *string                                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupSymbol   *string                                                    `json:"GroupSymbol,omitempty" xml:"GroupSymbol,omitempty"`
	GroupType     *string                                                    `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IsLocked      *string                                                    `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	Name          *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Status        *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Templates     *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	TranscodeMode *string                                                    `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroup) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroup) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetDefaultGroup() *string {
	return s.DefaultGroup
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetGroupSymbol() *string {
	return s.GroupSymbol
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetIsLocked() *string {
	return s.IsLocked
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetName() *string {
	return s.Name
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetStatus() *string {
	return s.Status
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetTemplates() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates {
	return s.Templates
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) GetTranscodeMode() *string {
	return s.TranscodeMode
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetDefaultGroup(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.DefaultGroup = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetGroupId(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.GroupId = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetGroupSymbol(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.GroupSymbol = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetGroupType(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.GroupType = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetIsLocked(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.IsLocked = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetName(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.Name = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetStatus(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.Status = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetTemplates(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.Templates = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) SetTranscodeMode(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroup {
	s.TranscodeMode = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroup) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates struct {
	Template []*GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates) GetTemplate() []*GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	return s.Template
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates) SetTemplate(v []*GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates {
	s.Template = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplates) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate struct {
	Audio            *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio            `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	Condition        *string                                                                            `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Container        *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer        `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	Definition       *string                                                                            `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Encrypt          *string                                                                            `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	EncryptionConfig *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig `json:"EncryptionConfig,omitempty" xml:"EncryptionConfig,omitempty" type:"Struct"`
	IsLocked         *string                                                                            `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	MediaDefinition  *string                                                                            `json:"MediaDefinition,omitempty" xml:"MediaDefinition,omitempty"`
	MuxConfig        *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig        `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	Name             *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	NarrowBand       *string                                                                            `json:"NarrowBand,omitempty" xml:"NarrowBand,omitempty"`
	Status           *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId       *string                                                                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TransConfig      *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig      `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	UseWaterMark     *string                                                                            `json:"UseWaterMark,omitempty" xml:"UseWaterMark,omitempty"`
	UserWaterMark    *string                                                                            `json:"UserWaterMark,omitempty" xml:"UserWaterMark,omitempty"`
	Video            *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo            `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetAudio() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	return s.Audio
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetCondition() *string {
	return s.Condition
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetContainer() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer {
	return s.Container
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetDefinition() *string {
	return s.Definition
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetEncrypt() *string {
	return s.Encrypt
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetEncryptionConfig() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig {
	return s.EncryptionConfig
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetIsLocked() *string {
	return s.IsLocked
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetMediaDefinition() *string {
	return s.MediaDefinition
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetMuxConfig() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig {
	return s.MuxConfig
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetName() *string {
	return s.Name
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetNarrowBand() *string {
	return s.NarrowBand
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetStatus() *string {
	return s.Status
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetTransConfig() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	return s.TransConfig
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetUseWaterMark() *string {
	return s.UseWaterMark
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetUserWaterMark() *string {
	return s.UserWaterMark
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) GetVideo() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	return s.Video
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetAudio(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Audio = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetCondition(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Condition = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetContainer(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Container = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetDefinition(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Definition = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetEncrypt(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Encrypt = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetEncryptionConfig(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.EncryptionConfig = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetIsLocked(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.IsLocked = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetMediaDefinition(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.MediaDefinition = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetMuxConfig(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.MuxConfig = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetName(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Name = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetNarrowBand(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.NarrowBand = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetStatus(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Status = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetTemplateId(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetTransConfig(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.TransConfig = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetUseWaterMark(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.UseWaterMark = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetUserWaterMark(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.UserWaterMark = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) SetVideo(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate {
	s.Video = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplate) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio struct {
	Bitrate    *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Channels   *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Profile    *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Remove     *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GetChannels() *string {
	return s.Channels
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GetCodec() *string {
	return s.Codec
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GetProfile() *string {
	return s.Profile
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GetRemove() *string {
	return s.Remove
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) SetBitrate(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	s.Bitrate = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) SetChannels(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	s.Channels = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) SetCodec(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	s.Codec = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) SetProfile(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	s.Profile = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) SetRemove(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	s.Remove = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) SetSamplerate(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio {
	s.Samplerate = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateAudio) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer) GetFormat() *string {
	return s.Format
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer) SetFormat(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer {
	s.Format = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateContainer) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig struct {
	DecryptKeyUri    *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	EncryptType      *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	KeyEncryptMethod *string `json:"KeyEncryptMethod,omitempty" xml:"KeyEncryptMethod,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) GetEncryptType() *string {
	return s.EncryptType
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) GetKeyEncryptMethod() *string {
	return s.KeyEncryptMethod
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) SetDecryptKeyUri(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig {
	s.DecryptKeyUri = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) SetEncryptType(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig {
	s.EncryptType = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) SetKeyEncryptMethod(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig {
	s.KeyEncryptMethod = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateEncryptionConfig) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig struct {
	Segment *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig) GetSegment() *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment {
	return s.Segment
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig) SetSegment(v *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig {
	s.Segment = v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfig) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment) SetDuration(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig struct {
	IsCheckAudioBitrate     *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	IsCheckReso             *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	IsCheckResoFail         *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	IsCheckVideoBitrate     *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	TransMode               *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetIsCheckAudioBitrate(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetIsCheckAudioBitrateFail(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetIsCheckReso(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetIsCheckResoFail(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetIsCheckVideoBitrate(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetIsCheckVideoBitrateFail(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) SetTransMode(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig {
	s.TransMode = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateTransConfig) Validate() error {
	return dara.Validate(s)
}

type GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo struct {
	Bitrate       *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Bufsize       *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	Codec         *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Crf           *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	Crop          *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Fps           *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Gop           *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Height        *string `json:"Height,omitempty" xml:"Height,omitempty"`
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	Maxrate       *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	Pad           *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	PixFmt        *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Preset        *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	Profile       *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Remove        *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	ScanMode      *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	Width         *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetCodec() *string {
	return s.Codec
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetCrf() *string {
	return s.Crf
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetCrop() *string {
	return s.Crop
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetFps() *string {
	return s.Fps
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetGop() *string {
	return s.Gop
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetHeight() *string {
	return s.Height
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetPad() *string {
	return s.Pad
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetPreset() *string {
	return s.Preset
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetProfile() *string {
	return s.Profile
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetRemove() *string {
	return s.Remove
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) GetWidth() *string {
	return s.Width
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetBitrate(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Bitrate = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetBufsize(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Bufsize = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetCodec(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Codec = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetCrf(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Crf = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetCrop(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Crop = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetFps(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Fps = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetGop(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Gop = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetHeight(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Height = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetLongShortMode(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.LongShortMode = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetMaxrate(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Maxrate = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetPad(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Pad = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetPixFmt(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.PixFmt = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetPreset(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Preset = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetProfile(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Profile = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetRemove(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Remove = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetScanMode(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.ScanMode = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) SetWidth(v string) *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo {
	s.Width = &v
	return s
}

func (s *GetTemplateGroupConsoleResponseBodyTemplateGroupTemplatesTemplateVideo) Validate() error {
	return dara.Validate(s)
}
