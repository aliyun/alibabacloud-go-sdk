// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Address struct {
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District    *string `json:"District,omitempty" xml:"District,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township    *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s Address) String() string {
	return tea.Prettify(s)
}

func (s Address) GoString() string {
	return s.String()
}

func (s *Address) SetAddressLine(v string) *Address {
	s.AddressLine = &v
	return s
}

func (s *Address) SetCity(v string) *Address {
	s.City = &v
	return s
}

func (s *Address) SetCountry(v string) *Address {
	s.Country = &v
	return s
}

func (s *Address) SetDistrict(v string) *Address {
	s.District = &v
	return s
}

func (s *Address) SetLanguage(v string) *Address {
	s.Language = &v
	return s
}

func (s *Address) SetProvince(v string) *Address {
	s.Province = &v
	return s
}

func (s *Address) SetTownship(v string) *Address {
	s.Township = &v
	return s
}

type AssumeRoleChain struct {
	Chain  []*AssumeRoleChainNode `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	Policy *string                `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s AssumeRoleChain) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleChain) GoString() string {
	return s.String()
}

func (s *AssumeRoleChain) SetChain(v []*AssumeRoleChainNode) *AssumeRoleChain {
	s.Chain = v
	return s
}

func (s *AssumeRoleChain) SetPolicy(v string) *AssumeRoleChain {
	s.Policy = &v
	return s
}

type AssumeRoleChainNode struct {
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AssumeRoleChainNode) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleChainNode) GoString() string {
	return s.String()
}

func (s *AssumeRoleChainNode) SetOwnerId(v string) *AssumeRoleChainNode {
	s.OwnerId = &v
	return s
}

func (s *AssumeRoleChainNode) SetRole(v string) *AssumeRoleChainNode {
	s.Role = &v
	return s
}

func (s *AssumeRoleChainNode) SetType(v string) *AssumeRoleChainNode {
	s.Type = &v
	return s
}

type AudioStream struct {
	Bitrate        *int64   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	ChannelLayout  *string  `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	Channels       *int64   `json:"Channels,omitempty" xml:"Channels,omitempty"`
	CodecLongName  *string  `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string  `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string  `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string  `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string  `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration       *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FrameCount     *int64   `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	Index          *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
	Language       *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	Lyric          *string  `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	SampleFormat   *string  `json:"SampleFormat,omitempty" xml:"SampleFormat,omitempty"`
	SampleRate     *int64   `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	StartTime      *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeBase       *string  `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
}

func (s AudioStream) String() string {
	return tea.Prettify(s)
}

func (s AudioStream) GoString() string {
	return s.String()
}

func (s *AudioStream) SetBitrate(v int64) *AudioStream {
	s.Bitrate = &v
	return s
}

func (s *AudioStream) SetChannelLayout(v string) *AudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *AudioStream) SetChannels(v int64) *AudioStream {
	s.Channels = &v
	return s
}

func (s *AudioStream) SetCodecLongName(v string) *AudioStream {
	s.CodecLongName = &v
	return s
}

func (s *AudioStream) SetCodecName(v string) *AudioStream {
	s.CodecName = &v
	return s
}

func (s *AudioStream) SetCodecTag(v string) *AudioStream {
	s.CodecTag = &v
	return s
}

func (s *AudioStream) SetCodecTagString(v string) *AudioStream {
	s.CodecTagString = &v
	return s
}

func (s *AudioStream) SetCodecTimeBase(v string) *AudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *AudioStream) SetDuration(v float64) *AudioStream {
	s.Duration = &v
	return s
}

func (s *AudioStream) SetFrameCount(v int64) *AudioStream {
	s.FrameCount = &v
	return s
}

func (s *AudioStream) SetIndex(v int64) *AudioStream {
	s.Index = &v
	return s
}

func (s *AudioStream) SetLanguage(v string) *AudioStream {
	s.Language = &v
	return s
}

func (s *AudioStream) SetLyric(v string) *AudioStream {
	s.Lyric = &v
	return s
}

func (s *AudioStream) SetSampleFormat(v string) *AudioStream {
	s.SampleFormat = &v
	return s
}

func (s *AudioStream) SetSampleRate(v int64) *AudioStream {
	s.SampleRate = &v
	return s
}

func (s *AudioStream) SetStartTime(v float64) *AudioStream {
	s.StartTime = &v
	return s
}

func (s *AudioStream) SetTimeBase(v string) *AudioStream {
	s.TimeBase = &v
	return s
}

type Binding struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Detail      *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Binding) String() string {
	return tea.Prettify(s)
}

func (s Binding) GoString() string {
	return s.String()
}

func (s *Binding) SetCreateTime(v string) *Binding {
	s.CreateTime = &v
	return s
}

func (s *Binding) SetDatasetName(v string) *Binding {
	s.DatasetName = &v
	return s
}

func (s *Binding) SetDetail(v string) *Binding {
	s.Detail = &v
	return s
}

func (s *Binding) SetPhase(v string) *Binding {
	s.Phase = &v
	return s
}

func (s *Binding) SetProjectName(v string) *Binding {
	s.ProjectName = &v
	return s
}

func (s *Binding) SetState(v string) *Binding {
	s.State = &v
	return s
}

func (s *Binding) SetURI(v string) *Binding {
	s.URI = &v
	return s
}

func (s *Binding) SetUpdateTime(v string) *Binding {
	s.UpdateTime = &v
	return s
}

type Body struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s Body) String() string {
	return tea.Prettify(s)
}

func (s Body) GoString() string {
	return s.String()
}

func (s *Body) SetBoundary(v *Boundary) *Body {
	s.Boundary = v
	return s
}

func (s *Body) SetConfidence(v float32) *Body {
	s.Confidence = &v
	return s
}

type Boundary struct {
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s Boundary) String() string {
	return tea.Prettify(s)
}

func (s Boundary) GoString() string {
	return s.String()
}

func (s *Boundary) SetHeight(v int64) *Boundary {
	s.Height = &v
	return s
}

func (s *Boundary) SetLeft(v int64) *Boundary {
	s.Left = &v
	return s
}

func (s *Boundary) SetTop(v int64) *Boundary {
	s.Top = &v
	return s
}

func (s *Boundary) SetWidth(v int64) *Boundary {
	s.Width = &v
	return s
}

type ClusterForReq struct {
	Cover        *ClusterForReqCover    `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectId     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
}

func (s ClusterForReq) String() string {
	return tea.Prettify(s)
}

func (s ClusterForReq) GoString() string {
	return s.String()
}

func (s *ClusterForReq) SetCover(v *ClusterForReqCover) *ClusterForReq {
	s.Cover = v
	return s
}

func (s *ClusterForReq) SetCustomId(v string) *ClusterForReq {
	s.CustomId = &v
	return s
}

func (s *ClusterForReq) SetCustomLabels(v map[string]interface{}) *ClusterForReq {
	s.CustomLabels = v
	return s
}

func (s *ClusterForReq) SetName(v string) *ClusterForReq {
	s.Name = &v
	return s
}

func (s *ClusterForReq) SetObjectId(v string) *ClusterForReq {
	s.ObjectId = &v
	return s
}

type ClusterForReqCover struct {
	Figures []*ClusterForReqCoverFigures `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
}

func (s ClusterForReqCover) String() string {
	return tea.Prettify(s)
}

func (s ClusterForReqCover) GoString() string {
	return s.String()
}

func (s *ClusterForReqCover) SetFigures(v []*ClusterForReqCoverFigures) *ClusterForReqCover {
	s.Figures = v
	return s
}

type ClusterForReqCoverFigures struct {
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
}

func (s ClusterForReqCoverFigures) String() string {
	return tea.Prettify(s)
}

func (s ClusterForReqCoverFigures) GoString() string {
	return s.String()
}

func (s *ClusterForReqCoverFigures) SetFigureId(v string) *ClusterForReqCoverFigures {
	s.FigureId = &v
	return s
}

type Codes struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Content    *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Codes) String() string {
	return tea.Prettify(s)
}

func (s Codes) GoString() string {
	return s.String()
}

func (s *Codes) SetBoundary(v *Boundary) *Codes {
	s.Boundary = v
	return s
}

func (s *Codes) SetConfidence(v float32) *Codes {
	s.Confidence = &v
	return s
}

func (s *Codes) SetContent(v string) *Codes {
	s.Content = &v
	return s
}

func (s *Codes) SetType(v string) *Codes {
	s.Type = &v
	return s
}

type CredentialConfig struct {
	Chain       []*CredentialConfigChain `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	Policy      *string                  `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ServiceRole *string                  `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
}

func (s CredentialConfig) String() string {
	return tea.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) SetChain(v []*CredentialConfigChain) *CredentialConfig {
	s.Chain = v
	return s
}

func (s *CredentialConfig) SetPolicy(v string) *CredentialConfig {
	s.Policy = &v
	return s
}

func (s *CredentialConfig) SetServiceRole(v string) *CredentialConfig {
	s.ServiceRole = &v
	return s
}

type CredentialConfigChain struct {
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CredentialConfigChain) String() string {
	return tea.Prettify(s)
}

func (s CredentialConfigChain) GoString() string {
	return s.String()
}

func (s *CredentialConfigChain) SetAssumeRoleFor(v string) *CredentialConfigChain {
	s.AssumeRoleFor = &v
	return s
}

func (s *CredentialConfigChain) SetRole(v string) *CredentialConfigChain {
	s.Role = &v
	return s
}

func (s *CredentialConfigChain) SetRoleType(v string) *CredentialConfigChain {
	s.RoleType = &v
	return s
}

type CroppingSuggestion struct {
	AspectRatio *string   `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	Boundary    *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence  *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s CroppingSuggestion) String() string {
	return tea.Prettify(s)
}

func (s CroppingSuggestion) GoString() string {
	return s.String()
}

func (s *CroppingSuggestion) SetAspectRatio(v string) *CroppingSuggestion {
	s.AspectRatio = &v
	return s
}

func (s *CroppingSuggestion) SetBoundary(v *Boundary) *CroppingSuggestion {
	s.Boundary = v
	return s
}

func (s *CroppingSuggestion) SetConfidence(v float32) *CroppingSuggestion {
	s.Confidence = &v
	return s
}

type Dataset struct {
	BindCount               *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileCount               *int64  `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TotalFileSize           *int64  `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	UpdateTime              *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Dataset) String() string {
	return tea.Prettify(s)
}

func (s Dataset) GoString() string {
	return s.String()
}

func (s *Dataset) SetBindCount(v int64) *Dataset {
	s.BindCount = &v
	return s
}

func (s *Dataset) SetCreateTime(v string) *Dataset {
	s.CreateTime = &v
	return s
}

func (s *Dataset) SetDatasetMaxBindCount(v int64) *Dataset {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxEntityCount(v int64) *Dataset {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxFileCount(v int64) *Dataset {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxRelationCount(v int64) *Dataset {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxTotalFileSize(v int64) *Dataset {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *Dataset) SetDatasetName(v string) *Dataset {
	s.DatasetName = &v
	return s
}

func (s *Dataset) SetDescription(v string) *Dataset {
	s.Description = &v
	return s
}

func (s *Dataset) SetFileCount(v int64) *Dataset {
	s.FileCount = &v
	return s
}

func (s *Dataset) SetProjectName(v string) *Dataset {
	s.ProjectName = &v
	return s
}

func (s *Dataset) SetTotalFileSize(v int64) *Dataset {
	s.TotalFileSize = &v
	return s
}

func (s *Dataset) SetUpdateTime(v string) *Dataset {
	s.UpdateTime = &v
	return s
}

type Figure struct {
	Age                     *int64    `json:"Age,omitempty" xml:"Age,omitempty"`
	AgeSD                   *float32  `json:"AgeSD,omitempty" xml:"AgeSD,omitempty"`
	Attractive              *float32  `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	Beard                   *string   `json:"Beard,omitempty" xml:"Beard,omitempty"`
	BeardConfidence         *float32  `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	Boundary                *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Emotion                 *string   `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	EmotionConfidence       *float32  `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	FaceQuality             *float32  `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	FigureClusterConfidence *float32  `json:"FigureClusterConfidence,omitempty" xml:"FigureClusterConfidence,omitempty"`
	FigureClusterId         *string   `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	FigureConfidence        *float32  `json:"FigureConfidence,omitempty" xml:"FigureConfidence,omitempty"`
	FigureId                *string   `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	FigureType              *string   `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
	Gender                  *string   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderConfidence        *float32  `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	Glasses                 *string   `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	GlassesConfidence       *float32  `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	Hat                     *string   `json:"Hat,omitempty" xml:"Hat,omitempty"`
	HatConfidence           *float32  `json:"HatConfidence,omitempty" xml:"HatConfidence,omitempty"`
	HeadPose                *HeadPose `json:"HeadPose,omitempty" xml:"HeadPose,omitempty"`
	Mask                    *string   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaskConfidence          *float32  `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	Mouth                   *string   `json:"Mouth,omitempty" xml:"Mouth,omitempty"`
	MouthConfidence         *float32  `json:"MouthConfidence,omitempty" xml:"MouthConfidence,omitempty"`
	Sharpness               *float32  `json:"Sharpness,omitempty" xml:"Sharpness,omitempty"`
}

func (s Figure) String() string {
	return tea.Prettify(s)
}

func (s Figure) GoString() string {
	return s.String()
}

func (s *Figure) SetAge(v int64) *Figure {
	s.Age = &v
	return s
}

func (s *Figure) SetAgeSD(v float32) *Figure {
	s.AgeSD = &v
	return s
}

func (s *Figure) SetAttractive(v float32) *Figure {
	s.Attractive = &v
	return s
}

func (s *Figure) SetBeard(v string) *Figure {
	s.Beard = &v
	return s
}

func (s *Figure) SetBeardConfidence(v float32) *Figure {
	s.BeardConfidence = &v
	return s
}

func (s *Figure) SetBoundary(v *Boundary) *Figure {
	s.Boundary = v
	return s
}

func (s *Figure) SetEmotion(v string) *Figure {
	s.Emotion = &v
	return s
}

func (s *Figure) SetEmotionConfidence(v float32) *Figure {
	s.EmotionConfidence = &v
	return s
}

func (s *Figure) SetFaceQuality(v float32) *Figure {
	s.FaceQuality = &v
	return s
}

func (s *Figure) SetFigureClusterConfidence(v float32) *Figure {
	s.FigureClusterConfidence = &v
	return s
}

func (s *Figure) SetFigureClusterId(v string) *Figure {
	s.FigureClusterId = &v
	return s
}

func (s *Figure) SetFigureConfidence(v float32) *Figure {
	s.FigureConfidence = &v
	return s
}

func (s *Figure) SetFigureId(v string) *Figure {
	s.FigureId = &v
	return s
}

func (s *Figure) SetFigureType(v string) *Figure {
	s.FigureType = &v
	return s
}

func (s *Figure) SetGender(v string) *Figure {
	s.Gender = &v
	return s
}

func (s *Figure) SetGenderConfidence(v float32) *Figure {
	s.GenderConfidence = &v
	return s
}

func (s *Figure) SetGlasses(v string) *Figure {
	s.Glasses = &v
	return s
}

func (s *Figure) SetGlassesConfidence(v float32) *Figure {
	s.GlassesConfidence = &v
	return s
}

func (s *Figure) SetHat(v string) *Figure {
	s.Hat = &v
	return s
}

func (s *Figure) SetHatConfidence(v float32) *Figure {
	s.HatConfidence = &v
	return s
}

func (s *Figure) SetHeadPose(v *HeadPose) *Figure {
	s.HeadPose = v
	return s
}

func (s *Figure) SetMask(v string) *Figure {
	s.Mask = &v
	return s
}

func (s *Figure) SetMaskConfidence(v float32) *Figure {
	s.MaskConfidence = &v
	return s
}

func (s *Figure) SetMouth(v string) *Figure {
	s.Mouth = &v
	return s
}

func (s *Figure) SetMouthConfidence(v float32) *Figure {
	s.MouthConfidence = &v
	return s
}

func (s *Figure) SetSharpness(v float32) *Figure {
	s.Sharpness = &v
	return s
}

type FigureCluster struct {
	AverageAge   *float32               `json:"AverageAge,omitempty" xml:"AverageAge,omitempty"`
	Cover        *File                  `json:"Cover,omitempty" xml:"Cover,omitempty"`
	CreateTime   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FaceCount    *int64                 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	Gender       *string                `json:"Gender,omitempty" xml:"Gender,omitempty"`
	ImageCount   *int64                 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	MaxAge       *float32               `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	MinAge       *float32               `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectId     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType   *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId      *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	UpdateTime   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VideoCount   *int64                 `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
}

func (s FigureCluster) String() string {
	return tea.Prettify(s)
}

func (s FigureCluster) GoString() string {
	return s.String()
}

func (s *FigureCluster) SetAverageAge(v float32) *FigureCluster {
	s.AverageAge = &v
	return s
}

func (s *FigureCluster) SetCover(v *File) *FigureCluster {
	s.Cover = v
	return s
}

func (s *FigureCluster) SetCreateTime(v string) *FigureCluster {
	s.CreateTime = &v
	return s
}

func (s *FigureCluster) SetCustomId(v string) *FigureCluster {
	s.CustomId = &v
	return s
}

func (s *FigureCluster) SetCustomLabels(v map[string]interface{}) *FigureCluster {
	s.CustomLabels = v
	return s
}

func (s *FigureCluster) SetDatasetName(v string) *FigureCluster {
	s.DatasetName = &v
	return s
}

func (s *FigureCluster) SetFaceCount(v int64) *FigureCluster {
	s.FaceCount = &v
	return s
}

func (s *FigureCluster) SetGender(v string) *FigureCluster {
	s.Gender = &v
	return s
}

func (s *FigureCluster) SetImageCount(v int64) *FigureCluster {
	s.ImageCount = &v
	return s
}

func (s *FigureCluster) SetMaxAge(v float32) *FigureCluster {
	s.MaxAge = &v
	return s
}

func (s *FigureCluster) SetMinAge(v float32) *FigureCluster {
	s.MinAge = &v
	return s
}

func (s *FigureCluster) SetName(v string) *FigureCluster {
	s.Name = &v
	return s
}

func (s *FigureCluster) SetObjectId(v string) *FigureCluster {
	s.ObjectId = &v
	return s
}

func (s *FigureCluster) SetObjectType(v string) *FigureCluster {
	s.ObjectType = &v
	return s
}

func (s *FigureCluster) SetOwnerId(v string) *FigureCluster {
	s.OwnerId = &v
	return s
}

func (s *FigureCluster) SetProjectName(v string) *FigureCluster {
	s.ProjectName = &v
	return s
}

func (s *FigureCluster) SetUpdateTime(v string) *FigureCluster {
	s.UpdateTime = &v
	return s
}

func (s *FigureCluster) SetVideoCount(v int64) *FigureCluster {
	s.VideoCount = &v
	return s
}

type FigureClusterForReq struct {
	Cover        *FigureClusterForReqCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	CustomId     *string                   `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{}    `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Name         *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectId     *string                   `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
}

func (s FigureClusterForReq) String() string {
	return tea.Prettify(s)
}

func (s FigureClusterForReq) GoString() string {
	return s.String()
}

func (s *FigureClusterForReq) SetCover(v *FigureClusterForReqCover) *FigureClusterForReq {
	s.Cover = v
	return s
}

func (s *FigureClusterForReq) SetCustomId(v string) *FigureClusterForReq {
	s.CustomId = &v
	return s
}

func (s *FigureClusterForReq) SetCustomLabels(v map[string]interface{}) *FigureClusterForReq {
	s.CustomLabels = v
	return s
}

func (s *FigureClusterForReq) SetName(v string) *FigureClusterForReq {
	s.Name = &v
	return s
}

func (s *FigureClusterForReq) SetObjectId(v string) *FigureClusterForReq {
	s.ObjectId = &v
	return s
}

type FigureClusterForReqCover struct {
	Figures []*FigureClusterForReqCoverFigures `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
}

func (s FigureClusterForReqCover) String() string {
	return tea.Prettify(s)
}

func (s FigureClusterForReqCover) GoString() string {
	return s.String()
}

func (s *FigureClusterForReqCover) SetFigures(v []*FigureClusterForReqCoverFigures) *FigureClusterForReqCover {
	s.Figures = v
	return s
}

type FigureClusterForReqCoverFigures struct {
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
}

func (s FigureClusterForReqCoverFigures) String() string {
	return tea.Prettify(s)
}

func (s FigureClusterForReqCoverFigures) GoString() string {
	return s.String()
}

func (s *FigureClusterForReqCoverFigures) SetFigureId(v string) *FigureClusterForReqCoverFigures {
	s.FigureId = &v
	return s
}

type File struct {
	AccessControlAllowOrigin              *string                `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	AccessControlRequestMethod            *string                `json:"AccessControlRequestMethod,omitempty" xml:"AccessControlRequestMethod,omitempty"`
	Addresses                             []*Address             `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	Album                                 *string                `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumArtist                           *string                `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	Artist                                *string                `json:"Artist,omitempty" xml:"Artist,omitempty"`
	AudioCovers                           []*Image               `json:"AudioCovers,omitempty" xml:"AudioCovers,omitempty" type:"Repeated"`
	AudioStreams                          []*AudioStream         `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	Bitrate                               *int64                 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CacheControl                          *string                `json:"CacheControl,omitempty" xml:"CacheControl,omitempty"`
	Composer                              *string                `json:"Composer,omitempty" xml:"Composer,omitempty"`
	ContentDisposition                    *string                `json:"ContentDisposition,omitempty" xml:"ContentDisposition,omitempty"`
	ContentEncoding                       *string                `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	ContentLanguage                       *string                `json:"ContentLanguage,omitempty" xml:"ContentLanguage,omitempty"`
	ContentMd5                            *string                `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	ContentType                           *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CreateTime                            *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CroppingSuggestions                   []*CroppingSuggestion  `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	CustomId                              *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels                          map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName                           *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DocumentContent                       *string                `json:"DocumentContent,omitempty" xml:"DocumentContent,omitempty"`
	DocumentLanguage                      *string                `json:"DocumentLanguage,omitempty" xml:"DocumentLanguage,omitempty"`
	Duration                              *float64               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ETag                                  *string                `json:"ETag,omitempty" xml:"ETag,omitempty"`
	EXIF                                  *string                `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	FigureCount                           *int64                 `json:"FigureCount,omitempty" xml:"FigureCount,omitempty"`
	Figures                               []*Figure              `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	FileAccessTime                        *string                `json:"FileAccessTime,omitempty" xml:"FileAccessTime,omitempty"`
	FileCreateTime                        *string                `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	FileHash                              *string                `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	FileModifiedTime                      *string                `json:"FileModifiedTime,omitempty" xml:"FileModifiedTime,omitempty"`
	Filename                              *string                `json:"Filename,omitempty" xml:"Filename,omitempty"`
	FormatLongName                        *string                `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName                            *string                `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	ImageHeight                           *int64                 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageScore                            *ImageScore            `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	ImageWidth                            *int64                 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	Labels                                []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Language                              *string                `json:"Language,omitempty" xml:"Language,omitempty"`
	LatLong                               *string                `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	MediaType                             *string                `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OCRContents                           []*OCRContents         `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
	OSSCRC64                              *string                `json:"OSSCRC64,omitempty" xml:"OSSCRC64,omitempty"`
	OSSDeleteMarker                       *string                `json:"OSSDeleteMarker,omitempty" xml:"OSSDeleteMarker,omitempty"`
	OSSExpiration                         *string                `json:"OSSExpiration,omitempty" xml:"OSSExpiration,omitempty"`
	OSSObjectType                         *string                `json:"OSSObjectType,omitempty" xml:"OSSObjectType,omitempty"`
	OSSStorageClass                       *string                `json:"OSSStorageClass,omitempty" xml:"OSSStorageClass,omitempty"`
	OSSTagging                            map[string]interface{} `json:"OSSTagging,omitempty" xml:"OSSTagging,omitempty"`
	OSSTaggingCount                       *int64                 `json:"OSSTaggingCount,omitempty" xml:"OSSTaggingCount,omitempty"`
	OSSURI                                *string                `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	OSSUserMeta                           map[string]interface{} `json:"OSSUserMeta,omitempty" xml:"OSSUserMeta,omitempty"`
	OSSVersionId                          *string                `json:"OSSVersionId,omitempty" xml:"OSSVersionId,omitempty"`
	ObjectACL                             *string                `json:"ObjectACL,omitempty" xml:"ObjectACL,omitempty"`
	ObjectId                              *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType                            *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Orientation                           *int64                 `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	OwnerId                               *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageCount                             *int64                 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Performer                             *string                `json:"Performer,omitempty" xml:"Performer,omitempty"`
	ProduceTime                           *string                `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	ProgramCount                          *int64                 `json:"ProgramCount,omitempty" xml:"ProgramCount,omitempty"`
	ProjectName                           *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ServerSideDataEncryption              *string                `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	ServerSideEncryption                  *string                `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	ServerSideEncryptionCustomerAlgorithm *string                `json:"ServerSideEncryptionCustomerAlgorithm,omitempty" xml:"ServerSideEncryptionCustomerAlgorithm,omitempty"`
	ServerSideEncryptionKeyId             *string                `json:"ServerSideEncryptionKeyId,omitempty" xml:"ServerSideEncryptionKeyId,omitempty"`
	Size                                  *int64                 `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime                             *float64               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamCount                           *int64                 `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	Subtitles                             []*SubtitleStream      `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	Timezone                              *string                `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	Title                                 *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	TravelClusterId                       *string                `json:"TravelClusterId,omitempty" xml:"TravelClusterId,omitempty"`
	URI                                   *string                `json:"URI,omitempty" xml:"URI,omitempty"`
	UpdateTime                            *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VideoHeight                           *int64                 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoStreams                          []*VideoStream         `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	VideoWidth                            *int64                 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s File) String() string {
	return tea.Prettify(s)
}

func (s File) GoString() string {
	return s.String()
}

func (s *File) SetAccessControlAllowOrigin(v string) *File {
	s.AccessControlAllowOrigin = &v
	return s
}

func (s *File) SetAccessControlRequestMethod(v string) *File {
	s.AccessControlRequestMethod = &v
	return s
}

func (s *File) SetAddresses(v []*Address) *File {
	s.Addresses = v
	return s
}

func (s *File) SetAlbum(v string) *File {
	s.Album = &v
	return s
}

func (s *File) SetAlbumArtist(v string) *File {
	s.AlbumArtist = &v
	return s
}

func (s *File) SetArtist(v string) *File {
	s.Artist = &v
	return s
}

func (s *File) SetAudioCovers(v []*Image) *File {
	s.AudioCovers = v
	return s
}

func (s *File) SetAudioStreams(v []*AudioStream) *File {
	s.AudioStreams = v
	return s
}

func (s *File) SetBitrate(v int64) *File {
	s.Bitrate = &v
	return s
}

func (s *File) SetCacheControl(v string) *File {
	s.CacheControl = &v
	return s
}

func (s *File) SetComposer(v string) *File {
	s.Composer = &v
	return s
}

func (s *File) SetContentDisposition(v string) *File {
	s.ContentDisposition = &v
	return s
}

func (s *File) SetContentEncoding(v string) *File {
	s.ContentEncoding = &v
	return s
}

func (s *File) SetContentLanguage(v string) *File {
	s.ContentLanguage = &v
	return s
}

func (s *File) SetContentMd5(v string) *File {
	s.ContentMd5 = &v
	return s
}

func (s *File) SetContentType(v string) *File {
	s.ContentType = &v
	return s
}

func (s *File) SetCreateTime(v string) *File {
	s.CreateTime = &v
	return s
}

func (s *File) SetCroppingSuggestions(v []*CroppingSuggestion) *File {
	s.CroppingSuggestions = v
	return s
}

func (s *File) SetCustomId(v string) *File {
	s.CustomId = &v
	return s
}

func (s *File) SetCustomLabels(v map[string]interface{}) *File {
	s.CustomLabels = v
	return s
}

func (s *File) SetDatasetName(v string) *File {
	s.DatasetName = &v
	return s
}

func (s *File) SetDocumentContent(v string) *File {
	s.DocumentContent = &v
	return s
}

func (s *File) SetDocumentLanguage(v string) *File {
	s.DocumentLanguage = &v
	return s
}

func (s *File) SetDuration(v float64) *File {
	s.Duration = &v
	return s
}

func (s *File) SetETag(v string) *File {
	s.ETag = &v
	return s
}

func (s *File) SetEXIF(v string) *File {
	s.EXIF = &v
	return s
}

func (s *File) SetFigureCount(v int64) *File {
	s.FigureCount = &v
	return s
}

func (s *File) SetFigures(v []*Figure) *File {
	s.Figures = v
	return s
}

func (s *File) SetFileAccessTime(v string) *File {
	s.FileAccessTime = &v
	return s
}

func (s *File) SetFileCreateTime(v string) *File {
	s.FileCreateTime = &v
	return s
}

func (s *File) SetFileHash(v string) *File {
	s.FileHash = &v
	return s
}

func (s *File) SetFileModifiedTime(v string) *File {
	s.FileModifiedTime = &v
	return s
}

func (s *File) SetFilename(v string) *File {
	s.Filename = &v
	return s
}

func (s *File) SetFormatLongName(v string) *File {
	s.FormatLongName = &v
	return s
}

func (s *File) SetFormatName(v string) *File {
	s.FormatName = &v
	return s
}

func (s *File) SetImageHeight(v int64) *File {
	s.ImageHeight = &v
	return s
}

func (s *File) SetImageScore(v *ImageScore) *File {
	s.ImageScore = v
	return s
}

func (s *File) SetImageWidth(v int64) *File {
	s.ImageWidth = &v
	return s
}

func (s *File) SetLabels(v []*Label) *File {
	s.Labels = v
	return s
}

func (s *File) SetLanguage(v string) *File {
	s.Language = &v
	return s
}

func (s *File) SetLatLong(v string) *File {
	s.LatLong = &v
	return s
}

func (s *File) SetMediaType(v string) *File {
	s.MediaType = &v
	return s
}

func (s *File) SetOCRContents(v []*OCRContents) *File {
	s.OCRContents = v
	return s
}

func (s *File) SetOSSCRC64(v string) *File {
	s.OSSCRC64 = &v
	return s
}

func (s *File) SetOSSDeleteMarker(v string) *File {
	s.OSSDeleteMarker = &v
	return s
}

func (s *File) SetOSSExpiration(v string) *File {
	s.OSSExpiration = &v
	return s
}

func (s *File) SetOSSObjectType(v string) *File {
	s.OSSObjectType = &v
	return s
}

func (s *File) SetOSSStorageClass(v string) *File {
	s.OSSStorageClass = &v
	return s
}

func (s *File) SetOSSTagging(v map[string]interface{}) *File {
	s.OSSTagging = v
	return s
}

func (s *File) SetOSSTaggingCount(v int64) *File {
	s.OSSTaggingCount = &v
	return s
}

func (s *File) SetOSSURI(v string) *File {
	s.OSSURI = &v
	return s
}

func (s *File) SetOSSUserMeta(v map[string]interface{}) *File {
	s.OSSUserMeta = v
	return s
}

func (s *File) SetOSSVersionId(v string) *File {
	s.OSSVersionId = &v
	return s
}

func (s *File) SetObjectACL(v string) *File {
	s.ObjectACL = &v
	return s
}

func (s *File) SetObjectId(v string) *File {
	s.ObjectId = &v
	return s
}

func (s *File) SetObjectType(v string) *File {
	s.ObjectType = &v
	return s
}

func (s *File) SetOrientation(v int64) *File {
	s.Orientation = &v
	return s
}

func (s *File) SetOwnerId(v string) *File {
	s.OwnerId = &v
	return s
}

func (s *File) SetPageCount(v int64) *File {
	s.PageCount = &v
	return s
}

func (s *File) SetPerformer(v string) *File {
	s.Performer = &v
	return s
}

func (s *File) SetProduceTime(v string) *File {
	s.ProduceTime = &v
	return s
}

func (s *File) SetProgramCount(v int64) *File {
	s.ProgramCount = &v
	return s
}

func (s *File) SetProjectName(v string) *File {
	s.ProjectName = &v
	return s
}

func (s *File) SetServerSideDataEncryption(v string) *File {
	s.ServerSideDataEncryption = &v
	return s
}

func (s *File) SetServerSideEncryption(v string) *File {
	s.ServerSideEncryption = &v
	return s
}

func (s *File) SetServerSideEncryptionCustomerAlgorithm(v string) *File {
	s.ServerSideEncryptionCustomerAlgorithm = &v
	return s
}

func (s *File) SetServerSideEncryptionKeyId(v string) *File {
	s.ServerSideEncryptionKeyId = &v
	return s
}

func (s *File) SetSize(v int64) *File {
	s.Size = &v
	return s
}

func (s *File) SetStartTime(v float64) *File {
	s.StartTime = &v
	return s
}

func (s *File) SetStreamCount(v int64) *File {
	s.StreamCount = &v
	return s
}

func (s *File) SetSubtitles(v []*SubtitleStream) *File {
	s.Subtitles = v
	return s
}

func (s *File) SetTimezone(v string) *File {
	s.Timezone = &v
	return s
}

func (s *File) SetTitle(v string) *File {
	s.Title = &v
	return s
}

func (s *File) SetTravelClusterId(v string) *File {
	s.TravelClusterId = &v
	return s
}

func (s *File) SetURI(v string) *File {
	s.URI = &v
	return s
}

func (s *File) SetUpdateTime(v string) *File {
	s.UpdateTime = &v
	return s
}

func (s *File) SetVideoHeight(v int64) *File {
	s.VideoHeight = &v
	return s
}

func (s *File) SetVideoStreams(v []*VideoStream) *File {
	s.VideoStreams = v
	return s
}

func (s *File) SetVideoWidth(v int64) *File {
	s.VideoWidth = &v
	return s
}

type FileForReq struct {
	ContentType  *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Figures      []*FileForReqFigures   `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	FileHash     *string                `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	MediaType    *string                `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OSSURI       *string                `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s FileForReq) String() string {
	return tea.Prettify(s)
}

func (s FileForReq) GoString() string {
	return s.String()
}

func (s *FileForReq) SetContentType(v string) *FileForReq {
	s.ContentType = &v
	return s
}

func (s *FileForReq) SetCustomId(v string) *FileForReq {
	s.CustomId = &v
	return s
}

func (s *FileForReq) SetCustomLabels(v map[string]interface{}) *FileForReq {
	s.CustomLabels = v
	return s
}

func (s *FileForReq) SetFigures(v []*FileForReqFigures) *FileForReq {
	s.Figures = v
	return s
}

func (s *FileForReq) SetFileHash(v string) *FileForReq {
	s.FileHash = &v
	return s
}

func (s *FileForReq) SetMediaType(v string) *FileForReq {
	s.MediaType = &v
	return s
}

func (s *FileForReq) SetOSSURI(v string) *FileForReq {
	s.OSSURI = &v
	return s
}

func (s *FileForReq) SetURI(v string) *FileForReq {
	s.URI = &v
	return s
}

type FileForReqFigures struct {
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	FigureId        *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	FigureType      *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
}

func (s FileForReqFigures) String() string {
	return tea.Prettify(s)
}

func (s FileForReqFigures) GoString() string {
	return s.String()
}

func (s *FileForReqFigures) SetFigureClusterId(v string) *FileForReqFigures {
	s.FigureClusterId = &v
	return s
}

func (s *FileForReqFigures) SetFigureId(v string) *FileForReqFigures {
	s.FigureId = &v
	return s
}

func (s *FileForReqFigures) SetFigureType(v string) *FileForReqFigures {
	s.FigureType = &v
	return s
}

type HeadPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s HeadPose) String() string {
	return tea.Prettify(s)
}

func (s HeadPose) GoString() string {
	return s.String()
}

func (s *HeadPose) SetPitch(v float32) *HeadPose {
	s.Pitch = &v
	return s
}

func (s *HeadPose) SetRoll(v float32) *HeadPose {
	s.Roll = &v
	return s
}

func (s *HeadPose) SetYaw(v float32) *HeadPose {
	s.Yaw = &v
	return s
}

type Image struct {
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	EXIF                *string               `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	ImageHeight         *int64                `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageScore          *ImageScore           `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	ImageWidth          *int64                `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	OCRContents         []*OCRContents        `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
}

func (s Image) String() string {
	return tea.Prettify(s)
}

func (s Image) GoString() string {
	return s.String()
}

func (s *Image) SetCroppingSuggestions(v []*CroppingSuggestion) *Image {
	s.CroppingSuggestions = v
	return s
}

func (s *Image) SetEXIF(v string) *Image {
	s.EXIF = &v
	return s
}

func (s *Image) SetImageHeight(v int64) *Image {
	s.ImageHeight = &v
	return s
}

func (s *Image) SetImageScore(v *ImageScore) *Image {
	s.ImageScore = v
	return s
}

func (s *Image) SetImageWidth(v int64) *Image {
	s.ImageWidth = &v
	return s
}

func (s *Image) SetOCRContents(v []*OCRContents) *Image {
	s.OCRContents = v
	return s
}

type ImageScore struct {
	OverallQualityScore *float32 `json:"OverallQualityScore,omitempty" xml:"OverallQualityScore,omitempty"`
}

func (s ImageScore) String() string {
	return tea.Prettify(s)
}

func (s ImageScore) GoString() string {
	return s.String()
}

func (s *ImageScore) SetOverallQualityScore(v float32) *ImageScore {
	s.OverallQualityScore = &v
	return s
}

type KdtreeOption struct {
	CompressionLevel *int32  `json:"CompressionLevel,omitempty" xml:"CompressionLevel,omitempty"`
	LibraryName      *string `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
	QuantizationBits *int32  `json:"QuantizationBits,omitempty" xml:"QuantizationBits,omitempty"`
}

func (s KdtreeOption) String() string {
	return tea.Prettify(s)
}

func (s KdtreeOption) GoString() string {
	return s.String()
}

func (s *KdtreeOption) SetCompressionLevel(v int32) *KdtreeOption {
	s.CompressionLevel = &v
	return s
}

func (s *KdtreeOption) SetLibraryName(v string) *KdtreeOption {
	s.LibraryName = &v
	return s
}

func (s *KdtreeOption) SetQuantizationBits(v int32) *KdtreeOption {
	s.QuantizationBits = &v
	return s
}

type KeyValuePair struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s KeyValuePair) String() string {
	return tea.Prettify(s)
}

func (s KeyValuePair) GoString() string {
	return s.String()
}

func (s *KeyValuePair) SetKey(v string) *KeyValuePair {
	s.Key = &v
	return s
}

func (s *KeyValuePair) SetValue(v string) *KeyValuePair {
	s.Value = &v
	return s
}

type Label struct {
	CentricScore    *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	LabelConfidence *float32 `json:"LabelConfidence,omitempty" xml:"LabelConfidence,omitempty"`
	LabelLevel      *int64   `json:"LabelLevel,omitempty" xml:"LabelLevel,omitempty"`
	LabelName       *string  `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	Language        *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	ParentLabelName *string  `json:"ParentLabelName,omitempty" xml:"ParentLabelName,omitempty"`
}

func (s Label) String() string {
	return tea.Prettify(s)
}

func (s Label) GoString() string {
	return s.String()
}

func (s *Label) SetCentricScore(v float32) *Label {
	s.CentricScore = &v
	return s
}

func (s *Label) SetLabelConfidence(v float32) *Label {
	s.LabelConfidence = &v
	return s
}

func (s *Label) SetLabelLevel(v int64) *Label {
	s.LabelLevel = &v
	return s
}

func (s *Label) SetLabelName(v string) *Label {
	s.LabelName = &v
	return s
}

func (s *Label) SetLanguage(v string) *Label {
	s.Language = &v
	return s
}

func (s *Label) SetParentLabelName(v string) *Label {
	s.ParentLabelName = &v
	return s
}

type OCRContents struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Contents   *string   `json:"Contents,omitempty" xml:"Contents,omitempty"`
	Language   *string   `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s OCRContents) String() string {
	return tea.Prettify(s)
}

func (s OCRContents) GoString() string {
	return s.String()
}

func (s *OCRContents) SetBoundary(v *Boundary) *OCRContents {
	s.Boundary = v
	return s
}

func (s *OCRContents) SetConfidence(v float32) *OCRContents {
	s.Confidence = &v
	return s
}

func (s *OCRContents) SetContents(v string) *OCRContents {
	s.Contents = &v
	return s
}

func (s *OCRContents) SetLanguage(v string) *OCRContents {
	s.Language = &v
	return s
}

type OctreeOption struct {
	DoVoxelGridDownDownSampling *bool    `json:"DoVoxelGridDownDownSampling,omitempty" xml:"DoVoxelGridDownDownSampling,omitempty"`
	LibraryName                 *string  `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
	OctreeResolution            *float64 `json:"OctreeResolution,omitempty" xml:"OctreeResolution,omitempty"`
	PointResolution             *float64 `json:"PointResolution,omitempty" xml:"PointResolution,omitempty"`
}

func (s OctreeOption) String() string {
	return tea.Prettify(s)
}

func (s OctreeOption) GoString() string {
	return s.String()
}

func (s *OctreeOption) SetDoVoxelGridDownDownSampling(v bool) *OctreeOption {
	s.DoVoxelGridDownDownSampling = &v
	return s
}

func (s *OctreeOption) SetLibraryName(v string) *OctreeOption {
	s.LibraryName = &v
	return s
}

func (s *OctreeOption) SetOctreeResolution(v float64) *OctreeOption {
	s.OctreeResolution = &v
	return s
}

func (s *OctreeOption) SetPointResolution(v float64) *OctreeOption {
	s.PointResolution = &v
	return s
}

type PresetReference struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PresetReference) String() string {
	return tea.Prettify(s)
}

func (s PresetReference) GoString() string {
	return s.String()
}

func (s *PresetReference) SetName(v string) *PresetReference {
	s.Name = &v
	return s
}

func (s *PresetReference) SetType(v string) *PresetReference {
	s.Type = &v
	return s
}

type Project struct {
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetCount            *int64  `json:"DatasetCount,omitempty" xml:"DatasetCount,omitempty"`
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineConcurrency       *int64  `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	FileCount               *int64  `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectQueriesPerSecond *int64  `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	ServiceRole             *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	TotalFileSize           *int64  `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	UpdateTime              *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Project) String() string {
	return tea.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) SetCreateTime(v string) *Project {
	s.CreateTime = &v
	return s
}

func (s *Project) SetDatasetCount(v int64) *Project {
	s.DatasetCount = &v
	return s
}

func (s *Project) SetDatasetMaxBindCount(v int64) *Project {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *Project) SetDatasetMaxEntityCount(v int64) *Project {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *Project) SetDatasetMaxFileCount(v int64) *Project {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *Project) SetDatasetMaxRelationCount(v int64) *Project {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *Project) SetDatasetMaxTotalFileSize(v int64) *Project {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetEngineConcurrency(v int64) *Project {
	s.EngineConcurrency = &v
	return s
}

func (s *Project) SetFileCount(v int64) *Project {
	s.FileCount = &v
	return s
}

func (s *Project) SetProjectMaxDatasetCount(v int64) *Project {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *Project) SetProjectName(v string) *Project {
	s.ProjectName = &v
	return s
}

func (s *Project) SetProjectQueriesPerSecond(v int64) *Project {
	s.ProjectQueriesPerSecond = &v
	return s
}

func (s *Project) SetServiceRole(v string) *Project {
	s.ServiceRole = &v
	return s
}

func (s *Project) SetTotalFileSize(v int64) *Project {
	s.TotalFileSize = &v
	return s
}

func (s *Project) SetUpdateTime(v string) *Project {
	s.UpdateTime = &v
	return s
}

type RegionType struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegionType) String() string {
	return tea.Prettify(s)
}

func (s RegionType) GoString() string {
	return s.String()
}

func (s *RegionType) SetLocalName(v string) *RegionType {
	s.LocalName = &v
	return s
}

func (s *RegionType) SetRegionId(v string) *RegionType {
	s.RegionId = &v
	return s
}

type Row struct {
	CustomLabels []*KeyValuePair `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty" type:"Repeated"`
	URI          *string         `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s Row) String() string {
	return tea.Prettify(s)
}

func (s Row) GoString() string {
	return s.String()
}

func (s *Row) SetCustomLabels(v []*KeyValuePair) *Row {
	s.CustomLabels = v
	return s
}

func (s *Row) SetURI(v string) *Row {
	s.URI = &v
	return s
}

type SimpleQuery struct {
	Field      *string        `json:"Field,omitempty" xml:"Field,omitempty"`
	Operation  *string        `json:"Operation,omitempty" xml:"Operation,omitempty"`
	SubQueries []*SimpleQuery `json:"SubQueries,omitempty" xml:"SubQueries,omitempty" type:"Repeated"`
	Value      *string        `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQuery) String() string {
	return tea.Prettify(s)
}

func (s SimpleQuery) GoString() string {
	return s.String()
}

func (s *SimpleQuery) SetField(v string) *SimpleQuery {
	s.Field = &v
	return s
}

func (s *SimpleQuery) SetOperation(v string) *SimpleQuery {
	s.Operation = &v
	return s
}

func (s *SimpleQuery) SetSubQueries(v []*SimpleQuery) *SimpleQuery {
	s.SubQueries = v
	return s
}

func (s *SimpleQuery) SetValue(v string) *SimpleQuery {
	s.Value = &v
	return s
}

type Story struct {
	Cover            *File                  `json:"Cover,omitempty" xml:"Cover,omitempty"`
	CreateTime       *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomId         *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels     map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName      *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureClusterIds []*string              `json:"FigureClusterIds,omitempty" xml:"FigureClusterIds,omitempty" type:"Repeated"`
	Files            []*File                `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ObjectId         *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType       *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId          *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName      *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryEndTime     *string                `json:"StoryEndTime,omitempty" xml:"StoryEndTime,omitempty"`
	StoryName        *string                `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StoryStartTime   *string                `json:"StoryStartTime,omitempty" xml:"StoryStartTime,omitempty"`
	StorySubType     *string                `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType        *string                `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	UpdateTime       *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Story) String() string {
	return tea.Prettify(s)
}

func (s Story) GoString() string {
	return s.String()
}

func (s *Story) SetCover(v *File) *Story {
	s.Cover = v
	return s
}

func (s *Story) SetCreateTime(v string) *Story {
	s.CreateTime = &v
	return s
}

func (s *Story) SetCustomId(v string) *Story {
	s.CustomId = &v
	return s
}

func (s *Story) SetCustomLabels(v map[string]interface{}) *Story {
	s.CustomLabels = v
	return s
}

func (s *Story) SetDatasetName(v string) *Story {
	s.DatasetName = &v
	return s
}

func (s *Story) SetFigureClusterIds(v []*string) *Story {
	s.FigureClusterIds = v
	return s
}

func (s *Story) SetFiles(v []*File) *Story {
	s.Files = v
	return s
}

func (s *Story) SetObjectId(v string) *Story {
	s.ObjectId = &v
	return s
}

func (s *Story) SetObjectType(v string) *Story {
	s.ObjectType = &v
	return s
}

func (s *Story) SetOwnerId(v string) *Story {
	s.OwnerId = &v
	return s
}

func (s *Story) SetProjectName(v string) *Story {
	s.ProjectName = &v
	return s
}

func (s *Story) SetStoryEndTime(v string) *Story {
	s.StoryEndTime = &v
	return s
}

func (s *Story) SetStoryName(v string) *Story {
	s.StoryName = &v
	return s
}

func (s *Story) SetStoryStartTime(v string) *Story {
	s.StoryStartTime = &v
	return s
}

func (s *Story) SetStorySubType(v string) *Story {
	s.StorySubType = &v
	return s
}

func (s *Story) SetStoryType(v string) *Story {
	s.StoryType = &v
	return s
}

func (s *Story) SetUpdateTime(v string) *Story {
	s.UpdateTime = &v
	return s
}

type SubtitleStream struct {
	Bitrate        *int64   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName  *string  `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string  `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string  `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string  `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	Content        *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	Duration       *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Height         *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Index          *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
	Language       *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	StartTime      *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Width          *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubtitleStream) String() string {
	return tea.Prettify(s)
}

func (s SubtitleStream) GoString() string {
	return s.String()
}

func (s *SubtitleStream) SetBitrate(v int64) *SubtitleStream {
	s.Bitrate = &v
	return s
}

func (s *SubtitleStream) SetCodecLongName(v string) *SubtitleStream {
	s.CodecLongName = &v
	return s
}

func (s *SubtitleStream) SetCodecName(v string) *SubtitleStream {
	s.CodecName = &v
	return s
}

func (s *SubtitleStream) SetCodecTag(v string) *SubtitleStream {
	s.CodecTag = &v
	return s
}

func (s *SubtitleStream) SetCodecTagString(v string) *SubtitleStream {
	s.CodecTagString = &v
	return s
}

func (s *SubtitleStream) SetContent(v string) *SubtitleStream {
	s.Content = &v
	return s
}

func (s *SubtitleStream) SetDuration(v float64) *SubtitleStream {
	s.Duration = &v
	return s
}

func (s *SubtitleStream) SetHeight(v int64) *SubtitleStream {
	s.Height = &v
	return s
}

func (s *SubtitleStream) SetIndex(v int64) *SubtitleStream {
	s.Index = &v
	return s
}

func (s *SubtitleStream) SetLanguage(v string) *SubtitleStream {
	s.Language = &v
	return s
}

func (s *SubtitleStream) SetStartTime(v float64) *SubtitleStream {
	s.StartTime = &v
	return s
}

func (s *SubtitleStream) SetWidth(v int64) *SubtitleStream {
	s.Width = &v
	return s
}

type TaskInfo struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime   *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	StartTime *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags      map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TaskId    *string                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType  *string                `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData  *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s TaskInfo) String() string {
	return tea.Prettify(s)
}

func (s TaskInfo) GoString() string {
	return s.String()
}

func (s *TaskInfo) SetCode(v string) *TaskInfo {
	s.Code = &v
	return s
}

func (s *TaskInfo) SetEndTime(v string) *TaskInfo {
	s.EndTime = &v
	return s
}

func (s *TaskInfo) SetMessage(v string) *TaskInfo {
	s.Message = &v
	return s
}

func (s *TaskInfo) SetStartTime(v string) *TaskInfo {
	s.StartTime = &v
	return s
}

func (s *TaskInfo) SetStatus(v string) *TaskInfo {
	s.Status = &v
	return s
}

func (s *TaskInfo) SetTags(v map[string]interface{}) *TaskInfo {
	s.Tags = v
	return s
}

func (s *TaskInfo) SetTaskId(v string) *TaskInfo {
	s.TaskId = &v
	return s
}

func (s *TaskInfo) SetTaskType(v string) *TaskInfo {
	s.TaskType = &v
	return s
}

func (s *TaskInfo) SetUserData(v string) *TaskInfo {
	s.UserData = &v
	return s
}

type TimeRange struct {
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s TimeRange) String() string {
	return tea.Prettify(s)
}

func (s TimeRange) GoString() string {
	return s.String()
}

func (s *TimeRange) SetEnd(v string) *TimeRange {
	s.End = &v
	return s
}

func (s *TimeRange) SetStart(v string) *TimeRange {
	s.Start = &v
	return s
}

type TrimPolicy struct {
	DisableDeleteEmptyCell     *bool `json:"DisableDeleteEmptyCell,omitempty" xml:"DisableDeleteEmptyCell,omitempty"`
	DisableDeleteRepeatedStyle *bool `json:"DisableDeleteRepeatedStyle,omitempty" xml:"DisableDeleteRepeatedStyle,omitempty"`
	DisableDeleteUnusedPicture *bool `json:"DisableDeleteUnusedPicture,omitempty" xml:"DisableDeleteUnusedPicture,omitempty"`
	DisableDeleteUnusedShape   *bool `json:"DisableDeleteUnusedShape,omitempty" xml:"DisableDeleteUnusedShape,omitempty"`
}

func (s TrimPolicy) String() string {
	return tea.Prettify(s)
}

func (s TrimPolicy) GoString() string {
	return s.String()
}

func (s *TrimPolicy) SetDisableDeleteEmptyCell(v bool) *TrimPolicy {
	s.DisableDeleteEmptyCell = &v
	return s
}

func (s *TrimPolicy) SetDisableDeleteRepeatedStyle(v bool) *TrimPolicy {
	s.DisableDeleteRepeatedStyle = &v
	return s
}

func (s *TrimPolicy) SetDisableDeleteUnusedPicture(v bool) *TrimPolicy {
	s.DisableDeleteUnusedPicture = &v
	return s
}

func (s *TrimPolicy) SetDisableDeleteUnusedShape(v bool) *TrimPolicy {
	s.DisableDeleteUnusedShape = &v
	return s
}

type VideoStream struct {
	AverageFrameRate   *string  `json:"AverageFrameRate,omitempty" xml:"AverageFrameRate,omitempty"`
	BitDepth           *int64   `json:"BitDepth,omitempty" xml:"BitDepth,omitempty"`
	Bitrate            *int64   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName      *string  `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName          *string  `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag           *string  `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString     *string  `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase      *string  `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	ColorPrimaries     *string  `json:"ColorPrimaries,omitempty" xml:"ColorPrimaries,omitempty"`
	ColorRange         *string  `json:"ColorRange,omitempty" xml:"ColorRange,omitempty"`
	ColorSpace         *string  `json:"ColorSpace,omitempty" xml:"ColorSpace,omitempty"`
	ColorTransfer      *string  `json:"ColorTransfer,omitempty" xml:"ColorTransfer,omitempty"`
	DisplayAspectRatio *string  `json:"DisplayAspectRatio,omitempty" xml:"DisplayAspectRatio,omitempty"`
	Duration           *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FrameCount         *int64   `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	FrameRate          *string  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	HasBFrames         *int64   `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height             *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Index              *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
	Language           *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	Level              *int64   `json:"Level,omitempty" xml:"Level,omitempty"`
	PixelFormat        *string  `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	Profile            *string  `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Rotate             *string  `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SampleAspectRatio  *string  `json:"SampleAspectRatio,omitempty" xml:"SampleAspectRatio,omitempty"`
	StartTime          *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeBase           *string  `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	Width              *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s VideoStream) String() string {
	return tea.Prettify(s)
}

func (s VideoStream) GoString() string {
	return s.String()
}

func (s *VideoStream) SetAverageFrameRate(v string) *VideoStream {
	s.AverageFrameRate = &v
	return s
}

func (s *VideoStream) SetBitDepth(v int64) *VideoStream {
	s.BitDepth = &v
	return s
}

func (s *VideoStream) SetBitrate(v int64) *VideoStream {
	s.Bitrate = &v
	return s
}

func (s *VideoStream) SetCodecLongName(v string) *VideoStream {
	s.CodecLongName = &v
	return s
}

func (s *VideoStream) SetCodecName(v string) *VideoStream {
	s.CodecName = &v
	return s
}

func (s *VideoStream) SetCodecTag(v string) *VideoStream {
	s.CodecTag = &v
	return s
}

func (s *VideoStream) SetCodecTagString(v string) *VideoStream {
	s.CodecTagString = &v
	return s
}

func (s *VideoStream) SetCodecTimeBase(v string) *VideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *VideoStream) SetColorPrimaries(v string) *VideoStream {
	s.ColorPrimaries = &v
	return s
}

func (s *VideoStream) SetColorRange(v string) *VideoStream {
	s.ColorRange = &v
	return s
}

func (s *VideoStream) SetColorSpace(v string) *VideoStream {
	s.ColorSpace = &v
	return s
}

func (s *VideoStream) SetColorTransfer(v string) *VideoStream {
	s.ColorTransfer = &v
	return s
}

func (s *VideoStream) SetDisplayAspectRatio(v string) *VideoStream {
	s.DisplayAspectRatio = &v
	return s
}

func (s *VideoStream) SetDuration(v float64) *VideoStream {
	s.Duration = &v
	return s
}

func (s *VideoStream) SetFrameCount(v int64) *VideoStream {
	s.FrameCount = &v
	return s
}

func (s *VideoStream) SetFrameRate(v string) *VideoStream {
	s.FrameRate = &v
	return s
}

func (s *VideoStream) SetHasBFrames(v int64) *VideoStream {
	s.HasBFrames = &v
	return s
}

func (s *VideoStream) SetHeight(v int64) *VideoStream {
	s.Height = &v
	return s
}

func (s *VideoStream) SetIndex(v int64) *VideoStream {
	s.Index = &v
	return s
}

func (s *VideoStream) SetLanguage(v string) *VideoStream {
	s.Language = &v
	return s
}

func (s *VideoStream) SetLevel(v int64) *VideoStream {
	s.Level = &v
	return s
}

func (s *VideoStream) SetPixelFormat(v string) *VideoStream {
	s.PixelFormat = &v
	return s
}

func (s *VideoStream) SetProfile(v string) *VideoStream {
	s.Profile = &v
	return s
}

func (s *VideoStream) SetRotate(v string) *VideoStream {
	s.Rotate = &v
	return s
}

func (s *VideoStream) SetSampleAspectRatio(v string) *VideoStream {
	s.SampleAspectRatio = &v
	return s
}

func (s *VideoStream) SetStartTime(v float64) *VideoStream {
	s.StartTime = &v
	return s
}

func (s *VideoStream) SetTimeBase(v string) *VideoStream {
	s.TimeBase = &v
	return s
}

func (s *VideoStream) SetWidth(v int64) *VideoStream {
	s.Width = &v
	return s
}

type WebofficePermission struct {
	Copy     *bool `json:"Copy,omitempty" xml:"Copy,omitempty"`
	Export   *bool `json:"Export,omitempty" xml:"Export,omitempty"`
	History  *bool `json:"History,omitempty" xml:"History,omitempty"`
	Print    *bool `json:"Print,omitempty" xml:"Print,omitempty"`
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	Rename   *bool `json:"Rename,omitempty" xml:"Rename,omitempty"`
}

func (s WebofficePermission) String() string {
	return tea.Prettify(s)
}

func (s WebofficePermission) GoString() string {
	return s.String()
}

func (s *WebofficePermission) SetCopy(v bool) *WebofficePermission {
	s.Copy = &v
	return s
}

func (s *WebofficePermission) SetExport(v bool) *WebofficePermission {
	s.Export = &v
	return s
}

func (s *WebofficePermission) SetHistory(v bool) *WebofficePermission {
	s.History = &v
	return s
}

func (s *WebofficePermission) SetPrint(v bool) *WebofficePermission {
	s.Print = &v
	return s
}

func (s *WebofficePermission) SetReadonly(v bool) *WebofficePermission {
	s.Readonly = &v
	return s
}

func (s *WebofficePermission) SetRename(v bool) *WebofficePermission {
	s.Rename = &v
	return s
}

type WebofficeUser struct {
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s WebofficeUser) String() string {
	return tea.Prettify(s)
}

func (s WebofficeUser) GoString() string {
	return s.String()
}

func (s *WebofficeUser) SetAvatar(v string) *WebofficeUser {
	s.Avatar = &v
	return s
}

func (s *WebofficeUser) SetId(v string) *WebofficeUser {
	s.Id = &v
	return s
}

func (s *WebofficeUser) SetName(v string) *WebofficeUser {
	s.Name = &v
	return s
}

type WebofficeWatermark struct {
	FillStyle  *string  `json:"FillStyle,omitempty" xml:"FillStyle,omitempty"`
	Font       *string  `json:"Font,omitempty" xml:"Font,omitempty"`
	Horizontal *int64   `json:"Horizontal,omitempty" xml:"Horizontal,omitempty"`
	Rotate     *float32 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	Type       *int64   `json:"Type,omitempty" xml:"Type,omitempty"`
	Value      *string  `json:"Value,omitempty" xml:"Value,omitempty"`
	Vertical   *int64   `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
}

func (s WebofficeWatermark) String() string {
	return tea.Prettify(s)
}

func (s WebofficeWatermark) GoString() string {
	return s.String()
}

func (s *WebofficeWatermark) SetFillStyle(v string) *WebofficeWatermark {
	s.FillStyle = &v
	return s
}

func (s *WebofficeWatermark) SetFont(v string) *WebofficeWatermark {
	s.Font = &v
	return s
}

func (s *WebofficeWatermark) SetHorizontal(v int64) *WebofficeWatermark {
	s.Horizontal = &v
	return s
}

func (s *WebofficeWatermark) SetRotate(v float32) *WebofficeWatermark {
	s.Rotate = &v
	return s
}

func (s *WebofficeWatermark) SetType(v int64) *WebofficeWatermark {
	s.Type = &v
	return s
}

func (s *WebofficeWatermark) SetValue(v string) *WebofficeWatermark {
	s.Value = &v
	return s
}

func (s *WebofficeWatermark) SetVertical(v int64) *WebofficeWatermark {
	s.Vertical = &v
	return s
}

type AddStoryFilesRequest struct {
	DatasetName *string                      `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files       []*AddStoryFilesRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ObjectId    *string                      `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string                      `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s AddStoryFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequest) SetDatasetName(v string) *AddStoryFilesRequest {
	s.DatasetName = &v
	return s
}

func (s *AddStoryFilesRequest) SetFiles(v []*AddStoryFilesRequestFiles) *AddStoryFilesRequest {
	s.Files = v
	return s
}

func (s *AddStoryFilesRequest) SetObjectId(v string) *AddStoryFilesRequest {
	s.ObjectId = &v
	return s
}

func (s *AddStoryFilesRequest) SetProjectName(v string) *AddStoryFilesRequest {
	s.ProjectName = &v
	return s
}

type AddStoryFilesRequestFiles struct {
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s AddStoryFilesRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequestFiles) SetURI(v string) *AddStoryFilesRequestFiles {
	s.URI = &v
	return s
}

type AddStoryFilesShrinkRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s AddStoryFilesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddStoryFilesShrinkRequest) SetDatasetName(v string) *AddStoryFilesShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) SetFilesShrink(v string) *AddStoryFilesShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) SetObjectId(v string) *AddStoryFilesShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *AddStoryFilesShrinkRequest) SetProjectName(v string) *AddStoryFilesShrinkRequest {
	s.ProjectName = &v
	return s
}

type AddStoryFilesResponseBody struct {
	Files     []*AddStoryFilesResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddStoryFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponseBody) SetFiles(v []*AddStoryFilesResponseBodyFiles) *AddStoryFilesResponseBody {
	s.Files = v
	return s
}

func (s *AddStoryFilesResponseBody) SetRequestId(v string) *AddStoryFilesResponseBody {
	s.RequestId = &v
	return s
}

type AddStoryFilesResponseBodyFiles struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	URI          *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s AddStoryFilesResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponseBodyFiles) SetErrorCode(v string) *AddStoryFilesResponseBodyFiles {
	s.ErrorCode = &v
	return s
}

func (s *AddStoryFilesResponseBodyFiles) SetErrorMessage(v string) *AddStoryFilesResponseBodyFiles {
	s.ErrorMessage = &v
	return s
}

func (s *AddStoryFilesResponseBodyFiles) SetURI(v string) *AddStoryFilesResponseBodyFiles {
	s.URI = &v
	return s
}

type AddStoryFilesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddStoryFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddStoryFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesResponse) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponse) SetHeaders(v map[string]*string) *AddStoryFilesResponse {
	s.Headers = v
	return s
}

func (s *AddStoryFilesResponse) SetStatusCode(v int32) *AddStoryFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddStoryFilesResponse) SetBody(v *AddStoryFilesResponseBody) *AddStoryFilesResponse {
	s.Body = v
	return s
}

type AttachOSSBucketRequest struct {
	OSSBucket   *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s AttachOSSBucketRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachOSSBucketRequest) GoString() string {
	return s.String()
}

func (s *AttachOSSBucketRequest) SetOSSBucket(v string) *AttachOSSBucketRequest {
	s.OSSBucket = &v
	return s
}

func (s *AttachOSSBucketRequest) SetProjectName(v string) *AttachOSSBucketRequest {
	s.ProjectName = &v
	return s
}

type AttachOSSBucketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachOSSBucketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachOSSBucketResponseBody) GoString() string {
	return s.String()
}

func (s *AttachOSSBucketResponseBody) SetRequestId(v string) *AttachOSSBucketResponseBody {
	s.RequestId = &v
	return s
}

type AttachOSSBucketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachOSSBucketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachOSSBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachOSSBucketResponse) GoString() string {
	return s.String()
}

func (s *AttachOSSBucketResponse) SetHeaders(v map[string]*string) *AttachOSSBucketResponse {
	s.Headers = v
	return s
}

func (s *AttachOSSBucketResponse) SetStatusCode(v int32) *AttachOSSBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachOSSBucketResponse) SetBody(v *AttachOSSBucketResponseBody) *AttachOSSBucketResponse {
	s.Body = v
	return s
}

type BatchDeleteFileMetaRequest struct {
	DatasetName *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URIs        []*string `json:"URIs,omitempty" xml:"URIs,omitempty" type:"Repeated"`
}

func (s BatchDeleteFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaRequest) SetDatasetName(v string) *BatchDeleteFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchDeleteFileMetaRequest) SetProjectName(v string) *BatchDeleteFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchDeleteFileMetaRequest) SetURIs(v []*string) *BatchDeleteFileMetaRequest {
	s.URIs = v
	return s
}

type BatchDeleteFileMetaShrinkRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URIsShrink  *string `json:"URIs,omitempty" xml:"URIs,omitempty"`
}

func (s BatchDeleteFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaShrinkRequest) SetDatasetName(v string) *BatchDeleteFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchDeleteFileMetaShrinkRequest) SetProjectName(v string) *BatchDeleteFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchDeleteFileMetaShrinkRequest) SetURIsShrink(v string) *BatchDeleteFileMetaShrinkRequest {
	s.URIsShrink = &v
	return s
}

type BatchDeleteFileMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaResponseBody) SetRequestId(v string) *BatchDeleteFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteFileMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaResponse) SetHeaders(v map[string]*string) *BatchDeleteFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteFileMetaResponse) SetStatusCode(v int32) *BatchDeleteFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteFileMetaResponse) SetBody(v *BatchDeleteFileMetaResponseBody) *BatchDeleteFileMetaResponse {
	s.Body = v
	return s
}

type BatchGetFileMetaRequest struct {
	DatasetName *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URIs        []*string `json:"URIs,omitempty" xml:"URIs,omitempty" type:"Repeated"`
}

func (s BatchGetFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaRequest) SetDatasetName(v string) *BatchGetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFileMetaRequest) SetProjectName(v string) *BatchGetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFileMetaRequest) SetURIs(v []*string) *BatchGetFileMetaRequest {
	s.URIs = v
	return s
}

type BatchGetFileMetaShrinkRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URIsShrink  *string `json:"URIs,omitempty" xml:"URIs,omitempty"`
}

func (s BatchGetFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaShrinkRequest) SetDatasetName(v string) *BatchGetFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetProjectName(v string) *BatchGetFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetURIsShrink(v string) *BatchGetFileMetaShrinkRequest {
	s.URIsShrink = &v
	return s
}

type BatchGetFileMetaResponseBody struct {
	Files     []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaResponseBody) SetFiles(v []*File) *BatchGetFileMetaResponseBody {
	s.Files = v
	return s
}

func (s *BatchGetFileMetaResponseBody) SetRequestId(v string) *BatchGetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type BatchGetFileMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaResponse) SetHeaders(v map[string]*string) *BatchGetFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchGetFileMetaResponse) SetStatusCode(v int32) *BatchGetFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetFileMetaResponse) SetBody(v *BatchGetFileMetaResponseBody) *BatchGetFileMetaResponse {
	s.Body = v
	return s
}

type BatchIndexFileMetaRequest struct {
	DatasetName     *string       `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files           []*FileForReq `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	NotifyTopicName *string       `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchIndexFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaRequest) SetDatasetName(v string) *BatchIndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetFiles(v []*FileForReq) *BatchIndexFileMetaRequest {
	s.Files = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetNotifyTopicName(v string) *BatchIndexFileMetaRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetProjectName(v string) *BatchIndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

type BatchIndexFileMetaShrinkRequest struct {
	DatasetName     *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink     *string `json:"Files,omitempty" xml:"Files,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchIndexFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaShrinkRequest) SetDatasetName(v string) *BatchIndexFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetFilesShrink(v string) *BatchIndexFileMetaShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetNotifyTopicName(v string) *BatchIndexFileMetaShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetProjectName(v string) *BatchIndexFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

type BatchIndexFileMetaResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchIndexFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaResponseBody) SetEventId(v string) *BatchIndexFileMetaResponseBody {
	s.EventId = &v
	return s
}

func (s *BatchIndexFileMetaResponseBody) SetRequestId(v string) *BatchIndexFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type BatchIndexFileMetaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchIndexFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchIndexFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaResponse) SetHeaders(v map[string]*string) *BatchIndexFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchIndexFileMetaResponse) SetStatusCode(v int32) *BatchIndexFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchIndexFileMetaResponse) SetBody(v *BatchIndexFileMetaResponseBody) *BatchIndexFileMetaResponse {
	s.Body = v
	return s
}

type BatchUpdateFileMetaRequest struct {
	DatasetName *string       `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files       []*FileForReq `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ProjectName *string       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchUpdateFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaRequest) SetDatasetName(v string) *BatchUpdateFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchUpdateFileMetaRequest) SetFiles(v []*FileForReq) *BatchUpdateFileMetaRequest {
	s.Files = v
	return s
}

func (s *BatchUpdateFileMetaRequest) SetProjectName(v string) *BatchUpdateFileMetaRequest {
	s.ProjectName = &v
	return s
}

type BatchUpdateFileMetaShrinkRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchUpdateFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaShrinkRequest) SetDatasetName(v string) *BatchUpdateFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) SetFilesShrink(v string) *BatchUpdateFileMetaShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) SetProjectName(v string) *BatchUpdateFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

type BatchUpdateFileMetaResponseBody struct {
	Files     []*BatchUpdateFileMetaResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponseBody) SetFiles(v []*BatchUpdateFileMetaResponseBodyFiles) *BatchUpdateFileMetaResponseBody {
	s.Files = v
	return s
}

func (s *BatchUpdateFileMetaResponseBody) SetRequestId(v string) *BatchUpdateFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type BatchUpdateFileMetaResponseBodyFiles struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	URI     *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s BatchUpdateFileMetaResponseBodyFiles) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponseBodyFiles) SetMessage(v string) *BatchUpdateFileMetaResponseBodyFiles {
	s.Message = &v
	return s
}

func (s *BatchUpdateFileMetaResponseBodyFiles) SetSuccess(v bool) *BatchUpdateFileMetaResponseBodyFiles {
	s.Success = &v
	return s
}

func (s *BatchUpdateFileMetaResponseBodyFiles) SetURI(v string) *BatchUpdateFileMetaResponseBodyFiles {
	s.URI = &v
	return s
}

type BatchUpdateFileMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponse) SetHeaders(v map[string]*string) *BatchUpdateFileMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFileMetaResponse) SetStatusCode(v int32) *BatchUpdateFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateFileMetaResponse) SetBody(v *BatchUpdateFileMetaResponseBody) *BatchUpdateFileMetaResponse {
	s.Body = v
	return s
}

type CreateBindingRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) SetDatasetName(v string) *CreateBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateBindingRequest) SetProjectName(v string) *CreateBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBindingRequest) SetURI(v string) *CreateBindingRequest {
	s.URI = &v
	return s
}

type CreateBindingResponseBody struct {
	Binding   *Binding `json:"Binding,omitempty" xml:"Binding,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBindingResponseBody) SetBinding(v *Binding) *CreateBindingResponseBody {
	s.Binding = v
	return s
}

func (s *CreateBindingResponseBody) SetRequestId(v string) *CreateBindingResponseBody {
	s.RequestId = &v
	return s
}

type CreateBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateBindingResponse) SetHeaders(v map[string]*string) *CreateBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateBindingResponse) SetStatusCode(v int32) *CreateBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBindingResponse) SetBody(v *CreateBindingResponseBody) *CreateBindingResponse {
	s.Body = v
	return s
}

type CreateCompressPointCloudTaskRequest struct {
	CompressMethod       *string                `json:"CompressMethod,omitempty" xml:"CompressMethod,omitempty"`
	CredentialConfig     *CredentialConfig      `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	KdtreeOption         *KdtreeOption          `json:"KdtreeOption,omitempty" xml:"KdtreeOption,omitempty"`
	NotifyTopicName      *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	OctreeOption         *OctreeOption          `json:"OctreeOption,omitempty" xml:"OctreeOption,omitempty"`
	PointCloudFields     []*string              `json:"PointCloudFields,omitempty" xml:"PointCloudFields,omitempty" type:"Repeated"`
	PointCloudFileFormat *string                `json:"PointCloudFileFormat,omitempty" xml:"PointCloudFileFormat,omitempty"`
	ProjectName          *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI            *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags                 map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI            *string                `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData             *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateCompressPointCloudTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCompressPointCloudTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskRequest) SetCompressMethod(v string) *CreateCompressPointCloudTaskRequest {
	s.CompressMethod = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateCompressPointCloudTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetKdtreeOption(v *KdtreeOption) *CreateCompressPointCloudTaskRequest {
	s.KdtreeOption = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetNotifyTopicName(v string) *CreateCompressPointCloudTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetOctreeOption(v *OctreeOption) *CreateCompressPointCloudTaskRequest {
	s.OctreeOption = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetPointCloudFields(v []*string) *CreateCompressPointCloudTaskRequest {
	s.PointCloudFields = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetPointCloudFileFormat(v string) *CreateCompressPointCloudTaskRequest {
	s.PointCloudFileFormat = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetProjectName(v string) *CreateCompressPointCloudTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetSourceURI(v string) *CreateCompressPointCloudTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetTags(v map[string]interface{}) *CreateCompressPointCloudTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetTargetURI(v string) *CreateCompressPointCloudTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetUserData(v string) *CreateCompressPointCloudTaskRequest {
	s.UserData = &v
	return s
}

type CreateCompressPointCloudTaskShrinkRequest struct {
	CompressMethod         *string `json:"CompressMethod,omitempty" xml:"CompressMethod,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	KdtreeOptionShrink     *string `json:"KdtreeOption,omitempty" xml:"KdtreeOption,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	OctreeOptionShrink     *string `json:"OctreeOption,omitempty" xml:"OctreeOption,omitempty"`
	PointCloudFieldsShrink *string `json:"PointCloudFields,omitempty" xml:"PointCloudFields,omitempty"`
	PointCloudFileFormat   *string `json:"PointCloudFileFormat,omitempty" xml:"PointCloudFileFormat,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI              *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateCompressPointCloudTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCompressPointCloudTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetCompressMethod(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.CompressMethod = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetKdtreeOptionShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.KdtreeOptionShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetNotifyTopicName(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetOctreeOptionShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.OctreeOptionShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetPointCloudFieldsShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.PointCloudFieldsShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetPointCloudFileFormat(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.PointCloudFileFormat = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetProjectName(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetSourceURI(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetTagsShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetTargetURI(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetUserData(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateCompressPointCloudTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateCompressPointCloudTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCompressPointCloudTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskResponseBody) SetEventId(v string) *CreateCompressPointCloudTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponseBody) SetRequestId(v string) *CreateCompressPointCloudTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponseBody) SetTaskId(v string) *CreateCompressPointCloudTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateCompressPointCloudTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCompressPointCloudTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCompressPointCloudTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCompressPointCloudTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskResponse) SetHeaders(v map[string]*string) *CreateCompressPointCloudTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCompressPointCloudTaskResponse) SetStatusCode(v int32) *CreateCompressPointCloudTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponse) SetBody(v *CreateCompressPointCloudTaskResponseBody) *CreateCompressPointCloudTaskResponse {
	s.Body = v
	return s
}

type CreateDatasetRequest struct {
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) SetDatasetMaxBindCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxEntityCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxFileCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxRelationCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxTotalFileSize(v int64) *CreateDatasetRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetName(v string) *CreateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetProjectName(v string) *CreateDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDatasetRequest) SetTemplateId(v string) *CreateDatasetRequest {
	s.TemplateId = &v
	return s
}

type CreateDatasetResponseBody struct {
	Dataset   *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) SetDataset(v *Dataset) *CreateDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponse) SetHeaders(v map[string]*string) *CreateDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetResponse) SetStatusCode(v int32) *CreateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateDetectVideoLabelsTaskRequest struct {
	CredentialConfig *CredentialConfig      `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotifyTopicName  *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName      *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData         *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateDetectVideoLabelsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectVideoLabelsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectVideoLabelsTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateDetectVideoLabelsTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetNotifyTopicName(v string) *CreateDetectVideoLabelsTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetProjectName(v string) *CreateDetectVideoLabelsTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetSourceURI(v string) *CreateDetectVideoLabelsTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetTags(v map[string]interface{}) *CreateDetectVideoLabelsTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetUserData(v string) *CreateDetectVideoLabelsTaskRequest {
	s.UserData = &v
	return s
}

type CreateDetectVideoLabelsTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateDetectVideoLabelsTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectVideoLabelsTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectVideoLabelsTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateDetectVideoLabelsTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskShrinkRequest) SetNotifyTopicName(v string) *CreateDetectVideoLabelsTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskShrinkRequest) SetProjectName(v string) *CreateDetectVideoLabelsTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskShrinkRequest) SetSourceURI(v string) *CreateDetectVideoLabelsTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskShrinkRequest) SetTagsShrink(v string) *CreateDetectVideoLabelsTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskShrinkRequest) SetUserData(v string) *CreateDetectVideoLabelsTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateDetectVideoLabelsTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDetectVideoLabelsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectVideoLabelsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetEventId(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetRequestId(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetTaskId(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDetectVideoLabelsTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDetectVideoLabelsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDetectVideoLabelsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectVideoLabelsTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDetectVideoLabelsTaskResponse) SetHeaders(v map[string]*string) *CreateDetectVideoLabelsTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponse) SetStatusCode(v int32) *CreateDetectVideoLabelsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponse) SetBody(v *CreateDetectVideoLabelsTaskResponseBody) *CreateDetectVideoLabelsTaskResponse {
	s.Body = v
	return s
}

type CreateFigureClusteringTaskRequest struct {
	DatasetName     *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	NotifyTopicName *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData        *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClusteringTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClusteringTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskRequest) SetDatasetName(v string) *CreateFigureClusteringTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetNotifyTopicName(v string) *CreateFigureClusteringTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetProjectName(v string) *CreateFigureClusteringTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetTags(v map[string]interface{}) *CreateFigureClusteringTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetUserData(v string) *CreateFigureClusteringTaskRequest {
	s.UserData = &v
	return s
}

type CreateFigureClusteringTaskShrinkRequest struct {
	DatasetName     *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClusteringTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetNotifyTopicName(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetProjectName(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetUserData(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateFigureClusteringTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFigureClusteringTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClusteringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskResponseBody) SetEventId(v string) *CreateFigureClusteringTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFigureClusteringTaskResponseBody) SetRequestId(v string) *CreateFigureClusteringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFigureClusteringTaskResponseBody) SetTaskId(v string) *CreateFigureClusteringTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateFigureClusteringTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFigureClusteringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFigureClusteringTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClusteringTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskResponse) SetHeaders(v map[string]*string) *CreateFigureClusteringTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFigureClusteringTaskResponse) SetStatusCode(v int32) *CreateFigureClusteringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFigureClusteringTaskResponse) SetBody(v *CreateFigureClusteringTaskResponseBody) *CreateFigureClusteringTaskResponse {
	s.Body = v
	return s
}

type CreateFigureClustersMergingTaskRequest struct {
	DatasetName     *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	From            *string                `json:"From,omitempty" xml:"From,omitempty"`
	NotifyTopicName *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	To              *string                `json:"To,omitempty" xml:"To,omitempty"`
	UserData        *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClustersMergingTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClustersMergingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskRequest) SetDatasetName(v string) *CreateFigureClustersMergingTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetFrom(v string) *CreateFigureClustersMergingTaskRequest {
	s.From = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetNotifyTopicName(v string) *CreateFigureClustersMergingTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetProjectName(v string) *CreateFigureClustersMergingTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetTags(v map[string]interface{}) *CreateFigureClustersMergingTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetTo(v string) *CreateFigureClustersMergingTaskRequest {
	s.To = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetUserData(v string) *CreateFigureClustersMergingTaskRequest {
	s.UserData = &v
	return s
}

type CreateFigureClustersMergingTaskShrinkRequest struct {
	DatasetName     *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	From            *string `json:"From,omitempty" xml:"From,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	To              *string `json:"To,omitempty" xml:"To,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClustersMergingTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClustersMergingTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetDatasetName(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetFrom(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.From = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetNotifyTopicName(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetProjectName(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetTagsShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetTo(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.To = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetUserData(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateFigureClustersMergingTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFigureClustersMergingTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClustersMergingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskResponseBody) SetEventId(v string) *CreateFigureClustersMergingTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponseBody) SetRequestId(v string) *CreateFigureClustersMergingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponseBody) SetTaskId(v string) *CreateFigureClustersMergingTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateFigureClustersMergingTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFigureClustersMergingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFigureClustersMergingTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFigureClustersMergingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskResponse) SetHeaders(v map[string]*string) *CreateFigureClustersMergingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFigureClustersMergingTaskResponse) SetStatusCode(v int32) *CreateFigureClustersMergingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponse) SetBody(v *CreateFigureClustersMergingTaskResponseBody) *CreateFigureClustersMergingTaskResponse {
	s.Body = v
	return s
}

type CreateImageModerationTaskRequest struct {
	CredentialConfig *CredentialConfig      `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Interval         *int64                 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames        *int64                 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	NotifyTopicName  *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName      *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer         *string                `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	Scenes           []*string              `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	SourceURI        *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData         *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageModerationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageModerationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateImageModerationTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetInterval(v int64) *CreateImageModerationTaskRequest {
	s.Interval = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetMaxFrames(v int64) *CreateImageModerationTaskRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetNotifyTopicName(v string) *CreateImageModerationTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetProjectName(v string) *CreateImageModerationTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetReviewer(v string) *CreateImageModerationTaskRequest {
	s.Reviewer = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetScenes(v []*string) *CreateImageModerationTaskRequest {
	s.Scenes = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetSourceURI(v string) *CreateImageModerationTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetTags(v map[string]interface{}) *CreateImageModerationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetUserData(v string) *CreateImageModerationTaskRequest {
	s.UserData = &v
	return s
}

type CreateImageModerationTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Interval               *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames              *int64  `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer               *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	ScenesShrink           *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageModerationTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageModerationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetInterval(v int64) *CreateImageModerationTaskShrinkRequest {
	s.Interval = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetMaxFrames(v int64) *CreateImageModerationTaskShrinkRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetNotifyTopicName(v string) *CreateImageModerationTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetProjectName(v string) *CreateImageModerationTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetReviewer(v string) *CreateImageModerationTaskShrinkRequest {
	s.Reviewer = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetScenesShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.ScenesShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetSourceURI(v string) *CreateImageModerationTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetTagsShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetUserData(v string) *CreateImageModerationTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateImageModerationTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageModerationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageModerationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskResponseBody) SetEventId(v string) *CreateImageModerationTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateImageModerationTaskResponseBody) SetRequestId(v string) *CreateImageModerationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageModerationTaskResponseBody) SetTaskId(v string) *CreateImageModerationTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateImageModerationTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageModerationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageModerationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageModerationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskResponse) SetHeaders(v map[string]*string) *CreateImageModerationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageModerationTaskResponse) SetStatusCode(v int32) *CreateImageModerationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageModerationTaskResponse) SetBody(v *CreateImageModerationTaskResponseBody) *CreateImageModerationTaskResponse {
	s.Body = v
	return s
}

type CreateImageSplicingTaskRequest struct {
	Align            *int64                                   `json:"Align,omitempty" xml:"Align,omitempty"`
	BackgroundColor  *string                                  `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	CredentialConfig *CredentialConfig                        `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Direction        *string                                  `json:"Direction,omitempty" xml:"Direction,omitempty"`
	ImageFormat      *string                                  `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	Margin           *int64                                   `json:"Margin,omitempty" xml:"Margin,omitempty"`
	NotifyTopicName  *string                                  `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Padding          *int64                                   `json:"Padding,omitempty" xml:"Padding,omitempty"`
	ProjectName      *string                                  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality          *int64                                   `json:"Quality,omitempty" xml:"Quality,omitempty"`
	ScaleType        *string                                  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	Sources          []*CreateImageSplicingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	Tags             map[string]interface{}                   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI        *string                                  `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData         *string                                  `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageSplicingTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageSplicingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskRequest) SetAlign(v int64) *CreateImageSplicingTaskRequest {
	s.Align = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetBackgroundColor(v string) *CreateImageSplicingTaskRequest {
	s.BackgroundColor = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateImageSplicingTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetDirection(v string) *CreateImageSplicingTaskRequest {
	s.Direction = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetImageFormat(v string) *CreateImageSplicingTaskRequest {
	s.ImageFormat = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetMargin(v int64) *CreateImageSplicingTaskRequest {
	s.Margin = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetNotifyTopicName(v string) *CreateImageSplicingTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetPadding(v int64) *CreateImageSplicingTaskRequest {
	s.Padding = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetProjectName(v string) *CreateImageSplicingTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetQuality(v int64) *CreateImageSplicingTaskRequest {
	s.Quality = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetScaleType(v string) *CreateImageSplicingTaskRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetSources(v []*CreateImageSplicingTaskRequestSources) *CreateImageSplicingTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetTags(v map[string]interface{}) *CreateImageSplicingTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetTargetURI(v string) *CreateImageSplicingTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetUserData(v string) *CreateImageSplicingTaskRequest {
	s.UserData = &v
	return s
}

type CreateImageSplicingTaskRequestSources struct {
	Rotate *int64  `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	URI    *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateImageSplicingTaskRequestSources) String() string {
	return tea.Prettify(s)
}

func (s CreateImageSplicingTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskRequestSources) SetRotate(v int64) *CreateImageSplicingTaskRequestSources {
	s.Rotate = &v
	return s
}

func (s *CreateImageSplicingTaskRequestSources) SetURI(v string) *CreateImageSplicingTaskRequestSources {
	s.URI = &v
	return s
}

type CreateImageSplicingTaskShrinkRequest struct {
	Align                  *int64  `json:"Align,omitempty" xml:"Align,omitempty"`
	BackgroundColor        *string `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Direction              *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	ImageFormat            *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	Margin                 *int64  `json:"Margin,omitempty" xml:"Margin,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Padding                *int64  `json:"Padding,omitempty" xml:"Padding,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality                *int64  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	ScaleType              *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SourcesShrink          *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI              *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageSplicingTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageSplicingTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskShrinkRequest) SetAlign(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Align = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetBackgroundColor(v string) *CreateImageSplicingTaskShrinkRequest {
	s.BackgroundColor = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetDirection(v string) *CreateImageSplicingTaskShrinkRequest {
	s.Direction = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetImageFormat(v string) *CreateImageSplicingTaskShrinkRequest {
	s.ImageFormat = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetMargin(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Margin = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetNotifyTopicName(v string) *CreateImageSplicingTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetPadding(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Padding = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetProjectName(v string) *CreateImageSplicingTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetQuality(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Quality = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetScaleType(v string) *CreateImageSplicingTaskShrinkRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetSourcesShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetTagsShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetTargetURI(v string) *CreateImageSplicingTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetUserData(v string) *CreateImageSplicingTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateImageSplicingTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageSplicingTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageSplicingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskResponseBody) SetEventId(v string) *CreateImageSplicingTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateImageSplicingTaskResponseBody) SetRequestId(v string) *CreateImageSplicingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageSplicingTaskResponseBody) SetTaskId(v string) *CreateImageSplicingTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateImageSplicingTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageSplicingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageSplicingTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageSplicingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskResponse) SetHeaders(v map[string]*string) *CreateImageSplicingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageSplicingTaskResponse) SetStatusCode(v int32) *CreateImageSplicingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageSplicingTaskResponse) SetBody(v *CreateImageSplicingTaskResponseBody) *CreateImageSplicingTaskResponse {
	s.Body = v
	return s
}

type CreateMediaConvertTaskRequest struct {
	CredentialConfig *CredentialConfig                       `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotifyTopicName  *string                                 `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName      *string                                 `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sources          []*CreateMediaConvertTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	Tags             map[string]interface{}                  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Targets          []*CreateMediaConvertTaskRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	UserData         *string                                 `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateMediaConvertTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateMediaConvertTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetNotifyTopicName(v string) *CreateMediaConvertTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetProjectName(v string) *CreateMediaConvertTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetSources(v []*CreateMediaConvertTaskRequestSources) *CreateMediaConvertTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetTags(v map[string]interface{}) *CreateMediaConvertTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetTargets(v []*CreateMediaConvertTaskRequestTargets) *CreateMediaConvertTaskRequest {
	s.Targets = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetUserData(v string) *CreateMediaConvertTaskRequest {
	s.UserData = &v
	return s
}

type CreateMediaConvertTaskRequestSources struct {
	Duration  *float64                                         `json:"Duration,omitempty" xml:"Duration,omitempty"`
	StartTime *float64                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Subtitles []*CreateMediaConvertTaskRequestSourcesSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	URI       *string                                          `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestSources) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestSources) SetDuration(v float64) *CreateMediaConvertTaskRequestSources {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetStartTime(v float64) *CreateMediaConvertTaskRequestSources {
	s.StartTime = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetSubtitles(v []*CreateMediaConvertTaskRequestSourcesSubtitles) *CreateMediaConvertTaskRequestSources {
	s.Subtitles = v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetURI(v string) *CreateMediaConvertTaskRequestSources {
	s.URI = &v
	return s
}

type CreateMediaConvertTaskRequestSourcesSubtitles struct {
	Language   *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	TimeOffset *float64 `json:"TimeOffset,omitempty" xml:"TimeOffset,omitempty"`
	URI        *string  `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestSourcesSubtitles) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestSourcesSubtitles) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) SetLanguage(v string) *CreateMediaConvertTaskRequestSourcesSubtitles {
	s.Language = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) SetTimeOffset(v float64) *CreateMediaConvertTaskRequestSourcesSubtitles {
	s.TimeOffset = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) SetURI(v string) *CreateMediaConvertTaskRequestSourcesSubtitles {
	s.URI = &v
	return s
}

type CreateMediaConvertTaskRequestTargets struct {
	Audio     *CreateMediaConvertTaskRequestTargetsAudio    `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	Container *string                                       `json:"Container,omitempty" xml:"Container,omitempty"`
	Image     *CreateMediaConvertTaskRequestTargetsImage    `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	Preset    *PresetReference                              `json:"Preset,omitempty" xml:"Preset,omitempty"`
	Segment   *CreateMediaConvertTaskRequestTargetsSegment  `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	Speed     *float32                                      `json:"Speed,omitempty" xml:"Speed,omitempty"`
	Subtitle  *CreateMediaConvertTaskRequestTargetsSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty" type:"Struct"`
	URI       *string                                       `json:"URI,omitempty" xml:"URI,omitempty"`
	Video     *CreateMediaConvertTaskRequestTargetsVideo    `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s CreateMediaConvertTaskRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargets) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargets) SetAudio(v *CreateMediaConvertTaskRequestTargetsAudio) *CreateMediaConvertTaskRequestTargets {
	s.Audio = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetContainer(v string) *CreateMediaConvertTaskRequestTargets {
	s.Container = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetImage(v *CreateMediaConvertTaskRequestTargetsImage) *CreateMediaConvertTaskRequestTargets {
	s.Image = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetPreset(v *PresetReference) *CreateMediaConvertTaskRequestTargets {
	s.Preset = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSegment(v *CreateMediaConvertTaskRequestTargetsSegment) *CreateMediaConvertTaskRequestTargets {
	s.Segment = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSpeed(v float32) *CreateMediaConvertTaskRequestTargets {
	s.Speed = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSubtitle(v *CreateMediaConvertTaskRequestTargetsSubtitle) *CreateMediaConvertTaskRequestTargets {
	s.Subtitle = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetURI(v string) *CreateMediaConvertTaskRequestTargets {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetVideo(v *CreateMediaConvertTaskRequestTargetsVideo) *CreateMediaConvertTaskRequestTargets {
	s.Video = v
	return s
}

type CreateMediaConvertTaskRequestTargetsAudio struct {
	DisableAudio   *bool                                                    `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	FilterAudio    *CreateMediaConvertTaskRequestTargetsAudioFilterAudio    `json:"FilterAudio,omitempty" xml:"FilterAudio,omitempty" type:"Struct"`
	TranscodeAudio *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio `json:"TranscodeAudio,omitempty" xml:"TranscodeAudio,omitempty" type:"Struct"`
}

func (s CreateMediaConvertTaskRequestTargetsAudio) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsAudio) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsAudio) SetDisableAudio(v bool) *CreateMediaConvertTaskRequestTargetsAudio {
	s.DisableAudio = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudio) SetFilterAudio(v *CreateMediaConvertTaskRequestTargetsAudioFilterAudio) *CreateMediaConvertTaskRequestTargetsAudio {
	s.FilterAudio = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudio) SetTranscodeAudio(v *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) *CreateMediaConvertTaskRequestTargetsAudio {
	s.TranscodeAudio = v
	return s
}

type CreateMediaConvertTaskRequestTargetsAudioFilterAudio struct {
	Mixing *bool `json:"Mixing,omitempty" xml:"Mixing,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsAudioFilterAudio) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsAudioFilterAudio) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsAudioFilterAudio) SetMixing(v bool) *CreateMediaConvertTaskRequestTargetsAudioFilterAudio {
	s.Mixing = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio struct {
	Bitrate          *int32  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateOption    *string `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	Channel          *int32  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Codec            *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Quality          *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	SampleRate       *int32  `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SampleRateOption *string `json:"SampleRateOption,omitempty" xml:"SampleRateOption,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetBitrate(v int32) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.Bitrate = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetBitrateOption(v string) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.BitrateOption = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetChannel(v int32) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.Channel = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetCodec(v string) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.Codec = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetQuality(v int32) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.Quality = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetSampleRate(v int32) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.SampleRate = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio) SetSampleRateOption(v string) *CreateMediaConvertTaskRequestTargetsAudioTranscodeAudio {
	s.SampleRateOption = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsImage struct {
	Snapshots []*CreateMediaConvertTaskRequestTargetsImageSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	Sprites   []*CreateMediaConvertTaskRequestTargetsImageSprites   `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Repeated"`
}

func (s CreateMediaConvertTaskRequestTargetsImage) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsImage) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsImage) SetSnapshots(v []*CreateMediaConvertTaskRequestTargetsImageSnapshots) *CreateMediaConvertTaskRequestTargetsImage {
	s.Snapshots = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImage) SetSprites(v []*CreateMediaConvertTaskRequestTargetsImageSprites) *CreateMediaConvertTaskRequestTargetsImage {
	s.Sprites = v
	return s
}

type CreateMediaConvertTaskRequestTargetsImageSnapshots struct {
	Format    *string  `json:"Format,omitempty" xml:"Format,omitempty"`
	Height    *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	Interval  *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Number    *int32   `json:"Number,omitempty" xml:"Number,omitempty"`
	ScaleType *string  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	URI       *string  `json:"URI,omitempty" xml:"URI,omitempty"`
	Width     *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsImageSnapshots) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsImageSnapshots) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetFormat(v string) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.Format = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetHeight(v int32) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.Height = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetInterval(v float64) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.Interval = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetNumber(v int32) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.Number = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetScaleType(v string) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.ScaleType = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetStartTime(v float64) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.StartTime = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetURI(v string) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSnapshots) SetWidth(v int32) *CreateMediaConvertTaskRequestTargetsImageSnapshots {
	s.Width = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsImageSprites struct {
	Format      *string  `json:"Format,omitempty" xml:"Format,omitempty"`
	Interval    *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Margin      *int32   `json:"Margin,omitempty" xml:"Margin,omitempty"`
	Number      *int32   `json:"Number,omitempty" xml:"Number,omitempty"`
	Pad         *int32   `json:"Pad,omitempty" xml:"Pad,omitempty"`
	ScaleHeight *float32 `json:"ScaleHeight,omitempty" xml:"ScaleHeight,omitempty"`
	ScaleType   *string  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	ScaleWidth  *float32 `json:"ScaleWidth,omitempty" xml:"ScaleWidth,omitempty"`
	StartTime   *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TileHeight  *int32   `json:"TileHeight,omitempty" xml:"TileHeight,omitempty"`
	TileWidth   *int32   `json:"TileWidth,omitempty" xml:"TileWidth,omitempty"`
	URI         *string  `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsImageSprites) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsImageSprites) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetFormat(v string) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.Format = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetInterval(v float64) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.Interval = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetMargin(v int32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.Margin = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetNumber(v int32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.Number = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetPad(v int32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.Pad = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetScaleHeight(v float32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.ScaleHeight = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetScaleType(v string) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.ScaleType = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetScaleWidth(v float32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.ScaleWidth = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetStartTime(v float64) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.StartTime = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetTileHeight(v int32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.TileHeight = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetTileWidth(v int32) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.TileWidth = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsImageSprites) SetURI(v string) *CreateMediaConvertTaskRequestTargetsImageSprites {
	s.URI = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsSegment struct {
	Duration    *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Format      *string  `json:"Format,omitempty" xml:"Format,omitempty"`
	StartNumber *int32   `json:"StartNumber,omitempty" xml:"StartNumber,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsSegment) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsSegment) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) SetDuration(v float64) *CreateMediaConvertTaskRequestTargetsSegment {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) SetFormat(v string) *CreateMediaConvertTaskRequestTargetsSegment {
	s.Format = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) SetStartNumber(v int32) *CreateMediaConvertTaskRequestTargetsSegment {
	s.StartNumber = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsSubtitle struct {
	DisableSubtitle *bool                                                        `json:"DisableSubtitle,omitempty" xml:"DisableSubtitle,omitempty"`
	ExtractSubtitle *CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle `json:"ExtractSubtitle,omitempty" xml:"ExtractSubtitle,omitempty" type:"Struct"`
}

func (s CreateMediaConvertTaskRequestTargetsSubtitle) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsSubtitle) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsSubtitle) SetDisableSubtitle(v bool) *CreateMediaConvertTaskRequestTargetsSubtitle {
	s.DisableSubtitle = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSubtitle) SetExtractSubtitle(v *CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle) *CreateMediaConvertTaskRequestTargetsSubtitle {
	s.ExtractSubtitle = v
	return s
}

type CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	URI    *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle) SetFormat(v string) *CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle {
	s.Format = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle) SetURI(v string) *CreateMediaConvertTaskRequestTargetsSubtitleExtractSubtitle {
	s.URI = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsVideo struct {
	DisableVideo   *bool                                                    `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	FilterVideo    *CreateMediaConvertTaskRequestTargetsVideoFilterVideo    `json:"FilterVideo,omitempty" xml:"FilterVideo,omitempty" type:"Struct"`
	TranscodeVideo *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo `json:"TranscodeVideo,omitempty" xml:"TranscodeVideo,omitempty" type:"Struct"`
}

func (s CreateMediaConvertTaskRequestTargetsVideo) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsVideo) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsVideo) SetDisableVideo(v bool) *CreateMediaConvertTaskRequestTargetsVideo {
	s.DisableVideo = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideo) SetFilterVideo(v *CreateMediaConvertTaskRequestTargetsVideoFilterVideo) *CreateMediaConvertTaskRequestTargetsVideo {
	s.FilterVideo = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideo) SetTranscodeVideo(v *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) *CreateMediaConvertTaskRequestTargetsVideo {
	s.TranscodeVideo = v
	return s
}

type CreateMediaConvertTaskRequestTargetsVideoFilterVideo struct {
	Delogos    []*CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos    `json:"Delogos,omitempty" xml:"Delogos,omitempty" type:"Repeated"`
	Watermarks []*CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s CreateMediaConvertTaskRequestTargetsVideoFilterVideo) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsVideoFilterVideo) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideo) SetDelogos(v []*CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) *CreateMediaConvertTaskRequestTargetsVideoFilterVideo {
	s.Delogos = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideo) SetWatermarks(v []*CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) *CreateMediaConvertTaskRequestTargetsVideoFilterVideo {
	s.Watermarks = v
	return s
}

type CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos struct {
	Duration  *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Dx        *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy        *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	ReferPos  *string  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetDuration(v float64) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetDx(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.Dx = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetDy(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.Dy = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetHeight(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.Height = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetReferPos(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.ReferPos = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetStartTime(v float64) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.StartTime = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos) SetWidth(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoDelogos {
	s.Width = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks struct {
	BorderColor *string  `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	BorderWidth *int32   `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	Content     *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	Duration    *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Dx          *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy          *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	FontApha    *float32 `json:"FontApha,omitempty" xml:"FontApha,omitempty"`
	FontColor   *string  `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontName    *string  `json:"FontName,omitempty" xml:"FontName,omitempty"`
	FontSize    *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	Height      *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	ReferPos    *string  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	StartTime   *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type        *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	URI         *string  `json:"URI,omitempty" xml:"URI,omitempty"`
	Width       *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetBorderColor(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.BorderColor = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetBorderWidth(v int32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.BorderWidth = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetContent(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Content = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetDuration(v float64) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetDx(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Dx = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetDy(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Dy = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetFontApha(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.FontApha = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetFontColor(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.FontColor = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetFontName(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.FontName = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetFontSize(v int32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.FontSize = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetHeight(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Height = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetReferPos(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.ReferPos = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetStartTime(v float64) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.StartTime = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetType(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Type = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetURI(v string) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks) SetWidth(v float32) *CreateMediaConvertTaskRequestTargetsVideoFilterVideoWatermarks {
	s.Width = &v
	return s
}

type CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo struct {
	AdaptiveResolutionDirection *bool    `json:"AdaptiveResolutionDirection,omitempty" xml:"AdaptiveResolutionDirection,omitempty"`
	BFrames                     *int32   `json:"BFrames,omitempty" xml:"BFrames,omitempty"`
	Bitrate                     *int32   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateOption               *string  `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	BufferSize                  *int32   `json:"BufferSize,omitempty" xml:"BufferSize,omitempty"`
	CRF                         *float32 `json:"CRF,omitempty" xml:"CRF,omitempty"`
	Codec                       *string  `json:"Codec,omitempty" xml:"Codec,omitempty"`
	FrameRate                   *float32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	FrameRateOption             *string  `json:"FrameRateOption,omitempty" xml:"FrameRateOption,omitempty"`
	GOPSize                     *int32   `json:"GOPSize,omitempty" xml:"GOPSize,omitempty"`
	MaxBitrate                  *int32   `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	PixelFormat                 *string  `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	Refs                        *int32   `json:"Refs,omitempty" xml:"Refs,omitempty"`
	Resolution                  *string  `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	ResolutionOption            *string  `json:"ResolutionOption,omitempty" xml:"ResolutionOption,omitempty"`
	Rotation                    *int32   `json:"Rotation,omitempty" xml:"Rotation,omitempty"`
	ScaleType                   *string  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetAdaptiveResolutionDirection(v bool) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.AdaptiveResolutionDirection = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetBFrames(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.BFrames = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetBitrate(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.Bitrate = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetBitrateOption(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.BitrateOption = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetBufferSize(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.BufferSize = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetCRF(v float32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.CRF = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetCodec(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.Codec = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetFrameRate(v float32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.FrameRate = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetFrameRateOption(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.FrameRateOption = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetGOPSize(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.GOPSize = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetMaxBitrate(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.MaxBitrate = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetPixelFormat(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.PixelFormat = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetRefs(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.Refs = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetResolution(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.Resolution = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetResolutionOption(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.ResolutionOption = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetRotation(v int32) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.Rotation = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo) SetScaleType(v string) *CreateMediaConvertTaskRequestTargetsVideoTranscodeVideo {
	s.ScaleType = &v
	return s
}

type CreateMediaConvertTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourcesShrink          *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetsShrink          *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateMediaConvertTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetNotifyTopicName(v string) *CreateMediaConvertTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetProjectName(v string) *CreateMediaConvertTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetSourcesShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetTagsShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetTargetsShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.TargetsShrink = &v
	return s
}

func (s *CreateMediaConvertTaskShrinkRequest) SetUserData(v string) *CreateMediaConvertTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateMediaConvertTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMediaConvertTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskResponseBody) SetEventId(v string) *CreateMediaConvertTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateMediaConvertTaskResponseBody) SetRequestId(v string) *CreateMediaConvertTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaConvertTaskResponseBody) SetTaskId(v string) *CreateMediaConvertTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateMediaConvertTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMediaConvertTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMediaConvertTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskResponse) SetHeaders(v map[string]*string) *CreateMediaConvertTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaConvertTaskResponse) SetStatusCode(v int32) *CreateMediaConvertTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaConvertTaskResponse) SetBody(v *CreateMediaConvertTaskResponseBody) *CreateMediaConvertTaskResponse {
	s.Body = v
	return s
}

type CreateOfficeConversionTaskRequest struct {
	CredentialConfig *CredentialConfig      `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	EndPage          *int64                 `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	FirstPage        *bool                  `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	FitToHeight      *bool                  `json:"FitToHeight,omitempty" xml:"FitToHeight,omitempty"`
	FitToWidth       *bool                  `json:"FitToWidth,omitempty" xml:"FitToWidth,omitempty"`
	HoldLineFeed     *bool                  `json:"HoldLineFeed,omitempty" xml:"HoldLineFeed,omitempty"`
	ImageDPI         *int64                 `json:"ImageDPI,omitempty" xml:"ImageDPI,omitempty"`
	LongPicture      *bool                  `json:"LongPicture,omitempty" xml:"LongPicture,omitempty"`
	LongText         *bool                  `json:"LongText,omitempty" xml:"LongText,omitempty"`
	MaxSheetColumn   *int64                 `json:"MaxSheetColumn,omitempty" xml:"MaxSheetColumn,omitempty"`
	MaxSheetRow      *int64                 `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	NotifyTopicName  *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Pages            *string                `json:"Pages,omitempty" xml:"Pages,omitempty"`
	PaperHorizontal  *bool                  `json:"PaperHorizontal,omitempty" xml:"PaperHorizontal,omitempty"`
	PaperSize        *string                `json:"PaperSize,omitempty" xml:"PaperSize,omitempty"`
	Password         *string                `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName      *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality          *int64                 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	ScalePercentage  *int64                 `json:"ScalePercentage,omitempty" xml:"ScalePercentage,omitempty"`
	SheetCount       *int64                 `json:"SheetCount,omitempty" xml:"SheetCount,omitempty"`
	SheetIndex       *int64                 `json:"SheetIndex,omitempty" xml:"SheetIndex,omitempty"`
	ShowComments     *bool                  `json:"ShowComments,omitempty" xml:"ShowComments,omitempty"`
	SourceType       *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceURI        *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	StartPage        *int64                 `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetType       *string                `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TargetURI        *string                `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	TargetURIPrefix  *string                `json:"TargetURIPrefix,omitempty" xml:"TargetURIPrefix,omitempty"`
	TrimPolicy       *TrimPolicy            `json:"TrimPolicy,omitempty" xml:"TrimPolicy,omitempty"`
	UserData         *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateOfficeConversionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateOfficeConversionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetEndPage(v int64) *CreateOfficeConversionTaskRequest {
	s.EndPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFirstPage(v bool) *CreateOfficeConversionTaskRequest {
	s.FirstPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFitToHeight(v bool) *CreateOfficeConversionTaskRequest {
	s.FitToHeight = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetFitToWidth(v bool) *CreateOfficeConversionTaskRequest {
	s.FitToWidth = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetHoldLineFeed(v bool) *CreateOfficeConversionTaskRequest {
	s.HoldLineFeed = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetImageDPI(v int64) *CreateOfficeConversionTaskRequest {
	s.ImageDPI = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetLongPicture(v bool) *CreateOfficeConversionTaskRequest {
	s.LongPicture = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetLongText(v bool) *CreateOfficeConversionTaskRequest {
	s.LongText = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetColumn(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetColumn = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetMaxSheetRow(v int64) *CreateOfficeConversionTaskRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetNotifyTopicName(v string) *CreateOfficeConversionTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPages(v string) *CreateOfficeConversionTaskRequest {
	s.Pages = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPaperHorizontal(v bool) *CreateOfficeConversionTaskRequest {
	s.PaperHorizontal = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPaperSize(v string) *CreateOfficeConversionTaskRequest {
	s.PaperSize = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetPassword(v string) *CreateOfficeConversionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetProjectName(v string) *CreateOfficeConversionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetQuality(v int64) *CreateOfficeConversionTaskRequest {
	s.Quality = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetScalePercentage(v int64) *CreateOfficeConversionTaskRequest {
	s.ScalePercentage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSheetCount(v int64) *CreateOfficeConversionTaskRequest {
	s.SheetCount = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSheetIndex(v int64) *CreateOfficeConversionTaskRequest {
	s.SheetIndex = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetShowComments(v bool) *CreateOfficeConversionTaskRequest {
	s.ShowComments = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSourceType(v string) *CreateOfficeConversionTaskRequest {
	s.SourceType = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetSourceURI(v string) *CreateOfficeConversionTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetStartPage(v int64) *CreateOfficeConversionTaskRequest {
	s.StartPage = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTags(v map[string]interface{}) *CreateOfficeConversionTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTargetType(v string) *CreateOfficeConversionTaskRequest {
	s.TargetType = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTargetURI(v string) *CreateOfficeConversionTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTargetURIPrefix(v string) *CreateOfficeConversionTaskRequest {
	s.TargetURIPrefix = &v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetTrimPolicy(v *TrimPolicy) *CreateOfficeConversionTaskRequest {
	s.TrimPolicy = v
	return s
}

func (s *CreateOfficeConversionTaskRequest) SetUserData(v string) *CreateOfficeConversionTaskRequest {
	s.UserData = &v
	return s
}

type CreateOfficeConversionTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	EndPage                *int64  `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	FirstPage              *bool   `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	FitToHeight            *bool   `json:"FitToHeight,omitempty" xml:"FitToHeight,omitempty"`
	FitToWidth             *bool   `json:"FitToWidth,omitempty" xml:"FitToWidth,omitempty"`
	HoldLineFeed           *bool   `json:"HoldLineFeed,omitempty" xml:"HoldLineFeed,omitempty"`
	ImageDPI               *int64  `json:"ImageDPI,omitempty" xml:"ImageDPI,omitempty"`
	LongPicture            *bool   `json:"LongPicture,omitempty" xml:"LongPicture,omitempty"`
	LongText               *bool   `json:"LongText,omitempty" xml:"LongText,omitempty"`
	MaxSheetColumn         *int64  `json:"MaxSheetColumn,omitempty" xml:"MaxSheetColumn,omitempty"`
	MaxSheetRow            *int64  `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Pages                  *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	PaperHorizontal        *bool   `json:"PaperHorizontal,omitempty" xml:"PaperHorizontal,omitempty"`
	PaperSize              *string `json:"PaperSize,omitempty" xml:"PaperSize,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality                *int64  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	ScalePercentage        *int64  `json:"ScalePercentage,omitempty" xml:"ScalePercentage,omitempty"`
	SheetCount             *int64  `json:"SheetCount,omitempty" xml:"SheetCount,omitempty"`
	SheetIndex             *int64  `json:"SheetIndex,omitempty" xml:"SheetIndex,omitempty"`
	ShowComments           *bool   `json:"ShowComments,omitempty" xml:"ShowComments,omitempty"`
	SourceType             *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	StartPage              *int64  `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetType             *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TargetURI              *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	TargetURIPrefix        *string `json:"TargetURIPrefix,omitempty" xml:"TargetURIPrefix,omitempty"`
	TrimPolicyShrink       *string `json:"TrimPolicy,omitempty" xml:"TrimPolicy,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateOfficeConversionTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetEndPage(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.EndPage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetFirstPage(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.FirstPage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetFitToHeight(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.FitToHeight = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetFitToWidth(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.FitToWidth = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetHoldLineFeed(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.HoldLineFeed = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetImageDPI(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.ImageDPI = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetLongPicture(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.LongPicture = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetLongText(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.LongText = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetMaxSheetColumn(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.MaxSheetColumn = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetMaxSheetRow(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetNotifyTopicName(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPages(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.Pages = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPaperHorizontal(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.PaperHorizontal = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPaperSize(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.PaperSize = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPassword(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetProjectName(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetQuality(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.Quality = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetScalePercentage(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.ScalePercentage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSheetCount(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.SheetCount = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSheetIndex(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.SheetIndex = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetShowComments(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.ShowComments = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSourceType(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSourceURI(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetStartPage(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.StartPage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTagsShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTargetType(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTargetURI(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTargetURIPrefix(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TargetURIPrefix = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTrimPolicyShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TrimPolicyShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetUserData(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateOfficeConversionTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateOfficeConversionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponseBody) SetEventId(v string) *CreateOfficeConversionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetRequestId(v string) *CreateOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetTaskId(v string) *CreateOfficeConversionTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateOfficeConversionTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOfficeConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOfficeConversionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOfficeConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponse) SetHeaders(v map[string]*string) *CreateOfficeConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetStatusCode(v int32) *CreateOfficeConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOfficeConversionTaskResponse) SetBody(v *CreateOfficeConversionTaskResponseBody) *CreateOfficeConversionTaskResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineConcurrency       *int64  `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectQueriesPerSecond *int64  `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	ServiceRole             *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetDatasetMaxBindCount(v int64) *CreateProjectRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxEntityCount(v int64) *CreateProjectRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxFileCount(v int64) *CreateProjectRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxRelationCount(v int64) *CreateProjectRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxTotalFileSize(v int64) *CreateProjectRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetEngineConcurrency(v int64) *CreateProjectRequest {
	s.EngineConcurrency = &v
	return s
}

func (s *CreateProjectRequest) SetProjectMaxDatasetCount(v int64) *CreateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetProjectQueriesPerSecond(v int64) *CreateProjectRequest {
	s.ProjectQueriesPerSecond = &v
	return s
}

func (s *CreateProjectRequest) SetServiceRole(v string) *CreateProjectRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateProjectRequest) SetTemplateId(v string) *CreateProjectRequest {
	s.TemplateId = &v
	return s
}

type CreateProjectResponseBody struct {
	Project   *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProject(v *Project) *CreateProjectResponseBody {
	s.Project = v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateStoryRequest struct {
	CustomId        *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels    map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName     *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxFileCount    *int64                 `json:"MaxFileCount,omitempty" xml:"MaxFileCount,omitempty"`
	MinFileCount    *int64                 `json:"MinFileCount,omitempty" xml:"MinFileCount,omitempty"`
	NotifyTopicName *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ObjectId        *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName     *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryEndTime    *string                `json:"StoryEndTime,omitempty" xml:"StoryEndTime,omitempty"`
	StoryName       *string                `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StoryStartTime  *string                `json:"StoryStartTime,omitempty" xml:"StoryStartTime,omitempty"`
	StorySubType    *string                `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType       *string                `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	Tags            map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData        *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateStoryRequest) SetCustomId(v string) *CreateStoryRequest {
	s.CustomId = &v
	return s
}

func (s *CreateStoryRequest) SetCustomLabels(v map[string]interface{}) *CreateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateStoryRequest) SetDatasetName(v string) *CreateStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateStoryRequest) SetMaxFileCount(v int64) *CreateStoryRequest {
	s.MaxFileCount = &v
	return s
}

func (s *CreateStoryRequest) SetMinFileCount(v int64) *CreateStoryRequest {
	s.MinFileCount = &v
	return s
}

func (s *CreateStoryRequest) SetNotifyTopicName(v string) *CreateStoryRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateStoryRequest) SetObjectId(v string) *CreateStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateStoryRequest) SetProjectName(v string) *CreateStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateStoryRequest) SetStoryEndTime(v string) *CreateStoryRequest {
	s.StoryEndTime = &v
	return s
}

func (s *CreateStoryRequest) SetStoryName(v string) *CreateStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateStoryRequest) SetStoryStartTime(v string) *CreateStoryRequest {
	s.StoryStartTime = &v
	return s
}

func (s *CreateStoryRequest) SetStorySubType(v string) *CreateStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateStoryRequest) SetStoryType(v string) *CreateStoryRequest {
	s.StoryType = &v
	return s
}

func (s *CreateStoryRequest) SetTags(v map[string]interface{}) *CreateStoryRequest {
	s.Tags = v
	return s
}

func (s *CreateStoryRequest) SetUserData(v string) *CreateStoryRequest {
	s.UserData = &v
	return s
}

type CreateStoryShrinkRequest struct {
	CustomId           *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxFileCount       *int64  `json:"MaxFileCount,omitempty" xml:"MaxFileCount,omitempty"`
	MinFileCount       *int64  `json:"MinFileCount,omitempty" xml:"MinFileCount,omitempty"`
	NotifyTopicName    *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ObjectId           *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryEndTime       *string `json:"StoryEndTime,omitempty" xml:"StoryEndTime,omitempty"`
	StoryName          *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StoryStartTime     *string `json:"StoryStartTime,omitempty" xml:"StoryStartTime,omitempty"`
	StorySubType       *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType          *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateStoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStoryShrinkRequest) SetCustomId(v string) *CreateStoryShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetCustomLabelsShrink(v string) *CreateStoryShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetDatasetName(v string) *CreateStoryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetMaxFileCount(v int64) *CreateStoryShrinkRequest {
	s.MaxFileCount = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetMinFileCount(v int64) *CreateStoryShrinkRequest {
	s.MinFileCount = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetNotifyTopicName(v string) *CreateStoryShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetObjectId(v string) *CreateStoryShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetProjectName(v string) *CreateStoryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryEndTime(v string) *CreateStoryShrinkRequest {
	s.StoryEndTime = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryName(v string) *CreateStoryShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryStartTime(v string) *CreateStoryShrinkRequest {
	s.StoryStartTime = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStorySubType(v string) *CreateStoryShrinkRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryType(v string) *CreateStoryShrinkRequest {
	s.StoryType = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetTagsShrink(v string) *CreateStoryShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetUserData(v string) *CreateStoryShrinkRequest {
	s.UserData = &v
	return s
}

type CreateStoryResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoryResponseBody) SetEventId(v string) *CreateStoryResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateStoryResponseBody) SetRequestId(v string) *CreateStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStoryResponseBody) SetTaskId(v string) *CreateStoryResponseBody {
	s.TaskId = &v
	return s
}

type CreateStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryResponse) GoString() string {
	return s.String()
}

func (s *CreateStoryResponse) SetHeaders(v map[string]*string) *CreateStoryResponse {
	s.Headers = v
	return s
}

func (s *CreateStoryResponse) SetStatusCode(v int32) *CreateStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoryResponse) SetBody(v *CreateStoryResponseBody) *CreateStoryResponse {
	s.Body = v
	return s
}

type CreateVideoModerationTaskRequest struct {
	CredentialConfig *CredentialConfig      `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Interval         *int64                 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames        *int64                 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	NotifyTopicName  *string                `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName      *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer         *string                `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	Scenes           []*string              `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	SourceURI        *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags             map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData         *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoModerationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoModerationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateVideoModerationTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetInterval(v int64) *CreateVideoModerationTaskRequest {
	s.Interval = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetMaxFrames(v int64) *CreateVideoModerationTaskRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetNotifyTopicName(v string) *CreateVideoModerationTaskRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetProjectName(v string) *CreateVideoModerationTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetReviewer(v string) *CreateVideoModerationTaskRequest {
	s.Reviewer = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetScenes(v []*string) *CreateVideoModerationTaskRequest {
	s.Scenes = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetSourceURI(v string) *CreateVideoModerationTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetTags(v map[string]interface{}) *CreateVideoModerationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetUserData(v string) *CreateVideoModerationTaskRequest {
	s.UserData = &v
	return s
}

type CreateVideoModerationTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Interval               *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames              *int64  `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer               *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	ScenesShrink           *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoModerationTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoModerationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetInterval(v int64) *CreateVideoModerationTaskShrinkRequest {
	s.Interval = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetMaxFrames(v int64) *CreateVideoModerationTaskShrinkRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetNotifyTopicName(v string) *CreateVideoModerationTaskShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetProjectName(v string) *CreateVideoModerationTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetReviewer(v string) *CreateVideoModerationTaskShrinkRequest {
	s.Reviewer = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetScenesShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.ScenesShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetSourceURI(v string) *CreateVideoModerationTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetTagsShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetUserData(v string) *CreateVideoModerationTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateVideoModerationTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoModerationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoModerationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskResponseBody) SetEventId(v string) *CreateVideoModerationTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateVideoModerationTaskResponseBody) SetRequestId(v string) *CreateVideoModerationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoModerationTaskResponseBody) SetTaskId(v string) *CreateVideoModerationTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVideoModerationTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVideoModerationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoModerationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoModerationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoModerationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoModerationTaskResponse) SetStatusCode(v int32) *CreateVideoModerationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoModerationTaskResponse) SetBody(v *CreateVideoModerationTaskResponseBody) *CreateVideoModerationTaskResponse {
	s.Body = v
	return s
}

type DeleteBindingRequest struct {
	Cleanup     *bool   `json:"Cleanup,omitempty" xml:"Cleanup,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DeleteBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) SetCleanup(v bool) *DeleteBindingRequest {
	s.Cleanup = &v
	return s
}

func (s *DeleteBindingRequest) SetDatasetName(v string) *DeleteBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteBindingRequest) SetProjectName(v string) *DeleteBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteBindingRequest) SetURI(v string) *DeleteBindingRequest {
	s.URI = &v
	return s
}

type DeleteBindingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponseBody) SetRequestId(v string) *DeleteBindingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponse) SetHeaders(v map[string]*string) *DeleteBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteBindingResponse) SetStatusCode(v int32) *DeleteBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBindingResponse) SetBody(v *DeleteBindingResponseBody) *DeleteBindingResponse {
	s.Body = v
	return s
}

type DeleteDatasetRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) SetDatasetName(v string) *DeleteDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteDatasetRequest) SetProjectName(v string) *DeleteDatasetRequest {
	s.ProjectName = &v
	return s
}

type DeleteDatasetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponseBody) SetRequestId(v string) *DeleteDatasetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponse) SetHeaders(v map[string]*string) *DeleteDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetResponse) SetStatusCode(v int32) *DeleteDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetResponse) SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse {
	s.Body = v
	return s
}

type DeleteFileMetaRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DeleteFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaRequest) SetDatasetName(v string) *DeleteFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteFileMetaRequest) SetProjectName(v string) *DeleteFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteFileMetaRequest) SetURI(v string) *DeleteFileMetaRequest {
	s.URI = &v
	return s
}

type DeleteFileMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaResponseBody) SetRequestId(v string) *DeleteFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaResponse) SetHeaders(v map[string]*string) *DeleteFileMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileMetaResponse) SetStatusCode(v int32) *DeleteFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileMetaResponse) SetBody(v *DeleteFileMetaResponseBody) *DeleteFileMetaResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectName(v string) *DeleteProjectRequest {
	s.ProjectName = &v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteStoryRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteStoryRequest) SetDatasetName(v string) *DeleteStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteStoryRequest) SetObjectId(v string) *DeleteStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *DeleteStoryRequest) SetProjectName(v string) *DeleteStoryRequest {
	s.ProjectName = &v
	return s
}

type DeleteStoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponseBody) SetRequestId(v string) *DeleteStoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponse) SetHeaders(v map[string]*string) *DeleteStoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteStoryResponse) SetStatusCode(v int32) *DeleteStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStoryResponse) SetBody(v *DeleteStoryResponseBody) *DeleteStoryResponse {
	s.Body = v
	return s
}

type DetachOSSBucketRequest struct {
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
}

func (s DetachOSSBucketRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachOSSBucketRequest) GoString() string {
	return s.String()
}

func (s *DetachOSSBucketRequest) SetOSSBucket(v string) *DetachOSSBucketRequest {
	s.OSSBucket = &v
	return s
}

type DetachOSSBucketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachOSSBucketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachOSSBucketResponseBody) GoString() string {
	return s.String()
}

func (s *DetachOSSBucketResponseBody) SetRequestId(v string) *DetachOSSBucketResponseBody {
	s.RequestId = &v
	return s
}

type DetachOSSBucketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachOSSBucketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachOSSBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachOSSBucketResponse) GoString() string {
	return s.String()
}

func (s *DetachOSSBucketResponse) SetHeaders(v map[string]*string) *DetachOSSBucketResponse {
	s.Headers = v
	return s
}

func (s *DetachOSSBucketResponse) SetStatusCode(v int32) *DetachOSSBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachOSSBucketResponse) SetBody(v *DetachOSSBucketResponseBody) *DetachOSSBucketResponse {
	s.Body = v
	return s
}

type DetectImageBodiesRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sensitivity      *float32          `json:"Sensitivity,omitempty" xml:"Sensitivity,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageBodiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageBodiesRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageBodiesRequest) SetProjectName(v string) *DetectImageBodiesRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageBodiesRequest) SetSensitivity(v float32) *DetectImageBodiesRequest {
	s.Sensitivity = &v
	return s
}

func (s *DetectImageBodiesRequest) SetSourceURI(v string) *DetectImageBodiesRequest {
	s.SourceURI = &v
	return s
}

type DetectImageBodiesShrinkRequest struct {
	CredentialConfigShrink *string  `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sensitivity            *float32 `json:"Sensitivity,omitempty" xml:"Sensitivity,omitempty"`
	SourceURI              *string  `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageBodiesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageBodiesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) SetProjectName(v string) *DetectImageBodiesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) SetSensitivity(v float32) *DetectImageBodiesShrinkRequest {
	s.Sensitivity = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) SetSourceURI(v string) *DetectImageBodiesShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectImageBodiesResponseBody struct {
	Bodies    []*Body `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageBodiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponseBody) SetBodies(v []*Body) *DetectImageBodiesResponseBody {
	s.Bodies = v
	return s
}

func (s *DetectImageBodiesResponseBody) SetRequestId(v string) *DetectImageBodiesResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageBodiesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageBodiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageBodiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageBodiesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesResponse) SetHeaders(v map[string]*string) *DetectImageBodiesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageBodiesResponse) SetStatusCode(v int32) *DetectImageBodiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageBodiesResponse) SetBody(v *DetectImageBodiesResponseBody) *DetectImageBodiesResponse {
	s.Body = v
	return s
}

type DetectImageCodesRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCodesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCodesRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageCodesRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageCodesRequest) SetProjectName(v string) *DetectImageCodesRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCodesRequest) SetSourceURI(v string) *DetectImageCodesRequest {
	s.SourceURI = &v
	return s
}

type DetectImageCodesShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCodesShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageCodesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageCodesShrinkRequest) SetProjectName(v string) *DetectImageCodesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCodesShrinkRequest) SetSourceURI(v string) *DetectImageCodesShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectImageCodesResponseBody struct {
	Codes     []*Codes `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageCodesResponseBody) SetCodes(v []*Codes) *DetectImageCodesResponseBody {
	s.Codes = v
	return s
}

func (s *DetectImageCodesResponseBody) SetRequestId(v string) *DetectImageCodesResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageCodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageCodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCodesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageCodesResponse) SetHeaders(v map[string]*string) *DetectImageCodesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageCodesResponse) SetStatusCode(v int32) *DetectImageCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageCodesResponse) SetBody(v *DetectImageCodesResponseBody) *DetectImageCodesResponse {
	s.Body = v
	return s
}

type DetectImageCroppingRequest struct {
	AspectRatios     *string           `json:"AspectRatios,omitempty" xml:"AspectRatios,omitempty"`
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCroppingRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCroppingRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingRequest) SetAspectRatios(v string) *DetectImageCroppingRequest {
	s.AspectRatios = &v
	return s
}

func (s *DetectImageCroppingRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageCroppingRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageCroppingRequest) SetProjectName(v string) *DetectImageCroppingRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCroppingRequest) SetSourceURI(v string) *DetectImageCroppingRequest {
	s.SourceURI = &v
	return s
}

type DetectImageCroppingShrinkRequest struct {
	AspectRatios           *string `json:"AspectRatios,omitempty" xml:"AspectRatios,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCroppingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCroppingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingShrinkRequest) SetAspectRatios(v string) *DetectImageCroppingShrinkRequest {
	s.AspectRatios = &v
	return s
}

func (s *DetectImageCroppingShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageCroppingShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageCroppingShrinkRequest) SetProjectName(v string) *DetectImageCroppingShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCroppingShrinkRequest) SetSourceURI(v string) *DetectImageCroppingShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectImageCroppingResponseBody struct {
	Croppings []*CroppingSuggestion `json:"Croppings,omitempty" xml:"Croppings,omitempty" type:"Repeated"`
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageCroppingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCroppingResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingResponseBody) SetCroppings(v []*CroppingSuggestion) *DetectImageCroppingResponseBody {
	s.Croppings = v
	return s
}

func (s *DetectImageCroppingResponseBody) SetRequestId(v string) *DetectImageCroppingResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageCroppingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageCroppingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageCroppingResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCroppingResponse) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingResponse) SetHeaders(v map[string]*string) *DetectImageCroppingResponse {
	s.Headers = v
	return s
}

func (s *DetectImageCroppingResponse) SetStatusCode(v int32) *DetectImageCroppingResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageCroppingResponse) SetBody(v *DetectImageCroppingResponseBody) *DetectImageCroppingResponse {
	s.Body = v
	return s
}

type DetectImageFacesRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageFacesRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageFacesRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageFacesRequest) SetProjectName(v string) *DetectImageFacesRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageFacesRequest) SetSourceURI(v string) *DetectImageFacesRequest {
	s.SourceURI = &v
	return s
}

type DetectImageFacesShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageFacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageFacesShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageFacesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageFacesShrinkRequest) SetProjectName(v string) *DetectImageFacesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageFacesShrinkRequest) SetSourceURI(v string) *DetectImageFacesShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectImageFacesResponseBody struct {
	Faces     []*Figure `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponseBody) SetFaces(v []*Figure) *DetectImageFacesResponseBody {
	s.Faces = v
	return s
}

func (s *DetectImageFacesResponseBody) SetRequestId(v string) *DetectImageFacesResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageFacesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageFacesResponse) GoString() string {
	return s.String()
}

func (s *DetectImageFacesResponse) SetHeaders(v map[string]*string) *DetectImageFacesResponse {
	s.Headers = v
	return s
}

func (s *DetectImageFacesResponse) SetStatusCode(v int32) *DetectImageFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageFacesResponse) SetBody(v *DetectImageFacesResponseBody) *DetectImageFacesResponse {
	s.Body = v
	return s
}

type DetectImageLabelsRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Threshold        *float32          `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectImageLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageLabelsRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageLabelsRequest) SetProjectName(v string) *DetectImageLabelsRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageLabelsRequest) SetSourceURI(v string) *DetectImageLabelsRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageLabelsRequest) SetThreshold(v float32) *DetectImageLabelsRequest {
	s.Threshold = &v
	return s
}

type DetectImageLabelsShrinkRequest struct {
	CredentialConfigShrink *string  `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string  `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Threshold              *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectImageLabelsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLabelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageLabelsShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageLabelsShrinkRequest) SetProjectName(v string) *DetectImageLabelsShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageLabelsShrinkRequest) SetSourceURI(v string) *DetectImageLabelsShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageLabelsShrinkRequest) SetThreshold(v float32) *DetectImageLabelsShrinkRequest {
	s.Threshold = &v
	return s
}

type DetectImageLabelsResponseBody struct {
	Labels    []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsResponseBody) SetLabels(v []*Label) *DetectImageLabelsResponseBody {
	s.Labels = v
	return s
}

func (s *DetectImageLabelsResponseBody) SetRequestId(v string) *DetectImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsResponse) SetHeaders(v map[string]*string) *DetectImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageLabelsResponse) SetStatusCode(v int32) *DetectImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageLabelsResponse) SetBody(v *DetectImageLabelsResponseBody) *DetectImageLabelsResponse {
	s.Body = v
	return s
}

type DetectImageScoreRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageScoreRequest) GoString() string {
	return s.String()
}

func (s *DetectImageScoreRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageScoreRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageScoreRequest) SetProjectName(v string) *DetectImageScoreRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageScoreRequest) SetSourceURI(v string) *DetectImageScoreRequest {
	s.SourceURI = &v
	return s
}

type DetectImageScoreShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageScoreShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageScoreShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageScoreShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageScoreShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageScoreShrinkRequest) SetProjectName(v string) *DetectImageScoreShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageScoreShrinkRequest) SetSourceURI(v string) *DetectImageScoreShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectImageScoreResponseBody struct {
	ImageScore *DetectImageScoreResponseBodyImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageScoreResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageScoreResponseBody) SetImageScore(v *DetectImageScoreResponseBodyImageScore) *DetectImageScoreResponseBody {
	s.ImageScore = v
	return s
}

func (s *DetectImageScoreResponseBody) SetRequestId(v string) *DetectImageScoreResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageScoreResponseBodyImageScore struct {
	OverallQualityScore *float32 `json:"OverallQualityScore,omitempty" xml:"OverallQualityScore,omitempty"`
}

func (s DetectImageScoreResponseBodyImageScore) String() string {
	return tea.Prettify(s)
}

func (s DetectImageScoreResponseBodyImageScore) GoString() string {
	return s.String()
}

func (s *DetectImageScoreResponseBodyImageScore) SetOverallQualityScore(v float32) *DetectImageScoreResponseBodyImageScore {
	s.OverallQualityScore = &v
	return s
}

type DetectImageScoreResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageScoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageScoreResponse) GoString() string {
	return s.String()
}

func (s *DetectImageScoreResponse) SetHeaders(v map[string]*string) *DetectImageScoreResponse {
	s.Headers = v
	return s
}

func (s *DetectImageScoreResponse) SetStatusCode(v int32) *DetectImageScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageScoreResponse) SetBody(v *DetectImageScoreResponseBody) *DetectImageScoreResponse {
	s.Body = v
	return s
}

type DetectTextAnomalyRequest struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DetectTextAnomalyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectTextAnomalyRequest) GoString() string {
	return s.String()
}

func (s *DetectTextAnomalyRequest) SetContent(v string) *DetectTextAnomalyRequest {
	s.Content = &v
	return s
}

func (s *DetectTextAnomalyRequest) SetProjectName(v string) *DetectTextAnomalyRequest {
	s.ProjectName = &v
	return s
}

type DetectTextAnomalyResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DetectTextAnomalyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectTextAnomalyResponseBody) GoString() string {
	return s.String()
}

func (s *DetectTextAnomalyResponseBody) SetRequestId(v string) *DetectTextAnomalyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectTextAnomalyResponseBody) SetSuggestion(v string) *DetectTextAnomalyResponseBody {
	s.Suggestion = &v
	return s
}

type DetectTextAnomalyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectTextAnomalyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectTextAnomalyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectTextAnomalyResponse) GoString() string {
	return s.String()
}

func (s *DetectTextAnomalyResponse) SetHeaders(v map[string]*string) *DetectTextAnomalyResponse {
	s.Headers = v
	return s
}

func (s *DetectTextAnomalyResponse) SetStatusCode(v int32) *DetectTextAnomalyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectTextAnomalyResponse) SetBody(v *DetectTextAnomalyResponseBody) *DetectTextAnomalyResponse {
	s.Body = v
	return s
}

type FuzzyQueryRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s FuzzyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FuzzyQueryRequest) GoString() string {
	return s.String()
}

func (s *FuzzyQueryRequest) SetDatasetName(v string) *FuzzyQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *FuzzyQueryRequest) SetMaxResults(v int64) *FuzzyQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *FuzzyQueryRequest) SetNextToken(v string) *FuzzyQueryRequest {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryRequest) SetProjectName(v string) *FuzzyQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *FuzzyQueryRequest) SetQuery(v string) *FuzzyQueryRequest {
	s.Query = &v
	return s
}

type FuzzyQueryResponseBody struct {
	Files     []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FuzzyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FuzzyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FuzzyQueryResponseBody) SetFiles(v []*File) *FuzzyQueryResponseBody {
	s.Files = v
	return s
}

func (s *FuzzyQueryResponseBody) SetNextToken(v string) *FuzzyQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryResponseBody) SetRequestId(v string) *FuzzyQueryResponseBody {
	s.RequestId = &v
	return s
}

type FuzzyQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FuzzyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FuzzyQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s FuzzyQueryResponse) GoString() string {
	return s.String()
}

func (s *FuzzyQueryResponse) SetHeaders(v map[string]*string) *FuzzyQueryResponse {
	s.Headers = v
	return s
}

func (s *FuzzyQueryResponse) SetStatusCode(v int32) *FuzzyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FuzzyQueryResponse) SetBody(v *FuzzyQueryResponseBody) *FuzzyQueryResponse {
	s.Body = v
	return s
}

type GetBindingRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBindingRequest) GoString() string {
	return s.String()
}

func (s *GetBindingRequest) SetDatasetName(v string) *GetBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *GetBindingRequest) SetProjectName(v string) *GetBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *GetBindingRequest) SetURI(v string) *GetBindingRequest {
	s.URI = &v
	return s
}

type GetBindingResponseBody struct {
	Binding   *Binding `json:"Binding,omitempty" xml:"Binding,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBindingResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindingResponseBody) SetBinding(v *Binding) *GetBindingResponseBody {
	s.Binding = v
	return s
}

func (s *GetBindingResponseBody) SetRequestId(v string) *GetBindingResponseBody {
	s.RequestId = &v
	return s
}

type GetBindingResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBindingResponse) GoString() string {
	return s.String()
}

func (s *GetBindingResponse) SetHeaders(v map[string]*string) *GetBindingResponse {
	s.Headers = v
	return s
}

func (s *GetBindingResponse) SetStatusCode(v int32) *GetBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBindingResponse) SetBody(v *GetBindingResponseBody) *GetBindingResponse {
	s.Body = v
	return s
}

type GetDatasetRequest struct {
	DatasetName    *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	WithStatistics *bool   `json:"WithStatistics,omitempty" xml:"WithStatistics,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) SetDatasetName(v string) *GetDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetRequest) SetProjectName(v string) *GetDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDatasetRequest) SetWithStatistics(v bool) *GetDatasetRequest {
	s.WithStatistics = &v
	return s
}

type GetDatasetResponseBody struct {
	Dataset   *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) SetDataset(v *Dataset) *GetDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

type GetDatasetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetResponse) SetHeaders(v map[string]*string) *GetDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetResponse) SetStatusCode(v int32) *GetDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

type GetDetectVideoLabelsResultRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetDetectVideoLabelsResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectVideoLabelsResultRequest) GoString() string {
	return s.String()
}

func (s *GetDetectVideoLabelsResultRequest) SetProjectName(v string) *GetDetectVideoLabelsResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDetectVideoLabelsResultRequest) SetTaskId(v string) *GetDetectVideoLabelsResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDetectVideoLabelsResultRequest) SetTaskType(v string) *GetDetectVideoLabelsResultRequest {
	s.TaskType = &v
	return s
}

type GetDetectVideoLabelsResultResponseBody struct {
	Code        *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime     *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventId     *string  `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Labels      []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message     *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectName *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId   *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId      *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData    *string  `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetDetectVideoLabelsResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectVideoLabelsResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectVideoLabelsResultResponseBody) SetCode(v string) *GetDetectVideoLabelsResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetEndTime(v string) *GetDetectVideoLabelsResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetEventId(v string) *GetDetectVideoLabelsResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetLabels(v []*Label) *GetDetectVideoLabelsResultResponseBody {
	s.Labels = v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetMessage(v string) *GetDetectVideoLabelsResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetProjectName(v string) *GetDetectVideoLabelsResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetRequestId(v string) *GetDetectVideoLabelsResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetStartTime(v string) *GetDetectVideoLabelsResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetStatus(v string) *GetDetectVideoLabelsResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetTaskId(v string) *GetDetectVideoLabelsResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetTaskType(v string) *GetDetectVideoLabelsResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponseBody) SetUserData(v string) *GetDetectVideoLabelsResultResponseBody {
	s.UserData = &v
	return s
}

type GetDetectVideoLabelsResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDetectVideoLabelsResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectVideoLabelsResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectVideoLabelsResultResponse) GoString() string {
	return s.String()
}

func (s *GetDetectVideoLabelsResultResponse) SetHeaders(v map[string]*string) *GetDetectVideoLabelsResultResponse {
	s.Headers = v
	return s
}

func (s *GetDetectVideoLabelsResultResponse) SetStatusCode(v int32) *GetDetectVideoLabelsResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectVideoLabelsResultResponse) SetBody(v *GetDetectVideoLabelsResultResponseBody) *GetDetectVideoLabelsResultResponse {
	s.Body = v
	return s
}

type GetFigureClusterRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetFigureClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *GetFigureClusterRequest) SetDatasetName(v string) *GetFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFigureClusterRequest) SetObjectId(v string) *GetFigureClusterRequest {
	s.ObjectId = &v
	return s
}

func (s *GetFigureClusterRequest) SetProjectName(v string) *GetFigureClusterRequest {
	s.ProjectName = &v
	return s
}

type GetFigureClusterResponseBody struct {
	FigureCluster *FigureCluster `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	RequestId     *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFigureClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetFigureClusterResponseBody) SetFigureCluster(v *FigureCluster) *GetFigureClusterResponseBody {
	s.FigureCluster = v
	return s
}

func (s *GetFigureClusterResponseBody) SetRequestId(v string) *GetFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

type GetFigureClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFigureClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *GetFigureClusterResponse) SetHeaders(v map[string]*string) *GetFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *GetFigureClusterResponse) SetStatusCode(v int32) *GetFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFigureClusterResponse) SetBody(v *GetFigureClusterResponseBody) *GetFigureClusterResponse {
	s.Body = v
	return s
}

type GetFileMetaRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *GetFileMetaRequest) SetDatasetName(v string) *GetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFileMetaRequest) SetProjectName(v string) *GetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileMetaRequest) SetURI(v string) *GetFileMetaRequest {
	s.URI = &v
	return s
}

type GetFileMetaResponseBody struct {
	Files     []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileMetaResponseBody) SetFiles(v []*File) *GetFileMetaResponseBody {
	s.Files = v
	return s
}

func (s *GetFileMetaResponseBody) SetRequestId(v string) *GetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type GetFileMetaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileMetaResponse) GoString() string {
	return s.String()
}

func (s *GetFileMetaResponse) SetHeaders(v map[string]*string) *GetFileMetaResponse {
	s.Headers = v
	return s
}

func (s *GetFileMetaResponse) SetStatusCode(v int32) *GetFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileMetaResponse) SetBody(v *GetFileMetaResponseBody) *GetFileMetaResponse {
	s.Body = v
	return s
}

type GetMediaMetaRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s GetMediaMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaRequest) GoString() string {
	return s.String()
}

func (s *GetMediaMetaRequest) SetCredentialConfig(v *CredentialConfig) *GetMediaMetaRequest {
	s.CredentialConfig = v
	return s
}

func (s *GetMediaMetaRequest) SetProjectName(v string) *GetMediaMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *GetMediaMetaRequest) SetSourceURI(v string) *GetMediaMetaRequest {
	s.SourceURI = &v
	return s
}

type GetMediaMetaShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s GetMediaMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMediaMetaShrinkRequest) SetCredentialConfigShrink(v string) *GetMediaMetaShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *GetMediaMetaShrinkRequest) SetProjectName(v string) *GetMediaMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GetMediaMetaShrinkRequest) SetSourceURI(v string) *GetMediaMetaShrinkRequest {
	s.SourceURI = &v
	return s
}

type GetMediaMetaResponseBody struct {
	Addresses      []*Address        `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	Album          *string           `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumArtist    *string           `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	Artist         *string           `json:"Artist,omitempty" xml:"Artist,omitempty"`
	AudioStreams   []*AudioStream    `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	Bitrate        *int64            `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Composer       *string           `json:"Composer,omitempty" xml:"Composer,omitempty"`
	Duration       *float64          `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatLongName *string           `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName     *string           `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	Language       *string           `json:"Language,omitempty" xml:"Language,omitempty"`
	LatLong        *string           `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	Performer      *string           `json:"Performer,omitempty" xml:"Performer,omitempty"`
	ProduceTime    *string           `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	ProgramCount   *int64            `json:"ProgramCount,omitempty" xml:"ProgramCount,omitempty"`
	RequestId      *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size           *int64            `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime      *float64          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamCount    *int64            `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	Subtitles      []*SubtitleStream `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	Title          *string           `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoHeight    *int64            `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoStreams   []*VideoStream    `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	VideoWidth     *int64            `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s GetMediaMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponseBody) SetAddresses(v []*Address) *GetMediaMetaResponseBody {
	s.Addresses = v
	return s
}

func (s *GetMediaMetaResponseBody) SetAlbum(v string) *GetMediaMetaResponseBody {
	s.Album = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetAlbumArtist(v string) *GetMediaMetaResponseBody {
	s.AlbumArtist = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetArtist(v string) *GetMediaMetaResponseBody {
	s.Artist = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetAudioStreams(v []*AudioStream) *GetMediaMetaResponseBody {
	s.AudioStreams = v
	return s
}

func (s *GetMediaMetaResponseBody) SetBitrate(v int64) *GetMediaMetaResponseBody {
	s.Bitrate = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetComposer(v string) *GetMediaMetaResponseBody {
	s.Composer = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetDuration(v float64) *GetMediaMetaResponseBody {
	s.Duration = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetFormatLongName(v string) *GetMediaMetaResponseBody {
	s.FormatLongName = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetFormatName(v string) *GetMediaMetaResponseBody {
	s.FormatName = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetLanguage(v string) *GetMediaMetaResponseBody {
	s.Language = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetLatLong(v string) *GetMediaMetaResponseBody {
	s.LatLong = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetPerformer(v string) *GetMediaMetaResponseBody {
	s.Performer = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetProduceTime(v string) *GetMediaMetaResponseBody {
	s.ProduceTime = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetProgramCount(v int64) *GetMediaMetaResponseBody {
	s.ProgramCount = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetRequestId(v string) *GetMediaMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetSize(v int64) *GetMediaMetaResponseBody {
	s.Size = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetStartTime(v float64) *GetMediaMetaResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetStreamCount(v int64) *GetMediaMetaResponseBody {
	s.StreamCount = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetSubtitles(v []*SubtitleStream) *GetMediaMetaResponseBody {
	s.Subtitles = v
	return s
}

func (s *GetMediaMetaResponseBody) SetTitle(v string) *GetMediaMetaResponseBody {
	s.Title = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetVideoHeight(v int64) *GetMediaMetaResponseBody {
	s.VideoHeight = &v
	return s
}

func (s *GetMediaMetaResponseBody) SetVideoStreams(v []*VideoStream) *GetMediaMetaResponseBody {
	s.VideoStreams = v
	return s
}

func (s *GetMediaMetaResponseBody) SetVideoWidth(v int64) *GetMediaMetaResponseBody {
	s.VideoWidth = &v
	return s
}

type GetMediaMetaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMediaMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaMetaResponse) GoString() string {
	return s.String()
}

func (s *GetMediaMetaResponse) SetHeaders(v map[string]*string) *GetMediaMetaResponse {
	s.Headers = v
	return s
}

func (s *GetMediaMetaResponse) SetStatusCode(v int32) *GetMediaMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaMetaResponse) SetBody(v *GetMediaMetaResponseBody) *GetMediaMetaResponse {
	s.Body = v
	return s
}

type GetOSSBucketAttachmentRequest struct {
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
}

func (s GetOSSBucketAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOSSBucketAttachmentRequest) GoString() string {
	return s.String()
}

func (s *GetOSSBucketAttachmentRequest) SetOSSBucket(v string) *GetOSSBucketAttachmentRequest {
	s.OSSBucket = &v
	return s
}

type GetOSSBucketAttachmentResponseBody struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOSSBucketAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOSSBucketAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSBucketAttachmentResponseBody) SetProjectName(v string) *GetOSSBucketAttachmentResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetOSSBucketAttachmentResponseBody) SetRequestId(v string) *GetOSSBucketAttachmentResponseBody {
	s.RequestId = &v
	return s
}

type GetOSSBucketAttachmentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOSSBucketAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOSSBucketAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOSSBucketAttachmentResponse) GoString() string {
	return s.String()
}

func (s *GetOSSBucketAttachmentResponse) SetHeaders(v map[string]*string) *GetOSSBucketAttachmentResponse {
	s.Headers = v
	return s
}

func (s *GetOSSBucketAttachmentResponse) SetStatusCode(v int32) *GetOSSBucketAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSBucketAttachmentResponse) SetBody(v *GetOSSBucketAttachmentResponseBody) *GetOSSBucketAttachmentResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	WithStatistics *bool   `json:"WithStatistics,omitempty" xml:"WithStatistics,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetProjectName(v string) *GetProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *GetProjectRequest) SetWithStatistics(v bool) *GetProjectRequest {
	s.WithStatistics = &v
	return s
}

type GetProjectResponseBody struct {
	Project   *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetProject(v *Project) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetStoryRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStoryRequest) GoString() string {
	return s.String()
}

func (s *GetStoryRequest) SetDatasetName(v string) *GetStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *GetStoryRequest) SetObjectId(v string) *GetStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *GetStoryRequest) SetProjectName(v string) *GetStoryRequest {
	s.ProjectName = &v
	return s
}

type GetStoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Story     *Story  `json:"Story,omitempty" xml:"Story,omitempty"`
}

func (s GetStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoryResponseBody) SetRequestId(v string) *GetStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStoryResponseBody) SetStory(v *Story) *GetStoryResponseBody {
	s.Story = v
	return s
}

type GetStoryResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoryResponse) GoString() string {
	return s.String()
}

func (s *GetStoryResponse) SetHeaders(v map[string]*string) *GetStoryResponse {
	s.Headers = v
	return s
}

func (s *GetStoryResponse) SetStatusCode(v int32) *GetStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoryResponse) SetBody(v *GetStoryResponseBody) *GetStoryResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetProjectName(v string) *GetTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskRequest) SetTaskType(v string) *GetTaskRequest {
	s.TaskType = &v
	return s
}

type GetTaskResponseBody struct {
	Code        *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime     *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventId     *string                `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Message     *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectName *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId   *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags        map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TaskId      *string                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string                `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData    *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetCode(v string) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetEndTime(v string) *GetTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBody) SetEventId(v string) *GetTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetProjectName(v string) *GetTaskResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetStartTime(v string) *GetTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBody) SetStatus(v string) *GetTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBody) SetTags(v map[string]interface{}) *GetTaskResponseBody {
	s.Tags = v
	return s
}

func (s *GetTaskResponseBody) SetTaskId(v string) *GetTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBody) SetTaskType(v string) *GetTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetTaskResponseBody) SetUserData(v string) *GetTaskResponseBody {
	s.UserData = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetWebofficeURLRequest struct {
	CachePreview     *bool                `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	CredentialConfig *CredentialConfig    `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ExternalUploaded *bool                `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	Filename         *string              `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Hidecmb          *bool                `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	NotifyTopicName  *string              `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Password         *string              `json:"Password,omitempty" xml:"Password,omitempty"`
	Permission       *WebofficePermission `json:"Permission,omitempty" xml:"Permission,omitempty"`
	PreviewPages     *int64               `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	ProjectName      *string              `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Referer          *string              `json:"Referer,omitempty" xml:"Referer,omitempty"`
	SourceURI        *string              `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	User             *WebofficeUser       `json:"User,omitempty" xml:"User,omitempty"`
	UserData         *string              `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Watermark        *WebofficeWatermark  `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GetWebofficeURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLRequest) SetCachePreview(v bool) *GetWebofficeURLRequest {
	s.CachePreview = &v
	return s
}

func (s *GetWebofficeURLRequest) SetCredentialConfig(v *CredentialConfig) *GetWebofficeURLRequest {
	s.CredentialConfig = v
	return s
}

func (s *GetWebofficeURLRequest) SetExternalUploaded(v bool) *GetWebofficeURLRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GetWebofficeURLRequest) SetFilename(v string) *GetWebofficeURLRequest {
	s.Filename = &v
	return s
}

func (s *GetWebofficeURLRequest) SetHidecmb(v bool) *GetWebofficeURLRequest {
	s.Hidecmb = &v
	return s
}

func (s *GetWebofficeURLRequest) SetNotifyTopicName(v string) *GetWebofficeURLRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetWebofficeURLRequest) SetPassword(v string) *GetWebofficeURLRequest {
	s.Password = &v
	return s
}

func (s *GetWebofficeURLRequest) SetPermission(v *WebofficePermission) *GetWebofficeURLRequest {
	s.Permission = v
	return s
}

func (s *GetWebofficeURLRequest) SetPreviewPages(v int64) *GetWebofficeURLRequest {
	s.PreviewPages = &v
	return s
}

func (s *GetWebofficeURLRequest) SetProjectName(v string) *GetWebofficeURLRequest {
	s.ProjectName = &v
	return s
}

func (s *GetWebofficeURLRequest) SetReferer(v string) *GetWebofficeURLRequest {
	s.Referer = &v
	return s
}

func (s *GetWebofficeURLRequest) SetSourceURI(v string) *GetWebofficeURLRequest {
	s.SourceURI = &v
	return s
}

func (s *GetWebofficeURLRequest) SetUser(v *WebofficeUser) *GetWebofficeURLRequest {
	s.User = v
	return s
}

func (s *GetWebofficeURLRequest) SetUserData(v string) *GetWebofficeURLRequest {
	s.UserData = &v
	return s
}

func (s *GetWebofficeURLRequest) SetWatermark(v *WebofficeWatermark) *GetWebofficeURLRequest {
	s.Watermark = v
	return s
}

type GetWebofficeURLShrinkRequest struct {
	CachePreview           *bool   `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ExternalUploaded       *bool   `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	Filename               *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Hidecmb                *bool   `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	NotifyTopicName        *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PermissionShrink       *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	PreviewPages           *int64  `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Referer                *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	UserShrink             *string `json:"User,omitempty" xml:"User,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WatermarkShrink        *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GetWebofficeURLShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLShrinkRequest) SetCachePreview(v bool) *GetWebofficeURLShrinkRequest {
	s.CachePreview = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetCredentialConfigShrink(v string) *GetWebofficeURLShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetExternalUploaded(v bool) *GetWebofficeURLShrinkRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetFilename(v string) *GetWebofficeURLShrinkRequest {
	s.Filename = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetHidecmb(v bool) *GetWebofficeURLShrinkRequest {
	s.Hidecmb = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetNotifyTopicName(v string) *GetWebofficeURLShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetPassword(v string) *GetWebofficeURLShrinkRequest {
	s.Password = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetPermissionShrink(v string) *GetWebofficeURLShrinkRequest {
	s.PermissionShrink = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetPreviewPages(v int64) *GetWebofficeURLShrinkRequest {
	s.PreviewPages = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetProjectName(v string) *GetWebofficeURLShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetReferer(v string) *GetWebofficeURLShrinkRequest {
	s.Referer = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetSourceURI(v string) *GetWebofficeURLShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetUserShrink(v string) *GetWebofficeURLShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetUserData(v string) *GetWebofficeURLShrinkRequest {
	s.UserData = &v
	return s
}

func (s *GetWebofficeURLShrinkRequest) SetWatermarkShrink(v string) *GetWebofficeURLShrinkRequest {
	s.WatermarkShrink = &v
	return s
}

type GetWebofficeURLResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebofficeURL            *string `json:"WebofficeURL,omitempty" xml:"WebofficeURL,omitempty"`
}

func (s GetWebofficeURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLResponseBody) SetAccessToken(v string) *GetWebofficeURLResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetAccessTokenExpiredTime(v string) *GetWebofficeURLResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRefreshToken(v string) *GetWebofficeURLResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRefreshTokenExpiredTime(v string) *GetWebofficeURLResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetRequestId(v string) *GetWebofficeURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebofficeURLResponseBody) SetWebofficeURL(v string) *GetWebofficeURLResponseBody {
	s.WebofficeURL = &v
	return s
}

type GetWebofficeURLResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWebofficeURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebofficeURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLResponse) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLResponse) SetHeaders(v map[string]*string) *GetWebofficeURLResponse {
	s.Headers = v
	return s
}

func (s *GetWebofficeURLResponse) SetStatusCode(v int32) *GetWebofficeURLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebofficeURLResponse) SetBody(v *GetWebofficeURLResponseBody) *GetWebofficeURLResponse {
	s.Body = v
	return s
}

type IndexFileMetaRequest struct {
	DatasetName     *string     `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	File            *FileForReq `json:"File,omitempty" xml:"File,omitempty"`
	NotifyTopicName *string     `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string     `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s IndexFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *IndexFileMetaRequest) SetDatasetName(v string) *IndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *IndexFileMetaRequest) SetFile(v *FileForReq) *IndexFileMetaRequest {
	s.File = v
	return s
}

func (s *IndexFileMetaRequest) SetNotifyTopicName(v string) *IndexFileMetaRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *IndexFileMetaRequest) SetProjectName(v string) *IndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

type IndexFileMetaShrinkRequest struct {
	DatasetName     *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FileShrink      *string `json:"File,omitempty" xml:"File,omitempty"`
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s IndexFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *IndexFileMetaShrinkRequest) SetDatasetName(v string) *IndexFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetFileShrink(v string) *IndexFileMetaShrinkRequest {
	s.FileShrink = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetNotifyTopicName(v string) *IndexFileMetaShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetProjectName(v string) *IndexFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

type IndexFileMetaResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IndexFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *IndexFileMetaResponseBody) SetEventId(v string) *IndexFileMetaResponseBody {
	s.EventId = &v
	return s
}

func (s *IndexFileMetaResponseBody) SetRequestId(v string) *IndexFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type IndexFileMetaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndexFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndexFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaResponse) GoString() string {
	return s.String()
}

func (s *IndexFileMetaResponse) SetHeaders(v map[string]*string) *IndexFileMetaResponse {
	s.Headers = v
	return s
}

func (s *IndexFileMetaResponse) SetStatusCode(v int32) *IndexFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *IndexFileMetaResponse) SetBody(v *IndexFileMetaResponseBody) *IndexFileMetaResponse {
	s.Body = v
	return s
}

type ListBindingsRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListBindingsRequest) SetDatasetName(v string) *ListBindingsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListBindingsRequest) SetMaxResults(v int64) *ListBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsRequest) SetNextToken(v string) *ListBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBindingsRequest) SetProjectName(v string) *ListBindingsRequest {
	s.ProjectName = &v
	return s
}

type ListBindingsResponseBody struct {
	Bindings  []*Binding `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	NextToken *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) SetBindings(v []*Binding) *ListBindingsResponseBody {
	s.Bindings = v
	return s
}

func (s *ListBindingsResponseBody) SetNextToken(v string) *ListBindingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

type ListBindingsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListBindingsResponse) SetHeaders(v map[string]*string) *ListBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListBindingsResponse) SetStatusCode(v int32) *ListBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindingsResponse) SetBody(v *ListBindingsResponseBody) *ListBindingsResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Prefix      *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) SetMaxResults(v int64) *ListDatasetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetsRequest) SetNextToken(v string) *ListDatasetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsRequest) SetPrefix(v string) *ListDatasetsRequest {
	s.Prefix = &v
	return s
}

func (s *ListDatasetsRequest) SetProjectName(v string) *ListDatasetsRequest {
	s.ProjectName = &v
	return s
}

type ListDatasetsResponseBody struct {
	Datasets  []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	NextToken *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) SetDatasets(v []*Dataset) *ListDatasetsResponseBody {
	s.Datasets = v
	return s
}

func (s *ListDatasetsResponseBody) SetNextToken(v string) *ListDatasetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

type ListDatasetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatasetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponse) SetHeaders(v map[string]*string) *ListDatasetsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetsResponse) SetStatusCode(v int32) *ListDatasetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Prefix     *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetMaxResults(v int64) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetPrefix(v string) *ListProjectsRequest {
	s.Prefix = &v
	return s
}

type ListProjectsResponseBody struct {
	NextToken *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Projects  []*Project `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*Project) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) SetAcceptLanguage(v string) *ListRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type ListRegionsResponseBody struct {
	Regions   []*RegionType `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*RegionType) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	EndTimeRange   *TimeRange `json:"EndTimeRange,omitempty" xml:"EndTimeRange,omitempty"`
	MaxResults     *int64     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order          *string    `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName    *string    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort           *string    `json:"Sort,omitempty" xml:"Sort,omitempty"`
	StartTimeRange *TimeRange `json:"StartTimeRange,omitempty" xml:"StartTimeRange,omitempty"`
	Status         *string    `json:"Status,omitempty" xml:"Status,omitempty"`
	TagSelector    *string    `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
	TaskTypes      []*string  `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetEndTimeRange(v *TimeRange) *ListTasksRequest {
	s.EndTimeRange = v
	return s
}

func (s *ListTasksRequest) SetMaxResults(v int64) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetOrder(v string) *ListTasksRequest {
	s.Order = &v
	return s
}

func (s *ListTasksRequest) SetProjectName(v string) *ListTasksRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTasksRequest) SetSort(v string) *ListTasksRequest {
	s.Sort = &v
	return s
}

func (s *ListTasksRequest) SetStartTimeRange(v *TimeRange) *ListTasksRequest {
	s.StartTimeRange = v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTagSelector(v string) *ListTasksRequest {
	s.TagSelector = &v
	return s
}

func (s *ListTasksRequest) SetTaskTypes(v []*string) *ListTasksRequest {
	s.TaskTypes = v
	return s
}

type ListTasksShrinkRequest struct {
	EndTimeRangeShrink   *string `json:"EndTimeRange,omitempty" xml:"EndTimeRange,omitempty"`
	MaxResults           *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort                 *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	StartTimeRangeShrink *string `json:"StartTimeRange,omitempty" xml:"StartTimeRange,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagSelector          *string `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
	TaskTypesShrink      *string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) SetEndTimeRangeShrink(v string) *ListTasksShrinkRequest {
	s.EndTimeRangeShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetMaxResults(v int64) *ListTasksShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksShrinkRequest) SetNextToken(v string) *ListTasksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksShrinkRequest) SetOrder(v string) *ListTasksShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectName(v string) *ListTasksShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTasksShrinkRequest) SetSort(v string) *ListTasksShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStartTimeRangeShrink(v string) *ListTasksShrinkRequest {
	s.StartTimeRangeShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStatus(v string) *ListTasksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTagSelector(v string) *ListTasksShrinkRequest {
	s.TagSelector = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskTypesShrink(v string) *ListTasksShrinkRequest {
	s.TaskTypesShrink = &v
	return s
}

type ListTasksResponseBody struct {
	MaxResults  *string     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProjectName *string     `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId   *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks       []*TaskInfo `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetMaxResults(v string) *ListTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetProjectName(v string) *ListTasksResponseBody {
	s.ProjectName = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*TaskInfo) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type MergeFigureClustersRequest struct {
	ClusterIdFrom       *string `json:"ClusterIdFrom,omitempty" xml:"ClusterIdFrom,omitempty"`
	ClusterIdTo         *string `json:"ClusterIdTo,omitempty" xml:"ClusterIdTo,omitempty"`
	CustomMessage       *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	DatasetName         *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureType          *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
	NotifyTopicEndpoint *string `json:"NotifyTopicEndpoint,omitempty" xml:"NotifyTopicEndpoint,omitempty"`
	NotifyTopicName     *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName         *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s MergeFigureClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeFigureClustersRequest) GoString() string {
	return s.String()
}

func (s *MergeFigureClustersRequest) SetClusterIdFrom(v string) *MergeFigureClustersRequest {
	s.ClusterIdFrom = &v
	return s
}

func (s *MergeFigureClustersRequest) SetClusterIdTo(v string) *MergeFigureClustersRequest {
	s.ClusterIdTo = &v
	return s
}

func (s *MergeFigureClustersRequest) SetCustomMessage(v string) *MergeFigureClustersRequest {
	s.CustomMessage = &v
	return s
}

func (s *MergeFigureClustersRequest) SetDatasetName(v string) *MergeFigureClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *MergeFigureClustersRequest) SetFigureType(v string) *MergeFigureClustersRequest {
	s.FigureType = &v
	return s
}

func (s *MergeFigureClustersRequest) SetNotifyTopicEndpoint(v string) *MergeFigureClustersRequest {
	s.NotifyTopicEndpoint = &v
	return s
}

func (s *MergeFigureClustersRequest) SetNotifyTopicName(v string) *MergeFigureClustersRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *MergeFigureClustersRequest) SetProjectName(v string) *MergeFigureClustersRequest {
	s.ProjectName = &v
	return s
}

type MergeFigureClustersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s MergeFigureClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeFigureClustersResponseBody) GoString() string {
	return s.String()
}

func (s *MergeFigureClustersResponseBody) SetRequestId(v string) *MergeFigureClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeFigureClustersResponseBody) SetTaskId(v string) *MergeFigureClustersResponseBody {
	s.TaskId = &v
	return s
}

type MergeFigureClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MergeFigureClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeFigureClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeFigureClustersResponse) GoString() string {
	return s.String()
}

func (s *MergeFigureClustersResponse) SetHeaders(v map[string]*string) *MergeFigureClustersResponse {
	s.Headers = v
	return s
}

func (s *MergeFigureClustersResponse) SetStatusCode(v int32) *MergeFigureClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeFigureClustersResponse) SetBody(v *MergeFigureClustersResponseBody) *MergeFigureClustersResponse {
	s.Body = v
	return s
}

type QueryFigureClustersRequest struct {
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults   *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort         *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s QueryFigureClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFigureClustersRequest) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersRequest) SetCustomLabels(v string) *QueryFigureClustersRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryFigureClustersRequest) SetDatasetName(v string) *QueryFigureClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryFigureClustersRequest) SetMaxResults(v int64) *QueryFigureClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryFigureClustersRequest) SetNextToken(v string) *QueryFigureClustersRequest {
	s.NextToken = &v
	return s
}

func (s *QueryFigureClustersRequest) SetOrder(v string) *QueryFigureClustersRequest {
	s.Order = &v
	return s
}

func (s *QueryFigureClustersRequest) SetProjectName(v string) *QueryFigureClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryFigureClustersRequest) SetSort(v string) *QueryFigureClustersRequest {
	s.Sort = &v
	return s
}

type QueryFigureClustersResponseBody struct {
	FigureClusters []*FigureCluster `json:"FigureClusters,omitempty" xml:"FigureClusters,omitempty" type:"Repeated"`
	NextToken      *string          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFigureClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFigureClustersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersResponseBody) SetFigureClusters(v []*FigureCluster) *QueryFigureClustersResponseBody {
	s.FigureClusters = v
	return s
}

func (s *QueryFigureClustersResponseBody) SetNextToken(v string) *QueryFigureClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryFigureClustersResponseBody) SetRequestId(v string) *QueryFigureClustersResponseBody {
	s.RequestId = &v
	return s
}

type QueryFigureClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFigureClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFigureClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFigureClustersResponse) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersResponse) SetHeaders(v map[string]*string) *QueryFigureClustersResponse {
	s.Headers = v
	return s
}

func (s *QueryFigureClustersResponse) SetStatusCode(v int32) *QueryFigureClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFigureClustersResponse) SetBody(v *QueryFigureClustersResponseBody) *QueryFigureClustersResponse {
	s.Body = v
	return s
}

type QueryStoriesRequest struct {
	CreateTimeRange     *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	CustomLabels        *string    `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName         *string    `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureClusterIds    []*string  `json:"FigureClusterIds,omitempty" xml:"FigureClusterIds,omitempty" type:"Repeated"`
	MaxResults          *int64     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ObjectId            *string    `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Order               *string    `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName         *string    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort                *string    `json:"Sort,omitempty" xml:"Sort,omitempty"`
	StoryEndTimeRange   *TimeRange `json:"StoryEndTimeRange,omitempty" xml:"StoryEndTimeRange,omitempty"`
	StoryName           *string    `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StoryStartTimeRange *TimeRange `json:"StoryStartTimeRange,omitempty" xml:"StoryStartTimeRange,omitempty"`
	StorySubType        *string    `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType           *string    `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	WithEmptyStories    *bool      `json:"WithEmptyStories,omitempty" xml:"WithEmptyStories,omitempty"`
}

func (s QueryStoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStoriesRequest) GoString() string {
	return s.String()
}

func (s *QueryStoriesRequest) SetCreateTimeRange(v *TimeRange) *QueryStoriesRequest {
	s.CreateTimeRange = v
	return s
}

func (s *QueryStoriesRequest) SetCustomLabels(v string) *QueryStoriesRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryStoriesRequest) SetDatasetName(v string) *QueryStoriesRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryStoriesRequest) SetFigureClusterIds(v []*string) *QueryStoriesRequest {
	s.FigureClusterIds = v
	return s
}

func (s *QueryStoriesRequest) SetMaxResults(v int64) *QueryStoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryStoriesRequest) SetNextToken(v string) *QueryStoriesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryStoriesRequest) SetObjectId(v string) *QueryStoriesRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryStoriesRequest) SetOrder(v string) *QueryStoriesRequest {
	s.Order = &v
	return s
}

func (s *QueryStoriesRequest) SetProjectName(v string) *QueryStoriesRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryStoriesRequest) SetSort(v string) *QueryStoriesRequest {
	s.Sort = &v
	return s
}

func (s *QueryStoriesRequest) SetStoryEndTimeRange(v *TimeRange) *QueryStoriesRequest {
	s.StoryEndTimeRange = v
	return s
}

func (s *QueryStoriesRequest) SetStoryName(v string) *QueryStoriesRequest {
	s.StoryName = &v
	return s
}

func (s *QueryStoriesRequest) SetStoryStartTimeRange(v *TimeRange) *QueryStoriesRequest {
	s.StoryStartTimeRange = v
	return s
}

func (s *QueryStoriesRequest) SetStorySubType(v string) *QueryStoriesRequest {
	s.StorySubType = &v
	return s
}

func (s *QueryStoriesRequest) SetStoryType(v string) *QueryStoriesRequest {
	s.StoryType = &v
	return s
}

func (s *QueryStoriesRequest) SetWithEmptyStories(v bool) *QueryStoriesRequest {
	s.WithEmptyStories = &v
	return s
}

type QueryStoriesShrinkRequest struct {
	CreateTimeRangeShrink     *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	CustomLabels              *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName               *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureClusterIdsShrink    *string `json:"FigureClusterIds,omitempty" xml:"FigureClusterIds,omitempty"`
	MaxResults                *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ObjectId                  *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Order                     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName               *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort                      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	StoryEndTimeRangeShrink   *string `json:"StoryEndTimeRange,omitempty" xml:"StoryEndTimeRange,omitempty"`
	StoryName                 *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StoryStartTimeRangeShrink *string `json:"StoryStartTimeRange,omitempty" xml:"StoryStartTimeRange,omitempty"`
	StorySubType              *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType                 *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	WithEmptyStories          *bool   `json:"WithEmptyStories,omitempty" xml:"WithEmptyStories,omitempty"`
}

func (s QueryStoriesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryStoriesShrinkRequest) SetCreateTimeRangeShrink(v string) *QueryStoriesShrinkRequest {
	s.CreateTimeRangeShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetCustomLabels(v string) *QueryStoriesShrinkRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetDatasetName(v string) *QueryStoriesShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetFigureClusterIdsShrink(v string) *QueryStoriesShrinkRequest {
	s.FigureClusterIdsShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetMaxResults(v int64) *QueryStoriesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetNextToken(v string) *QueryStoriesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetObjectId(v string) *QueryStoriesShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetOrder(v string) *QueryStoriesShrinkRequest {
	s.Order = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetProjectName(v string) *QueryStoriesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetSort(v string) *QueryStoriesShrinkRequest {
	s.Sort = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryEndTimeRangeShrink(v string) *QueryStoriesShrinkRequest {
	s.StoryEndTimeRangeShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryName(v string) *QueryStoriesShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryStartTimeRangeShrink(v string) *QueryStoriesShrinkRequest {
	s.StoryStartTimeRangeShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStorySubType(v string) *QueryStoriesShrinkRequest {
	s.StorySubType = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryType(v string) *QueryStoriesShrinkRequest {
	s.StoryType = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetWithEmptyStories(v bool) *QueryStoriesShrinkRequest {
	s.WithEmptyStories = &v
	return s
}

type QueryStoriesResponseBody struct {
	NextToken *string  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stories   []*Story `json:"Stories,omitempty" xml:"Stories,omitempty" type:"Repeated"`
}

func (s QueryStoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStoriesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStoriesResponseBody) SetNextToken(v string) *QueryStoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryStoriesResponseBody) SetRequestId(v string) *QueryStoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStoriesResponseBody) SetStories(v []*Story) *QueryStoriesResponseBody {
	s.Stories = v
	return s
}

type QueryStoriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStoriesResponse) GoString() string {
	return s.String()
}

func (s *QueryStoriesResponse) SetHeaders(v map[string]*string) *QueryStoriesResponse {
	s.Headers = v
	return s
}

func (s *QueryStoriesResponse) SetStatusCode(v int32) *QueryStoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStoriesResponse) SetBody(v *QueryStoriesResponseBody) *QueryStoriesResponse {
	s.Body = v
	return s
}

type RefreshWebofficeTokenRequest struct {
	AccessToken      *string           `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RefreshToken     *string           `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshWebofficeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenRequest) SetAccessToken(v string) *RefreshWebofficeTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetCredentialConfig(v *CredentialConfig) *RefreshWebofficeTokenRequest {
	s.CredentialConfig = v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetProjectName(v string) *RefreshWebofficeTokenRequest {
	s.ProjectName = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetRefreshToken(v string) *RefreshWebofficeTokenRequest {
	s.RefreshToken = &v
	return s
}

type RefreshWebofficeTokenShrinkRequest struct {
	AccessToken            *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RefreshToken           *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshWebofficeTokenShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenShrinkRequest) SetAccessToken(v string) *RefreshWebofficeTokenShrinkRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetCredentialConfigShrink(v string) *RefreshWebofficeTokenShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetProjectName(v string) *RefreshWebofficeTokenShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetRefreshToken(v string) *RefreshWebofficeTokenShrinkRequest {
	s.RefreshToken = &v
	return s
}

type RefreshWebofficeTokenResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshWebofficeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessToken(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshToken(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRequestId(v string) *RefreshWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshWebofficeTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshWebofficeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponse) SetHeaders(v map[string]*string) *RefreshWebofficeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetStatusCode(v int32) *RefreshWebofficeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshWebofficeTokenResponse) SetBody(v *RefreshWebofficeTokenResponseBody) *RefreshWebofficeTokenResponse {
	s.Body = v
	return s
}

type RemoveStoryFilesRequest struct {
	DatasetName *string                         `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files       []*RemoveStoryFilesRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ObjectId    *string                         `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string                         `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s RemoveStoryFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequest) SetDatasetName(v string) *RemoveStoryFilesRequest {
	s.DatasetName = &v
	return s
}

func (s *RemoveStoryFilesRequest) SetFiles(v []*RemoveStoryFilesRequestFiles) *RemoveStoryFilesRequest {
	s.Files = v
	return s
}

func (s *RemoveStoryFilesRequest) SetObjectId(v string) *RemoveStoryFilesRequest {
	s.ObjectId = &v
	return s
}

func (s *RemoveStoryFilesRequest) SetProjectName(v string) *RemoveStoryFilesRequest {
	s.ProjectName = &v
	return s
}

type RemoveStoryFilesRequestFiles struct {
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s RemoveStoryFilesRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequestFiles) SetURI(v string) *RemoveStoryFilesRequestFiles {
	s.URI = &v
	return s
}

type RemoveStoryFilesShrinkRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s RemoveStoryFilesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesShrinkRequest) SetDatasetName(v string) *RemoveStoryFilesShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) SetFilesShrink(v string) *RemoveStoryFilesShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) SetObjectId(v string) *RemoveStoryFilesShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *RemoveStoryFilesShrinkRequest) SetProjectName(v string) *RemoveStoryFilesShrinkRequest {
	s.ProjectName = &v
	return s
}

type RemoveStoryFilesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveStoryFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponseBody) SetRequestId(v string) *RemoveStoryFilesResponseBody {
	s.RequestId = &v
	return s
}

type RemoveStoryFilesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveStoryFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveStoryFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesResponse) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponse) SetHeaders(v map[string]*string) *RemoveStoryFilesResponse {
	s.Headers = v
	return s
}

func (s *RemoveStoryFilesResponse) SetStatusCode(v int32) *RemoveStoryFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveStoryFilesResponse) SetBody(v *RemoveStoryFilesResponseBody) *RemoveStoryFilesResponse {
	s.Body = v
	return s
}

type ResumeBindingRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ResumeBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeBindingRequest) GoString() string {
	return s.String()
}

func (s *ResumeBindingRequest) SetDatasetName(v string) *ResumeBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *ResumeBindingRequest) SetProjectName(v string) *ResumeBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *ResumeBindingRequest) SetURI(v string) *ResumeBindingRequest {
	s.URI = &v
	return s
}

type ResumeBindingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeBindingResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeBindingResponseBody) SetRequestId(v string) *ResumeBindingResponseBody {
	s.RequestId = &v
	return s
}

type ResumeBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeBindingResponse) GoString() string {
	return s.String()
}

func (s *ResumeBindingResponse) SetHeaders(v map[string]*string) *ResumeBindingResponse {
	s.Headers = v
	return s
}

func (s *ResumeBindingResponse) SetStatusCode(v int32) *ResumeBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeBindingResponse) SetBody(v *ResumeBindingResponseBody) *ResumeBindingResponse {
	s.Body = v
	return s
}

type SemanticQueryRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s SemanticQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s SemanticQueryRequest) GoString() string {
	return s.String()
}

func (s *SemanticQueryRequest) SetDatasetName(v string) *SemanticQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *SemanticQueryRequest) SetMaxResults(v int32) *SemanticQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *SemanticQueryRequest) SetNextToken(v string) *SemanticQueryRequest {
	s.NextToken = &v
	return s
}

func (s *SemanticQueryRequest) SetProjectName(v string) *SemanticQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *SemanticQueryRequest) SetQuery(v string) *SemanticQueryRequest {
	s.Query = &v
	return s
}

type SemanticQueryResponseBody struct {
	Aggregations []*SemanticQueryResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	Files        []*File                                  `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	NextToken    *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SemanticQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SemanticQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SemanticQueryResponseBody) SetAggregations(v []*SemanticQueryResponseBodyAggregations) *SemanticQueryResponseBody {
	s.Aggregations = v
	return s
}

func (s *SemanticQueryResponseBody) SetFiles(v []*File) *SemanticQueryResponseBody {
	s.Files = v
	return s
}

func (s *SemanticQueryResponseBody) SetNextToken(v string) *SemanticQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SemanticQueryResponseBody) SetRequestId(v string) *SemanticQueryResponseBody {
	s.RequestId = &v
	return s
}

type SemanticQueryResponseBodyAggregations struct {
	Field     *string                                        `json:"Field,omitempty" xml:"Field,omitempty"`
	Groups    []*SemanticQueryResponseBodyAggregationsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Operation *string                                        `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Value     *float32                                       `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SemanticQueryResponseBodyAggregations) String() string {
	return tea.Prettify(s)
}

func (s SemanticQueryResponseBodyAggregations) GoString() string {
	return s.String()
}

func (s *SemanticQueryResponseBodyAggregations) SetField(v string) *SemanticQueryResponseBodyAggregations {
	s.Field = &v
	return s
}

func (s *SemanticQueryResponseBodyAggregations) SetGroups(v []*SemanticQueryResponseBodyAggregationsGroups) *SemanticQueryResponseBodyAggregations {
	s.Groups = v
	return s
}

func (s *SemanticQueryResponseBodyAggregations) SetOperation(v string) *SemanticQueryResponseBodyAggregations {
	s.Operation = &v
	return s
}

func (s *SemanticQueryResponseBodyAggregations) SetValue(v float32) *SemanticQueryResponseBodyAggregations {
	s.Value = &v
	return s
}

type SemanticQueryResponseBodyAggregationsGroups struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SemanticQueryResponseBodyAggregationsGroups) String() string {
	return tea.Prettify(s)
}

func (s SemanticQueryResponseBodyAggregationsGroups) GoString() string {
	return s.String()
}

func (s *SemanticQueryResponseBodyAggregationsGroups) SetCount(v int64) *SemanticQueryResponseBodyAggregationsGroups {
	s.Count = &v
	return s
}

func (s *SemanticQueryResponseBodyAggregationsGroups) SetValue(v string) *SemanticQueryResponseBodyAggregationsGroups {
	s.Value = &v
	return s
}

type SemanticQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SemanticQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SemanticQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s SemanticQueryResponse) GoString() string {
	return s.String()
}

func (s *SemanticQueryResponse) SetHeaders(v map[string]*string) *SemanticQueryResponse {
	s.Headers = v
	return s
}

func (s *SemanticQueryResponse) SetStatusCode(v int32) *SemanticQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SemanticQueryResponse) SetBody(v *SemanticQueryResponseBody) *SemanticQueryResponse {
	s.Body = v
	return s
}

type SimpleQueryRequest struct {
	Aggregations []*SimpleQueryRequestAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	DatasetName  *string                           `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults   *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order        *string                           `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName  *string                           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Query        *SimpleQuery                      `json:"Query,omitempty" xml:"Query,omitempty"`
	Sort         *string                           `json:"Sort,omitempty" xml:"Sort,omitempty"`
	WithFields   []*string                         `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
}

func (s SimpleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryRequest) GoString() string {
	return s.String()
}

func (s *SimpleQueryRequest) SetAggregations(v []*SimpleQueryRequestAggregations) *SimpleQueryRequest {
	s.Aggregations = v
	return s
}

func (s *SimpleQueryRequest) SetDatasetName(v string) *SimpleQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *SimpleQueryRequest) SetMaxResults(v int32) *SimpleQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *SimpleQueryRequest) SetNextToken(v string) *SimpleQueryRequest {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryRequest) SetOrder(v string) *SimpleQueryRequest {
	s.Order = &v
	return s
}

func (s *SimpleQueryRequest) SetProjectName(v string) *SimpleQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *SimpleQueryRequest) SetQuery(v *SimpleQuery) *SimpleQueryRequest {
	s.Query = v
	return s
}

func (s *SimpleQueryRequest) SetSort(v string) *SimpleQueryRequest {
	s.Sort = &v
	return s
}

func (s *SimpleQueryRequest) SetWithFields(v []*string) *SimpleQueryRequest {
	s.WithFields = v
	return s
}

type SimpleQueryRequestAggregations struct {
	Field     *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s SimpleQueryRequestAggregations) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryRequestAggregations) GoString() string {
	return s.String()
}

func (s *SimpleQueryRequestAggregations) SetField(v string) *SimpleQueryRequestAggregations {
	s.Field = &v
	return s
}

func (s *SimpleQueryRequestAggregations) SetOperation(v string) *SimpleQueryRequestAggregations {
	s.Operation = &v
	return s
}

type SimpleQueryShrinkRequest struct {
	AggregationsShrink *string `json:"Aggregations,omitempty" xml:"Aggregations,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults         *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order              *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	QueryShrink        *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Sort               *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	WithFieldsShrink   *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
}

func (s SimpleQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *SimpleQueryShrinkRequest) SetAggregationsShrink(v string) *SimpleQueryShrinkRequest {
	s.AggregationsShrink = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetDatasetName(v string) *SimpleQueryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetMaxResults(v int32) *SimpleQueryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetNextToken(v string) *SimpleQueryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetOrder(v string) *SimpleQueryShrinkRequest {
	s.Order = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetProjectName(v string) *SimpleQueryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetQueryShrink(v string) *SimpleQueryShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetSort(v string) *SimpleQueryShrinkRequest {
	s.Sort = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetWithFieldsShrink(v string) *SimpleQueryShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

type SimpleQueryResponseBody struct {
	Aggregations []*SimpleQueryResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	Files        []*File                                `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	NextToken    *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SimpleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBody) SetAggregations(v []*SimpleQueryResponseBodyAggregations) *SimpleQueryResponseBody {
	s.Aggregations = v
	return s
}

func (s *SimpleQueryResponseBody) SetFiles(v []*File) *SimpleQueryResponseBody {
	s.Files = v
	return s
}

func (s *SimpleQueryResponseBody) SetNextToken(v string) *SimpleQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryResponseBody) SetRequestId(v string) *SimpleQueryResponseBody {
	s.RequestId = &v
	return s
}

type SimpleQueryResponseBodyAggregations struct {
	Field     *string                                      `json:"Field,omitempty" xml:"Field,omitempty"`
	Groups    []*SimpleQueryResponseBodyAggregationsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Operation *string                                      `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Value     *float64                                     `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQueryResponseBodyAggregations) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryResponseBodyAggregations) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBodyAggregations) SetField(v string) *SimpleQueryResponseBodyAggregations {
	s.Field = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetGroups(v []*SimpleQueryResponseBodyAggregationsGroups) *SimpleQueryResponseBodyAggregations {
	s.Groups = v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetOperation(v string) *SimpleQueryResponseBodyAggregations {
	s.Operation = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetValue(v float64) *SimpleQueryResponseBodyAggregations {
	s.Value = &v
	return s
}

type SimpleQueryResponseBodyAggregationsGroups struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQueryResponseBodyAggregationsGroups) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryResponseBodyAggregationsGroups) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBodyAggregationsGroups) SetCount(v int64) *SimpleQueryResponseBodyAggregationsGroups {
	s.Count = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregationsGroups) SetValue(v string) *SimpleQueryResponseBodyAggregationsGroups {
	s.Value = &v
	return s
}

type SimpleQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SimpleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SimpleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryResponse) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponse) SetHeaders(v map[string]*string) *SimpleQueryResponse {
	s.Headers = v
	return s
}

func (s *SimpleQueryResponse) SetStatusCode(v int32) *SimpleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SimpleQueryResponse) SetBody(v *SimpleQueryResponseBody) *SimpleQueryResponse {
	s.Body = v
	return s
}

type StopBindingRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s StopBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s StopBindingRequest) GoString() string {
	return s.String()
}

func (s *StopBindingRequest) SetDatasetName(v string) *StopBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *StopBindingRequest) SetProjectName(v string) *StopBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *StopBindingRequest) SetReason(v string) *StopBindingRequest {
	s.Reason = &v
	return s
}

func (s *StopBindingRequest) SetURI(v string) *StopBindingRequest {
	s.URI = &v
	return s
}

type StopBindingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBindingResponseBody) GoString() string {
	return s.String()
}

func (s *StopBindingResponseBody) SetRequestId(v string) *StopBindingResponseBody {
	s.RequestId = &v
	return s
}

type StopBindingResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s StopBindingResponse) GoString() string {
	return s.String()
}

func (s *StopBindingResponse) SetHeaders(v map[string]*string) *StopBindingResponse {
	s.Headers = v
	return s
}

func (s *StopBindingResponse) SetStatusCode(v int32) *StopBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBindingResponse) SetBody(v *StopBindingResponseBody) *StopBindingResponse {
	s.Body = v
	return s
}

type UpdateDatasetRequest struct {
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) SetDatasetMaxBindCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxEntityCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxFileCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxRelationCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateDatasetRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetName(v string) *UpdateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) SetProjectName(v string) *UpdateDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateDatasetRequest) SetTemplateId(v string) *UpdateDatasetRequest {
	s.TemplateId = &v
	return s
}

type UpdateDatasetResponseBody struct {
	Dataset   *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) SetDataset(v *Dataset) *UpdateDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponse) SetHeaders(v map[string]*string) *UpdateDatasetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetResponse) SetStatusCode(v int32) *UpdateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

type UpdateFigureClusterRequest struct {
	DatasetName   *string              `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureCluster *FigureClusterForReq `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	ProjectName   *string              `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFigureClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterRequest) SetDatasetName(v string) *UpdateFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFigureClusterRequest) SetFigureCluster(v *FigureClusterForReq) *UpdateFigureClusterRequest {
	s.FigureCluster = v
	return s
}

func (s *UpdateFigureClusterRequest) SetProjectName(v string) *UpdateFigureClusterRequest {
	s.ProjectName = &v
	return s
}

type UpdateFigureClusterShrinkRequest struct {
	DatasetName         *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureClusterShrink *string `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	ProjectName         *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFigureClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFigureClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterShrinkRequest) SetDatasetName(v string) *UpdateFigureClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFigureClusterShrinkRequest) SetFigureClusterShrink(v string) *UpdateFigureClusterShrinkRequest {
	s.FigureClusterShrink = &v
	return s
}

func (s *UpdateFigureClusterShrinkRequest) SetProjectName(v string) *UpdateFigureClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

type UpdateFigureClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFigureClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterResponseBody) SetRequestId(v string) *UpdateFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFigureClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFigureClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterResponse) SetHeaders(v map[string]*string) *UpdateFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateFigureClusterResponse) SetStatusCode(v int32) *UpdateFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFigureClusterResponse) SetBody(v *UpdateFigureClusterResponseBody) *UpdateFigureClusterResponse {
	s.Body = v
	return s
}

type UpdateFileMetaRequest struct {
	DatasetName *string     `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	File        *FileForReq `json:"File,omitempty" xml:"File,omitempty"`
	ProjectName *string     `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaRequest) SetDatasetName(v string) *UpdateFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFileMetaRequest) SetFile(v *FileForReq) *UpdateFileMetaRequest {
	s.File = v
	return s
}

func (s *UpdateFileMetaRequest) SetProjectName(v string) *UpdateFileMetaRequest {
	s.ProjectName = &v
	return s
}

type UpdateFileMetaShrinkRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FileShrink  *string `json:"File,omitempty" xml:"File,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaShrinkRequest) SetDatasetName(v string) *UpdateFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetFileShrink(v string) *UpdateFileMetaShrinkRequest {
	s.FileShrink = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetProjectName(v string) *UpdateFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

type UpdateFileMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaResponseBody) SetRequestId(v string) *UpdateFileMetaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFileMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFileMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaResponse) SetHeaders(v map[string]*string) *UpdateFileMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileMetaResponse) SetStatusCode(v int32) *UpdateFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileMetaResponse) SetBody(v *UpdateFileMetaResponseBody) *UpdateFileMetaResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineConcurrency       *int64  `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectQueriesPerSecond *int64  `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	ServiceRole             *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDatasetMaxBindCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxEntityCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxFileCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxRelationCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateProjectRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetEngineConcurrency(v int64) *UpdateProjectRequest {
	s.EngineConcurrency = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectMaxDatasetCount(v int64) *UpdateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectQueriesPerSecond(v int64) *UpdateProjectRequest {
	s.ProjectQueriesPerSecond = &v
	return s
}

func (s *UpdateProjectRequest) SetServiceRole(v string) *UpdateProjectRequest {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectRequest) SetTemplateId(v string) *UpdateProjectRequest {
	s.TemplateId = &v
	return s
}

type UpdateProjectResponseBody struct {
	Project   *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetProject(v *Project) *UpdateProjectResponseBody {
	s.Project = v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateStoryRequest struct {
	Cover        *UpdateStoryRequestCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	CustomId     *string                  `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{}   `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string                  `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId     *string                  `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName  *string                  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryName    *string                  `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
}

func (s UpdateStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequest) SetCover(v *UpdateStoryRequestCover) *UpdateStoryRequest {
	s.Cover = v
	return s
}

func (s *UpdateStoryRequest) SetCustomId(v string) *UpdateStoryRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateStoryRequest) SetCustomLabels(v map[string]interface{}) *UpdateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *UpdateStoryRequest) SetDatasetName(v string) *UpdateStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateStoryRequest) SetObjectId(v string) *UpdateStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateStoryRequest) SetProjectName(v string) *UpdateStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateStoryRequest) SetStoryName(v string) *UpdateStoryRequest {
	s.StoryName = &v
	return s
}

type UpdateStoryRequestCover struct {
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s UpdateStoryRequestCover) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryRequestCover) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequestCover) SetURI(v string) *UpdateStoryRequestCover {
	s.URI = &v
	return s
}

type UpdateStoryShrinkRequest struct {
	CoverShrink        *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	CustomId           *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId           *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryName          *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
}

func (s UpdateStoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoryShrinkRequest) SetCoverShrink(v string) *UpdateStoryShrinkRequest {
	s.CoverShrink = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetCustomId(v string) *UpdateStoryShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetCustomLabelsShrink(v string) *UpdateStoryShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetDatasetName(v string) *UpdateStoryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetObjectId(v string) *UpdateStoryShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetProjectName(v string) *UpdateStoryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateStoryShrinkRequest) SetStoryName(v string) *UpdateStoryShrinkRequest {
	s.StoryName = &v
	return s
}

type UpdateStoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponseBody) SetRequestId(v string) *UpdateStoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponse) SetHeaders(v map[string]*string) *UpdateStoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoryResponse) SetStatusCode(v int32) *UpdateStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStoryResponse) SetBody(v *UpdateStoryResponseBody) *UpdateStoryResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing-gov-1": tea.String("imm-vpc.cn-beijing-gov-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddStoryFilesWithOptions(tmpReq *AddStoryFilesRequest, runtime *util.RuntimeOptions) (_result *AddStoryFilesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddStoryFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Files)) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, tea.String("Files"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FilesShrink)) {
		body["Files"] = request.FilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddStoryFiles"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddStoryFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddStoryFiles(request *AddStoryFilesRequest) (_result *AddStoryFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddStoryFilesResponse{}
	_body, _err := client.AddStoryFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachOSSBucketWithOptions(request *AttachOSSBucketRequest, runtime *util.RuntimeOptions) (_result *AttachOSSBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OSSBucket)) {
		query["OSSBucket"] = request.OSSBucket
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachOSSBucket"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachOSSBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachOSSBucket(request *AttachOSSBucketRequest) (_result *AttachOSSBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachOSSBucketResponse{}
	_body, _err := client.AttachOSSBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteFileMetaWithOptions(tmpReq *BatchDeleteFileMetaRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchDeleteFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.URIs)) {
		request.URIsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.URIs, tea.String("URIs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URIsShrink)) {
		query["URIs"] = request.URIsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteFileMeta(request *BatchDeleteFileMetaRequest) (_result *BatchDeleteFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteFileMetaResponse{}
	_body, _err := client.BatchDeleteFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetFileMetaWithOptions(tmpReq *BatchGetFileMetaRequest, runtime *util.RuntimeOptions) (_result *BatchGetFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchGetFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.URIs)) {
		request.URIsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.URIs, tea.String("URIs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URIsShrink)) {
		query["URIs"] = request.URIsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetFileMeta(request *BatchGetFileMetaRequest) (_result *BatchGetFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchGetFileMetaResponse{}
	_body, _err := client.BatchGetFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchIndexFileMetaWithOptions(tmpReq *BatchIndexFileMetaRequest, runtime *util.RuntimeOptions) (_result *BatchIndexFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchIndexFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Files)) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, tea.String("Files"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FilesShrink)) {
		query["Files"] = request.FilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchIndexFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchIndexFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchIndexFileMeta(request *BatchIndexFileMetaRequest) (_result *BatchIndexFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchIndexFileMetaResponse{}
	_body, _err := client.BatchIndexFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateFileMetaWithOptions(tmpReq *BatchUpdateFileMetaRequest, runtime *util.RuntimeOptions) (_result *BatchUpdateFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchUpdateFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Files)) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, tea.String("Files"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FilesShrink)) {
		query["Files"] = request.FilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateFileMeta(request *BatchUpdateFileMetaRequest) (_result *BatchUpdateFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUpdateFileMetaResponse{}
	_body, _err := client.BatchUpdateFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBindingWithOptions(request *CreateBindingRequest, runtime *util.RuntimeOptions) (_result *CreateBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBinding"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBinding(request *CreateBindingRequest) (_result *CreateBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBindingResponse{}
	_body, _err := client.CreateBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCompressPointCloudTaskWithOptions(tmpReq *CreateCompressPointCloudTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCompressPointCloudTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCompressPointCloudTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.KdtreeOption))) {
		request.KdtreeOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.KdtreeOption), tea.String("KdtreeOption"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OctreeOption))) {
		request.OctreeOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OctreeOption), tea.String("OctreeOption"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PointCloudFields)) {
		request.PointCloudFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PointCloudFields, tea.String("PointCloudFields"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompressMethod)) {
		query["CompressMethod"] = request.CompressMethod
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.KdtreeOptionShrink)) {
		query["KdtreeOption"] = request.KdtreeOptionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.OctreeOptionShrink)) {
		query["OctreeOption"] = request.OctreeOptionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PointCloudFieldsShrink)) {
		query["PointCloudFields"] = request.PointCloudFieldsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PointCloudFileFormat)) {
		query["PointCloudFileFormat"] = request.PointCloudFileFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TargetURI)) {
		query["TargetURI"] = request.TargetURI
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCompressPointCloudTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCompressPointCloudTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCompressPointCloudTask(request *CreateCompressPointCloudTaskRequest) (_result *CreateCompressPointCloudTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCompressPointCloudTaskResponse{}
	_body, _err := client.CreateCompressPointCloudTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, runtime *util.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetMaxBindCount)) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxEntityCount)) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxFileCount)) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxRelationCount)) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxTotalFileSize)) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataset"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDetectVideoLabelsTaskWithOptions(tmpReq *CreateDetectVideoLabelsTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDetectVideoLabelsTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDetectVideoLabelsTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDetectVideoLabelsTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDetectVideoLabelsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDetectVideoLabelsTask(request *CreateDetectVideoLabelsTaskRequest) (_result *CreateDetectVideoLabelsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDetectVideoLabelsTaskResponse{}
	_body, _err := client.CreateDetectVideoLabelsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFigureClusteringTaskWithOptions(tmpReq *CreateFigureClusteringTaskRequest, runtime *util.RuntimeOptions) (_result *CreateFigureClusteringTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFigureClusteringTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFigureClusteringTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFigureClusteringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFigureClusteringTask(request *CreateFigureClusteringTaskRequest) (_result *CreateFigureClusteringTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFigureClusteringTaskResponse{}
	_body, _err := client.CreateFigureClusteringTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFigureClustersMergingTaskWithOptions(tmpReq *CreateFigureClustersMergingTaskRequest, runtime *util.RuntimeOptions) (_result *CreateFigureClustersMergingTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFigureClustersMergingTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFigureClustersMergingTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFigureClustersMergingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFigureClustersMergingTask(request *CreateFigureClustersMergingTaskRequest) (_result *CreateFigureClustersMergingTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFigureClustersMergingTaskResponse{}
	_body, _err := client.CreateFigureClustersMergingTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageModerationTaskWithOptions(tmpReq *CreateImageModerationTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageModerationTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateImageModerationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Scenes)) {
		request.ScenesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scenes, tea.String("Scenes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.MaxFrames)) {
		query["MaxFrames"] = request.MaxFrames
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Reviewer)) {
		query["Reviewer"] = request.Reviewer
	}

	if !tea.BoolValue(util.IsUnset(request.ScenesShrink)) {
		query["Scenes"] = request.ScenesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageModerationTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageModerationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageModerationTask(request *CreateImageModerationTaskRequest) (_result *CreateImageModerationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageModerationTaskResponse{}
	_body, _err := client.CreateImageModerationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageSplicingTaskWithOptions(tmpReq *CreateImageSplicingTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageSplicingTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateImageSplicingTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Align)) {
		query["Align"] = request.Align
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundColor)) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.ImageFormat)) {
		query["ImageFormat"] = request.ImageFormat
	}

	if !tea.BoolValue(util.IsUnset(request.Margin)) {
		query["Margin"] = request.Margin
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Padding)) {
		query["Padding"] = request.Padding
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Quality)) {
		query["Quality"] = request.Quality
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		query["ScaleType"] = request.ScaleType
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TargetURI)) {
		query["TargetURI"] = request.TargetURI
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageSplicingTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageSplicingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageSplicingTask(request *CreateImageSplicingTaskRequest) (_result *CreateImageSplicingTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageSplicingTaskResponse{}
	_body, _err := client.CreateImageSplicingTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMediaConvertTaskWithOptions(tmpReq *CreateMediaConvertTaskRequest, runtime *util.RuntimeOptions) (_result *CreateMediaConvertTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateMediaConvertTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Targets)) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, tea.String("Targets"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TargetsShrink)) {
		query["Targets"] = request.TargetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMediaConvertTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMediaConvertTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMediaConvertTask(request *CreateMediaConvertTaskRequest) (_result *CreateMediaConvertTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMediaConvertTaskResponse{}
	_body, _err := client.CreateMediaConvertTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOfficeConversionTaskWithOptions(tmpReq *CreateOfficeConversionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateOfficeConversionTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateOfficeConversionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.TrimPolicy))) {
		request.TrimPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.TrimPolicy), tea.String("TrimPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndPage)) {
		query["EndPage"] = request.EndPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirstPage)) {
		query["FirstPage"] = request.FirstPage
	}

	if !tea.BoolValue(util.IsUnset(request.FitToHeight)) {
		query["FitToHeight"] = request.FitToHeight
	}

	if !tea.BoolValue(util.IsUnset(request.FitToWidth)) {
		query["FitToWidth"] = request.FitToWidth
	}

	if !tea.BoolValue(util.IsUnset(request.HoldLineFeed)) {
		query["HoldLineFeed"] = request.HoldLineFeed
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDPI)) {
		query["ImageDPI"] = request.ImageDPI
	}

	if !tea.BoolValue(util.IsUnset(request.LongPicture)) {
		query["LongPicture"] = request.LongPicture
	}

	if !tea.BoolValue(util.IsUnset(request.LongText)) {
		query["LongText"] = request.LongText
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetColumn)) {
		query["MaxSheetColumn"] = request.MaxSheetColumn
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSheetRow)) {
		query["MaxSheetRow"] = request.MaxSheetRow
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Pages)) {
		query["Pages"] = request.Pages
	}

	if !tea.BoolValue(util.IsUnset(request.PaperHorizontal)) {
		query["PaperHorizontal"] = request.PaperHorizontal
	}

	if !tea.BoolValue(util.IsUnset(request.PaperSize)) {
		query["PaperSize"] = request.PaperSize
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Quality)) {
		query["Quality"] = request.Quality
	}

	if !tea.BoolValue(util.IsUnset(request.ScalePercentage)) {
		query["ScalePercentage"] = request.ScalePercentage
	}

	if !tea.BoolValue(util.IsUnset(request.SheetCount)) {
		query["SheetCount"] = request.SheetCount
	}

	if !tea.BoolValue(util.IsUnset(request.SheetIndex)) {
		query["SheetIndex"] = request.SheetIndex
	}

	if !tea.BoolValue(util.IsUnset(request.ShowComments)) {
		query["ShowComments"] = request.ShowComments
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.StartPage)) {
		query["StartPage"] = request.StartPage
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetURI)) {
		query["TargetURI"] = request.TargetURI
	}

	if !tea.BoolValue(util.IsUnset(request.TargetURIPrefix)) {
		query["TargetURIPrefix"] = request.TargetURIPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.TrimPolicyShrink)) {
		query["TrimPolicy"] = request.TrimPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOfficeConversionTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOfficeConversionTask(request *CreateOfficeConversionTaskRequest) (_result *CreateOfficeConversionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOfficeConversionTaskResponse{}
	_body, _err := client.CreateOfficeConversionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetMaxBindCount)) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxEntityCount)) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxFileCount)) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxRelationCount)) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxTotalFileSize)) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EngineConcurrency)) {
		query["EngineConcurrency"] = request.EngineConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectMaxDatasetCount)) {
		query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectQueriesPerSecond)) {
		query["ProjectQueriesPerSecond"] = request.ProjectQueriesPerSecond
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRole)) {
		query["ServiceRole"] = request.ServiceRole
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStoryWithOptions(tmpReq *CreateStoryRequest, runtime *util.RuntimeOptions) (_result *CreateStoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomId)) {
		body["CustomId"] = request.CustomId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabelsShrink)) {
		body["CustomLabels"] = request.CustomLabelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxFileCount)) {
		body["MaxFileCount"] = request.MaxFileCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinFileCount)) {
		body["MinFileCount"] = request.MinFileCount
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		body["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryEndTime)) {
		body["StoryEndTime"] = request.StoryEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["StoryName"] = request.StoryName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryStartTime)) {
		body["StoryStartTime"] = request.StoryStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StorySubType)) {
		body["StorySubType"] = request.StorySubType
	}

	if !tea.BoolValue(util.IsUnset(request.StoryType)) {
		body["StoryType"] = request.StoryType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStory(request *CreateStoryRequest) (_result *CreateStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStoryResponse{}
	_body, _err := client.CreateStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoModerationTaskWithOptions(tmpReq *CreateVideoModerationTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoModerationTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateVideoModerationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Scenes)) {
		request.ScenesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scenes, tea.String("Scenes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.MaxFrames)) {
		query["MaxFrames"] = request.MaxFrames
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Reviewer)) {
		query["Reviewer"] = request.Reviewer
	}

	if !tea.BoolValue(util.IsUnset(request.ScenesShrink)) {
		query["Scenes"] = request.ScenesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVideoModerationTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVideoModerationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoModerationTask(request *CreateVideoModerationTaskRequest) (_result *CreateVideoModerationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoModerationTaskResponse{}
	_body, _err := client.CreateVideoModerationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBindingWithOptions(request *DeleteBindingRequest, runtime *util.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cleanup)) {
		query["Cleanup"] = request.Cleanup
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBinding"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBinding(request *DeleteBindingRequest) (_result *DeleteBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DeleteBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatasetWithOptions(request *DeleteDatasetRequest, runtime *util.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataset"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataset(request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileMetaWithOptions(request *DeleteFileMetaRequest, runtime *util.RuntimeOptions) (_result *DeleteFileMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFileMeta(request *DeleteFileMetaRequest) (_result *DeleteFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileMetaResponse{}
	_body, _err := client.DeleteFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStoryWithOptions(request *DeleteStoryRequest, runtime *util.RuntimeOptions) (_result *DeleteStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStory(request *DeleteStoryRequest) (_result *DeleteStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStoryResponse{}
	_body, _err := client.DeleteStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachOSSBucketWithOptions(request *DetachOSSBucketRequest, runtime *util.RuntimeOptions) (_result *DetachOSSBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OSSBucket)) {
		query["OSSBucket"] = request.OSSBucket
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachOSSBucket"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachOSSBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachOSSBucket(request *DetachOSSBucketRequest) (_result *DetachOSSBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachOSSBucketResponse{}
	_body, _err := client.DetachOSSBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageBodiesWithOptions(tmpReq *DetectImageBodiesRequest, runtime *util.RuntimeOptions) (_result *DetectImageBodiesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageBodiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Sensitivity)) {
		query["Sensitivity"] = request.Sensitivity
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageBodies"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageBodies(request *DetectImageBodiesRequest) (_result *DetectImageBodiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageBodiesResponse{}
	_body, _err := client.DetectImageBodiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageCodesWithOptions(tmpReq *DetectImageCodesRequest, runtime *util.RuntimeOptions) (_result *DetectImageCodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageCodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageCodes"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageCodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageCodes(request *DetectImageCodesRequest) (_result *DetectImageCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageCodesResponse{}
	_body, _err := client.DetectImageCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageCroppingWithOptions(tmpReq *DetectImageCroppingRequest, runtime *util.RuntimeOptions) (_result *DetectImageCroppingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageCroppingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AspectRatios)) {
		query["AspectRatios"] = request.AspectRatios
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageCropping"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageCroppingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageCropping(request *DetectImageCroppingRequest) (_result *DetectImageCroppingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageCroppingResponse{}
	_body, _err := client.DetectImageCroppingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageFacesWithOptions(tmpReq *DetectImageFacesRequest, runtime *util.RuntimeOptions) (_result *DetectImageFacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageFacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageFaces"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageFaces(request *DetectImageFacesRequest) (_result *DetectImageFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageFacesResponse{}
	_body, _err := client.DetectImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageLabelsWithOptions(tmpReq *DetectImageLabelsRequest, runtime *util.RuntimeOptions) (_result *DetectImageLabelsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageLabelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageLabels"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageLabels(request *DetectImageLabelsRequest) (_result *DetectImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageLabelsResponse{}
	_body, _err := client.DetectImageLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageScoreWithOptions(tmpReq *DetectImageScoreRequest, runtime *util.RuntimeOptions) (_result *DetectImageScoreResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageScoreShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageScore"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageScore(request *DetectImageScoreRequest) (_result *DetectImageScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageScoreResponse{}
	_body, _err := client.DetectImageScoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectTextAnomalyWithOptions(request *DetectTextAnomalyRequest, runtime *util.RuntimeOptions) (_result *DetectTextAnomalyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectTextAnomaly"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectTextAnomalyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectTextAnomaly(request *DetectTextAnomalyRequest) (_result *DetectTextAnomalyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectTextAnomalyResponse{}
	_body, _err := client.DetectTextAnomalyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FuzzyQueryWithOptions(request *FuzzyQueryRequest, runtime *util.RuntimeOptions) (_result *FuzzyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FuzzyQuery"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FuzzyQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FuzzyQuery(request *FuzzyQueryRequest) (_result *FuzzyQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FuzzyQueryResponse{}
	_body, _err := client.FuzzyQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBindingWithOptions(request *GetBindingRequest, runtime *util.RuntimeOptions) (_result *GetBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBinding"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBinding(request *GetBindingRequest) (_result *GetBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBindingResponse{}
	_body, _err := client.GetBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatasetWithOptions(request *GetDatasetRequest, runtime *util.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.WithStatistics)) {
		query["WithStatistics"] = request.WithStatistics
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataset"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataset(request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectVideoLabelsResultWithOptions(request *GetDetectVideoLabelsResultRequest, runtime *util.RuntimeOptions) (_result *GetDetectVideoLabelsResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetectVideoLabelsResult"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetectVideoLabelsResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectVideoLabelsResult(request *GetDetectVideoLabelsResultRequest) (_result *GetDetectVideoLabelsResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectVideoLabelsResultResponse{}
	_body, _err := client.GetDetectVideoLabelsResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFigureClusterWithOptions(request *GetFigureClusterRequest, runtime *util.RuntimeOptions) (_result *GetFigureClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFigureCluster"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFigureCluster(request *GetFigureClusterRequest) (_result *GetFigureClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFigureClusterResponse{}
	_body, _err := client.GetFigureClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileMetaWithOptions(request *GetFileMetaRequest, runtime *util.RuntimeOptions) (_result *GetFileMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileMeta(request *GetFileMetaRequest) (_result *GetFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFileMetaResponse{}
	_body, _err := client.GetFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaMetaWithOptions(tmpReq *GetMediaMetaRequest, runtime *util.RuntimeOptions) (_result *GetMediaMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetMediaMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMediaMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaMeta(request *GetMediaMetaRequest) (_result *GetMediaMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaMetaResponse{}
	_body, _err := client.GetMediaMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOSSBucketAttachmentWithOptions(request *GetOSSBucketAttachmentRequest, runtime *util.RuntimeOptions) (_result *GetOSSBucketAttachmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OSSBucket)) {
		query["OSSBucket"] = request.OSSBucket
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOSSBucketAttachment"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOSSBucketAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOSSBucketAttachment(request *GetOSSBucketAttachmentRequest) (_result *GetOSSBucketAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOSSBucketAttachmentResponse{}
	_body, _err := client.GetOSSBucketAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.WithStatistics)) {
		query["WithStatistics"] = request.WithStatistics
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStoryWithOptions(request *GetStoryRequest, runtime *util.RuntimeOptions) (_result *GetStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStory(request *GetStoryRequest) (_result *GetStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStoryResponse{}
	_body, _err := client.GetStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebofficeURLWithOptions(tmpReq *GetWebofficeURLRequest, runtime *util.RuntimeOptions) (_result *GetWebofficeURLResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetWebofficeURLShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Permission))) {
		request.PermissionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Permission), tea.String("Permission"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.User))) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.User), tea.String("User"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Watermark))) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Watermark), tea.String("Watermark"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CachePreview)) {
		query["CachePreview"] = request.CachePreview
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUploaded)) {
		query["ExternalUploaded"] = request.ExternalUploaded
	}

	if !tea.BoolValue(util.IsUnset(request.Filename)) {
		query["Filename"] = request.Filename
	}

	if !tea.BoolValue(util.IsUnset(request.Hidecmb)) {
		query["Hidecmb"] = request.Hidecmb
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionShrink)) {
		query["Permission"] = request.PermissionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewPages)) {
		query["PreviewPages"] = request.PreviewPages
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Referer)) {
		query["Referer"] = request.Referer
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.UserShrink)) {
		query["User"] = request.UserShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkShrink)) {
		query["Watermark"] = request.WatermarkShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebofficeURL"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebofficeURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebofficeURL(request *GetWebofficeURLRequest) (_result *GetWebofficeURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebofficeURLResponse{}
	_body, _err := client.GetWebofficeURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndexFileMetaWithOptions(tmpReq *IndexFileMetaRequest, runtime *util.RuntimeOptions) (_result *IndexFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &IndexFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.File))) {
		request.FileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.File), tea.String("File"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FileShrink)) {
		query["File"] = request.FileShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IndexFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IndexFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndexFileMeta(request *IndexFileMetaRequest) (_result *IndexFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IndexFileMetaResponse{}
	_body, _err := client.IndexFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBindingsWithOptions(request *ListBindingsRequest, runtime *util.RuntimeOptions) (_result *ListBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBindings"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBindings(request *ListBindingsRequest) (_result *ListBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBindingsResponse{}
	_body, _err := client.ListBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, runtime *util.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasets"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(tmpReq *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.EndTimeRange))) {
		request.EndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.EndTimeRange), tea.String("EndTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.StartTimeRange))) {
		request.StartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.StartTimeRange), tea.String("StartTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskTypes)) {
		request.TaskTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTypes, tea.String("TaskTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimeRangeShrink)) {
		query["EndTimeRange"] = request.EndTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeRangeShrink)) {
		query["StartTimeRange"] = request.StartTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagSelector)) {
		query["TagSelector"] = request.TagSelector
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTypesShrink)) {
		query["TaskTypes"] = request.TaskTypesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeFigureClustersWithOptions(request *MergeFigureClustersRequest, runtime *util.RuntimeOptions) (_result *MergeFigureClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIdFrom)) {
		query["ClusterIdFrom"] = request.ClusterIdFrom
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterIdTo)) {
		query["ClusterIdTo"] = request.ClusterIdTo
	}

	if !tea.BoolValue(util.IsUnset(request.CustomMessage)) {
		query["CustomMessage"] = request.CustomMessage
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FigureType)) {
		query["FigureType"] = request.FigureType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicEndpoint)) {
		query["NotifyTopicEndpoint"] = request.NotifyTopicEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MergeFigureClusters"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MergeFigureClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeFigureClusters(request *MergeFigureClustersRequest) (_result *MergeFigureClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeFigureClustersResponse{}
	_body, _err := client.MergeFigureClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFigureClustersWithOptions(request *QueryFigureClustersRequest, runtime *util.RuntimeOptions) (_result *QueryFigureClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFigureClusters"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFigureClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFigureClusters(request *QueryFigureClustersRequest) (_result *QueryFigureClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFigureClustersResponse{}
	_body, _err := client.QueryFigureClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStoriesWithOptions(tmpReq *QueryStoriesRequest, runtime *util.RuntimeOptions) (_result *QueryStoriesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryStoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CreateTimeRange))) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CreateTimeRange), tea.String("CreateTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.FigureClusterIds)) {
		request.FigureClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FigureClusterIds, tea.String("FigureClusterIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.StoryEndTimeRange))) {
		request.StoryEndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.StoryEndTimeRange), tea.String("StoryEndTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.StoryStartTimeRange))) {
		request.StoryStartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.StoryStartTimeRange), tea.String("StoryStartTimeRange"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeRangeShrink)) {
		query["CreateTimeRange"] = request.CreateTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FigureClusterIdsShrink)) {
		query["FigureClusterIds"] = request.FigureClusterIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StoryEndTimeRangeShrink)) {
		query["StoryEndTimeRange"] = request.StoryEndTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		query["StoryName"] = request.StoryName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryStartTimeRangeShrink)) {
		query["StoryStartTimeRange"] = request.StoryStartTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StorySubType)) {
		query["StorySubType"] = request.StorySubType
	}

	if !tea.BoolValue(util.IsUnset(request.StoryType)) {
		query["StoryType"] = request.StoryType
	}

	if !tea.BoolValue(util.IsUnset(request.WithEmptyStories)) {
		query["WithEmptyStories"] = request.WithEmptyStories
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStories"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStories(request *QueryStoriesRequest) (_result *QueryStoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStoriesResponse{}
	_body, _err := client.QueryStoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshWebofficeTokenWithOptions(tmpReq *RefreshWebofficeTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshWebofficeTokenResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RefreshWebofficeTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CredentialConfig))) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CredentialConfig), tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		query["RefreshToken"] = request.RefreshToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshWebofficeToken"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshWebofficeToken(request *RefreshWebofficeTokenRequest) (_result *RefreshWebofficeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.RefreshWebofficeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveStoryFilesWithOptions(tmpReq *RemoveStoryFilesRequest, runtime *util.RuntimeOptions) (_result *RemoveStoryFilesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveStoryFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Files)) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, tea.String("Files"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FilesShrink)) {
		body["Files"] = request.FilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveStoryFiles"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveStoryFiles(request *RemoveStoryFilesRequest) (_result *RemoveStoryFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.RemoveStoryFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeBindingWithOptions(request *ResumeBindingRequest, runtime *util.RuntimeOptions) (_result *ResumeBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeBinding"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeBinding(request *ResumeBindingRequest) (_result *ResumeBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeBindingResponse{}
	_body, _err := client.ResumeBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SemanticQueryWithOptions(request *SemanticQueryRequest, runtime *util.RuntimeOptions) (_result *SemanticQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SemanticQuery"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SemanticQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SemanticQuery(request *SemanticQueryRequest) (_result *SemanticQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SemanticQueryResponse{}
	_body, _err := client.SemanticQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SimpleQueryWithOptions(tmpReq *SimpleQueryRequest, runtime *util.RuntimeOptions) (_result *SimpleQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SimpleQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Aggregations)) {
		request.AggregationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Aggregations, tea.String("Aggregations"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Query))) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Query), tea.String("Query"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.WithFields)) {
		request.WithFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WithFields, tea.String("WithFields"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregationsShrink)) {
		query["Aggregations"] = request.AggregationsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.QueryShrink)) {
		query["Query"] = request.QueryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.WithFieldsShrink)) {
		query["WithFields"] = request.WithFieldsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SimpleQuery"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SimpleQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SimpleQuery(request *SimpleQueryRequest) (_result *SimpleQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SimpleQueryResponse{}
	_body, _err := client.SimpleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopBindingWithOptions(request *StopBindingRequest, runtime *util.RuntimeOptions) (_result *StopBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopBinding"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopBinding(request *StopBindingRequest) (_result *StopBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopBindingResponse{}
	_body, _err := client.StopBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDatasetWithOptions(request *UpdateDatasetRequest, runtime *util.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetMaxBindCount)) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxEntityCount)) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxFileCount)) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxRelationCount)) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxTotalFileSize)) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataset"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataset(request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFigureClusterWithOptions(tmpReq *UpdateFigureClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateFigureClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateFigureClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.FigureCluster))) {
		request.FigureClusterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.FigureCluster), tea.String("FigureCluster"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FigureClusterShrink)) {
		query["FigureCluster"] = request.FigureClusterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFigureCluster"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFigureCluster(request *UpdateFigureClusterRequest) (_result *UpdateFigureClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFigureClusterResponse{}
	_body, _err := client.UpdateFigureClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFileMetaWithOptions(tmpReq *UpdateFileMetaRequest, runtime *util.RuntimeOptions) (_result *UpdateFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.File))) {
		request.FileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.File), tea.String("File"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FileShrink)) {
		query["File"] = request.FileShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFileMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFileMeta(request *UpdateFileMetaRequest) (_result *UpdateFileMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFileMetaResponse{}
	_body, _err := client.UpdateFileMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetMaxBindCount)) {
		query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxEntityCount)) {
		query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxFileCount)) {
		query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxRelationCount)) {
		query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetMaxTotalFileSize)) {
		query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EngineConcurrency)) {
		query["EngineConcurrency"] = request.EngineConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectMaxDatasetCount)) {
		query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectQueriesPerSecond)) {
		query["ProjectQueriesPerSecond"] = request.ProjectQueriesPerSecond
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRole)) {
		query["ServiceRole"] = request.ServiceRole
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStoryWithOptions(tmpReq *UpdateStoryRequest, runtime *util.RuntimeOptions) (_result *UpdateStoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Cover))) {
		request.CoverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Cover), tea.String("Cover"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverShrink)) {
		body["Cover"] = request.CoverShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomId)) {
		body["CustomId"] = request.CustomId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabelsShrink)) {
		body["CustomLabels"] = request.CustomLabelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["StoryName"] = request.StoryName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStory(request *UpdateStoryRequest) (_result *UpdateStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStoryResponse{}
	_body, _err := client.UpdateStoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
