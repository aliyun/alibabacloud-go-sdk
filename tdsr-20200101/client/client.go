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

type GetSingleConnDataRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// 关联信息
	List []*GetSingleConnDataResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s GetSingleConnDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSingleConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponseBody) SetRequestId(v string) *GetSingleConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetCode(v int64) *GetSingleConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetSuccess(v bool) *GetSingleConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetMessage(v string) *GetSingleConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetVersion(v string) *GetSingleConnDataResponseBody {
	s.Version = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetList(v []*GetSingleConnDataResponseBodyList) *GetSingleConnDataResponseBody {
	s.List = v
	return s
}

type GetSingleConnDataResponseBodyList struct {
	// ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 关联ID
	MapId *string `json:"MapId,omitempty" xml:"MapId,omitempty"`
	// outer:外关联 inner：内关联 stair：楼梯关联
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSingleConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSingleConnDataResponse) SetBody(v *GetSingleConnDataResponseBody) *GetSingleConnDataResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	// 任务ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 未开始 :init 处理中 : processing    失败 :failure   完成 :succeed  取消 :canceled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 墙线预测: wall_line 切图: cut_image   重建: build  直角优化：right_angle_optimization  其他：other
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 任务执行失败错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 任务执行失败错误消息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetCode(v int64) *GetTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetStatus(v string) *GetTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetType(v string) *GetTaskStatusResponseBody {
	s.Type = &v
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

type LinkImageRequest struct {
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// 图片或者视频名称xxx.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 相机高度 单位 cm
	CameraHeight *int32 `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
	// 平台标识，默认PC
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s LinkImageRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkImageRequest) GoString() string {
	return s.String()
}

func (s *LinkImageRequest) SetSubSceneId(v string) *LinkImageRequest {
	s.SubSceneId = &v
	return s
}

func (s *LinkImageRequest) SetFileName(v string) *LinkImageRequest {
	s.FileName = &v
	return s
}

func (s *LinkImageRequest) SetCameraHeight(v int32) *LinkImageRequest {
	s.CameraHeight = &v
	return s
}

func (s *LinkImageRequest) SetPlatform(v string) *LinkImageRequest {
	s.Platform = &v
	return s
}

type LinkImageResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 图片/视频ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s LinkImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LinkImageResponseBody) GoString() string {
	return s.String()
}

func (s *LinkImageResponseBody) SetRequestId(v string) *LinkImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *LinkImageResponseBody) SetCode(v int64) *LinkImageResponseBody {
	s.Code = &v
	return s
}

func (s *LinkImageResponseBody) SetSuccess(v bool) *LinkImageResponseBody {
	s.Success = &v
	return s
}

func (s *LinkImageResponseBody) SetMessage(v string) *LinkImageResponseBody {
	s.Message = &v
	return s
}

func (s *LinkImageResponseBody) SetResourceId(v string) *LinkImageResponseBody {
	s.ResourceId = &v
	return s
}

type LinkImageResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LinkImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *LinkImageResponse) SetBody(v *LinkImageResponseBody) *LinkImageResponse {
	s.Body = v
	return s
}

type AddSceneRequest struct {
	// 场景类型 3D模型：MODEL_3D  全景图片：PIC  全景视频：VIDEO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSceneRequest) GoString() string {
	return s.String()
}

func (s *AddSceneRequest) SetType(v string) *AddSceneRequest {
	s.Type = &v
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

type AddSceneResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 场景ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSceneResponseBody) GoString() string {
	return s.String()
}

func (s *AddSceneResponseBody) SetRequestId(v string) *AddSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSceneResponseBody) SetCode(v int64) *AddSceneResponseBody {
	s.Code = &v
	return s
}

func (s *AddSceneResponseBody) SetSuccess(v bool) *AddSceneResponseBody {
	s.Success = &v
	return s
}

func (s *AddSceneResponseBody) SetMessage(v string) *AddSceneResponseBody {
	s.Message = &v
	return s
}

func (s *AddSceneResponseBody) SetId(v string) *AddSceneResponseBody {
	s.Id = &v
	return s
}

type AddSceneResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddSceneResponse) SetBody(v *AddSceneResponseBody) *AddSceneResponse {
	s.Body = v
	return s
}

type UpdateConnDataRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 关联数据
	ConnData *string `json:"ConnData,omitempty" xml:"ConnData,omitempty"`
}

func (s UpdateConnDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnDataRequest) SetSceneId(v string) *UpdateConnDataRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateConnDataRequest) SetConnData(v string) *UpdateConnDataRequest {
	s.ConnData = &v
	return s
}

type UpdateConnDataResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateConnDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnDataResponseBody) SetRequestId(v string) *UpdateConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetCode(v int64) *UpdateConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetSuccess(v bool) *UpdateConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConnDataResponseBody) SetMessage(v string) *UpdateConnDataResponseBody {
	s.Message = &v
	return s
}

type UpdateConnDataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateConnDataResponse) SetBody(v *UpdateConnDataResponseBody) *UpdateConnDataResponse {
	s.Body = v
	return s
}

type RectifyImageRequest struct {
	// 图片地址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 相机高度 单位 cm
	CameraHeight *int64 `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
}

func (s RectifyImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RectifyImageRequest) GoString() string {
	return s.String()
}

func (s *RectifyImageRequest) SetUrl(v string) *RectifyImageRequest {
	s.Url = &v
	return s
}

func (s *RectifyImageRequest) SetCameraHeight(v int64) *RectifyImageRequest {
	s.CameraHeight = &v
	return s
}

type RectifyImageResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s RectifyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RectifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *RectifyImageResponseBody) SetRequestId(v string) *RectifyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectifyImageResponseBody) SetCode(v int64) *RectifyImageResponseBody {
	s.Code = &v
	return s
}

func (s *RectifyImageResponseBody) SetSuccess(v bool) *RectifyImageResponseBody {
	s.Success = &v
	return s
}

func (s *RectifyImageResponseBody) SetMessage(v string) *RectifyImageResponseBody {
	s.Message = &v
	return s
}

func (s *RectifyImageResponseBody) SetTaskId(v string) *RectifyImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *RectifyImageResponseBody) SetSubSceneId(v string) *RectifyImageResponseBody {
	s.SubSceneId = &v
	return s
}

type RectifyImageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RectifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RectifyImageResponse) SetBody(v *RectifyImageResponseBody) *RectifyImageResponse {
	s.Body = v
	return s
}

type LabelBuildRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s LabelBuildRequest) String() string {
	return tea.Prettify(s)
}

func (s LabelBuildRequest) GoString() string {
	return s.String()
}

func (s *LabelBuildRequest) SetSceneId(v string) *LabelBuildRequest {
	s.SceneId = &v
	return s
}

func (s *LabelBuildRequest) SetMode(v string) *LabelBuildRequest {
	s.Mode = &v
	return s
}

type LabelBuildResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 重建任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s LabelBuildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LabelBuildResponseBody) GoString() string {
	return s.String()
}

func (s *LabelBuildResponseBody) SetRequestId(v string) *LabelBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *LabelBuildResponseBody) SetCode(v int64) *LabelBuildResponseBody {
	s.Code = &v
	return s
}

func (s *LabelBuildResponseBody) SetSuccess(v bool) *LabelBuildResponseBody {
	s.Success = &v
	return s
}

func (s *LabelBuildResponseBody) SetMessage(v string) *LabelBuildResponseBody {
	s.Message = &v
	return s
}

func (s *LabelBuildResponseBody) SetTaskId(v string) *LabelBuildResponseBody {
	s.TaskId = &v
	return s
}

type LabelBuildResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LabelBuildResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *LabelBuildResponse) SetBody(v *LabelBuildResponseBody) *LabelBuildResponse {
	s.Body = v
	return s
}

type DropSceneRequest struct {
	// 主场景id
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DropSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DropSceneResponseBody) SetRequestId(v string) *DropSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSceneResponseBody) SetCode(v int64) *DropSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DropSceneResponseBody) SetSuccess(v bool) *DropSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DropSceneResponseBody) SetMessage(v string) *DropSceneResponseBody {
	s.Message = &v
	return s
}

type DropSceneResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DropSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DropSceneResponse) SetBody(v *DropSceneResponseBody) *DropSceneResponse {
	s.Body = v
	return s
}

type OptimizeRightAngleRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OptimizeRightAngleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OptimizeRightAngleResponseBody) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleResponseBody) SetRequestId(v string) *OptimizeRightAngleResponseBody {
	s.RequestId = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetCode(v int64) *OptimizeRightAngleResponseBody {
	s.Code = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetSuccess(v bool) *OptimizeRightAngleResponseBody {
	s.Success = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetMessage(v string) *OptimizeRightAngleResponseBody {
	s.Message = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetTaskId(v string) *OptimizeRightAngleResponseBody {
	s.TaskId = &v
	return s
}

type OptimizeRightAngleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OptimizeRightAngleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OptimizeRightAngleResponse) SetBody(v *OptimizeRightAngleResponseBody) *OptimizeRightAngleResponse {
	s.Body = v
	return s
}

type AddRelativePositionRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 相对位置信息
	RelativePosition *string `json:"RelativePosition,omitempty" xml:"RelativePosition,omitempty"`
}

func (s AddRelativePositionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRelativePositionRequest) GoString() string {
	return s.String()
}

func (s *AddRelativePositionRequest) SetSceneId(v string) *AddRelativePositionRequest {
	s.SceneId = &v
	return s
}

func (s *AddRelativePositionRequest) SetRelativePosition(v string) *AddRelativePositionRequest {
	s.RelativePosition = &v
	return s
}

type AddRelativePositionResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddRelativePositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRelativePositionResponseBody) GoString() string {
	return s.String()
}

func (s *AddRelativePositionResponseBody) SetRequestId(v string) *AddRelativePositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetCode(v int64) *AddRelativePositionResponseBody {
	s.Code = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetSuccess(v bool) *AddRelativePositionResponseBody {
	s.Success = &v
	return s
}

func (s *AddRelativePositionResponseBody) SetMessage(v string) *AddRelativePositionResponseBody {
	s.Message = &v
	return s
}

type AddRelativePositionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRelativePositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddRelativePositionResponse) SetBody(v *AddRelativePositionResponseBody) *AddRelativePositionResponse {
	s.Body = v
	return s
}

type DetailSceneRequest struct {
	// 场景Id
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 主场景Id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 场景类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 子场景数
	SubSceneNum *int64 `json:"SubSceneNum,omitempty" xml:"SubSceneNum,omitempty"`
	// 资源数
	SourceNum *int64 `json:"SourceNum,omitempty" xml:"SourceNum,omitempty"`
	// 是否已发布 true：已发布：false：未发布
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 预览Token
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s DetailSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DetailSceneResponseBody) SetRequestId(v string) *DetailSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailSceneResponseBody) SetCode(v int64) *DetailSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DetailSceneResponseBody) SetSuccess(v bool) *DetailSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DetailSceneResponseBody) SetMessage(v string) *DetailSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DetailSceneResponseBody) SetId(v string) *DetailSceneResponseBody {
	s.Id = &v
	return s
}

func (s *DetailSceneResponseBody) SetName(v string) *DetailSceneResponseBody {
	s.Name = &v
	return s
}

func (s *DetailSceneResponseBody) SetType(v string) *DetailSceneResponseBody {
	s.Type = &v
	return s
}

func (s *DetailSceneResponseBody) SetSubSceneNum(v int64) *DetailSceneResponseBody {
	s.SubSceneNum = &v
	return s
}

func (s *DetailSceneResponseBody) SetSourceNum(v int64) *DetailSceneResponseBody {
	s.SourceNum = &v
	return s
}

func (s *DetailSceneResponseBody) SetPublished(v bool) *DetailSceneResponseBody {
	s.Published = &v
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

func (s *DetailSceneResponseBody) SetPreviewToken(v string) *DetailSceneResponseBody {
	s.PreviewToken = &v
	return s
}

type DetailSceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetailSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetailSceneResponse) SetBody(v *DetailSceneResponseBody) *DetailSceneResponse {
	s.Body = v
	return s
}

type CreateSceneRequest struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneRequest) SetName(v string) *CreateSceneRequest {
	s.Name = &v
	return s
}

func (s *CreateSceneRequest) SetProjectId(v string) *CreateSceneRequest {
	s.ProjectId = &v
	return s
}

type CreateSceneResponseBody struct {
	SceneId      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneResponseBody) SetSceneId(v int64) *CreateSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *CreateSceneResponseBody) SetRequestId(v string) *CreateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneResponseBody) SetPreviewToken(v string) *CreateSceneResponseBody {
	s.PreviewToken = &v
	return s
}

func (s *CreateSceneResponseBody) SetErrMessage(v string) *CreateSceneResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSceneResponseBody) SetSuccess(v bool) *CreateSceneResponseBody {
	s.Success = &v
	return s
}

type CreateSceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneResponse) SetHeaders(v map[string]*string) *CreateSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneResponse) SetBody(v *CreateSceneResponseBody) *CreateSceneResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	ParamFile    *string `json:"ParamFile,omitempty" xml:"ParamFile,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetParamFile(v string) *DeleteFileRequest {
	s.ParamFile = &v
	return s
}

func (s *DeleteFileRequest) SetSubSceneUuid(v string) *DeleteFileRequest {
	s.SubSceneUuid = &v
	return s
}

type DeleteFileResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrMessage(v string) *DeleteFileResponseBody {
	s.ErrMessage = &v
	return s
}

type DeleteFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type CheckResourceRequest struct {
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceRequest) GoString() string {
	return s.String()
}

func (s *CheckResourceRequest) SetCountry(v string) *CheckResourceRequest {
	s.Country = &v
	return s
}

func (s *CheckResourceRequest) SetInterrupt(v bool) *CheckResourceRequest {
	s.Interrupt = &v
	return s
}

func (s *CheckResourceRequest) SetInvoker(v string) *CheckResourceRequest {
	s.Invoker = &v
	return s
}

func (s *CheckResourceRequest) SetPk(v string) *CheckResourceRequest {
	s.Pk = &v
	return s
}

func (s *CheckResourceRequest) SetBid(v string) *CheckResourceRequest {
	s.Bid = &v
	return s
}

func (s *CheckResourceRequest) SetHid(v int64) *CheckResourceRequest {
	s.Hid = &v
	return s
}

func (s *CheckResourceRequest) SetTaskIdentifier(v string) *CheckResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *CheckResourceRequest) SetTaskExtraData(v string) *CheckResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *CheckResourceRequest) SetGmtWakeup(v string) *CheckResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *CheckResourceRequest) SetSuccess(v bool) *CheckResourceRequest {
	s.Success = &v
	return s
}

func (s *CheckResourceRequest) SetMessage(v string) *CheckResourceRequest {
	s.Message = &v
	return s
}

func (s *CheckResourceRequest) SetLevel(v int64) *CheckResourceRequest {
	s.Level = &v
	return s
}

func (s *CheckResourceRequest) SetUrl(v string) *CheckResourceRequest {
	s.Url = &v
	return s
}

func (s *CheckResourceRequest) SetPrompt(v string) *CheckResourceRequest {
	s.Prompt = &v
	return s
}

type CheckResourceResponseBody struct {
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
}

func (s CheckResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourceResponseBody) SetGmtWakeup(v string) *CheckResourceResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *CheckResourceResponseBody) SetHid(v int64) *CheckResourceResponseBody {
	s.Hid = &v
	return s
}

func (s *CheckResourceResponseBody) SetMessage(v string) *CheckResourceResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResourceResponseBody) SetTaskIdentifier(v string) *CheckResourceResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *CheckResourceResponseBody) SetRequestId(v string) *CheckResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourceResponseBody) SetSuccess(v bool) *CheckResourceResponseBody {
	s.Success = &v
	return s
}

func (s *CheckResourceResponseBody) SetUrl(v string) *CheckResourceResponseBody {
	s.Url = &v
	return s
}

func (s *CheckResourceResponseBody) SetInterrupt(v bool) *CheckResourceResponseBody {
	s.Interrupt = &v
	return s
}

func (s *CheckResourceResponseBody) SetInvoker(v string) *CheckResourceResponseBody {
	s.Invoker = &v
	return s
}

func (s *CheckResourceResponseBody) SetTaskExtraData(v string) *CheckResourceResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *CheckResourceResponseBody) SetCountry(v string) *CheckResourceResponseBody {
	s.Country = &v
	return s
}

func (s *CheckResourceResponseBody) SetPrompt(v string) *CheckResourceResponseBody {
	s.Prompt = &v
	return s
}

func (s *CheckResourceResponseBody) SetLevel(v int64) *CheckResourceResponseBody {
	s.Level = &v
	return s
}

func (s *CheckResourceResponseBody) SetPk(v string) *CheckResourceResponseBody {
	s.Pk = &v
	return s
}

func (s *CheckResourceResponseBody) SetBid(v string) *CheckResourceResponseBody {
	s.Bid = &v
	return s
}

type CheckResourceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceResponse) GoString() string {
	return s.String()
}

func (s *CheckResourceResponse) SetHeaders(v map[string]*string) *CheckResourceResponse {
	s.Headers = v
	return s
}

func (s *CheckResourceResponse) SetBody(v *CheckResourceResponseBody) *CheckResourceResponse {
	s.Body = v
	return s
}

type ListSceneRequest struct {
	// 主场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 所有项目Id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 当前页
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页长
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListSceneRequest) SetProjectId(v string) *ListSceneRequest {
	s.ProjectId = &v
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

type ListSceneResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否有下一页
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// 当前页
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 总页数
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// 数据总数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 主场景数据
	List []*ListSceneResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBody) SetRequestId(v string) *ListSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSceneResponseBody) SetCode(v int64) *ListSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSceneResponseBody) SetSuccess(v bool) *ListSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListSceneResponseBody) SetMessage(v string) *ListSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSceneResponseBody) SetHasNext(v bool) *ListSceneResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListSceneResponseBody) SetCurrentPage(v int64) *ListSceneResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSceneResponseBody) SetTotalPage(v int64) *ListSceneResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSceneResponseBody) SetCount(v int64) *ListSceneResponseBody {
	s.Count = &v
	return s
}

func (s *ListSceneResponseBody) SetList(v []*ListSceneResponseBodyList) *ListSceneResponseBody {
	s.List = v
	return s
}

type ListSceneResponseBodyList struct {
	// 主场景Id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 场景类型 3D模型：MODEL_3D  全景图片：PIC  全景视频：VIDEO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 子场景数
	SubSceneNum *int64 `json:"SubSceneNum,omitempty" xml:"SubSceneNum,omitempty"`
	// 资源数
	SourceNum *int64 `json:"SourceNum,omitempty" xml:"SourceNum,omitempty"`
	// 是否已发布 true：已发布：false：未发布
	Published *bool `json:"Published,omitempty" xml:"Published,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 预览Token
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s ListSceneResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSceneResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSceneResponseBodyList) SetId(v string) *ListSceneResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSceneResponseBodyList) SetName(v string) *ListSceneResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListSceneResponseBodyList) SetType(v string) *ListSceneResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListSceneResponseBodyList) SetSubSceneNum(v int64) *ListSceneResponseBodyList {
	s.SubSceneNum = &v
	return s
}

func (s *ListSceneResponseBodyList) SetSourceNum(v int64) *ListSceneResponseBodyList {
	s.SourceNum = &v
	return s
}

func (s *ListSceneResponseBodyList) SetPublished(v bool) *ListSceneResponseBodyList {
	s.Published = &v
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

func (s *ListSceneResponseBodyList) SetPreviewToken(v string) *ListSceneResponseBodyList {
	s.PreviewToken = &v
	return s
}

type ListSceneResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSceneResponse) SetBody(v *ListSceneResponseBody) *ListSceneResponse {
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
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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

func (s *PublishHotspotResponseBody) SetRequestId(v string) *PublishHotspotResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishHotspotResponseBody) SetSuccess(v bool) *PublishHotspotResponseBody {
	s.Success = &v
	return s
}

func (s *PublishHotspotResponseBody) SetErrMessage(v string) *PublishHotspotResponseBody {
	s.ErrMessage = &v
	return s
}

type PublishHotspotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishHotspotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishHotspotResponse) SetBody(v *PublishHotspotResponseBody) *PublishHotspotResponse {
	s.Body = v
	return s
}

type UpdateSceneRequest struct {
	// 场景Id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 场景名称
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponseBody) SetRequestId(v string) *UpdateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSceneResponseBody) SetCode(v int64) *UpdateSceneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSceneResponseBody) SetSuccess(v bool) *UpdateSceneResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSceneResponseBody) SetMessage(v string) *UpdateSceneResponseBody {
	s.Message = &v
	return s
}

type UpdateSceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateSceneResponse) SetBody(v *UpdateSceneResponseBody) *UpdateSceneResponse {
	s.Body = v
	return s
}

type UpdateLayoutDataRequest struct {
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// 标注数据
	LayoutData *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
}

func (s UpdateLayoutDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataRequest) SetSubSceneId(v string) *UpdateLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

func (s *UpdateLayoutDataRequest) SetLayoutData(v string) *UpdateLayoutDataRequest {
	s.LayoutData = &v
	return s
}

type UpdateLayoutDataResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateLayoutDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataResponseBody) SetRequestId(v string) *UpdateLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetCode(v int64) *UpdateLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetSuccess(v bool) *UpdateLayoutDataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLayoutDataResponseBody) SetMessage(v string) *UpdateLayoutDataResponseBody {
	s.Message = &v
	return s
}

type UpdateLayoutDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateLayoutDataResponse) SetBody(v *UpdateLayoutDataResponseBody) *UpdateLayoutDataResponse {
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s SaveHotspotTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotTagResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponseBody) SetRequestId(v string) *SaveHotspotTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotTagResponseBody) SetSuccess(v bool) *SaveHotspotTagResponseBody {
	s.Success = &v
	return s
}

func (s *SaveHotspotTagResponseBody) SetErrMessage(v string) *SaveHotspotTagResponseBody {
	s.ErrMessage = &v
	return s
}

type SaveHotspotTagResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveHotspotTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SaveHotspotTagResponse) SetBody(v *SaveHotspotTagResponseBody) *SaveHotspotTagResponse {
	s.Body = v
	return s
}

type RecoveryOriginImageRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecoveryOriginImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoveryOriginImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageResponseBody) SetRequestId(v string) *RecoveryOriginImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetCode(v int64) *RecoveryOriginImageResponseBody {
	s.Code = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetSuccess(v bool) *RecoveryOriginImageResponseBody {
	s.Success = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetMessage(v string) *RecoveryOriginImageResponseBody {
	s.Message = &v
	return s
}

type RecoveryOriginImageResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecoveryOriginImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecoveryOriginImageResponse) SetBody(v *RecoveryOriginImageResponseBody) *RecoveryOriginImageResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteProjectResponseBody) SetErrMessage(v string) *DeleteProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
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

type RectVerticalRequest struct {
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// 矫正数据
	VerticalRect *string `json:"VerticalRect,omitempty" xml:"VerticalRect,omitempty"`
	// 是否开启门预测
	DetectDoor *bool `json:"DetectDoor,omitempty" xml:"DetectDoor,omitempty"`
	// 需要预测的门的数量
	CountDetectDoor *int32 `json:"CountDetectDoor,omitempty" xml:"CountDetectDoor,omitempty"`
}

func (s RectVerticalRequest) String() string {
	return tea.Prettify(s)
}

func (s RectVerticalRequest) GoString() string {
	return s.String()
}

func (s *RectVerticalRequest) SetSubSceneId(v string) *RectVerticalRequest {
	s.SubSceneId = &v
	return s
}

func (s *RectVerticalRequest) SetVerticalRect(v string) *RectVerticalRequest {
	s.VerticalRect = &v
	return s
}

func (s *RectVerticalRequest) SetDetectDoor(v bool) *RectVerticalRequest {
	s.DetectDoor = &v
	return s
}

func (s *RectVerticalRequest) SetCountDetectDoor(v int32) *RectVerticalRequest {
	s.CountDetectDoor = &v
	return s
}

type RectVerticalResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RectVerticalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RectVerticalResponseBody) GoString() string {
	return s.String()
}

func (s *RectVerticalResponseBody) SetRequestId(v string) *RectVerticalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectVerticalResponseBody) SetCode(v int64) *RectVerticalResponseBody {
	s.Code = &v
	return s
}

func (s *RectVerticalResponseBody) SetSuccess(v bool) *RectVerticalResponseBody {
	s.Success = &v
	return s
}

func (s *RectVerticalResponseBody) SetMessage(v string) *RectVerticalResponseBody {
	s.Message = &v
	return s
}

type RectVerticalResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RectVerticalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RectVerticalResponse) SetBody(v *RectVerticalResponseBody) *RectVerticalResponse {
	s.Body = v
	return s
}

type PredImageRequest struct {
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// 是否门预测
	DetectDoor *bool `json:"DetectDoor,omitempty" xml:"DetectDoor,omitempty"`
	// 门数量(DetectDoor为false时，可为0)
	CountDetectDoor *int64 `json:"CountDetectDoor,omitempty" xml:"CountDetectDoor,omitempty"`
	// 是否垂直矫正
	CorrectVertical *bool `json:"CorrectVertical,omitempty" xml:"CorrectVertical,omitempty"`
}

func (s PredImageRequest) String() string {
	return tea.Prettify(s)
}

func (s PredImageRequest) GoString() string {
	return s.String()
}

func (s *PredImageRequest) SetSubSceneId(v string) *PredImageRequest {
	s.SubSceneId = &v
	return s
}

func (s *PredImageRequest) SetDetectDoor(v bool) *PredImageRequest {
	s.DetectDoor = &v
	return s
}

func (s *PredImageRequest) SetCountDetectDoor(v int64) *PredImageRequest {
	s.CountDetectDoor = &v
	return s
}

func (s *PredImageRequest) SetCorrectVertical(v bool) *PredImageRequest {
	s.CorrectVertical = &v
	return s
}

type PredImageResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PredImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredImageResponseBody) GoString() string {
	return s.String()
}

func (s *PredImageResponseBody) SetRequestId(v string) *PredImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredImageResponseBody) SetCode(v int64) *PredImageResponseBody {
	s.Code = &v
	return s
}

func (s *PredImageResponseBody) SetSuccess(v bool) *PredImageResponseBody {
	s.Success = &v
	return s
}

func (s *PredImageResponseBody) SetMessage(v string) *PredImageResponseBody {
	s.Message = &v
	return s
}

func (s *PredImageResponseBody) SetTaskId(v string) *PredImageResponseBody {
	s.TaskId = &v
	return s
}

type PredImageResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PredImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PredImageResponse) SetBody(v *PredImageResponseBody) *PredImageResponse {
	s.Body = v
	return s
}

type GetOssPolicyRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// accessId
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// 授权
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 签名
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 授权路径
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// 上传地址
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 授权失效时间(s)
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// 上传回调
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
}

func (s GetOssPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssPolicyResponseBody) SetRequestId(v string) *GetOssPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetCode(v int64) *GetOssPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetSuccess(v bool) *GetOssPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetMessage(v string) *GetOssPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetAccessId(v string) *GetOssPolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetPolicy(v string) *GetOssPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetSignature(v string) *GetOssPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetDir(v string) *GetOssPolicyResponseBody {
	s.Dir = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetHost(v string) *GetOssPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetExpire(v string) *GetOssPolicyResponseBody {
	s.Expire = &v
	return s
}

func (s *GetOssPolicyResponseBody) SetCallback(v string) *GetOssPolicyResponseBody {
	s.Callback = &v
	return s
}

type GetOssPolicyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOssPolicyResponse) SetBody(v *GetOssPolicyResponseBody) *GetOssPolicyResponse {
	s.Body = v
	return s
}

type GetConnDataRequest struct {
	// 场景ID
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
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// 扩展信息
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// 关联信息
	List []*GetConnDataResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s GetConnDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnDataResponseBody) SetRequestId(v string) *GetConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnDataResponseBody) SetCode(v int64) *GetConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnDataResponseBody) SetSuccess(v bool) *GetConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetConnDataResponseBody) SetMessage(v string) *GetConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnDataResponseBody) SetVersion(v string) *GetConnDataResponseBody {
	s.Version = &v
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

type GetConnDataResponseBodyList struct {
	// ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 关联的ID
	MapId *string `json:"MapId,omitempty" xml:"MapId,omitempty"`
	// outer:外关联 inner：内关联 stair：楼梯关联
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConnDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetConnDataResponse) SetBody(v *GetConnDataResponseBody) *GetConnDataResponse {
	s.Body = v
	return s
}

type RollbackSubSceneRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RollbackSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneResponseBody) SetRequestId(v string) *RollbackSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetCode(v int64) *RollbackSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetSuccess(v bool) *RollbackSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetMessage(v string) *RollbackSubSceneResponseBody {
	s.Message = &v
	return s
}

type RollbackSubSceneResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RollbackSubSceneResponse) SetBody(v *RollbackSubSceneResponseBody) *RollbackSubSceneResponse {
	s.Body = v
	return s
}

type TempPreviewStatusRequest struct {
	// 任务ID
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s TempPreviewStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewStatusRequest) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusRequest) SetKey(v string) *TempPreviewStatusRequest {
	s.Key = &v
	return s
}

type TempPreviewStatusResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// processing：处理中 success：成功 failed：失败
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s TempPreviewStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewStatusResponseBody) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusResponseBody) SetRequestId(v string) *TempPreviewStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetCode(v int64) *TempPreviewStatusResponseBody {
	s.Code = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetSuccess(v bool) *TempPreviewStatusResponseBody {
	s.Success = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetMessage(v string) *TempPreviewStatusResponseBody {
	s.Message = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetStatus(v string) *TempPreviewStatusResponseBody {
	s.Status = &v
	return s
}

type TempPreviewStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TempPreviewStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TempPreviewStatusResponse) SetBody(v *TempPreviewStatusResponseBody) *TempPreviewStatusResponse {
	s.Body = v
	return s
}

type AddProjectRequest struct {
	// 业务id
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// 项目名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 项目ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProjectResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectResponseBody) SetRequestId(v string) *AddProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectResponseBody) SetCode(v int64) *AddProjectResponseBody {
	s.Code = &v
	return s
}

func (s *AddProjectResponseBody) SetSuccess(v bool) *AddProjectResponseBody {
	s.Success = &v
	return s
}

func (s *AddProjectResponseBody) SetMessage(v string) *AddProjectResponseBody {
	s.Message = &v
	return s
}

func (s *AddProjectResponseBody) SetId(v string) *AddProjectResponseBody {
	s.Id = &v
	return s
}

type AddProjectResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddProjectResponse) SetBody(v *AddProjectResponseBody) *AddProjectResponse {
	s.Body = v
	return s
}

type DetailSubSceneRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 子场景id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 子场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 图片ID/视频ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 图片路径/视频路径
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片路径/视频封面路径
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 子场景状态
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 切图路径
	CubemapPath *string `json:"CubemapPath,omitempty" xml:"CubemapPath,omitempty"`
}

func (s DetailSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DetailSubSceneResponseBody) SetRequestId(v string) *DetailSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetCode(v int64) *DetailSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetSuccess(v bool) *DetailSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetMessage(v string) *DetailSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetId(v string) *DetailSubSceneResponseBody {
	s.Id = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetName(v string) *DetailSubSceneResponseBody {
	s.Name = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetResourceId(v string) *DetailSubSceneResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetUrl(v string) *DetailSubSceneResponseBody {
	s.Url = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetCoverUrl(v string) *DetailSubSceneResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *DetailSubSceneResponseBody) SetStatus(v int64) *DetailSubSceneResponseBody {
	s.Status = &v
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

func (s *DetailSubSceneResponseBody) SetCubemapPath(v string) *DetailSubSceneResponseBody {
	s.CubemapPath = &v
	return s
}

type DetailSubSceneResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetailSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetailSubSceneResponse) SetBody(v *DetailSubSceneResponseBody) *DetailSubSceneResponse {
	s.Body = v
	return s
}

type ListSubSceneRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 页码
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页长
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneRequest) GoString() string {
	return s.String()
}

func (s *ListSubSceneRequest) SetSceneId(v string) *ListSubSceneRequest {
	s.SceneId = &v
	return s
}

func (s *ListSubSceneRequest) SetPageNum(v int64) *ListSubSceneRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubSceneRequest) SetPageSize(v int64) *ListSubSceneRequest {
	s.PageSize = &v
	return s
}

type ListSubSceneResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否有下一页
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// 当前页
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 总页数
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// 数据总条数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 子场景列表集
	List []*ListSubSceneResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBody) SetRequestId(v string) *ListSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCode(v int64) *ListSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubSceneResponseBody) SetSuccess(v bool) *ListSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubSceneResponseBody) SetMessage(v string) *ListSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubSceneResponseBody) SetHasNext(v bool) *ListSubSceneResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCurrentPage(v int64) *ListSubSceneResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSubSceneResponseBody) SetTotalPage(v int64) *ListSubSceneResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSubSceneResponseBody) SetCount(v int64) *ListSubSceneResponseBody {
	s.Count = &v
	return s
}

func (s *ListSubSceneResponseBody) SetList(v []*ListSubSceneResponseBodyList) *ListSubSceneResponseBody {
	s.List = v
	return s
}

type ListSubSceneResponseBodyList struct {
	// 子场景ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 子场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 图片ID/视频ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 图片路径/视频路径
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 图片路径/视频封面路径
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 子场景状态 1.未重建，      * 2.中间模型重建中，      * 3.中间模型重建完成，      * 4.待重建，      * 5.服务商重建中，      * 6.服务商重建完成，      * 7.已发布      * 8.发布中
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// 切图的路径
	CubemapPath *string `json:"CubemapPath,omitempty" xml:"CubemapPath,omitempty"`
	// 是否删除
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// 原图地址
	OriginUrl *string `json:"OriginUrl,omitempty" xml:"OriginUrl,omitempty"`
	// 2k基准图路径
	BaseImageUrl *string `json:"BaseImageUrl,omitempty" xml:"BaseImageUrl,omitempty"`
}

func (s ListSubSceneResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSubSceneResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponseBodyList) SetId(v string) *ListSubSceneResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetName(v string) *ListSubSceneResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetResourceId(v string) *ListSubSceneResponseBodyList {
	s.ResourceId = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetUrl(v string) *ListSubSceneResponseBodyList {
	s.Url = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetCoverUrl(v string) *ListSubSceneResponseBodyList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetStatus(v int64) *ListSubSceneResponseBodyList {
	s.Status = &v
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

func (s *ListSubSceneResponseBodyList) SetResourceName(v string) *ListSubSceneResponseBodyList {
	s.ResourceName = &v
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

func (s *ListSubSceneResponseBodyList) SetOriginUrl(v string) *ListSubSceneResponseBodyList {
	s.OriginUrl = &v
	return s
}

func (s *ListSubSceneResponseBodyList) SetBaseImageUrl(v string) *ListSubSceneResponseBodyList {
	s.BaseImageUrl = &v
	return s
}

type ListSubSceneResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSubSceneResponse) SetBody(v *ListSubSceneResponseBody) *ListSubSceneResponse {
	s.Body = v
	return s
}

type UpdateSubSceneRequest struct {
	// 子场景ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 子场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type UpdateSubSceneResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneResponseBody) SetRequestId(v string) *UpdateSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetCode(v int64) *UpdateSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetSuccess(v bool) *UpdateSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSubSceneResponseBody) SetMessage(v string) *UpdateSubSceneResponseBody {
	s.Message = &v
	return s
}

type UpdateSubSceneResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateSubSceneResponse) SetBody(v *UpdateSubSceneResponseBody) *UpdateSubSceneResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	// 任务实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetInstanceId(v string) *GetJobRequest {
	s.InstanceId = &v
	return s
}

type GetJobResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务运行状态
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetCode(v int64) *GetJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v bool) *GetJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobResponseBody) SetMessage(v string) *GetJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobResponseBody) SetStatus(v int64) *GetJobResponseBody {
	s.Status = &v
	return s
}

type GetJobResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BusinessId         *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessUserIdList *string `json:"BusinessUserIdList,omitempty" xml:"BusinessUserIdList,omitempty"`
	GatherUserIdList   *string `json:"GatherUserIdList,omitempty" xml:"GatherUserIdList,omitempty"`
	BuilderUserIdList  *string `json:"BuilderUserIdList,omitempty" xml:"BuilderUserIdList,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetBusinessId(v string) *CreateProjectRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateProjectRequest) SetBusinessUserIdList(v string) *CreateProjectRequest {
	s.BusinessUserIdList = &v
	return s
}

func (s *CreateProjectRequest) SetGatherUserIdList(v string) *CreateProjectRequest {
	s.GatherUserIdList = &v
	return s
}

func (s *CreateProjectRequest) SetBuilderUserIdList(v string) *CreateProjectRequest {
	s.BuilderUserIdList = &v
	return s
}

type CreateProjectResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetId(v int64) *CreateProjectResponseBody {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrMessage(v string) *CreateProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProjectResponseBody) SetName(v string) *CreateProjectResponseBody {
	s.Name = &v
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s SaveHotspotConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponseBody) SetRequestId(v string) *SaveHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetSuccess(v bool) *SaveHotspotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetErrMessage(v string) *SaveHotspotConfigResponseBody {
	s.ErrMessage = &v
	return s
}

type SaveHotspotConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SaveHotspotConfigResponse) SetBody(v *SaveHotspotConfigResponseBody) *SaveHotspotConfigResponse {
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
	ObjectString *string                `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage   *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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

func (s *GetWindowConfigResponseBody) SetErrMessage(v string) *GetWindowConfigResponseBody {
	s.ErrMessage = &v
	return s
}

type GetWindowConfigResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWindowConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWindowConfigResponse) SetBody(v *GetWindowConfigResponseBody) *GetWindowConfigResponse {
	s.Body = v
	return s
}

type GetHotspotConfigRequest struct {
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	Type         *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetHotspotConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigRequest) SetPreviewToken(v string) *GetHotspotConfigRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotConfigRequest) SetType(v int64) *GetHotspotConfigRequest {
	s.Type = &v
	return s
}

func (s *GetHotspotConfigRequest) SetEnabled(v bool) *GetHotspotConfigRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotConfigRequest) SetDomain(v string) *GetHotspotConfigRequest {
	s.Domain = &v
	return s
}

type GetHotspotConfigResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetHotspotConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponseBody) SetRequestId(v string) *GetHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetCode(v int64) *GetHotspotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetSuccess(v bool) *GetHotspotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetMessage(v string) *GetHotspotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetData(v string) *GetHotspotConfigResponseBody {
	s.Data = &v
	return s
}

type GetHotspotConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHotspotConfigResponse) SetBody(v *GetHotspotConfigResponseBody) *GetHotspotConfigResponse {
	s.Body = v
	return s
}

type GetSceneBuildTaskStatusRequest struct {
	// 场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 未开始  init 处理中 失败     failure   processing  完成     succeed 取消     canceled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 墙线预测: wall_line  切图: cut_image  重建: build  直角优化：right_angle_optimization 其他：other
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 任务失败错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 任务失败错误消息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s GetSceneBuildTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneBuildTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusResponseBody) SetRequestId(v string) *GetSceneBuildTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetCode(v int64) *GetSceneBuildTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetSuccess(v bool) *GetSceneBuildTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetMessage(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetId(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Id = &v
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

func (s *GetSceneBuildTaskStatusResponseBody) SetType(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Type = &v
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

type GetSceneBuildTaskStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSceneBuildTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSceneBuildTaskStatusResponse) SetBody(v *GetSceneBuildTaskStatusResponseBody) *GetSceneBuildTaskStatusResponse {
	s.Body = v
	return s
}

type TempPreviewRequest struct {
	// 场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 预览链接
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// 任务ID
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s TempPreviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TempPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *TempPreviewResponseBody) SetRequestId(v string) *TempPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempPreviewResponseBody) SetCode(v int64) *TempPreviewResponseBody {
	s.Code = &v
	return s
}

func (s *TempPreviewResponseBody) SetSuccess(v bool) *TempPreviewResponseBody {
	s.Success = &v
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

func (s *TempPreviewResponseBody) SetKey(v string) *TempPreviewResponseBody {
	s.Key = &v
	return s
}

type TempPreviewResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TempPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TempPreviewResponse) SetBody(v *TempPreviewResponseBody) *TempPreviewResponse {
	s.Body = v
	return s
}

type PublishSceneRequest struct {
	// 场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 预览链接
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
	// 任务实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PublishSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishSceneResponseBody) GoString() string {
	return s.String()
}

func (s *PublishSceneResponseBody) SetRequestId(v string) *PublishSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishSceneResponseBody) SetCode(v int64) *PublishSceneResponseBody {
	s.Code = &v
	return s
}

func (s *PublishSceneResponseBody) SetSuccess(v bool) *PublishSceneResponseBody {
	s.Success = &v
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

func (s *PublishSceneResponseBody) SetInstanceId(v string) *PublishSceneResponseBody {
	s.InstanceId = &v
	return s
}

type PublishSceneResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishSceneResponse) SetBody(v *PublishSceneResponseBody) *PublishSceneResponse {
	s.Body = v
	return s
}

type DetailProjectRequest struct {
	// 项目Id
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 项目ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 项目名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 业务ID
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// 业务名称
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DetailProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetailProjectResponseBody) SetRequestId(v string) *DetailProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailProjectResponseBody) SetCode(v int64) *DetailProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DetailProjectResponseBody) SetSuccess(v bool) *DetailProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DetailProjectResponseBody) SetMessage(v string) *DetailProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DetailProjectResponseBody) SetId(v string) *DetailProjectResponseBody {
	s.Id = &v
	return s
}

func (s *DetailProjectResponseBody) SetName(v string) *DetailProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DetailProjectResponseBody) SetBusinessId(v int64) *DetailProjectResponseBody {
	s.BusinessId = &v
	return s
}

func (s *DetailProjectResponseBody) SetBusinessName(v string) *DetailProjectResponseBody {
	s.BusinessName = &v
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

func (s *DetailProjectResponseBody) SetToken(v string) *DetailProjectResponseBody {
	s.Token = &v
	return s
}

type DetailProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetailProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetailProjectResponse) SetBody(v *DetailProjectResponseBody) *DetailProjectResponse {
	s.Body = v
	return s
}

type ListScenesRequest struct {
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IsPublishQuery *bool   `json:"IsPublishQuery,omitempty" xml:"IsPublishQuery,omitempty"`
}

func (s ListScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenesRequest) GoString() string {
	return s.String()
}

func (s *ListScenesRequest) SetProjectId(v string) *ListScenesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListScenesRequest) SetIsPublishQuery(v bool) *ListScenesRequest {
	s.IsPublishQuery = &v
	return s
}

type ListScenesResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data       []*ListScenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrMessage *string                       `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBody) SetRequestId(v string) *ListScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenesResponseBody) SetData(v []*ListScenesResponseBodyData) *ListScenesResponseBody {
	s.Data = v
	return s
}

func (s *ListScenesResponseBody) SetErrMessage(v string) *ListScenesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListScenesResponseBody) SetSuccess(v bool) *ListScenesResponseBody {
	s.Success = &v
	return s
}

type ListScenesResponseBodyData struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListScenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBodyData) SetSceneId(v string) *ListScenesResponseBodyData {
	s.SceneId = &v
	return s
}

type ListScenesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScenesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponse) GoString() string {
	return s.String()
}

func (s *ListScenesResponse) SetHeaders(v map[string]*string) *ListScenesResponse {
	s.Headers = v
	return s
}

func (s *ListScenesResponse) SetBody(v *ListScenesResponseBody) *ListScenesResponse {
	s.Body = v
	return s
}

type DropSubSceneRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DropSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DropSubSceneResponseBody) SetRequestId(v string) *DropSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSubSceneResponseBody) SetCode(v int64) *DropSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DropSubSceneResponseBody) SetSuccess(v bool) *DropSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DropSubSceneResponseBody) SetMessage(v string) *DropSubSceneResponseBody {
	s.Message = &v
	return s
}

type DropSubSceneResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DropSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DropSubSceneResponse) SetBody(v *DropSubSceneResponseBody) *DropSubSceneResponse {
	s.Body = v
	return s
}

type GetHotspotTagRequest struct {
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetHotspotTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagRequest) GoString() string {
	return s.String()
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

func (s *GetHotspotTagRequest) SetEnabled(v bool) *GetHotspotTagRequest {
	s.Enabled = &v
	return s
}

func (s *GetHotspotTagRequest) SetDomain(v string) *GetHotspotTagRequest {
	s.Domain = &v
	return s
}

type GetHotspotTagResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponseBody) SetRequestId(v string) *GetHotspotTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetObjectString(v string) *GetHotspotTagResponseBody {
	s.ObjectString = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetData(v string) *GetHotspotTagResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetErrMessage(v string) *GetHotspotTagResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetSuccess(v bool) *GetHotspotTagResponseBody {
	s.Success = &v
	return s
}

type GetHotspotTagResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotspotTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHotspotTagResponse) SetBody(v *GetHotspotTagResponseBody) *GetHotspotTagResponse {
	s.Body = v
	return s
}

type DropProjectRequest struct {
	// 项目ID
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
	// 请求ID与入参中requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DropProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DropProjectResponseBody) SetRequestId(v string) *DropProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropProjectResponseBody) SetCode(v int64) *DropProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DropProjectResponseBody) SetSuccess(v bool) *DropProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DropProjectResponseBody) SetMessage(v string) *DropProjectResponseBody {
	s.Message = &v
	return s
}

type DropProjectResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DropProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DropProjectResponse) SetBody(v *DropProjectResponseBody) *DropProjectResponse {
	s.Body = v
	return s
}

type ListProjectRequest struct {
	// 页码
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 页长
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目名称（使用name%搜索）
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) SetPageNum(v int64) *ListProjectRequest {
	s.PageNum = &v
	return s
}

func (s *ListProjectRequest) SetPageSize(v int64) *ListProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRequest) SetName(v string) *ListProjectRequest {
	s.Name = &v
	return s
}

type ListProjectResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否有下一页
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// 当前页
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 总页数
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// count
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 项目数据
	List []*ListProjectResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectResponseBody) SetCode(v int64) *ListProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectResponseBody) SetSuccess(v bool) *ListProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectResponseBody) SetMessage(v string) *ListProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectResponseBody) SetHasNext(v bool) *ListProjectResponseBody {
	s.HasNext = &v
	return s
}

func (s *ListProjectResponseBody) SetCurrentPage(v int64) *ListProjectResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListProjectResponseBody) SetTotalPage(v int64) *ListProjectResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetList(v []*ListProjectResponseBodyList) *ListProjectResponseBody {
	s.List = v
	return s
}

type ListProjectResponseBodyList struct {
	// 项目ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 项目名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 业务ID
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// 业务名称
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 最后修改时间
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListProjectResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyList) SetId(v string) *ListProjectResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListProjectResponseBodyList) SetName(v string) *ListProjectResponseBodyList {
	s.Name = &v
	return s
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

func (s *ListProjectResponseBodyList) SetModifiedTime(v int64) *ListProjectResponseBodyList {
	s.ModifiedTime = &v
	return s
}

func (s *ListProjectResponseBodyList) SetToken(v string) *ListProjectResponseBodyList {
	s.Token = &v
	return s
}

type ListProjectResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProjectResponse) SetBody(v *ListProjectResponseBody) *ListProjectResponse {
	s.Body = v
	return s
}

type GetOriginLayoutDataRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 标注数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetOriginLayoutDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOriginLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataResponseBody) SetRequestId(v string) *GetOriginLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetCode(v int64) *GetOriginLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetSuccess(v bool) *GetOriginLayoutDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetMessage(v string) *GetOriginLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetData(v string) *GetOriginLayoutDataResponseBody {
	s.Data = &v
	return s
}

type GetOriginLayoutDataResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOriginLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOriginLayoutDataResponse) SetBody(v *GetOriginLayoutDataResponseBody) *GetOriginLayoutDataResponse {
	s.Body = v
	return s
}

type GetHotspotSceneDataRequest struct {
	// 预览token
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// 0 未发布， 1 已发布
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 自定义oss域名（可为cdn域名）
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 是否开启自用资源访问
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetHotspotSceneDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataRequest) SetPreviewToken(v string) *GetHotspotSceneDataRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetType(v int64) *GetHotspotSceneDataRequest {
	s.Type = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetDomain(v string) *GetHotspotSceneDataRequest {
	s.Domain = &v
	return s
}

func (s *GetHotspotSceneDataRequest) SetEnabled(v bool) *GetHotspotSceneDataRequest {
	s.Enabled = &v
	return s
}

type GetHotspotSceneDataResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    *GetHotspotSceneDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetHotspotSceneDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBody) SetRequestId(v string) *GetHotspotSceneDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetCode(v int64) *GetHotspotSceneDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetSuccess(v bool) *GetHotspotSceneDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetMessage(v string) *GetHotspotSceneDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotSceneDataResponseBody) SetData(v *GetHotspotSceneDataResponseBodyData) *GetHotspotSceneDataResponseBody {
	s.Data = v
	return s
}

type GetHotspotSceneDataResponseBodyData struct {
	// 3D模型：MODEL_3D 全景图片：PIC 全景视频：VIDEO
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// 预览token
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
	// html转译后的预览数据，包含图片、子场景ID等信息
	PreviewData *string `json:"PreviewData,omitempty" xml:"PreviewData,omitempty"`
	// 模型token（sgm token）
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
}

func (s GetHotspotSceneDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotSceneDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotSceneDataResponseBodyData) SetSceneType(v string) *GetHotspotSceneDataResponseBodyData {
	s.SceneType = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetPreviewToken(v string) *GetHotspotSceneDataResponseBodyData {
	s.PreviewToken = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetPreviewData(v string) *GetHotspotSceneDataResponseBodyData {
	s.PreviewData = &v
	return s
}

func (s *GetHotspotSceneDataResponseBodyData) SetModelToken(v string) *GetHotspotSceneDataResponseBodyData {
	s.ModelToken = &v
	return s
}

type GetHotspotSceneDataResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotspotSceneDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHotspotSceneDataResponse) SetBody(v *GetHotspotSceneDataResponseBody) *GetHotspotSceneDataResponse {
	s.Body = v
	return s
}

type AddMosaicsRequest struct {
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// 马赛克位置数据
	MarkPosition *string `json:"MarkPosition,omitempty" xml:"MarkPosition,omitempty"`
}

func (s AddMosaicsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMosaicsRequest) GoString() string {
	return s.String()
}

func (s *AddMosaicsRequest) SetSubSceneId(v string) *AddMosaicsRequest {
	s.SubSceneId = &v
	return s
}

func (s *AddMosaicsRequest) SetMarkPosition(v string) *AddMosaicsRequest {
	s.MarkPosition = &v
	return s
}

type AddMosaicsResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddMosaicsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMosaicsResponseBody) GoString() string {
	return s.String()
}

func (s *AddMosaicsResponseBody) SetRequestId(v string) *AddMosaicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMosaicsResponseBody) SetCode(v int64) *AddMosaicsResponseBody {
	s.Code = &v
	return s
}

func (s *AddMosaicsResponseBody) SetSuccess(v bool) *AddMosaicsResponseBody {
	s.Success = &v
	return s
}

func (s *AddMosaicsResponseBody) SetMessage(v string) *AddMosaicsResponseBody {
	s.Message = &v
	return s
}

func (s *AddMosaicsResponseBody) SetTaskId(v string) *AddMosaicsResponseBody {
	s.TaskId = &v
	return s
}

type AddMosaicsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddMosaicsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddMosaicsResponse) SetBody(v *AddMosaicsResponseBody) *AddMosaicsResponse {
	s.Body = v
	return s
}

type ScenePublishRequest struct {
	// 场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 预览链接
	PreviewUrl *string `json:"PreviewUrl,omitempty" xml:"PreviewUrl,omitempty"`
}

func (s ScenePublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScenePublishResponseBody) GoString() string {
	return s.String()
}

func (s *ScenePublishResponseBody) SetRequestId(v string) *ScenePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScenePublishResponseBody) SetCode(v int64) *ScenePublishResponseBody {
	s.Code = &v
	return s
}

func (s *ScenePublishResponseBody) SetSuccess(v bool) *ScenePublishResponseBody {
	s.Success = &v
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

type ScenePublishResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScenePublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ScenePublishResponse) SetBody(v *ScenePublishResponseBody) *ScenePublishResponse {
	s.Body = v
	return s
}

type GetRectifyImageRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 图片地址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRectifyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRectifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetRectifyImageResponseBody) SetRequestId(v string) *GetRectifyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetCode(v int64) *GetRectifyImageResponseBody {
	s.Code = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetSuccess(v bool) *GetRectifyImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetMessage(v string) *GetRectifyImageResponseBody {
	s.Message = &v
	return s
}

func (s *GetRectifyImageResponseBody) SetUrl(v string) *GetRectifyImageResponseBody {
	s.Url = &v
	return s
}

type GetRectifyImageResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRectifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRectifyImageResponse) SetBody(v *GetRectifyImageResponseBody) *GetRectifyImageResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	// 项目id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 项目名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 业务Id
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetId(v string) *UpdateProjectRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectRequest) SetBusinessId(v string) *UpdateProjectRequest {
	s.BusinessId = &v
	return s
}

type UpdateProjectResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *UpdateProjectResponseBody) SetCode(v int64) *UpdateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectResponseBody) SetMessage(v string) *UpdateProjectResponseBody {
	s.Message = &v
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

type GetSubSceneTaskStatusRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务信息
	List []*GetSubSceneTaskStatusResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s GetSubSceneTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBody) SetRequestId(v string) *GetSubSceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetCode(v int64) *GetSubSceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetSuccess(v bool) *GetSubSceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetMessage(v string) *GetSubSceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetList(v []*GetSubSceneTaskStatusResponseBodyList) *GetSubSceneTaskStatusResponseBody {
	s.List = v
	return s
}

type GetSubSceneTaskStatusResponseBodyList struct {
	// 任务ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// 未开始  init 处理中   processing   失败     failure  完成     succeed  取消     canceled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 墙线预测: wall_line   切图: cut_image 重建: build  直角优化：right_angle_optimization 其他：other
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 任务失败错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 任务失败错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s GetSubSceneTaskStatusResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetSceneId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.SceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetSubSceneId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.SubSceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetStatus(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Status = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetType(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Type = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetErrorCode(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.ErrorCode = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetErrorMsg(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.ErrorMsg = &v
	return s
}

type GetSubSceneTaskStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSubSceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSubSceneTaskStatusResponse) SetBody(v *GetSubSceneTaskStatusResponseBody) *GetSubSceneTaskStatusResponse {
	s.Body = v
	return s
}

type PredictionWallLineRequest struct {
	// 图片地址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 相机高度 单位 cm
	CameraHeight *int64 `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
}

func (s PredictionWallLineRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictionWallLineRequest) GoString() string {
	return s.String()
}

func (s *PredictionWallLineRequest) SetUrl(v string) *PredictionWallLineRequest {
	s.Url = &v
	return s
}

func (s *PredictionWallLineRequest) SetCameraHeight(v int64) *PredictionWallLineRequest {
	s.CameraHeight = &v
	return s
}

type PredictionWallLineResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 子场景ID
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s PredictionWallLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictionWallLineResponseBody) GoString() string {
	return s.String()
}

func (s *PredictionWallLineResponseBody) SetRequestId(v string) *PredictionWallLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetCode(v int64) *PredictionWallLineResponseBody {
	s.Code = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetSuccess(v bool) *PredictionWallLineResponseBody {
	s.Success = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetMessage(v string) *PredictionWallLineResponseBody {
	s.Message = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetTaskId(v string) *PredictionWallLineResponseBody {
	s.TaskId = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetSubSceneId(v string) *PredictionWallLineResponseBody {
	s.SubSceneId = &v
	return s
}

type PredictionWallLineResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PredictionWallLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PredictionWallLineResponse) SetBody(v *PredictionWallLineResponseBody) *PredictionWallLineResponse {
	s.Body = v
	return s
}

type GetScenePreviewInfoRequest struct {
	// 模型token
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
	// 自定义oss域名（可为cdn域名）
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 是否开启自用资源访问
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetScenePreviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoRequest) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoRequest) SetModelToken(v string) *GetScenePreviewInfoRequest {
	s.ModelToken = &v
	return s
}

func (s *GetScenePreviewInfoRequest) SetDomain(v string) *GetScenePreviewInfoRequest {
	s.Domain = &v
	return s
}

func (s *GetScenePreviewInfoRequest) SetEnabled(v bool) *GetScenePreviewInfoRequest {
	s.Enabled = &v
	return s
}

type GetScenePreviewInfoResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    *GetScenePreviewInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetScenePreviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBody) SetRequestId(v string) *GetScenePreviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetCode(v int64) *GetScenePreviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetSuccess(v bool) *GetScenePreviewInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetMessage(v string) *GetScenePreviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetScenePreviewInfoResponseBody) SetData(v *GetScenePreviewInfoResponseBodyData) *GetScenePreviewInfoResponseBody {
	s.Data = v
	return s
}

type GetScenePreviewInfoResponseBodyData struct {
	// html转译后的预览数据
	PanoList *string `json:"PanoList,omitempty" xml:"PanoList,omitempty"`
	// 模型地址
	ModelPath *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
	// 模型的贴图路径
	TextureModelPath *string `json:"TextureModelPath,omitempty" xml:"TextureModelPath,omitempty"`
	// 漫游后预览图片路径
	TexturePanoPath *string `json:"TexturePanoPath,omitempty" xml:"TexturePanoPath,omitempty"`
}

func (s GetScenePreviewInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScenePreviewInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScenePreviewInfoResponseBodyData) SetPanoList(v string) *GetScenePreviewInfoResponseBodyData {
	s.PanoList = &v
	return s
}

func (s *GetScenePreviewInfoResponseBodyData) SetModelPath(v string) *GetScenePreviewInfoResponseBodyData {
	s.ModelPath = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScenePreviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetScenePreviewInfoResponse) SetBody(v *GetScenePreviewInfoResponseBody) *GetScenePreviewInfoResponse {
	s.Body = v
	return s
}

type GetPolicyRequest struct {
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) SetSubSceneUuid(v string) *GetPolicyRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *GetPolicyRequest) SetType(v string) *GetPolicyRequest {
	s.Type = &v
	return s
}

type GetPolicyResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ObjectString *string                `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage   *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) SetData(v map[string]interface{}) *GetPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyResponseBody) SetObjectString(v string) *GetPolicyResponseBody {
	s.ObjectString = &v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponseBody) SetSuccess(v bool) *GetPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GetPolicyResponseBody) SetErrMessage(v string) *GetPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

type GetPolicyResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) SetHeaders(v map[string]*string) *GetPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyResponse) SetBody(v *GetPolicyResponseBody) *GetPolicyResponse {
	s.Body = v
	return s
}

type AddSubSceneRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 子场景名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AddSubSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSubSceneRequest) GoString() string {
	return s.String()
}

func (s *AddSubSceneRequest) SetSceneId(v string) *AddSubSceneRequest {
	s.SceneId = &v
	return s
}

func (s *AddSubSceneRequest) SetName(v string) *AddSubSceneRequest {
	s.Name = &v
	return s
}

type AddSubSceneResponseBody struct {
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 子场景ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddSubSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *AddSubSceneResponseBody) SetRequestId(v string) *AddSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSubSceneResponseBody) SetCode(v int64) *AddSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *AddSubSceneResponseBody) SetSuccess(v bool) *AddSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *AddSubSceneResponseBody) SetMessage(v string) *AddSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *AddSubSceneResponseBody) SetId(v string) *AddSubSceneResponseBody {
	s.Id = &v
	return s
}

type AddSubSceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddSubSceneResponse) SetBody(v *AddSubSceneResponseBody) *AddSubSceneResponse {
	s.Body = v
	return s
}

type GetLayoutDataRequest struct {
	// 子场景ID
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
	// 请求ID，与入参requestId对应
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回码
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否请求成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 标注信息
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetLayoutDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetLayoutDataResponseBody) SetRequestId(v string) *GetLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetCode(v int64) *GetLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetSuccess(v bool) *GetLayoutDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetMessage(v string) *GetLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetData(v string) *GetLayoutDataResponseBody {
	s.Data = &v
	return s
}

type GetLayoutDataResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetLayoutDataResponse) SetBody(v *GetLayoutDataResponseBody) *GetLayoutDataResponse {
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

func (client *Client) GetSingleConnDataWithOptions(request *GetSingleConnDataRequest, runtime *util.RuntimeOptions) (_result *GetSingleConnDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSingleConnDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSingleConnData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) LinkImageWithOptions(request *LinkImageRequest, runtime *util.RuntimeOptions) (_result *LinkImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LinkImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LinkImage"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddSceneWithOptions(request *AddSceneRequest, runtime *util.RuntimeOptions) (_result *AddSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateConnDataWithOptions(request *UpdateConnDataRequest, runtime *util.RuntimeOptions) (_result *UpdateConnDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConnDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConnData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RectifyImageWithOptions(request *RectifyImageRequest, runtime *util.RuntimeOptions) (_result *RectifyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RectifyImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RectifyImage"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) LabelBuildWithOptions(request *LabelBuildRequest, runtime *util.RuntimeOptions) (_result *LabelBuildResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LabelBuildResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LabelBuild"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DropSceneWithOptions(request *DropSceneRequest, runtime *util.RuntimeOptions) (_result *DropSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DropSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DropScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) OptimizeRightAngleWithOptions(request *OptimizeRightAngleRequest, runtime *util.RuntimeOptions) (_result *OptimizeRightAngleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OptimizeRightAngleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OptimizeRightAngle"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddRelativePositionWithOptions(request *AddRelativePositionRequest, runtime *util.RuntimeOptions) (_result *AddRelativePositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddRelativePositionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddRelativePosition"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DetailSceneWithOptions(request *DetailSceneRequest, runtime *util.RuntimeOptions) (_result *DetailSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetailSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetailScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateSceneWithOptions(request *CreateSceneRequest, runtime *util.RuntimeOptions) (_result *CreateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScene(request *CreateSceneRequest) (_result *CreateSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSceneResponse{}
	_body, _err := client.CreateSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFile"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckResourceWithOptions(request *CheckResourceRequest, runtime *util.RuntimeOptions) (_result *CheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckResource"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckResource(request *CheckResourceRequest) (_result *CheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResourceResponse{}
	_body, _err := client.CheckResourceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PublishHotspotWithOptions(request *PublishHotspotRequest, runtime *util.RuntimeOptions) (_result *PublishHotspotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishHotspotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishHotspot"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateSceneWithOptions(request *UpdateSceneRequest, runtime *util.RuntimeOptions) (_result *UpdateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateLayoutDataWithOptions(request *UpdateLayoutDataRequest, runtime *util.RuntimeOptions) (_result *UpdateLayoutDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLayoutDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLayoutData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SaveHotspotTagWithOptions(request *SaveHotspotTagRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveHotspotTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveHotspotTag"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RecoveryOriginImageWithOptions(request *RecoveryOriginImageRequest, runtime *util.RuntimeOptions) (_result *RecoveryOriginImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecoveryOriginImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecoveryOriginImage"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RectVerticalWithOptions(request *RectVerticalRequest, runtime *util.RuntimeOptions) (_result *RectVerticalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RectVerticalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RectVertical"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PredImageWithOptions(request *PredImageRequest, runtime *util.RuntimeOptions) (_result *PredImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PredImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PredImage"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetOssPolicyWithOptions(request *GetOssPolicyRequest, runtime *util.RuntimeOptions) (_result *GetOssPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOssPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOssPolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetConnDataWithOptions(request *GetConnDataRequest, runtime *util.RuntimeOptions) (_result *GetConnDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConnDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConnData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RollbackSubSceneWithOptions(request *RollbackSubSceneRequest, runtime *util.RuntimeOptions) (_result *RollbackSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollbackSubSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollbackSubScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) TempPreviewStatusWithOptions(request *TempPreviewStatusRequest, runtime *util.RuntimeOptions) (_result *TempPreviewStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TempPreviewStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TempPreviewStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddProjectWithOptions(request *AddProjectRequest, runtime *util.RuntimeOptions) (_result *AddProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DetailSubSceneWithOptions(request *DetailSubSceneRequest, runtime *util.RuntimeOptions) (_result *DetailSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetailSubSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetailSubScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListSubSceneWithOptions(request *ListSubSceneRequest, runtime *util.RuntimeOptions) (_result *ListSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSubSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSubScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateSubSceneWithOptions(request *UpdateSubSceneRequest, runtime *util.RuntimeOptions) (_result *UpdateSubSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSubSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSubScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SaveHotspotConfigWithOptions(request *SaveHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *SaveHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveHotspotConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveHotspotConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetWindowConfigWithOptions(request *GetWindowConfigRequest, runtime *util.RuntimeOptions) (_result *GetWindowConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWindowConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWindowConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetHotspotConfigWithOptions(request *GetHotspotConfigRequest, runtime *util.RuntimeOptions) (_result *GetHotspotConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotspotConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotspotConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetSceneBuildTaskStatusWithOptions(request *GetSceneBuildTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetSceneBuildTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSceneBuildTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSceneBuildTaskStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) TempPreviewWithOptions(request *TempPreviewRequest, runtime *util.RuntimeOptions) (_result *TempPreviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TempPreviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TempPreview"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PublishSceneWithOptions(request *PublishSceneRequest, runtime *util.RuntimeOptions) (_result *PublishSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DetailProjectWithOptions(request *DetailProjectRequest, runtime *util.RuntimeOptions) (_result *DetailProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetailProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetailProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListScenesWithOptions(request *ListScenesRequest, runtime *util.RuntimeOptions) (_result *ListScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListScenesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScenes"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScenes(request *ListScenesRequest) (_result *ListScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScenesResponse{}
	_body, _err := client.ListScenesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DropSubSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DropSubScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetHotspotTagWithOptions(request *GetHotspotTagRequest, runtime *util.RuntimeOptions) (_result *GetHotspotTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotspotTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotspotTag"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DropProjectWithOptions(request *DropProjectRequest, runtime *util.RuntimeOptions) (_result *DropProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DropProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DropProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListProjectWithOptions(request *ListProjectRequest, runtime *util.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetOriginLayoutDataWithOptions(request *GetOriginLayoutDataRequest, runtime *util.RuntimeOptions) (_result *GetOriginLayoutDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOriginLayoutDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOriginLayoutData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetHotspotSceneDataWithOptions(request *GetHotspotSceneDataRequest, runtime *util.RuntimeOptions) (_result *GetHotspotSceneDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotspotSceneDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotspotSceneData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) AddMosaicsWithOptions(request *AddMosaicsRequest, runtime *util.RuntimeOptions) (_result *AddMosaicsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddMosaicsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddMosaics"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ScenePublishWithOptions(request *ScenePublishRequest, runtime *util.RuntimeOptions) (_result *ScenePublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ScenePublishResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ScenePublish"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetRectifyImageWithOptions(request *GetRectifyImageRequest, runtime *util.RuntimeOptions) (_result *GetRectifyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRectifyImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRectifyImage"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetSubSceneTaskStatusWithOptions(request *GetSubSceneTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetSubSceneTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSubSceneTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSubSceneTaskStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PredictionWallLineWithOptions(request *PredictionWallLineRequest, runtime *util.RuntimeOptions) (_result *PredictionWallLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PredictionWallLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PredictionWallLine"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetScenePreviewInfoWithOptions(request *GetScenePreviewInfoRequest, runtime *util.RuntimeOptions) (_result *GetScenePreviewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetScenePreviewInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetScenePreviewInfo"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSubSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSubScene"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetLayoutDataWithOptions(request *GetLayoutDataRequest, runtime *util.RuntimeOptions) (_result *GetLayoutDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLayoutDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLayoutData"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
