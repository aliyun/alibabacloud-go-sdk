// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddHotspotFileRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddHotspotFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHotspotFileRequest) GoString() string {
	return s.String()
}

func (s *AddHotspotFileRequest) SetFileName(v string) *AddHotspotFileRequest {
	s.FileName = &v
	return s
}

func (s *AddHotspotFileRequest) SetSceneId(v string) *AddHotspotFileRequest {
	s.SceneId = &v
	return s
}

func (s *AddHotspotFileRequest) SetType(v string) *AddHotspotFileRequest {
	s.Type = &v
	return s
}

type AddHotspotFileResponseBody struct {
	Code      *int64                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddHotspotFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHotspotFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddHotspotFileResponseBody) SetCode(v int64) *AddHotspotFileResponseBody {
	s.Code = &v
	return s
}

func (s *AddHotspotFileResponseBody) SetData(v map[string]interface{}) *AddHotspotFileResponseBody {
	s.Data = v
	return s
}

func (s *AddHotspotFileResponseBody) SetMessage(v string) *AddHotspotFileResponseBody {
	s.Message = &v
	return s
}

func (s *AddHotspotFileResponseBody) SetRequestId(v string) *AddHotspotFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHotspotFileResponseBody) SetSuccess(v bool) *AddHotspotFileResponseBody {
	s.Success = &v
	return s
}

type AddHotspotFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddHotspotFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHotspotFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHotspotFileResponse) GoString() string {
	return s.String()
}

func (s *AddHotspotFileResponse) SetHeaders(v map[string]*string) *AddHotspotFileResponse {
	s.Headers = v
	return s
}

func (s *AddHotspotFileResponse) SetStatusCode(v int32) *AddHotspotFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHotspotFileResponse) SetBody(v *AddHotspotFileResponseBody) *AddHotspotFileResponse {
	s.Body = v
	return s
}

type AddMosaicsRequest struct {
	MarkPosition *string `json:"MarkPosition,omitempty" xml:"MarkPosition,omitempty"`
	SubSceneId   *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s AddMosaicsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMosaicsRequest) GoString() string {
	return s.String()
}

func (s *AddMosaicsRequest) SetMarkPosition(v string) *AddMosaicsRequest {
	s.MarkPosition = &v
	return s
}

func (s *AddMosaicsRequest) SetSubSceneId(v string) *AddMosaicsRequest {
	s.SubSceneId = &v
	return s
}

type AddMosaicsResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddMosaicsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMosaicsResponseBody) GoString() string {
	return s.String()
}

func (s *AddMosaicsResponseBody) SetCode(v int64) *AddMosaicsResponseBody {
	s.Code = &v
	return s
}

func (s *AddMosaicsResponseBody) SetMessage(v string) *AddMosaicsResponseBody {
	s.Message = &v
	return s
}

func (s *AddMosaicsResponseBody) SetRequestId(v string) *AddMosaicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMosaicsResponseBody) SetSuccess(v bool) *AddMosaicsResponseBody {
	s.Success = &v
	return s
}

func (s *AddMosaicsResponseBody) SetTaskId(v string) *AddMosaicsResponseBody {
	s.TaskId = &v
	return s
}

type AddMosaicsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddMosaicsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMosaicsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMosaicsResponse) GoString() string {
	return s.String()
}

func (s *AddMosaicsResponse) SetHeaders(v map[string]*string) *AddMosaicsResponse {
	s.Headers = v
	return s
}

func (s *AddMosaicsResponse) SetStatusCode(v int32) *AddMosaicsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMosaicsResponse) SetBody(v *AddMosaicsResponseBody) *AddMosaicsResponse {
	s.Body = v
	return s
}

type AddProjectRequest struct {
	BusinessId *int64  `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AddProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProjectRequest) GoString() string {
	return s.String()
}

func (s *AddProjectRequest) SetBusinessId(v int64) *AddProjectRequest {
	s.BusinessId = &v
	return s
}

func (s *AddProjectRequest) SetName(v string) *AddProjectRequest {
	s.Name = &v
	return s
}

type AddProjectResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProjectResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectResponseBody) SetCode(v int64) *AddProjectResponseBody {
	s.Code = &v
	return s
}

func (s *AddProjectResponseBody) SetId(v string) *AddProjectResponseBody {
	s.Id = &v
	return s
}

func (s *AddProjectResponseBody) SetMessage(v string) *AddProjectResponseBody {
	s.Message = &v
	return s
}

func (s *AddProjectResponseBody) SetRequestId(v string) *AddProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectResponseBody) SetSuccess(v bool) *AddProjectResponseBody {
	s.Success = &v
	return s
}

type AddProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProjectResponse) GoString() string {
	return s.String()
}

func (s *AddProjectResponse) SetHeaders(v map[string]*string) *AddProjectResponse {
	s.Headers = v
	return s
}

func (s *AddProjectResponse) SetStatusCode(v int32) *AddProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectResponse) SetBody(v *AddProjectResponseBody) *AddProjectResponse {
	s.Body = v
	return s
}

type AddRelativePositionRequest struct {
	RelativePosition *string `json:"RelativePosition,omitempty" xml:"RelativePosition,omitempty"`
	SceneId          *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AddRelativePositionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRelativePositionRequest) GoString() string {
	return s.String()
}

func (s *AddRelativePositionRequest) SetRelativePosition(v string) *AddRelativePositionRequest {
	s.RelativePosition = &v
	return s
}

func (s *AddRelativePositionRequest) SetSceneId(v string) *AddRelativePositionRequest {
	s.SceneId = &v
	return s
}

type AddRelativePositionResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRelativePositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRelativePositionResponseBody) GoString() string {
	return s.String()
}

func (s *AddRelativePositionResponseBody) SetCode(v int64) *AddRelativePositionResponseBody {
	s.Code = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetMessage(v string) *AddRelativePositionResponseBody {
	s.Message = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetRequestId(v string) *AddRelativePositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetSuccess(v bool) *AddRelativePositionResponseBody {
	s.Success = &v
	return s
}

type AddRelativePositionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRelativePositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRelativePositionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRelativePositionResponse) GoString() string {
	return s.String()
}

func (s *AddRelativePositionResponse) SetHeaders(v map[string]*string) *AddRelativePositionResponse {
	s.Headers = v
	return s
}

func (s *AddRelativePositionResponse) SetStatusCode(v int32) *AddRelativePositionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRelativePositionResponse) SetBody(v *AddRelativePositionResponseBody) *AddRelativePositionResponse {
	s.Body = v
	return s
}

type AddRoomPlanRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AddRoomPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRoomPlanRequest) GoString() string {
	return s.String()
}

func (s *AddRoomPlanRequest) SetSceneId(v string) *AddRoomPlanRequest {
	s.SceneId = &v
	return s
}

type AddRoomPlanResponseBody struct {
	Code      *int64                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AddRoomPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRoomPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRoomPlanResponseBody) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponseBody) SetCode(v int64) *AddRoomPlanResponseBody {
	s.Code = &v
	return s
}

func (s *AddRoomPlanResponseBody) SetData(v *AddRoomPlanResponseBodyData) *AddRoomPlanResponseBody {
	s.Data = v
	return s
}

func (s *AddRoomPlanResponseBody) SetMessage(v string) *AddRoomPlanResponseBody {
	s.Message = &v
	return s
}

func (s *AddRoomPlanResponseBody) SetRequestId(v string) *AddRoomPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRoomPlanResponseBody) SetSuccess(v bool) *AddRoomPlanResponseBody {
	s.Success = &v
	return s
}

type AddRoomPlanResponseBodyData struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Callback  *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s AddRoomPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddRoomPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponseBodyData) SetAccessId(v string) *AddRoomPlanResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetCallback(v string) *AddRoomPlanResponseBodyData {
	s.Callback = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetDir(v string) *AddRoomPlanResponseBodyData {
	s.Dir = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetExpire(v string) *AddRoomPlanResponseBodyData {
	s.Expire = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetHost(v string) *AddRoomPlanResponseBodyData {
	s.Host = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetPolicy(v string) *AddRoomPlanResponseBodyData {
	s.Policy = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetSignature(v string) *AddRoomPlanResponseBodyData {
	s.Signature = &v
	return s
}

type AddRoomPlanResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRoomPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRoomPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRoomPlanResponse) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponse) SetHeaders(v map[string]*string) *AddRoomPlanResponse {
	s.Headers = v
	return s
}

func (s *AddRoomPlanResponse) SetStatusCode(v int32) *AddRoomPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRoomPlanResponse) SetBody(v *AddRoomPlanResponseBody) *AddRoomPlanResponse {
	s.Body = v
	return s
}

type AddSceneRequest struct {
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSceneRequest) GoString() string {
	return s.String()
}

func (s *AddSceneRequest) SetCustomerUid(v string) *AddSceneRequest {
	s.CustomerUid = &v
	return s
}

func (s *AddSceneRequest) SetName(v string) *AddSceneRequest {
	s.Name = &v
	return s
}

func (s *AddSceneRequest) SetProjectId(v string) *AddSceneRequest {
	s.ProjectId = &v
	return s
}

func (s *AddSceneRequest) SetType(v string) *AddSceneRequest {
	s.Type = &v
	return s
}

type AddSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSceneResponseBody) GoString() string {
	return s.String()
}

func (s *AddSceneResponseBody) SetCode(v int64) *AddSceneResponseBody {
	s.Code = &v
	return s
}

func (s *AddSceneResponseBody) SetId(v string) *AddSceneResponseBody {
	s.Id = &v
	return s
}

func (s *AddSceneResponseBody) SetMessage(v string) *AddSceneResponseBody {
	s.Message = &v
	return s
}

func (s *AddSceneResponseBody) SetRequestId(v string) *AddSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSceneResponseBody) SetSuccess(v bool) *AddSceneResponseBody {
	s.Success = &v
	return s
}

type AddSceneResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSceneResponse) GoString() string {
	return s.String()
}

func (s *AddSceneResponse) SetHeaders(v map[string]*string) *AddSceneResponse {
	s.Headers = v
	return s
}

func (s *AddSceneResponse) SetStatusCode(v int32) *AddSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSceneResponse) SetBody(v *AddSceneResponseBody) *AddSceneResponse {
	s.Body = v
	return s
}

type AddSubSceneRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s AddSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSubSceneRequest) GoString() string {
	return s.String()
}

func (s *AddSubSceneRequest) SetName(v string) *AddSubSceneRequest {
	s.Name = &v
	return s
}

func (s *AddSubSceneRequest) SetSceneId(v string) *AddSubSceneRequest {
	s.SceneId = &v
	return s
}

func (s *AddSubSceneRequest) SetUploadType(v string) *AddSubSceneRequest {
	s.UploadType = &v
	return s
}

type AddSubSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *AddSubSceneResponseBody) SetCode(v int64) *AddSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *AddSubSceneResponseBody) SetId(v string) *AddSubSceneResponseBody {
	s.Id = &v
	return s
}

func (s *AddSubSceneResponseBody) SetMessage(v string) *AddSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *AddSubSceneResponseBody) SetRequestId(v string) *AddSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSubSceneResponseBody) SetSuccess(v bool) *AddSubSceneResponseBody {
	s.Success = &v
	return s
}

type AddSubSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSubSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSubSceneResponse) GoString() string {
	return s.String()
}

func (s *AddSubSceneResponse) SetHeaders(v map[string]*string) *AddSubSceneResponse {
	s.Headers = v
	return s
}

func (s *AddSubSceneResponse) SetStatusCode(v int32) *AddSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSubSceneResponse) SetBody(v *AddSubSceneResponseBody) *AddSubSceneResponse {
	s.Body = v
	return s
}

type CheckUserPropertyRequest struct {
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CheckUserPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserPropertyRequest) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyRequest) SetUid(v string) *CheckUserPropertyRequest {
	s.Uid = &v
	return s
}

type CheckUserPropertyResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Match     *bool   `json:"Match,omitempty" xml:"Match,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUserPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyResponseBody) SetCode(v int64) *CheckUserPropertyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetMatch(v bool) *CheckUserPropertyResponseBody {
	s.Match = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetMessage(v string) *CheckUserPropertyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetRequestId(v string) *CheckUserPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserPropertyResponseBody) SetSuccess(v bool) *CheckUserPropertyResponseBody {
	s.Success = &v
	return s
}

type CheckUserPropertyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckUserPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUserPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserPropertyResponse) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyResponse) SetHeaders(v map[string]*string) *CheckUserPropertyResponse {
	s.Headers = v
	return s
}

func (s *CheckUserPropertyResponse) SetStatusCode(v int32) *CheckUserPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserPropertyResponse) SetBody(v *CheckUserPropertyResponseBody) *CheckUserPropertyResponse {
	s.Body = v
	return s
}

type CopySceneRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s CopySceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CopySceneRequest) GoString() string {
	return s.String()
}

func (s *CopySceneRequest) SetProjectId(v string) *CopySceneRequest {
	s.ProjectId = &v
	return s
}

func (s *CopySceneRequest) SetSceneId(v string) *CopySceneRequest {
	s.SceneId = &v
	return s
}

func (s *CopySceneRequest) SetSceneName(v string) *CopySceneRequest {
	s.SceneName = &v
	return s
}

type CopySceneResponseBody struct {
	Code      *int64                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CopySceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CopySceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopySceneResponseBody) GoString() string {
	return s.String()
}

func (s *CopySceneResponseBody) SetCode(v int64) *CopySceneResponseBody {
	s.Code = &v
	return s
}

func (s *CopySceneResponseBody) SetData(v *CopySceneResponseBodyData) *CopySceneResponseBody {
	s.Data = v
	return s
}

func (s *CopySceneResponseBody) SetMessage(v string) *CopySceneResponseBody {
	s.Message = &v
	return s
}

func (s *CopySceneResponseBody) SetRequestId(v string) *CopySceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopySceneResponseBody) SetSuccess(v bool) *CopySceneResponseBody {
	s.Success = &v
	return s
}

type CopySceneResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CopySceneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CopySceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopySceneResponseBodyData) SetTaskId(v string) *CopySceneResponseBodyData {
	s.TaskId = &v
	return s
}

type CopySceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CopySceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopySceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CopySceneResponse) GoString() string {
	return s.String()
}

func (s *CopySceneResponse) SetHeaders(v map[string]*string) *CopySceneResponse {
	s.Headers = v
	return s
}

func (s *CopySceneResponse) SetStatusCode(v int32) *CopySceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CopySceneResponse) SetBody(v *CopySceneResponseBody) *CopySceneResponse {
	s.Body = v
	return s
}

type CreateUploadPolicyRequest struct {
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateUploadPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyRequest) SetOption(v string) *CreateUploadPolicyRequest {
	s.Option = &v
	return s
}

func (s *CreateUploadPolicyRequest) SetType(v string) *CreateUploadPolicyRequest {
	s.Type = &v
	return s
}

type CreateUploadPolicyResponseBody struct {
	Code      *int64                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateUploadPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponseBody) SetCode(v int64) *CreateUploadPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetData(v *CreateUploadPolicyResponseBodyData) *CreateUploadPolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetMessage(v string) *CreateUploadPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetRequestId(v string) *CreateUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadPolicyResponseBody) SetSuccess(v bool) *CreateUploadPolicyResponseBody {
	s.Success = &v
	return s
}

type CreateUploadPolicyResponseBodyData struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Callback  *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateUploadPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponseBodyData) SetAccessId(v string) *CreateUploadPolicyResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetCallback(v string) *CreateUploadPolicyResponseBodyData {
	s.Callback = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetDir(v string) *CreateUploadPolicyResponseBodyData {
	s.Dir = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetExpire(v string) *CreateUploadPolicyResponseBodyData {
	s.Expire = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetHost(v string) *CreateUploadPolicyResponseBodyData {
	s.Host = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetPolicy(v string) *CreateUploadPolicyResponseBodyData {
	s.Policy = &v
	return s
}

func (s *CreateUploadPolicyResponseBodyData) SetSignature(v string) *CreateUploadPolicyResponseBodyData {
	s.Signature = &v
	return s
}

type CreateUploadPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadPolicyResponse) SetHeaders(v map[string]*string) *CreateUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadPolicyResponse) SetStatusCode(v int32) *CreateUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadPolicyResponse) SetBody(v *CreateUploadPolicyResponseBody) *CreateUploadPolicyResponse {
	s.Body = v
	return s
}

type DetailProjectRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetailProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailProjectRequest) GoString() string {
	return s.String()
}

func (s *DetailProjectRequest) SetId(v string) *DetailProjectRequest {
	s.Id = &v
	return s
}

type DetailProjectResponseBody struct {
	BusinessId   *int64  `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	Code         *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DetailProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetailProjectResponseBody) SetBusinessId(v int64) *DetailProjectResponseBody {
	s.BusinessId = &v
	return s
}

func (s *DetailProjectResponseBody) SetBusinessName(v string) *DetailProjectResponseBody {
	s.BusinessName = &v
	return s
}

func (s *DetailProjectResponseBody) SetCode(v int64) *DetailProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DetailProjectResponseBody) SetGmtCreate(v int64) *DetailProjectResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DetailProjectResponseBody) SetGmtModified(v int64) *DetailProjectResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DetailProjectResponseBody) SetId(v string) *DetailProjectResponseBody {
	s.Id = &v
	return s
}

func (s *DetailProjectResponseBody) SetMessage(v string) *DetailProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DetailProjectResponseBody) SetName(v string) *DetailProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DetailProjectResponseBody) SetRequestId(v string) *DetailProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailProjectResponseBody) SetSuccess(v bool) *DetailProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DetailProjectResponseBody) SetToken(v string) *DetailProjectResponseBody {
	s.Token = &v
	return s
}

type DetailProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetailProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetailProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailProjectResponse) GoString() string {
	return s.String()
}

func (s *DetailProjectResponse) SetHeaders(v map[string]*string) *DetailProjectResponse {
	s.Headers = v
	return s
}

func (s *DetailProjectResponse) SetStatusCode(v int32) *DetailProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailProjectResponse) SetBody(v *DetailProjectResponseBody) *DetailProjectResponse {
	s.Body = v
	return s
}

type DetailSceneRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetailSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailSceneRequest) GoString() string {
	return s.String()
}

func (s *DetailSceneRequest) SetId(v string) *DetailSceneRequest {
	s.Id = &v
	return s
}

type DetailSceneResponseBody struct {
	Captures     []*DetailSceneResponseBodyCaptures   `json:"Captures,omitempty" xml:"Captures,omitempty" type:"Repeated"`
	Code         *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	CoverUrl     *string                              `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	FloorPlans   []*DetailSceneResponseBodyFloorPlans `json:"FloorPlans,omitempty" xml:"FloorPlans,omitempty" type:"Repeated"`
	GmtCreate    *int64                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *string                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Message      *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Name         *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	PreviewToken *string                              `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	Published    *bool                                `json:"Published,omitempty" xml:"Published,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceNum    *int64                               `json:"SourceNum,omitempty" xml:"SourceNum,omitempty"`
	Status       *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName   *string                              `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	SubSceneNum  *int64                               `json:"SubSceneNum,omitempty" xml:"SubSceneNum,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Type         *string                              `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetailSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBody) SetCaptures(v []*DetailSceneResponseBodyCaptures) *DetailSceneResponseBody {
	s.Captures = v
	return s
}

func (s *DetailSceneResponseBody) SetCode(v int64) *DetailSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DetailSceneResponseBody) SetCoverUrl(v string) *DetailSceneResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *DetailSceneResponseBody) SetFloorPlans(v []*DetailSceneResponseBodyFloorPlans) *DetailSceneResponseBody {
	s.FloorPlans = v
	return s
}

func (s *DetailSceneResponseBody) SetGmtCreate(v int64) *DetailSceneResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DetailSceneResponseBody) SetGmtModified(v int64) *DetailSceneResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DetailSceneResponseBody) SetId(v string) *DetailSceneResponseBody {
	s.Id = &v
	return s
}

func (s *DetailSceneResponseBody) SetMessage(v string) *DetailSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DetailSceneResponseBody) SetName(v string) *DetailSceneResponseBody {
	s.Name = &v
	return s
}

func (s *DetailSceneResponseBody) SetPreviewToken(v string) *DetailSceneResponseBody {
	s.PreviewToken = &v
	return s
}

func (s *DetailSceneResponseBody) SetPublished(v bool) *DetailSceneResponseBody {
	s.Published = &v
	return s
}

func (s *DetailSceneResponseBody) SetRequestId(v string) *DetailSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailSceneResponseBody) SetSourceNum(v int64) *DetailSceneResponseBody {
	s.SourceNum = &v
	return s
}

func (s *DetailSceneResponseBody) SetStatus(v string) *DetailSceneResponseBody {
	s.Status = &v
	return s
}

func (s *DetailSceneResponseBody) SetStatusName(v string) *DetailSceneResponseBody {
	s.StatusName = &v
	return s
}

func (s *DetailSceneResponseBody) SetSubSceneNum(v int64) *DetailSceneResponseBody {
	s.SubSceneNum = &v
	return s
}

func (s *DetailSceneResponseBody) SetSuccess(v bool) *DetailSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DetailSceneResponseBody) SetType(v string) *DetailSceneResponseBody {
	s.Type = &v
	return s
}

type DetailSceneResponseBodyCaptures struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetailSceneResponseBodyCaptures) String() string {
	return tea.Prettify(s)
}

func (s DetailSceneResponseBodyCaptures) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBodyCaptures) SetTitle(v string) *DetailSceneResponseBodyCaptures {
	s.Title = &v
	return s
}

func (s *DetailSceneResponseBodyCaptures) SetUrl(v string) *DetailSceneResponseBodyCaptures {
	s.Url = &v
	return s
}

type DetailSceneResponseBodyFloorPlans struct {
	ColorMapUrl *string `json:"ColorMapUrl,omitempty" xml:"ColorMapUrl,omitempty"`
	FloorLabel  *string `json:"FloorLabel,omitempty" xml:"FloorLabel,omitempty"`
	FloorName   *string `json:"FloorName,omitempty" xml:"FloorName,omitempty"`
	MiniMapUrl  *string `json:"MiniMapUrl,omitempty" xml:"MiniMapUrl,omitempty"`
}

func (s DetailSceneResponseBodyFloorPlans) String() string {
	return tea.Prettify(s)
}

func (s DetailSceneResponseBodyFloorPlans) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBodyFloorPlans) SetColorMapUrl(v string) *DetailSceneResponseBodyFloorPlans {
	s.ColorMapUrl = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) SetFloorLabel(v string) *DetailSceneResponseBodyFloorPlans {
	s.FloorLabel = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) SetFloorName(v string) *DetailSceneResponseBodyFloorPlans {
	s.FloorName = &v
	return s
}

func (s *DetailSceneResponseBodyFloorPlans) SetMiniMapUrl(v string) *DetailSceneResponseBodyFloorPlans {
	s.MiniMapUrl = &v
	return s
}

type DetailSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetailSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetailSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailSceneResponse) GoString() string {
	return s.String()
}

func (s *DetailSceneResponse) SetHeaders(v map[string]*string) *DetailSceneResponse {
	s.Headers = v
	return s
}

func (s *DetailSceneResponse) SetStatusCode(v int32) *DetailSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailSceneResponse) SetBody(v *DetailSceneResponseBody) *DetailSceneResponse {
	s.Body = v
	return s
}

type DetailSubSceneRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DetailSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailSubSceneRequest) GoString() string {
	return s.String()
}

func (s *DetailSubSceneRequest) SetId(v string) *DetailSubSceneRequest {
	s.Id = &v
	return s
}

type DetailSubSceneResponseBody struct {
	Code        *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	CoverUrl    *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CubemapPath *string `json:"CubemapPath,omitempty" xml:"CubemapPath,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageUrl    *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	LayoutData  *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginUrl   *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	Position    *string `json:"Position,omitempty" xml:"Position,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status      *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetailSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DetailSubSceneResponseBody) SetCode(v int64) *DetailSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetCoverUrl(v string) *DetailSubSceneResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetCubemapPath(v string) *DetailSubSceneResponseBody {
	s.CubemapPath = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetGmtCreate(v int64) *DetailSubSceneResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetGmtModified(v int64) *DetailSubSceneResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetId(v string) *DetailSubSceneResponseBody {
	s.Id = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetImageUrl(v string) *DetailSubSceneResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetLayoutData(v string) *DetailSubSceneResponseBody {
	s.LayoutData = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetMessage(v string) *DetailSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetName(v string) *DetailSubSceneResponseBody {
	s.Name = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetOriginUrl(v string) *DetailSubSceneResponseBody {
	s.OriginUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetPosition(v string) *DetailSubSceneResponseBody {
	s.Position = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetRequestId(v string) *DetailSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetResourceId(v string) *DetailSubSceneResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetStatus(v int64) *DetailSubSceneResponseBody {
	s.Status = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetSuccess(v bool) *DetailSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetType(v string) *DetailSubSceneResponseBody {
	s.Type = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetUrl(v string) *DetailSubSceneResponseBody {
	s.Url = &v
	return s
}

type DetailSubSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetailSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetailSubSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailSubSceneResponse) GoString() string {
	return s.String()
}

func (s *DetailSubSceneResponse) SetHeaders(v map[string]*string) *DetailSubSceneResponse {
	s.Headers = v
	return s
}

func (s *DetailSubSceneResponse) SetStatusCode(v int32) *DetailSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailSubSceneResponse) SetBody(v *DetailSubSceneResponseBody) *DetailSubSceneResponse {
	s.Body = v
	return s
}

type DropProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DropProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DropProjectRequest) GoString() string {
	return s.String()
}

func (s *DropProjectRequest) SetProjectId(v string) *DropProjectRequest {
	s.ProjectId = &v
	return s
}

type DropProjectResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DropProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DropProjectResponseBody) SetCode(v int64) *DropProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DropProjectResponseBody) SetMessage(v string) *DropProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DropProjectResponseBody) SetRequestId(v string) *DropProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropProjectResponseBody) SetSuccess(v bool) *DropProjectResponseBody {
	s.Success = &v
	return s
}

type DropProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DropProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DropProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DropProjectResponse) GoString() string {
	return s.String()
}

func (s *DropProjectResponse) SetHeaders(v map[string]*string) *DropProjectResponse {
	s.Headers = v
	return s
}

func (s *DropProjectResponse) SetStatusCode(v int32) *DropProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DropProjectResponse) SetBody(v *DropProjectResponseBody) *DropProjectResponse {
	s.Body = v
	return s
}

type DropSceneRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DropSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DropSceneRequest) GoString() string {
	return s.String()
}

func (s *DropSceneRequest) SetId(v string) *DropSceneRequest {
	s.Id = &v
	return s
}

type DropSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DropSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DropSceneResponseBody) SetCode(v int64) *DropSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DropSceneResponseBody) SetMessage(v string) *DropSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DropSceneResponseBody) SetRequestId(v string) *DropSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSceneResponseBody) SetSuccess(v bool) *DropSceneResponseBody {
	s.Success = &v
	return s
}

type DropSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DropSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DropSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DropSceneResponse) GoString() string {
	return s.String()
}

func (s *DropSceneResponse) SetHeaders(v map[string]*string) *DropSceneResponse {
	s.Headers = v
	return s
}

func (s *DropSceneResponse) SetStatusCode(v int32) *DropSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DropSceneResponse) SetBody(v *DropSceneResponseBody) *DropSceneResponse {
	s.Body = v
	return s
}

type DropSubSceneRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DropSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DropSubSceneRequest) GoString() string {
	return s.String()
}

func (s *DropSubSceneRequest) SetId(v string) *DropSubSceneRequest {
	s.Id = &v
	return s
}

type DropSubSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DropSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DropSubSceneResponseBody) SetCode(v int64) *DropSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DropSubSceneResponseBody) SetMessage(v string) *DropSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DropSubSceneResponseBody) SetRequestId(v string) *DropSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSubSceneResponseBody) SetSuccess(v bool) *DropSubSceneResponseBody {
	s.Success = &v
	return s
}

type DropSubSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DropSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DropSubSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DropSubSceneResponse) GoString() string {
	return s.String()
}

func (s *DropSubSceneResponse) SetHeaders(v map[string]*string) *DropSubSceneResponse {
	s.Headers = v
	return s
}

func (s *DropSubSceneResponse) SetStatusCode(v int32) *DropSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DropSubSceneResponse) SetBody(v *DropSubSceneResponseBody) *DropSubSceneResponse {
	s.Body = v
	return s
}

type GetConnDataRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetConnDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnDataRequest) GoString() string {
	return s.String()
}

func (s *GetConnDataRequest) SetSceneId(v string) *GetConnDataRequest {
	s.SceneId = &v
	return s
}

type GetConnDataResponseBody struct {
	Code      *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Extend    *string                        `json:"Extend,omitempty" xml:"Extend,omitempty"`
	List      []*GetConnDataResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Version   *string                        `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetConnDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnDataResponseBody) SetCode(v int64) *GetConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnDataResponseBody) SetExtend(v string) *GetConnDataResponseBody {
	s.Extend = &v
	return s
}

func (s *GetConnDataResponseBody) SetList(v []*GetConnDataResponseBodyList) *GetConnDataResponseBody {
	s.List = v
	return s
}

func (s *GetConnDataResponseBody) SetMessage(v string) *GetConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnDataResponseBody) SetRequestId(v string) *GetConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnDataResponseBody) SetSuccess(v bool) *GetConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetConnDataResponseBody) SetVersion(v string) *GetConnDataResponseBody {
	s.Version = &v
	return s
}

type GetConnDataResponseBodyList struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MapId *string `json:"MapId,omitempty" xml:"MapId,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetConnDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetConnDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetConnDataResponseBodyList) SetId(v string) *GetConnDataResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetConnDataResponseBodyList) SetMapId(v string) *GetConnDataResponseBodyList {
	s.MapId = &v
	return s
}

func (s *GetConnDataResponseBodyList) SetType(v string) *GetConnDataResponseBodyList {
	s.Type = &v
	return s
}

type GetConnDataResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConnDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnDataResponse) GoString() string {
	return s.String()
}

func (s *GetConnDataResponse) SetHeaders(v map[string]*string) *GetConnDataResponse {
	s.Headers = v
	return s
}

func (s *GetConnDataResponse) SetStatusCode(v int32) *GetConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnDataResponse) SetBody(v *GetConnDataResponseBody) *GetConnDataResponse {
	s.Body = v
	return s
}

type GetCopySceneTaskStatusRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCopySceneTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCopySceneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusRequest) SetTaskId(v string) *GetCopySceneTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetCopySceneTaskStatusResponseBody struct {
	Code      *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCopySceneTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCopySceneTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCopySceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponseBody) SetCode(v int64) *GetCopySceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetData(v *GetCopySceneTaskStatusResponseBodyData) *GetCopySceneTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetMessage(v string) *GetCopySceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetRequestId(v string) *GetCopySceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetSuccess(v bool) *GetCopySceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

type GetCopySceneTaskStatusResponseBodyData struct {
	Progress *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCopySceneTaskStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCopySceneTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponseBodyData) SetProgress(v int64) *GetCopySceneTaskStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyData) SetStatus(v string) *GetCopySceneTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetCopySceneTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCopySceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCopySceneTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCopySceneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponse) SetHeaders(v map[string]*string) *GetCopySceneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCopySceneTaskStatusResponse) SetStatusCode(v int32) *GetCopySceneTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopySceneTaskStatusResponse) SetBody(v *GetCopySceneTaskStatusResponseBody) *GetCopySceneTaskStatusResponse {
	s.Body = v
	return s
}

type GetHotspotConfigRequest struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	Type         *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHotspotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigRequest) SetDomain(v string) *GetHotspotConfigRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotConfigRequest) SetEnabled(v bool) *GetHotspotConfigRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotConfigRequest) SetPreviewToken(v string) *GetHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotConfigRequest) SetType(v int64) *GetHotspotConfigRequest {
	s.Type = &v
	return s
}

type GetHotspotConfigResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponseBody) SetCode(v int64) *GetHotspotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetData(v string) *GetHotspotConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetMessage(v string) *GetHotspotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetRequestId(v string) *GetHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetSuccess(v bool) *GetHotspotConfigResponseBody {
	s.Success = &v
	return s
}

type GetHotspotConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotspotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponse) SetHeaders(v map[string]*string) *GetHotspotConfigResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotConfigResponse) SetStatusCode(v int32) *GetHotspotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotConfigResponse) SetBody(v *GetHotspotConfigResponseBody) *GetHotspotConfigResponse {
	s.Body = v
	return s
}

type GetHotspotSceneDataRequest struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	Type         *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHotspotSceneDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataRequest) SetDomain(v string) *GetHotspotSceneDataRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetEnabled(v bool) *GetHotspotSceneDataRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetPreviewToken(v string) *GetHotspotSceneDataRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetType(v int64) *GetHotspotSceneDataRequest {
	s.Type = &v
	return s
}

type GetHotspotSceneDataResponseBody struct {
	Code      *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHotspotSceneDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotSceneDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBody) SetCode(v int64) *GetHotspotSceneDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetData(v *GetHotspotSceneDataResponseBodyData) *GetHotspotSceneDataResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetMessage(v string) *GetHotspotSceneDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetRequestId(v string) *GetHotspotSceneDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetSuccess(v bool) *GetHotspotSceneDataResponseBody {
	s.Success = &v
	return s
}

type GetHotspotSceneDataResponseBodyData struct {
	ModelToken   *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
	PreviewData  *string `json:"PreviewData,omitempty" xml:"PreviewData,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	SceneType    *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetHotspotSceneDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBodyData) SetModelToken(v string) *GetHotspotSceneDataResponseBodyData {
	s.ModelToken = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetPreviewData(v string) *GetHotspotSceneDataResponseBodyData {
	s.PreviewData = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetPreviewToken(v string) *GetHotspotSceneDataResponseBodyData {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetSceneType(v string) *GetHotspotSceneDataResponseBodyData {
	s.SceneType = &v
	return s
}

type GetHotspotSceneDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotspotSceneDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotspotSceneDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponse) SetHeaders(v map[string]*string) *GetHotspotSceneDataResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotSceneDataResponse) SetStatusCode(v int32) *GetHotspotSceneDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotSceneDataResponse) SetBody(v *GetHotspotSceneDataResponseBody) *GetHotspotSceneDataResponse {
	s.Body = v
	return s
}

type GetHotspotTagRequest struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHotspotTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotTagRequest) SetDomain(v string) *GetHotspotTagRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotTagRequest) SetEnabled(v bool) *GetHotspotTagRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotTagRequest) SetPreviewToken(v string) *GetHotspotTagRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotTagRequest) SetSubSceneUuid(v string) *GetHotspotTagRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *GetHotspotTagRequest) SetType(v string) *GetHotspotTagRequest {
	s.Type = &v
	return s
}

type GetHotspotTagResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponseBody) SetData(v string) *GetHotspotTagResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetErrMessage(v string) *GetHotspotTagResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetObjectString(v string) *GetHotspotTagResponseBody {
	s.ObjectString = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetRequestId(v string) *GetHotspotTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetSuccess(v bool) *GetHotspotTagResponseBody {
	s.Success = &v
	return s
}

type GetHotspotTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotspotTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotspotTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponse) SetHeaders(v map[string]*string) *GetHotspotTagResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotTagResponse) SetStatusCode(v int32) *GetHotspotTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotTagResponse) SetBody(v *GetHotspotTagResponseBody) *GetHotspotTagResponse {
	s.Body = v
	return s
}

type GetLayoutDataRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetLayoutDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *GetLayoutDataRequest) SetSubSceneId(v string) *GetLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

type GetLayoutDataResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLayoutDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetLayoutDataResponseBody) SetCode(v int64) *GetLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetData(v string) *GetLayoutDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetMessage(v string) *GetLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetRequestId(v string) *GetLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetSuccess(v bool) *GetLayoutDataResponseBody {
	s.Success = &v
	return s
}

type GetLayoutDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLayoutDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLayoutDataResponse) GoString() string {
	return s.String()
}

func (s *GetLayoutDataResponse) SetHeaders(v map[string]*string) *GetLayoutDataResponse {
	s.Headers = v
	return s
}

func (s *GetLayoutDataResponse) SetStatusCode(v int32) *GetLayoutDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayoutDataResponse) SetBody(v *GetLayoutDataResponseBody) *GetLayoutDataResponse {
	s.Body = v
	return s
}

type GetOriginLayoutDataRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetOriginLayoutDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOriginLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataRequest) SetSubSceneId(v string) *GetOriginLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

type GetOriginLayoutDataResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOriginLayoutDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOriginLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataResponseBody) SetCode(v int64) *GetOriginLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetData(v string) *GetOriginLayoutDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetMessage(v string) *GetOriginLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetRequestId(v string) *GetOriginLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetSuccess(v bool) *GetOriginLayoutDataResponseBody {
	s.Success = &v
	return s
}

type GetOriginLayoutDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOriginLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOriginLayoutDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOriginLayoutDataResponse) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataResponse) SetHeaders(v map[string]*string) *GetOriginLayoutDataResponse {
	s.Headers = v
	return s
}

func (s *GetOriginLayoutDataResponse) SetStatusCode(v int32) *GetOriginLayoutDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginLayoutDataResponse) SetBody(v *GetOriginLayoutDataResponseBody) *GetOriginLayoutDataResponse {
	s.Body = v
	return s
}

type GetOssPolicyRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetOssPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetOssPolicyRequest) SetSubSceneId(v string) *GetOssPolicyRequest {
	s.SubSceneId = &v
	return s
}

type GetOssPolicyResponseBody struct {
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Callback  *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOssPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponseBody) SetAccessId(v string) *GetOssPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetCallback(v string) *GetOssPolicyResponseBody {
	s.Callback = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetCode(v int64) *GetOssPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetDir(v string) *GetOssPolicyResponseBody {
	s.Dir = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetExpire(v string) *GetOssPolicyResponseBody {
	s.Expire = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetHost(v string) *GetOssPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetMessage(v string) *GetOssPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetPolicy(v string) *GetOssPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetRequestId(v string) *GetOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetSignature(v string) *GetOssPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetSuccess(v bool) *GetOssPolicyResponseBody {
	s.Success = &v
	return s
}

type GetOssPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponse) SetHeaders(v map[string]*string) *GetOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetOssPolicyResponse) SetStatusCode(v int32) *GetOssPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssPolicyResponse) SetBody(v *GetOssPolicyResponseBody) *GetOssPolicyResponse {
	s.Body = v
	return s
}

type GetPackSceneTaskStatusRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPackSceneTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPackSceneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusRequest) SetTaskId(v string) *GetPackSceneTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetPackSceneTaskStatusRequest) SetType(v string) *GetPackSceneTaskStatusRequest {
	s.Type = &v
	return s
}

type GetPackSceneTaskStatusResponseBody struct {
	Code      *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetPackSceneTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPackSceneTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPackSceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponseBody) SetCode(v int64) *GetPackSceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetData(v *GetPackSceneTaskStatusResponseBodyData) *GetPackSceneTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetMessage(v string) *GetPackSceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetRequestId(v string) *GetPackSceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetSuccess(v bool) *GetPackSceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

type GetPackSceneTaskStatusResponseBodyData struct {
	Progress *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPackSceneTaskStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPackSceneTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponseBodyData) SetProgress(v int64) *GetPackSceneTaskStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyData) SetStatus(v string) *GetPackSceneTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetPackSceneTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPackSceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPackSceneTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPackSceneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponse) SetHeaders(v map[string]*string) *GetPackSceneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPackSceneTaskStatusResponse) SetStatusCode(v int32) *GetPackSceneTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackSceneTaskStatusResponse) SetBody(v *GetPackSceneTaskStatusResponseBody) *GetPackSceneTaskStatusResponse {
	s.Body = v
	return s
}

type GetRectifyImageRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetRectifyImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRectifyImageRequest) GoString() string {
	return s.String()
}

func (s *GetRectifyImageRequest) SetSubSceneId(v string) *GetRectifyImageRequest {
	s.SubSceneId = &v
	return s
}

type GetRectifyImageResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRectifyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRectifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetRectifyImageResponseBody) SetCode(v int64) *GetRectifyImageResponseBody {
	s.Code = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetMessage(v string) *GetRectifyImageResponseBody {
	s.Message = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetRequestId(v string) *GetRectifyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetSuccess(v bool) *GetRectifyImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetUrl(v string) *GetRectifyImageResponseBody {
	s.Url = &v
	return s
}

type GetRectifyImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRectifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRectifyImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRectifyImageResponse) GoString() string {
	return s.String()
}

func (s *GetRectifyImageResponse) SetHeaders(v map[string]*string) *GetRectifyImageResponse {
	s.Headers = v
	return s
}

func (s *GetRectifyImageResponse) SetStatusCode(v int32) *GetRectifyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRectifyImageResponse) SetBody(v *GetRectifyImageResponseBody) *GetRectifyImageResponse {
	s.Body = v
	return s
}

type GetSceneBuildTaskStatusRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetSceneBuildTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneBuildTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusRequest) SetSceneId(v string) *GetSceneBuildTaskStatusRequest {
	s.SceneId = &v
	return s
}

type GetSceneBuildTaskStatusResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSceneBuildTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneBuildTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusResponseBody) SetCode(v int64) *GetSceneBuildTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetErrorCode(v string) *GetSceneBuildTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetErrorMsg(v string) *GetSceneBuildTaskStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetMessage(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetRequestId(v string) *GetSceneBuildTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetSceneId(v string) *GetSceneBuildTaskStatusResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetStatus(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetSuccess(v bool) *GetSceneBuildTaskStatusResponseBody {
	s.Success = &v
	return s
}

type GetSceneBuildTaskStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSceneBuildTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSceneBuildTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneBuildTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusResponse) SetHeaders(v map[string]*string) *GetSceneBuildTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSceneBuildTaskStatusResponse) SetStatusCode(v int32) *GetSceneBuildTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponse) SetBody(v *GetSceneBuildTaskStatusResponseBody) *GetSceneBuildTaskStatusResponse {
	s.Body = v
	return s
}

type GetScenePackUrlRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetScenePackUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScenePackUrlRequest) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlRequest) SetSceneId(v string) *GetScenePackUrlRequest {
	s.SceneId = &v
	return s
}

type GetScenePackUrlResponseBody struct {
	Code      *int64                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetScenePackUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePackUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScenePackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponseBody) SetCode(v int64) *GetScenePackUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePackUrlResponseBody) SetData(v *GetScenePackUrlResponseBodyData) *GetScenePackUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePackUrlResponseBody) SetMessage(v string) *GetScenePackUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePackUrlResponseBody) SetRequestId(v string) *GetScenePackUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePackUrlResponseBody) SetSuccess(v bool) *GetScenePackUrlResponseBody {
	s.Success = &v
	return s
}

type GetScenePackUrlResponseBodyData struct {
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Valid  *bool   `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetScenePackUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScenePackUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponseBodyData) SetExpire(v string) *GetScenePackUrlResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetScenePackUrlResponseBodyData) SetUrl(v string) *GetScenePackUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetScenePackUrlResponseBodyData) SetValid(v bool) *GetScenePackUrlResponseBodyData {
	s.Valid = &v
	return s
}

type GetScenePackUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScenePackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScenePackUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScenePackUrlResponse) GoString() string {
	return s.String()
}

func (s *GetScenePackUrlResponse) SetHeaders(v map[string]*string) *GetScenePackUrlResponse {
	s.Headers = v
	return s
}

func (s *GetScenePackUrlResponse) SetStatusCode(v int32) *GetScenePackUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePackUrlResponse) SetBody(v *GetScenePackUrlResponseBody) *GetScenePackUrlResponse {
	s.Body = v
	return s
}

type GetScenePreviewDataRequest struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	ShowTag      *bool   `json:"ShowTag,omitempty" xml:"ShowTag,omitempty"`
}

func (s GetScenePreviewDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataRequest) SetDomain(v string) *GetScenePreviewDataRequest {
	s.Domain = &v
	return s
}

func (s *GetScenePreviewDataRequest) SetEnabled(v bool) *GetScenePreviewDataRequest {
	s.Enabled = &v
	return s
}

func (s *GetScenePreviewDataRequest) SetPreviewToken(v string) *GetScenePreviewDataRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetScenePreviewDataRequest) SetShowTag(v bool) *GetScenePreviewDataRequest {
	s.ShowTag = &v
	return s
}

type GetScenePreviewDataResponseBody struct {
	Code      *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetScenePreviewDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePreviewDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBody) SetCode(v int64) *GetScenePreviewDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetData(v *GetScenePreviewDataResponseBodyData) *GetScenePreviewDataResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetMessage(v string) *GetScenePreviewDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetRequestId(v string) *GetScenePreviewDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewDataResponseBody) SetSuccess(v bool) *GetScenePreviewDataResponseBody {
	s.Success = &v
	return s
}

type GetScenePreviewDataResponseBodyData struct {
	Model *GetScenePreviewDataResponseBodyDataModel  `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Tags  []*GetScenePreviewDataResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetScenePreviewDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyData) SetModel(v *GetScenePreviewDataResponseBodyDataModel) *GetScenePreviewDataResponseBodyData {
	s.Model = v
	return s
}

func (s *GetScenePreviewDataResponseBodyData) SetTags(v []*GetScenePreviewDataResponseBodyDataTags) *GetScenePreviewDataResponseBodyData {
	s.Tags = v
	return s
}

type GetScenePreviewDataResponseBodyDataModel struct {
	ModelPath        *string                                             `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	PanoList         []*GetScenePreviewDataResponseBodyDataModelPanoList `json:"PanoList,omitempty" xml:"PanoList,omitempty" type:"Repeated"`
	TextureModelPath *string                                             `json:"TextureModelPath,omitempty" xml:"TextureModelPath,omitempty"`
	TexturePanoPath  *string                                             `json:"TexturePanoPath,omitempty" xml:"TexturePanoPath,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataModel) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataModel) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetModelPath(v string) *GetScenePreviewDataResponseBodyDataModel {
	s.ModelPath = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetPanoList(v []*GetScenePreviewDataResponseBodyDataModelPanoList) *GetScenePreviewDataResponseBodyDataModel {
	s.PanoList = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetTextureModelPath(v string) *GetScenePreviewDataResponseBodyDataModel {
	s.TextureModelPath = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModel) SetTexturePanoPath(v string) *GetScenePreviewDataResponseBodyDataModel {
	s.TexturePanoPath = &v
	return s
}

type GetScenePreviewDataResponseBodyDataModelPanoList struct {
	CurRoomPicList []*string                                                 `json:"CurRoomPicList,omitempty" xml:"CurRoomPicList,omitempty" type:"Repeated"`
	Enabled        *bool                                                     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FloorIdx       *string                                                   `json:"FloorIdx,omitempty" xml:"FloorIdx,omitempty"`
	Id             *string                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	MainImage      *bool                                                     `json:"MainImage,omitempty" xml:"MainImage,omitempty"`
	Neighbours     []*string                                                 `json:"Neighbours,omitempty" xml:"Neighbours,omitempty" type:"Repeated"`
	Position       *GetScenePreviewDataResponseBodyDataModelPanoListPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Struct"`
	RawName        *string                                                   `json:"RawName,omitempty" xml:"RawName,omitempty"`
	Resource       *string                                                   `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RoomIdx        *string                                                   `json:"RoomIdx,omitempty" xml:"RoomIdx,omitempty"`
	SubSceneId     *string                                                   `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	Token          *string                                                   `json:"Token,omitempty" xml:"Token,omitempty"`
	VirtualId      *string                                                   `json:"VirtualId,omitempty" xml:"VirtualId,omitempty"`
	VirtualName    *string                                                   `json:"VirtualName,omitempty" xml:"VirtualName,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataModelPanoList) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataModelPanoList) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetCurRoomPicList(v []*string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.CurRoomPicList = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetEnabled(v bool) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Enabled = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetFloorIdx(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.FloorIdx = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetId(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Id = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetMainImage(v bool) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.MainImage = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetNeighbours(v []*string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Neighbours = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetPosition(v *GetScenePreviewDataResponseBodyDataModelPanoListPosition) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Position = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetRawName(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.RawName = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetResource(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Resource = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetRoomIdx(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.RoomIdx = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetSubSceneId(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.SubSceneId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetToken(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.Token = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetVirtualId(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.VirtualId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoList) SetVirtualName(v string) *GetScenePreviewDataResponseBodyDataModelPanoList {
	s.VirtualName = &v
	return s
}

type GetScenePreviewDataResponseBodyDataModelPanoListPosition struct {
	Rotation  []*float64 `json:"Rotation,omitempty" xml:"Rotation,omitempty" type:"Repeated"`
	Spot      []*float64 `json:"Spot,omitempty" xml:"Spot,omitempty" type:"Repeated"`
	Viewpoint []*float64 `json:"Viewpoint,omitempty" xml:"Viewpoint,omitempty" type:"Repeated"`
}

func (s GetScenePreviewDataResponseBodyDataModelPanoListPosition) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataModelPanoListPosition) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) SetRotation(v []*float64) *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	s.Rotation = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) SetSpot(v []*float64) *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	s.Spot = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataModelPanoListPosition) SetViewpoint(v []*float64) *GetScenePreviewDataResponseBodyDataModelPanoListPosition {
	s.Viewpoint = v
	return s
}

type GetScenePreviewDataResponseBodyDataTags struct {
	Config           *GetScenePreviewDataResponseBodyDataTagsConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Id               *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Position         []*float64                                     `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	PositionPanoCube []*float64                                     `json:"PositionPanoCube,omitempty" xml:"PositionPanoCube,omitempty" type:"Repeated"`
	Type             *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetConfig(v *GetScenePreviewDataResponseBodyDataTagsConfig) *GetScenePreviewDataResponseBodyDataTags {
	s.Config = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetId(v string) *GetScenePreviewDataResponseBodyDataTags {
	s.Id = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetPosition(v []*float64) *GetScenePreviewDataResponseBodyDataTags {
	s.Position = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetPositionPanoCube(v []*float64) *GetScenePreviewDataResponseBodyDataTags {
	s.PositionPanoCube = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTags) SetType(v string) *GetScenePreviewDataResponseBodyDataTags {
	s.Type = &v
	return s
}

type GetScenePreviewDataResponseBodyDataTagsConfig struct {
	BackgroundColor   *string                                                    `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	ButtonConfig      *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig `json:"ButtonConfig,omitempty" xml:"ButtonConfig,omitempty" type:"Struct"`
	Content           *string                                                    `json:"Content,omitempty" xml:"Content,omitempty"`
	FormImgSize       []*int64                                                   `json:"FormImgSize,omitempty" xml:"FormImgSize,omitempty" type:"Repeated"`
	FormJumpType      *bool                                                      `json:"FormJumpType,omitempty" xml:"FormJumpType,omitempty"`
	FormSelectImgType *string                                                    `json:"FormSelectImgType,omitempty" xml:"FormSelectImgType,omitempty"`
	Images            []*string                                                  `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	IsTagVisibleBy3d  *bool                                                      `json:"IsTagVisibleBy3d,omitempty" xml:"IsTagVisibleBy3d,omitempty"`
	Link              *string                                                    `json:"Link,omitempty" xml:"Link,omitempty"`
	PanoId            *string                                                    `json:"PanoId,omitempty" xml:"PanoId,omitempty"`
	Position          []*float64                                                 `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	PositionPanoCube  []*float64                                                 `json:"PositionPanoCube,omitempty" xml:"PositionPanoCube,omitempty" type:"Repeated"`
	RelatedPanoIds    []*string                                                  `json:"RelatedPanoIds,omitempty" xml:"RelatedPanoIds,omitempty" type:"Repeated"`
	SceneId           *int64                                                     `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Title             *string                                                    `json:"Title,omitempty" xml:"Title,omitempty"`
	Video             *string                                                    `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataTagsConfig) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataTagsConfig) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetBackgroundColor(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.BackgroundColor = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetButtonConfig(v *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.ButtonConfig = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetContent(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Content = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetFormImgSize(v []*int64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.FormImgSize = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetFormJumpType(v bool) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.FormJumpType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetFormSelectImgType(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.FormSelectImgType = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetImages(v []*string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Images = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetIsTagVisibleBy3d(v bool) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.IsTagVisibleBy3d = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetLink(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Link = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetPanoId(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.PanoId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetPosition(v []*float64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Position = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetPositionPanoCube(v []*float64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.PositionPanoCube = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetRelatedPanoIds(v []*string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.RelatedPanoIds = v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetSceneId(v int64) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.SceneId = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetTitle(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Title = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfig) SetVideo(v string) *GetScenePreviewDataResponseBodyDataTagsConfig {
	s.Video = &v
	return s
}

type GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig struct {
	CustomText *string `json:"CustomText,omitempty" xml:"CustomText,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) SetCustomText(v string) *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig {
	s.CustomText = &v
	return s
}

func (s *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig) SetType(v string) *GetScenePreviewDataResponseBodyDataTagsConfigButtonConfig {
	s.Type = &v
	return s
}

type GetScenePreviewDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScenePreviewDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScenePreviewDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewDataResponse) GoString() string {
	return s.String()
}

func (s *GetScenePreviewDataResponse) SetHeaders(v map[string]*string) *GetScenePreviewDataResponse {
	s.Headers = v
	return s
}

func (s *GetScenePreviewDataResponse) SetStatusCode(v int32) *GetScenePreviewDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePreviewDataResponse) SetBody(v *GetScenePreviewDataResponseBody) *GetScenePreviewDataResponse {
	s.Body = v
	return s
}

type GetScenePreviewInfoRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled    *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
}

func (s GetScenePreviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoRequest) SetDomain(v string) *GetScenePreviewInfoRequest {
	s.Domain = &v
	return s
}

func (s *GetScenePreviewInfoRequest) SetEnabled(v bool) *GetScenePreviewInfoRequest {
	s.Enabled = &v
	return s
}

func (s *GetScenePreviewInfoRequest) SetModelToken(v string) *GetScenePreviewInfoRequest {
	s.ModelToken = &v
	return s
}

type GetScenePreviewInfoResponseBody struct {
	Code      *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetScenePreviewInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePreviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBody) SetCode(v int64) *GetScenePreviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetData(v *GetScenePreviewInfoResponseBodyData) *GetScenePreviewInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetMessage(v string) *GetScenePreviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetRequestId(v string) *GetScenePreviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetSuccess(v bool) *GetScenePreviewInfoResponseBody {
	s.Success = &v
	return s
}

type GetScenePreviewInfoResponseBodyData struct {
	ModelPath        *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	PanoList         *string `json:"PanoList,omitempty" xml:"PanoList,omitempty"`
	TextureModelPath *string `json:"TextureModelPath,omitempty" xml:"TextureModelPath,omitempty"`
	TexturePanoPath  *string `json:"TexturePanoPath,omitempty" xml:"TexturePanoPath,omitempty"`
}

func (s GetScenePreviewInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBodyData) SetModelPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.ModelPath = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetPanoList(v string) *GetScenePreviewInfoResponseBodyData {
	s.PanoList = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetTextureModelPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.TextureModelPath = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetTexturePanoPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.TexturePanoPath = &v
	return s
}

type GetScenePreviewInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScenePreviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScenePreviewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoResponse) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponse) SetHeaders(v map[string]*string) *GetScenePreviewInfoResponse {
	s.Headers = v
	return s
}

func (s *GetScenePreviewInfoResponse) SetStatusCode(v int32) *GetScenePreviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePreviewInfoResponse) SetBody(v *GetScenePreviewInfoResponseBody) *GetScenePreviewInfoResponse {
	s.Body = v
	return s
}

type GetScenePreviewResourceRequest struct {
	Draft        *bool   `json:"Draft,omitempty" xml:"Draft,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s GetScenePreviewResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewResourceRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceRequest) SetDraft(v bool) *GetScenePreviewResourceRequest {
	s.Draft = &v
	return s
}

func (s *GetScenePreviewResourceRequest) SetPreviewToken(v string) *GetScenePreviewResourceRequest {
	s.PreviewToken = &v
	return s
}

type GetScenePreviewResourceResponseBody struct {
	Code      *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetScenePreviewResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScenePreviewResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBody) SetCode(v int64) *GetScenePreviewResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetData(v *GetScenePreviewResourceResponseBodyData) *GetScenePreviewResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetMessage(v string) *GetScenePreviewResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetRequestId(v string) *GetScenePreviewResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewResourceResponseBody) SetSuccess(v bool) *GetScenePreviewResourceResponseBody {
	s.Success = &v
	return s
}

type GetScenePreviewResourceResponseBodyData struct {
	Name              *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceDirectory *GetScenePreviewResourceResponseBodyDataResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
	Version           *string                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetScenePreviewResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBodyData) SetName(v string) *GetScenePreviewResourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyData) SetResourceDirectory(v *GetScenePreviewResourceResponseBodyDataResourceDirectory) *GetScenePreviewResourceResponseBodyData {
	s.ResourceDirectory = v
	return s
}

func (s *GetScenePreviewResourceResponseBodyData) SetVersion(v string) *GetScenePreviewResourceResponseBodyData {
	s.Version = &v
	return s
}

type GetScenePreviewResourceResponseBodyDataResourceDirectory struct {
	HotspotTagConfig *string `json:"HotspotTagConfig,omitempty" xml:"HotspotTagConfig,omitempty"`
	ModelConfig      *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	OrthomapConfig   *string `json:"OrthomapConfig,omitempty" xml:"OrthomapConfig,omitempty"`
	RootPath         *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
}

func (s GetScenePreviewResourceResponseBodyDataResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewResourceResponseBodyDataResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetHotspotTagConfig(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.HotspotTagConfig = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetModelConfig(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.ModelConfig = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetOrthomapConfig(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.OrthomapConfig = &v
	return s
}

func (s *GetScenePreviewResourceResponseBodyDataResourceDirectory) SetRootPath(v string) *GetScenePreviewResourceResponseBodyDataResourceDirectory {
	s.RootPath = &v
	return s
}

type GetScenePreviewResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScenePreviewResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScenePreviewResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewResourceResponse) GoString() string {
	return s.String()
}

func (s *GetScenePreviewResourceResponse) SetHeaders(v map[string]*string) *GetScenePreviewResourceResponse {
	s.Headers = v
	return s
}

func (s *GetScenePreviewResourceResponse) SetStatusCode(v int32) *GetScenePreviewResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScenePreviewResourceResponse) SetBody(v *GetScenePreviewResourceResponseBody) *GetScenePreviewResourceResponse {
	s.Body = v
	return s
}

type GetSingleConnDataRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetSingleConnDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSingleConnDataRequest) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataRequest) SetSubSceneId(v string) *GetSingleConnDataRequest {
	s.SubSceneId = &v
	return s
}

type GetSingleConnDataResponseBody struct {
	Code      *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*GetSingleConnDataResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Version   *string                              `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetSingleConnDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSingleConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponseBody) SetCode(v int64) *GetSingleConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetList(v []*GetSingleConnDataResponseBodyList) *GetSingleConnDataResponseBody {
	s.List = v
	return s
}

func (s *GetSingleConnDataResponseBody) SetMessage(v string) *GetSingleConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetRequestId(v string) *GetSingleConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetSuccess(v bool) *GetSingleConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetVersion(v string) *GetSingleConnDataResponseBody {
	s.Version = &v
	return s
}

type GetSingleConnDataResponseBodyList struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MapId *string `json:"MapId,omitempty" xml:"MapId,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSingleConnDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetSingleConnDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponseBodyList) SetId(v string) *GetSingleConnDataResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetSingleConnDataResponseBodyList) SetMapId(v string) *GetSingleConnDataResponseBodyList {
	s.MapId = &v
	return s
}

func (s *GetSingleConnDataResponseBodyList) SetType(v string) *GetSingleConnDataResponseBodyList {
	s.Type = &v
	return s
}

type GetSingleConnDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSingleConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSingleConnDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSingleConnDataResponse) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponse) SetHeaders(v map[string]*string) *GetSingleConnDataResponse {
	s.Headers = v
	return s
}

func (s *GetSingleConnDataResponse) SetStatusCode(v int32) *GetSingleConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSingleConnDataResponse) SetBody(v *GetSingleConnDataResponseBody) *GetSingleConnDataResponse {
	s.Body = v
	return s
}

type GetSourcePackStatusRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSourcePackStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSourcePackStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusRequest) SetTaskId(v string) *GetSourcePackStatusRequest {
	s.TaskId = &v
	return s
}

type GetSourcePackStatusResponseBody struct {
	Code      *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSourcePackStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Url       *string                              `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSourcePackStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSourcePackStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponseBody) SetCode(v int64) *GetSourcePackStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetData(v *GetSourcePackStatusResponseBodyData) *GetSourcePackStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetMessage(v string) *GetSourcePackStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetRequestId(v string) *GetSourcePackStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetSuccess(v bool) *GetSourcePackStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetUrl(v string) *GetSourcePackStatusResponseBody {
	s.Url = &v
	return s
}

type GetSourcePackStatusResponseBodyData struct {
	Progress *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSourcePackStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSourcePackStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponseBodyData) SetProgress(v int64) *GetSourcePackStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyData) SetStatus(v string) *GetSourcePackStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetSourcePackStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSourcePackStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSourcePackStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSourcePackStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponse) SetHeaders(v map[string]*string) *GetSourcePackStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSourcePackStatusResponse) SetStatusCode(v int32) *GetSourcePackStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourcePackStatusResponse) SetBody(v *GetSourcePackStatusResponseBody) *GetSourcePackStatusResponse {
	s.Body = v
	return s
}

type GetSubSceneTaskStatusRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetSubSceneTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubSceneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusRequest) SetSubSceneId(v string) *GetSubSceneTaskStatusRequest {
	s.SubSceneId = &v
	return s
}

type GetSubSceneTaskStatusResponseBody struct {
	Code      *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	List      []*GetSubSceneTaskStatusResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubSceneTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBody) SetCode(v int64) *GetSubSceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetList(v []*GetSubSceneTaskStatusResponseBodyList) *GetSubSceneTaskStatusResponseBody {
	s.List = v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetMessage(v string) *GetSubSceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetRequestId(v string) *GetSubSceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetSuccess(v bool) *GetSubSceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

type GetSubSceneTaskStatusResponseBodyList struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSubSceneTaskStatusResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetErrorCode(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.ErrorCode = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetErrorMsg(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.ErrorMsg = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetSceneId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.SceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetStatus(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Status = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetSubSceneId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.SubSceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetType(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Type = &v
	return s
}

type GetSubSceneTaskStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSubSceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubSceneTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubSceneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponse) SetHeaders(v map[string]*string) *GetSubSceneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSubSceneTaskStatusResponse) SetStatusCode(v int32) *GetSubSceneTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubSceneTaskStatusResponse) SetBody(v *GetSubSceneTaskStatusResponseBody) *GetSubSceneTaskStatusResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *GetTaskStatusResponseBody) SetErrorCode(v string) *GetTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrorMsg(v string) *GetTaskStatusResponseBody {
	s.ErrorMsg = &v
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

func (s *GetTaskStatusResponseBody) SetStatus(v string) *GetTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetType(v string) *GetTaskStatusResponseBody {
	s.Type = &v
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

type GetWindowConfigRequest struct {
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s GetWindowConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWindowConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWindowConfigRequest) SetPreviewToken(v string) *GetWindowConfigRequest {
	s.PreviewToken = &v
	return s
}

type GetWindowConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMessage   *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ObjectString *string                `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWindowConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWindowConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWindowConfigResponseBody) SetData(v map[string]interface{}) *GetWindowConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetWindowConfigResponseBody) SetErrMessage(v string) *GetWindowConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetWindowConfigResponseBody) SetObjectString(v string) *GetWindowConfigResponseBody {
	s.ObjectString = &v
	return s
}

func (s *GetWindowConfigResponseBody) SetRequestId(v string) *GetWindowConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWindowConfigResponseBody) SetSuccess(v bool) *GetWindowConfigResponseBody {
	s.Success = &v
	return s
}

type GetWindowConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWindowConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWindowConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWindowConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWindowConfigResponse) SetHeaders(v map[string]*string) *GetWindowConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWindowConfigResponse) SetStatusCode(v int32) *GetWindowConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWindowConfigResponse) SetBody(v *GetWindowConfigResponseBody) *GetWindowConfigResponse {
	s.Body = v
	return s
}

type LabelBuildRequest struct {
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ModelStyle        *string `json:"ModelStyle,omitempty" xml:"ModelStyle,omitempty"`
	OptimizeWallWidth *string `json:"OptimizeWallWidth,omitempty" xml:"OptimizeWallWidth,omitempty"`
	PlanStyle         *string `json:"PlanStyle,omitempty" xml:"PlanStyle,omitempty"`
	SceneId           *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	WallHeight        *int64  `json:"WallHeight,omitempty" xml:"WallHeight,omitempty"`
}

func (s LabelBuildRequest) String() string {
	return tea.Prettify(s)
}

func (s LabelBuildRequest) GoString() string {
	return s.String()
}

func (s *LabelBuildRequest) SetMode(v string) *LabelBuildRequest {
	s.Mode = &v
	return s
}

func (s *LabelBuildRequest) SetModelStyle(v string) *LabelBuildRequest {
	s.ModelStyle = &v
	return s
}

func (s *LabelBuildRequest) SetOptimizeWallWidth(v string) *LabelBuildRequest {
	s.OptimizeWallWidth = &v
	return s
}

func (s *LabelBuildRequest) SetPlanStyle(v string) *LabelBuildRequest {
	s.PlanStyle = &v
	return s
}

func (s *LabelBuildRequest) SetSceneId(v string) *LabelBuildRequest {
	s.SceneId = &v
	return s
}

func (s *LabelBuildRequest) SetWallHeight(v int64) *LabelBuildRequest {
	s.WallHeight = &v
	return s
}

type LabelBuildResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s LabelBuildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LabelBuildResponseBody) GoString() string {
	return s.String()
}

func (s *LabelBuildResponseBody) SetCode(v int64) *LabelBuildResponseBody {
	s.Code = &v
	return s
}

func (s *LabelBuildResponseBody) SetMessage(v string) *LabelBuildResponseBody {
	s.Message = &v
	return s
}

func (s *LabelBuildResponseBody) SetRequestId(v string) *LabelBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *LabelBuildResponseBody) SetSuccess(v bool) *LabelBuildResponseBody {
	s.Success = &v
	return s
}

func (s *LabelBuildResponseBody) SetTaskId(v string) *LabelBuildResponseBody {
	s.TaskId = &v
	return s
}

type LabelBuildResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LabelBuildResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LabelBuildResponse) String() string {
	return tea.Prettify(s)
}

func (s LabelBuildResponse) GoString() string {
	return s.String()
}

func (s *LabelBuildResponse) SetHeaders(v map[string]*string) *LabelBuildResponse {
	s.Headers = v
	return s
}

func (s *LabelBuildResponse) SetStatusCode(v int32) *LabelBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *LabelBuildResponse) SetBody(v *LabelBuildResponseBody) *LabelBuildResponse {
	s.Body = v
	return s
}

type LinkImageRequest struct {
	CameraHeight *int32  `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SubSceneId   *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s LinkImageRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkImageRequest) GoString() string {
	return s.String()
}

func (s *LinkImageRequest) SetCameraHeight(v int32) *LinkImageRequest {
	s.CameraHeight = &v
	return s
}

func (s *LinkImageRequest) SetFileName(v string) *LinkImageRequest {
	s.FileName = &v
	return s
}

func (s *LinkImageRequest) SetPlatform(v string) *LinkImageRequest {
	s.Platform = &v
	return s
}

func (s *LinkImageRequest) SetSubSceneId(v string) *LinkImageRequest {
	s.SubSceneId = &v
	return s
}

type LinkImageResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LinkImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LinkImageResponseBody) GoString() string {
	return s.String()
}

func (s *LinkImageResponseBody) SetCode(v int64) *LinkImageResponseBody {
	s.Code = &v
	return s
}

func (s *LinkImageResponseBody) SetMessage(v string) *LinkImageResponseBody {
	s.Message = &v
	return s
}

func (s *LinkImageResponseBody) SetRequestId(v string) *LinkImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *LinkImageResponseBody) SetResourceId(v string) *LinkImageResponseBody {
	s.ResourceId = &v
	return s
}

func (s *LinkImageResponseBody) SetSuccess(v bool) *LinkImageResponseBody {
	s.Success = &v
	return s
}

type LinkImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LinkImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LinkImageResponse) String() string {
	return tea.Prettify(s)
}

func (s LinkImageResponse) GoString() string {
	return s.String()
}

func (s *LinkImageResponse) SetHeaders(v map[string]*string) *LinkImageResponse {
	s.Headers = v
	return s
}

func (s *LinkImageResponse) SetStatusCode(v int32) *LinkImageResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkImageResponse) SetBody(v *LinkImageResponseBody) *LinkImageResponse {
	s.Body = v
	return s
}

type ListProjectRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNum  *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) SetName(v string) *ListProjectRequest {
	s.Name = &v
	return s
}

func (s *ListProjectRequest) SetPageNum(v int64) *ListProjectRequest {
	s.PageNum = &v
	return s
}

func (s *ListProjectRequest) SetPageSize(v int64) *ListProjectRequest {
	s.PageSize = &v
	return s
}

type ListProjectResponseBody struct {
	Code        *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Count       *int64                         `json:"Count,omitempty" xml:"Count,omitempty"`
	CurrentPage *int64                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasNext     *bool                          `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	List        []*ListProjectResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message     *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalPage   *int64                         `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) SetCode(v int64) *ListProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetCurrentPage(v int64) *ListProjectResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListProjectResponseBody) SetHasNext(v bool) *ListProjectResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListProjectResponseBody) SetList(v []*ListProjectResponseBodyList) *ListProjectResponseBody {
	s.List = v
	return s
}

func (s *ListProjectResponseBody) SetMessage(v string) *ListProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectResponseBody) SetSuccess(v bool) *ListProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectResponseBody) SetTotalPage(v int64) *ListProjectResponseBody {
	s.TotalPage = &v
	return s
}

type ListProjectResponseBodyList struct {
	BusinessId   *int64  `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListProjectResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyList) SetBusinessId(v int64) *ListProjectResponseBodyList {
	s.BusinessId = &v
	return s
}

func (s *ListProjectResponseBodyList) SetBusinessName(v string) *ListProjectResponseBodyList {
	s.BusinessName = &v
	return s
}

func (s *ListProjectResponseBodyList) SetCreateTime(v int64) *ListProjectResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListProjectResponseBodyList) SetId(v string) *ListProjectResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListProjectResponseBodyList) SetModifiedTime(v int64) *ListProjectResponseBodyList {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectResponseBodyList) SetName(v string) *ListProjectResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListProjectResponseBodyList) SetToken(v string) *ListProjectResponseBodyList {
	s.Token = &v
	return s
}

type ListProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponse) GoString() string {
	return s.String()
}

func (s *ListProjectResponse) SetHeaders(v map[string]*string) *ListProjectResponse {
	s.Headers = v
	return s
}

func (s *ListProjectResponse) SetStatusCode(v int32) *ListProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectResponse) SetBody(v *ListProjectResponseBody) *ListProjectResponse {
	s.Body = v
	return s
}

type ListSceneRequest struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNum   *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSceneRequest) GoString() string {
	return s.String()
}

func (s *ListSceneRequest) SetName(v string) *ListSceneRequest {
	s.Name = &v
	return s
}

func (s *ListSceneRequest) SetPageNum(v int64) *ListSceneRequest {
	s.PageNum = &v
	return s
}

func (s *ListSceneRequest) SetPageSize(v int64) *ListSceneRequest {
	s.PageSize = &v
	return s
}

func (s *ListSceneRequest) SetProjectId(v string) *ListSceneRequest {
	s.ProjectId = &v
	return s
}

type ListSceneResponseBody struct {
	Code        *int64                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Count       *int64                       `json:"Count,omitempty" xml:"Count,omitempty"`
	CurrentPage *int64                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasNext     *bool                        `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	List        []*ListSceneResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message     *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalPage   *int64                       `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBody) SetCode(v int64) *ListSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSceneResponseBody) SetCount(v int64) *ListSceneResponseBody {
	s.Count = &v
	return s
}

func (s *ListSceneResponseBody) SetCurrentPage(v int64) *ListSceneResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSceneResponseBody) SetHasNext(v bool) *ListSceneResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListSceneResponseBody) SetList(v []*ListSceneResponseBodyList) *ListSceneResponseBody {
	s.List = v
	return s
}

func (s *ListSceneResponseBody) SetMessage(v string) *ListSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSceneResponseBody) SetRequestId(v string) *ListSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSceneResponseBody) SetSuccess(v bool) *ListSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListSceneResponseBody) SetTotalPage(v int64) *ListSceneResponseBody {
	s.TotalPage = &v
	return s
}

type ListSceneResponseBodyList struct {
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	Published    *bool   `json:"Published,omitempty" xml:"Published,omitempty"`
	SourceNum    *int64  `json:"SourceNum,omitempty" xml:"SourceNum,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName   *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	SubSceneNum  *int64  `json:"SubSceneNum,omitempty" xml:"SubSceneNum,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSceneResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSceneResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBodyList) SetCoverUrl(v string) *ListSceneResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListSceneResponseBodyList) SetGmtCreate(v int64) *ListSceneResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListSceneResponseBodyList) SetGmtModified(v int64) *ListSceneResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListSceneResponseBodyList) SetId(v string) *ListSceneResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSceneResponseBodyList) SetName(v string) *ListSceneResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListSceneResponseBodyList) SetPreviewToken(v string) *ListSceneResponseBodyList {
	s.PreviewToken = &v
	return s
}

func (s *ListSceneResponseBodyList) SetPublished(v bool) *ListSceneResponseBodyList {
	s.Published = &v
	return s
}

func (s *ListSceneResponseBodyList) SetSourceNum(v int64) *ListSceneResponseBodyList {
	s.SourceNum = &v
	return s
}

func (s *ListSceneResponseBodyList) SetStatus(v string) *ListSceneResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListSceneResponseBodyList) SetStatusName(v string) *ListSceneResponseBodyList {
	s.StatusName = &v
	return s
}

func (s *ListSceneResponseBodyList) SetSubSceneNum(v int64) *ListSceneResponseBodyList {
	s.SubSceneNum = &v
	return s
}

func (s *ListSceneResponseBodyList) SetType(v string) *ListSceneResponseBodyList {
	s.Type = &v
	return s
}

type ListSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSceneResponse) GoString() string {
	return s.String()
}

func (s *ListSceneResponse) SetHeaders(v map[string]*string) *ListSceneResponse {
	s.Headers = v
	return s
}

func (s *ListSceneResponse) SetStatusCode(v int32) *ListSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSceneResponse) SetBody(v *ListSceneResponseBody) *ListSceneResponse {
	s.Body = v
	return s
}

type ListSubSceneRequest struct {
	PageNum        *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId        *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	ShowLayoutData *bool   `json:"ShowLayoutData,omitempty" xml:"ShowLayoutData,omitempty"`
	SortField      *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneRequest) GoString() string {
	return s.String()
}

func (s *ListSubSceneRequest) SetPageNum(v int64) *ListSubSceneRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubSceneRequest) SetPageSize(v int64) *ListSubSceneRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubSceneRequest) SetSceneId(v string) *ListSubSceneRequest {
	s.SceneId = &v
	return s
}

func (s *ListSubSceneRequest) SetShowLayoutData(v bool) *ListSubSceneRequest {
	s.ShowLayoutData = &v
	return s
}

func (s *ListSubSceneRequest) SetSortField(v string) *ListSubSceneRequest {
	s.SortField = &v
	return s
}

type ListSubSceneResponseBody struct {
	Code        *int64                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Count       *int64                          `json:"Count,omitempty" xml:"Count,omitempty"`
	CurrentPage *int64                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasNext     *bool                           `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	List        []*ListSubSceneResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message     *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalPage   *int64                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBody) SetCode(v int64) *ListSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCount(v int64) *ListSubSceneResponseBody {
	s.Count = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCurrentPage(v int64) *ListSubSceneResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSubSceneResponseBody) SetHasNext(v bool) *ListSubSceneResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListSubSceneResponseBody) SetList(v []*ListSubSceneResponseBodyList) *ListSubSceneResponseBody {
	s.List = v
	return s
}

func (s *ListSubSceneResponseBody) SetMessage(v string) *ListSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubSceneResponseBody) SetRequestId(v string) *ListSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubSceneResponseBody) SetSuccess(v bool) *ListSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubSceneResponseBody) SetTotalPage(v int64) *ListSubSceneResponseBody {
	s.TotalPage = &v
	return s
}

type ListSubSceneResponseBodyList struct {
	BaseImageUrl *string `json:"BaseImageUrl,omitempty" xml:"BaseImageUrl,omitempty"`
	CoverUrl     *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	CubemapPath  *string `json:"CubemapPath,omitempty" xml:"CubemapPath,omitempty"`
	Deleted      *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LayoutData   *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginUrl    *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSubSceneResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBodyList) SetBaseImageUrl(v string) *ListSubSceneResponseBodyList {
	s.BaseImageUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetCoverUrl(v string) *ListSubSceneResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetCubemapPath(v string) *ListSubSceneResponseBodyList {
	s.CubemapPath = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetDeleted(v bool) *ListSubSceneResponseBodyList {
	s.Deleted = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetGmtCreate(v int64) *ListSubSceneResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetGmtModified(v int64) *ListSubSceneResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetId(v string) *ListSubSceneResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetLayoutData(v string) *ListSubSceneResponseBodyList {
	s.LayoutData = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetName(v string) *ListSubSceneResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetOriginUrl(v string) *ListSubSceneResponseBodyList {
	s.OriginUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetResourceId(v string) *ListSubSceneResponseBodyList {
	s.ResourceId = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetResourceName(v string) *ListSubSceneResponseBodyList {
	s.ResourceName = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetStatus(v int64) *ListSubSceneResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetType(v string) *ListSubSceneResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetUrl(v string) *ListSubSceneResponseBodyList {
	s.Url = &v
	return s
}

type ListSubSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneResponse) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponse) SetHeaders(v map[string]*string) *ListSubSceneResponse {
	s.Headers = v
	return s
}

func (s *ListSubSceneResponse) SetStatusCode(v int32) *ListSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubSceneResponse) SetBody(v *ListSubSceneResponseBody) *ListSubSceneResponse {
	s.Body = v
	return s
}

type OptimizeRightAngleRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s OptimizeRightAngleRequest) String() string {
	return tea.Prettify(s)
}

func (s OptimizeRightAngleRequest) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleRequest) SetSubSceneId(v string) *OptimizeRightAngleRequest {
	s.SubSceneId = &v
	return s
}

type OptimizeRightAngleResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OptimizeRightAngleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OptimizeRightAngleResponseBody) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleResponseBody) SetCode(v int64) *OptimizeRightAngleResponseBody {
	s.Code = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetMessage(v string) *OptimizeRightAngleResponseBody {
	s.Message = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetRequestId(v string) *OptimizeRightAngleResponseBody {
	s.RequestId = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetSuccess(v bool) *OptimizeRightAngleResponseBody {
	s.Success = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetTaskId(v string) *OptimizeRightAngleResponseBody {
	s.TaskId = &v
	return s
}

type OptimizeRightAngleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OptimizeRightAngleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OptimizeRightAngleResponse) String() string {
	return tea.Prettify(s)
}

func (s OptimizeRightAngleResponse) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleResponse) SetHeaders(v map[string]*string) *OptimizeRightAngleResponse {
	s.Headers = v
	return s
}

func (s *OptimizeRightAngleResponse) SetStatusCode(v int32) *OptimizeRightAngleResponse {
	s.StatusCode = &v
	return s
}

func (s *OptimizeRightAngleResponse) SetBody(v *OptimizeRightAngleResponseBody) *OptimizeRightAngleResponse {
	s.Body = v
	return s
}

type PackSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PackSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s PackSceneRequest) GoString() string {
	return s.String()
}

func (s *PackSceneRequest) SetSceneId(v string) *PackSceneRequest {
	s.SceneId = &v
	return s
}

func (s *PackSceneRequest) SetType(v string) *PackSceneRequest {
	s.Type = &v
	return s
}

type PackSceneResponseBody struct {
	Code      *int64                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PackSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PackSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PackSceneResponseBody) GoString() string {
	return s.String()
}

func (s *PackSceneResponseBody) SetCode(v int64) *PackSceneResponseBody {
	s.Code = &v
	return s
}

func (s *PackSceneResponseBody) SetData(v *PackSceneResponseBodyData) *PackSceneResponseBody {
	s.Data = v
	return s
}

func (s *PackSceneResponseBody) SetMessage(v string) *PackSceneResponseBody {
	s.Message = &v
	return s
}

func (s *PackSceneResponseBody) SetRequestId(v string) *PackSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *PackSceneResponseBody) SetSuccess(v bool) *PackSceneResponseBody {
	s.Success = &v
	return s
}

type PackSceneResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PackSceneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PackSceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *PackSceneResponseBodyData) SetTaskId(v string) *PackSceneResponseBodyData {
	s.TaskId = &v
	return s
}

type PackSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PackSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PackSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s PackSceneResponse) GoString() string {
	return s.String()
}

func (s *PackSceneResponse) SetHeaders(v map[string]*string) *PackSceneResponse {
	s.Headers = v
	return s
}

func (s *PackSceneResponse) SetStatusCode(v int32) *PackSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *PackSceneResponse) SetBody(v *PackSceneResponseBody) *PackSceneResponse {
	s.Body = v
	return s
}

type PackSourceRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PackSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s PackSourceRequest) GoString() string {
	return s.String()
}

func (s *PackSourceRequest) SetSceneId(v string) *PackSourceRequest {
	s.SceneId = &v
	return s
}

type PackSourceResponseBody struct {
	Code      *int64                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PackSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PackSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PackSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PackSourceResponseBody) SetCode(v int64) *PackSourceResponseBody {
	s.Code = &v
	return s
}

func (s *PackSourceResponseBody) SetData(v *PackSourceResponseBodyData) *PackSourceResponseBody {
	s.Data = v
	return s
}

func (s *PackSourceResponseBody) SetMessage(v string) *PackSourceResponseBody {
	s.Message = &v
	return s
}

func (s *PackSourceResponseBody) SetRequestId(v string) *PackSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PackSourceResponseBody) SetSuccess(v bool) *PackSourceResponseBody {
	s.Success = &v
	return s
}

type PackSourceResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PackSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PackSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PackSourceResponseBodyData) SetTaskId(v string) *PackSourceResponseBodyData {
	s.TaskId = &v
	return s
}

type PackSourceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PackSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PackSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s PackSourceResponse) GoString() string {
	return s.String()
}

func (s *PackSourceResponse) SetHeaders(v map[string]*string) *PackSourceResponse {
	s.Headers = v
	return s
}

func (s *PackSourceResponse) SetStatusCode(v int32) *PackSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PackSourceResponse) SetBody(v *PackSourceResponseBody) *PackSourceResponse {
	s.Body = v
	return s
}

type PredImageRequest struct {
	CorrectVertical *bool   `json:"CorrectVertical,omitempty" xml:"CorrectVertical,omitempty"`
	CountDetectDoor *int64  `json:"CountDetectDoor,omitempty" xml:"CountDetectDoor,omitempty"`
	DetectDoor      *bool   `json:"DetectDoor,omitempty" xml:"DetectDoor,omitempty"`
	SubSceneId      *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s PredImageRequest) String() string {
	return tea.Prettify(s)
}

func (s PredImageRequest) GoString() string {
	return s.String()
}

func (s *PredImageRequest) SetCorrectVertical(v bool) *PredImageRequest {
	s.CorrectVertical = &v
	return s
}

func (s *PredImageRequest) SetCountDetectDoor(v int64) *PredImageRequest {
	s.CountDetectDoor = &v
	return s
}

func (s *PredImageRequest) SetDetectDoor(v bool) *PredImageRequest {
	s.DetectDoor = &v
	return s
}

func (s *PredImageRequest) SetSubSceneId(v string) *PredImageRequest {
	s.SubSceneId = &v
	return s
}

type PredImageResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PredImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredImageResponseBody) GoString() string {
	return s.String()
}

func (s *PredImageResponseBody) SetCode(v int64) *PredImageResponseBody {
	s.Code = &v
	return s
}

func (s *PredImageResponseBody) SetMessage(v string) *PredImageResponseBody {
	s.Message = &v
	return s
}

func (s *PredImageResponseBody) SetRequestId(v string) *PredImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredImageResponseBody) SetSuccess(v bool) *PredImageResponseBody {
	s.Success = &v
	return s
}

func (s *PredImageResponseBody) SetTaskId(v string) *PredImageResponseBody {
	s.TaskId = &v
	return s
}

type PredImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PredImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredImageResponse) String() string {
	return tea.Prettify(s)
}

func (s PredImageResponse) GoString() string {
	return s.String()
}

func (s *PredImageResponse) SetHeaders(v map[string]*string) *PredImageResponse {
	s.Headers = v
	return s
}

func (s *PredImageResponse) SetStatusCode(v int32) *PredImageResponse {
	s.StatusCode = &v
	return s
}

func (s *PredImageResponse) SetBody(v *PredImageResponseBody) *PredImageResponse {
	s.Body = v
	return s
}

type PredictionWallLineRequest struct {
	CameraHeight *int64  `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PredictionWallLineRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictionWallLineRequest) GoString() string {
	return s.String()
}

func (s *PredictionWallLineRequest) SetCameraHeight(v int64) *PredictionWallLineRequest {
	s.CameraHeight = &v
	return s
}

func (s *PredictionWallLineRequest) SetUrl(v string) *PredictionWallLineRequest {
	s.Url = &v
	return s
}

type PredictionWallLineResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PredictionWallLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictionWallLineResponseBody) GoString() string {
	return s.String()
}

func (s *PredictionWallLineResponseBody) SetCode(v int64) *PredictionWallLineResponseBody {
	s.Code = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetMessage(v string) *PredictionWallLineResponseBody {
	s.Message = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetRequestId(v string) *PredictionWallLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetSubSceneId(v string) *PredictionWallLineResponseBody {
	s.SubSceneId = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetSuccess(v bool) *PredictionWallLineResponseBody {
	s.Success = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetTaskId(v string) *PredictionWallLineResponseBody {
	s.TaskId = &v
	return s
}

type PredictionWallLineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PredictionWallLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredictionWallLineResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictionWallLineResponse) GoString() string {
	return s.String()
}

func (s *PredictionWallLineResponse) SetHeaders(v map[string]*string) *PredictionWallLineResponse {
	s.Headers = v
	return s
}

func (s *PredictionWallLineResponse) SetStatusCode(v int32) *PredictionWallLineResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictionWallLineResponse) SetBody(v *PredictionWallLineResponseBody) *PredictionWallLineResponse {
	s.Body = v
	return s
}

type PublishHotspotRequest struct {
	ParamTag     *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s PublishHotspotRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotRequest) GoString() string {
	return s.String()
}

func (s *PublishHotspotRequest) SetParamTag(v string) *PublishHotspotRequest {
	s.ParamTag = &v
	return s
}

func (s *PublishHotspotRequest) SetSubSceneUuid(v string) *PublishHotspotRequest {
	s.SubSceneUuid = &v
	return s
}

type PublishHotspotResponseBody struct {
	Data       map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMessage *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishHotspotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotResponseBody) GoString() string {
	return s.String()
}

func (s *PublishHotspotResponseBody) SetData(v map[string]interface{}) *PublishHotspotResponseBody {
	s.Data = v
	return s
}

func (s *PublishHotspotResponseBody) SetErrMessage(v string) *PublishHotspotResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PublishHotspotResponseBody) SetRequestId(v string) *PublishHotspotResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishHotspotResponseBody) SetSuccess(v bool) *PublishHotspotResponseBody {
	s.Success = &v
	return s
}

type PublishHotspotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishHotspotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishHotspotResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotResponse) GoString() string {
	return s.String()
}

func (s *PublishHotspotResponse) SetHeaders(v map[string]*string) *PublishHotspotResponse {
	s.Headers = v
	return s
}

func (s *PublishHotspotResponse) SetStatusCode(v int32) *PublishHotspotResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishHotspotResponse) SetBody(v *PublishHotspotResponseBody) *PublishHotspotResponse {
	s.Body = v
	return s
}

type PublishHotspotConfigRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishHotspotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigRequest) SetSceneId(v string) *PublishHotspotConfigRequest {
	s.SceneId = &v
	return s
}

type PublishHotspotConfigResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishHotspotConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigResponseBody) SetCode(v int64) *PublishHotspotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetMessage(v string) *PublishHotspotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetRequestId(v string) *PublishHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetSuccess(v bool) *PublishHotspotConfigResponseBody {
	s.Success = &v
	return s
}

type PublishHotspotConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishHotspotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigResponse) SetHeaders(v map[string]*string) *PublishHotspotConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishHotspotConfigResponse) SetStatusCode(v int32) *PublishHotspotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishHotspotConfigResponse) SetBody(v *PublishHotspotConfigResponseBody) *PublishHotspotConfigResponse {
	s.Body = v
	return s
}

type PublishSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishSceneRequest) GoString() string {
	return s.String()
}

func (s *PublishSceneRequest) SetSceneId(v string) *PublishSceneRequest {
	s.SceneId = &v
	return s
}

type PublishSceneResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishSceneResponseBody) GoString() string {
	return s.String()
}

func (s *PublishSceneResponseBody) SetCode(v int64) *PublishSceneResponseBody {
	s.Code = &v
	return s
}

func (s *PublishSceneResponseBody) SetMessage(v string) *PublishSceneResponseBody {
	s.Message = &v
	return s
}

func (s *PublishSceneResponseBody) SetPreviewUrl(v string) *PublishSceneResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *PublishSceneResponseBody) SetRequestId(v string) *PublishSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishSceneResponseBody) SetSuccess(v bool) *PublishSceneResponseBody {
	s.Success = &v
	return s
}

type PublishSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishSceneResponse) GoString() string {
	return s.String()
}

func (s *PublishSceneResponse) SetHeaders(v map[string]*string) *PublishSceneResponse {
	s.Headers = v
	return s
}

func (s *PublishSceneResponse) SetStatusCode(v int32) *PublishSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishSceneResponse) SetBody(v *PublishSceneResponseBody) *PublishSceneResponse {
	s.Body = v
	return s
}

type PublishStatusRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishStatusRequest) GoString() string {
	return s.String()
}

func (s *PublishStatusRequest) SetSceneId(v string) *PublishStatusRequest {
	s.SceneId = &v
	return s
}

type PublishStatusResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s PublishStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PublishStatusResponseBody) SetCode(v int64) *PublishStatusResponseBody {
	s.Code = &v
	return s
}

func (s *PublishStatusResponseBody) SetMessage(v string) *PublishStatusResponseBody {
	s.Message = &v
	return s
}

func (s *PublishStatusResponseBody) SetRequestId(v string) *PublishStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishStatusResponseBody) SetStatus(v string) *PublishStatusResponseBody {
	s.Status = &v
	return s
}

func (s *PublishStatusResponseBody) SetSuccess(v bool) *PublishStatusResponseBody {
	s.Success = &v
	return s
}

func (s *PublishStatusResponseBody) SetSyncStatus(v string) *PublishStatusResponseBody {
	s.SyncStatus = &v
	return s
}

type PublishStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishStatusResponse) GoString() string {
	return s.String()
}

func (s *PublishStatusResponse) SetHeaders(v map[string]*string) *PublishStatusResponse {
	s.Headers = v
	return s
}

func (s *PublishStatusResponse) SetStatusCode(v int32) *PublishStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishStatusResponse) SetBody(v *PublishStatusResponseBody) *PublishStatusResponse {
	s.Body = v
	return s
}

type RecoveryOriginImageRequest struct {
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s RecoveryOriginImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoveryOriginImageRequest) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageRequest) SetSubSceneId(v string) *RecoveryOriginImageRequest {
	s.SubSceneId = &v
	return s
}

type RecoveryOriginImageResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecoveryOriginImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoveryOriginImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageResponseBody) SetCode(v int64) *RecoveryOriginImageResponseBody {
	s.Code = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetMessage(v string) *RecoveryOriginImageResponseBody {
	s.Message = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetRequestId(v string) *RecoveryOriginImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetSuccess(v bool) *RecoveryOriginImageResponseBody {
	s.Success = &v
	return s
}

type RecoveryOriginImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoveryOriginImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoveryOriginImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoveryOriginImageResponse) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageResponse) SetHeaders(v map[string]*string) *RecoveryOriginImageResponse {
	s.Headers = v
	return s
}

func (s *RecoveryOriginImageResponse) SetStatusCode(v int32) *RecoveryOriginImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryOriginImageResponse) SetBody(v *RecoveryOriginImageResponseBody) *RecoveryOriginImageResponse {
	s.Body = v
	return s
}

type RectVerticalRequest struct {
	CountDetectDoor *int32  `json:"CountDetectDoor,omitempty" xml:"CountDetectDoor,omitempty"`
	DetectDoor      *bool   `json:"DetectDoor,omitempty" xml:"DetectDoor,omitempty"`
	SubSceneId      *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	VerticalRect    *string `json:"VerticalRect,omitempty" xml:"VerticalRect,omitempty"`
}

func (s RectVerticalRequest) String() string {
	return tea.Prettify(s)
}

func (s RectVerticalRequest) GoString() string {
	return s.String()
}

func (s *RectVerticalRequest) SetCountDetectDoor(v int32) *RectVerticalRequest {
	s.CountDetectDoor = &v
	return s
}

func (s *RectVerticalRequest) SetDetectDoor(v bool) *RectVerticalRequest {
	s.DetectDoor = &v
	return s
}

func (s *RectVerticalRequest) SetSubSceneId(v string) *RectVerticalRequest {
	s.SubSceneId = &v
	return s
}

func (s *RectVerticalRequest) SetVerticalRect(v string) *RectVerticalRequest {
	s.VerticalRect = &v
	return s
}

type RectVerticalResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RectVerticalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RectVerticalResponseBody) GoString() string {
	return s.String()
}

func (s *RectVerticalResponseBody) SetCode(v int64) *RectVerticalResponseBody {
	s.Code = &v
	return s
}

func (s *RectVerticalResponseBody) SetMessage(v string) *RectVerticalResponseBody {
	s.Message = &v
	return s
}

func (s *RectVerticalResponseBody) SetRequestId(v string) *RectVerticalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectVerticalResponseBody) SetSuccess(v bool) *RectVerticalResponseBody {
	s.Success = &v
	return s
}

func (s *RectVerticalResponseBody) SetTaskId(v string) *RectVerticalResponseBody {
	s.TaskId = &v
	return s
}

type RectVerticalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RectVerticalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RectVerticalResponse) String() string {
	return tea.Prettify(s)
}

func (s RectVerticalResponse) GoString() string {
	return s.String()
}

func (s *RectVerticalResponse) SetHeaders(v map[string]*string) *RectVerticalResponse {
	s.Headers = v
	return s
}

func (s *RectVerticalResponse) SetStatusCode(v int32) *RectVerticalResponse {
	s.StatusCode = &v
	return s
}

func (s *RectVerticalResponse) SetBody(v *RectVerticalResponseBody) *RectVerticalResponse {
	s.Body = v
	return s
}

type RectifyImageRequest struct {
	CameraHeight *int64  `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RectifyImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RectifyImageRequest) GoString() string {
	return s.String()
}

func (s *RectifyImageRequest) SetCameraHeight(v int64) *RectifyImageRequest {
	s.CameraHeight = &v
	return s
}

func (s *RectifyImageRequest) SetUrl(v string) *RectifyImageRequest {
	s.Url = &v
	return s
}

type RectifyImageResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RectifyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RectifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *RectifyImageResponseBody) SetCode(v int64) *RectifyImageResponseBody {
	s.Code = &v
	return s
}

func (s *RectifyImageResponseBody) SetMessage(v string) *RectifyImageResponseBody {
	s.Message = &v
	return s
}

func (s *RectifyImageResponseBody) SetRequestId(v string) *RectifyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectifyImageResponseBody) SetSubSceneId(v string) *RectifyImageResponseBody {
	s.SubSceneId = &v
	return s
}

func (s *RectifyImageResponseBody) SetSuccess(v bool) *RectifyImageResponseBody {
	s.Success = &v
	return s
}

func (s *RectifyImageResponseBody) SetTaskId(v string) *RectifyImageResponseBody {
	s.TaskId = &v
	return s
}

type RectifyImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RectifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RectifyImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RectifyImageResponse) GoString() string {
	return s.String()
}

func (s *RectifyImageResponse) SetHeaders(v map[string]*string) *RectifyImageResponse {
	s.Headers = v
	return s
}

func (s *RectifyImageResponse) SetStatusCode(v int32) *RectifyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RectifyImageResponse) SetBody(v *RectifyImageResponseBody) *RectifyImageResponse {
	s.Body = v
	return s
}

type RollbackSubSceneRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RollbackSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackSubSceneRequest) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneRequest) SetId(v string) *RollbackSubSceneRequest {
	s.Id = &v
	return s
}

type RollbackSubSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneResponseBody) SetCode(v int64) *RollbackSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetMessage(v string) *RollbackSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetRequestId(v string) *RollbackSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetSuccess(v bool) *RollbackSubSceneResponseBody {
	s.Success = &v
	return s
}

type RollbackSubSceneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackSubSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackSubSceneResponse) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneResponse) SetHeaders(v map[string]*string) *RollbackSubSceneResponse {
	s.Headers = v
	return s
}

func (s *RollbackSubSceneResponse) SetStatusCode(v int32) *RollbackSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackSubSceneResponse) SetBody(v *RollbackSubSceneResponseBody) *RollbackSubSceneResponse {
	s.Body = v
	return s
}

type SaveHotspotConfigRequest struct {
	ParamTag     *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s SaveHotspotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigRequest) SetParamTag(v string) *SaveHotspotConfigRequest {
	s.ParamTag = &v
	return s
}

func (s *SaveHotspotConfigRequest) SetPreviewToken(v string) *SaveHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

type SaveHotspotConfigResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveHotspotConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponseBody) SetErrMessage(v string) *SaveHotspotConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetRequestId(v string) *SaveHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetSuccess(v bool) *SaveHotspotConfigResponseBody {
	s.Success = &v
	return s
}

type SaveHotspotConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveHotspotConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponse) SetHeaders(v map[string]*string) *SaveHotspotConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveHotspotConfigResponse) SetStatusCode(v int32) *SaveHotspotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveHotspotConfigResponse) SetBody(v *SaveHotspotConfigResponseBody) *SaveHotspotConfigResponse {
	s.Body = v
	return s
}

type SaveHotspotTagRequest struct {
	ParamTag     *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s SaveHotspotTagRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagRequest) SetParamTag(v string) *SaveHotspotTagRequest {
	s.ParamTag = &v
	return s
}

func (s *SaveHotspotTagRequest) SetSubSceneUuid(v string) *SaveHotspotTagRequest {
	s.SubSceneUuid = &v
	return s
}

type SaveHotspotTagResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveHotspotTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponseBody) SetErrMessage(v string) *SaveHotspotTagResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SaveHotspotTagResponseBody) SetRequestId(v string) *SaveHotspotTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotTagResponseBody) SetSuccess(v bool) *SaveHotspotTagResponseBody {
	s.Success = &v
	return s
}

type SaveHotspotTagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveHotspotTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveHotspotTagResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponse) SetHeaders(v map[string]*string) *SaveHotspotTagResponse {
	s.Headers = v
	return s
}

func (s *SaveHotspotTagResponse) SetStatusCode(v int32) *SaveHotspotTagResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveHotspotTagResponse) SetBody(v *SaveHotspotTagResponseBody) *SaveHotspotTagResponse {
	s.Body = v
	return s
}

type SaveHotspotTagListRequest struct {
	HotspotListJson *string `json:"HotspotListJson,omitempty" xml:"HotspotListJson,omitempty"`
	SceneId         *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SaveHotspotTagListRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagListRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListRequest) SetHotspotListJson(v string) *SaveHotspotTagListRequest {
	s.HotspotListJson = &v
	return s
}

func (s *SaveHotspotTagListRequest) SetSceneId(v string) *SaveHotspotTagListRequest {
	s.SceneId = &v
	return s
}

type SaveHotspotTagListResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveHotspotTagListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagListResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListResponseBody) SetCode(v int64) *SaveHotspotTagListResponseBody {
	s.Code = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetMessage(v string) *SaveHotspotTagListResponseBody {
	s.Message = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetRequestId(v string) *SaveHotspotTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetSuccess(v bool) *SaveHotspotTagListResponseBody {
	s.Success = &v
	return s
}

type SaveHotspotTagListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveHotspotTagListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveHotspotTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagListResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListResponse) SetHeaders(v map[string]*string) *SaveHotspotTagListResponse {
	s.Headers = v
	return s
}

func (s *SaveHotspotTagListResponse) SetStatusCode(v int32) *SaveHotspotTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveHotspotTagListResponse) SetBody(v *SaveHotspotTagListResponseBody) *SaveHotspotTagListResponse {
	s.Body = v
	return s
}

type SaveMinimapRequest struct {
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SaveMinimapRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveMinimapRequest) GoString() string {
	return s.String()
}

func (s *SaveMinimapRequest) SetData(v string) *SaveMinimapRequest {
	s.Data = &v
	return s
}

func (s *SaveMinimapRequest) SetSceneId(v string) *SaveMinimapRequest {
	s.SceneId = &v
	return s
}

type SaveMinimapResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveMinimapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveMinimapResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMinimapResponseBody) SetCode(v int64) *SaveMinimapResponseBody {
	s.Code = &v
	return s
}

func (s *SaveMinimapResponseBody) SetMessage(v string) *SaveMinimapResponseBody {
	s.Message = &v
	return s
}

func (s *SaveMinimapResponseBody) SetRequestId(v string) *SaveMinimapResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMinimapResponseBody) SetSuccess(v bool) *SaveMinimapResponseBody {
	s.Success = &v
	return s
}

type SaveMinimapResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveMinimapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveMinimapResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveMinimapResponse) GoString() string {
	return s.String()
}

func (s *SaveMinimapResponse) SetHeaders(v map[string]*string) *SaveMinimapResponse {
	s.Headers = v
	return s
}

func (s *SaveMinimapResponse) SetStatusCode(v int32) *SaveMinimapResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMinimapResponse) SetBody(v *SaveMinimapResponseBody) *SaveMinimapResponse {
	s.Body = v
	return s
}

type SaveModelConfigRequest struct {
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SaveModelConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveModelConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveModelConfigRequest) SetData(v string) *SaveModelConfigRequest {
	s.Data = &v
	return s
}

func (s *SaveModelConfigRequest) SetSceneId(v string) *SaveModelConfigRequest {
	s.SceneId = &v
	return s
}

type SaveModelConfigResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveModelConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveModelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveModelConfigResponseBody) SetCode(v int64) *SaveModelConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveModelConfigResponseBody) SetMessage(v string) *SaveModelConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveModelConfigResponseBody) SetRequestId(v string) *SaveModelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveModelConfigResponseBody) SetSuccess(v bool) *SaveModelConfigResponseBody {
	s.Success = &v
	return s
}

type SaveModelConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveModelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveModelConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveModelConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveModelConfigResponse) SetHeaders(v map[string]*string) *SaveModelConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveModelConfigResponse) SetStatusCode(v int32) *SaveModelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveModelConfigResponse) SetBody(v *SaveModelConfigResponseBody) *SaveModelConfigResponse {
	s.Body = v
	return s
}

type ScenePublishRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ScenePublishRequest) String() string {
	return tea.Prettify(s)
}

func (s ScenePublishRequest) GoString() string {
	return s.String()
}

func (s *ScenePublishRequest) SetSceneId(v string) *ScenePublishRequest {
	s.SceneId = &v
	return s
}

type ScenePublishResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScenePublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScenePublishResponseBody) GoString() string {
	return s.String()
}

func (s *ScenePublishResponseBody) SetCode(v int64) *ScenePublishResponseBody {
	s.Code = &v
	return s
}

func (s *ScenePublishResponseBody) SetMessage(v string) *ScenePublishResponseBody {
	s.Message = &v
	return s
}

func (s *ScenePublishResponseBody) SetPreviewUrl(v string) *ScenePublishResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *ScenePublishResponseBody) SetRequestId(v string) *ScenePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScenePublishResponseBody) SetSuccess(v bool) *ScenePublishResponseBody {
	s.Success = &v
	return s
}

type ScenePublishResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScenePublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScenePublishResponse) String() string {
	return tea.Prettify(s)
}

func (s ScenePublishResponse) GoString() string {
	return s.String()
}

func (s *ScenePublishResponse) SetHeaders(v map[string]*string) *ScenePublishResponse {
	s.Headers = v
	return s
}

func (s *ScenePublishResponse) SetStatusCode(v int32) *ScenePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *ScenePublishResponse) SetBody(v *ScenePublishResponseBody) *ScenePublishResponse {
	s.Body = v
	return s
}

type TempPreviewRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s TempPreviewRequest) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewRequest) GoString() string {
	return s.String()
}

func (s *TempPreviewRequest) SetSceneId(v string) *TempPreviewRequest {
	s.SceneId = &v
	return s
}

type TempPreviewResponseBody struct {
	Code       *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TempPreviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *TempPreviewResponseBody) SetCode(v int64) *TempPreviewResponseBody {
	s.Code = &v
	return s
}

func (s *TempPreviewResponseBody) SetMessage(v string) *TempPreviewResponseBody {
	s.Message = &v
	return s
}

func (s *TempPreviewResponseBody) SetPreviewUrl(v string) *TempPreviewResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *TempPreviewResponseBody) SetRequestId(v string) *TempPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempPreviewResponseBody) SetSceneId(v string) *TempPreviewResponseBody {
	s.SceneId = &v
	return s
}

func (s *TempPreviewResponseBody) SetSuccess(v bool) *TempPreviewResponseBody {
	s.Success = &v
	return s
}

type TempPreviewResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TempPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TempPreviewResponse) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewResponse) GoString() string {
	return s.String()
}

func (s *TempPreviewResponse) SetHeaders(v map[string]*string) *TempPreviewResponse {
	s.Headers = v
	return s
}

func (s *TempPreviewResponse) SetStatusCode(v int32) *TempPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *TempPreviewResponse) SetBody(v *TempPreviewResponseBody) *TempPreviewResponse {
	s.Body = v
	return s
}

type TempPreviewStatusRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s TempPreviewStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewStatusRequest) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusRequest) SetSceneId(v string) *TempPreviewStatusRequest {
	s.SceneId = &v
	return s
}

type TempPreviewStatusResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TempPreviewStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewStatusResponseBody) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusResponseBody) SetCode(v int64) *TempPreviewStatusResponseBody {
	s.Code = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetMessage(v string) *TempPreviewStatusResponseBody {
	s.Message = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetRequestId(v string) *TempPreviewStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetStatus(v string) *TempPreviewStatusResponseBody {
	s.Status = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetSuccess(v bool) *TempPreviewStatusResponseBody {
	s.Success = &v
	return s
}

type TempPreviewStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TempPreviewStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TempPreviewStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewStatusResponse) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusResponse) SetHeaders(v map[string]*string) *TempPreviewStatusResponse {
	s.Headers = v
	return s
}

func (s *TempPreviewStatusResponse) SetStatusCode(v int32) *TempPreviewStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *TempPreviewStatusResponse) SetBody(v *TempPreviewStatusResponseBody) *TempPreviewStatusResponse {
	s.Body = v
	return s
}

type UpdateConnDataRequest struct {
	ConnData *string `json:"ConnData,omitempty" xml:"ConnData,omitempty"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdateConnDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnDataRequest) SetConnData(v string) *UpdateConnDataRequest {
	s.ConnData = &v
	return s
}

func (s *UpdateConnDataRequest) SetSceneId(v string) *UpdateConnDataRequest {
	s.SceneId = &v
	return s
}

type UpdateConnDataResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConnDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnDataResponseBody) SetCode(v int64) *UpdateConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetMessage(v string) *UpdateConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetRequestId(v string) *UpdateConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetSuccess(v bool) *UpdateConnDataResponseBody {
	s.Success = &v
	return s
}

type UpdateConnDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConnDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnDataResponse) SetHeaders(v map[string]*string) *UpdateConnDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnDataResponse) SetStatusCode(v int32) *UpdateConnDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnDataResponse) SetBody(v *UpdateConnDataResponseBody) *UpdateConnDataResponse {
	s.Body = v
	return s
}

type UpdateLayoutDataRequest struct {
	LayoutData *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s UpdateLayoutDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataRequest) SetLayoutData(v string) *UpdateLayoutDataRequest {
	s.LayoutData = &v
	return s
}

func (s *UpdateLayoutDataRequest) SetSubSceneId(v string) *UpdateLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

type UpdateLayoutDataResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLayoutDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataResponseBody) SetCode(v int64) *UpdateLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetMessage(v string) *UpdateLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetRequestId(v string) *UpdateLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetSuccess(v bool) *UpdateLayoutDataResponseBody {
	s.Success = &v
	return s
}

type UpdateLayoutDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLayoutDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayoutDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataResponse) SetHeaders(v map[string]*string) *UpdateLayoutDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateLayoutDataResponse) SetStatusCode(v int32) *UpdateLayoutDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLayoutDataResponse) SetBody(v *UpdateLayoutDataResponseBody) *UpdateLayoutDataResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetBusinessId(v string) *UpdateProjectRequest {
	s.BusinessId = &v
	return s
}

func (s *UpdateProjectRequest) SetId(v string) *UpdateProjectRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

type UpdateProjectResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetCode(v int64) *UpdateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProjectResponseBody) SetMessage(v string) *UpdateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
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

type UpdateSceneRequest struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSceneRequest) SetId(v string) *UpdateSceneRequest {
	s.Id = &v
	return s
}

func (s *UpdateSceneRequest) SetName(v string) *UpdateSceneRequest {
	s.Name = &v
	return s
}

type UpdateSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponseBody) SetCode(v int64) *UpdateSceneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSceneResponseBody) SetMessage(v string) *UpdateSceneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSceneResponseBody) SetRequestId(v string) *UpdateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSceneResponseBody) SetSuccess(v bool) *UpdateSceneResponseBody {
	s.Success = &v
	return s
}

type UpdateSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponse) SetHeaders(v map[string]*string) *UpdateSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSceneResponse) SetStatusCode(v int32) *UpdateSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSceneResponse) SetBody(v *UpdateSceneResponseBody) *UpdateSceneResponse {
	s.Body = v
	return s
}

type UpdateSubSceneRequest struct {
	Id        *string    `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string    `json:"Name,omitempty" xml:"Name,omitempty"`
	ViewPoint []*float64 `json:"ViewPoint,omitempty" xml:"ViewPoint,omitempty" type:"Repeated"`
}

func (s UpdateSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneRequest) SetId(v string) *UpdateSubSceneRequest {
	s.Id = &v
	return s
}

func (s *UpdateSubSceneRequest) SetName(v string) *UpdateSubSceneRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubSceneRequest) SetViewPoint(v []*float64) *UpdateSubSceneRequest {
	s.ViewPoint = v
	return s
}

type UpdateSubSceneShrinkRequest struct {
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ViewPointShrink *string `json:"ViewPoint,omitempty" xml:"ViewPoint,omitempty"`
}

func (s UpdateSubSceneShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneShrinkRequest) SetId(v string) *UpdateSubSceneShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateSubSceneShrinkRequest) SetName(v string) *UpdateSubSceneShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubSceneShrinkRequest) SetViewPointShrink(v string) *UpdateSubSceneShrinkRequest {
	s.ViewPointShrink = &v
	return s
}

type UpdateSubSceneResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneResponseBody) SetCode(v int64) *UpdateSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetMessage(v string) *UpdateSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetRequestId(v string) *UpdateSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetSuccess(v bool) *UpdateSubSceneResponseBody {
	s.Success = &v
	return s
}

type UpdateSubSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneResponse) SetHeaders(v map[string]*string) *UpdateSubSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubSceneResponse) SetStatusCode(v int32) *UpdateSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubSceneResponse) SetBody(v *UpdateSubSceneResponseBody) *UpdateSubSceneResponse {
	s.Body = v
	return s
}

type UpdateSubSceneSeqRequest struct {
	SceneId         *string   `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SortSubSceneIds []*string `json:"SortSubSceneIds,omitempty" xml:"SortSubSceneIds,omitempty" type:"Repeated"`
}

func (s UpdateSubSceneSeqRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneSeqRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqRequest) SetSceneId(v string) *UpdateSubSceneSeqRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateSubSceneSeqRequest) SetSortSubSceneIds(v []*string) *UpdateSubSceneSeqRequest {
	s.SortSubSceneIds = v
	return s
}

type UpdateSubSceneSeqShrinkRequest struct {
	SceneId               *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SortSubSceneIdsShrink *string `json:"SortSubSceneIds,omitempty" xml:"SortSubSceneIds,omitempty"`
}

func (s UpdateSubSceneSeqShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneSeqShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqShrinkRequest) SetSceneId(v string) *UpdateSubSceneSeqShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateSubSceneSeqShrinkRequest) SetSortSubSceneIdsShrink(v string) *UpdateSubSceneSeqShrinkRequest {
	s.SortSubSceneIdsShrink = &v
	return s
}

type UpdateSubSceneSeqResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSubSceneSeqResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneSeqResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqResponseBody) SetCode(v int64) *UpdateSubSceneSeqResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetMessage(v string) *UpdateSubSceneSeqResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetRequestId(v string) *UpdateSubSceneSeqResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubSceneSeqResponseBody) SetSuccess(v bool) *UpdateSubSceneSeqResponseBody {
	s.Success = &v
	return s
}

type UpdateSubSceneSeqResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSubSceneSeqResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubSceneSeqResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneSeqResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqResponse) SetHeaders(v map[string]*string) *UpdateSubSceneSeqResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubSceneSeqResponse) SetStatusCode(v int32) *UpdateSubSceneSeqResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubSceneSeqResponse) SetBody(v *UpdateSubSceneSeqResponseBody) *UpdateSubSceneSeqResponse {
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
		"cn-hangzhou": tea.String("lyj.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tdsr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddHotspotFileWithOptions(request *AddHotspotFileRequest, runtime *util.RuntimeOptions) (_result *AddHotspotFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddHotspotFile"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddHotspotFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHotspotFile(request *AddHotspotFileRequest) (_result *AddHotspotFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddHotspotFileResponse{}
	_body, _err := client.AddHotspotFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMosaicsWithOptions(request *AddMosaicsRequest, runtime *util.RuntimeOptions) (_result *AddMosaicsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MarkPosition)) {
		query["MarkPosition"] = request.MarkPosition
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMosaics"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMosaicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMosaics(request *AddMosaicsRequest) (_result *AddMosaicsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMosaicsResponse{}
	_body, _err := client.AddMosaicsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProjectWithOptions(request *AddProjectRequest, runtime *util.RuntimeOptions) (_result *AddProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddProject"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProject(request *AddProjectRequest) (_result *AddProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddProjectResponse{}
	_body, _err := client.AddProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRelativePositionWithOptions(request *AddRelativePositionRequest, runtime *util.RuntimeOptions) (_result *AddRelativePositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelativePosition)) {
		query["RelativePosition"] = request.RelativePosition
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRelativePosition"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRelativePositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRelativePosition(request *AddRelativePositionRequest) (_result *AddRelativePositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRelativePositionResponse{}
	_body, _err := client.AddRelativePositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRoomPlanWithOptions(request *AddRoomPlanRequest, runtime *util.RuntimeOptions) (_result *AddRoomPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRoomPlan"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRoomPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRoomPlan(request *AddRoomPlanRequest) (_result *AddRoomPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRoomPlanResponse{}
	_body, _err := client.AddRoomPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSceneWithOptions(request *AddSceneRequest, runtime *util.RuntimeOptions) (_result *AddSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerUid)) {
		query["CustomerUid"] = request.CustomerUid
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddScene(request *AddSceneRequest) (_result *AddSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSceneResponse{}
	_body, _err := client.AddSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSubSceneWithOptions(request *AddSubSceneRequest, runtime *util.RuntimeOptions) (_result *AddSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["UploadType"] = request.UploadType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSubScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSubSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSubScene(request *AddSubSceneRequest) (_result *AddSubSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSubSceneResponse{}
	_body, _err := client.AddSubSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserPropertyWithOptions(request *CheckUserPropertyRequest, runtime *util.RuntimeOptions) (_result *CheckUserPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUserProperty"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserProperty(request *CheckUserPropertyRequest) (_result *CheckUserPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserPropertyResponse{}
	_body, _err := client.CheckUserPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopySceneWithOptions(request *CopySceneRequest, runtime *util.RuntimeOptions) (_result *CopySceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopySceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyScene(request *CopySceneRequest) (_result *CopySceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopySceneResponse{}
	_body, _err := client.CopySceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadPolicyWithOptions(request *CreateUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUploadPolicy"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadPolicy(request *CreateUploadPolicyRequest) (_result *CreateUploadPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadPolicyResponse{}
	_body, _err := client.CreateUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetailProjectWithOptions(request *DetailProjectRequest, runtime *util.RuntimeOptions) (_result *DetailProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetailProject"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetailProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetailProject(request *DetailProjectRequest) (_result *DetailProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetailProjectResponse{}
	_body, _err := client.DetailProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetailSceneWithOptions(request *DetailSceneRequest, runtime *util.RuntimeOptions) (_result *DetailSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetailScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetailSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetailScene(request *DetailSceneRequest) (_result *DetailSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetailSceneResponse{}
	_body, _err := client.DetailSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetailSubSceneWithOptions(request *DetailSubSceneRequest, runtime *util.RuntimeOptions) (_result *DetailSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetailSubScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetailSubSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetailSubScene(request *DetailSubSceneRequest) (_result *DetailSubSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetailSubSceneResponse{}
	_body, _err := client.DetailSubSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DropProjectWithOptions(request *DropProjectRequest, runtime *util.RuntimeOptions) (_result *DropProjectResponse, _err error) {
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
		Action:      tea.String("DropProject"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DropProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DropProject(request *DropProjectRequest) (_result *DropProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DropProjectResponse{}
	_body, _err := client.DropProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DropSceneWithOptions(request *DropSceneRequest, runtime *util.RuntimeOptions) (_result *DropSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DropScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DropSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DropScene(request *DropSceneRequest) (_result *DropSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DropSceneResponse{}
	_body, _err := client.DropSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DropSubSceneWithOptions(request *DropSubSceneRequest, runtime *util.RuntimeOptions) (_result *DropSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DropSubScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DropSubSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DropSubScene(request *DropSubSceneRequest) (_result *DropSubSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DropSubSceneResponse{}
	_body, _err := client.DropSubSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnDataWithOptions(request *GetConnDataRequest, runtime *util.RuntimeOptions) (_result *GetConnDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnData(request *GetConnDataRequest) (_result *GetConnDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnDataResponse{}
	_body, _err := client.GetConnDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCopySceneTaskStatusWithOptions(request *GetCopySceneTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetCopySceneTaskStatusResponse, _err error) {
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
		Action:      tea.String("GetCopySceneTaskStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCopySceneTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCopySceneTaskStatus(request *GetCopySceneTaskStatusRequest) (_result *GetCopySceneTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCopySceneTaskStatusResponse{}
	_body, _err := client.GetCopySceneTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotspotConfigWithOptions(request *GetHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *GetHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotConfig"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotspotConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotspotConfig(request *GetHotspotConfigRequest) (_result *GetHotspotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotspotConfigResponse{}
	_body, _err := client.GetHotspotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotspotSceneDataWithOptions(request *GetHotspotSceneDataRequest, runtime *util.RuntimeOptions) (_result *GetHotspotSceneDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotSceneData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotspotSceneDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotspotSceneData(request *GetHotspotSceneDataRequest) (_result *GetHotspotSceneDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotspotSceneDataResponse{}
	_body, _err := client.GetHotspotSceneDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotspotTagWithOptions(request *GetHotspotTagRequest, runtime *util.RuntimeOptions) (_result *GetHotspotTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneUuid)) {
		query["SubSceneUuid"] = request.SubSceneUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotTag"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotspotTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotspotTag(request *GetHotspotTagRequest) (_result *GetHotspotTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotspotTagResponse{}
	_body, _err := client.GetHotspotTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLayoutDataWithOptions(request *GetLayoutDataRequest, runtime *util.RuntimeOptions) (_result *GetLayoutDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayoutData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLayoutDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLayoutData(request *GetLayoutDataRequest) (_result *GetLayoutDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLayoutDataResponse{}
	_body, _err := client.GetLayoutDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOriginLayoutDataWithOptions(request *GetOriginLayoutDataRequest, runtime *util.RuntimeOptions) (_result *GetOriginLayoutDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOriginLayoutData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOriginLayoutDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOriginLayoutData(request *GetOriginLayoutDataRequest) (_result *GetOriginLayoutDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOriginLayoutDataResponse{}
	_body, _err := client.GetOriginLayoutDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssPolicyWithOptions(request *GetOssPolicyRequest, runtime *util.RuntimeOptions) (_result *GetOssPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssPolicy"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssPolicy(request *GetOssPolicyRequest) (_result *GetOssPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssPolicyResponse{}
	_body, _err := client.GetOssPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPackSceneTaskStatusWithOptions(request *GetPackSceneTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetPackSceneTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPackSceneTaskStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPackSceneTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPackSceneTaskStatus(request *GetPackSceneTaskStatusRequest) (_result *GetPackSceneTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPackSceneTaskStatusResponse{}
	_body, _err := client.GetPackSceneTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRectifyImageWithOptions(request *GetRectifyImageRequest, runtime *util.RuntimeOptions) (_result *GetRectifyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRectifyImage"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRectifyImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRectifyImage(request *GetRectifyImageRequest) (_result *GetRectifyImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRectifyImageResponse{}
	_body, _err := client.GetRectifyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSceneBuildTaskStatusWithOptions(request *GetSceneBuildTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetSceneBuildTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSceneBuildTaskStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSceneBuildTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSceneBuildTaskStatus(request *GetSceneBuildTaskStatusRequest) (_result *GetSceneBuildTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSceneBuildTaskStatusResponse{}
	_body, _err := client.GetSceneBuildTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScenePackUrlWithOptions(request *GetScenePackUrlRequest, runtime *util.RuntimeOptions) (_result *GetScenePackUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScenePackUrl"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScenePackUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScenePackUrl(request *GetScenePackUrlRequest) (_result *GetScenePackUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScenePackUrlResponse{}
	_body, _err := client.GetScenePackUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScenePreviewDataWithOptions(request *GetScenePreviewDataRequest, runtime *util.RuntimeOptions) (_result *GetScenePreviewDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	if !tea.BoolValue(util.IsUnset(request.ShowTag)) {
		query["ShowTag"] = request.ShowTag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScenePreviewData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScenePreviewDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScenePreviewData(request *GetScenePreviewDataRequest) (_result *GetScenePreviewDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScenePreviewDataResponse{}
	_body, _err := client.GetScenePreviewDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScenePreviewInfoWithOptions(request *GetScenePreviewInfoRequest, runtime *util.RuntimeOptions) (_result *GetScenePreviewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.ModelToken)) {
		query["ModelToken"] = request.ModelToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScenePreviewInfo"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScenePreviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScenePreviewInfo(request *GetScenePreviewInfoRequest) (_result *GetScenePreviewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScenePreviewInfoResponse{}
	_body, _err := client.GetScenePreviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScenePreviewResourceWithOptions(request *GetScenePreviewResourceRequest, runtime *util.RuntimeOptions) (_result *GetScenePreviewResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Draft)) {
		query["Draft"] = request.Draft
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScenePreviewResource"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScenePreviewResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScenePreviewResource(request *GetScenePreviewResourceRequest) (_result *GetScenePreviewResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScenePreviewResourceResponse{}
	_body, _err := client.GetScenePreviewResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSingleConnDataWithOptions(request *GetSingleConnDataRequest, runtime *util.RuntimeOptions) (_result *GetSingleConnDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSingleConnData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSingleConnDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSingleConnData(request *GetSingleConnDataRequest) (_result *GetSingleConnDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSingleConnDataResponse{}
	_body, _err := client.GetSingleConnDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSourcePackStatusWithOptions(request *GetSourcePackStatusRequest, runtime *util.RuntimeOptions) (_result *GetSourcePackStatusResponse, _err error) {
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
		Action:      tea.String("GetSourcePackStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSourcePackStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSourcePackStatus(request *GetSourcePackStatusRequest) (_result *GetSourcePackStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSourcePackStatusResponse{}
	_body, _err := client.GetSourcePackStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubSceneTaskStatusWithOptions(request *GetSubSceneTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetSubSceneTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubSceneTaskStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubSceneTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubSceneTaskStatus(request *GetSubSceneTaskStatusRequest) (_result *GetSubSceneTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubSceneTaskStatusResponse{}
	_body, _err := client.GetSubSceneTaskStatusWithOptions(request, runtime)
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
		Version:     tea.String("2020-01-01"),
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

func (client *Client) GetWindowConfigWithOptions(request *GetWindowConfigRequest, runtime *util.RuntimeOptions) (_result *GetWindowConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWindowConfig"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWindowConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWindowConfig(request *GetWindowConfigRequest) (_result *GetWindowConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWindowConfigResponse{}
	_body, _err := client.GetWindowConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LabelBuildWithOptions(request *LabelBuildRequest, runtime *util.RuntimeOptions) (_result *LabelBuildResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ModelStyle)) {
		query["ModelStyle"] = request.ModelStyle
	}

	if !tea.BoolValue(util.IsUnset(request.OptimizeWallWidth)) {
		query["OptimizeWallWidth"] = request.OptimizeWallWidth
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStyle)) {
		query["PlanStyle"] = request.PlanStyle
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.WallHeight)) {
		query["WallHeight"] = request.WallHeight
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LabelBuild"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LabelBuildResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LabelBuild(request *LabelBuildRequest) (_result *LabelBuildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LabelBuildResponse{}
	_body, _err := client.LabelBuildWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LinkImageWithOptions(request *LinkImageRequest, runtime *util.RuntimeOptions) (_result *LinkImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CameraHeight)) {
		query["CameraHeight"] = request.CameraHeight
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LinkImage"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LinkImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LinkImage(request *LinkImageRequest) (_result *LinkImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LinkImageResponse{}
	_body, _err := client.LinkImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectWithOptions(request *ListProjectRequest, runtime *util.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProject"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSceneWithOptions(request *ListSceneRequest, runtime *util.RuntimeOptions) (_result *ListSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScene(request *ListSceneRequest) (_result *ListSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSceneResponse{}
	_body, _err := client.ListSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubSceneWithOptions(request *ListSubSceneRequest, runtime *util.RuntimeOptions) (_result *ListSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowLayoutData)) {
		query["ShowLayoutData"] = request.ShowLayoutData
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubScene(request *ListSubSceneRequest) (_result *ListSubSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubSceneResponse{}
	_body, _err := client.ListSubSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OptimizeRightAngleWithOptions(request *OptimizeRightAngleRequest, runtime *util.RuntimeOptions) (_result *OptimizeRightAngleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OptimizeRightAngle"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OptimizeRightAngleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OptimizeRightAngle(request *OptimizeRightAngleRequest) (_result *OptimizeRightAngleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OptimizeRightAngleResponse{}
	_body, _err := client.OptimizeRightAngleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PackSceneWithOptions(request *PackSceneRequest, runtime *util.RuntimeOptions) (_result *PackSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PackScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PackSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PackScene(request *PackSceneRequest) (_result *PackSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PackSceneResponse{}
	_body, _err := client.PackSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PackSourceWithOptions(request *PackSourceRequest, runtime *util.RuntimeOptions) (_result *PackSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PackSource"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PackSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PackSource(request *PackSourceRequest) (_result *PackSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PackSourceResponse{}
	_body, _err := client.PackSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PredImageWithOptions(request *PredImageRequest, runtime *util.RuntimeOptions) (_result *PredImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorrectVertical)) {
		query["CorrectVertical"] = request.CorrectVertical
	}

	if !tea.BoolValue(util.IsUnset(request.CountDetectDoor)) {
		query["CountDetectDoor"] = request.CountDetectDoor
	}

	if !tea.BoolValue(util.IsUnset(request.DetectDoor)) {
		query["DetectDoor"] = request.DetectDoor
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PredImage"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PredImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredImage(request *PredImageRequest) (_result *PredImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredImageResponse{}
	_body, _err := client.PredImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PredictionWallLineWithOptions(request *PredictionWallLineRequest, runtime *util.RuntimeOptions) (_result *PredictionWallLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CameraHeight)) {
		query["CameraHeight"] = request.CameraHeight
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictionWallLine"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictionWallLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictionWallLine(request *PredictionWallLineRequest) (_result *PredictionWallLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictionWallLineResponse{}
	_body, _err := client.PredictionWallLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishHotspotWithOptions(request *PublishHotspotRequest, runtime *util.RuntimeOptions) (_result *PublishHotspotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamTag)) {
		query["ParamTag"] = request.ParamTag
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneUuid)) {
		query["SubSceneUuid"] = request.SubSceneUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishHotspot"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishHotspotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishHotspot(request *PublishHotspotRequest) (_result *PublishHotspotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishHotspotResponse{}
	_body, _err := client.PublishHotspotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishHotspotConfigWithOptions(request *PublishHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *PublishHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishHotspotConfig"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishHotspotConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishHotspotConfig(request *PublishHotspotConfigRequest) (_result *PublishHotspotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishHotspotConfigResponse{}
	_body, _err := client.PublishHotspotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishSceneWithOptions(request *PublishSceneRequest, runtime *util.RuntimeOptions) (_result *PublishSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishScene(request *PublishSceneRequest) (_result *PublishSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishSceneResponse{}
	_body, _err := client.PublishSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishStatusWithOptions(request *PublishStatusRequest, runtime *util.RuntimeOptions) (_result *PublishStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishStatus(request *PublishStatusRequest) (_result *PublishStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishStatusResponse{}
	_body, _err := client.PublishStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoveryOriginImageWithOptions(request *RecoveryOriginImageRequest, runtime *util.RuntimeOptions) (_result *RecoveryOriginImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoveryOriginImage"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoveryOriginImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoveryOriginImage(request *RecoveryOriginImageRequest) (_result *RecoveryOriginImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoveryOriginImageResponse{}
	_body, _err := client.RecoveryOriginImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RectVerticalWithOptions(request *RectVerticalRequest, runtime *util.RuntimeOptions) (_result *RectVerticalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CountDetectDoor)) {
		query["CountDetectDoor"] = request.CountDetectDoor
	}

	if !tea.BoolValue(util.IsUnset(request.DetectDoor)) {
		query["DetectDoor"] = request.DetectDoor
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	if !tea.BoolValue(util.IsUnset(request.VerticalRect)) {
		query["VerticalRect"] = request.VerticalRect
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RectVertical"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RectVerticalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RectVertical(request *RectVerticalRequest) (_result *RectVerticalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RectVerticalResponse{}
	_body, _err := client.RectVerticalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RectifyImageWithOptions(request *RectifyImageRequest, runtime *util.RuntimeOptions) (_result *RectifyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CameraHeight)) {
		query["CameraHeight"] = request.CameraHeight
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RectifyImage"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RectifyImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RectifyImage(request *RectifyImageRequest) (_result *RectifyImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RectifyImageResponse{}
	_body, _err := client.RectifyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackSubSceneWithOptions(request *RollbackSubSceneRequest, runtime *util.RuntimeOptions) (_result *RollbackSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackSubScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackSubSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackSubScene(request *RollbackSubSceneRequest) (_result *RollbackSubSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackSubSceneResponse{}
	_body, _err := client.RollbackSubSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveHotspotConfigWithOptions(request *SaveHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamTag)) {
		query["ParamTag"] = request.ParamTag
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewToken)) {
		query["PreviewToken"] = request.PreviewToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveHotspotConfig"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveHotspotConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveHotspotConfig(request *SaveHotspotConfigRequest) (_result *SaveHotspotConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveHotspotConfigResponse{}
	_body, _err := client.SaveHotspotConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveHotspotTagWithOptions(request *SaveHotspotTagRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamTag)) {
		query["ParamTag"] = request.ParamTag
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneUuid)) {
		query["SubSceneUuid"] = request.SubSceneUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveHotspotTag"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveHotspotTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveHotspotTag(request *SaveHotspotTagRequest) (_result *SaveHotspotTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveHotspotTagResponse{}
	_body, _err := client.SaveHotspotTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveHotspotTagListWithOptions(request *SaveHotspotTagListRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotTagListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotspotListJson)) {
		query["HotspotListJson"] = request.HotspotListJson
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveHotspotTagList"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveHotspotTagListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveHotspotTagList(request *SaveHotspotTagListRequest) (_result *SaveHotspotTagListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveHotspotTagListResponse{}
	_body, _err := client.SaveHotspotTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveMinimapWithOptions(request *SaveMinimapRequest, runtime *util.RuntimeOptions) (_result *SaveMinimapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveMinimap"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveMinimapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveMinimap(request *SaveMinimapRequest) (_result *SaveMinimapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveMinimapResponse{}
	_body, _err := client.SaveMinimapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveModelConfigWithOptions(request *SaveModelConfigRequest, runtime *util.RuntimeOptions) (_result *SaveModelConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveModelConfig"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveModelConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveModelConfig(request *SaveModelConfigRequest) (_result *SaveModelConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveModelConfigResponse{}
	_body, _err := client.SaveModelConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScenePublishWithOptions(request *ScenePublishRequest, runtime *util.RuntimeOptions) (_result *ScenePublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScenePublish"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScenePublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScenePublish(request *ScenePublishRequest) (_result *ScenePublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScenePublishResponse{}
	_body, _err := client.ScenePublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TempPreviewWithOptions(request *TempPreviewRequest, runtime *util.RuntimeOptions) (_result *TempPreviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TempPreview"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TempPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TempPreview(request *TempPreviewRequest) (_result *TempPreviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TempPreviewResponse{}
	_body, _err := client.TempPreviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TempPreviewStatusWithOptions(request *TempPreviewStatusRequest, runtime *util.RuntimeOptions) (_result *TempPreviewStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TempPreviewStatus"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TempPreviewStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TempPreviewStatus(request *TempPreviewStatusRequest) (_result *TempPreviewStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TempPreviewStatusResponse{}
	_body, _err := client.TempPreviewStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConnDataWithOptions(request *UpdateConnDataRequest, runtime *util.RuntimeOptions) (_result *UpdateConnDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnData)) {
		query["ConnData"] = request.ConnData
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConnData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConnDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConnData(request *UpdateConnDataRequest) (_result *UpdateConnDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConnDataResponse{}
	_body, _err := client.UpdateConnDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLayoutDataWithOptions(request *UpdateLayoutDataRequest, runtime *util.RuntimeOptions) (_result *UpdateLayoutDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LayoutData)) {
		query["LayoutData"] = request.LayoutData
	}

	if !tea.BoolValue(util.IsUnset(request.SubSceneId)) {
		query["SubSceneId"] = request.SubSceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLayoutData"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLayoutDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLayoutData(request *UpdateLayoutDataRequest) (_result *UpdateLayoutDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLayoutDataResponse{}
	_body, _err := client.UpdateLayoutDataWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.BusinessId)) {
		query["BusinessId"] = request.BusinessId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2020-01-01"),
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

func (client *Client) UpdateSceneWithOptions(request *UpdateSceneRequest, runtime *util.RuntimeOptions) (_result *UpdateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateScene(request *UpdateSceneRequest) (_result *UpdateSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSceneResponse{}
	_body, _err := client.UpdateSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubSceneWithOptions(tmpReq *UpdateSubSceneRequest, runtime *util.RuntimeOptions) (_result *UpdateSubSceneResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSubSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ViewPoint)) {
		request.ViewPointShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewPoint, tea.String("ViewPoint"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ViewPointShrink)) {
		query["ViewPoint"] = request.ViewPointShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubScene"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubScene(request *UpdateSubSceneRequest) (_result *UpdateSubSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubSceneResponse{}
	_body, _err := client.UpdateSubSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubSceneSeqWithOptions(tmpReq *UpdateSubSceneSeqRequest, runtime *util.RuntimeOptions) (_result *UpdateSubSceneSeqResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSubSceneSeqShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SortSubSceneIds)) {
		request.SortSubSceneIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortSubSceneIds, tea.String("SortSubSceneIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.SortSubSceneIdsShrink)) {
		query["SortSubSceneIds"] = request.SortSubSceneIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubSceneSeq"),
		Version:     tea.String("2020-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubSceneSeqResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubSceneSeq(request *UpdateSubSceneSeqRequest) (_result *UpdateSubSceneSeqResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubSceneSeqResponse{}
	_body, _err := client.UpdateSubSceneSeqWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
