// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddLiveASRConfigRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName      *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	MnsTopic        *string `json:"MnsTopic,omitempty" xml:"MnsTopic,omitempty"`
	MnsRegion       *string `json:"MnsRegion,omitempty" xml:"MnsRegion,omitempty"`
	Period          *int    `json:"Period,omitempty" xml:"Period,omitempty"`
	HttpCallbackURL *string `json:"HttpCallbackURL,omitempty" xml:"HttpCallbackURL,omitempty"`
}

func (s AddLiveASRConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveASRConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveASRConfigRequest) SetDomainName(v string) *AddLiveASRConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveASRConfigRequest) SetAppName(v string) *AddLiveASRConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveASRConfigRequest) SetStreamName(v string) *AddLiveASRConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveASRConfigRequest) SetMnsTopic(v string) *AddLiveASRConfigRequest {
	s.MnsTopic = &v
	return s
}

func (s *AddLiveASRConfigRequest) SetMnsRegion(v string) *AddLiveASRConfigRequest {
	s.MnsRegion = &v
	return s
}

func (s *AddLiveASRConfigRequest) SetPeriod(v int) *AddLiveASRConfigRequest {
	s.Period = &v
	return s
}

func (s *AddLiveASRConfigRequest) SetHttpCallbackURL(v string) *AddLiveASRConfigRequest {
	s.HttpCallbackURL = &v
	return s
}

type AddLiveASRConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveASRConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveASRConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveASRConfigResponse) SetRequestId(v string) *AddLiveASRConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveAsrConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveAsrConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAsrConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAsrConfigRequest) SetDomainName(v string) *DescribeLiveAsrConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAsrConfigRequest) SetAppName(v string) *DescribeLiveAsrConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveAsrConfigRequest) SetStreamName(v string) *DescribeLiveAsrConfigRequest {
	s.StreamName = &v
	return s
}

type DescribeLiveAsrConfigResponse struct {
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveAsrConfig *DescribeLiveAsrConfigResponseLiveAsrConfig `json:"LiveAsrConfig,omitempty" xml:"LiveAsrConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveAsrConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAsrConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAsrConfigResponse) SetRequestId(v string) *DescribeLiveAsrConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAsrConfigResponse) SetLiveAsrConfig(v *DescribeLiveAsrConfigResponseLiveAsrConfig) *DescribeLiveAsrConfigResponse {
	s.LiveAsrConfig = v
	return s
}

type DescribeLiveAsrConfigResponseLiveAsrConfig struct {
	LiveAsrConfigList []*DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList `json:"LiveAsrConfigList,omitempty" xml:"LiveAsrConfigList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveAsrConfigResponseLiveAsrConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAsrConfigResponseLiveAsrConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfig) SetLiveAsrConfigList(v []*DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) *DescribeLiveAsrConfigResponseLiveAsrConfig {
	s.LiveAsrConfigList = v
	return s
}

type DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList struct {
	DomainName      *int    `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName      *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	Period          *int    `json:"Period,omitempty" xml:"Period,omitempty" require:"true"`
	MnsTopic        *string `json:"MnsTopic,omitempty" xml:"MnsTopic,omitempty" require:"true"`
	MnsRegion       *string `json:"MnsRegion,omitempty" xml:"MnsRegion,omitempty" require:"true"`
	HttpCallbackURL *string `json:"HttpCallbackURL,omitempty" xml:"HttpCallbackURL,omitempty" require:"true"`
}

func (s DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetDomainName(v int) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetAppName(v string) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.AppName = &v
	return s
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetStreamName(v string) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetPeriod(v int) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.Period = &v
	return s
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetMnsTopic(v string) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.MnsTopic = &v
	return s
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetMnsRegion(v string) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.MnsRegion = &v
	return s
}

func (s *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList) SetHttpCallbackURL(v string) *DescribeLiveAsrConfigResponseLiveAsrConfigLiveAsrConfigList {
	s.HttpCallbackURL = &v
	return s
}

type DeleteLiveASRConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
}

func (s DeleteLiveASRConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveASRConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveASRConfigRequest) SetDomainName(v string) *DeleteLiveASRConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveASRConfigRequest) SetAppName(v string) *DeleteLiveASRConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveASRConfigRequest) SetStreamName(v string) *DeleteLiveASRConfigRequest {
	s.StreamName = &v
	return s
}

type DeleteLiveASRConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveASRConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveASRConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveASRConfigResponse) SetRequestId(v string) *DeleteLiveASRConfigResponse {
	s.RequestId = &v
	return s
}

type UpdateLiveASRConfigRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName      *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	MnsTopic        *string `json:"MnsTopic,omitempty" xml:"MnsTopic,omitempty"`
	MnsRegion       *string `json:"MnsRegion,omitempty" xml:"MnsRegion,omitempty"`
	Period          *int    `json:"Period,omitempty" xml:"Period,omitempty"`
	HttpCallbackURL *string `json:"HttpCallbackURL,omitempty" xml:"HttpCallbackURL,omitempty"`
}

func (s UpdateLiveASRConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveASRConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveASRConfigRequest) SetDomainName(v string) *UpdateLiveASRConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveASRConfigRequest) SetAppName(v string) *UpdateLiveASRConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveASRConfigRequest) SetStreamName(v string) *UpdateLiveASRConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveASRConfigRequest) SetMnsTopic(v string) *UpdateLiveASRConfigRequest {
	s.MnsTopic = &v
	return s
}

func (s *UpdateLiveASRConfigRequest) SetMnsRegion(v string) *UpdateLiveASRConfigRequest {
	s.MnsRegion = &v
	return s
}

func (s *UpdateLiveASRConfigRequest) SetPeriod(v int) *UpdateLiveASRConfigRequest {
	s.Period = &v
	return s
}

func (s *UpdateLiveASRConfigRequest) SetHttpCallbackURL(v string) *UpdateLiveASRConfigRequest {
	s.HttpCallbackURL = &v
	return s
}

type UpdateLiveASRConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveASRConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveASRConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveASRConfigResponse) SetRequestId(v string) *UpdateLiveASRConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteMixStreamRequest struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName  *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty" require:"true"`
}

func (s DeleteMixStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixStreamRequest) GoString() string {
	return s.String()
}

func (s *DeleteMixStreamRequest) SetDomainName(v string) *DeleteMixStreamRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteMixStreamRequest) SetAppName(v string) *DeleteMixStreamRequest {
	s.AppName = &v
	return s
}

func (s *DeleteMixStreamRequest) SetStreamName(v string) *DeleteMixStreamRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteMixStreamRequest) SetMixStreamId(v string) *DeleteMixStreamRequest {
	s.MixStreamId = &v
	return s
}

type DeleteMixStreamResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty" require:"true"`
}

func (s DeleteMixStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteMixStreamResponse) SetRequestId(v string) *DeleteMixStreamResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteMixStreamResponse) SetMixStreamId(v string) *DeleteMixStreamResponse {
	s.MixStreamId = &v
	return s
}

type UpdateMixStreamRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	MixStreamId     *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty" require:"true"`
	InputStreamList *string `json:"InputStreamList,omitempty" xml:"InputStreamList,omitempty" require:"true"`
	LayoutId        *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
}

func (s UpdateMixStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMixStreamRequest) GoString() string {
	return s.String()
}

func (s *UpdateMixStreamRequest) SetDomainName(v string) *UpdateMixStreamRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateMixStreamRequest) SetMixStreamId(v string) *UpdateMixStreamRequest {
	s.MixStreamId = &v
	return s
}

func (s *UpdateMixStreamRequest) SetInputStreamList(v string) *UpdateMixStreamRequest {
	s.InputStreamList = &v
	return s
}

func (s *UpdateMixStreamRequest) SetLayoutId(v string) *UpdateMixStreamRequest {
	s.LayoutId = &v
	return s
}

type UpdateMixStreamResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty" require:"true"`
}

func (s UpdateMixStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMixStreamResponse) GoString() string {
	return s.String()
}

func (s *UpdateMixStreamResponse) SetRequestId(v string) *UpdateMixStreamResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateMixStreamResponse) SetMixStreamId(v string) *UpdateMixStreamResponse {
	s.MixStreamId = &v
	return s
}

type CreateMixStreamRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	LayoutId        *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
	InputStreamList *string `json:"InputStreamList,omitempty" xml:"InputStreamList,omitempty" require:"true"`
	OutputConfig    *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty" require:"true"`
	CallbackConfig  *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
}

func (s CreateMixStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMixStreamRequest) GoString() string {
	return s.String()
}

func (s *CreateMixStreamRequest) SetDomainName(v string) *CreateMixStreamRequest {
	s.DomainName = &v
	return s
}

func (s *CreateMixStreamRequest) SetLayoutId(v string) *CreateMixStreamRequest {
	s.LayoutId = &v
	return s
}

func (s *CreateMixStreamRequest) SetInputStreamList(v string) *CreateMixStreamRequest {
	s.InputStreamList = &v
	return s
}

func (s *CreateMixStreamRequest) SetOutputConfig(v string) *CreateMixStreamRequest {
	s.OutputConfig = &v
	return s
}

func (s *CreateMixStreamRequest) SetCallbackConfig(v string) *CreateMixStreamRequest {
	s.CallbackConfig = &v
	return s
}

type CreateMixStreamResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty" require:"true"`
}

func (s CreateMixStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMixStreamResponse) GoString() string {
	return s.String()
}

func (s *CreateMixStreamResponse) SetRequestId(v string) *CreateMixStreamResponse {
	s.RequestId = &v
	return s
}

func (s *CreateMixStreamResponse) SetMixStreamId(v string) *CreateMixStreamResponse {
	s.MixStreamId = &v
	return s
}

type DescribeMixStreamListRequest struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName  *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo      *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMixStreamListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixStreamListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListRequest) SetDomainName(v string) *DescribeMixStreamListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetAppName(v string) *DescribeMixStreamListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetStreamName(v string) *DescribeMixStreamListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetMixStreamId(v string) *DescribeMixStreamListRequest {
	s.MixStreamId = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetStartTime(v string) *DescribeMixStreamListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetEndTime(v string) *DescribeMixStreamListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetPageNo(v int) *DescribeMixStreamListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeMixStreamListRequest) SetPageSize(v int) *DescribeMixStreamListRequest {
	s.PageSize = &v
	return s
}

type DescribeMixStreamListResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total         *int                                          `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	MixStreamList []*DescribeMixStreamListResponseMixStreamList `json:"MixStreamList,omitempty" xml:"MixStreamList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeMixStreamListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixStreamListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListResponse) SetRequestId(v string) *DescribeMixStreamListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeMixStreamListResponse) SetTotal(v int) *DescribeMixStreamListResponse {
	s.Total = &v
	return s
}

func (s *DescribeMixStreamListResponse) SetMixStreamList(v []*DescribeMixStreamListResponseMixStreamList) *DescribeMixStreamListResponse {
	s.MixStreamList = v
	return s
}

type DescribeMixStreamListResponseMixStreamList struct {
	MixstreamId       *string `json:"MixstreamId,omitempty" xml:"MixstreamId,omitempty" require:"true"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	LayoutId          *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
	InputStreamNumber *int    `json:"InputStreamNumber,omitempty" xml:"InputStreamNumber,omitempty" require:"true"`
	MixStreamTemplate *string `json:"MixStreamTemplate,omitempty" xml:"MixStreamTemplate,omitempty" require:"true"`
	GmtCreate         *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	GmtModified       *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
}

func (s DescribeMixStreamListResponseMixStreamList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixStreamListResponseMixStreamList) GoString() string {
	return s.String()
}

func (s *DescribeMixStreamListResponseMixStreamList) SetMixstreamId(v string) *DescribeMixStreamListResponseMixStreamList {
	s.MixstreamId = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetDomainName(v string) *DescribeMixStreamListResponseMixStreamList {
	s.DomainName = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetAppName(v string) *DescribeMixStreamListResponseMixStreamList {
	s.AppName = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetStreamName(v string) *DescribeMixStreamListResponseMixStreamList {
	s.StreamName = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetLayoutId(v string) *DescribeMixStreamListResponseMixStreamList {
	s.LayoutId = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetInputStreamNumber(v int) *DescribeMixStreamListResponseMixStreamList {
	s.InputStreamNumber = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetMixStreamTemplate(v string) *DescribeMixStreamListResponseMixStreamList {
	s.MixStreamTemplate = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetGmtCreate(v string) *DescribeMixStreamListResponseMixStreamList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMixStreamListResponseMixStreamList) SetGmtModified(v string) *DescribeMixStreamListResponseMixStreamList {
	s.GmtModified = &v
	return s
}

type AddRtsLiveStreamTranscodeRequest struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	App             *string `json:"App,omitempty" xml:"App,omitempty" require:"true"`
	Template        *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
	TemplateType    *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty" require:"true"`
	Height          *int    `json:"Height,omitempty" xml:"Height,omitempty"`
	Width           *int    `json:"Width,omitempty" xml:"Width,omitempty"`
	FPS             *int    `json:"FPS,omitempty" xml:"FPS,omitempty"`
	VideoBitrate    *int    `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	AudioBitrate    *int    `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	Gop             *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	DeleteBframes   *bool   `json:"DeleteBframes,omitempty" xml:"DeleteBframes,omitempty"`
	Opus            *bool   `json:"Opus,omitempty" xml:"Opus,omitempty"`
	Profile         *int    `json:"Profile,omitempty" xml:"Profile,omitempty"`
	AudioProfile    *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	AudioCodec      *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	AudioRate       *int    `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	AudioChannelNum *int    `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	Lazy            *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
}

func (s AddRtsLiveStreamTranscodeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRtsLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *AddRtsLiveStreamTranscodeRequest) SetDomain(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetApp(v string) *AddRtsLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetTemplate(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetTemplateType(v string) *AddRtsLiveStreamTranscodeRequest {
	s.TemplateType = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetHeight(v int) *AddRtsLiveStreamTranscodeRequest {
	s.Height = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetWidth(v int) *AddRtsLiveStreamTranscodeRequest {
	s.Width = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetFPS(v int) *AddRtsLiveStreamTranscodeRequest {
	s.FPS = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetVideoBitrate(v int) *AddRtsLiveStreamTranscodeRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioBitrate(v int) *AddRtsLiveStreamTranscodeRequest {
	s.AudioBitrate = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetGop(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Gop = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetDeleteBframes(v bool) *AddRtsLiveStreamTranscodeRequest {
	s.DeleteBframes = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetOpus(v bool) *AddRtsLiveStreamTranscodeRequest {
	s.Opus = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetProfile(v int) *AddRtsLiveStreamTranscodeRequest {
	s.Profile = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioProfile(v string) *AddRtsLiveStreamTranscodeRequest {
	s.AudioProfile = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioCodec(v string) *AddRtsLiveStreamTranscodeRequest {
	s.AudioCodec = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioRate(v int) *AddRtsLiveStreamTranscodeRequest {
	s.AudioRate = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetAudioChannelNum(v int) *AddRtsLiveStreamTranscodeRequest {
	s.AudioChannelNum = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeRequest) SetLazy(v string) *AddRtsLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

type AddRtsLiveStreamTranscodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddRtsLiveStreamTranscodeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRtsLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *AddRtsLiveStreamTranscodeResponse) SetRequestId(v string) *AddRtsLiveStreamTranscodeResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainTimeShiftDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeLiveDomainTimeShiftDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetDomainName(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetStartTime(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetEndTime(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetInterval(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.Interval = &v
	return s
}

type DescribeLiveDomainTimeShiftDataResponse struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TimeShiftData *DescribeLiveDomainTimeShiftDataResponseTimeShiftData `json:"TimeShiftData,omitempty" xml:"TimeShiftData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainTimeShiftDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponse) SetRequestId(v string) *DescribeLiveDomainTimeShiftDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponse) SetTimeShiftData(v *DescribeLiveDomainTimeShiftDataResponseTimeShiftData) *DescribeLiveDomainTimeShiftDataResponse {
	s.TimeShiftData = v
	return s
}

type DescribeLiveDomainTimeShiftDataResponseTimeShiftData struct {
	DataModule []*DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainTimeShiftDataResponseTimeShiftData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponseTimeShiftData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponseTimeShiftData) SetDataModule(v []*DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule) *DescribeLiveDomainTimeShiftDataResponseTimeShiftData {
	s.DataModule = v
	return s
}

type DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Size      *string `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule) SetTimeStamp(v string) *DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule) SetSize(v string) *DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule {
	s.Size = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule) SetType(v string) *DescribeLiveDomainTimeShiftDataResponseTimeShiftDataDataModule {
	s.Type = &v
	return s
}

type DeleteHtmlResourceRequest struct {
	HtmlResourceId *string `json:"HtmlResourceId,omitempty" xml:"HtmlResourceId,omitempty"`
	HtmlUrl        *string `json:"htmlUrl,omitempty" xml:"htmlUrl,omitempty"`
	CasterId       *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
}

func (s DeleteHtmlResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHtmlResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteHtmlResourceRequest) SetHtmlResourceId(v string) *DeleteHtmlResourceRequest {
	s.HtmlResourceId = &v
	return s
}

func (s *DeleteHtmlResourceRequest) SetHtmlUrl(v string) *DeleteHtmlResourceRequest {
	s.HtmlUrl = &v
	return s
}

func (s *DeleteHtmlResourceRequest) SetCasterId(v string) *DeleteHtmlResourceRequest {
	s.CasterId = &v
	return s
}

type DeleteHtmlResourceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteHtmlResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHtmlResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteHtmlResourceResponse) SetRequestId(v string) *DeleteHtmlResourceResponse {
	s.RequestId = &v
	return s
}

type DescribeHtmlResourceRequest struct {
	HtmlResourceId *string `json:"HtmlResourceId,omitempty" xml:"HtmlResourceId,omitempty"`
	HtmlUrl        *string `json:"htmlUrl,omitempty" xml:"htmlUrl,omitempty"`
	CasterId       *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
}

func (s DescribeHtmlResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHtmlResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeHtmlResourceRequest) SetHtmlResourceId(v string) *DescribeHtmlResourceRequest {
	s.HtmlResourceId = &v
	return s
}

func (s *DescribeHtmlResourceRequest) SetHtmlUrl(v string) *DescribeHtmlResourceRequest {
	s.HtmlUrl = &v
	return s
}

func (s *DescribeHtmlResourceRequest) SetCasterId(v string) *DescribeHtmlResourceRequest {
	s.CasterId = &v
	return s
}

type DescribeHtmlResourceResponse struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HtmlResource *DescribeHtmlResourceResponseHtmlResource `json:"HtmlResource,omitempty" xml:"HtmlResource,omitempty" require:"true" type:"Struct"`
}

func (s DescribeHtmlResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHtmlResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeHtmlResourceResponse) SetRequestId(v string) *DescribeHtmlResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeHtmlResourceResponse) SetHtmlResource(v *DescribeHtmlResourceResponseHtmlResource) *DescribeHtmlResourceResponse {
	s.HtmlResource = v
	return s
}

type DescribeHtmlResourceResponseHtmlResource struct {
	HtmlResourceId *string `json:"HtmlResourceId,omitempty" xml:"HtmlResourceId,omitempty" require:"true"`
	HtmlUrl        *string `json:"HtmlUrl,omitempty" xml:"HtmlUrl,omitempty" require:"true"`
	HtmlContent    *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty" require:"true"`
	CasterId       *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	Config         *string `json:"Config,omitempty" xml:"Config,omitempty" require:"true"`
	StreamId       *string `json:"StreamId,omitempty" xml:"StreamId,omitempty" require:"true"`
}

func (s DescribeHtmlResourceResponseHtmlResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeHtmlResourceResponseHtmlResource) GoString() string {
	return s.String()
}

func (s *DescribeHtmlResourceResponseHtmlResource) SetHtmlResourceId(v string) *DescribeHtmlResourceResponseHtmlResource {
	s.HtmlResourceId = &v
	return s
}

func (s *DescribeHtmlResourceResponseHtmlResource) SetHtmlUrl(v string) *DescribeHtmlResourceResponseHtmlResource {
	s.HtmlUrl = &v
	return s
}

func (s *DescribeHtmlResourceResponseHtmlResource) SetHtmlContent(v string) *DescribeHtmlResourceResponseHtmlResource {
	s.HtmlContent = &v
	return s
}

func (s *DescribeHtmlResourceResponseHtmlResource) SetCasterId(v string) *DescribeHtmlResourceResponseHtmlResource {
	s.CasterId = &v
	return s
}

func (s *DescribeHtmlResourceResponseHtmlResource) SetConfig(v string) *DescribeHtmlResourceResponseHtmlResource {
	s.Config = &v
	return s
}

func (s *DescribeHtmlResourceResponseHtmlResource) SetStreamId(v string) *DescribeHtmlResourceResponseHtmlResource {
	s.StreamId = &v
	return s
}

type ControlHtmlResourceRequest struct {
	HtmlResourceId *string `json:"HtmlResourceId,omitempty" xml:"HtmlResourceId,omitempty"`
	HtmlUrl        *string `json:"htmlUrl,omitempty" xml:"htmlUrl,omitempty"`
	CasterId       *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	Operate        *string `json:"Operate,omitempty" xml:"Operate,omitempty" require:"true"`
}

func (s ControlHtmlResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ControlHtmlResourceRequest) GoString() string {
	return s.String()
}

func (s *ControlHtmlResourceRequest) SetHtmlResourceId(v string) *ControlHtmlResourceRequest {
	s.HtmlResourceId = &v
	return s
}

func (s *ControlHtmlResourceRequest) SetHtmlUrl(v string) *ControlHtmlResourceRequest {
	s.HtmlUrl = &v
	return s
}

func (s *ControlHtmlResourceRequest) SetCasterId(v string) *ControlHtmlResourceRequest {
	s.CasterId = &v
	return s
}

func (s *ControlHtmlResourceRequest) SetOperate(v string) *ControlHtmlResourceRequest {
	s.Operate = &v
	return s
}

type ControlHtmlResourceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	StreamId  *string `json:"StreamId,omitempty" xml:"StreamId,omitempty" require:"true"`
}

func (s ControlHtmlResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ControlHtmlResourceResponse) GoString() string {
	return s.String()
}

func (s *ControlHtmlResourceResponse) SetRequestId(v string) *ControlHtmlResourceResponse {
	s.RequestId = &v
	return s
}

func (s *ControlHtmlResourceResponse) SetStreamId(v string) *ControlHtmlResourceResponse {
	s.StreamId = &v
	return s
}

type EditHtmlResourceRequest struct {
	HtmlResourceId *string `json:"HtmlResourceId,omitempty" xml:"HtmlResourceId,omitempty"`
	CasterId       *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	HtmlUrl        *string `json:"HtmlUrl,omitempty" xml:"HtmlUrl,omitempty"`
	HtmlContent    *string `json:"htmlContent,omitempty" xml:"htmlContent,omitempty"`
	Config         *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s EditHtmlResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s EditHtmlResourceRequest) GoString() string {
	return s.String()
}

func (s *EditHtmlResourceRequest) SetHtmlResourceId(v string) *EditHtmlResourceRequest {
	s.HtmlResourceId = &v
	return s
}

func (s *EditHtmlResourceRequest) SetCasterId(v string) *EditHtmlResourceRequest {
	s.CasterId = &v
	return s
}

func (s *EditHtmlResourceRequest) SetHtmlUrl(v string) *EditHtmlResourceRequest {
	s.HtmlUrl = &v
	return s
}

func (s *EditHtmlResourceRequest) SetHtmlContent(v string) *EditHtmlResourceRequest {
	s.HtmlContent = &v
	return s
}

func (s *EditHtmlResourceRequest) SetConfig(v string) *EditHtmlResourceRequest {
	s.Config = &v
	return s
}

type EditHtmlResourceResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HtmlResourceId *string `json:"HtmlResourceId,omitempty" xml:"HtmlResourceId,omitempty" require:"true"`
}

func (s EditHtmlResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s EditHtmlResourceResponse) GoString() string {
	return s.String()
}

func (s *EditHtmlResourceResponse) SetRequestId(v string) *EditHtmlResourceResponse {
	s.RequestId = &v
	return s
}

func (s *EditHtmlResourceResponse) SetHtmlResourceId(v string) *EditHtmlResourceResponse {
	s.HtmlResourceId = &v
	return s
}

type DescribeLiveUserTagsRequest struct {
}

func (s DescribeLiveUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserTagsRequest) GoString() string {
	return s.String()
}

type DescribeLiveUserTagsResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Tags      []*DescribeLiveUserTagsResponseTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTagsResponse) SetRequestId(v string) *DescribeLiveUserTagsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserTagsResponse) SetTags(v []*DescribeLiveUserTagsResponseTags) *DescribeLiveUserTagsResponse {
	s.Tags = v
	return s
}

type DescribeLiveUserTagsResponseTags struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveUserTagsResponseTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserTagsResponseTags) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTagsResponseTags) SetKey(v string) *DescribeLiveUserTagsResponseTags {
	s.Key = &v
	return s
}

func (s *DescribeLiveUserTagsResponseTags) SetValue(v []*string) *DescribeLiveUserTagsResponseTags {
	s.Value = v
	return s
}

type UnTagLiveResourcesRequest struct {
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UnTagLiveResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagLiveResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagLiveResourcesRequest) SetResourceId(v []*string) *UnTagLiveResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagLiveResourcesRequest) SetResourceType(v string) *UnTagLiveResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagLiveResourcesRequest) SetTagKey(v []*string) *UnTagLiveResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagLiveResourcesRequest) SetAll(v bool) *UnTagLiveResourcesRequest {
	s.All = &v
	return s
}

type UnTagLiveResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnTagLiveResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagLiveResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagLiveResourcesResponse) SetRequestId(v string) *UnTagLiveResourcesResponse {
	s.RequestId = &v
	return s
}

type TagLiveResourcesRequest struct {
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Tag          []*TagLiveResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s TagLiveResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagLiveResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesRequest) SetResourceId(v []*string) *TagLiveResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagLiveResourcesRequest) SetResourceType(v string) *TagLiveResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagLiveResourcesRequest) SetTag(v []*TagLiveResourcesRequestTag) *TagLiveResourcesRequest {
	s.Tag = v
	return s
}

type TagLiveResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagLiveResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagLiveResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesRequestTag) SetKey(v string) *TagLiveResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagLiveResourcesRequestTag) SetValue(v string) *TagLiveResourcesRequestTag {
	s.Value = &v
	return s
}

type TagLiveResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagLiveResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagLiveResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagLiveResourcesResponse) SetRequestId(v string) *TagLiveResourcesResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveTagResourcesRequest struct {
	ResourceId   []*string                             `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	ResourceType *string                               `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Tag          []*DescribeLiveTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLiveTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesRequest) SetResourceId(v []*string) *DescribeLiveTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeLiveTagResourcesRequest) SetResourceType(v string) *DescribeLiveTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeLiveTagResourcesRequest) SetTag(v []*DescribeLiveTagResourcesRequestTag) *DescribeLiveTagResourcesRequest {
	s.Tag = v
	return s
}

type DescribeLiveTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesRequestTag) SetKey(v string) *DescribeLiveTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeLiveTagResourcesRequestTag) SetValue(v string) *DescribeLiveTagResourcesRequestTag {
	s.Value = &v
	return s
}

type DescribeLiveTagResourcesResponse struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TagResources []*DescribeLiveTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponse) SetRequestId(v string) *DescribeLiveTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveTagResourcesResponse) SetTagResources(v []*DescribeLiveTagResourcesResponseTagResources) *DescribeLiveTagResourcesResponse {
	s.TagResources = v
	return s
}

type DescribeLiveTagResourcesResponseTagResources struct {
	ResourceId *string                                            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	Tag        []*DescribeLiveTagResourcesResponseTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponseTagResources) SetResourceId(v string) *DescribeLiveTagResourcesResponseTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeLiveTagResourcesResponseTagResources) SetTag(v []*DescribeLiveTagResourcesResponseTagResourcesTag) *DescribeLiveTagResourcesResponseTagResources {
	s.Tag = v
	return s
}

type DescribeLiveTagResourcesResponseTagResourcesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeLiveTagResourcesResponseTagResourcesTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTagResourcesResponseTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponseTagResourcesTag) SetKey(v string) *DescribeLiveTagResourcesResponseTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeLiveTagResourcesResponseTagResourcesTag) SetValue(v string) *DescribeLiveTagResourcesResponseTagResourcesTag {
	s.Value = &v
	return s
}

type AddLiveAudioAuditConfigRequest struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName  *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s AddLiveAudioAuditConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditConfigRequest) SetDomainName(v string) *AddLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetAppName(v string) *AddLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetStreamName(v string) *AddLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOssBucket(v string) *AddLiveAudioAuditConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOssEndpoint(v string) *AddLiveAudioAuditConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOssObject(v string) *AddLiveAudioAuditConfigRequest {
	s.OssObject = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetBizType(v string) *AddLiveAudioAuditConfigRequest {
	s.BizType = &v
	return s
}

type AddLiveAudioAuditConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveAudioAuditConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditConfigResponse) SetRequestId(v string) *AddLiveAudioAuditConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveAudioAuditConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
}

func (s DeleteLiveAudioAuditConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditConfigRequest) SetDomainName(v string) *DeleteLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) SetAppName(v string) *DeleteLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) SetStreamName(v string) *DeleteLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

type DeleteLiveAudioAuditConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveAudioAuditConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditConfigResponse) SetRequestId(v string) *DeleteLiveAudioAuditConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveAudioAuditConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveAudioAuditConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigRequest) SetDomainName(v string) *DescribeLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) SetAppName(v string) *DescribeLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) SetStreamName(v string) *DescribeLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

type DescribeLiveAudioAuditConfigResponse struct {
	RequestId                *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveAudioAuditConfigList *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList `json:"LiveAudioAuditConfigList,omitempty" xml:"LiveAudioAuditConfigList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveAudioAuditConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponse) SetRequestId(v string) *DescribeLiveAudioAuditConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponse) SetLiveAudioAuditConfigList(v *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList) *DescribeLiveAudioAuditConfigResponse {
	s.LiveAudioAuditConfigList = v
	return s
}

type DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList struct {
	LiveAudioAuditConfig []*DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig `json:"LiveAudioAuditConfig,omitempty" xml:"LiveAudioAuditConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList) SetLiveAudioAuditConfig(v []*DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigList {
	s.LiveAudioAuditConfig = v
	return s
}

type DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig struct {
	DomainName *string                                                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string                                                                                 `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string                                                                                 `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	BizType    *string                                                                                 `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	Scenes     *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) SetDomainName(v string) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) SetAppName(v string) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) SetStreamName(v string) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) SetBizType(v string) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.BizType = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig) SetScenes(v *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.Scenes = v
	return s
}

type DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes struct {
	// scene
	Scene []*string `json:"scene,omitempty" xml:"scene,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes) SetScene(v []*string) *DescribeLiveAudioAuditConfigResponseLiveAudioAuditConfigListLiveAudioAuditConfigScenes {
	s.Scene = v
	return s
}

type UpdateLiveAudioAuditConfigRequest struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName  *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s UpdateLiveAudioAuditConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditConfigRequest) SetDomainName(v string) *UpdateLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetAppName(v string) *UpdateLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetStreamName(v string) *UpdateLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOssBucket(v string) *UpdateLiveAudioAuditConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOssEndpoint(v string) *UpdateLiveAudioAuditConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOssObject(v string) *UpdateLiveAudioAuditConfigRequest {
	s.OssObject = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetBizType(v string) *UpdateLiveAudioAuditConfigRequest {
	s.BizType = &v
	return s
}

type UpdateLiveAudioAuditConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveAudioAuditConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditConfigResponse) SetRequestId(v string) *UpdateLiveAudioAuditConfigResponse {
	s.RequestId = &v
	return s
}

type AddLiveAudioAuditNotifyConfigRequest struct {
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Callback         *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackTemplate *string `json:"CallbackTemplate,omitempty" xml:"CallbackTemplate,omitempty"`
}

func (s AddLiveAudioAuditNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetCallback(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.Callback = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigRequest) SetCallbackTemplate(v string) *AddLiveAudioAuditNotifyConfigRequest {
	s.CallbackTemplate = &v
	return s
}

type AddLiveAudioAuditNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveAudioAuditNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditNotifyConfigResponse) SetRequestId(v string) *AddLiveAudioAuditNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveAudioAuditNotifyConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DeleteLiveAudioAuditNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *DeleteLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

type DeleteLiveAudioAuditNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveAudioAuditNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) SetRequestId(v string) *DeleteLiveAudioAuditNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveAudioAuditNotifyConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveAudioAuditNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *DescribeLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveAudioAuditNotifyConfigResponse struct {
	RequestId                      *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveAudioAuditNotifyConfigList *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList `json:"LiveAudioAuditNotifyConfigList,omitempty" xml:"LiveAudioAuditNotifyConfigList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) SetRequestId(v string) *DescribeLiveAudioAuditNotifyConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) SetLiveAudioAuditNotifyConfigList(v *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList) *DescribeLiveAudioAuditNotifyConfigResponse {
	s.LiveAudioAuditNotifyConfigList = v
	return s
}

type DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList struct {
	LiveAudioAuditNotifyConfig []*DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig `json:"LiveAudioAuditNotifyConfig,omitempty" xml:"LiveAudioAuditNotifyConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList) SetLiveAudioAuditNotifyConfig(v []*DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigList {
	s.LiveAudioAuditNotifyConfig = v
	return s
}

type DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig struct {
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Callback         *string `json:"Callback,omitempty" xml:"Callback,omitempty" require:"true"`
	CallbackTemplate *string `json:"CallbackTemplate,omitempty" xml:"CallbackTemplate,omitempty" require:"true"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) SetDomainName(v string) *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) SetCallback(v string) *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	s.Callback = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) SetCallbackTemplate(v string) *DescribeLiveAudioAuditNotifyConfigResponseLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	s.CallbackTemplate = &v
	return s
}

type UpdateLiveAudioAuditNotifyConfigRequest struct {
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Callback         *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackTemplate *string `json:"CallbackTemplate,omitempty" xml:"CallbackTemplate,omitempty"`
}

func (s UpdateLiveAudioAuditNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetCallback(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.Callback = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigRequest) SetCallbackTemplate(v string) *UpdateLiveAudioAuditNotifyConfigRequest {
	s.CallbackTemplate = &v
	return s
}

type UpdateLiveAudioAuditNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveAudioAuditNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) SetRequestId(v string) *UpdateLiveAudioAuditNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainPushTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeLiveDomainPushTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetDomainName(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetStartTime(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetEndTime(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetInterval(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetIspNameEn(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeLiveDomainPushTrafficDataResponse struct {
	RequestId              *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName             *string                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime              *string                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime                *string                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval           *string                                                          `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	TrafficDataPerInterval *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainPushTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetRequestId(v string) *DescribeLiveDomainPushTrafficDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetDomainName(v string) *DescribeLiveDomainPushTrafficDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetStartTime(v string) *DescribeLiveDomainPushTrafficDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetEndTime(v string) *DescribeLiveDomainPushTrafficDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetDataInterval(v string) *DescribeLiveDomainPushTrafficDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponse) SetTrafficDataPerInterval(v *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval) *DescribeLiveDomainPushTrafficDataResponse {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval struct {
	DataModule []*DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval) SetDataModule(v []*DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule) *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule struct {
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty" require:"true"`
}

func (s DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeLiveDomainPushTrafficDataResponseTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

type DescribeLiveDomainPushBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeLiveDomainPushBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetDomainName(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetStartTime(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetEndTime(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetInterval(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetIspNameEn(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeLiveDomainPushBpsDataResponse struct {
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName         *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime          *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime            *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval       *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	BpsDataPerInterval *DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainPushBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetRequestId(v string) *DescribeLiveDomainPushBpsDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetDomainName(v string) *DescribeLiveDomainPushBpsDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetStartTime(v string) *DescribeLiveDomainPushBpsDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetEndTime(v string) *DescribeLiveDomainPushBpsDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetDataInterval(v string) *DescribeLiveDomainPushBpsDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetBpsDataPerInterval(v *DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval) *DescribeLiveDomainPushBpsDataResponse {
	s.BpsDataPerInterval = v
	return s
}

type DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval struct {
	DataModule []*DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval) SetDataModule(v []*DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule) *DescribeLiveDomainPushBpsDataResponseBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	BpsValue  *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty" require:"true"`
}

func (s DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeLiveDomainPushBpsDataResponseBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

type SetCasterSyncGroupRequest struct {
	CasterId  *string                               `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SyncGroup []*SetCasterSyncGroupRequestSyncGroup `json:"SyncGroup,omitempty" xml:"SyncGroup,omitempty" type:"Repeated"`
}

func (s SetCasterSyncGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCasterSyncGroupRequest) GoString() string {
	return s.String()
}

func (s *SetCasterSyncGroupRequest) SetCasterId(v string) *SetCasterSyncGroupRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterSyncGroupRequest) SetSyncGroup(v []*SetCasterSyncGroupRequestSyncGroup) *SetCasterSyncGroupRequest {
	s.SyncGroup = v
	return s
}

type SetCasterSyncGroupRequestSyncGroup struct {
	Mode               *int      `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SyncDelayThreshold *int64    `json:"SyncDelayThreshold,omitempty" xml:"SyncDelayThreshold,omitempty"`
	HostResourceId     *string   `json:"HostResourceId,omitempty" xml:"HostResourceId,omitempty"`
	ResourceIds        []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s SetCasterSyncGroupRequestSyncGroup) String() string {
	return tea.Prettify(s)
}

func (s SetCasterSyncGroupRequestSyncGroup) GoString() string {
	return s.String()
}

func (s *SetCasterSyncGroupRequestSyncGroup) SetMode(v int) *SetCasterSyncGroupRequestSyncGroup {
	s.Mode = &v
	return s
}

func (s *SetCasterSyncGroupRequestSyncGroup) SetSyncDelayThreshold(v int64) *SetCasterSyncGroupRequestSyncGroup {
	s.SyncDelayThreshold = &v
	return s
}

func (s *SetCasterSyncGroupRequestSyncGroup) SetHostResourceId(v string) *SetCasterSyncGroupRequestSyncGroup {
	s.HostResourceId = &v
	return s
}

func (s *SetCasterSyncGroupRequestSyncGroup) SetResourceIds(v []*string) *SetCasterSyncGroupRequestSyncGroup {
	s.ResourceIds = v
	return s
}

type SetCasterSyncGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetCasterSyncGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCasterSyncGroupResponse) GoString() string {
	return s.String()
}

func (s *SetCasterSyncGroupResponse) SetRequestId(v string) *SetCasterSyncGroupResponse {
	s.RequestId = &v
	return s
}

type DescribeCasterSyncGroupRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DescribeCasterSyncGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSyncGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterSyncGroupRequest) SetCasterId(v string) *DescribeCasterSyncGroupRequest {
	s.CasterId = &v
	return s
}

type DescribeCasterSyncGroupResponse struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId   *string                                    `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SyncGroups *DescribeCasterSyncGroupResponseSyncGroups `json:"SyncGroups,omitempty" xml:"SyncGroups,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterSyncGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSyncGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterSyncGroupResponse) SetRequestId(v string) *DescribeCasterSyncGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterSyncGroupResponse) SetCasterId(v string) *DescribeCasterSyncGroupResponse {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterSyncGroupResponse) SetSyncGroups(v *DescribeCasterSyncGroupResponseSyncGroups) *DescribeCasterSyncGroupResponse {
	s.SyncGroups = v
	return s
}

type DescribeCasterSyncGroupResponseSyncGroups struct {
	SyncGroup []*DescribeCasterSyncGroupResponseSyncGroupsSyncGroup `json:"SyncGroup,omitempty" xml:"SyncGroup,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterSyncGroupResponseSyncGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSyncGroupResponseSyncGroups) GoString() string {
	return s.String()
}

func (s *DescribeCasterSyncGroupResponseSyncGroups) SetSyncGroup(v []*DescribeCasterSyncGroupResponseSyncGroupsSyncGroup) *DescribeCasterSyncGroupResponseSyncGroups {
	s.SyncGroup = v
	return s
}

type DescribeCasterSyncGroupResponseSyncGroupsSyncGroup struct {
	Mode           *int                                                           `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	HostResourceId *string                                                        `json:"HostResourceId,omitempty" xml:"HostResourceId,omitempty" require:"true"`
	ResourceIds    *DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterSyncGroupResponseSyncGroupsSyncGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSyncGroupResponseSyncGroupsSyncGroup) GoString() string {
	return s.String()
}

func (s *DescribeCasterSyncGroupResponseSyncGroupsSyncGroup) SetMode(v int) *DescribeCasterSyncGroupResponseSyncGroupsSyncGroup {
	s.Mode = &v
	return s
}

func (s *DescribeCasterSyncGroupResponseSyncGroupsSyncGroup) SetHostResourceId(v string) *DescribeCasterSyncGroupResponseSyncGroupsSyncGroup {
	s.HostResourceId = &v
	return s
}

func (s *DescribeCasterSyncGroupResponseSyncGroupsSyncGroup) SetResourceIds(v *DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds) *DescribeCasterSyncGroupResponseSyncGroupsSyncGroup {
	s.ResourceIds = v
	return s
}

type DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds struct {
	// ResourceId
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds) SetResourceId(v []*string) *DescribeCasterSyncGroupResponseSyncGroupsSyncGroupResourceIds {
	s.ResourceId = v
	return s
}

type DescribeLiveDetectPornDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	App        *string `json:"App,omitempty" xml:"App,omitempty"`
	Stream     *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	Fee        *string `json:"Fee,omitempty" xml:"Fee,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SplitBy    *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
}

func (s DescribeLiveDetectPornDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectPornDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataRequest) SetDomainName(v string) *DescribeLiveDetectPornDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetStartTime(v string) *DescribeLiveDetectPornDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetEndTime(v string) *DescribeLiveDetectPornDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetApp(v string) *DescribeLiveDetectPornDataRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetStream(v string) *DescribeLiveDetectPornDataRequest {
	s.Stream = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetFee(v string) *DescribeLiveDetectPornDataRequest {
	s.Fee = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetScene(v string) *DescribeLiveDetectPornDataRequest {
	s.Scene = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetRegion(v string) *DescribeLiveDetectPornDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeLiveDetectPornDataRequest) SetSplitBy(v string) *DescribeLiveDetectPornDataRequest {
	s.SplitBy = &v
	return s
}

type DescribeLiveDetectPornDataResponse struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DetectPornData *DescribeLiveDetectPornDataResponseDetectPornData `json:"DetectPornData,omitempty" xml:"DetectPornData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDetectPornDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponse) SetRequestId(v string) *DescribeLiveDetectPornDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponse) SetDetectPornData(v *DescribeLiveDetectPornDataResponseDetectPornData) *DescribeLiveDetectPornDataResponse {
	s.DetectPornData = v
	return s
}

type DescribeLiveDetectPornDataResponseDetectPornData struct {
	DataModule []*DescribeLiveDetectPornDataResponseDetectPornDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDetectPornDataResponseDetectPornData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponseDetectPornData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponseDetectPornData) SetDataModule(v []*DescribeLiveDetectPornDataResponseDetectPornDataDataModule) *DescribeLiveDetectPornDataResponseDetectPornData {
	s.DataModule = v
	return s
}

type DescribeLiveDetectPornDataResponseDetectPornDataDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	App       *string `json:"App,omitempty" xml:"App,omitempty" require:"true"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Stream    *string `json:"Stream,omitempty" xml:"Stream,omitempty" require:"true"`
	Fee       *string `json:"Fee,omitempty" xml:"Fee,omitempty" require:"true"`
	Scene     *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
}

func (s DescribeLiveDetectPornDataResponseDetectPornDataDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponseDetectPornDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetTimeStamp(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetApp(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.App = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetDomain(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetStream(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.Stream = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetFee(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.Fee = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetScene(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.Scene = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetRegion(v string) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.Region = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponseDetectPornDataDataModule) SetCount(v int64) *DescribeLiveDetectPornDataResponseDetectPornDataDataModule {
	s.Count = &v
	return s
}

type DeleteLiveRealTimeLogLogstoreRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
}

func (s DeleteLiveRealTimeLogLogstoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRealTimeLogLogstoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetProject(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.Project = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetLogstore(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetRegion(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.Region = &v
	return s
}

type DeleteLiveRealTimeLogLogstoreResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveRealTimeLogLogstoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRealTimeLogLogstoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealTimeLogLogstoreResponse) SetRequestId(v string) *DeleteLiveRealTimeLogLogstoreResponse {
	s.RequestId = &v
	return s
}

type DisableLiveRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DisableLiveRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DisableLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *DisableLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

type DisableLiveRealtimeLogDeliveryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableLiveRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DisableLiveRealtimeLogDeliveryResponse) SetRequestId(v string) *DisableLiveRealtimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

type EnableLiveRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s EnableLiveRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *EnableLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *EnableLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

type EnableLiveRealtimeLogDeliveryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableLiveRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *EnableLiveRealtimeLogDeliveryResponse) SetRequestId(v string) *EnableLiveRealtimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

type ListLiveRealtimeLogDeliveryDomainsRequest struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
}

func (s ListLiveRealtimeLogDeliveryDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetProject(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.Project = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetLogstore(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.Logstore = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetRegion(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.Region = &v
	return s
}

type ListLiveRealtimeLogDeliveryDomainsResponse struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Content   *ListLiveRealtimeLogDeliveryDomainsResponseContent `json:"Content,omitempty" xml:"Content,omitempty" require:"true" type:"Struct"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) SetRequestId(v string) *ListLiveRealtimeLogDeliveryDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) SetContent(v *ListLiveRealtimeLogDeliveryDomainsResponseContent) *ListLiveRealtimeLogDeliveryDomainsResponse {
	s.Content = v
	return s
}

type ListLiveRealtimeLogDeliveryDomainsResponseContent struct {
	Domains []*ListLiveRealtimeLogDeliveryDomainsResponseContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" require:"true" type:"Repeated"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseContent) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseContent) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseContent) SetDomains(v []*ListLiveRealtimeLogDeliveryDomainsResponseContentDomains) *ListLiveRealtimeLogDeliveryDomainsResponseContent {
	s.Domains = v
	return s
}

type ListLiveRealtimeLogDeliveryDomainsResponseContentDomains struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseContentDomains) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponseContentDomains) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseContentDomains) SetDomainName(v string) *ListLiveRealtimeLogDeliveryDomainsResponseContentDomains {
	s.DomainName = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponseContentDomains) SetStatus(v string) *ListLiveRealtimeLogDeliveryDomainsResponseContentDomains {
	s.Status = &v
	return s
}

type ModifyLiveRealtimeLogDeliveryRequest struct {
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s ModifyLiveRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetProject(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetLogstore(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetRegion(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

type ModifyLiveRealtimeLogDeliveryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyLiveRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveRealtimeLogDeliveryResponse) SetRequestId(v string) *ModifyLiveRealtimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveRealtimeDeliveryAccRequest struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	LogStore  *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s DescribeLiveRealtimeDeliveryAccRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetStartTime(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetEndTime(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetInterval(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetProject(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.Project = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetLogStore(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.LogStore = &v
	return s
}

type DescribeLiveRealtimeDeliveryAccResponse struct {
	RequestId               *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RealTimeDeliveryAccData *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData `json:"RealTimeDeliveryAccData,omitempty" xml:"RealTimeDeliveryAccData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveRealtimeDeliveryAccResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) SetRequestId(v string) *DescribeLiveRealtimeDeliveryAccResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) SetRealTimeDeliveryAccData(v *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData) *DescribeLiveRealtimeDeliveryAccResponse {
	s.RealTimeDeliveryAccData = v
	return s
}

type DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData struct {
	AccData []*DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData `json:"AccData,omitempty" xml:"AccData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData) SetAccData(v []*DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData) *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccData {
	s.AccData = v
	return s
}

type DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData struct {
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	SuccessNum *int    `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty" require:"true"`
	FailedNum  *int    `json:"FailedNum,omitempty" xml:"FailedNum,omitempty" require:"true"`
}

func (s DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData) SetTimeStamp(v string) *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData) SetSuccessNum(v int) *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData {
	s.SuccessNum = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData) SetFailedNum(v int) *DescribeLiveRealtimeDeliveryAccResponseRealTimeDeliveryAccDataAccData {
	s.FailedNum = &v
	return s
}

type DescribeLiveRealtimeLogAuthorizedRequest struct {
	LiveOpenapiReserve *string `json:"LiveOpenapiReserve,omitempty" xml:"LiveOpenapiReserve,omitempty"`
}

func (s DescribeLiveRealtimeLogAuthorizedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRealtimeLogAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) SetLiveOpenapiReserve(v string) *DescribeLiveRealtimeLogAuthorizedRequest {
	s.LiveOpenapiReserve = &v
	return s
}

type DescribeLiveRealtimeLogAuthorizedResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AuthorizedStatus *string `json:"AuthorizedStatus,omitempty" xml:"AuthorizedStatus,omitempty" require:"true"`
}

func (s DescribeLiveRealtimeLogAuthorizedResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRealtimeLogAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) SetRequestId(v string) *DescribeLiveRealtimeLogAuthorizedResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) SetAuthorizedStatus(v string) *DescribeLiveRealtimeLogAuthorizedResponse {
	s.AuthorizedStatus = &v
	return s
}

type ListLiveRealtimeLogDeliveryRequest struct {
	LiveOpenapiReserve *string `json:"LiveOpenapiReserve,omitempty" xml:"LiveOpenapiReserve,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryRequest) SetLiveOpenapiReserve(v string) *ListLiveRealtimeLogDeliveryRequest {
	s.LiveOpenapiReserve = &v
	return s
}

type ListLiveRealtimeLogDeliveryResponse struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Content   *ListLiveRealtimeLogDeliveryResponseContent `json:"Content,omitempty" xml:"Content,omitempty" require:"true" type:"Struct"`
}

func (s ListLiveRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponse) SetRequestId(v string) *ListLiveRealtimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponse) SetContent(v *ListLiveRealtimeLogDeliveryResponseContent) *ListLiveRealtimeLogDeliveryResponse {
	s.Content = v
	return s
}

type ListLiveRealtimeLogDeliveryResponseContent struct {
	RealtimeLogDeliveryInfo []*ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo `json:"RealtimeLogDeliveryInfo,omitempty" xml:"RealtimeLogDeliveryInfo,omitempty" require:"true" type:"Repeated"`
}

func (s ListLiveRealtimeLogDeliveryResponseContent) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponseContent) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponseContent) SetRealtimeLogDeliveryInfo(v []*ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) *ListLiveRealtimeLogDeliveryResponseContent {
	s.RealtimeLogDeliveryInfo = v
	return s
}

type ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo struct {
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DmId       *int    `json:"DmId,omitempty" xml:"DmId,omitempty" require:"true"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) SetProject(v string) *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo {
	s.Project = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) SetLogstore(v string) *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo {
	s.Logstore = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) SetRegion(v string) *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo {
	s.Region = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) SetDomainName(v string) *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo {
	s.DomainName = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) SetDmId(v int) *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo {
	s.DmId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo) SetStatus(v string) *ListLiveRealtimeLogDeliveryResponseContentRealtimeLogDeliveryInfo {
	s.Status = &v
	return s
}

type ListLiveRealtimeLogDeliveryInfosRequest struct {
	LiveOpenapiReserve *string `json:"LiveOpenapiReserve,omitempty" xml:"LiveOpenapiReserve,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) SetLiveOpenapiReserve(v string) *ListLiveRealtimeLogDeliveryInfosRequest {
	s.LiveOpenapiReserve = &v
	return s
}

type ListLiveRealtimeLogDeliveryInfosResponse struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Content   *ListLiveRealtimeLogDeliveryInfosResponseContent `json:"Content,omitempty" xml:"Content,omitempty" require:"true" type:"Struct"`
}

func (s ListLiveRealtimeLogDeliveryInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) SetRequestId(v string) *ListLiveRealtimeLogDeliveryInfosResponse {
	s.RequestId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponse) SetContent(v *ListLiveRealtimeLogDeliveryInfosResponseContent) *ListLiveRealtimeLogDeliveryInfosResponse {
	s.Content = v
	return s
}

type ListLiveRealtimeLogDeliveryInfosResponseContent struct {
	RealtimeLogDeliveryInfos []*ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos `json:"RealtimeLogDeliveryInfos,omitempty" xml:"RealtimeLogDeliveryInfos,omitempty" require:"true" type:"Repeated"`
}

func (s ListLiveRealtimeLogDeliveryInfosResponseContent) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponseContent) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseContent) SetRealtimeLogDeliveryInfos(v []*ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos) *ListLiveRealtimeLogDeliveryInfosResponseContent {
	s.RealtimeLogDeliveryInfos = v
	return s
}

type ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos struct {
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
}

func (s ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos) SetProject(v string) *ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos {
	s.Project = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos) SetLogstore(v string) *ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos {
	s.Logstore = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos) SetRegion(v string) *ListLiveRealtimeLogDeliveryInfosResponseContentRealtimeLogDeliveryInfos {
	s.Region = &v
	return s
}

type DescribeLiveDomainRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveDomainRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) SetDomainName(v string) *DescribeLiveDomainRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveDomainRealtimeLogDeliveryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	Logstore  *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeLiveDomainRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetRequestId(v string) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetProject(v string) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.Project = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetRegion(v string) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetLogstore(v string) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.Logstore = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponse) SetStatus(v string) *DescribeLiveDomainRealtimeLogDeliveryResponse {
	s.Status = &v
	return s
}

type DeleteLiveRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
}

func (s DeleteLiveRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetProject(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetLogstore(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetRegion(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

type DeleteLiveRealtimeLogDeliveryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealtimeLogDeliveryResponse) SetRequestId(v string) *DeleteLiveRealtimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

type CreateLiveRealTimeLogDeliveryRequest struct {
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s CreateLiveRealTimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetProject(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetLogstore(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetRegion(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetDomainName(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

type CreateLiveRealTimeLogDeliveryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateLiveRealTimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveRealTimeLogDeliveryResponse) SetRequestId(v string) *CreateLiveRealTimeLogDeliveryResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainLimitRequest struct {
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	LiveapiRequestFrom *string `json:"LiveapiRequestFrom,omitempty" xml:"LiveapiRequestFrom,omitempty"`
}

func (s DescribeLiveDomainLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitRequest) SetDomainName(v string) *DescribeLiveDomainLimitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLimitRequest) SetLiveapiRequestFrom(v string) *DescribeLiveDomainLimitRequest {
	s.LiveapiRequestFrom = &v
	return s
}

type DescribeLiveDomainLimitResponse struct {
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveDomainLimitList *DescribeLiveDomainLimitResponseLiveDomainLimitList `json:"LiveDomainLimitList,omitempty" xml:"LiveDomainLimitList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainLimitResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponse) SetRequestId(v string) *DescribeLiveDomainLimitResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainLimitResponse) SetLiveDomainLimitList(v *DescribeLiveDomainLimitResponseLiveDomainLimitList) *DescribeLiveDomainLimitResponse {
	s.LiveDomainLimitList = v
	return s
}

type DescribeLiveDomainLimitResponseLiveDomainLimitList struct {
	LiveDomainLimit []*DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit `json:"LiveDomainLimit,omitempty" xml:"LiveDomainLimit,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainLimitResponseLiveDomainLimitList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainLimitResponseLiveDomainLimitList) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponseLiveDomainLimitList) SetLiveDomainLimit(v []*DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit) *DescribeLiveDomainLimitResponseLiveDomainLimitList {
	s.LiveDomainLimit = v
	return s
}

type DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit struct {
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	LimitNum          *int    `json:"LimitNum,omitempty" xml:"LimitNum,omitempty" require:"true"`
	LimitTranscodeNum *int    `json:"LimitTranscodeNum,omitempty" xml:"LimitTranscodeNum,omitempty" require:"true"`
}

func (s DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit) SetDomainName(v string) *DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit) SetLimitNum(v int) *DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit {
	s.LimitNum = &v
	return s
}

func (s *DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit) SetLimitTranscodeNum(v int) *DescribeLiveDomainLimitResponseLiveDomainLimitListLiveDomainLimit {
	s.LimitTranscodeNum = &v
	return s
}

type DescribeLiveDomainBpsDataByTimeStampRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	TimePoint     *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty" require:"true"`
	IspNames      *string `json:"IspNames,omitempty" xml:"IspNames,omitempty" require:"true"`
	LocationNames *string `json:"LocationNames,omitempty" xml:"LocationNames,omitempty" require:"true"`
}

func (s DescribeLiveDomainBpsDataByTimeStampRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByTimeStampRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByTimeStampRequest) SetDomainName(v string) *DescribeLiveDomainBpsDataByTimeStampRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampRequest) SetTimePoint(v string) *DescribeLiveDomainBpsDataByTimeStampRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampRequest) SetIspNames(v string) *DescribeLiveDomainBpsDataByTimeStampRequest {
	s.IspNames = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampRequest) SetLocationNames(v string) *DescribeLiveDomainBpsDataByTimeStampRequest {
	s.LocationNames = &v
	return s
}

type DescribeLiveDomainBpsDataByTimeStampResponse struct {
	RequestId   *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName  *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	TimeStamp   *string                                                  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	BpsDataList *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList `json:"BpsDataList,omitempty" xml:"BpsDataList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainBpsDataByTimeStampResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByTimeStampResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponse) SetRequestId(v string) *DescribeLiveDomainBpsDataByTimeStampResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponse) SetDomainName(v string) *DescribeLiveDomainBpsDataByTimeStampResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponse) SetTimeStamp(v string) *DescribeLiveDomainBpsDataByTimeStampResponse {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponse) SetBpsDataList(v *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList) *DescribeLiveDomainBpsDataByTimeStampResponse {
	s.BpsDataList = v
	return s
}

type DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList struct {
	BpsDataModel []*DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel `json:"BpsDataModel,omitempty" xml:"BpsDataModel,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList) SetBpsDataModel(v []*DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataList {
	s.BpsDataModel = v
	return s
}

type DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel struct {
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty" require:"true"`
	IspName      *string `json:"IspName,omitempty" xml:"IspName,omitempty" require:"true"`
	Bps          *int64  `json:"Bps,omitempty" xml:"Bps,omitempty" require:"true"`
}

func (s DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) SetTimeStamp(v string) *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) SetLocationName(v string) *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel {
	s.LocationName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) SetIspName(v string) *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel {
	s.IspName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel) SetBps(v int64) *DescribeLiveDomainBpsDataByTimeStampResponseBpsDataListBpsDataModel {
	s.Bps = &v
	return s
}

type DescribeLiveStreamTranscodeStreamNumRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveStreamTranscodeStreamNumRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeStreamNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) SetDomainName(v string) *DescribeLiveStreamTranscodeStreamNumRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveStreamTranscodeStreamNumResponse struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total                *int64  `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	TranscodedNumber     *int64  `json:"TranscodedNumber,omitempty" xml:"TranscodedNumber,omitempty" require:"true"`
	UntranscodeNumber    *int64  `json:"UntranscodeNumber,omitempty" xml:"UntranscodeNumber,omitempty" require:"true"`
	LazyTranscodedNumber *int64  `json:"LazyTranscodedNumber,omitempty" xml:"LazyTranscodedNumber,omitempty" require:"true"`
}

func (s DescribeLiveStreamTranscodeStreamNumResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeStreamNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetRequestId(v string) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetTotal(v int64) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.Total = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetTranscodedNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.TranscodedNumber = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetUntranscodeNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.UntranscodeNumber = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetLazyTranscodedNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.LazyTranscodedNumber = &v
	return s
}

type UpdateLiveTopLevelDomainRequest struct {
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty" require:"true"`
}

func (s UpdateLiveTopLevelDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveTopLevelDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveTopLevelDomainRequest) SetSecurityToken(v string) *UpdateLiveTopLevelDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveTopLevelDomainRequest) SetDomainName(v string) *UpdateLiveTopLevelDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveTopLevelDomainRequest) SetTopLevelDomain(v string) *UpdateLiveTopLevelDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type UpdateLiveTopLevelDomainResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveTopLevelDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveTopLevelDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveTopLevelDomainResponse) SetRequestId(v string) *UpdateLiveTopLevelDomainResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainCertificateInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveDomainCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoRequest) SetDomainName(v string) *DescribeLiveDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveDomainCertificateInfoResponse struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CertInfos *DescribeLiveDomainCertificateInfoResponseCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponse) SetRequestId(v string) *DescribeLiveDomainCertificateInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponse) SetCertInfos(v *DescribeLiveDomainCertificateInfoResponseCertInfos) *DescribeLiveDomainCertificateInfoResponse {
	s.CertInfos = v
	return s
}

type DescribeLiveDomainCertificateInfoResponseCertInfos struct {
	CertInfo []*DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainCertificateInfoResponseCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfos) SetCertInfo(v []*DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) *DescribeLiveDomainCertificateInfoResponseCertInfos {
	s.CertInfo = v
	return s
}

type DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty" require:"true"`
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty" require:"true"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty" require:"true"`
	CertLife       *string `json:"CertLife,omitempty" xml:"CertLife,omitempty" require:"true"`
	CertOrg        *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty" require:"true"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty" require:"true"`
	SSLProtocol    *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty" require:"true"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SSLPub         *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty" require:"true"`
}

func (s DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetDomainName(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetCertName(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetCertDomainName(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetCertExpireTime(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetCertLife(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetCertOrg(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetCertType(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetSSLProtocol(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetStatus(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo) SetSSLPub(v string) *DescribeLiveDomainCertificateInfoResponseCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

type ModifyLiveDomainSchdmByPropertyRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Property   *string `json:"Property,omitempty" xml:"Property,omitempty" require:"true"`
}

func (s ModifyLiveDomainSchdmByPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLiveDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyLiveDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyLiveDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

type ModifyLiveDomainSchdmByPropertyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyLiveDomainSchdmByPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLiveDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) SetRequestId(v string) *ModifyLiveDomainSchdmByPropertyResponse {
	s.RequestId = &v
	return s
}

type SetLiveStreamOptimizedFeatureConfigRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	ConfigName   *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty" require:"true"`
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty" require:"true"`
	ConfigValue  *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
}

func (s SetLiveStreamOptimizedFeatureConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLiveStreamOptimizedFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamOptimizedFeatureConfigRequest) SetDomainName(v string) *SetLiveStreamOptimizedFeatureConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamOptimizedFeatureConfigRequest) SetConfigName(v string) *SetLiveStreamOptimizedFeatureConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *SetLiveStreamOptimizedFeatureConfigRequest) SetConfigStatus(v string) *SetLiveStreamOptimizedFeatureConfigRequest {
	s.ConfigStatus = &v
	return s
}

func (s *SetLiveStreamOptimizedFeatureConfigRequest) SetConfigValue(v string) *SetLiveStreamOptimizedFeatureConfigRequest {
	s.ConfigValue = &v
	return s
}

type SetLiveStreamOptimizedFeatureConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetLiveStreamOptimizedFeatureConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLiveStreamOptimizedFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamOptimizedFeatureConfigResponse) SetRequestId(v string) *SetLiveStreamOptimizedFeatureConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveStreamOptimizedFeatureConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty" require:"true"`
}

func (s DescribeLiveStreamOptimizedFeatureConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOptimizedFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOptimizedFeatureConfigRequest) SetDomainName(v string) *DescribeLiveStreamOptimizedFeatureConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamOptimizedFeatureConfigRequest) SetConfigName(v string) *DescribeLiveStreamOptimizedFeatureConfigRequest {
	s.ConfigName = &v
	return s
}

type DescribeLiveStreamOptimizedFeatureConfigResponse struct {
	RequestId                            *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveStreamOptimizedFeatureConfigList *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList `json:"LiveStreamOptimizedFeatureConfigList,omitempty" xml:"LiveStreamOptimizedFeatureConfigList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamOptimizedFeatureConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOptimizedFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponse) SetRequestId(v string) *DescribeLiveStreamOptimizedFeatureConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponse) SetLiveStreamOptimizedFeatureConfigList(v *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList) *DescribeLiveStreamOptimizedFeatureConfigResponse {
	s.LiveStreamOptimizedFeatureConfigList = v
	return s
}

type DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList struct {
	LiveStreamOptimizedFeatureConfig []*DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig `json:"LiveStreamOptimizedFeatureConfig,omitempty" xml:"LiveStreamOptimizedFeatureConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList) SetLiveStreamOptimizedFeatureConfig(v []*DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigList {
	s.LiveStreamOptimizedFeatureConfig = v
	return s
}

type DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	ConfigName   *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty" require:"true"`
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty" require:"true"`
	ConfigValue  *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty" require:"true"`
}

func (s DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) SetDomainName(v string) *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) SetConfigName(v string) *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig {
	s.ConfigName = &v
	return s
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) SetConfigStatus(v string) *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig {
	s.ConfigStatus = &v
	return s
}

func (s *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig) SetConfigValue(v string) *DescribeLiveStreamOptimizedFeatureConfigResponseLiveStreamOptimizedFeatureConfigListLiveStreamOptimizedFeatureConfig {
	s.ConfigValue = &v
	return s
}

type SetLiveStreamDelayConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	HlsDelay   *int    `json:"HlsDelay,omitempty" xml:"HlsDelay,omitempty"`
	HlsLevel   *string `json:"HlsLevel,omitempty" xml:"HlsLevel,omitempty"`
	FlvDelay   *int    `json:"FlvDelay,omitempty" xml:"FlvDelay,omitempty"`
	FlvLevel   *string `json:"FlvLevel,omitempty" xml:"FlvLevel,omitempty"`
	RtmpDelay  *int    `json:"RtmpDelay,omitempty" xml:"RtmpDelay,omitempty"`
	RtmpLevel  *string `json:"RtmpLevel,omitempty" xml:"RtmpLevel,omitempty"`
}

func (s SetLiveStreamDelayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLiveStreamDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamDelayConfigRequest) SetDomainName(v string) *SetLiveStreamDelayConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetHlsDelay(v int) *SetLiveStreamDelayConfigRequest {
	s.HlsDelay = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetHlsLevel(v string) *SetLiveStreamDelayConfigRequest {
	s.HlsLevel = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetFlvDelay(v int) *SetLiveStreamDelayConfigRequest {
	s.FlvDelay = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetFlvLevel(v string) *SetLiveStreamDelayConfigRequest {
	s.FlvLevel = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetRtmpDelay(v int) *SetLiveStreamDelayConfigRequest {
	s.RtmpDelay = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetRtmpLevel(v string) *SetLiveStreamDelayConfigRequest {
	s.RtmpLevel = &v
	return s
}

type SetLiveStreamDelayConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetLiveStreamDelayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLiveStreamDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamDelayConfigResponse) SetRequestId(v string) *SetLiveStreamDelayConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveStreamDelayConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveStreamDelayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigRequest) SetDomainName(v string) *DescribeLiveStreamDelayConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveStreamDelayConfigResponse struct {
	RequestId                 *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveStreamHlsDelayConfig  *DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig  `json:"LiveStreamHlsDelayConfig,omitempty" xml:"LiveStreamHlsDelayConfig,omitempty" require:"true" type:"Struct"`
	LiveStreamFlvDelayConfig  *DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig  `json:"LiveStreamFlvDelayConfig,omitempty" xml:"LiveStreamFlvDelayConfig,omitempty" require:"true" type:"Struct"`
	LiveStreamRtmpDelayConfig *DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig `json:"LiveStreamRtmpDelayConfig,omitempty" xml:"LiveStreamRtmpDelayConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamDelayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponse) SetRequestId(v string) *DescribeLiveStreamDelayConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponse) SetLiveStreamHlsDelayConfig(v *DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig) *DescribeLiveStreamDelayConfigResponse {
	s.LiveStreamHlsDelayConfig = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponse) SetLiveStreamFlvDelayConfig(v *DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig) *DescribeLiveStreamDelayConfigResponse {
	s.LiveStreamFlvDelayConfig = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponse) SetLiveStreamRtmpDelayConfig(v *DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig) *DescribeLiveStreamDelayConfigResponse {
	s.LiveStreamRtmpDelayConfig = v
	return s
}

type DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig struct {
	Level *string `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	Delay *int    `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
}

func (s DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig) SetLevel(v string) *DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig {
	s.Level = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig) SetDelay(v int) *DescribeLiveStreamDelayConfigResponseLiveStreamHlsDelayConfig {
	s.Delay = &v
	return s
}

type DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig struct {
	Level *string `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	Delay *int    `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
}

func (s DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig) SetLevel(v string) *DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig {
	s.Level = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig) SetDelay(v int) *DescribeLiveStreamDelayConfigResponseLiveStreamFlvDelayConfig {
	s.Delay = &v
	return s
}

type DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig struct {
	Level *string `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	Delay *int    `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
}

func (s DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig) SetLevel(v string) *DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig {
	s.Level = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig) SetDelay(v int) *DescribeLiveStreamDelayConfigResponseLiveStreamRtmpDelayConfig {
	s.Delay = &v
	return s
}

type DescribeLiveDomainOnlineUserNumRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	QueryTime  *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
}

func (s DescribeLiveDomainOnlineUserNumRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumRequest) SetDomainName(v string) *DescribeLiveDomainOnlineUserNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumRequest) SetQueryTime(v string) *DescribeLiveDomainOnlineUserNumRequest {
	s.QueryTime = &v
	return s
}

type DescribeLiveDomainOnlineUserNumResponse struct {
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	StreamCount    *int                                                   `json:"StreamCount,omitempty" xml:"StreamCount,omitempty" require:"true"`
	UserCount      *int                                                   `json:"UserCount,omitempty" xml:"UserCount,omitempty" require:"true"`
	OnlineUserInfo *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo `json:"OnlineUserInfo,omitempty" xml:"OnlineUserInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainOnlineUserNumResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetRequestId(v string) *DescribeLiveDomainOnlineUserNumResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetStreamCount(v int) *DescribeLiveDomainOnlineUserNumResponse {
	s.StreamCount = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetUserCount(v int) *DescribeLiveDomainOnlineUserNumResponse {
	s.UserCount = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetOnlineUserInfo(v *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo) *DescribeLiveDomainOnlineUserNumResponse {
	s.OnlineUserInfo = v
	return s
}

type DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo struct {
	LiveStreamOnlineUserNumInfo []*DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo `json:"LiveStreamOnlineUserNumInfo,omitempty" xml:"LiveStreamOnlineUserNumInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo) SetLiveStreamOnlineUserNumInfo(v []*DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfo {
	s.LiveStreamOnlineUserNumInfo = v
	return s
}

type DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo struct {
	StreamName *string                                                                                `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	Infos      *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos `json:"Infos,omitempty" xml:"Infos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) SetStreamName(v string) *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) SetInfos(v *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.Infos = v
	return s
}

type DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos struct {
	Info []*DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) SetInfo(v []*DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfos {
	s.Info = v
	return s
}

type DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo struct {
	TranscodeTemplate *string `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty" require:"true"`
	UserNumber        *int64  `json:"UserNumber,omitempty" xml:"UserNumber,omitempty" require:"true"`
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) SetTranscodeTemplate(v string) *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	s.TranscodeTemplate = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) SetUserNumber(v int64) *DescribeLiveDomainOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	s.UserNumber = &v
	return s
}

type DescribeLiveDomainFrameRateAndBitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	QueryTime  *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty" require:"true"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) SetDomainName(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataRequest) SetQueryTime(v string) *DescribeLiveDomainFrameRateAndBitRateDataRequest {
	s.QueryTime = &v
	return s
}

type DescribeLiveDomainFrameRateAndBitRateDataResponse struct {
	RequestId                *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FrameRateAndBitRateInfos *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos `json:"FrameRateAndBitRateInfos,omitempty" xml:"FrameRateAndBitRateInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) SetRequestId(v string) *DescribeLiveDomainFrameRateAndBitRateDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) SetFrameRateAndBitRateInfos(v *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) *DescribeLiveDomainFrameRateAndBitRateDataResponse {
	s.FrameRateAndBitRateInfos = v
	return s
}

type DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos struct {
	FrameRateAndBitRateInfo []*DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo `json:"FrameRateAndBitRateInfo,omitempty" xml:"FrameRateAndBitRateInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) SetFrameRateAndBitRateInfo(v []*DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos {
	s.FrameRateAndBitRateInfo = v
	return s
}

type DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo struct {
	AudioFrameRate *float32 `json:"AudioFrameRate,omitempty" xml:"AudioFrameRate,omitempty" require:"true"`
	BitRate        *float32 `json:"BitRate,omitempty" xml:"BitRate,omitempty" require:"true"`
	VideoFrameRate *float32 `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty" require:"true"`
	StreamUrl      *string  `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetAudioFrameRate(v float32) *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.AudioFrameRate = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetBitRate(v float32) *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.BitRate = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetVideoFrameRate(v float32) *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.VideoFrameRate = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetStreamUrl(v string) *DescribeLiveDomainFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.StreamUrl = &v
	return s
}

type SetBoardCallbackRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AuthKey        *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSwitch     *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	CallbackEnable *int    `json:"CallbackEnable,omitempty" xml:"CallbackEnable,omitempty" require:"true"`
	CallbackUri    *string `json:"CallbackUri,omitempty" xml:"CallbackUri,omitempty"`
	CallbackEvents *string `json:"CallbackEvents,omitempty" xml:"CallbackEvents,omitempty"`
}

func (s SetBoardCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBoardCallbackRequest) GoString() string {
	return s.String()
}

func (s *SetBoardCallbackRequest) SetAppId(v string) *SetBoardCallbackRequest {
	s.AppId = &v
	return s
}

func (s *SetBoardCallbackRequest) SetAuthKey(v string) *SetBoardCallbackRequest {
	s.AuthKey = &v
	return s
}

func (s *SetBoardCallbackRequest) SetAuthSwitch(v string) *SetBoardCallbackRequest {
	s.AuthSwitch = &v
	return s
}

func (s *SetBoardCallbackRequest) SetCallbackEnable(v int) *SetBoardCallbackRequest {
	s.CallbackEnable = &v
	return s
}

func (s *SetBoardCallbackRequest) SetCallbackUri(v string) *SetBoardCallbackRequest {
	s.CallbackUri = &v
	return s
}

func (s *SetBoardCallbackRequest) SetCallbackEvents(v string) *SetBoardCallbackRequest {
	s.CallbackEvents = &v
	return s
}

type SetBoardCallbackResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetBoardCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBoardCallbackResponse) GoString() string {
	return s.String()
}

func (s *SetBoardCallbackResponse) SetRequestId(v string) *SetBoardCallbackResponse {
	s.RequestId = &v
	return s
}

type DescribeRecordsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	PageNum     *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RecordState *string `json:"RecordState,omitempty" xml:"RecordState,omitempty"`
}

func (s DescribeRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordsRequest) SetAppId(v string) *DescribeRecordsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordsRequest) SetPageNum(v int) *DescribeRecordsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordsRequest) SetPageSize(v int) *DescribeRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordsRequest) SetRecordState(v string) *DescribeRecordsRequest {
	s.RecordState = &v
	return s
}

type DescribeRecordsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Records   []*DescribeRecordsResponseRecords `json:"Records,omitempty" xml:"Records,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponse) SetRequestId(v string) *DescribeRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordsResponse) SetRecords(v []*DescribeRecordsResponseRecords) *DescribeRecordsResponse {
	s.Records = v
	return s
}

type DescribeRecordsResponseRecords struct {
	RecordId        *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardId         *int    `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
	RecordStartTime *int64  `json:"RecordStartTime,omitempty" xml:"RecordStartTime,omitempty" require:"true"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	State           *int    `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	OssPath         *string `json:"OssPath,omitempty" xml:"OssPath,omitempty" require:"true"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssEndpoint     *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
}

func (s DescribeRecordsResponseRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordsResponseRecords) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponseRecords) SetRecordId(v string) *DescribeRecordsResponseRecords {
	s.RecordId = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetAppId(v string) *DescribeRecordsResponseRecords {
	s.AppId = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetBoardId(v int) *DescribeRecordsResponseRecords {
	s.BoardId = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetRecordStartTime(v int64) *DescribeRecordsResponseRecords {
	s.RecordStartTime = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetStartTime(v int64) *DescribeRecordsResponseRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetEndTime(v int64) *DescribeRecordsResponseRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetState(v int) *DescribeRecordsResponseRecords {
	s.State = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetOssPath(v string) *DescribeRecordsResponseRecords {
	s.OssPath = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetOssBucket(v string) *DescribeRecordsResponseRecords {
	s.OssBucket = &v
	return s
}

func (s *DescribeRecordsResponseRecords) SetOssEndpoint(v string) *DescribeRecordsResponseRecords {
	s.OssEndpoint = &v
	return s
}

type DescribeRecordRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
}

func (s DescribeRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordRequest) SetAppId(v string) *DescribeRecordRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordRequest) SetRecordId(v string) *DescribeRecordRequest {
	s.RecordId = &v
	return s
}

type DescribeRecordResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RecordId        *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardId         *int    `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
	RecordStartTime *int64  `json:"RecordStartTime,omitempty" xml:"RecordStartTime,omitempty" require:"true"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	State           *int    `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	OssPath         *string `json:"OssPath,omitempty" xml:"OssPath,omitempty" require:"true"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssEndpoint     *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
}

func (s DescribeRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordResponse) SetRequestId(v string) *DescribeRecordResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordResponse) SetRecordId(v string) *DescribeRecordResponse {
	s.RecordId = &v
	return s
}

func (s *DescribeRecordResponse) SetAppId(v string) *DescribeRecordResponse {
	s.AppId = &v
	return s
}

func (s *DescribeRecordResponse) SetBoardId(v int) *DescribeRecordResponse {
	s.BoardId = &v
	return s
}

func (s *DescribeRecordResponse) SetRecordStartTime(v int64) *DescribeRecordResponse {
	s.RecordStartTime = &v
	return s
}

func (s *DescribeRecordResponse) SetStartTime(v int64) *DescribeRecordResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordResponse) SetEndTime(v int64) *DescribeRecordResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordResponse) SetState(v int) *DescribeRecordResponse {
	s.State = &v
	return s
}

func (s *DescribeRecordResponse) SetOssPath(v string) *DescribeRecordResponse {
	s.OssPath = &v
	return s
}

func (s *DescribeRecordResponse) SetOssBucket(v string) *DescribeRecordResponse {
	s.OssBucket = &v
	return s
}

func (s *DescribeRecordResponse) SetOssEndpoint(v string) *DescribeRecordResponse {
	s.OssEndpoint = &v
	return s
}

type CompleteBoardRecordRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s CompleteBoardRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteBoardRecordRequest) GoString() string {
	return s.String()
}

func (s *CompleteBoardRecordRequest) SetAppId(v string) *CompleteBoardRecordRequest {
	s.AppId = &v
	return s
}

func (s *CompleteBoardRecordRequest) SetRecordId(v string) *CompleteBoardRecordRequest {
	s.RecordId = &v
	return s
}

func (s *CompleteBoardRecordRequest) SetEndTime(v string) *CompleteBoardRecordRequest {
	s.EndTime = &v
	return s
}

type CompleteBoardRecordResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OssPath   *string `json:"OssPath,omitempty" xml:"OssPath,omitempty" require:"true"`
}

func (s CompleteBoardRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteBoardRecordResponse) GoString() string {
	return s.String()
}

func (s *CompleteBoardRecordResponse) SetRequestId(v string) *CompleteBoardRecordResponse {
	s.RequestId = &v
	return s
}

func (s *CompleteBoardRecordResponse) SetOssPath(v string) *CompleteBoardRecordResponse {
	s.OssPath = &v
	return s
}

type StartBoardRecordRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardId   *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
}

func (s StartBoardRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StartBoardRecordRequest) GoString() string {
	return s.String()
}

func (s *StartBoardRecordRequest) SetAppId(v string) *StartBoardRecordRequest {
	s.AppId = &v
	return s
}

func (s *StartBoardRecordRequest) SetBoardId(v string) *StartBoardRecordRequest {
	s.BoardId = &v
	return s
}

func (s *StartBoardRecordRequest) SetStartTime(v string) *StartBoardRecordRequest {
	s.StartTime = &v
	return s
}

type StartBoardRecordResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RecordId  *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
}

func (s StartBoardRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBoardRecordResponse) GoString() string {
	return s.String()
}

func (s *StartBoardRecordResponse) SetRequestId(v string) *StartBoardRecordResponse {
	s.RequestId = &v
	return s
}

func (s *StartBoardRecordResponse) SetRecordId(v string) *StartBoardRecordResponse {
	s.RecordId = &v
	return s
}

type ApplyRecordTokenRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ApplyRecordTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyRecordTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyRecordTokenRequest) SetAppId(v string) *ApplyRecordTokenRequest {
	s.AppId = &v
	return s
}

type ApplyRecordTokenResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty" require:"true"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty" require:"true"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty" require:"true"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty" require:"true"`
}

func (s ApplyRecordTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyRecordTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyRecordTokenResponse) SetRequestId(v string) *ApplyRecordTokenResponse {
	s.RequestId = &v
	return s
}

func (s *ApplyRecordTokenResponse) SetSecurityToken(v string) *ApplyRecordTokenResponse {
	s.SecurityToken = &v
	return s
}

func (s *ApplyRecordTokenResponse) SetAccessKeySecret(v string) *ApplyRecordTokenResponse {
	s.AccessKeySecret = &v
	return s
}

func (s *ApplyRecordTokenResponse) SetAccessKeyId(v string) *ApplyRecordTokenResponse {
	s.AccessKeyId = &v
	return s
}

func (s *ApplyRecordTokenResponse) SetExpiration(v string) *ApplyRecordTokenResponse {
	s.Expiration = &v
	return s
}

type UpdateBoardCallbackRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AuthKey        *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSwitch     *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	CallbackEnable *int    `json:"CallbackEnable,omitempty" xml:"CallbackEnable,omitempty" require:"true"`
	CallbackUri    *string `json:"CallbackUri,omitempty" xml:"CallbackUri,omitempty"`
	CallbackEvents *string `json:"CallbackEvents,omitempty" xml:"CallbackEvents,omitempty"`
}

func (s UpdateBoardCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBoardCallbackRequest) GoString() string {
	return s.String()
}

func (s *UpdateBoardCallbackRequest) SetAppId(v string) *UpdateBoardCallbackRequest {
	s.AppId = &v
	return s
}

func (s *UpdateBoardCallbackRequest) SetAuthKey(v string) *UpdateBoardCallbackRequest {
	s.AuthKey = &v
	return s
}

func (s *UpdateBoardCallbackRequest) SetAuthSwitch(v string) *UpdateBoardCallbackRequest {
	s.AuthSwitch = &v
	return s
}

func (s *UpdateBoardCallbackRequest) SetCallbackEnable(v int) *UpdateBoardCallbackRequest {
	s.CallbackEnable = &v
	return s
}

func (s *UpdateBoardCallbackRequest) SetCallbackUri(v string) *UpdateBoardCallbackRequest {
	s.CallbackUri = &v
	return s
}

func (s *UpdateBoardCallbackRequest) SetCallbackEvents(v string) *UpdateBoardCallbackRequest {
	s.CallbackEvents = &v
	return s
}

type UpdateBoardCallbackResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateBoardCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBoardCallbackResponse) GoString() string {
	return s.String()
}

func (s *UpdateBoardCallbackResponse) SetRequestId(v string) *UpdateBoardCallbackResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainMappingRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveDomainMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainMappingRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingRequest) SetDomainName(v string) *DescribeLiveDomainMappingRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveDomainMappingResponse struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveDomainModels *DescribeLiveDomainMappingResponseLiveDomainModels `json:"LiveDomainModels,omitempty" xml:"LiveDomainModels,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainMappingResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponse) SetRequestId(v string) *DescribeLiveDomainMappingResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainMappingResponse) SetLiveDomainModels(v *DescribeLiveDomainMappingResponseLiveDomainModels) *DescribeLiveDomainMappingResponse {
	s.LiveDomainModels = v
	return s
}

type DescribeLiveDomainMappingResponseLiveDomainModels struct {
	LiveDomainModel []*DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel `json:"LiveDomainModel,omitempty" xml:"LiveDomainModel,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainMappingResponseLiveDomainModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainMappingResponseLiveDomainModels) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponseLiveDomainModels) SetLiveDomainModel(v []*DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel) *DescribeLiveDomainMappingResponseLiveDomainModels {
	s.LiveDomainModel = v
	return s
}

type DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel) SetDomainName(v string) *DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel) SetType(v string) *DescribeLiveDomainMappingResponseLiveDomainModelsLiveDomainModel {
	s.Type = &v
	return s
}

type StopLiveIndexRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s StopLiveIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveIndexRequest) GoString() string {
	return s.String()
}

func (s *StopLiveIndexRequest) SetDomainName(v string) *StopLiveIndexRequest {
	s.DomainName = &v
	return s
}

func (s *StopLiveIndexRequest) SetAppName(v string) *StopLiveIndexRequest {
	s.AppName = &v
	return s
}

func (s *StopLiveIndexRequest) SetStreamName(v string) *StopLiveIndexRequest {
	s.StreamName = &v
	return s
}

func (s *StopLiveIndexRequest) SetTaskId(v string) *StopLiveIndexRequest {
	s.TaskId = &v
	return s
}

type StopLiveIndexResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopLiveIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveIndexResponse) GoString() string {
	return s.String()
}

func (s *StopLiveIndexResponse) SetRequestId(v string) *StopLiveIndexResponse {
	s.RequestId = &v
	return s
}

type StartLiveIndexRequest struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName  *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	TokenId     *string `json:"TokenId,omitempty" xml:"TokenId,omitempty" require:"true"`
	InputUrl    *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty" require:"true"`
	Interval    *int    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssUserId   *string `json:"OssUserId,omitempty" xml:"OssUserId,omitempty"`
	OssRamRole  *string `json:"OssRamRole,omitempty" xml:"OssRamRole,omitempty"`
}

func (s StartLiveIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s StartLiveIndexRequest) GoString() string {
	return s.String()
}

func (s *StartLiveIndexRequest) SetDomainName(v string) *StartLiveIndexRequest {
	s.DomainName = &v
	return s
}

func (s *StartLiveIndexRequest) SetAppName(v string) *StartLiveIndexRequest {
	s.AppName = &v
	return s
}

func (s *StartLiveIndexRequest) SetStreamName(v string) *StartLiveIndexRequest {
	s.StreamName = &v
	return s
}

func (s *StartLiveIndexRequest) SetTokenId(v string) *StartLiveIndexRequest {
	s.TokenId = &v
	return s
}

func (s *StartLiveIndexRequest) SetInputUrl(v string) *StartLiveIndexRequest {
	s.InputUrl = &v
	return s
}

func (s *StartLiveIndexRequest) SetInterval(v int) *StartLiveIndexRequest {
	s.Interval = &v
	return s
}

func (s *StartLiveIndexRequest) SetOssBucket(v string) *StartLiveIndexRequest {
	s.OssBucket = &v
	return s
}

func (s *StartLiveIndexRequest) SetOssEndpoint(v string) *StartLiveIndexRequest {
	s.OssEndpoint = &v
	return s
}

func (s *StartLiveIndexRequest) SetOssUserId(v string) *StartLiveIndexRequest {
	s.OssUserId = &v
	return s
}

func (s *StartLiveIndexRequest) SetOssRamRole(v string) *StartLiveIndexRequest {
	s.OssRamRole = &v
	return s
}

type StartLiveIndexResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s StartLiveIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s StartLiveIndexResponse) GoString() string {
	return s.String()
}

func (s *StartLiveIndexResponse) SetRequestId(v string) *StartLiveIndexResponse {
	s.RequestId = &v
	return s
}

func (s *StartLiveIndexResponse) SetTaskId(v string) *StartLiveIndexResponse {
	s.TaskId = &v
	return s
}

type RealTimeSnapshotCommandRequest struct {
	Command    *string `json:"Command,omitempty" xml:"Command,omitempty" require:"true"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	Mode       *int    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Interval   *int    `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s RealTimeSnapshotCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RealTimeSnapshotCommandRequest) GoString() string {
	return s.String()
}

func (s *RealTimeSnapshotCommandRequest) SetCommand(v string) *RealTimeSnapshotCommandRequest {
	s.Command = &v
	return s
}

func (s *RealTimeSnapshotCommandRequest) SetDomainName(v string) *RealTimeSnapshotCommandRequest {
	s.DomainName = &v
	return s
}

func (s *RealTimeSnapshotCommandRequest) SetAppName(v string) *RealTimeSnapshotCommandRequest {
	s.AppName = &v
	return s
}

func (s *RealTimeSnapshotCommandRequest) SetStreamName(v string) *RealTimeSnapshotCommandRequest {
	s.StreamName = &v
	return s
}

func (s *RealTimeSnapshotCommandRequest) SetMode(v int) *RealTimeSnapshotCommandRequest {
	s.Mode = &v
	return s
}

func (s *RealTimeSnapshotCommandRequest) SetInterval(v int) *RealTimeSnapshotCommandRequest {
	s.Interval = &v
	return s
}

type RealTimeSnapshotCommandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RealTimeSnapshotCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RealTimeSnapshotCommandResponse) GoString() string {
	return s.String()
}

func (s *RealTimeSnapshotCommandResponse) SetRequestId(v string) *RealTimeSnapshotCommandResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveTopDomainsByFlowRequest struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s DescribeLiveTopDomainsByFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetStartTime(v string) *DescribeLiveTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetEndTime(v string) *DescribeLiveTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetLimit(v int64) *DescribeLiveTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

type DescribeLiveTopDomainsByFlowResponse struct {
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	StartTime         *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime           *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DomainCount       *int64                                          `json:"DomainCount,omitempty" xml:"DomainCount,omitempty" require:"true"`
	DomainOnlineCount *int64                                          `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty" require:"true"`
	TopDomains        *DescribeLiveTopDomainsByFlowResponseTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveTopDomainsByFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetRequestId(v string) *DescribeLiveTopDomainsByFlowResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetStartTime(v string) *DescribeLiveTopDomainsByFlowResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetEndTime(v string) *DescribeLiveTopDomainsByFlowResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetDomainCount(v int64) *DescribeLiveTopDomainsByFlowResponse {
	s.DomainCount = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetDomainOnlineCount(v int64) *DescribeLiveTopDomainsByFlowResponse {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponse) SetTopDomains(v *DescribeLiveTopDomainsByFlowResponseTopDomains) *DescribeLiveTopDomainsByFlowResponse {
	s.TopDomains = v
	return s
}

type DescribeLiveTopDomainsByFlowResponseTopDomains struct {
	TopDomain []*DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveTopDomainsByFlowResponseTopDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponseTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomains) SetTopDomain(v []*DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) *DescribeLiveTopDomainsByFlowResponseTopDomains {
	s.TopDomain = v
	return s
}

type DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Rank           *int64  `json:"Rank,omitempty" xml:"Rank,omitempty" require:"true"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty" require:"true"`
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty" require:"true"`
	MaxBps         *int64  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty" require:"true"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty" require:"true"`
	TotalAccess    *int64  `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty" require:"true"`
}

func (s DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetDomainName(v string) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetRank(v int64) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetMaxBps(v int64) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeLiveTopDomainsByFlowResponseTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

type DescribeLiveDomainRealTimeBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveDomainRealTimeBpsDataResponse struct {
	RequestId                  *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName                 *string                                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime                  *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime                    *string                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval               *string                                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	RealTimeBpsDataPerInterval *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval `json:"RealTimeBpsDataPerInterval,omitempty" xml:"RealTimeBpsDataPerInterval,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetRequestId(v string) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetDomainName(v string) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetStartTime(v string) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetEndTime(v string) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetDataInterval(v string) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponse) SetRealTimeBpsDataPerInterval(v *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval) *DescribeLiveDomainRealTimeBpsDataResponse {
	s.RealTimeBpsDataPerInterval = v
	return s
}

type DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval struct {
	DataModule []*DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval) SetDataModule(v []*DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule) *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule) SetValue(v string) *DescribeLiveDomainRealTimeBpsDataResponseRealTimeBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeLiveDomainRealTimeHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeLiveDomainRealTimeHttpCodeDataResponse struct {
	RequestId            *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName           *string                                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime            *string                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime              *string                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval         *string                                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	RealTimeHttpCodeData *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData `json:"RealTimeHttpCodeData,omitempty" xml:"RealTimeHttpCodeData,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetRequestId(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetDomainName(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetStartTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetEndTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetDataInterval(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetRealTimeHttpCodeData(v *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.RealTimeHttpCodeData = v
	return s
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData struct {
	UsageData []*DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData) SetUsageData(v []*DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData struct {
	TimeStamp *string                                                                           `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Value     *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData) SetValue(v *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue struct {
	RealTimeCodeProportionData []*DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData `json:"RealTimeCodeProportionData,omitempty" xml:"RealTimeCodeProportionData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue) SetRealTimeCodeProportionData(v []*DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValue {
	s.RealTimeCodeProportionData = v
	return s
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty" require:"true"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCode(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetProportion(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCount(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Count = &v
	return s
}

type DescribeLiveDomainRealTimeTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetIspNameEn(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveDomainRealTimeTrafficDataResponse struct {
	RequestId                      *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName                     *string                                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime                      *string                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime                        *string                                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval                   *string                                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	RealTimeTrafficDataPerInterval *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetRequestId(v string) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetDomainName(v string) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetStartTime(v string) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetEndTime(v string) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetDataInterval(v string) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponse) SetRealTimeTrafficDataPerInterval(v *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval) *DescribeLiveDomainRealTimeTrafficDataResponse {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

type DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule) *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeLiveDomainRealTimeTrafficDataResponseRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type AddLiveDomainPlayMappingRequest struct {
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty" require:"true"`
	PullDomain *string `json:"PullDomain,omitempty" xml:"PullDomain,omitempty" require:"true"`
}

func (s AddLiveDomainPlayMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDomainPlayMappingRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDomainPlayMappingRequest) SetPlayDomain(v string) *AddLiveDomainPlayMappingRequest {
	s.PlayDomain = &v
	return s
}

func (s *AddLiveDomainPlayMappingRequest) SetPullDomain(v string) *AddLiveDomainPlayMappingRequest {
	s.PullDomain = &v
	return s
}

type AddLiveDomainPlayMappingResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveDomainPlayMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDomainPlayMappingResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDomainPlayMappingResponse) SetRequestId(v string) *AddLiveDomainPlayMappingResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveLazyPullStreamInfoConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
}

func (s DeleteLiveLazyPullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveLazyPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) SetDomainName(v string) *DeleteLiveLazyPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) SetAppName(v string) *DeleteLiveLazyPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

type DeleteLiveLazyPullStreamInfoConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveLazyPullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveLazyPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveLazyPullStreamInfoConfigResponse) SetRequestId(v string) *DeleteLiveLazyPullStreamInfoConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveLazyPullStreamConfigRequest struct {
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	LiveapiRequestFrom *string `json:"LiveapiRequestFrom,omitempty" xml:"LiveapiRequestFrom,omitempty"`
}

func (s DescribeLiveLazyPullStreamConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetDomainName(v string) *DescribeLiveLazyPullStreamConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetAppName(v string) *DescribeLiveLazyPullStreamConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetLiveapiRequestFrom(v string) *DescribeLiveLazyPullStreamConfigRequest {
	s.LiveapiRequestFrom = &v
	return s
}

type DescribeLiveLazyPullStreamConfigResponse struct {
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveLazyPullConfigList *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList `json:"LiveLazyPullConfigList,omitempty" xml:"LiveLazyPullConfigList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveLazyPullStreamConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponse) SetRequestId(v string) *DescribeLiveLazyPullStreamConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponse) SetLiveLazyPullConfigList(v *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList) *DescribeLiveLazyPullStreamConfigResponse {
	s.LiveLazyPullConfigList = v
	return s
}

type DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList struct {
	LiveLazyPullConfig []*DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig `json:"LiveLazyPullConfig,omitempty" xml:"LiveLazyPullConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList) SetLiveLazyPullConfig(v []*DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigList {
	s.LiveLazyPullConfig = v
	return s
}

type DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	PullDomainName *string `json:"PullDomainName,omitempty" xml:"PullDomainName,omitempty" require:"true"`
	PullAppName    *string `json:"PullAppName,omitempty" xml:"PullAppName,omitempty" require:"true"`
	PullProtocol   *string `json:"PullProtocol,omitempty" xml:"PullProtocol,omitempty" require:"true"`
	PullAuthType   *string `json:"PullAuthType,omitempty" xml:"PullAuthType,omitempty" require:"true"`
	PullAuthKey    *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty" require:"true"`
	PullArgs       *string `json:"PullArgs,omitempty" xml:"PullArgs,omitempty" require:"true"`
}

func (s DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetDomainName(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetAppName(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetPullDomainName(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullDomainName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetPullAppName(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullAppName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetPullProtocol(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullProtocol = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetPullAuthType(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullAuthType = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetPullAuthKey(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullAuthKey = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig) SetPullArgs(v string) *DescribeLiveLazyPullStreamConfigResponseLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullArgs = &v
	return s
}

type SetLiveLazyPullStreamInfoConfigRequest struct {
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	PullDomainName     *string `json:"PullDomainName,omitempty" xml:"PullDomainName,omitempty" require:"true"`
	PullAppName        *string `json:"PullAppName,omitempty" xml:"PullAppName,omitempty"`
	PullProtocol       *string `json:"PullProtocol,omitempty" xml:"PullProtocol,omitempty" require:"true"`
	PullAuthType       *string `json:"PullAuthType,omitempty" xml:"PullAuthType,omitempty"`
	PullAuthKey        *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullArgs           *string `json:"PullArgs,omitempty" xml:"PullArgs,omitempty"`
	LiveapiRequestFrom *string `json:"LiveapiRequestFrom,omitempty" xml:"LiveapiRequestFrom,omitempty"`
}

func (s SetLiveLazyPullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLiveLazyPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetDomainName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetAppName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullDomainName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullDomainName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullAppName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullAppName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullProtocol(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullProtocol = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullAuthType(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullAuthType = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullAuthKey(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullAuthKey = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullArgs(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullArgs = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetLiveapiRequestFrom(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.LiveapiRequestFrom = &v
	return s
}

type SetLiveLazyPullStreamInfoConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetLiveLazyPullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLiveLazyPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveLazyPullStreamInfoConfigResponse) SetRequestId(v string) *SetLiveLazyPullStreamInfoConfigResponse {
	s.RequestId = &v
	return s
}

type UpdateCasterSceneAudioRequest struct {
	CasterId     *string                                    `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId      *string                                    `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	FollowEnable *int                                       `json:"FollowEnable,omitempty" xml:"FollowEnable,omitempty"`
	AudioLayer   []*UpdateCasterSceneAudioRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
	MixList      []*string                                  `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
}

func (s UpdateCasterSceneAudioRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCasterSceneAudioRequest) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioRequest) SetCasterId(v string) *UpdateCasterSceneAudioRequest {
	s.CasterId = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetSceneId(v string) *UpdateCasterSceneAudioRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetFollowEnable(v int) *UpdateCasterSceneAudioRequest {
	s.FollowEnable = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetAudioLayer(v []*UpdateCasterSceneAudioRequestAudioLayer) *UpdateCasterSceneAudioRequest {
	s.AudioLayer = v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetMixList(v []*string) *UpdateCasterSceneAudioRequest {
	s.MixList = v
	return s
}

type UpdateCasterSceneAudioRequestAudioLayer struct {
	VolumeRate         *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
	ValidChannel       *string  `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	FixedDelayDuration *int     `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
}

func (s UpdateCasterSceneAudioRequestAudioLayer) String() string {
	return tea.Prettify(s)
}

func (s UpdateCasterSceneAudioRequestAudioLayer) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetVolumeRate(v float32) *UpdateCasterSceneAudioRequestAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetValidChannel(v string) *UpdateCasterSceneAudioRequestAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetFixedDelayDuration(v int) *UpdateCasterSceneAudioRequestAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

type UpdateCasterSceneAudioResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateCasterSceneAudioResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCasterSceneAudioResponse) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioResponse) SetRequestId(v string) *UpdateCasterSceneAudioResponse {
	s.RequestId = &v
	return s
}

type SetCasterChannelRequest struct {
	CasterId   *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SeekOffset *int    `json:"SeekOffset,omitempty" xml:"SeekOffset,omitempty"`
	PlayStatus *int    `json:"PlayStatus,omitempty" xml:"PlayStatus,omitempty"`
	ReloadFlag *int    `json:"ReloadFlag,omitempty" xml:"ReloadFlag,omitempty"`
}

func (s SetCasterChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCasterChannelRequest) GoString() string {
	return s.String()
}

func (s *SetCasterChannelRequest) SetCasterId(v string) *SetCasterChannelRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterChannelRequest) SetChannelId(v string) *SetCasterChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *SetCasterChannelRequest) SetResourceId(v string) *SetCasterChannelRequest {
	s.ResourceId = &v
	return s
}

func (s *SetCasterChannelRequest) SetSeekOffset(v int) *SetCasterChannelRequest {
	s.SeekOffset = &v
	return s
}

func (s *SetCasterChannelRequest) SetPlayStatus(v int) *SetCasterChannelRequest {
	s.PlayStatus = &v
	return s
}

func (s *SetCasterChannelRequest) SetReloadFlag(v int) *SetCasterChannelRequest {
	s.ReloadFlag = &v
	return s
}

type SetCasterChannelResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetCasterChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCasterChannelResponse) GoString() string {
	return s.String()
}

func (s *SetCasterChannelResponse) SetRequestId(v string) *SetCasterChannelResponse {
	s.RequestId = &v
	return s
}

type DescribeCasterSceneAudioRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
}

func (s DescribeCasterSceneAudioRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSceneAudioRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioRequest) SetCasterId(v string) *DescribeCasterSceneAudioRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterSceneAudioRequest) SetSceneId(v string) *DescribeCasterSceneAudioRequest {
	s.SceneId = &v
	return s
}

type DescribeCasterSceneAudioResponse struct {
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId     *string                                      `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	FollowEnable *int                                         `json:"FollowEnable,omitempty" xml:"FollowEnable,omitempty" require:"true"`
	AudioLayers  *DescribeCasterSceneAudioResponseAudioLayers `json:"AudioLayers,omitempty" xml:"AudioLayers,omitempty" require:"true" type:"Struct"`
	MixList      *DescribeCasterSceneAudioResponseMixList     `json:"MixList,omitempty" xml:"MixList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterSceneAudioResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSceneAudioResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponse) SetRequestId(v string) *DescribeCasterSceneAudioResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterSceneAudioResponse) SetCasterId(v string) *DescribeCasterSceneAudioResponse {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterSceneAudioResponse) SetFollowEnable(v int) *DescribeCasterSceneAudioResponse {
	s.FollowEnable = &v
	return s
}

func (s *DescribeCasterSceneAudioResponse) SetAudioLayers(v *DescribeCasterSceneAudioResponseAudioLayers) *DescribeCasterSceneAudioResponse {
	s.AudioLayers = v
	return s
}

func (s *DescribeCasterSceneAudioResponse) SetMixList(v *DescribeCasterSceneAudioResponseMixList) *DescribeCasterSceneAudioResponse {
	s.MixList = v
	return s
}

type DescribeCasterSceneAudioResponseAudioLayers struct {
	AudioLayer []*DescribeCasterSceneAudioResponseAudioLayersAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterSceneAudioResponseAudioLayers) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseAudioLayers) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseAudioLayers) SetAudioLayer(v []*DescribeCasterSceneAudioResponseAudioLayersAudioLayer) *DescribeCasterSceneAudioResponseAudioLayers {
	s.AudioLayer = v
	return s
}

type DescribeCasterSceneAudioResponseAudioLayersAudioLayer struct {
	VolumeRate         *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty" require:"true"`
	ValidChannel       *string  `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty" require:"true"`
	FixedDelayDuration *int     `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty" require:"true"`
}

func (s DescribeCasterSceneAudioResponseAudioLayersAudioLayer) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseAudioLayersAudioLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseAudioLayersAudioLayer) SetVolumeRate(v float32) *DescribeCasterSceneAudioResponseAudioLayersAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseAudioLayersAudioLayer) SetValidChannel(v string) *DescribeCasterSceneAudioResponseAudioLayersAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseAudioLayersAudioLayer) SetFixedDelayDuration(v int) *DescribeCasterSceneAudioResponseAudioLayersAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

type DescribeCasterSceneAudioResponseMixList struct {
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterSceneAudioResponseMixList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseMixList) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseMixList) SetLocationId(v []*string) *DescribeCasterSceneAudioResponseMixList {
	s.LocationId = v
	return s
}

type DescribeCasterChannelsRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DescribeCasterChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsRequest) SetCasterId(v string) *DescribeCasterChannelsRequest {
	s.CasterId = &v
	return s
}

type DescribeCasterChannelsResponse struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int                                    `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Channels  *DescribeCasterChannelsResponseChannels `json:"Channels,omitempty" xml:"Channels,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponse) SetRequestId(v string) *DescribeCasterChannelsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterChannelsResponse) SetTotal(v int) *DescribeCasterChannelsResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterChannelsResponse) SetChannels(v *DescribeCasterChannelsResponseChannels) *DescribeCasterChannelsResponse {
	s.Channels = v
	return s
}

type DescribeCasterChannelsResponseChannels struct {
	Channel []*DescribeCasterChannelsResponseChannelsChannel `json:"Channel,omitempty" xml:"Channel,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterChannelsResponseChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterChannelsResponseChannels) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponseChannels) SetChannel(v []*DescribeCasterChannelsResponseChannelsChannel) *DescribeCasterChannelsResponseChannels {
	s.Channel = v
	return s
}

type DescribeCasterChannelsResponseChannelsChannel struct {
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	StreamUrl  *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	RtmpUrl    *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty" require:"true"`
}

func (s DescribeCasterChannelsResponseChannelsChannel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterChannelsResponseChannelsChannel) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponseChannelsChannel) SetChannelId(v string) *DescribeCasterChannelsResponseChannelsChannel {
	s.ChannelId = &v
	return s
}

func (s *DescribeCasterChannelsResponseChannelsChannel) SetResourceId(v string) *DescribeCasterChannelsResponseChannelsChannel {
	s.ResourceId = &v
	return s
}

func (s *DescribeCasterChannelsResponseChannelsChannel) SetStreamUrl(v string) *DescribeCasterChannelsResponseChannelsChannel {
	s.StreamUrl = &v
	return s
}

func (s *DescribeCasterChannelsResponseChannelsChannel) SetRtmpUrl(v string) *DescribeCasterChannelsResponseChannelsChannel {
	s.RtmpUrl = &v
	return s
}

type UpdateBoardRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardData *string `json:"BoardData,omitempty" xml:"BoardData,omitempty" require:"true"`
}

func (s UpdateBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBoardRequest) GoString() string {
	return s.String()
}

func (s *UpdateBoardRequest) SetAppId(v string) *UpdateBoardRequest {
	s.AppId = &v
	return s
}

func (s *UpdateBoardRequest) SetBoardData(v string) *UpdateBoardRequest {
	s.BoardData = &v
	return s
}

type UpdateBoardResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBoardResponse) GoString() string {
	return s.String()
}

func (s *UpdateBoardResponse) SetRequestId(v string) *UpdateBoardResponse {
	s.RequestId = &v
	return s
}

type JoinBoardRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppUid  *string `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
	BoardId *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s JoinBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinBoardRequest) GoString() string {
	return s.String()
}

func (s *JoinBoardRequest) SetAppId(v string) *JoinBoardRequest {
	s.AppId = &v
	return s
}

func (s *JoinBoardRequest) SetAppUid(v string) *JoinBoardRequest {
	s.AppUid = &v
	return s
}

func (s *JoinBoardRequest) SetBoardId(v string) *JoinBoardRequest {
	s.BoardId = &v
	return s
}

type JoinBoardResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
	BoardId           *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
	TopicId           *string `json:"TopicId,omitempty" xml:"TopicId,omitempty" require:"true"`
	KeepaliveTopic    *string `json:"KeepaliveTopic,omitempty" xml:"KeepaliveTopic,omitempty" require:"true"`
	KeepaliveInterval *int    `json:"KeepaliveInterval,omitempty" xml:"KeepaliveInterval,omitempty" require:"true"`
}

func (s JoinBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinBoardResponse) GoString() string {
	return s.String()
}

func (s *JoinBoardResponse) SetRequestId(v string) *JoinBoardResponse {
	s.RequestId = &v
	return s
}

func (s *JoinBoardResponse) SetToken(v string) *JoinBoardResponse {
	s.Token = &v
	return s
}

func (s *JoinBoardResponse) SetBoardId(v string) *JoinBoardResponse {
	s.BoardId = &v
	return s
}

func (s *JoinBoardResponse) SetTopicId(v string) *JoinBoardResponse {
	s.TopicId = &v
	return s
}

func (s *JoinBoardResponse) SetKeepaliveTopic(v string) *JoinBoardResponse {
	s.KeepaliveTopic = &v
	return s
}

func (s *JoinBoardResponse) SetKeepaliveInterval(v int) *JoinBoardResponse {
	s.KeepaliveInterval = &v
	return s
}

type DescribeBoardSnapshotRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardId *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s DescribeBoardSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotRequest) SetAppId(v string) *DescribeBoardSnapshotRequest {
	s.AppId = &v
	return s
}

func (s *DescribeBoardSnapshotRequest) SetBoardId(v string) *DescribeBoardSnapshotRequest {
	s.BoardId = &v
	return s
}

type DescribeBoardSnapshotResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Snapshot  *DescribeBoardSnapshotResponseSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" require:"true" type:"Struct"`
}

func (s DescribeBoardSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotResponse) SetRequestId(v string) *DescribeBoardSnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBoardSnapshotResponse) SetSnapshot(v *DescribeBoardSnapshotResponseSnapshot) *DescribeBoardSnapshotResponse {
	s.Snapshot = v
	return s
}

type DescribeBoardSnapshotResponseSnapshot struct {
	Board *DescribeBoardSnapshotResponseSnapshotBoard `json:"Board,omitempty" xml:"Board,omitempty" require:"true" type:"Struct"`
}

func (s DescribeBoardSnapshotResponseSnapshot) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotResponseSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotResponseSnapshot) SetBoard(v *DescribeBoardSnapshotResponseSnapshotBoard) *DescribeBoardSnapshotResponseSnapshot {
	s.Board = v
	return s
}

type DescribeBoardSnapshotResponseSnapshotBoard struct {
	BoardId         *string                                              `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
	AppUid          *string                                              `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
	EventTimestamp  *int64                                               `json:"EventTimestamp,omitempty" xml:"EventTimestamp,omitempty" require:"true"`
	CreateTimestamp *int64                                               `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty" require:"true"`
	UpdateTimestamp *int64                                               `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty" require:"true"`
	Pages           []*DescribeBoardSnapshotResponseSnapshotBoardPages   `json:"Pages,omitempty" xml:"Pages,omitempty" require:"true" type:"Repeated"`
	Configs         []*DescribeBoardSnapshotResponseSnapshotBoardConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBoardSnapshotResponseSnapshotBoard) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotResponseSnapshotBoard) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetBoardId(v string) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.BoardId = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetAppUid(v string) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.AppUid = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetEventTimestamp(v int64) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.EventTimestamp = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetCreateTimestamp(v int64) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetUpdateTimestamp(v int64) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetPages(v []*DescribeBoardSnapshotResponseSnapshotBoardPages) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.Pages = v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoard) SetConfigs(v []*DescribeBoardSnapshotResponseSnapshotBoardConfigs) *DescribeBoardSnapshotResponseSnapshotBoard {
	s.Configs = v
	return s
}

type DescribeBoardSnapshotResponseSnapshotBoardPages struct {
	PageIndex *int                                                       `json:"PageIndex,omitempty" xml:"PageIndex,omitempty" require:"true"`
	Elements  []*DescribeBoardSnapshotResponseSnapshotBoardPagesElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBoardSnapshotResponseSnapshotBoardPages) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotResponseSnapshotBoardPages) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPages) SetPageIndex(v int) *DescribeBoardSnapshotResponseSnapshotBoardPages {
	s.PageIndex = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPages) SetElements(v []*DescribeBoardSnapshotResponseSnapshotBoardPagesElements) *DescribeBoardSnapshotResponseSnapshotBoardPages {
	s.Elements = v
	return s
}

type DescribeBoardSnapshotResponseSnapshotBoardPagesElements struct {
	ElementIndex    *string `json:"ElementIndex,omitempty" xml:"ElementIndex,omitempty" require:"true"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty" require:"true"`
	ElementType     *int    `json:"ElementType,omitempty" xml:"ElementType,omitempty" require:"true"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty" require:"true"`
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s DescribeBoardSnapshotResponseSnapshotBoardPagesElements) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotResponseSnapshotBoardPagesElements) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPagesElements) SetElementIndex(v string) *DescribeBoardSnapshotResponseSnapshotBoardPagesElements {
	s.ElementIndex = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPagesElements) SetOwnerId(v string) *DescribeBoardSnapshotResponseSnapshotBoardPagesElements {
	s.OwnerId = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPagesElements) SetElementType(v int) *DescribeBoardSnapshotResponseSnapshotBoardPagesElements {
	s.ElementType = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPagesElements) SetUpdateTimestamp(v int64) *DescribeBoardSnapshotResponseSnapshotBoardPagesElements {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardPagesElements) SetData(v string) *DescribeBoardSnapshotResponseSnapshotBoardPagesElements {
	s.Data = &v
	return s
}

type DescribeBoardSnapshotResponseSnapshotBoardConfigs struct {
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s DescribeBoardSnapshotResponseSnapshotBoardConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardSnapshotResponseSnapshotBoardConfigs) GoString() string {
	return s.String()
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardConfigs) SetAppUid(v string) *DescribeBoardSnapshotResponseSnapshotBoardConfigs {
	s.AppUid = &v
	return s
}

func (s *DescribeBoardSnapshotResponseSnapshotBoardConfigs) SetData(v string) *DescribeBoardSnapshotResponseSnapshotBoardConfigs {
	s.Data = &v
	return s
}

type DescribeBoardsRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	PageNum  *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeBoardsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBoardsRequest) SetAppId(v string) *DescribeBoardsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeBoardsRequest) SetPageNum(v int) *DescribeBoardsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBoardsRequest) SetPageSize(v int) *DescribeBoardsRequest {
	s.PageSize = &v
	return s
}

type DescribeBoardsResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Boards    []*DescribeBoardsResponseBoards `json:"Boards,omitempty" xml:"Boards,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBoardsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBoardsResponse) SetRequestId(v string) *DescribeBoardsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBoardsResponse) SetBoards(v []*DescribeBoardsResponseBoards) *DescribeBoardsResponse {
	s.Boards = v
	return s
}

type DescribeBoardsResponseBoards struct {
	BoardId *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
	Topic   *string `json:"Topic,omitempty" xml:"Topic,omitempty" require:"true"`
	State   *int    `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s DescribeBoardsResponseBoards) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardsResponseBoards) GoString() string {
	return s.String()
}

func (s *DescribeBoardsResponseBoards) SetBoardId(v string) *DescribeBoardsResponseBoards {
	s.BoardId = &v
	return s
}

func (s *DescribeBoardsResponseBoards) SetTopic(v string) *DescribeBoardsResponseBoards {
	s.Topic = &v
	return s
}

func (s *DescribeBoardsResponseBoards) SetState(v int) *DescribeBoardsResponseBoards {
	s.State = &v
	return s
}

func (s *DescribeBoardsResponseBoards) SetUserId(v string) *DescribeBoardsResponseBoards {
	s.UserId = &v
	return s
}

type DescribeBoardEventsRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BoardId   *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s DescribeBoardEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBoardEventsRequest) SetAppId(v string) *DescribeBoardEventsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeBoardEventsRequest) SetStartTime(v string) *DescribeBoardEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBoardEventsRequest) SetEndTime(v string) *DescribeBoardEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBoardEventsRequest) SetBoardId(v string) *DescribeBoardEventsRequest {
	s.BoardId = &v
	return s
}

type DescribeBoardEventsResponse struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Events    []*DescribeBoardEventsResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBoardEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBoardEventsResponse) SetRequestId(v string) *DescribeBoardEventsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBoardEventsResponse) SetEvents(v []*DescribeBoardEventsResponseEvents) *DescribeBoardEventsResponse {
	s.Events = v
	return s
}

type DescribeBoardEventsResponseEvents struct {
	EventId   *int64  `json:"EventId,omitempty" xml:"EventId,omitempty" require:"true"`
	EventType *int    `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	UserId    *int    `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty" require:"true"`
}

func (s DescribeBoardEventsResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeBoardEventsResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeBoardEventsResponseEvents) SetEventId(v int64) *DescribeBoardEventsResponseEvents {
	s.EventId = &v
	return s
}

func (s *DescribeBoardEventsResponseEvents) SetEventType(v int) *DescribeBoardEventsResponseEvents {
	s.EventType = &v
	return s
}

func (s *DescribeBoardEventsResponseEvents) SetUserId(v int) *DescribeBoardEventsResponseEvents {
	s.UserId = &v
	return s
}

func (s *DescribeBoardEventsResponseEvents) SetData(v string) *DescribeBoardEventsResponseEvents {
	s.Data = &v
	return s
}

func (s *DescribeBoardEventsResponseEvents) SetTimestamp(v int64) *DescribeBoardEventsResponseEvents {
	s.Timestamp = &v
	return s
}

type DeleteBoardRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardId *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s DeleteBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBoardRequest) GoString() string {
	return s.String()
}

func (s *DeleteBoardRequest) SetAppId(v string) *DeleteBoardRequest {
	s.AppId = &v
	return s
}

func (s *DeleteBoardRequest) SetBoardId(v string) *DeleteBoardRequest {
	s.BoardId = &v
	return s
}

type DeleteBoardResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBoardResponse) GoString() string {
	return s.String()
}

func (s *DeleteBoardResponse) SetRequestId(v string) *DeleteBoardResponse {
	s.RequestId = &v
	return s
}

type CreateBoardRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
}

func (s CreateBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBoardRequest) GoString() string {
	return s.String()
}

func (s *CreateBoardRequest) SetAppId(v string) *CreateBoardRequest {
	s.AppId = &v
	return s
}

func (s *CreateBoardRequest) SetAppUid(v string) *CreateBoardRequest {
	s.AppUid = &v
	return s
}

type CreateBoardResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BoardId   *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s CreateBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBoardResponse) GoString() string {
	return s.String()
}

func (s *CreateBoardResponse) SetRequestId(v string) *CreateBoardResponse {
	s.RequestId = &v
	return s
}

func (s *CreateBoardResponse) SetBoardId(v string) *CreateBoardResponse {
	s.BoardId = &v
	return s
}

type CompleteBoardRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BoardId *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s CompleteBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteBoardRequest) GoString() string {
	return s.String()
}

func (s *CompleteBoardRequest) SetAppId(v string) *CompleteBoardRequest {
	s.AppId = &v
	return s
}

func (s *CompleteBoardRequest) SetBoardId(v string) *CompleteBoardRequest {
	s.BoardId = &v
	return s
}

type CompleteBoardResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CompleteBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteBoardResponse) GoString() string {
	return s.String()
}

func (s *CompleteBoardResponse) SetRequestId(v string) *CompleteBoardResponse {
	s.RequestId = &v
	return s
}

type ApplyBoardTokenRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppUid  *string `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
	BoardId *string `json:"BoardId,omitempty" xml:"BoardId,omitempty" require:"true"`
}

func (s ApplyBoardTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyBoardTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyBoardTokenRequest) SetAppId(v string) *ApplyBoardTokenRequest {
	s.AppId = &v
	return s
}

func (s *ApplyBoardTokenRequest) SetAppUid(v string) *ApplyBoardTokenRequest {
	s.AppUid = &v
	return s
}

func (s *ApplyBoardTokenRequest) SetBoardId(v string) *ApplyBoardTokenRequest {
	s.BoardId = &v
	return s
}

type ApplyBoardTokenResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
	Expired   *string `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
}

func (s ApplyBoardTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyBoardTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyBoardTokenResponse) SetRequestId(v string) *ApplyBoardTokenResponse {
	s.RequestId = &v
	return s
}

func (s *ApplyBoardTokenResponse) SetToken(v string) *ApplyBoardTokenResponse {
	s.Token = &v
	return s
}

func (s *ApplyBoardTokenResponse) SetExpired(v string) *ApplyBoardTokenResponse {
	s.Expired = &v
	return s
}

type DescribeLiveStreamCountRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveStreamCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountRequest) SetDomainName(v string) *DescribeLiveStreamCountRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveStreamCountResponse struct {
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	StreamCountInfos *DescribeLiveStreamCountResponseStreamCountInfos `json:"StreamCountInfos,omitempty" xml:"StreamCountInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponse) SetRequestId(v string) *DescribeLiveStreamCountResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamCountResponse) SetStreamCountInfos(v *DescribeLiveStreamCountResponseStreamCountInfos) *DescribeLiveStreamCountResponse {
	s.StreamCountInfos = v
	return s
}

type DescribeLiveStreamCountResponseStreamCountInfos struct {
	StreamCountInfo []*DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo `json:"StreamCountInfo,omitempty" xml:"StreamCountInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamCountResponseStreamCountInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamCountResponseStreamCountInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseStreamCountInfos) SetStreamCountInfo(v []*DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) *DescribeLiveStreamCountResponseStreamCountInfos {
	s.StreamCountInfo = v
	return s
}

type DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo struct {
	Count              *int64                                                                            `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	Limit              *int64                                                                            `json:"Limit,omitempty" xml:"Limit,omitempty" require:"true"`
	Type               *string                                                                           `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	StreamCountDetails *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails `json:"StreamCountDetails,omitempty" xml:"StreamCountDetails,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) SetCount(v int64) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo {
	s.Count = &v
	return s
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) SetLimit(v int64) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo {
	s.Limit = &v
	return s
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) SetType(v string) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo {
	s.Type = &v
	return s
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo) SetStreamCountDetails(v *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfo {
	s.StreamCountDetails = v
	return s
}

type DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails struct {
	StreamCountDetail []*DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail `json:"StreamCountDetail,omitempty" xml:"StreamCountDetail,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails) SetStreamCountDetail(v []*DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetails {
	s.StreamCountDetail = v
	return s
}

type DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail struct {
	Format        *string `json:"Format,omitempty" xml:"Format,omitempty" require:"true"`
	VideoDataRate *int64  `json:"VideoDataRate,omitempty" xml:"VideoDataRate,omitempty" require:"true"`
	Count         *int64  `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
}

func (s DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) SetFormat(v string) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	s.Format = &v
	return s
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) SetVideoDataRate(v int64) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	s.VideoDataRate = &v
	return s
}

func (s *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) SetCount(v int64) *DescribeLiveStreamCountResponseStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	s.Count = &v
	return s
}

type DescribeLiveCertificateDetailRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty" require:"true"`
}

func (s DescribeLiveCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateDetailRequest) SetSecurityToken(v string) *DescribeLiveCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveCertificateDetailRequest) SetCertName(v string) *DescribeLiveCertificateDetailRequest {
	s.CertName = &v
	return s
}

type DescribeLiveCertificateDetailResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty" require:"true"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty" require:"true"`
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty" require:"true"`
}

func (s DescribeLiveCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateDetailResponse) SetRequestId(v string) *DescribeLiveCertificateDetailResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) SetCert(v string) *DescribeLiveCertificateDetailResponse {
	s.Cert = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) SetKey(v string) *DescribeLiveCertificateDetailResponse {
	s.Key = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) SetCertId(v int64) *DescribeLiveCertificateDetailResponse {
	s.CertId = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponse) SetCertName(v string) *DescribeLiveCertificateDetailResponse {
	s.CertName = &v
	return s
}

type DescribeHlsLiveStreamRealTimeBpsDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) SetDomainName(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataRequest) SetTime(v string) *DescribeHlsLiveStreamRealTimeBpsDataRequest {
	s.Time = &v
	return s
}

type DescribeHlsLiveStreamRealTimeBpsDataResponse struct {
	Time      *string                                                  `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UsageData []*DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) SetTime(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponse {
	s.Time = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) SetRequestId(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) SetUsageData(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData) *DescribeHlsLiveStreamRealTimeBpsDataResponse {
	s.UsageData = v
	return s
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData struct {
	DomainName  *string                                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StreamInfos []*DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData) SetDomainName(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData {
	s.DomainName = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData) SetStreamInfos(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageData {
	s.StreamInfos = v
	return s
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos struct {
	StreamName *string                                                                  `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	Infos      []*DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos `json:"Infos,omitempty" xml:"Infos,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos) SetStreamName(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos {
	s.StreamName = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos) SetInfos(v []*DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfos {
	s.Infos = v
	return s
}

type DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos struct {
	DownFlow *float32 `json:"DownFlow,omitempty" xml:"DownFlow,omitempty" require:"true"`
	Rate     *string  `json:"Rate,omitempty" xml:"Rate,omitempty" require:"true"`
	Online   *float32 `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos) SetDownFlow(v float32) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos {
	s.DownFlow = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos) SetRate(v string) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos {
	s.Rate = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos) SetOnline(v float32) *DescribeHlsLiveStreamRealTimeBpsDataResponseUsageDataStreamInfosInfos {
	s.Online = &v
	return s
}

type StopLiveDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s StopLiveDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *StopLiveDomainRequest) SetSecurityToken(v string) *StopLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopLiveDomainRequest) SetDomainName(v string) *StopLiveDomainRequest {
	s.DomainName = &v
	return s
}

type StopLiveDomainResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopLiveDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *StopLiveDomainResponse) SetRequestId(v string) *StopLiveDomainResponse {
	s.RequestId = &v
	return s
}

type StartLiveDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s StartLiveDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StartLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *StartLiveDomainRequest) SetSecurityToken(v string) *StartLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartLiveDomainRequest) SetDomainName(v string) *StartLiveDomainRequest {
	s.DomainName = &v
	return s
}

type StartLiveDomainResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StartLiveDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *StartLiveDomainResponse) SetRequestId(v string) *StartLiveDomainResponse {
	s.RequestId = &v
	return s
}

type SetLiveDomainCertificateRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType      *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty" require:"true"`
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SSLPri        *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	ForceSet      *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
}

func (s SetLiveDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLiveDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetLiveDomainCertificateRequest) SetSecurityToken(v string) *SetLiveDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetDomainName(v string) *SetLiveDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetCertName(v string) *SetLiveDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetCertType(v string) *SetLiveDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSSLProtocol(v string) *SetLiveDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSSLPub(v string) *SetLiveDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSSLPri(v string) *SetLiveDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetForceSet(v string) *SetLiveDomainCertificateRequest {
	s.ForceSet = &v
	return s
}

type SetLiveDomainCertificateResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetLiveDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLiveDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetLiveDomainCertificateResponse) SetRequestId(v string) *SetLiveDomainCertificateResponse {
	s.RequestId = &v
	return s
}

type BatchSetLiveDomainConfigsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" require:"true"`
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty" require:"true"`
}

func (s BatchSetLiveDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetLiveDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetLiveDomainConfigsRequest) SetSecurityToken(v string) *BatchSetLiveDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) SetDomainNames(v string) *BatchSetLiveDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) SetFunctions(v string) *BatchSetLiveDomainConfigsRequest {
	s.Functions = &v
	return s
}

type BatchSetLiveDomainConfigsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BatchSetLiveDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetLiveDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetLiveDomainConfigsResponse) SetRequestId(v string) *BatchSetLiveDomainConfigsResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveCertificateListRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeLiveCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListRequest) SetSecurityToken(v string) *DescribeLiveCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveCertificateListRequest) SetDomainName(v string) *DescribeLiveCertificateListRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveCertificateListResponse struct {
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CertificateListModel *DescribeLiveCertificateListResponseCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponse) SetRequestId(v string) *DescribeLiveCertificateListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCertificateListResponse) SetCertificateListModel(v *DescribeLiveCertificateListResponseCertificateListModel) *DescribeLiveCertificateListResponse {
	s.CertificateListModel = v
	return s
}

type DescribeLiveCertificateListResponseCertificateListModel struct {
	Count    *int                                                             `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	CertList *DescribeLiveCertificateListResponseCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveCertificateListResponseCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateListResponseCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseCertificateListModel) SetCount(v int) *DescribeLiveCertificateListResponseCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeLiveCertificateListResponseCertificateListModel) SetCertList(v *DescribeLiveCertificateListResponseCertificateListModelCertList) *DescribeLiveCertificateListResponseCertificateListModel {
	s.CertList = v
	return s
}

type DescribeLiveCertificateListResponseCertificateListModelCertList struct {
	Cert []*DescribeLiveCertificateListResponseCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveCertificateListResponseCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateListResponseCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertList) SetCert(v []*DescribeLiveCertificateListResponseCertificateListModelCertListCert) *DescribeLiveCertificateListResponseCertificateListModelCertList {
	s.Cert = v
	return s
}

type DescribeLiveCertificateListResponseCertificateListModelCertListCert struct {
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty" require:"true"`
	CertId      *int64  `json:"CertId,omitempty" xml:"CertId,omitempty" require:"true"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty" require:"true"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty" require:"true"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty" require:"true"`
	LastTime    *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty" require:"true"`
}

func (s DescribeLiveCertificateListResponseCertificateListModelCertListCert) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveCertificateListResponseCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertListCert) SetCertName(v string) *DescribeLiveCertificateListResponseCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertListCert) SetCertId(v int64) *DescribeLiveCertificateListResponseCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertListCert) SetFingerprint(v string) *DescribeLiveCertificateListResponseCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertListCert) SetCommon(v string) *DescribeLiveCertificateListResponseCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertListCert) SetIssuer(v string) *DescribeLiveCertificateListResponseCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeLiveCertificateListResponseCertificateListModelCertListCert) SetLastTime(v int64) *DescribeLiveCertificateListResponseCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

type DeleteLiveDomainRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DeleteLiveDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainRequest) SetSecurityToken(v string) *DeleteLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveDomainRequest) SetDomainName(v string) *DeleteLiveDomainRequest {
	s.DomainName = &v
	return s
}

type DeleteLiveDomainResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainResponse) SetRequestId(v string) *DeleteLiveDomainResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainConfigsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty" require:"true"`
}

func (s DescribeLiveDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsRequest) SetSecurityToken(v string) *DescribeLiveDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDomainConfigsRequest) SetDomainName(v string) *DescribeLiveDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainConfigsRequest) SetFunctionNames(v string) *DescribeLiveDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

type DescribeLiveDomainConfigsResponse struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainConfigs *DescribeLiveDomainConfigsResponseDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponse) SetRequestId(v string) *DescribeLiveDomainConfigsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponse) SetDomainConfigs(v *DescribeLiveDomainConfigsResponseDomainConfigs) *DescribeLiveDomainConfigsResponse {
	s.DomainConfigs = v
	return s
}

type DescribeLiveDomainConfigsResponseDomainConfigs struct {
	DomainConfig []*DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainConfigsResponseDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigs) SetDomainConfig(v []*DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) *DescribeLiveDomainConfigsResponseDomainConfigs {
	s.DomainConfig = v
	return s
}

type DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig struct {
	FunctionName *string                                                                 `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" require:"true"`
	ConfigId     *string                                                                 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty" require:"true"`
	Status       *string                                                                 `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	FunctionArgs *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) SetConfigId(v string) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) SetStatus(v string) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

type DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty" require:"true"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty" require:"true"`
}

func (s DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeLiveDomainConfigsResponseDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type AddLiveDomainRequest struct {
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty" require:"true"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	CheckUrl       *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	Scope          *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddLiveDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDomainRequest) SetSecurityToken(v string) *AddLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveDomainRequest) SetLiveDomainType(v string) *AddLiveDomainRequest {
	s.LiveDomainType = &v
	return s
}

func (s *AddLiveDomainRequest) SetDomainName(v string) *AddLiveDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveDomainRequest) SetRegion(v string) *AddLiveDomainRequest {
	s.Region = &v
	return s
}

func (s *AddLiveDomainRequest) SetCheckUrl(v string) *AddLiveDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddLiveDomainRequest) SetScope(v string) *AddLiveDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddLiveDomainRequest) SetTopLevelDomain(v string) *AddLiveDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type AddLiveDomainResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDomainResponse) SetRequestId(v string) *AddLiveDomainResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainDetailRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailRequest) SetSecurityToken(v string) *DescribeLiveDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDomainDetailRequest) SetDomainName(v string) *DescribeLiveDomainDetailRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveDomainDetailResponse struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainDetail *DescribeLiveDomainDetailResponseDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailResponse) SetRequestId(v string) *DescribeLiveDomainDetailResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainDetailResponse) SetDomainDetail(v *DescribeLiveDomainDetailResponseDomainDetail) *DescribeLiveDomainDetailResponse {
	s.DomainDetail = v
	return s
}

type DescribeLiveDomainDetailResponseDomainDetail struct {
	GmtCreated     *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty" require:"true"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
	DomainStatus   *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty" require:"true"`
	Cname          *string `json:"Cname,omitempty" xml:"Cname,omitempty" require:"true"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty" require:"true"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	SSLProtocol    *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty" require:"true"`
	SSLPub         *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty" require:"true"`
	Scope          *string `json:"Scope,omitempty" xml:"Scope,omitempty" require:"true"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty" require:"true"`
}

func (s DescribeLiveDomainDetailResponseDomainDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainDetailResponseDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetGmtCreated(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetGmtModified(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetDomainStatus(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetCname(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetDomainName(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetLiveDomainType(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.LiveDomainType = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetRegion(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetDescription(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetSSLProtocol(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetSSLPub(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetScope(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseDomainDetail) SetCertName(v string) *DescribeLiveDomainDetailResponseDomainDetail {
	s.CertName = &v
	return s
}

type BatchDeleteLiveDomainConfigsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" require:"true"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty" require:"true"`
}

func (s BatchDeleteLiveDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteLiveDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetSecurityToken(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

type BatchDeleteLiveDomainConfigsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BatchDeleteLiveDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteLiveDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteLiveDomainConfigsResponse) SetRequestId(v string) *BatchDeleteLiveDomainConfigsResponse {
	s.RequestId = &v
	return s
}

type DescribeRoomStatusRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
}

func (s DescribeRoomStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoomStatusRequest) SetAppId(v string) *DescribeRoomStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRoomStatusRequest) SetRoomId(v string) *DescribeRoomStatusRequest {
	s.RoomId = &v
	return s
}

type DescribeRoomStatusResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RoomStatus *int    `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty" require:"true"`
}

func (s DescribeRoomStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoomStatusResponse) SetRequestId(v string) *DescribeRoomStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRoomStatusResponse) SetRoomStatus(v int) *DescribeRoomStatusResponse {
	s.RoomStatus = &v
	return s
}

type DescribeRoomListRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId     *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	AnchorId   *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	RoomStatus *int    `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNum    *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRoomListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoomListRequest) SetAppId(v string) *DescribeRoomListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRoomListRequest) SetRoomId(v string) *DescribeRoomListRequest {
	s.RoomId = &v
	return s
}

func (s *DescribeRoomListRequest) SetAnchorId(v string) *DescribeRoomListRequest {
	s.AnchorId = &v
	return s
}

func (s *DescribeRoomListRequest) SetRoomStatus(v int) *DescribeRoomListRequest {
	s.RoomStatus = &v
	return s
}

func (s *DescribeRoomListRequest) SetStartTime(v string) *DescribeRoomListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRoomListRequest) SetEndTime(v string) *DescribeRoomListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRoomListRequest) SetOrder(v string) *DescribeRoomListRequest {
	s.Order = &v
	return s
}

func (s *DescribeRoomListRequest) SetPageNum(v int) *DescribeRoomListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRoomListRequest) SetPageSize(v int) *DescribeRoomListRequest {
	s.PageSize = &v
	return s
}

type DescribeRoomListResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalNum  *int                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage *int                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	RoomList  []*DescribeRoomListResponseRoomList `json:"RoomList,omitempty" xml:"RoomList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRoomListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoomListResponse) SetRequestId(v string) *DescribeRoomListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRoomListResponse) SetTotalNum(v int) *DescribeRoomListResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeRoomListResponse) SetTotalPage(v int) *DescribeRoomListResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeRoomListResponse) SetRoomList(v []*DescribeRoomListResponseRoomList) *DescribeRoomListResponse {
	s.RoomList = v
	return s
}

type DescribeRoomListResponseRoomList struct {
	RoomId       *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	AnchorId     *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty" require:"true"`
	RoomStatus   *int    `json:"RoomStatus,omitempty" xml:"RoomStatus,omitempty" require:"true"`
	ForbidStream *string `json:"ForbidStream,omitempty" xml:"ForbidStream,omitempty" require:"true"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s DescribeRoomListResponseRoomList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomListResponseRoomList) GoString() string {
	return s.String()
}

func (s *DescribeRoomListResponseRoomList) SetRoomId(v string) *DescribeRoomListResponseRoomList {
	s.RoomId = &v
	return s
}

func (s *DescribeRoomListResponseRoomList) SetAnchorId(v string) *DescribeRoomListResponseRoomList {
	s.AnchorId = &v
	return s
}

func (s *DescribeRoomListResponseRoomList) SetRoomStatus(v int) *DescribeRoomListResponseRoomList {
	s.RoomStatus = &v
	return s
}

func (s *DescribeRoomListResponseRoomList) SetForbidStream(v string) *DescribeRoomListResponseRoomList {
	s.ForbidStream = &v
	return s
}

func (s *DescribeRoomListResponseRoomList) SetCreateTime(v string) *DescribeRoomListResponseRoomList {
	s.CreateTime = &v
	return s
}

type DescribeRoomKickoutUserListRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId   *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	Order    *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNum  *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRoomKickoutUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomKickoutUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoomKickoutUserListRequest) SetAppId(v string) *DescribeRoomKickoutUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRoomKickoutUserListRequest) SetRoomId(v string) *DescribeRoomKickoutUserListRequest {
	s.RoomId = &v
	return s
}

func (s *DescribeRoomKickoutUserListRequest) SetOrder(v string) *DescribeRoomKickoutUserListRequest {
	s.Order = &v
	return s
}

func (s *DescribeRoomKickoutUserListRequest) SetPageNum(v int) *DescribeRoomKickoutUserListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRoomKickoutUserListRequest) SetPageSize(v int) *DescribeRoomKickoutUserListRequest {
	s.PageSize = &v
	return s
}

type DescribeRoomKickoutUserListResponse struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalNum  *int                                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage *int                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	UserList  []*DescribeRoomKickoutUserListResponseUserList `json:"UserList,omitempty" xml:"UserList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRoomKickoutUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomKickoutUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoomKickoutUserListResponse) SetRequestId(v string) *DescribeRoomKickoutUserListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRoomKickoutUserListResponse) SetTotalNum(v int) *DescribeRoomKickoutUserListResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeRoomKickoutUserListResponse) SetTotalPage(v int) *DescribeRoomKickoutUserListResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeRoomKickoutUserListResponse) SetUserList(v []*DescribeRoomKickoutUserListResponseUserList) *DescribeRoomKickoutUserListResponse {
	s.UserList = v
	return s
}

type DescribeRoomKickoutUserListResponseUserList struct {
	AppUid      *string `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
	OpStartTime *string `json:"OpStartTime,omitempty" xml:"OpStartTime,omitempty" require:"true"`
	OpEndTime   *string `json:"OpEndTime,omitempty" xml:"OpEndTime,omitempty" require:"true"`
}

func (s DescribeRoomKickoutUserListResponseUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoomKickoutUserListResponseUserList) GoString() string {
	return s.String()
}

func (s *DescribeRoomKickoutUserListResponseUserList) SetAppUid(v string) *DescribeRoomKickoutUserListResponseUserList {
	s.AppUid = &v
	return s
}

func (s *DescribeRoomKickoutUserListResponseUserList) SetOpStartTime(v string) *DescribeRoomKickoutUserListResponseUserList {
	s.OpStartTime = &v
	return s
}

func (s *DescribeRoomKickoutUserListResponseUserList) SetOpEndTime(v string) *DescribeRoomKickoutUserListResponseUserList {
	s.OpEndTime = &v
	return s
}

type SendRoomUserNotificationRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId   *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	AppUid   *string `json:"AppUid,omitempty" xml:"AppUid,omitempty" require:"true"`
	ToAppUid *string `json:"ToAppUid,omitempty" xml:"ToAppUid,omitempty" require:"true"`
	Data     *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Priority *int    `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SendRoomUserNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRoomUserNotificationRequest) GoString() string {
	return s.String()
}

func (s *SendRoomUserNotificationRequest) SetAppId(v string) *SendRoomUserNotificationRequest {
	s.AppId = &v
	return s
}

func (s *SendRoomUserNotificationRequest) SetRoomId(v string) *SendRoomUserNotificationRequest {
	s.RoomId = &v
	return s
}

func (s *SendRoomUserNotificationRequest) SetAppUid(v string) *SendRoomUserNotificationRequest {
	s.AppUid = &v
	return s
}

func (s *SendRoomUserNotificationRequest) SetToAppUid(v string) *SendRoomUserNotificationRequest {
	s.ToAppUid = &v
	return s
}

func (s *SendRoomUserNotificationRequest) SetData(v string) *SendRoomUserNotificationRequest {
	s.Data = &v
	return s
}

func (s *SendRoomUserNotificationRequest) SetPriority(v int) *SendRoomUserNotificationRequest {
	s.Priority = &v
	return s
}

type SendRoomUserNotificationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty" require:"true"`
}

func (s SendRoomUserNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRoomUserNotificationResponse) GoString() string {
	return s.String()
}

func (s *SendRoomUserNotificationResponse) SetRequestId(v string) *SendRoomUserNotificationResponse {
	s.RequestId = &v
	return s
}

func (s *SendRoomUserNotificationResponse) SetMessageId(v string) *SendRoomUserNotificationResponse {
	s.MessageId = &v
	return s
}

type DescribeForbidPushStreamRoomListRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Order    *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNum  *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeForbidPushStreamRoomListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeForbidPushStreamRoomListRequest) GoString() string {
	return s.String()
}

func (s *DescribeForbidPushStreamRoomListRequest) SetAppId(v string) *DescribeForbidPushStreamRoomListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListRequest) SetOrder(v string) *DescribeForbidPushStreamRoomListRequest {
	s.Order = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListRequest) SetPageNum(v int) *DescribeForbidPushStreamRoomListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListRequest) SetPageSize(v int) *DescribeForbidPushStreamRoomListRequest {
	s.PageSize = &v
	return s
}

type DescribeForbidPushStreamRoomListResponse struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalNum  *int                                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage *int                                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	RoomList  []*DescribeForbidPushStreamRoomListResponseRoomList `json:"RoomList,omitempty" xml:"RoomList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeForbidPushStreamRoomListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeForbidPushStreamRoomListResponse) GoString() string {
	return s.String()
}

func (s *DescribeForbidPushStreamRoomListResponse) SetRequestId(v string) *DescribeForbidPushStreamRoomListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListResponse) SetTotalNum(v int) *DescribeForbidPushStreamRoomListResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListResponse) SetTotalPage(v int) *DescribeForbidPushStreamRoomListResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListResponse) SetRoomList(v []*DescribeForbidPushStreamRoomListResponseRoomList) *DescribeForbidPushStreamRoomListResponse {
	s.RoomList = v
	return s
}

type DescribeForbidPushStreamRoomListResponseRoomList struct {
	RoomId      *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	AnchorId    *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty" require:"true"`
	OpStartTime *string `json:"OpStartTime,omitempty" xml:"OpStartTime,omitempty" require:"true"`
	OpEndTime   *string `json:"OpEndTime,omitempty" xml:"OpEndTime,omitempty" require:"true"`
}

func (s DescribeForbidPushStreamRoomListResponseRoomList) String() string {
	return tea.Prettify(s)
}

func (s DescribeForbidPushStreamRoomListResponseRoomList) GoString() string {
	return s.String()
}

func (s *DescribeForbidPushStreamRoomListResponseRoomList) SetRoomId(v string) *DescribeForbidPushStreamRoomListResponseRoomList {
	s.RoomId = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListResponseRoomList) SetAnchorId(v string) *DescribeForbidPushStreamRoomListResponseRoomList {
	s.AnchorId = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListResponseRoomList) SetOpStartTime(v string) *DescribeForbidPushStreamRoomListResponseRoomList {
	s.OpStartTime = &v
	return s
}

func (s *DescribeForbidPushStreamRoomListResponseRoomList) SetOpEndTime(v string) *DescribeForbidPushStreamRoomListResponseRoomList {
	s.OpEndTime = &v
	return s
}

type SendRoomNotificationRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId   *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	AppUid   *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	Data     *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Priority *int    `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SendRoomNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRoomNotificationRequest) GoString() string {
	return s.String()
}

func (s *SendRoomNotificationRequest) SetAppId(v string) *SendRoomNotificationRequest {
	s.AppId = &v
	return s
}

func (s *SendRoomNotificationRequest) SetRoomId(v string) *SendRoomNotificationRequest {
	s.RoomId = &v
	return s
}

func (s *SendRoomNotificationRequest) SetAppUid(v string) *SendRoomNotificationRequest {
	s.AppUid = &v
	return s
}

func (s *SendRoomNotificationRequest) SetData(v string) *SendRoomNotificationRequest {
	s.Data = &v
	return s
}

func (s *SendRoomNotificationRequest) SetPriority(v int) *SendRoomNotificationRequest {
	s.Priority = &v
	return s
}

type SendRoomNotificationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty" require:"true"`
}

func (s SendRoomNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRoomNotificationResponse) GoString() string {
	return s.String()
}

func (s *SendRoomNotificationResponse) SetRequestId(v string) *SendRoomNotificationResponse {
	s.RequestId = &v
	return s
}

func (s *SendRoomNotificationResponse) SetMessageId(v string) *SendRoomNotificationResponse {
	s.MessageId = &v
	return s
}

type ForbidPushStreamRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId   *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ForbidPushStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ForbidPushStreamRequest) GoString() string {
	return s.String()
}

func (s *ForbidPushStreamRequest) SetAppId(v string) *ForbidPushStreamRequest {
	s.AppId = &v
	return s
}

func (s *ForbidPushStreamRequest) SetRoomId(v string) *ForbidPushStreamRequest {
	s.RoomId = &v
	return s
}

func (s *ForbidPushStreamRequest) SetUserData(v string) *ForbidPushStreamRequest {
	s.UserData = &v
	return s
}

func (s *ForbidPushStreamRequest) SetEndTime(v string) *ForbidPushStreamRequest {
	s.EndTime = &v
	return s
}

type ForbidPushStreamResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ForbidPushStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ForbidPushStreamResponse) GoString() string {
	return s.String()
}

func (s *ForbidPushStreamResponse) SetRequestId(v string) *ForbidPushStreamResponse {
	s.RequestId = &v
	return s
}

type DeleteRoomRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
}

func (s DeleteRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoomRequest) SetAppId(v string) *DeleteRoomRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRoomRequest) SetRoomId(v string) *DeleteRoomRequest {
	s.RoomId = &v
	return s
}

type DeleteRoomResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoomResponse) SetRequestId(v string) *DeleteRoomResponse {
	s.RequestId = &v
	return s
}

type CreateRoomRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId          *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	AnchorId        *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty" require:"true"`
	TemplateIds     *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
	UseAppTranscode *bool   `json:"UseAppTranscode,omitempty" xml:"UseAppTranscode,omitempty"`
}

func (s CreateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequest) SetAppId(v string) *CreateRoomRequest {
	s.AppId = &v
	return s
}

func (s *CreateRoomRequest) SetRoomId(v string) *CreateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomRequest) SetAnchorId(v string) *CreateRoomRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateRoomRequest) SetTemplateIds(v string) *CreateRoomRequest {
	s.TemplateIds = &v
	return s
}

func (s *CreateRoomRequest) SetUseAppTranscode(v bool) *CreateRoomRequest {
	s.UseAppTranscode = &v
	return s
}

type CreateRoomResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId    *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	AnchorId  *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty" require:"true"`
}

func (s CreateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateRoomResponse) SetRequestId(v string) *CreateRoomResponse {
	s.RequestId = &v
	return s
}

func (s *CreateRoomResponse) SetAppId(v string) *CreateRoomResponse {
	s.AppId = &v
	return s
}

func (s *CreateRoomResponse) SetRoomId(v string) *CreateRoomResponse {
	s.RoomId = &v
	return s
}

func (s *CreateRoomResponse) SetAnchorId(v string) *CreateRoomResponse {
	s.AnchorId = &v
	return s
}

type AllowPushStreamRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
}

func (s AllowPushStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s AllowPushStreamRequest) GoString() string {
	return s.String()
}

func (s *AllowPushStreamRequest) SetAppId(v string) *AllowPushStreamRequest {
	s.AppId = &v
	return s
}

func (s *AllowPushStreamRequest) SetRoomId(v string) *AllowPushStreamRequest {
	s.RoomId = &v
	return s
}

type AllowPushStreamResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AllowPushStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s AllowPushStreamResponse) GoString() string {
	return s.String()
}

func (s *AllowPushStreamResponse) SetRequestId(v string) *AllowPushStreamResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveUserDomainsRequest struct {
	SecurityToken    *string                              `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	LiveDomainType   *string                              `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
	PageSize         *int                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DomainName       *string                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RegionName       *string                              `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	DomainSearchType *string                              `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	DomainStatus     *string                              `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	Tag              []*DescribeLiveUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsRequest) SetSecurityToken(v string) *DescribeLiveUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetLiveDomainType(v string) *DescribeLiveUserDomainsRequest {
	s.LiveDomainType = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetPageSize(v int) *DescribeLiveUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetPageNumber(v int) *DescribeLiveUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetDomainName(v string) *DescribeLiveUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetRegionName(v string) *DescribeLiveUserDomainsRequest {
	s.RegionName = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetDomainSearchType(v string) *DescribeLiveUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetDomainStatus(v string) *DescribeLiveUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeLiveUserDomainsRequest) SetTag(v []*DescribeLiveUserDomainsRequestTag) *DescribeLiveUserDomainsRequest {
	s.Tag = v
	return s
}

type DescribeLiveUserDomainsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUserDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsRequestTag) SetKey(v string) *DescribeLiveUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeLiveUserDomainsRequestTag) SetValue(v string) *DescribeLiveUserDomainsRequestTag {
	s.Value = &v
	return s
}

type DescribeLiveUserDomainsResponse struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Domains    *DescribeLiveUserDomainsResponseDomains `json:"Domains,omitempty" xml:"Domains,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponse) SetRequestId(v string) *DescribeLiveUserDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserDomainsResponse) SetPageNumber(v int64) *DescribeLiveUserDomainsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserDomainsResponse) SetPageSize(v int64) *DescribeLiveUserDomainsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserDomainsResponse) SetTotalCount(v int64) *DescribeLiveUserDomainsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeLiveUserDomainsResponse) SetDomains(v *DescribeLiveUserDomainsResponseDomains) *DescribeLiveUserDomainsResponse {
	s.Domains = v
	return s
}

type DescribeLiveUserDomainsResponseDomains struct {
	PageData []*DescribeLiveUserDomainsResponseDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveUserDomainsResponseDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserDomainsResponseDomains) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponseDomains) SetPageData(v []*DescribeLiveUserDomainsResponseDomainsPageData) *DescribeLiveUserDomainsResponseDomains {
	s.PageData = v
	return s
}

type DescribeLiveUserDomainsResponseDomainsPageData struct {
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Cname            *string `json:"Cname,omitempty" xml:"Cname,omitempty" require:"true"`
	LiveDomainType   *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty" require:"true"`
	GmtCreated       *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty" require:"true"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	LiveDomainStatus *string `json:"LiveDomainStatus,omitempty" xml:"LiveDomainStatus,omitempty" require:"true"`
	RegionName       *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
}

func (s DescribeLiveUserDomainsResponseDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveUserDomainsResponseDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetDomainName(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetCname(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetLiveDomainType(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.LiveDomainType = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetGmtCreated(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetGmtModified(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetDescription(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetLiveDomainStatus(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.LiveDomainStatus = &v
	return s
}

func (s *DescribeLiveUserDomainsResponseDomainsPageData) SetRegionName(v string) *DescribeLiveUserDomainsResponseDomainsPageData {
	s.RegionName = &v
	return s
}

type DescribeCasterRtcInfoRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DescribeCasterRtcInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterRtcInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterRtcInfoRequest) SetCasterId(v string) *DescribeCasterRtcInfoRequest {
	s.CasterId = &v
	return s
}

type DescribeCasterRtcInfoResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty" require:"true"`
}

func (s DescribeCasterRtcInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterRtcInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterRtcInfoResponse) SetRequestId(v string) *DescribeCasterRtcInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterRtcInfoResponse) SetCasterId(v string) *DescribeCasterRtcInfoResponse {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterRtcInfoResponse) SetAuthToken(v string) *DescribeCasterRtcInfoResponse {
	s.AuthToken = &v
	return s
}

type DescribeUpBpsPeakDataRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeUpBpsPeakDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataRequest) SetStartTime(v string) *DescribeUpBpsPeakDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetEndTime(v string) *DescribeUpBpsPeakDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetDomainSwitch(v string) *DescribeUpBpsPeakDataRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeUpBpsPeakDataRequest) SetDomainName(v string) *DescribeUpBpsPeakDataRequest {
	s.DomainName = &v
	return s
}

type DescribeUpBpsPeakDataResponse struct {
	RequestId              *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DescribeUpPeakTraffics *DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics `json:"DescribeUpPeakTraffics,omitempty" xml:"DescribeUpPeakTraffics,omitempty" require:"true" type:"Struct"`
}

func (s DescribeUpBpsPeakDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponse) SetRequestId(v string) *DescribeUpBpsPeakDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponse) SetDescribeUpPeakTraffics(v *DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics) *DescribeUpBpsPeakDataResponse {
	s.DescribeUpPeakTraffics = v
	return s
}

type DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics struct {
	DescribeUpPeakTraffic []*DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic `json:"DescribeUpPeakTraffic,omitempty" xml:"DescribeUpPeakTraffic,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics) SetDescribeUpPeakTraffic(v []*DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) *DescribeUpBpsPeakDataResponseDescribeUpPeakTraffics {
	s.DescribeUpPeakTraffic = v
	return s
}

type DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic struct {
	PeakTime  *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty" require:"true"`
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty" require:"true"`
	StatName  *string `json:"StatName,omitempty" xml:"StatName,omitempty" require:"true"`
	BandWidth *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty" require:"true"`
}

func (s DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetPeakTime(v string) *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.PeakTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetQueryTime(v string) *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.QueryTime = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetStatName(v string) *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.StatName = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic) SetBandWidth(v string) *DescribeUpBpsPeakDataResponseDescribeUpPeakTrafficsDescribeUpPeakTraffic {
	s.BandWidth = &v
	return s
}

type DescribeUpBpsPeakOfLineRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty" require:"true"`
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeUpBpsPeakOfLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineRequest) SetStartTime(v string) *DescribeUpBpsPeakOfLineRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetEndTime(v string) *DescribeUpBpsPeakOfLineRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetLine(v string) *DescribeUpBpsPeakOfLineRequest {
	s.Line = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetDomainSwitch(v string) *DescribeUpBpsPeakOfLineRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineRequest) SetDomainName(v string) *DescribeUpBpsPeakOfLineRequest {
	s.DomainName = &v
	return s
}

type DescribeUpBpsPeakOfLineResponse struct {
	RequestId                *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DescribeUpBpsPeakOfLines *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines `json:"DescribeUpBpsPeakOfLines,omitempty" xml:"DescribeUpBpsPeakOfLines,omitempty" require:"true" type:"Struct"`
}

func (s DescribeUpBpsPeakOfLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponse) SetRequestId(v string) *DescribeUpBpsPeakOfLineResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponse) SetDescribeUpBpsPeakOfLines(v *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines) *DescribeUpBpsPeakOfLineResponse {
	s.DescribeUpBpsPeakOfLines = v
	return s
}

type DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines struct {
	DescribeUpBpsPeakOfLine []*DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine `json:"DescribeUpBpsPeakOfLine,omitempty" xml:"DescribeUpBpsPeakOfLine,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines) SetDescribeUpBpsPeakOfLine(v []*DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLines {
	s.DescribeUpBpsPeakOfLine = v
	return s
}

type DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine struct {
	BandWidth *float32 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty" require:"true"`
	PeakTime  *string  `json:"PeakTime,omitempty" xml:"PeakTime,omitempty" require:"true"`
	QueryTime *string  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty" require:"true"`
	StatName  *string  `json:"StatName,omitempty" xml:"StatName,omitempty" require:"true"`
}

func (s DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetBandWidth(v float32) *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.BandWidth = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetPeakTime(v string) *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.PeakTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetQueryTime(v string) *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.QueryTime = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine) SetStatName(v string) *DescribeUpBpsPeakOfLineResponseDescribeUpBpsPeakOfLinesDescribeUpBpsPeakOfLine {
	s.StatName = &v
	return s
}

type DescribeUpPeakPublishStreamDataRequest struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DomainSwitch *string `json:"DomainSwitch,omitempty" xml:"DomainSwitch,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeUpPeakPublishStreamDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetStartTime(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetEndTime(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetDomainSwitch(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.DomainSwitch = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataRequest) SetDomainName(v string) *DescribeUpPeakPublishStreamDataRequest {
	s.DomainName = &v
	return s
}

type DescribeUpPeakPublishStreamDataResponse struct {
	RequestId                        *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DescribeUpPeakPublishStreamDatas *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas `json:"DescribeUpPeakPublishStreamDatas,omitempty" xml:"DescribeUpPeakPublishStreamDatas,omitempty" require:"true" type:"Struct"`
}

func (s DescribeUpPeakPublishStreamDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponse) SetRequestId(v string) *DescribeUpPeakPublishStreamDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponse) SetDescribeUpPeakPublishStreamDatas(v *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas) *DescribeUpPeakPublishStreamDataResponse {
	s.DescribeUpPeakPublishStreamDatas = v
	return s
}

type DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas struct {
	DescribeUpPeakPublishStreamData []*DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData `json:"DescribeUpPeakPublishStreamData,omitempty" xml:"DescribeUpPeakPublishStreamData,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas) SetDescribeUpPeakPublishStreamData(v []*DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatas {
	s.DescribeUpPeakPublishStreamData = v
	return s
}

type DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData struct {
	PublishStreamNum *int    `json:"PublishStreamNum,omitempty" xml:"PublishStreamNum,omitempty" require:"true"`
	PeakTime         *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty" require:"true"`
	QueryTime        *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty" require:"true"`
	StatName         *string `json:"StatName,omitempty" xml:"StatName,omitempty" require:"true"`
	BandWidth        *string `json:"BandWidth,omitempty" xml:"BandWidth,omitempty" require:"true"`
}

func (s DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) GoString() string {
	return s.String()
}

func (s *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetPublishStreamNum(v int) *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.PublishStreamNum = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetPeakTime(v string) *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.PeakTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetQueryTime(v string) *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.QueryTime = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetStatName(v string) *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.StatName = &v
	return s
}

func (s *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData) SetBandWidth(v string) *DescribeUpPeakPublishStreamDataResponseDescribeUpPeakPublishStreamDatasDescribeUpPeakPublishStreamData {
	s.BandWidth = &v
	return s
}

type DeleteLiveDomainMappingRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	PushDomain    *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty" require:"true"`
	PullDomain    *string `json:"PullDomain,omitempty" xml:"PullDomain,omitempty" require:"true"`
}

func (s DeleteLiveDomainMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveDomainMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainMappingRequest) SetSecurityToken(v string) *DeleteLiveDomainMappingRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveDomainMappingRequest) SetPushDomain(v string) *DeleteLiveDomainMappingRequest {
	s.PushDomain = &v
	return s
}

func (s *DeleteLiveDomainMappingRequest) SetPullDomain(v string) *DeleteLiveDomainMappingRequest {
	s.PullDomain = &v
	return s
}

type DeleteLiveDomainMappingResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveDomainMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveDomainMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainMappingResponse) SetRequestId(v string) *DeleteLiveDomainMappingResponse {
	s.RequestId = &v
	return s
}

type AddLiveDomainMappingRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	PushDomain    *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty" require:"true"`
	PullDomain    *string `json:"PullDomain,omitempty" xml:"PullDomain,omitempty" require:"true"`
}

func (s AddLiveDomainMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDomainMappingRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDomainMappingRequest) SetSecurityToken(v string) *AddLiveDomainMappingRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveDomainMappingRequest) SetPushDomain(v string) *AddLiveDomainMappingRequest {
	s.PushDomain = &v
	return s
}

func (s *AddLiveDomainMappingRequest) SetPullDomain(v string) *AddLiveDomainMappingRequest {
	s.PullDomain = &v
	return s
}

type AddLiveDomainMappingResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveDomainMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDomainMappingResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDomainMappingResponse) SetRequestId(v string) *AddLiveDomainMappingResponse {
	s.RequestId = &v
	return s
}

type AddCasterEpisodeGroupContentRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
}

func (s AddCasterEpisodeGroupContentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupContentRequest) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentRequest) SetClientToken(v string) *AddCasterEpisodeGroupContentRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCasterEpisodeGroupContentRequest) SetContent(v string) *AddCasterEpisodeGroupContentRequest {
	s.Content = &v
	return s
}

type AddCasterEpisodeGroupContentResponse struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ProgramId *string                                      `json:"ProgramId,omitempty" xml:"ProgramId,omitempty" require:"true"`
	ItemIds   *AddCasterEpisodeGroupContentResponseItemIds `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" require:"true" type:"Struct"`
}

func (s AddCasterEpisodeGroupContentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupContentResponse) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentResponse) SetRequestId(v string) *AddCasterEpisodeGroupContentResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterEpisodeGroupContentResponse) SetProgramId(v string) *AddCasterEpisodeGroupContentResponse {
	s.ProgramId = &v
	return s
}

func (s *AddCasterEpisodeGroupContentResponse) SetItemIds(v *AddCasterEpisodeGroupContentResponseItemIds) *AddCasterEpisodeGroupContentResponse {
	s.ItemIds = v
	return s
}

type AddCasterEpisodeGroupContentResponseItemIds struct {
	ItemId []*string `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true" type:"Repeated"`
}

func (s AddCasterEpisodeGroupContentResponseItemIds) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupContentResponseItemIds) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentResponseItemIds) SetItemId(v []*string) *AddCasterEpisodeGroupContentResponseItemIds {
	s.ItemId = v
	return s
}

type ModifyCasterProgramRequest struct {
	CasterId *string                              `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	Episode  []*ModifyCasterProgramRequestEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyCasterProgramRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramRequest) SetCasterId(v string) *ModifyCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterProgramRequest) SetEpisode(v []*ModifyCasterProgramRequestEpisode) *ModifyCasterProgramRequest {
	s.Episode = v
	return s
}

type ModifyCasterProgramRequestEpisode struct {
	EpisodeId   *string   `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	EpisodeType *string   `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	EpisodeName *string   `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	ResourceId  *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	StartTime   *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SwitchType  *string   `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s ModifyCasterProgramRequestEpisode) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterProgramRequestEpisode) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramRequestEpisode) SetEpisodeId(v string) *ModifyCasterProgramRequestEpisode {
	s.EpisodeId = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEpisodeType(v string) *ModifyCasterProgramRequestEpisode {
	s.EpisodeType = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEpisodeName(v string) *ModifyCasterProgramRequestEpisode {
	s.EpisodeName = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetResourceId(v string) *ModifyCasterProgramRequestEpisode {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetComponentId(v []*string) *ModifyCasterProgramRequestEpisode {
	s.ComponentId = v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetStartTime(v string) *ModifyCasterProgramRequestEpisode {
	s.StartTime = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEndTime(v string) *ModifyCasterProgramRequestEpisode {
	s.EndTime = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetSwitchType(v string) *ModifyCasterProgramRequestEpisode {
	s.SwitchType = &v
	return s
}

type ModifyCasterProgramResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s ModifyCasterProgramResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramResponse) SetRequestId(v string) *ModifyCasterProgramResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterProgramResponse) SetCasterId(v string) *ModifyCasterProgramResponse {
	s.CasterId = &v
	return s
}

type ModifyCasterEpisodeRequest struct {
	CasterId    *string   `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	EpisodeId   *string   `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
	EpisodeName *string   `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	ResourceId  *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	StartTime   *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SwitchType  *string   `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s ModifyCasterEpisodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterEpisodeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterEpisodeRequest) SetCasterId(v string) *ModifyCasterEpisodeRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetEpisodeId(v string) *ModifyCasterEpisodeRequest {
	s.EpisodeId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetEpisodeName(v string) *ModifyCasterEpisodeRequest {
	s.EpisodeName = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetResourceId(v string) *ModifyCasterEpisodeRequest {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetComponentId(v []*string) *ModifyCasterEpisodeRequest {
	s.ComponentId = v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetStartTime(v string) *ModifyCasterEpisodeRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetEndTime(v string) *ModifyCasterEpisodeRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetSwitchType(v string) *ModifyCasterEpisodeRequest {
	s.SwitchType = &v
	return s
}

type ModifyCasterEpisodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
}

func (s ModifyCasterEpisodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterEpisodeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterEpisodeResponse) SetRequestId(v string) *ModifyCasterEpisodeResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterEpisodeResponse) SetCasterId(v string) *ModifyCasterEpisodeResponse {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterEpisodeResponse) SetEpisodeId(v string) *ModifyCasterEpisodeResponse {
	s.EpisodeId = &v
	return s
}

type DescribeCasterProgramRequest struct {
	CasterId    *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	EpisodeId   *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNum     *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status      *int    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCasterProgramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramRequest) SetCasterId(v string) *DescribeCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetEpisodeId(v string) *DescribeCasterProgramRequest {
	s.EpisodeId = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetEpisodeType(v string) *DescribeCasterProgramRequest {
	s.EpisodeType = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetStartTime(v string) *DescribeCasterProgramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetEndTime(v string) *DescribeCasterProgramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetPageNum(v int) *DescribeCasterProgramRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetPageSize(v int) *DescribeCasterProgramRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCasterProgramRequest) SetStatus(v int) *DescribeCasterProgramRequest {
	s.Status = &v
	return s
}

type DescribeCasterProgramResponse struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId      *string                                `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ProgramName   *string                                `json:"ProgramName,omitempty" xml:"ProgramName,omitempty" require:"true"`
	ProgramEffect *int                                   `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty" require:"true"`
	Total         *int                                   `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Episodes      *DescribeCasterProgramResponseEpisodes `json:"Episodes,omitempty" xml:"Episodes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterProgramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponse) SetRequestId(v string) *DescribeCasterProgramResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterProgramResponse) SetCasterId(v string) *DescribeCasterProgramResponse {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterProgramResponse) SetProgramName(v string) *DescribeCasterProgramResponse {
	s.ProgramName = &v
	return s
}

func (s *DescribeCasterProgramResponse) SetProgramEffect(v int) *DescribeCasterProgramResponse {
	s.ProgramEffect = &v
	return s
}

func (s *DescribeCasterProgramResponse) SetTotal(v int) *DescribeCasterProgramResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterProgramResponse) SetEpisodes(v *DescribeCasterProgramResponseEpisodes) *DescribeCasterProgramResponse {
	s.Episodes = v
	return s
}

type DescribeCasterProgramResponseEpisodes struct {
	Episode []*DescribeCasterProgramResponseEpisodesEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterProgramResponseEpisodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterProgramResponseEpisodes) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseEpisodes) SetEpisode(v []*DescribeCasterProgramResponseEpisodesEpisode) *DescribeCasterProgramResponseEpisodes {
	s.Episode = v
	return s
}

type DescribeCasterProgramResponseEpisodesEpisode struct {
	EpisodeId    *string                                                   `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
	EpisodeType  *string                                                   `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty" require:"true"`
	EpisodeName  *string                                                   `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty" require:"true"`
	ResourceId   *string                                                   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	StartTime    *string                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	SwitchType   *string                                                   `json:"SwitchType,omitempty" xml:"SwitchType,omitempty" require:"true"`
	Status       *int                                                      `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ComponentIds *DescribeCasterProgramResponseEpisodesEpisodeComponentIds `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterProgramResponseEpisodesEpisode) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterProgramResponseEpisodesEpisode) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetEpisodeId(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.EpisodeId = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetEpisodeType(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.EpisodeType = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetEpisodeName(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.EpisodeName = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetResourceId(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.ResourceId = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetStartTime(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.StartTime = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetEndTime(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.EndTime = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetSwitchType(v string) *DescribeCasterProgramResponseEpisodesEpisode {
	s.SwitchType = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetStatus(v int) *DescribeCasterProgramResponseEpisodesEpisode {
	s.Status = &v
	return s
}

func (s *DescribeCasterProgramResponseEpisodesEpisode) SetComponentIds(v *DescribeCasterProgramResponseEpisodesEpisodeComponentIds) *DescribeCasterProgramResponseEpisodesEpisode {
	s.ComponentIds = v
	return s
}

type DescribeCasterProgramResponseEpisodesEpisodeComponentIds struct {
	// ComponentId
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterProgramResponseEpisodesEpisodeComponentIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterProgramResponseEpisodesEpisodeComponentIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseEpisodesEpisodeComponentIds) SetComponentId(v []*string) *DescribeCasterProgramResponseEpisodesEpisodeComponentIds {
	s.ComponentId = v
	return s
}

type DeleteCasterProgramRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DeleteCasterProgramRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterProgramRequest) SetCasterId(v string) *DeleteCasterProgramRequest {
	s.CasterId = &v
	return s
}

type DeleteCasterProgramResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DeleteCasterProgramResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterProgramResponse) SetRequestId(v string) *DeleteCasterProgramResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterProgramResponse) SetCasterId(v string) *DeleteCasterProgramResponse {
	s.CasterId = &v
	return s
}

type DeleteCasterEpisodeGroupRequest struct {
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty" require:"true"`
}

func (s DeleteCasterEpisodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterEpisodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeGroupRequest) SetProgramId(v string) *DeleteCasterEpisodeGroupRequest {
	s.ProgramId = &v
	return s
}

type DeleteCasterEpisodeGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteCasterEpisodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterEpisodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeGroupResponse) SetRequestId(v string) *DeleteCasterEpisodeGroupResponse {
	s.RequestId = &v
	return s
}

type DeleteCasterEpisodeRequest struct {
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
}

func (s DeleteCasterEpisodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterEpisodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeRequest) SetCasterId(v string) *DeleteCasterEpisodeRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterEpisodeRequest) SetEpisodeId(v string) *DeleteCasterEpisodeRequest {
	s.EpisodeId = &v
	return s
}

type DeleteCasterEpisodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
}

func (s DeleteCasterEpisodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterEpisodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeResponse) SetRequestId(v string) *DeleteCasterEpisodeResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterEpisodeResponse) SetCasterId(v string) *DeleteCasterEpisodeResponse {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterEpisodeResponse) SetEpisodeId(v string) *DeleteCasterEpisodeResponse {
	s.EpisodeId = &v
	return s
}

type AddCasterProgramRequest struct {
	CasterId *string                           `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	Episode  []*AddCasterProgramRequestEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" require:"true" type:"Repeated"`
}

func (s AddCasterProgramRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *AddCasterProgramRequest) SetCasterId(v string) *AddCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterProgramRequest) SetEpisode(v []*AddCasterProgramRequestEpisode) *AddCasterProgramRequest {
	s.Episode = v
	return s
}

type AddCasterProgramRequestEpisode struct {
	EpisodeType *string   `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	EpisodeName *string   `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	ResourceId  *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	StartTime   *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SwitchType  *string   `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s AddCasterProgramRequestEpisode) String() string {
	return tea.Prettify(s)
}

func (s AddCasterProgramRequestEpisode) GoString() string {
	return s.String()
}

func (s *AddCasterProgramRequestEpisode) SetEpisodeType(v string) *AddCasterProgramRequestEpisode {
	s.EpisodeType = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetEpisodeName(v string) *AddCasterProgramRequestEpisode {
	s.EpisodeName = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetResourceId(v string) *AddCasterProgramRequestEpisode {
	s.ResourceId = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetComponentId(v []*string) *AddCasterProgramRequestEpisode {
	s.ComponentId = v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetStartTime(v string) *AddCasterProgramRequestEpisode {
	s.StartTime = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetEndTime(v string) *AddCasterProgramRequestEpisode {
	s.EndTime = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetSwitchType(v string) *AddCasterProgramRequestEpisode {
	s.SwitchType = &v
	return s
}

type AddCasterProgramResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EpisodeIds *AddCasterProgramResponseEpisodeIds `json:"EpisodeIds,omitempty" xml:"EpisodeIds,omitempty" require:"true" type:"Struct"`
}

func (s AddCasterProgramResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponse) SetRequestId(v string) *AddCasterProgramResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterProgramResponse) SetEpisodeIds(v *AddCasterProgramResponseEpisodeIds) *AddCasterProgramResponse {
	s.EpisodeIds = v
	return s
}

type AddCasterProgramResponseEpisodeIds struct {
	EpisodeId []*AddCasterProgramResponseEpisodeIdsEpisodeId `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true" type:"Repeated"`
}

func (s AddCasterProgramResponseEpisodeIds) String() string {
	return tea.Prettify(s)
}

func (s AddCasterProgramResponseEpisodeIds) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponseEpisodeIds) SetEpisodeId(v []*AddCasterProgramResponseEpisodeIdsEpisodeId) *AddCasterProgramResponseEpisodeIds {
	s.EpisodeId = v
	return s
}

type AddCasterProgramResponseEpisodeIdsEpisodeId struct {
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
}

func (s AddCasterProgramResponseEpisodeIdsEpisodeId) String() string {
	return tea.Prettify(s)
}

func (s AddCasterProgramResponseEpisodeIdsEpisodeId) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponseEpisodeIdsEpisodeId) SetEpisodeId(v string) *AddCasterProgramResponseEpisodeIdsEpisodeId {
	s.EpisodeId = &v
	return s
}

type AddCasterEpisodeGroupRequest struct {
	ClientToken   *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
	DomainName    *string                             `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Item          []*AddCasterEpisodeGroupRequestItem `json:"Item,omitempty" xml:"Item,omitempty" require:"true" type:"Repeated"`
	StartTime     *string                             `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	RepeatNum     *int                                `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty" require:"true"`
	SideOutputUrl *string                             `json:"SideOutputUrl,omitempty" xml:"SideOutputUrl,omitempty" require:"true"`
	CallbackUrl   *string                             `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty" require:"true"`
}

func (s AddCasterEpisodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupRequest) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupRequest) SetClientToken(v string) *AddCasterEpisodeGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCasterEpisodeGroupRequest) SetDomainName(v string) *AddCasterEpisodeGroupRequest {
	s.DomainName = &v
	return s
}

func (s *AddCasterEpisodeGroupRequest) SetItem(v []*AddCasterEpisodeGroupRequestItem) *AddCasterEpisodeGroupRequest {
	s.Item = v
	return s
}

func (s *AddCasterEpisodeGroupRequest) SetStartTime(v string) *AddCasterEpisodeGroupRequest {
	s.StartTime = &v
	return s
}

func (s *AddCasterEpisodeGroupRequest) SetRepeatNum(v int) *AddCasterEpisodeGroupRequest {
	s.RepeatNum = &v
	return s
}

func (s *AddCasterEpisodeGroupRequest) SetSideOutputUrl(v string) *AddCasterEpisodeGroupRequest {
	s.SideOutputUrl = &v
	return s
}

func (s *AddCasterEpisodeGroupRequest) SetCallbackUrl(v string) *AddCasterEpisodeGroupRequest {
	s.CallbackUrl = &v
	return s
}

type AddCasterEpisodeGroupRequestItem struct {
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	VodUrl   *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s AddCasterEpisodeGroupRequestItem) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupRequestItem) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupRequestItem) SetItemName(v string) *AddCasterEpisodeGroupRequestItem {
	s.ItemName = &v
	return s
}

func (s *AddCasterEpisodeGroupRequestItem) SetVodUrl(v string) *AddCasterEpisodeGroupRequestItem {
	s.VodUrl = &v
	return s
}

type AddCasterEpisodeGroupResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ProgramId *string                               `json:"ProgramId,omitempty" xml:"ProgramId,omitempty" require:"true"`
	ItemIds   *AddCasterEpisodeGroupResponseItemIds `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" require:"true" type:"Struct"`
}

func (s AddCasterEpisodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupResponse) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupResponse) SetRequestId(v string) *AddCasterEpisodeGroupResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterEpisodeGroupResponse) SetProgramId(v string) *AddCasterEpisodeGroupResponse {
	s.ProgramId = &v
	return s
}

func (s *AddCasterEpisodeGroupResponse) SetItemIds(v *AddCasterEpisodeGroupResponseItemIds) *AddCasterEpisodeGroupResponse {
	s.ItemIds = v
	return s
}

type AddCasterEpisodeGroupResponseItemIds struct {
	ItemId []*string `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true" type:"Repeated"`
}

func (s AddCasterEpisodeGroupResponseItemIds) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeGroupResponseItemIds) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupResponseItemIds) SetItemId(v []*string) *AddCasterEpisodeGroupResponseItemIds {
	s.ItemId = v
	return s
}

type AddCasterEpisodeRequest struct {
	CasterId    *string   `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	EpisodeType *string   `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty" require:"true"`
	EpisodeName *string   `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	ResourceId  *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	StartTime   *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	SwitchType  *string   `json:"SwitchType,omitempty" xml:"SwitchType,omitempty" require:"true"`
}

func (s AddCasterEpisodeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeRequest) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeRequest) SetCasterId(v string) *AddCasterEpisodeRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetEpisodeType(v string) *AddCasterEpisodeRequest {
	s.EpisodeType = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetEpisodeName(v string) *AddCasterEpisodeRequest {
	s.EpisodeName = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetResourceId(v string) *AddCasterEpisodeRequest {
	s.ResourceId = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetComponentId(v []*string) *AddCasterEpisodeRequest {
	s.ComponentId = v
	return s
}

func (s *AddCasterEpisodeRequest) SetStartTime(v string) *AddCasterEpisodeRequest {
	s.StartTime = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetEndTime(v string) *AddCasterEpisodeRequest {
	s.EndTime = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetSwitchType(v string) *AddCasterEpisodeRequest {
	s.SwitchType = &v
	return s
}

type AddCasterEpisodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty" require:"true"`
}

func (s AddCasterEpisodeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterEpisodeResponse) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeResponse) SetRequestId(v string) *AddCasterEpisodeResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterEpisodeResponse) SetEpisodeId(v string) *AddCasterEpisodeResponse {
	s.EpisodeId = &v
	return s
}

type DescribeLiveDomainTranscodeDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeLiveDomainTranscodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTranscodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeDataRequest) SetDomainName(v string) *DescribeLiveDomainTranscodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTranscodeDataRequest) SetStartTime(v string) *DescribeLiveDomainTranscodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTranscodeDataRequest) SetEndTime(v string) *DescribeLiveDomainTranscodeDataRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveDomainTranscodeDataResponse struct {
	RequestId          *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TranscodeDataInfos *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos `json:"TranscodeDataInfos,omitempty" xml:"TranscodeDataInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainTranscodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTranscodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeDataResponse) SetRequestId(v string) *DescribeLiveDomainTranscodeDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainTranscodeDataResponse) SetTranscodeDataInfos(v *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos) *DescribeLiveDomainTranscodeDataResponse {
	s.TranscodeDataInfos = v
	return s
}

type DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos struct {
	TranscodeDataInfo []*DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo `json:"TranscodeDataInfo,omitempty" xml:"TranscodeDataInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos) SetTranscodeDataInfo(v []*DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo) *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfos {
	s.TranscodeDataInfo = v
	return s
}

type DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo struct {
	Date   *string `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Total  *int    `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty" require:"true"`
}

func (s DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo) SetDate(v string) *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo {
	s.Date = &v
	return s
}

func (s *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo) SetTotal(v int) *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo {
	s.Total = &v
	return s
}

func (s *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo) SetDetail(v string) *DescribeLiveDomainTranscodeDataResponseTranscodeDataInfosTranscodeDataInfo {
	s.Detail = &v
	return s
}

type DescribeLiveDomainSnapshotDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeLiveDomainSnapshotDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetDomainName(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetStartTime(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataRequest) SetEndTime(v string) *DescribeLiveDomainSnapshotDataRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveDomainSnapshotDataResponse struct {
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SnapshotDataInfos *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos `json:"SnapshotDataInfos,omitempty" xml:"SnapshotDataInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainSnapshotDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponse) SetRequestId(v string) *DescribeLiveDomainSnapshotDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponse) SetSnapshotDataInfos(v *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos) *DescribeLiveDomainSnapshotDataResponse {
	s.SnapshotDataInfos = v
	return s
}

type DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos struct {
	SnapshotDataInfo []*DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo `json:"SnapshotDataInfo,omitempty" xml:"SnapshotDataInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos) SetSnapshotDataInfo(v []*DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo) *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfos {
	s.SnapshotDataInfo = v
	return s
}

type DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo struct {
	Date  *string `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Total *int    `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
}

func (s DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo) SetDate(v string) *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo {
	s.Date = &v
	return s
}

func (s *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo) SetTotal(v int) *DescribeLiveDomainSnapshotDataResponseSnapshotDataInfosSnapshotDataInfo {
	s.Total = &v
	return s
}

type DescribeLiveDomainRecordDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
}

func (s DescribeLiveDomainRecordDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRecordDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordDataRequest) SetDomainName(v string) *DescribeLiveDomainRecordDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRecordDataRequest) SetStartTime(v string) *DescribeLiveDomainRecordDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRecordDataRequest) SetEndTime(v string) *DescribeLiveDomainRecordDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRecordDataRequest) SetRecordType(v string) *DescribeLiveDomainRecordDataRequest {
	s.RecordType = &v
	return s
}

type DescribeLiveDomainRecordDataResponse struct {
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RecordDataInfos *DescribeLiveDomainRecordDataResponseRecordDataInfos `json:"RecordDataInfos,omitempty" xml:"RecordDataInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainRecordDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRecordDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordDataResponse) SetRequestId(v string) *DescribeLiveDomainRecordDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRecordDataResponse) SetRecordDataInfos(v *DescribeLiveDomainRecordDataResponseRecordDataInfos) *DescribeLiveDomainRecordDataResponse {
	s.RecordDataInfos = v
	return s
}

type DescribeLiveDomainRecordDataResponseRecordDataInfos struct {
	RecordDataInfo []*DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo `json:"RecordDataInfo,omitempty" xml:"RecordDataInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainRecordDataResponseRecordDataInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRecordDataResponseRecordDataInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfos) SetRecordDataInfo(v []*DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo) *DescribeLiveDomainRecordDataResponseRecordDataInfos {
	s.RecordDataInfo = v
	return s
}

type DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo struct {
	Date   *string                                                                  `json:"Date,omitempty" xml:"Date,omitempty" require:"true"`
	Total  *int                                                                     `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Detail *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail `json:"Detail,omitempty" xml:"Detail,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo) SetDate(v string) *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo {
	s.Date = &v
	return s
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo) SetTotal(v int) *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo {
	s.Total = &v
	return s
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo) SetDetail(v *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail) *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfo {
	s.Detail = v
	return s
}

type DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail struct {
	MP4 *int `json:"MP4,omitempty" xml:"MP4,omitempty" require:"true"`
	FLV *int `json:"FLV,omitempty" xml:"FLV,omitempty" require:"true"`
	TS  *int `json:"TS,omitempty" xml:"TS,omitempty" require:"true"`
}

func (s DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail) SetMP4(v int) *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail {
	s.MP4 = &v
	return s
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail) SetFLV(v int) *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail {
	s.FLV = &v
	return s
}

func (s *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail) SetTS(v int) *DescribeLiveDomainRecordDataResponseRecordDataInfosRecordDataInfoDetail {
	s.TS = &v
	return s
}

type RealTimeRecordCommandRequest struct {
	Command    *string `json:"Command,omitempty" xml:"Command,omitempty" require:"true"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
}

func (s RealTimeRecordCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RealTimeRecordCommandRequest) GoString() string {
	return s.String()
}

func (s *RealTimeRecordCommandRequest) SetCommand(v string) *RealTimeRecordCommandRequest {
	s.Command = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetDomainName(v string) *RealTimeRecordCommandRequest {
	s.DomainName = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetAppName(v string) *RealTimeRecordCommandRequest {
	s.AppName = &v
	return s
}

func (s *RealTimeRecordCommandRequest) SetStreamName(v string) *RealTimeRecordCommandRequest {
	s.StreamName = &v
	return s
}

type RealTimeRecordCommandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RealTimeRecordCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RealTimeRecordCommandResponse) GoString() string {
	return s.String()
}

func (s *RealTimeRecordCommandResponse) SetRequestId(v string) *RealTimeRecordCommandResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDomainTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeLiveDomainTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataRequest) SetDomainName(v string) *DescribeLiveDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetStartTime(v string) *DescribeLiveDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetEndTime(v string) *DescribeLiveDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetInterval(v string) *DescribeLiveDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeLiveDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeLiveDomainTrafficDataResponse struct {
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName             *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime              *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime                *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval           *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	TrafficDataPerInterval *DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponse) SetRequestId(v string) *DescribeLiveDomainTrafficDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetDomainName(v string) *DescribeLiveDomainTrafficDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetStartTime(v string) *DescribeLiveDomainTrafficDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetEndTime(v string) *DescribeLiveDomainTrafficDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetDataInterval(v string) *DescribeLiveDomainTrafficDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetTrafficDataPerInterval(v *DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval) *DescribeLiveDomainTrafficDataResponse {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval struct {
	DataModule []*DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval) SetDataModule(v []*DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) *DescribeLiveDomainTrafficDataResponseTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule struct {
	TimeStamp         *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	TrafficValue      *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty" require:"true"`
	HttpTrafficValue  *string `json:"HttpTrafficValue,omitempty" xml:"HttpTrafficValue,omitempty" require:"true"`
	HttpsTrafficValue *string `json:"HttpsTrafficValue,omitempty" xml:"HttpsTrafficValue,omitempty" require:"true"`
}

func (s DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) SetHttpTrafficValue(v string) *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule {
	s.HttpTrafficValue = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule) SetHttpsTrafficValue(v string) *DescribeLiveDomainTrafficDataResponseTrafficDataPerIntervalDataModule {
	s.HttpsTrafficValue = &v
	return s
}

type DescribeLiveDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeLiveDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataRequest) SetDomainName(v string) *DescribeLiveDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetStartTime(v string) *DescribeLiveDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetEndTime(v string) *DescribeLiveDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetInterval(v string) *DescribeLiveDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetIspNameEn(v string) *DescribeLiveDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeLiveDomainBpsDataResponse struct {
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName         *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	StartTime          *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime            *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DataInterval       *string                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty" require:"true"`
	BpsDataPerInterval *DescribeLiveDomainBpsDataResponseBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponse) SetRequestId(v string) *DescribeLiveDomainBpsDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetDomainName(v string) *DescribeLiveDomainBpsDataResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetStartTime(v string) *DescribeLiveDomainBpsDataResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetEndTime(v string) *DescribeLiveDomainBpsDataResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetDataInterval(v string) *DescribeLiveDomainBpsDataResponse {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetBpsDataPerInterval(v *DescribeLiveDomainBpsDataResponseBpsDataPerInterval) *DescribeLiveDomainBpsDataResponse {
	s.BpsDataPerInterval = v
	return s
}

type DescribeLiveDomainBpsDataResponseBpsDataPerInterval struct {
	DataModule []*DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveDomainBpsDataResponseBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponseBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponseBpsDataPerInterval) SetDataModule(v []*DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) *DescribeLiveDomainBpsDataResponseBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule struct {
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	BpsValue      *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty" require:"true"`
	HttpBpsValue  *string `json:"HttpBpsValue,omitempty" xml:"HttpBpsValue,omitempty" require:"true"`
	HttpsBpsValue *string `json:"HttpsBpsValue,omitempty" xml:"HttpsBpsValue,omitempty" require:"true"`
}

func (s DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) SetHttpBpsValue(v string) *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule {
	s.HttpBpsValue = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule) SetHttpsBpsValue(v string) *DescribeLiveDomainBpsDataResponseBpsDataPerIntervalDataModule {
	s.HttpsBpsValue = &v
	return s
}

type AddTrancodeSEIRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	Text       *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	Pattern    *string `json:"Pattern,omitempty" xml:"Pattern,omitempty" require:"true"`
	Repeat     *int    `json:"Repeat,omitempty" xml:"Repeat,omitempty" require:"true"`
	Delay      *int    `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
}

func (s AddTrancodeSEIRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTrancodeSEIRequest) GoString() string {
	return s.String()
}

func (s *AddTrancodeSEIRequest) SetDomainName(v string) *AddTrancodeSEIRequest {
	s.DomainName = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetAppName(v string) *AddTrancodeSEIRequest {
	s.AppName = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetStreamName(v string) *AddTrancodeSEIRequest {
	s.StreamName = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetText(v string) *AddTrancodeSEIRequest {
	s.Text = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetPattern(v string) *AddTrancodeSEIRequest {
	s.Pattern = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetRepeat(v int) *AddTrancodeSEIRequest {
	s.Repeat = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetDelay(v int) *AddTrancodeSEIRequest {
	s.Delay = &v
	return s
}

type AddTrancodeSEIResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddTrancodeSEIResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTrancodeSEIResponse) GoString() string {
	return s.String()
}

func (s *AddTrancodeSEIResponse) SetRequestId(v string) *AddTrancodeSEIResponse {
	s.RequestId = &v
	return s
}

type DeleteCasterSceneConfigRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s DeleteCasterSceneConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterSceneConfigRequest) SetCasterId(v string) *DeleteCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) SetSceneId(v string) *DeleteCasterSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) SetType(v string) *DeleteCasterSceneConfigRequest {
	s.Type = &v
	return s
}

type DeleteCasterSceneConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteCasterSceneConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterSceneConfigResponse) SetRequestId(v string) *DeleteCasterSceneConfigResponse {
	s.RequestId = &v
	return s
}

type AddCustomLiveStreamTranscodeRequest struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	App             *string `json:"App,omitempty" xml:"App,omitempty" require:"true"`
	Template        *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
	TemplateType    *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty" require:"true"`
	Height          *int    `json:"Height,omitempty" xml:"Height,omitempty"`
	Width           *int    `json:"Width,omitempty" xml:"Width,omitempty"`
	FPS             *int    `json:"FPS,omitempty" xml:"FPS,omitempty"`
	VideoBitrate    *int    `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	AudioBitrate    *int    `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	Gop             *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	Profile         *int    `json:"Profile,omitempty" xml:"Profile,omitempty"`
	AudioProfile    *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty"`
	AudioCodec      *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty"`
	AudioRate       *int    `json:"AudioRate,omitempty" xml:"AudioRate,omitempty"`
	AudioChannelNum *int    `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty"`
	Lazy            *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
}

func (s AddCustomLiveStreamTranscodeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *AddCustomLiveStreamTranscodeRequest) SetDomain(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetApp(v string) *AddCustomLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetTemplate(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetTemplateType(v string) *AddCustomLiveStreamTranscodeRequest {
	s.TemplateType = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetHeight(v int) *AddCustomLiveStreamTranscodeRequest {
	s.Height = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetWidth(v int) *AddCustomLiveStreamTranscodeRequest {
	s.Width = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetFPS(v int) *AddCustomLiveStreamTranscodeRequest {
	s.FPS = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetVideoBitrate(v int) *AddCustomLiveStreamTranscodeRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioBitrate(v int) *AddCustomLiveStreamTranscodeRequest {
	s.AudioBitrate = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetGop(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Gop = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetProfile(v int) *AddCustomLiveStreamTranscodeRequest {
	s.Profile = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioProfile(v string) *AddCustomLiveStreamTranscodeRequest {
	s.AudioProfile = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioCodec(v string) *AddCustomLiveStreamTranscodeRequest {
	s.AudioCodec = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioRate(v int) *AddCustomLiveStreamTranscodeRequest {
	s.AudioRate = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetAudioChannelNum(v int) *AddCustomLiveStreamTranscodeRequest {
	s.AudioChannelNum = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeRequest) SetLazy(v string) *AddCustomLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

type AddCustomLiveStreamTranscodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddCustomLiveStreamTranscodeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *AddCustomLiveStreamTranscodeResponse) SetRequestId(v string) *AddCustomLiveStreamTranscodeResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveRecordVodConfigsRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	PageNum    *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLiveRecordVodConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsRequest) SetDomainName(v string) *DescribeLiveRecordVodConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetAppName(v string) *DescribeLiveRecordVodConfigsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetStreamName(v string) *DescribeLiveRecordVodConfigsRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetPageNum(v int64) *DescribeLiveRecordVodConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsRequest) SetPageSize(v int64) *DescribeLiveRecordVodConfigsRequest {
	s.PageSize = &v
	return s
}

type DescribeLiveRecordVodConfigsResponse struct {
	RequestId            *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum              *int                                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize             *int                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Total                *string                                                   `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	LiveRecordVodConfigs *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs `json:"LiveRecordVodConfigs,omitempty" xml:"LiveRecordVodConfigs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveRecordVodConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponse) SetRequestId(v string) *DescribeLiveRecordVodConfigsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) SetPageNum(v int) *DescribeLiveRecordVodConfigsResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) SetPageSize(v int) *DescribeLiveRecordVodConfigsResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) SetTotal(v string) *DescribeLiveRecordVodConfigsResponse {
	s.Total = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponse) SetLiveRecordVodConfigs(v *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs) *DescribeLiveRecordVodConfigsResponse {
	s.LiveRecordVodConfigs = v
	return s
}

type DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs struct {
	LiveRecordVodConfig []*DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig `json:"LiveRecordVodConfig,omitempty" xml:"LiveRecordVodConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs) SetLiveRecordVodConfig(v []*DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigs {
	s.LiveRecordVodConfig = v
	return s
}

type DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig struct {
	CreateTime                 *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	DomainName                 *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName                    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName                 *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	VodTranscodeGroupId        *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty" require:"true"`
	CycleDuration              *int    `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty" require:"true"`
	AutoCompose                *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty" require:"true"`
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty" require:"true"`
}

func (s DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetCreateTime(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetDomainName(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetAppName(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetStreamName(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetVodTranscodeGroupId(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetCycleDuration(v int) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.CycleDuration = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetAutoCompose(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.AutoCompose = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig) SetComposeVodTranscodeGroupId(v string) *DescribeLiveRecordVodConfigsResponseLiveRecordVodConfigsLiveRecordVodConfig {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

type DeleteLiveRecordVodConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveRecordVodConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRecordVodConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordVodConfigRequest) SetDomainName(v string) *DeleteLiveRecordVodConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) SetAppName(v string) *DeleteLiveRecordVodConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) SetStreamName(v string) *DeleteLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

type DeleteLiveRecordVodConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveRecordVodConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRecordVodConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordVodConfigResponse) SetRequestId(v string) *DeleteLiveRecordVodConfigResponse {
	s.RequestId = &v
	return s
}

type AddLiveRecordVodConfigRequest struct {
	DomainName                 *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName                    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName                 *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	VodTranscodeGroupId        *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty" require:"true"`
	CycleDuration              *int    `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	AutoCompose                *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	StorageLocation            *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
}

func (s AddLiveRecordVodConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveRecordVodConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigRequest) SetDomainName(v string) *AddLiveRecordVodConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetAppName(v string) *AddLiveRecordVodConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetStreamName(v string) *AddLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetCycleDuration(v int) *AddLiveRecordVodConfigRequest {
	s.CycleDuration = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetAutoCompose(v string) *AddLiveRecordVodConfigRequest {
	s.AutoCompose = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetStorageLocation(v string) *AddLiveRecordVodConfigRequest {
	s.StorageLocation = &v
	return s
}

func (s *AddLiveRecordVodConfigRequest) SetComposeVodTranscodeGroupId(v string) *AddLiveRecordVodConfigRequest {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

type AddLiveRecordVodConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveRecordVodConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveRecordVodConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveRecordVodConfigResponse) SetRequestId(v string) *AddLiveRecordVodConfigResponse {
	s.RequestId = &v
	return s
}

type ModifyCasterComponentRequest struct {
	CasterId            *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ComponentId         *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
	ComponentName       *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ComponentType       *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	Effect              *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	ComponentLayer      *string `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty"`
	TextLayerContent    *string `json:"TextLayerContent,omitempty" xml:"TextLayerContent,omitempty"`
	ImageLayerContent   *string `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty"`
	CaptionLayerContent *string `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty"`
}

func (s ModifyCasterComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterComponentRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterComponentRequest) SetCasterId(v string) *ModifyCasterComponentRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentId(v string) *ModifyCasterComponentRequest {
	s.ComponentId = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentName(v string) *ModifyCasterComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentType(v string) *ModifyCasterComponentRequest {
	s.ComponentType = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetEffect(v string) *ModifyCasterComponentRequest {
	s.Effect = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetComponentLayer(v string) *ModifyCasterComponentRequest {
	s.ComponentLayer = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetTextLayerContent(v string) *ModifyCasterComponentRequest {
	s.TextLayerContent = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetImageLayerContent(v string) *ModifyCasterComponentRequest {
	s.ImageLayerContent = &v
	return s
}

func (s *ModifyCasterComponentRequest) SetCaptionLayerContent(v string) *ModifyCasterComponentRequest {
	s.CaptionLayerContent = &v
	return s
}

type ModifyCasterComponentResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
}

func (s ModifyCasterComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterComponentResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterComponentResponse) SetRequestId(v string) *ModifyCasterComponentResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterComponentResponse) SetComponentId(v string) *ModifyCasterComponentResponse {
	s.ComponentId = &v
	return s
}

type DescribeCasterComponentsRequest struct {
	CasterId    *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
}

func (s DescribeCasterComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsRequest) SetCasterId(v string) *DescribeCasterComponentsRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterComponentsRequest) SetComponentId(v string) *DescribeCasterComponentsRequest {
	s.ComponentId = &v
	return s
}

type DescribeCasterComponentsResponse struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total      *int                                        `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Components *DescribeCasterComponentsResponseComponents `json:"Components,omitempty" xml:"Components,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponse) SetRequestId(v string) *DescribeCasterComponentsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterComponentsResponse) SetTotal(v int) *DescribeCasterComponentsResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterComponentsResponse) SetComponents(v *DescribeCasterComponentsResponseComponents) *DescribeCasterComponentsResponse {
	s.Components = v
	return s
}

type DescribeCasterComponentsResponseComponents struct {
	Component []*DescribeCasterComponentsResponseComponentsComponent `json:"Component,omitempty" xml:"Component,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterComponentsResponseComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponents) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponents) SetComponent(v []*DescribeCasterComponentsResponseComponentsComponent) *DescribeCasterComponentsResponseComponents {
	s.Component = v
	return s
}

type DescribeCasterComponentsResponseComponentsComponent struct {
	ComponentId         *string                                                                 `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
	ComponentName       *string                                                                 `json:"ComponentName,omitempty" xml:"ComponentName,omitempty" require:"true"`
	LocationId          *string                                                                 `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true"`
	ComponentType       *string                                                                 `json:"ComponentType,omitempty" xml:"ComponentType,omitempty" require:"true"`
	Effect              *string                                                                 `json:"Effect,omitempty" xml:"Effect,omitempty" require:"true"`
	ComponentLayer      *DescribeCasterComponentsResponseComponentsComponentComponentLayer      `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty" require:"true" type:"Struct"`
	TextLayerContent    *DescribeCasterComponentsResponseComponentsComponentTextLayerContent    `json:"TextLayerContent,omitempty" xml:"TextLayerContent,omitempty" require:"true" type:"Struct"`
	ImageLayerContent   *DescribeCasterComponentsResponseComponentsComponentImageLayerContent   `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty" require:"true" type:"Struct"`
	CaptionLayerContent *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterComponentsResponseComponentsComponent) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponentsComponent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetComponentId(v string) *DescribeCasterComponentsResponseComponentsComponent {
	s.ComponentId = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetComponentName(v string) *DescribeCasterComponentsResponseComponentsComponent {
	s.ComponentName = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetLocationId(v string) *DescribeCasterComponentsResponseComponentsComponent {
	s.LocationId = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetComponentType(v string) *DescribeCasterComponentsResponseComponentsComponent {
	s.ComponentType = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetEffect(v string) *DescribeCasterComponentsResponseComponentsComponent {
	s.Effect = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetComponentLayer(v *DescribeCasterComponentsResponseComponentsComponentComponentLayer) *DescribeCasterComponentsResponseComponentsComponent {
	s.ComponentLayer = v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetTextLayerContent(v *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) *DescribeCasterComponentsResponseComponentsComponent {
	s.TextLayerContent = v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetImageLayerContent(v *DescribeCasterComponentsResponseComponentsComponentImageLayerContent) *DescribeCasterComponentsResponseComponentsComponent {
	s.ImageLayerContent = v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponent) SetCaptionLayerContent(v *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) *DescribeCasterComponentsResponseComponentsComponent {
	s.CaptionLayerContent = v
	return s
}

type DescribeCasterComponentsResponseComponentsComponentComponentLayer struct {
	HeightNormalized    *float32                                                                              `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty" require:"true"`
	WidthNormalized     *float32                                                                              `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty" require:"true"`
	PositionRefer       *string                                                                               `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty" require:"true"`
	Transparency        *int                                                                                  `json:"Transparency,omitempty" xml:"Transparency,omitempty" require:"true"`
	PositionNormalizeds *DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds `json:"PositionNormalizeds,omitempty" xml:"PositionNormalizeds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterComponentsResponseComponentsComponentComponentLayer) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponentsComponentComponentLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponentsComponentComponentLayer) SetHeightNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentComponentLayer {
	s.HeightNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentComponentLayer) SetWidthNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentComponentLayer {
	s.WidthNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentComponentLayer) SetPositionRefer(v string) *DescribeCasterComponentsResponseComponentsComponentComponentLayer {
	s.PositionRefer = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentComponentLayer) SetTransparency(v int) *DescribeCasterComponentsResponseComponentsComponentComponentLayer {
	s.Transparency = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentComponentLayer) SetPositionNormalizeds(v *DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds) *DescribeCasterComponentsResponseComponentsComponentComponentLayer {
	s.PositionNormalizeds = v
	return s
}

type DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds struct {
	// Position
	Position []*float32 `json:"Position,omitempty" xml:"Position,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds) SetPosition(v []*float32) *DescribeCasterComponentsResponseComponentsComponentComponentLayerPositionNormalizeds {
	s.Position = v
	return s
}

type DescribeCasterComponentsResponseComponentsComponentTextLayerContent struct {
	Text                  *string  `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	Color                 *string  `json:"Color,omitempty" xml:"Color,omitempty" require:"true"`
	FontName              *string  `json:"FontName,omitempty" xml:"FontName,omitempty" require:"true"`
	SizeNormalized        *float32 `json:"SizeNormalized,omitempty" xml:"SizeNormalized,omitempty" require:"true"`
	BorderWidthNormalized *float32 `json:"BorderWidthNormalized,omitempty" xml:"BorderWidthNormalized,omitempty" require:"true"`
	BorderColor           *string  `json:"BorderColor,omitempty" xml:"BorderColor,omitempty" require:"true"`
}

func (s DescribeCasterComponentsResponseComponentsComponentTextLayerContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponentsComponentTextLayerContent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) SetText(v string) *DescribeCasterComponentsResponseComponentsComponentTextLayerContent {
	s.Text = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) SetColor(v string) *DescribeCasterComponentsResponseComponentsComponentTextLayerContent {
	s.Color = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) SetFontName(v string) *DescribeCasterComponentsResponseComponentsComponentTextLayerContent {
	s.FontName = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) SetSizeNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentTextLayerContent {
	s.SizeNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) SetBorderWidthNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentTextLayerContent {
	s.BorderWidthNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentTextLayerContent) SetBorderColor(v string) *DescribeCasterComponentsResponseComponentsComponentTextLayerContent {
	s.BorderColor = &v
	return s
}

type DescribeCasterComponentsResponseComponentsComponentImageLayerContent struct {
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty" require:"true"`
}

func (s DescribeCasterComponentsResponseComponentsComponentImageLayerContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponentsComponentImageLayerContent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponentsComponentImageLayerContent) SetMaterialId(v string) *DescribeCasterComponentsResponseComponentsComponentImageLayerContent {
	s.MaterialId = &v
	return s
}

type DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent struct {
	LocationId            *string  `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true"`
	PtsOffset             *int     `json:"PtsOffset,omitempty" xml:"PtsOffset,omitempty" require:"true"`
	WordsCount            *int     `json:"WordsCount,omitempty" xml:"WordsCount,omitempty" require:"true"`
	Color                 *string  `json:"Color,omitempty" xml:"Color,omitempty" require:"true"`
	FontName              *string  `json:"FontName,omitempty" xml:"FontName,omitempty" require:"true"`
	SizeNormalized        *float32 `json:"SizeNormalized,omitempty" xml:"SizeNormalized,omitempty" require:"true"`
	BorderWidthNormalized *float32 `json:"BorderWidthNormalized,omitempty" xml:"BorderWidthNormalized,omitempty" require:"true"`
	BorderColor           *string  `json:"BorderColor,omitempty" xml:"BorderColor,omitempty" require:"true"`
	WordCountPerLine      *int     `json:"WordCountPerLine,omitempty" xml:"WordCountPerLine,omitempty" require:"true"`
	WordSpaceNormalized   *float32 `json:"WordSpaceNormalized,omitempty" xml:"WordSpaceNormalized,omitempty" require:"true"`
	LineSpaceNormalized   *float32 `json:"LineSpaceNormalized,omitempty" xml:"LineSpaceNormalized,omitempty" require:"true"`
}

func (s DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetLocationId(v string) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.LocationId = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetPtsOffset(v int) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.PtsOffset = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetWordsCount(v int) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.WordsCount = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetColor(v string) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.Color = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetFontName(v string) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.FontName = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetSizeNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.SizeNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetBorderWidthNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.BorderWidthNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetBorderColor(v string) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.BorderColor = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetWordCountPerLine(v int) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.WordCountPerLine = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetWordSpaceNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.WordSpaceNormalized = &v
	return s
}

func (s *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent) SetLineSpaceNormalized(v float32) *DescribeCasterComponentsResponseComponentsComponentCaptionLayerContent {
	s.LineSpaceNormalized = &v
	return s
}

type DeleteCasterComponentRequest struct {
	CasterId    *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
}

func (s DeleteCasterComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterComponentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterComponentRequest) SetCasterId(v string) *DeleteCasterComponentRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterComponentRequest) SetComponentId(v string) *DeleteCasterComponentRequest {
	s.ComponentId = &v
	return s
}

type DeleteCasterComponentResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId    *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
}

func (s DeleteCasterComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterComponentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterComponentResponse) SetRequestId(v string) *DeleteCasterComponentResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterComponentResponse) SetCasterId(v string) *DeleteCasterComponentResponse {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterComponentResponse) SetComponentId(v string) *DeleteCasterComponentResponse {
	s.ComponentId = &v
	return s
}

type AddCasterComponentRequest struct {
	CasterId            *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ComponentName       *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	LocationId          *string `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true"`
	ComponentType       *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty" require:"true"`
	Effect              *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	ComponentLayer      *string `json:"ComponentLayer,omitempty" xml:"ComponentLayer,omitempty" require:"true"`
	LayerOrder          *string `json:"LayerOrder,omitempty" xml:"LayerOrder,omitempty"`
	TextLayerContent    *string `json:"TextLayerContent,omitempty" xml:"TextLayerContent,omitempty"`
	ImageLayerContent   *string `json:"ImageLayerContent,omitempty" xml:"ImageLayerContent,omitempty"`
	CaptionLayerContent *string `json:"CaptionLayerContent,omitempty" xml:"CaptionLayerContent,omitempty"`
	HtmlLayerContent    *string `json:"HtmlLayerContent,omitempty" xml:"HtmlLayerContent,omitempty"`
}

func (s AddCasterComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterComponentRequest) GoString() string {
	return s.String()
}

func (s *AddCasterComponentRequest) SetCasterId(v string) *AddCasterComponentRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterComponentRequest) SetComponentName(v string) *AddCasterComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *AddCasterComponentRequest) SetLocationId(v string) *AddCasterComponentRequest {
	s.LocationId = &v
	return s
}

func (s *AddCasterComponentRequest) SetComponentType(v string) *AddCasterComponentRequest {
	s.ComponentType = &v
	return s
}

func (s *AddCasterComponentRequest) SetEffect(v string) *AddCasterComponentRequest {
	s.Effect = &v
	return s
}

func (s *AddCasterComponentRequest) SetComponentLayer(v string) *AddCasterComponentRequest {
	s.ComponentLayer = &v
	return s
}

func (s *AddCasterComponentRequest) SetLayerOrder(v string) *AddCasterComponentRequest {
	s.LayerOrder = &v
	return s
}

func (s *AddCasterComponentRequest) SetTextLayerContent(v string) *AddCasterComponentRequest {
	s.TextLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) SetImageLayerContent(v string) *AddCasterComponentRequest {
	s.ImageLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) SetCaptionLayerContent(v string) *AddCasterComponentRequest {
	s.CaptionLayerContent = &v
	return s
}

func (s *AddCasterComponentRequest) SetHtmlLayerContent(v string) *AddCasterComponentRequest {
	s.HtmlLayerContent = &v
	return s
}

type AddCasterComponentResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
}

func (s AddCasterComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterComponentResponse) GoString() string {
	return s.String()
}

func (s *AddCasterComponentResponse) SetRequestId(v string) *AddCasterComponentResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterComponentResponse) SetComponentId(v string) *AddCasterComponentResponse {
	s.ComponentId = &v
	return s
}

type StopCasterRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s StopCasterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCasterRequest) GoString() string {
	return s.String()
}

func (s *StopCasterRequest) SetCasterId(v string) *StopCasterRequest {
	s.CasterId = &v
	return s
}

type StopCasterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopCasterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCasterResponse) GoString() string {
	return s.String()
}

func (s *StopCasterResponse) SetRequestId(v string) *StopCasterResponse {
	s.RequestId = &v
	return s
}

type StartCasterRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s StartCasterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCasterRequest) GoString() string {
	return s.String()
}

func (s *StartCasterRequest) SetCasterId(v string) *StartCasterRequest {
	s.CasterId = &v
	return s
}

type StartCasterResponse struct {
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PvwSceneInfos *StartCasterResponsePvwSceneInfos `json:"PvwSceneInfos,omitempty" xml:"PvwSceneInfos,omitempty" require:"true" type:"Struct"`
	PgmSceneInfos *StartCasterResponsePgmSceneInfos `json:"PgmSceneInfos,omitempty" xml:"PgmSceneInfos,omitempty" require:"true" type:"Struct"`
}

func (s StartCasterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponse) GoString() string {
	return s.String()
}

func (s *StartCasterResponse) SetRequestId(v string) *StartCasterResponse {
	s.RequestId = &v
	return s
}

func (s *StartCasterResponse) SetPvwSceneInfos(v *StartCasterResponsePvwSceneInfos) *StartCasterResponse {
	s.PvwSceneInfos = v
	return s
}

func (s *StartCasterResponse) SetPgmSceneInfos(v *StartCasterResponsePgmSceneInfos) *StartCasterResponse {
	s.PgmSceneInfos = v
	return s
}

type StartCasterResponsePvwSceneInfos struct {
	SceneInfo []*StartCasterResponsePvwSceneInfosSceneInfo `json:"SceneInfo,omitempty" xml:"SceneInfo,omitempty" require:"true" type:"Repeated"`
}

func (s StartCasterResponsePvwSceneInfos) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponsePvwSceneInfos) GoString() string {
	return s.String()
}

func (s *StartCasterResponsePvwSceneInfos) SetSceneInfo(v []*StartCasterResponsePvwSceneInfosSceneInfo) *StartCasterResponsePvwSceneInfos {
	s.SceneInfo = v
	return s
}

type StartCasterResponsePvwSceneInfosSceneInfo struct {
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
}

func (s StartCasterResponsePvwSceneInfosSceneInfo) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponsePvwSceneInfosSceneInfo) GoString() string {
	return s.String()
}

func (s *StartCasterResponsePvwSceneInfosSceneInfo) SetSceneId(v string) *StartCasterResponsePvwSceneInfosSceneInfo {
	s.SceneId = &v
	return s
}

func (s *StartCasterResponsePvwSceneInfosSceneInfo) SetStreamUrl(v string) *StartCasterResponsePvwSceneInfosSceneInfo {
	s.StreamUrl = &v
	return s
}

type StartCasterResponsePgmSceneInfos struct {
	SceneInfo []*StartCasterResponsePgmSceneInfosSceneInfo `json:"SceneInfo,omitempty" xml:"SceneInfo,omitempty" require:"true" type:"Repeated"`
}

func (s StartCasterResponsePgmSceneInfos) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponsePgmSceneInfos) GoString() string {
	return s.String()
}

func (s *StartCasterResponsePgmSceneInfos) SetSceneInfo(v []*StartCasterResponsePgmSceneInfosSceneInfo) *StartCasterResponsePgmSceneInfos {
	s.SceneInfo = v
	return s
}

type StartCasterResponsePgmSceneInfosSceneInfo struct {
	SceneId     *string                                               `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	StreamUrl   *string                                               `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	StreamInfos *StartCasterResponsePgmSceneInfosSceneInfoStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" require:"true" type:"Struct"`
}

func (s StartCasterResponsePgmSceneInfosSceneInfo) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponsePgmSceneInfosSceneInfo) GoString() string {
	return s.String()
}

func (s *StartCasterResponsePgmSceneInfosSceneInfo) SetSceneId(v string) *StartCasterResponsePgmSceneInfosSceneInfo {
	s.SceneId = &v
	return s
}

func (s *StartCasterResponsePgmSceneInfosSceneInfo) SetStreamUrl(v string) *StartCasterResponsePgmSceneInfosSceneInfo {
	s.StreamUrl = &v
	return s
}

func (s *StartCasterResponsePgmSceneInfosSceneInfo) SetStreamInfos(v *StartCasterResponsePgmSceneInfosSceneInfoStreamInfos) *StartCasterResponsePgmSceneInfosSceneInfo {
	s.StreamInfos = v
	return s
}

type StartCasterResponsePgmSceneInfosSceneInfoStreamInfos struct {
	StreamInfo []*StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" require:"true" type:"Repeated"`
}

func (s StartCasterResponsePgmSceneInfosSceneInfoStreamInfos) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponsePgmSceneInfosSceneInfoStreamInfos) GoString() string {
	return s.String()
}

func (s *StartCasterResponsePgmSceneInfosSceneInfoStreamInfos) SetStreamInfo(v []*StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo) *StartCasterResponsePgmSceneInfosSceneInfoStreamInfos {
	s.StreamInfo = v
	return s
}

type StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo struct {
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty" require:"true"`
	VideoFormat     *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty" require:"true"`
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty" require:"true"`
}

func (s StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo) String() string {
	return tea.Prettify(s)
}

func (s StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo) GoString() string {
	return s.String()
}

func (s *StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo) SetTranscodeConfig(v string) *StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo {
	s.TranscodeConfig = &v
	return s
}

func (s *StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo) SetVideoFormat(v string) *StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo {
	s.VideoFormat = &v
	return s
}

func (s *StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo) SetOutputStreamUrl(v string) *StartCasterResponsePgmSceneInfosSceneInfoStreamInfosStreamInfo {
	s.OutputStreamUrl = &v
	return s
}

type DescribeLiveStreamHistoryUserNumRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeLiveStreamHistoryUserNumRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetSecurityToken(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetDomainName(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetAppName(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetStreamName(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetStartTime(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumRequest) SetEndTime(v string) *DescribeLiveStreamHistoryUserNumRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveStreamHistoryUserNumResponse struct {
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveStreamUserNumInfos *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos `json:"LiveStreamUserNumInfos,omitempty" xml:"LiveStreamUserNumInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamHistoryUserNumResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponse) SetRequestId(v string) *DescribeLiveStreamHistoryUserNumResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponse) SetLiveStreamUserNumInfos(v *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos) *DescribeLiveStreamHistoryUserNumResponse {
	s.LiveStreamUserNumInfos = v
	return s
}

type DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos struct {
	LiveStreamUserNumInfo []*DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo `json:"LiveStreamUserNumInfo,omitempty" xml:"LiveStreamUserNumInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos) SetLiveStreamUserNumInfo(v []*DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo) *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfos {
	s.LiveStreamUserNumInfo = v
	return s
}

type DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo struct {
	StreamTime *string `json:"StreamTime,omitempty" xml:"StreamTime,omitempty" require:"true"`
	UserNum    *string `json:"UserNum,omitempty" xml:"UserNum,omitempty" require:"true"`
}

func (s DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo) SetStreamTime(v string) *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo {
	s.StreamTime = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo) SetUserNum(v string) *DescribeLiveStreamHistoryUserNumResponseLiveStreamUserNumInfosLiveStreamUserNumInfo {
	s.UserNum = &v
	return s
}

type UpdateCasterSceneConfigRequest struct {
	CasterId    *string   `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId     *string   `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	LayoutId    *string   `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
}

func (s UpdateCasterSceneConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneConfigRequest) SetCasterId(v string) *UpdateCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetSceneId(v string) *UpdateCasterSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetLayoutId(v string) *UpdateCasterSceneConfigRequest {
	s.LayoutId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetComponentId(v []*string) *UpdateCasterSceneConfigRequest {
	s.ComponentId = v
	return s
}

type UpdateCasterSceneConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateCasterSceneConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneConfigResponse) SetRequestId(v string) *UpdateCasterSceneConfigResponse {
	s.RequestId = &v
	return s
}

type StopCasterSceneRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
}

func (s StopCasterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCasterSceneRequest) GoString() string {
	return s.String()
}

func (s *StopCasterSceneRequest) SetCasterId(v string) *StopCasterSceneRequest {
	s.CasterId = &v
	return s
}

func (s *StopCasterSceneRequest) SetSceneId(v string) *StopCasterSceneRequest {
	s.SceneId = &v
	return s
}

type StopCasterSceneResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopCasterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCasterSceneResponse) GoString() string {
	return s.String()
}

func (s *StopCasterSceneResponse) SetRequestId(v string) *StopCasterSceneResponse {
	s.RequestId = &v
	return s
}

type StartCasterSceneRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
}

func (s StartCasterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCasterSceneRequest) GoString() string {
	return s.String()
}

func (s *StartCasterSceneRequest) SetCasterId(v string) *StartCasterSceneRequest {
	s.CasterId = &v
	return s
}

func (s *StartCasterSceneRequest) SetSceneId(v string) *StartCasterSceneRequest {
	s.SceneId = &v
	return s
}

type StartCasterSceneResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
}

func (s StartCasterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCasterSceneResponse) GoString() string {
	return s.String()
}

func (s *StartCasterSceneResponse) SetRequestId(v string) *StartCasterSceneResponse {
	s.RequestId = &v
	return s
}

func (s *StartCasterSceneResponse) SetStreamUrl(v string) *StartCasterSceneResponse {
	s.StreamUrl = &v
	return s
}

type SetCasterSceneConfigRequest struct {
	CasterId    *string   `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId     *string   `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	LayoutId    *string   `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
}

func (s SetCasterSceneConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCasterSceneConfigRequest) SetCasterId(v string) *SetCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetSceneId(v string) *SetCasterSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetLayoutId(v string) *SetCasterSceneConfigRequest {
	s.LayoutId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetComponentId(v []*string) *SetCasterSceneConfigRequest {
	s.ComponentId = v
	return s
}

type SetCasterSceneConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetCasterSceneConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCasterSceneConfigResponse) SetRequestId(v string) *SetCasterSceneConfigResponse {
	s.RequestId = &v
	return s
}

type SetCasterConfigRequest struct {
	CasterId         *string  `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	CasterName       *string  `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	DomainName       *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	TranscodeConfig  *string  `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty"`
	RecordConfig     *string  `json:"RecordConfig,omitempty" xml:"RecordConfig,omitempty"`
	Delay            *float32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	UrgentMaterialId *string  `json:"UrgentMaterialId,omitempty" xml:"UrgentMaterialId,omitempty"`
	SideOutputUrl    *string  `json:"SideOutputUrl,omitempty" xml:"SideOutputUrl,omitempty"`
	CallbackUrl      *string  `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ProgramEffect    *int     `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty"`
	ProgramName      *string  `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	ChannelEnable    *int     `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty"`
}

func (s SetCasterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCasterConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCasterConfigRequest) SetCasterId(v string) *SetCasterConfigRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterConfigRequest) SetCasterName(v string) *SetCasterConfigRequest {
	s.CasterName = &v
	return s
}

func (s *SetCasterConfigRequest) SetDomainName(v string) *SetCasterConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetCasterConfigRequest) SetTranscodeConfig(v string) *SetCasterConfigRequest {
	s.TranscodeConfig = &v
	return s
}

func (s *SetCasterConfigRequest) SetRecordConfig(v string) *SetCasterConfigRequest {
	s.RecordConfig = &v
	return s
}

func (s *SetCasterConfigRequest) SetDelay(v float32) *SetCasterConfigRequest {
	s.Delay = &v
	return s
}

func (s *SetCasterConfigRequest) SetUrgentMaterialId(v string) *SetCasterConfigRequest {
	s.UrgentMaterialId = &v
	return s
}

func (s *SetCasterConfigRequest) SetSideOutputUrl(v string) *SetCasterConfigRequest {
	s.SideOutputUrl = &v
	return s
}

func (s *SetCasterConfigRequest) SetCallbackUrl(v string) *SetCasterConfigRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SetCasterConfigRequest) SetProgramEffect(v int) *SetCasterConfigRequest {
	s.ProgramEffect = &v
	return s
}

func (s *SetCasterConfigRequest) SetProgramName(v string) *SetCasterConfigRequest {
	s.ProgramName = &v
	return s
}

func (s *SetCasterConfigRequest) SetChannelEnable(v int) *SetCasterConfigRequest {
	s.ChannelEnable = &v
	return s
}

type SetCasterConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s SetCasterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCasterConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCasterConfigResponse) SetRequestId(v string) *SetCasterConfigResponse {
	s.RequestId = &v
	return s
}

func (s *SetCasterConfigResponse) SetCasterId(v string) *SetCasterConfigResponse {
	s.CasterId = &v
	return s
}

type ModifyCasterVideoResourceRequest struct {
	CasterId            *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ResourceId          *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceName        *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	LiveStreamUrl       *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	MaterialId          *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	VodUrl              *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
	BeginOffset         *int    `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
	EndOffset           *int    `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	RepeatNum           *int    `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty"`
	PtsCallbackInterval *int    `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty"`
}

func (s ModifyCasterVideoResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterVideoResourceRequest) SetCasterId(v string) *ModifyCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetResourceId(v string) *ModifyCasterVideoResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetResourceName(v string) *ModifyCasterVideoResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetLiveStreamUrl(v string) *ModifyCasterVideoResourceRequest {
	s.LiveStreamUrl = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetMaterialId(v string) *ModifyCasterVideoResourceRequest {
	s.MaterialId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetVodUrl(v string) *ModifyCasterVideoResourceRequest {
	s.VodUrl = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetBeginOffset(v int) *ModifyCasterVideoResourceRequest {
	s.BeginOffset = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetEndOffset(v int) *ModifyCasterVideoResourceRequest {
	s.EndOffset = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetRepeatNum(v int) *ModifyCasterVideoResourceRequest {
	s.RepeatNum = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetPtsCallbackInterval(v int) *ModifyCasterVideoResourceRequest {
	s.PtsCallbackInterval = &v
	return s
}

type ModifyCasterVideoResourceResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId   *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s ModifyCasterVideoResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterVideoResourceResponse) SetRequestId(v string) *ModifyCasterVideoResourceResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterVideoResourceResponse) SetCasterId(v string) *ModifyCasterVideoResourceResponse {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterVideoResourceResponse) SetResourceId(v string) *ModifyCasterVideoResourceResponse {
	s.ResourceId = &v
	return s
}

type ModifyCasterLayoutRequest struct {
	CasterId   *string                                `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	LayoutId   *string                                `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
	VideoLayer []*ModifyCasterLayoutRequestVideoLayer `json:"VideoLayer,omitempty" xml:"VideoLayer,omitempty" require:"true" type:"Repeated"`
	AudioLayer []*ModifyCasterLayoutRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" require:"true" type:"Repeated"`
	BlendList  []*string                              `json:"BlendList,omitempty" xml:"BlendList,omitempty" require:"true" type:"Repeated"`
	MixList    []*string                              `json:"MixList,omitempty" xml:"MixList,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyCasterLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterLayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutRequest) SetCasterId(v string) *ModifyCasterLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterLayoutRequest) SetLayoutId(v string) *ModifyCasterLayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *ModifyCasterLayoutRequest) SetVideoLayer(v []*ModifyCasterLayoutRequestVideoLayer) *ModifyCasterLayoutRequest {
	s.VideoLayer = v
	return s
}

func (s *ModifyCasterLayoutRequest) SetAudioLayer(v []*ModifyCasterLayoutRequestAudioLayer) *ModifyCasterLayoutRequest {
	s.AudioLayer = v
	return s
}

func (s *ModifyCasterLayoutRequest) SetBlendList(v []*string) *ModifyCasterLayoutRequest {
	s.BlendList = v
	return s
}

func (s *ModifyCasterLayoutRequest) SetMixList(v []*string) *ModifyCasterLayoutRequest {
	s.MixList = v
	return s
}

type ModifyCasterLayoutRequestVideoLayer struct {
	FillMode           *string    `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	HeightNormalized   *float32   `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	WidthNormalized    *float32   `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
	PositionRefer      *string    `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	FixedDelayDuration *int       `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
}

func (s ModifyCasterLayoutRequestVideoLayer) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterLayoutRequestVideoLayer) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetFillMode(v string) *ModifyCasterLayoutRequestVideoLayer {
	s.FillMode = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetHeightNormalized(v float32) *ModifyCasterLayoutRequestVideoLayer {
	s.HeightNormalized = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetWidthNormalized(v float32) *ModifyCasterLayoutRequestVideoLayer {
	s.WidthNormalized = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetPositionRefer(v string) *ModifyCasterLayoutRequestVideoLayer {
	s.PositionRefer = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetPositionNormalized(v []*float32) *ModifyCasterLayoutRequestVideoLayer {
	s.PositionNormalized = v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetFixedDelayDuration(v int) *ModifyCasterLayoutRequestVideoLayer {
	s.FixedDelayDuration = &v
	return s
}

type ModifyCasterLayoutRequestAudioLayer struct {
	VolumeRate         *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
	ValidChannel       *string  `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	FixedDelayDuration *int     `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
}

func (s ModifyCasterLayoutRequestAudioLayer) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterLayoutRequestAudioLayer) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutRequestAudioLayer) SetVolumeRate(v float32) *ModifyCasterLayoutRequestAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *ModifyCasterLayoutRequestAudioLayer) SetValidChannel(v string) *ModifyCasterLayoutRequestAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *ModifyCasterLayoutRequestAudioLayer) SetFixedDelayDuration(v int) *ModifyCasterLayoutRequestAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

type ModifyCasterLayoutResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LayoutId  *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
}

func (s ModifyCasterLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCasterLayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutResponse) SetRequestId(v string) *ModifyCasterLayoutResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterLayoutResponse) SetLayoutId(v string) *ModifyCasterLayoutResponse {
	s.LayoutId = &v
	return s
}

type EffectCasterVideoResourceRequest struct {
	CasterId   *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s EffectCasterVideoResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s EffectCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *EffectCasterVideoResourceRequest) SetCasterId(v string) *EffectCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *EffectCasterVideoResourceRequest) SetSceneId(v string) *EffectCasterVideoResourceRequest {
	s.SceneId = &v
	return s
}

func (s *EffectCasterVideoResourceRequest) SetResourceId(v string) *EffectCasterVideoResourceRequest {
	s.ResourceId = &v
	return s
}

type EffectCasterVideoResourceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EffectCasterVideoResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s EffectCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *EffectCasterVideoResourceResponse) SetRequestId(v string) *EffectCasterVideoResourceResponse {
	s.RequestId = &v
	return s
}

type EffectCasterUrgentRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
}

func (s EffectCasterUrgentRequest) String() string {
	return tea.Prettify(s)
}

func (s EffectCasterUrgentRequest) GoString() string {
	return s.String()
}

func (s *EffectCasterUrgentRequest) SetCasterId(v string) *EffectCasterUrgentRequest {
	s.CasterId = &v
	return s
}

func (s *EffectCasterUrgentRequest) SetSceneId(v string) *EffectCasterUrgentRequest {
	s.SceneId = &v
	return s
}

type EffectCasterUrgentResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EffectCasterUrgentResponse) String() string {
	return tea.Prettify(s)
}

func (s EffectCasterUrgentResponse) GoString() string {
	return s.String()
}

func (s *EffectCasterUrgentResponse) SetRequestId(v string) *EffectCasterUrgentResponse {
	s.RequestId = &v
	return s
}

type DescribeCasterVideoResourcesRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DescribeCasterVideoResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterVideoResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesRequest) SetCasterId(v string) *DescribeCasterVideoResourcesRequest {
	s.CasterId = &v
	return s
}

type DescribeCasterVideoResourcesResponse struct {
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total          *int                                                `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	VideoResources *DescribeCasterVideoResourcesResponseVideoResources `json:"VideoResources,omitempty" xml:"VideoResources,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterVideoResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponse) SetRequestId(v string) *DescribeCasterVideoResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponse) SetTotal(v int) *DescribeCasterVideoResourcesResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponse) SetVideoResources(v *DescribeCasterVideoResourcesResponseVideoResources) *DescribeCasterVideoResourcesResponse {
	s.VideoResources = v
	return s
}

type DescribeCasterVideoResourcesResponseVideoResources struct {
	VideoResource []*DescribeCasterVideoResourcesResponseVideoResourcesVideoResource `json:"VideoResource,omitempty" xml:"VideoResource,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterVideoResourcesResponseVideoResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponseVideoResources) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponseVideoResources) SetVideoResource(v []*DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) *DescribeCasterVideoResourcesResponseVideoResources {
	s.VideoResource = v
	return s
}

type DescribeCasterVideoResourcesResponseVideoResourcesVideoResource struct {
	MaterialId          *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty" require:"true"`
	ResourceId          *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceName        *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	LocationId          *string `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true"`
	LiveStreamUrl       *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty" require:"true"`
	RepeatNum           *int    `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty" require:"true"`
	VodUrl              *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty" require:"true"`
	BeginOffset         *int    `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty" require:"true"`
	EndOffset           *int    `json:"EndOffset,omitempty" xml:"EndOffset,omitempty" require:"true"`
	PtsCallbackInterval *int    `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty" require:"true"`
}

func (s DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetMaterialId(v string) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.MaterialId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetResourceId(v string) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetResourceName(v string) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.ResourceName = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetLocationId(v string) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.LocationId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetLiveStreamUrl(v string) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.LiveStreamUrl = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetRepeatNum(v int) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.RepeatNum = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetVodUrl(v string) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.VodUrl = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetBeginOffset(v int) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.BeginOffset = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetEndOffset(v int) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.EndOffset = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource) SetPtsCallbackInterval(v int) *DescribeCasterVideoResourcesResponseVideoResourcesVideoResource {
	s.PtsCallbackInterval = &v
	return s
}

type DescribeCasterStreamUrlRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DescribeCasterStreamUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterStreamUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlRequest) SetCasterId(v string) *DescribeCasterStreamUrlRequest {
	s.CasterId = &v
	return s
}

type DescribeCasterStreamUrlResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId      *string                                       `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	Total         *int                                          `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	CasterStreams *DescribeCasterStreamUrlResponseCasterStreams `json:"CasterStreams,omitempty" xml:"CasterStreams,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterStreamUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterStreamUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponse) SetRequestId(v string) *DescribeCasterStreamUrlResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterStreamUrlResponse) SetCasterId(v string) *DescribeCasterStreamUrlResponse {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterStreamUrlResponse) SetTotal(v int) *DescribeCasterStreamUrlResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterStreamUrlResponse) SetCasterStreams(v *DescribeCasterStreamUrlResponseCasterStreams) *DescribeCasterStreamUrlResponse {
	s.CasterStreams = v
	return s
}

type DescribeCasterStreamUrlResponseCasterStreams struct {
	CasterStream []*DescribeCasterStreamUrlResponseCasterStreamsCasterStream `json:"CasterStream,omitempty" xml:"CasterStream,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterStreamUrlResponseCasterStreams) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseCasterStreams) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseCasterStreams) SetCasterStream(v []*DescribeCasterStreamUrlResponseCasterStreamsCasterStream) *DescribeCasterStreamUrlResponseCasterStreams {
	s.CasterStream = v
	return s
}

type DescribeCasterStreamUrlResponseCasterStreamsCasterStream struct {
	SceneId     *string                                                              `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	StreamUrl   *string                                                              `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	RtmpUrl     *string                                                              `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty" require:"true"`
	OutputType  *int                                                                 `json:"OutputType,omitempty" xml:"OutputType,omitempty" require:"true"`
	StreamInfos *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterStreamUrlResponseCasterStreamsCasterStream) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseCasterStreamsCasterStream) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStream) SetSceneId(v string) *DescribeCasterStreamUrlResponseCasterStreamsCasterStream {
	s.SceneId = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStream) SetStreamUrl(v string) *DescribeCasterStreamUrlResponseCasterStreamsCasterStream {
	s.StreamUrl = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStream) SetRtmpUrl(v string) *DescribeCasterStreamUrlResponseCasterStreamsCasterStream {
	s.RtmpUrl = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStream) SetOutputType(v int) *DescribeCasterStreamUrlResponseCasterStreamsCasterStream {
	s.OutputType = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStream) SetStreamInfos(v *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos) *DescribeCasterStreamUrlResponseCasterStreamsCasterStream {
	s.StreamInfos = v
	return s
}

type DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos struct {
	StreamInfo []*DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos) SetStreamInfo(v []*DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo) *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfos {
	s.StreamInfo = v
	return s
}

type DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo struct {
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty" require:"true"`
	VideoFormat     *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty" require:"true"`
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty" require:"true"`
}

func (s DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo) SetTranscodeConfig(v string) *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo {
	s.TranscodeConfig = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo) SetVideoFormat(v string) *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo {
	s.VideoFormat = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo) SetOutputStreamUrl(v string) *DescribeCasterStreamUrlResponseCasterStreamsCasterStreamStreamInfosStreamInfo {
	s.OutputStreamUrl = &v
	return s
}

type DescribeCasterScenesRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeCasterScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesRequest) SetCasterId(v string) *DescribeCasterScenesRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterScenesRequest) SetSceneId(v string) *DescribeCasterScenesRequest {
	s.SceneId = &v
	return s
}

type DescribeCasterScenesResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int                                   `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	SceneList *DescribeCasterScenesResponseSceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponse) SetRequestId(v string) *DescribeCasterScenesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterScenesResponse) SetTotal(v int) *DescribeCasterScenesResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterScenesResponse) SetSceneList(v *DescribeCasterScenesResponseSceneList) *DescribeCasterScenesResponse {
	s.SceneList = v
	return s
}

type DescribeCasterScenesResponseSceneList struct {
	Scene []*DescribeCasterScenesResponseSceneListScene `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterScenesResponseSceneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesResponseSceneList) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseSceneList) SetScene(v []*DescribeCasterScenesResponseSceneListScene) *DescribeCasterScenesResponseSceneList {
	s.Scene = v
	return s
}

type DescribeCasterScenesResponseSceneListScene struct {
	SceneId      *string                                                 `json:"SceneId,omitempty" xml:"SceneId,omitempty" require:"true"`
	SceneName    *string                                                 `json:"SceneName,omitempty" xml:"SceneName,omitempty" require:"true"`
	OutputType   *string                                                 `json:"OutputType,omitempty" xml:"OutputType,omitempty" require:"true"`
	LayoutId     *string                                                 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
	StreamUrl    *string                                                 `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	Status       *int                                                    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	StreamInfos  *DescribeCasterScenesResponseSceneListSceneStreamInfos  `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" require:"true" type:"Struct"`
	ComponentIds *DescribeCasterScenesResponseSceneListSceneComponentIds `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterScenesResponseSceneListScene) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesResponseSceneListScene) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseSceneListScene) SetSceneId(v string) *DescribeCasterScenesResponseSceneListScene {
	s.SceneId = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetSceneName(v string) *DescribeCasterScenesResponseSceneListScene {
	s.SceneName = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetOutputType(v string) *DescribeCasterScenesResponseSceneListScene {
	s.OutputType = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetLayoutId(v string) *DescribeCasterScenesResponseSceneListScene {
	s.LayoutId = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetStreamUrl(v string) *DescribeCasterScenesResponseSceneListScene {
	s.StreamUrl = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetStatus(v int) *DescribeCasterScenesResponseSceneListScene {
	s.Status = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetStreamInfos(v *DescribeCasterScenesResponseSceneListSceneStreamInfos) *DescribeCasterScenesResponseSceneListScene {
	s.StreamInfos = v
	return s
}

func (s *DescribeCasterScenesResponseSceneListScene) SetComponentIds(v *DescribeCasterScenesResponseSceneListSceneComponentIds) *DescribeCasterScenesResponseSceneListScene {
	s.ComponentIds = v
	return s
}

type DescribeCasterScenesResponseSceneListSceneStreamInfos struct {
	StreamInfo []*DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterScenesResponseSceneListSceneStreamInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesResponseSceneListSceneStreamInfos) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseSceneListSceneStreamInfos) SetStreamInfo(v []*DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo) *DescribeCasterScenesResponseSceneListSceneStreamInfos {
	s.StreamInfo = v
	return s
}

type DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo struct {
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty" require:"true"`
	VideoFormat     *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty" require:"true"`
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty" require:"true"`
}

func (s DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo) SetTranscodeConfig(v string) *DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo {
	s.TranscodeConfig = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo) SetVideoFormat(v string) *DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo {
	s.VideoFormat = &v
	return s
}

func (s *DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo) SetOutputStreamUrl(v string) *DescribeCasterScenesResponseSceneListSceneStreamInfosStreamInfo {
	s.OutputStreamUrl = &v
	return s
}

type DescribeCasterScenesResponseSceneListSceneComponentIds struct {
	// componentId
	ComponentId []*string `json:"componentId,omitempty" xml:"componentId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterScenesResponseSceneListSceneComponentIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterScenesResponseSceneListSceneComponentIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseSceneListSceneComponentIds) SetComponentId(v []*string) *DescribeCasterScenesResponseSceneListSceneComponentIds {
	s.ComponentId = v
	return s
}

type DescribeCastersRequest struct {
	CasterId   *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNum    *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status     *int    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCastersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCastersRequest) GoString() string {
	return s.String()
}

func (s *DescribeCastersRequest) SetCasterId(v string) *DescribeCastersRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCastersRequest) SetCasterName(v string) *DescribeCastersRequest {
	s.CasterName = &v
	return s
}

func (s *DescribeCastersRequest) SetStartTime(v string) *DescribeCastersRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCastersRequest) SetEndTime(v string) *DescribeCastersRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCastersRequest) SetPageNum(v int) *DescribeCastersRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCastersRequest) SetPageSize(v int) *DescribeCastersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCastersRequest) SetStatus(v int) *DescribeCastersRequest {
	s.Status = &v
	return s
}

type DescribeCastersResponse struct {
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total      *int                               `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	CasterList *DescribeCastersResponseCasterList `json:"CasterList,omitempty" xml:"CasterList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCastersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCastersResponse) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponse) SetRequestId(v string) *DescribeCastersResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCastersResponse) SetTotal(v int) *DescribeCastersResponse {
	s.Total = &v
	return s
}

func (s *DescribeCastersResponse) SetCasterList(v *DescribeCastersResponseCasterList) *DescribeCastersResponse {
	s.CasterList = v
	return s
}

type DescribeCastersResponseCasterList struct {
	Caster []*DescribeCastersResponseCasterListCaster `json:"Caster,omitempty" xml:"Caster,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCastersResponseCasterList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCastersResponseCasterList) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseCasterList) SetCaster(v []*DescribeCastersResponseCasterListCaster) *DescribeCastersResponseCasterList {
	s.Caster = v
	return s
}

type DescribeCastersResponseCasterListCaster struct {
	Status         *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	NormType       *int    `json:"NormType,omitempty" xml:"NormType,omitempty" require:"true"`
	CasterId       *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	CasterName     *string `json:"CasterName,omitempty" xml:"CasterName,omitempty" require:"true"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	PurchaseTime   *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty" require:"true"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	CasterTemplate *string `json:"CasterTemplate,omitempty" xml:"CasterTemplate,omitempty" require:"true"`
	ChannelEnable  *int    `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty" require:"true"`
}

func (s DescribeCastersResponseCasterListCaster) String() string {
	return tea.Prettify(s)
}

func (s DescribeCastersResponseCasterListCaster) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseCasterListCaster) SetStatus(v int) *DescribeCastersResponseCasterListCaster {
	s.Status = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetNormType(v int) *DescribeCastersResponseCasterListCaster {
	s.NormType = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetCasterId(v string) *DescribeCastersResponseCasterListCaster {
	s.CasterId = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetCasterName(v string) *DescribeCastersResponseCasterListCaster {
	s.CasterName = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetCreateTime(v string) *DescribeCastersResponseCasterListCaster {
	s.CreateTime = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetStartTime(v string) *DescribeCastersResponseCasterListCaster {
	s.StartTime = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetPurchaseTime(v string) *DescribeCastersResponseCasterListCaster {
	s.PurchaseTime = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetExpireTime(v string) *DescribeCastersResponseCasterListCaster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetChargeType(v string) *DescribeCastersResponseCasterListCaster {
	s.ChargeType = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetCasterTemplate(v string) *DescribeCastersResponseCasterListCaster {
	s.CasterTemplate = &v
	return s
}

func (s *DescribeCastersResponseCasterListCaster) SetChannelEnable(v int) *DescribeCastersResponseCasterListCaster {
	s.ChannelEnable = &v
	return s
}

type DescribeCasterLayoutsRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
}

func (s DescribeCasterLayoutsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsRequest) SetCasterId(v string) *DescribeCasterLayoutsRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterLayoutsRequest) SetLayoutId(v string) *DescribeCasterLayoutsRequest {
	s.LayoutId = &v
	return s
}

type DescribeCasterLayoutsResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int                                  `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Layouts   *DescribeCasterLayoutsResponseLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterLayoutsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponse) SetRequestId(v string) *DescribeCasterLayoutsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterLayoutsResponse) SetTotal(v int) *DescribeCasterLayoutsResponse {
	s.Total = &v
	return s
}

func (s *DescribeCasterLayoutsResponse) SetLayouts(v *DescribeCasterLayoutsResponseLayouts) *DescribeCasterLayoutsResponse {
	s.Layouts = v
	return s
}

type DescribeCasterLayoutsResponseLayouts struct {
	Layout []*DescribeCasterLayoutsResponseLayoutsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseLayouts) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayouts) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayouts) SetLayout(v []*DescribeCasterLayoutsResponseLayoutsLayout) *DescribeCasterLayoutsResponseLayouts {
	s.Layout = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayout struct {
	LayoutId    *string                                                `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
	VideoLayers *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers `json:"VideoLayers,omitempty" xml:"VideoLayers,omitempty" require:"true" type:"Struct"`
	AudioLayers *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers `json:"AudioLayers,omitempty" xml:"AudioLayers,omitempty" require:"true" type:"Struct"`
	BlendList   *DescribeCasterLayoutsResponseLayoutsLayoutBlendList   `json:"BlendList,omitempty" xml:"BlendList,omitempty" require:"true" type:"Struct"`
	MixList     *DescribeCasterLayoutsResponseLayoutsLayoutMixList     `json:"MixList,omitempty" xml:"MixList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayout) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayout) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayout) SetLayoutId(v string) *DescribeCasterLayoutsResponseLayoutsLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayout) SetVideoLayers(v *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers) *DescribeCasterLayoutsResponseLayoutsLayout {
	s.VideoLayers = v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayout) SetAudioLayers(v *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers) *DescribeCasterLayoutsResponseLayoutsLayout {
	s.AudioLayers = v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayout) SetBlendList(v *DescribeCasterLayoutsResponseLayoutsLayoutBlendList) *DescribeCasterLayoutsResponseLayoutsLayout {
	s.BlendList = v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayout) SetMixList(v *DescribeCasterLayoutsResponseLayoutsLayoutMixList) *DescribeCasterLayoutsResponseLayoutsLayout {
	s.MixList = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers struct {
	VideoLayer []*DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer `json:"VideoLayer,omitempty" xml:"VideoLayer,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers) SetVideoLayer(v []*DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayers {
	s.VideoLayer = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer struct {
	FillMode            *string                                                                             `json:"FillMode,omitempty" xml:"FillMode,omitempty" require:"true"`
	HeightNormalized    *float32                                                                            `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty" require:"true"`
	WidthNormalized     *float32                                                                            `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty" require:"true"`
	PositionRefer       *string                                                                             `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty" require:"true"`
	FixedDelayDuration  *int                                                                                `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty" require:"true"`
	PositionNormalizeds *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds `json:"PositionNormalizeds,omitempty" xml:"PositionNormalizeds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) SetFillMode(v string) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer {
	s.FillMode = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) SetHeightNormalized(v float32) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer {
	s.HeightNormalized = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) SetWidthNormalized(v float32) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer {
	s.WidthNormalized = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) SetPositionRefer(v string) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer {
	s.PositionRefer = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) SetFixedDelayDuration(v int) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer) SetPositionNormalizeds(v *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayer {
	s.PositionNormalizeds = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds struct {
	// Position
	Position []*float32 `json:"Position,omitempty" xml:"Position,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) SetPosition(v []*float32) *DescribeCasterLayoutsResponseLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds {
	s.Position = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers struct {
	AudioLayer []*DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers) SetAudioLayer(v []*DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer) *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayers {
	s.AudioLayer = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer struct {
	VolumeRate         *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty" require:"true"`
	ValidChannel       *string  `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty" require:"true"`
	FixedDelayDuration *int     `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty" require:"true"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer) SetVolumeRate(v float32) *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer) SetValidChannel(v string) *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer) SetFixedDelayDuration(v int) *DescribeCasterLayoutsResponseLayoutsLayoutAudioLayersAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutBlendList struct {
	// LocationId
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutBlendList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutBlendList) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutBlendList) SetLocationId(v []*string) *DescribeCasterLayoutsResponseLayoutsLayoutBlendList {
	s.LocationId = v
	return s
}

type DescribeCasterLayoutsResponseLayoutsLayoutMixList struct {
	// LocationId
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutMixList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterLayoutsResponseLayoutsLayoutMixList) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseLayoutsLayoutMixList) SetLocationId(v []*string) *DescribeCasterLayoutsResponseLayoutsLayoutMixList {
	s.LocationId = v
	return s
}

type DescribeCasterConfigRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DescribeCasterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigRequest) SetCasterId(v string) *DescribeCasterConfigRequest {
	s.CasterId = &v
	return s
}

type DescribeCasterConfigResponse struct {
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId         *string                                      `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	CasterName       *string                                      `json:"CasterName,omitempty" xml:"CasterName,omitempty" require:"true"`
	DomainName       *string                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	Delay            *float32                                     `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
	UrgentMaterialId *string                                      `json:"UrgentMaterialId,omitempty" xml:"UrgentMaterialId,omitempty" require:"true"`
	SideOutputUrl    *string                                      `json:"SideOutputUrl,omitempty" xml:"SideOutputUrl,omitempty" require:"true"`
	CallbackUrl      *string                                      `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty" require:"true"`
	ProgramName      *string                                      `json:"ProgramName,omitempty" xml:"ProgramName,omitempty" require:"true"`
	ProgramEffect    *int                                         `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty" require:"true"`
	ChannelEnable    *int                                         `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty" require:"true"`
	TranscodeConfig  *DescribeCasterConfigResponseTranscodeConfig `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty" require:"true" type:"Struct"`
	RecordConfig     *DescribeCasterConfigResponseRecordConfig    `json:"RecordConfig,omitempty" xml:"RecordConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponse) SetRequestId(v string) *DescribeCasterConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetCasterId(v string) *DescribeCasterConfigResponse {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetCasterName(v string) *DescribeCasterConfigResponse {
	s.CasterName = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetDomainName(v string) *DescribeCasterConfigResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetDelay(v float32) *DescribeCasterConfigResponse {
	s.Delay = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetUrgentMaterialId(v string) *DescribeCasterConfigResponse {
	s.UrgentMaterialId = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetSideOutputUrl(v string) *DescribeCasterConfigResponse {
	s.SideOutputUrl = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetCallbackUrl(v string) *DescribeCasterConfigResponse {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetProgramName(v string) *DescribeCasterConfigResponse {
	s.ProgramName = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetProgramEffect(v int) *DescribeCasterConfigResponse {
	s.ProgramEffect = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetChannelEnable(v int) *DescribeCasterConfigResponse {
	s.ChannelEnable = &v
	return s
}

func (s *DescribeCasterConfigResponse) SetTranscodeConfig(v *DescribeCasterConfigResponseTranscodeConfig) *DescribeCasterConfigResponse {
	s.TranscodeConfig = v
	return s
}

func (s *DescribeCasterConfigResponse) SetRecordConfig(v *DescribeCasterConfigResponseRecordConfig) *DescribeCasterConfigResponse {
	s.RecordConfig = v
	return s
}

type DescribeCasterConfigResponseTranscodeConfig struct {
	CasterTemplate  *string                                                     `json:"CasterTemplate,omitempty" xml:"CasterTemplate,omitempty" require:"true"`
	LiveTemplateIds *DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds `json:"LiveTemplateIds,omitempty" xml:"LiveTemplateIds,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterConfigResponseTranscodeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigResponseTranscodeConfig) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseTranscodeConfig) SetCasterTemplate(v string) *DescribeCasterConfigResponseTranscodeConfig {
	s.CasterTemplate = &v
	return s
}

func (s *DescribeCasterConfigResponseTranscodeConfig) SetLiveTemplateIds(v *DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds) *DescribeCasterConfigResponseTranscodeConfig {
	s.LiveTemplateIds = v
	return s
}

type DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds struct {
	// LocationId
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds) SetLocationId(v []*string) *DescribeCasterConfigResponseTranscodeConfigLiveTemplateIds {
	s.LocationId = v
	return s
}

type DescribeCasterConfigResponseRecordConfig struct {
	OssEndpoint  *string                                               `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket    *string                                               `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	RecordFormat *DescribeCasterConfigResponseRecordConfigRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCasterConfigResponseRecordConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigResponseRecordConfig) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseRecordConfig) SetOssEndpoint(v string) *DescribeCasterConfigResponseRecordConfig {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeCasterConfigResponseRecordConfig) SetOssBucket(v string) *DescribeCasterConfigResponseRecordConfig {
	s.OssBucket = &v
	return s
}

func (s *DescribeCasterConfigResponseRecordConfig) SetRecordFormat(v *DescribeCasterConfigResponseRecordConfigRecordFormat) *DescribeCasterConfigResponseRecordConfig {
	s.RecordFormat = v
	return s
}

type DescribeCasterConfigResponseRecordConfigRecordFormat struct {
	RecordFormat []*DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCasterConfigResponseRecordConfigRecordFormat) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigResponseRecordConfigRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseRecordConfigRecordFormat) SetRecordFormat(v []*DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) *DescribeCasterConfigResponseRecordConfigRecordFormat {
	s.RecordFormat = v
	return s
}

type DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat struct {
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty" require:"true"`
	OssObjectPrefix      *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty" require:"true"`
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty" require:"true"`
	CycleDuration        *int    `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty" require:"true"`
}

func (s DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) String() string {
	return tea.Prettify(s)
}

func (s DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) SetFormat(v string) *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat {
	s.Format = &v
	return s
}

func (s *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) SetOssObjectPrefix(v string) *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) SetSliceOssObjectPrefix(v string) *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat) SetCycleDuration(v int) *DescribeCasterConfigResponseRecordConfigRecordFormatRecordFormat {
	s.CycleDuration = &v
	return s
}

type DeleteCasterVideoResourceRequest struct {
	CasterId   *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s DeleteCasterVideoResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterVideoResourceRequest) SetCasterId(v string) *DeleteCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterVideoResourceRequest) SetResourceId(v string) *DeleteCasterVideoResourceRequest {
	s.ResourceId = &v
	return s
}

type DeleteCasterVideoResourceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteCasterVideoResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterVideoResourceResponse) SetRequestId(v string) *DeleteCasterVideoResourceResponse {
	s.RequestId = &v
	return s
}

type DeleteCasterLayoutRequest struct {
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
}

func (s DeleteCasterLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterLayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterLayoutRequest) SetCasterId(v string) *DeleteCasterLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterLayoutRequest) SetLayoutId(v string) *DeleteCasterLayoutRequest {
	s.LayoutId = &v
	return s
}

type DeleteCasterLayoutResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	LayoutId  *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
}

func (s DeleteCasterLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterLayoutResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterLayoutResponse) SetRequestId(v string) *DeleteCasterLayoutResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterLayoutResponse) SetCasterId(v string) *DeleteCasterLayoutResponse {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterLayoutResponse) SetLayoutId(v string) *DeleteCasterLayoutResponse {
	s.LayoutId = &v
	return s
}

type DeleteCasterRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	CasterId      *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DeleteCasterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterRequest) SetSecurityToken(v string) *DeleteCasterRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteCasterRequest) SetCasterId(v string) *DeleteCasterRequest {
	s.CasterId = &v
	return s
}

type DeleteCasterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s DeleteCasterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCasterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterResponse) SetRequestId(v string) *DeleteCasterResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterResponse) SetCasterId(v string) *DeleteCasterResponse {
	s.CasterId = &v
	return s
}

type CreateCasterRequest struct {
	CasterName     *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
	NormType       *int    `json:"NormType,omitempty" xml:"NormType,omitempty" require:"true"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	PurchaseTime   *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CasterTemplate *string `json:"CasterTemplate,omitempty" xml:"CasterTemplate,omitempty"`
}

func (s CreateCasterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCasterRequest) GoString() string {
	return s.String()
}

func (s *CreateCasterRequest) SetCasterName(v string) *CreateCasterRequest {
	s.CasterName = &v
	return s
}

func (s *CreateCasterRequest) SetClientToken(v string) *CreateCasterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCasterRequest) SetNormType(v int) *CreateCasterRequest {
	s.NormType = &v
	return s
}

func (s *CreateCasterRequest) SetChargeType(v string) *CreateCasterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCasterRequest) SetPurchaseTime(v string) *CreateCasterRequest {
	s.PurchaseTime = &v
	return s
}

func (s *CreateCasterRequest) SetExpireTime(v string) *CreateCasterRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateCasterRequest) SetCasterTemplate(v string) *CreateCasterRequest {
	s.CasterTemplate = &v
	return s
}

type CreateCasterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s CreateCasterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCasterResponse) GoString() string {
	return s.String()
}

func (s *CreateCasterResponse) SetRequestId(v string) *CreateCasterResponse {
	s.RequestId = &v
	return s
}

func (s *CreateCasterResponse) SetCasterId(v string) *CreateCasterResponse {
	s.CasterId = &v
	return s
}

type CopyCasterSceneConfigRequest struct {
	CasterId    *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	FromSceneId *string `json:"FromSceneId,omitempty" xml:"FromSceneId,omitempty" require:"true"`
	ToSceneId   *string `json:"ToSceneId,omitempty" xml:"ToSceneId,omitempty" require:"true"`
}

func (s CopyCasterSceneConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyCasterSceneConfigRequest) SetCasterId(v string) *CopyCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) SetFromSceneId(v string) *CopyCasterSceneConfigRequest {
	s.FromSceneId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) SetToSceneId(v string) *CopyCasterSceneConfigRequest {
	s.ToSceneId = &v
	return s
}

type CopyCasterSceneConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CopyCasterSceneConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyCasterSceneConfigResponse) SetRequestId(v string) *CopyCasterSceneConfigResponse {
	s.RequestId = &v
	return s
}

type CopyCasterRequest struct {
	CasterName  *string `json:"CasterName,omitempty" xml:"CasterName,omitempty" require:"true"`
	SrcCasterId *string `json:"SrcCasterId,omitempty" xml:"SrcCasterId,omitempty" require:"true"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CopyCasterRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyCasterRequest) GoString() string {
	return s.String()
}

func (s *CopyCasterRequest) SetCasterName(v string) *CopyCasterRequest {
	s.CasterName = &v
	return s
}

func (s *CopyCasterRequest) SetSrcCasterId(v string) *CopyCasterRequest {
	s.SrcCasterId = &v
	return s
}

func (s *CopyCasterRequest) SetClientToken(v string) *CopyCasterRequest {
	s.ClientToken = &v
	return s
}

type CopyCasterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CasterId  *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
}

func (s CopyCasterResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyCasterResponse) GoString() string {
	return s.String()
}

func (s *CopyCasterResponse) SetRequestId(v string) *CopyCasterResponse {
	s.RequestId = &v
	return s
}

func (s *CopyCasterResponse) SetCasterId(v string) *CopyCasterResponse {
	s.CasterId = &v
	return s
}

type AddCasterVideoResourceRequest struct {
	CasterId            *string `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	ResourceName        *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	LocationId          *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	LiveStreamUrl       *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	StreamId            *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	MaterialId          *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	VodUrl              *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
	BeginOffset         *int    `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
	EndOffset           *int    `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	RepeatNum           *int    `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty"`
	PtsCallbackInterval *int    `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty"`
}

func (s AddCasterVideoResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *AddCasterVideoResourceRequest) SetCasterId(v string) *AddCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetResourceName(v string) *AddCasterVideoResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetLocationId(v string) *AddCasterVideoResourceRequest {
	s.LocationId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetLiveStreamUrl(v string) *AddCasterVideoResourceRequest {
	s.LiveStreamUrl = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetStreamId(v string) *AddCasterVideoResourceRequest {
	s.StreamId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetMaterialId(v string) *AddCasterVideoResourceRequest {
	s.MaterialId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetVodUrl(v string) *AddCasterVideoResourceRequest {
	s.VodUrl = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetBeginOffset(v int) *AddCasterVideoResourceRequest {
	s.BeginOffset = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetEndOffset(v int) *AddCasterVideoResourceRequest {
	s.EndOffset = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetRepeatNum(v int) *AddCasterVideoResourceRequest {
	s.RepeatNum = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetPtsCallbackInterval(v int) *AddCasterVideoResourceRequest {
	s.PtsCallbackInterval = &v
	return s
}

type AddCasterVideoResourceResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s AddCasterVideoResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *AddCasterVideoResourceResponse) SetRequestId(v string) *AddCasterVideoResourceResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterVideoResourceResponse) SetResourceId(v string) *AddCasterVideoResourceResponse {
	s.ResourceId = &v
	return s
}

type AddCasterLayoutRequest struct {
	CasterId   *string                             `json:"CasterId,omitempty" xml:"CasterId,omitempty" require:"true"`
	VideoLayer []*AddCasterLayoutRequestVideoLayer `json:"VideoLayer,omitempty" xml:"VideoLayer,omitempty" require:"true" type:"Repeated"`
	AudioLayer []*AddCasterLayoutRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" require:"true" type:"Repeated"`
	BlendList  []*string                           `json:"BlendList,omitempty" xml:"BlendList,omitempty" require:"true" type:"Repeated"`
	MixList    []*string                           `json:"MixList,omitempty" xml:"MixList,omitempty" require:"true" type:"Repeated"`
}

func (s AddCasterLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCasterLayoutRequest) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutRequest) SetCasterId(v string) *AddCasterLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterLayoutRequest) SetVideoLayer(v []*AddCasterLayoutRequestVideoLayer) *AddCasterLayoutRequest {
	s.VideoLayer = v
	return s
}

func (s *AddCasterLayoutRequest) SetAudioLayer(v []*AddCasterLayoutRequestAudioLayer) *AddCasterLayoutRequest {
	s.AudioLayer = v
	return s
}

func (s *AddCasterLayoutRequest) SetBlendList(v []*string) *AddCasterLayoutRequest {
	s.BlendList = v
	return s
}

func (s *AddCasterLayoutRequest) SetMixList(v []*string) *AddCasterLayoutRequest {
	s.MixList = v
	return s
}

type AddCasterLayoutRequestVideoLayer struct {
	FillMode           *string    `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	HeightNormalized   *float32   `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	WidthNormalized    *float32   `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
	PositionRefer      *string    `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	FixedDelayDuration *int       `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
}

func (s AddCasterLayoutRequestVideoLayer) String() string {
	return tea.Prettify(s)
}

func (s AddCasterLayoutRequestVideoLayer) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutRequestVideoLayer) SetFillMode(v string) *AddCasterLayoutRequestVideoLayer {
	s.FillMode = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetHeightNormalized(v float32) *AddCasterLayoutRequestVideoLayer {
	s.HeightNormalized = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetWidthNormalized(v float32) *AddCasterLayoutRequestVideoLayer {
	s.WidthNormalized = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetPositionRefer(v string) *AddCasterLayoutRequestVideoLayer {
	s.PositionRefer = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetPositionNormalized(v []*float32) *AddCasterLayoutRequestVideoLayer {
	s.PositionNormalized = v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetFixedDelayDuration(v int) *AddCasterLayoutRequestVideoLayer {
	s.FixedDelayDuration = &v
	return s
}

type AddCasterLayoutRequestAudioLayer struct {
	VolumeRate         *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
	ValidChannel       *string  `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	FixedDelayDuration *int     `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
}

func (s AddCasterLayoutRequestAudioLayer) String() string {
	return tea.Prettify(s)
}

func (s AddCasterLayoutRequestAudioLayer) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutRequestAudioLayer) SetVolumeRate(v float32) *AddCasterLayoutRequestAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *AddCasterLayoutRequestAudioLayer) SetValidChannel(v string) *AddCasterLayoutRequestAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *AddCasterLayoutRequestAudioLayer) SetFixedDelayDuration(v int) *AddCasterLayoutRequestAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

type AddCasterLayoutResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LayoutId  *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" require:"true"`
}

func (s AddCasterLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCasterLayoutResponse) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutResponse) SetRequestId(v string) *AddCasterLayoutResponse {
	s.RequestId = &v
	return s
}

func (s *AddCasterLayoutResponse) SetLayoutId(v string) *AddCasterLayoutResponse {
	s.LayoutId = &v
	return s
}

type DescribeLivePullStreamConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLivePullStreamConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLivePullStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigRequest) SetDomainName(v string) *DescribeLivePullStreamConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeLivePullStreamConfigResponse struct {
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveAppRecordList *DescribeLivePullStreamConfigResponseLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLivePullStreamConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponse) SetRequestId(v string) *DescribeLivePullStreamConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponse) SetLiveAppRecordList(v *DescribeLivePullStreamConfigResponseLiveAppRecordList) *DescribeLivePullStreamConfigResponse {
	s.LiveAppRecordList = v
	return s
}

type DescribeLivePullStreamConfigResponseLiveAppRecordList struct {
	LiveAppRecord []*DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLivePullStreamConfigResponseLiveAppRecordList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponseLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordList) SetLiveAppRecord(v []*DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) *DescribeLivePullStreamConfigResponseLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

type DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty" require:"true"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) SetSourceUrl(v string) *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord {
	s.SourceUrl = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeLivePullStreamConfigResponseLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

type DeleteLivePullStreamInfoConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
}

func (s DeleteLivePullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivePullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetSecurityToken(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetDomainName(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetAppName(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetStreamName(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

type DeleteLivePullStreamInfoConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLivePullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLivePullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePullStreamInfoConfigResponse) SetRequestId(v string) *DeleteLivePullStreamInfoConfigResponse {
	s.RequestId = &v
	return s
}

type AddLivePullStreamInfoConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty" require:"true"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s AddLivePullStreamInfoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLivePullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLivePullStreamInfoConfigRequest) SetDomainName(v string) *AddLivePullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetAppName(v string) *AddLivePullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetStreamName(v string) *AddLivePullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetSourceUrl(v string) *AddLivePullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetStartTime(v string) *AddLivePullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetEndTime(v string) *AddLivePullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

type AddLivePullStreamInfoConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLivePullStreamInfoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLivePullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLivePullStreamInfoConfigResponse) SetRequestId(v string) *AddLivePullStreamInfoConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveStreamBitRateDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeLiveStreamBitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataRequest) SetSecurityToken(v string) *DescribeLiveStreamBitRateDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetDomainName(v string) *DescribeLiveStreamBitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetAppName(v string) *DescribeLiveStreamBitRateDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetStreamName(v string) *DescribeLiveStreamBitRateDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetStartTime(v string) *DescribeLiveStreamBitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataRequest) SetEndTime(v string) *DescribeLiveStreamBitRateDataRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveStreamBitRateDataResponse struct {
	RequestId                *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FrameRateAndBitRateInfos *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos `json:"FrameRateAndBitRateInfos,omitempty" xml:"FrameRateAndBitRateInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamBitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponse) SetRequestId(v string) *DescribeLiveStreamBitRateDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponse) SetFrameRateAndBitRateInfos(v *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos) *DescribeLiveStreamBitRateDataResponse {
	s.FrameRateAndBitRateInfos = v
	return s
}

type DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos struct {
	FrameRateAndBitRateInfo []*DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo `json:"FrameRateAndBitRateInfo,omitempty" xml:"FrameRateAndBitRateInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos) SetFrameRateAndBitRateInfo(v []*DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfos {
	s.FrameRateAndBitRateInfo = v
	return s
}

type DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo struct {
	StreamUrl      *string  `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	VideoFrameRate *float32 `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty" require:"true"`
	AudioFrameRate *float32 `json:"AudioFrameRate,omitempty" xml:"AudioFrameRate,omitempty" require:"true"`
	BitRate        *float32 `json:"BitRate,omitempty" xml:"BitRate,omitempty" require:"true"`
	Time           *string  `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetStreamUrl(v string) *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetVideoFrameRate(v float32) *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.VideoFrameRate = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetAudioFrameRate(v float32) *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.AudioFrameRate = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetBitRate(v float32) *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.BitRate = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetTime(v string) *DescribeLiveStreamBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.Time = &v
	return s
}

type AddLiveDetectNotifyConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl     *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
}

func (s AddLiveDetectNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *AddLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveDetectNotifyConfigRequest) SetDomainName(v string) *AddLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveDetectNotifyConfigRequest) SetNotifyUrl(v string) *AddLiveDetectNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

type AddLiveDetectNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveDetectNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDetectNotifyConfigResponse) SetRequestId(v string) *AddLiveDetectNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type AddLiveSnapshotDetectPornConfigRequest struct {
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string   `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	OssEndpoint   *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket     *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssObject     *string   `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	Scene         []*string `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
	Interval      *int      `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s AddLiveSnapshotDetectPornConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOssEndpoint(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOssBucket(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOssObject(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.OssObject = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetScene(v []*string) *AddLiveSnapshotDetectPornConfigRequest {
	s.Scene = v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetInterval(v int) *AddLiveSnapshotDetectPornConfigRequest {
	s.Interval = &v
	return s
}

type AddLiveSnapshotDetectPornConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveSnapshotDetectPornConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotDetectPornConfigResponse) SetRequestId(v string) *AddLiveSnapshotDetectPornConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveDetectNotifyConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DeleteLiveDetectNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *DeleteLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveDetectNotifyConfigRequest) SetDomainName(v string) *DeleteLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

type DeleteLiveDetectNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveDetectNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDetectNotifyConfigResponse) SetRequestId(v string) *DeleteLiveDetectNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveDetectNotifyConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveDetectNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *DescribeLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigRequest) SetDomainName(v string) *DescribeLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveDetectNotifyConfigResponse struct {
	RequestId              *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveDetectNotifyConfig *DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig `json:"LiveDetectNotifyConfig,omitempty" xml:"LiveDetectNotifyConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveDetectNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigResponse) SetRequestId(v string) *DescribeLiveDetectNotifyConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponse) SetLiveDetectNotifyConfig(v *DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig) *DescribeLiveDetectNotifyConfigResponse {
	s.LiveDetectNotifyConfig = v
	return s
}

type DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
}

func (s DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig) SetDomainName(v string) *DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig) SetNotifyUrl(v string) *DescribeLiveDetectNotifyConfigResponseLiveDetectNotifyConfig {
	s.NotifyUrl = &v
	return s
}

type DeleteLiveSnapshotDetectPornConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
}

func (s DeleteLiveSnapshotDetectPornConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

type DeleteLiveSnapshotDetectPornConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveSnapshotDetectPornConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) SetRequestId(v string) *DeleteLiveSnapshotDetectPornConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveSnapshotDetectPornConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	PageNum       *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DescribeLiveSnapshotDetectPornConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetPageNum(v int) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetPageSize(v int) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigRequest) SetOrder(v string) *DescribeLiveSnapshotDetectPornConfigRequest {
	s.Order = &v
	return s
}

type DescribeLiveSnapshotDetectPornConfigResponse struct {
	RequestId                        *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum                          *int                                                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize                         *int                                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Order                            *string                                                                       `json:"Order,omitempty" xml:"Order,omitempty" require:"true"`
	TotalNum                         *int                                                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage                        *int                                                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	LiveSnapshotDetectPornConfigList *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList `json:"LiveSnapshotDetectPornConfigList,omitempty" xml:"LiveSnapshotDetectPornConfigList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetRequestId(v string) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetPageNum(v int) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetPageSize(v int) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetOrder(v string) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.Order = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetTotalNum(v int) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetTotalPage(v int) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetLiveSnapshotDetectPornConfigList(v *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.LiveSnapshotDetectPornConfigList = v
	return s
}

type DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList struct {
	LiveSnapshotDetectPornConfig []*DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig `json:"LiveSnapshotDetectPornConfig,omitempty" xml:"LiveSnapshotDetectPornConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList) SetLiveSnapshotDetectPornConfig(v []*DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigList {
	s.LiveSnapshotDetectPornConfig = v
	return s
}

type DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig struct {
	DomainName  *string                                                                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string                                                                                                         `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	OssEndpoint *string                                                                                                         `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket   *string                                                                                                         `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssObject   *string                                                                                                         `json:"OssObject,omitempty" xml:"OssObject,omitempty" require:"true"`
	Interval    *int                                                                                                            `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	Scenes      *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetDomainName(v string) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetAppName(v string) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetOssEndpoint(v string) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetOssBucket(v string) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetOssObject(v string) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetInterval(v int) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.Interval = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetScenes(v *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.Scenes = v
	return s
}

type DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes struct {
	// scene
	Scene []*string `json:"scene,omitempty" xml:"scene,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) SetScene(v []*string) *DescribeLiveSnapshotDetectPornConfigResponseLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes {
	s.Scene = v
	return s
}

type UpdateLiveDetectNotifyConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl     *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
}

func (s UpdateLiveDetectNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *UpdateLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetDomainName(v string) *UpdateLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigRequest) SetNotifyUrl(v string) *UpdateLiveDetectNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

type UpdateLiveDetectNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveDetectNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveDetectNotifyConfigResponse) SetRequestId(v string) *UpdateLiveDetectNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type UpdateLiveSnapshotDetectPornConfigRequest struct {
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string   `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	OssEndpoint   *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssBucket     *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssObject     *string   `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	Interval      *int      `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Scene         []*string `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
}

func (s UpdateLiveSnapshotDetectPornConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOssEndpoint(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOssBucket(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOssObject(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OssObject = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetInterval(v int) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.Interval = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetScene(v []*string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.Scene = v
	return s
}

type UpdateLiveSnapshotDetectPornConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveSnapshotDetectPornConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotDetectPornConfigResponse) SetRequestId(v string) *UpdateLiveSnapshotDetectPornConfigResponse {
	s.RequestId = &v
	return s
}

type AddLiveRecordNotifyConfigRequest struct {
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl        *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
	NeedStatusNotify *bool   `json:"NeedStatusNotify,omitempty" xml:"NeedStatusNotify,omitempty"`
	OnDemandUrl      *string `json:"OnDemandUrl,omitempty" xml:"OnDemandUrl,omitempty"`
}

func (s AddLiveRecordNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *AddLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetDomainName(v string) *AddLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetNotifyUrl(v string) *AddLiveRecordNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetNeedStatusNotify(v bool) *AddLiveRecordNotifyConfigRequest {
	s.NeedStatusNotify = &v
	return s
}

func (s *AddLiveRecordNotifyConfigRequest) SetOnDemandUrl(v string) *AddLiveRecordNotifyConfigRequest {
	s.OnDemandUrl = &v
	return s
}

type AddLiveRecordNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveRecordNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveRecordNotifyConfigResponse) SetRequestId(v string) *AddLiveRecordNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveStreamsNotifyUrlConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DeleteLiveStreamsNotifyUrlConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DeleteLiveStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveStreamsNotifyUrlConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) SetRequestId(v string) *DeleteLiveStreamsNotifyUrlConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveRecordNotifyConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DeleteLiveRecordNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *DeleteLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveRecordNotifyConfigRequest) SetDomainName(v string) *DeleteLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

type DeleteLiveRecordNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveRecordNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordNotifyConfigResponse) SetRequestId(v string) *DeleteLiveRecordNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveRecordNotifyConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveRecordNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *DescribeLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigRequest) SetDomainName(v string) *DescribeLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveRecordNotifyConfigResponse struct {
	RequestId              *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveRecordNotifyConfig *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig `json:"LiveRecordNotifyConfig,omitempty" xml:"LiveRecordNotifyConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveRecordNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigResponse) SetRequestId(v string) *DescribeLiveRecordNotifyConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponse) SetLiveRecordNotifyConfig(v *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) *DescribeLiveRecordNotifyConfigResponse {
	s.LiveRecordNotifyConfig = v
	return s
}

type DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig struct {
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl        *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
	OnDemandUrl      *string `json:"OnDemandUrl,omitempty" xml:"OnDemandUrl,omitempty" require:"true"`
	NeedStatusNotify *bool   `json:"NeedStatusNotify,omitempty" xml:"NeedStatusNotify,omitempty" require:"true"`
}

func (s DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) SetDomainName(v string) *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) SetNotifyUrl(v string) *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) SetOnDemandUrl(v string) *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig {
	s.OnDemandUrl = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig) SetNeedStatusNotify(v bool) *DescribeLiveRecordNotifyConfigResponseLiveRecordNotifyConfig {
	s.NeedStatusNotify = &v
	return s
}

type DescribeLiveStreamsNotifyUrlConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
}

func (s DescribeLiveStreamsNotifyUrlConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DescribeLiveStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeLiveStreamsNotifyUrlConfigResponse struct {
	RequestId               *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LiveStreamsNotifyConfig *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig `json:"LiveStreamsNotifyConfig,omitempty" xml:"LiveStreamsNotifyConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamsNotifyUrlConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) SetRequestId(v string) *DescribeLiveStreamsNotifyUrlConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) SetLiveStreamsNotifyConfig(v *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) *DescribeLiveStreamsNotifyUrlConfigResponse {
	s.LiveStreamsNotifyConfig = v
	return s
}

type DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
	AuthType   *string `json:"AuthType,omitempty" xml:"AuthType,omitempty" require:"true"`
	AuthKey    *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty" require:"true"`
}

func (s DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) SetDomainName(v string) *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) SetNotifyUrl(v string) *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) SetAuthType(v string) *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig {
	s.AuthType = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig) SetAuthKey(v string) *DescribeLiveStreamsNotifyUrlConfigResponseLiveStreamsNotifyConfig {
	s.AuthKey = &v
	return s
}

type UpdateLiveRecordNotifyConfigRequest struct {
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl        *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OnDemandUrl      *string `json:"OnDemandUrl,omitempty" xml:"OnDemandUrl,omitempty"`
	NeedStatusNotify *bool   `json:"NeedStatusNotify,omitempty" xml:"NeedStatusNotify,omitempty"`
}

func (s UpdateLiveRecordNotifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetDomainName(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetNotifyUrl(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetOnDemandUrl(v string) *UpdateLiveRecordNotifyConfigRequest {
	s.OnDemandUrl = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigRequest) SetNeedStatusNotify(v bool) *UpdateLiveRecordNotifyConfigRequest {
	s.NeedStatusNotify = &v
	return s
}

type UpdateLiveRecordNotifyConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveRecordNotifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordNotifyConfigResponse) SetRequestId(v string) *UpdateLiveRecordNotifyConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveStreamsBlockListRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	PageNum       *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLiveStreamsBlockListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsBlockListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListRequest) SetSecurityToken(v string) *DescribeLiveStreamsBlockListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetDomainName(v string) *DescribeLiveStreamsBlockListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetPageNum(v int) *DescribeLiveStreamsBlockListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetPageSize(v int) *DescribeLiveStreamsBlockListRequest {
	s.PageSize = &v
	return s
}

type DescribeLiveStreamsBlockListResponse struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainName *string                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	PageNum    *int                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize   *int                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalNum   *int                                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage  *int                                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	StreamUrls *DescribeLiveStreamsBlockListResponseStreamUrls `json:"StreamUrls,omitempty" xml:"StreamUrls,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamsBlockListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsBlockListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListResponse) SetRequestId(v string) *DescribeLiveStreamsBlockListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetDomainName(v string) *DescribeLiveStreamsBlockListResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetPageNum(v int) *DescribeLiveStreamsBlockListResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetPageSize(v int) *DescribeLiveStreamsBlockListResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetTotalNum(v int) *DescribeLiveStreamsBlockListResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetTotalPage(v int) *DescribeLiveStreamsBlockListResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponse) SetStreamUrls(v *DescribeLiveStreamsBlockListResponseStreamUrls) *DescribeLiveStreamsBlockListResponse {
	s.StreamUrls = v
	return s
}

type DescribeLiveStreamsBlockListResponseStreamUrls struct {
	StreamUrl []*string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamsBlockListResponseStreamUrls) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsBlockListResponseStreamUrls) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListResponseStreamUrls) SetStreamUrl(v []*string) *DescribeLiveStreamsBlockListResponseStreamUrls {
	s.StreamUrl = v
	return s
}

type DescribeLiveStreamOnlineUserNumRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeLiveStreamOnlineUserNumRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOnlineUserNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOnlineUserNumRequest) SetSecurityToken(v string) *DescribeLiveStreamOnlineUserNumRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumRequest) SetDomainName(v string) *DescribeLiveStreamOnlineUserNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumRequest) SetAppName(v string) *DescribeLiveStreamOnlineUserNumRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumRequest) SetStreamName(v string) *DescribeLiveStreamOnlineUserNumRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumRequest) SetStartTime(v string) *DescribeLiveStreamOnlineUserNumRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumRequest) SetEndTime(v string) *DescribeLiveStreamOnlineUserNumRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveStreamOnlineUserNumResponse struct {
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalUserNumber *int64                                                 `json:"TotalUserNumber,omitempty" xml:"TotalUserNumber,omitempty" require:"true"`
	OnlineUserInfo  *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo `json:"OnlineUserInfo,omitempty" xml:"OnlineUserInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamOnlineUserNumResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOnlineUserNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOnlineUserNumResponse) SetRequestId(v string) *DescribeLiveStreamOnlineUserNumResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumResponse) SetTotalUserNumber(v int64) *DescribeLiveStreamOnlineUserNumResponse {
	s.TotalUserNumber = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumResponse) SetOnlineUserInfo(v *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo) *DescribeLiveStreamOnlineUserNumResponse {
	s.OnlineUserInfo = v
	return s
}

type DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo struct {
	LiveStreamOnlineUserNumInfo []*DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo `json:"LiveStreamOnlineUserNumInfo,omitempty" xml:"LiveStreamOnlineUserNumInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo) SetLiveStreamOnlineUserNumInfo(v []*DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfo {
	s.LiveStreamOnlineUserNumInfo = v
	return s
}

type DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo struct {
	StreamUrl  *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	UserNumber *int64  `json:"UserNumber,omitempty" xml:"UserNumber,omitempty" require:"true"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) SetStreamUrl(v string) *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) SetUserNumber(v int64) *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.UserNumber = &v
	return s
}

func (s *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo) SetTime(v string) *DescribeLiveStreamOnlineUserNumResponseOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.Time = &v
	return s
}

type DescribeLiveStreamsPublishListRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s DescribeLiveStreamsPublishListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsPublishListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListRequest) SetDomainName(v string) *DescribeLiveStreamsPublishListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetAppName(v string) *DescribeLiveStreamsPublishListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetStreamName(v string) *DescribeLiveStreamsPublishListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetStartTime(v string) *DescribeLiveStreamsPublishListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetEndTime(v string) *DescribeLiveStreamsPublishListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetPageSize(v int) *DescribeLiveStreamsPublishListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetPageNumber(v int) *DescribeLiveStreamsPublishListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetStreamType(v string) *DescribeLiveStreamsPublishListRequest {
	s.StreamType = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetQueryType(v string) *DescribeLiveStreamsPublishListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetOrderBy(v string) *DescribeLiveStreamsPublishListRequest {
	s.OrderBy = &v
	return s
}

type DescribeLiveStreamsPublishListResponse struct {
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum     *int                                               `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize    *int                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalNum    *int                                               `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage   *int                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	PublishInfo *DescribeLiveStreamsPublishListResponsePublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamsPublishListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponse) SetRequestId(v string) *DescribeLiveStreamsPublishListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetPageNum(v int) *DescribeLiveStreamsPublishListResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetPageSize(v int) *DescribeLiveStreamsPublishListResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetTotalNum(v int) *DescribeLiveStreamsPublishListResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetTotalPage(v int) *DescribeLiveStreamsPublishListResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponse) SetPublishInfo(v *DescribeLiveStreamsPublishListResponsePublishInfo) *DescribeLiveStreamsPublishListResponse {
	s.PublishInfo = v
	return s
}

type DescribeLiveStreamsPublishListResponsePublishInfo struct {
	LiveStreamPublishInfo []*DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo `json:"LiveStreamPublishInfo,omitempty" xml:"LiveStreamPublishInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamsPublishListResponsePublishInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponsePublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfo) SetLiveStreamPublishInfo(v []*DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) *DescribeLiveStreamsPublishListResponsePublishInfo {
	s.LiveStreamPublishInfo = v
	return s
}

type DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	StreamUrl     *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	PublishTime   *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty" require:"true"`
	StopTime      *string `json:"StopTime,omitempty" xml:"StopTime,omitempty" require:"true"`
	PublishUrl    *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty" require:"true"`
	ClientAddr    *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty" require:"true"`
	EdgeNodeAddr  *string `json:"EdgeNodeAddr,omitempty" xml:"EdgeNodeAddr,omitempty" require:"true"`
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty" require:"true"`
	PublishType   *string `json:"PublishType,omitempty" xml:"PublishType,omitempty" require:"true"`
	Transcoded    *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty" require:"true"`
	TranscodeId   *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty" require:"true"`
}

func (s DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetDomainName(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetAppName(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetStreamName(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetStreamUrl(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetPublishTime(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetStopTime(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.StopTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetPublishUrl(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetClientAddr(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.ClientAddr = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetEdgeNodeAddr(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.EdgeNodeAddr = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetPublishDomain(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetPublishType(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetTranscoded(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.Transcoded = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo) SetTranscodeId(v string) *DescribeLiveStreamsPublishListResponsePublishInfoLiveStreamPublishInfo {
	s.TranscodeId = &v
	return s
}

type DescribeLiveStreamsOnlineListRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum    *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
}

func (s DescribeLiveStreamsOnlineListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListRequest) SetDomainName(v string) *DescribeLiveStreamsOnlineListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetAppName(v string) *DescribeLiveStreamsOnlineListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetStreamName(v string) *DescribeLiveStreamsOnlineListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetPageSize(v int) *DescribeLiveStreamsOnlineListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetPageNum(v int) *DescribeLiveStreamsOnlineListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetStreamType(v string) *DescribeLiveStreamsOnlineListRequest {
	s.StreamType = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListRequest) SetQueryType(v string) *DescribeLiveStreamsOnlineListRequest {
	s.QueryType = &v
	return s
}

type DescribeLiveStreamsOnlineListResponse struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum    *int                                             `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize   *int                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalNum   *int                                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage  *int                                             `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	OnlineInfo *DescribeLiveStreamsOnlineListResponseOnlineInfo `json:"OnlineInfo,omitempty" xml:"OnlineInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamsOnlineListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponse) SetRequestId(v string) *DescribeLiveStreamsOnlineListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetPageNum(v int) *DescribeLiveStreamsOnlineListResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetPageSize(v int) *DescribeLiveStreamsOnlineListResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetTotalNum(v int) *DescribeLiveStreamsOnlineListResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetTotalPage(v int) *DescribeLiveStreamsOnlineListResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponse) SetOnlineInfo(v *DescribeLiveStreamsOnlineListResponseOnlineInfo) *DescribeLiveStreamsOnlineListResponse {
	s.OnlineInfo = v
	return s
}

type DescribeLiveStreamsOnlineListResponseOnlineInfo struct {
	LiveStreamOnlineInfo []*DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo `json:"LiveStreamOnlineInfo,omitempty" xml:"LiveStreamOnlineInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamsOnlineListResponseOnlineInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponseOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfo) SetLiveStreamOnlineInfo(v []*DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) *DescribeLiveStreamsOnlineListResponseOnlineInfo {
	s.LiveStreamOnlineInfo = v
	return s
}

type DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	PublishTime   *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty" require:"true"`
	PublishUrl    *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty" require:"true"`
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty" require:"true"`
	PublishType   *string `json:"PublishType,omitempty" xml:"PublishType,omitempty" require:"true"`
	Transcoded    *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty" require:"true"`
	TranscodeId   *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty" require:"true"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty" require:"true"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty" require:"true"`
	VideoCodecId  *int    `json:"VideoCodecId,omitempty" xml:"VideoCodecId,omitempty" require:"true"`
	VideoDataRate *int    `json:"VideoDataRate,omitempty" xml:"VideoDataRate,omitempty" require:"true"`
	FrameRate     *int    `json:"FrameRate,omitempty" xml:"FrameRate,omitempty" require:"true"`
	Width         *int    `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height        *int    `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	AudioCodecId  *int    `json:"AudioCodecId,omitempty" xml:"AudioCodecId,omitempty" require:"true"`
	AudioDataRate *int    `json:"AudioDataRate,omitempty" xml:"AudioDataRate,omitempty" require:"true"`
}

func (s DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetDomainName(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetAppName(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetStreamName(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetPublishTime(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetPublishUrl(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetPublishDomain(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetPublishType(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetTranscoded(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.Transcoded = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetTranscodeId(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.TranscodeId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetServerIp(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.ServerIp = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetClientIp(v string) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.ClientIp = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetVideoCodecId(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.VideoCodecId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetVideoDataRate(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.VideoDataRate = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetFrameRate(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.FrameRate = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetWidth(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetHeight(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetAudioCodecId(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.AudioCodecId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo) SetAudioDataRate(v int) *DescribeLiveStreamsOnlineListResponseOnlineInfoLiveStreamOnlineInfo {
	s.AudioDataRate = &v
	return s
}

type DescribeLiveStreamsControlHistoryRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeLiveStreamsControlHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetSecurityToken(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetDomainName(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetAppName(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetStartTime(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryRequest) SetEndTime(v string) *DescribeLiveStreamsControlHistoryRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveStreamsControlHistoryResponse struct {
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ControlInfo *DescribeLiveStreamsControlHistoryResponseControlInfo `json:"ControlInfo,omitempty" xml:"ControlInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamsControlHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponse) SetRequestId(v string) *DescribeLiveStreamsControlHistoryResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponse) SetControlInfo(v *DescribeLiveStreamsControlHistoryResponseControlInfo) *DescribeLiveStreamsControlHistoryResponse {
	s.ControlInfo = v
	return s
}

type DescribeLiveStreamsControlHistoryResponseControlInfo struct {
	LiveStreamControlInfo []*DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo `json:"LiveStreamControlInfo,omitempty" xml:"LiveStreamControlInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamsControlHistoryResponseControlInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponseControlInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponseControlInfo) SetLiveStreamControlInfo(v []*DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) *DescribeLiveStreamsControlHistoryResponseControlInfo {
	s.LiveStreamControlInfo = v
	return s
}

type DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo struct {
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	ClientIP   *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty" require:"true"`
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty" require:"true"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
}

func (s DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) SetStreamName(v string) *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) SetClientIP(v string) *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo {
	s.ClientIP = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) SetAction(v string) *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo {
	s.Action = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo) SetTimeStamp(v string) *DescribeLiveStreamsControlHistoryResponseControlInfoLiveStreamControlInfo {
	s.TimeStamp = &v
	return s
}

type AddLiveStreamTranscodeRequest struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	App          *string `json:"App,omitempty" xml:"App,omitempty" require:"true"`
	Template     *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
	Lazy         *string `json:"Lazy,omitempty" xml:"Lazy,omitempty"`
	Watermark    *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Mix          *string `json:"Mix,omitempty" xml:"Mix,omitempty"`
	OnlyAudio    *string `json:"OnlyAudio,omitempty" xml:"OnlyAudio,omitempty"`
	WaterPattern *string `json:"WaterPattern,omitempty" xml:"WaterPattern,omitempty"`
}

func (s AddLiveStreamTranscodeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *AddLiveStreamTranscodeRequest) SetDomain(v string) *AddLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetApp(v string) *AddLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetTemplate(v string) *AddLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetLazy(v string) *AddLiveStreamTranscodeRequest {
	s.Lazy = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetWatermark(v string) *AddLiveStreamTranscodeRequest {
	s.Watermark = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetMix(v string) *AddLiveStreamTranscodeRequest {
	s.Mix = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetOnlyAudio(v string) *AddLiveStreamTranscodeRequest {
	s.OnlyAudio = &v
	return s
}

func (s *AddLiveStreamTranscodeRequest) SetWaterPattern(v string) *AddLiveStreamTranscodeRequest {
	s.WaterPattern = &v
	return s
}

type AddLiveStreamTranscodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveStreamTranscodeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *AddLiveStreamTranscodeResponse) SetRequestId(v string) *AddLiveStreamTranscodeResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveStreamTranscodeRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	App           *string `json:"App,omitempty" xml:"App,omitempty" require:"true"`
	Template      *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
}

func (s DeleteLiveStreamTranscodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamTranscodeRequest) SetSecurityToken(v string) *DeleteLiveStreamTranscodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetDomain(v string) *DeleteLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetApp(v string) *DeleteLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetTemplate(v string) *DeleteLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

type DeleteLiveStreamTranscodeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveStreamTranscodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamTranscodeResponse) SetRequestId(v string) *DeleteLiveStreamTranscodeResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveStreamsFrameRateAndBitRateDataRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataRequest) SetSecurityToken(v string) *DescribeLiveStreamsFrameRateAndBitRateDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataRequest) SetDomainName(v string) *DescribeLiveStreamsFrameRateAndBitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataRequest) SetAppName(v string) *DescribeLiveStreamsFrameRateAndBitRateDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataRequest) SetStreamName(v string) *DescribeLiveStreamsFrameRateAndBitRateDataRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataRequest) SetStartTime(v string) *DescribeLiveStreamsFrameRateAndBitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataRequest) SetEndTime(v string) *DescribeLiveStreamsFrameRateAndBitRateDataRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveStreamsFrameRateAndBitRateDataResponse struct {
	RequestId                *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FrameRateAndBitRateInfos *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos `json:"FrameRateAndBitRateInfos,omitempty" xml:"FrameRateAndBitRateInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponse) SetRequestId(v string) *DescribeLiveStreamsFrameRateAndBitRateDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponse) SetFrameRateAndBitRateInfos(v *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) *DescribeLiveStreamsFrameRateAndBitRateDataResponse {
	s.FrameRateAndBitRateInfos = v
	return s
}

type DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos struct {
	FrameRateAndBitRateInfo []*DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo `json:"FrameRateAndBitRateInfo,omitempty" xml:"FrameRateAndBitRateInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos) SetFrameRateAndBitRateInfo(v []*DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfos {
	s.FrameRateAndBitRateInfo = v
	return s
}

type DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo struct {
	StreamUrl      *string  `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" require:"true"`
	VideoFrameRate *float32 `json:"VideoFrameRate,omitempty" xml:"VideoFrameRate,omitempty" require:"true"`
	AudioFrameRate *float32 `json:"AudioFrameRate,omitempty" xml:"AudioFrameRate,omitempty" require:"true"`
	BitRate        *float32 `json:"BitRate,omitempty" xml:"BitRate,omitempty" require:"true"`
	Time           *string  `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetStreamUrl(v string) *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetVideoFrameRate(v float32) *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.VideoFrameRate = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetAudioFrameRate(v float32) *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.AudioFrameRate = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetBitRate(v float32) *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.BitRate = &v
	return s
}

func (s *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo) SetTime(v string) *DescribeLiveStreamsFrameRateAndBitRateDataResponseFrameRateAndBitRateInfosFrameRateAndBitRateInfo {
	s.Time = &v
	return s
}

type ForbidLiveStreamRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName     *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty" require:"true"`
	Oneshot        *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	ResumeTime     *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
}

func (s ForbidLiveStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ForbidLiveStreamRequest) GoString() string {
	return s.String()
}

func (s *ForbidLiveStreamRequest) SetDomainName(v string) *ForbidLiveStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetAppName(v string) *ForbidLiveStreamRequest {
	s.AppName = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetStreamName(v string) *ForbidLiveStreamRequest {
	s.StreamName = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetLiveStreamType(v string) *ForbidLiveStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetOneshot(v string) *ForbidLiveStreamRequest {
	s.Oneshot = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetResumeTime(v string) *ForbidLiveStreamRequest {
	s.ResumeTime = &v
	return s
}

type ForbidLiveStreamResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ForbidLiveStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ForbidLiveStreamResponse) GoString() string {
	return s.String()
}

func (s *ForbidLiveStreamResponse) SetRequestId(v string) *ForbidLiveStreamResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveStreamTranscodeInfoRequest struct {
	DomainTranscodeName *string `json:"DomainTranscodeName,omitempty" xml:"DomainTranscodeName,omitempty" require:"true"`
}

func (s DescribeLiveStreamTranscodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoRequest) SetDomainTranscodeName(v string) *DescribeLiveStreamTranscodeInfoRequest {
	s.DomainTranscodeName = &v
	return s
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainTranscodeList *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList `json:"DomainTranscodeList,omitempty" xml:"DomainTranscodeList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamTranscodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponse) SetRequestId(v string) *DescribeLiveStreamTranscodeInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponse) SetDomainTranscodeList(v *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList) *DescribeLiveStreamTranscodeInfoResponse {
	s.DomainTranscodeList = v
	return s
}

type DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList struct {
	DomainTranscodeInfo []*DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo `json:"DomainTranscodeInfo,omitempty" xml:"DomainTranscodeInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList) SetDomainTranscodeInfo(v []*DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeList {
	s.DomainTranscodeInfo = v
	return s
}

type DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo struct {
	TranscodeApp              *string                                                                                                 `json:"TranscodeApp,omitempty" xml:"TranscodeApp,omitempty" require:"true"`
	TranscodeName             *string                                                                                                 `json:"TranscodeName,omitempty" xml:"TranscodeName,omitempty" require:"true"`
	TranscodeTemplate         *string                                                                                                 `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty" require:"true"`
	CustomTranscodeParameters *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters `json:"CustomTranscodeParameters,omitempty" xml:"CustomTranscodeParameters,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) SetTranscodeApp(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo {
	s.TranscodeApp = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) SetTranscodeName(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo {
	s.TranscodeName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) SetTranscodeTemplate(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo {
	s.TranscodeTemplate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo) SetCustomTranscodeParameters(v *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfo {
	s.CustomTranscodeParameters = v
	return s
}

type DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters struct {
	RtsFlag         *string `json:"RtsFlag,omitempty" xml:"RtsFlag,omitempty" require:"true"`
	Bframes         *string `json:"Bframes,omitempty" xml:"Bframes,omitempty" require:"true"`
	VideoBitrate    *int    `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty" require:"true"`
	FPS             *int    `json:"FPS,omitempty" xml:"FPS,omitempty" require:"true"`
	Height          *int    `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Width           *int    `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	TemplateType    *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty" require:"true"`
	VideoProfile    *string `json:"VideoProfile,omitempty" xml:"VideoProfile,omitempty" require:"true"`
	Gop             *string `json:"Gop,omitempty" xml:"Gop,omitempty" require:"true"`
	AudioBitrate    *int    `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty" require:"true"`
	AudioProfile    *string `json:"AudioProfile,omitempty" xml:"AudioProfile,omitempty" require:"true"`
	AudioCodec      *string `json:"AudioCodec,omitempty" xml:"AudioCodec,omitempty" require:"true"`
	AudioRate       *int    `json:"AudioRate,omitempty" xml:"AudioRate,omitempty" require:"true"`
	AudioChannelNum *int    `json:"AudioChannelNum,omitempty" xml:"AudioChannelNum,omitempty" require:"true"`
}

func (s DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetRtsFlag(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.RtsFlag = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetBframes(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Bframes = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetVideoBitrate(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.VideoBitrate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetFPS(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.FPS = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetHeight(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetWidth(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetTemplateType(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.TemplateType = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetVideoProfile(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.VideoProfile = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetGop(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.Gop = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioBitrate(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioBitrate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioProfile(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioProfile = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioCodec(v string) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioCodec = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioRate(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioRate = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters) SetAudioChannelNum(v int) *DescribeLiveStreamTranscodeInfoResponseDomainTranscodeListDomainTranscodeInfoCustomTranscodeParameters {
	s.AudioChannelNum = &v
	return s
}

type SetLiveStreamsNotifyUrlConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NotifyUrl  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty" require:"true"`
}

func (s SetLiveStreamsNotifyUrlConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLiveStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetDomainName(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigRequest) SetNotifyUrl(v string) *SetLiveStreamsNotifyUrlConfigRequest {
	s.NotifyUrl = &v
	return s
}

type SetLiveStreamsNotifyUrlConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetLiveStreamsNotifyUrlConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLiveStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) SetRequestId(v string) *SetLiveStreamsNotifyUrlConfigResponse {
	s.RequestId = &v
	return s
}

type ResumeLiveStreamRequest struct {
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty" require:"true"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName     *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
}

func (s ResumeLiveStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeLiveStreamRequest) GoString() string {
	return s.String()
}

func (s *ResumeLiveStreamRequest) SetSecurityToken(v string) *ResumeLiveStreamRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetDomainName(v string) *ResumeLiveStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetLiveStreamType(v string) *ResumeLiveStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetAppName(v string) *ResumeLiveStreamRequest {
	s.AppName = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetStreamName(v string) *ResumeLiveStreamRequest {
	s.StreamName = &v
	return s
}

type ResumeLiveStreamResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ResumeLiveStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeLiveStreamResponse) GoString() string {
	return s.String()
}

func (s *ResumeLiveStreamResponse) SetRequestId(v string) *ResumeLiveStreamResponse {
	s.RequestId = &v
	return s
}

type AddLiveAppSnapshotConfigRequest struct {
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	TimeInterval       *int    `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty" require:"true"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OverwriteOssObject *string `json:"OverwriteOssObject,omitempty" xml:"OverwriteOssObject,omitempty"`
	SequenceOssObject  *string `json:"SequenceOssObject,omitempty" xml:"SequenceOssObject,omitempty"`
	Callback           *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
}

func (s AddLiveAppSnapshotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAppSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAppSnapshotConfigRequest) SetSecurityToken(v string) *AddLiveAppSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetDomainName(v string) *AddLiveAppSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetAppName(v string) *AddLiveAppSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetTimeInterval(v int) *AddLiveAppSnapshotConfigRequest {
	s.TimeInterval = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOssEndpoint(v string) *AddLiveAppSnapshotConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOssBucket(v string) *AddLiveAppSnapshotConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOverwriteOssObject(v string) *AddLiveAppSnapshotConfigRequest {
	s.OverwriteOssObject = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetSequenceOssObject(v string) *AddLiveAppSnapshotConfigRequest {
	s.SequenceOssObject = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetCallback(v string) *AddLiveAppSnapshotConfigRequest {
	s.Callback = &v
	return s
}

type AddLiveAppSnapshotConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveAppSnapshotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAppSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAppSnapshotConfigResponse) SetRequestId(v string) *AddLiveAppSnapshotConfigResponse {
	s.RequestId = &v
	return s
}

type AddLiveAppRecordConfigRequest struct {
	SecurityToken *string                                      `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string                                      `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	OssEndpoint   *string                                      `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket     *string                                      `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	RecordFormat  []*AddLiveAppRecordConfigRequestRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" require:"true" type:"Repeated"`
	StreamName    *string                                      `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StartTime     *string                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OnDemand      *int                                         `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
}

func (s AddLiveAppRecordConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAppRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigRequest) SetSecurityToken(v string) *AddLiveAppRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetDomainName(v string) *AddLiveAppRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetAppName(v string) *AddLiveAppRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOssEndpoint(v string) *AddLiveAppRecordConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOssBucket(v string) *AddLiveAppRecordConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetRecordFormat(v []*AddLiveAppRecordConfigRequestRecordFormat) *AddLiveAppRecordConfigRequest {
	s.RecordFormat = v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetStreamName(v string) *AddLiveAppRecordConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetStartTime(v string) *AddLiveAppRecordConfigRequest {
	s.StartTime = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetEndTime(v string) *AddLiveAppRecordConfigRequest {
	s.EndTime = &v
	return s
}

func (s *AddLiveAppRecordConfigRequest) SetOnDemand(v int) *AddLiveAppRecordConfigRequest {
	s.OnDemand = &v
	return s
}

type AddLiveAppRecordConfigRequestRecordFormat struct {
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty"`
	OssObjectPrefix      *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty"`
	CycleDuration        *int    `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
}

func (s AddLiveAppRecordConfigRequestRecordFormat) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAppRecordConfigRequestRecordFormat) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetFormat(v string) *AddLiveAppRecordConfigRequestRecordFormat {
	s.Format = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetOssObjectPrefix(v string) *AddLiveAppRecordConfigRequestRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetSliceOssObjectPrefix(v string) *AddLiveAppRecordConfigRequestRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *AddLiveAppRecordConfigRequestRecordFormat) SetCycleDuration(v int) *AddLiveAppRecordConfigRequestRecordFormat {
	s.CycleDuration = &v
	return s
}

type AddLiveAppRecordConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLiveAppRecordConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveAppRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigResponse) SetRequestId(v string) *AddLiveAppRecordConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeLiveRecordConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	PageNum       *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DescribeLiveRecordConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigRequest) SetSecurityToken(v string) *DescribeLiveRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetDomainName(v string) *DescribeLiveRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetAppName(v string) *DescribeLiveRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetStreamName(v string) *DescribeLiveRecordConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetPageNum(v int) *DescribeLiveRecordConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetPageSize(v int) *DescribeLiveRecordConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordConfigRequest) SetOrder(v string) *DescribeLiveRecordConfigRequest {
	s.Order = &v
	return s
}

type DescribeLiveRecordConfigResponse struct {
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum           *int                                               `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize          *int                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Order             *string                                            `json:"Order,omitempty" xml:"Order,omitempty" require:"true"`
	TotalNum          *int                                               `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage         *int                                               `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	LiveAppRecordList *DescribeLiveRecordConfigResponseLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveRecordConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponse) SetRequestId(v string) *DescribeLiveRecordConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetPageNum(v int) *DescribeLiveRecordConfigResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetPageSize(v int) *DescribeLiveRecordConfigResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetOrder(v string) *DescribeLiveRecordConfigResponse {
	s.Order = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetTotalNum(v int) *DescribeLiveRecordConfigResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetTotalPage(v int) *DescribeLiveRecordConfigResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveRecordConfigResponse) SetLiveAppRecordList(v *DescribeLiveRecordConfigResponseLiveAppRecordList) *DescribeLiveRecordConfigResponse {
	s.LiveAppRecordList = v
	return s
}

type DescribeLiveRecordConfigResponseLiveAppRecordList struct {
	LiveAppRecord []*DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordList) SetLiveAppRecord(v []*DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) *DescribeLiveRecordConfigResponseLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

type DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord struct {
	DomainName       *string                                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName          *string                                                                         `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName       *string                                                                         `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssEndpoint      *string                                                                         `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket        *string                                                                         `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	CreateTime       *string                                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	StartTime        *string                                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime          *string                                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	OnDemond         *int                                                                            `json:"OnDemond,omitempty" xml:"OnDemond,omitempty" require:"true"`
	RecordFormatList *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList `json:"RecordFormatList,omitempty" xml:"RecordFormatList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetOssEndpoint(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetOssBucket(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetCreateTime(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetOnDemond(v int) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.OnDemond = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord) SetRecordFormatList(v *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecord {
	s.RecordFormatList = v
	return s
}

type DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList struct {
	RecordFormat []*DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList) SetRecordFormat(v []*DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatList {
	s.RecordFormat = v
	return s
}

type DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat struct {
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty" require:"true"`
	OssObjectPrefix      *string `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty" require:"true"`
	SliceOssObjectPrefix *string `json:"SliceOssObjectPrefix,omitempty" xml:"SliceOssObjectPrefix,omitempty" require:"true"`
	CycleDuration        *int    `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty" require:"true"`
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetFormat(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.Format = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetOssObjectPrefix(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetSliceOssObjectPrefix(v string) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.SliceOssObjectPrefix = &v
	return s
}

func (s *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat) SetCycleDuration(v int) *DescribeLiveRecordConfigResponseLiveAppRecordListLiveAppRecordRecordFormatListRecordFormat {
	s.CycleDuration = &v
	return s
}

type DeleteLiveAppSnapshotConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
}

func (s DeleteLiveAppSnapshotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAppSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetSecurityToken(v string) *DeleteLiveAppSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetDomainName(v string) *DeleteLiveAppSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetAppName(v string) *DeleteLiveAppSnapshotConfigRequest {
	s.AppName = &v
	return s
}

type DeleteLiveAppSnapshotConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveAppSnapshotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAppSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppSnapshotConfigResponse) SetRequestId(v string) *DeleteLiveAppSnapshotConfigResponse {
	s.RequestId = &v
	return s
}

type DeleteLiveAppRecordConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveAppRecordConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAppRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppRecordConfigRequest) SetSecurityToken(v string) *DeleteLiveAppRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetDomainName(v string) *DeleteLiveAppRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetAppName(v string) *DeleteLiveAppRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetStreamName(v string) *DeleteLiveAppRecordConfigRequest {
	s.StreamName = &v
	return s
}

type DeleteLiveAppRecordConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLiveAppRecordConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveAppRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppRecordConfigResponse) SetRequestId(v string) *DeleteLiveAppRecordConfigResponse {
	s.RequestId = &v
	return s
}

type CreateLiveStreamRecordIndexFilesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssEndpoint   *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket     *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssObject     *string `json:"OssObject,omitempty" xml:"OssObject,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s CreateLiveStreamRecordIndexFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetSecurityToken(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetDomainName(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.DomainName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetAppName(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.AppName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetStreamName(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.StreamName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOssEndpoint(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOssBucket(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOssObject(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.OssObject = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetStartTime(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.StartTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetEndTime(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.EndTime = &v
	return s
}

type CreateLiveStreamRecordIndexFilesResponse struct {
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RecordInfo *CreateLiveStreamRecordIndexFilesResponseRecordInfo `json:"RecordInfo,omitempty" xml:"RecordInfo,omitempty" require:"true" type:"Struct"`
}

func (s CreateLiveStreamRecordIndexFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesResponse) SetRequestId(v string) *CreateLiveStreamRecordIndexFilesResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponse) SetRecordInfo(v *CreateLiveStreamRecordIndexFilesResponseRecordInfo) *CreateLiveStreamRecordIndexFilesResponse {
	s.RecordInfo = v
	return s
}

type CreateLiveStreamRecordIndexFilesResponseRecordInfo struct {
	RecordId    *string  `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	RecordUrl   *string  `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty" require:"true"`
	DomainName  *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string  `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName  *string  `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssEndpoint *string  `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssObject   *string  `json:"OssObject,omitempty" xml:"OssObject,omitempty" require:"true"`
	StartTime   *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Duration    *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Height      *int     `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Width       *int     `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	CreateTime  *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s CreateLiveStreamRecordIndexFilesResponseRecordInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesResponseRecordInfo) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetRecordId(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.RecordId = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetRecordUrl(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.RecordUrl = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetDomainName(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.DomainName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetAppName(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.AppName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetStreamName(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.StreamName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetOssBucket(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.OssBucket = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetOssEndpoint(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.OssEndpoint = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetOssObject(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.OssObject = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetStartTime(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.StartTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetEndTime(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.EndTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetDuration(v float32) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.Duration = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetHeight(v int) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.Height = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetWidth(v int) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.Width = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseRecordInfo) SetCreateTime(v string) *CreateLiveStreamRecordIndexFilesResponseRecordInfo {
	s.CreateTime = &v
	return s
}

type DescribeLiveStreamSnapshotInfoRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Limit         *int    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DescribeLiveStreamSnapshotInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetSecurityToken(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetDomainName(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetAppName(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetStreamName(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetStartTime(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetEndTime(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetLimit(v int) *DescribeLiveStreamSnapshotInfoRequest {
	s.Limit = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetOrder(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.Order = &v
	return s
}

type DescribeLiveStreamSnapshotInfoResponse struct {
	RequestId                  *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextStartTime              *string                                                           `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty" require:"true"`
	LiveStreamSnapshotInfoList *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList `json:"LiveStreamSnapshotInfoList,omitempty" xml:"LiveStreamSnapshotInfoList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamSnapshotInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponse) SetRequestId(v string) *DescribeLiveStreamSnapshotInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponse) SetNextStartTime(v string) *DescribeLiveStreamSnapshotInfoResponse {
	s.NextStartTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponse) SetLiveStreamSnapshotInfoList(v *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList) *DescribeLiveStreamSnapshotInfoResponse {
	s.LiveStreamSnapshotInfoList = v
	return s
}

type DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList struct {
	LiveStreamSnapshotInfo []*DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo `json:"LiveStreamSnapshotInfo,omitempty" xml:"LiveStreamSnapshotInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList) SetLiveStreamSnapshotInfo(v []*DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoList {
	s.LiveStreamSnapshotInfo = v
	return s
}

type DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo struct {
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssObject   *string `json:"OssObject,omitempty" xml:"OssObject,omitempty" require:"true"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetOssEndpoint(v string) *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetOssBucket(v string) *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetOssObject(v string) *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetCreateTime(v string) *DescribeLiveStreamSnapshotInfoResponseLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.CreateTime = &v
	return s
}

type DescribeLiveStreamRecordIndexFilesRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNum       *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetSecurityToken(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetDomainName(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetAppName(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetStreamName(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetStartTime(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetEndTime(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetPageNum(v int) *DescribeLiveStreamRecordIndexFilesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetPageSize(v int) *DescribeLiveStreamRecordIndexFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesRequest) SetOrder(v string) *DescribeLiveStreamRecordIndexFilesRequest {
	s.Order = &v
	return s
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	RequestId           *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum             *int                                                           `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize            *int                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Order               *string                                                        `json:"Order,omitempty" xml:"Order,omitempty" require:"true"`
	TotalNum            *int                                                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage           *int                                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	RecordIndexInfoList *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList `json:"RecordIndexInfoList,omitempty" xml:"RecordIndexInfoList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamRecordIndexFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetRequestId(v string) *DescribeLiveStreamRecordIndexFilesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetPageNum(v int) *DescribeLiveStreamRecordIndexFilesResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetPageSize(v int) *DescribeLiveStreamRecordIndexFilesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetOrder(v string) *DescribeLiveStreamRecordIndexFilesResponse {
	s.Order = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetTotalNum(v int) *DescribeLiveStreamRecordIndexFilesResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetTotalPage(v int) *DescribeLiveStreamRecordIndexFilesResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetRecordIndexInfoList(v *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList) *DescribeLiveStreamRecordIndexFilesResponse {
	s.RecordIndexInfoList = v
	return s
}

type DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList struct {
	RecordIndexInfo []*DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo `json:"RecordIndexInfo,omitempty" xml:"RecordIndexInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList) SetRecordIndexInfo(v []*DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoList {
	s.RecordIndexInfo = v
	return s
}

type DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo struct {
	RecordId    *string  `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	RecordUrl   *string  `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty" require:"true"`
	DomainName  *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string  `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName  *string  `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssEndpoint *string  `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssObject   *string  `json:"OssObject,omitempty" xml:"OssObject,omitempty" require:"true"`
	StartTime   *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Duration    *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Height      *int     `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Width       *int     `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	CreateTime  *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetRecordId(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.RecordId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetRecordUrl(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.RecordUrl = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetDomainName(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetAppName(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetStreamName(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetOssBucket(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetOssEndpoint(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetOssObject(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetStartTime(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetEndTime(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetDuration(v float32) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.Duration = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetHeight(v int) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetWidth(v int) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo) SetCreateTime(v string) *DescribeLiveStreamRecordIndexFilesResponseRecordIndexInfoListRecordIndexInfo {
	s.CreateTime = &v
	return s
}

type DescribeLiveStreamRecordIndexFileRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	RecordId      *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
}

func (s DescribeLiveStreamRecordIndexFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetSecurityToken(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetDomainName(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetAppName(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetStreamName(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileRequest) SetRecordId(v string) *DescribeLiveStreamRecordIndexFileRequest {
	s.RecordId = &v
	return s
}

type DescribeLiveStreamRecordIndexFileResponse struct {
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RecordIndexInfo *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo `json:"RecordIndexInfo,omitempty" xml:"RecordIndexInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamRecordIndexFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileResponse) SetRequestId(v string) *DescribeLiveStreamRecordIndexFileResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponse) SetRecordIndexInfo(v *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) *DescribeLiveStreamRecordIndexFileResponse {
	s.RecordIndexInfo = v
	return s
}

type DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo struct {
	RecordId    *string  `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	RecordUrl   *string  `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty" require:"true"`
	DomainName  *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName     *string  `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName  *string  `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssEndpoint *string  `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssObject   *string  `json:"OssObject,omitempty" xml:"OssObject,omitempty" require:"true"`
	StartTime   *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Duration    *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Height      *int     `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Width       *int     `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	CreateTime  *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetRecordId(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.RecordId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetRecordUrl(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.RecordUrl = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetDomainName(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetAppName(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetStreamName(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetOssBucket(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetOssEndpoint(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetOssObject(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetStartTime(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetEndTime(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetDuration(v float32) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.Duration = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetHeight(v int) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetWidth(v int) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo) SetCreateTime(v string) *DescribeLiveStreamRecordIndexFileResponseRecordIndexInfo {
	s.CreateTime = &v
	return s
}

type DescribeLiveStreamRecordContentRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	StreamName    *string `json:"StreamName,omitempty" xml:"StreamName,omitempty" require:"true"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeLiveStreamRecordContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentRequest) SetSecurityToken(v string) *DescribeLiveStreamRecordContentRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetDomainName(v string) *DescribeLiveStreamRecordContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetAppName(v string) *DescribeLiveStreamRecordContentRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetStreamName(v string) *DescribeLiveStreamRecordContentRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetStartTime(v string) *DescribeLiveStreamRecordContentRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentRequest) SetEndTime(v string) *DescribeLiveStreamRecordContentRequest {
	s.EndTime = &v
	return s
}

type DescribeLiveStreamRecordContentResponse struct {
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RecordContentInfoList *DescribeLiveStreamRecordContentResponseRecordContentInfoList `json:"RecordContentInfoList,omitempty" xml:"RecordContentInfoList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveStreamRecordContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponse) SetRequestId(v string) *DescribeLiveStreamRecordContentResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponse) SetRecordContentInfoList(v *DescribeLiveStreamRecordContentResponseRecordContentInfoList) *DescribeLiveStreamRecordContentResponse {
	s.RecordContentInfoList = v
	return s
}

type DescribeLiveStreamRecordContentResponseRecordContentInfoList struct {
	RecordContentInfo []*DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo `json:"RecordContentInfo,omitempty" xml:"RecordContentInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveStreamRecordContentResponseRecordContentInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponseRecordContentInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoList) SetRecordContentInfo(v []*DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) *DescribeLiveStreamRecordContentResponseRecordContentInfoList {
	s.RecordContentInfo = v
	return s
}

type DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo struct {
	OssEndpoint     *string  `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket       *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OssObjectPrefix *string  `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty" require:"true"`
	StartTime       *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Duration        *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
}

func (s DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) SetOssEndpoint(v string) *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) SetOssBucket(v string) *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) SetOssObjectPrefix(v string) *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) SetStartTime(v string) *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) SetEndTime(v string) *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo) SetDuration(v float32) *DescribeLiveStreamRecordContentResponseRecordContentInfoListRecordContentInfo {
	s.Duration = &v
	return s
}

type DescribeLiveSnapshotConfigRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	PageNum       *int    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DescribeLiveSnapshotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigRequest) SetSecurityToken(v string) *DescribeLiveSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetDomainName(v string) *DescribeLiveSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetAppName(v string) *DescribeLiveSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetPageNum(v int) *DescribeLiveSnapshotConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetPageSize(v int) *DescribeLiveSnapshotConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotConfigRequest) SetOrder(v string) *DescribeLiveSnapshotConfigRequest {
	s.Order = &v
	return s
}

type DescribeLiveSnapshotConfigResponse struct {
	RequestId                    *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNum                      *int                                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty" require:"true"`
	PageSize                     *int                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Order                        *string                                                         `json:"Order,omitempty" xml:"Order,omitempty" require:"true"`
	TotalNum                     *int                                                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty" require:"true"`
	TotalPage                    *int                                                            `json:"TotalPage,omitempty" xml:"TotalPage,omitempty" require:"true"`
	LiveStreamSnapshotConfigList *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList `json:"LiveStreamSnapshotConfigList,omitempty" xml:"LiveStreamSnapshotConfigList,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLiveSnapshotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponse) SetRequestId(v string) *DescribeLiveSnapshotConfigResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetPageNum(v int) *DescribeLiveSnapshotConfigResponse {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetPageSize(v int) *DescribeLiveSnapshotConfigResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetOrder(v string) *DescribeLiveSnapshotConfigResponse {
	s.Order = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetTotalNum(v int) *DescribeLiveSnapshotConfigResponse {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetTotalPage(v int) *DescribeLiveSnapshotConfigResponse {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetLiveStreamSnapshotConfigList(v *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList) *DescribeLiveSnapshotConfigResponse {
	s.LiveStreamSnapshotConfigList = v
	return s
}

type DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList struct {
	LiveStreamSnapshotConfig []*DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig `json:"LiveStreamSnapshotConfig,omitempty" xml:"LiveStreamSnapshotConfig,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList) SetLiveStreamSnapshotConfig(v []*DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigList {
	s.LiveStreamSnapshotConfig = v
	return s
}

type DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig struct {
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	TimeInterval       *int    `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty" require:"true"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty" require:"true"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty" require:"true"`
	OverwriteOssObject *string `json:"OverwriteOssObject,omitempty" xml:"OverwriteOssObject,omitempty" require:"true"`
	SequenceOssObject  *string `json:"SequenceOssObject,omitempty" xml:"SequenceOssObject,omitempty" require:"true"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Callback           *string `json:"Callback,omitempty" xml:"Callback,omitempty" require:"true"`
}

func (s DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetDomainName(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetAppName(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetTimeInterval(v int) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.TimeInterval = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetOssEndpoint(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetOssBucket(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetOverwriteOssObject(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.OverwriteOssObject = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetSequenceOssObject(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.SequenceOssObject = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetCreateTime(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetCallback(v string) *DescribeLiveSnapshotConfigResponseLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.Callback = &v
	return s
}

type UpdateLiveAppSnapshotConfigRequest struct {
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	TimeInterval       *int    `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OverwriteOssObject *string `json:"OverwriteOssObject,omitempty" xml:"OverwriteOssObject,omitempty"`
	SequenceOssObject  *string `json:"SequenceOssObject,omitempty" xml:"SequenceOssObject,omitempty"`
	Callback           *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
}

func (s UpdateLiveAppSnapshotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveAppSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetSecurityToken(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetDomainName(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetAppName(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetTimeInterval(v int) *UpdateLiveAppSnapshotConfigRequest {
	s.TimeInterval = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOssEndpoint(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOssBucket(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOverwriteOssObject(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.OverwriteOssObject = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetSequenceOssObject(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.SequenceOssObject = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetCallback(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.Callback = &v
	return s
}

type UpdateLiveAppSnapshotConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateLiveAppSnapshotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveAppSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppSnapshotConfigResponse) SetRequestId(v string) *UpdateLiveAppSnapshotConfigResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("live.aliyuncs.com"),
		"cn-beijing":                  tea.String("live.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("live.aliyuncs.com"),
		"cn-shanghai":                 tea.String("live.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("live.aliyuncs.com"),
		"ap-southeast-1":              tea.String("live.aliyuncs.com"),
		"ap-southeast-5":              tea.String("live.aliyuncs.com"),
		"ap-northeast-1":              tea.String("live.aliyuncs.com"),
		"eu-central-1":                tea.String("live.aliyuncs.com"),
		"ap-south-1":                  tea.String("live.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("live.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("live.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("live.ap-southeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("live.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("live.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("live.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("live.aliyuncs.com"),
		"cn-chengdu":                  tea.String("live.aliyuncs.com"),
		"cn-edge-1":                   tea.String("live.aliyuncs.com"),
		"cn-fujian":                   tea.String("live.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("live.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("live.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("live.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("live.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("live.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("live.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("live.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("live.aliyuncs.com"),
		"cn-hongkong":                 tea.String("live.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("live.aliyuncs.com"),
		"cn-huhehaote":                tea.String("live.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("live.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("live.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("live.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("live.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("live.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("live.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("live.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("live.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("live.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("live.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("live.aliyuncs.com"),
		"cn-wuhan":                    tea.String("live.aliyuncs.com"),
		"cn-yushanfang":               tea.String("live.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("live.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("live.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("live.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("live.aliyuncs.com"),
		"eu-west-1":                   tea.String("live.ap-southeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("live.ap-southeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("live.ap-southeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("live.ap-southeast-1.aliyuncs.com"),
		"us-east-1":                   tea.String("live.ap-southeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("live.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("live"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) AddLiveASRConfigWithOptions(request *AddLiveASRConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveASRConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveASRConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveASRConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveASRConfig(request *AddLiveASRConfigRequest) (_result *AddLiveASRConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveASRConfigResponse{}
	_body, _err := client.AddLiveASRConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveAsrConfigWithOptions(request *DescribeLiveAsrConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveAsrConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveAsrConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveAsrConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveAsrConfig(request *DescribeLiveAsrConfigRequest) (_result *DescribeLiveAsrConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveAsrConfigResponse{}
	_body, _err := client.DescribeLiveAsrConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveASRConfigWithOptions(request *DeleteLiveASRConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveASRConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveASRConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveASRConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveASRConfig(request *DeleteLiveASRConfigRequest) (_result *DeleteLiveASRConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveASRConfigResponse{}
	_body, _err := client.DeleteLiveASRConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveASRConfigWithOptions(request *UpdateLiveASRConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveASRConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveASRConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveASRConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveASRConfig(request *UpdateLiveASRConfigRequest) (_result *UpdateLiveASRConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveASRConfigResponse{}
	_body, _err := client.UpdateLiveASRConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMixStreamWithOptions(request *DeleteMixStreamRequest, runtime *util.RuntimeOptions) (_result *DeleteMixStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteMixStreamResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteMixStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMixStream(request *DeleteMixStreamRequest) (_result *DeleteMixStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMixStreamResponse{}
	_body, _err := client.DeleteMixStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMixStreamWithOptions(request *UpdateMixStreamRequest, runtime *util.RuntimeOptions) (_result *UpdateMixStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateMixStreamResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateMixStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMixStream(request *UpdateMixStreamRequest) (_result *UpdateMixStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMixStreamResponse{}
	_body, _err := client.UpdateMixStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMixStreamWithOptions(request *CreateMixStreamRequest, runtime *util.RuntimeOptions) (_result *CreateMixStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateMixStreamResponse{}
	_body, _err := client.DoRequest(tea.String("CreateMixStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMixStream(request *CreateMixStreamRequest) (_result *CreateMixStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMixStreamResponse{}
	_body, _err := client.CreateMixStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMixStreamListWithOptions(request *DescribeMixStreamListRequest, runtime *util.RuntimeOptions) (_result *DescribeMixStreamListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeMixStreamListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeMixStreamList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMixStreamList(request *DescribeMixStreamListRequest) (_result *DescribeMixStreamListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMixStreamListResponse{}
	_body, _err := client.DescribeMixStreamListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRtsLiveStreamTranscodeWithOptions(request *AddRtsLiveStreamTranscodeRequest, runtime *util.RuntimeOptions) (_result *AddRtsLiveStreamTranscodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddRtsLiveStreamTranscodeResponse{}
	_body, _err := client.DoRequest(tea.String("AddRtsLiveStreamTranscode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRtsLiveStreamTranscode(request *AddRtsLiveStreamTranscodeRequest) (_result *AddRtsLiveStreamTranscodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRtsLiveStreamTranscodeResponse{}
	_body, _err := client.AddRtsLiveStreamTranscodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainTimeShiftDataWithOptions(request *DescribeLiveDomainTimeShiftDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainTimeShiftDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainTimeShiftDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainTimeShiftData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainTimeShiftData(request *DescribeLiveDomainTimeShiftDataRequest) (_result *DescribeLiveDomainTimeShiftDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainTimeShiftDataResponse{}
	_body, _err := client.DescribeLiveDomainTimeShiftDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHtmlResourceWithOptions(request *DeleteHtmlResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteHtmlResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteHtmlResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteHtmlResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHtmlResource(request *DeleteHtmlResourceRequest) (_result *DeleteHtmlResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHtmlResourceResponse{}
	_body, _err := client.DeleteHtmlResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHtmlResourceWithOptions(request *DescribeHtmlResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeHtmlResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeHtmlResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeHtmlResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHtmlResource(request *DescribeHtmlResourceRequest) (_result *DescribeHtmlResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHtmlResourceResponse{}
	_body, _err := client.DescribeHtmlResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ControlHtmlResourceWithOptions(request *ControlHtmlResourceRequest, runtime *util.RuntimeOptions) (_result *ControlHtmlResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ControlHtmlResourceResponse{}
	_body, _err := client.DoRequest(tea.String("ControlHtmlResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ControlHtmlResource(request *ControlHtmlResourceRequest) (_result *ControlHtmlResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ControlHtmlResourceResponse{}
	_body, _err := client.ControlHtmlResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditHtmlResourceWithOptions(request *EditHtmlResourceRequest, runtime *util.RuntimeOptions) (_result *EditHtmlResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EditHtmlResourceResponse{}
	_body, _err := client.DoRequest(tea.String("EditHtmlResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditHtmlResource(request *EditHtmlResourceRequest) (_result *EditHtmlResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditHtmlResourceResponse{}
	_body, _err := client.EditHtmlResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveUserTagsWithOptions(request *DescribeLiveUserTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveUserTagsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveUserTags"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveUserTags(request *DescribeLiveUserTagsRequest) (_result *DescribeLiveUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveUserTagsResponse{}
	_body, _err := client.DescribeLiveUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagLiveResourcesWithOptions(request *UnTagLiveResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagLiveResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnTagLiveResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UnTagLiveResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagLiveResources(request *UnTagLiveResourcesRequest) (_result *UnTagLiveResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagLiveResourcesResponse{}
	_body, _err := client.UnTagLiveResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagLiveResourcesWithOptions(request *TagLiveResourcesRequest, runtime *util.RuntimeOptions) (_result *TagLiveResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagLiveResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagLiveResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagLiveResources(request *TagLiveResourcesRequest) (_result *TagLiveResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagLiveResourcesResponse{}
	_body, _err := client.TagLiveResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveTagResourcesWithOptions(request *DescribeLiveTagResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveTagResources(request *DescribeLiveTagResourcesRequest) (_result *DescribeLiveTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveTagResourcesResponse{}
	_body, _err := client.DescribeLiveTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveAudioAuditConfigWithOptions(request *AddLiveAudioAuditConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveAudioAuditConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveAudioAuditConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveAudioAuditConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveAudioAuditConfig(request *AddLiveAudioAuditConfigRequest) (_result *AddLiveAudioAuditConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveAudioAuditConfigResponse{}
	_body, _err := client.AddLiveAudioAuditConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveAudioAuditConfigWithOptions(request *DeleteLiveAudioAuditConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveAudioAuditConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveAudioAuditConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveAudioAuditConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveAudioAuditConfig(request *DeleteLiveAudioAuditConfigRequest) (_result *DeleteLiveAudioAuditConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveAudioAuditConfigResponse{}
	_body, _err := client.DeleteLiveAudioAuditConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveAudioAuditConfigWithOptions(request *DescribeLiveAudioAuditConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveAudioAuditConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveAudioAuditConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveAudioAuditConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveAudioAuditConfig(request *DescribeLiveAudioAuditConfigRequest) (_result *DescribeLiveAudioAuditConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveAudioAuditConfigResponse{}
	_body, _err := client.DescribeLiveAudioAuditConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveAudioAuditConfigWithOptions(request *UpdateLiveAudioAuditConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveAudioAuditConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveAudioAuditConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveAudioAuditConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveAudioAuditConfig(request *UpdateLiveAudioAuditConfigRequest) (_result *UpdateLiveAudioAuditConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveAudioAuditConfigResponse{}
	_body, _err := client.UpdateLiveAudioAuditConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveAudioAuditNotifyConfigWithOptions(request *AddLiveAudioAuditNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveAudioAuditNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveAudioAuditNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveAudioAuditNotifyConfig(request *AddLiveAudioAuditNotifyConfigRequest) (_result *AddLiveAudioAuditNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.AddLiveAudioAuditNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveAudioAuditNotifyConfigWithOptions(request *DeleteLiveAudioAuditNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveAudioAuditNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveAudioAuditNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveAudioAuditNotifyConfig(request *DeleteLiveAudioAuditNotifyConfigRequest) (_result *DeleteLiveAudioAuditNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.DeleteLiveAudioAuditNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveAudioAuditNotifyConfigWithOptions(request *DescribeLiveAudioAuditNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveAudioAuditNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveAudioAuditNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveAudioAuditNotifyConfig(request *DescribeLiveAudioAuditNotifyConfigRequest) (_result *DescribeLiveAudioAuditNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.DescribeLiveAudioAuditNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveAudioAuditNotifyConfigWithOptions(request *UpdateLiveAudioAuditNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveAudioAuditNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveAudioAuditNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveAudioAuditNotifyConfig(request *UpdateLiveAudioAuditNotifyConfigRequest) (_result *UpdateLiveAudioAuditNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveAudioAuditNotifyConfigResponse{}
	_body, _err := client.UpdateLiveAudioAuditNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainPushTrafficDataWithOptions(request *DescribeLiveDomainPushTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainPushTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainPushTrafficDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainPushTrafficData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainPushTrafficData(request *DescribeLiveDomainPushTrafficDataRequest) (_result *DescribeLiveDomainPushTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainPushTrafficDataResponse{}
	_body, _err := client.DescribeLiveDomainPushTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainPushBpsDataWithOptions(request *DescribeLiveDomainPushBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainPushBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainPushBpsDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainPushBpsData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainPushBpsData(request *DescribeLiveDomainPushBpsDataRequest) (_result *DescribeLiveDomainPushBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainPushBpsDataResponse{}
	_body, _err := client.DescribeLiveDomainPushBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCasterSyncGroupWithOptions(request *SetCasterSyncGroupRequest, runtime *util.RuntimeOptions) (_result *SetCasterSyncGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetCasterSyncGroupResponse{}
	_body, _err := client.DoRequest(tea.String("SetCasterSyncGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCasterSyncGroup(request *SetCasterSyncGroupRequest) (_result *SetCasterSyncGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCasterSyncGroupResponse{}
	_body, _err := client.SetCasterSyncGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterSyncGroupWithOptions(request *DescribeCasterSyncGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterSyncGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterSyncGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterSyncGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterSyncGroup(request *DescribeCasterSyncGroupRequest) (_result *DescribeCasterSyncGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterSyncGroupResponse{}
	_body, _err := client.DescribeCasterSyncGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDetectPornDataWithOptions(request *DescribeLiveDetectPornDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDetectPornDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDetectPornDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDetectPornData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDetectPornData(request *DescribeLiveDetectPornDataRequest) (_result *DescribeLiveDetectPornDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDetectPornDataResponse{}
	_body, _err := client.DescribeLiveDetectPornDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveRealTimeLogLogstoreWithOptions(request *DeleteLiveRealTimeLogLogstoreRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveRealTimeLogLogstoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveRealTimeLogLogstoreResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveRealTimeLogLogstore"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveRealTimeLogLogstore(request *DeleteLiveRealTimeLogLogstoreRequest) (_result *DeleteLiveRealTimeLogLogstoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveRealTimeLogLogstoreResponse{}
	_body, _err := client.DeleteLiveRealTimeLogLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLiveRealtimeLogDeliveryWithOptions(request *DisableLiveRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DisableLiveRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("DisableLiveRealtimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLiveRealtimeLogDelivery(request *DisableLiveRealtimeLogDeliveryRequest) (_result *DisableLiveRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DisableLiveRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLiveRealtimeLogDeliveryWithOptions(request *EnableLiveRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *EnableLiveRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("EnableLiveRealtimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLiveRealtimeLogDelivery(request *EnableLiveRealtimeLogDeliveryRequest) (_result *EnableLiveRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.EnableLiveRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRealtimeLogDeliveryDomainsWithOptions(request *ListLiveRealtimeLogDeliveryDomainsRequest, runtime *util.RuntimeOptions) (_result *ListLiveRealtimeLogDeliveryDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListLiveRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.DoRequest(tea.String("ListLiveRealtimeLogDeliveryDomains"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRealtimeLogDeliveryDomains(request *ListLiveRealtimeLogDeliveryDomainsRequest) (_result *ListLiveRealtimeLogDeliveryDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.ListLiveRealtimeLogDeliveryDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLiveRealtimeLogDeliveryWithOptions(request *ModifyLiveRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *ModifyLiveRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyLiveRealtimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLiveRealtimeLogDelivery(request *ModifyLiveRealtimeLogDeliveryRequest) (_result *ModifyLiveRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.ModifyLiveRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveRealtimeDeliveryAccWithOptions(request *DescribeLiveRealtimeDeliveryAccRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveRealtimeDeliveryAccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveRealtimeDeliveryAccResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveRealtimeDeliveryAcc"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveRealtimeDeliveryAcc(request *DescribeLiveRealtimeDeliveryAccRequest) (_result *DescribeLiveRealtimeDeliveryAccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveRealtimeDeliveryAccResponse{}
	_body, _err := client.DescribeLiveRealtimeDeliveryAccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveRealtimeLogAuthorizedWithOptions(request *DescribeLiveRealtimeLogAuthorizedRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveRealtimeLogAuthorizedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveRealtimeLogAuthorizedResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveRealtimeLogAuthorized"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveRealtimeLogAuthorized(request *DescribeLiveRealtimeLogAuthorizedRequest) (_result *DescribeLiveRealtimeLogAuthorizedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveRealtimeLogAuthorizedResponse{}
	_body, _err := client.DescribeLiveRealtimeLogAuthorizedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRealtimeLogDeliveryWithOptions(request *ListLiveRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *ListLiveRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("ListLiveRealtimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRealtimeLogDelivery(request *ListLiveRealtimeLogDeliveryRequest) (_result *ListLiveRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.ListLiveRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRealtimeLogDeliveryInfosWithOptions(request *ListLiveRealtimeLogDeliveryInfosRequest, runtime *util.RuntimeOptions) (_result *ListLiveRealtimeLogDeliveryInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListLiveRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.DoRequest(tea.String("ListLiveRealtimeLogDeliveryInfos"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRealtimeLogDeliveryInfos(request *ListLiveRealtimeLogDeliveryInfosRequest) (_result *ListLiveRealtimeLogDeliveryInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.ListLiveRealtimeLogDeliveryInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealtimeLogDeliveryWithOptions(request *DescribeLiveDomainRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainRealtimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealtimeLogDelivery(request *DescribeLiveDomainRealtimeLogDeliveryRequest) (_result *DescribeLiveDomainRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DescribeLiveDomainRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveRealtimeLogDeliveryWithOptions(request *DeleteLiveRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveRealtimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveRealtimeLogDelivery(request *DeleteLiveRealtimeLogDeliveryRequest) (_result *DeleteLiveRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveRealtimeLogDeliveryResponse{}
	_body, _err := client.DeleteLiveRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveRealTimeLogDeliveryWithOptions(request *CreateLiveRealTimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *CreateLiveRealTimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateLiveRealTimeLogDeliveryResponse{}
	_body, _err := client.DoRequest(tea.String("CreateLiveRealTimeLogDelivery"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLiveRealTimeLogDelivery(request *CreateLiveRealTimeLogDeliveryRequest) (_result *CreateLiveRealTimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveRealTimeLogDeliveryResponse{}
	_body, _err := client.CreateLiveRealTimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainLimitWithOptions(request *DescribeLiveDomainLimitRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainLimitResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainLimit"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainLimit(request *DescribeLiveDomainLimitRequest) (_result *DescribeLiveDomainLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainLimitResponse{}
	_body, _err := client.DescribeLiveDomainLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainBpsDataByTimeStampWithOptions(request *DescribeLiveDomainBpsDataByTimeStampRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainBpsDataByTimeStampResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainBpsDataByTimeStampResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainBpsDataByTimeStamp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainBpsDataByTimeStamp(request *DescribeLiveDomainBpsDataByTimeStampRequest) (_result *DescribeLiveDomainBpsDataByTimeStampResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainBpsDataByTimeStampResponse{}
	_body, _err := client.DescribeLiveDomainBpsDataByTimeStampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamTranscodeStreamNumWithOptions(request *DescribeLiveStreamTranscodeStreamNumRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamTranscodeStreamNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamTranscodeStreamNumResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamTranscodeStreamNum"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamTranscodeStreamNum(request *DescribeLiveStreamTranscodeStreamNumRequest) (_result *DescribeLiveStreamTranscodeStreamNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamTranscodeStreamNumResponse{}
	_body, _err := client.DescribeLiveStreamTranscodeStreamNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveTopLevelDomainWithOptions(request *UpdateLiveTopLevelDomainRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveTopLevelDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveTopLevelDomainResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveTopLevelDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveTopLevelDomain(request *UpdateLiveTopLevelDomainRequest) (_result *UpdateLiveTopLevelDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveTopLevelDomainResponse{}
	_body, _err := client.UpdateLiveTopLevelDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainCertificateInfoWithOptions(request *DescribeLiveDomainCertificateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainCertificateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainCertificateInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainCertificateInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainCertificateInfo(request *DescribeLiveDomainCertificateInfoRequest) (_result *DescribeLiveDomainCertificateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainCertificateInfoResponse{}
	_body, _err := client.DescribeLiveDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLiveDomainSchdmByPropertyWithOptions(request *ModifyLiveDomainSchdmByPropertyRequest, runtime *util.RuntimeOptions) (_result *ModifyLiveDomainSchdmByPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyLiveDomainSchdmByPropertyResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyLiveDomainSchdmByProperty"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLiveDomainSchdmByProperty(request *ModifyLiveDomainSchdmByPropertyRequest) (_result *ModifyLiveDomainSchdmByPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLiveDomainSchdmByPropertyResponse{}
	_body, _err := client.ModifyLiveDomainSchdmByPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLiveStreamOptimizedFeatureConfigWithOptions(request *SetLiveStreamOptimizedFeatureConfigRequest, runtime *util.RuntimeOptions) (_result *SetLiveStreamOptimizedFeatureConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetLiveStreamOptimizedFeatureConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SetLiveStreamOptimizedFeatureConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLiveStreamOptimizedFeatureConfig(request *SetLiveStreamOptimizedFeatureConfigRequest) (_result *SetLiveStreamOptimizedFeatureConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLiveStreamOptimizedFeatureConfigResponse{}
	_body, _err := client.SetLiveStreamOptimizedFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamOptimizedFeatureConfigWithOptions(request *DescribeLiveStreamOptimizedFeatureConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamOptimizedFeatureConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamOptimizedFeatureConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamOptimizedFeatureConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamOptimizedFeatureConfig(request *DescribeLiveStreamOptimizedFeatureConfigRequest) (_result *DescribeLiveStreamOptimizedFeatureConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamOptimizedFeatureConfigResponse{}
	_body, _err := client.DescribeLiveStreamOptimizedFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLiveStreamDelayConfigWithOptions(request *SetLiveStreamDelayConfigRequest, runtime *util.RuntimeOptions) (_result *SetLiveStreamDelayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetLiveStreamDelayConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SetLiveStreamDelayConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLiveStreamDelayConfig(request *SetLiveStreamDelayConfigRequest) (_result *SetLiveStreamDelayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLiveStreamDelayConfigResponse{}
	_body, _err := client.SetLiveStreamDelayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamDelayConfigWithOptions(request *DescribeLiveStreamDelayConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamDelayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamDelayConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamDelayConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamDelayConfig(request *DescribeLiveStreamDelayConfigRequest) (_result *DescribeLiveStreamDelayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamDelayConfigResponse{}
	_body, _err := client.DescribeLiveStreamDelayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainOnlineUserNumWithOptions(request *DescribeLiveDomainOnlineUserNumRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainOnlineUserNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainOnlineUserNumResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainOnlineUserNum"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainOnlineUserNum(request *DescribeLiveDomainOnlineUserNumRequest) (_result *DescribeLiveDomainOnlineUserNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainOnlineUserNumResponse{}
	_body, _err := client.DescribeLiveDomainOnlineUserNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainFrameRateAndBitRateDataWithOptions(request *DescribeLiveDomainFrameRateAndBitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainFrameRateAndBitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainFrameRateAndBitRateDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainFrameRateAndBitRateData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainFrameRateAndBitRateData(request *DescribeLiveDomainFrameRateAndBitRateDataRequest) (_result *DescribeLiveDomainFrameRateAndBitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainFrameRateAndBitRateDataResponse{}
	_body, _err := client.DescribeLiveDomainFrameRateAndBitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetBoardCallbackWithOptions(request *SetBoardCallbackRequest, runtime *util.RuntimeOptions) (_result *SetBoardCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetBoardCallbackResponse{}
	_body, _err := client.DoRequest(tea.String("SetBoardCallback"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetBoardCallback(request *SetBoardCallbackRequest) (_result *SetBoardCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetBoardCallbackResponse{}
	_body, _err := client.SetBoardCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordsWithOptions(request *DescribeRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecords(request *DescribeRecordsRequest) (_result *DescribeRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordsResponse{}
	_body, _err := client.DescribeRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordWithOptions(request *DescribeRecordRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRecordResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRecord"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecord(request *DescribeRecordRequest) (_result *DescribeRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordResponse{}
	_body, _err := client.DescribeRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteBoardRecordWithOptions(request *CompleteBoardRecordRequest, runtime *util.RuntimeOptions) (_result *CompleteBoardRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompleteBoardRecordResponse{}
	_body, _err := client.DoRequest(tea.String("CompleteBoardRecord"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteBoardRecord(request *CompleteBoardRecordRequest) (_result *CompleteBoardRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteBoardRecordResponse{}
	_body, _err := client.CompleteBoardRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBoardRecordWithOptions(request *StartBoardRecordRequest, runtime *util.RuntimeOptions) (_result *StartBoardRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartBoardRecordResponse{}
	_body, _err := client.DoRequest(tea.String("StartBoardRecord"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBoardRecord(request *StartBoardRecordRequest) (_result *StartBoardRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartBoardRecordResponse{}
	_body, _err := client.StartBoardRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyRecordTokenWithOptions(request *ApplyRecordTokenRequest, runtime *util.RuntimeOptions) (_result *ApplyRecordTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyRecordTokenResponse{}
	_body, _err := client.DoRequest(tea.String("ApplyRecordToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyRecordToken(request *ApplyRecordTokenRequest) (_result *ApplyRecordTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyRecordTokenResponse{}
	_body, _err := client.ApplyRecordTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBoardCallbackWithOptions(request *UpdateBoardCallbackRequest, runtime *util.RuntimeOptions) (_result *UpdateBoardCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateBoardCallbackResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateBoardCallback"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBoardCallback(request *UpdateBoardCallbackRequest) (_result *UpdateBoardCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBoardCallbackResponse{}
	_body, _err := client.UpdateBoardCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainMappingWithOptions(request *DescribeLiveDomainMappingRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainMappingResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainMapping"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainMapping(request *DescribeLiveDomainMappingRequest) (_result *DescribeLiveDomainMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainMappingResponse{}
	_body, _err := client.DescribeLiveDomainMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLiveIndexWithOptions(request *StopLiveIndexRequest, runtime *util.RuntimeOptions) (_result *StopLiveIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopLiveIndexResponse{}
	_body, _err := client.DoRequest(tea.String("StopLiveIndex"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLiveIndex(request *StopLiveIndexRequest) (_result *StopLiveIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveIndexResponse{}
	_body, _err := client.StopLiveIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartLiveIndexWithOptions(request *StartLiveIndexRequest, runtime *util.RuntimeOptions) (_result *StartLiveIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartLiveIndexResponse{}
	_body, _err := client.DoRequest(tea.String("StartLiveIndex"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartLiveIndex(request *StartLiveIndexRequest) (_result *StartLiveIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartLiveIndexResponse{}
	_body, _err := client.StartLiveIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RealTimeSnapshotCommandWithOptions(request *RealTimeSnapshotCommandRequest, runtime *util.RuntimeOptions) (_result *RealTimeSnapshotCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RealTimeSnapshotCommandResponse{}
	_body, _err := client.DoRequest(tea.String("RealTimeSnapshotCommand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RealTimeSnapshotCommand(request *RealTimeSnapshotCommandRequest) (_result *RealTimeSnapshotCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RealTimeSnapshotCommandResponse{}
	_body, _err := client.RealTimeSnapshotCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveTopDomainsByFlowWithOptions(request *DescribeLiveTopDomainsByFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveTopDomainsByFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveTopDomainsByFlowResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveTopDomainsByFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveTopDomainsByFlow(request *DescribeLiveTopDomainsByFlowRequest) (_result *DescribeLiveTopDomainsByFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveTopDomainsByFlowResponse{}
	_body, _err := client.DescribeLiveTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealTimeBpsDataWithOptions(request *DescribeLiveDomainRealTimeBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainRealTimeBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainRealTimeBpsDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainRealTimeBpsData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealTimeBpsData(request *DescribeLiveDomainRealTimeBpsDataRequest) (_result *DescribeLiveDomainRealTimeBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainRealTimeBpsDataResponse{}
	_body, _err := client.DescribeLiveDomainRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealTimeHttpCodeDataWithOptions(request *DescribeLiveDomainRealTimeHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainRealTimeHttpCodeData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealTimeHttpCodeData(request *DescribeLiveDomainRealTimeHttpCodeDataRequest) (_result *DescribeLiveDomainRealTimeHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DescribeLiveDomainRealTimeHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealTimeTrafficDataWithOptions(request *DescribeLiveDomainRealTimeTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainRealTimeTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainRealTimeTrafficData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainRealTimeTrafficData(request *DescribeLiveDomainRealTimeTrafficDataRequest) (_result *DescribeLiveDomainRealTimeTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DescribeLiveDomainRealTimeTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveDomainPlayMappingWithOptions(request *AddLiveDomainPlayMappingRequest, runtime *util.RuntimeOptions) (_result *AddLiveDomainPlayMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveDomainPlayMappingResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveDomainPlayMapping"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveDomainPlayMapping(request *AddLiveDomainPlayMappingRequest) (_result *AddLiveDomainPlayMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveDomainPlayMappingResponse{}
	_body, _err := client.AddLiveDomainPlayMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveLazyPullStreamInfoConfigWithOptions(request *DeleteLiveLazyPullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveLazyPullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveLazyPullStreamInfoConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveLazyPullStreamInfoConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveLazyPullStreamInfoConfig(request *DeleteLiveLazyPullStreamInfoConfigRequest) (_result *DeleteLiveLazyPullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveLazyPullStreamInfoConfigResponse{}
	_body, _err := client.DeleteLiveLazyPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveLazyPullStreamConfigWithOptions(request *DescribeLiveLazyPullStreamConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveLazyPullStreamConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveLazyPullStreamConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveLazyPullStreamConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveLazyPullStreamConfig(request *DescribeLiveLazyPullStreamConfigRequest) (_result *DescribeLiveLazyPullStreamConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveLazyPullStreamConfigResponse{}
	_body, _err := client.DescribeLiveLazyPullStreamConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLiveLazyPullStreamInfoConfigWithOptions(request *SetLiveLazyPullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *SetLiveLazyPullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetLiveLazyPullStreamInfoConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SetLiveLazyPullStreamInfoConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLiveLazyPullStreamInfoConfig(request *SetLiveLazyPullStreamInfoConfigRequest) (_result *SetLiveLazyPullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLiveLazyPullStreamInfoConfigResponse{}
	_body, _err := client.SetLiveLazyPullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCasterSceneAudioWithOptions(request *UpdateCasterSceneAudioRequest, runtime *util.RuntimeOptions) (_result *UpdateCasterSceneAudioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateCasterSceneAudioResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateCasterSceneAudio"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCasterSceneAudio(request *UpdateCasterSceneAudioRequest) (_result *UpdateCasterSceneAudioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCasterSceneAudioResponse{}
	_body, _err := client.UpdateCasterSceneAudioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCasterChannelWithOptions(request *SetCasterChannelRequest, runtime *util.RuntimeOptions) (_result *SetCasterChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetCasterChannelResponse{}
	_body, _err := client.DoRequest(tea.String("SetCasterChannel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCasterChannel(request *SetCasterChannelRequest) (_result *SetCasterChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCasterChannelResponse{}
	_body, _err := client.SetCasterChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterSceneAudioWithOptions(request *DescribeCasterSceneAudioRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterSceneAudioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterSceneAudioResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterSceneAudio"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterSceneAudio(request *DescribeCasterSceneAudioRequest) (_result *DescribeCasterSceneAudioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterSceneAudioResponse{}
	_body, _err := client.DescribeCasterSceneAudioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterChannelsWithOptions(request *DescribeCasterChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterChannelsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterChannels"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterChannels(request *DescribeCasterChannelsRequest) (_result *DescribeCasterChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterChannelsResponse{}
	_body, _err := client.DescribeCasterChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBoardWithOptions(request *UpdateBoardRequest, runtime *util.RuntimeOptions) (_result *UpdateBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateBoardResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateBoard"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBoard(request *UpdateBoardRequest) (_result *UpdateBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBoardResponse{}
	_body, _err := client.UpdateBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinBoardWithOptions(request *JoinBoardRequest, runtime *util.RuntimeOptions) (_result *JoinBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &JoinBoardResponse{}
	_body, _err := client.DoRequest(tea.String("JoinBoard"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinBoard(request *JoinBoardRequest) (_result *JoinBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinBoardResponse{}
	_body, _err := client.JoinBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBoardSnapshotWithOptions(request *DescribeBoardSnapshotRequest, runtime *util.RuntimeOptions) (_result *DescribeBoardSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBoardSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBoardSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBoardSnapshot(request *DescribeBoardSnapshotRequest) (_result *DescribeBoardSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBoardSnapshotResponse{}
	_body, _err := client.DescribeBoardSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBoardsWithOptions(request *DescribeBoardsRequest, runtime *util.RuntimeOptions) (_result *DescribeBoardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBoardsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBoards"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBoards(request *DescribeBoardsRequest) (_result *DescribeBoardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBoardsResponse{}
	_body, _err := client.DescribeBoardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBoardEventsWithOptions(request *DescribeBoardEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeBoardEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBoardEventsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBoardEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBoardEvents(request *DescribeBoardEventsRequest) (_result *DescribeBoardEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBoardEventsResponse{}
	_body, _err := client.DescribeBoardEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBoardWithOptions(request *DeleteBoardRequest, runtime *util.RuntimeOptions) (_result *DeleteBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteBoardResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteBoard"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBoard(request *DeleteBoardRequest) (_result *DeleteBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBoardResponse{}
	_body, _err := client.DeleteBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBoardWithOptions(request *CreateBoardRequest, runtime *util.RuntimeOptions) (_result *CreateBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBoardResponse{}
	_body, _err := client.DoRequest(tea.String("CreateBoard"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBoard(request *CreateBoardRequest) (_result *CreateBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBoardResponse{}
	_body, _err := client.CreateBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteBoardWithOptions(request *CompleteBoardRequest, runtime *util.RuntimeOptions) (_result *CompleteBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CompleteBoardResponse{}
	_body, _err := client.DoRequest(tea.String("CompleteBoard"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteBoard(request *CompleteBoardRequest) (_result *CompleteBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteBoardResponse{}
	_body, _err := client.CompleteBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyBoardTokenWithOptions(request *ApplyBoardTokenRequest, runtime *util.RuntimeOptions) (_result *ApplyBoardTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyBoardTokenResponse{}
	_body, _err := client.DoRequest(tea.String("ApplyBoardToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyBoardToken(request *ApplyBoardTokenRequest) (_result *ApplyBoardTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyBoardTokenResponse{}
	_body, _err := client.ApplyBoardTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamCountWithOptions(request *DescribeLiveStreamCountRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamCountResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamCount"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamCount(request *DescribeLiveStreamCountRequest) (_result *DescribeLiveStreamCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamCountResponse{}
	_body, _err := client.DescribeLiveStreamCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveCertificateDetailWithOptions(request *DescribeLiveCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveCertificateDetailResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveCertificateDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveCertificateDetail(request *DescribeLiveCertificateDetailRequest) (_result *DescribeLiveCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveCertificateDetailResponse{}
	_body, _err := client.DescribeLiveCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHlsLiveStreamRealTimeBpsDataWithOptions(request *DescribeHlsLiveStreamRealTimeBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeHlsLiveStreamRealTimeBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeHlsLiveStreamRealTimeBpsDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeHlsLiveStreamRealTimeBpsData"), tea.String("HTTPS"), tea.String("GET"), tea.String("2016-11-01"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHlsLiveStreamRealTimeBpsData(request *DescribeHlsLiveStreamRealTimeBpsDataRequest) (_result *DescribeHlsLiveStreamRealTimeBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHlsLiveStreamRealTimeBpsDataResponse{}
	_body, _err := client.DescribeHlsLiveStreamRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLiveDomainWithOptions(request *StopLiveDomainRequest, runtime *util.RuntimeOptions) (_result *StopLiveDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopLiveDomainResponse{}
	_body, _err := client.DoRequest(tea.String("StopLiveDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLiveDomain(request *StopLiveDomainRequest) (_result *StopLiveDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveDomainResponse{}
	_body, _err := client.StopLiveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartLiveDomainWithOptions(request *StartLiveDomainRequest, runtime *util.RuntimeOptions) (_result *StartLiveDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartLiveDomainResponse{}
	_body, _err := client.DoRequest(tea.String("StartLiveDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartLiveDomain(request *StartLiveDomainRequest) (_result *StartLiveDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartLiveDomainResponse{}
	_body, _err := client.StartLiveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLiveDomainCertificateWithOptions(request *SetLiveDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *SetLiveDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetLiveDomainCertificateResponse{}
	_body, _err := client.DoRequest(tea.String("SetLiveDomainCertificate"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLiveDomainCertificate(request *SetLiveDomainCertificateRequest) (_result *SetLiveDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLiveDomainCertificateResponse{}
	_body, _err := client.SetLiveDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetLiveDomainConfigsWithOptions(request *BatchSetLiveDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchSetLiveDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BatchSetLiveDomainConfigsResponse{}
	_body, _err := client.DoRequest(tea.String("BatchSetLiveDomainConfigs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetLiveDomainConfigs(request *BatchSetLiveDomainConfigsRequest) (_result *BatchSetLiveDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetLiveDomainConfigsResponse{}
	_body, _err := client.BatchSetLiveDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveCertificateListWithOptions(request *DescribeLiveCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveCertificateListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveCertificateList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveCertificateList(request *DescribeLiveCertificateListRequest) (_result *DescribeLiveCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveCertificateListResponse{}
	_body, _err := client.DescribeLiveCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveDomainWithOptions(request *DeleteLiveDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveDomainResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveDomain(request *DeleteLiveDomainRequest) (_result *DeleteLiveDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveDomainResponse{}
	_body, _err := client.DeleteLiveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainConfigsWithOptions(request *DescribeLiveDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainConfigsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainConfigs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainConfigs(request *DescribeLiveDomainConfigsRequest) (_result *DescribeLiveDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainConfigsResponse{}
	_body, _err := client.DescribeLiveDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveDomainWithOptions(request *AddLiveDomainRequest, runtime *util.RuntimeOptions) (_result *AddLiveDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveDomainResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveDomain(request *AddLiveDomainRequest) (_result *AddLiveDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveDomainResponse{}
	_body, _err := client.AddLiveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainDetailWithOptions(request *DescribeLiveDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainDetailResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainDetail(request *DescribeLiveDomainDetailRequest) (_result *DescribeLiveDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainDetailResponse{}
	_body, _err := client.DescribeLiveDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteLiveDomainConfigsWithOptions(request *BatchDeleteLiveDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteLiveDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BatchDeleteLiveDomainConfigsResponse{}
	_body, _err := client.DoRequest(tea.String("BatchDeleteLiveDomainConfigs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteLiveDomainConfigs(request *BatchDeleteLiveDomainConfigsRequest) (_result *BatchDeleteLiveDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteLiveDomainConfigsResponse{}
	_body, _err := client.BatchDeleteLiveDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoomStatusWithOptions(request *DescribeRoomStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeRoomStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRoomStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRoomStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoomStatus(request *DescribeRoomStatusRequest) (_result *DescribeRoomStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoomStatusResponse{}
	_body, _err := client.DescribeRoomStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoomListWithOptions(request *DescribeRoomListRequest, runtime *util.RuntimeOptions) (_result *DescribeRoomListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRoomListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRoomList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoomList(request *DescribeRoomListRequest) (_result *DescribeRoomListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoomListResponse{}
	_body, _err := client.DescribeRoomListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoomKickoutUserListWithOptions(request *DescribeRoomKickoutUserListRequest, runtime *util.RuntimeOptions) (_result *DescribeRoomKickoutUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRoomKickoutUserListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRoomKickoutUserList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoomKickoutUserList(request *DescribeRoomKickoutUserListRequest) (_result *DescribeRoomKickoutUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoomKickoutUserListResponse{}
	_body, _err := client.DescribeRoomKickoutUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendRoomUserNotificationWithOptions(request *SendRoomUserNotificationRequest, runtime *util.RuntimeOptions) (_result *SendRoomUserNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SendRoomUserNotificationResponse{}
	_body, _err := client.DoRequest(tea.String("SendRoomUserNotification"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendRoomUserNotification(request *SendRoomUserNotificationRequest) (_result *SendRoomUserNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendRoomUserNotificationResponse{}
	_body, _err := client.SendRoomUserNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeForbidPushStreamRoomListWithOptions(request *DescribeForbidPushStreamRoomListRequest, runtime *util.RuntimeOptions) (_result *DescribeForbidPushStreamRoomListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeForbidPushStreamRoomListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeForbidPushStreamRoomList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeForbidPushStreamRoomList(request *DescribeForbidPushStreamRoomListRequest) (_result *DescribeForbidPushStreamRoomListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeForbidPushStreamRoomListResponse{}
	_body, _err := client.DescribeForbidPushStreamRoomListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendRoomNotificationWithOptions(request *SendRoomNotificationRequest, runtime *util.RuntimeOptions) (_result *SendRoomNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SendRoomNotificationResponse{}
	_body, _err := client.DoRequest(tea.String("SendRoomNotification"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendRoomNotification(request *SendRoomNotificationRequest) (_result *SendRoomNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendRoomNotificationResponse{}
	_body, _err := client.SendRoomNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ForbidPushStreamWithOptions(request *ForbidPushStreamRequest, runtime *util.RuntimeOptions) (_result *ForbidPushStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ForbidPushStreamResponse{}
	_body, _err := client.DoRequest(tea.String("ForbidPushStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ForbidPushStream(request *ForbidPushStreamRequest) (_result *ForbidPushStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForbidPushStreamResponse{}
	_body, _err := client.ForbidPushStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoomWithOptions(request *DeleteRoomRequest, runtime *util.RuntimeOptions) (_result *DeleteRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteRoomResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteRoom"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRoom(request *DeleteRoomRequest) (_result *DeleteRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoomResponse{}
	_body, _err := client.DeleteRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoomWithOptions(request *CreateRoomRequest, runtime *util.RuntimeOptions) (_result *CreateRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateRoomResponse{}
	_body, _err := client.DoRequest(tea.String("CreateRoom"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRoom(request *CreateRoomRequest) (_result *CreateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoomResponse{}
	_body, _err := client.CreateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllowPushStreamWithOptions(request *AllowPushStreamRequest, runtime *util.RuntimeOptions) (_result *AllowPushStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AllowPushStreamResponse{}
	_body, _err := client.DoRequest(tea.String("AllowPushStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllowPushStream(request *AllowPushStreamRequest) (_result *AllowPushStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllowPushStreamResponse{}
	_body, _err := client.AllowPushStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveUserDomainsWithOptions(request *DescribeLiveUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveUserDomainsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveUserDomains"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveUserDomains(request *DescribeLiveUserDomainsRequest) (_result *DescribeLiveUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveUserDomainsResponse{}
	_body, _err := client.DescribeLiveUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterRtcInfoWithOptions(request *DescribeCasterRtcInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterRtcInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterRtcInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterRtcInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterRtcInfo(request *DescribeCasterRtcInfoRequest) (_result *DescribeCasterRtcInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterRtcInfoResponse{}
	_body, _err := client.DescribeCasterRtcInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpBpsPeakDataWithOptions(request *DescribeUpBpsPeakDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUpBpsPeakDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUpBpsPeakDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUpBpsPeakData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpBpsPeakData(request *DescribeUpBpsPeakDataRequest) (_result *DescribeUpBpsPeakDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpBpsPeakDataResponse{}
	_body, _err := client.DescribeUpBpsPeakDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpBpsPeakOfLineWithOptions(request *DescribeUpBpsPeakOfLineRequest, runtime *util.RuntimeOptions) (_result *DescribeUpBpsPeakOfLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUpBpsPeakOfLineResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUpBpsPeakOfLine"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpBpsPeakOfLine(request *DescribeUpBpsPeakOfLineRequest) (_result *DescribeUpBpsPeakOfLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpBpsPeakOfLineResponse{}
	_body, _err := client.DescribeUpBpsPeakOfLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpPeakPublishStreamDataWithOptions(request *DescribeUpPeakPublishStreamDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUpPeakPublishStreamDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUpPeakPublishStreamDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUpPeakPublishStreamData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpPeakPublishStreamData(request *DescribeUpPeakPublishStreamDataRequest) (_result *DescribeUpPeakPublishStreamDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpPeakPublishStreamDataResponse{}
	_body, _err := client.DescribeUpPeakPublishStreamDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveDomainMappingWithOptions(request *DeleteLiveDomainMappingRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveDomainMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveDomainMappingResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveDomainMapping"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveDomainMapping(request *DeleteLiveDomainMappingRequest) (_result *DeleteLiveDomainMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveDomainMappingResponse{}
	_body, _err := client.DeleteLiveDomainMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveDomainMappingWithOptions(request *AddLiveDomainMappingRequest, runtime *util.RuntimeOptions) (_result *AddLiveDomainMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveDomainMappingResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveDomainMapping"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveDomainMapping(request *AddLiveDomainMappingRequest) (_result *AddLiveDomainMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveDomainMappingResponse{}
	_body, _err := client.AddLiveDomainMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterEpisodeGroupContentWithOptions(request *AddCasterEpisodeGroupContentRequest, runtime *util.RuntimeOptions) (_result *AddCasterEpisodeGroupContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterEpisodeGroupContentResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterEpisodeGroupContent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterEpisodeGroupContent(request *AddCasterEpisodeGroupContentRequest) (_result *AddCasterEpisodeGroupContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterEpisodeGroupContentResponse{}
	_body, _err := client.AddCasterEpisodeGroupContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCasterProgramWithOptions(request *ModifyCasterProgramRequest, runtime *util.RuntimeOptions) (_result *ModifyCasterProgramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyCasterProgramResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyCasterProgram"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCasterProgram(request *ModifyCasterProgramRequest) (_result *ModifyCasterProgramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCasterProgramResponse{}
	_body, _err := client.ModifyCasterProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCasterEpisodeWithOptions(request *ModifyCasterEpisodeRequest, runtime *util.RuntimeOptions) (_result *ModifyCasterEpisodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyCasterEpisodeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyCasterEpisode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCasterEpisode(request *ModifyCasterEpisodeRequest) (_result *ModifyCasterEpisodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCasterEpisodeResponse{}
	_body, _err := client.ModifyCasterEpisodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterProgramWithOptions(request *DescribeCasterProgramRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterProgramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterProgramResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterProgram"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterProgram(request *DescribeCasterProgramRequest) (_result *DescribeCasterProgramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterProgramResponse{}
	_body, _err := client.DescribeCasterProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterProgramWithOptions(request *DeleteCasterProgramRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterProgramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterProgramResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterProgram"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterProgram(request *DeleteCasterProgramRequest) (_result *DeleteCasterProgramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterProgramResponse{}
	_body, _err := client.DeleteCasterProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterEpisodeGroupWithOptions(request *DeleteCasterEpisodeGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterEpisodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterEpisodeGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterEpisodeGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterEpisodeGroup(request *DeleteCasterEpisodeGroupRequest) (_result *DeleteCasterEpisodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterEpisodeGroupResponse{}
	_body, _err := client.DeleteCasterEpisodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterEpisodeWithOptions(request *DeleteCasterEpisodeRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterEpisodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterEpisodeResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterEpisode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterEpisode(request *DeleteCasterEpisodeRequest) (_result *DeleteCasterEpisodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterEpisodeResponse{}
	_body, _err := client.DeleteCasterEpisodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterProgramWithOptions(request *AddCasterProgramRequest, runtime *util.RuntimeOptions) (_result *AddCasterProgramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterProgramResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterProgram"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterProgram(request *AddCasterProgramRequest) (_result *AddCasterProgramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterProgramResponse{}
	_body, _err := client.AddCasterProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterEpisodeGroupWithOptions(request *AddCasterEpisodeGroupRequest, runtime *util.RuntimeOptions) (_result *AddCasterEpisodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterEpisodeGroupResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterEpisodeGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterEpisodeGroup(request *AddCasterEpisodeGroupRequest) (_result *AddCasterEpisodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterEpisodeGroupResponse{}
	_body, _err := client.AddCasterEpisodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterEpisodeWithOptions(request *AddCasterEpisodeRequest, runtime *util.RuntimeOptions) (_result *AddCasterEpisodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterEpisodeResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterEpisode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterEpisode(request *AddCasterEpisodeRequest) (_result *AddCasterEpisodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterEpisodeResponse{}
	_body, _err := client.AddCasterEpisodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainTranscodeDataWithOptions(request *DescribeLiveDomainTranscodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainTranscodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainTranscodeDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainTranscodeData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainTranscodeData(request *DescribeLiveDomainTranscodeDataRequest) (_result *DescribeLiveDomainTranscodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainTranscodeDataResponse{}
	_body, _err := client.DescribeLiveDomainTranscodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainSnapshotDataWithOptions(request *DescribeLiveDomainSnapshotDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainSnapshotDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainSnapshotDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainSnapshotData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainSnapshotData(request *DescribeLiveDomainSnapshotDataRequest) (_result *DescribeLiveDomainSnapshotDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainSnapshotDataResponse{}
	_body, _err := client.DescribeLiveDomainSnapshotDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainRecordDataWithOptions(request *DescribeLiveDomainRecordDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainRecordDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainRecordDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainRecordData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainRecordData(request *DescribeLiveDomainRecordDataRequest) (_result *DescribeLiveDomainRecordDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainRecordDataResponse{}
	_body, _err := client.DescribeLiveDomainRecordDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RealTimeRecordCommandWithOptions(request *RealTimeRecordCommandRequest, runtime *util.RuntimeOptions) (_result *RealTimeRecordCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RealTimeRecordCommandResponse{}
	_body, _err := client.DoRequest(tea.String("RealTimeRecordCommand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RealTimeRecordCommand(request *RealTimeRecordCommandRequest) (_result *RealTimeRecordCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RealTimeRecordCommandResponse{}
	_body, _err := client.RealTimeRecordCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainTrafficDataWithOptions(request *DescribeLiveDomainTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainTrafficDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainTrafficData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainTrafficData(request *DescribeLiveDomainTrafficDataRequest) (_result *DescribeLiveDomainTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainTrafficDataResponse{}
	_body, _err := client.DescribeLiveDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDomainBpsDataWithOptions(request *DescribeLiveDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDomainBpsDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDomainBpsData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDomainBpsData(request *DescribeLiveDomainBpsDataRequest) (_result *DescribeLiveDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDomainBpsDataResponse{}
	_body, _err := client.DescribeLiveDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTrancodeSEIWithOptions(request *AddTrancodeSEIRequest, runtime *util.RuntimeOptions) (_result *AddTrancodeSEIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddTrancodeSEIResponse{}
	_body, _err := client.DoRequest(tea.String("AddTrancodeSEI"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTrancodeSEI(request *AddTrancodeSEIRequest) (_result *AddTrancodeSEIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTrancodeSEIResponse{}
	_body, _err := client.AddTrancodeSEIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterSceneConfigWithOptions(request *DeleteCasterSceneConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterSceneConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterSceneConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterSceneConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterSceneConfig(request *DeleteCasterSceneConfigRequest) (_result *DeleteCasterSceneConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterSceneConfigResponse{}
	_body, _err := client.DeleteCasterSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCustomLiveStreamTranscodeWithOptions(request *AddCustomLiveStreamTranscodeRequest, runtime *util.RuntimeOptions) (_result *AddCustomLiveStreamTranscodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCustomLiveStreamTranscodeResponse{}
	_body, _err := client.DoRequest(tea.String("AddCustomLiveStreamTranscode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCustomLiveStreamTranscode(request *AddCustomLiveStreamTranscodeRequest) (_result *AddCustomLiveStreamTranscodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCustomLiveStreamTranscodeResponse{}
	_body, _err := client.AddCustomLiveStreamTranscodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveRecordVodConfigsWithOptions(request *DescribeLiveRecordVodConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveRecordVodConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveRecordVodConfigsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveRecordVodConfigs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveRecordVodConfigs(request *DescribeLiveRecordVodConfigsRequest) (_result *DescribeLiveRecordVodConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveRecordVodConfigsResponse{}
	_body, _err := client.DescribeLiveRecordVodConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveRecordVodConfigWithOptions(request *DeleteLiveRecordVodConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveRecordVodConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveRecordVodConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveRecordVodConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveRecordVodConfig(request *DeleteLiveRecordVodConfigRequest) (_result *DeleteLiveRecordVodConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveRecordVodConfigResponse{}
	_body, _err := client.DeleteLiveRecordVodConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveRecordVodConfigWithOptions(request *AddLiveRecordVodConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveRecordVodConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveRecordVodConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveRecordVodConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveRecordVodConfig(request *AddLiveRecordVodConfigRequest) (_result *AddLiveRecordVodConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveRecordVodConfigResponse{}
	_body, _err := client.AddLiveRecordVodConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCasterComponentWithOptions(request *ModifyCasterComponentRequest, runtime *util.RuntimeOptions) (_result *ModifyCasterComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyCasterComponentResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyCasterComponent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCasterComponent(request *ModifyCasterComponentRequest) (_result *ModifyCasterComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCasterComponentResponse{}
	_body, _err := client.ModifyCasterComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterComponentsWithOptions(request *DescribeCasterComponentsRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterComponentsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterComponents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterComponents(request *DescribeCasterComponentsRequest) (_result *DescribeCasterComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterComponentsResponse{}
	_body, _err := client.DescribeCasterComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterComponentWithOptions(request *DeleteCasterComponentRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterComponentResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterComponent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterComponent(request *DeleteCasterComponentRequest) (_result *DeleteCasterComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterComponentResponse{}
	_body, _err := client.DeleteCasterComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterComponentWithOptions(request *AddCasterComponentRequest, runtime *util.RuntimeOptions) (_result *AddCasterComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterComponentResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterComponent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterComponent(request *AddCasterComponentRequest) (_result *AddCasterComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterComponentResponse{}
	_body, _err := client.AddCasterComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCasterWithOptions(request *StopCasterRequest, runtime *util.RuntimeOptions) (_result *StopCasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopCasterResponse{}
	_body, _err := client.DoRequest(tea.String("StopCaster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCaster(request *StopCasterRequest) (_result *StopCasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCasterResponse{}
	_body, _err := client.StopCasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCasterWithOptions(request *StartCasterRequest, runtime *util.RuntimeOptions) (_result *StartCasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartCasterResponse{}
	_body, _err := client.DoRequest(tea.String("StartCaster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCaster(request *StartCasterRequest) (_result *StartCasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCasterResponse{}
	_body, _err := client.StartCasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamHistoryUserNumWithOptions(request *DescribeLiveStreamHistoryUserNumRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamHistoryUserNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamHistoryUserNumResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamHistoryUserNum"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamHistoryUserNum(request *DescribeLiveStreamHistoryUserNumRequest) (_result *DescribeLiveStreamHistoryUserNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamHistoryUserNumResponse{}
	_body, _err := client.DescribeLiveStreamHistoryUserNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCasterSceneConfigWithOptions(request *UpdateCasterSceneConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateCasterSceneConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateCasterSceneConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateCasterSceneConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCasterSceneConfig(request *UpdateCasterSceneConfigRequest) (_result *UpdateCasterSceneConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCasterSceneConfigResponse{}
	_body, _err := client.UpdateCasterSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCasterSceneWithOptions(request *StopCasterSceneRequest, runtime *util.RuntimeOptions) (_result *StopCasterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopCasterSceneResponse{}
	_body, _err := client.DoRequest(tea.String("StopCasterScene"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCasterScene(request *StopCasterSceneRequest) (_result *StopCasterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCasterSceneResponse{}
	_body, _err := client.StopCasterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCasterSceneWithOptions(request *StartCasterSceneRequest, runtime *util.RuntimeOptions) (_result *StartCasterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartCasterSceneResponse{}
	_body, _err := client.DoRequest(tea.String("StartCasterScene"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCasterScene(request *StartCasterSceneRequest) (_result *StartCasterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCasterSceneResponse{}
	_body, _err := client.StartCasterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCasterSceneConfigWithOptions(request *SetCasterSceneConfigRequest, runtime *util.RuntimeOptions) (_result *SetCasterSceneConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetCasterSceneConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SetCasterSceneConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCasterSceneConfig(request *SetCasterSceneConfigRequest) (_result *SetCasterSceneConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCasterSceneConfigResponse{}
	_body, _err := client.SetCasterSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCasterConfigWithOptions(request *SetCasterConfigRequest, runtime *util.RuntimeOptions) (_result *SetCasterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetCasterConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SetCasterConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCasterConfig(request *SetCasterConfigRequest) (_result *SetCasterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCasterConfigResponse{}
	_body, _err := client.SetCasterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCasterVideoResourceWithOptions(request *ModifyCasterVideoResourceRequest, runtime *util.RuntimeOptions) (_result *ModifyCasterVideoResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyCasterVideoResourceResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyCasterVideoResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCasterVideoResource(request *ModifyCasterVideoResourceRequest) (_result *ModifyCasterVideoResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCasterVideoResourceResponse{}
	_body, _err := client.ModifyCasterVideoResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCasterLayoutWithOptions(request *ModifyCasterLayoutRequest, runtime *util.RuntimeOptions) (_result *ModifyCasterLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyCasterLayoutResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyCasterLayout"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCasterLayout(request *ModifyCasterLayoutRequest) (_result *ModifyCasterLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCasterLayoutResponse{}
	_body, _err := client.ModifyCasterLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EffectCasterVideoResourceWithOptions(request *EffectCasterVideoResourceRequest, runtime *util.RuntimeOptions) (_result *EffectCasterVideoResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EffectCasterVideoResourceResponse{}
	_body, _err := client.DoRequest(tea.String("EffectCasterVideoResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EffectCasterVideoResource(request *EffectCasterVideoResourceRequest) (_result *EffectCasterVideoResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EffectCasterVideoResourceResponse{}
	_body, _err := client.EffectCasterVideoResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EffectCasterUrgentWithOptions(request *EffectCasterUrgentRequest, runtime *util.RuntimeOptions) (_result *EffectCasterUrgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EffectCasterUrgentResponse{}
	_body, _err := client.DoRequest(tea.String("EffectCasterUrgent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EffectCasterUrgent(request *EffectCasterUrgentRequest) (_result *EffectCasterUrgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EffectCasterUrgentResponse{}
	_body, _err := client.EffectCasterUrgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterVideoResourcesWithOptions(request *DescribeCasterVideoResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterVideoResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterVideoResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterVideoResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterVideoResources(request *DescribeCasterVideoResourcesRequest) (_result *DescribeCasterVideoResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterVideoResourcesResponse{}
	_body, _err := client.DescribeCasterVideoResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterStreamUrlWithOptions(request *DescribeCasterStreamUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterStreamUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterStreamUrlResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterStreamUrl"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterStreamUrl(request *DescribeCasterStreamUrlRequest) (_result *DescribeCasterStreamUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterStreamUrlResponse{}
	_body, _err := client.DescribeCasterStreamUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterScenesWithOptions(request *DescribeCasterScenesRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterScenesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterScenes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterScenes(request *DescribeCasterScenesRequest) (_result *DescribeCasterScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterScenesResponse{}
	_body, _err := client.DescribeCasterScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCastersWithOptions(request *DescribeCastersRequest, runtime *util.RuntimeOptions) (_result *DescribeCastersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCastersResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasters(request *DescribeCastersRequest) (_result *DescribeCastersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCastersResponse{}
	_body, _err := client.DescribeCastersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterLayoutsWithOptions(request *DescribeCasterLayoutsRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterLayoutsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterLayoutsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterLayouts"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterLayouts(request *DescribeCasterLayoutsRequest) (_result *DescribeCasterLayoutsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterLayoutsResponse{}
	_body, _err := client.DescribeCasterLayoutsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCasterConfigWithOptions(request *DescribeCasterConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeCasterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCasterConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCasterConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCasterConfig(request *DescribeCasterConfigRequest) (_result *DescribeCasterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCasterConfigResponse{}
	_body, _err := client.DescribeCasterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterVideoResourceWithOptions(request *DeleteCasterVideoResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterVideoResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterVideoResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterVideoResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterVideoResource(request *DeleteCasterVideoResourceRequest) (_result *DeleteCasterVideoResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterVideoResourceResponse{}
	_body, _err := client.DeleteCasterVideoResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterLayoutWithOptions(request *DeleteCasterLayoutRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterLayoutResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCasterLayout"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCasterLayout(request *DeleteCasterLayoutRequest) (_result *DeleteCasterLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterLayoutResponse{}
	_body, _err := client.DeleteCasterLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCasterWithOptions(request *DeleteCasterRequest, runtime *util.RuntimeOptions) (_result *DeleteCasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteCasterResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCaster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCaster(request *DeleteCasterRequest) (_result *DeleteCasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCasterResponse{}
	_body, _err := client.DeleteCasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCasterWithOptions(request *CreateCasterRequest, runtime *util.RuntimeOptions) (_result *CreateCasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCasterResponse{}
	_body, _err := client.DoRequest(tea.String("CreateCaster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCaster(request *CreateCasterRequest) (_result *CreateCasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCasterResponse{}
	_body, _err := client.CreateCasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyCasterSceneConfigWithOptions(request *CopyCasterSceneConfigRequest, runtime *util.RuntimeOptions) (_result *CopyCasterSceneConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CopyCasterSceneConfigResponse{}
	_body, _err := client.DoRequest(tea.String("CopyCasterSceneConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyCasterSceneConfig(request *CopyCasterSceneConfigRequest) (_result *CopyCasterSceneConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyCasterSceneConfigResponse{}
	_body, _err := client.CopyCasterSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyCasterWithOptions(request *CopyCasterRequest, runtime *util.RuntimeOptions) (_result *CopyCasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CopyCasterResponse{}
	_body, _err := client.DoRequest(tea.String("CopyCaster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyCaster(request *CopyCasterRequest) (_result *CopyCasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyCasterResponse{}
	_body, _err := client.CopyCasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterVideoResourceWithOptions(request *AddCasterVideoResourceRequest, runtime *util.RuntimeOptions) (_result *AddCasterVideoResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterVideoResourceResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterVideoResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterVideoResource(request *AddCasterVideoResourceRequest) (_result *AddCasterVideoResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterVideoResourceResponse{}
	_body, _err := client.AddCasterVideoResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCasterLayoutWithOptions(request *AddCasterLayoutRequest, runtime *util.RuntimeOptions) (_result *AddCasterLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCasterLayoutResponse{}
	_body, _err := client.DoRequest(tea.String("AddCasterLayout"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCasterLayout(request *AddCasterLayoutRequest) (_result *AddCasterLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCasterLayoutResponse{}
	_body, _err := client.AddCasterLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLivePullStreamConfigWithOptions(request *DescribeLivePullStreamConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLivePullStreamConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLivePullStreamConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLivePullStreamConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLivePullStreamConfig(request *DescribeLivePullStreamConfigRequest) (_result *DescribeLivePullStreamConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLivePullStreamConfigResponse{}
	_body, _err := client.DescribeLivePullStreamConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLivePullStreamInfoConfigWithOptions(request *DeleteLivePullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLivePullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLivePullStreamInfoConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLivePullStreamInfoConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLivePullStreamInfoConfig(request *DeleteLivePullStreamInfoConfigRequest) (_result *DeleteLivePullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLivePullStreamInfoConfigResponse{}
	_body, _err := client.DeleteLivePullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLivePullStreamInfoConfigWithOptions(request *AddLivePullStreamInfoConfigRequest, runtime *util.RuntimeOptions) (_result *AddLivePullStreamInfoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLivePullStreamInfoConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLivePullStreamInfoConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLivePullStreamInfoConfig(request *AddLivePullStreamInfoConfigRequest) (_result *AddLivePullStreamInfoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLivePullStreamInfoConfigResponse{}
	_body, _err := client.AddLivePullStreamInfoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamBitRateDataWithOptions(request *DescribeLiveStreamBitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamBitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamBitRateDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamBitRateData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamBitRateData(request *DescribeLiveStreamBitRateDataRequest) (_result *DescribeLiveStreamBitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamBitRateDataResponse{}
	_body, _err := client.DescribeLiveStreamBitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveDetectNotifyConfigWithOptions(request *AddLiveDetectNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveDetectNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveDetectNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveDetectNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveDetectNotifyConfig(request *AddLiveDetectNotifyConfigRequest) (_result *AddLiveDetectNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveDetectNotifyConfigResponse{}
	_body, _err := client.AddLiveDetectNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveSnapshotDetectPornConfigWithOptions(request *AddLiveSnapshotDetectPornConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveSnapshotDetectPornConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveSnapshotDetectPornConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveSnapshotDetectPornConfig(request *AddLiveSnapshotDetectPornConfigRequest) (_result *AddLiveSnapshotDetectPornConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.AddLiveSnapshotDetectPornConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveDetectNotifyConfigWithOptions(request *DeleteLiveDetectNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveDetectNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveDetectNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveDetectNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveDetectNotifyConfig(request *DeleteLiveDetectNotifyConfigRequest) (_result *DeleteLiveDetectNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveDetectNotifyConfigResponse{}
	_body, _err := client.DeleteLiveDetectNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveDetectNotifyConfigWithOptions(request *DescribeLiveDetectNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveDetectNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveDetectNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveDetectNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveDetectNotifyConfig(request *DescribeLiveDetectNotifyConfigRequest) (_result *DescribeLiveDetectNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveDetectNotifyConfigResponse{}
	_body, _err := client.DescribeLiveDetectNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveSnapshotDetectPornConfigWithOptions(request *DeleteLiveSnapshotDetectPornConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveSnapshotDetectPornConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveSnapshotDetectPornConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveSnapshotDetectPornConfig(request *DeleteLiveSnapshotDetectPornConfigRequest) (_result *DeleteLiveSnapshotDetectPornConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.DeleteLiveSnapshotDetectPornConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveSnapshotDetectPornConfigWithOptions(request *DescribeLiveSnapshotDetectPornConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveSnapshotDetectPornConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveSnapshotDetectPornConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveSnapshotDetectPornConfig(request *DescribeLiveSnapshotDetectPornConfigRequest) (_result *DescribeLiveSnapshotDetectPornConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.DescribeLiveSnapshotDetectPornConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveDetectNotifyConfigWithOptions(request *UpdateLiveDetectNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveDetectNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveDetectNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveDetectNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveDetectNotifyConfig(request *UpdateLiveDetectNotifyConfigRequest) (_result *UpdateLiveDetectNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveDetectNotifyConfigResponse{}
	_body, _err := client.UpdateLiveDetectNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveSnapshotDetectPornConfigWithOptions(request *UpdateLiveSnapshotDetectPornConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveSnapshotDetectPornConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveSnapshotDetectPornConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveSnapshotDetectPornConfig(request *UpdateLiveSnapshotDetectPornConfigRequest) (_result *UpdateLiveSnapshotDetectPornConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveSnapshotDetectPornConfigResponse{}
	_body, _err := client.UpdateLiveSnapshotDetectPornConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveRecordNotifyConfigWithOptions(request *AddLiveRecordNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveRecordNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveRecordNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveRecordNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveRecordNotifyConfig(request *AddLiveRecordNotifyConfigRequest) (_result *AddLiveRecordNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveRecordNotifyConfigResponse{}
	_body, _err := client.AddLiveRecordNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveStreamsNotifyUrlConfigWithOptions(request *DeleteLiveStreamsNotifyUrlConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveStreamsNotifyUrlConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveStreamsNotifyUrlConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveStreamsNotifyUrlConfig(request *DeleteLiveStreamsNotifyUrlConfigRequest) (_result *DeleteLiveStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DeleteLiveStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveRecordNotifyConfigWithOptions(request *DeleteLiveRecordNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveRecordNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveRecordNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveRecordNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveRecordNotifyConfig(request *DeleteLiveRecordNotifyConfigRequest) (_result *DeleteLiveRecordNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveRecordNotifyConfigResponse{}
	_body, _err := client.DeleteLiveRecordNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveRecordNotifyConfigWithOptions(request *DescribeLiveRecordNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveRecordNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveRecordNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveRecordNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveRecordNotifyConfig(request *DescribeLiveRecordNotifyConfigRequest) (_result *DescribeLiveRecordNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveRecordNotifyConfigResponse{}
	_body, _err := client.DescribeLiveRecordNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamsNotifyUrlConfigWithOptions(request *DescribeLiveStreamsNotifyUrlConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamsNotifyUrlConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamsNotifyUrlConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamsNotifyUrlConfig(request *DescribeLiveStreamsNotifyUrlConfigRequest) (_result *DescribeLiveStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DescribeLiveStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveRecordNotifyConfigWithOptions(request *UpdateLiveRecordNotifyConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveRecordNotifyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveRecordNotifyConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveRecordNotifyConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveRecordNotifyConfig(request *UpdateLiveRecordNotifyConfigRequest) (_result *UpdateLiveRecordNotifyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveRecordNotifyConfigResponse{}
	_body, _err := client.UpdateLiveRecordNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamsBlockListWithOptions(request *DescribeLiveStreamsBlockListRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamsBlockListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamsBlockListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamsBlockList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamsBlockList(request *DescribeLiveStreamsBlockListRequest) (_result *DescribeLiveStreamsBlockListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamsBlockListResponse{}
	_body, _err := client.DescribeLiveStreamsBlockListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamOnlineUserNumWithOptions(request *DescribeLiveStreamOnlineUserNumRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamOnlineUserNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamOnlineUserNumResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamOnlineUserNum"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamOnlineUserNum(request *DescribeLiveStreamOnlineUserNumRequest) (_result *DescribeLiveStreamOnlineUserNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamOnlineUserNumResponse{}
	_body, _err := client.DescribeLiveStreamOnlineUserNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamsPublishListWithOptions(request *DescribeLiveStreamsPublishListRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamsPublishListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamsPublishListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamsPublishList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamsPublishList(request *DescribeLiveStreamsPublishListRequest) (_result *DescribeLiveStreamsPublishListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamsPublishListResponse{}
	_body, _err := client.DescribeLiveStreamsPublishListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamsOnlineListWithOptions(request *DescribeLiveStreamsOnlineListRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamsOnlineListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamsOnlineListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamsOnlineList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamsOnlineList(request *DescribeLiveStreamsOnlineListRequest) (_result *DescribeLiveStreamsOnlineListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamsOnlineListResponse{}
	_body, _err := client.DescribeLiveStreamsOnlineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamsControlHistoryWithOptions(request *DescribeLiveStreamsControlHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamsControlHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamsControlHistoryResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamsControlHistory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamsControlHistory(request *DescribeLiveStreamsControlHistoryRequest) (_result *DescribeLiveStreamsControlHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamsControlHistoryResponse{}
	_body, _err := client.DescribeLiveStreamsControlHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveStreamTranscodeWithOptions(request *AddLiveStreamTranscodeRequest, runtime *util.RuntimeOptions) (_result *AddLiveStreamTranscodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveStreamTranscodeResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveStreamTranscode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveStreamTranscode(request *AddLiveStreamTranscodeRequest) (_result *AddLiveStreamTranscodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveStreamTranscodeResponse{}
	_body, _err := client.AddLiveStreamTranscodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveStreamTranscodeWithOptions(request *DeleteLiveStreamTranscodeRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveStreamTranscodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveStreamTranscodeResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveStreamTranscode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveStreamTranscode(request *DeleteLiveStreamTranscodeRequest) (_result *DeleteLiveStreamTranscodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveStreamTranscodeResponse{}
	_body, _err := client.DeleteLiveStreamTranscodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamsFrameRateAndBitRateDataWithOptions(request *DescribeLiveStreamsFrameRateAndBitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamsFrameRateAndBitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamsFrameRateAndBitRateDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamsFrameRateAndBitRateData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamsFrameRateAndBitRateData(request *DescribeLiveStreamsFrameRateAndBitRateDataRequest) (_result *DescribeLiveStreamsFrameRateAndBitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamsFrameRateAndBitRateDataResponse{}
	_body, _err := client.DescribeLiveStreamsFrameRateAndBitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ForbidLiveStreamWithOptions(request *ForbidLiveStreamRequest, runtime *util.RuntimeOptions) (_result *ForbidLiveStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ForbidLiveStreamResponse{}
	_body, _err := client.DoRequest(tea.String("ForbidLiveStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (_result *ForbidLiveStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForbidLiveStreamResponse{}
	_body, _err := client.ForbidLiveStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamTranscodeInfoWithOptions(request *DescribeLiveStreamTranscodeInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamTranscodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamTranscodeInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamTranscodeInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamTranscodeInfo(request *DescribeLiveStreamTranscodeInfoRequest) (_result *DescribeLiveStreamTranscodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamTranscodeInfoResponse{}
	_body, _err := client.DescribeLiveStreamTranscodeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLiveStreamsNotifyUrlConfigWithOptions(request *SetLiveStreamsNotifyUrlConfigRequest, runtime *util.RuntimeOptions) (_result *SetLiveStreamsNotifyUrlConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.DoRequest(tea.String("SetLiveStreamsNotifyUrlConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLiveStreamsNotifyUrlConfig(request *SetLiveStreamsNotifyUrlConfigRequest) (_result *SetLiveStreamsNotifyUrlConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLiveStreamsNotifyUrlConfigResponse{}
	_body, _err := client.SetLiveStreamsNotifyUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeLiveStreamWithOptions(request *ResumeLiveStreamRequest, runtime *util.RuntimeOptions) (_result *ResumeLiveStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResumeLiveStreamResponse{}
	_body, _err := client.DoRequest(tea.String("ResumeLiveStream"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeLiveStream(request *ResumeLiveStreamRequest) (_result *ResumeLiveStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeLiveStreamResponse{}
	_body, _err := client.ResumeLiveStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveAppSnapshotConfigWithOptions(request *AddLiveAppSnapshotConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveAppSnapshotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveAppSnapshotConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveAppSnapshotConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveAppSnapshotConfig(request *AddLiveAppSnapshotConfigRequest) (_result *AddLiveAppSnapshotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveAppSnapshotConfigResponse{}
	_body, _err := client.AddLiveAppSnapshotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLiveAppRecordConfigWithOptions(request *AddLiveAppRecordConfigRequest, runtime *util.RuntimeOptions) (_result *AddLiveAppRecordConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLiveAppRecordConfigResponse{}
	_body, _err := client.DoRequest(tea.String("AddLiveAppRecordConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLiveAppRecordConfig(request *AddLiveAppRecordConfigRequest) (_result *AddLiveAppRecordConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLiveAppRecordConfigResponse{}
	_body, _err := client.AddLiveAppRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveRecordConfigWithOptions(request *DescribeLiveRecordConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveRecordConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveRecordConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveRecordConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveRecordConfig(request *DescribeLiveRecordConfigRequest) (_result *DescribeLiveRecordConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveRecordConfigResponse{}
	_body, _err := client.DescribeLiveRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveAppSnapshotConfigWithOptions(request *DeleteLiveAppSnapshotConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveAppSnapshotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveAppSnapshotConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveAppSnapshotConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveAppSnapshotConfig(request *DeleteLiveAppSnapshotConfigRequest) (_result *DeleteLiveAppSnapshotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveAppSnapshotConfigResponse{}
	_body, _err := client.DeleteLiveAppSnapshotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveAppRecordConfigWithOptions(request *DeleteLiveAppRecordConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveAppRecordConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLiveAppRecordConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLiveAppRecordConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveAppRecordConfig(request *DeleteLiveAppRecordConfigRequest) (_result *DeleteLiveAppRecordConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveAppRecordConfigResponse{}
	_body, _err := client.DeleteLiveAppRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveStreamRecordIndexFilesWithOptions(request *CreateLiveStreamRecordIndexFilesRequest, runtime *util.RuntimeOptions) (_result *CreateLiveStreamRecordIndexFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.DoRequest(tea.String("CreateLiveStreamRecordIndexFiles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLiveStreamRecordIndexFiles(request *CreateLiveStreamRecordIndexFilesRequest) (_result *CreateLiveStreamRecordIndexFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.CreateLiveStreamRecordIndexFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamSnapshotInfoWithOptions(request *DescribeLiveStreamSnapshotInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamSnapshotInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamSnapshotInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamSnapshotInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamSnapshotInfo(request *DescribeLiveStreamSnapshotInfoRequest) (_result *DescribeLiveStreamSnapshotInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamSnapshotInfoResponse{}
	_body, _err := client.DescribeLiveStreamSnapshotInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamRecordIndexFilesWithOptions(request *DescribeLiveStreamRecordIndexFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamRecordIndexFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamRecordIndexFiles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamRecordIndexFiles(request *DescribeLiveStreamRecordIndexFilesRequest) (_result *DescribeLiveStreamRecordIndexFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamRecordIndexFilesResponse{}
	_body, _err := client.DescribeLiveStreamRecordIndexFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamRecordIndexFileWithOptions(request *DescribeLiveStreamRecordIndexFileRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamRecordIndexFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamRecordIndexFileResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamRecordIndexFile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamRecordIndexFile(request *DescribeLiveStreamRecordIndexFileRequest) (_result *DescribeLiveStreamRecordIndexFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamRecordIndexFileResponse{}
	_body, _err := client.DescribeLiveStreamRecordIndexFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveStreamRecordContentWithOptions(request *DescribeLiveStreamRecordContentRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveStreamRecordContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveStreamRecordContentResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveStreamRecordContent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveStreamRecordContent(request *DescribeLiveStreamRecordContentRequest) (_result *DescribeLiveStreamRecordContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveStreamRecordContentResponse{}
	_body, _err := client.DescribeLiveStreamRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLiveSnapshotConfigWithOptions(request *DescribeLiveSnapshotConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLiveSnapshotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLiveSnapshotConfigResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLiveSnapshotConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLiveSnapshotConfig(request *DescribeLiveSnapshotConfigRequest) (_result *DescribeLiveSnapshotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLiveSnapshotConfigResponse{}
	_body, _err := client.DescribeLiveSnapshotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveAppSnapshotConfigWithOptions(request *UpdateLiveAppSnapshotConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveAppSnapshotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLiveAppSnapshotConfigResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLiveAppSnapshotConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-11-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveAppSnapshotConfig(request *UpdateLiveAppSnapshotConfigRequest) (_result *UpdateLiveAppSnapshotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveAppSnapshotConfigResponse{}
	_body, _err := client.UpdateLiveAppSnapshotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
