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

type CroppingSuggestion struct {
	// AspectRatio
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// Confidence
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Boundary
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
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

func (s *CroppingSuggestion) SetConfidence(v float32) *CroppingSuggestion {
	s.Confidence = &v
	return s
}

func (s *CroppingSuggestion) SetBoundary(v *Boundary) *CroppingSuggestion {
	s.Boundary = v
	return s
}

type Address struct {
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// AddressLine
	AddressLine *string `json:"AddressLine,omitempty" xml:"AddressLine,omitempty"`
	// Country
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Province
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// City
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// District
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	// Township
	Township *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s Address) String() string {
	return tea.Prettify(s)
}

func (s Address) GoString() string {
	return s.String()
}

func (s *Address) SetLanguage(v string) *Address {
	s.Language = &v
	return s
}

func (s *Address) SetAddressLine(v string) *Address {
	s.AddressLine = &v
	return s
}

func (s *Address) SetCountry(v string) *Address {
	s.Country = &v
	return s
}

func (s *Address) SetProvince(v string) *Address {
	s.Province = &v
	return s
}

func (s *Address) SetCity(v string) *Address {
	s.City = &v
	return s
}

func (s *Address) SetDistrict(v string) *Address {
	s.District = &v
	return s
}

func (s *Address) SetTownship(v string) *Address {
	s.Township = &v
	return s
}

type SubtitleStream struct {
	// Index
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SubtitleStream) String() string {
	return tea.Prettify(s)
}

func (s SubtitleStream) GoString() string {
	return s.String()
}

func (s *SubtitleStream) SetIndex(v int64) *SubtitleStream {
	s.Index = &v
	return s
}

func (s *SubtitleStream) SetLanguage(v string) *SubtitleStream {
	s.Language = &v
	return s
}

func (s *SubtitleStream) SetContent(v string) *SubtitleStream {
	s.Content = &v
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

type Label struct {
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// LabelName
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// LabelLevel
	LabelLevel *int64 `json:"LabelLevel,omitempty" xml:"LabelLevel,omitempty"`
	// LabelConfidence
	LabelConfidence *float32 `json:"LabelConfidence,omitempty" xml:"LabelConfidence,omitempty"`
}

func (s Label) String() string {
	return tea.Prettify(s)
}

func (s Label) GoString() string {
	return s.String()
}

func (s *Label) SetLanguage(v string) *Label {
	s.Language = &v
	return s
}

func (s *Label) SetLabelName(v string) *Label {
	s.LabelName = &v
	return s
}

func (s *Label) SetLabelLevel(v int64) *Label {
	s.LabelLevel = &v
	return s
}

func (s *Label) SetLabelConfidence(v float32) *Label {
	s.LabelConfidence = &v
	return s
}

type VideoStream struct {
	// Index
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// CodecName
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// CodecLongName
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// Profile
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// CodecTimeBase
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// CodecTagString
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// CodecTag
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// Width
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// Height
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// HasBFrames
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// SampleAspectRatio
	SampleAspectRatio *string `json:"SampleAspectRatio,omitempty" xml:"SampleAspectRatio,omitempty"`
	// DisplayAspectRatio
	DisplayAspectRatio *string `json:"DisplayAspectRatio,omitempty" xml:"DisplayAspectRatio,omitempty"`
	// PixelFormat
	PixelFormat *string `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	// Level
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// FrameRate
	FrameRate *float32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// AverageFrameRate
	AverageFrameRate *float32 `json:"AverageFrameRate,omitempty" xml:"AverageFrameRate,omitempty"`
	// TimeBase
	TimeBase *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	// StartTime
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Duration
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Bitrate
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// FrameCount
	FrameCount *int64 `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
}

func (s VideoStream) String() string {
	return tea.Prettify(s)
}

func (s VideoStream) GoString() string {
	return s.String()
}

func (s *VideoStream) SetIndex(v int64) *VideoStream {
	s.Index = &v
	return s
}

func (s *VideoStream) SetLanguage(v string) *VideoStream {
	s.Language = &v
	return s
}

func (s *VideoStream) SetCodecName(v string) *VideoStream {
	s.CodecName = &v
	return s
}

func (s *VideoStream) SetCodecLongName(v string) *VideoStream {
	s.CodecLongName = &v
	return s
}

func (s *VideoStream) SetProfile(v string) *VideoStream {
	s.Profile = &v
	return s
}

func (s *VideoStream) SetCodecTimeBase(v string) *VideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *VideoStream) SetCodecTagString(v string) *VideoStream {
	s.CodecTagString = &v
	return s
}

func (s *VideoStream) SetCodecTag(v string) *VideoStream {
	s.CodecTag = &v
	return s
}

func (s *VideoStream) SetWidth(v int64) *VideoStream {
	s.Width = &v
	return s
}

func (s *VideoStream) SetHeight(v int64) *VideoStream {
	s.Height = &v
	return s
}

func (s *VideoStream) SetHasBFrames(v string) *VideoStream {
	s.HasBFrames = &v
	return s
}

func (s *VideoStream) SetSampleAspectRatio(v string) *VideoStream {
	s.SampleAspectRatio = &v
	return s
}

func (s *VideoStream) SetDisplayAspectRatio(v string) *VideoStream {
	s.DisplayAspectRatio = &v
	return s
}

func (s *VideoStream) SetPixelFormat(v string) *VideoStream {
	s.PixelFormat = &v
	return s
}

func (s *VideoStream) SetLevel(v int64) *VideoStream {
	s.Level = &v
	return s
}

func (s *VideoStream) SetFrameRate(v float32) *VideoStream {
	s.FrameRate = &v
	return s
}

func (s *VideoStream) SetAverageFrameRate(v float32) *VideoStream {
	s.AverageFrameRate = &v
	return s
}

func (s *VideoStream) SetTimeBase(v string) *VideoStream {
	s.TimeBase = &v
	return s
}

func (s *VideoStream) SetStartTime(v float32) *VideoStream {
	s.StartTime = &v
	return s
}

func (s *VideoStream) SetDuration(v float32) *VideoStream {
	s.Duration = &v
	return s
}

func (s *VideoStream) SetBitrate(v int64) *VideoStream {
	s.Bitrate = &v
	return s
}

func (s *VideoStream) SetFrameCount(v int64) *VideoStream {
	s.FrameCount = &v
	return s
}

type Image struct {
	// ImageWidth
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// ImageHeight
	ImageHeight *int64 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// EXIF
	EXIF       *string     `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	ImageScore *ImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	// CroppingSuggestions
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	// OCRContents
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
}

func (s Image) String() string {
	return tea.Prettify(s)
}

func (s Image) GoString() string {
	return s.String()
}

func (s *Image) SetImageWidth(v int64) *Image {
	s.ImageWidth = &v
	return s
}

func (s *Image) SetImageHeight(v int64) *Image {
	s.ImageHeight = &v
	return s
}

func (s *Image) SetEXIF(v string) *Image {
	s.EXIF = &v
	return s
}

func (s *Image) SetImageScore(v *ImageScore) *Image {
	s.ImageScore = v
	return s
}

func (s *Image) SetCroppingSuggestions(v []*CroppingSuggestion) *Image {
	s.CroppingSuggestions = v
	return s
}

func (s *Image) SetOCRContents(v []*OCRContents) *Image {
	s.OCRContents = v
	return s
}

type Boundary struct {
	// Width
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// Height
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Left
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// Top
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s Boundary) String() string {
	return tea.Prettify(s)
}

func (s Boundary) GoString() string {
	return s.String()
}

func (s *Boundary) SetWidth(v int64) *Boundary {
	s.Width = &v
	return s
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

type AudioStream struct {
	// Index
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// CodecName
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// CodecLongName
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// CodecTimeBase
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// CodecTagString
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// CodecTag
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// SampleFormat
	SampleFormat *string `json:"SampleFormat,omitempty" xml:"SampleFormat,omitempty"`
	// SampleRate
	SampleRate *int64 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// Channels
	Channels *int64 `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// ChannelLayout
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// TimeBase
	TimeBase *string `json:"TimeBase,omitempty" xml:"TimeBase,omitempty"`
	// StartTime
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Duration
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Bitrate
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Lyric
	Lyric *string `json:"Lyric,omitempty" xml:"Lyric,omitempty"`
}

func (s AudioStream) String() string {
	return tea.Prettify(s)
}

func (s AudioStream) GoString() string {
	return s.String()
}

func (s *AudioStream) SetIndex(v int64) *AudioStream {
	s.Index = &v
	return s
}

func (s *AudioStream) SetLanguage(v string) *AudioStream {
	s.Language = &v
	return s
}

func (s *AudioStream) SetCodecName(v string) *AudioStream {
	s.CodecName = &v
	return s
}

func (s *AudioStream) SetCodecLongName(v string) *AudioStream {
	s.CodecLongName = &v
	return s
}

func (s *AudioStream) SetCodecTimeBase(v string) *AudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *AudioStream) SetCodecTagString(v string) *AudioStream {
	s.CodecTagString = &v
	return s
}

func (s *AudioStream) SetCodecTag(v string) *AudioStream {
	s.CodecTag = &v
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

func (s *AudioStream) SetChannels(v int64) *AudioStream {
	s.Channels = &v
	return s
}

func (s *AudioStream) SetChannelLayout(v string) *AudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *AudioStream) SetTimeBase(v string) *AudioStream {
	s.TimeBase = &v
	return s
}

func (s *AudioStream) SetStartTime(v float32) *AudioStream {
	s.StartTime = &v
	return s
}

func (s *AudioStream) SetDuration(v float32) *AudioStream {
	s.Duration = &v
	return s
}

func (s *AudioStream) SetBitrate(v int64) *AudioStream {
	s.Bitrate = &v
	return s
}

func (s *AudioStream) SetLyric(v string) *AudioStream {
	s.Lyric = &v
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

type OCRContents struct {
	// Language
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Contents
	Contents *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	// Confidence
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Boundary
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
}

func (s OCRContents) String() string {
	return tea.Prettify(s)
}

func (s OCRContents) GoString() string {
	return s.String()
}

func (s *OCRContents) SetLanguage(v string) *OCRContents {
	s.Language = &v
	return s
}

func (s *OCRContents) SetContents(v string) *OCRContents {
	s.Contents = &v
	return s
}

func (s *OCRContents) SetConfidence(v float32) *OCRContents {
	s.Confidence = &v
	return s
}

func (s *OCRContents) SetBoundary(v *Boundary) *OCRContents {
	s.Boundary = v
	return s
}

type Face struct {
	// FaceId
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// FaceConfidence
	FaceConfidence *float32 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	// Age
	Age *int64 `json:"Age,omitempty" xml:"Age,omitempty"`
	// AgeConfidence
	AgeConfidence *float32 `json:"AgeConfidence,omitempty" xml:"AgeConfidence,omitempty"`
	// Gender
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// GenderConfidence
	GenderConfidence *float32 `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	// Emotion
	Emotion *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// EmotionConfidence
	EmotionConfidence *float32 `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	// FaceClusterId
	FaceClusterId *string `json:"FaceClusterId,omitempty" xml:"FaceClusterId,omitempty"`
	// Mouth
	Mouth *string `json:"Mouth,omitempty" xml:"Mouth,omitempty"`
	// MouthConfidence
	MouthConfidence *float32 `json:"MouthConfidence,omitempty" xml:"MouthConfidence,omitempty"`
	// Beard
	Beard *string `json:"Beard,omitempty" xml:"Beard,omitempty"`
	// BeardConfidence
	BeardConfidence *float32 `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	// Hat
	Hat *string `json:"Hat,omitempty" xml:"Hat,omitempty"`
	// HatConfidence
	HatConfidence *float32 `json:"HatConfidence,omitempty" xml:"HatConfidence,omitempty"`
	// Race
	Race *string `json:"Race,omitempty" xml:"Race,omitempty"`
	// RaceConfidence
	RaceConfidence *float32 `json:"RaceConfidence,omitempty" xml:"RaceConfidence,omitempty"`
	// Mask
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// MaskConfidence
	MaskConfidence *float32 `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	// Glasses
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// GlassesConfidence
	GlassesConfidence *float32 `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	// LeftEye
	LeftEye *string `json:"LeftEye,omitempty" xml:"LeftEye,omitempty"`
	// LeftEyeConfidence
	LeftEyeConfidence *float32 `json:"LeftEyeConfidence,omitempty" xml:"LeftEyeConfidence,omitempty"`
	// RightEye
	RightEye *string `json:"RightEye,omitempty" xml:"RightEye,omitempty"`
	// RightEyeConfidence
	RightEyeConfidence *float32  `json:"RightEyeConfidence,omitempty" xml:"RightEyeConfidence,omitempty"`
	HeadPose           *HeadPose `json:"HeadPose,omitempty" xml:"HeadPose,omitempty"`
	Boundary           *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// EmbeddingsFloat32
	EmbeddingsFloat32 []*float32 `json:"EmbeddingsFloat32,omitempty" xml:"EmbeddingsFloat32,omitempty" type:"Repeated"`
	// EmbeddingsInt8
	EmbeddingsInt8 []*int32 `json:"EmbeddingsInt8,omitempty" xml:"EmbeddingsInt8,omitempty" type:"Repeated"`
}

func (s Face) String() string {
	return tea.Prettify(s)
}

func (s Face) GoString() string {
	return s.String()
}

func (s *Face) SetFaceId(v string) *Face {
	s.FaceId = &v
	return s
}

func (s *Face) SetFaceConfidence(v float32) *Face {
	s.FaceConfidence = &v
	return s
}

func (s *Face) SetAge(v int64) *Face {
	s.Age = &v
	return s
}

func (s *Face) SetAgeConfidence(v float32) *Face {
	s.AgeConfidence = &v
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

func (s *Face) SetMouth(v string) *Face {
	s.Mouth = &v
	return s
}

func (s *Face) SetMouthConfidence(v float32) *Face {
	s.MouthConfidence = &v
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

func (s *Face) SetHat(v string) *Face {
	s.Hat = &v
	return s
}

func (s *Face) SetHatConfidence(v float32) *Face {
	s.HatConfidence = &v
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

func (s *Face) SetMask(v string) *Face {
	s.Mask = &v
	return s
}

func (s *Face) SetMaskConfidence(v float32) *Face {
	s.MaskConfidence = &v
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

func (s *Face) SetLeftEye(v string) *Face {
	s.LeftEye = &v
	return s
}

func (s *Face) SetLeftEyeConfidence(v float32) *Face {
	s.LeftEyeConfidence = &v
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

func (s *Face) SetHeadPose(v *HeadPose) *Face {
	s.HeadPose = v
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

type File struct {
	// OwnerId
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// ObjectType
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// ObjectId
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// UpdateTime
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// CreateTime
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// Filename
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// MediaType
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// ContentType
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Size
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// FileHash
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// FileModifiedTime
	FileModifiedTime *string `json:"FileModifiedTime,omitempty" xml:"FileModifiedTime,omitempty"`
	// FileCreateTime
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// FileAccessTime
	FileAccessTime *string `json:"FileAccessTime,omitempty" xml:"FileAccessTime,omitempty"`
	// ProduceTime
	ProduceTime *string `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	// LatLong
	LatLong *string `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	// Timezone
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// Addresses
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// TravelClusterId
	TravelClusterId *string `json:"TravelClusterId,omitempty" xml:"TravelClusterId,omitempty"`
	// Orientation
	Orientation *string `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	// Faces
	Faces []*Face `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	// FaceCount
	FaceCount *int64 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// Labels
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// ImageWidth
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// ImageHeight
	ImageHeight *int64 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// EXIF
	EXIF       *string     `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	ImageScore *ImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	// CroppingSuggestions
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	// OCRContents
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
	// ImageEmbeddingsFloat32
	ImageEmbeddingsFloat32 []*float32 `json:"ImageEmbeddingsFloat32,omitempty" xml:"ImageEmbeddingsFloat32,omitempty" type:"Repeated"`
	// ImageEmbeddingsInt8
	ImageEmbeddingsInt8 []*int32 `json:"ImageEmbeddingsInt8,omitempty" xml:"ImageEmbeddingsInt8,omitempty" type:"Repeated"`
	// VideoWidth
	VideoWidth *int64 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
	// VideoHeight
	VideoHeight *int64 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// VideoTakenTime
	VideoTakenTime *string `json:"VideoTakenTime,omitempty" xml:"VideoTakenTime,omitempty"`
	// VideoDuration
	VideoDuration *float32 `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	// VideoBitrate
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// VideoStartTime
	VideoStartTime *float32 `json:"VideoStartTime,omitempty" xml:"VideoStartTime,omitempty"`
	// VideoStreams
	VideoStreams []*VideoStream `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	// Subtitles
	Subtitles []*SubtitleStream `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// VideoEmbeddingsFloat32
	VideoEmbeddingsFloat32 []*float32 `json:"VideoEmbeddingsFloat32,omitempty" xml:"VideoEmbeddingsFloat32,omitempty" type:"Repeated"`
	// VideoEmbeddingsInt8
	VideoEmbeddingsInt8 []*int32 `json:"VideoEmbeddingsInt8,omitempty" xml:"VideoEmbeddingsInt8,omitempty" type:"Repeated"`
	// AudioTakenTime
	AudioTakenTime *string `json:"AudioTakenTime,omitempty" xml:"AudioTakenTime,omitempty"`
	// AudioDuration
	AudioDuration *float32 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	// AudioBitrate
	AudioBitrate *float32 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// AudioStreams
	AudioStreams []*AudioStream `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	// Artists
	Artists []*string `json:"Artists,omitempty" xml:"Artists,omitempty" type:"Repeated"`
	// AudioCovers
	AudioCovers []*Image `json:"AudioCovers,omitempty" xml:"AudioCovers,omitempty" type:"Repeated"`
	// Composer
	Composer *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	// Performer
	Performer *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	// AudioLanguage
	AudioLanguage *string `json:"AudioLanguage,omitempty" xml:"AudioLanguage,omitempty"`
	// Album
	Album *string `json:"Album,omitempty" xml:"Album,omitempty"`
	// AlbumArtist
	AlbumArtist *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	// AudioEmbeddingsFloat32
	AudioEmbeddingsFloat32 []*float32 `json:"AudioEmbeddingsFloat32,omitempty" xml:"AudioEmbeddingsFloat32,omitempty" type:"Repeated"`
	// AudioEmbeddingsInt8
	AudioEmbeddingsInt8 []*int32 `json:"AudioEmbeddingsInt8,omitempty" xml:"AudioEmbeddingsInt8,omitempty" type:"Repeated"`
	// DocumentLanguage
	DocumentLanguage *string `json:"DocumentLanguage,omitempty" xml:"DocumentLanguage,omitempty"`
	// PageCount
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// DocumentContent
	DocumentContent *string `json:"DocumentContent,omitempty" xml:"DocumentContent,omitempty"`
	// DocumentEmbeddingsFloat32
	DocumentEmbeddingsFloat32 []*float32 `json:"DocumentEmbeddingsFloat32,omitempty" xml:"DocumentEmbeddingsFloat32,omitempty" type:"Repeated"`
	// DocumentEmbeddingsInt8
	DocumentEmbeddingsInt8 []*int32 `json:"DocumentEmbeddingsInt8,omitempty" xml:"DocumentEmbeddingsInt8,omitempty" type:"Repeated"`
	// ETag
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// CacheControl
	CacheControl *string `json:"CacheControl,omitempty" xml:"CacheControl,omitempty"`
	// ContentDisposition
	ContentDisposition *string `json:"ContentDisposition,omitempty" xml:"ContentDisposition,omitempty"`
	// ContentEncoding
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// ContentLanguage
	ContentLanguage *string `json:"ContentLanguage,omitempty" xml:"ContentLanguage,omitempty"`
	// AccessControlAllowOrigin
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	// AccessControlRequestMethod
	AccessControlRequestMethod *string `json:"AccessControlRequestMethod,omitempty" xml:"AccessControlRequestMethod,omitempty"`
	// ServerSideEncryptionCustomerAlgorithm
	ServerSideEncryptionCustomerAlgorithm *string `json:"ServerSideEncryptionCustomerAlgorithm,omitempty" xml:"ServerSideEncryptionCustomerAlgorithm,omitempty"`
	// ServerSideEncryption
	ServerSideEncryption *string `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	// ServerSideDataEncryption
	ServerSideDataEncryption *string `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	// ServerSideEncryptionKeyId
	ServerSideEncryptionKeyId *string `json:"ServerSideEncryptionKeyId,omitempty" xml:"ServerSideEncryptionKeyId,omitempty"`
	// StorageClass
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// ObjectACL
	ObjectACL *string `json:"ObjectACL,omitempty" xml:"ObjectACL,omitempty"`
	// ContentMd5
	ContentMd5 *string `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	// OSSUserMeta
	OSSUserMeta map[string]interface{} `json:"OSSUserMeta,omitempty" xml:"OSSUserMeta,omitempty"`
	// OSSTaggingCount
	OSSTaggingCount *int64 `json:"OSSTaggingCount,omitempty" xml:"OSSTaggingCount,omitempty"`
	// OSSTagging
	OSSTagging map[string]interface{} `json:"OSSTagging,omitempty" xml:"OSSTagging,omitempty"`
	// OSSExpiration
	OSSExpiration *string `json:"OSSExpiration,omitempty" xml:"OSSExpiration,omitempty"`
	// OSSVersionId
	OSSVersionId *string `json:"OSSVersionId,omitempty" xml:"OSSVersionId,omitempty"`
	// OSSDeleteMarker
	OSSDeleteMarker *string `json:"OSSDeleteMarker,omitempty" xml:"OSSDeleteMarker,omitempty"`
	// OSSObjectType
	OSSObjectType *string `json:"OSSObjectType,omitempty" xml:"OSSObjectType,omitempty"`
	// CustomId
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// CustomLabels
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s File) String() string {
	return tea.Prettify(s)
}

func (s File) GoString() string {
	return s.String()
}

func (s *File) SetOwnerId(v string) *File {
	s.OwnerId = &v
	return s
}

func (s *File) SetProjectName(v string) *File {
	s.ProjectName = &v
	return s
}

func (s *File) SetDatasetName(v string) *File {
	s.DatasetName = &v
	return s
}

func (s *File) SetObjectType(v string) *File {
	s.ObjectType = &v
	return s
}

func (s *File) SetObjectId(v string) *File {
	s.ObjectId = &v
	return s
}

func (s *File) SetUpdateTime(v string) *File {
	s.UpdateTime = &v
	return s
}

func (s *File) SetCreateTime(v string) *File {
	s.CreateTime = &v
	return s
}

func (s *File) SetURI(v string) *File {
	s.URI = &v
	return s
}

func (s *File) SetFilename(v string) *File {
	s.Filename = &v
	return s
}

func (s *File) SetMediaType(v string) *File {
	s.MediaType = &v
	return s
}

func (s *File) SetContentType(v string) *File {
	s.ContentType = &v
	return s
}

func (s *File) SetSize(v int64) *File {
	s.Size = &v
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

func (s *File) SetFileCreateTime(v string) *File {
	s.FileCreateTime = &v
	return s
}

func (s *File) SetFileAccessTime(v string) *File {
	s.FileAccessTime = &v
	return s
}

func (s *File) SetProduceTime(v string) *File {
	s.ProduceTime = &v
	return s
}

func (s *File) SetLatLong(v string) *File {
	s.LatLong = &v
	return s
}

func (s *File) SetTimezone(v string) *File {
	s.Timezone = &v
	return s
}

func (s *File) SetAddresses(v []*Address) *File {
	s.Addresses = v
	return s
}

func (s *File) SetTravelClusterId(v string) *File {
	s.TravelClusterId = &v
	return s
}

func (s *File) SetOrientation(v string) *File {
	s.Orientation = &v
	return s
}

func (s *File) SetFaces(v []*Face) *File {
	s.Faces = v
	return s
}

func (s *File) SetFaceCount(v int64) *File {
	s.FaceCount = &v
	return s
}

func (s *File) SetLabels(v []*Label) *File {
	s.Labels = v
	return s
}

func (s *File) SetTitle(v string) *File {
	s.Title = &v
	return s
}

func (s *File) SetImageWidth(v int64) *File {
	s.ImageWidth = &v
	return s
}

func (s *File) SetImageHeight(v int64) *File {
	s.ImageHeight = &v
	return s
}

func (s *File) SetEXIF(v string) *File {
	s.EXIF = &v
	return s
}

func (s *File) SetImageScore(v *ImageScore) *File {
	s.ImageScore = v
	return s
}

func (s *File) SetCroppingSuggestions(v []*CroppingSuggestion) *File {
	s.CroppingSuggestions = v
	return s
}

func (s *File) SetOCRContents(v []*OCRContents) *File {
	s.OCRContents = v
	return s
}

func (s *File) SetImageEmbeddingsFloat32(v []*float32) *File {
	s.ImageEmbeddingsFloat32 = v
	return s
}

func (s *File) SetImageEmbeddingsInt8(v []*int32) *File {
	s.ImageEmbeddingsInt8 = v
	return s
}

func (s *File) SetVideoWidth(v int64) *File {
	s.VideoWidth = &v
	return s
}

func (s *File) SetVideoHeight(v int64) *File {
	s.VideoHeight = &v
	return s
}

func (s *File) SetVideoTakenTime(v string) *File {
	s.VideoTakenTime = &v
	return s
}

func (s *File) SetVideoDuration(v float32) *File {
	s.VideoDuration = &v
	return s
}

func (s *File) SetVideoBitrate(v int64) *File {
	s.VideoBitrate = &v
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

func (s *File) SetSubtitles(v []*SubtitleStream) *File {
	s.Subtitles = v
	return s
}

func (s *File) SetVideoEmbeddingsFloat32(v []*float32) *File {
	s.VideoEmbeddingsFloat32 = v
	return s
}

func (s *File) SetVideoEmbeddingsInt8(v []*int32) *File {
	s.VideoEmbeddingsInt8 = v
	return s
}

func (s *File) SetAudioTakenTime(v string) *File {
	s.AudioTakenTime = &v
	return s
}

func (s *File) SetAudioDuration(v float32) *File {
	s.AudioDuration = &v
	return s
}

func (s *File) SetAudioBitrate(v float32) *File {
	s.AudioBitrate = &v
	return s
}

func (s *File) SetAudioStreams(v []*AudioStream) *File {
	s.AudioStreams = v
	return s
}

func (s *File) SetArtists(v []*string) *File {
	s.Artists = v
	return s
}

func (s *File) SetAudioCovers(v []*Image) *File {
	s.AudioCovers = v
	return s
}

func (s *File) SetComposer(v string) *File {
	s.Composer = &v
	return s
}

func (s *File) SetPerformer(v string) *File {
	s.Performer = &v
	return s
}

func (s *File) SetAudioLanguage(v string) *File {
	s.AudioLanguage = &v
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

func (s *File) SetAudioEmbeddingsFloat32(v []*float32) *File {
	s.AudioEmbeddingsFloat32 = v
	return s
}

func (s *File) SetAudioEmbeddingsInt8(v []*int32) *File {
	s.AudioEmbeddingsInt8 = v
	return s
}

func (s *File) SetDocumentLanguage(v string) *File {
	s.DocumentLanguage = &v
	return s
}

func (s *File) SetPageCount(v int64) *File {
	s.PageCount = &v
	return s
}

func (s *File) SetDocumentContent(v string) *File {
	s.DocumentContent = &v
	return s
}

func (s *File) SetDocumentEmbeddingsFloat32(v []*float32) *File {
	s.DocumentEmbeddingsFloat32 = v
	return s
}

func (s *File) SetDocumentEmbeddingsInt8(v []*int32) *File {
	s.DocumentEmbeddingsInt8 = v
	return s
}

func (s *File) SetETag(v string) *File {
	s.ETag = &v
	return s
}

func (s *File) SetCacheControl(v string) *File {
	s.CacheControl = &v
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

func (s *File) SetAccessControlAllowOrigin(v string) *File {
	s.AccessControlAllowOrigin = &v
	return s
}

func (s *File) SetAccessControlRequestMethod(v string) *File {
	s.AccessControlRequestMethod = &v
	return s
}

func (s *File) SetServerSideEncryptionCustomerAlgorithm(v string) *File {
	s.ServerSideEncryptionCustomerAlgorithm = &v
	return s
}

func (s *File) SetServerSideEncryption(v string) *File {
	s.ServerSideEncryption = &v
	return s
}

func (s *File) SetServerSideDataEncryption(v string) *File {
	s.ServerSideDataEncryption = &v
	return s
}

func (s *File) SetServerSideEncryptionKeyId(v string) *File {
	s.ServerSideEncryptionKeyId = &v
	return s
}

func (s *File) SetStorageClass(v string) *File {
	s.StorageClass = &v
	return s
}

func (s *File) SetObjectACL(v string) *File {
	s.ObjectACL = &v
	return s
}

func (s *File) SetContentMd5(v string) *File {
	s.ContentMd5 = &v
	return s
}

func (s *File) SetOSSUserMeta(v map[string]interface{}) *File {
	s.OSSUserMeta = v
	return s
}

func (s *File) SetOSSTaggingCount(v int64) *File {
	s.OSSTaggingCount = &v
	return s
}

func (s *File) SetOSSTagging(v map[string]interface{}) *File {
	s.OSSTagging = v
	return s
}

func (s *File) SetOSSExpiration(v string) *File {
	s.OSSExpiration = &v
	return s
}

func (s *File) SetOSSVersionId(v string) *File {
	s.OSSVersionId = &v
	return s
}

func (s *File) SetOSSDeleteMarker(v string) *File {
	s.OSSDeleteMarker = &v
	return s
}

func (s *File) SetOSSObjectType(v string) *File {
	s.OSSObjectType = &v
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

type WebofficeUser struct {
	// Id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 头像
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
}

func (s WebofficeUser) String() string {
	return tea.Prettify(s)
}

func (s WebofficeUser) GoString() string {
	return s.String()
}

func (s *WebofficeUser) SetId(v string) *WebofficeUser {
	s.Id = &v
	return s
}

func (s *WebofficeUser) SetName(v string) *WebofficeUser {
	s.Name = &v
	return s
}

func (s *WebofficeUser) SetAvatar(v string) *WebofficeUser {
	s.Avatar = &v
	return s
}

type AssumeRoleChain struct {
	// 当前用户 policy
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 链式授权节点
	Chain []*AssumeRoleChainNode `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
}

func (s AssumeRoleChain) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleChain) GoString() string {
	return s.String()
}

func (s *AssumeRoleChain) SetPolicy(v string) *AssumeRoleChain {
	s.Policy = &v
	return s
}

func (s *AssumeRoleChain) SetChain(v []*AssumeRoleChainNode) *AssumeRoleChain {
	s.Chain = v
	return s
}

type WebofficeWatermark struct {
	// 水印类型，目前仅支持文字水印，0: 无水印；1: 文字水印
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 水印文字
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 旋转角度
	Rotate *float32 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// 垂直间距
	Vertical *int64 `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
	// 水平间距
	Horizontal *int64 `json:"Horizontal,omitempty" xml:"Horizontal,omitempty"`
	// 字体样式
	Font *string `json:"Font,omitempty" xml:"Font,omitempty"`
	// 字体颜色
	FillStyle *string `json:"FillStyle,omitempty" xml:"FillStyle,omitempty"`
}

func (s WebofficeWatermark) String() string {
	return tea.Prettify(s)
}

func (s WebofficeWatermark) GoString() string {
	return s.String()
}

func (s *WebofficeWatermark) SetType(v int64) *WebofficeWatermark {
	s.Type = &v
	return s
}

func (s *WebofficeWatermark) SetValue(v string) *WebofficeWatermark {
	s.Value = &v
	return s
}

func (s *WebofficeWatermark) SetRotate(v float32) *WebofficeWatermark {
	s.Rotate = &v
	return s
}

func (s *WebofficeWatermark) SetVertical(v int64) *WebofficeWatermark {
	s.Vertical = &v
	return s
}

func (s *WebofficeWatermark) SetHorizontal(v int64) *WebofficeWatermark {
	s.Horizontal = &v
	return s
}

func (s *WebofficeWatermark) SetFont(v string) *WebofficeWatermark {
	s.Font = &v
	return s
}

func (s *WebofficeWatermark) SetFillStyle(v string) *WebofficeWatermark {
	s.FillStyle = &v
	return s
}

type AssumeRoleChainNode struct {
	// 账号类型，普通账号填 user，服务账号填 service
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 账号id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 授权角色名
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AssumeRoleChainNode) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleChainNode) GoString() string {
	return s.String()
}

func (s *AssumeRoleChainNode) SetType(v string) *AssumeRoleChainNode {
	s.Type = &v
	return s
}

func (s *AssumeRoleChainNode) SetOwnerId(v string) *AssumeRoleChainNode {
	s.OwnerId = &v
	return s
}

func (s *AssumeRoleChainNode) SetRole(v string) *AssumeRoleChainNode {
	s.Role = &v
	return s
}

type WebofficePermission struct {
	// 重命名
	Rename *bool `json:"Rename,omitempty" xml:"Rename,omitempty"`
	// 只读模式
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// 查看历史版本
	History *bool `json:"History,omitempty" xml:"History,omitempty"`
	// 打印
	Print *bool `json:"Print,omitempty" xml:"Print,omitempty"`
	// 导出
	Export *bool `json:"Export,omitempty" xml:"Export,omitempty"`
	// 拷贝
	Copy *bool `json:"Copy,omitempty" xml:"Copy,omitempty"`
}

func (s WebofficePermission) String() string {
	return tea.Prettify(s)
}

func (s WebofficePermission) GoString() string {
	return s.String()
}

func (s *WebofficePermission) SetRename(v bool) *WebofficePermission {
	s.Rename = &v
	return s
}

func (s *WebofficePermission) SetReadonly(v bool) *WebofficePermission {
	s.Readonly = &v
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

func (s *WebofficePermission) SetExport(v bool) *WebofficePermission {
	s.Export = &v
	return s
}

func (s *WebofficePermission) SetCopy(v bool) *WebofficePermission {
	s.Copy = &v
	return s
}

type SimpleQuery struct {
	// 需要查询的字段名
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// 需要查询的字段值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 运算符
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 由 SimpleQuery 结构体组成的子查询数组
	SubQueries []*SimpleQuery `json:"SubQueries,omitempty" xml:"SubQueries,omitempty" type:"Repeated"`
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

func (s *SimpleQuery) SetValue(v string) *SimpleQuery {
	s.Value = &v
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

type BatchDeleteFileMetaRequest struct {
	ProjectName *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URIs        []*string `json:"URIs,omitempty" xml:"URIs,omitempty" type:"Repeated"`
}

func (s BatchDeleteFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaRequest) SetProjectName(v string) *BatchDeleteFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchDeleteFileMetaRequest) SetDatasetName(v string) *BatchDeleteFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchDeleteFileMetaRequest) SetURIs(v []*string) *BatchDeleteFileMetaRequest {
	s.URIs = v
	return s
}

type BatchDeleteFileMetaShrinkRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URIsShrink  *string `json:"URIs,omitempty" xml:"URIs,omitempty"`
}

func (s BatchDeleteFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteFileMetaShrinkRequest) SetProjectName(v string) *BatchDeleteFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchDeleteFileMetaShrinkRequest) SetDatasetName(v string) *BatchDeleteFileMetaShrinkRequest {
	s.DatasetName = &v
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
	ProjectName *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URIs        []*string `json:"URIs,omitempty" xml:"URIs,omitempty" type:"Repeated"`
}

func (s BatchGetFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaRequest) SetProjectName(v string) *BatchGetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFileMetaRequest) SetDatasetName(v string) *BatchGetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFileMetaRequest) SetURIs(v []*string) *BatchGetFileMetaRequest {
	s.URIs = v
	return s
}

type BatchGetFileMetaShrinkRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URIsShrink  *string `json:"URIs,omitempty" xml:"URIs,omitempty"`
}

func (s BatchGetFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaShrinkRequest) SetProjectName(v string) *BatchGetFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetDatasetName(v string) *BatchGetFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFileMetaShrinkRequest) SetURIsShrink(v string) *BatchGetFileMetaShrinkRequest {
	s.URIsShrink = &v
	return s
}

type BatchGetFileMetaResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Files     []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s BatchGetFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFileMetaResponseBody) SetRequestId(v string) *BatchGetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetFileMetaResponseBody) SetFiles(v []*File) *BatchGetFileMetaResponseBody {
	s.Files = v
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
	ProjectName *string                           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string                           `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files       []*BatchIndexFileMetaRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s BatchIndexFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaRequest) SetProjectName(v string) *BatchIndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetDatasetName(v string) *BatchIndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetFiles(v []*BatchIndexFileMetaRequestFiles) *BatchIndexFileMetaRequest {
	s.Files = v
	return s
}

type BatchIndexFileMetaRequestFiles struct {
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s BatchIndexFileMetaRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaRequestFiles) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaRequestFiles) SetURI(v string) *BatchIndexFileMetaRequestFiles {
	s.URI = &v
	return s
}

func (s *BatchIndexFileMetaRequestFiles) SetCustomLabels(v map[string]interface{}) *BatchIndexFileMetaRequestFiles {
	s.CustomLabels = v
	return s
}

type BatchIndexFileMetaShrinkRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
}

func (s BatchIndexFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaShrinkRequest) SetProjectName(v string) *BatchIndexFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetDatasetName(v string) *BatchIndexFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetFilesShrink(v string) *BatchIndexFileMetaShrinkRequest {
	s.FilesShrink = &v
	return s
}

type BatchIndexFileMetaResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchIndexFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchIndexFileMetaResponseBody) GoString() string {
	return s.String()
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
	ProjectName *string                            `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string                            `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files       []*BatchUpdateFileMetaRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s BatchUpdateFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaRequest) SetProjectName(v string) *BatchUpdateFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchUpdateFileMetaRequest) SetDatasetName(v string) *BatchUpdateFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchUpdateFileMetaRequest) SetFiles(v []*BatchUpdateFileMetaRequestFiles) *BatchUpdateFileMetaRequest {
	s.Files = v
	return s
}

type BatchUpdateFileMetaRequestFiles struct {
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s BatchUpdateFileMetaRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaRequestFiles) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaRequestFiles) SetURI(v string) *BatchUpdateFileMetaRequestFiles {
	s.URI = &v
	return s
}

func (s *BatchUpdateFileMetaRequestFiles) SetCustomLabels(v map[string]interface{}) *BatchUpdateFileMetaRequestFiles {
	s.CustomLabels = v
	return s
}

type BatchUpdateFileMetaShrinkRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
}

func (s BatchUpdateFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaShrinkRequest) SetProjectName(v string) *BatchUpdateFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) SetDatasetName(v string) *BatchUpdateFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchUpdateFileMetaShrinkRequest) SetFilesShrink(v string) *BatchUpdateFileMetaShrinkRequest {
	s.FilesShrink = &v
	return s
}

type BatchUpdateFileMetaResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileMetaResponseBody) SetRequestId(v string) *BatchUpdateFileMetaResponseBody {
	s.RequestId = &v
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

type CreateBindingRequest struct {
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// DatasetName
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// URI
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) SetProjectName(v string) *CreateBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBindingRequest) SetDatasetName(v string) *CreateBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateBindingRequest) SetURI(v string) *CreateBindingRequest {
	s.URI = &v
	return s
}

type CreateBindingResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponseBody) GoString() string {
	return s.String()
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
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 对数据集的描述
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	DatasetMaxOSSBindCount  *int64  `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) SetProjectName(v string) *CreateDatasetRequest {
	s.ProjectName = &v
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

func (s *CreateDatasetRequest) SetTemplateId(v string) *CreateDatasetRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxOSSBindCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxFileCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxEntityCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxEntityCount = &v
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

type CreateDatasetResponseBody struct {
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 数据集创建时间
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetProjectName(v string) *CreateDatasetResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CreateDatasetResponseBody) SetDatasetName(v string) *CreateDatasetResponseBody {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetResponseBody) SetCreateTime(v int64) *CreateDatasetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateDatasetResponseBody) SetDescription(v string) *CreateDatasetResponseBody {
	s.Description = &v
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

type CreateProjectRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 项目描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 服务角色
	ServiceRole             *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ProjectQPS              *int64  `json:"ProjectQPS,omitempty" xml:"ProjectQPS,omitempty"`
	ProjectTPS              *int64  `json:"ProjectTPS,omitempty" xml:"ProjectTPS,omitempty"`
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	DatasetMaxOSSBindCount  *int64  `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
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

func (s *CreateProjectRequest) SetProjectQPS(v int64) *CreateProjectRequest {
	s.ProjectQPS = &v
	return s
}

func (s *CreateProjectRequest) SetProjectTPS(v int64) *CreateProjectRequest {
	s.ProjectTPS = &v
	return s
}

func (s *CreateProjectRequest) SetProjectMaxDatasetCount(v int64) *CreateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxOSSBindCount(v int64) *CreateProjectRequest {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxFileCount(v int64) *CreateProjectRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateProjectRequest) SetDatasetMaxEntityCount(v int64) *CreateProjectRequest {
	s.DatasetMaxEntityCount = &v
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

type CreateProjectResponseBody struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 项目创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 本次请求的唯一 ID
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectName(v string) *CreateProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectResponseBody) SetCreateTime(v int64) *CreateProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetDescription(v string) *CreateProjectResponseBody {
	s.Description = &v
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
	// A short description of struct
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DeleteBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) SetProjectName(v string) *DeleteBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteBindingRequest) SetDatasetName(v string) *DeleteBindingRequest {
	s.DatasetName = &v
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
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) SetProjectName(v string) *DeleteDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteDatasetRequest) SetDatasetName(v string) *DeleteDatasetRequest {
	s.DatasetName = &v
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
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DeleteFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaRequest) SetProjectName(v string) *DeleteFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteFileMetaRequest) SetDatasetName(v string) *DeleteFileMetaRequest {
	s.DatasetName = &v
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

type FuzzyQueryRequest struct {
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Dataset 名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 用于搜索的字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s FuzzyQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FuzzyQueryRequest) GoString() string {
	return s.String()
}

func (s *FuzzyQueryRequest) SetNextToken(v string) *FuzzyQueryRequest {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryRequest) SetMaxResults(v int64) *FuzzyQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *FuzzyQueryRequest) SetProjectName(v string) *FuzzyQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *FuzzyQueryRequest) SetDatasetName(v string) *FuzzyQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *FuzzyQueryRequest) SetQuery(v string) *FuzzyQueryRequest {
	s.Query = &v
	return s
}

type FuzzyQueryResponseBody struct {
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求的唯一 Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Files     []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s FuzzyQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FuzzyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FuzzyQueryResponseBody) SetNextToken(v string) *FuzzyQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *FuzzyQueryResponseBody) SetRequestId(v string) *FuzzyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FuzzyQueryResponseBody) SetFiles(v []*File) *FuzzyQueryResponseBody {
	s.Files = v
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
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBindingRequest) GoString() string {
	return s.String()
}

func (s *GetBindingRequest) SetProjectName(v string) *GetBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *GetBindingRequest) SetDatasetName(v string) *GetBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *GetBindingRequest) SetURI(v string) *GetBindingRequest {
	s.URI = &v
	return s
}

type GetBindingResponseBody struct {
	// Id of the request
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Binding   *GetBindingResponseBodyBinding `json:"Binding,omitempty" xml:"Binding,omitempty" type:"Struct"`
}

func (s GetBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBindingResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindingResponseBody) SetRequestId(v string) *GetBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBindingResponseBody) SetBinding(v *GetBindingResponseBodyBinding) *GetBindingResponseBody {
	s.Binding = v
	return s
}

type GetBindingResponseBodyBinding struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetBindingResponseBodyBinding) String() string {
	return tea.Prettify(s)
}

func (s GetBindingResponseBodyBinding) GoString() string {
	return s.String()
}

func (s *GetBindingResponseBodyBinding) SetProjectName(v string) *GetBindingResponseBodyBinding {
	s.ProjectName = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetDatasetName(v string) *GetBindingResponseBodyBinding {
	s.DatasetName = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetURI(v string) *GetBindingResponseBodyBinding {
	s.URI = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetState(v string) *GetBindingResponseBodyBinding {
	s.State = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetPhase(v string) *GetBindingResponseBodyBinding {
	s.Phase = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetReason(v string) *GetBindingResponseBodyBinding {
	s.Reason = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetCreateTime(v int64) *GetBindingResponseBodyBinding {
	s.CreateTime = &v
	return s
}

func (s *GetBindingResponseBodyBinding) SetUpdateTime(v int64) *GetBindingResponseBodyBinding {
	s.UpdateTime = &v
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
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) SetProjectName(v string) *GetDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDatasetRequest) SetDatasetName(v string) *GetDatasetRequest {
	s.DatasetName = &v
	return s
}

type GetDatasetResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CreateTime              *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime              *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DatasetMaxOSSBindCount  *int64  `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	BindCount               *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	FileCount               *int64  `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	TotalFileSize           *int64  `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetProjectName(v string) *GetDatasetResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetName(v string) *GetDatasetResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetResponseBody) SetTemplateId(v string) *GetDatasetResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetDatasetResponseBody) SetCreateTime(v int64) *GetDatasetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetUpdateTime(v int64) *GetDatasetResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetDescription(v string) *GetDatasetResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetMaxOSSBindCount(v int64) *GetDatasetResponseBody {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetMaxFileCount(v int64) *GetDatasetResponseBody {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetMaxEntityCount(v int64) *GetDatasetResponseBody {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetMaxRelationCount(v int64) *GetDatasetResponseBody {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetMaxTotalFileSize(v int64) *GetDatasetResponseBody {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *GetDatasetResponseBody) SetBindCount(v int64) *GetDatasetResponseBody {
	s.BindCount = &v
	return s
}

func (s *GetDatasetResponseBody) SetFileCount(v int64) *GetDatasetResponseBody {
	s.FileCount = &v
	return s
}

func (s *GetDatasetResponseBody) SetTotalFileSize(v int64) *GetDatasetResponseBody {
	s.TotalFileSize = &v
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

type GetFileMetaRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *GetFileMetaRequest) SetProjectName(v string) *GetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileMetaRequest) SetDatasetName(v string) *GetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFileMetaRequest) SetURI(v string) *GetFileMetaRequest {
	s.URI = &v
	return s
}

type GetFileMetaResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// File list.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s GetFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileMetaResponseBody) SetRequestId(v string) *GetFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileMetaResponseBody) SetFiles(v []*File) *GetFileMetaResponseBody {
	s.Files = v
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

type GetProjectResponseBody struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 项目描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 服务角色
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// 工作流
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 项目创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 项目修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 项目QPS
	ProjectQPS *int64 `json:"ProjectQPS,omitempty" xml:"ProjectQPS,omitempty"`
	// 项目TPS
	ProjectTPS *int64 `json:"ProjectTPS,omitempty" xml:"ProjectTPS,omitempty"`
	// 最大媒体集数量
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// 当前项目每个媒体集最大绑定数
	DatasetMaxOSSBindCount *int64 `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	// 当前项目每个媒体集最大文件数
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 当前项目每个媒体集最大实体数
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 当前项目每个媒体集最大关系数
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 当前项目每个媒体集最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 媒体集数量
	DatasetCount *int64 `json:"DatasetCount,omitempty" xml:"DatasetCount,omitempty"`
	// 项目当前文件数量
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// 项目当前文件总大小
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// 本次请求的唯一 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetProjectName(v string) *GetProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetProjectResponseBody) SetDescription(v string) *GetProjectResponseBody {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBody) SetServiceRole(v string) *GetProjectResponseBody {
	s.ServiceRole = &v
	return s
}

func (s *GetProjectResponseBody) SetTemplateId(v string) *GetProjectResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetProjectResponseBody) SetCreateTime(v int64) *GetProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetUpdateTime(v int64) *GetProjectResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectQPS(v int64) *GetProjectResponseBody {
	s.ProjectQPS = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectTPS(v int64) *GetProjectResponseBody {
	s.ProjectTPS = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectMaxDatasetCount(v int64) *GetProjectResponseBody {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *GetProjectResponseBody) SetDatasetMaxOSSBindCount(v int64) *GetProjectResponseBody {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *GetProjectResponseBody) SetDatasetMaxFileCount(v int64) *GetProjectResponseBody {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *GetProjectResponseBody) SetDatasetMaxEntityCount(v int64) *GetProjectResponseBody {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *GetProjectResponseBody) SetDatasetMaxRelationCount(v int64) *GetProjectResponseBody {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *GetProjectResponseBody) SetDatasetMaxTotalFileSize(v int64) *GetProjectResponseBody {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *GetProjectResponseBody) SetDatasetCount(v int64) *GetProjectResponseBody {
	s.DatasetCount = &v
	return s
}

func (s *GetProjectResponseBody) SetFileCount(v int64) *GetProjectResponseBody {
	s.FileCount = &v
	return s
}

func (s *GetProjectResponseBody) SetTotalFileSize(v int64) *GetProjectResponseBody {
	s.TotalFileSize = &v
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

type GetWebofficeUrlRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 预览编辑地址
	SourceUri *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	// 文件名，必须带文件名后缀，默认是 SourceUri 的最后一级
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// 用户自定义数据，在消息通知中返回
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 预览前几页
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// 文件密码
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 是否支持外部上传
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// mns 消息通知地址
	NotifyEndpoint *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	// mns 消息通知 topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// 隐藏工具栏，预览模式下使用
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// 权限
	Permission *WebofficePermission `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// 用户
	User *WebofficeUser `json:"User,omitempty" xml:"User,omitempty"`
	// 水印
	Watermark *WebofficeWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// 链式授权
	AssumeRoleChain *AssumeRoleChain `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
}

func (s GetWebofficeUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeUrlRequest) SetProjectName(v string) *GetWebofficeUrlRequest {
	s.ProjectName = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetSourceUri(v string) *GetWebofficeUrlRequest {
	s.SourceUri = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetFilename(v string) *GetWebofficeUrlRequest {
	s.Filename = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetUserData(v string) *GetWebofficeUrlRequest {
	s.UserData = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetPreviewPages(v int64) *GetWebofficeUrlRequest {
	s.PreviewPages = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetPassword(v string) *GetWebofficeUrlRequest {
	s.Password = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetExternalUploaded(v bool) *GetWebofficeUrlRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetNotifyEndpoint(v string) *GetWebofficeUrlRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetNotifyTopicName(v string) *GetWebofficeUrlRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetHidecmb(v bool) *GetWebofficeUrlRequest {
	s.Hidecmb = &v
	return s
}

func (s *GetWebofficeUrlRequest) SetPermission(v *WebofficePermission) *GetWebofficeUrlRequest {
	s.Permission = v
	return s
}

func (s *GetWebofficeUrlRequest) SetUser(v *WebofficeUser) *GetWebofficeUrlRequest {
	s.User = v
	return s
}

func (s *GetWebofficeUrlRequest) SetWatermark(v *WebofficeWatermark) *GetWebofficeUrlRequest {
	s.Watermark = v
	return s
}

func (s *GetWebofficeUrlRequest) SetAssumeRoleChain(v *AssumeRoleChain) *GetWebofficeUrlRequest {
	s.AssumeRoleChain = v
	return s
}

type GetWebofficeUrlShrinkRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 预览编辑地址
	SourceUri *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	// 文件名，必须带文件名后缀，默认是 SourceUri 的最后一级
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// 用户自定义数据，在消息通知中返回
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 预览前几页
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// 文件密码
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 是否支持外部上传
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// mns 消息通知地址
	NotifyEndpoint *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	// mns 消息通知 topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// 隐藏工具栏，预览模式下使用
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// 权限
	PermissionShrink *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// 用户
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
	// 水印
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	// 链式授权
	AssumeRoleChainShrink *string `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
}

func (s GetWebofficeUrlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWebofficeUrlShrinkRequest) SetProjectName(v string) *GetWebofficeUrlShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetSourceUri(v string) *GetWebofficeUrlShrinkRequest {
	s.SourceUri = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetFilename(v string) *GetWebofficeUrlShrinkRequest {
	s.Filename = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetUserData(v string) *GetWebofficeUrlShrinkRequest {
	s.UserData = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetPreviewPages(v int64) *GetWebofficeUrlShrinkRequest {
	s.PreviewPages = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetPassword(v string) *GetWebofficeUrlShrinkRequest {
	s.Password = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetExternalUploaded(v bool) *GetWebofficeUrlShrinkRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetNotifyEndpoint(v string) *GetWebofficeUrlShrinkRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetNotifyTopicName(v string) *GetWebofficeUrlShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetHidecmb(v bool) *GetWebofficeUrlShrinkRequest {
	s.Hidecmb = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetPermissionShrink(v string) *GetWebofficeUrlShrinkRequest {
	s.PermissionShrink = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetUserShrink(v string) *GetWebofficeUrlShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetWatermarkShrink(v string) *GetWebofficeUrlShrinkRequest {
	s.WatermarkShrink = &v
	return s
}

func (s *GetWebofficeUrlShrinkRequest) SetAssumeRoleChainShrink(v string) *GetWebofficeUrlShrinkRequest {
	s.AssumeRoleChainShrink = &v
	return s
}

type GetWebofficeUrlResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 预览编辑地址
	WebofficeUrl *string `json:"WebofficeUrl,omitempty" xml:"WebofficeUrl,omitempty"`
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// access token 过期时间
	AccessTokenExpiredTime *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// refresh token 过期时间
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
}

func (s GetWebofficeUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebofficeUrlResponseBody) SetRequestId(v string) *GetWebofficeUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebofficeUrlResponseBody) SetWebofficeUrl(v string) *GetWebofficeUrlResponseBody {
	s.WebofficeUrl = &v
	return s
}

func (s *GetWebofficeUrlResponseBody) SetAccessToken(v string) *GetWebofficeUrlResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetWebofficeUrlResponseBody) SetRefreshToken(v string) *GetWebofficeUrlResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetWebofficeUrlResponseBody) SetAccessTokenExpiredTime(v string) *GetWebofficeUrlResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetWebofficeUrlResponseBody) SetRefreshTokenExpiredTime(v string) *GetWebofficeUrlResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

type GetWebofficeUrlResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebofficeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebofficeUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebofficeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetWebofficeUrlResponse) SetHeaders(v map[string]*string) *GetWebofficeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetWebofficeUrlResponse) SetBody(v *GetWebofficeUrlResponseBody) *GetWebofficeUrlResponse {
	s.Body = v
	return s
}

type IndexFileMetaRequest struct {
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s IndexFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *IndexFileMetaRequest) SetProjectName(v string) *IndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *IndexFileMetaRequest) SetDatasetName(v string) *IndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *IndexFileMetaRequest) SetURI(v string) *IndexFileMetaRequest {
	s.URI = &v
	return s
}

func (s *IndexFileMetaRequest) SetCustomLabels(v map[string]interface{}) *IndexFileMetaRequest {
	s.CustomLabels = v
	return s
}

type IndexFileMetaShrinkRequest struct {
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI                *string `json:"URI,omitempty" xml:"URI,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s IndexFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *IndexFileMetaShrinkRequest) SetProjectName(v string) *IndexFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetDatasetName(v string) *IndexFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetURI(v string) *IndexFileMetaShrinkRequest {
	s.URI = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetCustomLabelsShrink(v string) *IndexFileMetaShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

type IndexFileMetaResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IndexFileMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndexFileMetaResponseBody) GoString() string {
	return s.String()
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
	// A short description of struct
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListBindingsRequest) SetProjectName(v string) *ListBindingsRequest {
	s.ProjectName = &v
	return s
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

type ListBindingsResponseBody struct {
	// Id of the request
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Bindings  []*ListBindingsResponseBodyBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
}

func (s ListBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindingsResponseBody) SetNextToken(v string) *ListBindingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBindingsResponseBody) SetBindings(v []*ListBindingsResponseBodyBindings) *ListBindingsResponseBody {
	s.Bindings = v
	return s
}

type ListBindingsResponseBodyBindings struct {
	Binding *ListBindingsResponseBodyBindingsBinding `json:"Binding,omitempty" xml:"Binding,omitempty" type:"Struct"`
}

func (s ListBindingsResponseBodyBindings) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyBindings) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyBindings) SetBinding(v *ListBindingsResponseBodyBindingsBinding) *ListBindingsResponseBodyBindings {
	s.Binding = v
	return s
}

type ListBindingsResponseBodyBindingsBinding struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListBindingsResponseBodyBindingsBinding) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyBindingsBinding) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyBindingsBinding) SetProjectName(v string) *ListBindingsResponseBodyBindingsBinding {
	s.ProjectName = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetDatasetName(v string) *ListBindingsResponseBodyBindingsBinding {
	s.DatasetName = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetURI(v string) *ListBindingsResponseBodyBindingsBinding {
	s.URI = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetState(v string) *ListBindingsResponseBodyBindingsBinding {
	s.State = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetPhase(v string) *ListBindingsResponseBodyBindingsBinding {
	s.Phase = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetReason(v string) *ListBindingsResponseBodyBindingsBinding {
	s.Reason = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetCreateTime(v int64) *ListBindingsResponseBodyBindingsBinding {
	s.CreateTime = &v
	return s
}

func (s *ListBindingsResponseBodyBindingsBinding) SetUpdateTime(v int64) *ListBindingsResponseBodyBindingsBinding {
	s.UpdateTime = &v
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
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 返回最大个数
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) SetProjectName(v string) *ListDatasetsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListDatasetsRequest) SetMaxResults(v int64) *ListDatasetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetsRequest) SetNextToken(v string) *ListDatasetsRequest {
	s.NextToken = &v
	return s
}

type ListDatasetsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Responses
	Datasets []*ListDatasetsResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
}

func (s ListDatasetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) SetNextToken(v string) *ListDatasetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsResponseBody) SetDatasets(v []*ListDatasetsResponseBodyDatasets) *ListDatasetsResponseBody {
	s.Datasets = v
	return s
}

type ListDatasetsResponseBodyDatasets struct {
	// ProjectName
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// DatasetName
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Workflow                *string `json:"Workflow,omitempty" xml:"Workflow,omitempty"`
	CreateTime              *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime              *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DatasetMaxOSSBindCount  *int64  `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	BindCount               *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
}

func (s ListDatasetsResponseBodyDatasets) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyDatasets) SetProjectName(v string) *ListDatasetsResponseBodyDatasets {
	s.ProjectName = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetName(v string) *ListDatasetsResponseBodyDatasets {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetWorkflow(v string) *ListDatasetsResponseBodyDatasets {
	s.Workflow = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetCreateTime(v int64) *ListDatasetsResponseBodyDatasets {
	s.CreateTime = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetUpdateTime(v int64) *ListDatasetsResponseBodyDatasets {
	s.UpdateTime = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDescription(v string) *ListDatasetsResponseBodyDatasets {
	s.Description = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetMaxOSSBindCount(v int64) *ListDatasetsResponseBodyDatasets {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetMaxFileCount(v int64) *ListDatasetsResponseBodyDatasets {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetMaxEntityCount(v int64) *ListDatasetsResponseBodyDatasets {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetMaxRelationCount(v int64) *ListDatasetsResponseBodyDatasets {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetMaxTotalFileSize(v int64) *ListDatasetsResponseBodyDatasets {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetBindCount(v int64) *ListDatasetsResponseBodyDatasets {
	s.BindCount = &v
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

type ListProjectsRequest struct {
	// 返回结果的最大个数
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 当总结果个数大于MaxResults时，用于翻页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

type ListProjectsResponseBody struct {
	// 当总结果个数大于MaxResults时，用于翻页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 由ProjectItem组成的数组
	Projects []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
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

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyProjects struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 服务角色
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// 工作流
	Workflow *string `json:"Workflow,omitempty" xml:"Workflow,omitempty"`
	// 项目描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 项目创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 项目上次修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 项目QPS
	ProjectQPS *int64 `json:"ProjectQPS,omitempty" xml:"ProjectQPS,omitempty"`
	// 项目TPS
	ProjectTPS *int64 `json:"ProjectTPS,omitempty" xml:"ProjectTPS,omitempty"`
	// 最大媒体集数
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// 项目下每个媒体集最多绑定数
	DatasetMaxOSSBindCount *int64 `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	// 项目下每个媒体集最大文件数量
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// 项目下每个媒体集最大实体数
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// 项目下每个媒体集最大关系数
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// 项目下每个媒体集最大文件总大小
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// 媒体集数量
	DatasetCount *int64 `json:"DatasetCount,omitempty" xml:"DatasetCount,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetProjectName(v string) *ListProjectsResponseBodyProjects {
	s.ProjectName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetServiceRole(v string) *ListProjectsResponseBodyProjects {
	s.ServiceRole = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetWorkflow(v string) *ListProjectsResponseBodyProjects {
	s.Workflow = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreateTime(v int64) *ListProjectsResponseBodyProjects {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetUpdateTime(v int64) *ListProjectsResponseBodyProjects {
	s.UpdateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectQPS(v int64) *ListProjectsResponseBodyProjects {
	s.ProjectQPS = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectTPS(v int64) *ListProjectsResponseBodyProjects {
	s.ProjectTPS = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectMaxDatasetCount(v int64) *ListProjectsResponseBodyProjects {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDatasetMaxOSSBindCount(v int64) *ListProjectsResponseBodyProjects {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDatasetMaxFileCount(v int64) *ListProjectsResponseBodyProjects {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDatasetMaxEntityCount(v int64) *ListProjectsResponseBodyProjects {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDatasetMaxRelationCount(v int64) *ListProjectsResponseBodyProjects {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDatasetMaxTotalFileSize(v int64) *ListProjectsResponseBodyProjects {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDatasetCount(v int64) *ListProjectsResponseBodyProjects {
	s.DatasetCount = &v
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

type RefreshWebofficeTokenRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// 链式授权
	AssumeRoleChain *AssumeRoleChain `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
}

func (s RefreshWebofficeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenRequest) SetProjectName(v string) *RefreshWebofficeTokenRequest {
	s.ProjectName = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetAccessToken(v string) *RefreshWebofficeTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetRefreshToken(v string) *RefreshWebofficeTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetAssumeRoleChain(v *AssumeRoleChain) *RefreshWebofficeTokenRequest {
	s.AssumeRoleChain = v
	return s
}

type RefreshWebofficeTokenShrinkRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// 链式授权
	AssumeRoleChainShrink *string `json:"AssumeRoleChain,omitempty" xml:"AssumeRoleChain,omitempty"`
}

func (s RefreshWebofficeTokenShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenShrinkRequest) SetProjectName(v string) *RefreshWebofficeTokenShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetAccessToken(v string) *RefreshWebofficeTokenShrinkRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetRefreshToken(v string) *RefreshWebofficeTokenShrinkRequest {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetAssumeRoleChainShrink(v string) *RefreshWebofficeTokenShrinkRequest {
	s.AssumeRoleChainShrink = &v
	return s
}

type RefreshWebofficeTokenResponseBody struct {
	// 请求 Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// refresh token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// access token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// refresh token 过期时间
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	// access token 过期时间
	AccessTokenExpiredTime *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
}

func (s RefreshWebofficeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponseBody) SetRequestId(v string) *RefreshWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshToken(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessToken(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
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
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ResumeBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeBindingRequest) GoString() string {
	return s.String()
}

func (s *ResumeBindingRequest) SetProjectName(v string) *ResumeBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *ResumeBindingRequest) SetDatasetName(v string) *ResumeBindingRequest {
	s.DatasetName = &v
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

type SimpleQueryRequest struct {
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Dataset 名称
	DatasetName *string      `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Query       *SimpleQuery `json:"Query,omitempty" xml:"Query,omitempty"`
	// 排序方式，默认 DESC
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// 排序字段
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 聚合字段
	Aggregations []*SimpleQueryRequestAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
}

func (s SimpleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryRequest) GoString() string {
	return s.String()
}

func (s *SimpleQueryRequest) SetNextToken(v string) *SimpleQueryRequest {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryRequest) SetMaxResults(v int32) *SimpleQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *SimpleQueryRequest) SetProjectName(v string) *SimpleQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *SimpleQueryRequest) SetDatasetName(v string) *SimpleQueryRequest {
	s.DatasetName = &v
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

func (s *SimpleQueryRequest) SetOrder(v string) *SimpleQueryRequest {
	s.Order = &v
	return s
}

func (s *SimpleQueryRequest) SetAggregations(v []*SimpleQueryRequestAggregations) *SimpleQueryRequest {
	s.Aggregations = v
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
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Dataset 名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// 排序方式，默认 DESC
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// 排序字段
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 聚合字段
	AggregationsShrink *string `json:"Aggregations,omitempty" xml:"Aggregations,omitempty"`
}

func (s SimpleQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *SimpleQueryShrinkRequest) SetNextToken(v string) *SimpleQueryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetMaxResults(v int32) *SimpleQueryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetProjectName(v string) *SimpleQueryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetDatasetName(v string) *SimpleQueryShrinkRequest {
	s.DatasetName = &v
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

func (s *SimpleQueryShrinkRequest) SetOrder(v string) *SimpleQueryShrinkRequest {
	s.Order = &v
	return s
}

func (s *SimpleQueryShrinkRequest) SetAggregationsShrink(v string) *SimpleQueryShrinkRequest {
	s.AggregationsShrink = &v
	return s
}

type SimpleQueryResponseBody struct {
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求的唯一 Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 文件列表
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// 聚合字段的字段名
	Aggregations []*SimpleQueryResponseBodyAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Repeated"`
}

func (s SimpleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBody) SetNextToken(v string) *SimpleQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SimpleQueryResponseBody) SetRequestId(v string) *SimpleQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SimpleQueryResponseBody) SetFiles(v []*File) *SimpleQueryResponseBody {
	s.Files = v
	return s
}

func (s *SimpleQueryResponseBody) SetAggregations(v []*SimpleQueryResponseBodyAggregations) *SimpleQueryResponseBody {
	s.Aggregations = v
	return s
}

type SimpleQueryResponseBodyAggregations struct {
	// 聚合字段名
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// 聚合字段的聚合操作符
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 聚合的统计结果
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
	// 分组聚合的结果
	Groups []*SimpleQueryResponseBodyAggregationsGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
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

func (s *SimpleQueryResponseBodyAggregations) SetOperation(v string) *SimpleQueryResponseBodyAggregations {
	s.Operation = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetValue(v float32) *SimpleQueryResponseBodyAggregations {
	s.Value = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregations) SetGroups(v []*SimpleQueryResponseBodyAggregationsGroups) *SimpleQueryResponseBodyAggregations {
	s.Groups = v
	return s
}

type SimpleQueryResponseBodyAggregationsGroups struct {
	// 分组聚合的值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 分组聚合的计数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s SimpleQueryResponseBodyAggregationsGroups) String() string {
	return tea.Prettify(s)
}

func (s SimpleQueryResponseBodyAggregationsGroups) GoString() string {
	return s.String()
}

func (s *SimpleQueryResponseBodyAggregationsGroups) SetValue(v string) *SimpleQueryResponseBodyAggregationsGroups {
	s.Value = &v
	return s
}

func (s *SimpleQueryResponseBodyAggregationsGroups) SetCount(v int64) *SimpleQueryResponseBodyAggregationsGroups {
	s.Count = &v
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
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s StopBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s StopBindingRequest) GoString() string {
	return s.String()
}

func (s *StopBindingRequest) SetProjectName(v string) *StopBindingRequest {
	s.ProjectName = &v
	return s
}

func (s *StopBindingRequest) SetDatasetName(v string) *StopBindingRequest {
	s.DatasetName = &v
	return s
}

func (s *StopBindingRequest) SetURI(v string) *StopBindingRequest {
	s.URI = &v
	return s
}

func (s *StopBindingRequest) SetReason(v string) *StopBindingRequest {
	s.Reason = &v
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
	ProjectName             *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName             *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateId              *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	DatasetMaxOSSBindCount  *int64    `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64    `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64    `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64    `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64    `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	ResetItems              []*string `json:"ResetItems,omitempty" xml:"ResetItems,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) SetProjectName(v string) *UpdateDatasetRequest {
	s.ProjectName = &v
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

func (s *UpdateDatasetRequest) SetTemplateId(v string) *UpdateDatasetRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxOSSBindCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxFileCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxEntityCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxEntityCount = &v
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

func (s *UpdateDatasetRequest) SetResetItems(v []*string) *UpdateDatasetRequest {
	s.ResetItems = v
	return s
}

type UpdateDatasetShrinkRequest struct {
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	DatasetMaxOSSBindCount  *int64  `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	ResetItemsShrink        *string `json:"ResetItems,omitempty" xml:"ResetItems,omitempty"`
}

func (s UpdateDatasetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetShrinkRequest) SetProjectName(v string) *UpdateDatasetShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetName(v string) *UpdateDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDescription(v string) *UpdateDatasetShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetTemplateId(v string) *UpdateDatasetShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxOSSBindCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxFileCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxEntityCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxRelationCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetResetItemsShrink(v string) *UpdateDatasetShrinkRequest {
	s.ResetItemsShrink = &v
	return s
}

type UpdateDatasetResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
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

type UpdateFileMetaRequest struct {
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s UpdateFileMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaRequest) SetProjectName(v string) *UpdateFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateFileMetaRequest) SetDatasetName(v string) *UpdateFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFileMetaRequest) SetURI(v string) *UpdateFileMetaRequest {
	s.URI = &v
	return s
}

func (s *UpdateFileMetaRequest) SetCustomLabels(v map[string]interface{}) *UpdateFileMetaRequest {
	s.CustomLabels = v
	return s
}

type UpdateFileMetaShrinkRequest struct {
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	URI                *string `json:"URI,omitempty" xml:"URI,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
}

func (s UpdateFileMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaShrinkRequest) SetProjectName(v string) *UpdateFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetDatasetName(v string) *UpdateFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetURI(v string) *UpdateFileMetaShrinkRequest {
	s.URI = &v
	return s
}

func (s *UpdateFileMetaShrinkRequest) SetCustomLabelsShrink(v string) *UpdateFileMetaShrinkRequest {
	s.CustomLabelsShrink = &v
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
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 服务角色
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// 项目描述
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateId              *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ProjectQPS              *int64    `json:"ProjectQPS,omitempty" xml:"ProjectQPS,omitempty"`
	ProjectTPS              *int64    `json:"ProjectTPS,omitempty" xml:"ProjectTPS,omitempty"`
	ProjectMaxDatasetCount  *int64    `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	DatasetMaxOSSBindCount  *int64    `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64    `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64    `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64    `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	ResetItems              []*string `json:"ResetItems,omitempty" xml:"ResetItems,omitempty" type:"Repeated"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectRequest) SetServiceRole(v string) *UpdateProjectRequest {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetTemplateId(v string) *UpdateProjectRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectQPS(v int64) *UpdateProjectRequest {
	s.ProjectQPS = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectTPS(v int64) *UpdateProjectRequest {
	s.ProjectTPS = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectMaxDatasetCount(v int64) *UpdateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxOSSBindCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxFileCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxEntityCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateProjectRequest) SetDatasetMaxRelationCount(v int64) *UpdateProjectRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateProjectRequest) SetResetItems(v []*string) *UpdateProjectRequest {
	s.ResetItems = v
	return s
}

type UpdateProjectShrinkRequest struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// 服务角色
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// 项目描述
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ProjectQPS              *int64  `json:"ProjectQPS,omitempty" xml:"ProjectQPS,omitempty"`
	ProjectTPS              *int64  `json:"ProjectTPS,omitempty" xml:"ProjectTPS,omitempty"`
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	DatasetMaxOSSBindCount  *int64  `json:"DatasetMaxOSSBindCount,omitempty" xml:"DatasetMaxOSSBindCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	ResetItemsShrink        *string `json:"ResetItems,omitempty" xml:"ResetItems,omitempty"`
}

func (s UpdateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectShrinkRequest) SetProjectName(v string) *UpdateProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetServiceRole(v string) *UpdateProjectShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDescription(v string) *UpdateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetTemplateId(v string) *UpdateProjectShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProjectQPS(v int64) *UpdateProjectShrinkRequest {
	s.ProjectQPS = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProjectTPS(v int64) *UpdateProjectShrinkRequest {
	s.ProjectTPS = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProjectMaxDatasetCount(v int64) *UpdateProjectShrinkRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxOSSBindCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxOSSBindCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxFileCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxEntityCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetDatasetMaxRelationCount(v int64) *UpdateProjectShrinkRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetResetItemsShrink(v string) *UpdateProjectShrinkRequest {
	s.ResetItemsShrink = &v
	return s
}

type UpdateProjectResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchDeleteFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchDeleteFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchGetFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchGetFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchIndexFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchIndexFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchUpdateFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchUpdateFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBinding"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataset"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBinding"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDataset"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) FuzzyQueryWithOptions(request *FuzzyQueryRequest, runtime *util.RuntimeOptions) (_result *FuzzyQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FuzzyQueryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FuzzyQuery"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBinding"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataset"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetFileMetaWithOptions(request *GetFileMetaRequest, runtime *util.RuntimeOptions) (_result *GetFileMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFileSignedURIResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFileSignedURI"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProject"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetWebofficeUrlWithOptions(tmpReq *GetWebofficeUrlRequest, runtime *util.RuntimeOptions) (_result *GetWebofficeUrlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetWebofficeUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Permission))) {
		request.PermissionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Permission), tea.String("Permission"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.User))) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.User), tea.String("User"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Watermark))) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Watermark), tea.String("Watermark"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.AssumeRoleChain))) {
		request.AssumeRoleChainShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.AssumeRoleChain), tea.String("AssumeRoleChain"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebofficeUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebofficeUrl"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebofficeUrl(request *GetWebofficeUrlRequest) (_result *GetWebofficeUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebofficeUrlResponse{}
	_body, _err := client.GetWebofficeUrlWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IndexFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IndexFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBindings"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDatasets"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjects"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshWebofficeTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshWebofficeToken"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResumeBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResumeBinding"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SimpleQueryWithOptions(tmpReq *SimpleQueryRequest, runtime *util.RuntimeOptions) (_result *SimpleQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SimpleQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Query))) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Query), tea.String("Query"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Aggregations)) {
		request.AggregationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Aggregations, tea.String("Aggregations"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SimpleQueryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SimpleQuery"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopBinding"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateDatasetWithOptions(tmpReq *UpdateDatasetRequest, runtime *util.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResetItems)) {
		request.ResetItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResetItems, tea.String("ResetItems"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDataset"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateFileMetaWithOptions(tmpReq *UpdateFileMetaRequest, runtime *util.RuntimeOptions) (_result *UpdateFileMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateFileMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFileMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFileMeta"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateProjectWithOptions(tmpReq *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResetItems)) {
		request.ResetItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResetItems, tea.String("ResetItems"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
