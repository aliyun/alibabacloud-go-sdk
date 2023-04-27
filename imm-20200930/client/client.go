// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
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

type AddressForStory struct {
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Township *string `json:"Township,omitempty" xml:"Township,omitempty"`
}

func (s AddressForStory) String() string {
	return tea.Prettify(s)
}

func (s AddressForStory) GoString() string {
	return s.String()
}

func (s *AddressForStory) SetCity(v string) *AddressForStory {
	s.City = &v
	return s
}

func (s *AddressForStory) SetCountry(v string) *AddressForStory {
	s.Country = &v
	return s
}

func (s *AddressForStory) SetDistrict(v string) *AddressForStory {
	s.District = &v
	return s
}

func (s *AddressForStory) SetProvince(v string) *AddressForStory {
	s.Province = &v
	return s
}

func (s *AddressForStory) SetTownship(v string) *AddressForStory {
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

type Car struct {
	Boundary           *Boundary       `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	CarColor           *string         `json:"CarColor,omitempty" xml:"CarColor,omitempty"`
	CarColorConfidence *float64        `json:"CarColorConfidence,omitempty" xml:"CarColorConfidence,omitempty"`
	CarType            *string         `json:"CarType,omitempty" xml:"CarType,omitempty"`
	CarTypeConfidence  *float64        `json:"CarTypeConfidence,omitempty" xml:"CarTypeConfidence,omitempty"`
	Confidence         *float64        `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	LicensePlates      []*LicensePlate `json:"LicensePlates,omitempty" xml:"LicensePlates,omitempty" type:"Repeated"`
}

func (s Car) String() string {
	return tea.Prettify(s)
}

func (s Car) GoString() string {
	return s.String()
}

func (s *Car) SetBoundary(v *Boundary) *Car {
	s.Boundary = v
	return s
}

func (s *Car) SetCarColor(v string) *Car {
	s.CarColor = &v
	return s
}

func (s *Car) SetCarColorConfidence(v float64) *Car {
	s.CarColorConfidence = &v
	return s
}

func (s *Car) SetCarType(v string) *Car {
	s.CarType = &v
	return s
}

func (s *Car) SetCarTypeConfidence(v float64) *Car {
	s.CarTypeConfidence = &v
	return s
}

func (s *Car) SetConfidence(v float64) *Car {
	s.Confidence = &v
	return s
}

func (s *Car) SetLicensePlates(v []*LicensePlate) *Car {
	s.LicensePlates = v
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

type DataIngestion struct {
	Actions      []*DataIngestionActions    `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	CreateTime   *string                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error        *string                    `json:"Error,omitempty" xml:"Error,omitempty"`
	Id           *string                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Input        *Input                     `json:"Input,omitempty" xml:"Input,omitempty"`
	Marker       *string                    `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Notification *DataIngestionNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	State        *string                    `json:"State,omitempty" xml:"State,omitempty"`
	Statistic    *DataIngestionStatistic    `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	Tags         map[string]interface{}     `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UpdateTime   *string                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DataIngestion) String() string {
	return tea.Prettify(s)
}

func (s DataIngestion) GoString() string {
	return s.String()
}

func (s *DataIngestion) SetActions(v []*DataIngestionActions) *DataIngestion {
	s.Actions = v
	return s
}

func (s *DataIngestion) SetCreateTime(v string) *DataIngestion {
	s.CreateTime = &v
	return s
}

func (s *DataIngestion) SetError(v string) *DataIngestion {
	s.Error = &v
	return s
}

func (s *DataIngestion) SetId(v string) *DataIngestion {
	s.Id = &v
	return s
}

func (s *DataIngestion) SetInput(v *Input) *DataIngestion {
	s.Input = v
	return s
}

func (s *DataIngestion) SetMarker(v string) *DataIngestion {
	s.Marker = &v
	return s
}

func (s *DataIngestion) SetNotification(v *DataIngestionNotification) *DataIngestion {
	s.Notification = v
	return s
}

func (s *DataIngestion) SetState(v string) *DataIngestion {
	s.State = &v
	return s
}

func (s *DataIngestion) SetStatistic(v *DataIngestionStatistic) *DataIngestion {
	s.Statistic = v
	return s
}

func (s *DataIngestion) SetTags(v map[string]interface{}) *DataIngestion {
	s.Tags = v
	return s
}

func (s *DataIngestion) SetUpdateTime(v string) *DataIngestion {
	s.UpdateTime = &v
	return s
}

type DataIngestionActions struct {
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DataIngestionActions) String() string {
	return tea.Prettify(s)
}

func (s DataIngestionActions) GoString() string {
	return s.String()
}

func (s *DataIngestionActions) SetName(v string) *DataIngestionActions {
	s.Name = &v
	return s
}

func (s *DataIngestionActions) SetParameters(v []*string) *DataIngestionActions {
	s.Parameters = v
	return s
}

type DataIngestionNotification struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Topic    *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DataIngestionNotification) String() string {
	return tea.Prettify(s)
}

func (s DataIngestionNotification) GoString() string {
	return s.String()
}

func (s *DataIngestionNotification) SetEndpoint(v string) *DataIngestionNotification {
	s.Endpoint = &v
	return s
}

func (s *DataIngestionNotification) SetTopic(v string) *DataIngestionNotification {
	s.Topic = &v
	return s
}

type DataIngestionStatistic struct {
	SubmitFailure *int64 `json:"SubmitFailure,omitempty" xml:"SubmitFailure,omitempty"`
	SubmitSuccess *int64 `json:"SubmitSuccess,omitempty" xml:"SubmitSuccess,omitempty"`
}

func (s DataIngestionStatistic) String() string {
	return tea.Prettify(s)
}

func (s DataIngestionStatistic) GoString() string {
	return s.String()
}

func (s *DataIngestionStatistic) SetSubmitFailure(v int64) *DataIngestionStatistic {
	s.SubmitFailure = &v
	return s
}

func (s *DataIngestionStatistic) SetSubmitSuccess(v int64) *DataIngestionStatistic {
	s.SubmitSuccess = &v
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
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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

func (s *Dataset) SetTemplateId(v string) *Dataset {
	s.TemplateId = &v
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
	AverageAge      *float32               `json:"AverageAge,omitempty" xml:"AverageAge,omitempty"`
	Cover           *File                  `json:"Cover,omitempty" xml:"Cover,omitempty"`
	CreateTime      *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomId        *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels    map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName     *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FaceCount       *int64                 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	Gender          *string                `json:"Gender,omitempty" xml:"Gender,omitempty"`
	ImageCount      *int64                 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	MaxAge          *float32               `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	MetaLockVersion *int64                 `json:"MetaLockVersion,omitempty" xml:"MetaLockVersion,omitempty"`
	MinAge          *float32               `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectId        *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType      *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId         *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName     *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	UpdateTime      *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VideoCount      *int64                 `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
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

func (s *FigureCluster) SetMetaLockVersion(v int64) *FigureCluster {
	s.MetaLockVersion = &v
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
	Cover           *FigureClusterForReqCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	CustomId        *string                   `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels    map[string]interface{}    `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	MetaLockVersion *int64                    `json:"MetaLockVersion,omitempty" xml:"MetaLockVersion,omitempty"`
	Name            *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectId        *string                   `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
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

func (s *FigureClusterForReq) SetMetaLockVersion(v int64) *FigureClusterForReq {
	s.MetaLockVersion = &v
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

type Input struct {
	OSS *InputOSS `json:"OSS,omitempty" xml:"OSS,omitempty"`
}

func (s Input) String() string {
	return tea.Prettify(s)
}

func (s Input) GoString() string {
	return s.String()
}

func (s *Input) SetOSS(v *InputOSS) *Input {
	s.OSS = v
	return s
}

type InputFile struct {
	ContentType  *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Figures      []*InputFileFigures    `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	FileHash     *string                `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	MediaType    *string                `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OSSURI       *string                `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s InputFile) String() string {
	return tea.Prettify(s)
}

func (s InputFile) GoString() string {
	return s.String()
}

func (s *InputFile) SetContentType(v string) *InputFile {
	s.ContentType = &v
	return s
}

func (s *InputFile) SetCustomId(v string) *InputFile {
	s.CustomId = &v
	return s
}

func (s *InputFile) SetCustomLabels(v map[string]interface{}) *InputFile {
	s.CustomLabels = v
	return s
}

func (s *InputFile) SetFigures(v []*InputFileFigures) *InputFile {
	s.Figures = v
	return s
}

func (s *InputFile) SetFileHash(v string) *InputFile {
	s.FileHash = &v
	return s
}

func (s *InputFile) SetMediaType(v string) *InputFile {
	s.MediaType = &v
	return s
}

func (s *InputFile) SetOSSURI(v string) *InputFile {
	s.OSSURI = &v
	return s
}

func (s *InputFile) SetURI(v string) *InputFile {
	s.URI = &v
	return s
}

type InputFileFigures struct {
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	FigureId        *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	FigureType      *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
}

func (s InputFileFigures) String() string {
	return tea.Prettify(s)
}

func (s InputFileFigures) GoString() string {
	return s.String()
}

func (s *InputFileFigures) SetFigureClusterId(v string) *InputFileFigures {
	s.FigureClusterId = &v
	return s
}

func (s *InputFileFigures) SetFigureId(v string) *InputFileFigures {
	s.FigureId = &v
	return s
}

func (s *InputFileFigures) SetFigureType(v string) *InputFileFigures {
	s.FigureType = &v
	return s
}

type InputOSS struct {
	Bucket           *string   `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	MatchExpressions []*string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty" type:"Repeated"`
	Prefix           *string   `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s InputOSS) String() string {
	return tea.Prettify(s)
}

func (s InputOSS) GoString() string {
	return s.String()
}

func (s *InputOSS) SetBucket(v string) *InputOSS {
	s.Bucket = &v
	return s
}

func (s *InputOSS) SetMatchExpressions(v []*string) *InputOSS {
	s.MatchExpressions = v
	return s
}

func (s *InputOSS) SetPrefix(v string) *InputOSS {
	s.Prefix = &v
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

type LicensePlate struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	Confidence *float64  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Content    *string   `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s LicensePlate) String() string {
	return tea.Prettify(s)
}

func (s LicensePlate) GoString() string {
	return s.String()
}

func (s *LicensePlate) SetBoundary(v *Boundary) *LicensePlate {
	s.Boundary = v
	return s
}

func (s *LicensePlate) SetConfidence(v float64) *LicensePlate {
	s.Confidence = &v
	return s
}

func (s *LicensePlate) SetContent(v string) *LicensePlate {
	s.Content = &v
	return s
}

type LocationDateCluster struct {
	Addresses                    []*Address             `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	CreateTime                   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomId                     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels                 map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	LocationDateClusterEndTime   *string                `json:"LocationDateClusterEndTime,omitempty" xml:"LocationDateClusterEndTime,omitempty"`
	LocationDateClusterLevel     *string                `json:"LocationDateClusterLevel,omitempty" xml:"LocationDateClusterLevel,omitempty"`
	LocationDateClusterStartTime *string                `json:"LocationDateClusterStartTime,omitempty" xml:"LocationDateClusterStartTime,omitempty"`
	ObjectId                     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Title                        *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime                   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s LocationDateCluster) String() string {
	return tea.Prettify(s)
}

func (s LocationDateCluster) GoString() string {
	return s.String()
}

func (s *LocationDateCluster) SetAddresses(v []*Address) *LocationDateCluster {
	s.Addresses = v
	return s
}

func (s *LocationDateCluster) SetCreateTime(v string) *LocationDateCluster {
	s.CreateTime = &v
	return s
}

func (s *LocationDateCluster) SetCustomId(v string) *LocationDateCluster {
	s.CustomId = &v
	return s
}

func (s *LocationDateCluster) SetCustomLabels(v map[string]interface{}) *LocationDateCluster {
	s.CustomLabels = v
	return s
}

func (s *LocationDateCluster) SetLocationDateClusterEndTime(v string) *LocationDateCluster {
	s.LocationDateClusterEndTime = &v
	return s
}

func (s *LocationDateCluster) SetLocationDateClusterLevel(v string) *LocationDateCluster {
	s.LocationDateClusterLevel = &v
	return s
}

func (s *LocationDateCluster) SetLocationDateClusterStartTime(v string) *LocationDateCluster {
	s.LocationDateClusterStartTime = &v
	return s
}

func (s *LocationDateCluster) SetObjectId(v string) *LocationDateCluster {
	s.ObjectId = &v
	return s
}

func (s *LocationDateCluster) SetTitle(v string) *LocationDateCluster {
	s.Title = &v
	return s
}

func (s *LocationDateCluster) SetUpdateTime(v string) *LocationDateCluster {
	s.UpdateTime = &v
	return s
}

type MNS struct {
	Endpoint  *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s MNS) String() string {
	return tea.Prettify(s)
}

func (s MNS) GoString() string {
	return s.String()
}

func (s *MNS) SetEndpoint(v string) *MNS {
	s.Endpoint = &v
	return s
}

func (s *MNS) SetTopicName(v string) *MNS {
	s.TopicName = &v
	return s
}

type Notification struct {
	MNS      *MNS      `json:"MNS,omitempty" xml:"MNS,omitempty"`
	RocketMQ *RocketMQ `json:"RocketMQ,omitempty" xml:"RocketMQ,omitempty"`
}

func (s Notification) String() string {
	return tea.Prettify(s)
}

func (s Notification) GoString() string {
	return s.String()
}

func (s *Notification) SetMNS(v *MNS) *Notification {
	s.MNS = v
	return s
}

func (s *Notification) SetRocketMQ(v *RocketMQ) *Notification {
	s.RocketMQ = v
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
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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

func (s *Project) SetTemplateId(v string) *Project {
	s.TemplateId = &v
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

type RocketMQ struct {
	Endpoint   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TopicName  *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s RocketMQ) String() string {
	return tea.Prettify(s)
}

func (s RocketMQ) GoString() string {
	return s.String()
}

func (s *RocketMQ) SetEndpoint(v string) *RocketMQ {
	s.Endpoint = &v
	return s
}

func (s *RocketMQ) SetInstanceId(v string) *RocketMQ {
	s.InstanceId = &v
	return s
}

func (s *RocketMQ) SetTopicName(v string) *RocketMQ {
	s.TopicName = &v
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

type SimilarImage struct {
	ImageScore *float64 `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	URI        *string  `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s SimilarImage) String() string {
	return tea.Prettify(s)
}

func (s SimilarImage) GoString() string {
	return s.String()
}

func (s *SimilarImage) SetImageScore(v float64) *SimilarImage {
	s.ImageScore = &v
	return s
}

func (s *SimilarImage) SetURI(v string) *SimilarImage {
	s.URI = &v
	return s
}

type SimilarImageCluster struct {
	CreateTime   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Files        []*SimilarImage        `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ObjectId     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	UpdateTime   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s SimilarImageCluster) String() string {
	return tea.Prettify(s)
}

func (s SimilarImageCluster) GoString() string {
	return s.String()
}

func (s *SimilarImageCluster) SetCreateTime(v string) *SimilarImageCluster {
	s.CreateTime = &v
	return s
}

func (s *SimilarImageCluster) SetCustomLabels(v map[string]interface{}) *SimilarImageCluster {
	s.CustomLabels = v
	return s
}

func (s *SimilarImageCluster) SetFiles(v []*SimilarImage) *SimilarImageCluster {
	s.Files = v
	return s
}

func (s *SimilarImageCluster) SetObjectId(v string) *SimilarImageCluster {
	s.ObjectId = &v
	return s
}

func (s *SimilarImageCluster) SetUpdateTime(v string) *SimilarImageCluster {
	s.UpdateTime = &v
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
	Addresses        []*Address             `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
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

func (s *Story) SetAddresses(v []*Address) *Story {
	s.Addresses = v
	return s
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

type TargetAudio struct {
	DisableAudio   *bool                      `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	FilterAudio    *TargetAudioFilterAudio    `json:"FilterAudio,omitempty" xml:"FilterAudio,omitempty" type:"Struct"`
	Stream         []*int64                   `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
	TranscodeAudio *TargetAudioTranscodeAudio `json:"TranscodeAudio,omitempty" xml:"TranscodeAudio,omitempty" type:"Struct"`
}

func (s TargetAudio) String() string {
	return tea.Prettify(s)
}

func (s TargetAudio) GoString() string {
	return s.String()
}

func (s *TargetAudio) SetDisableAudio(v bool) *TargetAudio {
	s.DisableAudio = &v
	return s
}

func (s *TargetAudio) SetFilterAudio(v *TargetAudioFilterAudio) *TargetAudio {
	s.FilterAudio = v
	return s
}

func (s *TargetAudio) SetStream(v []*int64) *TargetAudio {
	s.Stream = v
	return s
}

func (s *TargetAudio) SetTranscodeAudio(v *TargetAudioTranscodeAudio) *TargetAudio {
	s.TranscodeAudio = v
	return s
}

type TargetAudioFilterAudio struct {
	Mixing *bool `json:"Mixing,omitempty" xml:"Mixing,omitempty"`
}

func (s TargetAudioFilterAudio) String() string {
	return tea.Prettify(s)
}

func (s TargetAudioFilterAudio) GoString() string {
	return s.String()
}

func (s *TargetAudioFilterAudio) SetMixing(v bool) *TargetAudioFilterAudio {
	s.Mixing = &v
	return s
}

type TargetAudioTranscodeAudio struct {
	Bitrate          *int32  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateOption    *string `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	Channel          *int32  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Codec            *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	Quality          *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	SampleRate       *int32  `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SampleRateOption *string `json:"SampleRateOption,omitempty" xml:"SampleRateOption,omitempty"`
}

func (s TargetAudioTranscodeAudio) String() string {
	return tea.Prettify(s)
}

func (s TargetAudioTranscodeAudio) GoString() string {
	return s.String()
}

func (s *TargetAudioTranscodeAudio) SetBitrate(v int32) *TargetAudioTranscodeAudio {
	s.Bitrate = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetBitrateOption(v string) *TargetAudioTranscodeAudio {
	s.BitrateOption = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetChannel(v int32) *TargetAudioTranscodeAudio {
	s.Channel = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetCodec(v string) *TargetAudioTranscodeAudio {
	s.Codec = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetQuality(v int32) *TargetAudioTranscodeAudio {
	s.Quality = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetSampleRate(v int32) *TargetAudioTranscodeAudio {
	s.SampleRate = &v
	return s
}

func (s *TargetAudioTranscodeAudio) SetSampleRateOption(v string) *TargetAudioTranscodeAudio {
	s.SampleRateOption = &v
	return s
}

type TargetImage struct {
	Animations []*TargetImageAnimations `json:"Animations,omitempty" xml:"Animations,omitempty" type:"Repeated"`
	Snapshots  []*TargetImageSnapshots  `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	Sprites    []*TargetImageSprites    `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Repeated"`
}

func (s TargetImage) String() string {
	return tea.Prettify(s)
}

func (s TargetImage) GoString() string {
	return s.String()
}

func (s *TargetImage) SetAnimations(v []*TargetImageAnimations) *TargetImage {
	s.Animations = v
	return s
}

func (s *TargetImage) SetSnapshots(v []*TargetImageSnapshots) *TargetImage {
	s.Snapshots = v
	return s
}

func (s *TargetImage) SetSprites(v []*TargetImageSprites) *TargetImage {
	s.Sprites = v
	return s
}

type TargetImageAnimations struct {
	Format    *string  `json:"Format,omitempty" xml:"Format,omitempty"`
	FrameRate *float64 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Height    *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	Interval  *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Number    *int32   `json:"Number,omitempty" xml:"Number,omitempty"`
	ScaleType *string  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	URI       *string  `json:"URI,omitempty" xml:"URI,omitempty"`
	Width     *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetImageAnimations) String() string {
	return tea.Prettify(s)
}

func (s TargetImageAnimations) GoString() string {
	return s.String()
}

func (s *TargetImageAnimations) SetFormat(v string) *TargetImageAnimations {
	s.Format = &v
	return s
}

func (s *TargetImageAnimations) SetFrameRate(v float64) *TargetImageAnimations {
	s.FrameRate = &v
	return s
}

func (s *TargetImageAnimations) SetHeight(v int32) *TargetImageAnimations {
	s.Height = &v
	return s
}

func (s *TargetImageAnimations) SetInterval(v float64) *TargetImageAnimations {
	s.Interval = &v
	return s
}

func (s *TargetImageAnimations) SetNumber(v int32) *TargetImageAnimations {
	s.Number = &v
	return s
}

func (s *TargetImageAnimations) SetScaleType(v string) *TargetImageAnimations {
	s.ScaleType = &v
	return s
}

func (s *TargetImageAnimations) SetStartTime(v float64) *TargetImageAnimations {
	s.StartTime = &v
	return s
}

func (s *TargetImageAnimations) SetURI(v string) *TargetImageAnimations {
	s.URI = &v
	return s
}

func (s *TargetImageAnimations) SetWidth(v int32) *TargetImageAnimations {
	s.Width = &v
	return s
}

type TargetImageSnapshots struct {
	Format    *string  `json:"Format,omitempty" xml:"Format,omitempty"`
	Height    *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	Interval  *float64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Number    *int32   `json:"Number,omitempty" xml:"Number,omitempty"`
	ScaleType *string  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	URI       *string  `json:"URI,omitempty" xml:"URI,omitempty"`
	Width     *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetImageSnapshots) String() string {
	return tea.Prettify(s)
}

func (s TargetImageSnapshots) GoString() string {
	return s.String()
}

func (s *TargetImageSnapshots) SetFormat(v string) *TargetImageSnapshots {
	s.Format = &v
	return s
}

func (s *TargetImageSnapshots) SetHeight(v int32) *TargetImageSnapshots {
	s.Height = &v
	return s
}

func (s *TargetImageSnapshots) SetInterval(v float64) *TargetImageSnapshots {
	s.Interval = &v
	return s
}

func (s *TargetImageSnapshots) SetNumber(v int32) *TargetImageSnapshots {
	s.Number = &v
	return s
}

func (s *TargetImageSnapshots) SetScaleType(v string) *TargetImageSnapshots {
	s.ScaleType = &v
	return s
}

func (s *TargetImageSnapshots) SetStartTime(v float64) *TargetImageSnapshots {
	s.StartTime = &v
	return s
}

func (s *TargetImageSnapshots) SetURI(v string) *TargetImageSnapshots {
	s.URI = &v
	return s
}

func (s *TargetImageSnapshots) SetWidth(v int32) *TargetImageSnapshots {
	s.Width = &v
	return s
}

type TargetImageSprites struct {
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

func (s TargetImageSprites) String() string {
	return tea.Prettify(s)
}

func (s TargetImageSprites) GoString() string {
	return s.String()
}

func (s *TargetImageSprites) SetFormat(v string) *TargetImageSprites {
	s.Format = &v
	return s
}

func (s *TargetImageSprites) SetInterval(v float64) *TargetImageSprites {
	s.Interval = &v
	return s
}

func (s *TargetImageSprites) SetMargin(v int32) *TargetImageSprites {
	s.Margin = &v
	return s
}

func (s *TargetImageSprites) SetNumber(v int32) *TargetImageSprites {
	s.Number = &v
	return s
}

func (s *TargetImageSprites) SetPad(v int32) *TargetImageSprites {
	s.Pad = &v
	return s
}

func (s *TargetImageSprites) SetScaleHeight(v float32) *TargetImageSprites {
	s.ScaleHeight = &v
	return s
}

func (s *TargetImageSprites) SetScaleType(v string) *TargetImageSprites {
	s.ScaleType = &v
	return s
}

func (s *TargetImageSprites) SetScaleWidth(v float32) *TargetImageSprites {
	s.ScaleWidth = &v
	return s
}

func (s *TargetImageSprites) SetStartTime(v float64) *TargetImageSprites {
	s.StartTime = &v
	return s
}

func (s *TargetImageSprites) SetTileHeight(v int32) *TargetImageSprites {
	s.TileHeight = &v
	return s
}

func (s *TargetImageSprites) SetTileWidth(v int32) *TargetImageSprites {
	s.TileWidth = &v
	return s
}

func (s *TargetImageSprites) SetURI(v string) *TargetImageSprites {
	s.URI = &v
	return s
}

type TargetSubtitle struct {
	DisableSubtitle *bool                          `json:"DisableSubtitle,omitempty" xml:"DisableSubtitle,omitempty"`
	ExtractSubtitle *TargetSubtitleExtractSubtitle `json:"ExtractSubtitle,omitempty" xml:"ExtractSubtitle,omitempty" type:"Struct"`
	Stream          []*int32                       `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s TargetSubtitle) String() string {
	return tea.Prettify(s)
}

func (s TargetSubtitle) GoString() string {
	return s.String()
}

func (s *TargetSubtitle) SetDisableSubtitle(v bool) *TargetSubtitle {
	s.DisableSubtitle = &v
	return s
}

func (s *TargetSubtitle) SetExtractSubtitle(v *TargetSubtitleExtractSubtitle) *TargetSubtitle {
	s.ExtractSubtitle = v
	return s
}

func (s *TargetSubtitle) SetStream(v []*int32) *TargetSubtitle {
	s.Stream = v
	return s
}

type TargetSubtitleExtractSubtitle struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	URI    *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s TargetSubtitleExtractSubtitle) String() string {
	return tea.Prettify(s)
}

func (s TargetSubtitleExtractSubtitle) GoString() string {
	return s.String()
}

func (s *TargetSubtitleExtractSubtitle) SetFormat(v string) *TargetSubtitleExtractSubtitle {
	s.Format = &v
	return s
}

func (s *TargetSubtitleExtractSubtitle) SetURI(v string) *TargetSubtitleExtractSubtitle {
	s.URI = &v
	return s
}

type TargetVideo struct {
	DisableVideo   *bool                      `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	FilterVideo    *TargetVideoFilterVideo    `json:"FilterVideo,omitempty" xml:"FilterVideo,omitempty" type:"Struct"`
	Stream         []*int32                   `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
	TranscodeVideo *TargetVideoTranscodeVideo `json:"TranscodeVideo,omitempty" xml:"TranscodeVideo,omitempty" type:"Struct"`
}

func (s TargetVideo) String() string {
	return tea.Prettify(s)
}

func (s TargetVideo) GoString() string {
	return s.String()
}

func (s *TargetVideo) SetDisableVideo(v bool) *TargetVideo {
	s.DisableVideo = &v
	return s
}

func (s *TargetVideo) SetFilterVideo(v *TargetVideoFilterVideo) *TargetVideo {
	s.FilterVideo = v
	return s
}

func (s *TargetVideo) SetStream(v []*int32) *TargetVideo {
	s.Stream = v
	return s
}

func (s *TargetVideo) SetTranscodeVideo(v *TargetVideoTranscodeVideo) *TargetVideo {
	s.TranscodeVideo = v
	return s
}

type TargetVideoFilterVideo struct {
	Delogos    []*TargetVideoFilterVideoDelogos    `json:"Delogos,omitempty" xml:"Delogos,omitempty" type:"Repeated"`
	Watermarks []*TargetVideoFilterVideoWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s TargetVideoFilterVideo) String() string {
	return tea.Prettify(s)
}

func (s TargetVideoFilterVideo) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideo) SetDelogos(v []*TargetVideoFilterVideoDelogos) *TargetVideoFilterVideo {
	s.Delogos = v
	return s
}

func (s *TargetVideoFilterVideo) SetWatermarks(v []*TargetVideoFilterVideoWatermarks) *TargetVideoFilterVideo {
	s.Watermarks = v
	return s
}

type TargetVideoFilterVideoDelogos struct {
	Duration  *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Dx        *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy        *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	ReferPos  *string  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetVideoFilterVideoDelogos) String() string {
	return tea.Prettify(s)
}

func (s TargetVideoFilterVideoDelogos) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDelogos) SetDuration(v float64) *TargetVideoFilterVideoDelogos {
	s.Duration = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetDx(v float32) *TargetVideoFilterVideoDelogos {
	s.Dx = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetDy(v float32) *TargetVideoFilterVideoDelogos {
	s.Dy = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetHeight(v float32) *TargetVideoFilterVideoDelogos {
	s.Height = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetReferPos(v string) *TargetVideoFilterVideoDelogos {
	s.ReferPos = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetStartTime(v float64) *TargetVideoFilterVideoDelogos {
	s.StartTime = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetWidth(v float32) *TargetVideoFilterVideoDelogos {
	s.Width = &v
	return s
}

type TargetVideoFilterVideoWatermarks struct {
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

func (s TargetVideoFilterVideoWatermarks) String() string {
	return tea.Prettify(s)
}

func (s TargetVideoFilterVideoWatermarks) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoWatermarks) SetBorderColor(v string) *TargetVideoFilterVideoWatermarks {
	s.BorderColor = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetBorderWidth(v int32) *TargetVideoFilterVideoWatermarks {
	s.BorderWidth = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetContent(v string) *TargetVideoFilterVideoWatermarks {
	s.Content = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetDuration(v float64) *TargetVideoFilterVideoWatermarks {
	s.Duration = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetDx(v float32) *TargetVideoFilterVideoWatermarks {
	s.Dx = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetDy(v float32) *TargetVideoFilterVideoWatermarks {
	s.Dy = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontApha(v float32) *TargetVideoFilterVideoWatermarks {
	s.FontApha = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontColor(v string) *TargetVideoFilterVideoWatermarks {
	s.FontColor = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontName(v string) *TargetVideoFilterVideoWatermarks {
	s.FontName = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontSize(v int32) *TargetVideoFilterVideoWatermarks {
	s.FontSize = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetHeight(v float32) *TargetVideoFilterVideoWatermarks {
	s.Height = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetReferPos(v string) *TargetVideoFilterVideoWatermarks {
	s.ReferPos = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetStartTime(v float64) *TargetVideoFilterVideoWatermarks {
	s.StartTime = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetType(v string) *TargetVideoFilterVideoWatermarks {
	s.Type = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetURI(v string) *TargetVideoFilterVideoWatermarks {
	s.URI = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetWidth(v float32) *TargetVideoFilterVideoWatermarks {
	s.Width = &v
	return s
}

type TargetVideoTranscodeVideo struct {
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

func (s TargetVideoTranscodeVideo) String() string {
	return tea.Prettify(s)
}

func (s TargetVideoTranscodeVideo) GoString() string {
	return s.String()
}

func (s *TargetVideoTranscodeVideo) SetAdaptiveResolutionDirection(v bool) *TargetVideoTranscodeVideo {
	s.AdaptiveResolutionDirection = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBFrames(v int32) *TargetVideoTranscodeVideo {
	s.BFrames = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBitrate(v int32) *TargetVideoTranscodeVideo {
	s.Bitrate = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBitrateOption(v string) *TargetVideoTranscodeVideo {
	s.BitrateOption = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBufferSize(v int32) *TargetVideoTranscodeVideo {
	s.BufferSize = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetCRF(v float32) *TargetVideoTranscodeVideo {
	s.CRF = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetCodec(v string) *TargetVideoTranscodeVideo {
	s.Codec = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetFrameRate(v float32) *TargetVideoTranscodeVideo {
	s.FrameRate = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetFrameRateOption(v string) *TargetVideoTranscodeVideo {
	s.FrameRateOption = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetGOPSize(v int32) *TargetVideoTranscodeVideo {
	s.GOPSize = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetMaxBitrate(v int32) *TargetVideoTranscodeVideo {
	s.MaxBitrate = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetPixelFormat(v string) *TargetVideoTranscodeVideo {
	s.PixelFormat = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetRefs(v int32) *TargetVideoTranscodeVideo {
	s.Refs = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetResolution(v string) *TargetVideoTranscodeVideo {
	s.Resolution = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetResolutionOption(v string) *TargetVideoTranscodeVideo {
	s.ResolutionOption = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetRotation(v int32) *TargetVideoTranscodeVideo {
	s.Rotation = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetScaleType(v string) *TargetVideoTranscodeVideo {
	s.ScaleType = &v
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

type AddImageMosaicRequest struct {
	CredentialConfig *CredentialConfig               `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ImageFormat      *string                         `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ProjectName      *string                         `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality          *int32                          `json:"Quality,omitempty" xml:"Quality,omitempty"`
	SourceURI        *string                         `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TargetURI        *string                         `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	Targets          []*AddImageMosaicRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s AddImageMosaicRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageMosaicRequest) GoString() string {
	return s.String()
}

func (s *AddImageMosaicRequest) SetCredentialConfig(v *CredentialConfig) *AddImageMosaicRequest {
	s.CredentialConfig = v
	return s
}

func (s *AddImageMosaicRequest) SetImageFormat(v string) *AddImageMosaicRequest {
	s.ImageFormat = &v
	return s
}

func (s *AddImageMosaicRequest) SetProjectName(v string) *AddImageMosaicRequest {
	s.ProjectName = &v
	return s
}

func (s *AddImageMosaicRequest) SetQuality(v int32) *AddImageMosaicRequest {
	s.Quality = &v
	return s
}

func (s *AddImageMosaicRequest) SetSourceURI(v string) *AddImageMosaicRequest {
	s.SourceURI = &v
	return s
}

func (s *AddImageMosaicRequest) SetTargetURI(v string) *AddImageMosaicRequest {
	s.TargetURI = &v
	return s
}

func (s *AddImageMosaicRequest) SetTargets(v []*AddImageMosaicRequestTargets) *AddImageMosaicRequest {
	s.Targets = v
	return s
}

type AddImageMosaicRequestTargets struct {
	BlurRadius   *int32                                `json:"BlurRadius,omitempty" xml:"BlurRadius,omitempty"`
	Boundary     *AddImageMosaicRequestTargetsBoundary `json:"Boundary,omitempty" xml:"Boundary,omitempty" type:"Struct"`
	Color        *string                               `json:"Color,omitempty" xml:"Color,omitempty"`
	MosaicRadius *int32                                `json:"MosaicRadius,omitempty" xml:"MosaicRadius,omitempty"`
	Sigma        *int32                                `json:"Sigma,omitempty" xml:"Sigma,omitempty"`
	Type         *string                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddImageMosaicRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s AddImageMosaicRequestTargets) GoString() string {
	return s.String()
}

func (s *AddImageMosaicRequestTargets) SetBlurRadius(v int32) *AddImageMosaicRequestTargets {
	s.BlurRadius = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetBoundary(v *AddImageMosaicRequestTargetsBoundary) *AddImageMosaicRequestTargets {
	s.Boundary = v
	return s
}

func (s *AddImageMosaicRequestTargets) SetColor(v string) *AddImageMosaicRequestTargets {
	s.Color = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetMosaicRadius(v int32) *AddImageMosaicRequestTargets {
	s.MosaicRadius = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetSigma(v int32) *AddImageMosaicRequestTargets {
	s.Sigma = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetType(v string) *AddImageMosaicRequestTargets {
	s.Type = &v
	return s
}

type AddImageMosaicRequestTargetsBoundary struct {
	Height   *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	ReferPos *string  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	Width    *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X        *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y        *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s AddImageMosaicRequestTargetsBoundary) String() string {
	return tea.Prettify(s)
}

func (s AddImageMosaicRequestTargetsBoundary) GoString() string {
	return s.String()
}

func (s *AddImageMosaicRequestTargetsBoundary) SetHeight(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.Height = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetReferPos(v string) *AddImageMosaicRequestTargetsBoundary {
	s.ReferPos = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetWidth(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.Width = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetX(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.X = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetY(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.Y = &v
	return s
}

type AddImageMosaicShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ImageFormat            *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality                *int32  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TargetURI              *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	TargetsShrink          *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s AddImageMosaicShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageMosaicShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddImageMosaicShrinkRequest) SetCredentialConfigShrink(v string) *AddImageMosaicShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetImageFormat(v string) *AddImageMosaicShrinkRequest {
	s.ImageFormat = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetProjectName(v string) *AddImageMosaicShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetQuality(v int32) *AddImageMosaicShrinkRequest {
	s.Quality = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetSourceURI(v string) *AddImageMosaicShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetTargetURI(v string) *AddImageMosaicShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetTargetsShrink(v string) *AddImageMosaicShrinkRequest {
	s.TargetsShrink = &v
	return s
}

type AddImageMosaicResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddImageMosaicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageMosaicResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageMosaicResponseBody) SetRequestId(v string) *AddImageMosaicResponseBody {
	s.RequestId = &v
	return s
}

type AddImageMosaicResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddImageMosaicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddImageMosaicResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageMosaicResponse) GoString() string {
	return s.String()
}

func (s *AddImageMosaicResponse) SetHeaders(v map[string]*string) *AddImageMosaicResponse {
	s.Headers = v
	return s
}

func (s *AddImageMosaicResponse) SetStatusCode(v int32) *AddImageMosaicResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageMosaicResponse) SetBody(v *AddImageMosaicResponseBody) *AddImageMosaicResponse {
	s.Body = v
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

type BatchGetFigureClusterRequest struct {
	DatasetName *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectIds   []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	ProjectName *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchGetFigureClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterRequest) SetDatasetName(v string) *BatchGetFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFigureClusterRequest) SetObjectIds(v []*string) *BatchGetFigureClusterRequest {
	s.ObjectIds = v
	return s
}

func (s *BatchGetFigureClusterRequest) SetProjectName(v string) *BatchGetFigureClusterRequest {
	s.ProjectName = &v
	return s
}

type BatchGetFigureClusterShrinkRequest struct {
	DatasetName     *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	ProjectName     *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s BatchGetFigureClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFigureClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterShrinkRequest) SetDatasetName(v string) *BatchGetFigureClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchGetFigureClusterShrinkRequest) SetObjectIdsShrink(v string) *BatchGetFigureClusterShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *BatchGetFigureClusterShrinkRequest) SetProjectName(v string) *BatchGetFigureClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

type BatchGetFigureClusterResponseBody struct {
	FigureClusters []*FigureCluster `json:"FigureClusters,omitempty" xml:"FigureClusters,omitempty" type:"Repeated"`
	RequestId      *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetFigureClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterResponseBody) SetFigureClusters(v []*FigureCluster) *BatchGetFigureClusterResponseBody {
	s.FigureClusters = v
	return s
}

func (s *BatchGetFigureClusterResponseBody) SetRequestId(v string) *BatchGetFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

type BatchGetFigureClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGetFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetFigureClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterResponse) SetHeaders(v map[string]*string) *BatchGetFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *BatchGetFigureClusterResponse) SetStatusCode(v int32) *BatchGetFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetFigureClusterResponse) SetBody(v *BatchGetFigureClusterResponseBody) *BatchGetFigureClusterResponse {
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
	DatasetName  *string       `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files        []*InputFile  `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *BatchIndexFileMetaRequest) SetFiles(v []*InputFile) *BatchIndexFileMetaRequest {
	s.Files = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetNotification(v *Notification) *BatchIndexFileMetaRequest {
	s.Notification = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetProjectName(v string) *BatchIndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

type BatchIndexFileMetaShrinkRequest struct {
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink        *string `json:"Files,omitempty" xml:"Files,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *BatchIndexFileMetaShrinkRequest) SetNotificationShrink(v string) *BatchIndexFileMetaShrinkRequest {
	s.NotificationShrink = &v
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
	DatasetName *string      `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files       []*InputFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ProjectName *string      `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *BatchUpdateFileMetaRequest) SetFiles(v []*InputFile) *BatchUpdateFileMetaRequest {
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

type CompareImageFacesRequest struct {
	CredentialConfig *CredentialConfig               `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string                         `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Source           *CompareImageFacesRequestSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
}

func (s CompareImageFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareImageFacesRequest) SetCredentialConfig(v *CredentialConfig) *CompareImageFacesRequest {
	s.CredentialConfig = v
	return s
}

func (s *CompareImageFacesRequest) SetProjectName(v string) *CompareImageFacesRequest {
	s.ProjectName = &v
	return s
}

func (s *CompareImageFacesRequest) SetSource(v *CompareImageFacesRequestSource) *CompareImageFacesRequest {
	s.Source = v
	return s
}

type CompareImageFacesRequestSource struct {
	URI1 *string `json:"URI1,omitempty" xml:"URI1,omitempty"`
	URI2 *string `json:"URI2,omitempty" xml:"URI2,omitempty"`
}

func (s CompareImageFacesRequestSource) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesRequestSource) GoString() string {
	return s.String()
}

func (s *CompareImageFacesRequestSource) SetURI1(v string) *CompareImageFacesRequestSource {
	s.URI1 = &v
	return s
}

func (s *CompareImageFacesRequestSource) SetURI2(v string) *CompareImageFacesRequestSource {
	s.URI2 = &v
	return s
}

type CompareImageFacesShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceShrink           *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CompareImageFacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CompareImageFacesShrinkRequest) SetCredentialConfigShrink(v string) *CompareImageFacesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CompareImageFacesShrinkRequest) SetProjectName(v string) *CompareImageFacesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CompareImageFacesShrinkRequest) SetSourceShrink(v string) *CompareImageFacesShrinkRequest {
	s.SourceShrink = &v
	return s
}

type CompareImageFacesResponseBody struct {
	RequestId  *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Similarity *float32 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s CompareImageFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponseBody) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponseBody) SetRequestId(v string) *CompareImageFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareImageFacesResponseBody) SetSimilarity(v float32) *CompareImageFacesResponseBody {
	s.Similarity = &v
	return s
}

type CompareImageFacesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompareImageFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareImageFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareImageFacesResponse) GoString() string {
	return s.String()
}

func (s *CompareImageFacesResponse) SetHeaders(v map[string]*string) *CompareImageFacesResponse {
	s.Headers = v
	return s
}

func (s *CompareImageFacesResponse) SetStatusCode(v int32) *CompareImageFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareImageFacesResponse) SetBody(v *CompareImageFacesResponseBody) *CompareImageFacesResponse {
	s.Body = v
	return s
}

type CreateArchiveFileInspectionTaskRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Notification     *Notification     `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Password         *string           `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TargetURI        *string           `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData         *string           `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateArchiveFileInspectionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateArchiveFileInspectionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetNotification(v *Notification) *CreateArchiveFileInspectionTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetPassword(v string) *CreateArchiveFileInspectionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetProjectName(v string) *CreateArchiveFileInspectionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetSourceURI(v string) *CreateArchiveFileInspectionTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetTargetURI(v string) *CreateArchiveFileInspectionTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskRequest) SetUserData(v string) *CreateArchiveFileInspectionTaskRequest {
	s.UserData = &v
	return s
}

type CreateArchiveFileInspectionTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotificationShrink     *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TargetURI              *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateArchiveFileInspectionTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetNotificationShrink(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetPassword(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetProjectName(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetSourceURI(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetTargetURI(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskShrinkRequest) SetUserData(v string) *CreateArchiveFileInspectionTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateArchiveFileInspectionTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateArchiveFileInspectionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskResponseBody) SetEventId(v string) *CreateArchiveFileInspectionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponseBody) SetRequestId(v string) *CreateArchiveFileInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponseBody) SetTaskId(v string) *CreateArchiveFileInspectionTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateArchiveFileInspectionTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateArchiveFileInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateArchiveFileInspectionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskResponse) SetHeaders(v map[string]*string) *CreateArchiveFileInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponse) SetStatusCode(v int32) *CreateArchiveFileInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponse) SetBody(v *CreateArchiveFileInspectionTaskResponseBody) *CreateArchiveFileInspectionTaskResponse {
	s.Body = v
	return s
}

type CreateBatchRequest struct {
	Actions      []*CreateBatchRequestActions    `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Input        *Input                          `json:"Input,omitempty" xml:"Input,omitempty"`
	Notification *CreateBatchRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	ProjectName  *string                         `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ServiceRole  *string                         `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Tags         map[string]interface{}          `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchRequest) SetActions(v []*CreateBatchRequestActions) *CreateBatchRequest {
	s.Actions = v
	return s
}

func (s *CreateBatchRequest) SetInput(v *Input) *CreateBatchRequest {
	s.Input = v
	return s
}

func (s *CreateBatchRequest) SetNotification(v *CreateBatchRequestNotification) *CreateBatchRequest {
	s.Notification = v
	return s
}

func (s *CreateBatchRequest) SetProjectName(v string) *CreateBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBatchRequest) SetServiceRole(v string) *CreateBatchRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateBatchRequest) SetTags(v map[string]interface{}) *CreateBatchRequest {
	s.Tags = v
	return s
}

type CreateBatchRequestActions struct {
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s CreateBatchRequestActions) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchRequestActions) GoString() string {
	return s.String()
}

func (s *CreateBatchRequestActions) SetName(v string) *CreateBatchRequestActions {
	s.Name = &v
	return s
}

func (s *CreateBatchRequestActions) SetParameters(v []*string) *CreateBatchRequestActions {
	s.Parameters = v
	return s
}

type CreateBatchRequestNotification struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Topic    *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateBatchRequestNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchRequestNotification) GoString() string {
	return s.String()
}

func (s *CreateBatchRequestNotification) SetEndpoint(v string) *CreateBatchRequestNotification {
	s.Endpoint = &v
	return s
}

func (s *CreateBatchRequestNotification) SetTopic(v string) *CreateBatchRequestNotification {
	s.Topic = &v
	return s
}

type CreateBatchShrinkRequest struct {
	ActionsShrink      *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	InputShrink        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ServiceRole        *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateBatchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchShrinkRequest) SetActionsShrink(v string) *CreateBatchShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetInputShrink(v string) *CreateBatchShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetNotificationShrink(v string) *CreateBatchShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetProjectName(v string) *CreateBatchShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetServiceRole(v string) *CreateBatchShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateBatchShrinkRequest) SetTagsShrink(v string) *CreateBatchShrinkRequest {
	s.TagsShrink = &v
	return s
}

type CreateBatchResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchResponseBody) SetId(v string) *CreateBatchResponseBody {
	s.Id = &v
	return s
}

func (s *CreateBatchResponseBody) SetRequestId(v string) *CreateBatchResponseBody {
	s.RequestId = &v
	return s
}

type CreateBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchResponse) SetHeaders(v map[string]*string) *CreateBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchResponse) SetStatusCode(v int32) *CreateBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchResponse) SetBody(v *CreateBatchResponseBody) *CreateBatchResponse {
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
	Notification         *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s *CreateCompressPointCloudTaskRequest) SetNotification(v *Notification) *CreateCompressPointCloudTaskRequest {
	s.Notification = v
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
	NotificationShrink     *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s *CreateCompressPointCloudTaskShrinkRequest) SetNotificationShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.NotificationShrink = &v
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

type CreateCustomizedStoryRequest struct {
	Cover        *CreateCustomizedStoryRequestCover   `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	CustomLabels map[string]interface{}               `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string                              `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Files        []*CreateCustomizedStoryRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ProjectName  *string                              `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryName    *string                              `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StorySubType *string                              `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType    *string                              `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
}

func (s CreateCustomizedStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequest) SetCover(v *CreateCustomizedStoryRequestCover) *CreateCustomizedStoryRequest {
	s.Cover = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetCustomLabels(v map[string]interface{}) *CreateCustomizedStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetDatasetName(v string) *CreateCustomizedStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetFiles(v []*CreateCustomizedStoryRequestFiles) *CreateCustomizedStoryRequest {
	s.Files = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetProjectName(v string) *CreateCustomizedStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryName(v string) *CreateCustomizedStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStorySubType(v string) *CreateCustomizedStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryType(v string) *CreateCustomizedStoryRequest {
	s.StoryType = &v
	return s
}

type CreateCustomizedStoryRequestCover struct {
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateCustomizedStoryRequestCover) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryRequestCover) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestCover) SetURI(v string) *CreateCustomizedStoryRequestCover {
	s.URI = &v
	return s
}

type CreateCustomizedStoryRequestFiles struct {
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateCustomizedStoryRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryRequestFiles) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestFiles) SetURI(v string) *CreateCustomizedStoryRequestFiles {
	s.URI = &v
	return s
}

type CreateCustomizedStoryShrinkRequest struct {
	CoverShrink        *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FilesShrink        *string `json:"Files,omitempty" xml:"Files,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StoryName          *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	StorySubType       *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	StoryType          *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
}

func (s CreateCustomizedStoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryShrinkRequest) SetCoverShrink(v string) *CreateCustomizedStoryShrinkRequest {
	s.CoverShrink = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetCustomLabelsShrink(v string) *CreateCustomizedStoryShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetDatasetName(v string) *CreateCustomizedStoryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetFilesShrink(v string) *CreateCustomizedStoryShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetProjectName(v string) *CreateCustomizedStoryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetStoryName(v string) *CreateCustomizedStoryShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetStorySubType(v string) *CreateCustomizedStoryShrinkRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateCustomizedStoryShrinkRequest) SetStoryType(v string) *CreateCustomizedStoryShrinkRequest {
	s.StoryType = &v
	return s
}

type CreateCustomizedStoryResponseBody struct {
	ObjectId  *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomizedStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponseBody) SetObjectId(v string) *CreateCustomizedStoryResponseBody {
	s.ObjectId = &v
	return s
}

func (s *CreateCustomizedStoryResponseBody) SetRequestId(v string) *CreateCustomizedStoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomizedStoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomizedStoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomizedStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponse) SetHeaders(v map[string]*string) *CreateCustomizedStoryResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedStoryResponse) SetStatusCode(v int32) *CreateCustomizedStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedStoryResponse) SetBody(v *CreateCustomizedStoryResponseBody) *CreateCustomizedStoryResponse {
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

type CreateFacesSearchingTaskRequest struct {
	DatasetName  *string                                   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResult    *string                                   `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	Notification *Notification                             `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                                   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sources      []*CreateFacesSearchingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	TopK         *int64                                    `json:"TopK,omitempty" xml:"TopK,omitempty"`
	UserData     *string                                   `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFacesSearchingTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFacesSearchingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskRequest) SetDatasetName(v string) *CreateFacesSearchingTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetMaxResult(v string) *CreateFacesSearchingTaskRequest {
	s.MaxResult = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetNotification(v *Notification) *CreateFacesSearchingTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetProjectName(v string) *CreateFacesSearchingTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetSources(v []*CreateFacesSearchingTaskRequestSources) *CreateFacesSearchingTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetTopK(v int64) *CreateFacesSearchingTaskRequest {
	s.TopK = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetUserData(v string) *CreateFacesSearchingTaskRequest {
	s.UserData = &v
	return s
}

type CreateFacesSearchingTaskRequestSources struct {
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateFacesSearchingTaskRequestSources) String() string {
	return tea.Prettify(s)
}

func (s CreateFacesSearchingTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskRequestSources) SetURI(v string) *CreateFacesSearchingTaskRequestSources {
	s.URI = &v
	return s
}

type CreateFacesSearchingTaskShrinkRequest struct {
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResult          *string `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourcesShrink      *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopK               *int64  `json:"TopK,omitempty" xml:"TopK,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFacesSearchingTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFacesSearchingTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetDatasetName(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetMaxResult(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.MaxResult = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetNotificationShrink(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetProjectName(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetSourcesShrink(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetTopK(v int64) *CreateFacesSearchingTaskShrinkRequest {
	s.TopK = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetUserData(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateFacesSearchingTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFacesSearchingTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFacesSearchingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskResponseBody) SetEventId(v string) *CreateFacesSearchingTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFacesSearchingTaskResponseBody) SetRequestId(v string) *CreateFacesSearchingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFacesSearchingTaskResponseBody) SetTaskId(v string) *CreateFacesSearchingTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateFacesSearchingTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFacesSearchingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFacesSearchingTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFacesSearchingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskResponse) SetHeaders(v map[string]*string) *CreateFacesSearchingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFacesSearchingTaskResponse) SetStatusCode(v int32) *CreateFacesSearchingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFacesSearchingTaskResponse) SetBody(v *CreateFacesSearchingTaskResponseBody) *CreateFacesSearchingTaskResponse {
	s.Body = v
	return s
}

type CreateFigureClusteringTaskRequest struct {
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Notification *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData     *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateFigureClusteringTaskRequest) SetNotification(v *Notification) *CreateFigureClusteringTaskRequest {
	s.Notification = v
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
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateFigureClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
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
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	From         *string                `json:"From,omitempty" xml:"From,omitempty"`
	Froms        []*string              `json:"Froms,omitempty" xml:"Froms,omitempty" type:"Repeated"`
	Notification *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	To           *string                `json:"To,omitempty" xml:"To,omitempty"`
	UserData     *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateFigureClustersMergingTaskRequest) SetFroms(v []*string) *CreateFigureClustersMergingTaskRequest {
	s.Froms = v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetNotification(v *Notification) *CreateFigureClustersMergingTaskRequest {
	s.Notification = v
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
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	From               *string `json:"From,omitempty" xml:"From,omitempty"`
	FromsShrink        *string `json:"Froms,omitempty" xml:"Froms,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	To                 *string `json:"To,omitempty" xml:"To,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetFromsShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.FromsShrink = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetNotificationShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.NotificationShrink = &v
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

type CreateFileCompressionTaskRequest struct {
	CompressedFormat  *string                                    `json:"CompressedFormat,omitempty" xml:"CompressedFormat,omitempty"`
	CredentialConfig  *CredentialConfig                          `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Notification      *Notification                              `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Password          *string                                    `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName       *string                                    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceManifestURI *string                                    `json:"SourceManifestURI,omitempty" xml:"SourceManifestURI,omitempty"`
	Sources           []*CreateFileCompressionTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	TargetURI         *string                                    `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData          *string                                    `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFileCompressionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileCompressionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskRequest) SetCompressedFormat(v string) *CreateFileCompressionTaskRequest {
	s.CompressedFormat = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateFileCompressionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetNotification(v *Notification) *CreateFileCompressionTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetPassword(v string) *CreateFileCompressionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetProjectName(v string) *CreateFileCompressionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetSourceManifestURI(v string) *CreateFileCompressionTaskRequest {
	s.SourceManifestURI = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetSources(v []*CreateFileCompressionTaskRequestSources) *CreateFileCompressionTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetTargetURI(v string) *CreateFileCompressionTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateFileCompressionTaskRequest) SetUserData(v string) *CreateFileCompressionTaskRequest {
	s.UserData = &v
	return s
}

type CreateFileCompressionTaskRequestSources struct {
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	URI   *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateFileCompressionTaskRequestSources) String() string {
	return tea.Prettify(s)
}

func (s CreateFileCompressionTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskRequestSources) SetAlias(v string) *CreateFileCompressionTaskRequestSources {
	s.Alias = &v
	return s
}

func (s *CreateFileCompressionTaskRequestSources) SetURI(v string) *CreateFileCompressionTaskRequestSources {
	s.URI = &v
	return s
}

type CreateFileCompressionTaskShrinkRequest struct {
	CompressedFormat       *string `json:"CompressedFormat,omitempty" xml:"CompressedFormat,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotificationShrink     *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceManifestURI      *string `json:"SourceManifestURI,omitempty" xml:"SourceManifestURI,omitempty"`
	SourcesShrink          *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TargetURI              *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFileCompressionTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileCompressionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskShrinkRequest) SetCompressedFormat(v string) *CreateFileCompressionTaskShrinkRequest {
	s.CompressedFormat = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateFileCompressionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetNotificationShrink(v string) *CreateFileCompressionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetPassword(v string) *CreateFileCompressionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetProjectName(v string) *CreateFileCompressionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetSourceManifestURI(v string) *CreateFileCompressionTaskShrinkRequest {
	s.SourceManifestURI = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetSourcesShrink(v string) *CreateFileCompressionTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetTargetURI(v string) *CreateFileCompressionTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateFileCompressionTaskShrinkRequest) SetUserData(v string) *CreateFileCompressionTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateFileCompressionTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFileCompressionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileCompressionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskResponseBody) SetEventId(v string) *CreateFileCompressionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFileCompressionTaskResponseBody) SetRequestId(v string) *CreateFileCompressionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileCompressionTaskResponseBody) SetTaskId(v string) *CreateFileCompressionTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateFileCompressionTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileCompressionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileCompressionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileCompressionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskResponse) SetHeaders(v map[string]*string) *CreateFileCompressionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFileCompressionTaskResponse) SetStatusCode(v int32) *CreateFileCompressionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileCompressionTaskResponse) SetBody(v *CreateFileCompressionTaskResponseBody) *CreateFileCompressionTaskResponse {
	s.Body = v
	return s
}

type CreateFileUncompressionTaskRequest struct {
	CredentialConfig *CredentialConfig                         `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Notification     *Notification                             `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Password         *string                                   `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName      *string                                   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SelectedFiles    []*string                                 `json:"SelectedFiles,omitempty" xml:"SelectedFiles,omitempty" type:"Repeated"`
	SourceURI        *string                                   `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Target           *CreateFileUncompressionTaskRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	UserData         *string                                   `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFileUncompressionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileUncompressionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateFileUncompressionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetNotification(v *Notification) *CreateFileUncompressionTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetPassword(v string) *CreateFileUncompressionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetProjectName(v string) *CreateFileUncompressionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetSelectedFiles(v []*string) *CreateFileUncompressionTaskRequest {
	s.SelectedFiles = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetSourceURI(v string) *CreateFileUncompressionTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetTarget(v *CreateFileUncompressionTaskRequestTarget) *CreateFileUncompressionTaskRequest {
	s.Target = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetUserData(v string) *CreateFileUncompressionTaskRequest {
	s.UserData = &v
	return s
}

type CreateFileUncompressionTaskRequestTarget struct {
	ManifestURI *string `json:"ManifestURI,omitempty" xml:"ManifestURI,omitempty"`
	URI         *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateFileUncompressionTaskRequestTarget) String() string {
	return tea.Prettify(s)
}

func (s CreateFileUncompressionTaskRequestTarget) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskRequestTarget) SetManifestURI(v string) *CreateFileUncompressionTaskRequestTarget {
	s.ManifestURI = &v
	return s
}

func (s *CreateFileUncompressionTaskRequestTarget) SetURI(v string) *CreateFileUncompressionTaskRequestTarget {
	s.URI = &v
	return s
}

type CreateFileUncompressionTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	NotificationShrink     *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SelectedFilesShrink    *string `json:"SelectedFiles,omitempty" xml:"SelectedFiles,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TargetShrink           *string `json:"Target,omitempty" xml:"Target,omitempty"`
	UserData               *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFileUncompressionTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileUncompressionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetNotificationShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetPassword(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetProjectName(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetSelectedFilesShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.SelectedFilesShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetSourceURI(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetTargetShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.TargetShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetUserData(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateFileUncompressionTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFileUncompressionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileUncompressionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskResponseBody) SetEventId(v string) *CreateFileUncompressionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFileUncompressionTaskResponseBody) SetRequestId(v string) *CreateFileUncompressionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileUncompressionTaskResponseBody) SetTaskId(v string) *CreateFileUncompressionTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateFileUncompressionTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileUncompressionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileUncompressionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileUncompressionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskResponse) SetHeaders(v map[string]*string) *CreateFileUncompressionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFileUncompressionTaskResponse) SetStatusCode(v int32) *CreateFileUncompressionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileUncompressionTaskResponse) SetBody(v *CreateFileUncompressionTaskResponseBody) *CreateFileUncompressionTaskResponse {
	s.Body = v
	return s
}

type CreateImageModerationTaskRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Interval         *int64            `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames        *int64            `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer     *string                `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	Scenes       []*string              `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	SourceURI    *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData     *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateImageModerationTaskRequest) SetNotification(v *Notification) *CreateImageModerationTaskRequest {
	s.Notification = v
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
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer           *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	ScenesShrink       *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	SourceURI          *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateImageModerationTaskShrinkRequest) SetNotificationShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.NotificationShrink = &v
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
	Align            *int64            `json:"Align,omitempty" xml:"Align,omitempty"`
	BackgroundColor  *string           `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Direction        *string           `json:"Direction,omitempty" xml:"Direction,omitempty"`
	ImageFormat      *string           `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	Margin           *int64            `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification *Notification                            `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Padding      *int64                                   `json:"Padding,omitempty" xml:"Padding,omitempty"`
	ProjectName  *string                                  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality      *int64                                   `json:"Quality,omitempty" xml:"Quality,omitempty"`
	ScaleType    *string                                  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	Sources      []*CreateImageSplicingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	Tags         map[string]interface{}                   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI    *string                                  `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData     *string                                  `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateImageSplicingTaskRequest) SetNotification(v *Notification) *CreateImageSplicingTaskRequest {
	s.Notification = v
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
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	Padding            *int64  `json:"Padding,omitempty" xml:"Padding,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Quality            *int64  `json:"Quality,omitempty" xml:"Quality,omitempty"`
	ScaleType          *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SourcesShrink      *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI          *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateImageSplicingTaskShrinkRequest) SetNotificationShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.NotificationShrink = &v
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

type CreateImageToPDFTaskRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification *Notification                         `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                               `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sources      []*CreateImageToPDFTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	Tags         map[string]interface{}                `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI    *string                               `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData     *string                               `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageToPDFTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageToPDFTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateImageToPDFTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetNotification(v *Notification) *CreateImageToPDFTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetProjectName(v string) *CreateImageToPDFTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetSources(v []*CreateImageToPDFTaskRequestSources) *CreateImageToPDFTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetTags(v map[string]interface{}) *CreateImageToPDFTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetTargetURI(v string) *CreateImageToPDFTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageToPDFTaskRequest) SetUserData(v string) *CreateImageToPDFTaskRequest {
	s.UserData = &v
	return s
}

type CreateImageToPDFTaskRequestSources struct {
	Rotate *int64  `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	URI    *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateImageToPDFTaskRequestSources) String() string {
	return tea.Prettify(s)
}

func (s CreateImageToPDFTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskRequestSources) SetRotate(v int64) *CreateImageToPDFTaskRequestSources {
	s.Rotate = &v
	return s
}

func (s *CreateImageToPDFTaskRequestSources) SetURI(v string) *CreateImageToPDFTaskRequestSources {
	s.URI = &v
	return s
}

type CreateImageToPDFTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourcesShrink      *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetURI          *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageToPDFTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageToPDFTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetNotificationShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetProjectName(v string) *CreateImageToPDFTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetSourcesShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetTagsShrink(v string) *CreateImageToPDFTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetTargetURI(v string) *CreateImageToPDFTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageToPDFTaskShrinkRequest) SetUserData(v string) *CreateImageToPDFTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateImageToPDFTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageToPDFTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageToPDFTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskResponseBody) SetEventId(v string) *CreateImageToPDFTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateImageToPDFTaskResponseBody) SetRequestId(v string) *CreateImageToPDFTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageToPDFTaskResponseBody) SetTaskId(v string) *CreateImageToPDFTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateImageToPDFTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageToPDFTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageToPDFTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageToPDFTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskResponse) SetHeaders(v map[string]*string) *CreateImageToPDFTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageToPDFTaskResponse) SetStatusCode(v int32) *CreateImageToPDFTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageToPDFTaskResponse) SetBody(v *CreateImageToPDFTaskResponseBody) *CreateImageToPDFTaskResponse {
	s.Body = v
	return s
}

type CreateLocationDateClusteringTaskRequest struct {
	DatasetName     *string                                                 `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DateOptions     *CreateLocationDateClusteringTaskRequestDateOptions     `json:"DateOptions,omitempty" xml:"DateOptions,omitempty" type:"Struct"`
	LocationOptions *CreateLocationDateClusteringTaskRequestLocationOptions `json:"LocationOptions,omitempty" xml:"LocationOptions,omitempty" type:"Struct"`
	Notification    *Notification                                           `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName     *string                                                 `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags            map[string]interface{}                                  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData        *string                                                 `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateLocationDateClusteringTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLocationDateClusteringTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskRequest) SetDatasetName(v string) *CreateLocationDateClusteringTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetDateOptions(v *CreateLocationDateClusteringTaskRequestDateOptions) *CreateLocationDateClusteringTaskRequest {
	s.DateOptions = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetLocationOptions(v *CreateLocationDateClusteringTaskRequestLocationOptions) *CreateLocationDateClusteringTaskRequest {
	s.LocationOptions = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetNotification(v *Notification) *CreateLocationDateClusteringTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetProjectName(v string) *CreateLocationDateClusteringTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetTags(v map[string]interface{}) *CreateLocationDateClusteringTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetUserData(v string) *CreateLocationDateClusteringTaskRequest {
	s.UserData = &v
	return s
}

type CreateLocationDateClusteringTaskRequestDateOptions struct {
	GapDays *int64 `json:"GapDays,omitempty" xml:"GapDays,omitempty"`
	MaxDays *int64 `json:"MaxDays,omitempty" xml:"MaxDays,omitempty"`
	MinDays *int64 `json:"MinDays,omitempty" xml:"MinDays,omitempty"`
}

func (s CreateLocationDateClusteringTaskRequestDateOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateLocationDateClusteringTaskRequestDateOptions) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) SetGapDays(v int64) *CreateLocationDateClusteringTaskRequestDateOptions {
	s.GapDays = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) SetMaxDays(v int64) *CreateLocationDateClusteringTaskRequestDateOptions {
	s.MaxDays = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) SetMinDays(v int64) *CreateLocationDateClusteringTaskRequestDateOptions {
	s.MinDays = &v
	return s
}

type CreateLocationDateClusteringTaskRequestLocationOptions struct {
	LocationDateClusterLevels []*string `json:"LocationDateClusterLevels,omitempty" xml:"LocationDateClusterLevels,omitempty" type:"Repeated"`
}

func (s CreateLocationDateClusteringTaskRequestLocationOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateLocationDateClusteringTaskRequestLocationOptions) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskRequestLocationOptions) SetLocationDateClusterLevels(v []*string) *CreateLocationDateClusteringTaskRequestLocationOptions {
	s.LocationDateClusterLevels = v
	return s
}

type CreateLocationDateClusteringTaskShrinkRequest struct {
	DatasetName           *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DateOptionsShrink     *string `json:"DateOptions,omitempty" xml:"DateOptions,omitempty"`
	LocationOptionsShrink *string `json:"LocationOptions,omitempty" xml:"LocationOptions,omitempty"`
	NotificationShrink    *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName           *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData              *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateLocationDateClusteringTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLocationDateClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetDateOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.DateOptionsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetLocationOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.LocationOptionsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetProjectName(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetUserData(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateLocationDateClusteringTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateLocationDateClusteringTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLocationDateClusteringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskResponseBody) SetEventId(v string) *CreateLocationDateClusteringTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponseBody) SetRequestId(v string) *CreateLocationDateClusteringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponseBody) SetTaskId(v string) *CreateLocationDateClusteringTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateLocationDateClusteringTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLocationDateClusteringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLocationDateClusteringTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLocationDateClusteringTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskResponse) SetHeaders(v map[string]*string) *CreateLocationDateClusteringTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLocationDateClusteringTaskResponse) SetStatusCode(v int32) *CreateLocationDateClusteringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponse) SetBody(v *CreateLocationDateClusteringTaskResponseBody) *CreateLocationDateClusteringTaskResponse {
	s.Body = v
	return s
}

type CreateMediaConvertTaskRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification *Notification                           `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                                 `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sources      []*CreateMediaConvertTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	Tags         map[string]interface{}                  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Targets      []*CreateMediaConvertTaskRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	UserData     *string                                 `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateMediaConvertTaskRequest) SetNotification(v *Notification) *CreateMediaConvertTaskRequest {
	s.Notification = v
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
	Audio         *TargetAudio                                 `json:"Audio,omitempty" xml:"Audio,omitempty"`
	Container     *string                                      `json:"Container,omitempty" xml:"Container,omitempty"`
	Image         *TargetImage                                 `json:"Image,omitempty" xml:"Image,omitempty"`
	Segment       *CreateMediaConvertTaskRequestTargetsSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	Speed         *float32                                     `json:"Speed,omitempty" xml:"Speed,omitempty"`
	StripMetadata *bool                                        `json:"StripMetadata,omitempty" xml:"StripMetadata,omitempty"`
	Subtitle      *TargetSubtitle                              `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	URI           *string                                      `json:"URI,omitempty" xml:"URI,omitempty"`
	Video         *TargetVideo                                 `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargets) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargets) SetAudio(v *TargetAudio) *CreateMediaConvertTaskRequestTargets {
	s.Audio = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetContainer(v string) *CreateMediaConvertTaskRequestTargets {
	s.Container = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetImage(v *TargetImage) *CreateMediaConvertTaskRequestTargets {
	s.Image = v
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

func (s *CreateMediaConvertTaskRequestTargets) SetStripMetadata(v bool) *CreateMediaConvertTaskRequestTargets {
	s.StripMetadata = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSubtitle(v *TargetSubtitle) *CreateMediaConvertTaskRequestTargets {
	s.Subtitle = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetURI(v string) *CreateMediaConvertTaskRequestTargets {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetVideo(v *TargetVideo) *CreateMediaConvertTaskRequestTargets {
	s.Video = v
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

type CreateMediaConvertTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourcesShrink      *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetsShrink      *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateMediaConvertTaskShrinkRequest) SetNotificationShrink(v string) *CreateMediaConvertTaskShrinkRequest {
	s.NotificationShrink = &v
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
	Notification     *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s *CreateOfficeConversionTaskRequest) SetNotification(v *Notification) *CreateOfficeConversionTaskRequest {
	s.Notification = v
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
	NotificationShrink     *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s *CreateOfficeConversionTaskShrinkRequest) SetNotificationShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.NotificationShrink = &v
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
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *CreateProjectRequest) SetProjectMaxDatasetCount(v int64) *CreateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
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

type CreateSimilarImageClusteringTaskRequest struct {
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Notification *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData     *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateSimilarImageClusteringTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskRequest) SetDatasetName(v string) *CreateSimilarImageClusteringTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetNotification(v *Notification) *CreateSimilarImageClusteringTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetProjectName(v string) *CreateSimilarImageClusteringTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetTags(v map[string]interface{}) *CreateSimilarImageClusteringTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetUserData(v string) *CreateSimilarImageClusteringTaskRequest {
	s.UserData = &v
	return s
}

type CreateSimilarImageClusteringTaskShrinkRequest struct {
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateSimilarImageClusteringTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetProjectName(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetUserData(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateSimilarImageClusteringTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSimilarImageClusteringTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskResponseBody) SetEventId(v string) *CreateSimilarImageClusteringTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponseBody) SetRequestId(v string) *CreateSimilarImageClusteringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponseBody) SetTaskId(v string) *CreateSimilarImageClusteringTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateSimilarImageClusteringTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSimilarImageClusteringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSimilarImageClusteringTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskResponse) SetHeaders(v map[string]*string) *CreateSimilarImageClusteringTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponse) SetStatusCode(v int32) *CreateSimilarImageClusteringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponse) SetBody(v *CreateSimilarImageClusteringTaskResponseBody) *CreateSimilarImageClusteringTaskResponse {
	s.Body = v
	return s
}

type CreateStoryRequest struct {
	Address      *AddressForStory       `json:"Address,omitempty" xml:"Address,omitempty"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxFileCount *int64                 `json:"MaxFileCount,omitempty" xml:"MaxFileCount,omitempty"`
	MinFileCount *int64                 `json:"MinFileCount,omitempty" xml:"MinFileCount,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification    *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s *CreateStoryRequest) SetAddress(v *AddressForStory) *CreateStoryRequest {
	s.Address = v
	return s
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

func (s *CreateStoryRequest) SetNotification(v *Notification) *CreateStoryRequest {
	s.Notification = v
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
	AddressShrink      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	CustomId           *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxFileCount       *int64  `json:"MaxFileCount,omitempty" xml:"MaxFileCount,omitempty"`
	MinFileCount       *int64  `json:"MinFileCount,omitempty" xml:"MinFileCount,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s *CreateStoryShrinkRequest) SetAddressShrink(v string) *CreateStoryShrinkRequest {
	s.AddressShrink = &v
	return s
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

func (s *CreateStoryShrinkRequest) SetNotificationShrink(v string) *CreateStoryShrinkRequest {
	s.NotificationShrink = &v
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

type CreateTriggerRequest struct {
	Actions      []*CreateTriggerRequestActions    `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Input        *Input                            `json:"Input,omitempty" xml:"Input,omitempty"`
	Notification *CreateTriggerRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	ProjectName  *string                           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ServiceRole  *string                           `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	Tags         map[string]interface{}            `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetActions(v []*CreateTriggerRequestActions) *CreateTriggerRequest {
	s.Actions = v
	return s
}

func (s *CreateTriggerRequest) SetInput(v *Input) *CreateTriggerRequest {
	s.Input = v
	return s
}

func (s *CreateTriggerRequest) SetNotification(v *CreateTriggerRequestNotification) *CreateTriggerRequest {
	s.Notification = v
	return s
}

func (s *CreateTriggerRequest) SetProjectName(v string) *CreateTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateTriggerRequest) SetServiceRole(v string) *CreateTriggerRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateTriggerRequest) SetTags(v map[string]interface{}) *CreateTriggerRequest {
	s.Tags = v
	return s
}

type CreateTriggerRequestActions struct {
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s CreateTriggerRequestActions) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequestActions) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequestActions) SetName(v string) *CreateTriggerRequestActions {
	s.Name = &v
	return s
}

func (s *CreateTriggerRequestActions) SetParameters(v []*string) *CreateTriggerRequestActions {
	s.Parameters = v
	return s
}

type CreateTriggerRequestNotification struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Topic    *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateTriggerRequestNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequestNotification) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequestNotification) SetEndpoint(v string) *CreateTriggerRequestNotification {
	s.Endpoint = &v
	return s
}

func (s *CreateTriggerRequestNotification) SetTopic(v string) *CreateTriggerRequestNotification {
	s.Topic = &v
	return s
}

type CreateTriggerShrinkRequest struct {
	ActionsShrink      *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	InputShrink        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ServiceRole        *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateTriggerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerShrinkRequest) SetActionsShrink(v string) *CreateTriggerShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetInputShrink(v string) *CreateTriggerShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetNotificationShrink(v string) *CreateTriggerShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetProjectName(v string) *CreateTriggerShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetServiceRole(v string) *CreateTriggerShrinkRequest {
	s.ServiceRole = &v
	return s
}

func (s *CreateTriggerShrinkRequest) SetTagsShrink(v string) *CreateTriggerShrinkRequest {
	s.TagsShrink = &v
	return s
}

type CreateTriggerResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) SetId(v string) *CreateTriggerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTriggerResponseBody) SetRequestId(v string) *CreateTriggerResponseBody {
	s.RequestId = &v
	return s
}

type CreateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetStatusCode(v int32) *CreateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse {
	s.Body = v
	return s
}

type CreateVideoLabelClassificationTaskRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI    *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData     *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoLabelClassificationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateVideoLabelClassificationTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetNotification(v *Notification) *CreateVideoLabelClassificationTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetProjectName(v string) *CreateVideoLabelClassificationTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetSourceURI(v string) *CreateVideoLabelClassificationTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetTags(v map[string]interface{}) *CreateVideoLabelClassificationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetUserData(v string) *CreateVideoLabelClassificationTaskRequest {
	s.UserData = &v
	return s
}

type CreateVideoLabelClassificationTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI          *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoLabelClassificationTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetNotificationShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetProjectName(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetSourceURI(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetTagsShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetUserData(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.UserData = &v
	return s
}

type CreateVideoLabelClassificationTaskResponseBody struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoLabelClassificationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskResponseBody) SetEventId(v string) *CreateVideoLabelClassificationTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponseBody) SetRequestId(v string) *CreateVideoLabelClassificationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponseBody) SetTaskId(v string) *CreateVideoLabelClassificationTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVideoLabelClassificationTaskResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVideoLabelClassificationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoLabelClassificationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoLabelClassificationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponse) SetStatusCode(v int32) *CreateVideoLabelClassificationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponse) SetBody(v *CreateVideoLabelClassificationTaskResponseBody) *CreateVideoLabelClassificationTaskResponse {
	s.Body = v
	return s
}

type CreateVideoModerationTaskRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Interval         *int64            `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames        *int64            `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification *Notification          `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer     *string                `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	Scenes       []*string              `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	SourceURI    *string                `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags         map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData     *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateVideoModerationTaskRequest) SetNotification(v *Notification) *CreateVideoModerationTaskRequest {
	s.Notification = v
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
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reviewer           *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	ScenesShrink       *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	SourceURI          *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *CreateVideoModerationTaskShrinkRequest) SetNotificationShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.NotificationShrink = &v
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

type DeleteBatchRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteBatchRequest) SetId(v string) *DeleteBatchRequest {
	s.Id = &v
	return s
}

func (s *DeleteBatchRequest) SetProjectName(v string) *DeleteBatchRequest {
	s.ProjectName = &v
	return s
}

type DeleteBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBatchResponseBody) SetRequestId(v string) *DeleteBatchResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteBatchResponse) SetHeaders(v map[string]*string) *DeleteBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteBatchResponse) SetStatusCode(v int32) *DeleteBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBatchResponse) SetBody(v *DeleteBatchResponseBody) *DeleteBatchResponse {
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

type DeleteLocationDateClusterRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteLocationDateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocationDateClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteLocationDateClusterRequest) SetDatasetName(v string) *DeleteLocationDateClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteLocationDateClusterRequest) SetObjectId(v string) *DeleteLocationDateClusterRequest {
	s.ObjectId = &v
	return s
}

func (s *DeleteLocationDateClusterRequest) SetProjectName(v string) *DeleteLocationDateClusterRequest {
	s.ProjectName = &v
	return s
}

type DeleteLocationDateClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLocationDateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocationDateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLocationDateClusterResponseBody) SetRequestId(v string) *DeleteLocationDateClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLocationDateClusterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLocationDateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLocationDateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocationDateClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteLocationDateClusterResponse) SetHeaders(v map[string]*string) *DeleteLocationDateClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteLocationDateClusterResponse) SetStatusCode(v int32) *DeleteLocationDateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLocationDateClusterResponse) SetBody(v *DeleteLocationDateClusterResponseBody) *DeleteLocationDateClusterResponse {
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

type DeleteTriggerRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteTriggerRequest) SetId(v string) *DeleteTriggerRequest {
	s.Id = &v
	return s
}

func (s *DeleteTriggerRequest) SetProjectName(v string) *DeleteTriggerRequest {
	s.ProjectName = &v
	return s
}

type DeleteTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponseBody) SetRequestId(v string) *DeleteTriggerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponse) SetHeaders(v map[string]*string) *DeleteTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteTriggerResponse) SetStatusCode(v int32) *DeleteTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTriggerResponse) SetBody(v *DeleteTriggerResponseBody) *DeleteTriggerResponse {
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

type DetectImageCarsRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCarsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCarsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCarsRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageCarsRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageCarsRequest) SetProjectName(v string) *DetectImageCarsRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCarsRequest) SetSourceURI(v string) *DetectImageCarsRequest {
	s.SourceURI = &v
	return s
}

type DetectImageCarsShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCarsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCarsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCarsShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageCarsShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageCarsShrinkRequest) SetProjectName(v string) *DetectImageCarsShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCarsShrinkRequest) SetSourceURI(v string) *DetectImageCarsShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectImageCarsResponseBody struct {
	Cars      []*Car  `json:"Cars,omitempty" xml:"Cars,omitempty" type:"Repeated"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageCarsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCarsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageCarsResponseBody) SetCars(v []*Car) *DetectImageCarsResponseBody {
	s.Cars = v
	return s
}

func (s *DetectImageCarsResponseBody) SetRequestId(v string) *DetectImageCarsResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageCarsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageCarsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageCarsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageCarsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageCarsResponse) SetHeaders(v map[string]*string) *DetectImageCarsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageCarsResponse) SetStatusCode(v int32) *DetectImageCarsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageCarsResponse) SetBody(v *DetectImageCarsResponseBody) *DetectImageCarsResponse {
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

type DetectMediaMetaRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectMediaMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectMediaMetaRequest) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaRequest) SetCredentialConfig(v *CredentialConfig) *DetectMediaMetaRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectMediaMetaRequest) SetProjectName(v string) *DetectMediaMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectMediaMetaRequest) SetSourceURI(v string) *DetectMediaMetaRequest {
	s.SourceURI = &v
	return s
}

type DetectMediaMetaShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectMediaMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectMediaMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaShrinkRequest) SetCredentialConfigShrink(v string) *DetectMediaMetaShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectMediaMetaShrinkRequest) SetProjectName(v string) *DetectMediaMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectMediaMetaShrinkRequest) SetSourceURI(v string) *DetectMediaMetaShrinkRequest {
	s.SourceURI = &v
	return s
}

type DetectMediaMetaResponseBody struct {
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

func (s DetectMediaMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectMediaMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaResponseBody) SetAddresses(v []*Address) *DetectMediaMetaResponseBody {
	s.Addresses = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetAlbum(v string) *DetectMediaMetaResponseBody {
	s.Album = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetAlbumArtist(v string) *DetectMediaMetaResponseBody {
	s.AlbumArtist = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetArtist(v string) *DetectMediaMetaResponseBody {
	s.Artist = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetAudioStreams(v []*AudioStream) *DetectMediaMetaResponseBody {
	s.AudioStreams = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetBitrate(v int64) *DetectMediaMetaResponseBody {
	s.Bitrate = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetComposer(v string) *DetectMediaMetaResponseBody {
	s.Composer = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetDuration(v float64) *DetectMediaMetaResponseBody {
	s.Duration = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetFormatLongName(v string) *DetectMediaMetaResponseBody {
	s.FormatLongName = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetFormatName(v string) *DetectMediaMetaResponseBody {
	s.FormatName = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetLanguage(v string) *DetectMediaMetaResponseBody {
	s.Language = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetLatLong(v string) *DetectMediaMetaResponseBody {
	s.LatLong = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetPerformer(v string) *DetectMediaMetaResponseBody {
	s.Performer = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetProduceTime(v string) *DetectMediaMetaResponseBody {
	s.ProduceTime = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetProgramCount(v int64) *DetectMediaMetaResponseBody {
	s.ProgramCount = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetRequestId(v string) *DetectMediaMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetSize(v int64) *DetectMediaMetaResponseBody {
	s.Size = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetStartTime(v float64) *DetectMediaMetaResponseBody {
	s.StartTime = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetStreamCount(v int64) *DetectMediaMetaResponseBody {
	s.StreamCount = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetSubtitles(v []*SubtitleStream) *DetectMediaMetaResponseBody {
	s.Subtitles = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetTitle(v string) *DetectMediaMetaResponseBody {
	s.Title = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetVideoHeight(v int64) *DetectMediaMetaResponseBody {
	s.VideoHeight = &v
	return s
}

func (s *DetectMediaMetaResponseBody) SetVideoStreams(v []*VideoStream) *DetectMediaMetaResponseBody {
	s.VideoStreams = v
	return s
}

func (s *DetectMediaMetaResponseBody) SetVideoWidth(v int64) *DetectMediaMetaResponseBody {
	s.VideoWidth = &v
	return s
}

type DetectMediaMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectMediaMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectMediaMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectMediaMetaResponse) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaResponse) SetHeaders(v map[string]*string) *DetectMediaMetaResponse {
	s.Headers = v
	return s
}

func (s *DetectMediaMetaResponse) SetStatusCode(v int32) *DetectMediaMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectMediaMetaResponse) SetBody(v *DetectMediaMetaResponseBody) *DetectMediaMetaResponse {
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

type ExtractDocumentTextRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s ExtractDocumentTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractDocumentTextRequest) GoString() string {
	return s.String()
}

func (s *ExtractDocumentTextRequest) SetCredentialConfig(v *CredentialConfig) *ExtractDocumentTextRequest {
	s.CredentialConfig = v
	return s
}

func (s *ExtractDocumentTextRequest) SetProjectName(v string) *ExtractDocumentTextRequest {
	s.ProjectName = &v
	return s
}

func (s *ExtractDocumentTextRequest) SetSourceURI(v string) *ExtractDocumentTextRequest {
	s.SourceURI = &v
	return s
}

type ExtractDocumentTextShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s ExtractDocumentTextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractDocumentTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExtractDocumentTextShrinkRequest) SetCredentialConfigShrink(v string) *ExtractDocumentTextShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *ExtractDocumentTextShrinkRequest) SetProjectName(v string) *ExtractDocumentTextShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *ExtractDocumentTextShrinkRequest) SetSourceURI(v string) *ExtractDocumentTextShrinkRequest {
	s.SourceURI = &v
	return s
}

type ExtractDocumentTextResponseBody struct {
	DocumentText *string `json:"DocumentText,omitempty" xml:"DocumentText,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractDocumentTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractDocumentTextResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractDocumentTextResponseBody) SetDocumentText(v string) *ExtractDocumentTextResponseBody {
	s.DocumentText = &v
	return s
}

func (s *ExtractDocumentTextResponseBody) SetRequestId(v string) *ExtractDocumentTextResponseBody {
	s.RequestId = &v
	return s
}

type ExtractDocumentTextResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractDocumentTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractDocumentTextResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractDocumentTextResponse) GoString() string {
	return s.String()
}

func (s *ExtractDocumentTextResponse) SetHeaders(v map[string]*string) *ExtractDocumentTextResponse {
	s.Headers = v
	return s
}

func (s *ExtractDocumentTextResponse) SetStatusCode(v int32) *ExtractDocumentTextResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractDocumentTextResponse) SetBody(v *ExtractDocumentTextResponseBody) *ExtractDocumentTextResponse {
	s.Body = v
	return s
}

type FuzzyQueryRequest struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

func (s *FuzzyQueryRequest) SetOrder(v string) *FuzzyQueryRequest {
	s.Order = &v
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

func (s *FuzzyQueryRequest) SetSort(v string) *FuzzyQueryRequest {
	s.Sort = &v
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

type GenerateDRMLicenseRequest struct {
	KeyId            *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	NotifyEndpoint   *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName  *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName      *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProtectionSystem *string `json:"ProtectionSystem,omitempty" xml:"ProtectionSystem,omitempty"`
}

func (s GenerateDRMLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDRMLicenseRequest) GoString() string {
	return s.String()
}

func (s *GenerateDRMLicenseRequest) SetKeyId(v string) *GenerateDRMLicenseRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDRMLicenseRequest) SetNotifyEndpoint(v string) *GenerateDRMLicenseRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GenerateDRMLicenseRequest) SetNotifyTopicName(v string) *GenerateDRMLicenseRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GenerateDRMLicenseRequest) SetProjectName(v string) *GenerateDRMLicenseRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateDRMLicenseRequest) SetProtectionSystem(v string) *GenerateDRMLicenseRequest {
	s.ProtectionSystem = &v
	return s
}

type GenerateDRMLicenseResponseBody struct {
	DeviceInfo *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	License    *string `json:"License,omitempty" xml:"License,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	States     *int64  `json:"States,omitempty" xml:"States,omitempty"`
}

func (s GenerateDRMLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDRMLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDRMLicenseResponseBody) SetDeviceInfo(v string) *GenerateDRMLicenseResponseBody {
	s.DeviceInfo = &v
	return s
}

func (s *GenerateDRMLicenseResponseBody) SetLicense(v string) *GenerateDRMLicenseResponseBody {
	s.License = &v
	return s
}

func (s *GenerateDRMLicenseResponseBody) SetRequestId(v string) *GenerateDRMLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDRMLicenseResponseBody) SetStates(v int64) *GenerateDRMLicenseResponseBody {
	s.States = &v
	return s
}

type GenerateDRMLicenseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDRMLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateDRMLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDRMLicenseResponse) GoString() string {
	return s.String()
}

func (s *GenerateDRMLicenseResponse) SetHeaders(v map[string]*string) *GenerateDRMLicenseResponse {
	s.Headers = v
	return s
}

func (s *GenerateDRMLicenseResponse) SetStatusCode(v int32) *GenerateDRMLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDRMLicenseResponse) SetBody(v *GenerateDRMLicenseResponseBody) *GenerateDRMLicenseResponse {
	s.Body = v
	return s
}

type GenerateVideoPlaylistRequest struct {
	CredentialConfig *CredentialConfig                              `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	MasterURI        *string                                        `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	ProjectName      *string                                        `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceDuration   *float32                                       `json:"SourceDuration,omitempty" xml:"SourceDuration,omitempty"`
	SourceStartTime  *float32                                       `json:"SourceStartTime,omitempty" xml:"SourceStartTime,omitempty"`
	SourceSubtitles  []*GenerateVideoPlaylistRequestSourceSubtitles `json:"SourceSubtitles,omitempty" xml:"SourceSubtitles,omitempty" type:"Repeated"`
	SourceURI        *string                                        `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Tags             map[string]interface{}                         `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Targets          []*GenerateVideoPlaylistRequestTargets         `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s GenerateVideoPlaylistRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistRequest) SetCredentialConfig(v *CredentialConfig) *GenerateVideoPlaylistRequest {
	s.CredentialConfig = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetMasterURI(v string) *GenerateVideoPlaylistRequest {
	s.MasterURI = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetProjectName(v string) *GenerateVideoPlaylistRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceDuration(v float32) *GenerateVideoPlaylistRequest {
	s.SourceDuration = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceStartTime(v float32) *GenerateVideoPlaylistRequest {
	s.SourceStartTime = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceSubtitles(v []*GenerateVideoPlaylistRequestSourceSubtitles) *GenerateVideoPlaylistRequest {
	s.SourceSubtitles = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceURI(v string) *GenerateVideoPlaylistRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetTags(v map[string]interface{}) *GenerateVideoPlaylistRequest {
	s.Tags = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetTargets(v []*GenerateVideoPlaylistRequestTargets) *GenerateVideoPlaylistRequest {
	s.Targets = v
	return s
}

type GenerateVideoPlaylistRequestSourceSubtitles struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	URI      *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistRequestSourceSubtitles) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistRequestSourceSubtitles) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) SetLanguage(v string) *GenerateVideoPlaylistRequestSourceSubtitles {
	s.Language = &v
	return s
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) SetURI(v string) *GenerateVideoPlaylistRequestSourceSubtitles {
	s.URI = &v
	return s
}

type GenerateVideoPlaylistRequestTargets struct {
	Audio            *TargetAudio    `json:"Audio,omitempty" xml:"Audio,omitempty"`
	Duration         *float32        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InitialSegments  []*float32      `json:"InitialSegments,omitempty" xml:"InitialSegments,omitempty" type:"Repeated"`
	InitialTranscode *float32        `json:"InitialTranscode,omitempty" xml:"InitialTranscode,omitempty"`
	Speed            *float32        `json:"Speed,omitempty" xml:"Speed,omitempty"`
	Subtitle         *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	TranscodeAhead   *int32          `json:"TranscodeAhead,omitempty" xml:"TranscodeAhead,omitempty"`
	URI              *string         `json:"URI,omitempty" xml:"URI,omitempty"`
	Video            *TargetVideo    `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GenerateVideoPlaylistRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistRequestTargets) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistRequestTargets) SetAudio(v *TargetAudio) *GenerateVideoPlaylistRequestTargets {
	s.Audio = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetDuration(v float32) *GenerateVideoPlaylistRequestTargets {
	s.Duration = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetInitialSegments(v []*float32) *GenerateVideoPlaylistRequestTargets {
	s.InitialSegments = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetInitialTranscode(v float32) *GenerateVideoPlaylistRequestTargets {
	s.InitialTranscode = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetSpeed(v float32) *GenerateVideoPlaylistRequestTargets {
	s.Speed = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetSubtitle(v *TargetSubtitle) *GenerateVideoPlaylistRequestTargets {
	s.Subtitle = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetTranscodeAhead(v int32) *GenerateVideoPlaylistRequestTargets {
	s.TranscodeAhead = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetURI(v string) *GenerateVideoPlaylistRequestTargets {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetVideo(v *TargetVideo) *GenerateVideoPlaylistRequestTargets {
	s.Video = v
	return s
}

type GenerateVideoPlaylistShrinkRequest struct {
	CredentialConfigShrink *string  `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	MasterURI              *string  `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	ProjectName            *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceDuration         *float32 `json:"SourceDuration,omitempty" xml:"SourceDuration,omitempty"`
	SourceStartTime        *float32 `json:"SourceStartTime,omitempty" xml:"SourceStartTime,omitempty"`
	SourceSubtitlesShrink  *string  `json:"SourceSubtitles,omitempty" xml:"SourceSubtitles,omitempty"`
	SourceURI              *string  `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	TagsShrink             *string  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetsShrink          *string  `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s GenerateVideoPlaylistShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistShrinkRequest) SetCredentialConfigShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetMasterURI(v string) *GenerateVideoPlaylistShrinkRequest {
	s.MasterURI = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetProjectName(v string) *GenerateVideoPlaylistShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceDuration(v float32) *GenerateVideoPlaylistShrinkRequest {
	s.SourceDuration = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceStartTime(v float32) *GenerateVideoPlaylistShrinkRequest {
	s.SourceStartTime = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceSubtitlesShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.SourceSubtitlesShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetSourceURI(v string) *GenerateVideoPlaylistShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetTagsShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *GenerateVideoPlaylistShrinkRequest) SetTargetsShrink(v string) *GenerateVideoPlaylistShrinkRequest {
	s.TargetsShrink = &v
	return s
}

type GenerateVideoPlaylistResponseBody struct {
	// 转码文件列表。
	AudioPlaylist []*GenerateVideoPlaylistResponseBodyAudioPlaylist `json:"AudioPlaylist,omitempty" xml:"AudioPlaylist,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 转码文件列表。
	SubtitlePlaylist []*GenerateVideoPlaylistResponseBodySubtitlePlaylist `json:"SubtitlePlaylist,omitempty" xml:"SubtitlePlaylist,omitempty" type:"Repeated"`
	Token            *string                                              `json:"Token,omitempty" xml:"Token,omitempty"`
	URI              *string                                              `json:"URI,omitempty" xml:"URI,omitempty"`
	// 转码文件列表。
	VideoPlaylist []*GenerateVideoPlaylistResponseBodyVideoPlaylist `json:"VideoPlaylist,omitempty" xml:"VideoPlaylist,omitempty" type:"Repeated"`
}

func (s GenerateVideoPlaylistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBody) SetAudioPlaylist(v []*GenerateVideoPlaylistResponseBodyAudioPlaylist) *GenerateVideoPlaylistResponseBody {
	s.AudioPlaylist = v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetRequestId(v string) *GenerateVideoPlaylistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetSubtitlePlaylist(v []*GenerateVideoPlaylistResponseBodySubtitlePlaylist) *GenerateVideoPlaylistResponseBody {
	s.SubtitlePlaylist = v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetToken(v string) *GenerateVideoPlaylistResponseBody {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetURI(v string) *GenerateVideoPlaylistResponseBody {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBody) SetVideoPlaylist(v []*GenerateVideoPlaylistResponseBodyVideoPlaylist) *GenerateVideoPlaylistResponseBody {
	s.VideoPlaylist = v
	return s
}

type GenerateVideoPlaylistResponseBodyAudioPlaylist struct {
	// 转码生成的Token。用于LiveTranscoding访问的参数。
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// 输出m3u8的OSS地址。地址规则为 Target.URI + ".m3u8“， 其中Target.URI为输入参数中视频转码输出地址前缀。
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistResponseBodyAudioPlaylist) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBodyAudioPlaylist) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) SetToken(v string) *GenerateVideoPlaylistResponseBodyAudioPlaylist {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyAudioPlaylist) SetURI(v string) *GenerateVideoPlaylistResponseBodyAudioPlaylist {
	s.URI = &v
	return s
}

type GenerateVideoPlaylistResponseBodySubtitlePlaylist struct {
	// 字幕流编号，从0开始。
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// 视频源中字幕流的语言。
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 转码生成的Token。用于LiveTranscoding访问的参数。
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// 输出m3u8的OSS地址。地址规则为 Target.URI + “_” + Index + ".m3u8“， 其中Target.URI为输入参数中视频转码输出地址前缀。
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistResponseBodySubtitlePlaylist) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBodySubtitlePlaylist) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetIndex(v int32) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.Index = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetLanguage(v string) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.Language = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetToken(v string) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodySubtitlePlaylist) SetURI(v string) *GenerateVideoPlaylistResponseBodySubtitlePlaylist {
	s.URI = &v
	return s
}

type GenerateVideoPlaylistResponseBodyVideoPlaylist struct {
	// 转码生成的Token。用于LiveTranscoding访问的参数。
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// 输出m3u8的OSS地址。地址规则为 Target.URI + ".m3u8“， 其中Target.URI为输入参数中视频转码输出地址前缀。
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistResponseBodyVideoPlaylist) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistResponseBodyVideoPlaylist) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) SetToken(v string) *GenerateVideoPlaylistResponseBodyVideoPlaylist {
	s.Token = &v
	return s
}

func (s *GenerateVideoPlaylistResponseBodyVideoPlaylist) SetURI(v string) *GenerateVideoPlaylistResponseBodyVideoPlaylist {
	s.URI = &v
	return s
}

type GenerateVideoPlaylistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateVideoPlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVideoPlaylistResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoPlaylistResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponse) SetHeaders(v map[string]*string) *GenerateVideoPlaylistResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoPlaylistResponse) SetStatusCode(v int32) *GenerateVideoPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateVideoPlaylistResponse) SetBody(v *GenerateVideoPlaylistResponseBody) *GenerateVideoPlaylistResponse {
	s.Body = v
	return s
}

type GenerateWebofficeTokenRequest struct {
	CachePreview     *bool             `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ExternalUploaded *bool             `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	Filename         *string           `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Hidecmb          *bool             `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	Notification    *Notification        `json:"Notification,omitempty" xml:"Notification,omitempty"`
	NotifyTopicName *string              `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Password        *string              `json:"Password,omitempty" xml:"Password,omitempty"`
	Permission      *WebofficePermission `json:"Permission,omitempty" xml:"Permission,omitempty"`
	PreviewPages    *int64               `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	ProjectName     *string              `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Referer         *string              `json:"Referer,omitempty" xml:"Referer,omitempty"`
	SourceURI       *string              `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	User            *WebofficeUser       `json:"User,omitempty" xml:"User,omitempty"`
	UserData        *string              `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Watermark       *WebofficeWatermark  `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GenerateWebofficeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenRequest) SetCachePreview(v bool) *GenerateWebofficeTokenRequest {
	s.CachePreview = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetCredentialConfig(v *CredentialConfig) *GenerateWebofficeTokenRequest {
	s.CredentialConfig = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetExternalUploaded(v bool) *GenerateWebofficeTokenRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetFilename(v string) *GenerateWebofficeTokenRequest {
	s.Filename = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetHidecmb(v bool) *GenerateWebofficeTokenRequest {
	s.Hidecmb = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetNotification(v *Notification) *GenerateWebofficeTokenRequest {
	s.Notification = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetNotifyTopicName(v string) *GenerateWebofficeTokenRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetPassword(v string) *GenerateWebofficeTokenRequest {
	s.Password = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetPermission(v *WebofficePermission) *GenerateWebofficeTokenRequest {
	s.Permission = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetPreviewPages(v int64) *GenerateWebofficeTokenRequest {
	s.PreviewPages = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetProjectName(v string) *GenerateWebofficeTokenRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetReferer(v string) *GenerateWebofficeTokenRequest {
	s.Referer = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetSourceURI(v string) *GenerateWebofficeTokenRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetUser(v *WebofficeUser) *GenerateWebofficeTokenRequest {
	s.User = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetUserData(v string) *GenerateWebofficeTokenRequest {
	s.UserData = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetWatermark(v *WebofficeWatermark) *GenerateWebofficeTokenRequest {
	s.Watermark = v
	return s
}

type GenerateWebofficeTokenShrinkRequest struct {
	CachePreview           *bool   `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ExternalUploaded       *bool   `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	Filename               *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Hidecmb                *bool   `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// 消息通知配置，支持使用MNS、RocketMQ接收异步消息通知。
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	NotifyTopicName    *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	Password           *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PermissionShrink   *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	PreviewPages       *int64  `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Referer            *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	SourceURI          *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	UserShrink         *string `json:"User,omitempty" xml:"User,omitempty"`
	UserData           *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WatermarkShrink    *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GenerateWebofficeTokenShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebofficeTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenShrinkRequest) SetCachePreview(v bool) *GenerateWebofficeTokenShrinkRequest {
	s.CachePreview = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetCredentialConfigShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetExternalUploaded(v bool) *GenerateWebofficeTokenShrinkRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetFilename(v string) *GenerateWebofficeTokenShrinkRequest {
	s.Filename = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetHidecmb(v bool) *GenerateWebofficeTokenShrinkRequest {
	s.Hidecmb = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetNotificationShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetNotifyTopicName(v string) *GenerateWebofficeTokenShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetPassword(v string) *GenerateWebofficeTokenShrinkRequest {
	s.Password = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetPermissionShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.PermissionShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetPreviewPages(v int64) *GenerateWebofficeTokenShrinkRequest {
	s.PreviewPages = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetProjectName(v string) *GenerateWebofficeTokenShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetReferer(v string) *GenerateWebofficeTokenShrinkRequest {
	s.Referer = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetSourceURI(v string) *GenerateWebofficeTokenShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetUserShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetUserData(v string) *GenerateWebofficeTokenShrinkRequest {
	s.UserData = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetWatermarkShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.WatermarkShrink = &v
	return s
}

type GenerateWebofficeTokenResponseBody struct {
	AccessToken             *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessTokenExpiredTime  *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	RefreshToken            *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebofficeURL            *string `json:"WebofficeURL,omitempty" xml:"WebofficeURL,omitempty"`
}

func (s GenerateWebofficeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenResponseBody) SetAccessToken(v string) *GenerateWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *GenerateWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetRefreshToken(v string) *GenerateWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *GenerateWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetRequestId(v string) *GenerateWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebofficeTokenResponseBody) SetWebofficeURL(v string) *GenerateWebofficeTokenResponseBody {
	s.WebofficeURL = &v
	return s
}

type GenerateWebofficeTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateWebofficeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateWebofficeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebofficeTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenResponse) SetHeaders(v map[string]*string) *GenerateWebofficeTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebofficeTokenResponse) SetStatusCode(v int32) *GenerateWebofficeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWebofficeTokenResponse) SetBody(v *GenerateWebofficeTokenResponseBody) *GenerateWebofficeTokenResponse {
	s.Body = v
	return s
}

type GetBatchRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBatchRequest) GoString() string {
	return s.String()
}

func (s *GetBatchRequest) SetId(v string) *GetBatchRequest {
	s.Id = &v
	return s
}

func (s *GetBatchRequest) SetProjectName(v string) *GetBatchRequest {
	s.ProjectName = &v
	return s
}

type GetBatchResponseBody struct {
	Batch     *DataIngestion `json:"Batch,omitempty" xml:"Batch,omitempty"`
	RequestId *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBatchResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchResponseBody) SetBatch(v *DataIngestion) *GetBatchResponseBody {
	s.Batch = v
	return s
}

func (s *GetBatchResponseBody) SetRequestId(v string) *GetBatchResponseBody {
	s.RequestId = &v
	return s
}

type GetBatchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBatchResponse) GoString() string {
	return s.String()
}

func (s *GetBatchResponse) SetHeaders(v map[string]*string) *GetBatchResponse {
	s.Headers = v
	return s
}

func (s *GetBatchResponse) SetStatusCode(v int32) *GetBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchResponse) SetBody(v *GetBatchResponseBody) *GetBatchResponse {
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

type GetDRMLicenseRequest struct {
	KeyId            *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	NotifyEndpoint   *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	NotifyTopicName  *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	ProjectName      *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProtectionSystem *string `json:"ProtectionSystem,omitempty" xml:"ProtectionSystem,omitempty"`
}

func (s GetDRMLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDRMLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseRequest) SetKeyId(v string) *GetDRMLicenseRequest {
	s.KeyId = &v
	return s
}

func (s *GetDRMLicenseRequest) SetNotifyEndpoint(v string) *GetDRMLicenseRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetDRMLicenseRequest) SetNotifyTopicName(v string) *GetDRMLicenseRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetDRMLicenseRequest) SetProjectName(v string) *GetDRMLicenseRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDRMLicenseRequest) SetProtectionSystem(v string) *GetDRMLicenseRequest {
	s.ProtectionSystem = &v
	return s
}

type GetDRMLicenseResponseBody struct {
	DeviceInfo *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	License    *string `json:"License,omitempty" xml:"License,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	States     *int64  `json:"States,omitempty" xml:"States,omitempty"`
}

func (s GetDRMLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDRMLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponseBody) SetDeviceInfo(v string) *GetDRMLicenseResponseBody {
	s.DeviceInfo = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetLicense(v string) *GetDRMLicenseResponseBody {
	s.License = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetRequestId(v string) *GetDRMLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetStates(v int64) *GetDRMLicenseResponseBody {
	s.States = &v
	return s
}

type GetDRMLicenseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDRMLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDRMLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDRMLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponse) SetHeaders(v map[string]*string) *GetDRMLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetDRMLicenseResponse) SetStatusCode(v int32) *GetDRMLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDRMLicenseResponse) SetBody(v *GetDRMLicenseResponseBody) *GetDRMLicenseResponse {
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

type GetImageModerationResultRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetImageModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageModerationResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultRequest) SetProjectName(v string) *GetImageModerationResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetImageModerationResultRequest) SetTaskId(v string) *GetImageModerationResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageModerationResultRequest) SetTaskType(v string) *GetImageModerationResultRequest {
	s.TaskType = &v
	return s
}

type GetImageModerationResultResponseBody struct {
	Code             *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime          *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventId          *string                                               `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ModerationResult *GetImageModerationResultResponseBodyModerationResult `json:"ModerationResult,omitempty" xml:"ModerationResult,omitempty" type:"Struct"`
	ProjectName      *string                                               `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime        *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId           *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType         *string                                               `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData         *string                                               `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetImageModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBody) SetCode(v string) *GetImageModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetEndTime(v string) *GetImageModerationResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetEventId(v string) *GetImageModerationResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetMessage(v string) *GetImageModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetModerationResult(v *GetImageModerationResultResponseBodyModerationResult) *GetImageModerationResultResponseBody {
	s.ModerationResult = v
	return s
}

func (s *GetImageModerationResultResponseBody) SetProjectName(v string) *GetImageModerationResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetRequestId(v string) *GetImageModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetStartTime(v string) *GetImageModerationResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetStatus(v string) *GetImageModerationResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetTaskId(v string) *GetImageModerationResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetTaskType(v string) *GetImageModerationResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetUserData(v string) *GetImageModerationResultResponseBody {
	s.UserData = &v
	return s
}

type GetImageModerationResultResponseBodyModerationResult struct {
	Categories []*string                                                   `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Frames     *GetImageModerationResultResponseBodyModerationResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Struct"`
	Suggestion *string                                                     `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	URI        *string                                                     `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetImageModerationResultResponseBodyModerationResult) String() string {
	return tea.Prettify(s)
}

func (s GetImageModerationResultResponseBodyModerationResult) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetCategories(v []*string) *GetImageModerationResultResponseBodyModerationResult {
	s.Categories = v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetFrames(v *GetImageModerationResultResponseBodyModerationResultFrames) *GetImageModerationResultResponseBodyModerationResult {
	s.Frames = v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetSuggestion(v string) *GetImageModerationResultResponseBodyModerationResult {
	s.Suggestion = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetURI(v string) *GetImageModerationResultResponseBodyModerationResult {
	s.URI = &v
	return s
}

type GetImageModerationResultResponseBodyModerationResultFrames struct {
	BlockFrames []*GetImageModerationResultResponseBodyModerationResultFramesBlockFrames `json:"BlockFrames,omitempty" xml:"BlockFrames,omitempty" type:"Repeated"`
	TotalCount  *int32                                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetImageModerationResultResponseBodyModerationResultFrames) String() string {
	return tea.Prettify(s)
}

func (s GetImageModerationResultResponseBodyModerationResultFrames) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) SetBlockFrames(v []*GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) *GetImageModerationResultResponseBodyModerationResultFrames {
	s.BlockFrames = v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) SetTotalCount(v int32) *GetImageModerationResultResponseBodyModerationResultFrames {
	s.TotalCount = &v
	return s
}

type GetImageModerationResultResponseBodyModerationResultFramesBlockFrames struct {
	Label  *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Offset *int32   `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Rate   *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) String() string {
	return tea.Prettify(s)
}

func (s GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) SetLabel(v string) *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Label = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) SetOffset(v int32) *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Offset = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) SetRate(v float64) *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Rate = &v
	return s
}

type GetImageModerationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageModerationResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponse) SetHeaders(v map[string]*string) *GetImageModerationResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageModerationResultResponse) SetStatusCode(v int32) *GetImageModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageModerationResultResponse) SetBody(v *GetImageModerationResultResponseBody) *GetImageModerationResultResponse {
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

type GetTriggerRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerRequest) GoString() string {
	return s.String()
}

func (s *GetTriggerRequest) SetId(v string) *GetTriggerRequest {
	s.Id = &v
	return s
}

func (s *GetTriggerRequest) SetProjectName(v string) *GetTriggerRequest {
	s.ProjectName = &v
	return s
}

type GetTriggerResponseBody struct {
	RequestId *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Trigger   *DataIngestion `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s GetTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *GetTriggerResponseBody) SetRequestId(v string) *GetTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTriggerResponseBody) SetTrigger(v *DataIngestion) *GetTriggerResponseBody {
	s.Trigger = v
	return s
}

type GetTriggerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetTriggerResponse) SetHeaders(v map[string]*string) *GetTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetTriggerResponse) SetStatusCode(v int32) *GetTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTriggerResponse) SetBody(v *GetTriggerResponseBody) *GetTriggerResponse {
	s.Body = v
	return s
}

type GetVideoLabelClassificationResultRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVideoLabelClassificationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoLabelClassificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoLabelClassificationResultRequest) SetProjectName(v string) *GetVideoLabelClassificationResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetVideoLabelClassificationResultRequest) SetTaskId(v string) *GetVideoLabelClassificationResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoLabelClassificationResultRequest) SetTaskType(v string) *GetVideoLabelClassificationResultRequest {
	s.TaskType = &v
	return s
}

type GetVideoLabelClassificationResultResponseBody struct {
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

func (s GetVideoLabelClassificationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoLabelClassificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoLabelClassificationResultResponseBody) SetCode(v string) *GetVideoLabelClassificationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetEndTime(v string) *GetVideoLabelClassificationResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetEventId(v string) *GetVideoLabelClassificationResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetLabels(v []*Label) *GetVideoLabelClassificationResultResponseBody {
	s.Labels = v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetMessage(v string) *GetVideoLabelClassificationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetProjectName(v string) *GetVideoLabelClassificationResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetRequestId(v string) *GetVideoLabelClassificationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetStartTime(v string) *GetVideoLabelClassificationResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetStatus(v string) *GetVideoLabelClassificationResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetTaskId(v string) *GetVideoLabelClassificationResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetTaskType(v string) *GetVideoLabelClassificationResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetUserData(v string) *GetVideoLabelClassificationResultResponseBody {
	s.UserData = &v
	return s
}

type GetVideoLabelClassificationResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoLabelClassificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoLabelClassificationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoLabelClassificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoLabelClassificationResultResponse) SetHeaders(v map[string]*string) *GetVideoLabelClassificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoLabelClassificationResultResponse) SetStatusCode(v int32) *GetVideoLabelClassificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponse) SetBody(v *GetVideoLabelClassificationResultResponseBody) *GetVideoLabelClassificationResultResponse {
	s.Body = v
	return s
}

type GetVideoModerationResultRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVideoModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoModerationResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultRequest) SetProjectName(v string) *GetVideoModerationResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetVideoModerationResultRequest) SetTaskId(v string) *GetVideoModerationResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoModerationResultRequest) SetTaskType(v string) *GetVideoModerationResultRequest {
	s.TaskType = &v
	return s
}

type GetVideoModerationResultResponseBody struct {
	Code             *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime          *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventId          *string                                               `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ModerationResult *GetVideoModerationResultResponseBodyModerationResult `json:"ModerationResult,omitempty" xml:"ModerationResult,omitempty" type:"Struct"`
	ProjectName      *string                                               `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime        *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId           *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType         *string                                               `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData         *string                                               `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetVideoModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBody) SetCode(v string) *GetVideoModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetEndTime(v string) *GetVideoModerationResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetEventId(v string) *GetVideoModerationResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetMessage(v string) *GetVideoModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetModerationResult(v *GetVideoModerationResultResponseBodyModerationResult) *GetVideoModerationResultResponseBody {
	s.ModerationResult = v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetProjectName(v string) *GetVideoModerationResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetRequestId(v string) *GetVideoModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetStartTime(v string) *GetVideoModerationResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetStatus(v string) *GetVideoModerationResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetTaskId(v string) *GetVideoModerationResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetTaskType(v string) *GetVideoModerationResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetUserData(v string) *GetVideoModerationResultResponseBody {
	s.UserData = &v
	return s
}

type GetVideoModerationResultResponseBodyModerationResult struct {
	Categories []*string                                                   `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Frames     *GetVideoModerationResultResponseBodyModerationResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Struct"`
	Suggestion *string                                                     `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	URI        *string                                                     `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetVideoModerationResultResponseBodyModerationResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoModerationResultResponseBodyModerationResult) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetCategories(v []*string) *GetVideoModerationResultResponseBodyModerationResult {
	s.Categories = v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetFrames(v *GetVideoModerationResultResponseBodyModerationResultFrames) *GetVideoModerationResultResponseBodyModerationResult {
	s.Frames = v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetSuggestion(v string) *GetVideoModerationResultResponseBodyModerationResult {
	s.Suggestion = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetURI(v string) *GetVideoModerationResultResponseBodyModerationResult {
	s.URI = &v
	return s
}

type GetVideoModerationResultResponseBodyModerationResultFrames struct {
	BlockFrames []*GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames `json:"BlockFrames,omitempty" xml:"BlockFrames,omitempty" type:"Repeated"`
	TotalCount  *int32                                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVideoModerationResultResponseBodyModerationResultFrames) String() string {
	return tea.Prettify(s)
}

func (s GetVideoModerationResultResponseBodyModerationResultFrames) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) SetBlockFrames(v []*GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) *GetVideoModerationResultResponseBodyModerationResultFrames {
	s.BlockFrames = v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) SetTotalCount(v int32) *GetVideoModerationResultResponseBodyModerationResultFrames {
	s.TotalCount = &v
	return s
}

type GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames struct {
	Label  *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Offset *int32   `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Rate   *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) String() string {
	return tea.Prettify(s)
}

func (s GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) SetLabel(v string) *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Label = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) SetOffset(v int32) *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Offset = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) SetRate(v float64) *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Rate = &v
	return s
}

type GetVideoModerationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoModerationResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponse) SetHeaders(v map[string]*string) *GetVideoModerationResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoModerationResultResponse) SetStatusCode(v int32) *GetVideoModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoModerationResultResponse) SetBody(v *GetVideoModerationResultResponseBody) *GetVideoModerationResultResponse {
	s.Body = v
	return s
}

type IndexFileMetaRequest struct {
	DatasetName  *string       `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	File         *InputFile    `json:"File,omitempty" xml:"File,omitempty"`
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName  *string       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *IndexFileMetaRequest) SetFile(v *InputFile) *IndexFileMetaRequest {
	s.File = v
	return s
}

func (s *IndexFileMetaRequest) SetNotification(v *Notification) *IndexFileMetaRequest {
	s.Notification = v
	return s
}

func (s *IndexFileMetaRequest) SetProjectName(v string) *IndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

type IndexFileMetaShrinkRequest struct {
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	FileShrink         *string `json:"File,omitempty" xml:"File,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *IndexFileMetaShrinkRequest) SetNotificationShrink(v string) *IndexFileMetaShrinkRequest {
	s.NotificationShrink = &v
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

type ListBatchesRequest struct {
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	TagSelector *string `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
}

func (s ListBatchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBatchesRequest) GoString() string {
	return s.String()
}

func (s *ListBatchesRequest) SetMaxResults(v int32) *ListBatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBatchesRequest) SetNextToken(v string) *ListBatchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListBatchesRequest) SetOrder(v string) *ListBatchesRequest {
	s.Order = &v
	return s
}

func (s *ListBatchesRequest) SetProjectName(v string) *ListBatchesRequest {
	s.ProjectName = &v
	return s
}

func (s *ListBatchesRequest) SetSort(v string) *ListBatchesRequest {
	s.Sort = &v
	return s
}

func (s *ListBatchesRequest) SetState(v string) *ListBatchesRequest {
	s.State = &v
	return s
}

func (s *ListBatchesRequest) SetTagSelector(v string) *ListBatchesRequest {
	s.TagSelector = &v
	return s
}

type ListBatchesResponseBody struct {
	Batches   []*DataIngestion `json:"Batches,omitempty" xml:"Batches,omitempty" type:"Repeated"`
	NextToken *string          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBatchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchesResponseBody) SetBatches(v []*DataIngestion) *ListBatchesResponseBody {
	s.Batches = v
	return s
}

func (s *ListBatchesResponseBody) SetNextToken(v string) *ListBatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBatchesResponseBody) SetRequestId(v string) *ListBatchesResponseBody {
	s.RequestId = &v
	return s
}

type ListBatchesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBatchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBatchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBatchesResponse) GoString() string {
	return s.String()
}

func (s *ListBatchesResponse) SetHeaders(v map[string]*string) *ListBatchesResponse {
	s.Headers = v
	return s
}

func (s *ListBatchesResponse) SetStatusCode(v int32) *ListBatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchesResponse) SetBody(v *ListBatchesResponseBody) *ListBatchesResponse {
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

type ListTriggersRequest struct {
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	TagSelector *string `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) SetMaxResults(v int32) *ListTriggersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTriggersRequest) SetNextToken(v string) *ListTriggersRequest {
	s.NextToken = &v
	return s
}

func (s *ListTriggersRequest) SetOrder(v string) *ListTriggersRequest {
	s.Order = &v
	return s
}

func (s *ListTriggersRequest) SetProjectName(v string) *ListTriggersRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTriggersRequest) SetSort(v string) *ListTriggersRequest {
	s.Sort = &v
	return s
}

func (s *ListTriggersRequest) SetState(v string) *ListTriggersRequest {
	s.State = &v
	return s
}

func (s *ListTriggersRequest) SetTagSelector(v string) *ListTriggersRequest {
	s.TagSelector = &v
	return s
}

type ListTriggersResponseBody struct {
	NextToken *string          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Triggers  []*DataIngestion `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListTriggersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBody) SetNextToken(v string) *ListTriggersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTriggersResponseBody) SetRequestId(v string) *ListTriggersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTriggersResponseBody) SetTriggers(v []*DataIngestion) *ListTriggersResponseBody {
	s.Triggers = v
	return s
}

type ListTriggersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListTriggersResponse) SetHeaders(v map[string]*string) *ListTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListTriggersResponse) SetStatusCode(v int32) *ListTriggersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTriggersResponse) SetBody(v *ListTriggersResponseBody) *ListTriggersResponse {
	s.Body = v
	return s
}

type LiveTranscodingRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Token            *string           `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s LiveTranscodingRequest) String() string {
	return tea.Prettify(s)
}

func (s LiveTranscodingRequest) GoString() string {
	return s.String()
}

func (s *LiveTranscodingRequest) SetCredentialConfig(v *CredentialConfig) *LiveTranscodingRequest {
	s.CredentialConfig = v
	return s
}

func (s *LiveTranscodingRequest) SetProjectName(v string) *LiveTranscodingRequest {
	s.ProjectName = &v
	return s
}

func (s *LiveTranscodingRequest) SetSourceURI(v string) *LiveTranscodingRequest {
	s.SourceURI = &v
	return s
}

func (s *LiveTranscodingRequest) SetToken(v string) *LiveTranscodingRequest {
	s.Token = &v
	return s
}

type LiveTranscodingShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Token                  *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s LiveTranscodingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s LiveTranscodingShrinkRequest) GoString() string {
	return s.String()
}

func (s *LiveTranscodingShrinkRequest) SetCredentialConfigShrink(v string) *LiveTranscodingShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *LiveTranscodingShrinkRequest) SetProjectName(v string) *LiveTranscodingShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *LiveTranscodingShrinkRequest) SetSourceURI(v string) *LiveTranscodingShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *LiveTranscodingShrinkRequest) SetToken(v string) *LiveTranscodingShrinkRequest {
	s.Token = &v
	return s
}

type LiveTranscodingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	URI       *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s LiveTranscodingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiveTranscodingResponseBody) GoString() string {
	return s.String()
}

func (s *LiveTranscodingResponseBody) SetRequestId(v string) *LiveTranscodingResponseBody {
	s.RequestId = &v
	return s
}

func (s *LiveTranscodingResponseBody) SetURI(v string) *LiveTranscodingResponseBody {
	s.URI = &v
	return s
}

type LiveTranscodingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LiveTranscodingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LiveTranscodingResponse) String() string {
	return tea.Prettify(s)
}

func (s LiveTranscodingResponse) GoString() string {
	return s.String()
}

func (s *LiveTranscodingResponse) SetHeaders(v map[string]*string) *LiveTranscodingResponse {
	s.Headers = v
	return s
}

func (s *LiveTranscodingResponse) SetStatusCode(v int32) *LiveTranscodingResponse {
	s.StatusCode = &v
	return s
}

func (s *LiveTranscodingResponse) SetBody(v *LiveTranscodingResponseBody) *LiveTranscodingResponse {
	s.Body = v
	return s
}

type QueryFigureClustersRequest struct {
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	CustomLabels    *string    `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName     *string    `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults      *int64     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order           *string    `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName     *string    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort            *string    `json:"Sort,omitempty" xml:"Sort,omitempty"`
	UpdateTimeRange *TimeRange `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
	WithTotalCount  *bool      `json:"WithTotalCount,omitempty" xml:"WithTotalCount,omitempty"`
}

func (s QueryFigureClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFigureClustersRequest) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersRequest) SetCreateTimeRange(v *TimeRange) *QueryFigureClustersRequest {
	s.CreateTimeRange = v
	return s
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

func (s *QueryFigureClustersRequest) SetUpdateTimeRange(v *TimeRange) *QueryFigureClustersRequest {
	s.UpdateTimeRange = v
	return s
}

func (s *QueryFigureClustersRequest) SetWithTotalCount(v bool) *QueryFigureClustersRequest {
	s.WithTotalCount = &v
	return s
}

type QueryFigureClustersShrinkRequest struct {
	CreateTimeRangeShrink *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	CustomLabels          *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName           *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults            *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order                 *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName           *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort                  *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	UpdateTimeRangeShrink *string `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
	WithTotalCount        *bool   `json:"WithTotalCount,omitempty" xml:"WithTotalCount,omitempty"`
}

func (s QueryFigureClustersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFigureClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersShrinkRequest) SetCreateTimeRangeShrink(v string) *QueryFigureClustersShrinkRequest {
	s.CreateTimeRangeShrink = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetCustomLabels(v string) *QueryFigureClustersShrinkRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetDatasetName(v string) *QueryFigureClustersShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetMaxResults(v int64) *QueryFigureClustersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetNextToken(v string) *QueryFigureClustersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetOrder(v string) *QueryFigureClustersShrinkRequest {
	s.Order = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetProjectName(v string) *QueryFigureClustersShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetSort(v string) *QueryFigureClustersShrinkRequest {
	s.Sort = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetUpdateTimeRangeShrink(v string) *QueryFigureClustersShrinkRequest {
	s.UpdateTimeRangeShrink = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetWithTotalCount(v bool) *QueryFigureClustersShrinkRequest {
	s.WithTotalCount = &v
	return s
}

type QueryFigureClustersResponseBody struct {
	FigureClusters []*FigureCluster `json:"FigureClusters,omitempty" xml:"FigureClusters,omitempty" type:"Repeated"`
	NextToken      *string          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int64           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *QueryFigureClustersResponseBody) SetTotalCount(v int64) *QueryFigureClustersResponseBody {
	s.TotalCount = &v
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

type QueryLocationDateClustersRequest struct {
	Address                           *Address   `json:"Address,omitempty" xml:"Address,omitempty"`
	CreateTimeRange                   *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	CustomLabels                      *string    `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName                       *string    `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	LocationDateClusterEndTimeRange   *TimeRange `json:"LocationDateClusterEndTimeRange,omitempty" xml:"LocationDateClusterEndTimeRange,omitempty"`
	LocationDateClusterLevels         []*string  `json:"LocationDateClusterLevels,omitempty" xml:"LocationDateClusterLevels,omitempty" type:"Repeated"`
	LocationDateClusterStartTimeRange *TimeRange `json:"LocationDateClusterStartTimeRange,omitempty" xml:"LocationDateClusterStartTimeRange,omitempty"`
	MaxResults                        *int32     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                         *string    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ObjectId                          *string    `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Order                             *string    `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName                       *string    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort                              *string    `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Title                             *string    `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTimeRange                   *TimeRange `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
}

func (s QueryLocationDateClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLocationDateClustersRequest) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersRequest) SetAddress(v *Address) *QueryLocationDateClustersRequest {
	s.Address = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetCreateTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.CreateTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetCustomLabels(v string) *QueryLocationDateClustersRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetDatasetName(v string) *QueryLocationDateClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetLocationDateClusterEndTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.LocationDateClusterEndTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetLocationDateClusterLevels(v []*string) *QueryLocationDateClustersRequest {
	s.LocationDateClusterLevels = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetLocationDateClusterStartTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.LocationDateClusterStartTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetMaxResults(v int32) *QueryLocationDateClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetNextToken(v string) *QueryLocationDateClustersRequest {
	s.NextToken = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetObjectId(v string) *QueryLocationDateClustersRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetOrder(v string) *QueryLocationDateClustersRequest {
	s.Order = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetProjectName(v string) *QueryLocationDateClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetSort(v string) *QueryLocationDateClustersRequest {
	s.Sort = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetTitle(v string) *QueryLocationDateClustersRequest {
	s.Title = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetUpdateTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.UpdateTimeRange = v
	return s
}

type QueryLocationDateClustersShrinkRequest struct {
	AddressShrink                           *string `json:"Address,omitempty" xml:"Address,omitempty"`
	CreateTimeRangeShrink                   *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	CustomLabels                            *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName                             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	LocationDateClusterEndTimeRangeShrink   *string `json:"LocationDateClusterEndTimeRange,omitempty" xml:"LocationDateClusterEndTimeRange,omitempty"`
	LocationDateClusterLevelsShrink         *string `json:"LocationDateClusterLevels,omitempty" xml:"LocationDateClusterLevels,omitempty"`
	LocationDateClusterStartTimeRangeShrink *string `json:"LocationDateClusterStartTimeRange,omitempty" xml:"LocationDateClusterStartTimeRange,omitempty"`
	MaxResults                              *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                               *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ObjectId                                *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Order                                   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName                             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort                                    *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Title                                   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTimeRangeShrink                   *string `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
}

func (s QueryLocationDateClustersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLocationDateClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersShrinkRequest) SetAddressShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.AddressShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetCreateTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.CreateTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetCustomLabels(v string) *QueryLocationDateClustersShrinkRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetDatasetName(v string) *QueryLocationDateClustersShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetLocationDateClusterEndTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.LocationDateClusterEndTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetLocationDateClusterLevelsShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.LocationDateClusterLevelsShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetLocationDateClusterStartTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.LocationDateClusterStartTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetMaxResults(v int32) *QueryLocationDateClustersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetNextToken(v string) *QueryLocationDateClustersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetObjectId(v string) *QueryLocationDateClustersShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetOrder(v string) *QueryLocationDateClustersShrinkRequest {
	s.Order = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetProjectName(v string) *QueryLocationDateClustersShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetSort(v string) *QueryLocationDateClustersShrinkRequest {
	s.Sort = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetTitle(v string) *QueryLocationDateClustersShrinkRequest {
	s.Title = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetUpdateTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.UpdateTimeRangeShrink = &v
	return s
}

type QueryLocationDateClustersResponseBody struct {
	LocationDateClusters []*LocationDateCluster `json:"LocationDateClusters,omitempty" xml:"LocationDateClusters,omitempty" type:"Repeated"`
	NextToken            *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryLocationDateClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLocationDateClustersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersResponseBody) SetLocationDateClusters(v []*LocationDateCluster) *QueryLocationDateClustersResponseBody {
	s.LocationDateClusters = v
	return s
}

func (s *QueryLocationDateClustersResponseBody) SetNextToken(v string) *QueryLocationDateClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryLocationDateClustersResponseBody) SetRequestId(v string) *QueryLocationDateClustersResponseBody {
	s.RequestId = &v
	return s
}

type QueryLocationDateClustersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryLocationDateClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryLocationDateClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLocationDateClustersResponse) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersResponse) SetHeaders(v map[string]*string) *QueryLocationDateClustersResponse {
	s.Headers = v
	return s
}

func (s *QueryLocationDateClustersResponse) SetStatusCode(v int32) *QueryLocationDateClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocationDateClustersResponse) SetBody(v *QueryLocationDateClustersResponseBody) *QueryLocationDateClustersResponse {
	s.Body = v
	return s
}

type QuerySimilarImageClustersRequest struct {
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Sort         *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s QuerySimilarImageClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySimilarImageClustersRequest) GoString() string {
	return s.String()
}

func (s *QuerySimilarImageClustersRequest) SetCustomLabels(v string) *QuerySimilarImageClustersRequest {
	s.CustomLabels = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetDatasetName(v string) *QuerySimilarImageClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetMaxResults(v int32) *QuerySimilarImageClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetNextToken(v string) *QuerySimilarImageClustersRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetOrder(v string) *QuerySimilarImageClustersRequest {
	s.Order = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetProjectName(v string) *QuerySimilarImageClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *QuerySimilarImageClustersRequest) SetSort(v string) *QuerySimilarImageClustersRequest {
	s.Sort = &v
	return s
}

type QuerySimilarImageClustersResponseBody struct {
	NextToken            *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimilarImageClusters []*SimilarImageCluster `json:"SimilarImageClusters,omitempty" xml:"SimilarImageClusters,omitempty" type:"Repeated"`
}

func (s QuerySimilarImageClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySimilarImageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySimilarImageClustersResponseBody) SetNextToken(v string) *QuerySimilarImageClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QuerySimilarImageClustersResponseBody) SetRequestId(v string) *QuerySimilarImageClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySimilarImageClustersResponseBody) SetSimilarImageClusters(v []*SimilarImageCluster) *QuerySimilarImageClustersResponseBody {
	s.SimilarImageClusters = v
	return s
}

type QuerySimilarImageClustersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySimilarImageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySimilarImageClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySimilarImageClustersResponse) GoString() string {
	return s.String()
}

func (s *QuerySimilarImageClustersResponse) SetHeaders(v map[string]*string) *QuerySimilarImageClustersResponse {
	s.Headers = v
	return s
}

func (s *QuerySimilarImageClustersResponse) SetStatusCode(v int32) *QuerySimilarImageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySimilarImageClustersResponse) SetBody(v *QuerySimilarImageClustersResponseBody) *QuerySimilarImageClustersResponse {
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

type ResumeBatchRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ResumeBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeBatchRequest) GoString() string {
	return s.String()
}

func (s *ResumeBatchRequest) SetId(v string) *ResumeBatchRequest {
	s.Id = &v
	return s
}

func (s *ResumeBatchRequest) SetProjectName(v string) *ResumeBatchRequest {
	s.ProjectName = &v
	return s
}

type ResumeBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeBatchResponseBody) SetRequestId(v string) *ResumeBatchResponseBody {
	s.RequestId = &v
	return s
}

type ResumeBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeBatchResponse) GoString() string {
	return s.String()
}

func (s *ResumeBatchResponse) SetHeaders(v map[string]*string) *ResumeBatchResponse {
	s.Headers = v
	return s
}

func (s *ResumeBatchResponse) SetStatusCode(v int32) *ResumeBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeBatchResponse) SetBody(v *ResumeBatchResponseBody) *ResumeBatchResponse {
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

type ResumeTriggerRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ResumeTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeTriggerRequest) GoString() string {
	return s.String()
}

func (s *ResumeTriggerRequest) SetId(v string) *ResumeTriggerRequest {
	s.Id = &v
	return s
}

func (s *ResumeTriggerRequest) SetProjectName(v string) *ResumeTriggerRequest {
	s.ProjectName = &v
	return s
}

type ResumeTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeTriggerResponseBody) SetRequestId(v string) *ResumeTriggerResponseBody {
	s.RequestId = &v
	return s
}

type ResumeTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeTriggerResponse) GoString() string {
	return s.String()
}

func (s *ResumeTriggerResponse) SetHeaders(v map[string]*string) *ResumeTriggerResponse {
	s.Headers = v
	return s
}

func (s *ResumeTriggerResponse) SetStatusCode(v int32) *ResumeTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTriggerResponse) SetBody(v *ResumeTriggerResponseBody) *ResumeTriggerResponse {
	s.Body = v
	return s
}

type SearchImageFigureClusterRequest struct {
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	DatasetName      *string           `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName      *string           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI        *string           `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s SearchImageFigureClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterRequest) SetCredentialConfig(v *CredentialConfig) *SearchImageFigureClusterRequest {
	s.CredentialConfig = v
	return s
}

func (s *SearchImageFigureClusterRequest) SetDatasetName(v string) *SearchImageFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchImageFigureClusterRequest) SetProjectName(v string) *SearchImageFigureClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *SearchImageFigureClusterRequest) SetSourceURI(v string) *SearchImageFigureClusterRequest {
	s.SourceURI = &v
	return s
}

type SearchImageFigureClusterShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	DatasetName            *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ProjectName            *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceURI              *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s SearchImageFigureClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageFigureClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterShrinkRequest) SetCredentialConfigShrink(v string) *SearchImageFigureClusterShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) SetDatasetName(v string) *SearchImageFigureClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) SetProjectName(v string) *SearchImageFigureClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) SetSourceURI(v string) *SearchImageFigureClusterShrinkRequest {
	s.SourceURI = &v
	return s
}

type SearchImageFigureClusterResponseBody struct {
	Clusters  []*SearchImageFigureClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchImageFigureClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterResponseBody) SetClusters(v []*SearchImageFigureClusterResponseBodyClusters) *SearchImageFigureClusterResponseBody {
	s.Clusters = v
	return s
}

func (s *SearchImageFigureClusterResponseBody) SetRequestId(v string) *SearchImageFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

type SearchImageFigureClusterResponseBodyClusters struct {
	Boundary   *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	ClusterId  *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Similarity *float32  `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s SearchImageFigureClusterResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s SearchImageFigureClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterResponseBodyClusters) SetBoundary(v *Boundary) *SearchImageFigureClusterResponseBodyClusters {
	s.Boundary = v
	return s
}

func (s *SearchImageFigureClusterResponseBodyClusters) SetClusterId(v string) *SearchImageFigureClusterResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *SearchImageFigureClusterResponseBodyClusters) SetSimilarity(v float32) *SearchImageFigureClusterResponseBodyClusters {
	s.Similarity = &v
	return s
}

type SearchImageFigureClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchImageFigureClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchImageFigureClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageFigureClusterResponse) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterResponse) SetHeaders(v map[string]*string) *SearchImageFigureClusterResponse {
	s.Headers = v
	return s
}

func (s *SearchImageFigureClusterResponse) SetStatusCode(v int32) *SearchImageFigureClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageFigureClusterResponse) SetBody(v *SearchImageFigureClusterResponseBody) *SearchImageFigureClusterResponse {
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
	Files     []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SemanticQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SemanticQueryResponseBody) GoString() string {
	return s.String()
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

type SuspendBatchRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s SuspendBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendBatchRequest) GoString() string {
	return s.String()
}

func (s *SuspendBatchRequest) SetId(v string) *SuspendBatchRequest {
	s.Id = &v
	return s
}

func (s *SuspendBatchRequest) SetProjectName(v string) *SuspendBatchRequest {
	s.ProjectName = &v
	return s
}

type SuspendBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendBatchResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendBatchResponseBody) SetRequestId(v string) *SuspendBatchResponseBody {
	s.RequestId = &v
	return s
}

type SuspendBatchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendBatchResponse) GoString() string {
	return s.String()
}

func (s *SuspendBatchResponse) SetHeaders(v map[string]*string) *SuspendBatchResponse {
	s.Headers = v
	return s
}

func (s *SuspendBatchResponse) SetStatusCode(v int32) *SuspendBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendBatchResponse) SetBody(v *SuspendBatchResponseBody) *SuspendBatchResponse {
	s.Body = v
	return s
}

type SuspendTriggerRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s SuspendTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendTriggerRequest) GoString() string {
	return s.String()
}

func (s *SuspendTriggerRequest) SetId(v string) *SuspendTriggerRequest {
	s.Id = &v
	return s
}

func (s *SuspendTriggerRequest) SetProjectName(v string) *SuspendTriggerRequest {
	s.ProjectName = &v
	return s
}

type SuspendTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendTriggerResponseBody) SetRequestId(v string) *SuspendTriggerResponseBody {
	s.RequestId = &v
	return s
}

type SuspendTriggerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendTriggerResponse) GoString() string {
	return s.String()
}

func (s *SuspendTriggerResponse) SetHeaders(v map[string]*string) *SuspendTriggerResponse {
	s.Headers = v
	return s
}

func (s *SuspendTriggerResponse) SetStatusCode(v int32) *SuspendTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendTriggerResponse) SetBody(v *SuspendTriggerResponseBody) *SuspendTriggerResponse {
	s.Body = v
	return s
}

type UpdateBatchRequest struct {
	Actions      []*UpdateBatchRequestActions    `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Id           *string                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Input        *Input                          `json:"Input,omitempty" xml:"Input,omitempty"`
	Notification *UpdateBatchRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	ProjectName  *string                         `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags         map[string]interface{}          `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchRequest) SetActions(v []*UpdateBatchRequestActions) *UpdateBatchRequest {
	s.Actions = v
	return s
}

func (s *UpdateBatchRequest) SetId(v string) *UpdateBatchRequest {
	s.Id = &v
	return s
}

func (s *UpdateBatchRequest) SetInput(v *Input) *UpdateBatchRequest {
	s.Input = v
	return s
}

func (s *UpdateBatchRequest) SetNotification(v *UpdateBatchRequestNotification) *UpdateBatchRequest {
	s.Notification = v
	return s
}

func (s *UpdateBatchRequest) SetProjectName(v string) *UpdateBatchRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateBatchRequest) SetTags(v map[string]interface{}) *UpdateBatchRequest {
	s.Tags = v
	return s
}

type UpdateBatchRequestActions struct {
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateBatchRequestActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchRequestActions) GoString() string {
	return s.String()
}

func (s *UpdateBatchRequestActions) SetName(v string) *UpdateBatchRequestActions {
	s.Name = &v
	return s
}

func (s *UpdateBatchRequestActions) SetParameters(v []*string) *UpdateBatchRequestActions {
	s.Parameters = v
	return s
}

type UpdateBatchRequestNotification struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Topic    *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateBatchRequestNotification) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchRequestNotification) GoString() string {
	return s.String()
}

func (s *UpdateBatchRequestNotification) SetEndpoint(v string) *UpdateBatchRequestNotification {
	s.Endpoint = &v
	return s
}

func (s *UpdateBatchRequestNotification) SetTopic(v string) *UpdateBatchRequestNotification {
	s.Topic = &v
	return s
}

type UpdateBatchShrinkRequest struct {
	ActionsShrink      *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InputShrink        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateBatchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchShrinkRequest) SetActionsShrink(v string) *UpdateBatchShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetId(v string) *UpdateBatchShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetInputShrink(v string) *UpdateBatchShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetNotificationShrink(v string) *UpdateBatchShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetProjectName(v string) *UpdateBatchShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateBatchShrinkRequest) SetTagsShrink(v string) *UpdateBatchShrinkRequest {
	s.TagsShrink = &v
	return s
}

type UpdateBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBatchResponseBody) SetRequestId(v string) *UpdateBatchResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBatchResponse) GoString() string {
	return s.String()
}

func (s *UpdateBatchResponse) SetHeaders(v map[string]*string) *UpdateBatchResponse {
	s.Headers = v
	return s
}

func (s *UpdateBatchResponse) SetStatusCode(v int32) *UpdateBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBatchResponse) SetBody(v *UpdateBatchResponseBody) *UpdateBatchResponse {
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
	DatasetName *string    `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	File        *InputFile `json:"File,omitempty" xml:"File,omitempty"`
	ProjectName *string    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *UpdateFileMetaRequest) SetFile(v *InputFile) *UpdateFileMetaRequest {
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

type UpdateLocationDateClusterRequest struct {
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName  *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName  *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Title        *string                `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLocationDateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLocationDateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterRequest) SetCustomId(v string) *UpdateLocationDateClusterRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetCustomLabels(v map[string]interface{}) *UpdateLocationDateClusterRequest {
	s.CustomLabels = v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetDatasetName(v string) *UpdateLocationDateClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetObjectId(v string) *UpdateLocationDateClusterRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetProjectName(v string) *UpdateLocationDateClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetTitle(v string) *UpdateLocationDateClusterRequest {
	s.Title = &v
	return s
}

type UpdateLocationDateClusterShrinkRequest struct {
	CustomId           *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName        *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ObjectId           *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Title              *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLocationDateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLocationDateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterShrinkRequest) SetCustomId(v string) *UpdateLocationDateClusterShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetCustomLabelsShrink(v string) *UpdateLocationDateClusterShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetDatasetName(v string) *UpdateLocationDateClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetObjectId(v string) *UpdateLocationDateClusterShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetProjectName(v string) *UpdateLocationDateClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetTitle(v string) *UpdateLocationDateClusterShrinkRequest {
	s.Title = &v
	return s
}

type UpdateLocationDateClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLocationDateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLocationDateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterResponseBody) SetRequestId(v string) *UpdateLocationDateClusterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLocationDateClusterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLocationDateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLocationDateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLocationDateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterResponse) SetHeaders(v map[string]*string) *UpdateLocationDateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateLocationDateClusterResponse) SetStatusCode(v int32) *UpdateLocationDateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLocationDateClusterResponse) SetBody(v *UpdateLocationDateClusterResponseBody) *UpdateLocationDateClusterResponse {
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
	ProjectMaxDatasetCount  *int64  `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s *UpdateProjectRequest) SetProjectMaxDatasetCount(v int64) *UpdateProjectRequest {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
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

type UpdateTriggerRequest struct {
	Actions      []*UpdateTriggerRequestActions    `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Id           *string                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Input        *Input                            `json:"Input,omitempty" xml:"Input,omitempty"`
	Notification *UpdateTriggerRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	ProjectName  *string                           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Tags         map[string]interface{}            `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) SetActions(v []*UpdateTriggerRequestActions) *UpdateTriggerRequest {
	s.Actions = v
	return s
}

func (s *UpdateTriggerRequest) SetId(v string) *UpdateTriggerRequest {
	s.Id = &v
	return s
}

func (s *UpdateTriggerRequest) SetInput(v *Input) *UpdateTriggerRequest {
	s.Input = v
	return s
}

func (s *UpdateTriggerRequest) SetNotification(v *UpdateTriggerRequestNotification) *UpdateTriggerRequest {
	s.Notification = v
	return s
}

func (s *UpdateTriggerRequest) SetProjectName(v string) *UpdateTriggerRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateTriggerRequest) SetTags(v map[string]interface{}) *UpdateTriggerRequest {
	s.Tags = v
	return s
}

type UpdateTriggerRequestActions struct {
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s UpdateTriggerRequestActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequestActions) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequestActions) SetName(v string) *UpdateTriggerRequestActions {
	s.Name = &v
	return s
}

func (s *UpdateTriggerRequestActions) SetParameters(v []*string) *UpdateTriggerRequestActions {
	s.Parameters = v
	return s
}

type UpdateTriggerRequestNotification struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Topic    *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateTriggerRequestNotification) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequestNotification) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequestNotification) SetEndpoint(v string) *UpdateTriggerRequestNotification {
	s.Endpoint = &v
	return s
}

func (s *UpdateTriggerRequestNotification) SetTopic(v string) *UpdateTriggerRequestNotification {
	s.Topic = &v
	return s
}

type UpdateTriggerShrinkRequest struct {
	ActionsShrink      *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InputShrink        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateTriggerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerShrinkRequest) SetActionsShrink(v string) *UpdateTriggerShrinkRequest {
	s.ActionsShrink = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetId(v string) *UpdateTriggerShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetInputShrink(v string) *UpdateTriggerShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetNotificationShrink(v string) *UpdateTriggerShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetProjectName(v string) *UpdateTriggerShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateTriggerShrinkRequest) SetTagsShrink(v string) *UpdateTriggerShrinkRequest {
	s.TagsShrink = &v
	return s
}

type UpdateTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponseBody) SetRequestId(v string) *UpdateTriggerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponse) SetHeaders(v map[string]*string) *UpdateTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTriggerResponse) SetStatusCode(v int32) *UpdateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTriggerResponse) SetBody(v *UpdateTriggerResponseBody) *UpdateTriggerResponse {
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

func (client *Client) AddImageMosaicWithOptions(tmpReq *AddImageMosaicRequest, runtime *util.RuntimeOptions) (_result *AddImageMosaicResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddImageMosaicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Targets)) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, tea.String("Targets"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageFormat)) {
		query["ImageFormat"] = request.ImageFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Quality)) {
		query["Quality"] = request.Quality
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TargetURI)) {
		query["TargetURI"] = request.TargetURI
	}

	if !tea.BoolValue(util.IsUnset(request.TargetsShrink)) {
		query["Targets"] = request.TargetsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImageMosaic"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageMosaicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImageMosaic(request *AddImageMosaicRequest) (_result *AddImageMosaicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImageMosaicResponse{}
	_body, _err := client.AddImageMosaicWithOptions(request, runtime)
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

func (client *Client) BatchGetFigureClusterWithOptions(tmpReq *BatchGetFigureClusterRequest, runtime *util.RuntimeOptions) (_result *BatchGetFigureClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchGetFigureClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ObjectIds)) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, tea.String("ObjectIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectIdsShrink)) {
		query["ObjectIds"] = request.ObjectIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetFigureCluster"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetFigureCluster(request *BatchGetFigureClusterRequest) (_result *BatchGetFigureClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchGetFigureClusterResponse{}
	_body, _err := client.BatchGetFigureClusterWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FilesShrink)) {
		query["Files"] = request.FilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

func (client *Client) CompareImageFacesWithOptions(tmpReq *CompareImageFacesRequest, runtime *util.RuntimeOptions) (_result *CompareImageFacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CompareImageFacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Source)) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, tea.String("Source"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceShrink)) {
		query["Source"] = request.SourceShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareImageFaces"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareImageFaces(request *CompareImageFacesRequest) (_result *CompareImageFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareImageFacesResponse{}
	_body, _err := client.CompareImageFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateArchiveFileInspectionTaskWithOptions(tmpReq *CreateArchiveFileInspectionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateArchiveFileInspectionTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateArchiveFileInspectionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
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
		Action:      tea.String("CreateArchiveFileInspectionTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArchiveFileInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateArchiveFileInspectionTask(request *CreateArchiveFileInspectionTaskRequest) (_result *CreateArchiveFileInspectionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArchiveFileInspectionTaskResponse{}
	_body, _err := client.CreateArchiveFileInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBatchWithOptions(tmpReq *CreateBatchRequest, runtime *util.RuntimeOptions) (_result *CreateBatchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Actions)) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, tea.String("Actions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionsShrink)) {
		body["Actions"] = request.ActionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		body["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		body["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRole)) {
		body["ServiceRole"] = request.ServiceRole
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBatch"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBatch(request *CreateBatchRequest) (_result *CreateBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBatchResponse{}
	_body, _err := client.CreateBatchWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.KdtreeOption)) {
		request.KdtreeOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KdtreeOption, tea.String("KdtreeOption"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OctreeOption)) {
		request.OctreeOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OctreeOption, tea.String("OctreeOption"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

func (client *Client) CreateCustomizedStoryWithOptions(tmpReq *CreateCustomizedStoryRequest, runtime *util.RuntimeOptions) (_result *CreateCustomizedStoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCustomizedStoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Cover)) {
		request.CoverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cover, tea.String("Cover"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Files)) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, tea.String("Files"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverShrink)) {
		body["Cover"] = request.CoverShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabelsShrink)) {
		body["CustomLabels"] = request.CustomLabelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		body["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FilesShrink)) {
		body["Files"] = request.FilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["StoryName"] = request.StoryName
	}

	if !tea.BoolValue(util.IsUnset(request.StorySubType)) {
		body["StorySubType"] = request.StorySubType
	}

	if !tea.BoolValue(util.IsUnset(request.StoryType)) {
		body["StoryType"] = request.StoryType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomizedStory"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomizedStory(request *CreateCustomizedStoryRequest) (_result *CreateCustomizedStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.CreateCustomizedStoryWithOptions(request, runtime)
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

func (client *Client) CreateFacesSearchingTaskWithOptions(tmpReq *CreateFacesSearchingTaskRequest, runtime *util.RuntimeOptions) (_result *CreateFacesSearchingTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFacesSearchingTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		query["TopK"] = request.TopK
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFacesSearchingTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFacesSearchingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFacesSearchingTask(request *CreateFacesSearchingTaskRequest) (_result *CreateFacesSearchingTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFacesSearchingTaskResponse{}
	_body, _err := client.CreateFacesSearchingTaskWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
	if !tea.BoolValue(util.IsUnset(tmpReq.Froms)) {
		request.FromsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Froms, tea.String("Froms"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

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

	if !tea.BoolValue(util.IsUnset(request.FromsShrink)) {
		query["Froms"] = request.FromsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

func (client *Client) CreateFileCompressionTaskWithOptions(tmpReq *CreateFileCompressionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateFileCompressionTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFileCompressionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompressedFormat)) {
		query["CompressedFormat"] = request.CompressedFormat
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceManifestURI)) {
		query["SourceManifestURI"] = request.SourceManifestURI
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
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
		Action:      tea.String("CreateFileCompressionTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileCompressionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFileCompressionTask(request *CreateFileCompressionTaskRequest) (_result *CreateFileCompressionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileCompressionTaskResponse{}
	_body, _err := client.CreateFileCompressionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileUncompressionTaskWithOptions(tmpReq *CreateFileUncompressionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateFileUncompressionTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFileUncompressionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SelectedFiles)) {
		request.SelectedFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedFiles, tea.String("SelectedFiles"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Target)) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, tea.String("Target"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SelectedFilesShrink)) {
		query["SelectedFiles"] = request.SelectedFilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TargetShrink)) {
		query["Target"] = request.TargetShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileUncompressionTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileUncompressionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFileUncompressionTask(request *CreateFileUncompressionTaskRequest) (_result *CreateFileUncompressionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileUncompressionTaskResponse{}
	_body, _err := client.CreateFileUncompressionTaskWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

func (client *Client) CreateImageToPDFTaskWithOptions(tmpReq *CreateImageToPDFTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageToPDFTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateImageToPDFTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
		Action:      tea.String("CreateImageToPDFTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageToPDFTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageToPDFTask(request *CreateImageToPDFTaskRequest) (_result *CreateImageToPDFTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageToPDFTaskResponse{}
	_body, _err := client.CreateImageToPDFTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLocationDateClusteringTaskWithOptions(tmpReq *CreateLocationDateClusteringTaskRequest, runtime *util.RuntimeOptions) (_result *CreateLocationDateClusteringTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateLocationDateClusteringTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DateOptions)) {
		request.DateOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DateOptions, tea.String("DateOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LocationOptions)) {
		request.LocationOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationOptions, tea.String("LocationOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.DateOptionsShrink)) {
		query["DateOptions"] = request.DateOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LocationOptionsShrink)) {
		query["LocationOptions"] = request.LocationOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
		Action:      tea.String("CreateLocationDateClusteringTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLocationDateClusteringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLocationDateClusteringTask(request *CreateLocationDateClusteringTaskRequest) (_result *CreateLocationDateClusteringTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLocationDateClusteringTaskResponse{}
	_body, _err := client.CreateLocationDateClusteringTaskWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TrimPolicy)) {
		request.TrimPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrimPolicy, tea.String("TrimPolicy"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

	if !tea.BoolValue(util.IsUnset(request.ProjectMaxDatasetCount)) {
		query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
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

func (client *Client) CreateSimilarImageClusteringTaskWithOptions(tmpReq *CreateSimilarImageClusteringTaskRequest, runtime *util.RuntimeOptions) (_result *CreateSimilarImageClusteringTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSimilarImageClusteringTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
		Action:      tea.String("CreateSimilarImageClusteringTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSimilarImageClusteringTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSimilarImageClusteringTask(request *CreateSimilarImageClusteringTaskRequest) (_result *CreateSimilarImageClusteringTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSimilarImageClusteringTaskResponse{}
	_body, _err := client.CreateSimilarImageClusteringTaskWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.Address)) {
		request.AddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Address, tea.String("Address"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressShrink)) {
		body["Address"] = request.AddressShrink
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

func (client *Client) CreateTriggerWithOptions(tmpReq *CreateTriggerRequest, runtime *util.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTriggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Actions)) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, tea.String("Actions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionsShrink)) {
		body["Actions"] = request.ActionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		body["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		body["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRole)) {
		body["ServiceRole"] = request.ServiceRole
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrigger"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrigger(request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoLabelClassificationTaskWithOptions(tmpReq *CreateVideoLabelClassificationTaskRequest, runtime *util.RuntimeOptions) (_result *CreateVideoLabelClassificationTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateVideoLabelClassificationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
		Action:      tea.String("CreateVideoLabelClassificationTask"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVideoLabelClassificationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoLabelClassificationTask(request *CreateVideoLabelClassificationTaskRequest) (_result *CreateVideoLabelClassificationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVideoLabelClassificationTaskResponse{}
	_body, _err := client.CreateVideoLabelClassificationTaskWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

func (client *Client) DeleteBatchWithOptions(request *DeleteBatchRequest, runtime *util.RuntimeOptions) (_result *DeleteBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBatch"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBatch(request *DeleteBatchRequest) (_result *DeleteBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBatchResponse{}
	_body, _err := client.DeleteBatchWithOptions(request, runtime)
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

func (client *Client) DeleteLocationDateClusterWithOptions(request *DeleteLocationDateClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteLocationDateClusterResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["ObjectId"] = request.ObjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLocationDateCluster"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLocationDateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLocationDateCluster(request *DeleteLocationDateClusterRequest) (_result *DeleteLocationDateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLocationDateClusterResponse{}
	_body, _err := client.DeleteLocationDateClusterWithOptions(request, runtime)
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

func (client *Client) DeleteTriggerWithOptions(request *DeleteTriggerRequest, runtime *util.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrigger"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrigger(request *DeleteTriggerRequest) (_result *DeleteTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DeleteTriggerWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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

func (client *Client) DetectImageCarsWithOptions(tmpReq *DetectImageCarsRequest, runtime *util.RuntimeOptions) (_result *DetectImageCarsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectImageCarsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
		Action:      tea.String("DetectImageCars"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageCarsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageCars(request *DetectImageCarsRequest) (_result *DetectImageCarsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageCarsResponse{}
	_body, _err := client.DetectImageCarsWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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

func (client *Client) DetectMediaMetaWithOptions(tmpReq *DetectMediaMetaRequest, runtime *util.RuntimeOptions) (_result *DetectMediaMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectMediaMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
		Action:      tea.String("DetectMediaMeta"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectMediaMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectMediaMeta(request *DetectMediaMetaRequest) (_result *DetectMediaMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectMediaMetaResponse{}
	_body, _err := client.DetectMediaMetaWithOptions(request, runtime)
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

func (client *Client) ExtractDocumentTextWithOptions(tmpReq *ExtractDocumentTextRequest, runtime *util.RuntimeOptions) (_result *ExtractDocumentTextResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExtractDocumentTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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
		Action:      tea.String("ExtractDocumentText"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractDocumentTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractDocumentText(request *ExtractDocumentTextRequest) (_result *ExtractDocumentTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractDocumentTextResponse{}
	_body, _err := client.ExtractDocumentTextWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
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

func (client *Client) GenerateDRMLicenseWithOptions(request *GenerateDRMLicenseRequest, runtime *util.RuntimeOptions) (_result *GenerateDRMLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectionSystem)) {
		query["ProtectionSystem"] = request.ProtectionSystem
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDRMLicense"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDRMLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDRMLicense(request *GenerateDRMLicenseRequest) (_result *GenerateDRMLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDRMLicenseResponse{}
	_body, _err := client.GenerateDRMLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateVideoPlaylistWithOptions(tmpReq *GenerateVideoPlaylistRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoPlaylistResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateVideoPlaylistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceSubtitles)) {
		request.SourceSubtitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceSubtitles, tea.String("SourceSubtitles"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.MasterURI)) {
		query["MasterURI"] = request.MasterURI
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDuration)) {
		query["SourceDuration"] = request.SourceDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SourceStartTime)) {
		query["SourceStartTime"] = request.SourceStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SourceSubtitlesShrink)) {
		query["SourceSubtitles"] = request.SourceSubtitlesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceURI)) {
		query["SourceURI"] = request.SourceURI
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TargetsShrink)) {
		query["Targets"] = request.TargetsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateVideoPlaylist"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateVideoPlaylistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVideoPlaylist(request *GenerateVideoPlaylistRequest) (_result *GenerateVideoPlaylistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVideoPlaylistResponse{}
	_body, _err := client.GenerateVideoPlaylistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateWebofficeTokenWithOptions(tmpReq *GenerateWebofficeTokenRequest, runtime *util.RuntimeOptions) (_result *GenerateWebofficeTokenResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateWebofficeTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Permission)) {
		request.PermissionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permission, tea.String("Permission"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.User)) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, tea.String("User"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Watermark)) {
		request.WatermarkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Watermark, tea.String("Watermark"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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
		Action:      tea.String("GenerateWebofficeToken"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateWebofficeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateWebofficeToken(request *GenerateWebofficeTokenRequest) (_result *GenerateWebofficeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateWebofficeTokenResponse{}
	_body, _err := client.GenerateWebofficeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBatchWithOptions(request *GetBatchRequest, runtime *util.RuntimeOptions) (_result *GetBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBatch"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBatch(request *GetBatchRequest) (_result *GetBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBatchResponse{}
	_body, _err := client.GetBatchWithOptions(request, runtime)
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

/**
 * @deprecated
 *
 * @param request GetDRMLicenseRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDRMLicenseResponse
 */
// Deprecated
func (client *Client) GetDRMLicenseWithOptions(request *GetDRMLicenseRequest, runtime *util.RuntimeOptions) (_result *GetDRMLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyEndpoint)) {
		query["NotifyEndpoint"] = request.NotifyEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTopicName)) {
		query["NotifyTopicName"] = request.NotifyTopicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectionSystem)) {
		query["ProtectionSystem"] = request.ProtectionSystem
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDRMLicense"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDRMLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetDRMLicenseRequest
 * @return GetDRMLicenseResponse
 */
// Deprecated
func (client *Client) GetDRMLicense(request *GetDRMLicenseRequest) (_result *GetDRMLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDRMLicenseResponse{}
	_body, _err := client.GetDRMLicenseWithOptions(request, runtime)
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

func (client *Client) GetImageModerationResultWithOptions(request *GetImageModerationResultRequest, runtime *util.RuntimeOptions) (_result *GetImageModerationResultResponse, _err error) {
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
		Action:      tea.String("GetImageModerationResult"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageModerationResult(request *GetImageModerationResultRequest) (_result *GetImageModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageModerationResultResponse{}
	_body, _err := client.GetImageModerationResultWithOptions(request, runtime)
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

func (client *Client) GetTriggerWithOptions(request *GetTriggerRequest, runtime *util.RuntimeOptions) (_result *GetTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrigger"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrigger(request *GetTriggerRequest) (_result *GetTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTriggerResponse{}
	_body, _err := client.GetTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoLabelClassificationResultWithOptions(request *GetVideoLabelClassificationResultRequest, runtime *util.RuntimeOptions) (_result *GetVideoLabelClassificationResultResponse, _err error) {
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
		Action:      tea.String("GetVideoLabelClassificationResult"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoLabelClassificationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoLabelClassificationResult(request *GetVideoLabelClassificationResultRequest) (_result *GetVideoLabelClassificationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoLabelClassificationResultResponse{}
	_body, _err := client.GetVideoLabelClassificationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoModerationResultWithOptions(request *GetVideoModerationResultRequest, runtime *util.RuntimeOptions) (_result *GetVideoModerationResultResponse, _err error) {
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
		Action:      tea.String("GetVideoModerationResult"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoModerationResult(request *GetVideoModerationResultRequest) (_result *GetVideoModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoModerationResultResponse{}
	_body, _err := client.GetVideoModerationResultWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.File)) {
		request.FileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.File, tea.String("File"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.FileShrink)) {
		query["File"] = request.FileShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		query["Notification"] = request.NotificationShrink
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

func (client *Client) ListBatchesWithOptions(request *ListBatchesRequest, runtime *util.RuntimeOptions) (_result *ListBatchesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.TagSelector)) {
		query["TagSelector"] = request.TagSelector
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBatches"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBatches(request *ListBatchesRequest) (_result *ListBatchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBatchesResponse{}
	_body, _err := client.ListBatchesWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.EndTimeRange)) {
		request.EndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EndTimeRange, tea.String("EndTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StartTimeRange)) {
		request.StartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartTimeRange, tea.String("StartTimeRange"), tea.String("json"))
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

func (client *Client) ListTriggersWithOptions(request *ListTriggersRequest, runtime *util.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.TagSelector)) {
		query["TagSelector"] = request.TagSelector
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTriggers"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTriggers(request *ListTriggersRequest) (_result *ListTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTriggersResponse{}
	_body, _err := client.ListTriggersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LiveTranscodingWithOptions(tmpReq *LiveTranscodingRequest, runtime *util.RuntimeOptions) (_result *LiveTranscodingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &LiveTranscodingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LiveTranscoding"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LiveTranscodingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LiveTranscoding(request *LiveTranscodingRequest) (_result *LiveTranscodingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LiveTranscodingResponse{}
	_body, _err := client.LiveTranscodingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFigureClustersWithOptions(tmpReq *QueryFigureClustersRequest, runtime *util.RuntimeOptions) (_result *QueryFigureClustersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryFigureClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateTimeRange)) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimeRange, tea.String("CreateTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateTimeRange)) {
		request.UpdateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateTimeRange, tea.String("UpdateTimeRange"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeRangeShrink)) {
		query["UpdateTimeRange"] = request.UpdateTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WithTotalCount)) {
		query["WithTotalCount"] = request.WithTotalCount
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

func (client *Client) QueryLocationDateClustersWithOptions(tmpReq *QueryLocationDateClustersRequest, runtime *util.RuntimeOptions) (_result *QueryLocationDateClustersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryLocationDateClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Address)) {
		request.AddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Address, tea.String("Address"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CreateTimeRange)) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimeRange, tea.String("CreateTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LocationDateClusterEndTimeRange)) {
		request.LocationDateClusterEndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDateClusterEndTimeRange, tea.String("LocationDateClusterEndTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LocationDateClusterLevels)) {
		request.LocationDateClusterLevelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDateClusterLevels, tea.String("LocationDateClusterLevels"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LocationDateClusterStartTimeRange)) {
		request.LocationDateClusterStartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LocationDateClusterStartTimeRange, tea.String("LocationDateClusterStartTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateTimeRange)) {
		request.UpdateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateTimeRange, tea.String("UpdateTimeRange"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressShrink)) {
		query["Address"] = request.AddressShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeRangeShrink)) {
		query["CreateTimeRange"] = request.CreateTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		query["CustomLabels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.LocationDateClusterEndTimeRangeShrink)) {
		query["LocationDateClusterEndTimeRange"] = request.LocationDateClusterEndTimeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LocationDateClusterLevelsShrink)) {
		query["LocationDateClusterLevels"] = request.LocationDateClusterLevelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LocationDateClusterStartTimeRangeShrink)) {
		query["LocationDateClusterStartTimeRange"] = request.LocationDateClusterStartTimeRangeShrink
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

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeRangeShrink)) {
		query["UpdateTimeRange"] = request.UpdateTimeRangeShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryLocationDateClusters"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLocationDateClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLocationDateClusters(request *QueryLocationDateClustersRequest) (_result *QueryLocationDateClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLocationDateClustersResponse{}
	_body, _err := client.QueryLocationDateClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySimilarImageClustersWithOptions(request *QuerySimilarImageClustersRequest, runtime *util.RuntimeOptions) (_result *QuerySimilarImageClustersResponse, _err error) {
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
		Action:      tea.String("QuerySimilarImageClusters"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySimilarImageClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySimilarImageClusters(request *QuerySimilarImageClustersRequest) (_result *QuerySimilarImageClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySimilarImageClustersResponse{}
	_body, _err := client.QuerySimilarImageClustersWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateTimeRange)) {
		request.CreateTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimeRange, tea.String("CreateTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.FigureClusterIds)) {
		request.FigureClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FigureClusterIds, tea.String("FigureClusterIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StoryEndTimeRange)) {
		request.StoryEndTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoryEndTimeRange, tea.String("StoryEndTimeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StoryStartTimeRange)) {
		request.StoryStartTimeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoryStartTimeRange, tea.String("StoryStartTimeRange"), tea.String("json"))
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
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
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

func (client *Client) ResumeBatchWithOptions(request *ResumeBatchRequest, runtime *util.RuntimeOptions) (_result *ResumeBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeBatch"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeBatch(request *ResumeBatchRequest) (_result *ResumeBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeBatchResponse{}
	_body, _err := client.ResumeBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ResumeBindingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResumeBindingResponse
 */
// Deprecated
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

/**
 * @deprecated
 *
 * @param request ResumeBindingRequest
 * @return ResumeBindingResponse
 */
// Deprecated
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

func (client *Client) ResumeTriggerWithOptions(request *ResumeTriggerRequest, runtime *util.RuntimeOptions) (_result *ResumeTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeTrigger"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeTrigger(request *ResumeTriggerRequest) (_result *ResumeTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeTriggerResponse{}
	_body, _err := client.ResumeTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageFigureClusterWithOptions(tmpReq *SearchImageFigureClusterRequest, runtime *util.RuntimeOptions) (_result *SearchImageFigureClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchImageFigureClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CredentialConfig)) {
		request.CredentialConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CredentialConfig, tea.String("CredentialConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialConfigShrink)) {
		query["CredentialConfig"] = request.CredentialConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
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
		Action:      tea.String("SearchImageFigureCluster"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchImageFigureClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchImageFigureCluster(request *SearchImageFigureClusterRequest) (_result *SearchImageFigureClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImageFigureClusterResponse{}
	_body, _err := client.SearchImageFigureClusterWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(tmpReq.Query)) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, tea.String("Query"), tea.String("json"))
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

/**
 * @deprecated
 *
 * @param request StopBindingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopBindingResponse
 */
// Deprecated
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

/**
 * @deprecated
 *
 * @param request StopBindingRequest
 * @return StopBindingResponse
 */
// Deprecated
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

func (client *Client) SuspendBatchWithOptions(request *SuspendBatchRequest, runtime *util.RuntimeOptions) (_result *SuspendBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendBatch"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendBatch(request *SuspendBatchRequest) (_result *SuspendBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendBatchResponse{}
	_body, _err := client.SuspendBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendTriggerWithOptions(request *SuspendTriggerRequest, runtime *util.RuntimeOptions) (_result *SuspendTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendTrigger"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendTrigger(request *SuspendTriggerRequest) (_result *SuspendTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendTriggerResponse{}
	_body, _err := client.SuspendTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBatchWithOptions(tmpReq *UpdateBatchRequest, runtime *util.RuntimeOptions) (_result *UpdateBatchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Actions)) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, tea.String("Actions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionsShrink)) {
		body["Actions"] = request.ActionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		body["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		body["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBatch"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBatch(request *UpdateBatchRequest) (_result *UpdateBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBatchResponse{}
	_body, _err := client.UpdateBatchWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.FigureCluster)) {
		request.FigureClusterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FigureCluster, tea.String("FigureCluster"), tea.String("json"))
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
	if !tea.BoolValue(util.IsUnset(tmpReq.File)) {
		request.FileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.File, tea.String("File"), tea.String("json"))
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

func (client *Client) UpdateLocationDateClusterWithOptions(tmpReq *UpdateLocationDateClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateLocationDateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLocationDateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomLabels)) {
		request.CustomLabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomLabels, tea.String("CustomLabels"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomId)) {
		query["CustomId"] = request.CustomId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabelsShrink)) {
		query["CustomLabels"] = request.CustomLabelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetName)) {
		query["DatasetName"] = request.DatasetName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLocationDateCluster"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLocationDateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLocationDateCluster(request *UpdateLocationDateClusterRequest) (_result *UpdateLocationDateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLocationDateClusterResponse{}
	_body, _err := client.UpdateLocationDateClusterWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ProjectMaxDatasetCount)) {
		query["ProjectMaxDatasetCount"] = request.ProjectMaxDatasetCount
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
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
	if !tea.BoolValue(util.IsUnset(tmpReq.Cover)) {
		request.CoverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cover, tea.String("Cover"), tea.String("json"))
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

func (client *Client) UpdateTriggerWithOptions(tmpReq *UpdateTriggerRequest, runtime *util.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTriggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Actions)) {
		request.ActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Actions, tea.String("Actions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Notification)) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, tea.String("Notification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionsShrink)) {
		body["Actions"] = request.ActionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		body["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationShrink)) {
		body["Notification"] = request.NotificationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrigger"),
		Version:     tea.String("2020-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrigger(request *UpdateTriggerRequest) (_result *UpdateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.UpdateTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
