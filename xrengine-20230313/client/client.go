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

type AuthUserRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s AuthUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthUserRequest) GoString() string {
	return s.String()
}

func (s *AuthUserRequest) SetJwtToken(v string) *AuthUserRequest {
	s.JwtToken = &v
	return s
}

type AuthUserResponseBody struct {
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AuthUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                   `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                    `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthUserResponseBody) GoString() string {
	return s.String()
}

func (s *AuthUserResponseBody) SetCode(v string) *AuthUserResponseBody {
	s.Code = &v
	return s
}

func (s *AuthUserResponseBody) SetData(v *AuthUserResponseBodyData) *AuthUserResponseBody {
	s.Data = v
	return s
}

func (s *AuthUserResponseBody) SetErrorName(v string) *AuthUserResponseBody {
	s.ErrorName = &v
	return s
}

func (s *AuthUserResponseBody) SetHttpCode(v int32) *AuthUserResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AuthUserResponseBody) SetMessage(v string) *AuthUserResponseBody {
	s.Message = &v
	return s
}

func (s *AuthUserResponseBody) SetRequestId(v string) *AuthUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthUserResponseBody) SetSuccess(v bool) *AuthUserResponseBody {
	s.Success = &v
	return s
}

type AuthUserResponseBodyData struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AuthUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AuthUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *AuthUserResponseBodyData) SetJwtToken(v string) *AuthUserResponseBodyData {
	s.JwtToken = &v
	return s
}

func (s *AuthUserResponseBodyData) SetType(v string) *AuthUserResponseBodyData {
	s.Type = &v
	return s
}

type AuthUserResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthUserResponse) GoString() string {
	return s.String()
}

func (s *AuthUserResponse) SetHeaders(v map[string]*string) *AuthUserResponse {
	s.Headers = v
	return s
}

func (s *AuthUserResponse) SetStatusCode(v int32) *AuthUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthUserResponse) SetBody(v *AuthUserResponseBody) *AuthUserResponse {
	s.Body = v
	return s
}

type BatchQueryMotionShopTaskStatusRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchQueryMotionShopTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryMotionShopTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryMotionShopTaskStatusRequest) SetJwtToken(v string) *BatchQueryMotionShopTaskStatusRequest {
	s.JwtToken = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusRequest) SetTaskId(v string) *BatchQueryMotionShopTaskStatusRequest {
	s.TaskId = &v
	return s
}

type BatchQueryMotionShopTaskStatusResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchQueryMotionShopTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchQueryMotionShopTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryMotionShopTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryMotionShopTaskStatusResponseBody) SetCode(v string) *BatchQueryMotionShopTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBody) SetData(v *BatchQueryMotionShopTaskStatusResponseBodyData) *BatchQueryMotionShopTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBody) SetMessage(v string) *BatchQueryMotionShopTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBody) SetRequestId(v string) *BatchQueryMotionShopTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBody) SetSuccess(v bool) *BatchQueryMotionShopTaskStatusResponseBody {
	s.Success = &v
	return s
}

type BatchQueryMotionShopTaskStatusResponseBodyData struct {
	Tasks []*BatchQueryMotionShopTaskStatusResponseBodyDataTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s BatchQueryMotionShopTaskStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryMotionShopTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyData) SetTasks(v []*BatchQueryMotionShopTaskStatusResponseBodyDataTasks) *BatchQueryMotionShopTaskStatusResponseBodyData {
	s.Tasks = v
	return s
}

type BatchQueryMotionShopTaskStatusResponseBodyDataTasks struct {
	Result *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId *string                                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchQueryMotionShopTaskStatusResponseBodyDataTasks) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryMotionShopTaskStatusResponseBodyDataTasks) GoString() string {
	return s.String()
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyDataTasks) SetResult(v *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult) *BatchQueryMotionShopTaskStatusResponseBodyDataTasks {
	s.Result = v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyDataTasks) SetStatus(v string) *BatchQueryMotionShopTaskStatusResponseBodyDataTasks {
	s.Status = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyDataTasks) SetTaskId(v string) *BatchQueryMotionShopTaskStatusResponseBodyDataTasks {
	s.TaskId = &v
	return s
}

type BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	ShareUrl *string `json:"ShareUrl,omitempty" xml:"ShareUrl,omitempty"`
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult) GoString() string {
	return s.String()
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult) SetCoverUrl(v string) *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult {
	s.CoverUrl = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult) SetShareUrl(v string) *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult {
	s.ShareUrl = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult) SetVideoUrl(v string) *BatchQueryMotionShopTaskStatusResponseBodyDataTasksResult {
	s.VideoUrl = &v
	return s
}

type BatchQueryMotionShopTaskStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchQueryMotionShopTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQueryMotionShopTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryMotionShopTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryMotionShopTaskStatusResponse) SetHeaders(v map[string]*string) *BatchQueryMotionShopTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponse) SetStatusCode(v int32) *BatchQueryMotionShopTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryMotionShopTaskStatusResponse) SetBody(v *BatchQueryMotionShopTaskStatusResponseBody) *BatchQueryMotionShopTaskStatusResponse {
	s.Body = v
	return s
}

type CreateAvatarTalkProjectRequest struct {
	AvatarProjectId *string `json:"AvatarProjectId,omitempty" xml:"AvatarProjectId,omitempty"`
	JwtToken        *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	TtsVoice        *string `json:"TtsVoice,omitempty" xml:"TtsVoice,omitempty"`
	TxtContent      *string `json:"TxtContent,omitempty" xml:"TxtContent,omitempty"`
}

func (s CreateAvatarTalkProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAvatarTalkProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateAvatarTalkProjectRequest) SetAvatarProjectId(v string) *CreateAvatarTalkProjectRequest {
	s.AvatarProjectId = &v
	return s
}

func (s *CreateAvatarTalkProjectRequest) SetJwtToken(v string) *CreateAvatarTalkProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateAvatarTalkProjectRequest) SetTitle(v string) *CreateAvatarTalkProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateAvatarTalkProjectRequest) SetTtsVoice(v string) *CreateAvatarTalkProjectRequest {
	s.TtsVoice = &v
	return s
}

func (s *CreateAvatarTalkProjectRequest) SetTxtContent(v string) *CreateAvatarTalkProjectRequest {
	s.TxtContent = &v
	return s
}

type CreateAvatarTalkProjectResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAvatarTalkProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAvatarTalkProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAvatarTalkProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAvatarTalkProjectResponseBody) SetCode(v string) *CreateAvatarTalkProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAvatarTalkProjectResponseBody) SetData(v *CreateAvatarTalkProjectResponseBodyData) *CreateAvatarTalkProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateAvatarTalkProjectResponseBody) SetMessage(v string) *CreateAvatarTalkProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAvatarTalkProjectResponseBody) SetRequestId(v string) *CreateAvatarTalkProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAvatarTalkProjectResponseBody) SetSuccess(v bool) *CreateAvatarTalkProjectResponseBody {
	s.Success = &v
	return s
}

type CreateAvatarTalkProjectResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateAvatarTalkProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAvatarTalkProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAvatarTalkProjectResponseBodyData) SetId(v string) *CreateAvatarTalkProjectResponseBodyData {
	s.Id = &v
	return s
}

type CreateAvatarTalkProjectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAvatarTalkProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAvatarTalkProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAvatarTalkProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateAvatarTalkProjectResponse) SetHeaders(v map[string]*string) *CreateAvatarTalkProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateAvatarTalkProjectResponse) SetStatusCode(v int32) *CreateAvatarTalkProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAvatarTalkProjectResponse) SetBody(v *CreateAvatarTalkProjectResponseBody) *CreateAvatarTalkProjectResponse {
	s.Body = v
	return s
}

type CreateDigitalHumanProjectRequest struct {
	AudioId           *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	AudioUrl          *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	BackgroundId      *string `json:"BackgroundId,omitempty" xml:"BackgroundId,omitempty"`
	BackgroundUrl     *string `json:"BackgroundUrl,omitempty" xml:"BackgroundUrl,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ForegroundId      *string `json:"ForegroundId,omitempty" xml:"ForegroundId,omitempty"`
	ForegroundUrl     *string `json:"ForegroundUrl,omitempty" xml:"ForegroundUrl,omitempty"`
	HumanLayerStyle   *string `json:"HumanLayerStyle,omitempty" xml:"HumanLayerStyle,omitempty"`
	Intro             *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken          *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ModelId           *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	OutputConfig      *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	TtsVoiceId        *string `json:"TtsVoiceId,omitempty" xml:"TtsVoiceId,omitempty"`
	WatermarkImageUrl *string `json:"WatermarkImageUrl,omitempty" xml:"WatermarkImageUrl,omitempty"`
	WatermarkStyle    *string `json:"WatermarkStyle,omitempty" xml:"WatermarkStyle,omitempty"`
}

func (s CreateDigitalHumanProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDigitalHumanProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalHumanProjectRequest) SetAudioId(v string) *CreateDigitalHumanProjectRequest {
	s.AudioId = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetAudioUrl(v string) *CreateDigitalHumanProjectRequest {
	s.AudioUrl = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetBackgroundId(v string) *CreateDigitalHumanProjectRequest {
	s.BackgroundId = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetBackgroundUrl(v string) *CreateDigitalHumanProjectRequest {
	s.BackgroundUrl = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetContent(v string) *CreateDigitalHumanProjectRequest {
	s.Content = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetForegroundId(v string) *CreateDigitalHumanProjectRequest {
	s.ForegroundId = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetForegroundUrl(v string) *CreateDigitalHumanProjectRequest {
	s.ForegroundUrl = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetHumanLayerStyle(v string) *CreateDigitalHumanProjectRequest {
	s.HumanLayerStyle = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetIntro(v string) *CreateDigitalHumanProjectRequest {
	s.Intro = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetJwtToken(v string) *CreateDigitalHumanProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetMode(v string) *CreateDigitalHumanProjectRequest {
	s.Mode = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetModelId(v string) *CreateDigitalHumanProjectRequest {
	s.ModelId = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetOutputConfig(v string) *CreateDigitalHumanProjectRequest {
	s.OutputConfig = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetTitle(v string) *CreateDigitalHumanProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetTtsVoiceId(v string) *CreateDigitalHumanProjectRequest {
	s.TtsVoiceId = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetWatermarkImageUrl(v string) *CreateDigitalHumanProjectRequest {
	s.WatermarkImageUrl = &v
	return s
}

func (s *CreateDigitalHumanProjectRequest) SetWatermarkStyle(v string) *CreateDigitalHumanProjectRequest {
	s.WatermarkStyle = &v
	return s
}

type CreateDigitalHumanProjectResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateDigitalHumanProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDigitalHumanProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDigitalHumanProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDigitalHumanProjectResponseBody) SetCode(v string) *CreateDigitalHumanProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDigitalHumanProjectResponseBody) SetData(v *CreateDigitalHumanProjectResponseBodyData) *CreateDigitalHumanProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateDigitalHumanProjectResponseBody) SetMessage(v string) *CreateDigitalHumanProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDigitalHumanProjectResponseBody) SetRequestId(v string) *CreateDigitalHumanProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDigitalHumanProjectResponseBody) SetSuccess(v bool) *CreateDigitalHumanProjectResponseBody {
	s.Success = &v
	return s
}

type CreateDigitalHumanProjectResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateDigitalHumanProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDigitalHumanProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDigitalHumanProjectResponseBodyData) SetId(v string) *CreateDigitalHumanProjectResponseBodyData {
	s.Id = &v
	return s
}

type CreateDigitalHumanProjectResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDigitalHumanProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDigitalHumanProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDigitalHumanProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateDigitalHumanProjectResponse) SetHeaders(v map[string]*string) *CreateDigitalHumanProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateDigitalHumanProjectResponse) SetStatusCode(v int32) *CreateDigitalHumanProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDigitalHumanProjectResponse) SetBody(v *CreateDigitalHumanProjectResponseBody) *CreateDigitalHumanProjectResponse {
	s.Body = v
	return s
}

type CreateLivePortraitProjectRequest struct {
	AudioId           *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	AudioUrl          *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ImageId           *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUrl          *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Intro             *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken          *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	LightModel        *bool   `json:"LightModel,omitempty" xml:"LightModel,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OutputConfig      *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	TtsVoiceId        *string `json:"TtsVoiceId,omitempty" xml:"TtsVoiceId,omitempty"`
	WatermarkImageUrl *string `json:"WatermarkImageUrl,omitempty" xml:"WatermarkImageUrl,omitempty"`
	WatermarkStyle    *string `json:"WatermarkStyle,omitempty" xml:"WatermarkStyle,omitempty"`
}

func (s CreateLivePortraitProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLivePortraitProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePortraitProjectRequest) SetAudioId(v string) *CreateLivePortraitProjectRequest {
	s.AudioId = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetAudioUrl(v string) *CreateLivePortraitProjectRequest {
	s.AudioUrl = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetContent(v string) *CreateLivePortraitProjectRequest {
	s.Content = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetImageId(v string) *CreateLivePortraitProjectRequest {
	s.ImageId = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetImageUrl(v string) *CreateLivePortraitProjectRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetIntro(v string) *CreateLivePortraitProjectRequest {
	s.Intro = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetJwtToken(v string) *CreateLivePortraitProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetLightModel(v bool) *CreateLivePortraitProjectRequest {
	s.LightModel = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetMode(v string) *CreateLivePortraitProjectRequest {
	s.Mode = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetOutputConfig(v string) *CreateLivePortraitProjectRequest {
	s.OutputConfig = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetTitle(v string) *CreateLivePortraitProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetTtsVoiceId(v string) *CreateLivePortraitProjectRequest {
	s.TtsVoiceId = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetWatermarkImageUrl(v string) *CreateLivePortraitProjectRequest {
	s.WatermarkImageUrl = &v
	return s
}

func (s *CreateLivePortraitProjectRequest) SetWatermarkStyle(v string) *CreateLivePortraitProjectRequest {
	s.WatermarkStyle = &v
	return s
}

type CreateLivePortraitProjectResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateLivePortraitProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLivePortraitProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLivePortraitProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivePortraitProjectResponseBody) SetCode(v string) *CreateLivePortraitProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLivePortraitProjectResponseBody) SetData(v *CreateLivePortraitProjectResponseBodyData) *CreateLivePortraitProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateLivePortraitProjectResponseBody) SetMessage(v string) *CreateLivePortraitProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLivePortraitProjectResponseBody) SetRequestId(v string) *CreateLivePortraitProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivePortraitProjectResponseBody) SetSuccess(v bool) *CreateLivePortraitProjectResponseBody {
	s.Success = &v
	return s
}

type CreateLivePortraitProjectResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateLivePortraitProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLivePortraitProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLivePortraitProjectResponseBodyData) SetId(v string) *CreateLivePortraitProjectResponseBodyData {
	s.Id = &v
	return s
}

type CreateLivePortraitProjectResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLivePortraitProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLivePortraitProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLivePortraitProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateLivePortraitProjectResponse) SetHeaders(v map[string]*string) *CreateLivePortraitProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateLivePortraitProjectResponse) SetStatusCode(v int32) *CreateLivePortraitProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivePortraitProjectResponse) SetBody(v *CreateLivePortraitProjectResponseBody) *CreateLivePortraitProjectResponse {
	s.Body = v
	return s
}

type GenerateMotionShopVideoUploadUrlRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GenerateMotionShopVideoUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateMotionShopVideoUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateMotionShopVideoUploadUrlRequest) SetJwtToken(v string) *GenerateMotionShopVideoUploadUrlRequest {
	s.JwtToken = &v
	return s
}

type GenerateMotionShopVideoUploadUrlResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GenerateMotionShopVideoUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateMotionShopVideoUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateMotionShopVideoUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMotionShopVideoUploadUrlResponseBody) SetCode(v string) *GenerateMotionShopVideoUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponseBody) SetData(v *GenerateMotionShopVideoUploadUrlResponseBodyData) *GenerateMotionShopVideoUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponseBody) SetMessage(v string) *GenerateMotionShopVideoUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponseBody) SetRequestId(v string) *GenerateMotionShopVideoUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponseBody) SetSuccess(v bool) *GenerateMotionShopVideoUploadUrlResponseBody {
	s.Success = &v
	return s
}

type GenerateMotionShopVideoUploadUrlResponseBodyData struct {
	OssKey    *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s GenerateMotionShopVideoUploadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateMotionShopVideoUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateMotionShopVideoUploadUrlResponseBodyData) SetOssKey(v string) *GenerateMotionShopVideoUploadUrlResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponseBodyData) SetUploadUrl(v string) *GenerateMotionShopVideoUploadUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

type GenerateMotionShopVideoUploadUrlResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateMotionShopVideoUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateMotionShopVideoUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateMotionShopVideoUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateMotionShopVideoUploadUrlResponse) SetHeaders(v map[string]*string) *GenerateMotionShopVideoUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponse) SetStatusCode(v int32) *GenerateMotionShopVideoUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateMotionShopVideoUploadUrlResponse) SetBody(v *GenerateMotionShopVideoUploadUrlResponseBody) *GenerateMotionShopVideoUploadUrlResponse {
	s.Body = v
	return s
}

type GetMapDataRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetMapDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMapDataRequest) GoString() string {
	return s.String()
}

func (s *GetMapDataRequest) SetAppId(v string) *GetMapDataRequest {
	s.AppId = &v
	return s
}

func (s *GetMapDataRequest) SetJwtToken(v string) *GetMapDataRequest {
	s.JwtToken = &v
	return s
}

type GetMapDataResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                  `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMapDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMapDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMapDataResponseBody) SetCode(v string) *GetMapDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetMapDataResponseBody) SetData(v []map[string]interface{}) *GetMapDataResponseBody {
	s.Data = v
	return s
}

func (s *GetMapDataResponseBody) SetErrorName(v string) *GetMapDataResponseBody {
	s.ErrorName = &v
	return s
}

func (s *GetMapDataResponseBody) SetHttpCode(v int32) *GetMapDataResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetMapDataResponseBody) SetMessage(v string) *GetMapDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetMapDataResponseBody) SetRequestId(v string) *GetMapDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMapDataResponseBody) SetSuccess(v bool) *GetMapDataResponseBody {
	s.Success = &v
	return s
}

type GetMapDataResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMapDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMapDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMapDataResponse) GoString() string {
	return s.String()
}

func (s *GetMapDataResponse) SetHeaders(v map[string]*string) *GetMapDataResponse {
	s.Headers = v
	return s
}

func (s *GetMapDataResponse) SetStatusCode(v int32) *GetMapDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMapDataResponse) SetBody(v *GetMapDataResponseBody) *GetMapDataResponse {
	s.Body = v
	return s
}

type GetMapPublishDataRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetMapPublishDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMapPublishDataRequest) GoString() string {
	return s.String()
}

func (s *GetMapPublishDataRequest) SetAppId(v string) *GetMapPublishDataRequest {
	s.AppId = &v
	return s
}

func (s *GetMapPublishDataRequest) SetJwtToken(v string) *GetMapPublishDataRequest {
	s.JwtToken = &v
	return s
}

type GetMapPublishDataResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                  `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMapPublishDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMapPublishDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMapPublishDataResponseBody) SetCode(v string) *GetMapPublishDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetMapPublishDataResponseBody) SetData(v []map[string]interface{}) *GetMapPublishDataResponseBody {
	s.Data = v
	return s
}

func (s *GetMapPublishDataResponseBody) SetErrorName(v string) *GetMapPublishDataResponseBody {
	s.ErrorName = &v
	return s
}

func (s *GetMapPublishDataResponseBody) SetHttpCode(v int32) *GetMapPublishDataResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetMapPublishDataResponseBody) SetMessage(v string) *GetMapPublishDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetMapPublishDataResponseBody) SetRequestId(v string) *GetMapPublishDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMapPublishDataResponseBody) SetSuccess(v bool) *GetMapPublishDataResponseBody {
	s.Success = &v
	return s
}

type GetMapPublishDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMapPublishDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMapPublishDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMapPublishDataResponse) GoString() string {
	return s.String()
}

func (s *GetMapPublishDataResponse) SetHeaders(v map[string]*string) *GetMapPublishDataResponse {
	s.Headers = v
	return s
}

func (s *GetMapPublishDataResponse) SetStatusCode(v int32) *GetMapPublishDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMapPublishDataResponse) SetBody(v *GetMapPublishDataResponseBody) *GetMapPublishDataResponse {
	s.Body = v
	return s
}

type InitLocateRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Params   *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s InitLocateRequest) String() string {
	return tea.Prettify(s)
}

func (s InitLocateRequest) GoString() string {
	return s.String()
}

func (s *InitLocateRequest) SetJwtToken(v string) *InitLocateRequest {
	s.JwtToken = &v
	return s
}

func (s *InitLocateRequest) SetParams(v string) *InitLocateRequest {
	s.Params = &v
	return s
}

type InitLocateResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitLocateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitLocateResponseBody) GoString() string {
	return s.String()
}

func (s *InitLocateResponseBody) SetCode(v string) *InitLocateResponseBody {
	s.Code = &v
	return s
}

func (s *InitLocateResponseBody) SetData(v string) *InitLocateResponseBody {
	s.Data = &v
	return s
}

func (s *InitLocateResponseBody) SetErrorName(v string) *InitLocateResponseBody {
	s.ErrorName = &v
	return s
}

func (s *InitLocateResponseBody) SetHttpCode(v int32) *InitLocateResponseBody {
	s.HttpCode = &v
	return s
}

func (s *InitLocateResponseBody) SetMessage(v string) *InitLocateResponseBody {
	s.Message = &v
	return s
}

func (s *InitLocateResponseBody) SetRequestId(v string) *InitLocateResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitLocateResponseBody) SetSuccess(v bool) *InitLocateResponseBody {
	s.Success = &v
	return s
}

type InitLocateResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitLocateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitLocateResponse) String() string {
	return tea.Prettify(s)
}

func (s InitLocateResponse) GoString() string {
	return s.String()
}

func (s *InitLocateResponse) SetHeaders(v map[string]*string) *InitLocateResponse {
	s.Headers = v
	return s
}

func (s *InitLocateResponse) SetStatusCode(v int32) *InitLocateResponse {
	s.StatusCode = &v
	return s
}

func (s *InitLocateResponse) SetBody(v *InitLocateResponseBody) *InitLocateResponse {
	s.Body = v
	return s
}

type ListCommonMaterialsRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCommonMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonMaterialsRequest) GoString() string {
	return s.String()
}

func (s *ListCommonMaterialsRequest) SetJwtToken(v string) *ListCommonMaterialsRequest {
	s.JwtToken = &v
	return s
}

func (s *ListCommonMaterialsRequest) SetType(v string) *ListCommonMaterialsRequest {
	s.Type = &v
	return s
}

type ListCommonMaterialsResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListCommonMaterialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCommonMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommonMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonMaterialsResponseBody) SetCode(v string) *ListCommonMaterialsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCommonMaterialsResponseBody) SetData(v []*ListCommonMaterialsResponseBodyData) *ListCommonMaterialsResponseBody {
	s.Data = v
	return s
}

func (s *ListCommonMaterialsResponseBody) SetMessage(v string) *ListCommonMaterialsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCommonMaterialsResponseBody) SetRequestId(v string) *ListCommonMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonMaterialsResponseBody) SetSuccess(v bool) *ListCommonMaterialsResponseBody {
	s.Success = &v
	return s
}

type ListCommonMaterialsResponseBodyData struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Ext      *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCommonMaterialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCommonMaterialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCommonMaterialsResponseBodyData) SetCoverUrl(v string) *ListCommonMaterialsResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *ListCommonMaterialsResponseBodyData) SetExt(v string) *ListCommonMaterialsResponseBodyData {
	s.Ext = &v
	return s
}

func (s *ListCommonMaterialsResponseBodyData) SetFileUrl(v string) *ListCommonMaterialsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *ListCommonMaterialsResponseBodyData) SetId(v string) *ListCommonMaterialsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCommonMaterialsResponseBodyData) SetName(v string) *ListCommonMaterialsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListCommonMaterialsResponseBodyData) SetOssKey(v string) *ListCommonMaterialsResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *ListCommonMaterialsResponseBodyData) SetType(v string) *ListCommonMaterialsResponseBodyData {
	s.Type = &v
	return s
}

type ListCommonMaterialsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommonMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommonMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommonMaterialsResponse) GoString() string {
	return s.String()
}

func (s *ListCommonMaterialsResponse) SetHeaders(v map[string]*string) *ListCommonMaterialsResponse {
	s.Headers = v
	return s
}

func (s *ListCommonMaterialsResponse) SetStatusCode(v int32) *ListCommonMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonMaterialsResponse) SetBody(v *ListCommonMaterialsResponseBody) *ListCommonMaterialsResponse {
	s.Body = v
	return s
}

type ListDigitalHumanMaterialsRequest struct {
	Current               *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken              *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	OnlyPersonalMaterials *bool   `json:"OnlyPersonalMaterials,omitempty" xml:"OnlyPersonalMaterials,omitempty"`
	Size                  *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Types                 *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListDigitalHumanMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDigitalHumanMaterialsRequest) GoString() string {
	return s.String()
}

func (s *ListDigitalHumanMaterialsRequest) SetCurrent(v int32) *ListDigitalHumanMaterialsRequest {
	s.Current = &v
	return s
}

func (s *ListDigitalHumanMaterialsRequest) SetJwtToken(v string) *ListDigitalHumanMaterialsRequest {
	s.JwtToken = &v
	return s
}

func (s *ListDigitalHumanMaterialsRequest) SetOnlyPersonalMaterials(v bool) *ListDigitalHumanMaterialsRequest {
	s.OnlyPersonalMaterials = &v
	return s
}

func (s *ListDigitalHumanMaterialsRequest) SetSize(v int32) *ListDigitalHumanMaterialsRequest {
	s.Size = &v
	return s
}

func (s *ListDigitalHumanMaterialsRequest) SetTypes(v string) *ListDigitalHumanMaterialsRequest {
	s.Types = &v
	return s
}

type ListDigitalHumanMaterialsResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                       `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*ListDigitalHumanMaterialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                       `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                       `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDigitalHumanMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDigitalHumanMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDigitalHumanMaterialsResponseBody) SetCode(v string) *ListDigitalHumanMaterialsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetCurrent(v int32) *ListDigitalHumanMaterialsResponseBody {
	s.Current = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetData(v []*ListDigitalHumanMaterialsResponseBodyData) *ListDigitalHumanMaterialsResponseBody {
	s.Data = v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetMessage(v string) *ListDigitalHumanMaterialsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetPages(v int32) *ListDigitalHumanMaterialsResponseBody {
	s.Pages = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetRequestId(v string) *ListDigitalHumanMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetSize(v int32) *ListDigitalHumanMaterialsResponseBody {
	s.Size = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetSuccess(v bool) *ListDigitalHumanMaterialsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBody) SetTotal(v int32) *ListDigitalHumanMaterialsResponseBody {
	s.Total = &v
	return s
}

type ListDigitalHumanMaterialsResponseBodyData struct {
	Components []*ListDigitalHumanMaterialsResponseBodyDataComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	Ext        *string                                                `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl    *string                                                `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Files      map[string]interface{}                                 `json:"Files,omitempty" xml:"Files,omitempty"`
	Id         *string                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDigitalHumanMaterialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDigitalHumanMaterialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetComponents(v []*ListDigitalHumanMaterialsResponseBodyDataComponents) *ListDigitalHumanMaterialsResponseBodyData {
	s.Components = v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetExt(v string) *ListDigitalHumanMaterialsResponseBodyData {
	s.Ext = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetFileUrl(v string) *ListDigitalHumanMaterialsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetFiles(v map[string]interface{}) *ListDigitalHumanMaterialsResponseBodyData {
	s.Files = v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetId(v string) *ListDigitalHumanMaterialsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetName(v string) *ListDigitalHumanMaterialsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyData) SetType(v string) *ListDigitalHumanMaterialsResponseBodyData {
	s.Type = &v
	return s
}

type ListDigitalHumanMaterialsResponseBodyDataComponents struct {
	Ext     *string                `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl *string                `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Files   map[string]interface{} `json:"Files,omitempty" xml:"Files,omitempty"`
	Id      *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Type    *string                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDigitalHumanMaterialsResponseBodyDataComponents) String() string {
	return tea.Prettify(s)
}

func (s ListDigitalHumanMaterialsResponseBodyDataComponents) GoString() string {
	return s.String()
}

func (s *ListDigitalHumanMaterialsResponseBodyDataComponents) SetExt(v string) *ListDigitalHumanMaterialsResponseBodyDataComponents {
	s.Ext = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyDataComponents) SetFileUrl(v string) *ListDigitalHumanMaterialsResponseBodyDataComponents {
	s.FileUrl = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyDataComponents) SetFiles(v map[string]interface{}) *ListDigitalHumanMaterialsResponseBodyDataComponents {
	s.Files = v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyDataComponents) SetId(v string) *ListDigitalHumanMaterialsResponseBodyDataComponents {
	s.Id = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyDataComponents) SetName(v string) *ListDigitalHumanMaterialsResponseBodyDataComponents {
	s.Name = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponseBodyDataComponents) SetType(v string) *ListDigitalHumanMaterialsResponseBodyDataComponents {
	s.Type = &v
	return s
}

type ListDigitalHumanMaterialsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDigitalHumanMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDigitalHumanMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDigitalHumanMaterialsResponse) GoString() string {
	return s.String()
}

func (s *ListDigitalHumanMaterialsResponse) SetHeaders(v map[string]*string) *ListDigitalHumanMaterialsResponse {
	s.Headers = v
	return s
}

func (s *ListDigitalHumanMaterialsResponse) SetStatusCode(v int32) *ListDigitalHumanMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDigitalHumanMaterialsResponse) SetBody(v *ListDigitalHumanMaterialsResponseBody) *ListDigitalHumanMaterialsResponse {
	s.Body = v
	return s
}

type ListLocationServiceRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Sort      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListLocationServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *ListLocationServiceRequest) SetCurrent(v int32) *ListLocationServiceRequest {
	s.Current = &v
	return s
}

func (s *ListLocationServiceRequest) SetJwtToken(v string) *ListLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *ListLocationServiceRequest) SetSize(v int32) *ListLocationServiceRequest {
	s.Size = &v
	return s
}

func (s *ListLocationServiceRequest) SetSort(v string) *ListLocationServiceRequest {
	s.Sort = &v
	return s
}

func (s *ListLocationServiceRequest) SetSortField(v string) *ListLocationServiceRequest {
	s.SortField = &v
	return s
}

type ListLocationServiceResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                 `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*ListLocationServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                                `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                 `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLocationServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLocationServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListLocationServiceResponseBody) SetCode(v string) *ListLocationServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetCurrent(v int32) *ListLocationServiceResponseBody {
	s.Current = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetData(v []*ListLocationServiceResponseBodyData) *ListLocationServiceResponseBody {
	s.Data = v
	return s
}

func (s *ListLocationServiceResponseBody) SetErrorName(v string) *ListLocationServiceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetHttpCode(v int32) *ListLocationServiceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetMessage(v string) *ListLocationServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetPages(v int32) *ListLocationServiceResponseBody {
	s.Pages = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetRequestId(v string) *ListLocationServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetSize(v int32) *ListLocationServiceResponseBody {
	s.Size = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetSuccess(v bool) *ListLocationServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ListLocationServiceResponseBody) SetTotal(v int32) *ListLocationServiceResponseBody {
	s.Total = &v
	return s
}

type ListLocationServiceResponseBodyData struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Qps         *int64  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SvcState    *string `json:"SvcState,omitempty" xml:"SvcState,omitempty"`
}

func (s ListLocationServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLocationServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLocationServiceResponseBodyData) SetAppId(v string) *ListLocationServiceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetExpireTime(v string) *ListLocationServiceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetGmtCreate(v string) *ListLocationServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetGmtModified(v string) *ListLocationServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetId(v int64) *ListLocationServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetName(v string) *ListLocationServiceResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetNote(v string) *ListLocationServiceResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetQps(v int64) *ListLocationServiceResponseBodyData {
	s.Qps = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetStartTime(v string) *ListLocationServiceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetStatus(v string) *ListLocationServiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListLocationServiceResponseBodyData) SetSvcState(v string) *ListLocationServiceResponseBodyData {
	s.SvcState = &v
	return s
}

type ListLocationServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLocationServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *ListLocationServiceResponse) SetHeaders(v map[string]*string) *ListLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *ListLocationServiceResponse) SetStatusCode(v int32) *ListLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLocationServiceResponse) SetBody(v *ListLocationServiceResponseBody) *ListLocationServiceResponse {
	s.Body = v
	return s
}

type ListMotionShopTasksRequest struct {
	Current  *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListMotionShopTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMotionShopTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMotionShopTasksRequest) SetCurrent(v int32) *ListMotionShopTasksRequest {
	s.Current = &v
	return s
}

func (s *ListMotionShopTasksRequest) SetJwtToken(v string) *ListMotionShopTasksRequest {
	s.JwtToken = &v
	return s
}

func (s *ListMotionShopTasksRequest) SetSize(v int32) *ListMotionShopTasksRequest {
	s.Size = &v
	return s
}

type ListMotionShopTasksResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                 `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*ListMotionShopTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                 `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMotionShopTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMotionShopTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMotionShopTasksResponseBody) SetCode(v string) *ListMotionShopTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetCurrent(v int32) *ListMotionShopTasksResponseBody {
	s.Current = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetData(v []*ListMotionShopTasksResponseBodyData) *ListMotionShopTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetMessage(v string) *ListMotionShopTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetPages(v int32) *ListMotionShopTasksResponseBody {
	s.Pages = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetRequestId(v string) *ListMotionShopTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetSize(v int32) *ListMotionShopTasksResponseBody {
	s.Size = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetSuccess(v bool) *ListMotionShopTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListMotionShopTasksResponseBody) SetTotal(v int32) *ListMotionShopTasksResponseBody {
	s.Total = &v
	return s
}

type ListMotionShopTasksResponseBodyData struct {
	Material *ListMotionShopTasksResponseBodyDataMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Struct"`
	Result   *ListMotionShopTasksResponseBodyDataResult   `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status   *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId   *string                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListMotionShopTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMotionShopTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMotionShopTasksResponseBodyData) SetMaterial(v *ListMotionShopTasksResponseBodyDataMaterial) *ListMotionShopTasksResponseBodyData {
	s.Material = v
	return s
}

func (s *ListMotionShopTasksResponseBodyData) SetResult(v *ListMotionShopTasksResponseBodyDataResult) *ListMotionShopTasksResponseBodyData {
	s.Result = v
	return s
}

func (s *ListMotionShopTasksResponseBodyData) SetStatus(v string) *ListMotionShopTasksResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListMotionShopTasksResponseBodyData) SetTaskId(v string) *ListMotionShopTasksResponseBodyData {
	s.TaskId = &v
	return s
}

type ListMotionShopTasksResponseBodyDataMaterial struct {
	AvatarId *string    `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	Box      []*float64 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
	CoverUrl *string    `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
}

func (s ListMotionShopTasksResponseBodyDataMaterial) String() string {
	return tea.Prettify(s)
}

func (s ListMotionShopTasksResponseBodyDataMaterial) GoString() string {
	return s.String()
}

func (s *ListMotionShopTasksResponseBodyDataMaterial) SetAvatarId(v string) *ListMotionShopTasksResponseBodyDataMaterial {
	s.AvatarId = &v
	return s
}

func (s *ListMotionShopTasksResponseBodyDataMaterial) SetBox(v []*float64) *ListMotionShopTasksResponseBodyDataMaterial {
	s.Box = v
	return s
}

func (s *ListMotionShopTasksResponseBodyDataMaterial) SetCoverUrl(v string) *ListMotionShopTasksResponseBodyDataMaterial {
	s.CoverUrl = &v
	return s
}

type ListMotionShopTasksResponseBodyDataResult struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	ShareUrl *string `json:"ShareUrl,omitempty" xml:"ShareUrl,omitempty"`
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ListMotionShopTasksResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListMotionShopTasksResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListMotionShopTasksResponseBodyDataResult) SetCoverUrl(v string) *ListMotionShopTasksResponseBodyDataResult {
	s.CoverUrl = &v
	return s
}

func (s *ListMotionShopTasksResponseBodyDataResult) SetShareUrl(v string) *ListMotionShopTasksResponseBodyDataResult {
	s.ShareUrl = &v
	return s
}

func (s *ListMotionShopTasksResponseBodyDataResult) SetVideoUrl(v string) *ListMotionShopTasksResponseBodyDataResult {
	s.VideoUrl = &v
	return s
}

type ListMotionShopTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMotionShopTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMotionShopTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMotionShopTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMotionShopTasksResponse) SetHeaders(v map[string]*string) *ListMotionShopTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMotionShopTasksResponse) SetStatusCode(v int32) *ListMotionShopTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMotionShopTasksResponse) SetBody(v *ListMotionShopTasksResponseBody) *ListMotionShopTasksResponse {
	s.Body = v
	return s
}

type LivePortraitFaceDetectRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s LivePortraitFaceDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s LivePortraitFaceDetectRequest) GoString() string {
	return s.String()
}

func (s *LivePortraitFaceDetectRequest) SetImageUrl(v string) *LivePortraitFaceDetectRequest {
	s.ImageUrl = &v
	return s
}

func (s *LivePortraitFaceDetectRequest) SetJwtToken(v string) *LivePortraitFaceDetectRequest {
	s.JwtToken = &v
	return s
}

type LivePortraitFaceDetectResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LivePortraitFaceDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LivePortraitFaceDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LivePortraitFaceDetectResponseBody) GoString() string {
	return s.String()
}

func (s *LivePortraitFaceDetectResponseBody) SetCode(v string) *LivePortraitFaceDetectResponseBody {
	s.Code = &v
	return s
}

func (s *LivePortraitFaceDetectResponseBody) SetData(v *LivePortraitFaceDetectResponseBodyData) *LivePortraitFaceDetectResponseBody {
	s.Data = v
	return s
}

func (s *LivePortraitFaceDetectResponseBody) SetMessage(v string) *LivePortraitFaceDetectResponseBody {
	s.Message = &v
	return s
}

func (s *LivePortraitFaceDetectResponseBody) SetRequestId(v string) *LivePortraitFaceDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *LivePortraitFaceDetectResponseBody) SetSuccess(v bool) *LivePortraitFaceDetectResponseBody {
	s.Success = &v
	return s
}

type LivePortraitFaceDetectResponseBodyData struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s LivePortraitFaceDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LivePortraitFaceDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *LivePortraitFaceDetectResponseBodyData) SetCode(v int32) *LivePortraitFaceDetectResponseBodyData {
	s.Code = &v
	return s
}

func (s *LivePortraitFaceDetectResponseBodyData) SetMessage(v string) *LivePortraitFaceDetectResponseBodyData {
	s.Message = &v
	return s
}

type LivePortraitFaceDetectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LivePortraitFaceDetectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LivePortraitFaceDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s LivePortraitFaceDetectResponse) GoString() string {
	return s.String()
}

func (s *LivePortraitFaceDetectResponse) SetHeaders(v map[string]*string) *LivePortraitFaceDetectResponse {
	s.Headers = v
	return s
}

func (s *LivePortraitFaceDetectResponse) SetStatusCode(v int32) *LivePortraitFaceDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *LivePortraitFaceDetectResponse) SetBody(v *LivePortraitFaceDetectResponseBody) *LivePortraitFaceDetectResponse {
	s.Body = v
	return s
}

type LocateRequest struct {
	Image    *string `json:"Image,omitempty" xml:"Image,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Params   *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s LocateRequest) String() string {
	return tea.Prettify(s)
}

func (s LocateRequest) GoString() string {
	return s.String()
}

func (s *LocateRequest) SetImage(v string) *LocateRequest {
	s.Image = &v
	return s
}

func (s *LocateRequest) SetJwtToken(v string) *LocateRequest {
	s.JwtToken = &v
	return s
}

func (s *LocateRequest) SetParams(v string) *LocateRequest {
	s.Params = &v
	return s
}

type LocateResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LocateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LocateResponseBody) GoString() string {
	return s.String()
}

func (s *LocateResponseBody) SetCode(v string) *LocateResponseBody {
	s.Code = &v
	return s
}

func (s *LocateResponseBody) SetData(v string) *LocateResponseBody {
	s.Data = &v
	return s
}

func (s *LocateResponseBody) SetErrorName(v string) *LocateResponseBody {
	s.ErrorName = &v
	return s
}

func (s *LocateResponseBody) SetHttpCode(v int32) *LocateResponseBody {
	s.HttpCode = &v
	return s
}

func (s *LocateResponseBody) SetMessage(v string) *LocateResponseBody {
	s.Message = &v
	return s
}

func (s *LocateResponseBody) SetRequestId(v string) *LocateResponseBody {
	s.RequestId = &v
	return s
}

func (s *LocateResponseBody) SetSuccess(v bool) *LocateResponseBody {
	s.Success = &v
	return s
}

type LocateResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LocateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LocateResponse) String() string {
	return tea.Prettify(s)
}

func (s LocateResponse) GoString() string {
	return s.String()
}

func (s *LocateResponse) SetHeaders(v map[string]*string) *LocateResponse {
	s.Headers = v
	return s
}

func (s *LocateResponse) SetStatusCode(v int32) *LocateResponse {
	s.StatusCode = &v
	return s
}

func (s *LocateResponse) SetBody(v *LocateResponseBody) *LocateResponse {
	s.Body = v
	return s
}

type LoginModelScopeRequest struct {
	EmpId   *string `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	EmpName *string `json:"EmpName,omitempty" xml:"EmpName,omitempty"`
	Token   *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LoginModelScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginModelScopeRequest) GoString() string {
	return s.String()
}

func (s *LoginModelScopeRequest) SetEmpId(v string) *LoginModelScopeRequest {
	s.EmpId = &v
	return s
}

func (s *LoginModelScopeRequest) SetEmpName(v string) *LoginModelScopeRequest {
	s.EmpName = &v
	return s
}

func (s *LoginModelScopeRequest) SetToken(v string) *LoginModelScopeRequest {
	s.Token = &v
	return s
}

func (s *LoginModelScopeRequest) SetType(v string) *LoginModelScopeRequest {
	s.Type = &v
	return s
}

type LoginModelScopeResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LoginModelScopeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                          `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                           `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LoginModelScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginModelScopeResponseBody) GoString() string {
	return s.String()
}

func (s *LoginModelScopeResponseBody) SetCode(v string) *LoginModelScopeResponseBody {
	s.Code = &v
	return s
}

func (s *LoginModelScopeResponseBody) SetData(v *LoginModelScopeResponseBodyData) *LoginModelScopeResponseBody {
	s.Data = v
	return s
}

func (s *LoginModelScopeResponseBody) SetErrorName(v string) *LoginModelScopeResponseBody {
	s.ErrorName = &v
	return s
}

func (s *LoginModelScopeResponseBody) SetHttpCode(v int32) *LoginModelScopeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *LoginModelScopeResponseBody) SetMessage(v string) *LoginModelScopeResponseBody {
	s.Message = &v
	return s
}

func (s *LoginModelScopeResponseBody) SetRequestId(v string) *LoginModelScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginModelScopeResponseBody) SetSuccess(v bool) *LoginModelScopeResponseBody {
	s.Success = &v
	return s
}

type LoginModelScopeResponseBodyData struct {
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	Uid      *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s LoginModelScopeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LoginModelScopeResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoginModelScopeResponseBodyData) SetEmail(v string) *LoginModelScopeResponseBodyData {
	s.Email = &v
	return s
}

func (s *LoginModelScopeResponseBodyData) SetJwtToken(v string) *LoginModelScopeResponseBodyData {
	s.JwtToken = &v
	return s
}

func (s *LoginModelScopeResponseBodyData) SetNickname(v string) *LoginModelScopeResponseBodyData {
	s.Nickname = &v
	return s
}

func (s *LoginModelScopeResponseBodyData) SetUid(v string) *LoginModelScopeResponseBodyData {
	s.Uid = &v
	return s
}

type LoginModelScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoginModelScopeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoginModelScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginModelScopeResponse) GoString() string {
	return s.String()
}

func (s *LoginModelScopeResponse) SetHeaders(v map[string]*string) *LoginModelScopeResponse {
	s.Headers = v
	return s
}

func (s *LoginModelScopeResponse) SetStatusCode(v int32) *LoginModelScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginModelScopeResponse) SetBody(v *LoginModelScopeResponseBody) *LoginModelScopeResponse {
	s.Body = v
	return s
}

type MotionShopVideoDetectRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s MotionShopVideoDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s MotionShopVideoDetectRequest) GoString() string {
	return s.String()
}

func (s *MotionShopVideoDetectRequest) SetJwtToken(v string) *MotionShopVideoDetectRequest {
	s.JwtToken = &v
	return s
}

func (s *MotionShopVideoDetectRequest) SetOssKey(v string) *MotionShopVideoDetectRequest {
	s.OssKey = &v
	return s
}

type MotionShopVideoDetectResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MotionShopVideoDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MotionShopVideoDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MotionShopVideoDetectResponseBody) GoString() string {
	return s.String()
}

func (s *MotionShopVideoDetectResponseBody) SetCode(v string) *MotionShopVideoDetectResponseBody {
	s.Code = &v
	return s
}

func (s *MotionShopVideoDetectResponseBody) SetData(v *MotionShopVideoDetectResponseBodyData) *MotionShopVideoDetectResponseBody {
	s.Data = v
	return s
}

func (s *MotionShopVideoDetectResponseBody) SetMessage(v string) *MotionShopVideoDetectResponseBody {
	s.Message = &v
	return s
}

func (s *MotionShopVideoDetectResponseBody) SetRequestId(v string) *MotionShopVideoDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *MotionShopVideoDetectResponseBody) SetSuccess(v bool) *MotionShopVideoDetectResponseBody {
	s.Success = &v
	return s
}

type MotionShopVideoDetectResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s MotionShopVideoDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MotionShopVideoDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *MotionShopVideoDetectResponseBodyData) SetJobId(v string) *MotionShopVideoDetectResponseBodyData {
	s.JobId = &v
	return s
}

type MotionShopVideoDetectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MotionShopVideoDetectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MotionShopVideoDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s MotionShopVideoDetectResponse) GoString() string {
	return s.String()
}

func (s *MotionShopVideoDetectResponse) SetHeaders(v map[string]*string) *MotionShopVideoDetectResponse {
	s.Headers = v
	return s
}

func (s *MotionShopVideoDetectResponse) SetStatusCode(v int32) *MotionShopVideoDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *MotionShopVideoDetectResponse) SetBody(v *MotionShopVideoDetectResponseBody) *MotionShopVideoDetectResponse {
	s.Body = v
	return s
}

type PopBatchQueryObjectGenerationProjectStatusRequest struct {
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s PopBatchQueryObjectGenerationProjectStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectGenerationProjectStatusRequest) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectGenerationProjectStatusRequest) SetJwtToken(v string) *PopBatchQueryObjectGenerationProjectStatusRequest {
	s.JwtToken = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusRequest) SetProjectIds(v string) *PopBatchQueryObjectGenerationProjectStatusRequest {
	s.ProjectIds = &v
	return s
}

type PopBatchQueryObjectGenerationProjectStatusResponseBody struct {
	Code      *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PopBatchQueryObjectGenerationProjectStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBatchQueryObjectGenerationProjectStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectGenerationProjectStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBody) SetCode(v string) *PopBatchQueryObjectGenerationProjectStatusResponseBody {
	s.Code = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBody) SetData(v []*PopBatchQueryObjectGenerationProjectStatusResponseBodyData) *PopBatchQueryObjectGenerationProjectStatusResponseBody {
	s.Data = v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBody) SetMessage(v string) *PopBatchQueryObjectGenerationProjectStatusResponseBody {
	s.Message = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBody) SetRequestId(v string) *PopBatchQueryObjectGenerationProjectStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBody) SetSuccess(v bool) *PopBatchQueryObjectGenerationProjectStatusResponseBody {
	s.Success = &v
	return s
}

type PopBatchQueryObjectGenerationProjectStatusResponseBodyData struct {
	BizUsage *string                                                            `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Dataset  *PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Id       *string                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Status   *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopBatchQueryObjectGenerationProjectStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectGenerationProjectStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBodyData) SetBizUsage(v string) *PopBatchQueryObjectGenerationProjectStatusResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBodyData) SetDataset(v *PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset) *PopBatchQueryObjectGenerationProjectStatusResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBodyData) SetId(v string) *PopBatchQueryObjectGenerationProjectStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBodyData) SetStatus(v string) *PopBatchQueryObjectGenerationProjectStatusResponseBodyData {
	s.Status = &v
	return s
}

type PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset struct {
	BuildResultUrl map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
}

func (s PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopBatchQueryObjectGenerationProjectStatusResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

type PopBatchQueryObjectGenerationProjectStatusResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBatchQueryObjectGenerationProjectStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBatchQueryObjectGenerationProjectStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectGenerationProjectStatusResponse) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponse) SetHeaders(v map[string]*string) *PopBatchQueryObjectGenerationProjectStatusResponse {
	s.Headers = v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponse) SetStatusCode(v int32) *PopBatchQueryObjectGenerationProjectStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBatchQueryObjectGenerationProjectStatusResponse) SetBody(v *PopBatchQueryObjectGenerationProjectStatusResponseBody) *PopBatchQueryObjectGenerationProjectStatusResponse {
	s.Body = v
	return s
}

type PopBatchQueryObjectProjectStatusRequest struct {
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s PopBatchQueryObjectProjectStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectProjectStatusRequest) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectProjectStatusRequest) SetJwtToken(v string) *PopBatchQueryObjectProjectStatusRequest {
	s.JwtToken = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusRequest) SetProjectIds(v string) *PopBatchQueryObjectProjectStatusRequest {
	s.ProjectIds = &v
	return s
}

type PopBatchQueryObjectProjectStatusResponseBody struct {
	Code      *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PopBatchQueryObjectProjectStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                                             `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                              `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBatchQueryObjectProjectStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectProjectStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetCode(v string) *PopBatchQueryObjectProjectStatusResponseBody {
	s.Code = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetData(v []*PopBatchQueryObjectProjectStatusResponseBodyData) *PopBatchQueryObjectProjectStatusResponseBody {
	s.Data = v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetErrorName(v string) *PopBatchQueryObjectProjectStatusResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetHttpCode(v int32) *PopBatchQueryObjectProjectStatusResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetMessage(v string) *PopBatchQueryObjectProjectStatusResponseBody {
	s.Message = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetRequestId(v string) *PopBatchQueryObjectProjectStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBody) SetSuccess(v bool) *PopBatchQueryObjectProjectStatusResponseBody {
	s.Success = &v
	return s
}

type PopBatchQueryObjectProjectStatusResponseBodyData struct {
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopBatchQueryObjectProjectStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectProjectStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectProjectStatusResponseBodyData) SetCheckStatus(v string) *PopBatchQueryObjectProjectStatusResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBodyData) SetId(v string) *PopBatchQueryObjectProjectStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponseBodyData) SetStatus(v string) *PopBatchQueryObjectProjectStatusResponseBodyData {
	s.Status = &v
	return s
}

type PopBatchQueryObjectProjectStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBatchQueryObjectProjectStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBatchQueryObjectProjectStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBatchQueryObjectProjectStatusResponse) GoString() string {
	return s.String()
}

func (s *PopBatchQueryObjectProjectStatusResponse) SetHeaders(v map[string]*string) *PopBatchQueryObjectProjectStatusResponse {
	s.Headers = v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponse) SetStatusCode(v int32) *PopBatchQueryObjectProjectStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBatchQueryObjectProjectStatusResponse) SetBody(v *PopBatchQueryObjectProjectStatusResponseBody) *PopBatchQueryObjectProjectStatusResponse {
	s.Body = v
	return s
}

type PopBuildFeatureToAvatarProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopBuildFeatureToAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBuildFeatureToAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *PopBuildFeatureToAvatarProjectRequest) SetProjectId(v string) *PopBuildFeatureToAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

type PopBuildFeatureToAvatarProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBuildFeatureToAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBuildFeatureToAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopBuildFeatureToAvatarProjectResponseBody) SetCode(v string) *PopBuildFeatureToAvatarProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopBuildFeatureToAvatarProjectResponseBody) SetMessage(v string) *PopBuildFeatureToAvatarProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopBuildFeatureToAvatarProjectResponseBody) SetRequestId(v string) *PopBuildFeatureToAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBuildFeatureToAvatarProjectResponseBody) SetSuccess(v bool) *PopBuildFeatureToAvatarProjectResponseBody {
	s.Success = &v
	return s
}

type PopBuildFeatureToAvatarProjectResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBuildFeatureToAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBuildFeatureToAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBuildFeatureToAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *PopBuildFeatureToAvatarProjectResponse) SetHeaders(v map[string]*string) *PopBuildFeatureToAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *PopBuildFeatureToAvatarProjectResponse) SetStatusCode(v int32) *PopBuildFeatureToAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBuildFeatureToAvatarProjectResponse) SetBody(v *PopBuildFeatureToAvatarProjectResponseBody) *PopBuildFeatureToAvatarProjectResponse {
	s.Body = v
	return s
}

type PopBuildLivePortraitModelScopeProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopBuildLivePortraitModelScopeProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBuildLivePortraitModelScopeProjectRequest) GoString() string {
	return s.String()
}

func (s *PopBuildLivePortraitModelScopeProjectRequest) SetProjectId(v string) *PopBuildLivePortraitModelScopeProjectRequest {
	s.ProjectId = &v
	return s
}

type PopBuildLivePortraitModelScopeProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBuildLivePortraitModelScopeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBuildLivePortraitModelScopeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopBuildLivePortraitModelScopeProjectResponseBody) SetCode(v string) *PopBuildLivePortraitModelScopeProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopBuildLivePortraitModelScopeProjectResponseBody) SetMessage(v string) *PopBuildLivePortraitModelScopeProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopBuildLivePortraitModelScopeProjectResponseBody) SetRequestId(v string) *PopBuildLivePortraitModelScopeProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBuildLivePortraitModelScopeProjectResponseBody) SetSuccess(v bool) *PopBuildLivePortraitModelScopeProjectResponseBody {
	s.Success = &v
	return s
}

type PopBuildLivePortraitModelScopeProjectResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBuildLivePortraitModelScopeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBuildLivePortraitModelScopeProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBuildLivePortraitModelScopeProjectResponse) GoString() string {
	return s.String()
}

func (s *PopBuildLivePortraitModelScopeProjectResponse) SetHeaders(v map[string]*string) *PopBuildLivePortraitModelScopeProjectResponse {
	s.Headers = v
	return s
}

func (s *PopBuildLivePortraitModelScopeProjectResponse) SetStatusCode(v int32) *PopBuildLivePortraitModelScopeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBuildLivePortraitModelScopeProjectResponse) SetBody(v *PopBuildLivePortraitModelScopeProjectResponseBody) *PopBuildLivePortraitModelScopeProjectResponse {
	s.Body = v
	return s
}

type PopBuildObjectGenerationProjectRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopBuildObjectGenerationProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBuildObjectGenerationProjectRequest) GoString() string {
	return s.String()
}

func (s *PopBuildObjectGenerationProjectRequest) SetJwtToken(v string) *PopBuildObjectGenerationProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopBuildObjectGenerationProjectRequest) SetProjectId(v string) *PopBuildObjectGenerationProjectRequest {
	s.ProjectId = &v
	return s
}

type PopBuildObjectGenerationProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBuildObjectGenerationProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBuildObjectGenerationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopBuildObjectGenerationProjectResponseBody) SetCode(v string) *PopBuildObjectGenerationProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopBuildObjectGenerationProjectResponseBody) SetMessage(v string) *PopBuildObjectGenerationProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopBuildObjectGenerationProjectResponseBody) SetRequestId(v string) *PopBuildObjectGenerationProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBuildObjectGenerationProjectResponseBody) SetSuccess(v bool) *PopBuildObjectGenerationProjectResponseBody {
	s.Success = &v
	return s
}

type PopBuildObjectGenerationProjectResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBuildObjectGenerationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBuildObjectGenerationProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBuildObjectGenerationProjectResponse) GoString() string {
	return s.String()
}

func (s *PopBuildObjectGenerationProjectResponse) SetHeaders(v map[string]*string) *PopBuildObjectGenerationProjectResponse {
	s.Headers = v
	return s
}

func (s *PopBuildObjectGenerationProjectResponse) SetStatusCode(v int32) *PopBuildObjectGenerationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBuildObjectGenerationProjectResponse) SetBody(v *PopBuildObjectGenerationProjectResponseBody) *PopBuildObjectGenerationProjectResponse {
	s.Body = v
	return s
}

type PopBuildObjectProjectRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopBuildObjectProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBuildObjectProjectRequest) GoString() string {
	return s.String()
}

func (s *PopBuildObjectProjectRequest) SetJwtToken(v string) *PopBuildObjectProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopBuildObjectProjectRequest) SetProjectId(v string) *PopBuildObjectProjectRequest {
	s.ProjectId = &v
	return s
}

type PopBuildObjectProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBuildObjectProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBuildObjectProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopBuildObjectProjectResponseBody) SetCode(v string) *PopBuildObjectProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopBuildObjectProjectResponseBody) SetErrorName(v string) *PopBuildObjectProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopBuildObjectProjectResponseBody) SetHttpCode(v int32) *PopBuildObjectProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopBuildObjectProjectResponseBody) SetMessage(v string) *PopBuildObjectProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopBuildObjectProjectResponseBody) SetRequestId(v string) *PopBuildObjectProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBuildObjectProjectResponseBody) SetSuccess(v bool) *PopBuildObjectProjectResponseBody {
	s.Success = &v
	return s
}

type PopBuildObjectProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBuildObjectProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBuildObjectProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBuildObjectProjectResponse) GoString() string {
	return s.String()
}

func (s *PopBuildObjectProjectResponse) SetHeaders(v map[string]*string) *PopBuildObjectProjectResponse {
	s.Headers = v
	return s
}

func (s *PopBuildObjectProjectResponse) SetStatusCode(v int32) *PopBuildObjectProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBuildObjectProjectResponse) SetBody(v *PopBuildObjectProjectResponseBody) *PopBuildObjectProjectResponse {
	s.Body = v
	return s
}

type PopBuildPakRenderProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopBuildPakRenderProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBuildPakRenderProjectRequest) GoString() string {
	return s.String()
}

func (s *PopBuildPakRenderProjectRequest) SetProjectId(v string) *PopBuildPakRenderProjectRequest {
	s.ProjectId = &v
	return s
}

type PopBuildPakRenderProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBuildPakRenderProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBuildPakRenderProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopBuildPakRenderProjectResponseBody) SetCode(v string) *PopBuildPakRenderProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopBuildPakRenderProjectResponseBody) SetMessage(v string) *PopBuildPakRenderProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopBuildPakRenderProjectResponseBody) SetRequestId(v string) *PopBuildPakRenderProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBuildPakRenderProjectResponseBody) SetSuccess(v bool) *PopBuildPakRenderProjectResponseBody {
	s.Success = &v
	return s
}

type PopBuildPakRenderProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBuildPakRenderProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBuildPakRenderProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBuildPakRenderProjectResponse) GoString() string {
	return s.String()
}

func (s *PopBuildPakRenderProjectResponse) SetHeaders(v map[string]*string) *PopBuildPakRenderProjectResponse {
	s.Headers = v
	return s
}

func (s *PopBuildPakRenderProjectResponse) SetStatusCode(v int32) *PopBuildPakRenderProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBuildPakRenderProjectResponse) SetBody(v *PopBuildPakRenderProjectResponseBody) *PopBuildPakRenderProjectResponse {
	s.Body = v
	return s
}

type PopBuildTextToAvatarProjectRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopBuildTextToAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopBuildTextToAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *PopBuildTextToAvatarProjectRequest) SetJwtToken(v string) *PopBuildTextToAvatarProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopBuildTextToAvatarProjectRequest) SetProjectId(v string) *PopBuildTextToAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

type PopBuildTextToAvatarProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopBuildTextToAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopBuildTextToAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopBuildTextToAvatarProjectResponseBody) SetCode(v string) *PopBuildTextToAvatarProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopBuildTextToAvatarProjectResponseBody) SetMessage(v string) *PopBuildTextToAvatarProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopBuildTextToAvatarProjectResponseBody) SetRequestId(v string) *PopBuildTextToAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopBuildTextToAvatarProjectResponseBody) SetSuccess(v bool) *PopBuildTextToAvatarProjectResponseBody {
	s.Success = &v
	return s
}

type PopBuildTextToAvatarProjectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopBuildTextToAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopBuildTextToAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopBuildTextToAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *PopBuildTextToAvatarProjectResponse) SetHeaders(v map[string]*string) *PopBuildTextToAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *PopBuildTextToAvatarProjectResponse) SetStatusCode(v int32) *PopBuildTextToAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopBuildTextToAvatarProjectResponse) SetBody(v *PopBuildTextToAvatarProjectResponseBody) *PopBuildTextToAvatarProjectResponse {
	s.Body = v
	return s
}

type PopCreateFeatureToAvatarProjectRequest struct {
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Intro   *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopCreateFeatureToAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectRequest) SetExtInfo(v string) *PopCreateFeatureToAvatarProjectRequest {
	s.ExtInfo = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectRequest) SetIntro(v string) *PopCreateFeatureToAvatarProjectRequest {
	s.Intro = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectRequest) SetTitle(v string) *PopCreateFeatureToAvatarProjectRequest {
	s.Title = &v
	return s
}

type PopCreateFeatureToAvatarProjectResponseBody struct {
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreateFeatureToAvatarProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreateFeatureToAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectResponseBody) SetCode(v string) *PopCreateFeatureToAvatarProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBody) SetData(v *PopCreateFeatureToAvatarProjectResponseBodyData) *PopCreateFeatureToAvatarProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBody) SetMessage(v string) *PopCreateFeatureToAvatarProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBody) SetRequestId(v string) *PopCreateFeatureToAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBody) SetSuccess(v bool) *PopCreateFeatureToAvatarProjectResponseBody {
	s.Success = &v
	return s
}

type PopCreateFeatureToAvatarProjectResponseBodyData struct {
	BizUsage         *string                                                     `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                     `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                     `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset          *PopCreateFeatureToAvatarProjectResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                       `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext              *string                                                     `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                     `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                     `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                     `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status           *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreateFeatureToAvatarProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetBizUsage(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetBuildDetail(v *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetCheckStatus(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetCreateMode(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetCreateTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetDataset(v *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetDeleted(v bool) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetExt(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetId(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetIntro(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetMaterialCoverUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetModifiedTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetStatus(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetTitle(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyData) SetType(v string) *PopCreateFeatureToAvatarProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetCreateTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetDeleted(v bool) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetId(v int64) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.Id = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetRunningTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetStatus(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.Status = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopCreateFeatureToAvatarProjectResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                                        `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                                       `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                         `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                                       `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	Id              *int64                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string                                                       `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                                       `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                                       `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                                       `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                                       `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                                       `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopCreateFeatureToAvatarProjectResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetCoverUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetCreateTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetDeleted(v bool) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetErrorCode(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetErrorMessage(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetGlbModelUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetId(v int64) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.Id = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetModelUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetModifiedTime(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetOriginResultUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetOssKey(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetPolicy(v *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetPoseUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDataset) SetPreviewUrl(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

type PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetDir(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetExpire(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetHost(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetSignature(v string) *PopCreateFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopCreateFeatureToAvatarProjectResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreateFeatureToAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreateFeatureToAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreateFeatureToAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *PopCreateFeatureToAvatarProjectResponse) SetHeaders(v map[string]*string) *PopCreateFeatureToAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponse) SetStatusCode(v int32) *PopCreateFeatureToAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreateFeatureToAvatarProjectResponse) SetBody(v *PopCreateFeatureToAvatarProjectResponseBody) *PopCreateFeatureToAvatarProjectResponse {
	s.Body = v
	return s
}

type PopCreateLivePortraitModelScopeProjectRequest struct {
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Intro   *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopCreateLivePortraitModelScopeProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreateLivePortraitModelScopeProjectRequest) GoString() string {
	return s.String()
}

func (s *PopCreateLivePortraitModelScopeProjectRequest) SetExtInfo(v string) *PopCreateLivePortraitModelScopeProjectRequest {
	s.ExtInfo = &v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectRequest) SetIntro(v string) *PopCreateLivePortraitModelScopeProjectRequest {
	s.Intro = &v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectRequest) SetTitle(v string) *PopCreateLivePortraitModelScopeProjectRequest {
	s.Title = &v
	return s
}

type PopCreateLivePortraitModelScopeProjectResponseBody struct {
	Code      *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreateLivePortraitModelScopeProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreateLivePortraitModelScopeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreateLivePortraitModelScopeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreateLivePortraitModelScopeProjectResponseBody) SetCode(v string) *PopCreateLivePortraitModelScopeProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectResponseBody) SetData(v *PopCreateLivePortraitModelScopeProjectResponseBodyData) *PopCreateLivePortraitModelScopeProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectResponseBody) SetMessage(v string) *PopCreateLivePortraitModelScopeProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectResponseBody) SetRequestId(v string) *PopCreateLivePortraitModelScopeProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectResponseBody) SetSuccess(v bool) *PopCreateLivePortraitModelScopeProjectResponseBody {
	s.Success = &v
	return s
}

type PopCreateLivePortraitModelScopeProjectResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PopCreateLivePortraitModelScopeProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreateLivePortraitModelScopeProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreateLivePortraitModelScopeProjectResponseBodyData) SetId(v string) *PopCreateLivePortraitModelScopeProjectResponseBodyData {
	s.Id = &v
	return s
}

type PopCreateLivePortraitModelScopeProjectResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreateLivePortraitModelScopeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreateLivePortraitModelScopeProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreateLivePortraitModelScopeProjectResponse) GoString() string {
	return s.String()
}

func (s *PopCreateLivePortraitModelScopeProjectResponse) SetHeaders(v map[string]*string) *PopCreateLivePortraitModelScopeProjectResponse {
	s.Headers = v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectResponse) SetStatusCode(v int32) *PopCreateLivePortraitModelScopeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreateLivePortraitModelScopeProjectResponse) SetBody(v *PopCreateLivePortraitModelScopeProjectResponseBody) *PopCreateLivePortraitModelScopeProjectResponse {
	s.Body = v
	return s
}

type PopCreateMaterialRequest struct {
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Ext         *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Intro       *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken    *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ListStatus  *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey      *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreateMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreateMaterialRequest) GoString() string {
	return s.String()
}

func (s *PopCreateMaterialRequest) SetCheckStatus(v string) *PopCreateMaterialRequest {
	s.CheckStatus = &v
	return s
}

func (s *PopCreateMaterialRequest) SetExt(v string) *PopCreateMaterialRequest {
	s.Ext = &v
	return s
}

func (s *PopCreateMaterialRequest) SetIntro(v string) *PopCreateMaterialRequest {
	s.Intro = &v
	return s
}

func (s *PopCreateMaterialRequest) SetJwtToken(v string) *PopCreateMaterialRequest {
	s.JwtToken = &v
	return s
}

func (s *PopCreateMaterialRequest) SetListStatus(v string) *PopCreateMaterialRequest {
	s.ListStatus = &v
	return s
}

func (s *PopCreateMaterialRequest) SetName(v string) *PopCreateMaterialRequest {
	s.Name = &v
	return s
}

func (s *PopCreateMaterialRequest) SetOssKey(v string) *PopCreateMaterialRequest {
	s.OssKey = &v
	return s
}

func (s *PopCreateMaterialRequest) SetType(v string) *PopCreateMaterialRequest {
	s.Type = &v
	return s
}

type PopCreateMaterialResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreateMaterialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreateMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreateMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreateMaterialResponseBody) SetCode(v string) *PopCreateMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreateMaterialResponseBody) SetData(v *PopCreateMaterialResponseBodyData) *PopCreateMaterialResponseBody {
	s.Data = v
	return s
}

func (s *PopCreateMaterialResponseBody) SetMessage(v string) *PopCreateMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreateMaterialResponseBody) SetRequestId(v string) *PopCreateMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreateMaterialResponseBody) SetSuccess(v bool) *PopCreateMaterialResponseBody {
	s.Success = &v
	return s
}

type PopCreateMaterialResponseBodyData struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreateMaterialResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreateMaterialResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreateMaterialResponseBodyData) SetCheckStatus(v string) *PopCreateMaterialResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetCommon(v bool) *PopCreateMaterialResponseBodyData {
	s.Common = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetCoverUrl(v string) *PopCreateMaterialResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetCreateTime(v string) *PopCreateMaterialResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetDeleted(v bool) *PopCreateMaterialResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetExt(v string) *PopCreateMaterialResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetFileUrl(v string) *PopCreateMaterialResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetId(v string) *PopCreateMaterialResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetIntro(v string) *PopCreateMaterialResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetListStatus(v string) *PopCreateMaterialResponseBodyData {
	s.ListStatus = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetModifiedTime(v string) *PopCreateMaterialResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetName(v string) *PopCreateMaterialResponseBodyData {
	s.Name = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetOssKey(v string) *PopCreateMaterialResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetPreviewUrl(v string) *PopCreateMaterialResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *PopCreateMaterialResponseBodyData) SetType(v string) *PopCreateMaterialResponseBodyData {
	s.Type = &v
	return s
}

type PopCreateMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreateMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreateMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreateMaterialResponse) GoString() string {
	return s.String()
}

func (s *PopCreateMaterialResponse) SetHeaders(v map[string]*string) *PopCreateMaterialResponse {
	s.Headers = v
	return s
}

func (s *PopCreateMaterialResponse) SetStatusCode(v int32) *PopCreateMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreateMaterialResponse) SetBody(v *PopCreateMaterialResponseBody) *PopCreateMaterialResponse {
	s.Body = v
	return s
}

type PopCreateObjectGenerationProjectRequest struct {
	BizUsage *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	ExtInfo  *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Intro    *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopCreateObjectGenerationProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectGenerationProjectRequest) GoString() string {
	return s.String()
}

func (s *PopCreateObjectGenerationProjectRequest) SetBizUsage(v string) *PopCreateObjectGenerationProjectRequest {
	s.BizUsage = &v
	return s
}

func (s *PopCreateObjectGenerationProjectRequest) SetExtInfo(v string) *PopCreateObjectGenerationProjectRequest {
	s.ExtInfo = &v
	return s
}

func (s *PopCreateObjectGenerationProjectRequest) SetIntro(v string) *PopCreateObjectGenerationProjectRequest {
	s.Intro = &v
	return s
}

func (s *PopCreateObjectGenerationProjectRequest) SetJwtToken(v string) *PopCreateObjectGenerationProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopCreateObjectGenerationProjectRequest) SetTitle(v string) *PopCreateObjectGenerationProjectRequest {
	s.Title = &v
	return s
}

type PopCreateObjectGenerationProjectResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreateObjectGenerationProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreateObjectGenerationProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectGenerationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreateObjectGenerationProjectResponseBody) SetCode(v string) *PopCreateObjectGenerationProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreateObjectGenerationProjectResponseBody) SetData(v *PopCreateObjectGenerationProjectResponseBodyData) *PopCreateObjectGenerationProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopCreateObjectGenerationProjectResponseBody) SetMessage(v string) *PopCreateObjectGenerationProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreateObjectGenerationProjectResponseBody) SetRequestId(v string) *PopCreateObjectGenerationProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreateObjectGenerationProjectResponseBody) SetSuccess(v bool) *PopCreateObjectGenerationProjectResponseBody {
	s.Success = &v
	return s
}

type PopCreateObjectGenerationProjectResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PopCreateObjectGenerationProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectGenerationProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreateObjectGenerationProjectResponseBodyData) SetId(v string) *PopCreateObjectGenerationProjectResponseBodyData {
	s.Id = &v
	return s
}

type PopCreateObjectGenerationProjectResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreateObjectGenerationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreateObjectGenerationProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectGenerationProjectResponse) GoString() string {
	return s.String()
}

func (s *PopCreateObjectGenerationProjectResponse) SetHeaders(v map[string]*string) *PopCreateObjectGenerationProjectResponse {
	s.Headers = v
	return s
}

func (s *PopCreateObjectGenerationProjectResponse) SetStatusCode(v int32) *PopCreateObjectGenerationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreateObjectGenerationProjectResponse) SetBody(v *PopCreateObjectGenerationProjectResponseBody) *PopCreateObjectGenerationProjectResponse {
	s.Body = v
	return s
}

type PopCreateObjectProjectRequest struct {
	AutoBuild       *bool   `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage        *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	CustomSource    *string `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dependencies    *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	ForeignUid      *string `json:"ForeignUid,omitempty" xml:"ForeignUid,omitempty"`
	Intro           *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken        *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RecommendStatus *string `json:"RecommendStatus,omitempty" xml:"RecommendStatus,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopCreateObjectProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectRequest) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectRequest) SetAutoBuild(v bool) *PopCreateObjectProjectRequest {
	s.AutoBuild = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetBizUsage(v string) *PopCreateObjectProjectRequest {
	s.BizUsage = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetCustomSource(v string) *PopCreateObjectProjectRequest {
	s.CustomSource = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetDependencies(v string) *PopCreateObjectProjectRequest {
	s.Dependencies = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetForeignUid(v string) *PopCreateObjectProjectRequest {
	s.ForeignUid = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetIntro(v string) *PopCreateObjectProjectRequest {
	s.Intro = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetJwtToken(v string) *PopCreateObjectProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetMode(v string) *PopCreateObjectProjectRequest {
	s.Mode = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetRecommendStatus(v string) *PopCreateObjectProjectRequest {
	s.RecommendStatus = &v
	return s
}

func (s *PopCreateObjectProjectRequest) SetTitle(v string) *PopCreateObjectProjectRequest {
	s.Title = &v
	return s
}

type PopCreateObjectProjectResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreateObjectProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                                 `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreateObjectProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBody) SetCode(v string) *PopCreateObjectProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreateObjectProjectResponseBody) SetData(v *PopCreateObjectProjectResponseBodyData) *PopCreateObjectProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopCreateObjectProjectResponseBody) SetErrorName(v string) *PopCreateObjectProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopCreateObjectProjectResponseBody) SetHttpCode(v int32) *PopCreateObjectProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopCreateObjectProjectResponseBody) SetMessage(v string) *PopCreateObjectProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreateObjectProjectResponseBody) SetRequestId(v string) *PopCreateObjectProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreateObjectProjectResponseBody) SetSuccess(v bool) *PopCreateObjectProjectResponseBody {
	s.Success = &v
	return s
}

type PopCreateObjectProjectResponseBodyData struct {
	AuditStatus     *string                                            `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild       *bool                                              `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage        *string                                            `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail     *PopCreateObjectProjectResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus     *string                                            `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode      *string                                            `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime      *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource    *string                                            `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset         *PopCreateObjectProjectResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted         *bool                                              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies    *string                                            `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext             *string                                            `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id              *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro           *string                                            `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ModifiedTime    *string                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RecommendStatus *string                                            `json:"RecommendStatus,omitempty" xml:"RecommendStatus,omitempty"`
	Source          *PopCreateObjectProjectResponseBodyDataSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status          *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Title           *string                                            `json:"Title,omitempty" xml:"Title,omitempty"`
	Type            *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyData) SetAuditStatus(v string) *PopCreateObjectProjectResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetAutoBuild(v bool) *PopCreateObjectProjectResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetBizUsage(v string) *PopCreateObjectProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetBuildDetail(v *PopCreateObjectProjectResponseBodyDataBuildDetail) *PopCreateObjectProjectResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetCheckStatus(v string) *PopCreateObjectProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetCreateMode(v string) *PopCreateObjectProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetCreateTime(v string) *PopCreateObjectProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetCustomSource(v string) *PopCreateObjectProjectResponseBodyData {
	s.CustomSource = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetDataset(v *PopCreateObjectProjectResponseBodyDataDataset) *PopCreateObjectProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetDeleted(v bool) *PopCreateObjectProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetDependencies(v string) *PopCreateObjectProjectResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetExt(v string) *PopCreateObjectProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetId(v string) *PopCreateObjectProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetIntro(v string) *PopCreateObjectProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetModifiedTime(v string) *PopCreateObjectProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetRecommendStatus(v string) *PopCreateObjectProjectResponseBodyData {
	s.RecommendStatus = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetSource(v *PopCreateObjectProjectResponseBodyDataSource) *PopCreateObjectProjectResponseBodyData {
	s.Source = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetStatus(v string) *PopCreateObjectProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetTitle(v string) *PopCreateObjectProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyData) SetType(v string) *PopCreateObjectProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetCreateTime(v string) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetDeleted(v bool) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetRunningTime(v string) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopCreateObjectProjectResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                               `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                              `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage    *string                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                              `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                                              `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                              `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                              `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopCreateObjectProjectResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                              `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                              `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Token           *PopCreateObjectProjectResponseBodyDataDatasetToken  `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s PopCreateObjectProjectResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopCreateObjectProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetCoverUrl(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetCreateTime(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetDeleted(v bool) *PopCreateObjectProjectResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetErrorMessage(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetGlbModelUrl(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetModelUrl(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetModifiedTime(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetOriginResultUrl(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetOssKey(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetPolicy(v *PopCreateObjectProjectResponseBodyDataDatasetPolicy) *PopCreateObjectProjectResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetPoseUrl(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetPreviewUrl(v string) *PopCreateObjectProjectResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDataset) SetToken(v *PopCreateObjectProjectResponseBodyDataDatasetToken) *PopCreateObjectProjectResponseBodyDataDataset {
	s.Token = v
	return s
}

type PopCreateObjectProjectResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopCreateObjectProjectResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetPolicy) SetDir(v string) *PopCreateObjectProjectResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetPolicy) SetExpire(v string) *PopCreateObjectProjectResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetPolicy) SetHost(v string) *PopCreateObjectProjectResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopCreateObjectProjectResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetPolicy) SetSignature(v string) *PopCreateObjectProjectResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataDatasetToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Dir             *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Host            *string `json:"Host,omitempty" xml:"Host,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataDatasetToken) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataDatasetToken) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetToken) SetAccessKeyId(v string) *PopCreateObjectProjectResponseBodyDataDatasetToken {
	s.AccessKeyId = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetToken) SetAccessKeySecret(v string) *PopCreateObjectProjectResponseBodyDataDatasetToken {
	s.AccessKeySecret = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetToken) SetDir(v string) *PopCreateObjectProjectResponseBodyDataDatasetToken {
	s.Dir = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetToken) SetExpiration(v string) *PopCreateObjectProjectResponseBodyDataDatasetToken {
	s.Expiration = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetToken) SetHost(v string) *PopCreateObjectProjectResponseBodyDataDatasetToken {
	s.Host = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataDatasetToken) SetSecurityToken(v string) *PopCreateObjectProjectResponseBodyDataDatasetToken {
	s.SecurityToken = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataSource struct {
	Clothes      []*PopCreateObjectProjectResponseBodyDataSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                      `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string                                                    `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                    `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopCreateObjectProjectResponseBodyDataSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopCreateObjectProjectResponseBodyDataSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
	Token        *PopCreateObjectProjectResponseBodyDataSourceToken         `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s PopCreateObjectProjectResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetClothes(v []*PopCreateObjectProjectResponseBodyDataSourceClothes) *PopCreateObjectProjectResponseBodyDataSource {
	s.Clothes = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetCreateTime(v string) *PopCreateObjectProjectResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetDeleted(v bool) *PopCreateObjectProjectResponseBodyDataSource {
	s.Deleted = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetModifiedTime(v string) *PopCreateObjectProjectResponseBodyDataSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetOssKey(v string) *PopCreateObjectProjectResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetPolicy(v *PopCreateObjectProjectResponseBodyDataSourcePolicy) *PopCreateObjectProjectResponseBodyDataSource {
	s.Policy = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetSourceFiles(v []*PopCreateObjectProjectResponseBodyDataSourceSourceFiles) *PopCreateObjectProjectResponseBodyDataSource {
	s.SourceFiles = v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSource) SetToken(v *PopCreateObjectProjectResponseBodyDataSourceToken) *PopCreateObjectProjectResponseBodyDataSource {
	s.Token = v
	return s
}

type PopCreateObjectProjectResponseBodyDataSourceClothes struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataSourceClothes) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetCoverUrl(v string) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetCreateTime(v string) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetDeleted(v bool) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetModifiedTime(v string) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetName(v string) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.Name = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetOssKey(v string) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceClothes) SetType(v string) *PopCreateObjectProjectResponseBodyDataSourceClothes {
	s.Type = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataSourcePolicy) SetAccessId(v string) *PopCreateObjectProjectResponseBodyDataSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourcePolicy) SetDir(v string) *PopCreateObjectProjectResponseBodyDataSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourcePolicy) SetExpire(v string) *PopCreateObjectProjectResponseBodyDataSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourcePolicy) SetHost(v string) *PopCreateObjectProjectResponseBodyDataSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourcePolicy) SetPolicy(v string) *PopCreateObjectProjectResponseBodyDataSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourcePolicy) SetSignature(v string) *PopCreateObjectProjectResponseBodyDataSourcePolicy {
	s.Signature = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetCoverUrl(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetCreateTime(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetDeleted(v bool) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetFileName(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetFilesize(v int64) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetModifiedTime(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetOssKey(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetType(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceSourceFiles) SetUrl(v string) *PopCreateObjectProjectResponseBodyDataSourceSourceFiles {
	s.Url = &v
	return s
}

type PopCreateObjectProjectResponseBodyDataSourceToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Dir             *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Host            *string `json:"Host,omitempty" xml:"Host,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PopCreateObjectProjectResponseBodyDataSourceToken) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponseBodyDataSourceToken) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponseBodyDataSourceToken) SetAccessKeyId(v string) *PopCreateObjectProjectResponseBodyDataSourceToken {
	s.AccessKeyId = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceToken) SetAccessKeySecret(v string) *PopCreateObjectProjectResponseBodyDataSourceToken {
	s.AccessKeySecret = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceToken) SetDir(v string) *PopCreateObjectProjectResponseBodyDataSourceToken {
	s.Dir = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceToken) SetExpiration(v string) *PopCreateObjectProjectResponseBodyDataSourceToken {
	s.Expiration = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceToken) SetHost(v string) *PopCreateObjectProjectResponseBodyDataSourceToken {
	s.Host = &v
	return s
}

func (s *PopCreateObjectProjectResponseBodyDataSourceToken) SetSecurityToken(v string) *PopCreateObjectProjectResponseBodyDataSourceToken {
	s.SecurityToken = &v
	return s
}

type PopCreateObjectProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreateObjectProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreateObjectProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreateObjectProjectResponse) GoString() string {
	return s.String()
}

func (s *PopCreateObjectProjectResponse) SetHeaders(v map[string]*string) *PopCreateObjectProjectResponse {
	s.Headers = v
	return s
}

func (s *PopCreateObjectProjectResponse) SetStatusCode(v int32) *PopCreateObjectProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreateObjectProjectResponse) SetBody(v *PopCreateObjectProjectResponseBody) *PopCreateObjectProjectResponse {
	s.Body = v
	return s
}

type PopCreatePakRenderProjectRequest struct {
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Intro   *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopCreatePakRenderProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreatePakRenderProjectRequest) GoString() string {
	return s.String()
}

func (s *PopCreatePakRenderProjectRequest) SetExtInfo(v string) *PopCreatePakRenderProjectRequest {
	s.ExtInfo = &v
	return s
}

func (s *PopCreatePakRenderProjectRequest) SetIntro(v string) *PopCreatePakRenderProjectRequest {
	s.Intro = &v
	return s
}

func (s *PopCreatePakRenderProjectRequest) SetTitle(v string) *PopCreatePakRenderProjectRequest {
	s.Title = &v
	return s
}

type PopCreatePakRenderProjectResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreatePakRenderProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreatePakRenderProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreatePakRenderProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreatePakRenderProjectResponseBody) SetCode(v string) *PopCreatePakRenderProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBody) SetData(v *PopCreatePakRenderProjectResponseBodyData) *PopCreatePakRenderProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopCreatePakRenderProjectResponseBody) SetMessage(v string) *PopCreatePakRenderProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBody) SetRequestId(v string) *PopCreatePakRenderProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBody) SetSuccess(v bool) *PopCreatePakRenderProjectResponseBody {
	s.Success = &v
	return s
}

type PopCreatePakRenderProjectResponseBodyData struct {
	AutoBuild        *bool   `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	CheckStatus      *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted          *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreatePakRenderProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreatePakRenderProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetAutoBuild(v bool) *PopCreatePakRenderProjectResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetBizUsage(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetCheckStatus(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetCreateMode(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetCreateTime(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetDeleted(v bool) *PopCreatePakRenderProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetDependencies(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetExt(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetId(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetIntro(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetMaterialCoverUrl(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetModifiedTime(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetStatus(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetTitle(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopCreatePakRenderProjectResponseBodyData) SetType(v string) *PopCreatePakRenderProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopCreatePakRenderProjectResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreatePakRenderProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreatePakRenderProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreatePakRenderProjectResponse) GoString() string {
	return s.String()
}

func (s *PopCreatePakRenderProjectResponse) SetHeaders(v map[string]*string) *PopCreatePakRenderProjectResponse {
	s.Headers = v
	return s
}

func (s *PopCreatePakRenderProjectResponse) SetStatusCode(v int32) *PopCreatePakRenderProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreatePakRenderProjectResponse) SetBody(v *PopCreatePakRenderProjectResponseBody) *PopCreatePakRenderProjectResponse {
	s.Body = v
	return s
}

type PopCreateTextToAvatarProjectRequest struct {
	ExtInfo  *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Intro    *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopCreateTextToAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopCreateTextToAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *PopCreateTextToAvatarProjectRequest) SetExtInfo(v string) *PopCreateTextToAvatarProjectRequest {
	s.ExtInfo = &v
	return s
}

func (s *PopCreateTextToAvatarProjectRequest) SetIntro(v string) *PopCreateTextToAvatarProjectRequest {
	s.Intro = &v
	return s
}

func (s *PopCreateTextToAvatarProjectRequest) SetJwtToken(v string) *PopCreateTextToAvatarProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopCreateTextToAvatarProjectRequest) SetTitle(v string) *PopCreateTextToAvatarProjectRequest {
	s.Title = &v
	return s
}

type PopCreateTextToAvatarProjectResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopCreateTextToAvatarProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopCreateTextToAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopCreateTextToAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopCreateTextToAvatarProjectResponseBody) SetCode(v string) *PopCreateTextToAvatarProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBody) SetData(v *PopCreateTextToAvatarProjectResponseBodyData) *PopCreateTextToAvatarProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBody) SetMessage(v string) *PopCreateTextToAvatarProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBody) SetRequestId(v string) *PopCreateTextToAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBody) SetSuccess(v bool) *PopCreateTextToAvatarProjectResponseBody {
	s.Success = &v
	return s
}

type PopCreateTextToAvatarProjectResponseBodyData struct {
	AutoBuild        *bool   `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	CheckStatus      *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted          *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopCreateTextToAvatarProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopCreateTextToAvatarProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetAutoBuild(v bool) *PopCreateTextToAvatarProjectResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetBizUsage(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetCheckStatus(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetCreateMode(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetCreateTime(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetDeleted(v bool) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetDependencies(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetExt(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetId(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetIntro(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetMaterialCoverUrl(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetModifiedTime(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetStatus(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetTitle(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponseBodyData) SetType(v string) *PopCreateTextToAvatarProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopCreateTextToAvatarProjectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopCreateTextToAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopCreateTextToAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopCreateTextToAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *PopCreateTextToAvatarProjectResponse) SetHeaders(v map[string]*string) *PopCreateTextToAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *PopCreateTextToAvatarProjectResponse) SetStatusCode(v int32) *PopCreateTextToAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopCreateTextToAvatarProjectResponse) SetBody(v *PopCreateTextToAvatarProjectResponseBody) *PopCreateTextToAvatarProjectResponse {
	s.Body = v
	return s
}

type PopDeleteMaterialRequest struct {
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
}

func (s PopDeleteMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s PopDeleteMaterialRequest) GoString() string {
	return s.String()
}

func (s *PopDeleteMaterialRequest) SetJwtToken(v string) *PopDeleteMaterialRequest {
	s.JwtToken = &v
	return s
}

func (s *PopDeleteMaterialRequest) SetMaterialId(v string) *PopDeleteMaterialRequest {
	s.MaterialId = &v
	return s
}

type PopDeleteMaterialResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopDeleteMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopDeleteMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *PopDeleteMaterialResponseBody) SetCode(v string) *PopDeleteMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *PopDeleteMaterialResponseBody) SetMessage(v string) *PopDeleteMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *PopDeleteMaterialResponseBody) SetRequestId(v string) *PopDeleteMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopDeleteMaterialResponseBody) SetSuccess(v bool) *PopDeleteMaterialResponseBody {
	s.Success = &v
	return s
}

type PopDeleteMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopDeleteMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopDeleteMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s PopDeleteMaterialResponse) GoString() string {
	return s.String()
}

func (s *PopDeleteMaterialResponse) SetHeaders(v map[string]*string) *PopDeleteMaterialResponse {
	s.Headers = v
	return s
}

func (s *PopDeleteMaterialResponse) SetStatusCode(v int32) *PopDeleteMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *PopDeleteMaterialResponse) SetBody(v *PopDeleteMaterialResponseBody) *PopDeleteMaterialResponse {
	s.Body = v
	return s
}

type PopGetAITryOnJobRequest struct {
	JwtToken     *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WithMaterial *bool   `json:"WithMaterial,omitempty" xml:"WithMaterial,omitempty"`
	WithResult   *bool   `json:"WithResult,omitempty" xml:"WithResult,omitempty"`
}

func (s PopGetAITryOnJobRequest) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobRequest) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobRequest) SetJwtToken(v string) *PopGetAITryOnJobRequest {
	s.JwtToken = &v
	return s
}

func (s *PopGetAITryOnJobRequest) SetProjectId(v string) *PopGetAITryOnJobRequest {
	s.ProjectId = &v
	return s
}

func (s *PopGetAITryOnJobRequest) SetWithMaterial(v bool) *PopGetAITryOnJobRequest {
	s.WithMaterial = &v
	return s
}

func (s *PopGetAITryOnJobRequest) SetWithResult(v bool) *PopGetAITryOnJobRequest {
	s.WithResult = &v
	return s
}

type PopGetAITryOnJobResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopGetAITryOnJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopGetAITryOnJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBody) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBody) SetCode(v string) *PopGetAITryOnJobResponseBody {
	s.Code = &v
	return s
}

func (s *PopGetAITryOnJobResponseBody) SetData(v *PopGetAITryOnJobResponseBodyData) *PopGetAITryOnJobResponseBody {
	s.Data = v
	return s
}

func (s *PopGetAITryOnJobResponseBody) SetMessage(v string) *PopGetAITryOnJobResponseBody {
	s.Message = &v
	return s
}

func (s *PopGetAITryOnJobResponseBody) SetRequestId(v string) *PopGetAITryOnJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBody) SetSuccess(v bool) *PopGetAITryOnJobResponseBody {
	s.Success = &v
	return s
}

type PopGetAITryOnJobResponseBodyData struct {
	DummyProjectInfo *PopGetAITryOnJobResponseBodyDataDummyProjectInfo `json:"DummyProjectInfo,omitempty" xml:"DummyProjectInfo,omitempty" type:"Struct"`
	MaterialInfo     *PopGetAITryOnJobResponseBodyDataMaterialInfo     `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" type:"Struct"`
	SubTasks         []*PopGetAITryOnJobResponseBodyDataSubTasks       `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
}

func (s PopGetAITryOnJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyData) SetDummyProjectInfo(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) *PopGetAITryOnJobResponseBodyData {
	s.DummyProjectInfo = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyData) SetMaterialInfo(v *PopGetAITryOnJobResponseBodyDataMaterialInfo) *PopGetAITryOnJobResponseBodyData {
	s.MaterialInfo = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyData) SetSubTasks(v []*PopGetAITryOnJobResponseBodyDataSubTasks) *PopGetAITryOnJobResponseBodyData {
	s.SubTasks = v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfo struct {
	AuditStatus      *string                                                      `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild        *bool                                                        `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string                                                      `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                      `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                      `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource     *string                                                      `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset          *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                        `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                                      `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                                      `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                      `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                      `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                      `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source           *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status           *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                      `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfo) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetAuditStatus(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.AuditStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetAutoBuild(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.AutoBuild = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetBizUsage(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.BizUsage = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetBuildDetail(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.BuildDetail = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetCheckStatus(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.CheckStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetCreateMode(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.CreateMode = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetCustomSource(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.CustomSource = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetDataset(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Dataset = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetDependencies(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Dependencies = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetExt(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Ext = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetId(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetIntro(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Intro = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetMaterialCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetSource(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Source = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetStatus(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Status = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetTitle(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Title = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfo) SetType(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfo {
	s.Type = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetCompletedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetErrorMessage(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetEstimatedDuration(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetId(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetRunningTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetStatus(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.Status = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail) SetSubmitTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset struct {
	BuildResultUrl  map[string]interface{}                                         `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                                        `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                          `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                                        `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	Id              *int64                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string                                                        `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                                        `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                                        `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                                        `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                                        `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetBuildResultUrl(v map[string]interface{}) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetErrorCode(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetErrorMessage(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetGlbModelUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetId(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetModelUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetOriginResultUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetPolicy(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.Policy = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetPoseUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset) SetPreviewUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDataset {
	s.PreviewUrl = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) SetAccessId(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) SetDir(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) SetExpire(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) SetHost(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) SetPolicy(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy) SetSignature(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Signature = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource struct {
	Clothes      []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                              `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
	Token        *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken         `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetClothes(v []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.Clothes = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetId(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetPolicy(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.Policy = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetSourceFiles(v []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.SourceFiles = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource) SetToken(v *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSource {
	s.Token = v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes struct {
	CoverUrl     *string                                                                  `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                    `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string                                                                  `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Part         *string                                                                  `json:"Part,omitempty" xml:"Part,omitempty"`
	Size         *string                                                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	SkuProps     []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps `json:"SkuProps,omitempty" xml:"SkuProps,omitempty" type:"Repeated"`
	Skus         []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus     `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	Status       map[string]*string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	Version      *int32                                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetId(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetName(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetPart(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Part = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetSize(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Size = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetSkuProps(v []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.SkuProps = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetSkus(v []*PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Skus = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetStatus(v map[string]*string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Status = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetType(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Type = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes) SetVersion(v int32) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothes {
	s.Version = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Options []*string `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps) SetName(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps) SetOptions(v []*string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkuProps {
	s.Options = v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus struct {
	Color  *string `json:"Color,omitempty" xml:"Color,omitempty"`
	Cover  *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	Size   *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) SetColor(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Color = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) SetCover(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Cover = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) SetSize(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Size = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus) SetStatus(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Status = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) SetAccessId(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) SetDir(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) SetExpire(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) SetHost(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) SetPolicy(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy) SetSignature(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Signature = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetFileName(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetFilesize(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetId(v int64) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetType(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles) SetUrl(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Url = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Dir             *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Host            *string `json:"Host,omitempty" xml:"Host,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) SetAccessKeyId(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken {
	s.AccessKeyId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) SetAccessKeySecret(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken {
	s.AccessKeySecret = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) SetDir(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken {
	s.Dir = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) SetExpiration(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken {
	s.Expiration = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) SetHost(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken {
	s.Host = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken) SetSecurityToken(v string) *PopGetAITryOnJobResponseBodyDataDummyProjectInfoSourceToken {
	s.SecurityToken = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataMaterialInfo struct {
	Bottoms      *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms `json:"Bottoms,omitempty" xml:"Bottoms,omitempty" type:"Struct"`
	ClothingType *string                                              `json:"ClothingType,omitempty" xml:"ClothingType,omitempty"`
	Model        *PopGetAITryOnJobResponseBodyDataMaterialInfoModel   `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	ShoeType     *string                                              `json:"ShoeType,omitempty" xml:"ShoeType,omitempty"`
	Suit         *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit    `json:"Suit,omitempty" xml:"Suit,omitempty" type:"Struct"`
	Tops         *PopGetAITryOnJobResponseBodyDataMaterialInfoTops    `json:"Tops,omitempty" xml:"Tops,omitempty" type:"Struct"`
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfo) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfo) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfo) SetBottoms(v *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) *PopGetAITryOnJobResponseBodyDataMaterialInfo {
	s.Bottoms = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfo) SetClothingType(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfo {
	s.ClothingType = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfo) SetModel(v *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) *PopGetAITryOnJobResponseBodyDataMaterialInfo {
	s.Model = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfo) SetShoeType(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfo {
	s.ShoeType = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfo) SetSuit(v *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) *PopGetAITryOnJobResponseBodyDataMaterialInfo {
	s.Suit = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfo) SetTops(v *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) *PopGetAITryOnJobResponseBodyDataMaterialInfo {
	s.Tops = v
	return s
}

type PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetCheckStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.CheckStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetCommon(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Common = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetExt(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Ext = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetFileUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.FileUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetId(v int64) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetIntro(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Intro = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetListStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.ListStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetName(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetPreviewUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.PreviewUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms) SetType(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoBottoms {
	s.Type = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataMaterialInfoModel struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoModel) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoModel) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetCheckStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.CheckStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetCommon(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Common = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetExt(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Ext = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetFileUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.FileUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetId(v int64) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetIntro(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Intro = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetListStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.ListStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetName(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetPreviewUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.PreviewUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoModel) SetType(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoModel {
	s.Type = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataMaterialInfoSuit struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetCheckStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.CheckStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetCommon(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Common = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetExt(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Ext = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetFileUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.FileUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetId(v int64) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetIntro(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Intro = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetListStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.ListStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetName(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetPreviewUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.PreviewUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit) SetType(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoSuit {
	s.Type = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataMaterialInfoTops struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoTops) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataMaterialInfoTops) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetCheckStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.CheckStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetCommon(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Common = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetExt(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Ext = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetFileUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.FileUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetId(v int64) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetIntro(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Intro = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetListStatus(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.ListStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetName(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetPreviewUrl(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.PreviewUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataMaterialInfoTops) SetType(v string) *PopGetAITryOnJobResponseBodyDataMaterialInfoTops {
	s.Type = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasks struct {
	Feedback       *PopGetAITryOnJobResponseBodyDataSubTasksFeedback       `json:"Feedback,omitempty" xml:"Feedback,omitempty" type:"Struct"`
	SubProjectInfo *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo `json:"SubProjectInfo,omitempty" xml:"SubProjectInfo,omitempty" type:"Struct"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasks) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasks) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasks) SetFeedback(v *PopGetAITryOnJobResponseBodyDataSubTasksFeedback) *PopGetAITryOnJobResponseBodyDataSubTasks {
	s.Feedback = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasks) SetSubProjectInfo(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) *PopGetAITryOnJobResponseBodyDataSubTasks {
	s.SubProjectInfo = v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksFeedback struct {
	DislikeTags []*int32 `json:"DislikeTags,omitempty" xml:"DislikeTags,omitempty" type:"Repeated"`
	OtherReason *string  `json:"OtherReason,omitempty" xml:"OtherReason,omitempty"`
	ProjectId   *int64   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Rating      *int32   `json:"Rating,omitempty" xml:"Rating,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksFeedback) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksFeedback) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksFeedback) SetDislikeTags(v []*int32) *PopGetAITryOnJobResponseBodyDataSubTasksFeedback {
	s.DislikeTags = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksFeedback) SetOtherReason(v string) *PopGetAITryOnJobResponseBodyDataSubTasksFeedback {
	s.OtherReason = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksFeedback) SetProjectId(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksFeedback {
	s.ProjectId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksFeedback) SetRating(v int32) *PopGetAITryOnJobResponseBodyDataSubTasksFeedback {
	s.Rating = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo struct {
	AuditStatus      *string                                                            `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild        *bool                                                              `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string                                                            `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                            `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                            `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource     *string                                                            `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset          *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                                            `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                                            `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                            `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                            `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source           *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status           *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                            `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetAuditStatus(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.AuditStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetAutoBuild(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.AutoBuild = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetBizUsage(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.BizUsage = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetBuildDetail(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.BuildDetail = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetCheckStatus(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.CheckStatus = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetCreateMode(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.CreateMode = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetCustomSource(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.CustomSource = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetDataset(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Dataset = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetDependencies(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Dependencies = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetExt(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Ext = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetId(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetIntro(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Intro = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetMaterialCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetSource(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Source = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetStatus(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Status = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetTitle(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Title = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo) SetType(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfo {
	s.Type = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetCompletedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetErrorMessage(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetEstimatedDuration(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetId(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetRunningTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetStatus(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.Status = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetSubmitTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset struct {
	BuildResultUrl  map[string]interface{}                                               `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                                              `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                                `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                                              `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	Id              *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string                                                              `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                                              `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                                              `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                                              `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                                              `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetBuildResultUrl(v map[string]interface{}) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetErrorCode(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetErrorMessage(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetGlbModelUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetId(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetModelUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetOriginResultUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetPolicy(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.Policy = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetPoseUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset) SetPreviewUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDataset {
	s.PreviewUrl = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetAccessId(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetDir(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetExpire(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetHost(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetPolicy(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetSignature(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Signature = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource struct {
	Clothes      []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                      `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                    `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                                    `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
	Token        *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken         `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetClothes(v []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.Clothes = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetId(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetPolicy(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.Policy = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetSourceFiles(v []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.SourceFiles = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource) SetToken(v *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSource {
	s.Token = v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes struct {
	CoverUrl     *string                                                                        `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                          `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string                                                                        `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Part         *string                                                                        `json:"Part,omitempty" xml:"Part,omitempty"`
	Size         *string                                                                        `json:"Size,omitempty" xml:"Size,omitempty"`
	SkuProps     []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps `json:"SkuProps,omitempty" xml:"SkuProps,omitempty" type:"Repeated"`
	Skus         []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus     `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	Status       map[string]*string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Version      *int32                                                                         `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetId(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetName(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetPart(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Part = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetSize(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Size = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetSkuProps(v []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.SkuProps = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetSkus(v []*PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Skus = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetStatus(v map[string]*string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Status = v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetType(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Type = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetVersion(v int32) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Version = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Options []*string `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) SetName(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps {
	s.Name = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) SetOptions(v []*string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps {
	s.Options = v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus struct {
	Color  *string `json:"Color,omitempty" xml:"Color,omitempty"`
	Cover  *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	Size   *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetColor(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Color = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetCover(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Cover = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetSize(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Size = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetStatus(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Status = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetAccessId(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetDir(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetExpire(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetHost(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetPolicy(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetSignature(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Signature = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetCoverUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetCreateTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetDeleted(v bool) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetFileName(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetFilesize(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetId(v int64) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Id = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetModifiedTime(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetOssKey(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetType(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetUrl(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Url = &v
	return s
}

type PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Dir             *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Host            *string `json:"Host,omitempty" xml:"Host,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) SetAccessKeyId(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.AccessKeyId = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) SetAccessKeySecret(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.AccessKeySecret = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) SetDir(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.Dir = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) SetExpiration(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.Expiration = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) SetHost(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.Host = &v
	return s
}

func (s *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken) SetSecurityToken(v string) *PopGetAITryOnJobResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.SecurityToken = &v
	return s
}

type PopGetAITryOnJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopGetAITryOnJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopGetAITryOnJobResponse) String() string {
	return tea.Prettify(s)
}

func (s PopGetAITryOnJobResponse) GoString() string {
	return s.String()
}

func (s *PopGetAITryOnJobResponse) SetHeaders(v map[string]*string) *PopGetAITryOnJobResponse {
	s.Headers = v
	return s
}

func (s *PopGetAITryOnJobResponse) SetStatusCode(v int32) *PopGetAITryOnJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PopGetAITryOnJobResponse) SetBody(v *PopGetAITryOnJobResponseBody) *PopGetAITryOnJobResponse {
	s.Body = v
	return s
}

type PopListAITryOnJobsRequest struct {
	Current  *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s PopListAITryOnJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsRequest) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsRequest) SetCurrent(v int32) *PopListAITryOnJobsRequest {
	s.Current = &v
	return s
}

func (s *PopListAITryOnJobsRequest) SetJwtToken(v string) *PopListAITryOnJobsRequest {
	s.JwtToken = &v
	return s
}

func (s *PopListAITryOnJobsRequest) SetSize(v int32) *PopListAITryOnJobsRequest {
	s.Size = &v
	return s
}

type PopListAITryOnJobsResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListAITryOnJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListAITryOnJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBody) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBody) SetCode(v string) *PopListAITryOnJobsResponseBody {
	s.Code = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetCurrent(v int32) *PopListAITryOnJobsResponseBody {
	s.Current = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetData(v []*PopListAITryOnJobsResponseBodyData) *PopListAITryOnJobsResponseBody {
	s.Data = v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetMessage(v string) *PopListAITryOnJobsResponseBody {
	s.Message = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetPages(v int32) *PopListAITryOnJobsResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetRequestId(v string) *PopListAITryOnJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetSize(v int32) *PopListAITryOnJobsResponseBody {
	s.Size = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetSuccess(v bool) *PopListAITryOnJobsResponseBody {
	s.Success = &v
	return s
}

func (s *PopListAITryOnJobsResponseBody) SetTotal(v int32) *PopListAITryOnJobsResponseBody {
	s.Total = &v
	return s
}

type PopListAITryOnJobsResponseBodyData struct {
	DummyProjectInfo *PopListAITryOnJobsResponseBodyDataDummyProjectInfo `json:"DummyProjectInfo,omitempty" xml:"DummyProjectInfo,omitempty" type:"Struct"`
	MaterialInfo     *PopListAITryOnJobsResponseBodyDataMaterialInfo     `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty" type:"Struct"`
	SubTasks         []*PopListAITryOnJobsResponseBodyDataSubTasks       `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
}

func (s PopListAITryOnJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyData) SetDummyProjectInfo(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) *PopListAITryOnJobsResponseBodyData {
	s.DummyProjectInfo = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyData) SetMaterialInfo(v *PopListAITryOnJobsResponseBodyDataMaterialInfo) *PopListAITryOnJobsResponseBodyData {
	s.MaterialInfo = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyData) SetSubTasks(v []*PopListAITryOnJobsResponseBodyDataSubTasks) *PopListAITryOnJobsResponseBodyData {
	s.SubTasks = v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfo struct {
	AuditStatus      *string                                                        `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild        *bool                                                          `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string                                                        `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                        `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                        `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource     *string                                                        `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset          *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                          `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                                        `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                                        `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                        `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                        `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source           *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status           *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfo) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetAuditStatus(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.AuditStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetAutoBuild(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.AutoBuild = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetBizUsage(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.BizUsage = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetBuildDetail(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.BuildDetail = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetCheckStatus(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.CheckStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetCreateMode(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.CreateMode = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetCustomSource(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.CustomSource = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetDataset(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Dataset = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetDependencies(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Dependencies = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetExt(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Ext = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetId(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetIntro(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Intro = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetMaterialCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetSource(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Source = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetStatus(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Status = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetTitle(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Title = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfo) SetType(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfo {
	s.Type = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetCompletedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetErrorMessage(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetEstimatedDuration(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetId(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetRunningTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetStatus(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.Status = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail) SetSubmitTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset struct {
	BuildResultUrl  map[string]interface{}                                           `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                                          `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                            `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                                          `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	Id              *int64                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string                                                          `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                                          `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                                          `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                                          `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                                          `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                                          `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetBuildResultUrl(v map[string]interface{}) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetErrorCode(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetErrorMessage(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetGlbModelUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetId(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetModelUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetOriginResultUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetPolicy(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.Policy = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetPoseUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset) SetPreviewUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDataset {
	s.PreviewUrl = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) SetAccessId(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) SetDir(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) SetExpire(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) SetHost(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) SetPolicy(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy) SetSignature(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoDatasetPolicy {
	s.Signature = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource struct {
	Clothes      []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                                `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
	Token        *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken         `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetClothes(v []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.Clothes = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetId(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetPolicy(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.Policy = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetSourceFiles(v []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.SourceFiles = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource) SetToken(v *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSource {
	s.Token = v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes struct {
	CoverUrl     *string                                                                    `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string                                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                      `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                    `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string                                                                    `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Part         *string                                                                    `json:"Part,omitempty" xml:"Part,omitempty"`
	Size         *string                                                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	SkuProps     []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps `json:"SkuProps,omitempty" xml:"SkuProps,omitempty" type:"Repeated"`
	Skus         []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus     `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	Status       map[string]*string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Version      *int32                                                                     `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetId(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetName(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetPart(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Part = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetSize(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Size = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetSkuProps(v []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.SkuProps = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetSkus(v []*PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Skus = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetStatus(v map[string]*string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Status = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetType(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Type = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes) SetVersion(v int32) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothes {
	s.Version = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Options []*string `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps) SetName(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps) SetOptions(v []*string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkuProps {
	s.Options = v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus struct {
	Color  *string `json:"Color,omitempty" xml:"Color,omitempty"`
	Cover  *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	Size   *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) SetColor(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Color = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) SetCover(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Cover = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) SetSize(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Size = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus) SetStatus(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceClothesSkus {
	s.Status = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) SetAccessId(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) SetDir(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) SetExpire(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) SetHost(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) SetPolicy(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy) SetSignature(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourcePolicy {
	s.Signature = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetFileName(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetFilesize(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetId(v int64) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetType(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles) SetUrl(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceSourceFiles {
	s.Url = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Dir             *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Host            *string `json:"Host,omitempty" xml:"Host,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) SetAccessKeyId(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken {
	s.AccessKeyId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) SetAccessKeySecret(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken {
	s.AccessKeySecret = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) SetDir(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken {
	s.Dir = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) SetExpiration(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken {
	s.Expiration = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) SetHost(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken {
	s.Host = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken) SetSecurityToken(v string) *PopListAITryOnJobsResponseBodyDataDummyProjectInfoSourceToken {
	s.SecurityToken = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataMaterialInfo struct {
	Bottoms      *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms `json:"Bottoms,omitempty" xml:"Bottoms,omitempty" type:"Struct"`
	ClothingType *string                                                `json:"ClothingType,omitempty" xml:"ClothingType,omitempty"`
	Model        *PopListAITryOnJobsResponseBodyDataMaterialInfoModel   `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	ShoeType     *string                                                `json:"ShoeType,omitempty" xml:"ShoeType,omitempty"`
	Suit         *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit    `json:"Suit,omitempty" xml:"Suit,omitempty" type:"Struct"`
	Tops         *PopListAITryOnJobsResponseBodyDataMaterialInfoTops    `json:"Tops,omitempty" xml:"Tops,omitempty" type:"Struct"`
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfo) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfo) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfo) SetBottoms(v *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) *PopListAITryOnJobsResponseBodyDataMaterialInfo {
	s.Bottoms = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfo) SetClothingType(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfo {
	s.ClothingType = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfo) SetModel(v *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) *PopListAITryOnJobsResponseBodyDataMaterialInfo {
	s.Model = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfo) SetShoeType(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfo {
	s.ShoeType = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfo) SetSuit(v *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) *PopListAITryOnJobsResponseBodyDataMaterialInfo {
	s.Suit = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfo) SetTops(v *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) *PopListAITryOnJobsResponseBodyDataMaterialInfo {
	s.Tops = v
	return s
}

type PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetCheckStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.CheckStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetCommon(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Common = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetExt(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Ext = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetFileUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.FileUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetId(v int64) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetIntro(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Intro = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetListStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.ListStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetName(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetPreviewUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.PreviewUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms) SetType(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoBottoms {
	s.Type = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataMaterialInfoModel struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoModel) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoModel) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetCheckStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.CheckStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetCommon(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Common = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetExt(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Ext = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetFileUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.FileUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetId(v int64) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetIntro(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Intro = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetListStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.ListStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetName(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetPreviewUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.PreviewUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoModel) SetType(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoModel {
	s.Type = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataMaterialInfoSuit struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetCheckStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.CheckStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetCommon(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Common = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetExt(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Ext = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetFileUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.FileUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetId(v int64) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetIntro(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Intro = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetListStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.ListStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetName(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetPreviewUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.PreviewUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit) SetType(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoSuit {
	s.Type = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataMaterialInfoTops struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoTops) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataMaterialInfoTops) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetCheckStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.CheckStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetCommon(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Common = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetExt(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Ext = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetFileUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.FileUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetId(v int64) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetIntro(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Intro = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetListStatus(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.ListStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetName(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetPreviewUrl(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.PreviewUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataMaterialInfoTops) SetType(v string) *PopListAITryOnJobsResponseBodyDataMaterialInfoTops {
	s.Type = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasks struct {
	Feedback       *PopListAITryOnJobsResponseBodyDataSubTasksFeedback       `json:"Feedback,omitempty" xml:"Feedback,omitempty" type:"Struct"`
	SubProjectInfo *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo `json:"SubProjectInfo,omitempty" xml:"SubProjectInfo,omitempty" type:"Struct"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasks) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasks) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasks) SetFeedback(v *PopListAITryOnJobsResponseBodyDataSubTasksFeedback) *PopListAITryOnJobsResponseBodyDataSubTasks {
	s.Feedback = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasks) SetSubProjectInfo(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) *PopListAITryOnJobsResponseBodyDataSubTasks {
	s.SubProjectInfo = v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksFeedback struct {
	DislikeTags []*int32 `json:"DislikeTags,omitempty" xml:"DislikeTags,omitempty" type:"Repeated"`
	OtherReason *string  `json:"OtherReason,omitempty" xml:"OtherReason,omitempty"`
	ProjectId   *int64   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Rating      *int32   `json:"Rating,omitempty" xml:"Rating,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksFeedback) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksFeedback) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksFeedback) SetDislikeTags(v []*int32) *PopListAITryOnJobsResponseBodyDataSubTasksFeedback {
	s.DislikeTags = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksFeedback) SetOtherReason(v string) *PopListAITryOnJobsResponseBodyDataSubTasksFeedback {
	s.OtherReason = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksFeedback) SetProjectId(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksFeedback {
	s.ProjectId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksFeedback) SetRating(v int32) *PopListAITryOnJobsResponseBodyDataSubTasksFeedback {
	s.Rating = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo struct {
	AuditStatus      *string                                                              `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild        *bool                                                                `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string                                                              `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                              `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                              `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource     *string                                                              `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset          *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                                `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                                              `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                                              `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                              `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                              `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source           *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status           *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                              `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetAuditStatus(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.AuditStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetAutoBuild(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.AutoBuild = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetBizUsage(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.BizUsage = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetBuildDetail(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.BuildDetail = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetCheckStatus(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.CheckStatus = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetCreateMode(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.CreateMode = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetCustomSource(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.CustomSource = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetDataset(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Dataset = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetDependencies(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Dependencies = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetExt(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Ext = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetId(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetIntro(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Intro = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetMaterialCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetSource(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Source = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetStatus(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Status = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetTitle(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Title = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo) SetType(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfo {
	s.Type = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetCompletedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetErrorMessage(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetEstimatedDuration(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetId(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetRunningTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetStatus(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.Status = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail) SetSubmitTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset struct {
	BuildResultUrl  map[string]interface{}                                                 `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                                                `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                                  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                                                `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	Id              *int64                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl        *string                                                                `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                                                `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                                                `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                                                `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                                                `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetBuildResultUrl(v map[string]interface{}) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetErrorCode(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetErrorMessage(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetGlbModelUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetId(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetModelUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetOriginResultUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetPolicy(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.Policy = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetPoseUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset) SetPreviewUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDataset {
	s.PreviewUrl = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetAccessId(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetDir(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetExpire(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetHost(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetPolicy(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy) SetSignature(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoDatasetPolicy {
	s.Signature = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource struct {
	Clothes      []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                        `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                      `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                                      `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
	Token        *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken         `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetClothes(v []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.Clothes = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetId(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetPolicy(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.Policy = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetSourceFiles(v []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.SourceFiles = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource) SetToken(v *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSource {
	s.Token = v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes struct {
	CoverUrl     *string                                                                          `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string                                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                                            `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Id           *int64                                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string                                                                          `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string                                                                          `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Part         *string                                                                          `json:"Part,omitempty" xml:"Part,omitempty"`
	Size         *string                                                                          `json:"Size,omitempty" xml:"Size,omitempty"`
	SkuProps     []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps `json:"SkuProps,omitempty" xml:"SkuProps,omitempty" type:"Repeated"`
	Skus         []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus     `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	Status       map[string]*string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string                                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Version      *int32                                                                           `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetId(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetName(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetPart(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Part = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetSize(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Size = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetSkuProps(v []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.SkuProps = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetSkus(v []*PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Skus = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetStatus(v map[string]*string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Status = v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetType(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Type = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes) SetVersion(v int32) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothes {
	s.Version = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Options []*string `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) SetName(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps {
	s.Name = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps) SetOptions(v []*string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkuProps {
	s.Options = v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus struct {
	Color  *string `json:"Color,omitempty" xml:"Color,omitempty"`
	Cover  *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	Size   *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetColor(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Color = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetCover(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Cover = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetSize(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Size = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus) SetStatus(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceClothesSkus {
	s.Status = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetAccessId(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetDir(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetExpire(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetHost(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetPolicy(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy) SetSignature(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourcePolicy {
	s.Signature = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetCoverUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetCreateTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetDeleted(v bool) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetFileName(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetFilesize(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetId(v int64) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Id = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetModifiedTime(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetOssKey(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetType(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles) SetUrl(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceSourceFiles {
	s.Url = &v
	return s
}

type PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Dir             *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Host            *string `json:"Host,omitempty" xml:"Host,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) SetAccessKeyId(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.AccessKeyId = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) SetAccessKeySecret(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.AccessKeySecret = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) SetDir(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.Dir = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) SetExpiration(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.Expiration = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) SetHost(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.Host = &v
	return s
}

func (s *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken) SetSecurityToken(v string) *PopListAITryOnJobsResponseBodyDataSubTasksSubProjectInfoSourceToken {
	s.SecurityToken = &v
	return s
}

type PopListAITryOnJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListAITryOnJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListAITryOnJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListAITryOnJobsResponse) GoString() string {
	return s.String()
}

func (s *PopListAITryOnJobsResponse) SetHeaders(v map[string]*string) *PopListAITryOnJobsResponse {
	s.Headers = v
	return s
}

func (s *PopListAITryOnJobsResponse) SetStatusCode(v int32) *PopListAITryOnJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListAITryOnJobsResponse) SetBody(v *PopListAITryOnJobsResponseBody) *PopListAITryOnJobsResponse {
	s.Body = v
	return s
}

type PopListCommonMaterialsAllRequest struct {
	Current    *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ListStatus *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Tags       *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListCommonMaterialsAllRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListCommonMaterialsAllRequest) GoString() string {
	return s.String()
}

func (s *PopListCommonMaterialsAllRequest) SetCurrent(v int32) *PopListCommonMaterialsAllRequest {
	s.Current = &v
	return s
}

func (s *PopListCommonMaterialsAllRequest) SetJwtToken(v string) *PopListCommonMaterialsAllRequest {
	s.JwtToken = &v
	return s
}

func (s *PopListCommonMaterialsAllRequest) SetListStatus(v string) *PopListCommonMaterialsAllRequest {
	s.ListStatus = &v
	return s
}

func (s *PopListCommonMaterialsAllRequest) SetName(v string) *PopListCommonMaterialsAllRequest {
	s.Name = &v
	return s
}

func (s *PopListCommonMaterialsAllRequest) SetSize(v int32) *PopListCommonMaterialsAllRequest {
	s.Size = &v
	return s
}

func (s *PopListCommonMaterialsAllRequest) SetTags(v string) *PopListCommonMaterialsAllRequest {
	s.Tags = &v
	return s
}

func (s *PopListCommonMaterialsAllRequest) SetType(v string) *PopListCommonMaterialsAllRequest {
	s.Type = &v
	return s
}

type PopListCommonMaterialsAllResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PopListCommonMaterialsAllResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopListCommonMaterialsAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListCommonMaterialsAllResponseBody) GoString() string {
	return s.String()
}

func (s *PopListCommonMaterialsAllResponseBody) SetCode(v string) *PopListCommonMaterialsAllResponseBody {
	s.Code = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBody) SetData(v []*PopListCommonMaterialsAllResponseBodyData) *PopListCommonMaterialsAllResponseBody {
	s.Data = v
	return s
}

func (s *PopListCommonMaterialsAllResponseBody) SetMessage(v string) *PopListCommonMaterialsAllResponseBody {
	s.Message = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBody) SetRequestId(v string) *PopListCommonMaterialsAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBody) SetSuccess(v bool) *PopListCommonMaterialsAllResponseBody {
	s.Success = &v
	return s
}

type PopListCommonMaterialsAllResponseBodyData struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common       *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext          *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus   *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PreviewUrl   *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListCommonMaterialsAllResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListCommonMaterialsAllResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetCheckStatus(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetCommon(v bool) *PopListCommonMaterialsAllResponseBodyData {
	s.Common = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetCoverUrl(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetCreateTime(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetDeleted(v bool) *PopListCommonMaterialsAllResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetExt(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetFileUrl(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetId(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetIntro(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetListStatus(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.ListStatus = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetModifiedTime(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetName(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.Name = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetOssKey(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetPreviewUrl(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *PopListCommonMaterialsAllResponseBodyData) SetType(v string) *PopListCommonMaterialsAllResponseBodyData {
	s.Type = &v
	return s
}

type PopListCommonMaterialsAllResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListCommonMaterialsAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListCommonMaterialsAllResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListCommonMaterialsAllResponse) GoString() string {
	return s.String()
}

func (s *PopListCommonMaterialsAllResponse) SetHeaders(v map[string]*string) *PopListCommonMaterialsAllResponse {
	s.Headers = v
	return s
}

func (s *PopListCommonMaterialsAllResponse) SetStatusCode(v int32) *PopListCommonMaterialsAllResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListCommonMaterialsAllResponse) SetBody(v *PopListCommonMaterialsAllResponseBody) *PopListCommonMaterialsAllResponse {
	s.Body = v
	return s
}

type PopListFeatureToAvatarMaterialsRequest struct {
	Current    *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	ListStatus *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Tags       *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s PopListFeatureToAvatarMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarMaterialsRequest) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarMaterialsRequest) SetCurrent(v int32) *PopListFeatureToAvatarMaterialsRequest {
	s.Current = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsRequest) SetListStatus(v string) *PopListFeatureToAvatarMaterialsRequest {
	s.ListStatus = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsRequest) SetSize(v int32) *PopListFeatureToAvatarMaterialsRequest {
	s.Size = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsRequest) SetTags(v string) *PopListFeatureToAvatarMaterialsRequest {
	s.Tags = &v
	return s
}

type PopListFeatureToAvatarMaterialsResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                             `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListFeatureToAvatarMaterialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                             `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                             `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListFeatureToAvatarMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetCode(v string) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Code = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetCurrent(v int32) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Current = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetData(v []*PopListFeatureToAvatarMaterialsResponseBodyData) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Data = v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetMessage(v string) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Message = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetPages(v int32) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetRequestId(v string) *PopListFeatureToAvatarMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetSize(v int32) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Size = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetSuccess(v bool) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Success = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBody) SetTotal(v int32) *PopListFeatureToAvatarMaterialsResponseBody {
	s.Total = &v
	return s
}

type PopListFeatureToAvatarMaterialsResponseBodyData struct {
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Common      *bool   `json:"Common,omitempty" xml:"Common,omitempty"`
	CoverUrl    *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Deleted     *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext         *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl     *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro       *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ListStatus  *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListFeatureToAvatarMaterialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarMaterialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetCheckStatus(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetCommon(v bool) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Common = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetCoverUrl(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetDeleted(v bool) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetExt(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetFileUrl(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetId(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetIntro(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetListStatus(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.ListStatus = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetName(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Name = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponseBodyData) SetType(v string) *PopListFeatureToAvatarMaterialsResponseBodyData {
	s.Type = &v
	return s
}

type PopListFeatureToAvatarMaterialsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListFeatureToAvatarMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListFeatureToAvatarMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarMaterialsResponse) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarMaterialsResponse) SetHeaders(v map[string]*string) *PopListFeatureToAvatarMaterialsResponse {
	s.Headers = v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponse) SetStatusCode(v int32) *PopListFeatureToAvatarMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListFeatureToAvatarMaterialsResponse) SetBody(v *PopListFeatureToAvatarMaterialsResponseBody) *PopListFeatureToAvatarMaterialsResponse {
	s.Body = v
	return s
}

type PopListFeatureToAvatarProjectRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopListFeatureToAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectRequest) SetCurrent(v int32) *PopListFeatureToAvatarProjectRequest {
	s.Current = &v
	return s
}

func (s *PopListFeatureToAvatarProjectRequest) SetSize(v int32) *PopListFeatureToAvatarProjectRequest {
	s.Size = &v
	return s
}

func (s *PopListFeatureToAvatarProjectRequest) SetSortField(v string) *PopListFeatureToAvatarProjectRequest {
	s.SortField = &v
	return s
}

func (s *PopListFeatureToAvatarProjectRequest) SetStatus(v string) *PopListFeatureToAvatarProjectRequest {
	s.Status = &v
	return s
}

func (s *PopListFeatureToAvatarProjectRequest) SetTitle(v string) *PopListFeatureToAvatarProjectRequest {
	s.Title = &v
	return s
}

type PopListFeatureToAvatarProjectResponseBody struct {
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                           `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListFeatureToAvatarProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                           `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListFeatureToAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetCode(v string) *PopListFeatureToAvatarProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetCurrent(v int32) *PopListFeatureToAvatarProjectResponseBody {
	s.Current = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetData(v []*PopListFeatureToAvatarProjectResponseBodyData) *PopListFeatureToAvatarProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetMessage(v string) *PopListFeatureToAvatarProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetPages(v int32) *PopListFeatureToAvatarProjectResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetRequestId(v string) *PopListFeatureToAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetSize(v int32) *PopListFeatureToAvatarProjectResponseBody {
	s.Size = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetSuccess(v bool) *PopListFeatureToAvatarProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBody) SetTotal(v int32) *PopListFeatureToAvatarProjectResponseBody {
	s.Total = &v
	return s
}

type PopListFeatureToAvatarProjectResponseBodyData struct {
	BizUsage         *string                                                   `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                   `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                   `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset          *PopListFeatureToAvatarProjectResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                     `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Ext              *string                                                   `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                   `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                   `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status           *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListFeatureToAvatarProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetBizUsage(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetBuildDetail(v *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) *PopListFeatureToAvatarProjectResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetCheckStatus(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetCreateMode(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetCreateTime(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetDataset(v *PopListFeatureToAvatarProjectResponseBodyDataDataset) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetDeleted(v bool) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetExt(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetId(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetIntro(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetMaterialCoverUrl(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetModifiedTime(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetStatus(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetTitle(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyData) SetType(v string) *PopListFeatureToAvatarProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopListFeatureToAvatarProjectResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetCreateTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetDeleted(v bool) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetRunningTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetStatus(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.Status = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListFeatureToAvatarProjectResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                                      `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                                     `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                       `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                                     `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                                                     `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                                     `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                                     `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                                     `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                                     `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                                     `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopListFeatureToAvatarProjectResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetCoverUrl(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetCreateTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetDeleted(v bool) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetErrorCode(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetErrorMessage(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetGlbModelUrl(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetModelUrl(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetModifiedTime(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetOriginResultUrl(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetOssKey(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetPolicy(v *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetPoseUrl(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDataset) SetPreviewUrl(v string) *PopListFeatureToAvatarProjectResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

type PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetDir(v string) *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetExpire(v string) *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetHost(v string) *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy) SetSignature(v string) *PopListFeatureToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopListFeatureToAvatarProjectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListFeatureToAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListFeatureToAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListFeatureToAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *PopListFeatureToAvatarProjectResponse) SetHeaders(v map[string]*string) *PopListFeatureToAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *PopListFeatureToAvatarProjectResponse) SetStatusCode(v int32) *PopListFeatureToAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListFeatureToAvatarProjectResponse) SetBody(v *PopListFeatureToAvatarProjectResponseBody) *PopListFeatureToAvatarProjectResponse {
	s.Body = v
	return s
}

type PopListLivePortraitModelScopeMaterialsRequest struct {
	Current *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	Size    *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Types   *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s PopListLivePortraitModelScopeMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListLivePortraitModelScopeMaterialsRequest) GoString() string {
	return s.String()
}

func (s *PopListLivePortraitModelScopeMaterialsRequest) SetCurrent(v int32) *PopListLivePortraitModelScopeMaterialsRequest {
	s.Current = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsRequest) SetSize(v int32) *PopListLivePortraitModelScopeMaterialsRequest {
	s.Size = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsRequest) SetTypes(v string) *PopListLivePortraitModelScopeMaterialsRequest {
	s.Types = &v
	return s
}

type PopListLivePortraitModelScopeMaterialsResponseBody struct {
	Code      *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                                    `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListLivePortraitModelScopeMaterialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                                    `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListLivePortraitModelScopeMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListLivePortraitModelScopeMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetCode(v string) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Code = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetCurrent(v int32) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Current = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetData(v []*PopListLivePortraitModelScopeMaterialsResponseBodyData) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Data = v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetMessage(v string) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Message = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetPages(v int32) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetRequestId(v string) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetSize(v int32) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Size = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetSuccess(v bool) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Success = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBody) SetTotal(v int32) *PopListLivePortraitModelScopeMaterialsResponseBody {
	s.Total = &v
	return s
}

type PopListLivePortraitModelScopeMaterialsResponseBodyData struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Ext      *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro    *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListLivePortraitModelScopeMaterialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListLivePortraitModelScopeMaterialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetCoverUrl(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetExt(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetFileUrl(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetId(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetIntro(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetName(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.Name = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponseBodyData) SetType(v string) *PopListLivePortraitModelScopeMaterialsResponseBodyData {
	s.Type = &v
	return s
}

type PopListLivePortraitModelScopeMaterialsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListLivePortraitModelScopeMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListLivePortraitModelScopeMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListLivePortraitModelScopeMaterialsResponse) GoString() string {
	return s.String()
}

func (s *PopListLivePortraitModelScopeMaterialsResponse) SetHeaders(v map[string]*string) *PopListLivePortraitModelScopeMaterialsResponse {
	s.Headers = v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponse) SetStatusCode(v int32) *PopListLivePortraitModelScopeMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListLivePortraitModelScopeMaterialsResponse) SetBody(v *PopListLivePortraitModelScopeMaterialsResponseBody) *PopListLivePortraitModelScopeMaterialsResponse {
	s.Body = v
	return s
}

type PopListObjectCaseRequest struct {
	Current  *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s PopListObjectCaseRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseRequest) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseRequest) SetCurrent(v int32) *PopListObjectCaseRequest {
	s.Current = &v
	return s
}

func (s *PopListObjectCaseRequest) SetJwtToken(v string) *PopListObjectCaseRequest {
	s.JwtToken = &v
	return s
}

func (s *PopListObjectCaseRequest) SetSize(v int32) *PopListObjectCaseRequest {
	s.Size = &v
	return s
}

type PopListObjectCaseResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                               `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListObjectCaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                              `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                               `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                               `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListObjectCaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBody) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBody) SetCode(v string) *PopListObjectCaseResponseBody {
	s.Code = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetCurrent(v int32) *PopListObjectCaseResponseBody {
	s.Current = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetData(v []*PopListObjectCaseResponseBodyData) *PopListObjectCaseResponseBody {
	s.Data = v
	return s
}

func (s *PopListObjectCaseResponseBody) SetErrorName(v string) *PopListObjectCaseResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetHttpCode(v int32) *PopListObjectCaseResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetMessage(v string) *PopListObjectCaseResponseBody {
	s.Message = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetPages(v int32) *PopListObjectCaseResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetRequestId(v string) *PopListObjectCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetSize(v int32) *PopListObjectCaseResponseBody {
	s.Size = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetSuccess(v bool) *PopListObjectCaseResponseBody {
	s.Success = &v
	return s
}

func (s *PopListObjectCaseResponseBody) SetTotal(v int32) *PopListObjectCaseResponseBody {
	s.Total = &v
	return s
}

type PopListObjectCaseResponseBodyData struct {
	AuditStatus  *string                                       `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild    *bool                                         `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage     *string                                       `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail  *PopListObjectCaseResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus  *string                                       `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode   *string                                       `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime   *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource *string                                       `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset      *PopListObjectCaseResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted      *bool                                         `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies *string                                       `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext          *string                                       `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id           *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string                                       `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ModifiedTime *string                                       `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source       *PopListObjectCaseResponseBodyDataSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status       *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string                                       `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListObjectCaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyData) SetAuditStatus(v string) *PopListObjectCaseResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetAutoBuild(v bool) *PopListObjectCaseResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetBizUsage(v string) *PopListObjectCaseResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetBuildDetail(v *PopListObjectCaseResponseBodyDataBuildDetail) *PopListObjectCaseResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetCheckStatus(v string) *PopListObjectCaseResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetCreateMode(v string) *PopListObjectCaseResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetCreateTime(v string) *PopListObjectCaseResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetCustomSource(v string) *PopListObjectCaseResponseBodyData {
	s.CustomSource = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetDataset(v *PopListObjectCaseResponseBodyDataDataset) *PopListObjectCaseResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetDeleted(v bool) *PopListObjectCaseResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetDependencies(v string) *PopListObjectCaseResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetExt(v string) *PopListObjectCaseResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetId(v string) *PopListObjectCaseResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetIntro(v string) *PopListObjectCaseResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetModifiedTime(v string) *PopListObjectCaseResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetSource(v *PopListObjectCaseResponseBodyDataSource) *PopListObjectCaseResponseBodyData {
	s.Source = v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetStatus(v string) *PopListObjectCaseResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetTitle(v string) *PopListObjectCaseResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopListObjectCaseResponseBodyData) SetType(v string) *PopListObjectCaseResponseBodyData {
	s.Type = &v
	return s
}

type PopListObjectCaseResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListObjectCaseResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetCreateTime(v string) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetDeleted(v bool) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetRunningTime(v string) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopListObjectCaseResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListObjectCaseResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                          `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                         `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                           `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage    *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                         `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                                         `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                         `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                         `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                         `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopListObjectCaseResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                         `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                         `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopListObjectCaseResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopListObjectCaseResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetCoverUrl(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetCreateTime(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetDeleted(v bool) *PopListObjectCaseResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetErrorMessage(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetGlbModelUrl(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetModelUrl(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetModifiedTime(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetOriginResultUrl(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetOssKey(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetPolicy(v *PopListObjectCaseResponseBodyDataDatasetPolicy) *PopListObjectCaseResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetPoseUrl(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDataset) SetPreviewUrl(v string) *PopListObjectCaseResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

type PopListObjectCaseResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListObjectCaseResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopListObjectCaseResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDatasetPolicy) SetDir(v string) *PopListObjectCaseResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDatasetPolicy) SetExpire(v string) *PopListObjectCaseResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDatasetPolicy) SetHost(v string) *PopListObjectCaseResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopListObjectCaseResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataDatasetPolicy) SetSignature(v string) *PopListObjectCaseResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopListObjectCaseResponseBodyDataSource struct {
	Clothes      []*PopListObjectCaseResponseBodyDataSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string                                               `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                               `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopListObjectCaseResponseBodyDataSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopListObjectCaseResponseBodyDataSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
}

func (s PopListObjectCaseResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataSource) SetClothes(v []*PopListObjectCaseResponseBodyDataSourceClothes) *PopListObjectCaseResponseBodyDataSource {
	s.Clothes = v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSource) SetCreateTime(v string) *PopListObjectCaseResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSource) SetDeleted(v bool) *PopListObjectCaseResponseBodyDataSource {
	s.Deleted = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSource) SetModifiedTime(v string) *PopListObjectCaseResponseBodyDataSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSource) SetOssKey(v string) *PopListObjectCaseResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSource) SetPolicy(v *PopListObjectCaseResponseBodyDataSourcePolicy) *PopListObjectCaseResponseBodyDataSource {
	s.Policy = v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSource) SetSourceFiles(v []*PopListObjectCaseResponseBodyDataSourceSourceFiles) *PopListObjectCaseResponseBodyDataSource {
	s.SourceFiles = v
	return s
}

type PopListObjectCaseResponseBodyDataSourceClothes struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListObjectCaseResponseBodyDataSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataSourceClothes) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetCoverUrl(v string) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetCreateTime(v string) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetDeleted(v bool) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetModifiedTime(v string) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetName(v string) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.Name = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetOssKey(v string) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceClothes) SetType(v string) *PopListObjectCaseResponseBodyDataSourceClothes {
	s.Type = &v
	return s
}

type PopListObjectCaseResponseBodyDataSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListObjectCaseResponseBodyDataSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataSourcePolicy) SetAccessId(v string) *PopListObjectCaseResponseBodyDataSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourcePolicy) SetDir(v string) *PopListObjectCaseResponseBodyDataSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourcePolicy) SetExpire(v string) *PopListObjectCaseResponseBodyDataSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourcePolicy) SetHost(v string) *PopListObjectCaseResponseBodyDataSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourcePolicy) SetPolicy(v string) *PopListObjectCaseResponseBodyDataSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourcePolicy) SetSignature(v string) *PopListObjectCaseResponseBodyDataSourcePolicy {
	s.Signature = &v
	return s
}

type PopListObjectCaseResponseBodyDataSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopListObjectCaseResponseBodyDataSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponseBodyDataSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetCoverUrl(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetCreateTime(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetDeleted(v bool) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetFileName(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetFilesize(v int64) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetModifiedTime(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetOssKey(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetType(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopListObjectCaseResponseBodyDataSourceSourceFiles) SetUrl(v string) *PopListObjectCaseResponseBodyDataSourceSourceFiles {
	s.Url = &v
	return s
}

type PopListObjectCaseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListObjectCaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListObjectCaseResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectCaseResponse) GoString() string {
	return s.String()
}

func (s *PopListObjectCaseResponse) SetHeaders(v map[string]*string) *PopListObjectCaseResponse {
	s.Headers = v
	return s
}

func (s *PopListObjectCaseResponse) SetStatusCode(v int32) *PopListObjectCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListObjectCaseResponse) SetBody(v *PopListObjectCaseResponseBody) *PopListObjectCaseResponse {
	s.Body = v
	return s
}

type PopListObjectGenerationProjectRequest struct {
	Current  *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s PopListObjectGenerationProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectGenerationProjectRequest) GoString() string {
	return s.String()
}

func (s *PopListObjectGenerationProjectRequest) SetCurrent(v int32) *PopListObjectGenerationProjectRequest {
	s.Current = &v
	return s
}

func (s *PopListObjectGenerationProjectRequest) SetJwtToken(v string) *PopListObjectGenerationProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopListObjectGenerationProjectRequest) SetSize(v int32) *PopListObjectGenerationProjectRequest {
	s.Size = &v
	return s
}

type PopListObjectGenerationProjectResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                            `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListObjectGenerationProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                            `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                            `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListObjectGenerationProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectGenerationProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopListObjectGenerationProjectResponseBody) SetCode(v string) *PopListObjectGenerationProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetCurrent(v int32) *PopListObjectGenerationProjectResponseBody {
	s.Current = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetData(v []*PopListObjectGenerationProjectResponseBodyData) *PopListObjectGenerationProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetMessage(v string) *PopListObjectGenerationProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetPages(v int32) *PopListObjectGenerationProjectResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetRequestId(v string) *PopListObjectGenerationProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetSize(v int32) *PopListObjectGenerationProjectResponseBody {
	s.Size = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetSuccess(v bool) *PopListObjectGenerationProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBody) SetTotal(v int32) *PopListObjectGenerationProjectResponseBody {
	s.Total = &v
	return s
}

type PopListObjectGenerationProjectResponseBodyData struct {
	BizUsage    *string                                                    `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail *PopListObjectGenerationProjectResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	Dataset     *PopListObjectGenerationProjectResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Ext         *string                                                    `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id          *string                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro       *string                                                    `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Status      *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string                                                    `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopListObjectGenerationProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectGenerationProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetBizUsage(v string) *PopListObjectGenerationProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetBuildDetail(v *PopListObjectGenerationProjectResponseBodyDataBuildDetail) *PopListObjectGenerationProjectResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetDataset(v *PopListObjectGenerationProjectResponseBodyDataDataset) *PopListObjectGenerationProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetExt(v string) *PopListObjectGenerationProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetId(v string) *PopListObjectGenerationProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetIntro(v string) *PopListObjectGenerationProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetStatus(v string) *PopListObjectGenerationProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyData) SetTitle(v string) *PopListObjectGenerationProjectResponseBodyData {
	s.Title = &v
	return s
}

type PopListObjectGenerationProjectResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListObjectGenerationProjectResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectGenerationProjectResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListObjectGenerationProjectResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopListObjectGenerationProjectResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopListObjectGenerationProjectResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopListObjectGenerationProjectResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyDataBuildDetail) SetRunningTime(v string) *PopListObjectGenerationProjectResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListObjectGenerationProjectResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopListObjectGenerationProjectResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListObjectGenerationProjectResponseBodyDataDataset struct {
	BuildResultUrl map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
}

func (s PopListObjectGenerationProjectResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectGenerationProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopListObjectGenerationProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopListObjectGenerationProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

type PopListObjectGenerationProjectResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListObjectGenerationProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListObjectGenerationProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectGenerationProjectResponse) GoString() string {
	return s.String()
}

func (s *PopListObjectGenerationProjectResponse) SetHeaders(v map[string]*string) *PopListObjectGenerationProjectResponse {
	s.Headers = v
	return s
}

func (s *PopListObjectGenerationProjectResponse) SetStatusCode(v int32) *PopListObjectGenerationProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListObjectGenerationProjectResponse) SetBody(v *PopListObjectGenerationProjectResponseBody) *PopListObjectGenerationProjectResponse {
	s.Body = v
	return s
}

type PopListObjectProjectRequest struct {
	AuditStatus  *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	Current      *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	CustomSource *string `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	JwtToken     *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SortField    *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	WithSource   *bool   `json:"WithSource,omitempty" xml:"WithSource,omitempty"`
}

func (s PopListObjectProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectRequest) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectRequest) SetAuditStatus(v string) *PopListObjectProjectRequest {
	s.AuditStatus = &v
	return s
}

func (s *PopListObjectProjectRequest) SetCurrent(v int32) *PopListObjectProjectRequest {
	s.Current = &v
	return s
}

func (s *PopListObjectProjectRequest) SetCustomSource(v string) *PopListObjectProjectRequest {
	s.CustomSource = &v
	return s
}

func (s *PopListObjectProjectRequest) SetJwtToken(v string) *PopListObjectProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopListObjectProjectRequest) SetSize(v int32) *PopListObjectProjectRequest {
	s.Size = &v
	return s
}

func (s *PopListObjectProjectRequest) SetSortField(v string) *PopListObjectProjectRequest {
	s.SortField = &v
	return s
}

func (s *PopListObjectProjectRequest) SetStatus(v string) *PopListObjectProjectRequest {
	s.Status = &v
	return s
}

func (s *PopListObjectProjectRequest) SetTitle(v string) *PopListObjectProjectRequest {
	s.Title = &v
	return s
}

func (s *PopListObjectProjectRequest) SetWithSource(v bool) *PopListObjectProjectRequest {
	s.WithSource = &v
	return s
}

type PopListObjectProjectResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                  `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListObjectProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                                 `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                  `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListObjectProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBody) SetCode(v string) *PopListObjectProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetCurrent(v int32) *PopListObjectProjectResponseBody {
	s.Current = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetData(v []*PopListObjectProjectResponseBodyData) *PopListObjectProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopListObjectProjectResponseBody) SetErrorName(v string) *PopListObjectProjectResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetHttpCode(v int32) *PopListObjectProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetMessage(v string) *PopListObjectProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetPages(v int32) *PopListObjectProjectResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetRequestId(v string) *PopListObjectProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetSize(v int32) *PopListObjectProjectResponseBody {
	s.Size = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetSuccess(v bool) *PopListObjectProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PopListObjectProjectResponseBody) SetTotal(v int32) *PopListObjectProjectResponseBody {
	s.Total = &v
	return s
}

type PopListObjectProjectResponseBodyData struct {
	AuditStatus  *string                                          `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoBuild    *bool                                            `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage     *string                                          `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail  *PopListObjectProjectResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus  *string                                          `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode   *string                                          `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime   *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomSource *string                                          `json:"CustomSource,omitempty" xml:"CustomSource,omitempty"`
	Dataset      *PopListObjectProjectResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted      *bool                                            `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies *string                                          `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext          *string                                          `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id           *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string                                          `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ModifiedTime *string                                          `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source       *PopListObjectProjectResponseBodyDataSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status       *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListObjectProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyData) SetAuditStatus(v string) *PopListObjectProjectResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetAutoBuild(v bool) *PopListObjectProjectResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetBizUsage(v string) *PopListObjectProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetBuildDetail(v *PopListObjectProjectResponseBodyDataBuildDetail) *PopListObjectProjectResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetCheckStatus(v string) *PopListObjectProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetCreateMode(v string) *PopListObjectProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetCreateTime(v string) *PopListObjectProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetCustomSource(v string) *PopListObjectProjectResponseBodyData {
	s.CustomSource = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetDataset(v *PopListObjectProjectResponseBodyDataDataset) *PopListObjectProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetDeleted(v bool) *PopListObjectProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetDependencies(v string) *PopListObjectProjectResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetExt(v string) *PopListObjectProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetId(v string) *PopListObjectProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetIntro(v string) *PopListObjectProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetModifiedTime(v string) *PopListObjectProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetSource(v *PopListObjectProjectResponseBodyDataSource) *PopListObjectProjectResponseBodyData {
	s.Source = v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetStatus(v string) *PopListObjectProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetTitle(v string) *PopListObjectProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopListObjectProjectResponseBodyData) SetType(v string) *PopListObjectProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopListObjectProjectResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListObjectProjectResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetCreateTime(v string) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetDeleted(v bool) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetRunningTime(v string) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopListObjectProjectResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListObjectProjectResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                             `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                            `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage    *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                            `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                                            `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                            `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                            `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopListObjectProjectResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                            `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                            `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopListObjectProjectResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopListObjectProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetCoverUrl(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetCreateTime(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetDeleted(v bool) *PopListObjectProjectResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetErrorMessage(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetGlbModelUrl(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetModelUrl(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetModifiedTime(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetOriginResultUrl(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetOssKey(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetPolicy(v *PopListObjectProjectResponseBodyDataDatasetPolicy) *PopListObjectProjectResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetPoseUrl(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDataset) SetPreviewUrl(v string) *PopListObjectProjectResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

type PopListObjectProjectResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListObjectProjectResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopListObjectProjectResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDatasetPolicy) SetDir(v string) *PopListObjectProjectResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDatasetPolicy) SetExpire(v string) *PopListObjectProjectResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDatasetPolicy) SetHost(v string) *PopListObjectProjectResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopListObjectProjectResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataDatasetPolicy) SetSignature(v string) *PopListObjectProjectResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopListObjectProjectResponseBodyDataSource struct {
	Clothes      []*PopListObjectProjectResponseBodyDataSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                    `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string                                                  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                  `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopListObjectProjectResponseBodyDataSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopListObjectProjectResponseBodyDataSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
}

func (s PopListObjectProjectResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataSource) SetClothes(v []*PopListObjectProjectResponseBodyDataSourceClothes) *PopListObjectProjectResponseBodyDataSource {
	s.Clothes = v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSource) SetCreateTime(v string) *PopListObjectProjectResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSource) SetDeleted(v bool) *PopListObjectProjectResponseBodyDataSource {
	s.Deleted = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSource) SetModifiedTime(v string) *PopListObjectProjectResponseBodyDataSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSource) SetOssKey(v string) *PopListObjectProjectResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSource) SetPolicy(v *PopListObjectProjectResponseBodyDataSourcePolicy) *PopListObjectProjectResponseBodyDataSource {
	s.Policy = v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSource) SetSourceFiles(v []*PopListObjectProjectResponseBodyDataSourceSourceFiles) *PopListObjectProjectResponseBodyDataSource {
	s.SourceFiles = v
	return s
}

type PopListObjectProjectResponseBodyDataSourceClothes struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListObjectProjectResponseBodyDataSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataSourceClothes) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetCoverUrl(v string) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetCreateTime(v string) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetDeleted(v bool) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetModifiedTime(v string) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetName(v string) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.Name = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetOssKey(v string) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceClothes) SetType(v string) *PopListObjectProjectResponseBodyDataSourceClothes {
	s.Type = &v
	return s
}

type PopListObjectProjectResponseBodyDataSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListObjectProjectResponseBodyDataSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataSourcePolicy) SetAccessId(v string) *PopListObjectProjectResponseBodyDataSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourcePolicy) SetDir(v string) *PopListObjectProjectResponseBodyDataSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourcePolicy) SetExpire(v string) *PopListObjectProjectResponseBodyDataSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourcePolicy) SetHost(v string) *PopListObjectProjectResponseBodyDataSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourcePolicy) SetPolicy(v string) *PopListObjectProjectResponseBodyDataSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourcePolicy) SetSignature(v string) *PopListObjectProjectResponseBodyDataSourcePolicy {
	s.Signature = &v
	return s
}

type PopListObjectProjectResponseBodyDataSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopListObjectProjectResponseBodyDataSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponseBodyDataSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetCoverUrl(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetCreateTime(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetDeleted(v bool) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetFileName(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetFilesize(v int64) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetModifiedTime(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetOssKey(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetType(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopListObjectProjectResponseBodyDataSourceSourceFiles) SetUrl(v string) *PopListObjectProjectResponseBodyDataSourceSourceFiles {
	s.Url = &v
	return s
}

type PopListObjectProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListObjectProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListObjectProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListObjectProjectResponse) GoString() string {
	return s.String()
}

func (s *PopListObjectProjectResponse) SetHeaders(v map[string]*string) *PopListObjectProjectResponse {
	s.Headers = v
	return s
}

func (s *PopListObjectProjectResponse) SetStatusCode(v int32) *PopListObjectProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListObjectProjectResponse) SetBody(v *PopListObjectProjectResponseBody) *PopListObjectProjectResponse {
	s.Body = v
	return s
}

type PopListPakRenderExpressionRequest struct {
	Current    *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	ListStatus *string `json:"ListStatus,omitempty" xml:"ListStatus,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s PopListPakRenderExpressionRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListPakRenderExpressionRequest) GoString() string {
	return s.String()
}

func (s *PopListPakRenderExpressionRequest) SetCurrent(v int32) *PopListPakRenderExpressionRequest {
	s.Current = &v
	return s
}

func (s *PopListPakRenderExpressionRequest) SetListStatus(v string) *PopListPakRenderExpressionRequest {
	s.ListStatus = &v
	return s
}

func (s *PopListPakRenderExpressionRequest) SetSize(v int32) *PopListPakRenderExpressionRequest {
	s.Size = &v
	return s
}

type PopListPakRenderExpressionResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                        `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListPakRenderExpressionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                        `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListPakRenderExpressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListPakRenderExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *PopListPakRenderExpressionResponseBody) SetCode(v string) *PopListPakRenderExpressionResponseBody {
	s.Code = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetCurrent(v int32) *PopListPakRenderExpressionResponseBody {
	s.Current = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetData(v []*PopListPakRenderExpressionResponseBodyData) *PopListPakRenderExpressionResponseBody {
	s.Data = v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetMessage(v string) *PopListPakRenderExpressionResponseBody {
	s.Message = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetPages(v int32) *PopListPakRenderExpressionResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetRequestId(v string) *PopListPakRenderExpressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetSize(v int32) *PopListPakRenderExpressionResponseBody {
	s.Size = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetSuccess(v bool) *PopListPakRenderExpressionResponseBody {
	s.Success = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBody) SetTotal(v int32) *PopListPakRenderExpressionResponseBody {
	s.Total = &v
	return s
}

type PopListPakRenderExpressionResponseBodyData struct {
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Ext      *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro    *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PopListPakRenderExpressionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListPakRenderExpressionResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListPakRenderExpressionResponseBodyData) SetCoverUrl(v string) *PopListPakRenderExpressionResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBodyData) SetExt(v string) *PopListPakRenderExpressionResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBodyData) SetFileUrl(v string) *PopListPakRenderExpressionResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBodyData) SetId(v string) *PopListPakRenderExpressionResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBodyData) SetIntro(v string) *PopListPakRenderExpressionResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListPakRenderExpressionResponseBodyData) SetName(v string) *PopListPakRenderExpressionResponseBodyData {
	s.Name = &v
	return s
}

type PopListPakRenderExpressionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListPakRenderExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListPakRenderExpressionResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListPakRenderExpressionResponse) GoString() string {
	return s.String()
}

func (s *PopListPakRenderExpressionResponse) SetHeaders(v map[string]*string) *PopListPakRenderExpressionResponse {
	s.Headers = v
	return s
}

func (s *PopListPakRenderExpressionResponse) SetStatusCode(v int32) *PopListPakRenderExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListPakRenderExpressionResponse) SetBody(v *PopListPakRenderExpressionResponseBody) *PopListPakRenderExpressionResponse {
	s.Body = v
	return s
}

type PopListTextToAvatarProjectRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopListTextToAvatarProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectRequest) SetCurrent(v int32) *PopListTextToAvatarProjectRequest {
	s.Current = &v
	return s
}

func (s *PopListTextToAvatarProjectRequest) SetJwtToken(v string) *PopListTextToAvatarProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PopListTextToAvatarProjectRequest) SetSize(v int32) *PopListTextToAvatarProjectRequest {
	s.Size = &v
	return s
}

func (s *PopListTextToAvatarProjectRequest) SetSortField(v string) *PopListTextToAvatarProjectRequest {
	s.SortField = &v
	return s
}

func (s *PopListTextToAvatarProjectRequest) SetStatus(v string) *PopListTextToAvatarProjectRequest {
	s.Status = &v
	return s
}

func (s *PopListTextToAvatarProjectRequest) SetTitle(v string) *PopListTextToAvatarProjectRequest {
	s.Title = &v
	return s
}

type PopListTextToAvatarProjectResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                                        `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      []*PopListTextToAvatarProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                                        `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PopListTextToAvatarProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectResponseBody) SetCode(v string) *PopListTextToAvatarProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetCurrent(v int32) *PopListTextToAvatarProjectResponseBody {
	s.Current = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetData(v []*PopListTextToAvatarProjectResponseBodyData) *PopListTextToAvatarProjectResponseBody {
	s.Data = v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetMessage(v string) *PopListTextToAvatarProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetPages(v int32) *PopListTextToAvatarProjectResponseBody {
	s.Pages = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetRequestId(v string) *PopListTextToAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetSize(v int32) *PopListTextToAvatarProjectResponseBody {
	s.Size = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetSuccess(v bool) *PopListTextToAvatarProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBody) SetTotal(v int32) *PopListTextToAvatarProjectResponseBody {
	s.Total = &v
	return s
}

type PopListTextToAvatarProjectResponseBodyData struct {
	AutoBuild        *bool                                                  `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string                                                `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopListTextToAvatarProjectResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset          *PopListTextToAvatarProjectResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                                `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                                `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status           *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopListTextToAvatarProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetAutoBuild(v bool) *PopListTextToAvatarProjectResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetBizUsage(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetBuildDetail(v *PopListTextToAvatarProjectResponseBodyDataBuildDetail) *PopListTextToAvatarProjectResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetCheckStatus(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetCreateMode(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetCreateTime(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetDataset(v *PopListTextToAvatarProjectResponseBodyDataDataset) *PopListTextToAvatarProjectResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetDeleted(v bool) *PopListTextToAvatarProjectResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetDependencies(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetExt(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetId(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetIntro(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetMaterialCoverUrl(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetModifiedTime(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetStatus(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetTitle(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyData) SetType(v string) *PopListTextToAvatarProjectResponseBodyData {
	s.Type = &v
	return s
}

type PopListTextToAvatarProjectResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopListTextToAvatarProjectResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetCreateTime(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetDeleted(v bool) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetRunningTime(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetStatus(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.Status = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopListTextToAvatarProjectResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopListTextToAvatarProjectResponseBodyDataDataset struct {
	BuildResultUrl map[string]interface{}                                   `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CreateTime     *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted        *bool                                                    `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode      *string                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ModifiedTime   *string                                                  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey         *string                                                  `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy         *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s PopListTextToAvatarProjectResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetCreateTime(v string) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetDeleted(v bool) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetErrorCode(v string) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetErrorMessage(v string) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetModifiedTime(v string) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetOssKey(v string) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDataset) SetPolicy(v *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) *PopListTextToAvatarProjectResponseBodyDataDataset {
	s.Policy = v
	return s
}

type PopListTextToAvatarProjectResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) SetDir(v string) *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) SetExpire(v string) *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) SetHost(v string) *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy) SetSignature(v string) *PopListTextToAvatarProjectResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopListTextToAvatarProjectResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopListTextToAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopListTextToAvatarProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PopListTextToAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *PopListTextToAvatarProjectResponse) SetHeaders(v map[string]*string) *PopListTextToAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *PopListTextToAvatarProjectResponse) SetStatusCode(v int32) *PopListTextToAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PopListTextToAvatarProjectResponse) SetBody(v *PopListTextToAvatarProjectResponseBody) *PopListTextToAvatarProjectResponse {
	s.Body = v
	return s
}

type PopObjectProjectDetailRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopObjectProjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailRequest) SetJwtToken(v string) *PopObjectProjectDetailRequest {
	s.JwtToken = &v
	return s
}

func (s *PopObjectProjectDetailRequest) SetProjectId(v string) *PopObjectProjectDetailRequest {
	s.ProjectId = &v
	return s
}

type PopObjectProjectDetailResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopObjectProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                                 `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopObjectProjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBody) SetCode(v string) *PopObjectProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *PopObjectProjectDetailResponseBody) SetData(v *PopObjectProjectDetailResponseBodyData) *PopObjectProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *PopObjectProjectDetailResponseBody) SetErrorName(v string) *PopObjectProjectDetailResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopObjectProjectDetailResponseBody) SetHttpCode(v int32) *PopObjectProjectDetailResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopObjectProjectDetailResponseBody) SetMessage(v string) *PopObjectProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *PopObjectProjectDetailResponseBody) SetRequestId(v string) *PopObjectProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopObjectProjectDetailResponseBody) SetSuccess(v bool) *PopObjectProjectDetailResponseBody {
	s.Success = &v
	return s
}

type PopObjectProjectDetailResponseBodyData struct {
	AutoBuild    *bool                                              `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage     *string                                            `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail  *PopObjectProjectDetailResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus  *string                                            `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode   *string                                            `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime   *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset      *PopObjectProjectDetailResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted      *bool                                              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies *string                                            `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext          *string                                            `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id           *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro        *string                                            `json:"Intro,omitempty" xml:"Intro,omitempty"`
	ModifiedTime *string                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Source       *PopObjectProjectDetailResponseBodyDataSource      `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status       *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string                                            `json:"Title,omitempty" xml:"Title,omitempty"`
	Type         *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyData) SetAutoBuild(v bool) *PopObjectProjectDetailResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetBizUsage(v string) *PopObjectProjectDetailResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetBuildDetail(v *PopObjectProjectDetailResponseBodyDataBuildDetail) *PopObjectProjectDetailResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetCheckStatus(v string) *PopObjectProjectDetailResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetCreateMode(v string) *PopObjectProjectDetailResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetCreateTime(v string) *PopObjectProjectDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetDataset(v *PopObjectProjectDetailResponseBodyDataDataset) *PopObjectProjectDetailResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetDeleted(v bool) *PopObjectProjectDetailResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetDependencies(v string) *PopObjectProjectDetailResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetExt(v string) *PopObjectProjectDetailResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetId(v string) *PopObjectProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetIntro(v string) *PopObjectProjectDetailResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetModifiedTime(v string) *PopObjectProjectDetailResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetSource(v *PopObjectProjectDetailResponseBodyDataSource) *PopObjectProjectDetailResponseBodyData {
	s.Source = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetStatus(v string) *PopObjectProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetTitle(v string) *PopObjectProjectDetailResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyData) SetType(v string) *PopObjectProjectDetailResponseBodyData {
	s.Type = &v
	return s
}

type PopObjectProjectDetailResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetCreateTime(v string) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetDeleted(v bool) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetRunningTime(v string) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopObjectProjectDetailResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopObjectProjectDetailResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{}                               `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                                              `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                                                `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage    *string                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                                              `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                                              `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                                              `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                                              `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy          *PopObjectProjectDetailResponseBodyDataDatasetPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	PoseUrl         *string                                              `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                                              `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopObjectProjectDetailResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetCoverUrl(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetCreateTime(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetDeleted(v bool) *PopObjectProjectDetailResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetErrorMessage(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetGlbModelUrl(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetModelUrl(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetModifiedTime(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetOriginResultUrl(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetOssKey(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetPolicy(v *PopObjectProjectDetailResponseBodyDataDatasetPolicy) *PopObjectProjectDetailResponseBodyDataDataset {
	s.Policy = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetPoseUrl(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDataset) SetPreviewUrl(v string) *PopObjectProjectDetailResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

type PopObjectProjectDetailResponseBodyDataDatasetPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyDataDatasetPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataDatasetPolicy) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataDatasetPolicy) SetAccessId(v string) *PopObjectProjectDetailResponseBodyDataDatasetPolicy {
	s.AccessId = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDatasetPolicy) SetDir(v string) *PopObjectProjectDetailResponseBodyDataDatasetPolicy {
	s.Dir = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDatasetPolicy) SetExpire(v string) *PopObjectProjectDetailResponseBodyDataDatasetPolicy {
	s.Expire = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDatasetPolicy) SetHost(v string) *PopObjectProjectDetailResponseBodyDataDatasetPolicy {
	s.Host = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDatasetPolicy) SetPolicy(v string) *PopObjectProjectDetailResponseBodyDataDatasetPolicy {
	s.Policy = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataDatasetPolicy) SetSignature(v string) *PopObjectProjectDetailResponseBodyDataDatasetPolicy {
	s.Signature = &v
	return s
}

type PopObjectProjectDetailResponseBodyDataSource struct {
	Clothes      []*PopObjectProjectDetailResponseBodyDataSourceClothes     `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Repeated"`
	CreateTime   *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool                                                      `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string                                                    `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string                                                    `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy       *PopObjectProjectDetailResponseBodyDataSourcePolicy        `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	SourceFiles  []*PopObjectProjectDetailResponseBodyDataSourceSourceFiles `json:"SourceFiles,omitempty" xml:"SourceFiles,omitempty" type:"Repeated"`
}

func (s PopObjectProjectDetailResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetClothes(v []*PopObjectProjectDetailResponseBodyDataSourceClothes) *PopObjectProjectDetailResponseBodyDataSource {
	s.Clothes = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetCreateTime(v string) *PopObjectProjectDetailResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetDeleted(v bool) *PopObjectProjectDetailResponseBodyDataSource {
	s.Deleted = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetModifiedTime(v string) *PopObjectProjectDetailResponseBodyDataSource {
	s.ModifiedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetOssKey(v string) *PopObjectProjectDetailResponseBodyDataSource {
	s.OssKey = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetPolicy(v *PopObjectProjectDetailResponseBodyDataSourcePolicy) *PopObjectProjectDetailResponseBodyDataSource {
	s.Policy = v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSource) SetSourceFiles(v []*PopObjectProjectDetailResponseBodyDataSourceSourceFiles) *PopObjectProjectDetailResponseBodyDataSource {
	s.SourceFiles = v
	return s
}

type PopObjectProjectDetailResponseBodyDataSourceClothes struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyDataSourceClothes) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataSourceClothes) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetCoverUrl(v string) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.CoverUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetCreateTime(v string) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.CreateTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetDeleted(v bool) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.Deleted = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetModifiedTime(v string) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.ModifiedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetName(v string) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.Name = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetOssKey(v string) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.OssKey = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceClothes) SetType(v string) *PopObjectProjectDetailResponseBodyDataSourceClothes {
	s.Type = &v
	return s
}

type PopObjectProjectDetailResponseBodyDataSourcePolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyDataSourcePolicy) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataSourcePolicy) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataSourcePolicy) SetAccessId(v string) *PopObjectProjectDetailResponseBodyDataSourcePolicy {
	s.AccessId = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourcePolicy) SetDir(v string) *PopObjectProjectDetailResponseBodyDataSourcePolicy {
	s.Dir = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourcePolicy) SetExpire(v string) *PopObjectProjectDetailResponseBodyDataSourcePolicy {
	s.Expire = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourcePolicy) SetHost(v string) *PopObjectProjectDetailResponseBodyDataSourcePolicy {
	s.Host = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourcePolicy) SetPolicy(v string) *PopObjectProjectDetailResponseBodyDataSourcePolicy {
	s.Policy = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourcePolicy) SetSignature(v string) *PopObjectProjectDetailResponseBodyDataSourcePolicy {
	s.Signature = &v
	return s
}

type PopObjectProjectDetailResponseBodyDataSourceSourceFiles struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Filesize     *int64  `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PopObjectProjectDetailResponseBodyDataSourceSourceFiles) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponseBodyDataSourceSourceFiles) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetCoverUrl(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.CoverUrl = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetCreateTime(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.CreateTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetDeleted(v bool) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.Deleted = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetFileName(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.FileName = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetFilesize(v int64) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.Filesize = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetModifiedTime(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.ModifiedTime = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetOssKey(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.OssKey = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetType(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.Type = &v
	return s
}

func (s *PopObjectProjectDetailResponseBodyDataSourceSourceFiles) SetUrl(v string) *PopObjectProjectDetailResponseBodyDataSourceSourceFiles {
	s.Url = &v
	return s
}

type PopObjectProjectDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopObjectProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopObjectProjectDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PopObjectProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *PopObjectProjectDetailResponse) SetHeaders(v map[string]*string) *PopObjectProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *PopObjectProjectDetailResponse) SetStatusCode(v int32) *PopObjectProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *PopObjectProjectDetailResponse) SetBody(v *PopObjectProjectDetailResponseBody) *PopObjectProjectDetailResponse {
	s.Body = v
	return s
}

type PopObjectRetrievalRequest struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TopK       *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s PopObjectRetrievalRequest) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalRequest) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalRequest) SetContent(v string) *PopObjectRetrievalRequest {
	s.Content = &v
	return s
}

func (s *PopObjectRetrievalRequest) SetJwtToken(v string) *PopObjectRetrievalRequest {
	s.JwtToken = &v
	return s
}

func (s *PopObjectRetrievalRequest) SetSourceType(v string) *PopObjectRetrievalRequest {
	s.SourceType = &v
	return s
}

func (s *PopObjectRetrievalRequest) SetTopK(v int32) *PopObjectRetrievalRequest {
	s.TopK = &v
	return s
}

type PopObjectRetrievalResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*PopObjectRetrievalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopObjectRetrievalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalResponseBody) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalResponseBody) SetCode(v string) *PopObjectRetrievalResponseBody {
	s.Code = &v
	return s
}

func (s *PopObjectRetrievalResponseBody) SetData(v []*PopObjectRetrievalResponseBodyData) *PopObjectRetrievalResponseBody {
	s.Data = v
	return s
}

func (s *PopObjectRetrievalResponseBody) SetMessage(v string) *PopObjectRetrievalResponseBody {
	s.Message = &v
	return s
}

func (s *PopObjectRetrievalResponseBody) SetRequestId(v string) *PopObjectRetrievalResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopObjectRetrievalResponseBody) SetSuccess(v bool) *PopObjectRetrievalResponseBody {
	s.Success = &v
	return s
}

type PopObjectRetrievalResponseBodyData struct {
	CoverUrl   *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	GlbUrl     *string `json:"GlbUrl,omitempty" xml:"GlbUrl,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelUrl   *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopObjectRetrievalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalResponseBodyData) SetCoverUrl(v string) *PopObjectRetrievalResponseBodyData {
	s.CoverUrl = &v
	return s
}

func (s *PopObjectRetrievalResponseBodyData) SetGlbUrl(v string) *PopObjectRetrievalResponseBodyData {
	s.GlbUrl = &v
	return s
}

func (s *PopObjectRetrievalResponseBodyData) SetId(v string) *PopObjectRetrievalResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopObjectRetrievalResponseBodyData) SetModelUrl(v string) *PopObjectRetrievalResponseBodyData {
	s.ModelUrl = &v
	return s
}

func (s *PopObjectRetrievalResponseBodyData) SetPreviewUrl(v string) *PopObjectRetrievalResponseBodyData {
	s.PreviewUrl = &v
	return s
}

func (s *PopObjectRetrievalResponseBodyData) SetTitle(v string) *PopObjectRetrievalResponseBodyData {
	s.Title = &v
	return s
}

type PopObjectRetrievalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopObjectRetrievalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopObjectRetrievalResponse) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalResponse) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalResponse) SetHeaders(v map[string]*string) *PopObjectRetrievalResponse {
	s.Headers = v
	return s
}

func (s *PopObjectRetrievalResponse) SetStatusCode(v int32) *PopObjectRetrievalResponse {
	s.StatusCode = &v
	return s
}

func (s *PopObjectRetrievalResponse) SetBody(v *PopObjectRetrievalResponseBody) *PopObjectRetrievalResponse {
	s.Body = v
	return s
}

type PopObjectRetrievalUploadDataRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s PopObjectRetrievalUploadDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalUploadDataRequest) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalUploadDataRequest) SetJwtToken(v string) *PopObjectRetrievalUploadDataRequest {
	s.JwtToken = &v
	return s
}

type PopObjectRetrievalUploadDataResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopObjectRetrievalUploadDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopObjectRetrievalUploadDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalUploadDataResponseBody) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalUploadDataResponseBody) SetCode(v string) *PopObjectRetrievalUploadDataResponseBody {
	s.Code = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBody) SetData(v *PopObjectRetrievalUploadDataResponseBodyData) *PopObjectRetrievalUploadDataResponseBody {
	s.Data = v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBody) SetMessage(v string) *PopObjectRetrievalUploadDataResponseBody {
	s.Message = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBody) SetRequestId(v string) *PopObjectRetrievalUploadDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBody) SetSuccess(v bool) *PopObjectRetrievalUploadDataResponseBody {
	s.Success = &v
	return s
}

type PopObjectRetrievalUploadDataResponseBodyData struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopObjectRetrievalUploadDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalUploadDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalUploadDataResponseBodyData) SetAccessId(v string) *PopObjectRetrievalUploadDataResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBodyData) SetDir(v string) *PopObjectRetrievalUploadDataResponseBodyData {
	s.Dir = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBodyData) SetExpire(v string) *PopObjectRetrievalUploadDataResponseBodyData {
	s.Expire = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBodyData) SetHost(v string) *PopObjectRetrievalUploadDataResponseBodyData {
	s.Host = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBodyData) SetPolicy(v string) *PopObjectRetrievalUploadDataResponseBodyData {
	s.Policy = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponseBodyData) SetSignature(v string) *PopObjectRetrievalUploadDataResponseBodyData {
	s.Signature = &v
	return s
}

type PopObjectRetrievalUploadDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopObjectRetrievalUploadDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopObjectRetrievalUploadDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PopObjectRetrievalUploadDataResponse) GoString() string {
	return s.String()
}

func (s *PopObjectRetrievalUploadDataResponse) SetHeaders(v map[string]*string) *PopObjectRetrievalUploadDataResponse {
	s.Headers = v
	return s
}

func (s *PopObjectRetrievalUploadDataResponse) SetStatusCode(v int32) *PopObjectRetrievalUploadDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PopObjectRetrievalUploadDataResponse) SetBody(v *PopObjectRetrievalUploadDataResponseBody) *PopObjectRetrievalUploadDataResponse {
	s.Body = v
	return s
}

type PopQueryAvatarProjectDetailRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopQueryAvatarProjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PopQueryAvatarProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *PopQueryAvatarProjectDetailRequest) SetProjectId(v string) *PopQueryAvatarProjectDetailRequest {
	s.ProjectId = &v
	return s
}

type PopQueryAvatarProjectDetailResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopQueryAvatarProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopQueryAvatarProjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopQueryAvatarProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *PopQueryAvatarProjectDetailResponseBody) SetCode(v string) *PopQueryAvatarProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBody) SetData(v *PopQueryAvatarProjectDetailResponseBodyData) *PopQueryAvatarProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBody) SetMessage(v string) *PopQueryAvatarProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBody) SetRequestId(v string) *PopQueryAvatarProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBody) SetSuccess(v bool) *PopQueryAvatarProjectDetailResponseBody {
	s.Success = &v
	return s
}

type PopQueryAvatarProjectDetailResponseBodyData struct {
	AutoBuild        *bool                                                   `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	BizUsage         *string                                                 `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail      *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CheckStatus      *string                                                 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CreateMode       *string                                                 `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreateTime       *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset          *PopQueryAvatarProjectDetailResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Deleted          *bool                                                   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Dependencies     *string                                                 `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	Ext              *string                                                 `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                 `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                 `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	ModifiedTime     *string                                                 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status           *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopQueryAvatarProjectDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopQueryAvatarProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetAutoBuild(v bool) *PopQueryAvatarProjectDetailResponseBodyData {
	s.AutoBuild = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetBizUsage(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetBuildDetail(v *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) *PopQueryAvatarProjectDetailResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetCheckStatus(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.CheckStatus = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetCreateMode(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetCreateTime(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetDataset(v *PopQueryAvatarProjectDetailResponseBodyDataDataset) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetDeleted(v bool) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetDependencies(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Dependencies = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetExt(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetId(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetIntro(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetMaterialCoverUrl(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetModifiedTime(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetStatus(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetTitle(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyData) SetType(v string) *PopQueryAvatarProjectDetailResponseBodyData {
	s.Type = &v
	return s
}

type PopQueryAvatarProjectDetailResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted           *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetCreateTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.CreateTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetDeleted(v bool) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.Deleted = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetModifiedTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.ModifiedTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetRunningTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetStatus(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.Status = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopQueryAvatarProjectDetailResponseBodyDataDataset struct {
	BuildResultUrl  map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	CoverUrl        *string                `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CreateTime      *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *bool                  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	ErrorCode       *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GlbModelUrl     *string                `json:"GlbModelUrl,omitempty" xml:"GlbModelUrl,omitempty"`
	ModelUrl        *string                `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	ModifiedTime    *string                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OriginResultUrl *string                `json:"OriginResultUrl,omitempty" xml:"OriginResultUrl,omitempty"`
	OssKey          *string                `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PoseUrl         *string                `json:"PoseUrl,omitempty" xml:"PoseUrl,omitempty"`
	PreviewUrl      *string                `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s PopQueryAvatarProjectDetailResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopQueryAvatarProjectDetailResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetCoverUrl(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.CoverUrl = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetCreateTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.CreateTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetDeleted(v bool) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.Deleted = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetErrorCode(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetErrorMessage(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetGlbModelUrl(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.GlbModelUrl = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetModelUrl(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.ModelUrl = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetModifiedTime(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.ModifiedTime = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetOriginResultUrl(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.OriginResultUrl = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetOssKey(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.OssKey = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetPoseUrl(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.PoseUrl = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponseBodyDataDataset) SetPreviewUrl(v string) *PopQueryAvatarProjectDetailResponseBodyDataDataset {
	s.PreviewUrl = &v
	return s
}

type PopQueryAvatarProjectDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopQueryAvatarProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopQueryAvatarProjectDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PopQueryAvatarProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *PopQueryAvatarProjectDetailResponse) SetHeaders(v map[string]*string) *PopQueryAvatarProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *PopQueryAvatarProjectDetailResponse) SetStatusCode(v int32) *PopQueryAvatarProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *PopQueryAvatarProjectDetailResponse) SetBody(v *PopQueryAvatarProjectDetailResponseBody) *PopQueryAvatarProjectDetailResponse {
	s.Body = v
	return s
}

type PopQueryLatestAvatarProjectDetailByUserRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s PopQueryLatestAvatarProjectDetailByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLatestAvatarProjectDetailByUserRequest) GoString() string {
	return s.String()
}

func (s *PopQueryLatestAvatarProjectDetailByUserRequest) SetJwtToken(v string) *PopQueryLatestAvatarProjectDetailByUserRequest {
	s.JwtToken = &v
	return s
}

type PopQueryLatestAvatarProjectDetailByUserResponseBody struct {
	Code      *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopQueryLatestAvatarProjectDetailByUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBody) GoString() string {
	return s.String()
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBody) SetCode(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBody {
	s.Code = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBody) SetData(v *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) *PopQueryLatestAvatarProjectDetailByUserResponseBody {
	s.Data = v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBody) SetMessage(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBody {
	s.Message = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBody) SetRequestId(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBody) SetSuccess(v bool) *PopQueryLatestAvatarProjectDetailByUserResponseBody {
	s.Success = &v
	return s
}

type PopQueryLatestAvatarProjectDetailByUserResponseBodyData struct {
	BizUsage    *string                                                             `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	CreateTime  *string                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Dataset     *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Ext         *string                                                             `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id          *string                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro       *string                                                             `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Status      *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string                                                             `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetBizUsage(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetBuildDetail(v *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetCreateTime(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetDataset(v *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetExt(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetId(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetIntro(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetStatus(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyData) SetTitle(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyData {
	s.Title = &v
	return s
}

type PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail struct {
	RunningTime *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail) SetRunningTime(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail) SetStatus(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataBuildDetail {
	s.Status = &v
	return s
}

type PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset struct {
	BuildResultUrl map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	ErrorCode      *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset) SetErrorCode(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset) SetErrorMessage(v string) *PopQueryLatestAvatarProjectDetailByUserResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

type PopQueryLatestAvatarProjectDetailByUserResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopQueryLatestAvatarProjectDetailByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopQueryLatestAvatarProjectDetailByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLatestAvatarProjectDetailByUserResponse) GoString() string {
	return s.String()
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponse) SetHeaders(v map[string]*string) *PopQueryLatestAvatarProjectDetailByUserResponse {
	s.Headers = v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponse) SetStatusCode(v int32) *PopQueryLatestAvatarProjectDetailByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *PopQueryLatestAvatarProjectDetailByUserResponse) SetBody(v *PopQueryLatestAvatarProjectDetailByUserResponseBody) *PopQueryLatestAvatarProjectDetailByUserResponse {
	s.Body = v
	return s
}

type PopQueryLivePortraitModelScopeProjectDetailRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopQueryLivePortraitModelScopeProjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLivePortraitModelScopeProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *PopQueryLivePortraitModelScopeProjectDetailRequest) SetProjectId(v string) *PopQueryLivePortraitModelScopeProjectDetailRequest {
	s.ProjectId = &v
	return s
}

type PopQueryLivePortraitModelScopeProjectDetailResponseBody struct {
	Code      *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBody) SetCode(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBody) SetData(v *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) *PopQueryLivePortraitModelScopeProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBody) SetMessage(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBody) SetRequestId(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBody) SetSuccess(v bool) *PopQueryLivePortraitModelScopeProjectDetailResponseBody {
	s.Success = &v
	return s
}

type PopQueryLivePortraitModelScopeProjectDetailResponseBodyData struct {
	BizUsage         *string                                                             `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Dataset          *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Ext              *string                                                             `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id               *string                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro            *string                                                             `json:"Intro,omitempty" xml:"Intro,omitempty"`
	MaterialCoverUrl *string                                                             `json:"MaterialCoverUrl,omitempty" xml:"MaterialCoverUrl,omitempty"`
	Status           *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string                                                             `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetBizUsage(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetDataset(v *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetExt(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetId(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetIntro(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetMaterialCoverUrl(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.MaterialCoverUrl = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetStatus(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetTitle(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Title = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData) SetType(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyData {
	s.Type = &v
	return s
}

type PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset struct {
	BuildResultUrl map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
	ErrorCode      *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset) SetErrorCode(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset {
	s.ErrorCode = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset) SetErrorMessage(v string) *PopQueryLivePortraitModelScopeProjectDetailResponseBodyDataDataset {
	s.ErrorMessage = &v
	return s
}

type PopQueryLivePortraitModelScopeProjectDetailResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopQueryLivePortraitModelScopeProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PopQueryLivePortraitModelScopeProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponse) SetHeaders(v map[string]*string) *PopQueryLivePortraitModelScopeProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponse) SetStatusCode(v int32) *PopQueryLivePortraitModelScopeProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *PopQueryLivePortraitModelScopeProjectDetailResponse) SetBody(v *PopQueryLivePortraitModelScopeProjectDetailResponseBody) *PopQueryLivePortraitModelScopeProjectDetailResponse {
	s.Body = v
	return s
}

type PopQueryObjectGenerationProjectDetailRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopQueryObjectGenerationProjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PopQueryObjectGenerationProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *PopQueryObjectGenerationProjectDetailRequest) SetJwtToken(v string) *PopQueryObjectGenerationProjectDetailRequest {
	s.JwtToken = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailRequest) SetProjectId(v string) *PopQueryObjectGenerationProjectDetailRequest {
	s.ProjectId = &v
	return s
}

type PopQueryObjectGenerationProjectDetailResponseBody struct {
	Code      *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopQueryObjectGenerationProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopQueryObjectGenerationProjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopQueryObjectGenerationProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *PopQueryObjectGenerationProjectDetailResponseBody) SetCode(v string) *PopQueryObjectGenerationProjectDetailResponseBody {
	s.Code = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBody) SetData(v *PopQueryObjectGenerationProjectDetailResponseBodyData) *PopQueryObjectGenerationProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBody) SetMessage(v string) *PopQueryObjectGenerationProjectDetailResponseBody {
	s.Message = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBody) SetRequestId(v string) *PopQueryObjectGenerationProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBody) SetSuccess(v bool) *PopQueryObjectGenerationProjectDetailResponseBody {
	s.Success = &v
	return s
}

type PopQueryObjectGenerationProjectDetailResponseBodyData struct {
	BizUsage    *string                                                           `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	BuildDetail *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail `json:"BuildDetail,omitempty" xml:"BuildDetail,omitempty" type:"Struct"`
	Dataset     *PopQueryObjectGenerationProjectDetailResponseBodyDataDataset     `json:"Dataset,omitempty" xml:"Dataset,omitempty" type:"Struct"`
	Ext         *string                                                           `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Id          *string                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro       *string                                                           `json:"Intro,omitempty" xml:"Intro,omitempty"`
	Status      *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string                                                           `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PopQueryObjectGenerationProjectDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopQueryObjectGenerationProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetBizUsage(v string) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.BizUsage = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetBuildDetail(v *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.BuildDetail = v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetDataset(v *PopQueryObjectGenerationProjectDetailResponseBodyDataDataset) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.Dataset = v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetExt(v string) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.Ext = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetId(v string) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetIntro(v string) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.Intro = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetStatus(v string) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyData) SetTitle(v string) *PopQueryObjectGenerationProjectDetailResponseBodyData {
	s.Title = &v
	return s
}

type PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail struct {
	CompletedTime     *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int64  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SubmitTime        *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) String() string {
	return tea.Prettify(s)
}

func (s PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) GoString() string {
	return s.String()
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) SetCompletedTime(v string) *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail {
	s.CompletedTime = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) SetErrorMessage(v string) *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail {
	s.ErrorMessage = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) SetEstimatedDuration(v int64) *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail {
	s.EstimatedDuration = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) SetRunningTime(v string) *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail {
	s.RunningTime = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail) SetSubmitTime(v string) *PopQueryObjectGenerationProjectDetailResponseBodyDataBuildDetail {
	s.SubmitTime = &v
	return s
}

type PopQueryObjectGenerationProjectDetailResponseBodyDataDataset struct {
	BuildResultUrl map[string]interface{} `json:"BuildResultUrl,omitempty" xml:"BuildResultUrl,omitempty"`
}

func (s PopQueryObjectGenerationProjectDetailResponseBodyDataDataset) String() string {
	return tea.Prettify(s)
}

func (s PopQueryObjectGenerationProjectDetailResponseBodyDataDataset) GoString() string {
	return s.String()
}

func (s *PopQueryObjectGenerationProjectDetailResponseBodyDataDataset) SetBuildResultUrl(v map[string]interface{}) *PopQueryObjectGenerationProjectDetailResponseBodyDataDataset {
	s.BuildResultUrl = v
	return s
}

type PopQueryObjectGenerationProjectDetailResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopQueryObjectGenerationProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopQueryObjectGenerationProjectDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PopQueryObjectGenerationProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *PopQueryObjectGenerationProjectDetailResponse) SetHeaders(v map[string]*string) *PopQueryObjectGenerationProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponse) SetStatusCode(v int32) *PopQueryObjectGenerationProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *PopQueryObjectGenerationProjectDetailResponse) SetBody(v *PopQueryObjectGenerationProjectDetailResponseBody) *PopQueryObjectGenerationProjectDetailResponse {
	s.Body = v
	return s
}

type PopRetryAITryOnTaskRequest struct {
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopRetryAITryOnTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PopRetryAITryOnTaskRequest) GoString() string {
	return s.String()
}

func (s *PopRetryAITryOnTaskRequest) SetJwtToken(v string) *PopRetryAITryOnTaskRequest {
	s.JwtToken = &v
	return s
}

func (s *PopRetryAITryOnTaskRequest) SetProjectId(v string) *PopRetryAITryOnTaskRequest {
	s.ProjectId = &v
	return s
}

type PopRetryAITryOnTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopRetryAITryOnTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopRetryAITryOnTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PopRetryAITryOnTaskResponseBody) SetCode(v string) *PopRetryAITryOnTaskResponseBody {
	s.Code = &v
	return s
}

func (s *PopRetryAITryOnTaskResponseBody) SetMessage(v string) *PopRetryAITryOnTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PopRetryAITryOnTaskResponseBody) SetRequestId(v string) *PopRetryAITryOnTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopRetryAITryOnTaskResponseBody) SetSuccess(v bool) *PopRetryAITryOnTaskResponseBody {
	s.Success = &v
	return s
}

type PopRetryAITryOnTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopRetryAITryOnTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopRetryAITryOnTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PopRetryAITryOnTaskResponse) GoString() string {
	return s.String()
}

func (s *PopRetryAITryOnTaskResponse) SetHeaders(v map[string]*string) *PopRetryAITryOnTaskResponse {
	s.Headers = v
	return s
}

func (s *PopRetryAITryOnTaskResponse) SetStatusCode(v int32) *PopRetryAITryOnTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PopRetryAITryOnTaskResponse) SetBody(v *PopRetryAITryOnTaskResponseBody) *PopRetryAITryOnTaskResponse {
	s.Body = v
	return s
}

type PopSubmitAITryOnJobRequest struct {
	BottomsId    *string `json:"BottomsId,omitempty" xml:"BottomsId,omitempty"`
	ClothingType *string `json:"ClothingType,omitempty" xml:"ClothingType,omitempty"`
	JwtToken     *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ModelId      *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ShoeType     *string `json:"ShoeType,omitempty" xml:"ShoeType,omitempty"`
	SuitId       *string `json:"SuitId,omitempty" xml:"SuitId,omitempty"`
	TopsId       *string `json:"TopsId,omitempty" xml:"TopsId,omitempty"`
}

func (s PopSubmitAITryOnJobRequest) String() string {
	return tea.Prettify(s)
}

func (s PopSubmitAITryOnJobRequest) GoString() string {
	return s.String()
}

func (s *PopSubmitAITryOnJobRequest) SetBottomsId(v string) *PopSubmitAITryOnJobRequest {
	s.BottomsId = &v
	return s
}

func (s *PopSubmitAITryOnJobRequest) SetClothingType(v string) *PopSubmitAITryOnJobRequest {
	s.ClothingType = &v
	return s
}

func (s *PopSubmitAITryOnJobRequest) SetJwtToken(v string) *PopSubmitAITryOnJobRequest {
	s.JwtToken = &v
	return s
}

func (s *PopSubmitAITryOnJobRequest) SetModelId(v string) *PopSubmitAITryOnJobRequest {
	s.ModelId = &v
	return s
}

func (s *PopSubmitAITryOnJobRequest) SetShoeType(v string) *PopSubmitAITryOnJobRequest {
	s.ShoeType = &v
	return s
}

func (s *PopSubmitAITryOnJobRequest) SetSuitId(v string) *PopSubmitAITryOnJobRequest {
	s.SuitId = &v
	return s
}

func (s *PopSubmitAITryOnJobRequest) SetTopsId(v string) *PopSubmitAITryOnJobRequest {
	s.TopsId = &v
	return s
}

type PopSubmitAITryOnJobResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopSubmitAITryOnJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopSubmitAITryOnJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopSubmitAITryOnJobResponseBody) GoString() string {
	return s.String()
}

func (s *PopSubmitAITryOnJobResponseBody) SetCode(v string) *PopSubmitAITryOnJobResponseBody {
	s.Code = &v
	return s
}

func (s *PopSubmitAITryOnJobResponseBody) SetData(v *PopSubmitAITryOnJobResponseBodyData) *PopSubmitAITryOnJobResponseBody {
	s.Data = v
	return s
}

func (s *PopSubmitAITryOnJobResponseBody) SetMessage(v string) *PopSubmitAITryOnJobResponseBody {
	s.Message = &v
	return s
}

func (s *PopSubmitAITryOnJobResponseBody) SetRequestId(v string) *PopSubmitAITryOnJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopSubmitAITryOnJobResponseBody) SetSuccess(v bool) *PopSubmitAITryOnJobResponseBody {
	s.Success = &v
	return s
}

type PopSubmitAITryOnJobResponseBodyData struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PopSubmitAITryOnJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopSubmitAITryOnJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopSubmitAITryOnJobResponseBodyData) SetProjectId(v string) *PopSubmitAITryOnJobResponseBodyData {
	s.ProjectId = &v
	return s
}

type PopSubmitAITryOnJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopSubmitAITryOnJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopSubmitAITryOnJobResponse) String() string {
	return tea.Prettify(s)
}

func (s PopSubmitAITryOnJobResponse) GoString() string {
	return s.String()
}

func (s *PopSubmitAITryOnJobResponse) SetHeaders(v map[string]*string) *PopSubmitAITryOnJobResponse {
	s.Headers = v
	return s
}

func (s *PopSubmitAITryOnJobResponse) SetStatusCode(v int32) *PopSubmitAITryOnJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PopSubmitAITryOnJobResponse) SetBody(v *PopSubmitAITryOnJobResponseBody) *PopSubmitAITryOnJobResponse {
	s.Body = v
	return s
}

type PopUploadMaterialRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s PopUploadMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s PopUploadMaterialRequest) GoString() string {
	return s.String()
}

func (s *PopUploadMaterialRequest) SetJwtToken(v string) *PopUploadMaterialRequest {
	s.JwtToken = &v
	return s
}

type PopUploadMaterialResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PopUploadMaterialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopUploadMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopUploadMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *PopUploadMaterialResponseBody) SetCode(v string) *PopUploadMaterialResponseBody {
	s.Code = &v
	return s
}

func (s *PopUploadMaterialResponseBody) SetData(v *PopUploadMaterialResponseBodyData) *PopUploadMaterialResponseBody {
	s.Data = v
	return s
}

func (s *PopUploadMaterialResponseBody) SetMessage(v string) *PopUploadMaterialResponseBody {
	s.Message = &v
	return s
}

func (s *PopUploadMaterialResponseBody) SetRequestId(v string) *PopUploadMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopUploadMaterialResponseBody) SetSuccess(v bool) *PopUploadMaterialResponseBody {
	s.Success = &v
	return s
}

type PopUploadMaterialResponseBodyData struct {
	OssKey *string                                  `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Policy *PopUploadMaterialResponseBodyDataPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
}

func (s PopUploadMaterialResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PopUploadMaterialResponseBodyData) GoString() string {
	return s.String()
}

func (s *PopUploadMaterialResponseBodyData) SetOssKey(v string) *PopUploadMaterialResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *PopUploadMaterialResponseBodyData) SetPolicy(v *PopUploadMaterialResponseBodyDataPolicy) *PopUploadMaterialResponseBodyData {
	s.Policy = v
	return s
}

type PopUploadMaterialResponseBodyDataPolicy struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s PopUploadMaterialResponseBodyDataPolicy) String() string {
	return tea.Prettify(s)
}

func (s PopUploadMaterialResponseBodyDataPolicy) GoString() string {
	return s.String()
}

func (s *PopUploadMaterialResponseBodyDataPolicy) SetAccessId(v string) *PopUploadMaterialResponseBodyDataPolicy {
	s.AccessId = &v
	return s
}

func (s *PopUploadMaterialResponseBodyDataPolicy) SetDir(v string) *PopUploadMaterialResponseBodyDataPolicy {
	s.Dir = &v
	return s
}

func (s *PopUploadMaterialResponseBodyDataPolicy) SetExpire(v string) *PopUploadMaterialResponseBodyDataPolicy {
	s.Expire = &v
	return s
}

func (s *PopUploadMaterialResponseBodyDataPolicy) SetHost(v string) *PopUploadMaterialResponseBodyDataPolicy {
	s.Host = &v
	return s
}

func (s *PopUploadMaterialResponseBodyDataPolicy) SetPolicy(v string) *PopUploadMaterialResponseBodyDataPolicy {
	s.Policy = &v
	return s
}

func (s *PopUploadMaterialResponseBodyDataPolicy) SetSignature(v string) *PopUploadMaterialResponseBodyDataPolicy {
	s.Signature = &v
	return s
}

type PopUploadMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopUploadMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopUploadMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s PopUploadMaterialResponse) GoString() string {
	return s.String()
}

func (s *PopUploadMaterialResponse) SetHeaders(v map[string]*string) *PopUploadMaterialResponse {
	s.Headers = v
	return s
}

func (s *PopUploadMaterialResponse) SetStatusCode(v int32) *PopUploadMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *PopUploadMaterialResponse) SetBody(v *PopUploadMaterialResponseBody) *PopUploadMaterialResponse {
	s.Body = v
	return s
}

type PopVideoSaveSourceRequest struct {
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s PopVideoSaveSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s PopVideoSaveSourceRequest) GoString() string {
	return s.String()
}

func (s *PopVideoSaveSourceRequest) SetJwtToken(v string) *PopVideoSaveSourceRequest {
	s.JwtToken = &v
	return s
}

func (s *PopVideoSaveSourceRequest) SetProjectId(v string) *PopVideoSaveSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *PopVideoSaveSourceRequest) SetSourceType(v string) *PopVideoSaveSourceRequest {
	s.SourceType = &v
	return s
}

type PopVideoSaveSourceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PopVideoSaveSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PopVideoSaveSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PopVideoSaveSourceResponseBody) SetCode(v string) *PopVideoSaveSourceResponseBody {
	s.Code = &v
	return s
}

func (s *PopVideoSaveSourceResponseBody) SetErrorName(v string) *PopVideoSaveSourceResponseBody {
	s.ErrorName = &v
	return s
}

func (s *PopVideoSaveSourceResponseBody) SetHttpCode(v int32) *PopVideoSaveSourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *PopVideoSaveSourceResponseBody) SetMessage(v string) *PopVideoSaveSourceResponseBody {
	s.Message = &v
	return s
}

func (s *PopVideoSaveSourceResponseBody) SetRequestId(v string) *PopVideoSaveSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PopVideoSaveSourceResponseBody) SetSuccess(v bool) *PopVideoSaveSourceResponseBody {
	s.Success = &v
	return s
}

type PopVideoSaveSourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PopVideoSaveSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PopVideoSaveSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s PopVideoSaveSourceResponse) GoString() string {
	return s.String()
}

func (s *PopVideoSaveSourceResponse) SetHeaders(v map[string]*string) *PopVideoSaveSourceResponse {
	s.Headers = v
	return s
}

func (s *PopVideoSaveSourceResponse) SetStatusCode(v int32) *PopVideoSaveSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PopVideoSaveSourceResponse) SetBody(v *PopVideoSaveSourceResponseBody) *PopVideoSaveSourceResponse {
	s.Body = v
	return s
}

type QueryDigitalHumanProjectRequest struct {
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s QueryDigitalHumanProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalHumanProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryDigitalHumanProjectRequest) SetIds(v string) *QueryDigitalHumanProjectRequest {
	s.Ids = &v
	return s
}

func (s *QueryDigitalHumanProjectRequest) SetJwtToken(v string) *QueryDigitalHumanProjectRequest {
	s.JwtToken = &v
	return s
}

type QueryDigitalHumanProjectResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryDigitalHumanProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDigitalHumanProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalHumanProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDigitalHumanProjectResponseBody) SetCode(v string) *QueryDigitalHumanProjectResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBody) SetData(v []*QueryDigitalHumanProjectResponseBodyData) *QueryDigitalHumanProjectResponseBody {
	s.Data = v
	return s
}

func (s *QueryDigitalHumanProjectResponseBody) SetMessage(v string) *QueryDigitalHumanProjectResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBody) SetRequestId(v string) *QueryDigitalHumanProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBody) SetSuccess(v bool) *QueryDigitalHumanProjectResponseBody {
	s.Success = &v
	return s
}

type QueryDigitalHumanProjectResponseBodyData struct {
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EstimatedDuration *int32  `json:"EstimatedDuration,omitempty" xml:"EstimatedDuration,omitempty"`
	FileUrl           *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Intro             *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	RunningTime       *string `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubtitleUrl       *string `json:"SubtitleUrl,omitempty" xml:"SubtitleUrl,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoLength       *int32  `json:"VideoLength,omitempty" xml:"VideoLength,omitempty"`
	WaitingTime       *int32  `json:"WaitingTime,omitempty" xml:"WaitingTime,omitempty"`
}

func (s QueryDigitalHumanProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalHumanProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetErrorCode(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetErrorMessage(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetEstimatedDuration(v int32) *QueryDigitalHumanProjectResponseBodyData {
	s.EstimatedDuration = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetFileUrl(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetId(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetIntro(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.Intro = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetRunningTime(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetStatus(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetSubtitleUrl(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.SubtitleUrl = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetTitle(v string) *QueryDigitalHumanProjectResponseBodyData {
	s.Title = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetVideoLength(v int32) *QueryDigitalHumanProjectResponseBodyData {
	s.VideoLength = &v
	return s
}

func (s *QueryDigitalHumanProjectResponseBodyData) SetWaitingTime(v int32) *QueryDigitalHumanProjectResponseBodyData {
	s.WaitingTime = &v
	return s
}

type QueryDigitalHumanProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDigitalHumanProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDigitalHumanProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDigitalHumanProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryDigitalHumanProjectResponse) SetHeaders(v map[string]*string) *QueryDigitalHumanProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryDigitalHumanProjectResponse) SetStatusCode(v int32) *QueryDigitalHumanProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDigitalHumanProjectResponse) SetBody(v *QueryDigitalHumanProjectResponseBody) *QueryDigitalHumanProjectResponse {
	s.Body = v
	return s
}

type QueryLongTtsResultRequest struct {
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s QueryLongTtsResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLongTtsResultRequest) GoString() string {
	return s.String()
}

func (s *QueryLongTtsResultRequest) SetJobId(v string) *QueryLongTtsResultRequest {
	s.JobId = &v
	return s
}

func (s *QueryLongTtsResultRequest) SetJwtToken(v string) *QueryLongTtsResultRequest {
	s.JwtToken = &v
	return s
}

type QueryLongTtsResultResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryLongTtsResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLongTtsResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLongTtsResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLongTtsResultResponseBody) SetCode(v string) *QueryLongTtsResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLongTtsResultResponseBody) SetData(v *QueryLongTtsResultResponseBodyData) *QueryLongTtsResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryLongTtsResultResponseBody) SetMessage(v string) *QueryLongTtsResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryLongTtsResultResponseBody) SetRequestId(v string) *QueryLongTtsResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLongTtsResultResponseBody) SetSuccess(v bool) *QueryLongTtsResultResponseBody {
	s.Success = &v
	return s
}

type QueryLongTtsResultResponseBodyData struct {
	AudioUrl *string  `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Status   *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryLongTtsResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryLongTtsResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryLongTtsResultResponseBodyData) SetAudioUrl(v string) *QueryLongTtsResultResponseBodyData {
	s.AudioUrl = &v
	return s
}

func (s *QueryLongTtsResultResponseBodyData) SetDuration(v float32) *QueryLongTtsResultResponseBodyData {
	s.Duration = &v
	return s
}

func (s *QueryLongTtsResultResponseBodyData) SetStatus(v string) *QueryLongTtsResultResponseBodyData {
	s.Status = &v
	return s
}

type QueryLongTtsResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryLongTtsResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryLongTtsResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLongTtsResultResponse) GoString() string {
	return s.String()
}

func (s *QueryLongTtsResultResponse) SetHeaders(v map[string]*string) *QueryLongTtsResultResponse {
	s.Headers = v
	return s
}

func (s *QueryLongTtsResultResponse) SetStatusCode(v int32) *QueryLongTtsResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLongTtsResultResponse) SetBody(v *QueryLongTtsResultResponseBody) *QueryLongTtsResultResponse {
	s.Body = v
	return s
}

type QueryMotionShopVideoDetectResultRequest struct {
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s QueryMotionShopVideoDetectResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMotionShopVideoDetectResultRequest) GoString() string {
	return s.String()
}

func (s *QueryMotionShopVideoDetectResultRequest) SetJobId(v string) *QueryMotionShopVideoDetectResultRequest {
	s.JobId = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultRequest) SetJwtToken(v string) *QueryMotionShopVideoDetectResultRequest {
	s.JwtToken = &v
	return s
}

type QueryMotionShopVideoDetectResultResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryMotionShopVideoDetectResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMotionShopVideoDetectResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMotionShopVideoDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMotionShopVideoDetectResultResponseBody) SetCode(v string) *QueryMotionShopVideoDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBody) SetData(v *QueryMotionShopVideoDetectResultResponseBodyData) *QueryMotionShopVideoDetectResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBody) SetMessage(v string) *QueryMotionShopVideoDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBody) SetRequestId(v string) *QueryMotionShopVideoDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBody) SetSuccess(v bool) *QueryMotionShopVideoDetectResultResponseBody {
	s.Success = &v
	return s
}

type QueryMotionShopVideoDetectResultResponseBodyData struct {
	DetectResult *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult `json:"DetectResult,omitempty" xml:"DetectResult,omitempty" type:"Struct"`
	Status       *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	VideoId      *string                                                       `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s QueryMotionShopVideoDetectResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMotionShopVideoDetectResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMotionShopVideoDetectResultResponseBodyData) SetDetectResult(v *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) *QueryMotionShopVideoDetectResultResponseBodyData {
	s.DetectResult = v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBodyData) SetStatus(v string) *QueryMotionShopVideoDetectResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBodyData) SetVideoId(v string) *QueryMotionShopVideoDetectResultResponseBodyData {
	s.VideoId = &v
	return s
}

type QueryMotionShopVideoDetectResultResponseBodyDataDetectResult struct {
	Box                []*float64 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
	Code               *int32     `json:"Code,omitempty" xml:"Code,omitempty"`
	CoverUrl           *string    `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	Message            *string    `json:"Message,omitempty" xml:"Message,omitempty"`
	SelectedFrameIndex *int32     `json:"SelectedFrameIndex,omitempty" xml:"SelectedFrameIndex,omitempty"`
}

func (s QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) GoString() string {
	return s.String()
}

func (s *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) SetBox(v []*float64) *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult {
	s.Box = v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) SetCode(v int32) *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult {
	s.Code = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) SetCoverUrl(v string) *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult {
	s.CoverUrl = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) SetMessage(v string) *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult {
	s.Message = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult) SetSelectedFrameIndex(v int32) *QueryMotionShopVideoDetectResultResponseBodyDataDetectResult {
	s.SelectedFrameIndex = &v
	return s
}

type QueryMotionShopVideoDetectResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMotionShopVideoDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMotionShopVideoDetectResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMotionShopVideoDetectResultResponse) GoString() string {
	return s.String()
}

func (s *QueryMotionShopVideoDetectResultResponse) SetHeaders(v map[string]*string) *QueryMotionShopVideoDetectResultResponse {
	s.Headers = v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponse) SetStatusCode(v int32) *QueryMotionShopVideoDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMotionShopVideoDetectResultResponse) SetBody(v *QueryMotionShopVideoDetectResultResponseBody) *QueryMotionShopVideoDetectResultResponse {
	s.Body = v
	return s
}

type SubmitLongTtsTaskRequest struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	TtsVoiceId *string `json:"TtsVoiceId,omitempty" xml:"TtsVoiceId,omitempty"`
}

func (s SubmitLongTtsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitLongTtsTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitLongTtsTaskRequest) SetContent(v string) *SubmitLongTtsTaskRequest {
	s.Content = &v
	return s
}

func (s *SubmitLongTtsTaskRequest) SetJwtToken(v string) *SubmitLongTtsTaskRequest {
	s.JwtToken = &v
	return s
}

func (s *SubmitLongTtsTaskRequest) SetTtsVoiceId(v string) *SubmitLongTtsTaskRequest {
	s.TtsVoiceId = &v
	return s
}

type SubmitLongTtsTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitLongTtsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitLongTtsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLongTtsTaskResponseBody) SetCode(v string) *SubmitLongTtsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitLongTtsTaskResponseBody) SetData(v string) *SubmitLongTtsTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitLongTtsTaskResponseBody) SetMessage(v string) *SubmitLongTtsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitLongTtsTaskResponseBody) SetRequestId(v string) *SubmitLongTtsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLongTtsTaskResponseBody) SetSuccess(v bool) *SubmitLongTtsTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitLongTtsTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitLongTtsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitLongTtsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitLongTtsTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitLongTtsTaskResponse) SetHeaders(v map[string]*string) *SubmitLongTtsTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitLongTtsTaskResponse) SetStatusCode(v int32) *SubmitLongTtsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLongTtsTaskResponse) SetBody(v *SubmitLongTtsTaskResponseBody) *SubmitLongTtsTaskResponse {
	s.Body = v
	return s
}

type SubmitMotionShopTaskRequest struct {
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId  *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitMotionShopTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMotionShopTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitMotionShopTaskRequest) SetAvatarId(v string) *SubmitMotionShopTaskRequest {
	s.AvatarId = &v
	return s
}

func (s *SubmitMotionShopTaskRequest) SetJwtToken(v string) *SubmitMotionShopTaskRequest {
	s.JwtToken = &v
	return s
}

func (s *SubmitMotionShopTaskRequest) SetTitle(v string) *SubmitMotionShopTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitMotionShopTaskRequest) SetVideoId(v string) *SubmitMotionShopTaskRequest {
	s.VideoId = &v
	return s
}

type SubmitMotionShopTaskResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitMotionShopTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitMotionShopTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitMotionShopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMotionShopTaskResponseBody) SetCode(v string) *SubmitMotionShopTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitMotionShopTaskResponseBody) SetData(v *SubmitMotionShopTaskResponseBodyData) *SubmitMotionShopTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitMotionShopTaskResponseBody) SetMessage(v string) *SubmitMotionShopTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitMotionShopTaskResponseBody) SetRequestId(v string) *SubmitMotionShopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMotionShopTaskResponseBody) SetSuccess(v bool) *SubmitMotionShopTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitMotionShopTaskResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitMotionShopTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitMotionShopTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitMotionShopTaskResponseBodyData) SetTaskId(v string) *SubmitMotionShopTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitMotionShopTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitMotionShopTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitMotionShopTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitMotionShopTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitMotionShopTaskResponse) SetHeaders(v map[string]*string) *SubmitMotionShopTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitMotionShopTaskResponse) SetStatusCode(v int32) *SubmitMotionShopTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMotionShopTaskResponse) SetBody(v *SubmitMotionShopTaskResponseBody) *SubmitMotionShopTaskResponse {
	s.Body = v
	return s
}

type UpdateUserEmailRequest struct {
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s UpdateUserEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserEmailRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserEmailRequest) SetEmail(v string) *UpdateUserEmailRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserEmailRequest) SetJwtToken(v string) *UpdateUserEmailRequest {
	s.JwtToken = &v
	return s
}

type UpdateUserEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserEmailResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserEmailResponseBody) SetCode(v string) *UpdateUserEmailResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserEmailResponseBody) SetErrorName(v string) *UpdateUserEmailResponseBody {
	s.ErrorName = &v
	return s
}

func (s *UpdateUserEmailResponseBody) SetHttpCode(v int32) *UpdateUserEmailResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateUserEmailResponseBody) SetMessage(v string) *UpdateUserEmailResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserEmailResponseBody) SetRequestId(v string) *UpdateUserEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserEmailResponseBody) SetSuccess(v bool) *UpdateUserEmailResponseBody {
	s.Success = &v
	return s
}

type UpdateUserEmailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserEmailResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserEmailResponse) SetHeaders(v map[string]*string) *UpdateUserEmailResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserEmailResponse) SetStatusCode(v int32) *UpdateUserEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserEmailResponse) SetBody(v *UpdateUserEmailResponseBody) *UpdateUserEmailResponse {
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
		"ap-northeast-1":              tea.String("xrengine-daily.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("xrengine-daily.aliyuncs.com"),
		"ap-south-1":                  tea.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-1":              tea.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-2":              tea.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-3":              tea.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-5":              tea.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing":                  tea.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("xrengine-daily.aliyuncs.com"),
		"cn-chengdu":                  tea.String("xrengine-daily.aliyuncs.com"),
		"cn-edge-1":                   tea.String("xrengine-daily.aliyuncs.com"),
		"cn-fujian":                   tea.String("xrengine-daily.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hongkong":                 tea.String("xrengine-daily.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("xrengine-daily.aliyuncs.com"),
		"cn-huhehaote":                tea.String("xrengine-daily.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("xrengine-daily.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("xrengine-daily.aliyuncs.com"),
		"cn-qingdao":                  tea.String("xrengine-daily.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai":                 tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("xrengine-daily.aliyuncs.com"),
		"cn-wuhan":                    tea.String("xrengine-daily.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("xrengine-daily.aliyuncs.com"),
		"cn-yushanfang":               tea.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("xrengine-daily.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("xrengine-daily.aliyuncs.com"),
		"eu-central-1":                tea.String("xrengine-daily.aliyuncs.com"),
		"eu-west-1":                   tea.String("xrengine-daily.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("xrengine-daily.aliyuncs.com"),
		"me-east-1":                   tea.String("xrengine-daily.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("xrengine-daily.aliyuncs.com"),
		"us-east-1":                   tea.String("xrengine-daily.aliyuncs.com"),
		"us-west-1":                   tea.String("xrengine-daily.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("xrengine"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AuthUserWithOptions(request *AuthUserRequest, runtime *util.RuntimeOptions) (_result *AuthUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthUser"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthUser(request *AuthUserRequest) (_result *AuthUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthUserResponse{}
	_body, _err := client.AuthUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQueryMotionShopTaskStatusWithOptions(request *BatchQueryMotionShopTaskStatusRequest, runtime *util.RuntimeOptions) (_result *BatchQueryMotionShopTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchQueryMotionShopTaskStatus"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryMotionShopTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQueryMotionShopTaskStatus(request *BatchQueryMotionShopTaskStatusRequest) (_result *BatchQueryMotionShopTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchQueryMotionShopTaskStatusResponse{}
	_body, _err := client.BatchQueryMotionShopTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAvatarTalkProjectWithOptions(request *CreateAvatarTalkProjectRequest, runtime *util.RuntimeOptions) (_result *CreateAvatarTalkProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvatarProjectId)) {
		body["AvatarProjectId"] = request.AvatarProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.TtsVoice)) {
		body["TtsVoice"] = request.TtsVoice
	}

	if !tea.BoolValue(util.IsUnset(request.TxtContent)) {
		body["TxtContent"] = request.TxtContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAvatarTalkProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAvatarTalkProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAvatarTalkProject(request *CreateAvatarTalkProjectRequest) (_result *CreateAvatarTalkProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAvatarTalkProjectResponse{}
	_body, _err := client.CreateAvatarTalkProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDigitalHumanProjectWithOptions(request *CreateDigitalHumanProjectRequest, runtime *util.RuntimeOptions) (_result *CreateDigitalHumanProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AudioId)) {
		body["AudioId"] = request.AudioId
	}

	if !tea.BoolValue(util.IsUnset(request.AudioUrl)) {
		body["AudioUrl"] = request.AudioUrl
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundId)) {
		body["BackgroundId"] = request.BackgroundId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundUrl)) {
		body["BackgroundUrl"] = request.BackgroundUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ForegroundId)) {
		body["ForegroundId"] = request.ForegroundId
	}

	if !tea.BoolValue(util.IsUnset(request.ForegroundUrl)) {
		body["ForegroundUrl"] = request.ForegroundUrl
	}

	if !tea.BoolValue(util.IsUnset(request.HumanLayerStyle)) {
		body["HumanLayerStyle"] = request.HumanLayerStyle
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputConfig)) {
		body["OutputConfig"] = request.OutputConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.TtsVoiceId)) {
		body["TtsVoiceId"] = request.TtsVoiceId
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkImageUrl)) {
		body["WatermarkImageUrl"] = request.WatermarkImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkStyle)) {
		body["WatermarkStyle"] = request.WatermarkStyle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDigitalHumanProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDigitalHumanProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDigitalHumanProject(request *CreateDigitalHumanProjectRequest) (_result *CreateDigitalHumanProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDigitalHumanProjectResponse{}
	_body, _err := client.CreateDigitalHumanProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLivePortraitProjectWithOptions(request *CreateLivePortraitProjectRequest, runtime *util.RuntimeOptions) (_result *CreateLivePortraitProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AudioId)) {
		body["AudioId"] = request.AudioId
	}

	if !tea.BoolValue(util.IsUnset(request.AudioUrl)) {
		body["AudioUrl"] = request.AudioUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.LightModel)) {
		body["LightModel"] = request.LightModel
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OutputConfig)) {
		body["OutputConfig"] = request.OutputConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.TtsVoiceId)) {
		body["TtsVoiceId"] = request.TtsVoiceId
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkImageUrl)) {
		body["WatermarkImageUrl"] = request.WatermarkImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkStyle)) {
		body["WatermarkStyle"] = request.WatermarkStyle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLivePortraitProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLivePortraitProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLivePortraitProject(request *CreateLivePortraitProjectRequest) (_result *CreateLivePortraitProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLivePortraitProjectResponse{}
	_body, _err := client.CreateLivePortraitProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateMotionShopVideoUploadUrlWithOptions(request *GenerateMotionShopVideoUploadUrlRequest, runtime *util.RuntimeOptions) (_result *GenerateMotionShopVideoUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateMotionShopVideoUploadUrl"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateMotionShopVideoUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateMotionShopVideoUploadUrl(request *GenerateMotionShopVideoUploadUrlRequest) (_result *GenerateMotionShopVideoUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateMotionShopVideoUploadUrlResponse{}
	_body, _err := client.GenerateMotionShopVideoUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMapDataWithOptions(request *GetMapDataRequest, runtime *util.RuntimeOptions) (_result *GetMapDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMapData"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMapDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMapData(request *GetMapDataRequest) (_result *GetMapDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMapDataResponse{}
	_body, _err := client.GetMapDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMapPublishDataWithOptions(request *GetMapPublishDataRequest, runtime *util.RuntimeOptions) (_result *GetMapPublishDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMapPublishData"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMapPublishDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMapPublishData(request *GetMapPublishDataRequest) (_result *GetMapPublishDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMapPublishDataResponse{}
	_body, _err := client.GetMapPublishDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitLocateWithOptions(request *InitLocateRequest, runtime *util.RuntimeOptions) (_result *InitLocateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitLocate"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitLocateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitLocate(request *InitLocateRequest) (_result *InitLocateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitLocateResponse{}
	_body, _err := client.InitLocateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommonMaterialsWithOptions(request *ListCommonMaterialsRequest, runtime *util.RuntimeOptions) (_result *ListCommonMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommonMaterials"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommonMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommonMaterials(request *ListCommonMaterialsRequest) (_result *ListCommonMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommonMaterialsResponse{}
	_body, _err := client.ListCommonMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDigitalHumanMaterialsWithOptions(request *ListDigitalHumanMaterialsRequest, runtime *util.RuntimeOptions) (_result *ListDigitalHumanMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyPersonalMaterials)) {
		body["OnlyPersonalMaterials"] = request.OnlyPersonalMaterials
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		body["Types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDigitalHumanMaterials"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDigitalHumanMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDigitalHumanMaterials(request *ListDigitalHumanMaterialsRequest) (_result *ListDigitalHumanMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDigitalHumanMaterialsResponse{}
	_body, _err := client.ListDigitalHumanMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLocationServiceWithOptions(request *ListLocationServiceRequest, runtime *util.RuntimeOptions) (_result *ListLocationServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		body["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		body["SortField"] = request.SortField
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLocationService"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLocationService(request *ListLocationServiceRequest) (_result *ListLocationServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLocationServiceResponse{}
	_body, _err := client.ListLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMotionShopTasksWithOptions(request *ListMotionShopTasksRequest, runtime *util.RuntimeOptions) (_result *ListMotionShopTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMotionShopTasks"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMotionShopTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMotionShopTasks(request *ListMotionShopTasksRequest) (_result *ListMotionShopTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMotionShopTasksResponse{}
	_body, _err := client.ListMotionShopTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LivePortraitFaceDetectWithOptions(request *LivePortraitFaceDetectRequest, runtime *util.RuntimeOptions) (_result *LivePortraitFaceDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LivePortraitFaceDetect"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LivePortraitFaceDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LivePortraitFaceDetect(request *LivePortraitFaceDetectRequest) (_result *LivePortraitFaceDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivePortraitFaceDetectResponse{}
	_body, _err := client.LivePortraitFaceDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LocateWithOptions(request *LocateRequest, runtime *util.RuntimeOptions) (_result *LocateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Image)) {
		body["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Locate"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LocateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Locate(request *LocateRequest) (_result *LocateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LocateResponse{}
	_body, _err := client.LocateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoginModelScopeWithOptions(request *LoginModelScopeRequest, runtime *util.RuntimeOptions) (_result *LoginModelScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EmpId)) {
		body["EmpId"] = request.EmpId
	}

	if !tea.BoolValue(util.IsUnset(request.EmpName)) {
		body["EmpName"] = request.EmpName
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LoginModelScope"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoginModelScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoginModelScope(request *LoginModelScopeRequest) (_result *LoginModelScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoginModelScopeResponse{}
	_body, _err := client.LoginModelScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MotionShopVideoDetectWithOptions(request *MotionShopVideoDetectRequest, runtime *util.RuntimeOptions) (_result *MotionShopVideoDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OssKey)) {
		body["OssKey"] = request.OssKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MotionShopVideoDetect"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MotionShopVideoDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MotionShopVideoDetect(request *MotionShopVideoDetectRequest) (_result *MotionShopVideoDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MotionShopVideoDetectResponse{}
	_body, _err := client.MotionShopVideoDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBatchQueryObjectGenerationProjectStatusWithOptions(request *PopBatchQueryObjectGenerationProjectStatusRequest, runtime *util.RuntimeOptions) (_result *PopBatchQueryObjectGenerationProjectStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectIds)) {
		body["ProjectIds"] = request.ProjectIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBatchQueryObjectGenerationProjectStatus"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBatchQueryObjectGenerationProjectStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBatchQueryObjectGenerationProjectStatus(request *PopBatchQueryObjectGenerationProjectStatusRequest) (_result *PopBatchQueryObjectGenerationProjectStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBatchQueryObjectGenerationProjectStatusResponse{}
	_body, _err := client.PopBatchQueryObjectGenerationProjectStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBatchQueryObjectProjectStatusWithOptions(request *PopBatchQueryObjectProjectStatusRequest, runtime *util.RuntimeOptions) (_result *PopBatchQueryObjectProjectStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectIds)) {
		body["ProjectIds"] = request.ProjectIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBatchQueryObjectProjectStatus"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBatchQueryObjectProjectStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBatchQueryObjectProjectStatus(request *PopBatchQueryObjectProjectStatusRequest) (_result *PopBatchQueryObjectProjectStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBatchQueryObjectProjectStatusResponse{}
	_body, _err := client.PopBatchQueryObjectProjectStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBuildFeatureToAvatarProjectWithOptions(request *PopBuildFeatureToAvatarProjectRequest, runtime *util.RuntimeOptions) (_result *PopBuildFeatureToAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBuildFeatureToAvatarProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBuildFeatureToAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBuildFeatureToAvatarProject(request *PopBuildFeatureToAvatarProjectRequest) (_result *PopBuildFeatureToAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBuildFeatureToAvatarProjectResponse{}
	_body, _err := client.PopBuildFeatureToAvatarProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBuildLivePortraitModelScopeProjectWithOptions(request *PopBuildLivePortraitModelScopeProjectRequest, runtime *util.RuntimeOptions) (_result *PopBuildLivePortraitModelScopeProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBuildLivePortraitModelScopeProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBuildLivePortraitModelScopeProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBuildLivePortraitModelScopeProject(request *PopBuildLivePortraitModelScopeProjectRequest) (_result *PopBuildLivePortraitModelScopeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBuildLivePortraitModelScopeProjectResponse{}
	_body, _err := client.PopBuildLivePortraitModelScopeProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBuildObjectGenerationProjectWithOptions(request *PopBuildObjectGenerationProjectRequest, runtime *util.RuntimeOptions) (_result *PopBuildObjectGenerationProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBuildObjectGenerationProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBuildObjectGenerationProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBuildObjectGenerationProject(request *PopBuildObjectGenerationProjectRequest) (_result *PopBuildObjectGenerationProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBuildObjectGenerationProjectResponse{}
	_body, _err := client.PopBuildObjectGenerationProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBuildObjectProjectWithOptions(request *PopBuildObjectProjectRequest, runtime *util.RuntimeOptions) (_result *PopBuildObjectProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBuildObjectProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBuildObjectProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBuildObjectProject(request *PopBuildObjectProjectRequest) (_result *PopBuildObjectProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBuildObjectProjectResponse{}
	_body, _err := client.PopBuildObjectProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBuildPakRenderProjectWithOptions(request *PopBuildPakRenderProjectRequest, runtime *util.RuntimeOptions) (_result *PopBuildPakRenderProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBuildPakRenderProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBuildPakRenderProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBuildPakRenderProject(request *PopBuildPakRenderProjectRequest) (_result *PopBuildPakRenderProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBuildPakRenderProjectResponse{}
	_body, _err := client.PopBuildPakRenderProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopBuildTextToAvatarProjectWithOptions(request *PopBuildTextToAvatarProjectRequest, runtime *util.RuntimeOptions) (_result *PopBuildTextToAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopBuildTextToAvatarProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopBuildTextToAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopBuildTextToAvatarProject(request *PopBuildTextToAvatarProjectRequest) (_result *PopBuildTextToAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopBuildTextToAvatarProjectResponse{}
	_body, _err := client.PopBuildTextToAvatarProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreateFeatureToAvatarProjectWithOptions(request *PopCreateFeatureToAvatarProjectRequest, runtime *util.RuntimeOptions) (_result *PopCreateFeatureToAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreateFeatureToAvatarProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreateFeatureToAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreateFeatureToAvatarProject(request *PopCreateFeatureToAvatarProjectRequest) (_result *PopCreateFeatureToAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreateFeatureToAvatarProjectResponse{}
	_body, _err := client.PopCreateFeatureToAvatarProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreateLivePortraitModelScopeProjectWithOptions(request *PopCreateLivePortraitModelScopeProjectRequest, runtime *util.RuntimeOptions) (_result *PopCreateLivePortraitModelScopeProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreateLivePortraitModelScopeProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreateLivePortraitModelScopeProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreateLivePortraitModelScopeProject(request *PopCreateLivePortraitModelScopeProjectRequest) (_result *PopCreateLivePortraitModelScopeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreateLivePortraitModelScopeProjectResponse{}
	_body, _err := client.PopCreateLivePortraitModelScopeProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreateMaterialWithOptions(request *PopCreateMaterialRequest, runtime *util.RuntimeOptions) (_result *PopCreateMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckStatus)) {
		body["CheckStatus"] = request.CheckStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["Ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.ListStatus)) {
		body["ListStatus"] = request.ListStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssKey)) {
		body["OssKey"] = request.OssKey
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreateMaterial"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreateMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreateMaterial(request *PopCreateMaterialRequest) (_result *PopCreateMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreateMaterialResponse{}
	_body, _err := client.PopCreateMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreateObjectGenerationProjectWithOptions(request *PopCreateObjectGenerationProjectRequest, runtime *util.RuntimeOptions) (_result *PopCreateObjectGenerationProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizUsage)) {
		body["BizUsage"] = request.BizUsage
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreateObjectGenerationProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreateObjectGenerationProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreateObjectGenerationProject(request *PopCreateObjectGenerationProjectRequest) (_result *PopCreateObjectGenerationProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreateObjectGenerationProjectResponse{}
	_body, _err := client.PopCreateObjectGenerationProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreateObjectProjectWithOptions(request *PopCreateObjectProjectRequest, runtime *util.RuntimeOptions) (_result *PopCreateObjectProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoBuild)) {
		body["AutoBuild"] = request.AutoBuild
	}

	if !tea.BoolValue(util.IsUnset(request.BizUsage)) {
		body["BizUsage"] = request.BizUsage
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSource)) {
		body["CustomSource"] = request.CustomSource
	}

	if !tea.BoolValue(util.IsUnset(request.Dependencies)) {
		body["Dependencies"] = request.Dependencies
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignUid)) {
		body["ForeignUid"] = request.ForeignUid
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RecommendStatus)) {
		body["RecommendStatus"] = request.RecommendStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreateObjectProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreateObjectProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreateObjectProject(request *PopCreateObjectProjectRequest) (_result *PopCreateObjectProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreateObjectProjectResponse{}
	_body, _err := client.PopCreateObjectProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreatePakRenderProjectWithOptions(request *PopCreatePakRenderProjectRequest, runtime *util.RuntimeOptions) (_result *PopCreatePakRenderProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreatePakRenderProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreatePakRenderProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreatePakRenderProject(request *PopCreatePakRenderProjectRequest) (_result *PopCreatePakRenderProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreatePakRenderProjectResponse{}
	_body, _err := client.PopCreatePakRenderProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopCreateTextToAvatarProjectWithOptions(request *PopCreateTextToAvatarProjectRequest, runtime *util.RuntimeOptions) (_result *PopCreateTextToAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["Intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopCreateTextToAvatarProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopCreateTextToAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopCreateTextToAvatarProject(request *PopCreateTextToAvatarProjectRequest) (_result *PopCreateTextToAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopCreateTextToAvatarProjectResponse{}
	_body, _err := client.PopCreateTextToAvatarProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopDeleteMaterialWithOptions(request *PopDeleteMaterialRequest, runtime *util.RuntimeOptions) (_result *PopDeleteMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialId)) {
		query["MaterialId"] = request.MaterialId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopDeleteMaterial"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopDeleteMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopDeleteMaterial(request *PopDeleteMaterialRequest) (_result *PopDeleteMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopDeleteMaterialResponse{}
	_body, _err := client.PopDeleteMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopGetAITryOnJobWithOptions(request *PopGetAITryOnJobRequest, runtime *util.RuntimeOptions) (_result *PopGetAITryOnJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WithMaterial)) {
		query["WithMaterial"] = request.WithMaterial
	}

	if !tea.BoolValue(util.IsUnset(request.WithResult)) {
		query["WithResult"] = request.WithResult
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopGetAITryOnJob"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopGetAITryOnJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopGetAITryOnJob(request *PopGetAITryOnJobRequest) (_result *PopGetAITryOnJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopGetAITryOnJobResponse{}
	_body, _err := client.PopGetAITryOnJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListAITryOnJobsWithOptions(request *PopListAITryOnJobsRequest, runtime *util.RuntimeOptions) (_result *PopListAITryOnJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListAITryOnJobs"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListAITryOnJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListAITryOnJobs(request *PopListAITryOnJobsRequest) (_result *PopListAITryOnJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListAITryOnJobsResponse{}
	_body, _err := client.PopListAITryOnJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListCommonMaterialsAllWithOptions(request *PopListCommonMaterialsAllRequest, runtime *util.RuntimeOptions) (_result *PopListCommonMaterialsAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.ListStatus)) {
		query["ListStatus"] = request.ListStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListCommonMaterialsAll"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListCommonMaterialsAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListCommonMaterialsAll(request *PopListCommonMaterialsAllRequest) (_result *PopListCommonMaterialsAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListCommonMaterialsAllResponse{}
	_body, _err := client.PopListCommonMaterialsAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListFeatureToAvatarMaterialsWithOptions(request *PopListFeatureToAvatarMaterialsRequest, runtime *util.RuntimeOptions) (_result *PopListFeatureToAvatarMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.ListStatus)) {
		body["ListStatus"] = request.ListStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListFeatureToAvatarMaterials"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListFeatureToAvatarMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListFeatureToAvatarMaterials(request *PopListFeatureToAvatarMaterialsRequest) (_result *PopListFeatureToAvatarMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListFeatureToAvatarMaterialsResponse{}
	_body, _err := client.PopListFeatureToAvatarMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListFeatureToAvatarProjectWithOptions(request *PopListFeatureToAvatarProjectRequest, runtime *util.RuntimeOptions) (_result *PopListFeatureToAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		body["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListFeatureToAvatarProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListFeatureToAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListFeatureToAvatarProject(request *PopListFeatureToAvatarProjectRequest) (_result *PopListFeatureToAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListFeatureToAvatarProjectResponse{}
	_body, _err := client.PopListFeatureToAvatarProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListLivePortraitModelScopeMaterialsWithOptions(request *PopListLivePortraitModelScopeMaterialsRequest, runtime *util.RuntimeOptions) (_result *PopListLivePortraitModelScopeMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		body["Types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListLivePortraitModelScopeMaterials"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListLivePortraitModelScopeMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListLivePortraitModelScopeMaterials(request *PopListLivePortraitModelScopeMaterialsRequest) (_result *PopListLivePortraitModelScopeMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListLivePortraitModelScopeMaterialsResponse{}
	_body, _err := client.PopListLivePortraitModelScopeMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListObjectCaseWithOptions(request *PopListObjectCaseRequest, runtime *util.RuntimeOptions) (_result *PopListObjectCaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListObjectCase"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListObjectCaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListObjectCase(request *PopListObjectCaseRequest) (_result *PopListObjectCaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListObjectCaseResponse{}
	_body, _err := client.PopListObjectCaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListObjectGenerationProjectWithOptions(request *PopListObjectGenerationProjectRequest, runtime *util.RuntimeOptions) (_result *PopListObjectGenerationProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListObjectGenerationProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListObjectGenerationProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListObjectGenerationProject(request *PopListObjectGenerationProjectRequest) (_result *PopListObjectGenerationProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListObjectGenerationProjectResponse{}
	_body, _err := client.PopListObjectGenerationProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListObjectProjectWithOptions(request *PopListObjectProjectRequest, runtime *util.RuntimeOptions) (_result *PopListObjectProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		body["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSource)) {
		body["CustomSource"] = request.CustomSource
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		body["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.WithSource)) {
		body["WithSource"] = request.WithSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListObjectProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListObjectProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListObjectProject(request *PopListObjectProjectRequest) (_result *PopListObjectProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListObjectProjectResponse{}
	_body, _err := client.PopListObjectProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListPakRenderExpressionWithOptions(request *PopListPakRenderExpressionRequest, runtime *util.RuntimeOptions) (_result *PopListPakRenderExpressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListStatus)) {
		query["ListStatus"] = request.ListStatus
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListPakRenderExpression"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListPakRenderExpressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListPakRenderExpression(request *PopListPakRenderExpressionRequest) (_result *PopListPakRenderExpressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListPakRenderExpressionResponse{}
	_body, _err := client.PopListPakRenderExpressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopListTextToAvatarProjectWithOptions(request *PopListTextToAvatarProjectRequest, runtime *util.RuntimeOptions) (_result *PopListTextToAvatarProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		body["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopListTextToAvatarProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopListTextToAvatarProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopListTextToAvatarProject(request *PopListTextToAvatarProjectRequest) (_result *PopListTextToAvatarProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopListTextToAvatarProjectResponse{}
	_body, _err := client.PopListTextToAvatarProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopObjectProjectDetailWithOptions(request *PopObjectProjectDetailRequest, runtime *util.RuntimeOptions) (_result *PopObjectProjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopObjectProjectDetail"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopObjectProjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopObjectProjectDetail(request *PopObjectProjectDetailRequest) (_result *PopObjectProjectDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopObjectProjectDetailResponse{}
	_body, _err := client.PopObjectProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopObjectRetrievalWithOptions(request *PopObjectRetrievalRequest, runtime *util.RuntimeOptions) (_result *PopObjectRetrievalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["TopK"] = request.TopK
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopObjectRetrieval"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopObjectRetrievalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopObjectRetrieval(request *PopObjectRetrievalRequest) (_result *PopObjectRetrievalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopObjectRetrievalResponse{}
	_body, _err := client.PopObjectRetrievalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopObjectRetrievalUploadDataWithOptions(request *PopObjectRetrievalUploadDataRequest, runtime *util.RuntimeOptions) (_result *PopObjectRetrievalUploadDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopObjectRetrievalUploadData"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopObjectRetrievalUploadDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopObjectRetrievalUploadData(request *PopObjectRetrievalUploadDataRequest) (_result *PopObjectRetrievalUploadDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopObjectRetrievalUploadDataResponse{}
	_body, _err := client.PopObjectRetrievalUploadDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopQueryAvatarProjectDetailWithOptions(request *PopQueryAvatarProjectDetailRequest, runtime *util.RuntimeOptions) (_result *PopQueryAvatarProjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopQueryAvatarProjectDetail"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopQueryAvatarProjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopQueryAvatarProjectDetail(request *PopQueryAvatarProjectDetailRequest) (_result *PopQueryAvatarProjectDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopQueryAvatarProjectDetailResponse{}
	_body, _err := client.PopQueryAvatarProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopQueryLatestAvatarProjectDetailByUserWithOptions(request *PopQueryLatestAvatarProjectDetailByUserRequest, runtime *util.RuntimeOptions) (_result *PopQueryLatestAvatarProjectDetailByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopQueryLatestAvatarProjectDetailByUser"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopQueryLatestAvatarProjectDetailByUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopQueryLatestAvatarProjectDetailByUser(request *PopQueryLatestAvatarProjectDetailByUserRequest) (_result *PopQueryLatestAvatarProjectDetailByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopQueryLatestAvatarProjectDetailByUserResponse{}
	_body, _err := client.PopQueryLatestAvatarProjectDetailByUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopQueryLivePortraitModelScopeProjectDetailWithOptions(request *PopQueryLivePortraitModelScopeProjectDetailRequest, runtime *util.RuntimeOptions) (_result *PopQueryLivePortraitModelScopeProjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopQueryLivePortraitModelScopeProjectDetail"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopQueryLivePortraitModelScopeProjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopQueryLivePortraitModelScopeProjectDetail(request *PopQueryLivePortraitModelScopeProjectDetailRequest) (_result *PopQueryLivePortraitModelScopeProjectDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopQueryLivePortraitModelScopeProjectDetailResponse{}
	_body, _err := client.PopQueryLivePortraitModelScopeProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopQueryObjectGenerationProjectDetailWithOptions(request *PopQueryObjectGenerationProjectDetailRequest, runtime *util.RuntimeOptions) (_result *PopQueryObjectGenerationProjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopQueryObjectGenerationProjectDetail"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopQueryObjectGenerationProjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopQueryObjectGenerationProjectDetail(request *PopQueryObjectGenerationProjectDetailRequest) (_result *PopQueryObjectGenerationProjectDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopQueryObjectGenerationProjectDetailResponse{}
	_body, _err := client.PopQueryObjectGenerationProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopRetryAITryOnTaskWithOptions(request *PopRetryAITryOnTaskRequest, runtime *util.RuntimeOptions) (_result *PopRetryAITryOnTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopRetryAITryOnTask"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopRetryAITryOnTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopRetryAITryOnTask(request *PopRetryAITryOnTaskRequest) (_result *PopRetryAITryOnTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopRetryAITryOnTaskResponse{}
	_body, _err := client.PopRetryAITryOnTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopSubmitAITryOnJobWithOptions(request *PopSubmitAITryOnJobRequest, runtime *util.RuntimeOptions) (_result *PopSubmitAITryOnJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BottomsId)) {
		query["BottomsId"] = request.BottomsId
	}

	if !tea.BoolValue(util.IsUnset(request.ClothingType)) {
		query["ClothingType"] = request.ClothingType
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ShoeType)) {
		query["ShoeType"] = request.ShoeType
	}

	if !tea.BoolValue(util.IsUnset(request.SuitId)) {
		query["SuitId"] = request.SuitId
	}

	if !tea.BoolValue(util.IsUnset(request.TopsId)) {
		query["TopsId"] = request.TopsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopSubmitAITryOnJob"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopSubmitAITryOnJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopSubmitAITryOnJob(request *PopSubmitAITryOnJobRequest) (_result *PopSubmitAITryOnJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopSubmitAITryOnJobResponse{}
	_body, _err := client.PopSubmitAITryOnJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopUploadMaterialWithOptions(request *PopUploadMaterialRequest, runtime *util.RuntimeOptions) (_result *PopUploadMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PopUploadMaterial"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopUploadMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopUploadMaterial(request *PopUploadMaterialRequest) (_result *PopUploadMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopUploadMaterialResponse{}
	_body, _err := client.PopUploadMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PopVideoSaveSourceWithOptions(request *PopVideoSaveSourceRequest, runtime *util.RuntimeOptions) (_result *PopVideoSaveSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PopVideoSaveSource"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PopVideoSaveSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PopVideoSaveSource(request *PopVideoSaveSourceRequest) (_result *PopVideoSaveSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PopVideoSaveSourceResponse{}
	_body, _err := client.PopVideoSaveSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDigitalHumanProjectWithOptions(request *QueryDigitalHumanProjectRequest, runtime *util.RuntimeOptions) (_result *QueryDigitalHumanProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["Ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDigitalHumanProject"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDigitalHumanProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDigitalHumanProject(request *QueryDigitalHumanProjectRequest) (_result *QueryDigitalHumanProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDigitalHumanProjectResponse{}
	_body, _err := client.QueryDigitalHumanProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryLongTtsResultWithOptions(request *QueryLongTtsResultRequest, runtime *util.RuntimeOptions) (_result *QueryLongTtsResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryLongTtsResult"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLongTtsResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLongTtsResult(request *QueryLongTtsResultRequest) (_result *QueryLongTtsResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLongTtsResultResponse{}
	_body, _err := client.QueryLongTtsResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMotionShopVideoDetectResultWithOptions(request *QueryMotionShopVideoDetectResultRequest, runtime *util.RuntimeOptions) (_result *QueryMotionShopVideoDetectResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMotionShopVideoDetectResult"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMotionShopVideoDetectResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMotionShopVideoDetectResult(request *QueryMotionShopVideoDetectResultRequest) (_result *QueryMotionShopVideoDetectResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMotionShopVideoDetectResultResponse{}
	_body, _err := client.QueryMotionShopVideoDetectResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitLongTtsTaskWithOptions(request *SubmitLongTtsTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitLongTtsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.TtsVoiceId)) {
		body["TtsVoiceId"] = request.TtsVoiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitLongTtsTask"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitLongTtsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitLongTtsTask(request *SubmitLongTtsTaskRequest) (_result *SubmitLongTtsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitLongTtsTaskResponse{}
	_body, _err := client.SubmitLongTtsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitMotionShopTaskWithOptions(request *SubmitMotionShopTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitMotionShopTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvatarId)) {
		body["AvatarId"] = request.AvatarId
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.VideoId)) {
		body["VideoId"] = request.VideoId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitMotionShopTask"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitMotionShopTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitMotionShopTask(request *SubmitMotionShopTaskRequest) (_result *SubmitMotionShopTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitMotionShopTaskResponse{}
	_body, _err := client.SubmitMotionShopTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserEmailWithOptions(request *UpdateUserEmailRequest, runtime *util.RuntimeOptions) (_result *UpdateUserEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.JwtToken)) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserEmail"),
		Version:     tea.String("2023-03-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserEmail(request *UpdateUserEmailRequest) (_result *UpdateUserEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserEmailResponse{}
	_body, _err := client.UpdateUserEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
