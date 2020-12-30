// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateCoverTaskRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	VideoName     *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	TemplateId    *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Scales        *string `json:"Scales,omitempty" xml:"Scales,omitempty"`
}

func (s CreateCoverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCoverTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCoverTaskRequest) SetApplicationId(v string) *CreateCoverTaskRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateCoverTaskRequest) SetVideoName(v string) *CreateCoverTaskRequest {
	s.VideoName = &v
	return s
}

func (s *CreateCoverTaskRequest) SetVideoUrl(v string) *CreateCoverTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *CreateCoverTaskRequest) SetTemplateId(v int64) *CreateCoverTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateCoverTaskRequest) SetCallbackUrl(v string) *CreateCoverTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateCoverTaskRequest) SetScales(v string) *CreateCoverTaskRequest {
	s.Scales = &v
	return s
}

type CreateCoverTaskResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCoverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCoverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCoverTaskResponseBody) SetTaskId(v int64) *CreateCoverTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateCoverTaskResponseBody) SetRequestId(v string) *CreateCoverTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateCoverTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCoverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCoverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCoverTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCoverTaskResponse) SetHeaders(v map[string]*string) *CreateCoverTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCoverTaskResponse) SetBody(v *CreateCoverTaskResponseBody) *CreateCoverTaskResponse {
	s.Body = v
	return s
}

type CreateFaceGroupRequest struct {
	FaceGroupName *string `json:"FaceGroupName,omitempty" xml:"FaceGroupName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateFaceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceGroupRequest) SetFaceGroupName(v string) *CreateFaceGroupRequest {
	s.FaceGroupName = &v
	return s
}

func (s *CreateFaceGroupRequest) SetDescription(v string) *CreateFaceGroupRequest {
	s.Description = &v
	return s
}

type CreateFaceGroupResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FaceGroupId *int64  `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
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

func (s *CreateFaceGroupResponseBody) SetFaceGroupId(v int64) *CreateFaceGroupResponseBody {
	s.FaceGroupId = &v
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

type CreateFacePersonRequest struct {
	FaceGroupId    *int64  `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	FacePersonName *string `json:"FacePersonName,omitempty" xml:"FacePersonName,omitempty"`
	ImageUrls      *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
}

func (s CreateFacePersonRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFacePersonRequest) GoString() string {
	return s.String()
}

func (s *CreateFacePersonRequest) SetFaceGroupId(v int64) *CreateFacePersonRequest {
	s.FaceGroupId = &v
	return s
}

func (s *CreateFacePersonRequest) SetFacePersonName(v string) *CreateFacePersonRequest {
	s.FacePersonName = &v
	return s
}

func (s *CreateFacePersonRequest) SetImageUrls(v string) *CreateFacePersonRequest {
	s.ImageUrls = &v
	return s
}

type CreateFacePersonResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FacePersonId *int64  `json:"FacePersonId,omitempty" xml:"FacePersonId,omitempty"`
}

func (s CreateFacePersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFacePersonResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFacePersonResponseBody) SetRequestId(v string) *CreateFacePersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFacePersonResponseBody) SetFacePersonId(v int64) *CreateFacePersonResponseBody {
	s.FacePersonId = &v
	return s
}

type CreateFacePersonResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFacePersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFacePersonResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFacePersonResponse) GoString() string {
	return s.String()
}

func (s *CreateFacePersonResponse) SetHeaders(v map[string]*string) *CreateFacePersonResponse {
	s.Headers = v
	return s
}

func (s *CreateFacePersonResponse) SetBody(v *CreateFacePersonResponseBody) *CreateFacePersonResponse {
	s.Body = v
	return s
}

type CreateGifTaskRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	VideoName     *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	TemplateId    *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Scales        *string `json:"Scales,omitempty" xml:"Scales,omitempty"`
}

func (s CreateGifTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGifTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGifTaskRequest) SetApplicationId(v string) *CreateGifTaskRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateGifTaskRequest) SetVideoName(v string) *CreateGifTaskRequest {
	s.VideoName = &v
	return s
}

func (s *CreateGifTaskRequest) SetVideoUrl(v string) *CreateGifTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *CreateGifTaskRequest) SetTemplateId(v int64) *CreateGifTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateGifTaskRequest) SetCallbackUrl(v string) *CreateGifTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateGifTaskRequest) SetScales(v string) *CreateGifTaskRequest {
	s.Scales = &v
	return s
}

type CreateGifTaskResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGifTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGifTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGifTaskResponseBody) SetTaskId(v int64) *CreateGifTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateGifTaskResponseBody) SetRequestId(v string) *CreateGifTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateGifTaskResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGifTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGifTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGifTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGifTaskResponse) SetHeaders(v map[string]*string) *CreateGifTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGifTaskResponse) SetBody(v *CreateGifTaskResponseBody) *CreateGifTaskResponse {
	s.Body = v
	return s
}

type CreateLabelTaskRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	VideoName     *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	TemplateId    *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
}

func (s CreateLabelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskRequest) SetApplicationId(v string) *CreateLabelTaskRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateLabelTaskRequest) SetVideoName(v string) *CreateLabelTaskRequest {
	s.VideoName = &v
	return s
}

func (s *CreateLabelTaskRequest) SetVideoUrl(v string) *CreateLabelTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *CreateLabelTaskRequest) SetTemplateId(v int64) *CreateLabelTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateLabelTaskRequest) SetCallbackUrl(v string) *CreateLabelTaskRequest {
	s.CallbackUrl = &v
	return s
}

type CreateLabelTaskResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLabelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskResponseBody) SetTaskId(v int64) *CreateLabelTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateLabelTaskResponseBody) SetRequestId(v string) *CreateLabelTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateLabelTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLabelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLabelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskResponse) SetHeaders(v map[string]*string) *CreateLabelTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLabelTaskResponse) SetBody(v *CreateLabelTaskResponseBody) *CreateLabelTaskResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetIsDefault(v bool) *CreateTemplateRequest {
	s.IsDefault = &v
	return s
}

func (s *CreateTemplateRequest) SetType(v int32) *CreateTemplateRequest {
	s.Type = &v
	return s
}

type CreateTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteFaceGroupRequest struct {
	FaceGroupId *int64 `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
}

func (s DeleteFaceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupRequest) SetFaceGroupId(v int64) *DeleteFaceGroupRequest {
	s.FaceGroupId = &v
	return s
}

type DeleteFaceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DeleteFaceImageRequest struct {
	FaceGroupId  *int64 `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	FacePersonId *int64 `json:"FacePersonId,omitempty" xml:"FacePersonId,omitempty"`
	FaceImageId  *int64 `json:"FaceImageId,omitempty" xml:"FaceImageId,omitempty"`
}

func (s DeleteFaceImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageRequest) SetFaceGroupId(v int64) *DeleteFaceImageRequest {
	s.FaceGroupId = &v
	return s
}

func (s *DeleteFaceImageRequest) SetFacePersonId(v int64) *DeleteFaceImageRequest {
	s.FacePersonId = &v
	return s
}

func (s *DeleteFaceImageRequest) SetFaceImageId(v int64) *DeleteFaceImageRequest {
	s.FaceImageId = &v
	return s
}

type DeleteFaceImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageResponseBody) SetRequestId(v string) *DeleteFaceImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceImageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFaceImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageResponse) SetHeaders(v map[string]*string) *DeleteFaceImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceImageResponse) SetBody(v *DeleteFaceImageResponseBody) *DeleteFaceImageResponse {
	s.Body = v
	return s
}

type DeleteFacePersonRequest struct {
	FaceGroupId  *int64 `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	FacePersonId *int64 `json:"FacePersonId,omitempty" xml:"FacePersonId,omitempty"`
}

func (s DeleteFacePersonRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacePersonRequest) GoString() string {
	return s.String()
}

func (s *DeleteFacePersonRequest) SetFaceGroupId(v int64) *DeleteFacePersonRequest {
	s.FaceGroupId = &v
	return s
}

func (s *DeleteFacePersonRequest) SetFacePersonId(v int64) *DeleteFacePersonRequest {
	s.FacePersonId = &v
	return s
}

type DeleteFacePersonResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFacePersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacePersonResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFacePersonResponseBody) SetRequestId(v string) *DeleteFacePersonResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFacePersonResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFacePersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFacePersonResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacePersonResponse) GoString() string {
	return s.String()
}

func (s *DeleteFacePersonResponse) SetHeaders(v map[string]*string) *DeleteFacePersonResponse {
	s.Headers = v
	return s
}

func (s *DeleteFacePersonResponse) SetBody(v *DeleteFacePersonResponseBody) *DeleteFacePersonResponse {
	s.Body = v
	return s
}

type GetTaskResultRequest struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultRequest) SetTaskId(v int64) *GetTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetTaskResultResponseBody struct {
	Status    *int32                           `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetTaskResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) SetStatus(v int32) *GetTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetResult(v *GetTaskResultResponseBodyResult) *GetTaskResultResponseBody {
	s.Result = v
	return s
}

type GetTaskResultResponseBodyResult struct {
	ErrorName        *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	VideoName        *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	AnalysisUseTime  *int64  `json:"AnalysisUseTime,omitempty" xml:"AnalysisUseTime,omitempty"`
	ProcessResultUrl *string `json:"ProcessResultUrl,omitempty" xml:"ProcessResultUrl,omitempty"`
	ApplicationId    *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ErrorReason      *string `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty"`
	VideoUrl         *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetTaskResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyResult) SetErrorName(v string) *GetTaskResultResponseBodyResult {
	s.ErrorName = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetErrorMessage(v string) *GetTaskResultResponseBodyResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetErrorCode(v string) *GetTaskResultResponseBodyResult {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetVideoName(v string) *GetTaskResultResponseBodyResult {
	s.VideoName = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetAnalysisUseTime(v int64) *GetTaskResultResponseBodyResult {
	s.AnalysisUseTime = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetProcessResultUrl(v string) *GetTaskResultResponseBodyResult {
	s.ProcessResultUrl = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetApplicationId(v string) *GetTaskResultResponseBodyResult {
	s.ApplicationId = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetErrorReason(v string) *GetTaskResultResponseBodyResult {
	s.ErrorReason = &v
	return s
}

func (s *GetTaskResultResponseBodyResult) SetVideoUrl(v string) *GetTaskResultResponseBodyResult {
	s.VideoUrl = &v
	return s
}

type GetTaskResultResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponse) SetHeaders(v map[string]*string) *GetTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResultResponse) SetBody(v *GetTaskResultResponseBody) *GetTaskResultResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetTaskId(v int64) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetStatus(v int32) *GetTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetTemplateId(v int64) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetTemplateResponseBody struct {
	IsDefault    *bool                  `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Category     *int32                 `json:"Category,omitempty" xml:"Category,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content      map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	TemplateName *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateId   *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetIsDefault(v bool) *GetTemplateResponseBody {
	s.IsDefault = &v
	return s
}

func (s *GetTemplateResponseBody) SetCategory(v int32) *GetTemplateResponseBody {
	s.Category = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetContent(v map[string]interface{}) *GetTemplateResponseBody {
	s.Content = v
	return s
}

func (s *GetTemplateResponseBody) SetCreateTime(v string) *GetTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetUpdateTime(v string) *GetTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateName(v string) *GetTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateId(v string) *GetTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type ListFaceGroupsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFaceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsRequest) SetPageNumber(v int32) *ListFaceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFaceGroupsRequest) SetPageSize(v int32) *ListFaceGroupsRequest {
	s.PageSize = &v
	return s
}

type ListFaceGroupsResponseBody struct {
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FaceGroups []*ListFaceGroupsResponseBodyFaceGroups `json:"FaceGroups,omitempty" xml:"FaceGroups,omitempty" type:"Repeated"`
}

func (s ListFaceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBody) SetTotalCount(v int64) *ListFaceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetRequestId(v string) *ListFaceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetPageSize(v int32) *ListFaceGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetPageNumber(v int32) *ListFaceGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFaceGroupsResponseBody) SetFaceGroups(v []*ListFaceGroupsResponseBodyFaceGroups) *ListFaceGroupsResponseBody {
	s.FaceGroups = v
	return s
}

type ListFaceGroupsResponseBodyFaceGroups struct {
	Description   *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	FaceGroupName *string                                          `json:"FaceGroupName,omitempty" xml:"FaceGroupName,omitempty"`
	PersonCount   *int64                                           `json:"PersonCount,omitempty" xml:"PersonCount,omitempty"`
	ImageCount    *int64                                           `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	FaceGroupId   *int64                                           `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	Templates     []*ListFaceGroupsResponseBodyFaceGroupsTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListFaceGroupsResponseBodyFaceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroups) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetDescription(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.Description = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetFaceGroupName(v string) *ListFaceGroupsResponseBodyFaceGroups {
	s.FaceGroupName = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetPersonCount(v int64) *ListFaceGroupsResponseBodyFaceGroups {
	s.PersonCount = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetImageCount(v int64) *ListFaceGroupsResponseBodyFaceGroups {
	s.ImageCount = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetFaceGroupId(v int64) *ListFaceGroupsResponseBodyFaceGroups {
	s.FaceGroupId = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroups) SetTemplates(v []*ListFaceGroupsResponseBodyFaceGroupsTemplates) *ListFaceGroupsResponseBodyFaceGroups {
	s.Templates = v
	return s
}

type ListFaceGroupsResponseBodyFaceGroupsTemplates struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFaceGroupsResponseBodyFaceGroupsTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponseBodyFaceGroupsTemplates) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponseBodyFaceGroupsTemplates) SetName(v string) *ListFaceGroupsResponseBodyFaceGroupsTemplates {
	s.Name = &v
	return s
}

func (s *ListFaceGroupsResponseBodyFaceGroupsTemplates) SetId(v string) *ListFaceGroupsResponseBodyFaceGroupsTemplates {
	s.Id = &v
	return s
}

type ListFaceGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceGroupsResponse) SetHeaders(v map[string]*string) *ListFaceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceGroupsResponse) SetBody(v *ListFaceGroupsResponseBody) *ListFaceGroupsResponse {
	s.Body = v
	return s
}

type ListFaceImagesRequest struct {
	FaceGroupId  *int64 `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	FacePersonId *int64 `json:"FacePersonId,omitempty" xml:"FacePersonId,omitempty"`
	PageNumber   *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFaceImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceImagesRequest) GoString() string {
	return s.String()
}

func (s *ListFaceImagesRequest) SetFaceGroupId(v int64) *ListFaceImagesRequest {
	s.FaceGroupId = &v
	return s
}

func (s *ListFaceImagesRequest) SetFacePersonId(v int64) *ListFaceImagesRequest {
	s.FacePersonId = &v
	return s
}

func (s *ListFaceImagesRequest) SetPageNumber(v int32) *ListFaceImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFaceImagesRequest) SetPageSize(v int32) *ListFaceImagesRequest {
	s.PageSize = &v
	return s
}

type ListFaceImagesResponseBody struct {
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FaceImages []*ListFaceImagesResponseBodyFaceImages `json:"FaceImages,omitempty" xml:"FaceImages,omitempty" type:"Repeated"`
}

func (s ListFaceImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceImagesResponseBody) SetTotalCount(v int64) *ListFaceImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFaceImagesResponseBody) SetRequestId(v string) *ListFaceImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceImagesResponseBody) SetPageSize(v int32) *ListFaceImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFaceImagesResponseBody) SetPageNumber(v int32) *ListFaceImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFaceImagesResponseBody) SetFaceImages(v []*ListFaceImagesResponseBodyFaceImages) *ListFaceImagesResponseBody {
	s.FaceImages = v
	return s
}

type ListFaceImagesResponseBodyFaceImages struct {
	FaceRectangle []*float32 `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Repeated"`
	ImageUrl      *string    `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	FaceImageId   *int64     `json:"FaceImageId,omitempty" xml:"FaceImageId,omitempty"`
}

func (s ListFaceImagesResponseBodyFaceImages) String() string {
	return tea.Prettify(s)
}

func (s ListFaceImagesResponseBodyFaceImages) GoString() string {
	return s.String()
}

func (s *ListFaceImagesResponseBodyFaceImages) SetFaceRectangle(v []*float32) *ListFaceImagesResponseBodyFaceImages {
	s.FaceRectangle = v
	return s
}

func (s *ListFaceImagesResponseBodyFaceImages) SetImageUrl(v string) *ListFaceImagesResponseBodyFaceImages {
	s.ImageUrl = &v
	return s
}

func (s *ListFaceImagesResponseBodyFaceImages) SetFaceImageId(v int64) *ListFaceImagesResponseBodyFaceImages {
	s.FaceImageId = &v
	return s
}

type ListFaceImagesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFaceImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceImagesResponse) GoString() string {
	return s.String()
}

func (s *ListFaceImagesResponse) SetHeaders(v map[string]*string) *ListFaceImagesResponse {
	s.Headers = v
	return s
}

func (s *ListFaceImagesResponse) SetBody(v *ListFaceImagesResponseBody) *ListFaceImagesResponse {
	s.Body = v
	return s
}

type ListFacePersonsRequest struct {
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	FaceGroupId    *int64  `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	FacePersonName *string `json:"FacePersonName,omitempty" xml:"FacePersonName,omitempty"`
}

func (s ListFacePersonsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFacePersonsRequest) GoString() string {
	return s.String()
}

func (s *ListFacePersonsRequest) SetPageNumber(v int32) *ListFacePersonsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFacePersonsRequest) SetPageSize(v int32) *ListFacePersonsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFacePersonsRequest) SetFaceGroupId(v int64) *ListFacePersonsRequest {
	s.FaceGroupId = &v
	return s
}

func (s *ListFacePersonsRequest) SetFacePersonName(v string) *ListFacePersonsRequest {
	s.FacePersonName = &v
	return s
}

type ListFacePersonsResponseBody struct {
	TotalCount  *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FacePersons []*ListFacePersonsResponseBodyFacePersons `json:"FacePersons,omitempty" xml:"FacePersons,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListFacePersonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFacePersonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFacePersonsResponseBody) SetTotalCount(v int64) *ListFacePersonsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFacePersonsResponseBody) SetFacePersons(v []*ListFacePersonsResponseBodyFacePersons) *ListFacePersonsResponseBody {
	s.FacePersons = v
	return s
}

func (s *ListFacePersonsResponseBody) SetRequestId(v string) *ListFacePersonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFacePersonsResponseBody) SetPageSize(v int32) *ListFacePersonsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFacePersonsResponseBody) SetPageNumber(v int32) *ListFacePersonsResponseBody {
	s.PageNumber = &v
	return s
}

type ListFacePersonsResponseBodyFacePersons struct {
	ImageUrl       *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ImageCount     *int64  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	FacePersonId   *int64  `json:"FacePersonId,omitempty" xml:"FacePersonId,omitempty"`
	FacePersonName *string `json:"FacePersonName,omitempty" xml:"FacePersonName,omitempty"`
}

func (s ListFacePersonsResponseBodyFacePersons) String() string {
	return tea.Prettify(s)
}

func (s ListFacePersonsResponseBodyFacePersons) GoString() string {
	return s.String()
}

func (s *ListFacePersonsResponseBodyFacePersons) SetImageUrl(v string) *ListFacePersonsResponseBodyFacePersons {
	s.ImageUrl = &v
	return s
}

func (s *ListFacePersonsResponseBodyFacePersons) SetImageCount(v int64) *ListFacePersonsResponseBodyFacePersons {
	s.ImageCount = &v
	return s
}

func (s *ListFacePersonsResponseBodyFacePersons) SetFacePersonId(v int64) *ListFacePersonsResponseBodyFacePersons {
	s.FacePersonId = &v
	return s
}

func (s *ListFacePersonsResponseBodyFacePersons) SetFacePersonName(v string) *ListFacePersonsResponseBodyFacePersons {
	s.FacePersonName = &v
	return s
}

type ListFacePersonsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFacePersonsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFacePersonsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFacePersonsResponse) GoString() string {
	return s.String()
}

func (s *ListFacePersonsResponse) SetHeaders(v map[string]*string) *ListFacePersonsResponse {
	s.Headers = v
	return s
}

func (s *ListFacePersonsResponse) SetBody(v *ListFacePersonsResponseBody) *ListFacePersonsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateId(v int64) *ListTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) SetType(v int32) *ListTemplatesRequest {
	s.Type = &v
	return s
}

type ListTemplatesResponseBody struct {
	TotalCount  *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber  *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Templates   []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	TotalAmount *int64                                `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int64) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageNumber(v int32) *ListTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalAmount(v int64) *ListTemplatesResponseBody {
	s.TotalAmount = &v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetType(v int32) *ListTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetIsDefault(v bool) *ListTemplatesResponseBodyTemplates {
	s.IsDefault = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ProcessFaceAlgorithmRequest struct {
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ProcessFaceAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessFaceAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessFaceAlgorithmRequest) SetData(v string) *ProcessFaceAlgorithmRequest {
	s.Data = &v
	return s
}

func (s *ProcessFaceAlgorithmRequest) SetAppKey(v string) *ProcessFaceAlgorithmRequest {
	s.AppKey = &v
	return s
}

type ProcessFaceAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessFaceAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessFaceAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessFaceAlgorithmResponseBody) SetMessage(v string) *ProcessFaceAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessFaceAlgorithmResponseBody) SetRequestId(v string) *ProcessFaceAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessFaceAlgorithmResponseBody) SetData(v string) *ProcessFaceAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessFaceAlgorithmResponseBody) SetCode(v int32) *ProcessFaceAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessFaceAlgorithmResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessFaceAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessFaceAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessFaceAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessFaceAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessFaceAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessFaceAlgorithmResponse) SetBody(v *ProcessFaceAlgorithmResponseBody) *ProcessFaceAlgorithmResponse {
	s.Body = v
	return s
}

type ProcessImageTagAlgorithmRequest struct {
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ProcessImageTagAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessImageTagAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessImageTagAlgorithmRequest) SetData(v string) *ProcessImageTagAlgorithmRequest {
	s.Data = &v
	return s
}

func (s *ProcessImageTagAlgorithmRequest) SetAppKey(v string) *ProcessImageTagAlgorithmRequest {
	s.AppKey = &v
	return s
}

type ProcessImageTagAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessImageTagAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessImageTagAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessImageTagAlgorithmResponseBody) SetMessage(v string) *ProcessImageTagAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessImageTagAlgorithmResponseBody) SetRequestId(v string) *ProcessImageTagAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessImageTagAlgorithmResponseBody) SetData(v string) *ProcessImageTagAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessImageTagAlgorithmResponseBody) SetCode(v int32) *ProcessImageTagAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessImageTagAlgorithmResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessImageTagAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessImageTagAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessImageTagAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessImageTagAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessImageTagAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessImageTagAlgorithmResponse) SetBody(v *ProcessImageTagAlgorithmResponseBody) *ProcessImageTagAlgorithmResponse {
	s.Body = v
	return s
}

type ProcessLandmarkAlgorithmRequest struct {
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ProcessLandmarkAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessLandmarkAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessLandmarkAlgorithmRequest) SetData(v string) *ProcessLandmarkAlgorithmRequest {
	s.Data = &v
	return s
}

func (s *ProcessLandmarkAlgorithmRequest) SetAppKey(v string) *ProcessLandmarkAlgorithmRequest {
	s.AppKey = &v
	return s
}

type ProcessLandmarkAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessLandmarkAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessLandmarkAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessLandmarkAlgorithmResponseBody) SetMessage(v string) *ProcessLandmarkAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessLandmarkAlgorithmResponseBody) SetRequestId(v string) *ProcessLandmarkAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessLandmarkAlgorithmResponseBody) SetData(v string) *ProcessLandmarkAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessLandmarkAlgorithmResponseBody) SetCode(v int32) *ProcessLandmarkAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessLandmarkAlgorithmResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessLandmarkAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessLandmarkAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessLandmarkAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessLandmarkAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessLandmarkAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessLandmarkAlgorithmResponse) SetBody(v *ProcessLandmarkAlgorithmResponseBody) *ProcessLandmarkAlgorithmResponse {
	s.Body = v
	return s
}

type ProcessLogoAlgorithmRequest struct {
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ProcessLogoAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessLogoAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessLogoAlgorithmRequest) SetData(v string) *ProcessLogoAlgorithmRequest {
	s.Data = &v
	return s
}

func (s *ProcessLogoAlgorithmRequest) SetAppKey(v string) *ProcessLogoAlgorithmRequest {
	s.AppKey = &v
	return s
}

type ProcessLogoAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessLogoAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessLogoAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessLogoAlgorithmResponseBody) SetMessage(v string) *ProcessLogoAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessLogoAlgorithmResponseBody) SetRequestId(v string) *ProcessLogoAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessLogoAlgorithmResponseBody) SetData(v string) *ProcessLogoAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessLogoAlgorithmResponseBody) SetCode(v int32) *ProcessLogoAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessLogoAlgorithmResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessLogoAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessLogoAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessLogoAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessLogoAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessLogoAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessLogoAlgorithmResponse) SetBody(v *ProcessLogoAlgorithmResponseBody) *ProcessLogoAlgorithmResponse {
	s.Body = v
	return s
}

type ProcessNewsAlgorithmRequest struct {
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ProcessNewsAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessNewsAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessNewsAlgorithmRequest) SetData(v string) *ProcessNewsAlgorithmRequest {
	s.Data = &v
	return s
}

func (s *ProcessNewsAlgorithmRequest) SetAppKey(v string) *ProcessNewsAlgorithmRequest {
	s.AppKey = &v
	return s
}

type ProcessNewsAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessNewsAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessNewsAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessNewsAlgorithmResponseBody) SetMessage(v string) *ProcessNewsAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessNewsAlgorithmResponseBody) SetRequestId(v string) *ProcessNewsAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessNewsAlgorithmResponseBody) SetData(v string) *ProcessNewsAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessNewsAlgorithmResponseBody) SetCode(v int32) *ProcessNewsAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessNewsAlgorithmResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessNewsAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessNewsAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessNewsAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessNewsAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessNewsAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessNewsAlgorithmResponse) SetBody(v *ProcessNewsAlgorithmResponseBody) *ProcessNewsAlgorithmResponse {
	s.Body = v
	return s
}

type ProcessNlpAlgorithmRequest struct {
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ProcessNlpAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessNlpAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessNlpAlgorithmRequest) SetAppKey(v string) *ProcessNlpAlgorithmRequest {
	s.AppKey = &v
	return s
}

func (s *ProcessNlpAlgorithmRequest) SetData(v string) *ProcessNlpAlgorithmRequest {
	s.Data = &v
	return s
}

type ProcessNlpAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessNlpAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessNlpAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessNlpAlgorithmResponseBody) SetMessage(v string) *ProcessNlpAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessNlpAlgorithmResponseBody) SetRequestId(v string) *ProcessNlpAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessNlpAlgorithmResponseBody) SetData(v string) *ProcessNlpAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessNlpAlgorithmResponseBody) SetCode(v int32) *ProcessNlpAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessNlpAlgorithmResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessNlpAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessNlpAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessNlpAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessNlpAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessNlpAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessNlpAlgorithmResponse) SetBody(v *ProcessNlpAlgorithmResponseBody) *ProcessNlpAlgorithmResponse {
	s.Body = v
	return s
}

type ProcessOcrAlgorithmRequest struct {
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Data   *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ProcessOcrAlgorithmRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessOcrAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ProcessOcrAlgorithmRequest) SetAppKey(v string) *ProcessOcrAlgorithmRequest {
	s.AppKey = &v
	return s
}

func (s *ProcessOcrAlgorithmRequest) SetData(v string) *ProcessOcrAlgorithmRequest {
	s.Data = &v
	return s
}

type ProcessOcrAlgorithmResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ProcessOcrAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessOcrAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessOcrAlgorithmResponseBody) SetMessage(v string) *ProcessOcrAlgorithmResponseBody {
	s.Message = &v
	return s
}

func (s *ProcessOcrAlgorithmResponseBody) SetRequestId(v string) *ProcessOcrAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProcessOcrAlgorithmResponseBody) SetData(v string) *ProcessOcrAlgorithmResponseBody {
	s.Data = &v
	return s
}

func (s *ProcessOcrAlgorithmResponseBody) SetCode(v int32) *ProcessOcrAlgorithmResponseBody {
	s.Code = &v
	return s
}

type ProcessOcrAlgorithmResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessOcrAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessOcrAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessOcrAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ProcessOcrAlgorithmResponse) SetHeaders(v map[string]*string) *ProcessOcrAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ProcessOcrAlgorithmResponse) SetBody(v *ProcessOcrAlgorithmResponseBody) *ProcessOcrAlgorithmResponse {
	s.Body = v
	return s
}

type RegisterFaceImageRequest struct {
	FaceGroupId  *int64  `json:"FaceGroupId,omitempty" xml:"FaceGroupId,omitempty"`
	FacePersonId *int64  `json:"FacePersonId,omitempty" xml:"FacePersonId,omitempty"`
	ImageUrl     *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RegisterFaceImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceImageRequest) GoString() string {
	return s.String()
}

func (s *RegisterFaceImageRequest) SetFaceGroupId(v int64) *RegisterFaceImageRequest {
	s.FaceGroupId = &v
	return s
}

func (s *RegisterFaceImageRequest) SetFacePersonId(v int64) *RegisterFaceImageRequest {
	s.FacePersonId = &v
	return s
}

func (s *RegisterFaceImageRequest) SetImageUrl(v string) *RegisterFaceImageRequest {
	s.ImageUrl = &v
	return s
}

type RegisterFaceImageResponseBody struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FaceImages []*RegisterFaceImageResponseBodyFaceImages `json:"FaceImages,omitempty" xml:"FaceImages,omitempty" type:"Repeated"`
}

func (s RegisterFaceImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceImageResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterFaceImageResponseBody) SetRequestId(v string) *RegisterFaceImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterFaceImageResponseBody) SetFaceImages(v []*RegisterFaceImageResponseBodyFaceImages) *RegisterFaceImageResponseBody {
	s.FaceImages = v
	return s
}

type RegisterFaceImageResponseBodyFaceImages struct {
	FaceImageId *int64 `json:"FaceImageId,omitempty" xml:"FaceImageId,omitempty"`
}

func (s RegisterFaceImageResponseBodyFaceImages) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceImageResponseBodyFaceImages) GoString() string {
	return s.String()
}

func (s *RegisterFaceImageResponseBodyFaceImages) SetFaceImageId(v int64) *RegisterFaceImageResponseBodyFaceImages {
	s.FaceImageId = &v
	return s
}

type RegisterFaceImageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterFaceImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterFaceImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterFaceImageResponse) GoString() string {
	return s.String()
}

func (s *RegisterFaceImageResponse) SetHeaders(v map[string]*string) *RegisterFaceImageResponse {
	s.Headers = v
	return s
}

func (s *RegisterFaceImageResponse) SetBody(v *RegisterFaceImageResponseBody) *RegisterFaceImageResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetContent(v string) *UpdateTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateTemplateRequest) SetIsDefault(v bool) *UpdateTemplateRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateTemplateRequest) SetType(v int32) *UpdateTemplateRequest {
	s.Type = &v
	return s
}

type UpdateTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("multimediaai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateCoverTaskWithOptions(request *CreateCoverTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCoverTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCoverTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCoverTask"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCoverTask(request *CreateCoverTaskRequest) (_result *CreateCoverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCoverTaskResponse{}
	_body, _err := client.CreateCoverTaskWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFaceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFaceGroup"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateFacePersonWithOptions(request *CreateFacePersonRequest, runtime *util.RuntimeOptions) (_result *CreateFacePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFacePersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFacePerson"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFacePerson(request *CreateFacePersonRequest) (_result *CreateFacePersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFacePersonResponse{}
	_body, _err := client.CreateFacePersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGifTaskWithOptions(request *CreateGifTaskRequest, runtime *util.RuntimeOptions) (_result *CreateGifTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGifTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGifTask"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGifTask(request *CreateGifTaskRequest) (_result *CreateGifTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGifTaskResponse{}
	_body, _err := client.CreateGifTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLabelTaskWithOptions(request *CreateLabelTaskRequest, runtime *util.RuntimeOptions) (_result *CreateLabelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLabelTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLabelTask"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLabelTask(request *CreateLabelTaskRequest) (_result *CreateLabelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLabelTaskResponse{}
	_body, _err := client.CreateLabelTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTemplate"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceGroup"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteFaceImageWithOptions(request *DeleteFaceImageRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFaceImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaceImage"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceImage(request *DeleteFaceImageRequest) (_result *DeleteFaceImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceImageResponse{}
	_body, _err := client.DeleteFaceImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFacePersonWithOptions(request *DeleteFacePersonRequest, runtime *util.RuntimeOptions) (_result *DeleteFacePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFacePersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFacePerson"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFacePerson(request *DeleteFacePersonRequest) (_result *DeleteFacePersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFacePersonResponse{}
	_body, _err := client.DeleteFacePersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskResultWithOptions(request *GetTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskResult"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskResult(request *GetTaskResultRequest) (_result *GetTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResultResponse{}
	_body, _err := client.GetTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskStatus"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplate"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceGroupsWithOptions(request *ListFaceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListFaceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFaceGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaceGroups"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceGroups(request *ListFaceGroupsRequest) (_result *ListFaceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceGroupsResponse{}
	_body, _err := client.ListFaceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceImagesWithOptions(request *ListFaceImagesRequest, runtime *util.RuntimeOptions) (_result *ListFaceImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFaceImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaceImages"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceImages(request *ListFaceImagesRequest) (_result *ListFaceImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceImagesResponse{}
	_body, _err := client.ListFaceImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFacePersonsWithOptions(request *ListFacePersonsRequest, runtime *util.RuntimeOptions) (_result *ListFacePersonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFacePersonsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFacePersons"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFacePersons(request *ListFacePersonsRequest) (_result *ListFacePersonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFacePersonsResponse{}
	_body, _err := client.ListFacePersonsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplates"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessFaceAlgorithmWithOptions(request *ProcessFaceAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessFaceAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessFaceAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessFaceAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessFaceAlgorithm(request *ProcessFaceAlgorithmRequest) (_result *ProcessFaceAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessFaceAlgorithmResponse{}
	_body, _err := client.ProcessFaceAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessImageTagAlgorithmWithOptions(request *ProcessImageTagAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessImageTagAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessImageTagAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessImageTagAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessImageTagAlgorithm(request *ProcessImageTagAlgorithmRequest) (_result *ProcessImageTagAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessImageTagAlgorithmResponse{}
	_body, _err := client.ProcessImageTagAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessLandmarkAlgorithmWithOptions(request *ProcessLandmarkAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessLandmarkAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessLandmarkAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessLandmarkAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessLandmarkAlgorithm(request *ProcessLandmarkAlgorithmRequest) (_result *ProcessLandmarkAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessLandmarkAlgorithmResponse{}
	_body, _err := client.ProcessLandmarkAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessLogoAlgorithmWithOptions(request *ProcessLogoAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessLogoAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessLogoAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessLogoAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessLogoAlgorithm(request *ProcessLogoAlgorithmRequest) (_result *ProcessLogoAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessLogoAlgorithmResponse{}
	_body, _err := client.ProcessLogoAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessNewsAlgorithmWithOptions(request *ProcessNewsAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessNewsAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessNewsAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessNewsAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessNewsAlgorithm(request *ProcessNewsAlgorithmRequest) (_result *ProcessNewsAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessNewsAlgorithmResponse{}
	_body, _err := client.ProcessNewsAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessNlpAlgorithmWithOptions(request *ProcessNlpAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessNlpAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessNlpAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessNlpAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessNlpAlgorithm(request *ProcessNlpAlgorithmRequest) (_result *ProcessNlpAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessNlpAlgorithmResponse{}
	_body, _err := client.ProcessNlpAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessOcrAlgorithmWithOptions(request *ProcessOcrAlgorithmRequest, runtime *util.RuntimeOptions) (_result *ProcessOcrAlgorithmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProcessOcrAlgorithmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProcessOcrAlgorithm"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessOcrAlgorithm(request *ProcessOcrAlgorithmRequest) (_result *ProcessOcrAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProcessOcrAlgorithmResponse{}
	_body, _err := client.ProcessOcrAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterFaceImageWithOptions(request *RegisterFaceImageRequest, runtime *util.RuntimeOptions) (_result *RegisterFaceImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterFaceImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterFaceImage"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterFaceImage(request *RegisterFaceImageRequest) (_result *RegisterFaceImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterFaceImageResponse{}
	_body, _err := client.RegisterFaceImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(request *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTemplate"), tea.String("2019-08-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
