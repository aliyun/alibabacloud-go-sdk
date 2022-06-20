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

type AePredictCategoryRequest struct {
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s AePredictCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AePredictCategoryRequest) GoString() string {
	return s.String()
}

func (s *AePredictCategoryRequest) SetPicUrl(v string) *AePredictCategoryRequest {
	s.PicUrl = &v
	return s
}

type AePredictCategoryAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
}

func (s AePredictCategoryAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AePredictCategoryAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AePredictCategoryAdvanceRequest) SetPicUrlObject(v io.Reader) *AePredictCategoryAdvanceRequest {
	s.PicUrlObject = v
	return s
}

type AePredictCategoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AePredictCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AePredictCategoryResponse) GoString() string {
	return s.String()
}

func (s *AePredictCategoryResponse) SetHeaders(v map[string]*string) *AePredictCategoryResponse {
	s.Headers = v
	return s
}

func (s *AePredictCategoryResponse) SetStatusCode(v int32) *AePredictCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AePredictCategoryResponse) SetBody(v map[string]interface{}) *AePredictCategoryResponse {
	s.Body = v
	return s
}

type AePropRecRequest struct {
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s AePropRecRequest) String() string {
	return tea.Prettify(s)
}

func (s AePropRecRequest) GoString() string {
	return s.String()
}

func (s *AePropRecRequest) SetPicUrl(v string) *AePropRecRequest {
	s.PicUrl = &v
	return s
}

type AePropRecAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
}

func (s AePropRecAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AePropRecAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AePropRecAdvanceRequest) SetPicUrlObject(v io.Reader) *AePropRecAdvanceRequest {
	s.PicUrlObject = v
	return s
}

type AePropRecResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AePropRecResponse) String() string {
	return tea.Prettify(s)
}

func (s AePropRecResponse) GoString() string {
	return s.String()
}

func (s *AePropRecResponse) SetHeaders(v map[string]*string) *AePropRecResponse {
	s.Headers = v
	return s
}

func (s *AePropRecResponse) SetStatusCode(v int32) *AePropRecResponse {
	s.StatusCode = &v
	return s
}

func (s *AePropRecResponse) SetBody(v map[string]interface{}) *AePropRecResponse {
	s.Body = v
	return s
}

type AlivisionImgdupRequest struct {
	ImageHeight    *int32  `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageWidth     *int32  `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	OutputImageNum *int32  `json:"OutputImageNum,omitempty" xml:"OutputImageNum,omitempty"`
	PicNumList     *string `json:"PicNumList,omitempty" xml:"PicNumList,omitempty"`
	PicUrlList     *string `json:"PicUrlList,omitempty" xml:"PicUrlList,omitempty"`
}

func (s AlivisionImgdupRequest) String() string {
	return tea.Prettify(s)
}

func (s AlivisionImgdupRequest) GoString() string {
	return s.String()
}

func (s *AlivisionImgdupRequest) SetImageHeight(v int32) *AlivisionImgdupRequest {
	s.ImageHeight = &v
	return s
}

func (s *AlivisionImgdupRequest) SetImageWidth(v int32) *AlivisionImgdupRequest {
	s.ImageWidth = &v
	return s
}

func (s *AlivisionImgdupRequest) SetOutputImageNum(v int32) *AlivisionImgdupRequest {
	s.OutputImageNum = &v
	return s
}

func (s *AlivisionImgdupRequest) SetPicNumList(v string) *AlivisionImgdupRequest {
	s.PicNumList = &v
	return s
}

func (s *AlivisionImgdupRequest) SetPicUrlList(v string) *AlivisionImgdupRequest {
	s.PicUrlList = &v
	return s
}

type AlivisionImgdupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AlivisionImgdupResponse) String() string {
	return tea.Prettify(s)
}

func (s AlivisionImgdupResponse) GoString() string {
	return s.String()
}

func (s *AlivisionImgdupResponse) SetHeaders(v map[string]*string) *AlivisionImgdupResponse {
	s.Headers = v
	return s
}

func (s *AlivisionImgdupResponse) SetStatusCode(v int32) *AlivisionImgdupResponse {
	s.StatusCode = &v
	return s
}

func (s *AlivisionImgdupResponse) SetBody(v map[string]interface{}) *AlivisionImgdupResponse {
	s.Body = v
	return s
}

type CreateImageAmazonTaskRequest struct {
	Gif          *bool     `json:"Gif,omitempty" xml:"Gif,omitempty"`
	ImgUrlList   []*string `json:"ImgUrlList,omitempty" xml:"ImgUrlList,omitempty" type:"Repeated"`
	TemplateMode *string   `json:"TemplateMode,omitempty" xml:"TemplateMode,omitempty"`
	TextList     []*string `json:"TextList,omitempty" xml:"TextList,omitempty" type:"Repeated"`
}

func (s CreateImageAmazonTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskRequest) SetGif(v bool) *CreateImageAmazonTaskRequest {
	s.Gif = &v
	return s
}

func (s *CreateImageAmazonTaskRequest) SetImgUrlList(v []*string) *CreateImageAmazonTaskRequest {
	s.ImgUrlList = v
	return s
}

func (s *CreateImageAmazonTaskRequest) SetTemplateMode(v string) *CreateImageAmazonTaskRequest {
	s.TemplateMode = &v
	return s
}

func (s *CreateImageAmazonTaskRequest) SetTextList(v []*string) *CreateImageAmazonTaskRequest {
	s.TextList = v
	return s
}

type CreateImageAmazonTaskShrinkRequest struct {
	Gif              *bool   `json:"Gif,omitempty" xml:"Gif,omitempty"`
	ImgUrlListShrink *string `json:"ImgUrlList,omitempty" xml:"ImgUrlList,omitempty"`
	TemplateMode     *string `json:"TemplateMode,omitempty" xml:"TemplateMode,omitempty"`
	TextListShrink   *string `json:"TextList,omitempty" xml:"TextList,omitempty"`
}

func (s CreateImageAmazonTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskShrinkRequest) SetGif(v bool) *CreateImageAmazonTaskShrinkRequest {
	s.Gif = &v
	return s
}

func (s *CreateImageAmazonTaskShrinkRequest) SetImgUrlListShrink(v string) *CreateImageAmazonTaskShrinkRequest {
	s.ImgUrlListShrink = &v
	return s
}

func (s *CreateImageAmazonTaskShrinkRequest) SetTemplateMode(v string) *CreateImageAmazonTaskShrinkRequest {
	s.TemplateMode = &v
	return s
}

func (s *CreateImageAmazonTaskShrinkRequest) SetTextListShrink(v string) *CreateImageAmazonTaskShrinkRequest {
	s.TextListShrink = &v
	return s
}

type CreateImageAmazonTaskResponseBody struct {
	Code    *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s CreateImageAmazonTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskResponseBody) SetCode(v int64) *CreateImageAmazonTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetData(v int64) *CreateImageAmazonTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetMessage(v string) *CreateImageAmazonTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetRequestId(v string) *CreateImageAmazonTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageAmazonTaskResponseBody) SetSuccessResponse(v bool) *CreateImageAmazonTaskResponseBody {
	s.SuccessResponse = &v
	return s
}

type CreateImageAmazonTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageAmazonTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageAmazonTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageAmazonTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageAmazonTaskResponse) SetHeaders(v map[string]*string) *CreateImageAmazonTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageAmazonTaskResponse) SetStatusCode(v int32) *CreateImageAmazonTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageAmazonTaskResponse) SetBody(v *CreateImageAmazonTaskResponseBody) *CreateImageAmazonTaskResponse {
	s.Body = v
	return s
}

type CreateRemoveWorkTaskRequest struct {
	ItemIdentity *string `json:"ItemIdentity,omitempty" xml:"ItemIdentity,omitempty"`
	PicUrl       *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s CreateRemoveWorkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoveWorkTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRemoveWorkTaskRequest) SetItemIdentity(v string) *CreateRemoveWorkTaskRequest {
	s.ItemIdentity = &v
	return s
}

func (s *CreateRemoveWorkTaskRequest) SetPicUrl(v string) *CreateRemoveWorkTaskRequest {
	s.PicUrl = &v
	return s
}

type CreateRemoveWorkTaskAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
	ItemIdentity *string   `json:"ItemIdentity,omitempty" xml:"ItemIdentity,omitempty"`
}

func (s CreateRemoveWorkTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoveWorkTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateRemoveWorkTaskAdvanceRequest) SetPicUrlObject(v io.Reader) *CreateRemoveWorkTaskAdvanceRequest {
	s.PicUrlObject = v
	return s
}

func (s *CreateRemoveWorkTaskAdvanceRequest) SetItemIdentity(v string) *CreateRemoveWorkTaskAdvanceRequest {
	s.ItemIdentity = &v
	return s
}

type CreateRemoveWorkTaskResponseBody struct {
	Code    *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s CreateRemoveWorkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoveWorkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemoveWorkTaskResponseBody) SetCode(v int64) *CreateRemoveWorkTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRemoveWorkTaskResponseBody) SetData(v int64) *CreateRemoveWorkTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRemoveWorkTaskResponseBody) SetMessage(v string) *CreateRemoveWorkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRemoveWorkTaskResponseBody) SetRequestId(v string) *CreateRemoveWorkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRemoveWorkTaskResponseBody) SetSuccessResponse(v bool) *CreateRemoveWorkTaskResponseBody {
	s.SuccessResponse = &v
	return s
}

type CreateRemoveWorkTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRemoveWorkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRemoveWorkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoveWorkTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRemoveWorkTaskResponse) SetHeaders(v map[string]*string) *CreateRemoveWorkTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRemoveWorkTaskResponse) SetStatusCode(v int32) *CreateRemoveWorkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRemoveWorkTaskResponse) SetBody(v *CreateRemoveWorkTaskResponseBody) *CreateRemoveWorkTaskResponse {
	s.Body = v
	return s
}

type FaceshifterTRequest struct {
	Age    *int32  `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender *int32  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Race   *int32  `json:"Race,omitempty" xml:"Race,omitempty"`
}

func (s FaceshifterTRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceshifterTRequest) GoString() string {
	return s.String()
}

func (s *FaceshifterTRequest) SetAge(v int32) *FaceshifterTRequest {
	s.Age = &v
	return s
}

func (s *FaceshifterTRequest) SetGender(v int32) *FaceshifterTRequest {
	s.Gender = &v
	return s
}

func (s *FaceshifterTRequest) SetPicUrl(v string) *FaceshifterTRequest {
	s.PicUrl = &v
	return s
}

func (s *FaceshifterTRequest) SetRace(v int32) *FaceshifterTRequest {
	s.Race = &v
	return s
}

type FaceshifterTAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
	Age          *int32    `json:"Age,omitempty" xml:"Age,omitempty"`
	Gender       *int32    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Race         *int32    `json:"Race,omitempty" xml:"Race,omitempty"`
}

func (s FaceshifterTAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceshifterTAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceshifterTAdvanceRequest) SetPicUrlObject(v io.Reader) *FaceshifterTAdvanceRequest {
	s.PicUrlObject = v
	return s
}

func (s *FaceshifterTAdvanceRequest) SetAge(v int32) *FaceshifterTAdvanceRequest {
	s.Age = &v
	return s
}

func (s *FaceshifterTAdvanceRequest) SetGender(v int32) *FaceshifterTAdvanceRequest {
	s.Gender = &v
	return s
}

func (s *FaceshifterTAdvanceRequest) SetRace(v int32) *FaceshifterTAdvanceRequest {
	s.Race = &v
	return s
}

type FaceshifterTResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceshifterTResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceshifterTResponse) GoString() string {
	return s.String()
}

func (s *FaceshifterTResponse) SetHeaders(v map[string]*string) *FaceshifterTResponse {
	s.Headers = v
	return s
}

func (s *FaceshifterTResponse) SetStatusCode(v int32) *FaceshifterTResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceshifterTResponse) SetBody(v map[string]interface{}) *FaceshifterTResponse {
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
	Code    *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) SetCode(v int64) *GetTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResultResponseBody) SetData(v *GetTaskResultResponseBodyData) *GetTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResultResponseBody) SetMessage(v string) *GetTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetSuccessResponse(v bool) *GetTaskResultResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTaskResultResponseBodyData struct {
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyData) SetResult(v string) *GetTaskResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetStatus(v int64) *GetTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetStatusName(v string) *GetTaskResultResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskId(v int64) *GetTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

type GetTaskResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskResultResponse) SetStatusCode(v int32) *GetTaskResultResponse {
	s.StatusCode = &v
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
	Code    *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetCode(v int64) *GetTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetData(v *GetTaskStatusResponseBodyData) *GetTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccessResponse(v bool) *GetTaskStatusResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTaskStatusResponseBodyData struct {
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBodyData) SetStatus(v int64) *GetTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBodyData) SetStatusName(v string) *GetTaskStatusResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *GetTaskStatusResponseBodyData) SetTaskId(v int64) *GetTaskStatusResponseBodyData {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type KuajingSegRequest struct {
	PicUrl          *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ReturnPicFormat *string `json:"ReturnPicFormat,omitempty" xml:"ReturnPicFormat,omitempty"`
	ReturnPicType   *string `json:"ReturnPicType,omitempty" xml:"ReturnPicType,omitempty"`
}

func (s KuajingSegRequest) String() string {
	return tea.Prettify(s)
}

func (s KuajingSegRequest) GoString() string {
	return s.String()
}

func (s *KuajingSegRequest) SetPicUrl(v string) *KuajingSegRequest {
	s.PicUrl = &v
	return s
}

func (s *KuajingSegRequest) SetReturnPicFormat(v string) *KuajingSegRequest {
	s.ReturnPicFormat = &v
	return s
}

func (s *KuajingSegRequest) SetReturnPicType(v string) *KuajingSegRequest {
	s.ReturnPicType = &v
	return s
}

type KuajingSegAdvanceRequest struct {
	PicUrlObject    io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
	ReturnPicFormat *string   `json:"ReturnPicFormat,omitempty" xml:"ReturnPicFormat,omitempty"`
	ReturnPicType   *string   `json:"ReturnPicType,omitempty" xml:"ReturnPicType,omitempty"`
}

func (s KuajingSegAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s KuajingSegAdvanceRequest) GoString() string {
	return s.String()
}

func (s *KuajingSegAdvanceRequest) SetPicUrlObject(v io.Reader) *KuajingSegAdvanceRequest {
	s.PicUrlObject = v
	return s
}

func (s *KuajingSegAdvanceRequest) SetReturnPicFormat(v string) *KuajingSegAdvanceRequest {
	s.ReturnPicFormat = &v
	return s
}

func (s *KuajingSegAdvanceRequest) SetReturnPicType(v string) *KuajingSegAdvanceRequest {
	s.ReturnPicType = &v
	return s
}

type KuajingSegResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KuajingSegResponse) String() string {
	return tea.Prettify(s)
}

func (s KuajingSegResponse) GoString() string {
	return s.String()
}

func (s *KuajingSegResponse) SetHeaders(v map[string]*string) *KuajingSegResponse {
	s.Headers = v
	return s
}

func (s *KuajingSegResponse) SetStatusCode(v int32) *KuajingSegResponse {
	s.StatusCode = &v
	return s
}

func (s *KuajingSegResponse) SetBody(v map[string]interface{}) *KuajingSegResponse {
	s.Body = v
	return s
}

type RemoveWordsRequest struct {
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s RemoveWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWordsRequest) GoString() string {
	return s.String()
}

func (s *RemoveWordsRequest) SetPicUrl(v string) *RemoveWordsRequest {
	s.PicUrl = &v
	return s
}

type RemoveWordsAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
}

func (s RemoveWordsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWordsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveWordsAdvanceRequest) SetPicUrlObject(v io.Reader) *RemoveWordsAdvanceRequest {
	s.PicUrlObject = v
	return s
}

type RemoveWordsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWordsResponse) GoString() string {
	return s.String()
}

func (s *RemoveWordsResponse) SetHeaders(v map[string]*string) *RemoveWordsResponse {
	s.Headers = v
	return s
}

func (s *RemoveWordsResponse) SetStatusCode(v int32) *RemoveWordsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWordsResponse) SetBody(v map[string]interface{}) *RemoveWordsResponse {
	s.Body = v
	return s
}

type ReplaceBackgroundRequest struct {
	// 返回的图片背景图片ID
	BackgroundId     *string `json:"BackgroundId,omitempty" xml:"BackgroundId,omitempty"`
	Num              *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicBackgroundUrl *string `json:"PicBackgroundUrl,omitempty" xml:"PicBackgroundUrl,omitempty"`
	// 图片地址
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s ReplaceBackgroundRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBackgroundRequest) GoString() string {
	return s.String()
}

func (s *ReplaceBackgroundRequest) SetBackgroundId(v string) *ReplaceBackgroundRequest {
	s.BackgroundId = &v
	return s
}

func (s *ReplaceBackgroundRequest) SetNum(v int32) *ReplaceBackgroundRequest {
	s.Num = &v
	return s
}

func (s *ReplaceBackgroundRequest) SetPicBackgroundUrl(v string) *ReplaceBackgroundRequest {
	s.PicBackgroundUrl = &v
	return s
}

func (s *ReplaceBackgroundRequest) SetPicUrl(v string) *ReplaceBackgroundRequest {
	s.PicUrl = &v
	return s
}

type ReplaceBackgroundAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
	// 返回的图片背景图片ID
	BackgroundId     *string `json:"BackgroundId,omitempty" xml:"BackgroundId,omitempty"`
	Num              *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	PicBackgroundUrl *string `json:"PicBackgroundUrl,omitempty" xml:"PicBackgroundUrl,omitempty"`
}

func (s ReplaceBackgroundAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBackgroundAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ReplaceBackgroundAdvanceRequest) SetPicUrlObject(v io.Reader) *ReplaceBackgroundAdvanceRequest {
	s.PicUrlObject = v
	return s
}

func (s *ReplaceBackgroundAdvanceRequest) SetBackgroundId(v string) *ReplaceBackgroundAdvanceRequest {
	s.BackgroundId = &v
	return s
}

func (s *ReplaceBackgroundAdvanceRequest) SetNum(v int32) *ReplaceBackgroundAdvanceRequest {
	s.Num = &v
	return s
}

func (s *ReplaceBackgroundAdvanceRequest) SetPicBackgroundUrl(v string) *ReplaceBackgroundAdvanceRequest {
	s.PicBackgroundUrl = &v
	return s
}

type ReplaceBackgroundResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceBackgroundResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBackgroundResponse) GoString() string {
	return s.String()
}

func (s *ReplaceBackgroundResponse) SetHeaders(v map[string]*string) *ReplaceBackgroundResponse {
	s.Headers = v
	return s
}

func (s *ReplaceBackgroundResponse) SetStatusCode(v int32) *ReplaceBackgroundResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceBackgroundResponse) SetBody(v map[string]interface{}) *ReplaceBackgroundResponse {
	s.Body = v
	return s
}

type SeleteCommodityRequest struct {
	Num   *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Start *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SeleteCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s SeleteCommodityRequest) GoString() string {
	return s.String()
}

func (s *SeleteCommodityRequest) SetNum(v int32) *SeleteCommodityRequest {
	s.Num = &v
	return s
}

func (s *SeleteCommodityRequest) SetPid(v string) *SeleteCommodityRequest {
	s.Pid = &v
	return s
}

func (s *SeleteCommodityRequest) SetQuery(v string) *SeleteCommodityRequest {
	s.Query = &v
	return s
}

func (s *SeleteCommodityRequest) SetStart(v int32) *SeleteCommodityRequest {
	s.Start = &v
	return s
}

type SeleteCommodityResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SeleteCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s SeleteCommodityResponse) GoString() string {
	return s.String()
}

func (s *SeleteCommodityResponse) SetHeaders(v map[string]*string) *SeleteCommodityResponse {
	s.Headers = v
	return s
}

func (s *SeleteCommodityResponse) SetStatusCode(v int32) *SeleteCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *SeleteCommodityResponse) SetBody(v map[string]interface{}) *SeleteCommodityResponse {
	s.Body = v
	return s
}

type SeleteCommodityByBToBRequest struct {
	Num   *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Start *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SeleteCommodityByBToBRequest) String() string {
	return tea.Prettify(s)
}

func (s SeleteCommodityByBToBRequest) GoString() string {
	return s.String()
}

func (s *SeleteCommodityByBToBRequest) SetNum(v int32) *SeleteCommodityByBToBRequest {
	s.Num = &v
	return s
}

func (s *SeleteCommodityByBToBRequest) SetPid(v string) *SeleteCommodityByBToBRequest {
	s.Pid = &v
	return s
}

func (s *SeleteCommodityByBToBRequest) SetQuery(v string) *SeleteCommodityByBToBRequest {
	s.Query = &v
	return s
}

func (s *SeleteCommodityByBToBRequest) SetStart(v int32) *SeleteCommodityByBToBRequest {
	s.Start = &v
	return s
}

type SeleteCommodityByBToBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SeleteCommodityByBToBResponse) String() string {
	return tea.Prettify(s)
}

func (s SeleteCommodityByBToBResponse) GoString() string {
	return s.String()
}

func (s *SeleteCommodityByBToBResponse) SetHeaders(v map[string]*string) *SeleteCommodityByBToBResponse {
	s.Headers = v
	return s
}

func (s *SeleteCommodityByBToBResponse) SetStatusCode(v int32) *SeleteCommodityByBToBResponse {
	s.StatusCode = &v
	return s
}

func (s *SeleteCommodityByBToBResponse) SetBody(v map[string]interface{}) *SeleteCommodityByBToBResponse {
	s.Body = v
	return s
}

type TbPredictCategoryRequest struct {
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s TbPredictCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s TbPredictCategoryRequest) GoString() string {
	return s.String()
}

func (s *TbPredictCategoryRequest) SetPicUrl(v string) *TbPredictCategoryRequest {
	s.PicUrl = &v
	return s
}

type TbPredictCategoryAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
}

func (s TbPredictCategoryAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TbPredictCategoryAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TbPredictCategoryAdvanceRequest) SetPicUrlObject(v io.Reader) *TbPredictCategoryAdvanceRequest {
	s.PicUrlObject = v
	return s
}

type TbPredictCategoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TbPredictCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s TbPredictCategoryResponse) GoString() string {
	return s.String()
}

func (s *TbPredictCategoryResponse) SetHeaders(v map[string]*string) *TbPredictCategoryResponse {
	s.Headers = v
	return s
}

func (s *TbPredictCategoryResponse) SetStatusCode(v int32) *TbPredictCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *TbPredictCategoryResponse) SetBody(v map[string]interface{}) *TbPredictCategoryResponse {
	s.Body = v
	return s
}

type TbPropRecRequest struct {
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s TbPropRecRequest) String() string {
	return tea.Prettify(s)
}

func (s TbPropRecRequest) GoString() string {
	return s.String()
}

func (s *TbPropRecRequest) SetPicUrl(v string) *TbPropRecRequest {
	s.PicUrl = &v
	return s
}

type TbPropRecAdvanceRequest struct {
	PicUrlObject io.Reader `json:"PicUrlObject,omitempty" xml:"PicUrlObject,omitempty" require:"true"`
}

func (s TbPropRecAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TbPropRecAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TbPropRecAdvanceRequest) SetPicUrlObject(v io.Reader) *TbPropRecAdvanceRequest {
	s.PicUrlObject = v
	return s
}

type TbPropRecResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TbPropRecResponse) String() string {
	return tea.Prettify(s)
}

func (s TbPropRecResponse) GoString() string {
	return s.String()
}

func (s *TbPropRecResponse) SetHeaders(v map[string]*string) *TbPropRecResponse {
	s.Headers = v
	return s
}

func (s *TbPropRecResponse) SetStatusCode(v int32) *TbPropRecResponse {
	s.StatusCode = &v
	return s
}

func (s *TbPropRecResponse) SetBody(v map[string]interface{}) *TbPropRecResponse {
	s.Body = v
	return s
}

type TransferUrlByBtoBRequest struct {
	OfferId *int64  `json:"OfferId,omitempty" xml:"OfferId,omitempty"`
	Pid     *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s TransferUrlByBtoBRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferUrlByBtoBRequest) GoString() string {
	return s.String()
}

func (s *TransferUrlByBtoBRequest) SetOfferId(v int64) *TransferUrlByBtoBRequest {
	s.OfferId = &v
	return s
}

func (s *TransferUrlByBtoBRequest) SetPid(v string) *TransferUrlByBtoBRequest {
	s.Pid = &v
	return s
}

type TransferUrlByBtoBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferUrlByBtoBResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferUrlByBtoBResponse) GoString() string {
	return s.String()
}

func (s *TransferUrlByBtoBResponse) SetHeaders(v map[string]*string) *TransferUrlByBtoBResponse {
	s.Headers = v
	return s
}

func (s *TransferUrlByBtoBResponse) SetStatusCode(v int32) *TransferUrlByBtoBResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferUrlByBtoBResponse) SetBody(v map[string]interface{}) *TransferUrlByBtoBResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AePredictCategoryWithOptions(request *AePredictCategoryRequest, runtime *util.RuntimeOptions) (_result *AePredictCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AePredictCategory"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AePredictCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AePredictCategory(request *AePredictCategoryRequest) (_result *AePredictCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AePredictCategoryResponse{}
	_body, _err := client.AePredictCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AePredictCategoryAdvance(request *AePredictCategoryAdvanceRequest, runtime *util.RuntimeOptions) (_result *AePredictCategoryResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	aePredictCategoryReq := &AePredictCategoryRequest{}
	openapiutil.Convert(request, aePredictCategoryReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		aePredictCategoryReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	aePredictCategoryResp, _err := client.AePredictCategoryWithOptions(aePredictCategoryReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = aePredictCategoryResp
	return _result, _err
}

func (client *Client) AePropRecWithOptions(request *AePropRecRequest, runtime *util.RuntimeOptions) (_result *AePropRecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AePropRec"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AePropRecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AePropRec(request *AePropRecRequest) (_result *AePropRecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AePropRecResponse{}
	_body, _err := client.AePropRecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AePropRecAdvance(request *AePropRecAdvanceRequest, runtime *util.RuntimeOptions) (_result *AePropRecResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	aePropRecReq := &AePropRecRequest{}
	openapiutil.Convert(request, aePropRecReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		aePropRecReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	aePropRecResp, _err := client.AePropRecWithOptions(aePropRecReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = aePropRecResp
	return _result, _err
}

func (client *Client) AlivisionImgdupWithOptions(request *AlivisionImgdupRequest, runtime *util.RuntimeOptions) (_result *AlivisionImgdupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageHeight)) {
		body["ImageHeight"] = request.ImageHeight
	}

	if !tea.BoolValue(util.IsUnset(request.ImageWidth)) {
		body["ImageWidth"] = request.ImageWidth
	}

	if !tea.BoolValue(util.IsUnset(request.OutputImageNum)) {
		body["OutputImageNum"] = request.OutputImageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PicNumList)) {
		body["PicNumList"] = request.PicNumList
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrlList)) {
		body["PicUrlList"] = request.PicUrlList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AlivisionImgdup"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AlivisionImgdupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AlivisionImgdup(request *AlivisionImgdupRequest) (_result *AlivisionImgdupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AlivisionImgdupResponse{}
	_body, _err := client.AlivisionImgdupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageAmazonTaskWithOptions(tmpReq *CreateImageAmazonTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageAmazonTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateImageAmazonTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImgUrlList)) {
		request.ImgUrlListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImgUrlList, tea.String("ImgUrlList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TextList)) {
		request.TextListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextList, tea.String("TextList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Gif)) {
		query["Gif"] = request.Gif
	}

	if !tea.BoolValue(util.IsUnset(request.ImgUrlListShrink)) {
		query["ImgUrlList"] = request.ImgUrlListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateMode)) {
		query["TemplateMode"] = request.TemplateMode
	}

	if !tea.BoolValue(util.IsUnset(request.TextListShrink)) {
		query["TextList"] = request.TextListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageAmazonTask"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageAmazonTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageAmazonTask(request *CreateImageAmazonTaskRequest) (_result *CreateImageAmazonTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageAmazonTaskResponse{}
	_body, _err := client.CreateImageAmazonTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRemoveWorkTaskWithOptions(request *CreateRemoveWorkTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRemoveWorkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemIdentity)) {
		query["ItemIdentity"] = request.ItemIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		query["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRemoveWorkTask"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRemoveWorkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRemoveWorkTask(request *CreateRemoveWorkTaskRequest) (_result *CreateRemoveWorkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRemoveWorkTaskResponse{}
	_body, _err := client.CreateRemoveWorkTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRemoveWorkTaskAdvance(request *CreateRemoveWorkTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *CreateRemoveWorkTaskResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	createRemoveWorkTaskReq := &CreateRemoveWorkTaskRequest{}
	openapiutil.Convert(request, createRemoveWorkTaskReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		createRemoveWorkTaskReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	createRemoveWorkTaskResp, _err := client.CreateRemoveWorkTaskWithOptions(createRemoveWorkTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createRemoveWorkTaskResp
	return _result, _err
}

func (client *Client) FaceshifterTWithOptions(request *FaceshifterTRequest, runtime *util.RuntimeOptions) (_result *FaceshifterTResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Age)) {
		body["Age"] = request.Age
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["Gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["PicUrl"] = request.PicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Race)) {
		body["Race"] = request.Race
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceshifterT"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceshifterTResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceshifterT(request *FaceshifterTRequest) (_result *FaceshifterTResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceshifterTResponse{}
	_body, _err := client.FaceshifterTWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceshifterTAdvance(request *FaceshifterTAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceshifterTResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	faceshifterTReq := &FaceshifterTRequest{}
	openapiutil.Convert(request, faceshifterTReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		faceshifterTReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	faceshifterTResp, _err := client.FaceshifterTWithOptions(faceshifterTReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceshifterTResp
	return _result, _err
}

func (client *Client) GetTaskResultWithOptions(request *GetTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskResult"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) KuajingSegWithOptions(request *KuajingSegRequest, runtime *util.RuntimeOptions) (_result *KuajingSegResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["PicUrl"] = request.PicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnPicFormat)) {
		body["ReturnPicFormat"] = request.ReturnPicFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnPicType)) {
		body["ReturnPicType"] = request.ReturnPicType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KuajingSeg"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KuajingSegResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KuajingSeg(request *KuajingSegRequest) (_result *KuajingSegResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KuajingSegResponse{}
	_body, _err := client.KuajingSegWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KuajingSegAdvance(request *KuajingSegAdvanceRequest, runtime *util.RuntimeOptions) (_result *KuajingSegResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	kuajingSegReq := &KuajingSegRequest{}
	openapiutil.Convert(request, kuajingSegReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		kuajingSegReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	kuajingSegResp, _err := client.KuajingSegWithOptions(kuajingSegReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = kuajingSegResp
	return _result, _err
}

func (client *Client) RemoveWordsWithOptions(request *RemoveWordsRequest, runtime *util.RuntimeOptions) (_result *RemoveWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		query["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveWords"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveWordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveWords(request *RemoveWordsRequest) (_result *RemoveWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveWordsResponse{}
	_body, _err := client.RemoveWordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveWordsAdvance(request *RemoveWordsAdvanceRequest, runtime *util.RuntimeOptions) (_result *RemoveWordsResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	removeWordsReq := &RemoveWordsRequest{}
	openapiutil.Convert(request, removeWordsReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		removeWordsReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	removeWordsResp, _err := client.RemoveWordsWithOptions(removeWordsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = removeWordsResp
	return _result, _err
}

func (client *Client) ReplaceBackgroundWithOptions(request *ReplaceBackgroundRequest, runtime *util.RuntimeOptions) (_result *ReplaceBackgroundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackgroundId)) {
		query["BackgroundId"] = request.BackgroundId
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicBackgroundUrl)) {
		query["PicBackgroundUrl"] = request.PicBackgroundUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		query["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceBackground"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceBackgroundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceBackground(request *ReplaceBackgroundRequest) (_result *ReplaceBackgroundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceBackgroundResponse{}
	_body, _err := client.ReplaceBackgroundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceBackgroundAdvance(request *ReplaceBackgroundAdvanceRequest, runtime *util.RuntimeOptions) (_result *ReplaceBackgroundResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	replaceBackgroundReq := &ReplaceBackgroundRequest{}
	openapiutil.Convert(request, replaceBackgroundReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		replaceBackgroundReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	replaceBackgroundResp, _err := client.ReplaceBackgroundWithOptions(replaceBackgroundReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = replaceBackgroundResp
	return _result, _err
}

func (client *Client) SeleteCommodityWithOptions(request *SeleteCommodityRequest, runtime *util.RuntimeOptions) (_result *SeleteCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SeleteCommodity"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SeleteCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SeleteCommodity(request *SeleteCommodityRequest) (_result *SeleteCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SeleteCommodityResponse{}
	_body, _err := client.SeleteCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SeleteCommodityByBToBWithOptions(request *SeleteCommodityByBToBRequest, runtime *util.RuntimeOptions) (_result *SeleteCommodityByBToBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SeleteCommodityByBToB"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SeleteCommodityByBToBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SeleteCommodityByBToB(request *SeleteCommodityByBToBRequest) (_result *SeleteCommodityByBToBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SeleteCommodityByBToBResponse{}
	_body, _err := client.SeleteCommodityByBToBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TbPredictCategoryWithOptions(request *TbPredictCategoryRequest, runtime *util.RuntimeOptions) (_result *TbPredictCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TbPredictCategory"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TbPredictCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TbPredictCategory(request *TbPredictCategoryRequest) (_result *TbPredictCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TbPredictCategoryResponse{}
	_body, _err := client.TbPredictCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TbPredictCategoryAdvance(request *TbPredictCategoryAdvanceRequest, runtime *util.RuntimeOptions) (_result *TbPredictCategoryResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	tbPredictCategoryReq := &TbPredictCategoryRequest{}
	openapiutil.Convert(request, tbPredictCategoryReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		tbPredictCategoryReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	tbPredictCategoryResp, _err := client.TbPredictCategoryWithOptions(tbPredictCategoryReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = tbPredictCategoryResp
	return _result, _err
}

func (client *Client) TbPropRecWithOptions(request *TbPropRecRequest, runtime *util.RuntimeOptions) (_result *TbPropRecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PicUrl)) {
		body["PicUrl"] = request.PicUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TbPropRec"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TbPropRecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TbPropRec(request *TbPropRecRequest) (_result *TbPropRecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TbPropRecResponse{}
	_body, _err := client.TbPropRecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TbPropRecAdvance(request *TbPropRecAdvanceRequest, runtime *util.RuntimeOptions) (_result *TbPropRecResponse, _err error) {
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
		Product:  tea.String("dplus"),
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
	tbPropRecReq := &TbPropRecRequest{}
	openapiutil.Convert(request, tbPropRecReq)
	if !tea.BoolValue(util.IsUnset(request.PicUrlObject)) {
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
			Content:     request.PicUrlObject,
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
		tbPropRecReq.PicUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	tbPropRecResp, _err := client.TbPropRecWithOptions(tbPropRecReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = tbPropRecResp
	return _result, _err
}

func (client *Client) TransferUrlByBtoBWithOptions(request *TransferUrlByBtoBRequest, runtime *util.RuntimeOptions) (_result *TransferUrlByBtoBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OfferId)) {
		query["OfferId"] = request.OfferId
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferUrlByBtoB"),
		Version:     tea.String("2020-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferUrlByBtoBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferUrlByBtoB(request *TransferUrlByBtoBRequest) (_result *TransferUrlByBtoBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferUrlByBtoBResponse{}
	_body, _err := client.TransferUrlByBtoBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
