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
	// AddressLine
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	// City
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Country
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// District
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Province
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// Township
	Township *string `json:"Township,omitempty" xml:"Township,omitempty"`
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
	// 链式授权节点
	Chain []*AssumeRoleChainNode `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	// 当前用户 policy
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
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
	// 账号id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 授权角色名
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 账号类型，普通账号填 user，服务账号填 service
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// Bitrate
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// ChannelLayout
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// Channels
	Channels *int64 `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// CodecLongName
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// CodecName
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// CodecTag
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// CodecTagString
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// CodecTimeBase
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// Duration
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Index
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Lyric
	Lyric *string `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
	// SampleFormat
	SampleFormat *string `json:"SampleFormat,omitempty" xml:"SampleFormat,omitempty"`
	// SampleRate
	SampleRate *int64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// StartTime
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// TimeBase
	TimeBase *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
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

func (s *AudioStream) SetDuration(v float32) *AudioStream {
	s.Duration = &v
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

func (s *AudioStream) SetStartTime(v float32) *AudioStream {
	s.StartTime = &v
	return s
}

func (s *AudioStream) SetTimeBase(v string) *AudioStream {
	s.TimeBase = &v
	return s
}

type Binding struct {
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// Detail
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Phase
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// State
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// UpdateTime
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

type Boundary struct {
	// Height
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Left
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// Top
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// Width
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	// Cover
	Cover *ClusterForReqCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// CustomId
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// CustomLabels
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// ObjectId
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
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
	// Figures
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
	// FigureId
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

type CroppingSuggestion struct {
	// AspectRatio
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// Boundary
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// Confidence
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
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
	// 媒体集当前绑定数
	BindCount *int64 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒体集最大绑定数
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// 媒体集最多实体数量
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 媒体集最多文件数量
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 媒体集最多关系数量
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 媒体集最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 媒体集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 媒体集当前文件数
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 媒体集当前文件总大小
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

type Face struct {
	// Age
	Age *int64 `json:"Age,omitempty" xml:"Age,omitempty"`
	// AgeConfidence
	AgeConfidence *float32 `json:"AgeConfidence,omitempty" xml:"AgeConfidence,omitempty"`
	// Beard
	Beard *string `json:"Beard,omitempty" xml:"Beard,omitempty"`
	// BeardConfidence
	BeardConfidence *float32  `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	Boundary        *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// EmbeddingsFloat32
	EmbeddingsFloat32 []*float32 `json:"EmbeddingsFloat32,omitempty" xml:"EmbeddingsFloat32,omitempty" type:"Repeated"`
	// EmbeddingsInt8
	EmbeddingsInt8 []*int32 `json:"EmbeddingsInt8,omitempty" xml:"EmbeddingsInt8,omitempty" type:"Repeated"`
	// Emotion
	Emotion *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// EmotionConfidence
	EmotionConfidence *float32 `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	// FaceClusterId
	FaceClusterId *string `json:"FaceClusterId,omitempty" xml:"FaceClusterId,omitempty"`
	// FaceConfidence
	FaceConfidence *float32 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	// FaceId
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// Gender
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// GenderConfidence
	GenderConfidence *float32 `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	// Glasses
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// GlassesConfidence
	GlassesConfidence *float32 `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	// Hat
	Hat *string `json:"Hat,omitempty" xml:"Hat,omitempty"`
	// HatConfidence
	HatConfidence *float32  `json:"HatConfidence,omitempty" xml:"HatConfidence,omitempty"`
	HeadPose      *HeadPose `json:"HeadPose,omitempty" xml:"HeadPose,omitempty"`
	// LeftEye
	LeftEye *string `json:"LeftEye,omitempty" xml:"LeftEye,omitempty"`
	// LeftEyeConfidence
	LeftEyeConfidence *float32 `json:"LeftEyeConfidence,omitempty" xml:"LeftEyeConfidence,omitempty"`
	// Mask
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// MaskConfidence
	MaskConfidence *float32 `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	// Mouth
	Mouth *string `json:"Mouth,omitempty" xml:"Mouth,omitempty"`
	// MouthConfidence
	MouthConfidence *float32 `json:"MouthConfidence,omitempty" xml:"MouthConfidence,omitempty"`
	// Race
	Race *string `json:"Race,omitempty" xml:"Race,omitempty"`
	// RaceConfidence
	RaceConfidence *float32 `json:"RaceConfidence,omitempty" xml:"RaceConfidence,omitempty"`
	// RightEye
	RightEye *string `json:"RightEye,omitempty" xml:"RightEye,omitempty"`
	// RightEyeConfidence
	RightEyeConfidence *float32 `json:"RightEyeConfidence,omitempty" xml:"RightEyeConfidence,omitempty"`
}

func (s Face) String() string {
	return tea.Prettify(s)
}

func (s Face) GoString() string {
	return s.String()
}

func (s *Face) SetAge(v int64) *Face {
	s.Age = &v
	return s
}

func (s *Face) SetAgeConfidence(v float32) *Face {
	s.AgeConfidence = &v
	return s
}

func (s *Face) SetBeard(v string) *Face {
	s.Beard = &v
	return s
}

func (s *Face) SetBeardConfidence(v float32) *Face {
	s.BeardConfidence = &v
	return s
}

func (s *Face) SetBoundary(v *Boundary) *Face {
	s.Boundary = v
	return s
}

func (s *Face) SetEmbeddingsFloat32(v []*float32) *Face {
	s.EmbeddingsFloat32 = v
	return s
}

func (s *Face) SetEmbeddingsInt8(v []*int32) *Face {
	s.EmbeddingsInt8 = v
	return s
}

func (s *Face) SetEmotion(v string) *Face {
	s.Emotion = &v
	return s
}

func (s *Face) SetEmotionConfidence(v float32) *Face {
	s.EmotionConfidence = &v
	return s
}

func (s *Face) SetFaceClusterId(v string) *Face {
	s.FaceClusterId = &v
	return s
}

func (s *Face) SetFaceConfidence(v float32) *Face {
	s.FaceConfidence = &v
	return s
}

func (s *Face) SetFaceId(v string) *Face {
	s.FaceId = &v
	return s
}

func (s *Face) SetGender(v string) *Face {
	s.Gender = &v
	return s
}

func (s *Face) SetGenderConfidence(v float32) *Face {
	s.GenderConfidence = &v
	return s
}

func (s *Face) SetGlasses(v string) *Face {
	s.Glasses = &v
	return s
}

func (s *Face) SetGlassesConfidence(v float32) *Face {
	s.GlassesConfidence = &v
	return s
}

func (s *Face) SetHat(v string) *Face {
	s.Hat = &v
	return s
}

func (s *Face) SetHatConfidence(v float32) *Face {
	s.HatConfidence = &v
	return s
}

func (s *Face) SetHeadPose(v *HeadPose) *Face {
	s.HeadPose = v
	return s
}

func (s *Face) SetLeftEye(v string) *Face {
	s.LeftEye = &v
	return s
}

func (s *Face) SetLeftEyeConfidence(v float32) *Face {
	s.LeftEyeConfidence = &v
	return s
}

func (s *Face) SetMask(v string) *Face {
	s.Mask = &v
	return s
}

func (s *Face) SetMaskConfidence(v float32) *Face {
	s.MaskConfidence = &v
	return s
}

func (s *Face) SetMouth(v string) *Face {
	s.Mouth = &v
	return s
}

func (s *Face) SetMouthConfidence(v float32) *Face {
	s.MouthConfidence = &v
	return s
}

func (s *Face) SetRace(v string) *Face {
	s.Race = &v
	return s
}

func (s *Face) SetRaceConfidence(v float32) *Face {
	s.RaceConfidence = &v
	return s
}

func (s *Face) SetRightEye(v string) *Face {
	s.RightEye = &v
	return s
}

func (s *Face) SetRightEyeConfidence(v float32) *Face {
	s.RightEyeConfidence = &v
	return s
}

type Figure struct {
	// Age
	Age *int64 `json:"Age,omitempty" xml:"Age,omitempty"`
	// AgeSD
	AgeSD *float32 `json:"AgeSD,omitempty" xml:"AgeSD,omitempty"`
	// Attractive
	Attractive *float32 `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	// Beard
	Beard *string `json:"Beard,omitempty" xml:"Beard,omitempty"`
	// BeardConfidence
	BeardConfidence *float32 `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	// Boundary
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// Emotion
	Emotion *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// EmotionConfidence
	EmotionConfidence *float32 `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	// FaceQuality
	FaceQuality *float32 `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// FigureClusterConfidence
	FigureClusterConfidence *float32 `json:"FigureClusterConfidence,omitempty" xml:"FigureClusterConfidence,omitempty"`
	// FigureClusterId
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	// FigureConfidence
	FigureConfidence *float32 `json:"FigureConfidence,omitempty" xml:"FigureConfidence,omitempty"`
	// FigureId
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// FigureType
	FigureType *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
	// Gender
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// GenderConfidence
	GenderConfidence *float32 `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	// Glasses
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// GlassesConfidence
	GlassesConfidence *float32 `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	// Hat
	Hat *string `json:"Hat,omitempty" xml:"Hat,omitempty"`
	// HatConfidence
	HatConfidence *float32  `json:"HatConfidence,omitempty" xml:"HatConfidence,omitempty"`
	HeadPose      *HeadPose `json:"HeadPose,omitempty" xml:"HeadPose,omitempty"`
	// Mask
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// MaskConfidence
	MaskConfidence *float32 `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	// Mouth
	Mouth *string `json:"Mouth,omitempty" xml:"Mouth,omitempty"`
	// MouthConfidence
	MouthConfidence *float32 `json:"MouthConfidence,omitempty" xml:"MouthConfidence,omitempty"`
	// Sharpness
	Sharpness *float32 `json:"Sharpness,omitempty" xml:"Sharpness,omitempty"`
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
	// AverageAge
	AverageAge *float32 `json:"AverageAge,omitempty" xml:"AverageAge,omitempty"`
	// Cover
	Cover *File `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// CustomId
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// CustomLabels
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// FaceCount
	FaceCount *int64 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// Gender
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// ImageCount
	ImageCount *int64 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// MaxAge
	MaxAge *float32 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	// MinAge
	MinAge *float32 `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// ObjectId
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// ObjectType
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// OwnerId
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// UpdateTime
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Version
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

func (s *FigureCluster) SetVersion(v string) *FigureCluster {
	s.Version = &v
	return s
}

type File struct {
	// AccessControlAllowOrigin
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	// AccessControlRequestMethod
	AccessControlRequestMethod *string `json:"AccessControlRequestMethod,omitempty" xml:"AccessControlRequestMethod,omitempty"`
	// Addresses
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// Album
	Album *string `json:"Album,omitempty" xml:"Album,omitempty"`
	// AlbumArtist
	AlbumArtist *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	// Artists
	Artists []*string `json:"Artists,omitempty" xml:"Artists,omitempty" type:"Repeated"`
	// AudioBitrate
	AudioBitrate *float32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// AudioCovers
	AudioCovers []*Image `json:"AudioCovers,omitempty" xml:"AudioCovers,omitempty" type:"Repeated"`
	// AudioDuration
	AudioDuration *float32 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	// AudioLanguage
	AudioLanguage *string `json:"AudioLanguage,omitempty" xml:"AudioLanguage,omitempty"`
	// AudioStreams
	AudioStreams []*AudioStream `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	// AudioTakenTime
	AudioTakenTime *string `json:"AudioTakenTime,omitempty" xml:"AudioTakenTime,omitempty"`
	// CacheControl
	CacheControl *string `json:"CacheControl,omitempty" xml:"CacheControl,omitempty"`
	// Composer
	Composer *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	// ContentDisposition
	ContentDisposition *string `json:"ContentDisposition,omitempty" xml:"ContentDisposition,omitempty"`
	// ContentEncoding
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// ContentLanguage
	ContentLanguage *string `json:"ContentLanguage,omitempty" xml:"ContentLanguage,omitempty"`
	// ContentMd5
	ContentMd5 *string `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	// ContentType
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// CroppingSuggestions
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	// CustomId
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// CustomLabels
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// DocumentContent
	DocumentContent *string `json:"DocumentContent,omitempty" xml:"DocumentContent,omitempty"`
	// DocumentLanguage
	DocumentLanguage *string `json:"DocumentLanguage,omitempty" xml:"DocumentLanguage,omitempty"`
	// ETag
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// EXIF
	EXIF *string `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	// FigureCount
	FigureCount *int64 `json:"FigureCount,omitempty" xml:"FigureCount,omitempty"`
	// Figures
	Figures []*Figure `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	// FileAccessTime
	FileAccessTime *string `json:"FileAccessTime,omitempty" xml:"FileAccessTime,omitempty"`
	// FileCreateTime
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// FileHash
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// FileModifiedTime
	FileModifiedTime *string `json:"FileModifiedTime,omitempty" xml:"FileModifiedTime,omitempty"`
	// Filename
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// ImageHeight
	ImageHeight *int64      `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageScore  *ImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	// ImageWidth
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// Labels
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// LatLong
	LatLong *string `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	// MediaType
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// OCRContents
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
	// OSSCRC64
	OSSCRC64 *string `json:"OSSCRC64,omitempty" xml:"OSSCRC64,omitempty"`
	// OSSDeleteMarker
	OSSDeleteMarker *string `json:"OSSDeleteMarker,omitempty" xml:"OSSDeleteMarker,omitempty"`
	// OSSExpiration
	OSSExpiration *string `json:"OSSExpiration,omitempty" xml:"OSSExpiration,omitempty"`
	// OSSObjectType
	OSSObjectType *string `json:"OSSObjectType,omitempty" xml:"OSSObjectType,omitempty"`
	// OSSStorageClass
	OSSStorageClass *string `json:"OSSStorageClass,omitempty" xml:"OSSStorageClass,omitempty"`
	// OSSTagging
	OSSTagging map[string]interface{} `json:"OSSTagging,omitempty" xml:"OSSTagging,omitempty"`
	// OSSTaggingCount
	OSSTaggingCount *int64 `json:"OSSTaggingCount,omitempty" xml:"OSSTaggingCount,omitempty"`
	// OSSURI
	OSSURI *string `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	// OSSUserMeta
	OSSUserMeta map[string]interface{} `json:"OSSUserMeta,omitempty" xml:"OSSUserMeta,omitempty"`
	// OSSVersionId
	OSSVersionId *string `json:"OSSVersionId,omitempty" xml:"OSSVersionId,omitempty"`
	// ObjectACL
	ObjectACL *string `json:"ObjectACL,omitempty" xml:"ObjectACL,omitempty"`
	// ObjectId
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// ObjectType
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// Orientation
	Orientation *string `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	// OwnerId
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// PageCount
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// Performer
	Performer *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	// ProduceTime
	ProduceTime *string `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// ServerSideDataEncryption
	ServerSideDataEncryption *string `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	// ServerSideEncryption
	ServerSideEncryption *string `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	// ServerSideEncryptionCustomerAlgorithm
	ServerSideEncryptionCustomerAlgorithm *string `json:"ServerSideEncryptionCustomerAlgorithm,omitempty" xml:"ServerSideEncryptionCustomerAlgorithm,omitempty"`
	// ServerSideEncryptionKeyId
	ServerSideEncryptionKeyId *string `json:"ServerSideEncryptionKeyId,omitempty" xml:"ServerSideEncryptionKeyId,omitempty"`
	// Size
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Subtitles
	Subtitles []*SubtitleStream `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// Timezone
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// TravelClusterId
	TravelClusterId *string `json:"TravelClusterId,omitempty" xml:"TravelClusterId,omitempty"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// UpdateTime
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// VideoBitrate
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// VideoDuration
	VideoDuration *float32 `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	// VideoHeight
	VideoHeight *int64 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// VideoStartTime
	VideoStartTime *float32 `json:"VideoStartTime,omitempty" xml:"VideoStartTime,omitempty"`
	// VideoStreams
	VideoStreams []*VideoStream `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	// VideoTakenTime
	VideoTakenTime *string `json:"VideoTakenTime,omitempty" xml:"VideoTakenTime,omitempty"`
	// VideoWidth
	VideoWidth *int64 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
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

func (s *File) SetArtists(v []*string) *File {
	s.Artists = v
	return s
}

func (s *File) SetAudioBitrate(v float32) *File {
	s.AudioBitrate = &v
	return s
}

func (s *File) SetAudioCovers(v []*Image) *File {
	s.AudioCovers = v
	return s
}

func (s *File) SetAudioDuration(v float32) *File {
	s.AudioDuration = &v
	return s
}

func (s *File) SetAudioLanguage(v string) *File {
	s.AudioLanguage = &v
	return s
}

func (s *File) SetAudioStreams(v []*AudioStream) *File {
	s.AudioStreams = v
	return s
}

func (s *File) SetAudioTakenTime(v string) *File {
	s.AudioTakenTime = &v
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

func (s *File) SetOrientation(v string) *File {
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

func (s *File) SetVideoBitrate(v int64) *File {
	s.VideoBitrate = &v
	return s
}

func (s *File) SetVideoDuration(v float32) *File {
	s.VideoDuration = &v
	return s
}

func (s *File) SetVideoHeight(v int64) *File {
	s.VideoHeight = &v
	return s
}

func (s *File) SetVideoStartTime(v float32) *File {
	s.VideoStartTime = &v
	return s
}

func (s *File) SetVideoStreams(v []*VideoStream) *File {
	s.VideoStreams = v
	return s
}

func (s *File) SetVideoTakenTime(v string) *File {
	s.VideoTakenTime = &v
	return s
}

func (s *File) SetVideoWidth(v int64) *File {
	s.VideoWidth = &v
	return s
}

type FileForReq struct {
	// ContentType
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// CustomId
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// CustomLabels
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// Figures
	Figures []*FileForReqFigures `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	// FileHash
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// MediaType
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// OSSURI
	OSSURI *string `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
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
	// FigureClusterId
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	// FigureId
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// FigureType
	FigureType *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
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
	// Pitch
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// Roll
	Roll *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	// Yaw
	Yaw *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
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
	// CroppingSuggestions
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	// EXIF
	EXIF *string `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	// ImageHeight
	ImageHeight *int64      `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageScore  *ImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	// ImageWidth
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// OCRContents
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
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
	// OverallQualityScore
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

type KeyValuePair struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 值
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
	// CentricScore
	CentricScore *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	// LabelConfidence
	LabelConfidence *float32 `json:"LabelConfidence,omitempty" xml:"LabelConfidence,omitempty"`
	// LabelLevel
	LabelLevel *int64 `json:"LabelLevel,omitempty" xml:"LabelLevel,omitempty"`
	// LabelName
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// ParentLabelName
	ParentLabelName *string `json:"ParentLabelName,omitempty" xml:"ParentLabelName,omitempty"`
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
	// Boundary
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// Confidence
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Contents
	Contents *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

type Project struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 项目当前媒体集数
	DatasetCount *int64 `json:"DatasetCount,omitempty" xml:"DatasetCount,omitempty"`
	// 项目最多绑定数
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// 项目最多实体数
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 项目最多文件数
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 项目最多关系数
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 项目最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 项目最大并发数
	EngineConcurrency *int64 `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	// 项目当前文件数
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// 项目最多媒体集数量
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 项目QPS
	ProjectQueriesPerSecond *int64 `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	// 服务角色
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// 项目当前文件总大小
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

type Row struct {
	// CustomLabels
	CustomLabels []*KeyValuePair `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty" type:"Repeated"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
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
	// 需要查询的字段名
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// 运算符
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 由 SimpleQuery 结构体组成的子查询数组
	SubQueries []*SimpleQuery `json:"SubQueries,omitempty" xml:"SubQueries,omitempty" type:"Repeated"`
	// 需要查询的字段值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type SubtitleStream struct {
	// Content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Index
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s SubtitleStream) String() string {
	return tea.Prettify(s)
}

func (s SubtitleStream) GoString() string {
	return s.String()
}

func (s *SubtitleStream) SetContent(v string) *SubtitleStream {
	s.Content = &v
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

type TaskInfo struct {
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 用户自定义信息
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

type VideoStream struct {
	// AverageFrameRate
	AverageFrameRate *float32 `json:"AverageFrameRate,omitempty" xml:"AverageFrameRate,omitempty"`
	// Bitrate
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// CodecLongName
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// CodecName
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// CodecTag
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// CodecTagString
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// CodecTimeBase
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// DisplayAspectRatio
	DisplayAspectRatio *string `json:"DisplayAspectRatio,omitempty" xml:"DisplayAspectRatio,omitempty"`
	// Duration
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// FrameCount
	FrameCount *int64 `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	// FrameRate
	FrameRate *float32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// HasBFrames
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// Height
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Index
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Level
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// PixelFormat
	PixelFormat *string `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	// Profile
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// SampleAspectRatio
	SampleAspectRatio *string `json:"SampleAspectRatio,omitempty" xml:"SampleAspectRatio,omitempty"`
	// StartTime
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// TimeBase
	TimeBase *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	// Width
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s VideoStream) String() string {
	return tea.Prettify(s)
}

func (s VideoStream) GoString() string {
	return s.String()
}

func (s *VideoStream) SetAverageFrameRate(v float32) *VideoStream {
	s.AverageFrameRate = &v
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

func (s *VideoStream) SetDisplayAspectRatio(v string) *VideoStream {
	s.DisplayAspectRatio = &v
	return s
}

func (s *VideoStream) SetDuration(v float32) *VideoStream {
	s.Duration = &v
	return s
}

func (s *VideoStream) SetFrameCount(v int64) *VideoStream {
	s.FrameCount = &v
	return s
}

func (s *VideoStream) SetFrameRate(v float32) *VideoStream {
	s.FrameRate = &v
	return s
}

func (s *VideoStream) SetHasBFrames(v string) *VideoStream {
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

func (s *VideoStream) SetSampleAspectRatio(v string) *VideoStream {
	s.SampleAspectRatio = &v
	return s
}

func (s *VideoStream) SetStartTime(v float32) *VideoStream {
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
	// 拷贝
	Copy *bool `json:"Copy,omitempty" xml:"Copy,omitempty"`
	// 导出
	Export *bool `json:"Export,omitempty" xml:"Export,omitempty"`
	// 查看历史版本
	History *bool `json:"History,omitempty" xml:"History,omitempty"`
	// 打印
	Print *bool `json:"Print,omitempty" xml:"Print,omitempty"`
	// 只读模式
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// 重命名
	Rename *bool `json:"Rename,omitempty" xml:"Rename,omitempty"`
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
	// 头像
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// Id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 字体颜色
	FillStyle *string `json:"FillStyle,omitempty" xml:"FillStyle,omitempty"`
	// 字体样式
	Font *string `json:"Font,omitempty" xml:"Font,omitempty"`
	// 水平间距
	Horizontal *int64 `json:"Horizontal,omitempty" xml:"Horizontal,omitempty"`
	// 旋转角度
	Rotate *float32 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// 水印类型，目前仅支持文字水印，0: 无水印；1: 文字水印
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 水印文字
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 垂直间距
	Vertical *int64 `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
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
	// Id of the request
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Id of the request
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchGetFileMetaResponse) SetBody(v *BatchGetFileMetaResponseBody) *BatchGetFileMetaResponse {
	s.Body = v
	return s
}

type BatchIndexFileMetaRequest struct {
	DatasetName     *string       `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files           []*FileForReq `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	NotifyEndpoint  *string       `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
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

func (s *BatchIndexFileMetaRequest) SetNotifyEndpoint(v string) *BatchIndexFileMetaRequest {
	s.NotifyEndpoint = &v
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
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
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

func (s *BatchIndexFileMetaShrinkRequest) SetNotifyEndpoint(v string) *BatchIndexFileMetaShrinkRequest {
	s.NotifyEndpoint = &v
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
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// Id of the request
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchIndexFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Files []*BatchUpdateFileMetaResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUpdateFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchUpdateFileMetaResponse) SetBody(v *BatchUpdateFileMetaResponseBody) *BatchUpdateFileMetaResponse {
	s.Body = v
	return s
}

type ClusterFiguresRequest struct {
	CustomMessage       *string `json:"CustomMessage,omitempty" xml:"CustomMessage,omitempty"`
	DatasetName         *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureType          *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
	NotifyTopicEndpoint *string `json:"NotifyTopicEndpoint,omitempty" xml:"NotifyTopicEndpoint,omitempty"`
	NotifyTopicName     *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName         *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ClusterFiguresRequest) String() string {
	return tea.Prettify(s)
}

func (s ClusterFiguresRequest) GoString() string {
	return s.String()
}

func (s *ClusterFiguresRequest) SetCustomMessage(v string) *ClusterFiguresRequest {
	s.CustomMessage = &v
	return s
}

func (s *ClusterFiguresRequest) SetDatasetName(v string) *ClusterFiguresRequest {
	s.DatasetName = &v
	return s
}

func (s *ClusterFiguresRequest) SetFigureType(v string) *ClusterFiguresRequest {
	s.FigureType = &v
	return s
}

func (s *ClusterFiguresRequest) SetNotifyTopicEndpoint(v string) *ClusterFiguresRequest {
	s.NotifyTopicEndpoint = &v
	return s
}

func (s *ClusterFiguresRequest) SetNotifyTopicName(v string) *ClusterFiguresRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *ClusterFiguresRequest) SetProjectName(v string) *ClusterFiguresRequest {
	s.ProjectName = &v
	return s
}

type ClusterFiguresResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ClusterFiguresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClusterFiguresResponseBody) GoString() string {
	return s.String()
}

func (s *ClusterFiguresResponseBody) SetRequestId(v string) *ClusterFiguresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClusterFiguresResponseBody) SetTaskId(v string) *ClusterFiguresResponseBody {
	s.TaskId = &v
	return s
}

type ClusterFiguresResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClusterFiguresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClusterFiguresResponse) String() string {
	return tea.Prettify(s)
}

func (s ClusterFiguresResponse) GoString() string {
	return s.String()
}

func (s *ClusterFiguresResponse) SetHeaders(v map[string]*string) *ClusterFiguresResponse {
	s.Headers = v
	return s
}

func (s *ClusterFiguresResponse) SetBody(v *ClusterFiguresResponseBody) *ClusterFiguresResponse {
	s.Body = v
	return s
}

type CreateBindingRequest struct {
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
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
	Binding *Binding `json:"Binding,omitempty" xml:"Binding,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateBindingResponse) SetBody(v *CreateBindingResponseBody) *CreateBindingResponse {
	s.Body = v
	return s
}

type CreateDatasetRequest struct {
	// 媒体集最多帮定数
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// 媒体集最多实体数
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 媒体集最多文件数
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 媒体集最多关系数
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 媒体集最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 对数据集的描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 模板Id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
	Dataset *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateDetectVideoLabelsTaskRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// SourceURI
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// UserData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateDetectVideoLabelsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectVideoLabelsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectVideoLabelsTaskRequest) SetProjectName(v string) *CreateDetectVideoLabelsTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetSourceURI(v string) *CreateDetectVideoLabelsTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskRequest) SetUserData(v string) *CreateDetectVideoLabelsTaskRequest {
	s.UserData = &v
	return s
}

type CreateDetectVideoLabelsTaskResponseBody struct {
	// 任务错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 事件Id
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// 任务错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 请求唯一Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务运行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 用户自定义信息
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateDetectVideoLabelsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectVideoLabelsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetCode(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetEndTime(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetEventId(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetMessage(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetProjectName(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetRequestId(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetStartTime(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetStatus(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetTaskId(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetTaskType(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *CreateDetectVideoLabelsTaskResponseBody) SetUserData(v string) *CreateDetectVideoLabelsTaskResponseBody {
	s.UserData = &v
	return s
}

type CreateDetectVideoLabelsTaskResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDetectVideoLabelsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDetectVideoLabelsTaskResponse) SetBody(v *CreateDetectVideoLabelsTaskResponseBody) *CreateDetectVideoLabelsTaskResponse {
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
	// 项目名称
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
	Project *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	// 本次请求的唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type DeleteBindingRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// A short description of struct
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DeleteBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
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
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFileMetaResponse) SetBody(v *DeleteFileMetaResponseBody) *DeleteFileMetaResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// 项目名称
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
	// 本次请求的唯一 ID
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DetectImageLabelsRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// SourceURI
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Threshold
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectImageLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageLabelsRequest) GoString() string {
	return s.String()
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

type DetectImageLabelsResponseBody struct {
	// 内容标签列表
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 请求唯一ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectImageLabelsResponse) SetBody(v *DetectImageLabelsResponseBody) *DetectImageLabelsResponse {
	s.Body = v
	return s
}

type FuzzyQueryRequest struct {
	// Dataset 名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 用于搜索的字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
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
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求的唯一 Id
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FuzzyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Binding *Binding `json:"Binding,omitempty" xml:"Binding,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

type GetDetectVideoLabelsResultRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// TaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// TaskType
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	// 任务错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 事件Id
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// 标签列表
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 任务错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 请求唯一Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务运行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 用户自定义信息
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectVideoLabelsResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDetectVideoLabelsResultResponse) SetBody(v *GetDetectVideoLabelsResultResponseBody) *GetDetectVideoLabelsResultResponse {
	s.Body = v
	return s
}

type GetFigureClusterRequest struct {
	DatasetName     *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *GetFigureClusterRequest) SetFigureClusterId(v string) *GetFigureClusterRequest {
	s.FigureClusterId = &v
	return s
}

func (s *GetFigureClusterRequest) SetProjectName(v string) *GetFigureClusterRequest {
	s.ProjectName = &v
	return s
}

type GetFigureClusterResponseBody struct {
	FigureCluster *FigureCluster `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// File list.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFileMetaResponse) SetBody(v *GetFileMetaResponseBody) *GetFileMetaResponse {
	s.Body = v
	return s
}

type GetFileSignedURIRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetFileSignedURIRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileSignedURIRequest) GoString() string {
	return s.String()
}

func (s *GetFileSignedURIRequest) SetProjectName(v string) *GetFileSignedURIRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileSignedURIRequest) SetURI(v string) *GetFileSignedURIRequest {
	s.URI = &v
	return s
}

type GetFileSignedURIResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 签名地址
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetFileSignedURIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileSignedURIResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileSignedURIResponseBody) SetRequestId(v string) *GetFileSignedURIResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileSignedURIResponseBody) SetURI(v string) *GetFileSignedURIResponseBody {
	s.URI = &v
	return s
}

type GetFileSignedURIResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileSignedURIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileSignedURIResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileSignedURIResponse) GoString() string {
	return s.String()
}

func (s *GetFileSignedURIResponse) SetHeaders(v map[string]*string) *GetFileSignedURIResponse {
	s.Headers = v
	return s
}

func (s *GetFileSignedURIResponse) SetBody(v *GetFileSignedURIResponseBody) *GetFileSignedURIResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 是否获取详细信息
	WithStatistics *bool `json:"WithStatistics,omitempty" xml:"WithStatistics,omitempty"`
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
	Project *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// TaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// TaskType
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	// 任务错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 事件Id
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// 任务错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 请求唯一Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务运行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务类型
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 用户自定义信息
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetWebofficeURLRequest struct {
	// 链式授权
	AssumeRoleChain *AssumeRoleChain `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
	// 是否支持外部上传
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// 文件名，必须带文件名后缀，默认是 SourceUri 的最后一级
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// 隐藏工具栏，预览模式下使用
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// mns 消息通知地址
	NotifyEndpoint *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	// mns 消息通知 topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// 文件密码
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 权限
	Permission *WebofficePermission `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// 预览前几页
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// oss 防盗链 referer
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// 预览编辑地址
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// 用户
	User *WebofficeUser `json:"User,omitempty" xml:"User,omitempty"`
	// 用户自定义数据，在消息通知中返回
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 水印
	Watermark *WebofficeWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GetWebofficeURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLRequest) SetAssumeRoleChain(v *AssumeRoleChain) *GetWebofficeURLRequest {
	s.AssumeRoleChain = v
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

func (s *GetWebofficeURLRequest) SetNotifyEndpoint(v string) *GetWebofficeURLRequest {
	s.NotifyEndpoint = &v
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
	// 链式授权
	AssumeRoleChainShrink *string `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
	// 是否支持外部上传
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// 文件名，必须带文件名后缀，默认是 SourceUri 的最后一级
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// 隐藏工具栏，预览模式下使用
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// mns 消息通知地址
	NotifyEndpoint *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	// mns 消息通知 topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// 文件密码
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 权限
	PermissionShrink *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// 预览前几页
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// oss 防盗链 referer
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// 预览编辑地址
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// 用户
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
	// 用户自定义数据，在消息通知中返回
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 水印
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GetWebofficeURLShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeURLShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeURLShrinkRequest) SetAssumeRoleChainShrink(v string) *GetWebofficeURLShrinkRequest {
	s.AssumeRoleChainShrink = &v
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

func (s *GetWebofficeURLShrinkRequest) SetNotifyEndpoint(v string) *GetWebofficeURLShrinkRequest {
	s.NotifyEndpoint = &v
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
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// access token 过期时间
	AccessTokenExpiredTime *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// refresh token 过期时间
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 预览编辑地址
	WebofficeURL *string `json:"WebofficeURL,omitempty" xml:"WebofficeURL,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebofficeURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWebofficeURLResponse) SetBody(v *GetWebofficeURLResponseBody) *GetWebofficeURLResponse {
	s.Body = v
	return s
}

type IndexFileMetaRequest struct {
	DatasetName     *string     `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	File            *FileForReq `json:"File,omitempty" xml:"File,omitempty"`
	NotifyEndpoint  *string     `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
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

func (s *IndexFileMetaRequest) SetNotifyEndpoint(v string) *IndexFileMetaRequest {
	s.NotifyEndpoint = &v
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
	NotifyEndpoint  *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
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

func (s *IndexFileMetaShrinkRequest) SetNotifyEndpoint(v string) *IndexFileMetaShrinkRequest {
	s.NotifyEndpoint = &v
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
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndexFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *IndexFileMetaResponse) SetBody(v *IndexFileMetaResponseBody) *IndexFileMetaResponse {
	s.Body = v
	return s
}

type ListBindingsRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A short description of struct
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
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListBindingsResponse) SetBody(v *ListBindingsResponseBody) *ListBindingsResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	// 返回最大个数
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Prefix    *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// 项目名称
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
	// Datasets
	Datasets  []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	NextToken *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
	s.Body = v
	return s
}

type ListFigureClustersRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListFigureClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFigureClustersRequest) GoString() string {
	return s.String()
}

func (s *ListFigureClustersRequest) SetDatasetName(v string) *ListFigureClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *ListFigureClustersRequest) SetLabels(v string) *ListFigureClustersRequest {
	s.Labels = &v
	return s
}

func (s *ListFigureClustersRequest) SetMaxResults(v int64) *ListFigureClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFigureClustersRequest) SetNextToken(v string) *ListFigureClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListFigureClustersRequest) SetOrder(v string) *ListFigureClustersRequest {
	s.Order = &v
	return s
}

func (s *ListFigureClustersRequest) SetProjectName(v string) *ListFigureClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *ListFigureClustersRequest) SetSort(v string) *ListFigureClustersRequest {
	s.Sort = &v
	return s
}

type ListFigureClustersResponseBody struct {
	FigureClusters []*FigureCluster `json:"FigureClusters,omitempty" xml:"FigureClusters,omitempty" type:"Repeated"`
	NextToken      *string          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFigureClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFigureClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFigureClustersResponseBody) SetFigureClusters(v []*FigureCluster) *ListFigureClustersResponseBody {
	s.FigureClusters = v
	return s
}

func (s *ListFigureClustersResponseBody) SetNextToken(v string) *ListFigureClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFigureClustersResponseBody) SetRequestId(v string) *ListFigureClustersResponseBody {
	s.RequestId = &v
	return s
}

type ListFigureClustersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFigureClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFigureClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFigureClustersResponse) GoString() string {
	return s.String()
}

func (s *ListFigureClustersResponse) SetHeaders(v map[string]*string) *ListFigureClustersResponse {
	s.Headers = v
	return s
}

func (s *ListFigureClustersResponse) SetBody(v *ListFigureClustersResponseBody) *ListFigureClustersResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// 返回结果的最大个数
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 列出包含某前缀的project
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
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
	// 当总结果个数大于MaxResults时，用于翻页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 由ProjectItem组成的数组
	Projects []*Project `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// 本次请求的唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// MaxResults
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// TaskType
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetMaxResults(v int64) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetProjectName(v string) *ListTasksRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

type ListTasksResponseBody struct {
	// 最大结果数量
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 翻页标记
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 请求唯一Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务信息
	Tasks []*TaskInfo `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MergeFigureClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MergeFigureClustersResponse) SetBody(v *MergeFigureClustersResponseBody) *MergeFigureClustersResponse {
	s.Body = v
	return s
}

type RefreshWebofficeTokenRequest struct {
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 链式授权
	AssumeRoleChain *AssumeRoleChain `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
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

func (s *RefreshWebofficeTokenRequest) SetAssumeRoleChain(v *AssumeRoleChain) *RefreshWebofficeTokenRequest {
	s.AssumeRoleChain = v
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
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 链式授权
	AssumeRoleChainShrink *string `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
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

func (s *RefreshWebofficeTokenShrinkRequest) SetAssumeRoleChainShrink(v string) *RefreshWebofficeTokenShrinkRequest {
	s.AssumeRoleChainShrink = &v
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
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// access token 过期时间
	AccessTokenExpiredTime *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// refresh token 过期时间
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	// 请求 Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RefreshWebofficeTokenResponse) SetBody(v *RefreshWebofficeTokenResponseBody) *RefreshWebofficeTokenResponse {
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
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResumeBindingResponse) SetBody(v *ResumeBindingResponseBody) *ResumeBindingResponse {
	s.Body = v
	return s
}

type SemanticQueryRequest struct {
	// Dataset 名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 需要搜索的内容，使用自然语言描述
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
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
	// 聚合字段的字段名
	Aggregations []*SemanticQueryResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	// 文件列表
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求的唯一 Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 聚合字段名
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// 分组聚合的结果
	Groups []*SemanticQueryResponseBodyAggregationsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// 聚合字段的聚合操作符
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 聚合的统计结果
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// 分组聚合的计数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 分组聚合的值
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SemanticQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SemanticQueryResponse) SetBody(v *SemanticQueryResponseBody) *SemanticQueryResponse {
	s.Body = v
	return s
}

type SimpleQueryRequest struct {
	// 聚合字段
	Aggregations []*SimpleQueryRequestAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	// Dataset 名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 排序字段
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 项目名称
	ProjectName *string      `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Query       *SimpleQuery `json:"Query,omitempty" xml:"Query,omitempty"`
	// 排序方式，默认 DESC
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

type SimpleQueryRequestAggregations struct {
	// 聚合字段的字段名
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// 聚合字段的聚合操作符
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
	// 聚合字段
	AggregationsShrink *string `json:"Aggregations,omitempty" xml:"Aggregations,omitempty"`
	// Dataset 名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 排序字段
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// 排序方式，默认 DESC
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

type SimpleQueryResponseBody struct {
	// 聚合字段的字段名
	Aggregations []*SimpleQueryResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
	// 文件列表
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求的唯一 Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 聚合字段名
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// 分组聚合的结果
	Groups []*SimpleQueryResponseBodyAggregationsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// 聚合字段的聚合操作符
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 聚合的统计结果
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
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

func (s *SimpleQueryResponseBodyAggregations) SetValue(v float32) *SimpleQueryResponseBodyAggregations {
	s.Value = &v
	return s
}

type SimpleQueryResponseBodyAggregationsGroups struct {
	// 分组聚合的计数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 分组聚合的值
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SimpleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopBindingResponse) SetBody(v *StopBindingResponseBody) *StopBindingResponse {
	s.Body = v
	return s
}

type UpdateDatasetRequest struct {
	// 媒体集最多绑定数
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// 媒体集最多实体数
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 媒体集最多文件数
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 媒体集最多关系数
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 媒体集最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 媒体集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 模板Id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
	Dataset *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

type UpdateFigureClusterRequest struct {
	DatasetName   *string        `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FigureCluster *FigureCluster `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	ProjectName   *string        `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *UpdateFigureClusterRequest) SetFigureCluster(v *FigureCluster) *UpdateFigureClusterRequest {
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
	// Id of the request
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateFileMetaResponse) SetBody(v *UpdateFileMetaResponseBody) *UpdateFileMetaResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	// 媒体集最多绑定数
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// 媒体集最多实体数
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 媒体集最多文件数
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 媒体集最多关系数
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 媒体集最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 项目描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 项目并发数
	EngineConcurrency *int64 `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	// 项目最多媒体集数
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 项目QPS
	ProjectQueriesPerSecond *int64 `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	// 服务角色
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// 模板Id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
	Project *Project `json:"Project,omitempty" xml:"Project,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URIs"] = request.URIsShrink
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URIs"] = request.URIsShrink
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
	query["DatasetName"] = request.DatasetName
	query["Files"] = request.FilesShrink
	query["NotifyEndpoint"] = request.NotifyEndpoint
	query["NotifyTopicName"] = request.NotifyTopicName
	query["ProjectName"] = request.ProjectName
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
	query["DatasetName"] = request.DatasetName
	query["Files"] = request.FilesShrink
	query["ProjectName"] = request.ProjectName
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

func (client *Client) ClusterFiguresWithOptions(request *ClusterFiguresRequest, runtime *util.RuntimeOptions) (_result *ClusterFiguresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CustomMessage"] = request.CustomMessage
	query["DatasetName"] = request.DatasetName
	query["FigureType"] = request.FigureType
	query["NotifyTopicEndpoint"] = request.NotifyTopicEndpoint
	query["NotifyTopicName"] = request.NotifyTopicName
	query["ProjectName"] = request.ProjectName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClusterFigures"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClusterFiguresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClusterFigures(request *ClusterFiguresRequest) (_result *ClusterFiguresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClusterFiguresResponse{}
	_body, _err := client.ClusterFiguresWithOptions(request, runtime)
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
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

func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, runtime *util.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	query["DatasetName"] = request.DatasetName
	query["Description"] = request.Description
	query["ProjectName"] = request.ProjectName
	query["TemplateId"] = request.TemplateId
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

func (client *Client) CreateDetectVideoLabelsTaskWithOptions(request *CreateDetectVideoLabelsTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDetectVideoLabelsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectName"] = request.ProjectName
	query["SourceURI"] = request.SourceURI
	query["UserData"] = request.UserData
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

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	query["Description"] = request.Description
	query["EngineConcurrency"] = request.EngineConcurrency
	query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	query["ProjectName"] = request.ProjectName
	query["ProjectQueriesPerSecond"] = request.ProjectQueriesPerSecond
	query["ServiceRole"] = request.ServiceRole
	query["TemplateId"] = request.TemplateId
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

func (client *Client) DeleteBindingWithOptions(request *DeleteBindingRequest, runtime *util.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
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
	query["ProjectName"] = request.ProjectName
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

func (client *Client) DetectImageLabelsWithOptions(request *DetectImageLabelsRequest, runtime *util.RuntimeOptions) (_result *DetectImageLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectName"] = request.ProjectName
	query["SourceURI"] = request.SourceURI
	query["Threshold"] = request.Threshold
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

func (client *Client) FuzzyQueryWithOptions(request *FuzzyQueryRequest, runtime *util.RuntimeOptions) (_result *FuzzyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DatasetName"] = request.DatasetName
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProjectName"] = request.ProjectName
	query["Query"] = request.Query
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["WithStatistics"] = request.WithStatistics
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
	query["ProjectName"] = request.ProjectName
	query["TaskId"] = request.TaskId
	query["TaskType"] = request.TaskType
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
	query["DatasetName"] = request.DatasetName
	query["FigureClusterId"] = request.FigureClusterId
	query["ProjectName"] = request.ProjectName
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
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

func (client *Client) GetFileSignedURIWithOptions(request *GetFileSignedURIRequest, runtime *util.RuntimeOptions) (_result *GetFileSignedURIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileSignedURI"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileSignedURIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileSignedURI(request *GetFileSignedURIRequest) (_result *GetFileSignedURIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFileSignedURIResponse{}
	_body, _err := client.GetFileSignedURIWithOptions(request, runtime)
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
	query["ProjectName"] = request.ProjectName
	query["WithStatistics"] = request.WithStatistics
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

func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectName"] = request.ProjectName
	query["TaskId"] = request.TaskId
	query["TaskType"] = request.TaskType
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
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.AssumeRoleChain))) {
		request.AssumeRoleChainShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.AssumeRoleChain), tea.String("AssumeRoleChain"), tea.String("json"))
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
	query["AssumeRoleChain"] = request.AssumeRoleChainShrink
	query["ExternalUploaded"] = request.ExternalUploaded
	query["Filename"] = request.Filename
	query["Hidecmb"] = request.Hidecmb
	query["NotifyEndpoint"] = request.NotifyEndpoint
	query["NotifyTopicName"] = request.NotifyTopicName
	query["Password"] = request.Password
	query["Permission"] = request.PermissionShrink
	query["PreviewPages"] = request.PreviewPages
	query["ProjectName"] = request.ProjectName
	query["Referer"] = request.Referer
	query["SourceURI"] = request.SourceURI
	query["User"] = request.UserShrink
	query["UserData"] = request.UserData
	query["Watermark"] = request.WatermarkShrink
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
	query["DatasetName"] = request.DatasetName
	query["File"] = request.FileShrink
	query["NotifyEndpoint"] = request.NotifyEndpoint
	query["NotifyTopicName"] = request.NotifyTopicName
	query["ProjectName"] = request.ProjectName
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
	query["DatasetName"] = request.DatasetName
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProjectName"] = request.ProjectName
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
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["Prefix"] = request.Prefix
	query["ProjectName"] = request.ProjectName
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

func (client *Client) ListFigureClustersWithOptions(request *ListFigureClustersRequest, runtime *util.RuntimeOptions) (_result *ListFigureClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DatasetName"] = request.DatasetName
	query["Labels"] = request.Labels
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["Order"] = request.Order
	query["ProjectName"] = request.ProjectName
	query["Sort"] = request.Sort
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFigureClusters"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFigureClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFigureClusters(request *ListFigureClustersRequest) (_result *ListFigureClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFigureClustersResponse{}
	_body, _err := client.ListFigureClustersWithOptions(request, runtime)
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
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["Prefix"] = request.Prefix
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

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProjectName"] = request.ProjectName
	query["TaskType"] = request.TaskType
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
	query["ClusterIdFrom"] = request.ClusterIdFrom
	query["ClusterIdTo"] = request.ClusterIdTo
	query["CustomMessage"] = request.CustomMessage
	query["DatasetName"] = request.DatasetName
	query["FigureType"] = request.FigureType
	query["NotifyTopicEndpoint"] = request.NotifyTopicEndpoint
	query["NotifyTopicName"] = request.NotifyTopicName
	query["ProjectName"] = request.ProjectName
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

func (client *Client) RefreshWebofficeTokenWithOptions(tmpReq *RefreshWebofficeTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshWebofficeTokenResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RefreshWebofficeTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.AssumeRoleChain))) {
		request.AssumeRoleChainShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.AssumeRoleChain), tea.String("AssumeRoleChain"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["AccessToken"] = request.AccessToken
	query["AssumeRoleChain"] = request.AssumeRoleChainShrink
	query["ProjectName"] = request.ProjectName
	query["RefreshToken"] = request.RefreshToken
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

func (client *Client) ResumeBindingWithOptions(request *ResumeBindingRequest, runtime *util.RuntimeOptions) (_result *ResumeBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["URI"] = request.URI
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
	query["DatasetName"] = request.DatasetName
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["ProjectName"] = request.ProjectName
	query["Query"] = request.Query
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

	query := map[string]interface{}{}
	query["Aggregations"] = request.AggregationsShrink
	query["DatasetName"] = request.DatasetName
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["Order"] = request.Order
	query["ProjectName"] = request.ProjectName
	query["Query"] = request.QueryShrink
	query["Sort"] = request.Sort
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
	query["DatasetName"] = request.DatasetName
	query["ProjectName"] = request.ProjectName
	query["Reason"] = request.Reason
	query["URI"] = request.URI
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
	query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	query["DatasetName"] = request.DatasetName
	query["Description"] = request.Description
	query["ProjectName"] = request.ProjectName
	query["TemplateId"] = request.TemplateId
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
	query["DatasetName"] = request.DatasetName
	query["FigureCluster"] = request.FigureClusterShrink
	query["ProjectName"] = request.ProjectName
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
	query["DatasetName"] = request.DatasetName
	query["File"] = request.FileShrink
	query["ProjectName"] = request.ProjectName
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
	query["DatasetMaxBindCount"] = request.DatasetMaxBindCount
	query["DatasetMaxEntityCount"] = request.DatasetMaxEntityCount
	query["DatasetMaxFileCount"] = request.DatasetMaxFileCount
	query["DatasetMaxRelationCount"] = request.DatasetMaxRelationCount
	query["DatasetMaxTotalFileSize"] = request.DatasetMaxTotalFileSize
	query["Description"] = request.Description
	query["EngineConcurrency"] = request.EngineConcurrency
	query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	query["ProjectName"] = request.ProjectName
	query["ProjectQueriesPerSecond"] = request.ProjectQueriesPerSecond
	query["ServiceRole"] = request.ServiceRole
	query["TemplateId"] = request.TemplateId
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
