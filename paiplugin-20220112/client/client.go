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

type CreateCampaignRequest struct {
	// 运营活动名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequest) SetName(v string) *CreateCampaignRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignRequest) SetRemark(v string) *CreateCampaignRequest {
	s.Remark = &v
	return s
}

type CreateCampaignResponseBody struct {
	// 返回数据。
	Data *CreateCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponseBody) SetData(v *CreateCampaignResponseBodyData) *CreateCampaignResponseBody {
	s.Data = v
	return s
}

func (s *CreateCampaignResponseBody) SetErrorCode(v int32) *CreateCampaignResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCampaignResponseBody) SetErrorMessage(v string) *CreateCampaignResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCampaignResponseBody) SetRequestId(v string) *CreateCampaignResponseBody {
	s.RequestId = &v
	return s
}

type CreateCampaignResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 运营活动Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 运营活动名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateCampaignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponseBodyData) SetCreatedTime(v string) *CreateCampaignResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateCampaignResponseBodyData) SetId(v string) *CreateCampaignResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateCampaignResponseBodyData) SetName(v string) *CreateCampaignResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCampaignResponseBodyData) SetRemark(v string) *CreateCampaignResponseBodyData {
	s.Remark = &v
	return s
}

func (s *CreateCampaignResponseBodyData) SetUpdatedTime(v string) *CreateCampaignResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignResponse) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponse) SetHeaders(v map[string]*string) *CreateCampaignResponse {
	s.Headers = v
	return s
}

func (s *CreateCampaignResponse) SetStatusCode(v int32) *CreateCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCampaignResponse) SetBody(v *CreateCampaignResponseBody) *CreateCampaignResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// 关联算法，人群来源为算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 手机号列名，人群来源为CSV文件，MaxCompute，并且包含手机号时需指定。
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// 过滤条件，人群来源为MaxCompute时可指定。
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// 预测任务Id，人群来源为算法。
	InferenceJobId *string `json:"InferenceJobId,omitempty" xml:"InferenceJobId,omitempty"`
	// 人群名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否包含手机号，包含手机号的人群可用于触达计划。
	PhoneNumber *bool `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// MaxCompute(ODPS)项目名，人群来源为MaxCompute时需指定。
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// 人群备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 人群来源。
	// - 0: 文本，每行一个手机号，最多100个。
	// - 1: 文本文件，每行一个手机号，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 2: CSV文件，需指定手机号列名，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 3: MaxCompute(ODPS)表，需指定手机号列名。
	// - 4: 算法。
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// MaxCompute(ODPS)表名，人群来源为MaxCompute时需指定。
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// 文本，人群来源为文本时需指定。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 文件地址，人群来源为文本文件，CSV文件，MaxCompute时需指定。
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetAlgorithm(v string) *CreateGroupRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateGroupRequest) SetColumn(v string) *CreateGroupRequest {
	s.Column = &v
	return s
}

func (s *CreateGroupRequest) SetFilter(v string) *CreateGroupRequest {
	s.Filter = &v
	return s
}

func (s *CreateGroupRequest) SetInferenceJobId(v string) *CreateGroupRequest {
	s.InferenceJobId = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetPhoneNumber(v bool) *CreateGroupRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateGroupRequest) SetProject(v string) *CreateGroupRequest {
	s.Project = &v
	return s
}

func (s *CreateGroupRequest) SetRemark(v string) *CreateGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateGroupRequest) SetSource(v int32) *CreateGroupRequest {
	s.Source = &v
	return s
}

func (s *CreateGroupRequest) SetTable(v string) *CreateGroupRequest {
	s.Table = &v
	return s
}

func (s *CreateGroupRequest) SetText(v string) *CreateGroupRequest {
	s.Text = &v
	return s
}

func (s *CreateGroupRequest) SetUri(v string) *CreateGroupRequest {
	s.Uri = &v
	return s
}

type CreateGroupResponseBody struct {
	// 返回数据。
	Data *CreateGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetData(v *CreateGroupResponseBodyData) *CreateGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateGroupResponseBody) SetErrorCode(v int32) *CreateGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateGroupResponseBody) SetErrorMessage(v string) *CreateGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponseBodyData struct {
	// 关联算法，人群来源为算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 人群数量。
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// 手机号列名，人群来源为CSV文件，MaxCompute，并且包含手机号时需指定。
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 过滤条件，人群来源为MaxCompute时可指定。
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// 人群Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测任务Id，人群来源为算法。
	InferenceJobId *string `json:"InferenceJobId,omitempty" xml:"InferenceJobId,omitempty"`
	// 人群名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否包含手机号，包含手机号的人群可用于触达计划。
	PhoneNumber *bool `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// MaxCompute(ODPS)项目名，人群来源为MaxCompute时需指定。
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// 人群备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 人群来源。
	// - 0: 文本，每行一个手机号，最多100个。
	// - 1: 文本文件，每行一个手机号，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 2: CSV文件，需指定手机号列名，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 3: MaxCompute(ODPS)表，需指定手机号列名。
	// - 4: 算法。
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// 人群状态。
	// - 0: 检查中。
	// - 1: 已通过。
	// - 2: 未通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// MaxCompute(ODPS)表名，人群来源为MaxCompute时需指定。
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// 文本，人群来源为文本时需指定。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 文件地址，人群来源为文本文件，CSV文件，MaxCompute时需指定。
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyData) SetAlgorithm(v string) *CreateGroupResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetAmount(v int32) *CreateGroupResponseBodyData {
	s.Amount = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetColumn(v string) *CreateGroupResponseBodyData {
	s.Column = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetCreatedTime(v string) *CreateGroupResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetFilter(v string) *CreateGroupResponseBodyData {
	s.Filter = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetId(v string) *CreateGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetInferenceJobId(v string) *CreateGroupResponseBodyData {
	s.InferenceJobId = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetName(v string) *CreateGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetPhoneNumber(v bool) *CreateGroupResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetProject(v string) *CreateGroupResponseBodyData {
	s.Project = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetRemark(v string) *CreateGroupResponseBodyData {
	s.Remark = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetSource(v int32) *CreateGroupResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetStatus(v int32) *CreateGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetTable(v string) *CreateGroupResponseBodyData {
	s.Table = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetText(v string) *CreateGroupResponseBodyData {
	s.Text = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetUpdatedTime(v string) *CreateGroupResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *CreateGroupResponseBodyData) SetUri(v string) *CreateGroupResponseBodyData {
	s.Uri = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateInferenceJobRequest struct {
	// 关联算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动Id。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 预测数据路径。
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// 预测任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 输出数据路径，需要为空目录。
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// 关联训练任务。
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s CreateInferenceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *CreateInferenceJobRequest) SetAlgorithm(v string) *CreateInferenceJobRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateInferenceJobRequest) SetCampaignId(v string) *CreateInferenceJobRequest {
	s.CampaignId = &v
	return s
}

func (s *CreateInferenceJobRequest) SetDataPath(v string) *CreateInferenceJobRequest {
	s.DataPath = &v
	return s
}

func (s *CreateInferenceJobRequest) SetName(v string) *CreateInferenceJobRequest {
	s.Name = &v
	return s
}

func (s *CreateInferenceJobRequest) SetRemark(v string) *CreateInferenceJobRequest {
	s.Remark = &v
	return s
}

func (s *CreateInferenceJobRequest) SetTargetPath(v string) *CreateInferenceJobRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateInferenceJobRequest) SetTrainingJobId(v string) *CreateInferenceJobRequest {
	s.TrainingJobId = &v
	return s
}

func (s *CreateInferenceJobRequest) SetUserConfig(v string) *CreateInferenceJobRequest {
	s.UserConfig = &v
	return s
}

type CreateInferenceJobResponseBody struct {
	// 返回数据。
	Data *CreateInferenceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInferenceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInferenceJobResponseBody) SetData(v *CreateInferenceJobResponseBodyData) *CreateInferenceJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateInferenceJobResponseBody) SetErrorCode(v int32) *CreateInferenceJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInferenceJobResponseBody) SetErrorMessage(v string) *CreateInferenceJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateInferenceJobResponseBody) SetRequestId(v string) *CreateInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateInferenceJobResponseBodyData struct {
	// 关联算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动Id。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 预测数据路径。
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// 关联人群Id，如果任务失败则人群无效。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 预测任务日志。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 预测任务Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 预测任务状态。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 输出数据路径，需要为空目录。
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// 关联训练任务。
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s CreateInferenceJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateInferenceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInferenceJobResponseBodyData) SetAlgorithm(v string) *CreateInferenceJobResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetCampaignId(v string) *CreateInferenceJobResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetCreatedTime(v string) *CreateInferenceJobResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetDataPath(v string) *CreateInferenceJobResponseBodyData {
	s.DataPath = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetGroupId(v string) *CreateInferenceJobResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetHistory(v string) *CreateInferenceJobResponseBodyData {
	s.History = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetId(v string) *CreateInferenceJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetName(v string) *CreateInferenceJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetRemark(v string) *CreateInferenceJobResponseBodyData {
	s.Remark = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetStatus(v int32) *CreateInferenceJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetTargetPath(v string) *CreateInferenceJobResponseBodyData {
	s.TargetPath = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetTrainingJobId(v string) *CreateInferenceJobResponseBodyData {
	s.TrainingJobId = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetUpdatedTime(v string) *CreateInferenceJobResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *CreateInferenceJobResponseBodyData) SetUserConfig(v string) *CreateInferenceJobResponseBodyData {
	s.UserConfig = &v
	return s
}

type CreateInferenceJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *CreateInferenceJobResponse) SetHeaders(v map[string]*string) *CreateInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *CreateInferenceJobResponse) SetStatusCode(v int32) *CreateInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInferenceJobResponse) SetBody(v *CreateInferenceJobResponseBody) *CreateInferenceJobResponse {
	s.Body = v
	return s
}

type CreateScheduleRequest struct {
	AISendEndDate   *string `json:"AISendEndDate,omitempty" xml:"AISendEndDate,omitempty"`
	AISendStartDate *string `json:"AISendStartDate,omitempty" xml:"AISendStartDate,omitempty"`
	// 终止时间（UTC+8）。
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 执行时间 (UTC+8)，为空立即执行。
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// 人群ID。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 触达计划名称。
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 重复周期，按重复周期与重复周期单位执行。
	RepeatCycle *int32 `json:"RepeatCycle,omitempty" xml:"RepeatCycle,omitempty"`
	// 重复周期单位，若指定执行时间，则重复周期生效。
	// - 0: 从不（默认）。
	// - 1: 小时。
	// - 2: 天。
	// - 3: 周。
	// - 4: 月。
	RepeatCycleUnit *int32 `json:"RepeatCycleUnit,omitempty" xml:"RepeatCycleUnit,omitempty"`
	// 重复次数。
	// - 0: 不设终止时间（默认）。
	// - N: 重复N次后终止。
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// 签名。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名ID，或指定签名。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板ID，或指定模板Code。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleRequest) SetAISendEndDate(v string) *CreateScheduleRequest {
	s.AISendEndDate = &v
	return s
}

func (s *CreateScheduleRequest) SetAISendStartDate(v string) *CreateScheduleRequest {
	s.AISendStartDate = &v
	return s
}

func (s *CreateScheduleRequest) SetEndTime(v int32) *CreateScheduleRequest {
	s.EndTime = &v
	return s
}

func (s *CreateScheduleRequest) SetExecuteTime(v string) *CreateScheduleRequest {
	s.ExecuteTime = &v
	return s
}

func (s *CreateScheduleRequest) SetGroupId(v string) *CreateScheduleRequest {
	s.GroupId = &v
	return s
}

func (s *CreateScheduleRequest) SetName(v string) *CreateScheduleRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduleRequest) SetPaymentType(v string) *CreateScheduleRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateScheduleRequest) SetRepeatCycle(v int32) *CreateScheduleRequest {
	s.RepeatCycle = &v
	return s
}

func (s *CreateScheduleRequest) SetRepeatCycleUnit(v int32) *CreateScheduleRequest {
	s.RepeatCycleUnit = &v
	return s
}

func (s *CreateScheduleRequest) SetRepeatTimes(v int32) *CreateScheduleRequest {
	s.RepeatTimes = &v
	return s
}

func (s *CreateScheduleRequest) SetSignName(v string) *CreateScheduleRequest {
	s.SignName = &v
	return s
}

func (s *CreateScheduleRequest) SetSignatureId(v string) *CreateScheduleRequest {
	s.SignatureId = &v
	return s
}

func (s *CreateScheduleRequest) SetTemplateCode(v string) *CreateScheduleRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateScheduleRequest) SetTemplateId(v string) *CreateScheduleRequest {
	s.TemplateId = &v
	return s
}

type CreateScheduleResponseBody struct {
	// 返回数据。
	Data *CreateScheduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBody) SetData(v *CreateScheduleResponseBodyData) *CreateScheduleResponseBody {
	s.Data = v
	return s
}

func (s *CreateScheduleResponseBody) SetErrorCode(v int32) *CreateScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateScheduleResponseBody) SetErrorMessage(v string) *CreateScheduleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateScheduleResponseBody) SetRequestId(v string) *CreateScheduleResponseBody {
	s.RequestId = &v
	return s
}

type CreateScheduleResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 终止时间（UTC+8）。
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 执行时间 (UTC+8)，为空立即执行。
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// 人群ID。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 触达计划ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 触达计划名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 重复周期，按重复周期与重复周期单位执行。
	RepeatCycle *int32 `json:"RepeatCycle,omitempty" xml:"RepeatCycle,omitempty"`
	// 重复周期单位，若指定执行时间，则重复周期生效。
	// - 0: 从不（默认）。
	// - 1: 小时。
	// - 2: 天。
	// - 3: 周。
	// - 4: 月。
	RepeatCycleUnit *int32 `json:"RepeatCycleUnit,omitempty" xml:"RepeatCycleUnit,omitempty"`
	// 重复次数。
	// - 0: 不设终止时间（默认）。
	// - N: 重复N次后终止。
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// 签名。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名ID，或指定签名。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 状态。
	// - 0: 检查中。
	// - 1: 检查成功。
	// - 2: 检查失败。
	// - 3: 发送中。
	// - 4: 发送成功。
	// - 5: 发送失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板ID，或指定模板Code。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBodyData) SetCreatedTime(v string) *CreateScheduleResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetEndTime(v int32) *CreateScheduleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetExecuteTime(v string) *CreateScheduleResponseBodyData {
	s.ExecuteTime = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetGroupId(v string) *CreateScheduleResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetId(v string) *CreateScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetName(v string) *CreateScheduleResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetRepeatCycle(v int32) *CreateScheduleResponseBodyData {
	s.RepeatCycle = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetRepeatCycleUnit(v int32) *CreateScheduleResponseBodyData {
	s.RepeatCycleUnit = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetRepeatTimes(v int32) *CreateScheduleResponseBodyData {
	s.RepeatTimes = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetSignName(v string) *CreateScheduleResponseBodyData {
	s.SignName = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetSignatureId(v string) *CreateScheduleResponseBodyData {
	s.SignatureId = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetStatus(v int32) *CreateScheduleResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetTemplateCode(v string) *CreateScheduleResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetTemplateId(v string) *CreateScheduleResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetUpdatedTime(v string) *CreateScheduleResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateScheduleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponse) SetHeaders(v map[string]*string) *CreateScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleResponse) SetStatusCode(v int32) *CreateScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleResponse) SetBody(v *CreateScheduleResponseBody) *CreateScheduleResponse {
	s.Body = v
	return s
}

type CreateSignatureRequest struct {
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureRequest) GoString() string {
	return s.String()
}

func (s *CreateSignatureRequest) SetDescription(v string) *CreateSignatureRequest {
	s.Description = &v
	return s
}

func (s *CreateSignatureRequest) SetName(v string) *CreateSignatureRequest {
	s.Name = &v
	return s
}

type CreateSignatureResponseBody struct {
	// 返回数据。
	Data *CreateSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBody) SetData(v *CreateSignatureResponseBodyData) *CreateSignatureResponseBody {
	s.Data = v
	return s
}

func (s *CreateSignatureResponseBody) SetErrorCode(v int32) *CreateSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSignatureResponseBody) SetErrorMessage(v string) *CreateSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSignatureResponseBody) SetRequestId(v string) *CreateSignatureResponseBody {
	s.RequestId = &v
	return s
}

type CreateSignatureResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 签名ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 签名审核状态。
	// - 0：审核中。
	// - 1：审核通过。
	// - 2：审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBodyData) SetCreatedTime(v string) *CreateSignatureResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetId(v string) *CreateSignatureResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetName(v string) *CreateSignatureResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetStatus(v int32) *CreateSignatureResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetUpdatedTime(v string) *CreateSignatureResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateSignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponse) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponse) SetHeaders(v map[string]*string) *CreateSignatureResponse {
	s.Headers = v
	return s
}

func (s *CreateSignatureResponse) SetStatusCode(v int32) *CreateSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSignatureResponse) SetBody(v *CreateSignatureResponseBody) *CreateSignatureResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// 模板内容，请注意控制总字数在70个字以内，超出部分按长短信收费，按67个字为单位记一条短信，营销短信必须在结尾添加“回T退订”。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 签名名称，同时只能指定签名名称或签名ID其中之一。
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 签名ID，可通过ListSignatures获取审核状态为已通过的签名列表，获取签名ID。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 模板类型。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetName(v string) *CreateTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplateRequest) SetSignature(v string) *CreateTemplateRequest {
	s.Signature = &v
	return s
}

func (s *CreateTemplateRequest) SetSignatureId(v string) *CreateTemplateRequest {
	s.SignatureId = &v
	return s
}

func (s *CreateTemplateRequest) SetType(v int32) *CreateTemplateRequest {
	s.Type = &v
	return s
}

type CreateTemplateResponseBody struct {
	// 返回数据。
	Data *CreateTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetData(v *CreateTemplateResponseBodyData) *CreateTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateTemplateResponseBody) SetErrorCode(v int32) *CreateTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTemplateResponseBody) SetErrorMessage(v string) *CreateTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateTemplateResponseBodyData struct {
	// 模板内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 模板名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核意见。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名。
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 签名ID。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 审核状态。
	// - 0 : 审核中。
	// - 1 : 审核通过。
	// - 2 : 审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板类型。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBodyData) SetContent(v string) *CreateTemplateResponseBodyData {
	s.Content = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetCreatedTime(v string) *CreateTemplateResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetDescription(v string) *CreateTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetId(v string) *CreateTemplateResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetName(v string) *CreateTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetReason(v string) *CreateTemplateResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetSignature(v string) *CreateTemplateResponseBodyData {
	s.Signature = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetSignatureId(v string) *CreateTemplateResponseBodyData {
	s.SignatureId = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetStatus(v int32) *CreateTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetTemplateCode(v string) *CreateTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetType(v int32) *CreateTemplateResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetUpdatedTime(v string) *CreateTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type CreateTrainingJobRequest struct {
	// 关联算法ID。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动ID。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 训练数据路径，指定路径前需确保已在控制台完成一键授权。
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// 训练任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s CreateTrainingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequest) SetAlgorithm(v string) *CreateTrainingJobRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateTrainingJobRequest) SetCampaignId(v string) *CreateTrainingJobRequest {
	s.CampaignId = &v
	return s
}

func (s *CreateTrainingJobRequest) SetDataPath(v string) *CreateTrainingJobRequest {
	s.DataPath = &v
	return s
}

func (s *CreateTrainingJobRequest) SetName(v string) *CreateTrainingJobRequest {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequest) SetRemark(v string) *CreateTrainingJobRequest {
	s.Remark = &v
	return s
}

func (s *CreateTrainingJobRequest) SetUserConfig(v string) *CreateTrainingJobRequest {
	s.UserConfig = &v
	return s
}

type CreateTrainingJobResponseBody struct {
	// 返回数据。
	Data *CreateTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponseBody) SetData(v *CreateTrainingJobResponseBodyData) *CreateTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateTrainingJobResponseBody) SetErrorCode(v int32) *CreateTrainingJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTrainingJobResponseBody) SetErrorMessage(v string) *CreateTrainingJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTrainingJobResponseBody) SetRequestId(v string) *CreateTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateTrainingJobResponseBodyData struct {
	// 关联算法ID。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动ID。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 训练数据路径，指定路径前需确保已在控制台完成一键授权。
	DataPath     *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	HasModelInfo *bool   `json:"HasModelInfo,omitempty" xml:"HasModelInfo,omitempty"`
	// 训练任务日志。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 训练任务ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 训练任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 训练任务状态。
	// - 0: 队列中。
	// - 1: 已提交。
	// - 2: 运行中。
	// - 3: 成功。
	// - 4: 失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 关联训练计划ID。
	TrainingScheduleId *string `json:"TrainingScheduleId,omitempty" xml:"TrainingScheduleId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s CreateTrainingJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponseBodyData) SetAlgorithm(v string) *CreateTrainingJobResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetCampaignId(v string) *CreateTrainingJobResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetCreatedTime(v string) *CreateTrainingJobResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetDataPath(v string) *CreateTrainingJobResponseBodyData {
	s.DataPath = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetHasModelInfo(v bool) *CreateTrainingJobResponseBodyData {
	s.HasModelInfo = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetHistory(v string) *CreateTrainingJobResponseBodyData {
	s.History = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetId(v string) *CreateTrainingJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetName(v string) *CreateTrainingJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetRemark(v string) *CreateTrainingJobResponseBodyData {
	s.Remark = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetStatus(v int32) *CreateTrainingJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetTrainingScheduleId(v string) *CreateTrainingJobResponseBodyData {
	s.TrainingScheduleId = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetUpdatedTime(v string) *CreateTrainingJobResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *CreateTrainingJobResponseBodyData) SetUserConfig(v string) *CreateTrainingJobResponseBodyData {
	s.UserConfig = &v
	return s
}

type CreateTrainingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponse) SetHeaders(v map[string]*string) *CreateTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainingJobResponse) SetStatusCode(v int32) *CreateTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainingJobResponse) SetBody(v *CreateTrainingJobResponseBody) *CreateTrainingJobResponse {
	s.Body = v
	return s
}

type DeleteCampaignResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCampaignResponseBody) SetData(v string) *DeleteCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCampaignResponseBody) SetErrorCode(v int32) *DeleteCampaignResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCampaignResponseBody) SetErrorMessage(v string) *DeleteCampaignResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCampaignResponseBody) SetRequestId(v string) *DeleteCampaignResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCampaignResponse) GoString() string {
	return s.String()
}

func (s *DeleteCampaignResponse) SetHeaders(v map[string]*string) *DeleteCampaignResponse {
	s.Headers = v
	return s
}

func (s *DeleteCampaignResponse) SetStatusCode(v int32) *DeleteCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCampaignResponse) SetBody(v *DeleteCampaignResponseBody) *DeleteCampaignResponse {
	s.Body = v
	return s
}

type DeleteGroupResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetData(v string) *DeleteGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGroupResponseBody) SetErrorCode(v int32) *DeleteGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGroupResponseBody) SetErrorMessage(v string) *DeleteGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteInferenceJobResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInferenceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInferenceJobResponseBody) SetData(v string) *DeleteInferenceJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInferenceJobResponseBody) SetErrorCode(v int32) *DeleteInferenceJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteInferenceJobResponseBody) SetErrorMessage(v string) *DeleteInferenceJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteInferenceJobResponseBody) SetRequestId(v string) *DeleteInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInferenceJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteInferenceJobResponse) SetHeaders(v map[string]*string) *DeleteInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteInferenceJobResponse) SetStatusCode(v int32) *DeleteInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInferenceJobResponse) SetBody(v *DeleteInferenceJobResponseBody) *DeleteInferenceJobResponse {
	s.Body = v
	return s
}

type DeleteScheduleResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponseBody) SetData(v string) *DeleteScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteScheduleResponseBody) SetErrorCode(v int32) *DeleteScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteScheduleResponseBody) SetErrorMessage(v string) *DeleteScheduleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteScheduleResponseBody) SetRequestId(v string) *DeleteScheduleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponse) SetHeaders(v map[string]*string) *DeleteScheduleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleResponse) SetStatusCode(v int32) *DeleteScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduleResponse) SetBody(v *DeleteScheduleResponseBody) *DeleteScheduleResponse {
	s.Body = v
	return s
}

type DeleteSignatureResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponseBody) SetData(v string) *DeleteSignatureResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSignatureResponseBody) SetErrorCode(v int32) *DeleteSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSignatureResponseBody) SetErrorMessage(v string) *DeleteSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSignatureResponseBody) SetRequestId(v string) *DeleteSignatureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponse) SetHeaders(v map[string]*string) *DeleteSignatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteSignatureResponse) SetStatusCode(v int32) *DeleteSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSignatureResponse) SetBody(v *DeleteSignatureResponseBody) *DeleteSignatureResponse {
	s.Body = v
	return s
}

type DeleteTemplateResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetData(v string) *DeleteTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetErrorCode(v int32) *DeleteTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetErrorMessage(v string) *DeleteTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteTrainingJobResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobResponseBody) SetData(v string) *DeleteTrainingJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTrainingJobResponseBody) SetErrorCode(v int32) *DeleteTrainingJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTrainingJobResponseBody) SetErrorMessage(v string) *DeleteTrainingJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTrainingJobResponseBody) SetRequestId(v string) *DeleteTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrainingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobResponse) SetHeaders(v map[string]*string) *DeleteTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrainingJobResponse) SetStatusCode(v int32) *DeleteTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrainingJobResponse) SetBody(v *DeleteTrainingJobResponseBody) *DeleteTrainingJobResponse {
	s.Body = v
	return s
}

type GetAlgorithmResponseBody struct {
	// 返回数据。
	Data *GetAlgorithmResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlgorithmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponseBody) SetData(v *GetAlgorithmResponseBodyData) *GetAlgorithmResponseBody {
	s.Data = v
	return s
}

func (s *GetAlgorithmResponseBody) SetErrorCode(v int32) *GetAlgorithmResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetErrorMessage(v string) *GetAlgorithmResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetRequestId(v string) *GetAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

type GetAlgorithmResponseBodyData struct {
	// 算法说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 算法Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测所需参数名与对应的参数说明。
	InferUserConfigMap *string `json:"InferUserConfigMap,omitempty" xml:"InferUserConfigMap,omitempty"`
	// 算法名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 训练所需参数名与对应的参数说明。
	TrainUserConfigMap *string `json:"TrainUserConfigMap,omitempty" xml:"TrainUserConfigMap,omitempty"`
}

func (s GetAlgorithmResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponseBodyData) SetDescription(v string) *GetAlgorithmResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAlgorithmResponseBodyData) SetId(v string) *GetAlgorithmResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAlgorithmResponseBodyData) SetInferUserConfigMap(v string) *GetAlgorithmResponseBodyData {
	s.InferUserConfigMap = &v
	return s
}

func (s *GetAlgorithmResponseBodyData) SetName(v string) *GetAlgorithmResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAlgorithmResponseBodyData) SetTrainUserConfigMap(v string) *GetAlgorithmResponseBodyData {
	s.TrainUserConfigMap = &v
	return s
}

type GetAlgorithmResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponse) SetHeaders(v map[string]*string) *GetAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmResponse) SetStatusCode(v int32) *GetAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmResponse) SetBody(v *GetAlgorithmResponseBody) *GetAlgorithmResponse {
	s.Body = v
	return s
}

type GetCampaignResponseBody struct {
	// 返回数据。
	Data *GetCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBody) SetData(v *GetCampaignResponseBodyData) *GetCampaignResponseBody {
	s.Data = v
	return s
}

func (s *GetCampaignResponseBody) SetErrorCode(v int32) *GetCampaignResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCampaignResponseBody) SetErrorMessage(v string) *GetCampaignResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCampaignResponseBody) SetRequestId(v string) *GetCampaignResponseBody {
	s.RequestId = &v
	return s
}

type GetCampaignResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 运营活动Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 运营活动名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetCampaignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBodyData) SetCreatedTime(v string) *GetCampaignResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetId(v string) *GetCampaignResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetName(v string) *GetCampaignResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetRemark(v string) *GetCampaignResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetUpdatedTime(v string) *GetCampaignResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type GetCampaignResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignResponse) GoString() string {
	return s.String()
}

func (s *GetCampaignResponse) SetHeaders(v map[string]*string) *GetCampaignResponse {
	s.Headers = v
	return s
}

func (s *GetCampaignResponse) SetStatusCode(v int32) *GetCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCampaignResponse) SetBody(v *GetCampaignResponseBody) *GetCampaignResponse {
	s.Body = v
	return s
}

type GetGroupResponseBody struct {
	// 返回数据。
	Data *GetGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) SetData(v *GetGroupResponseBodyData) *GetGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetGroupResponseBody) SetErrorCode(v int32) *GetGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGroupResponseBody) SetErrorMessage(v string) *GetGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupResponseBodyData struct {
	// 关联算法，人群来源为算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 人群数量。
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// 关联运营活动Id。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 手机号列名，人群来源为CSV文件，MaxCompute，并且包含手机号时需指定。
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 过滤条件，人群来源为MaxCompute时可指定。
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// 历史记录。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 人群Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测任务Id，人群来源为算法。
	InferenceJobId *string `json:"InferenceJobId,omitempty" xml:"InferenceJobId,omitempty"`
	// 人群名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否包含手机号，包含手机号的人群可用于触达计划。
	PhoneNumber *bool `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// MaxCompute(ODPS)项目名，人群来源为MaxCompute时需指定。
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// 人群备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 人群来源。
	// - 0: 文本，每行一个手机号，最多100个。
	// - 1: 文本文件，每行一个手机号，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 2: 多列CSV文件，需指定手机号列名，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 3: MaxCompute表，需指定手机号列名。
	// - 4: 算法。
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// 人群状态。
	// - 0: 检查中。
	// - 1: 已通过。
	// - 2: 未通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// MaxCompute(ODPS)表名，人群来源为MaxCompute时需指定。
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// 文本，人群来源为文本时需指定。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 文件地址，人群来源为文本文件，CSV文件时需指定。
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyData) SetAlgorithm(v string) *GetGroupResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetGroupResponseBodyData) SetAmount(v int32) *GetGroupResponseBodyData {
	s.Amount = &v
	return s
}

func (s *GetGroupResponseBodyData) SetCampaignId(v string) *GetGroupResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *GetGroupResponseBodyData) SetColumn(v string) *GetGroupResponseBodyData {
	s.Column = &v
	return s
}

func (s *GetGroupResponseBodyData) SetCreatedTime(v string) *GetGroupResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetGroupResponseBodyData) SetFilter(v string) *GetGroupResponseBodyData {
	s.Filter = &v
	return s
}

func (s *GetGroupResponseBodyData) SetHistory(v string) *GetGroupResponseBodyData {
	s.History = &v
	return s
}

func (s *GetGroupResponseBodyData) SetId(v string) *GetGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGroupResponseBodyData) SetInferenceJobId(v string) *GetGroupResponseBodyData {
	s.InferenceJobId = &v
	return s
}

func (s *GetGroupResponseBodyData) SetName(v string) *GetGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGroupResponseBodyData) SetPhoneNumber(v bool) *GetGroupResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetGroupResponseBodyData) SetProject(v string) *GetGroupResponseBodyData {
	s.Project = &v
	return s
}

func (s *GetGroupResponseBodyData) SetRemark(v string) *GetGroupResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetGroupResponseBodyData) SetSource(v int32) *GetGroupResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetGroupResponseBodyData) SetStatus(v int32) *GetGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGroupResponseBodyData) SetTable(v string) *GetGroupResponseBodyData {
	s.Table = &v
	return s
}

func (s *GetGroupResponseBodyData) SetText(v string) *GetGroupResponseBodyData {
	s.Text = &v
	return s
}

func (s *GetGroupResponseBodyData) SetUpdatedTime(v string) *GetGroupResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetGroupResponseBodyData) SetUri(v string) *GetGroupResponseBodyData {
	s.Uri = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetInferenceJobResponseBody struct {
	// 返回数据。
	Data *GetInferenceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInferenceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetInferenceJobResponseBody) SetData(v *GetInferenceJobResponseBodyData) *GetInferenceJobResponseBody {
	s.Data = v
	return s
}

func (s *GetInferenceJobResponseBody) SetErrorCode(v int32) *GetInferenceJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInferenceJobResponseBody) SetErrorMessage(v string) *GetInferenceJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInferenceJobResponseBody) SetRequestId(v string) *GetInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

type GetInferenceJobResponseBodyData struct {
	// 关联算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动Id。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 预测数据路径。
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// 关联人群Id，如果任务失败则人群无效。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 预测任务日志。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 预测任务Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 预测任务状态。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 输出数据路径，需要为空目录。
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// 关联训练任务。
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s GetInferenceJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInferenceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInferenceJobResponseBodyData) SetAlgorithm(v string) *GetInferenceJobResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetCampaignId(v string) *GetInferenceJobResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetCreatedTime(v string) *GetInferenceJobResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetDataPath(v string) *GetInferenceJobResponseBodyData {
	s.DataPath = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetGroupId(v string) *GetInferenceJobResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetHistory(v string) *GetInferenceJobResponseBodyData {
	s.History = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetId(v string) *GetInferenceJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetName(v string) *GetInferenceJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetRemark(v string) *GetInferenceJobResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetStatus(v int32) *GetInferenceJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetTargetPath(v string) *GetInferenceJobResponseBodyData {
	s.TargetPath = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetTrainingJobId(v string) *GetInferenceJobResponseBodyData {
	s.TrainingJobId = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetUpdatedTime(v string) *GetInferenceJobResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetInferenceJobResponseBodyData) SetUserConfig(v string) *GetInferenceJobResponseBodyData {
	s.UserConfig = &v
	return s
}

type GetInferenceJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *GetInferenceJobResponse) SetHeaders(v map[string]*string) *GetInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *GetInferenceJobResponse) SetStatusCode(v int32) *GetInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInferenceJobResponse) SetBody(v *GetInferenceJobResponseBody) *GetInferenceJobResponse {
	s.Body = v
	return s
}

type GetMessageConfigResponseBody struct {
	// 返回数据。
	Data *GetMessageConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageConfigResponseBody) SetData(v *GetMessageConfigResponseBodyData) *GetMessageConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetMessageConfigResponseBody) SetErrorCode(v int32) *GetMessageConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMessageConfigResponseBody) SetErrorMessage(v string) *GetMessageConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMessageConfigResponseBody) SetRequestId(v string) *GetMessageConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetMessageConfigResponseBodyData struct {
	// 短信发送状态回执接收服务地址。
	SmsReportUrl *string `json:"SmsReportUrl,omitempty" xml:"SmsReportUrl,omitempty"`
	// 上行短信消息接收服务地址。
	SmsUpUrl *string `json:"SmsUpUrl,omitempty" xml:"SmsUpUrl,omitempty"`
}

func (s GetMessageConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMessageConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessageConfigResponseBodyData) SetSmsReportUrl(v string) *GetMessageConfigResponseBodyData {
	s.SmsReportUrl = &v
	return s
}

func (s *GetMessageConfigResponseBodyData) SetSmsUpUrl(v string) *GetMessageConfigResponseBodyData {
	s.SmsUpUrl = &v
	return s
}

type GetMessageConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMessageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMessageConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageConfigResponse) GoString() string {
	return s.String()
}

func (s *GetMessageConfigResponse) SetHeaders(v map[string]*string) *GetMessageConfigResponse {
	s.Headers = v
	return s
}

func (s *GetMessageConfigResponse) SetStatusCode(v int32) *GetMessageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageConfigResponse) SetBody(v *GetMessageConfigResponseBody) *GetMessageConfigResponse {
	s.Body = v
	return s
}

type GetScheduleResponseBody struct {
	// 返回数据。
	Data *GetScheduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBody) SetData(v *GetScheduleResponseBodyData) *GetScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetScheduleResponseBody) SetErrorCode(v int32) *GetScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetScheduleResponseBody) SetErrorMessage(v string) *GetScheduleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetScheduleResponseBody) SetRequestId(v string) *GetScheduleResponseBody {
	s.RequestId = &v
	return s
}

type GetScheduleResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 终止时间（UTC+8）。
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 执行时间 (UTC+8)，为空立即执行。
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// 人群ID。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 历史记录。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 触达计划ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 触达计划名称。
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 重复周期，按重复周期与重复周期单位执行。
	RepeatCycle *int32 `json:"RepeatCycle,omitempty" xml:"RepeatCycle,omitempty"`
	// 重复周期单位，若指定执行时间，则重复周期生效。
	// - 0: 从不（默认）。
	// - 1: 小时。
	// - 2: 天。
	// - 3: 周。
	// - 4: 月。
	RepeatCycleUnit *int32 `json:"RepeatCycleUnit,omitempty" xml:"RepeatCycleUnit,omitempty"`
	// 重复次数。
	// - 0: 不设终止时间（默认）。
	// - N: 重复N次后终止。
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// 签名。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名ID，或指定签名。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 状态。
	// - 0: 检查中。
	// - 1: 检查成功。
	// - 2: 检查失败。
	// - 3: 发送中。
	// - 4: 发送成功。
	// - 5: 发送失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板ID，或指定模板Code。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyData) SetCreatedTime(v string) *GetScheduleResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetEndTime(v int32) *GetScheduleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetExecuteTime(v string) *GetScheduleResponseBodyData {
	s.ExecuteTime = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetGroupId(v string) *GetScheduleResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetHistory(v string) *GetScheduleResponseBodyData {
	s.History = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetId(v string) *GetScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetName(v string) *GetScheduleResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetPaymentType(v string) *GetScheduleResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetRepeatCycle(v int32) *GetScheduleResponseBodyData {
	s.RepeatCycle = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetRepeatCycleUnit(v int32) *GetScheduleResponseBodyData {
	s.RepeatCycleUnit = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetRepeatTimes(v int32) *GetScheduleResponseBodyData {
	s.RepeatTimes = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetSignName(v string) *GetScheduleResponseBodyData {
	s.SignName = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetSignatureId(v string) *GetScheduleResponseBodyData {
	s.SignatureId = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetStatus(v int32) *GetScheduleResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetTemplateCode(v string) *GetScheduleResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetTemplateId(v string) *GetScheduleResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetScheduleResponseBodyData) SetUpdatedTime(v string) *GetScheduleResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type GetScheduleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleResponse) SetHeaders(v map[string]*string) *GetScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleResponse) SetStatusCode(v int32) *GetScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduleResponse) SetBody(v *GetScheduleResponseBody) *GetScheduleResponse {
	s.Body = v
	return s
}

type GetSignatureResponseBody struct {
	// 返回数据。
	Data *GetSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignatureResponseBody) SetData(v *GetSignatureResponseBodyData) *GetSignatureResponseBody {
	s.Data = v
	return s
}

func (s *GetSignatureResponseBody) SetErrorCode(v int32) *GetSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSignatureResponseBody) SetErrorMessage(v string) *GetSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSignatureResponseBody) SetRequestId(v string) *GetSignatureResponseBody {
	s.RequestId = &v
	return s
}

type GetSignatureResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 签名Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核建议。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名审核状态。
	// - 0：审核中。
	// - 1：审核通过。
	// - 2：审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSignatureResponseBodyData) SetCreatedTime(v string) *GetSignatureResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetDescription(v string) *GetSignatureResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetId(v string) *GetSignatureResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetName(v string) *GetSignatureResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetReason(v string) *GetSignatureResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetStatus(v int32) *GetSignatureResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetUpdatedTime(v string) *GetSignatureResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type GetSignatureResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponse) GoString() string {
	return s.String()
}

func (s *GetSignatureResponse) SetHeaders(v map[string]*string) *GetSignatureResponse {
	s.Headers = v
	return s
}

func (s *GetSignatureResponse) SetStatusCode(v int32) *GetSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSignatureResponse) SetBody(v *GetSignatureResponseBody) *GetSignatureResponse {
	s.Body = v
	return s
}

type GetTemplateResponseBody struct {
	// 返回数据。
	Data *GetTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetErrorCode(v int32) *GetTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetErrorMessage(v string) *GetTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

type GetTemplateResponseBodyData struct {
	// 模板内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核意见。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名Id。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 审核状态。
	// - 0 : 审核中。
	// - 1 : 审核通过。
	// - 2 : 审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板类型。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyData) SetContent(v string) *GetTemplateResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetCreatedTime(v string) *GetTemplateResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetDescription(v string) *GetTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetId(v string) *GetTemplateResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetName(v string) *GetTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetReason(v string) *GetTemplateResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetSignatureId(v string) *GetTemplateResponseBodyData {
	s.SignatureId = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetStatus(v int32) *GetTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetTemplateCode(v string) *GetTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetType(v int32) *GetTemplateResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetUpdatedTime(v string) *GetTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetTrainingJobResponseBody struct {
	// 返回数据。
	Data *GetTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBody) SetData(v *GetTrainingJobResponseBodyData) *GetTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *GetTrainingJobResponseBody) SetErrorCode(v int32) *GetTrainingJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetErrorMessage(v string) *GetTrainingJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetRequestId(v string) *GetTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

type GetTrainingJobResponseBodyData struct {
	// 关联算法ID。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动ID。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 训练数据路径，指定路径前需确保已在控制台完成一键授权。
	DataPath     *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	HasModelInfo *bool   `json:"HasModelInfo,omitempty" xml:"HasModelInfo,omitempty"`
	// 训练任务日志。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 训练任务ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 训练任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 训练任务状态。
	// - 0: 队列中。
	// - 1: 已提交。
	// - 2: 运行中。
	// - 3: 成功。
	// - 4: 失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 关联训练计划ID。
	TrainingScheduleId *string `json:"TrainingScheduleId,omitempty" xml:"TrainingScheduleId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s GetTrainingJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyData) SetAlgorithm(v string) *GetTrainingJobResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetCampaignId(v string) *GetTrainingJobResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetCreatedTime(v string) *GetTrainingJobResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetDataPath(v string) *GetTrainingJobResponseBodyData {
	s.DataPath = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetHasModelInfo(v bool) *GetTrainingJobResponseBodyData {
	s.HasModelInfo = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetHistory(v string) *GetTrainingJobResponseBodyData {
	s.History = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetId(v string) *GetTrainingJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetName(v string) *GetTrainingJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetRemark(v string) *GetTrainingJobResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetStatus(v int32) *GetTrainingJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetTrainingScheduleId(v string) *GetTrainingJobResponseBodyData {
	s.TrainingScheduleId = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetUpdatedTime(v string) *GetTrainingJobResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetTrainingJobResponseBodyData) SetUserConfig(v string) *GetTrainingJobResponseBodyData {
	s.UserConfig = &v
	return s
}

type GetTrainingJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrainingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponse) SetHeaders(v map[string]*string) *GetTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobResponse) SetStatusCode(v int32) *GetTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobResponse) SetBody(v *GetTrainingJobResponseBody) *GetTrainingJobResponse {
	s.Body = v
	return s
}

type GetUserResponseBody struct {
	// 返回数据。
	Data *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetErrorCode(v int32) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetErrorMessage(v string) *GetUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResponseBodyData struct {
	// 账号状态。
	// - 0 : 正常使用。
	// - 1 : 因欠费等原因暂时停用。
	// - 2 : 已释放产品。
	AccountStatus *int32 `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetAccountStatus(v int32) *GetUserResponseBodyData {
	s.AccountStatus = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListAlgorithmsRequest struct {
	// 算法Id过滤。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 算法名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlgorithmsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsRequest) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsRequest) SetId(v string) *ListAlgorithmsRequest {
	s.Id = &v
	return s
}

func (s *ListAlgorithmsRequest) SetName(v string) *ListAlgorithmsRequest {
	s.Name = &v
	return s
}

func (s *ListAlgorithmsRequest) SetPageNumber(v int32) *ListAlgorithmsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlgorithmsRequest) SetPageSize(v int32) *ListAlgorithmsRequest {
	s.PageSize = &v
	return s
}

type ListAlgorithmsResponseBody struct {
	// 返回数据。
	Data *ListAlgorithmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlgorithmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBody) SetData(v *ListAlgorithmsResponseBodyData) *ListAlgorithmsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlgorithmsResponseBody) SetErrorCode(v int32) *ListAlgorithmsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAlgorithmsResponseBody) SetErrorMessage(v string) *ListAlgorithmsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAlgorithmsResponseBody) SetRequestId(v string) *ListAlgorithmsResponseBody {
	s.RequestId = &v
	return s
}

type ListAlgorithmsResponseBodyData struct {
	// 算法列表。
	Algorithms []*ListAlgorithmsResponseBodyDataAlgorithms `json:"Algorithms,omitempty" xml:"Algorithms,omitempty" type:"Repeated"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总算法数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlgorithmsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBodyData) SetAlgorithms(v []*ListAlgorithmsResponseBodyDataAlgorithms) *ListAlgorithmsResponseBodyData {
	s.Algorithms = v
	return s
}

func (s *ListAlgorithmsResponseBodyData) SetPageNumber(v int32) *ListAlgorithmsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAlgorithmsResponseBodyData) SetPageSize(v int32) *ListAlgorithmsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAlgorithmsResponseBodyData) SetTotalCount(v int32) *ListAlgorithmsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListAlgorithmsResponseBodyDataAlgorithms struct {
	// 算法Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 算法名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAlgorithmsResponseBodyDataAlgorithms) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponseBodyDataAlgorithms) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBodyDataAlgorithms) SetId(v string) *ListAlgorithmsResponseBodyDataAlgorithms {
	s.Id = &v
	return s
}

func (s *ListAlgorithmsResponseBodyDataAlgorithms) SetName(v string) *ListAlgorithmsResponseBodyDataAlgorithms {
	s.Name = &v
	return s
}

type ListAlgorithmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlgorithmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlgorithmsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlgorithmsResponse) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponse) SetHeaders(v map[string]*string) *ListAlgorithmsResponse {
	s.Headers = v
	return s
}

func (s *ListAlgorithmsResponse) SetStatusCode(v int32) *ListAlgorithmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlgorithmsResponse) SetBody(v *ListAlgorithmsResponseBody) *ListAlgorithmsResponse {
	s.Body = v
	return s
}

type ListCampaignsRequest struct {
	// 运营活动名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 运营活动备注过滤。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ListCampaignsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsRequest) GoString() string {
	return s.String()
}

func (s *ListCampaignsRequest) SetName(v string) *ListCampaignsRequest {
	s.Name = &v
	return s
}

func (s *ListCampaignsRequest) SetPageNumber(v int32) *ListCampaignsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsRequest) SetPageSize(v int32) *ListCampaignsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsRequest) SetRemark(v string) *ListCampaignsRequest {
	s.Remark = &v
	return s
}

type ListCampaignsResponseBody struct {
	// 返回数据。
	Data *ListCampaignsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCampaignsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBody) SetData(v *ListCampaignsResponseBodyData) *ListCampaignsResponseBody {
	s.Data = v
	return s
}

func (s *ListCampaignsResponseBody) SetErrorCode(v int32) *ListCampaignsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCampaignsResponseBody) SetErrorMessage(v string) *ListCampaignsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCampaignsResponseBody) SetRequestId(v string) *ListCampaignsResponseBody {
	s.RequestId = &v
	return s
}

type ListCampaignsResponseBodyData struct {
	// 运营活动列表。
	Campaigns []*ListCampaignsResponseBodyDataCampaigns `json:"Campaigns,omitempty" xml:"Campaigns,omitempty" type:"Repeated"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总运营活动数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCampaignsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyData) SetCampaigns(v []*ListCampaignsResponseBodyDataCampaigns) *ListCampaignsResponseBodyData {
	s.Campaigns = v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageNumber(v int32) *ListCampaignsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageSize(v int32) *ListCampaignsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetTotalCount(v int32) *ListCampaignsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCampaignsResponseBodyDataCampaigns struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 运营活动Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 运营活动名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListCampaignsResponseBodyDataCampaigns) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponseBodyDataCampaigns) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyDataCampaigns) SetCreatedTime(v string) *ListCampaignsResponseBodyDataCampaigns {
	s.CreatedTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataCampaigns) SetId(v string) *ListCampaignsResponseBodyDataCampaigns {
	s.Id = &v
	return s
}

func (s *ListCampaignsResponseBodyDataCampaigns) SetName(v string) *ListCampaignsResponseBodyDataCampaigns {
	s.Name = &v
	return s
}

func (s *ListCampaignsResponseBodyDataCampaigns) SetRemark(v string) *ListCampaignsResponseBodyDataCampaigns {
	s.Remark = &v
	return s
}

func (s *ListCampaignsResponseBodyDataCampaigns) SetUpdatedTime(v string) *ListCampaignsResponseBodyDataCampaigns {
	s.UpdatedTime = &v
	return s
}

type ListCampaignsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCampaignsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCampaignsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponse) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponse) SetHeaders(v map[string]*string) *ListCampaignsResponse {
	s.Headers = v
	return s
}

func (s *ListCampaignsResponse) SetStatusCode(v int32) *ListCampaignsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCampaignsResponse) SetBody(v *ListCampaignsResponseBody) *ListCampaignsResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// 人群名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 是否包含手机号过滤。
	PhoneNumber *bool `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 人群备注过滤。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 来源过滤。
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// 审核状态过滤。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetName(v string) *ListGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListGroupsRequest) SetPageNumber(v int32) *ListGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v int32) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsRequest) SetPhoneNumber(v bool) *ListGroupsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListGroupsRequest) SetRemark(v string) *ListGroupsRequest {
	s.Remark = &v
	return s
}

func (s *ListGroupsRequest) SetSource(v int32) *ListGroupsRequest {
	s.Source = &v
	return s
}

func (s *ListGroupsRequest) SetStatus(v int32) *ListGroupsRequest {
	s.Status = &v
	return s
}

type ListGroupsResponseBody struct {
	// 返回数据。
	Data *ListGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetData(v *ListGroupsResponseBodyData) *ListGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsResponseBody) SetErrorCode(v int32) *ListGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupsResponseBody) SetErrorMessage(v string) *ListGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupsResponseBodyData struct {
	// 人群列表。
	Groups []*ListGroupsResponseBodyDataGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总人群数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyData) SetGroups(v []*ListGroupsResponseBodyDataGroups) *ListGroupsResponseBodyData {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBodyData) SetPageNumber(v int32) *ListGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetPageSize(v int32) *ListGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetTotalCount(v int32) *ListGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGroupsResponseBodyDataGroups struct {
	// 关联算法，人群来源为算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 人群数量。
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// 手机号列名，人群来源为CSV文件，MaxCompute，并且包含手机号时需指定。
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 过滤条件，人群来源为MaxCompute时可指定。
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// 人群Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测任务Id，人群来源为算法。
	InferenceJobId *string `json:"InferenceJobId,omitempty" xml:"InferenceJobId,omitempty"`
	// 人群名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否包含手机号，包含手机号的人群可用于触达计划。
	PhoneNumber *bool `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// MaxCompute(ODPS)项目名，人群来源为MaxCompute时需指定。
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// 人群备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 人群来源。
	// - 0: 文本，每行一个手机号，最多100个。
	// - 1: 文本文件，每行一个手机号，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 2: CSV文件，需指定手机号列名，可通过控制台上传或指定自定义OSS地址，指定自定义OSS地址前需确保已在控制台完成一键授权。
	// - 3: MaxCompute(ODPS)表，需指定手机号列名。
	// - 4: 算法。
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// 人群状态。
	// - 0: 检查中。
	// - 1: 已通过。
	// - 2: 未通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// MaxCompute(ODPS)表名，人群来源为MaxCompute时需指定。
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// 文本，人群来源为文本时需指定。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 文件地址，人群来源为文本文件，CSV文件，MaxCompute时需指定。
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListGroupsResponseBodyDataGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyDataGroups) SetAlgorithm(v string) *ListGroupsResponseBodyDataGroups {
	s.Algorithm = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetAmount(v int32) *ListGroupsResponseBodyDataGroups {
	s.Amount = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetColumn(v string) *ListGroupsResponseBodyDataGroups {
	s.Column = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetCreatedTime(v string) *ListGroupsResponseBodyDataGroups {
	s.CreatedTime = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetFilter(v string) *ListGroupsResponseBodyDataGroups {
	s.Filter = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetId(v string) *ListGroupsResponseBodyDataGroups {
	s.Id = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetInferenceJobId(v string) *ListGroupsResponseBodyDataGroups {
	s.InferenceJobId = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetName(v string) *ListGroupsResponseBodyDataGroups {
	s.Name = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetPhoneNumber(v bool) *ListGroupsResponseBodyDataGroups {
	s.PhoneNumber = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetProject(v string) *ListGroupsResponseBodyDataGroups {
	s.Project = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetRemark(v string) *ListGroupsResponseBodyDataGroups {
	s.Remark = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetSource(v int32) *ListGroupsResponseBodyDataGroups {
	s.Source = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetStatus(v int32) *ListGroupsResponseBodyDataGroups {
	s.Status = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetTable(v string) *ListGroupsResponseBodyDataGroups {
	s.Table = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetText(v string) *ListGroupsResponseBodyDataGroups {
	s.Text = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetUpdatedTime(v string) *ListGroupsResponseBodyDataGroups {
	s.UpdatedTime = &v
	return s
}

func (s *ListGroupsResponseBodyDataGroups) SetUri(v string) *ListGroupsResponseBodyDataGroups {
	s.Uri = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListInferenceJobsRequest struct {
	// 归属运营活动过滤。
	CampaignId   *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CampaignName *string `json:"CampaignName,omitempty" xml:"CampaignName,omitempty"`
	// 预测任务名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 预测任务备注过滤。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 预测任务状态过滤。
	// - 0: 队列中。
	// - 1: 已提交。
	// - 2: 运行中。
	// - 3: 成功。
	// - 4: 失败。
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
}

func (s ListInferenceJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInferenceJobsRequest) GoString() string {
	return s.String()
}

func (s *ListInferenceJobsRequest) SetCampaignId(v string) *ListInferenceJobsRequest {
	s.CampaignId = &v
	return s
}

func (s *ListInferenceJobsRequest) SetCampaignName(v string) *ListInferenceJobsRequest {
	s.CampaignName = &v
	return s
}

func (s *ListInferenceJobsRequest) SetName(v string) *ListInferenceJobsRequest {
	s.Name = &v
	return s
}

func (s *ListInferenceJobsRequest) SetPageNumber(v int32) *ListInferenceJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInferenceJobsRequest) SetPageSize(v int32) *ListInferenceJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInferenceJobsRequest) SetRemark(v string) *ListInferenceJobsRequest {
	s.Remark = &v
	return s
}

func (s *ListInferenceJobsRequest) SetStatus(v int32) *ListInferenceJobsRequest {
	s.Status = &v
	return s
}

func (s *ListInferenceJobsRequest) SetTrainingJobName(v string) *ListInferenceJobsRequest {
	s.TrainingJobName = &v
	return s
}

type ListInferenceJobsResponseBody struct {
	// 返回数据。
	Data *ListInferenceJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInferenceJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInferenceJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInferenceJobsResponseBody) SetData(v *ListInferenceJobsResponseBodyData) *ListInferenceJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListInferenceJobsResponseBody) SetErrorCode(v int32) *ListInferenceJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInferenceJobsResponseBody) SetErrorMessage(v string) *ListInferenceJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInferenceJobsResponseBody) SetRequestId(v string) *ListInferenceJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListInferenceJobsResponseBodyData struct {
	// 预测任务列表。
	InferenceJobs []*ListInferenceJobsResponseBodyDataInferenceJobs `json:"InferenceJobs,omitempty" xml:"InferenceJobs,omitempty" type:"Repeated"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总预测任务数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInferenceJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInferenceJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInferenceJobsResponseBodyData) SetInferenceJobs(v []*ListInferenceJobsResponseBodyDataInferenceJobs) *ListInferenceJobsResponseBodyData {
	s.InferenceJobs = v
	return s
}

func (s *ListInferenceJobsResponseBodyData) SetPageNumber(v int32) *ListInferenceJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInferenceJobsResponseBodyData) SetPageSize(v int32) *ListInferenceJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInferenceJobsResponseBodyData) SetTotalCount(v int32) *ListInferenceJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListInferenceJobsResponseBodyDataInferenceJobs struct {
	// 关联算法。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动ID。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 预测数据路径，当预测人群数据分布在多个csv文件时可指定目录，指定路径前需确保已在控制台完成一键授权。
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// 预测人群，人群来源必须为多列csv，当同时指定DataPath与GroupId时，以GroupId为准。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 预测任务日志。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 预测任务ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 预测任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 预测任务状态。
	// - 0: 队列中。
	// - 1: 已提交。
	// - 2: 运行中。
	// - 3: 成功。
	// - 4: 失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 关联目标人群ID，如果任务失败则人群无效。
	TargetGroupId *string `json:"TargetGroupId,omitempty" xml:"TargetGroupId,omitempty"`
	// 输出数据路径，需要为空目录，指定路径前需确保已在控制台完成一键授权。
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// 关联训练任务。
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s ListInferenceJobsResponseBodyDataInferenceJobs) String() string {
	return tea.Prettify(s)
}

func (s ListInferenceJobsResponseBodyDataInferenceJobs) GoString() string {
	return s.String()
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetAlgorithm(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.Algorithm = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetCampaignId(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.CampaignId = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetCreatedTime(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.CreatedTime = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetDataPath(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.DataPath = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetGroupId(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.GroupId = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetHistory(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.History = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetId(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.Id = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetName(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.Name = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetRemark(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.Remark = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetStatus(v int32) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.Status = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetTargetGroupId(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.TargetGroupId = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetTargetPath(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.TargetPath = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetTrainingJobId(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.TrainingJobId = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetUpdatedTime(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.UpdatedTime = &v
	return s
}

func (s *ListInferenceJobsResponseBodyDataInferenceJobs) SetUserConfig(v string) *ListInferenceJobsResponseBodyDataInferenceJobs {
	s.UserConfig = &v
	return s
}

type ListInferenceJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInferenceJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInferenceJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInferenceJobsResponse) GoString() string {
	return s.String()
}

func (s *ListInferenceJobsResponse) SetHeaders(v map[string]*string) *ListInferenceJobsResponse {
	s.Headers = v
	return s
}

func (s *ListInferenceJobsResponse) SetStatusCode(v int32) *ListInferenceJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInferenceJobsResponse) SetBody(v *ListInferenceJobsResponseBody) *ListInferenceJobsResponse {
	s.Body = v
	return s
}

type ListMessageMetricsRequest struct {
	// 结束日期，格式20220102。
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// 关联人群Id。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 关联触达计划Id。
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// 签名名称。
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 签名Id，同时只能指定签名名称或签名Id其中之一。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 开始日期，格式20220102。
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// 模板号。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板Id，同时只能指定模板Code或模板Id其中之一。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板类型。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListMessageMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessageMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListMessageMetricsRequest) SetEndDate(v string) *ListMessageMetricsRequest {
	s.EndDate = &v
	return s
}

func (s *ListMessageMetricsRequest) SetGroupId(v string) *ListMessageMetricsRequest {
	s.GroupId = &v
	return s
}

func (s *ListMessageMetricsRequest) SetPageNumber(v int32) *ListMessageMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessageMetricsRequest) SetPageSize(v int32) *ListMessageMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageMetricsRequest) SetScheduleId(v string) *ListMessageMetricsRequest {
	s.ScheduleId = &v
	return s
}

func (s *ListMessageMetricsRequest) SetSignature(v string) *ListMessageMetricsRequest {
	s.Signature = &v
	return s
}

func (s *ListMessageMetricsRequest) SetSignatureId(v string) *ListMessageMetricsRequest {
	s.SignatureId = &v
	return s
}

func (s *ListMessageMetricsRequest) SetStartDate(v string) *ListMessageMetricsRequest {
	s.StartDate = &v
	return s
}

func (s *ListMessageMetricsRequest) SetTemplateCode(v string) *ListMessageMetricsRequest {
	s.TemplateCode = &v
	return s
}

func (s *ListMessageMetricsRequest) SetTemplateId(v string) *ListMessageMetricsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListMessageMetricsRequest) SetTemplateType(v int32) *ListMessageMetricsRequest {
	s.TemplateType = &v
	return s
}

type ListMessageMetricsResponseBody struct {
	// 返回数据。
	Data *ListMessageMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMessageMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessageMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageMetricsResponseBody) SetData(v *ListMessageMetricsResponseBodyData) *ListMessageMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ListMessageMetricsResponseBody) SetErrorCode(v int32) *ListMessageMetricsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMessageMetricsResponseBody) SetErrorMessage(v string) *ListMessageMetricsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMessageMetricsResponseBody) SetRequestId(v string) *ListMessageMetricsResponseBody {
	s.RequestId = &v
	return s
}

type ListMessageMetricsResponseBodyData struct {
	// 分页返回的统计数据条目列表。
	Metrics []*ListMessageMetricsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总统计数据条目数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessageMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMessageMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessageMetricsResponseBodyData) SetMetrics(v []*ListMessageMetricsResponseBodyDataMetrics) *ListMessageMetricsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListMessageMetricsResponseBodyData) SetPageNumber(v int32) *ListMessageMetricsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMessageMetricsResponseBodyData) SetPageSize(v int32) *ListMessageMetricsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMessageMetricsResponseBodyData) SetTotalCount(v int32) *ListMessageMetricsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListMessageMetricsResponseBodyDataMetrics struct {
	// 发送日期。
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// 发送失败。
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// 发送中。
	Pending *int32 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	// 发送成功率。
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// 发送成功。
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总计短信数量。
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMessageMetricsResponseBodyDataMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListMessageMetricsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListMessageMetricsResponseBodyDataMetrics) SetDate(v string) *ListMessageMetricsResponseBodyDataMetrics {
	s.Date = &v
	return s
}

func (s *ListMessageMetricsResponseBodyDataMetrics) SetFail(v int32) *ListMessageMetricsResponseBodyDataMetrics {
	s.Fail = &v
	return s
}

func (s *ListMessageMetricsResponseBodyDataMetrics) SetPending(v int32) *ListMessageMetricsResponseBodyDataMetrics {
	s.Pending = &v
	return s
}

func (s *ListMessageMetricsResponseBodyDataMetrics) SetRate(v float32) *ListMessageMetricsResponseBodyDataMetrics {
	s.Rate = &v
	return s
}

func (s *ListMessageMetricsResponseBodyDataMetrics) SetSuccess(v int32) *ListMessageMetricsResponseBodyDataMetrics {
	s.Success = &v
	return s
}

func (s *ListMessageMetricsResponseBodyDataMetrics) SetTotal(v int32) *ListMessageMetricsResponseBodyDataMetrics {
	s.Total = &v
	return s
}

type ListMessageMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMessageMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMessageMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessageMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListMessageMetricsResponse) SetHeaders(v map[string]*string) *ListMessageMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListMessageMetricsResponse) SetStatusCode(v int32) *ListMessageMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageMetricsResponse) SetBody(v *ListMessageMetricsResponseBody) *ListMessageMetricsResponse {
	s.Body = v
	return s
}

type ListMessagesRequest struct {
	// 发送日期，格式为20220101。
	Datetime *string `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	// 短信错误码过滤。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 关联人群Id过滤。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 短信Id过滤，短信Id为SendMessage成功返回的Id。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 手机号码过滤。
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 短信批处理Id过滤，短信批处理Id为SendMessage成功返回的RequestId。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 关联触达计划Id过滤。
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// 签名名称过滤。
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 签名Id过滤，同时只能指定签名名称或签名Id其中之一。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 短信发送状态过滤。
	// - 0 : 发送中。
	// - 1 : 发送成功。
	// - 2 : 发送失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板号过滤。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板Id过滤，同时只能指定模板Code或模板Id其中之一。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板类型过滤。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) SetDatetime(v string) *ListMessagesRequest {
	s.Datetime = &v
	return s
}

func (s *ListMessagesRequest) SetErrorCode(v string) *ListMessagesRequest {
	s.ErrorCode = &v
	return s
}

func (s *ListMessagesRequest) SetGroupId(v string) *ListMessagesRequest {
	s.GroupId = &v
	return s
}

func (s *ListMessagesRequest) SetMessageId(v string) *ListMessagesRequest {
	s.MessageId = &v
	return s
}

func (s *ListMessagesRequest) SetPageNumber(v int32) *ListMessagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessagesRequest) SetPageSize(v int32) *ListMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessagesRequest) SetPhoneNumber(v string) *ListMessagesRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListMessagesRequest) SetRequestId(v string) *ListMessagesRequest {
	s.RequestId = &v
	return s
}

func (s *ListMessagesRequest) SetScheduleId(v string) *ListMessagesRequest {
	s.ScheduleId = &v
	return s
}

func (s *ListMessagesRequest) SetSignature(v string) *ListMessagesRequest {
	s.Signature = &v
	return s
}

func (s *ListMessagesRequest) SetSignatureId(v string) *ListMessagesRequest {
	s.SignatureId = &v
	return s
}

func (s *ListMessagesRequest) SetStatus(v int32) *ListMessagesRequest {
	s.Status = &v
	return s
}

func (s *ListMessagesRequest) SetTemplateCode(v string) *ListMessagesRequest {
	s.TemplateCode = &v
	return s
}

func (s *ListMessagesRequest) SetTemplateId(v string) *ListMessagesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListMessagesRequest) SetTemplateType(v int32) *ListMessagesRequest {
	s.TemplateType = &v
	return s
}

type ListMessagesResponseBody struct {
	// 返回数据。
	Data *ListMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBody) SetData(v *ListMessagesResponseBodyData) *ListMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListMessagesResponseBody) SetErrorCode(v int32) *ListMessagesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMessagesResponseBody) SetErrorMessage(v string) *ListMessagesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMessagesResponseBody) SetRequestId(v string) *ListMessagesResponseBody {
	s.RequestId = &v
	return s
}

type ListMessagesResponseBodyData struct {
	// 短信列表。
	Messages []*ListMessagesResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 短信数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMessagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyData) SetMessages(v []*ListMessagesResponseBodyDataMessages) *ListMessagesResponseBodyData {
	s.Messages = v
	return s
}

func (s *ListMessagesResponseBodyData) SetPageNumber(v int32) *ListMessagesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetPageSize(v int32) *ListMessagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetTotalCount(v int32) *ListMessagesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListMessagesResponseBodyDataMessages struct {
	// 短信错误码。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 关联人群Id，未关联则为空。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 短信序列号。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 外部拓展字段。
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// 手机号码。
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 关联触达计划Id，未关联则为空。
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// 签名名称。
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 短信发送状态。
	// - 0 : 发送中。
	// - 1 : 发送成功。
	// - 2 : 发送失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板号。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板参数。
	TemplateParams *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// 模板类型。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListMessagesResponseBodyDataMessages) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataMessages) SetErrorCode(v string) *ListMessagesResponseBodyDataMessages {
	s.ErrorCode = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetGroupId(v string) *ListMessagesResponseBodyDataMessages {
	s.GroupId = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetId(v string) *ListMessagesResponseBodyDataMessages {
	s.Id = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetOutId(v string) *ListMessagesResponseBodyDataMessages {
	s.OutId = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetPhoneNumber(v string) *ListMessagesResponseBodyDataMessages {
	s.PhoneNumber = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetScheduleId(v string) *ListMessagesResponseBodyDataMessages {
	s.ScheduleId = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetSignature(v string) *ListMessagesResponseBodyDataMessages {
	s.Signature = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetStatus(v int32) *ListMessagesResponseBodyDataMessages {
	s.Status = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetTemplateCode(v string) *ListMessagesResponseBodyDataMessages {
	s.TemplateCode = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetTemplateParams(v string) *ListMessagesResponseBodyDataMessages {
	s.TemplateParams = &v
	return s
}

func (s *ListMessagesResponseBodyDataMessages) SetTemplateType(v int32) *ListMessagesResponseBodyDataMessages {
	s.TemplateType = &v
	return s
}

type ListMessagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListMessagesResponse) SetHeaders(v map[string]*string) *ListMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListMessagesResponse) SetStatusCode(v int32) *ListMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessagesResponse) SetBody(v *ListMessagesResponseBody) *ListMessagesResponse {
	s.Body = v
	return s
}

type ListSchedulesRequest struct {
	// 触达计划名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 发送状态过滤。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSchedulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListSchedulesRequest) SetName(v string) *ListSchedulesRequest {
	s.Name = &v
	return s
}

func (s *ListSchedulesRequest) SetPageNumber(v int32) *ListSchedulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSchedulesRequest) SetPageSize(v int32) *ListSchedulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchedulesRequest) SetStatus(v int32) *ListSchedulesRequest {
	s.Status = &v
	return s
}

type ListSchedulesResponseBody struct {
	// 返回数据。
	Data *ListSchedulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSchedulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBody) SetData(v *ListSchedulesResponseBodyData) *ListSchedulesResponseBody {
	s.Data = v
	return s
}

func (s *ListSchedulesResponseBody) SetErrorCode(v int32) *ListSchedulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSchedulesResponseBody) SetErrorMessage(v string) *ListSchedulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSchedulesResponseBody) SetRequestId(v string) *ListSchedulesResponseBody {
	s.RequestId = &v
	return s
}

type ListSchedulesResponseBodyData struct {
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 触达计划列表。
	Schedules []*ListSchedulesResponseBodyDataSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
	// 触达计划数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSchedulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodyData) SetPageNumber(v int32) *ListSchedulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSchedulesResponseBodyData) SetPageSize(v int32) *ListSchedulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSchedulesResponseBodyData) SetSchedules(v []*ListSchedulesResponseBodyDataSchedules) *ListSchedulesResponseBodyData {
	s.Schedules = v
	return s
}

func (s *ListSchedulesResponseBodyData) SetTotalCount(v int32) *ListSchedulesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSchedulesResponseBodyDataSchedules struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 终止时间（UTC+8）。
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 执行时间 (UTC+8)，为空立即执行。
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// 人群ID。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 触达计划ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 触达计划名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 重复周期，按重复周期与重复周期单位执行。
	RepeatCycle *int32 `json:"RepeatCycle,omitempty" xml:"RepeatCycle,omitempty"`
	// 重复周期单位，若指定执行时间，则重复周期生效。
	// - 0: 从不（默认）。
	// - 1: 小时。
	// - 2: 天。
	// - 3: 周。
	// - 4: 月。
	RepeatCycleUnit *int32 `json:"RepeatCycleUnit,omitempty" xml:"RepeatCycleUnit,omitempty"`
	// 重复次数。
	// - 0: 不设终止时间（默认）。
	// - N: 重复N次后终止。
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// 签名。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名ID，或指定签名。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 状态。
	// - 0: 检查中。
	// - 1: 检查成功。
	// - 2: 检查失败。
	// - 3: 发送中。
	// - 4: 发送成功。
	// - 5: 发送失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板ID，或指定模板Code。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListSchedulesResponseBodyDataSchedules) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBodyDataSchedules) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodyDataSchedules) SetCreatedTime(v string) *ListSchedulesResponseBodyDataSchedules {
	s.CreatedTime = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetEndTime(v int32) *ListSchedulesResponseBodyDataSchedules {
	s.EndTime = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetExecuteTime(v string) *ListSchedulesResponseBodyDataSchedules {
	s.ExecuteTime = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetGroupId(v string) *ListSchedulesResponseBodyDataSchedules {
	s.GroupId = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetId(v string) *ListSchedulesResponseBodyDataSchedules {
	s.Id = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetName(v string) *ListSchedulesResponseBodyDataSchedules {
	s.Name = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetRepeatCycle(v int32) *ListSchedulesResponseBodyDataSchedules {
	s.RepeatCycle = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetRepeatCycleUnit(v int32) *ListSchedulesResponseBodyDataSchedules {
	s.RepeatCycleUnit = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetRepeatTimes(v int32) *ListSchedulesResponseBodyDataSchedules {
	s.RepeatTimes = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetSignName(v string) *ListSchedulesResponseBodyDataSchedules {
	s.SignName = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetSignatureId(v string) *ListSchedulesResponseBodyDataSchedules {
	s.SignatureId = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetStatus(v int32) *ListSchedulesResponseBodyDataSchedules {
	s.Status = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetTemplateCode(v string) *ListSchedulesResponseBodyDataSchedules {
	s.TemplateCode = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetTemplateId(v string) *ListSchedulesResponseBodyDataSchedules {
	s.TemplateId = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetUpdatedTime(v string) *ListSchedulesResponseBodyDataSchedules {
	s.UpdatedTime = &v
	return s
}

type ListSchedulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSchedulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponse) SetHeaders(v map[string]*string) *ListSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListSchedulesResponse) SetStatusCode(v int32) *ListSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchedulesResponse) SetBody(v *ListSchedulesResponseBody) *ListSchedulesResponse {
	s.Body = v
	return s
}

type ListSignaturesRequest struct {
	// 签名名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 签名审核状态过滤。
	// - 0：审核中。
	// - 1：审核通过。
	// - 2：审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSignaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesRequest) GoString() string {
	return s.String()
}

func (s *ListSignaturesRequest) SetName(v string) *ListSignaturesRequest {
	s.Name = &v
	return s
}

func (s *ListSignaturesRequest) SetPageNumber(v int32) *ListSignaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSignaturesRequest) SetPageSize(v int32) *ListSignaturesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSignaturesRequest) SetStatus(v int32) *ListSignaturesRequest {
	s.Status = &v
	return s
}

type ListSignaturesResponseBody struct {
	// 返回数据。
	Data *ListSignaturesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSignaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponseBody) SetData(v *ListSignaturesResponseBodyData) *ListSignaturesResponseBody {
	s.Data = v
	return s
}

func (s *ListSignaturesResponseBody) SetErrorCode(v int32) *ListSignaturesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSignaturesResponseBody) SetErrorMessage(v string) *ListSignaturesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSignaturesResponseBody) SetRequestId(v string) *ListSignaturesResponseBody {
	s.RequestId = &v
	return s
}

type ListSignaturesResponseBodyData struct {
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 分页返回的签名列表。
	Signatures []*ListSignaturesResponseBodyDataSignatures `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
	// 账号下全部签名注册记录数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSignaturesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponseBodyData) SetPageNumber(v int32) *ListSignaturesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSignaturesResponseBodyData) SetPageSize(v int32) *ListSignaturesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSignaturesResponseBodyData) SetSignatures(v []*ListSignaturesResponseBodyDataSignatures) *ListSignaturesResponseBodyData {
	s.Signatures = v
	return s
}

func (s *ListSignaturesResponseBodyData) SetTotalCount(v int32) *ListSignaturesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSignaturesResponseBodyDataSignatures struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 签名Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 签名审核状态。
	// - 0：审核中。
	// - 1：审核通过。
	// - 2：审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListSignaturesResponseBodyDataSignatures) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponseBodyDataSignatures) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponseBodyDataSignatures) SetCreatedTime(v string) *ListSignaturesResponseBodyDataSignatures {
	s.CreatedTime = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetId(v string) *ListSignaturesResponseBodyDataSignatures {
	s.Id = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetName(v string) *ListSignaturesResponseBodyDataSignatures {
	s.Name = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetStatus(v int32) *ListSignaturesResponseBodyDataSignatures {
	s.Status = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetUpdatedTime(v string) *ListSignaturesResponseBodyDataSignatures {
	s.UpdatedTime = &v
	return s
}

type ListSignaturesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSignaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponse) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponse) SetHeaders(v map[string]*string) *ListSignaturesResponse {
	s.Headers = v
	return s
}

func (s *ListSignaturesResponse) SetStatusCode(v int32) *ListSignaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSignaturesResponse) SetBody(v *ListSignaturesResponseBody) *ListSignaturesResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// 模板内容过滤。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 模板名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 审核状态过滤。
	// - 0 : 审核中。
	// - 1 : 审核通过。
	// - 2 : 审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板类型过滤。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetContent(v string) *ListTemplatesRequest {
	s.Content = &v
	return s
}

func (s *ListTemplatesRequest) SetName(v string) *ListTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetStatus(v int32) *ListTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplatesRequest) SetType(v int32) *ListTemplatesRequest {
	s.Type = &v
	return s
}

type ListTemplatesResponseBody struct {
	// 返回数据。
	Data *ListTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetData(v *ListTemplatesResponseBodyData) *ListTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplatesResponseBody) SetErrorCode(v int32) *ListTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTemplatesResponseBody) SetErrorMessage(v string) *ListTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListTemplatesResponseBodyData struct {
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 分页返回的模板列表。
	Templates []*ListTemplatesResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// 全部模板注册记录数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyData) SetPageNumber(v int32) *ListTemplatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBodyData) SetPageSize(v int32) *ListTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBodyData) SetTemplates(v []*ListTemplatesResponseBodyDataTemplates) *ListTemplatesResponseBodyData {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBodyData) SetTotalCount(v int32) *ListTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyDataTemplates struct {
	// 模板内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核意见。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名Id。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 审核状态。
	// - 0 : 审核中。
	// - 1 : 审核通过。
	// - 2 : 审核不通过。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板类型。
	// - 0 : 验证码。
	// - 1 : 短信通知。
	// - 2 : 推广短信。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListTemplatesResponseBodyDataTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyDataTemplates) SetContent(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Content = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetCreatedTime(v string) *ListTemplatesResponseBodyDataTemplates {
	s.CreatedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetDescription(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetId(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Id = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetName(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetReason(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Reason = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetSignatureId(v string) *ListTemplatesResponseBodyDataTemplates {
	s.SignatureId = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetStatus(v int32) *ListTemplatesResponseBodyDataTemplates {
	s.Status = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetTemplateCode(v string) *ListTemplatesResponseBodyDataTemplates {
	s.TemplateCode = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetType(v int32) *ListTemplatesResponseBodyDataTemplates {
	s.Type = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetUpdatedTime(v string) *ListTemplatesResponseBodyDataTemplates {
	s.UpdatedTime = &v
	return s
}

type ListTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListTrainingJobsRequest struct {
	// 归属运营活动过滤。
	CampaignId   *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CampaignName *string `json:"CampaignName,omitempty" xml:"CampaignName,omitempty"`
	// 训练任务名称过滤。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 训练任务备注过滤。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 训练任务状态过滤。
	// - 0: 队列中。
	// - 1: 已提交。
	// - 2: 运行中。
	// - 3: 成功。
	// - 4: 失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 归属训练计划过滤。
	TrainingScheduleId *string `json:"TrainingScheduleId,omitempty" xml:"TrainingScheduleId,omitempty"`
}

func (s ListTrainingJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsRequest) SetCampaignId(v string) *ListTrainingJobsRequest {
	s.CampaignId = &v
	return s
}

func (s *ListTrainingJobsRequest) SetCampaignName(v string) *ListTrainingJobsRequest {
	s.CampaignName = &v
	return s
}

func (s *ListTrainingJobsRequest) SetName(v string) *ListTrainingJobsRequest {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsRequest) SetPageNumber(v int32) *ListTrainingJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobsRequest) SetPageSize(v int32) *ListTrainingJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobsRequest) SetRemark(v string) *ListTrainingJobsRequest {
	s.Remark = &v
	return s
}

func (s *ListTrainingJobsRequest) SetStatus(v int32) *ListTrainingJobsRequest {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsRequest) SetTrainingScheduleId(v string) *ListTrainingJobsRequest {
	s.TrainingScheduleId = &v
	return s
}

type ListTrainingJobsResponseBody struct {
	// 返回数据。
	Data *ListTrainingJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTrainingJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBody) SetData(v *ListTrainingJobsResponseBodyData) *ListTrainingJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListTrainingJobsResponseBody) SetErrorCode(v int32) *ListTrainingJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTrainingJobsResponseBody) SetErrorMessage(v string) *ListTrainingJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTrainingJobsResponseBody) SetRequestId(v string) *ListTrainingJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListTrainingJobsResponseBodyData struct {
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总训练任务数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 训练任务列表。
	TrainingJobs []*ListTrainingJobsResponseBodyDataTrainingJobs `json:"TrainingJobs,omitempty" xml:"TrainingJobs,omitempty" type:"Repeated"`
}

func (s ListTrainingJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyData) SetPageNumber(v int32) *ListTrainingJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTrainingJobsResponseBodyData) SetPageSize(v int32) *ListTrainingJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTrainingJobsResponseBodyData) SetTotalCount(v int32) *ListTrainingJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobsResponseBodyData) SetTrainingJobs(v []*ListTrainingJobsResponseBodyDataTrainingJobs) *ListTrainingJobsResponseBodyData {
	s.TrainingJobs = v
	return s
}

type ListTrainingJobsResponseBodyDataTrainingJobs struct {
	// 关联算法ID。
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// 关联运营活动ID。
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 训练数据路径，指定路径前需确保已在控制台完成一键授权。
	DataPath     *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	HasModelInfo *bool   `json:"HasModelInfo,omitempty" xml:"HasModelInfo,omitempty"`
	// 训练任务日志。
	History *string `json:"History,omitempty" xml:"History,omitempty"`
	// 训练任务ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 训练任务名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 训练任务状态。
	// - 0: 队列中。
	// - 1: 已提交。
	// - 2: 运行中。
	// - 3: 成功。
	// - 4: 失败。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 关联训练计划ID。
	TrainingScheduleId *string `json:"TrainingScheduleId,omitempty" xml:"TrainingScheduleId,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 用户配置。
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s ListTrainingJobsResponseBodyDataTrainingJobs) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponseBodyDataTrainingJobs) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetAlgorithm(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.Algorithm = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetCampaignId(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.CampaignId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetCreatedTime(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.CreatedTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetDataPath(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.DataPath = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetHasModelInfo(v bool) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.HasModelInfo = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetHistory(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.History = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetId(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.Id = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetName(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetRemark(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.Remark = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetStatus(v int32) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetTrainingScheduleId(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.TrainingScheduleId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetUpdatedTime(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.UpdatedTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyDataTrainingJobs) SetUserConfig(v string) *ListTrainingJobsResponseBodyDataTrainingJobs {
	s.UserConfig = &v
	return s
}

type ListTrainingJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrainingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTrainingJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponse) SetHeaders(v map[string]*string) *ListTrainingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobsResponse) SetStatusCode(v int32) *ListTrainingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobsResponse) SetBody(v *ListTrainingJobsResponseBody) *ListTrainingJobsResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	// 人群Id，用于关联人群。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 外部拓展字段，示例：["1234567890"]。
	OutIds      []*string `json:"OutIds,omitempty" xml:"OutIds,omitempty" type:"Repeated"`
	PaymentType *string   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 手机号，每个手机号对应一个模板变量、上行拓展码和外部拓展字段，示例：["1234567890"]。
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// 触达计划Id，用于关联触达计划。
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// 签名名称。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名Id，同时只能指定签名名称或签名Id其中之一。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 短信上行拓展码，示例：["1234567890"]。
	SmsUpExtendCodes []*string `json:"SmsUpExtendCodes,omitempty" xml:"SmsUpExtendCodes,omitempty" type:"Repeated"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板Id，同时只能指定模板Code或模板Id其中之一。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 短信模板变量对应的实际值，JSON格式。支持传入多个参数，示例：[{"name":"张三","number":"15038****76"}]。
	TemplateParams []*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty" type:"Repeated"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetGroupId(v string) *SendMessageRequest {
	s.GroupId = &v
	return s
}

func (s *SendMessageRequest) SetOutIds(v []*string) *SendMessageRequest {
	s.OutIds = v
	return s
}

func (s *SendMessageRequest) SetPaymentType(v string) *SendMessageRequest {
	s.PaymentType = &v
	return s
}

func (s *SendMessageRequest) SetPhoneNumbers(v []*string) *SendMessageRequest {
	s.PhoneNumbers = v
	return s
}

func (s *SendMessageRequest) SetScheduleId(v string) *SendMessageRequest {
	s.ScheduleId = &v
	return s
}

func (s *SendMessageRequest) SetSignName(v string) *SendMessageRequest {
	s.SignName = &v
	return s
}

func (s *SendMessageRequest) SetSignatureId(v string) *SendMessageRequest {
	s.SignatureId = &v
	return s
}

func (s *SendMessageRequest) SetSmsUpExtendCodes(v []*string) *SendMessageRequest {
	s.SmsUpExtendCodes = v
	return s
}

func (s *SendMessageRequest) SetTemplateCode(v string) *SendMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendMessageRequest) SetTemplateId(v string) *SendMessageRequest {
	s.TemplateId = &v
	return s
}

func (s *SendMessageRequest) SetTemplateParams(v []*string) *SendMessageRequest {
	s.TemplateParams = v
	return s
}

type SendMessageResponseBody struct {
	// 返回数据。
	Data *SendMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetData(v *SendMessageResponseBodyData) *SendMessageResponseBody {
	s.Data = v
	return s
}

func (s *SendMessageResponseBody) SetErrorCode(v int32) *SendMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendMessageResponseBody) SetErrorMessage(v string) *SendMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendMessageResponseBodyData struct {
	// 短信结果列表，列表中手机号的顺序与输入请求手机号顺序一一对应。
	Messages []*SendMessageResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// 短信批处理Id，可使用ListMessages查询短信状态。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendMessageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBodyData) SetMessages(v []*SendMessageResponseBodyDataMessages) *SendMessageResponseBodyData {
	s.Messages = v
	return s
}

func (s *SendMessageResponseBodyData) SetRequestId(v string) *SendMessageResponseBodyData {
	s.RequestId = &v
	return s
}

type SendMessageResponseBodyDataMessages struct {
	// 短信Id，可使用ListMessages查询短信状态。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 手机号码。
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s SendMessageResponseBodyDataMessages) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBodyDataMessages) SetId(v string) *SendMessageResponseBodyDataMessages {
	s.Id = &v
	return s
}

func (s *SendMessageResponseBodyDataMessages) SetPhoneNumber(v string) *SendMessageResponseBodyDataMessages {
	s.PhoneNumber = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type SmsReportRequest struct {
	// 请求参数的主体信息。
	Body []*SmsReportRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s SmsReportRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsReportRequest) GoString() string {
	return s.String()
}

func (s *SmsReportRequest) SetBody(v []*SmsReportRequestBody) *SmsReportRequest {
	s.Body = v
	return s
}

type SmsReportRequestBody struct {
	// 发送回执ID，即发送流水号。
	BizId *string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 状态报告编码。
	ErrCode *string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 状态报告说明。
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 短信Id。调用发送接口SendMessage发送短信时，返回值中的Id字段。可使用短信Id在接口ListMessages查询具体的发送状态。
	MessageId *string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// 外部拓展字段。
	OutId *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 手机号码。
	PhoneNumber *string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 状态报告时间。
	ReportTime *string `json:"report_time,omitempty" xml:"report_time,omitempty"`
	// 短信批处理Id。调用发送接口SendMessage发送短信时，返回值中的RequestId字段。可使用短信批处理Id在接口ListMessages查询具体的发送状态。
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发送时间。
	SendTime *string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 签名。
	SignName *string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
	// 短信长度。短信长度不超过70个字，按照一条短信计费；超过70个字，即为长短信，按照67字/条拆分成多条计费。
	SmsSize *string `json:"sms_size,omitempty" xml:"sms_size,omitempty"`
	// 是否接收成功。
	// - true : 接收成功。
	// - false : 接收失败。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 模板号。
	TemplateCode *string `json:"template_code,omitempty" xml:"template_code,omitempty"`
}

func (s SmsReportRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SmsReportRequestBody) GoString() string {
	return s.String()
}

func (s *SmsReportRequestBody) SetBizId(v string) *SmsReportRequestBody {
	s.BizId = &v
	return s
}

func (s *SmsReportRequestBody) SetErrCode(v string) *SmsReportRequestBody {
	s.ErrCode = &v
	return s
}

func (s *SmsReportRequestBody) SetErrMsg(v string) *SmsReportRequestBody {
	s.ErrMsg = &v
	return s
}

func (s *SmsReportRequestBody) SetMessageId(v string) *SmsReportRequestBody {
	s.MessageId = &v
	return s
}

func (s *SmsReportRequestBody) SetOutId(v string) *SmsReportRequestBody {
	s.OutId = &v
	return s
}

func (s *SmsReportRequestBody) SetPhoneNumber(v string) *SmsReportRequestBody {
	s.PhoneNumber = &v
	return s
}

func (s *SmsReportRequestBody) SetReportTime(v string) *SmsReportRequestBody {
	s.ReportTime = &v
	return s
}

func (s *SmsReportRequestBody) SetRequestId(v string) *SmsReportRequestBody {
	s.RequestId = &v
	return s
}

func (s *SmsReportRequestBody) SetSendTime(v string) *SmsReportRequestBody {
	s.SendTime = &v
	return s
}

func (s *SmsReportRequestBody) SetSignName(v string) *SmsReportRequestBody {
	s.SignName = &v
	return s
}

func (s *SmsReportRequestBody) SetSmsSize(v string) *SmsReportRequestBody {
	s.SmsSize = &v
	return s
}

func (s *SmsReportRequestBody) SetSuccess(v bool) *SmsReportRequestBody {
	s.Success = &v
	return s
}

func (s *SmsReportRequestBody) SetTemplateCode(v string) *SmsReportRequestBody {
	s.TemplateCode = &v
	return s
}

type SmsReportResponseBody struct {
	// 应答编码。
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// 描述信息。
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s SmsReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsReportResponseBody) GoString() string {
	return s.String()
}

func (s *SmsReportResponseBody) SetCode(v int32) *SmsReportResponseBody {
	s.Code = &v
	return s
}

func (s *SmsReportResponseBody) SetMsg(v string) *SmsReportResponseBody {
	s.Msg = &v
	return s
}

type SmsReportResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SmsReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SmsReportResponse) String() string {
	return tea.Prettify(s)
}

func (s SmsReportResponse) GoString() string {
	return s.String()
}

func (s *SmsReportResponse) SetHeaders(v map[string]*string) *SmsReportResponse {
	s.Headers = v
	return s
}

func (s *SmsReportResponse) SetStatusCode(v int32) *SmsReportResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsReportResponse) SetBody(v *SmsReportResponseBody) *SmsReportResponse {
	s.Body = v
	return s
}

type SmsUpRequest struct {
	// 请求参数的主体信息。
	Body []*SmsUpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s SmsUpRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsUpRequest) GoString() string {
	return s.String()
}

func (s *SmsUpRequest) SetBody(v []*SmsUpRequestBody) *SmsUpRequest {
	s.Body = v
	return s
}

type SmsUpRequestBody struct {
	// 发送内容。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 上行短信扩展号码，系统后台自动生成，不支持自定义传入。
	DestCode *string `json:"dest_code,omitempty" xml:"dest_code,omitempty"`
	// 手机号码。
	PhoneNumber *string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 发送时间。
	SendTime *string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 序列号。
	SequenceId *int32 `json:"sequence_id,omitempty" xml:"sequence_id,omitempty"`
	// 签名信息。
	SignName *string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
}

func (s SmsUpRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SmsUpRequestBody) GoString() string {
	return s.String()
}

func (s *SmsUpRequestBody) SetContent(v string) *SmsUpRequestBody {
	s.Content = &v
	return s
}

func (s *SmsUpRequestBody) SetDestCode(v string) *SmsUpRequestBody {
	s.DestCode = &v
	return s
}

func (s *SmsUpRequestBody) SetPhoneNumber(v string) *SmsUpRequestBody {
	s.PhoneNumber = &v
	return s
}

func (s *SmsUpRequestBody) SetSendTime(v string) *SmsUpRequestBody {
	s.SendTime = &v
	return s
}

func (s *SmsUpRequestBody) SetSequenceId(v int32) *SmsUpRequestBody {
	s.SequenceId = &v
	return s
}

func (s *SmsUpRequestBody) SetSignName(v string) *SmsUpRequestBody {
	s.SignName = &v
	return s
}

type SmsUpResponseBody struct {
	// 应答编码。
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// 描述信息。
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s SmsUpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsUpResponseBody) GoString() string {
	return s.String()
}

func (s *SmsUpResponseBody) SetCode(v int32) *SmsUpResponseBody {
	s.Code = &v
	return s
}

func (s *SmsUpResponseBody) SetMsg(v string) *SmsUpResponseBody {
	s.Msg = &v
	return s
}

type SmsUpResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SmsUpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SmsUpResponse) String() string {
	return tea.Prettify(s)
}

func (s SmsUpResponse) GoString() string {
	return s.String()
}

func (s *SmsUpResponse) SetHeaders(v map[string]*string) *SmsUpResponse {
	s.Headers = v
	return s
}

func (s *SmsUpResponse) SetStatusCode(v int32) *SmsUpResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsUpResponse) SetBody(v *SmsUpResponseBody) *SmsUpResponse {
	s.Body = v
	return s
}

type UpdateCampaignRequest struct {
	// 运营活动名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCampaignRequest) GoString() string {
	return s.String()
}

func (s *UpdateCampaignRequest) SetName(v string) *UpdateCampaignRequest {
	s.Name = &v
	return s
}

func (s *UpdateCampaignRequest) SetRemark(v string) *UpdateCampaignRequest {
	s.Remark = &v
	return s
}

type UpdateCampaignResponseBody struct {
	// 返回数据。
	Data *UpdateCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCampaignResponseBody) SetData(v *UpdateCampaignResponseBodyData) *UpdateCampaignResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCampaignResponseBody) SetErrorCode(v int32) *UpdateCampaignResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateCampaignResponseBody) SetErrorMessage(v string) *UpdateCampaignResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateCampaignResponseBody) SetRequestId(v string) *UpdateCampaignResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCampaignResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 运营活动Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 运营活动名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 备注。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s UpdateCampaignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCampaignResponseBodyData) SetCreatedTime(v string) *UpdateCampaignResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *UpdateCampaignResponseBodyData) SetId(v string) *UpdateCampaignResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateCampaignResponseBodyData) SetName(v string) *UpdateCampaignResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateCampaignResponseBodyData) SetRemark(v string) *UpdateCampaignResponseBodyData {
	s.Remark = &v
	return s
}

func (s *UpdateCampaignResponseBodyData) SetUpdatedTime(v string) *UpdateCampaignResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type UpdateCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCampaignResponse) GoString() string {
	return s.String()
}

func (s *UpdateCampaignResponse) SetHeaders(v map[string]*string) *UpdateCampaignResponse {
	s.Headers = v
	return s
}

func (s *UpdateCampaignResponse) SetStatusCode(v int32) *UpdateCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCampaignResponse) SetBody(v *UpdateCampaignResponseBody) *UpdateCampaignResponse {
	s.Body = v
	return s
}

type UpdateReportUrlRequest struct {
	// 可公开访问的地址。
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateReportUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReportUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateReportUrlRequest) SetUrl(v string) *UpdateReportUrlRequest {
	s.Url = &v
	return s
}

type UpdateReportUrlResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateReportUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateReportUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReportUrlResponseBody) SetData(v string) *UpdateReportUrlResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateReportUrlResponseBody) SetErrorCode(v int32) *UpdateReportUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateReportUrlResponseBody) SetErrorMessage(v string) *UpdateReportUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateReportUrlResponseBody) SetRequestId(v string) *UpdateReportUrlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateReportUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateReportUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateReportUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateReportUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateReportUrlResponse) SetHeaders(v map[string]*string) *UpdateReportUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateReportUrlResponse) SetStatusCode(v int32) *UpdateReportUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReportUrlResponse) SetBody(v *UpdateReportUrlResponseBody) *UpdateReportUrlResponse {
	s.Body = v
	return s
}

type UpdateUploadUrlRequest struct {
	// 可公开访问的地址。
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateUploadUrlRequest) SetUrl(v string) *UpdateUploadUrlRequest {
	s.Url = &v
	return s
}

type UpdateUploadUrlResponseBody struct {
	// 返回数据。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码。
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUploadUrlResponseBody) SetData(v string) *UpdateUploadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUploadUrlResponseBody) SetErrorCode(v int32) *UpdateUploadUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUploadUrlResponseBody) SetErrorMessage(v string) *UpdateUploadUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateUploadUrlResponseBody) SetRequestId(v string) *UpdateUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUploadUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateUploadUrlResponse) SetHeaders(v map[string]*string) *UpdateUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateUploadUrlResponse) SetStatusCode(v int32) *UpdateUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUploadUrlResponse) SetBody(v *UpdateUploadUrlResponseBody) *UpdateUploadUrlResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("paiplugin"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * 注册运营活动。
 *
 * @param request CreateCampaignRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCampaignResponse
 */
func (client *Client) CreateCampaignWithOptions(request *CreateCampaignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCampaign"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/campaigns"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 注册运营活动。
 *
 * @param request CreateCampaignRequest
 * @return CreateCampaignResponse
 */
func (client *Client) CreateCampaign(request *CreateCampaignRequest) (_result *CreateCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCampaignResponse{}
	_body, _err := client.CreateCampaignWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 注册人群。
 *
 * @param request CreateGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateGroupResponse
 */
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		body["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Column)) {
		body["Column"] = request.Column
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InferenceJobId)) {
		body["InferenceJobId"] = request.InferenceJobId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		body["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		body["Table"] = request.Table
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 注册人群。
 *
 * @param request CreateGroupRequest
 * @return CreateGroupResponse
 */
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 注册预测任务。
 *
 * @param request CreateInferenceJobRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateInferenceJobResponse
 */
func (client *Client) CreateInferenceJobWithOptions(request *CreateInferenceJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInferenceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		body["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		body["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.DataPath)) {
		body["DataPath"] = request.DataPath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPath)) {
		body["TargetPath"] = request.TargetPath
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobId)) {
		body["TrainingJobId"] = request.TrainingJobId
	}

	if !tea.BoolValue(util.IsUnset(request.UserConfig)) {
		body["UserConfig"] = request.UserConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInferenceJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/inference/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 注册预测任务。
 *
 * @param request CreateInferenceJobRequest
 * @return CreateInferenceJobResponse
 */
func (client *Client) CreateInferenceJob(request *CreateInferenceJobRequest) (_result *CreateInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInferenceJobResponse{}
	_body, _err := client.CreateInferenceJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduleWithOptions(request *CreateScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AISendEndDate)) {
		body["AISendEndDate"] = request.AISendEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.AISendStartDate)) {
		body["AISendStartDate"] = request.AISendStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteTime)) {
		body["ExecuteTime"] = request.ExecuteTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatCycle)) {
		body["RepeatCycle"] = request.RepeatCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatCycleUnit)) {
		body["RepeatCycleUnit"] = request.RepeatCycleUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatTimes)) {
		body["RepeatTimes"] = request.RepeatTimes
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		body["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		body["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchedule"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/schedules"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchedule(request *CreateScheduleRequest) (_result *CreateScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScheduleResponse{}
	_body, _err := client.CreateScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSignatureWithOptions(request *CreateSignatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSignature"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSignature(request *CreateSignatureRequest) (_result *CreateSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSignatureResponse{}
	_body, _err := client.CreateSignatureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		body["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrainingJobWithOptions(request *CreateTrainingJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTrainingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		body["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		body["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.DataPath)) {
		body["DataPath"] = request.DataPath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserConfig)) {
		body["UserConfig"] = request.UserConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/training/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrainingJob(request *CreateTrainingJobRequest) (_result *CreateTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrainingJobResponse{}
	_body, _err := client.CreateTrainingJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除运营活动
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteCampaignResponse
 */
func (client *Client) DeleteCampaignWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCampaignResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCampaign"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/campaigns/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 删除运营活动
 *
 * @return DeleteCampaignResponse
 */
func (client *Client) DeleteCampaign(Id *string) (_result *DeleteCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCampaignResponse{}
	_body, _err := client.DeleteCampaignWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除人群
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteGroupResponse
 */
func (client *Client) DeleteGroupWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/groups/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 删除人群
 *
 * @return DeleteGroupResponse
 */
func (client *Client) DeleteGroup(Id *string) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除预测任务。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteInferenceJobResponse
 */
func (client *Client) DeleteInferenceJobWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInferenceJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInferenceJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/inference/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 删除预测任务。
 *
 * @return DeleteInferenceJobResponse
 */
func (client *Client) DeleteInferenceJob(Id *string) (_result *DeleteInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInferenceJobResponse{}
	_body, _err := client.DeleteInferenceJobWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除触达计划。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteScheduleResponse
 */
func (client *Client) DeleteScheduleWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteScheduleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchedule"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/schedules/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 删除触达计划。
 *
 * @return DeleteScheduleResponse
 */
func (client *Client) DeleteSchedule(Id *string) (_result *DeleteScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteScheduleResponse{}
	_body, _err := client.DeleteScheduleWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSignatureWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSignatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSignature"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSignature(Id *string) (_result *DeleteSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSignatureResponse{}
	_body, _err := client.DeleteSignatureWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除模板
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTemplateResponse
 */
func (client *Client) DeleteTemplateWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 删除模板
 *
 * @return DeleteTemplateResponse
 */
func (client *Client) DeleteTemplate(Id *string) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 删除训练任务。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTrainingJobResponse
 */
func (client *Client) DeleteTrainingJobWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTrainingJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/training/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 删除训练任务。
 *
 * @return DeleteTrainingJobResponse
 */
func (client *Client) DeleteTrainingJob(Id *string) (_result *DeleteTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrainingJobResponse{}
	_body, _err := client.DeleteTrainingJobWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取算法详情。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAlgorithmResponse
 */
func (client *Client) GetAlgorithmWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgorithmResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgorithm"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/algorithms/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取算法详情。
 *
 * @return GetAlgorithmResponse
 */
func (client *Client) GetAlgorithm(Id *string) (_result *GetAlgorithmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmResponse{}
	_body, _err := client.GetAlgorithmWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取运营活动详情。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetCampaignResponse
 */
func (client *Client) GetCampaignWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCampaignResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCampaign"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/campaigns/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取运营活动详情。
 *
 * @return GetCampaignResponse
 */
func (client *Client) GetCampaign(Id *string) (_result *GetCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCampaignResponse{}
	_body, _err := client.GetCampaignWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取人群详情。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetGroupResponse
 */
func (client *Client) GetGroupWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/groups/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取人群详情。
 *
 * @return GetGroupResponse
 */
func (client *Client) GetGroup(Id *string) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取预测任务详情。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetInferenceJobResponse
 */
func (client *Client) GetInferenceJobWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInferenceJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInferenceJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/inference/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取预测任务详情。
 *
 * @return GetInferenceJobResponse
 */
func (client *Client) GetInferenceJob(Id *string) (_result *GetInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInferenceJobResponse{}
	_body, _err := client.GetInferenceJobWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取短信配置。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetMessageConfigResponse
 */
func (client *Client) GetMessageConfigWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMessageConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMessageConfig"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/users/messageConfig"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMessageConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取短信配置。
 *
 * @return GetMessageConfigResponse
 */
func (client *Client) GetMessageConfig() (_result *GetMessageConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMessageConfigResponse{}
	_body, _err := client.GetMessageConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScheduleWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetScheduleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSchedule"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/schedules/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSchedule(Id *string) (_result *GetScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduleResponse{}
	_body, _err := client.GetScheduleWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取签名详情。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetSignatureResponse
 */
func (client *Client) GetSignatureWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSignatureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSignature"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取签名详情。
 *
 * @return GetSignatureResponse
 */
func (client *Client) GetSignature(Id *string) (_result *GetSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSignatureResponse{}
	_body, _err := client.GetSignatureWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取模板详情。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTemplateResponse
 */
func (client *Client) GetTemplateWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取模板详情。
 *
 * @return GetTemplateResponse
 */
func (client *Client) GetTemplate(Id *string) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrainingJobWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrainingJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainingJob"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/training/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrainingJob(Id *string) (_result *GetTrainingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobResponse{}
	_body, _err := client.GetTrainingJobWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取账号状态。
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetUserResponse
 */
func (client *Client) GetUserWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取账号状态。
 *
 * @return GetUserResponse
 */
func (client *Client) GetUser() (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取算法列表。
 *
 * @param request ListAlgorithmsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAlgorithmsResponse
 */
func (client *Client) ListAlgorithmsWithOptions(request *ListAlgorithmsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlgorithmsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlgorithms"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/algorithms"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlgorithmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取算法列表。
 *
 * @param request ListAlgorithmsRequest
 * @return ListAlgorithmsResponse
 */
func (client *Client) ListAlgorithms(request *ListAlgorithmsRequest) (_result *ListAlgorithmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlgorithmsResponse{}
	_body, _err := client.ListAlgorithmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取运营活动列表。
 *
 * @param request ListCampaignsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListCampaignsResponse
 */
func (client *Client) ListCampaignsWithOptions(request *ListCampaignsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCampaignsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCampaigns"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/campaigns"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCampaignsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取运营活动列表。
 *
 * @param request ListCampaignsRequest
 * @return ListCampaignsResponse
 */
func (client *Client) ListCampaigns(request *ListCampaignsRequest) (_result *ListCampaignsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCampaignsResponse{}
	_body, _err := client.ListCampaignsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取人群列表。
 *
 * @param request ListGroupsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListGroupsResponse
 */
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取人群列表。
 *
 * @param request ListGroupsRequest
 * @return ListGroupsResponse
 */
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInferenceJobsWithOptions(request *ListInferenceJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInferenceJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.CampaignName)) {
		query["CampaignName"] = request.CampaignName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingJobName)) {
		query["TrainingJobName"] = request.TrainingJobName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInferenceJobs"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/inference/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInferenceJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInferenceJobs(request *ListInferenceJobsRequest) (_result *ListInferenceJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInferenceJobsResponse{}
	_body, _err := client.ListInferenceJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取短信发送统计列表。
 * 获取短信发送统计数据，可按指定条件获取分类别详细数据，返回数据按日期顺序排列，发送统计为空的日期默认不返回。
 * 发送数据在48小时内会随实际短信发送状态不断更新，最终数据以48小时后数据为准。
 *
 * @param request ListMessageMetricsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListMessageMetricsResponse
 */
func (client *Client) ListMessageMetricsWithOptions(request *ListMessageMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMessageMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleId)) {
		query["ScheduleId"] = request.ScheduleId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		query["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessageMetrics"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/messages/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessageMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取短信发送统计列表。
 * 获取短信发送统计数据，可按指定条件获取分类别详细数据，返回数据按日期顺序排列，发送统计为空的日期默认不返回。
 * 发送数据在48小时内会随实际短信发送状态不断更新，最终数据以48小时后数据为准。
 *
 * @param request ListMessageMetricsRequest
 * @return ListMessageMetricsResponse
 */
func (client *Client) ListMessageMetrics(request *ListMessageMetricsRequest) (_result *ListMessageMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMessageMetricsResponse{}
	_body, _err := client.ListMessageMetricsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 查询短信发送详情列表。
 *
 * @param request ListMessagesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListMessagesResponse
 */
func (client *Client) ListMessagesWithOptions(request *ListMessagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Datetime)) {
		query["Datetime"] = request.Datetime
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		query["ErrorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleId)) {
		query["ScheduleId"] = request.ScheduleId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		query["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessages"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/messages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 查询短信发送详情列表。
 *
 * @param request ListMessagesRequest
 * @return ListMessagesResponse
 */
func (client *Client) ListMessages(request *ListMessagesRequest) (_result *ListMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMessagesResponse{}
	_body, _err := client.ListMessagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSchedulesWithOptions(request *ListSchedulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSchedulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSchedules"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/schedules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSchedules(request *ListSchedulesRequest) (_result *ListSchedulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSchedulesResponse{}
	_body, _err := client.ListSchedulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取签名列表。
 *
 * @param request ListSignaturesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSignaturesResponse
 */
func (client *Client) ListSignaturesWithOptions(request *ListSignaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSignaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSignatures"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSignaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取签名列表。
 *
 * @param request ListSignaturesRequest
 * @return ListSignaturesResponse
 */
func (client *Client) ListSignatures(request *ListSignaturesRequest) (_result *ListSignaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSignaturesResponse{}
	_body, _err := client.ListSignaturesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 获取模板列表。
 *
 * @param request ListTemplatesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTemplatesResponse
 */
func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 获取模板列表。
 *
 * @param request ListTemplatesRequest
 * @return ListTemplatesResponse
 */
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTrainingJobsWithOptions(request *ListTrainingJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTrainingJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.CampaignName)) {
		query["CampaignName"] = request.CampaignName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingScheduleId)) {
		query["TrainingScheduleId"] = request.TrainingScheduleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainingJobs"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/training/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrainingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrainingJobs(request *ListTrainingJobsRequest) (_result *ListTrainingJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobsResponse{}
	_body, _err := client.ListTrainingJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OutIds)) {
		body["OutIds"] = request.OutIds
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		body["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleId)) {
		body["ScheduleId"] = request.ScheduleId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		body["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		body["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCodes)) {
		body["SmsUpExtendCodes"] = request.SmsUpExtendCodes
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParams)) {
		body["TemplateParams"] = request.TemplateParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/messages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 短信回执。
 *
 * @param request SmsReportRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return SmsReportResponse
 */
func (client *Client) SmsReportWithOptions(request *SmsReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SmsReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("SmsReport"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/recall/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SmsReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 短信回执。
 *
 * @param request SmsReportRequest
 * @return SmsReportResponse
 */
func (client *Client) SmsReport(request *SmsReportRequest) (_result *SmsReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SmsReportResponse{}
	_body, _err := client.SmsReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 短信上行。
 *
 * @param request SmsUpRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return SmsUpResponse
 */
func (client *Client) SmsUpWithOptions(request *SmsUpRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SmsUpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("SmsUp"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/recall/up"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SmsUpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 短信上行。
 *
 * @param request SmsUpRequest
 * @return SmsUpResponse
 */
func (client *Client) SmsUp(request *SmsUpRequest) (_result *SmsUpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SmsUpResponse{}
	_body, _err := client.SmsUpWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 更新运营活动
 *
 * @param request UpdateCampaignRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateCampaignResponse
 */
func (client *Client) UpdateCampaignWithOptions(Id *string, request *UpdateCampaignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCampaign"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/campaigns/" + tea.StringValue(openapiutil.GetEncodeParam(Id))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 更新运营活动
 *
 * @param request UpdateCampaignRequest
 * @return UpdateCampaignResponse
 */
func (client *Client) UpdateCampaign(Id *string, request *UpdateCampaignRequest) (_result *UpdateCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCampaignResponse{}
	_body, _err := client.UpdateCampaignWithOptions(Id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 更新回执Url。
 *
 * @param request UpdateReportUrlRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateReportUrlResponse
 */
func (client *Client) UpdateReportUrlWithOptions(request *UpdateReportUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateReportUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateReportUrl"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/users/reportUrl"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateReportUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 更新回执Url。
 *
 * @param request UpdateReportUrlRequest
 * @return UpdateReportUrlResponse
 */
func (client *Client) UpdateReportUrl(request *UpdateReportUrlRequest) (_result *UpdateReportUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateReportUrlResponse{}
	_body, _err := client.UpdateReportUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * 更新上行Url。
 *
 * @param request UpdateUploadUrlRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateUploadUrlResponse
 */
func (client *Client) UpdateUploadUrlWithOptions(request *UpdateUploadUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUploadUrl"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/users/uploadUrl"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * 更新上行Url。
 *
 * @param request UpdateUploadUrlRequest
 * @return UpdateUploadUrlResponse
 */
func (client *Client) UpdateUploadUrl(request *UpdateUploadUrlRequest) (_result *UpdateUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUploadUrlResponse{}
	_body, _err := client.UpdateUploadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
