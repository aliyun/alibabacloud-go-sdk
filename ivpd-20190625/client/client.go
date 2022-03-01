// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type ChangeImageSizeRequest struct {
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeImageSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeRequest) SetHeight(v int32) *ChangeImageSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeRequest) SetUrl(v string) *ChangeImageSizeRequest {
	s.Url = &v
	return s
}

func (s *ChangeImageSizeRequest) SetWidth(v int32) *ChangeImageSizeRequest {
	s.Width = &v
	return s
}

type ChangeImageSizeResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ChangeImageSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeImageSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBody) SetCode(v string) *ChangeImageSizeResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeImageSizeResponseBody) SetData(v *ChangeImageSizeResponseBodyData) *ChangeImageSizeResponseBody {
	s.Data = v
	return s
}

func (s *ChangeImageSizeResponseBody) SetMessage(v string) *ChangeImageSizeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeImageSizeResponseBody) SetRequestId(v string) *ChangeImageSizeResponseBody {
	s.RequestId = &v
	return s
}

type ChangeImageSizeResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeImageSizeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyData) SetUrl(v string) *ChangeImageSizeResponseBodyData {
	s.Url = &v
	return s
}

type ChangeImageSizeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeImageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeImageSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponse) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponse) SetHeaders(v map[string]*string) *ChangeImageSizeResponse {
	s.Headers = v
	return s
}

func (s *ChangeImageSizeResponse) SetBody(v *ChangeImageSizeResponseBody) *ChangeImageSizeResponse {
	s.Body = v
	return s
}

type CreateSegmentBodyJobRequest struct {
	DataList   []*CreateSegmentBodyJobRequestDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	JobId      *string                                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	TimeToLive *int32                                 `json:"TimeToLive,omitempty" xml:"TimeToLive,omitempty"`
}

func (s CreateSegmentBodyJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobRequest) SetDataList(v []*CreateSegmentBodyJobRequestDataList) *CreateSegmentBodyJobRequest {
	s.DataList = v
	return s
}

func (s *CreateSegmentBodyJobRequest) SetJobId(v string) *CreateSegmentBodyJobRequest {
	s.JobId = &v
	return s
}

func (s *CreateSegmentBodyJobRequest) SetTimeToLive(v int32) *CreateSegmentBodyJobRequest {
	s.TimeToLive = &v
	return s
}

type CreateSegmentBodyJobRequestDataList struct {
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s CreateSegmentBodyJobRequestDataList) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobRequestDataList) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobRequestDataList) SetDataId(v string) *CreateSegmentBodyJobRequestDataList {
	s.DataId = &v
	return s
}

func (s *CreateSegmentBodyJobRequestDataList) SetImageUrl(v string) *CreateSegmentBodyJobRequestDataList {
	s.ImageUrl = &v
	return s
}

type CreateSegmentBodyJobResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateSegmentBodyJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSegmentBodyJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobResponseBody) SetCode(v string) *CreateSegmentBodyJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBody) SetData(v *CreateSegmentBodyJobResponseBodyData) *CreateSegmentBodyJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateSegmentBodyJobResponseBody) SetMessage(v string) *CreateSegmentBodyJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBody) SetRequestId(v string) *CreateSegmentBodyJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateSegmentBodyJobResponseBodyData struct {
	BatchSize     *int32                                            `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Completed     *bool                                             `json:"Completed,omitempty" xml:"Completed,omitempty"`
	JobId         *string                                           `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Progress      *int32                                            `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ResultList    []*CreateSegmentBodyJobResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	Status        *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalUsedTime *int64                                            `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s CreateSegmentBodyJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobResponseBodyData) SetBatchSize(v int32) *CreateSegmentBodyJobResponseBodyData {
	s.BatchSize = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyData) SetCompleted(v bool) *CreateSegmentBodyJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyData) SetJobId(v string) *CreateSegmentBodyJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyData) SetProgress(v int32) *CreateSegmentBodyJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyData) SetResultList(v []*CreateSegmentBodyJobResponseBodyDataResultList) *CreateSegmentBodyJobResponseBodyData {
	s.ResultList = v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyData) SetStatus(v string) *CreateSegmentBodyJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyData) SetTotalUsedTime(v int64) *CreateSegmentBodyJobResponseBodyData {
	s.TotalUsedTime = &v
	return s
}

type CreateSegmentBodyJobResponseBodyDataResultList struct {
	Code       *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DataId     *string                                                   `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Message    *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	ResultData *CreateSegmentBodyJobResponseBodyDataResultListResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Struct"`
	Success    *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSegmentBodyJobResponseBodyDataResultList) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobResponseBodyDataResultList) SetCode(v string) *CreateSegmentBodyJobResponseBodyDataResultList {
	s.Code = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyDataResultList) SetDataId(v string) *CreateSegmentBodyJobResponseBodyDataResultList {
	s.DataId = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyDataResultList) SetMessage(v string) *CreateSegmentBodyJobResponseBodyDataResultList {
	s.Message = &v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyDataResultList) SetResultData(v *CreateSegmentBodyJobResponseBodyDataResultListResultData) *CreateSegmentBodyJobResponseBodyDataResultList {
	s.ResultData = v
	return s
}

func (s *CreateSegmentBodyJobResponseBodyDataResultList) SetSuccess(v bool) *CreateSegmentBodyJobResponseBodyDataResultList {
	s.Success = &v
	return s
}

type CreateSegmentBodyJobResponseBodyDataResultListResultData struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s CreateSegmentBodyJobResponseBodyDataResultListResultData) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobResponseBodyDataResultListResultData) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobResponseBodyDataResultListResultData) SetImageUrl(v string) *CreateSegmentBodyJobResponseBodyDataResultListResultData {
	s.ImageUrl = &v
	return s
}

type CreateSegmentBodyJobResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSegmentBodyJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSegmentBodyJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSegmentBodyJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSegmentBodyJobResponse) SetHeaders(v map[string]*string) *CreateSegmentBodyJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSegmentBodyJobResponse) SetBody(v *CreateSegmentBodyJobResponseBody) *CreateSegmentBodyJobResponse {
	s.Body = v
	return s
}

type DetectImageElementsRequest struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectImageElementsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageElementsRequest) SetUrl(v string) *DetectImageElementsRequest {
	s.Url = &v
	return s
}

type DetectImageElementsResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DetectImageElementsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageElementsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponseBody) SetCode(v string) *DetectImageElementsResponseBody {
	s.Code = &v
	return s
}

func (s *DetectImageElementsResponseBody) SetData(v *DetectImageElementsResponseBodyData) *DetectImageElementsResponseBody {
	s.Data = v
	return s
}

func (s *DetectImageElementsResponseBody) SetMessage(v string) *DetectImageElementsResponseBody {
	s.Message = &v
	return s
}

func (s *DetectImageElementsResponseBody) SetRequestId(v string) *DetectImageElementsResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageElementsResponseBodyData struct {
	Elements []*DetectImageElementsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectImageElementsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponseBodyData) SetElements(v []*DetectImageElementsResponseBodyDataElements) *DetectImageElementsResponseBodyData {
	s.Elements = v
	return s
}

type DetectImageElementsResponseBodyDataElements struct {
	Height *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32   `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectImageElementsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponseBodyDataElements) SetHeight(v int32) *DetectImageElementsResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetScore(v float32) *DetectImageElementsResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetType(v string) *DetectImageElementsResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetWidth(v int32) *DetectImageElementsResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetX(v int32) *DetectImageElementsResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetY(v int32) *DetectImageElementsResponseBodyDataElements {
	s.Y = &v
	return s
}

type DetectImageElementsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectImageElementsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageElementsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponse) SetHeaders(v map[string]*string) *DetectImageElementsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageElementsResponse) SetBody(v *DetectImageElementsResponseBody) *DetectImageElementsResponse {
	s.Body = v
	return s
}

type EraseLogoInVideoRequest struct {
	Boxes    []*EraseLogoInVideoRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	JobId    *string                         `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VideoUrl *string                         `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseLogoInVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseLogoInVideoRequest) GoString() string {
	return s.String()
}

func (s *EraseLogoInVideoRequest) SetBoxes(v []*EraseLogoInVideoRequestBoxes) *EraseLogoInVideoRequest {
	s.Boxes = v
	return s
}

func (s *EraseLogoInVideoRequest) SetJobId(v string) *EraseLogoInVideoRequest {
	s.JobId = &v
	return s
}

func (s *EraseLogoInVideoRequest) SetVideoUrl(v string) *EraseLogoInVideoRequest {
	s.VideoUrl = &v
	return s
}

type EraseLogoInVideoRequestBoxes struct {
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s EraseLogoInVideoRequestBoxes) String() string {
	return tea.Prettify(s)
}

func (s EraseLogoInVideoRequestBoxes) GoString() string {
	return s.String()
}

func (s *EraseLogoInVideoRequestBoxes) SetH(v float32) *EraseLogoInVideoRequestBoxes {
	s.H = &v
	return s
}

func (s *EraseLogoInVideoRequestBoxes) SetW(v float32) *EraseLogoInVideoRequestBoxes {
	s.W = &v
	return s
}

func (s *EraseLogoInVideoRequestBoxes) SetX(v float32) *EraseLogoInVideoRequestBoxes {
	s.X = &v
	return s
}

func (s *EraseLogoInVideoRequestBoxes) SetY(v float32) *EraseLogoInVideoRequestBoxes {
	s.Y = &v
	return s
}

type EraseLogoInVideoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *EraseLogoInVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EraseLogoInVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EraseLogoInVideoResponseBody) GoString() string {
	return s.String()
}

func (s *EraseLogoInVideoResponseBody) SetCode(v string) *EraseLogoInVideoResponseBody {
	s.Code = &v
	return s
}

func (s *EraseLogoInVideoResponseBody) SetData(v *EraseLogoInVideoResponseBodyData) *EraseLogoInVideoResponseBody {
	s.Data = v
	return s
}

func (s *EraseLogoInVideoResponseBody) SetMessage(v string) *EraseLogoInVideoResponseBody {
	s.Message = &v
	return s
}

func (s *EraseLogoInVideoResponseBody) SetRequestId(v string) *EraseLogoInVideoResponseBody {
	s.RequestId = &v
	return s
}

type EraseLogoInVideoResponseBodyData struct {
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseLogoInVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EraseLogoInVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *EraseLogoInVideoResponseBodyData) SetJobId(v string) *EraseLogoInVideoResponseBodyData {
	s.JobId = &v
	return s
}

func (s *EraseLogoInVideoResponseBodyData) SetVideoUrl(v string) *EraseLogoInVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type EraseLogoInVideoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EraseLogoInVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EraseLogoInVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s EraseLogoInVideoResponse) GoString() string {
	return s.String()
}

func (s *EraseLogoInVideoResponse) SetHeaders(v map[string]*string) *EraseLogoInVideoResponse {
	s.Headers = v
	return s
}

func (s *EraseLogoInVideoResponse) SetBody(v *EraseLogoInVideoResponseBody) *EraseLogoInVideoResponse {
	s.Body = v
	return s
}

type ExtendImageStyleRequest struct {
	MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
	StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ExtendImageStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleRequest) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleRequest) SetMajorUrl(v string) *ExtendImageStyleRequest {
	s.MajorUrl = &v
	return s
}

func (s *ExtendImageStyleRequest) SetStyleUrl(v string) *ExtendImageStyleRequest {
	s.StyleUrl = &v
	return s
}

type ExtendImageStyleResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ExtendImageStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtendImageStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponseBody) SetCode(v string) *ExtendImageStyleResponseBody {
	s.Code = &v
	return s
}

func (s *ExtendImageStyleResponseBody) SetData(v *ExtendImageStyleResponseBodyData) *ExtendImageStyleResponseBody {
	s.Data = v
	return s
}

func (s *ExtendImageStyleResponseBody) SetMessage(v string) *ExtendImageStyleResponseBody {
	s.Message = &v
	return s
}

func (s *ExtendImageStyleResponseBody) SetRequestId(v string) *ExtendImageStyleResponseBody {
	s.RequestId = &v
	return s
}

type ExtendImageStyleResponseBodyData struct {
	MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExtendImageStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponseBodyData) SetMajorUrl(v string) *ExtendImageStyleResponseBodyData {
	s.MajorUrl = &v
	return s
}

func (s *ExtendImageStyleResponseBodyData) SetUrl(v string) *ExtendImageStyleResponseBodyData {
	s.Url = &v
	return s
}

type ExtendImageStyleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtendImageStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtendImageStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponse) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponse) SetHeaders(v map[string]*string) *ExtendImageStyleResponse {
	s.Headers = v
	return s
}

func (s *ExtendImageStyleResponse) SetBody(v *ExtendImageStyleResponseBody) *ExtendImageStyleResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type GetAsyncResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncResultRequest) SetJobId(v string) *GetAsyncResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncResultResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAsyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponseBody) SetCode(v string) *GetAsyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetData(v *GetAsyncResultResponseBodyData) *GetAsyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncResultResponseBody) SetMessage(v string) *GetAsyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncResultResponseBody) SetRequestId(v string) *GetAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncResultResponseBodyData struct {
	BatchSize     *string                `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Code          *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed     *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Finish        *bool                  `json:"Finish,omitempty" xml:"Finish,omitempty"`
	Message       *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Progress      *float32               `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Result        map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	Status        *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalUsedTime *int64                 `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s GetAsyncResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponseBodyData) SetBatchSize(v string) *GetAsyncResultResponseBodyData {
	s.BatchSize = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetCode(v string) *GetAsyncResultResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetCompleted(v bool) *GetAsyncResultResponseBodyData {
	s.Completed = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetFinish(v bool) *GetAsyncResultResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetMessage(v string) *GetAsyncResultResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetProgress(v float32) *GetAsyncResultResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetResult(v map[string]interface{}) *GetAsyncResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetStatus(v string) *GetAsyncResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncResultResponseBodyData) SetTotalUsedTime(v int64) *GetAsyncResultResponseBodyData {
	s.TotalUsedTime = &v
	return s
}

type GetAsyncResultResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncResultResponse) SetHeaders(v map[string]*string) *GetAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncResultResponse) SetBody(v *GetAsyncResultResponseBody) *GetAsyncResultResponse {
	s.Body = v
	return s
}

type GetJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetJobResultRequest) SetJobId(v string) *GetJobResultRequest {
	s.JobId = &v
	return s
}

type GetJobResultResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResultResponseBody) SetCode(v string) *GetJobResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobResultResponseBody) SetData(v *GetJobResultResponseBodyData) *GetJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResultResponseBody) SetMessage(v string) *GetJobResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobResultResponseBody) SetRequestId(v string) *GetJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetJobResultResponseBodyData struct {
	BatchSize     *string                  `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Completed     *bool                    `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Finish        *bool                    `json:"Finish,omitempty" xml:"Finish,omitempty"`
	Message       *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Progress      *float32                 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ResultList    []map[string]interface{} `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	Status        *string                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalUsedTime *int64                   `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s GetJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobResultResponseBodyData) SetBatchSize(v string) *GetJobResultResponseBodyData {
	s.BatchSize = &v
	return s
}

func (s *GetJobResultResponseBodyData) SetCompleted(v bool) *GetJobResultResponseBodyData {
	s.Completed = &v
	return s
}

func (s *GetJobResultResponseBodyData) SetFinish(v bool) *GetJobResultResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetJobResultResponseBodyData) SetMessage(v string) *GetJobResultResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetJobResultResponseBodyData) SetProgress(v float32) *GetJobResultResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetJobResultResponseBodyData) SetResultList(v []map[string]interface{}) *GetJobResultResponseBodyData {
	s.ResultList = v
	return s
}

func (s *GetJobResultResponseBodyData) SetStatus(v string) *GetJobResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetJobResultResponseBodyData) SetTotalUsedTime(v int64) *GetJobResultResponseBodyData {
	s.TotalUsedTime = &v
	return s
}

type GetJobResultResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetJobResultResponse) SetHeaders(v map[string]*string) *GetJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetJobResultResponse) SetBody(v *GetJobResultResponseBody) *GetJobResultResponse {
	s.Body = v
	return s
}

type GetJobStatusRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

type GetJobStatusResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) SetCode(v string) *GetJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobStatusResponseBody) SetData(v *GetJobStatusResponseBodyData) *GetJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetJobStatusResponseBody) SetMessage(v string) *GetJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetJobStatusResponseBodyData struct {
	BatchSize     *string  `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Completed     *bool    `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Finish        *bool    `json:"Finish,omitempty" xml:"Finish,omitempty"`
	Message       *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	Progress      *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status        *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeToLive    *int32   `json:"TimeToLive,omitempty" xml:"TimeToLive,omitempty"`
	TotalUsedTime *int64   `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s GetJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBodyData) SetBatchSize(v string) *GetJobStatusResponseBodyData {
	s.BatchSize = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetCompleted(v bool) *GetJobStatusResponseBodyData {
	s.Completed = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetFinish(v bool) *GetJobStatusResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetMessage(v string) *GetJobStatusResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetProgress(v float32) *GetJobStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetStatus(v string) *GetJobStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetTimeToLive(v int32) *GetJobStatusResponseBodyData {
	s.TimeToLive = &v
	return s
}

func (s *GetJobStatusResponseBodyData) SetTotalUsedTime(v int64) *GetJobStatusResponseBodyData {
	s.TotalUsedTime = &v
	return s
}

type GetJobStatusResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetHeaders(v map[string]*string) *GetJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetJobStatusResponse) SetBody(v *GetJobStatusResponseBody) *GetJobStatusResponse {
	s.Body = v
	return s
}

type GetUserBucketConfigResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetUserBucketConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserBucketConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserBucketConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBucketConfigResponseBody) SetCode(v string) *GetUserBucketConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserBucketConfigResponseBody) SetData(v []*GetUserBucketConfigResponseBodyData) *GetUserBucketConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetUserBucketConfigResponseBody) SetMessage(v string) *GetUserBucketConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserBucketConfigResponseBody) SetRequestId(v string) *GetUserBucketConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetUserBucketConfigResponseBodyData struct {
	Bucket     *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s GetUserBucketConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserBucketConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserBucketConfigResponseBodyData) SetBucket(v string) *GetUserBucketConfigResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetUserBucketConfigResponseBodyData) SetRegion(v string) *GetUserBucketConfigResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetUserBucketConfigResponseBodyData) SetRegionName(v string) *GetUserBucketConfigResponseBodyData {
	s.RegionName = &v
	return s
}

type GetUserBucketConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserBucketConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserBucketConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserBucketConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUserBucketConfigResponse) SetHeaders(v map[string]*string) *GetUserBucketConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUserBucketConfigResponse) SetBody(v *GetUserBucketConfigResponseBody) *GetUserBucketConfigResponse {
	s.Body = v
	return s
}

type HighlightGameVideoRequest struct {
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s HighlightGameVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoRequest) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoRequest) SetAsync(v bool) *HighlightGameVideoRequest {
	s.Async = &v
	return s
}

func (s *HighlightGameVideoRequest) SetVideoUrl(v string) *HighlightGameVideoRequest {
	s.VideoUrl = &v
	return s
}

type HighlightGameVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s HighlightGameVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *HighlightGameVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *HighlightGameVideoAdvanceRequest) SetAsync(v bool) *HighlightGameVideoAdvanceRequest {
	s.Async = &v
	return s
}

type HighlightGameVideoResponseBody struct {
	Data      *HighlightGameVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HighlightGameVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoResponseBody) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoResponseBody) SetData(v *HighlightGameVideoResponseBodyData) *HighlightGameVideoResponseBody {
	s.Data = v
	return s
}

func (s *HighlightGameVideoResponseBody) SetRequestId(v string) *HighlightGameVideoResponseBody {
	s.RequestId = &v
	return s
}

type HighlightGameVideoResponseBodyData struct {
	GameList      []*HighlightGameVideoResponseBodyDataGameList      `json:"GameList,omitempty" xml:"GameList,omitempty" type:"Repeated"`
	HighlightList []*HighlightGameVideoResponseBodyDataHighlightList `json:"HighlightList,omitempty" xml:"HighlightList,omitempty" type:"Repeated"`
}

func (s HighlightGameVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoResponseBodyData) SetGameList(v []*HighlightGameVideoResponseBodyDataGameList) *HighlightGameVideoResponseBodyData {
	s.GameList = v
	return s
}

func (s *HighlightGameVideoResponseBodyData) SetHighlightList(v []*HighlightGameVideoResponseBodyDataHighlightList) *HighlightGameVideoResponseBodyData {
	s.HighlightList = v
	return s
}

type HighlightGameVideoResponseBodyDataGameList struct {
	End      *float32               `json:"End,omitempty" xml:"End,omitempty"`
	GameInfo map[string]interface{} `json:"GameInfo,omitempty" xml:"GameInfo,omitempty"`
	Id       *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	Start    *float32               `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s HighlightGameVideoResponseBodyDataGameList) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoResponseBodyDataGameList) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoResponseBodyDataGameList) SetEnd(v float32) *HighlightGameVideoResponseBodyDataGameList {
	s.End = &v
	return s
}

func (s *HighlightGameVideoResponseBodyDataGameList) SetGameInfo(v map[string]interface{}) *HighlightGameVideoResponseBodyDataGameList {
	s.GameInfo = v
	return s
}

func (s *HighlightGameVideoResponseBodyDataGameList) SetId(v string) *HighlightGameVideoResponseBodyDataGameList {
	s.Id = &v
	return s
}

func (s *HighlightGameVideoResponseBodyDataGameList) SetStart(v float32) *HighlightGameVideoResponseBodyDataGameList {
	s.Start = &v
	return s
}

type HighlightGameVideoResponseBodyDataHighlightList struct {
	End   *float32 `json:"End,omitempty" xml:"End,omitempty"`
	Start *float32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s HighlightGameVideoResponseBodyDataHighlightList) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoResponseBodyDataHighlightList) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoResponseBodyDataHighlightList) SetEnd(v float32) *HighlightGameVideoResponseBodyDataHighlightList {
	s.End = &v
	return s
}

func (s *HighlightGameVideoResponseBodyDataHighlightList) SetStart(v float32) *HighlightGameVideoResponseBodyDataHighlightList {
	s.Start = &v
	return s
}

type HighlightGameVideoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HighlightGameVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HighlightGameVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s HighlightGameVideoResponse) GoString() string {
	return s.String()
}

func (s *HighlightGameVideoResponse) SetHeaders(v map[string]*string) *HighlightGameVideoResponse {
	s.Headers = v
	return s
}

func (s *HighlightGameVideoResponse) SetBody(v *HighlightGameVideoResponseBody) *HighlightGameVideoResponse {
	s.Body = v
	return s
}

type ListPackageDesignModelTypesResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListPackageDesignModelTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPackageDesignModelTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPackageDesignModelTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPackageDesignModelTypesResponseBody) SetCode(v string) *ListPackageDesignModelTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPackageDesignModelTypesResponseBody) SetData(v *ListPackageDesignModelTypesResponseBodyData) *ListPackageDesignModelTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListPackageDesignModelTypesResponseBody) SetMessage(v string) *ListPackageDesignModelTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPackageDesignModelTypesResponseBody) SetRequestId(v string) *ListPackageDesignModelTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListPackageDesignModelTypesResponseBodyData struct {
	ModelTypeList []*ListPackageDesignModelTypesResponseBodyDataModelTypeList `json:"ModelTypeList,omitempty" xml:"ModelTypeList,omitempty" type:"Repeated"`
}

func (s ListPackageDesignModelTypesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPackageDesignModelTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPackageDesignModelTypesResponseBodyData) SetModelTypeList(v []*ListPackageDesignModelTypesResponseBodyDataModelTypeList) *ListPackageDesignModelTypesResponseBodyData {
	s.ModelTypeList = v
	return s
}

type ListPackageDesignModelTypesResponseBodyDataModelTypeList struct {
	Elements  []*ListPackageDesignModelTypesResponseBodyDataModelTypeListElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	ModelType *string                                                             `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
}

func (s ListPackageDesignModelTypesResponseBodyDataModelTypeList) String() string {
	return tea.Prettify(s)
}

func (s ListPackageDesignModelTypesResponseBodyDataModelTypeList) GoString() string {
	return s.String()
}

func (s *ListPackageDesignModelTypesResponseBodyDataModelTypeList) SetElements(v []*ListPackageDesignModelTypesResponseBodyDataModelTypeListElements) *ListPackageDesignModelTypesResponseBodyDataModelTypeList {
	s.Elements = v
	return s
}

func (s *ListPackageDesignModelTypesResponseBodyDataModelTypeList) SetModelType(v string) *ListPackageDesignModelTypesResponseBodyDataModelTypeList {
	s.ModelType = &v
	return s
}

type ListPackageDesignModelTypesResponseBodyDataModelTypeListElements struct {
	SideName *string `json:"SideName,omitempty" xml:"SideName,omitempty"`
}

func (s ListPackageDesignModelTypesResponseBodyDataModelTypeListElements) String() string {
	return tea.Prettify(s)
}

func (s ListPackageDesignModelTypesResponseBodyDataModelTypeListElements) GoString() string {
	return s.String()
}

func (s *ListPackageDesignModelTypesResponseBodyDataModelTypeListElements) SetSideName(v string) *ListPackageDesignModelTypesResponseBodyDataModelTypeListElements {
	s.SideName = &v
	return s
}

type ListPackageDesignModelTypesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPackageDesignModelTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPackageDesignModelTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPackageDesignModelTypesResponse) GoString() string {
	return s.String()
}

func (s *ListPackageDesignModelTypesResponse) SetHeaders(v map[string]*string) *ListPackageDesignModelTypesResponse {
	s.Headers = v
	return s
}

func (s *ListPackageDesignModelTypesResponse) SetBody(v *ListPackageDesignModelTypesResponseBody) *ListPackageDesignModelTypesResponse {
	s.Body = v
	return s
}

type ListUserBucketsRequest struct {
	Data []*ListUserBucketsRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListUserBucketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListUserBucketsRequest) SetData(v []*ListUserBucketsRequestData) *ListUserBucketsRequest {
	s.Data = v
	return s
}

type ListUserBucketsRequestData struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserBucketsRequestData) String() string {
	return tea.Prettify(s)
}

func (s ListUserBucketsRequestData) GoString() string {
	return s.String()
}

func (s *ListUserBucketsRequestData) SetRegionId(v string) *ListUserBucketsRequestData {
	s.RegionId = &v
	return s
}

type ListUserBucketsResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserBucketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserBucketsResponseBody) SetCode(v string) *ListUserBucketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserBucketsResponseBody) SetData(v []*string) *ListUserBucketsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserBucketsResponseBody) SetMessage(v string) *ListUserBucketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserBucketsResponseBody) SetRequestId(v string) *ListUserBucketsResponseBody {
	s.RequestId = &v
	return s
}

type ListUserBucketsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserBucketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListUserBucketsResponse) SetHeaders(v map[string]*string) *ListUserBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListUserBucketsResponse) SetBody(v *ListUserBucketsResponseBody) *ListUserBucketsResponse {
	s.Body = v
	return s
}

type MakeSuperResolutionImageRequest struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageRequest) SetUrl(v string) *MakeSuperResolutionImageRequest {
	s.Url = &v
	return s
}

type MakeSuperResolutionImageResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MakeSuperResolutionImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MakeSuperResolutionImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBody) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBody) SetCode(v string) *MakeSuperResolutionImageResponseBody {
	s.Code = &v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) SetData(v *MakeSuperResolutionImageResponseBodyData) *MakeSuperResolutionImageResponseBody {
	s.Data = v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) SetMessage(v string) *MakeSuperResolutionImageResponseBody {
	s.Message = &v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) SetRequestId(v string) *MakeSuperResolutionImageResponseBody {
	s.RequestId = &v
	return s
}

type MakeSuperResolutionImageResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBodyData) SetUrl(v string) *MakeSuperResolutionImageResponseBodyData {
	s.Url = &v
	return s
}

type MakeSuperResolutionImageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MakeSuperResolutionImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MakeSuperResolutionImageResponse) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponse) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponse) SetHeaders(v map[string]*string) *MakeSuperResolutionImageResponse {
	s.Headers = v
	return s
}

func (s *MakeSuperResolutionImageResponse) SetBody(v *MakeSuperResolutionImageResponseBody) *MakeSuperResolutionImageResponse {
	s.Body = v
	return s
}

type RecognizeImageColorRequest struct {
	ColorCount *string `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeImageColorRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorRequest) SetColorCount(v string) *RecognizeImageColorRequest {
	s.ColorCount = &v
	return s
}

func (s *RecognizeImageColorRequest) SetUrl(v string) *RecognizeImageColorRequest {
	s.Url = &v
	return s
}

type RecognizeImageColorResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RecognizeImageColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeImageColorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponseBody) SetCode(v string) *RecognizeImageColorResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeImageColorResponseBody) SetData(v *RecognizeImageColorResponseBodyData) *RecognizeImageColorResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeImageColorResponseBody) SetMessage(v string) *RecognizeImageColorResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeImageColorResponseBody) SetRequestId(v string) *RecognizeImageColorResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeImageColorResponseBodyData struct {
	ColorTemplateList []*RecognizeImageColorResponseBodyDataColorTemplateList `json:"ColorTemplateList,omitempty" xml:"ColorTemplateList,omitempty" type:"Repeated"`
}

func (s RecognizeImageColorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponseBodyData) SetColorTemplateList(v []*RecognizeImageColorResponseBodyDataColorTemplateList) *RecognizeImageColorResponseBodyData {
	s.ColorTemplateList = v
	return s
}

type RecognizeImageColorResponseBodyDataColorTemplateList struct {
	Color      *string  `json:"Color,omitempty" xml:"Color,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Percentage *float32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s RecognizeImageColorResponseBodyDataColorTemplateList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponseBodyDataColorTemplateList) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponseBodyDataColorTemplateList) SetColor(v string) *RecognizeImageColorResponseBodyDataColorTemplateList {
	s.Color = &v
	return s
}

func (s *RecognizeImageColorResponseBodyDataColorTemplateList) SetLabel(v string) *RecognizeImageColorResponseBodyDataColorTemplateList {
	s.Label = &v
	return s
}

func (s *RecognizeImageColorResponseBodyDataColorTemplateList) SetPercentage(v float32) *RecognizeImageColorResponseBodyDataColorTemplateList {
	s.Percentage = &v
	return s
}

type RecognizeImageColorResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeImageColorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeImageColorResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponse) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponse) SetHeaders(v map[string]*string) *RecognizeImageColorResponse {
	s.Headers = v
	return s
}

func (s *RecognizeImageColorResponse) SetBody(v *RecognizeImageColorResponseBody) *RecognizeImageColorResponse {
	s.Body = v
	return s
}

type RecognizeImageStyleRequest struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeImageStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleRequest) SetUrl(v string) *RecognizeImageStyleRequest {
	s.Url = &v
	return s
}

type RecognizeImageStyleResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RecognizeImageStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeImageStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleResponseBody) SetCode(v string) *RecognizeImageStyleResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeImageStyleResponseBody) SetData(v *RecognizeImageStyleResponseBodyData) *RecognizeImageStyleResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeImageStyleResponseBody) SetMessage(v string) *RecognizeImageStyleResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeImageStyleResponseBody) SetRequestId(v string) *RecognizeImageStyleResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeImageStyleResponseBodyData struct {
	Styles []*string `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
}

func (s RecognizeImageStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleResponseBodyData) SetStyles(v []*string) *RecognizeImageStyleResponseBodyData {
	s.Styles = v
	return s
}

type RecognizeImageStyleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeImageStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeImageStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleResponse) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleResponse) SetHeaders(v map[string]*string) *RecognizeImageStyleResponse {
	s.Headers = v
	return s
}

func (s *RecognizeImageStyleResponse) SetBody(v *RecognizeImageStyleResponseBody) *RecognizeImageStyleResponse {
	s.Body = v
	return s
}

type RecolorImageRequest struct {
	ColorCount    *int32                              `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	ColorTemplate []*RecolorImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	Mode          *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrl        *string                             `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	Url           *string                             `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorImageRequest) SetColorCount(v int32) *RecolorImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorImageRequest) SetColorTemplate(v []*RecolorImageRequestColorTemplate) *RecolorImageRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorImageRequest) SetMode(v string) *RecolorImageRequest {
	s.Mode = &v
	return s
}

func (s *RecolorImageRequest) SetRefUrl(v string) *RecolorImageRequest {
	s.RefUrl = &v
	return s
}

func (s *RecolorImageRequest) SetUrl(v string) *RecolorImageRequest {
	s.Url = &v
	return s
}

type RecolorImageRequestColorTemplate struct {
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorImageRequestColorTemplate) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorImageRequestColorTemplate) SetColor(v string) *RecolorImageRequestColorTemplate {
	s.Color = &v
	return s
}

type RecolorImageResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RecolorImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecolorImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBody) SetCode(v string) *RecolorImageResponseBody {
	s.Code = &v
	return s
}

func (s *RecolorImageResponseBody) SetData(v *RecolorImageResponseBodyData) *RecolorImageResponseBody {
	s.Data = v
	return s
}

func (s *RecolorImageResponseBody) SetMessage(v string) *RecolorImageResponseBody {
	s.Message = &v
	return s
}

func (s *RecolorImageResponseBody) SetRequestId(v string) *RecolorImageResponseBody {
	s.RequestId = &v
	return s
}

type RecolorImageResponseBodyData struct {
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
}

func (s RecolorImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBodyData) SetImageList(v []*string) *RecolorImageResponseBodyData {
	s.ImageList = v
	return s
}

type RecolorImageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecolorImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecolorImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponse) GoString() string {
	return s.String()
}

func (s *RecolorImageResponse) SetHeaders(v map[string]*string) *RecolorImageResponse {
	s.Headers = v
	return s
}

func (s *RecolorImageResponse) SetBody(v *RecolorImageResponseBody) *RecolorImageResponse {
	s.Body = v
	return s
}

type SegmentBodyRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyRequest) SetImageUrl(v string) *SegmentBodyRequest {
	s.ImageUrl = &v
	return s
}

type SegmentBodyResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SegmentBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseBody) SetCode(v string) *SegmentBodyResponseBody {
	s.Code = &v
	return s
}

func (s *SegmentBodyResponseBody) SetData(v *SegmentBodyResponseBodyData) *SegmentBodyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentBodyResponseBody) SetMessage(v string) *SegmentBodyResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentBodyResponseBody) SetRequestId(v string) *SegmentBodyResponseBody {
	s.RequestId = &v
	return s
}

type SegmentBodyResponseBodyData struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseBodyData) SetImageUrl(v string) *SegmentBodyResponseBodyData {
	s.ImageUrl = &v
	return s
}

type SegmentBodyResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SegmentBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponse) SetHeaders(v map[string]*string) *SegmentBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentBodyResponse) SetBody(v *SegmentBodyResponseBody) *SegmentBodyResponse {
	s.Body = v
	return s
}

type SegmentCommodityRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityRequest) SetImageURL(v string) *SegmentCommodityRequest {
	s.ImageURL = &v
	return s
}

type SegmentCommodityResponseBody struct {
	Data      *SegmentCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseBody) SetData(v *SegmentCommodityResponseBodyData) *SegmentCommodityResponseBody {
	s.Data = v
	return s
}

func (s *SegmentCommodityResponseBody) SetRequestId(v string) *SegmentCommodityResponseBody {
	s.RequestId = &v
	return s
}

type SegmentCommodityResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentCommodityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseBodyData) SetImageURL(v string) *SegmentCommodityResponseBodyData {
	s.ImageURL = &v
	return s
}

type SegmentCommodityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SegmentCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentCommodityResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponse) SetHeaders(v map[string]*string) *SegmentCommodityResponse {
	s.Headers = v
	return s
}

func (s *SegmentCommodityResponse) SetBody(v *SegmentCommodityResponseBody) *SegmentCommodityResponse {
	s.Body = v
	return s
}

type SegmentImageRequest struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SegmentImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentImageRequest) SetUrl(v string) *SegmentImageRequest {
	s.Url = &v
	return s
}

type SegmentImageResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SegmentImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentImageResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentImageResponseBody) SetCode(v string) *SegmentImageResponseBody {
	s.Code = &v
	return s
}

func (s *SegmentImageResponseBody) SetData(v *SegmentImageResponseBodyData) *SegmentImageResponseBody {
	s.Data = v
	return s
}

func (s *SegmentImageResponseBody) SetMessage(v string) *SegmentImageResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentImageResponseBody) SetRequestId(v string) *SegmentImageResponseBody {
	s.RequestId = &v
	return s
}

type SegmentImageResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SegmentImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentImageResponseBodyData) SetUrl(v string) *SegmentImageResponseBodyData {
	s.Url = &v
	return s
}

type SegmentImageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SegmentImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentImageResponse) SetHeaders(v map[string]*string) *SegmentImageResponse {
	s.Headers = v
	return s
}

func (s *SegmentImageResponse) SetBody(v *SegmentImageResponseBody) *SegmentImageResponse {
	s.Body = v
	return s
}

type UpdateUserBucketConfigRequest struct {
	Data []*UpdateUserBucketConfigRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s UpdateUserBucketConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserBucketConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserBucketConfigRequest) SetData(v []*UpdateUserBucketConfigRequestData) *UpdateUserBucketConfigRequest {
	s.Data = v
	return s
}

type UpdateUserBucketConfigRequestData struct {
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s UpdateUserBucketConfigRequestData) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserBucketConfigRequestData) GoString() string {
	return s.String()
}

func (s *UpdateUserBucketConfigRequestData) SetBucket(v string) *UpdateUserBucketConfigRequestData {
	s.Bucket = &v
	return s
}

func (s *UpdateUserBucketConfigRequestData) SetRegion(v string) *UpdateUserBucketConfigRequestData {
	s.Region = &v
	return s
}

type UpdateUserBucketConfigResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*UpdateUserBucketConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserBucketConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserBucketConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserBucketConfigResponseBody) SetCode(v string) *UpdateUserBucketConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserBucketConfigResponseBody) SetData(v []*UpdateUserBucketConfigResponseBodyData) *UpdateUserBucketConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUserBucketConfigResponseBody) SetMessage(v string) *UpdateUserBucketConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserBucketConfigResponseBody) SetRequestId(v string) *UpdateUserBucketConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserBucketConfigResponseBodyData struct {
	Bucket     *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s UpdateUserBucketConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserBucketConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateUserBucketConfigResponseBodyData) SetBucket(v string) *UpdateUserBucketConfigResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *UpdateUserBucketConfigResponseBodyData) SetRegion(v string) *UpdateUserBucketConfigResponseBodyData {
	s.Region = &v
	return s
}

func (s *UpdateUserBucketConfigResponseBodyData) SetRegionName(v string) *UpdateUserBucketConfigResponseBodyData {
	s.RegionName = &v
	return s
}

type UpdateUserBucketConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserBucketConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserBucketConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserBucketConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserBucketConfigResponse) SetHeaders(v map[string]*string) *UpdateUserBucketConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserBucketConfigResponse) SetBody(v *UpdateUserBucketConfigResponseBody) *UpdateUserBucketConfigResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ivpd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ChangeImageSizeWithOptions(request *ChangeImageSizeRequest, runtime *util.RuntimeOptions) (_result *ChangeImageSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeImageSize"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeImageSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeImageSize(request *ChangeImageSizeRequest) (_result *ChangeImageSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeImageSizeResponse{}
	_body, _err := client.ChangeImageSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSegmentBodyJobWithOptions(request *CreateSegmentBodyJobRequest, runtime *util.RuntimeOptions) (_result *CreateSegmentBodyJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataList)) {
		body["DataList"] = request.DataList
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeToLive)) {
		body["TimeToLive"] = request.TimeToLive
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSegmentBodyJob"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSegmentBodyJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSegmentBodyJob(request *CreateSegmentBodyJobRequest) (_result *CreateSegmentBodyJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSegmentBodyJobResponse{}
	_body, _err := client.CreateSegmentBodyJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageElementsWithOptions(request *DetectImageElementsRequest, runtime *util.RuntimeOptions) (_result *DetectImageElementsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageElements"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageElementsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageElements(request *DetectImageElementsRequest) (_result *DetectImageElementsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageElementsResponse{}
	_body, _err := client.DetectImageElementsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EraseLogoInVideoWithOptions(request *EraseLogoInVideoRequest, runtime *util.RuntimeOptions) (_result *EraseLogoInVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Boxes)) {
		body["Boxes"] = request.Boxes
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EraseLogoInVideo"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EraseLogoInVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EraseLogoInVideo(request *EraseLogoInVideoRequest) (_result *EraseLogoInVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EraseLogoInVideoResponse{}
	_body, _err := client.EraseLogoInVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtendImageStyleWithOptions(request *ExtendImageStyleRequest, runtime *util.RuntimeOptions) (_result *ExtendImageStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MajorUrl)) {
		body["MajorUrl"] = request.MajorUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StyleUrl)) {
		body["StyleUrl"] = request.StyleUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtendImageStyle"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtendImageStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtendImageStyle(request *ExtendImageStyleRequest) (_result *ExtendImageStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtendImageStyleResponse{}
	_body, _err := client.ExtendImageStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncResultWithOptions(request *GetAsyncResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncResult"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncResult(request *GetAsyncResultRequest) (_result *GetAsyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncResultResponse{}
	_body, _err := client.GetAsyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobResultWithOptions(request *GetJobResultRequest, runtime *util.RuntimeOptions) (_result *GetJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobResult"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobResult(request *GetJobResultRequest) (_result *GetJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobResultResponse{}
	_body, _err := client.GetJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobStatus"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserBucketConfigWithOptions(runtime *util.RuntimeOptions) (_result *GetUserBucketConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUserBucketConfig"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserBucketConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserBucketConfig() (_result *GetUserBucketConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserBucketConfigResponse{}
	_body, _err := client.GetUserBucketConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HighlightGameVideoWithOptions(request *HighlightGameVideoRequest, runtime *util.RuntimeOptions) (_result *HighlightGameVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		body["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HighlightGameVideo"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HighlightGameVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HighlightGameVideo(request *HighlightGameVideoRequest) (_result *HighlightGameVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HighlightGameVideoResponse{}
	_body, _err := client.HighlightGameVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HighlightGameVideoAdvance(request *HighlightGameVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *HighlightGameVideoResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ivpd"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	highlightGameVideoReq := &HighlightGameVideoRequest{}
	openapiutil.Convert(request, highlightGameVideoReq)
	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		highlightGameVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	highlightGameVideoResp, _err := client.HighlightGameVideoWithOptions(highlightGameVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = highlightGameVideoResp
	return _result, _err
}

func (client *Client) ListPackageDesignModelTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListPackageDesignModelTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListPackageDesignModelTypes"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPackageDesignModelTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPackageDesignModelTypes() (_result *ListPackageDesignModelTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPackageDesignModelTypesResponse{}
	_body, _err := client.ListPackageDesignModelTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserBucketsWithOptions(request *ListUserBucketsRequest, runtime *util.RuntimeOptions) (_result *ListUserBucketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserBuckets"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserBucketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserBuckets(request *ListUserBucketsRequest) (_result *ListUserBucketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserBucketsResponse{}
	_body, _err := client.ListUserBucketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MakeSuperResolutionImageWithOptions(request *MakeSuperResolutionImageRequest, runtime *util.RuntimeOptions) (_result *MakeSuperResolutionImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MakeSuperResolutionImage"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MakeSuperResolutionImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MakeSuperResolutionImage(request *MakeSuperResolutionImageRequest) (_result *MakeSuperResolutionImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MakeSuperResolutionImageResponse{}
	_body, _err := client.MakeSuperResolutionImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeImageColorWithOptions(request *RecognizeImageColorRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColorCount)) {
		body["ColorCount"] = request.ColorCount
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeImageColor"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeImageColorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeImageColor(request *RecognizeImageColorRequest) (_result *RecognizeImageColorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeImageColorResponse{}
	_body, _err := client.RecognizeImageColorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeImageStyleWithOptions(request *RecognizeImageStyleRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeImageStyle"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeImageStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeImageStyle(request *RecognizeImageStyleRequest) (_result *RecognizeImageStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeImageStyleResponse{}
	_body, _err := client.RecognizeImageStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecolorImageWithOptions(request *RecolorImageRequest, runtime *util.RuntimeOptions) (_result *RecolorImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColorCount)) {
		body["ColorCount"] = request.ColorCount
	}

	if !tea.BoolValue(util.IsUnset(request.ColorTemplate)) {
		body["ColorTemplate"] = request.ColorTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RefUrl)) {
		body["RefUrl"] = request.RefUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecolorImage"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecolorImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecolorImage(request *RecolorImageRequest) (_result *RecolorImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecolorImageResponse{}
	_body, _err := client.RecolorImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentBodyWithOptions(request *SegmentBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentBody"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentBodyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentBody(request *SegmentBodyRequest) (_result *SegmentBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentBodyResponse{}
	_body, _err := client.SegmentBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentCommodityWithOptions(request *SegmentCommodityRequest, runtime *util.RuntimeOptions) (_result *SegmentCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentCommodity"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentCommodity(request *SegmentCommodityRequest) (_result *SegmentCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentCommodityResponse{}
	_body, _err := client.SegmentCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentImageWithOptions(request *SegmentImageRequest, runtime *util.RuntimeOptions) (_result *SegmentImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SegmentImage"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentImage(request *SegmentImageRequest) (_result *SegmentImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentImageResponse{}
	_body, _err := client.SegmentImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserBucketConfigWithOptions(request *UpdateUserBucketConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateUserBucketConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserBucketConfig"),
		Version:     tea.String("2019-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserBucketConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserBucketConfig(request *UpdateUserBucketConfigRequest) (_result *UpdateUserBucketConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserBucketConfigResponse{}
	_body, _err := client.UpdateUserBucketConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
