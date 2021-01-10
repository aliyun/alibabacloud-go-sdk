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

type CreateFaceGroupRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateFaceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceGroupRequest) SetOwnerId(v int64) *CreateFaceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFaceGroupRequest) SetShowLog(v string) *CreateFaceGroupRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateFaceGroupRequest) SetName(v string) *CreateFaceGroupRequest {
	s.Name = &v
	return s
}

type CreateFaceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateFaceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaceGroupResponseBody) SetRequestId(v string) *CreateFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFaceGroupResponseBody) SetGroupId(v string) *CreateFaceGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateFaceGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFaceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateFaceGroupResponse) SetHeaders(v map[string]*string) *CreateFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateFaceGroupResponse) SetBody(v *CreateFaceGroupResponseBody) *CreateFaceGroupResponse {
	s.Body = v
	return s
}

type CreateStreamPredictRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog               *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StreamType            *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	StreamId              *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	PredictTemplateId     *string `json:"PredictTemplateId,omitempty" xml:"PredictTemplateId,omitempty"`
	ModelIds              *string `json:"ModelIds,omitempty" xml:"ModelIds,omitempty"`
	ProbabilityThresholds *string `json:"ProbabilityThresholds,omitempty" xml:"ProbabilityThresholds,omitempty"`
	DetectIntervals       *string `json:"DetectIntervals,omitempty" xml:"DetectIntervals,omitempty"`
	Output                *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Notify                *string `json:"Notify,omitempty" xml:"Notify,omitempty"`
	AutoStart             *string `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	FaceGroupId           *string `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	ModelUserData         *string `json:"ModelUserData,omitempty" xml:"ModelUserData,omitempty"`
}

func (s CreateStreamPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPredictRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamPredictRequest) SetOwnerId(v int64) *CreateStreamPredictRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateStreamPredictRequest) SetShowLog(v string) *CreateStreamPredictRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateStreamPredictRequest) SetClientToken(v string) *CreateStreamPredictRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStreamPredictRequest) SetStreamType(v string) *CreateStreamPredictRequest {
	s.StreamType = &v
	return s
}

func (s *CreateStreamPredictRequest) SetStreamId(v string) *CreateStreamPredictRequest {
	s.StreamId = &v
	return s
}

func (s *CreateStreamPredictRequest) SetPredictTemplateId(v string) *CreateStreamPredictRequest {
	s.PredictTemplateId = &v
	return s
}

func (s *CreateStreamPredictRequest) SetModelIds(v string) *CreateStreamPredictRequest {
	s.ModelIds = &v
	return s
}

func (s *CreateStreamPredictRequest) SetProbabilityThresholds(v string) *CreateStreamPredictRequest {
	s.ProbabilityThresholds = &v
	return s
}

func (s *CreateStreamPredictRequest) SetDetectIntervals(v string) *CreateStreamPredictRequest {
	s.DetectIntervals = &v
	return s
}

func (s *CreateStreamPredictRequest) SetOutput(v string) *CreateStreamPredictRequest {
	s.Output = &v
	return s
}

func (s *CreateStreamPredictRequest) SetNotify(v string) *CreateStreamPredictRequest {
	s.Notify = &v
	return s
}

func (s *CreateStreamPredictRequest) SetAutoStart(v string) *CreateStreamPredictRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateStreamPredictRequest) SetFaceGroupId(v string) *CreateStreamPredictRequest {
	s.FaceGroupId = &v
	return s
}

func (s *CreateStreamPredictRequest) SetModelUserData(v string) *CreateStreamPredictRequest {
	s.ModelUserData = &v
	return s
}

type CreateStreamPredictResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s CreateStreamPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPredictResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamPredictResponseBody) SetRequestId(v string) *CreateStreamPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamPredictResponseBody) SetPredictId(v string) *CreateStreamPredictResponseBody {
	s.PredictId = &v
	return s
}

type CreateStreamPredictResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStreamPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStreamPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStreamPredictResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamPredictResponse) SetHeaders(v map[string]*string) *CreateStreamPredictResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamPredictResponse) SetBody(v *CreateStreamPredictResponseBody) *CreateStreamPredictResponse {
	s.Body = v
	return s
}

type DeleteFaceGroupRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteFaceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupRequest) SetOwnerId(v int64) *DeleteFaceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFaceGroupRequest) SetShowLog(v string) *DeleteFaceGroupRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteFaceGroupRequest) SetGroupId(v string) *DeleteFaceGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteFaceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteFaceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupResponseBody) SetRequestId(v string) *DeleteFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceGroupResponseBody) SetGroupId(v string) *DeleteFaceGroupResponseBody {
	s.GroupId = &v
	return s
}

type DeleteFaceGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupResponse) SetHeaders(v map[string]*string) *DeleteFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceGroupResponse) SetBody(v *DeleteFaceGroupResponseBody) *DeleteFaceGroupResponse {
	s.Body = v
	return s
}

type DeleteStreamPredictRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s DeleteStreamPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamPredictRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamPredictRequest) SetOwnerId(v int64) *DeleteStreamPredictRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteStreamPredictRequest) SetShowLog(v string) *DeleteStreamPredictRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteStreamPredictRequest) SetPredictId(v string) *DeleteStreamPredictRequest {
	s.PredictId = &v
	return s
}

type DeleteStreamPredictResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s DeleteStreamPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamPredictResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamPredictResponseBody) SetRequestId(v string) *DeleteStreamPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamPredictResponseBody) SetPredictId(v string) *DeleteStreamPredictResponseBody {
	s.PredictId = &v
	return s
}

type DeleteStreamPredictResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStreamPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStreamPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamPredictResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamPredictResponse) SetHeaders(v map[string]*string) *DeleteStreamPredictResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamPredictResponse) SetBody(v *DeleteStreamPredictResponseBody) *DeleteStreamPredictResponse {
	s.Body = v
	return s
}

type DescribeFaceGroupsRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog       *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	CurrentPage   *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFaceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaceGroupsRequest) SetOwnerId(v int64) *DescribeFaceGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFaceGroupsRequest) SetShowLog(v string) *DescribeFaceGroupsRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeFaceGroupsRequest) SetNextPageToken(v string) *DescribeFaceGroupsRequest {
	s.NextPageToken = &v
	return s
}

func (s *DescribeFaceGroupsRequest) SetCurrentPage(v int64) *DescribeFaceGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeFaceGroupsRequest) SetPageSize(v int64) *DescribeFaceGroupsRequest {
	s.PageSize = &v
	return s
}

type DescribeFaceGroupsResponseBody struct {
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage   *int64                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	NextPageToken *string                                 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	TotalNum      *int64                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	Groups        []*DescribeFaceGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s DescribeFaceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceGroupsResponseBody) SetRequestId(v string) *DescribeFaceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceGroupsResponseBody) SetCurrentPage(v int64) *DescribeFaceGroupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeFaceGroupsResponseBody) SetPageSize(v int64) *DescribeFaceGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFaceGroupsResponseBody) SetNextPageToken(v string) *DescribeFaceGroupsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *DescribeFaceGroupsResponseBody) SetTotalNum(v int64) *DescribeFaceGroupsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeFaceGroupsResponseBody) SetGroups(v []*DescribeFaceGroupsResponseBodyGroups) *DescribeFaceGroupsResponseBody {
	s.Groups = v
	return s
}

type DescribeFaceGroupsResponseBodyGroups struct {
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
}

func (s DescribeFaceGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeFaceGroupsResponseBodyGroups) SetGroupId(v string) *DescribeFaceGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeFaceGroupsResponseBodyGroups) SetName(v string) *DescribeFaceGroupsResponseBodyGroups {
	s.Name = &v
	return s
}

func (s *DescribeFaceGroupsResponseBodyGroups) SetCreationTime(v string) *DescribeFaceGroupsResponseBodyGroups {
	s.CreationTime = &v
	return s
}

type DescribeFaceGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFaceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFaceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaceGroupsResponse) SetHeaders(v map[string]*string) *DescribeFaceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaceGroupsResponse) SetBody(v *DescribeFaceGroupsResponseBody) *DescribeFaceGroupsResponse {
	s.Body = v
	return s
}

type DescribeStreamPredictResultRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog              *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	PredictId            *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
	ModelId              *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ProbabilityThreshold *string `json:"ProbabilityThreshold,omitempty" xml:"ProbabilityThreshold,omitempty"`
	NextPageToken        *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	CurrentPage          *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeStreamPredictResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictResultRequest) SetOwnerId(v int64) *DescribeStreamPredictResultRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetShowLog(v string) *DescribeStreamPredictResultRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetPredictId(v string) *DescribeStreamPredictResultRequest {
	s.PredictId = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetModelId(v string) *DescribeStreamPredictResultRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetStartTime(v string) *DescribeStreamPredictResultRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetEndTime(v string) *DescribeStreamPredictResultRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetProbabilityThreshold(v string) *DescribeStreamPredictResultRequest {
	s.ProbabilityThreshold = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetNextPageToken(v string) *DescribeStreamPredictResultRequest {
	s.NextPageToken = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetCurrentPage(v int64) *DescribeStreamPredictResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeStreamPredictResultRequest) SetPageSize(v int64) *DescribeStreamPredictResultRequest {
	s.PageSize = &v
	return s
}

type DescribeStreamPredictResultResponseBody struct {
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum           *int64                                                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	CurrentPage        *int64                                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize           *int64                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	NextPageToken      *string                                                      `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	StreamPredictDatas []*DescribeStreamPredictResultResponseBodyStreamPredictDatas `json:"StreamPredictDatas,omitempty" xml:"StreamPredictDatas,omitempty" type:"Repeated"`
}

func (s DescribeStreamPredictResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictResultResponseBody) SetRequestId(v string) *DescribeStreamPredictResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBody) SetTotalNum(v int64) *DescribeStreamPredictResultResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBody) SetCurrentPage(v int64) *DescribeStreamPredictResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBody) SetPageSize(v int64) *DescribeStreamPredictResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBody) SetNextPageToken(v string) *DescribeStreamPredictResultResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBody) SetStreamPredictDatas(v []*DescribeStreamPredictResultResponseBodyStreamPredictDatas) *DescribeStreamPredictResultResponseBody {
	s.StreamPredictDatas = v
	return s
}

type DescribeStreamPredictResultResponseBodyStreamPredictDatas struct {
	PredictId     *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	DataUrl       *string `json:"DataUrl,omitempty" xml:"DataUrl,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	PredictTime   *string `json:"PredictTime,omitempty" xml:"PredictTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PredictResult *string `json:"PredictResult,omitempty" xml:"PredictResult,omitempty"`
}

func (s DescribeStreamPredictResultResponseBodyStreamPredictDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictResultResponseBodyStreamPredictDatas) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetPredictId(v string) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.PredictId = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetModelId(v string) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.ModelId = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetDataUrl(v string) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.DataUrl = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetTimestamp(v int64) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.Timestamp = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetPredictTime(v string) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.PredictTime = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetStatus(v string) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.Status = &v
	return s
}

func (s *DescribeStreamPredictResultResponseBodyStreamPredictDatas) SetPredictResult(v string) *DescribeStreamPredictResultResponseBodyStreamPredictDatas {
	s.PredictResult = &v
	return s
}

type DescribeStreamPredictResultResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamPredictResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamPredictResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictResultResponse) SetHeaders(v map[string]*string) *DescribeStreamPredictResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamPredictResultResponse) SetBody(v *DescribeStreamPredictResultResponseBody) *DescribeStreamPredictResultResponse {
	s.Body = v
	return s
}

type DescribeStreamPredictsRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog       *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	PredictIds    *string `json:"PredictIds,omitempty" xml:"PredictIds,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	CurrentPage   *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeStreamPredictsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictsRequest) SetOwnerId(v int64) *DescribeStreamPredictsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamPredictsRequest) SetShowLog(v string) *DescribeStreamPredictsRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeStreamPredictsRequest) SetPredictIds(v string) *DescribeStreamPredictsRequest {
	s.PredictIds = &v
	return s
}

func (s *DescribeStreamPredictsRequest) SetModelId(v string) *DescribeStreamPredictsRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeStreamPredictsRequest) SetNextPageToken(v string) *DescribeStreamPredictsRequest {
	s.NextPageToken = &v
	return s
}

func (s *DescribeStreamPredictsRequest) SetCurrentPage(v int64) *DescribeStreamPredictsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeStreamPredictsRequest) SetPageSize(v int64) *DescribeStreamPredictsRequest {
	s.PageSize = &v
	return s
}

type DescribeStreamPredictsResponseBody struct {
	TotalNum       *int64                                              `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	PageSize       *int64                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage    *int64                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	NextPageToken  *string                                             `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	StreamPredicts []*DescribeStreamPredictsResponseBodyStreamPredicts `json:"StreamPredicts,omitempty" xml:"StreamPredicts,omitempty" type:"Repeated"`
}

func (s DescribeStreamPredictsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictsResponseBody) SetTotalNum(v int64) *DescribeStreamPredictsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeStreamPredictsResponseBody) SetPageSize(v int64) *DescribeStreamPredictsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamPredictsResponseBody) SetRequestId(v string) *DescribeStreamPredictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamPredictsResponseBody) SetCurrentPage(v int64) *DescribeStreamPredictsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeStreamPredictsResponseBody) SetNextPageToken(v string) *DescribeStreamPredictsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *DescribeStreamPredictsResponseBody) SetStreamPredicts(v []*DescribeStreamPredictsResponseBodyStreamPredicts) *DescribeStreamPredictsResponseBody {
	s.StreamPredicts = v
	return s
}

type DescribeStreamPredictsResponseBodyStreamPredicts struct {
	CreationTime          *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Notify                *string `json:"Notify,omitempty" xml:"Notify,omitempty"`
	PredictId             *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
	ModelUserData         *string `json:"ModelUserData,omitempty" xml:"ModelUserData,omitempty"`
	Output                *string `json:"Output,omitempty" xml:"Output,omitempty"`
	PredictTemplateId     *string `json:"PredictTemplateId,omitempty" xml:"PredictTemplateId,omitempty"`
	StreamId              *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	AutoStart             *string `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	ProbabilityThresholds *string `json:"ProbabilityThresholds,omitempty" xml:"ProbabilityThresholds,omitempty"`
	DetectIntervals       *string `json:"DetectIntervals,omitempty" xml:"DetectIntervals,omitempty"`
	StreamType            *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	UserData              *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ModelIds              *string `json:"ModelIds,omitempty" xml:"ModelIds,omitempty"`
	FaceGroupId           *string `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
}

func (s DescribeStreamPredictsResponseBodyStreamPredicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictsResponseBodyStreamPredicts) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetCreationTime(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.CreationTime = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetStatus(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.Status = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetNotify(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.Notify = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetPredictId(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.PredictId = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetModelUserData(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.ModelUserData = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetOutput(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.Output = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetPredictTemplateId(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.PredictTemplateId = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetStreamId(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.StreamId = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetAutoStart(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.AutoStart = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetProbabilityThresholds(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.ProbabilityThresholds = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetDetectIntervals(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.DetectIntervals = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetStreamType(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.StreamType = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetUserData(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.UserData = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetModelIds(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.ModelIds = &v
	return s
}

func (s *DescribeStreamPredictsResponseBodyStreamPredicts) SetFaceGroupId(v string) *DescribeStreamPredictsResponseBodyStreamPredicts {
	s.FaceGroupId = &v
	return s
}

type DescribeStreamPredictsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStreamPredictsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStreamPredictsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStreamPredictsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamPredictsResponse) SetHeaders(v map[string]*string) *DescribeStreamPredictsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamPredictsResponse) SetBody(v *DescribeStreamPredictsResponseBody) *DescribeStreamPredictsResponse {
	s.Body = v
	return s
}

type GetAlgorithmDetailRequest struct {
	AlgorithmCode *string `json:"AlgorithmCode,omitempty" xml:"AlgorithmCode,omitempty"`
}

func (s GetAlgorithmDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailRequest) SetAlgorithmCode(v string) *GetAlgorithmDetailRequest {
	s.AlgorithmCode = &v
	return s
}

type GetAlgorithmDetailResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAlgorithmDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlgorithmDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailResponseBody) SetMessage(v string) *GetAlgorithmDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlgorithmDetailResponseBody) SetRequestId(v string) *GetAlgorithmDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDetailResponseBody) SetData(v *GetAlgorithmDetailResponseBodyData) *GetAlgorithmDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetAlgorithmDetailResponseBody) SetCode(v string) *GetAlgorithmDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlgorithmDetailResponseBody) SetSuccess(v bool) *GetAlgorithmDetailResponseBody {
	s.Success = &v
	return s
}

type GetAlgorithmDetailResponseBodyData struct {
	AlgorithmName            *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	DeployRegion             *string `json:"DeployRegion,omitempty" xml:"DeployRegion,omitempty"`
	CurrentMonthCount        *int32  `json:"CurrentMonthCount,omitempty" xml:"CurrentMonthCount,omitempty"`
	AlgorithmCode            *string `json:"AlgorithmCode,omitempty" xml:"AlgorithmCode,omitempty"`
	ApiDocUrl                *string `json:"ApiDocUrl,omitempty" xml:"ApiDocUrl,omitempty"`
	CurrentMonthSuccessCount *int32  `json:"CurrentMonthSuccessCount,omitempty" xml:"CurrentMonthSuccessCount,omitempty"`
}

func (s GetAlgorithmDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailResponseBodyData) SetAlgorithmName(v string) *GetAlgorithmDetailResponseBodyData {
	s.AlgorithmName = &v
	return s
}

func (s *GetAlgorithmDetailResponseBodyData) SetDeployRegion(v string) *GetAlgorithmDetailResponseBodyData {
	s.DeployRegion = &v
	return s
}

func (s *GetAlgorithmDetailResponseBodyData) SetCurrentMonthCount(v int32) *GetAlgorithmDetailResponseBodyData {
	s.CurrentMonthCount = &v
	return s
}

func (s *GetAlgorithmDetailResponseBodyData) SetAlgorithmCode(v string) *GetAlgorithmDetailResponseBodyData {
	s.AlgorithmCode = &v
	return s
}

func (s *GetAlgorithmDetailResponseBodyData) SetApiDocUrl(v string) *GetAlgorithmDetailResponseBodyData {
	s.ApiDocUrl = &v
	return s
}

func (s *GetAlgorithmDetailResponseBodyData) SetCurrentMonthSuccessCount(v int32) *GetAlgorithmDetailResponseBodyData {
	s.CurrentMonthSuccessCount = &v
	return s
}

type GetAlgorithmDetailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAlgorithmDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDetailResponse) SetHeaders(v map[string]*string) *GetAlgorithmDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDetailResponse) SetBody(v *GetAlgorithmDetailResponseBody) *GetAlgorithmDetailResponse {
	s.Body = v
	return s
}

type GetAlgorithmHistogramsRequest struct {
	AlgorithmCode *string `json:"AlgorithmCode,omitempty" xml:"AlgorithmCode,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	AggregateType *string `json:"AggregateType,omitempty" xml:"AggregateType,omitempty"`
}

func (s GetAlgorithmHistogramsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmHistogramsRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmHistogramsRequest) SetAlgorithmCode(v string) *GetAlgorithmHistogramsRequest {
	s.AlgorithmCode = &v
	return s
}

func (s *GetAlgorithmHistogramsRequest) SetStartDate(v string) *GetAlgorithmHistogramsRequest {
	s.StartDate = &v
	return s
}

func (s *GetAlgorithmHistogramsRequest) SetEndDate(v string) *GetAlgorithmHistogramsRequest {
	s.EndDate = &v
	return s
}

func (s *GetAlgorithmHistogramsRequest) SetAggregateType(v string) *GetAlgorithmHistogramsRequest {
	s.AggregateType = &v
	return s
}

type GetAlgorithmHistogramsResponseBody struct {
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAlgorithmHistogramsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAlgorithmHistogramsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmHistogramsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmHistogramsResponseBody) SetSuccess(v bool) *GetAlgorithmHistogramsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBody) SetCode(v string) *GetAlgorithmHistogramsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBody) SetMessage(v string) *GetAlgorithmHistogramsResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBody) SetRequestId(v string) *GetAlgorithmHistogramsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBody) SetData(v *GetAlgorithmHistogramsResponseBodyData) *GetAlgorithmHistogramsResponseBody {
	s.Data = v
	return s
}

type GetAlgorithmHistogramsResponseBodyData struct {
	SuccessCount *int32                                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	FailureCount *int32                                              `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	Histograms   []*GetAlgorithmHistogramsResponseBodyDataHistograms `json:"Histograms,omitempty" xml:"Histograms,omitempty" type:"Repeated"`
}

func (s GetAlgorithmHistogramsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmHistogramsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlgorithmHistogramsResponseBodyData) SetSuccessCount(v int32) *GetAlgorithmHistogramsResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBodyData) SetFailureCount(v int32) *GetAlgorithmHistogramsResponseBodyData {
	s.FailureCount = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBodyData) SetHistograms(v []*GetAlgorithmHistogramsResponseBodyDataHistograms) *GetAlgorithmHistogramsResponseBodyData {
	s.Histograms = v
	return s
}

type GetAlgorithmHistogramsResponseBodyDataHistograms struct {
	Time         *string `json:"Time,omitempty" xml:"Time,omitempty"`
	SuccessCount *int32  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	FailureCount *int32  `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
}

func (s GetAlgorithmHistogramsResponseBodyDataHistograms) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmHistogramsResponseBodyDataHistograms) GoString() string {
	return s.String()
}

func (s *GetAlgorithmHistogramsResponseBodyDataHistograms) SetTime(v string) *GetAlgorithmHistogramsResponseBodyDataHistograms {
	s.Time = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBodyDataHistograms) SetSuccessCount(v int32) *GetAlgorithmHistogramsResponseBodyDataHistograms {
	s.SuccessCount = &v
	return s
}

func (s *GetAlgorithmHistogramsResponseBodyDataHistograms) SetFailureCount(v int32) *GetAlgorithmHistogramsResponseBodyDataHistograms {
	s.FailureCount = &v
	return s
}

type GetAlgorithmHistogramsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAlgorithmHistogramsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmHistogramsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmHistogramsResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmHistogramsResponse) SetHeaders(v map[string]*string) *GetAlgorithmHistogramsResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmHistogramsResponse) SetBody(v *GetAlgorithmHistogramsResponseBody) *GetAlgorithmHistogramsResponse {
	s.Body = v
	return s
}

type ImagePredictRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	DataUrl *string `json:"DataUrl,omitempty" xml:"DataUrl,omitempty"`
}

func (s ImagePredictRequest) String() string {
	return tea.Prettify(s)
}

func (s ImagePredictRequest) GoString() string {
	return s.String()
}

func (s *ImagePredictRequest) SetOwnerId(v int64) *ImagePredictRequest {
	s.OwnerId = &v
	return s
}

func (s *ImagePredictRequest) SetShowLog(v string) *ImagePredictRequest {
	s.ShowLog = &v
	return s
}

func (s *ImagePredictRequest) SetModelId(v string) *ImagePredictRequest {
	s.ModelId = &v
	return s
}

func (s *ImagePredictRequest) SetDataUrl(v string) *ImagePredictRequest {
	s.DataUrl = &v
	return s
}

type ImagePredictResponseBody struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImagePredict *ImagePredictResponseBodyImagePredict `json:"ImagePredict,omitempty" xml:"ImagePredict,omitempty" type:"Struct"`
}

func (s ImagePredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImagePredictResponseBody) GoString() string {
	return s.String()
}

func (s *ImagePredictResponseBody) SetRequestId(v string) *ImagePredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImagePredictResponseBody) SetImagePredict(v *ImagePredictResponseBodyImagePredict) *ImagePredictResponseBody {
	s.ImagePredict = v
	return s
}

type ImagePredictResponseBodyImagePredict struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PredictResult *string `json:"PredictResult,omitempty" xml:"PredictResult,omitempty"`
	PredictId     *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
	PredictTime   *string `json:"PredictTime,omitempty" xml:"PredictTime,omitempty"`
	DataUrl       *string `json:"DataUrl,omitempty" xml:"DataUrl,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
}

func (s ImagePredictResponseBodyImagePredict) String() string {
	return tea.Prettify(s)
}

func (s ImagePredictResponseBodyImagePredict) GoString() string {
	return s.String()
}

func (s *ImagePredictResponseBodyImagePredict) SetStatus(v string) *ImagePredictResponseBodyImagePredict {
	s.Status = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetPredictResult(v string) *ImagePredictResponseBodyImagePredict {
	s.PredictResult = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetPredictId(v string) *ImagePredictResponseBodyImagePredict {
	s.PredictId = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetPredictTime(v string) *ImagePredictResponseBodyImagePredict {
	s.PredictTime = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetDataUrl(v string) *ImagePredictResponseBodyImagePredict {
	s.DataUrl = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetCode(v string) *ImagePredictResponseBodyImagePredict {
	s.Code = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetMessage(v string) *ImagePredictResponseBodyImagePredict {
	s.Message = &v
	return s
}

func (s *ImagePredictResponseBodyImagePredict) SetModelId(v string) *ImagePredictResponseBodyImagePredict {
	s.ModelId = &v
	return s
}

type ImagePredictResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImagePredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImagePredictResponse) String() string {
	return tea.Prettify(s)
}

func (s ImagePredictResponse) GoString() string {
	return s.String()
}

func (s *ImagePredictResponse) SetHeaders(v map[string]*string) *ImagePredictResponse {
	s.Headers = v
	return s
}

func (s *ImagePredictResponse) SetBody(v *ImagePredictResponseBody) *ImagePredictResponse {
	s.Body = v
	return s
}

type ListMyAlgorithmRequest struct {
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListMyAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMyAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ListMyAlgorithmRequest) SetAlgorithmName(v string) *ListMyAlgorithmRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListMyAlgorithmRequest) SetPageSize(v int32) *ListMyAlgorithmRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyAlgorithmRequest) SetPageNumber(v int32) *ListMyAlgorithmRequest {
	s.PageNumber = &v
	return s
}

type ListMyAlgorithmResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListMyAlgorithmResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMyAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMyAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyAlgorithmResponseBody) SetMessage(v string) *ListMyAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ListMyAlgorithmResponseBody) SetRequestId(v string) *ListMyAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyAlgorithmResponseBody) SetData(v *ListMyAlgorithmResponseBodyData) *ListMyAlgorithmResponseBody {
	s.Data = v
	return s
}

func (s *ListMyAlgorithmResponseBody) SetCode(v string) *ListMyAlgorithmResponseBody {
	s.Code = &v
	return s
}

func (s *ListMyAlgorithmResponseBody) SetSuccess(v bool) *ListMyAlgorithmResponseBody {
	s.Success = &v
	return s
}

type ListMyAlgorithmResponseBodyData struct {
	AlgorithmList []*ListMyAlgorithmResponseBodyDataAlgorithmList `json:"AlgorithmList,omitempty" xml:"AlgorithmList,omitempty" type:"Repeated"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMyAlgorithmResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMyAlgorithmResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMyAlgorithmResponseBodyData) SetAlgorithmList(v []*ListMyAlgorithmResponseBodyDataAlgorithmList) *ListMyAlgorithmResponseBodyData {
	s.AlgorithmList = v
	return s
}

func (s *ListMyAlgorithmResponseBodyData) SetPageSize(v int32) *ListMyAlgorithmResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyData) SetPageNumber(v int32) *ListMyAlgorithmResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyData) SetTotalCount(v int32) *ListMyAlgorithmResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListMyAlgorithmResponseBodyDataAlgorithmList struct {
	AlgorithmName     *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	DeployRegion      *string `json:"DeployRegion,omitempty" xml:"DeployRegion,omitempty"`
	CurrentMonthCount *int32  `json:"CurrentMonthCount,omitempty" xml:"CurrentMonthCount,omitempty"`
	AlgorithmCode     *string `json:"AlgorithmCode,omitempty" xml:"AlgorithmCode,omitempty"`
	ApiDocUrl         *string `json:"ApiDocUrl,omitempty" xml:"ApiDocUrl,omitempty"`
	YesterdayCount    *int32  `json:"YesterdayCount,omitempty" xml:"YesterdayCount,omitempty"`
	AlgorithmOrder    *int32  `json:"AlgorithmOrder,omitempty" xml:"AlgorithmOrder,omitempty"`
}

func (s ListMyAlgorithmResponseBodyDataAlgorithmList) String() string {
	return tea.Prettify(s)
}

func (s ListMyAlgorithmResponseBodyDataAlgorithmList) GoString() string {
	return s.String()
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetAlgorithmName(v string) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.AlgorithmName = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetDeployRegion(v string) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.DeployRegion = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetCurrentMonthCount(v int32) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.CurrentMonthCount = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetAlgorithmCode(v string) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.AlgorithmCode = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetApiDocUrl(v string) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.ApiDocUrl = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetYesterdayCount(v int32) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.YesterdayCount = &v
	return s
}

func (s *ListMyAlgorithmResponseBodyDataAlgorithmList) SetAlgorithmOrder(v int32) *ListMyAlgorithmResponseBodyDataAlgorithmList {
	s.AlgorithmOrder = &v
	return s
}

type ListMyAlgorithmResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMyAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMyAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMyAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ListMyAlgorithmResponse) SetHeaders(v map[string]*string) *ListMyAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ListMyAlgorithmResponse) SetBody(v *ListMyAlgorithmResponseBody) *ListMyAlgorithmResponse {
	s.Body = v
	return s
}

type PredictPictureRequest struct {
	AlgorithmCode *string `json:"AlgorithmCode,omitempty" xml:"AlgorithmCode,omitempty"`
	OssPath       *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s PredictPictureRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictPictureRequest) GoString() string {
	return s.String()
}

func (s *PredictPictureRequest) SetAlgorithmCode(v string) *PredictPictureRequest {
	s.AlgorithmCode = &v
	return s
}

func (s *PredictPictureRequest) SetOssPath(v string) *PredictPictureRequest {
	s.OssPath = &v
	return s
}

type PredictPictureResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *PredictPictureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s PredictPictureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictPictureResponseBody) GoString() string {
	return s.String()
}

func (s *PredictPictureResponseBody) SetCode(v string) *PredictPictureResponseBody {
	s.Code = &v
	return s
}

func (s *PredictPictureResponseBody) SetMessage(v string) *PredictPictureResponseBody {
	s.Message = &v
	return s
}

func (s *PredictPictureResponseBody) SetRequestId(v string) *PredictPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictPictureResponseBody) SetData(v *PredictPictureResponseBodyData) *PredictPictureResponseBody {
	s.Data = v
	return s
}

type PredictPictureResponseBodyData struct {
	PredictResult *string `json:"PredictResult,omitempty" xml:"PredictResult,omitempty"`
}

func (s PredictPictureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PredictPictureResponseBodyData) GoString() string {
	return s.String()
}

func (s *PredictPictureResponseBodyData) SetPredictResult(v string) *PredictPictureResponseBodyData {
	s.PredictResult = &v
	return s
}

type PredictPictureResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PredictPictureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredictPictureResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictPictureResponse) GoString() string {
	return s.String()
}

func (s *PredictPictureResponse) SetHeaders(v map[string]*string) *PredictPictureResponse {
	s.Headers = v
	return s
}

func (s *PredictPictureResponse) SetBody(v *PredictPictureResponseBody) *PredictPictureResponse {
	s.Body = v
	return s
}

type RegisterFaceRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog  *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RegisterFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceRequest) GoString() string {
	return s.String()
}

func (s *RegisterFaceRequest) SetOwnerId(v int64) *RegisterFaceRequest {
	s.OwnerId = &v
	return s
}

func (s *RegisterFaceRequest) SetShowLog(v string) *RegisterFaceRequest {
	s.ShowLog = &v
	return s
}

func (s *RegisterFaceRequest) SetGroupId(v string) *RegisterFaceRequest {
	s.GroupId = &v
	return s
}

func (s *RegisterFaceRequest) SetDataType(v string) *RegisterFaceRequest {
	s.DataType = &v
	return s
}

func (s *RegisterFaceRequest) SetContent(v string) *RegisterFaceRequest {
	s.Content = &v
	return s
}

type RegisterFaceResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Faces     []*RegisterFaceResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s RegisterFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterFaceResponseBody) SetRequestId(v string) *RegisterFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterFaceResponseBody) SetGroupId(v string) *RegisterFaceResponseBody {
	s.GroupId = &v
	return s
}

func (s *RegisterFaceResponseBody) SetFaces(v []*RegisterFaceResponseBodyFaces) *RegisterFaceResponseBody {
	s.Faces = v
	return s
}

type RegisterFaceResponseBodyFaces struct {
	FaceToken *string                            `json:"FaceToken,omitempty" xml:"FaceToken,omitempty"`
	Rect      *RegisterFaceResponseBodyFacesRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
}

func (s RegisterFaceResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *RegisterFaceResponseBodyFaces) SetFaceToken(v string) *RegisterFaceResponseBodyFaces {
	s.FaceToken = &v
	return s
}

func (s *RegisterFaceResponseBodyFaces) SetRect(v *RegisterFaceResponseBodyFacesRect) *RegisterFaceResponseBodyFaces {
	s.Rect = v
	return s
}

type RegisterFaceResponseBodyFacesRect struct {
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s RegisterFaceResponseBodyFacesRect) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceResponseBodyFacesRect) GoString() string {
	return s.String()
}

func (s *RegisterFaceResponseBodyFacesRect) SetLeft(v int32) *RegisterFaceResponseBodyFacesRect {
	s.Left = &v
	return s
}

func (s *RegisterFaceResponseBodyFacesRect) SetTop(v int32) *RegisterFaceResponseBodyFacesRect {
	s.Top = &v
	return s
}

func (s *RegisterFaceResponseBodyFacesRect) SetWidth(v int32) *RegisterFaceResponseBodyFacesRect {
	s.Width = &v
	return s
}

func (s *RegisterFaceResponseBodyFacesRect) SetHeight(v int32) *RegisterFaceResponseBodyFacesRect {
	s.Height = &v
	return s
}

type RegisterFaceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceResponse) GoString() string {
	return s.String()
}

func (s *RegisterFaceResponse) SetHeaders(v map[string]*string) *RegisterFaceResponse {
	s.Headers = v
	return s
}

func (s *RegisterFaceResponse) SetBody(v *RegisterFaceResponseBody) *RegisterFaceResponse {
	s.Body = v
	return s
}

type SearchFaceRequest struct {
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog              *string  `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	GroupId              *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ProbabilityThreshold *float32 `json:"ProbabilityThreshold,omitempty" xml:"ProbabilityThreshold,omitempty"`
	Count                *int32   `json:"Count,omitempty" xml:"Count,omitempty"`
	DataType             *string  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Content              *string  `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SearchFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceRequest) SetOwnerId(v int64) *SearchFaceRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchFaceRequest) SetShowLog(v string) *SearchFaceRequest {
	s.ShowLog = &v
	return s
}

func (s *SearchFaceRequest) SetGroupId(v string) *SearchFaceRequest {
	s.GroupId = &v
	return s
}

func (s *SearchFaceRequest) SetProbabilityThreshold(v float32) *SearchFaceRequest {
	s.ProbabilityThreshold = &v
	return s
}

func (s *SearchFaceRequest) SetCount(v int32) *SearchFaceRequest {
	s.Count = &v
	return s
}

func (s *SearchFaceRequest) SetDataType(v string) *SearchFaceRequest {
	s.DataType = &v
	return s
}

func (s *SearchFaceRequest) SetContent(v string) *SearchFaceRequest {
	s.Content = &v
	return s
}

type SearchFaceResponseBody struct {
	Rect        *SearchFaceResponseBodyRect          `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FaceResults []*SearchFaceResponseBodyFaceResults `json:"FaceResults,omitempty" xml:"FaceResults,omitempty" type:"Repeated"`
	GroupId     *string                              `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s SearchFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBody) SetRect(v *SearchFaceResponseBodyRect) *SearchFaceResponseBody {
	s.Rect = v
	return s
}

func (s *SearchFaceResponseBody) SetRequestId(v string) *SearchFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFaceResponseBody) SetFaceResults(v []*SearchFaceResponseBodyFaceResults) *SearchFaceResponseBody {
	s.FaceResults = v
	return s
}

func (s *SearchFaceResponseBody) SetGroupId(v string) *SearchFaceResponseBody {
	s.GroupId = &v
	return s
}

type SearchFaceResponseBodyRect struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s SearchFaceResponseBodyRect) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyRect) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyRect) SetTop(v int32) *SearchFaceResponseBodyRect {
	s.Top = &v
	return s
}

func (s *SearchFaceResponseBodyRect) SetWidth(v int32) *SearchFaceResponseBodyRect {
	s.Width = &v
	return s
}

func (s *SearchFaceResponseBodyRect) SetHeight(v int32) *SearchFaceResponseBodyRect {
	s.Height = &v
	return s
}

func (s *SearchFaceResponseBodyRect) SetLeft(v int32) *SearchFaceResponseBodyRect {
	s.Left = &v
	return s
}

type SearchFaceResponseBodyFaceResults struct {
	FaceToken   *string  `json:"FaceToken,omitempty" xml:"FaceToken,omitempty"`
	Probability *float32 `json:"Probability,omitempty" xml:"Probability,omitempty"`
}

func (s SearchFaceResponseBodyFaceResults) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyFaceResults) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyFaceResults) SetFaceToken(v string) *SearchFaceResponseBodyFaceResults {
	s.FaceToken = &v
	return s
}

func (s *SearchFaceResponseBodyFaceResults) SetProbability(v float32) *SearchFaceResponseBodyFaceResults {
	s.Probability = &v
	return s
}

type SearchFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchFaceResponse) SetHeaders(v map[string]*string) *SearchFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchFaceResponse) SetBody(v *SearchFaceResponseBody) *SearchFaceResponse {
	s.Body = v
	return s
}

type StartStreamPredictRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s StartStreamPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s StartStreamPredictRequest) GoString() string {
	return s.String()
}

func (s *StartStreamPredictRequest) SetOwnerId(v int64) *StartStreamPredictRequest {
	s.OwnerId = &v
	return s
}

func (s *StartStreamPredictRequest) SetShowLog(v string) *StartStreamPredictRequest {
	s.ShowLog = &v
	return s
}

func (s *StartStreamPredictRequest) SetPredictId(v string) *StartStreamPredictRequest {
	s.PredictId = &v
	return s
}

type StartStreamPredictResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s StartStreamPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartStreamPredictResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamPredictResponseBody) SetRequestId(v string) *StartStreamPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartStreamPredictResponseBody) SetPredictId(v string) *StartStreamPredictResponseBody {
	s.PredictId = &v
	return s
}

type StartStreamPredictResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartStreamPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartStreamPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s StartStreamPredictResponse) GoString() string {
	return s.String()
}

func (s *StartStreamPredictResponse) SetHeaders(v map[string]*string) *StartStreamPredictResponse {
	s.Headers = v
	return s
}

func (s *StartStreamPredictResponse) SetBody(v *StartStreamPredictResponseBody) *StartStreamPredictResponse {
	s.Body = v
	return s
}

type StopStreamPredictRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s StopStreamPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStreamPredictRequest) GoString() string {
	return s.String()
}

func (s *StopStreamPredictRequest) SetOwnerId(v int64) *StopStreamPredictRequest {
	s.OwnerId = &v
	return s
}

func (s *StopStreamPredictRequest) SetShowLog(v string) *StopStreamPredictRequest {
	s.ShowLog = &v
	return s
}

func (s *StopStreamPredictRequest) SetPredictId(v string) *StopStreamPredictRequest {
	s.PredictId = &v
	return s
}

type StopStreamPredictResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PredictId *string `json:"PredictId,omitempty" xml:"PredictId,omitempty"`
}

func (s StopStreamPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStreamPredictResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamPredictResponseBody) SetRequestId(v string) *StopStreamPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopStreamPredictResponseBody) SetPredictId(v string) *StopStreamPredictResponseBody {
	s.PredictId = &v
	return s
}

type StopStreamPredictResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopStreamPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStreamPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStreamPredictResponse) GoString() string {
	return s.String()
}

func (s *StopStreamPredictResponse) SetHeaders(v map[string]*string) *StopStreamPredictResponse {
	s.Headers = v
	return s
}

func (s *StopStreamPredictResponse) SetBody(v *StopStreamPredictResponseBody) *StopStreamPredictResponse {
	s.Body = v
	return s
}

type UnregisterFaceRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	FaceToken *string `json:"FaceToken,omitempty" xml:"FaceToken,omitempty"`
}

func (s UnregisterFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnregisterFaceRequest) GoString() string {
	return s.String()
}

func (s *UnregisterFaceRequest) SetOwnerId(v int64) *UnregisterFaceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnregisterFaceRequest) SetShowLog(v string) *UnregisterFaceRequest {
	s.ShowLog = &v
	return s
}

func (s *UnregisterFaceRequest) SetGroupId(v string) *UnregisterFaceRequest {
	s.GroupId = &v
	return s
}

func (s *UnregisterFaceRequest) SetFaceToken(v string) *UnregisterFaceRequest {
	s.FaceToken = &v
	return s
}

type UnregisterFaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	FaceToken *string `json:"FaceToken,omitempty" xml:"FaceToken,omitempty"`
}

func (s UnregisterFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnregisterFaceResponseBody) GoString() string {
	return s.String()
}

func (s *UnregisterFaceResponseBody) SetRequestId(v string) *UnregisterFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnregisterFaceResponseBody) SetGroupId(v string) *UnregisterFaceResponseBody {
	s.GroupId = &v
	return s
}

func (s *UnregisterFaceResponseBody) SetFaceToken(v string) *UnregisterFaceResponseBody {
	s.FaceToken = &v
	return s
}

type UnregisterFaceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnregisterFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnregisterFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnregisterFaceResponse) GoString() string {
	return s.String()
}

func (s *UnregisterFaceResponse) SetHeaders(v map[string]*string) *UnregisterFaceResponse {
	s.Headers = v
	return s
}

func (s *UnregisterFaceResponse) SetBody(v *UnregisterFaceResponseBody) *UnregisterFaceResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ivision"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateFaceGroupWithOptions(request *CreateFaceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateFaceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateFaceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFaceGroup"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFaceGroup(request *CreateFaceGroupRequest) (_result *CreateFaceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFaceGroupResponse{}
	_body, _err := client.CreateFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStreamPredictWithOptions(request *CreateStreamPredictRequest, runtime *util.RuntimeOptions) (_result *CreateStreamPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStreamPredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStreamPredict"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStreamPredict(request *CreateStreamPredictRequest) (_result *CreateStreamPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStreamPredictResponse{}
	_body, _err := client.CreateStreamPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceGroupWithOptions(request *DeleteFaceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteFaceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceGroup"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceGroup(request *DeleteFaceGroupRequest) (_result *DeleteFaceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceGroupResponse{}
	_body, _err := client.DeleteFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStreamPredictWithOptions(request *DeleteStreamPredictRequest, runtime *util.RuntimeOptions) (_result *DeleteStreamPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteStreamPredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteStreamPredict"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStreamPredict(request *DeleteStreamPredictRequest) (_result *DeleteStreamPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStreamPredictResponse{}
	_body, _err := client.DeleteStreamPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFaceGroupsWithOptions(request *DescribeFaceGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeFaceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeFaceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFaceGroups"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFaceGroups(request *DescribeFaceGroupsRequest) (_result *DescribeFaceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaceGroupsResponse{}
	_body, _err := client.DescribeFaceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamPredictResultWithOptions(request *DescribeStreamPredictResultRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamPredictResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStreamPredictResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStreamPredictResult"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStreamPredictResult(request *DescribeStreamPredictResultRequest) (_result *DescribeStreamPredictResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamPredictResultResponse{}
	_body, _err := client.DescribeStreamPredictResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStreamPredictsWithOptions(request *DescribeStreamPredictsRequest, runtime *util.RuntimeOptions) (_result *DescribeStreamPredictsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStreamPredictsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStreamPredicts"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStreamPredicts(request *DescribeStreamPredictsRequest) (_result *DescribeStreamPredictsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStreamPredictsResponse{}
	_body, _err := client.DescribeStreamPredictsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmDetailWithOptions(request *GetAlgorithmDetailRequest, runtime *util.RuntimeOptions) (_result *GetAlgorithmDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAlgorithmDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAlgorithmDetail"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmDetail(request *GetAlgorithmDetailRequest) (_result *GetAlgorithmDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlgorithmDetailResponse{}
	_body, _err := client.GetAlgorithmDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmHistogramsWithOptions(request *GetAlgorithmHistogramsRequest, runtime *util.RuntimeOptions) (_result *GetAlgorithmHistogramsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAlgorithmHistogramsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAlgorithmHistograms"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmHistograms(request *GetAlgorithmHistogramsRequest) (_result *GetAlgorithmHistogramsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlgorithmHistogramsResponse{}
	_body, _err := client.GetAlgorithmHistogramsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImagePredictWithOptions(request *ImagePredictRequest, runtime *util.RuntimeOptions) (_result *ImagePredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ImagePredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImagePredict"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImagePredict(request *ImagePredictRequest) (_result *ImagePredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImagePredictResponse{}
	_body, _err := client.ImagePredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMyAlgorithmWithOptions(request *ListMyAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ListMyAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMyAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMyAlgorithm"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMyAlgorithm(request *ListMyAlgorithmRequest) (_result *ListMyAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMyAlgorithmResponse{}
	_body, _err := client.ListMyAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PredictPictureWithOptions(request *PredictPictureRequest, runtime *util.RuntimeOptions) (_result *PredictPictureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PredictPictureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PredictPicture"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictPicture(request *PredictPictureRequest) (_result *PredictPictureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictPictureResponse{}
	_body, _err := client.PredictPictureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterFaceWithOptions(request *RegisterFaceRequest, runtime *util.RuntimeOptions) (_result *RegisterFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RegisterFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterFace"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterFace(request *RegisterFaceRequest) (_result *RegisterFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterFaceResponse{}
	_body, _err := client.RegisterFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFaceWithOptions(request *SearchFaceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SearchFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchFace"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFace(request *SearchFaceRequest) (_result *SearchFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchFaceResponse{}
	_body, _err := client.SearchFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartStreamPredictWithOptions(request *StartStreamPredictRequest, runtime *util.RuntimeOptions) (_result *StartStreamPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartStreamPredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartStreamPredict"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartStreamPredict(request *StartStreamPredictRequest) (_result *StartStreamPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartStreamPredictResponse{}
	_body, _err := client.StartStreamPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStreamPredictWithOptions(request *StopStreamPredictRequest, runtime *util.RuntimeOptions) (_result *StopStreamPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopStreamPredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopStreamPredict"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStreamPredict(request *StopStreamPredictRequest) (_result *StopStreamPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopStreamPredictResponse{}
	_body, _err := client.StopStreamPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnregisterFaceWithOptions(request *UnregisterFaceRequest, runtime *util.RuntimeOptions) (_result *UnregisterFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UnregisterFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnregisterFace"), tea.String("2019-03-08"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnregisterFace(request *UnregisterFaceRequest) (_result *UnregisterFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnregisterFaceResponse{}
	_body, _err := client.UnregisterFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
