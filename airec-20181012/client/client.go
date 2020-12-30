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

type AttachDatasetResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *AttachDatasetResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AttachDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDatasetResponseBody) SetMessage(v string) *AttachDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *AttachDatasetResponseBody) SetRequestId(v string) *AttachDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDatasetResponseBody) SetCode(v string) *AttachDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *AttachDatasetResponseBody) SetResult(v *AttachDatasetResponseBodyResult) *AttachDatasetResponseBody {
	s.Result = v
	return s
}

type AttachDatasetResponseBodyResult struct {
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s AttachDatasetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AttachDatasetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AttachDatasetResponseBodyResult) SetState(v string) *AttachDatasetResponseBodyResult {
	s.State = &v
	return s
}

func (s *AttachDatasetResponseBodyResult) SetGmtCreate(v int64) *AttachDatasetResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *AttachDatasetResponseBodyResult) SetVersionId(v string) *AttachDatasetResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *AttachDatasetResponseBodyResult) SetInstanceId(v string) *AttachDatasetResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *AttachDatasetResponseBodyResult) SetGmtModified(v int64) *AttachDatasetResponseBodyResult {
	s.GmtModified = &v
	return s
}

type AttachDatasetResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDatasetResponse) GoString() string {
	return s.String()
}

func (s *AttachDatasetResponse) SetHeaders(v map[string]*string) *AttachDatasetResponse {
	s.Headers = v
	return s
}

func (s *AttachDatasetResponse) SetBody(v *AttachDatasetResponseBody) *AttachDatasetResponse {
	s.Body = v
	return s
}

type CreateDiversifyResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *CreateDiversifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateDiversifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiversifyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiversifyResponseBody) SetMessage(v string) *CreateDiversifyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDiversifyResponseBody) SetRequestId(v string) *CreateDiversifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiversifyResponseBody) SetCode(v string) *CreateDiversifyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiversifyResponseBody) SetResult(v *CreateDiversifyResponseBodyResult) *CreateDiversifyResponseBody {
	s.Result = v
	return s
}

type CreateDiversifyResponseBodyResult struct {
	GmtCreate   *string                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *CreateDiversifyResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s CreateDiversifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDiversifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDiversifyResponseBodyResult) SetGmtCreate(v string) *CreateDiversifyResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreateDiversifyResponseBodyResult) SetName(v string) *CreateDiversifyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateDiversifyResponseBodyResult) SetGmtModified(v string) *CreateDiversifyResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CreateDiversifyResponseBodyResult) SetParameter(v *CreateDiversifyResponseBodyResultParameter) *CreateDiversifyResponseBodyResult {
	s.Parameter = v
	return s
}

type CreateDiversifyResponseBodyResultParameter struct {
	Window        *int32 `json:"Window,omitempty" xml:"Window,omitempty"`
	CategoryIndex *int32 `json:"CategoryIndex,omitempty" xml:"CategoryIndex,omitempty"`
}

func (s CreateDiversifyResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateDiversifyResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *CreateDiversifyResponseBodyResultParameter) SetWindow(v int32) *CreateDiversifyResponseBodyResultParameter {
	s.Window = &v
	return s
}

func (s *CreateDiversifyResponseBodyResultParameter) SetCategoryIndex(v int32) *CreateDiversifyResponseBodyResultParameter {
	s.CategoryIndex = &v
	return s
}

type CreateDiversifyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDiversifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiversifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiversifyResponse) GoString() string {
	return s.String()
}

func (s *CreateDiversifyResponse) SetHeaders(v map[string]*string) *CreateDiversifyResponse {
	s.Headers = v
	return s
}

func (s *CreateDiversifyResponse) SetBody(v *CreateDiversifyResponseBody) *CreateDiversifyResponse {
	s.Body = v
	return s
}

type CreateInstanceResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *CreateInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody {
	s.Result = v
	return s
}

type CreateInstanceResponseBodyResult struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyResult) SetInstanceId(v string) *CreateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateMixResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *CreateMixResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateMixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMixResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMixResponseBody) SetMessage(v string) *CreateMixResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMixResponseBody) SetRequestId(v string) *CreateMixResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMixResponseBody) SetCode(v string) *CreateMixResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMixResponseBody) SetResult(v *CreateMixResponseBodyResult) *CreateMixResponseBody {
	s.Result = v
	return s
}

type CreateMixResponseBodyResult struct {
	GmtCreate   *string                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *CreateMixResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s CreateMixResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateMixResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMixResponseBodyResult) SetGmtCreate(v string) *CreateMixResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreateMixResponseBodyResult) SetName(v string) *CreateMixResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateMixResponseBodyResult) SetGmtModified(v string) *CreateMixResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CreateMixResponseBodyResult) SetParameter(v *CreateMixResponseBodyResultParameter) *CreateMixResponseBodyResult {
	s.Parameter = v
	return s
}

type CreateMixResponseBodyResultParameter struct {
	Settings []*CreateMixResponseBodyResultParameterSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s CreateMixResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateMixResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *CreateMixResponseBodyResultParameter) SetSettings(v []*CreateMixResponseBodyResultParameterSettings) *CreateMixResponseBodyResultParameter {
	s.Settings = v
	return s
}

type CreateMixResponseBodyResultParameterSettings struct {
	Value *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMixResponseBodyResultParameterSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateMixResponseBodyResultParameterSettings) GoString() string {
	return s.String()
}

func (s *CreateMixResponseBodyResultParameterSettings) SetValue(v int32) *CreateMixResponseBodyResultParameterSettings {
	s.Value = &v
	return s
}

func (s *CreateMixResponseBodyResultParameterSettings) SetName(v string) *CreateMixResponseBodyResultParameterSettings {
	s.Name = &v
	return s
}

type CreateMixResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMixResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMixResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMixResponse) GoString() string {
	return s.String()
}

func (s *CreateMixResponse) SetHeaders(v map[string]*string) *CreateMixResponse {
	s.Headers = v
	return s
}

func (s *CreateMixResponse) SetBody(v *CreateMixResponseBody) *CreateMixResponse {
	s.Body = v
	return s
}

type CreateRuleResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetResult(v *CreateRuleResponseBodyResult) *CreateRuleResponseBody {
	s.Result = v
	return s
}

type CreateRuleResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBodyResult) SetStatus(v string) *CreateRuleResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateRuleResponseBodyResult) SetGmtCreate(v string) *CreateRuleResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreateRuleResponseBodyResult) SetGmtModified(v string) *CreateRuleResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *CreateRuleResponseBodyResult) SetRuleId(v string) *CreateRuleResponseBodyResult {
	s.RuleId = &v
	return s
}

type CreateRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateSceneRequest struct {
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s CreateSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneRequest) SetDryRun(v bool) *CreateSceneRequest {
	s.DryRun = &v
	return s
}

type CreateSceneResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateSceneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneResponseBody) SetRequestId(v string) *CreateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneResponseBody) SetResult(v *CreateSceneResponseBodyResult) *CreateSceneResponseBody {
	s.Result = v
	return s
}

type CreateSceneResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s CreateSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSceneResponseBodyResult) SetStatus(v string) *CreateSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateSceneResponseBodyResult) SetGmtCreate(v string) *CreateSceneResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *CreateSceneResponseBodyResult) SetSceneId(v string) *CreateSceneResponseBodyResult {
	s.SceneId = &v
	return s
}

func (s *CreateSceneResponseBodyResult) SetGmtModified(v string) *CreateSceneResponseBodyResult {
	s.GmtModified = &v
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

type DeleteDataSetResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DeleteDataSetResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponseBody) SetMessage(v string) *DeleteDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetRequestId(v string) *DeleteDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetCode(v string) *DeleteDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetResult(v *DeleteDataSetResponseBodyResult) *DeleteDataSetResponseBody {
	s.Result = v
	return s
}

type DeleteDataSetResponseBodyResult struct {
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DeleteDataSetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponseBodyResult) SetState(v string) *DeleteDataSetResponseBodyResult {
	s.State = &v
	return s
}

func (s *DeleteDataSetResponseBodyResult) SetGmtCreate(v int64) *DeleteDataSetResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DeleteDataSetResponseBodyResult) SetVersionId(v string) *DeleteDataSetResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *DeleteDataSetResponseBodyResult) SetInstanceId(v string) *DeleteDataSetResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DeleteDataSetResponseBodyResult) SetGmtModified(v int64) *DeleteDataSetResponseBodyResult {
	s.GmtModified = &v
	return s
}

type DeleteDataSetResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponse) SetHeaders(v map[string]*string) *DeleteDataSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSetResponse) SetBody(v *DeleteDataSetResponseBody) *DeleteDataSetResponse {
	s.Body = v
	return s
}

type DeleteDiversifyResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DeleteDiversifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteDiversifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiversifyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiversifyResponseBody) SetMessage(v string) *DeleteDiversifyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDiversifyResponseBody) SetRequestId(v string) *DeleteDiversifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiversifyResponseBody) SetCode(v string) *DeleteDiversifyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDiversifyResponseBody) SetResult(v *DeleteDiversifyResponseBodyResult) *DeleteDiversifyResponseBody {
	s.Result = v
	return s
}

type DeleteDiversifyResponseBodyResult struct {
	GmtCreate   *string                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *DeleteDiversifyResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s DeleteDiversifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiversifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteDiversifyResponseBodyResult) SetGmtCreate(v string) *DeleteDiversifyResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DeleteDiversifyResponseBodyResult) SetName(v string) *DeleteDiversifyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DeleteDiversifyResponseBodyResult) SetGmtModified(v string) *DeleteDiversifyResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *DeleteDiversifyResponseBodyResult) SetParameter(v *DeleteDiversifyResponseBodyResultParameter) *DeleteDiversifyResponseBodyResult {
	s.Parameter = v
	return s
}

type DeleteDiversifyResponseBodyResultParameter struct {
	Window        *int32 `json:"Window,omitempty" xml:"Window,omitempty"`
	CategoryIndex *int32 `json:"CategoryIndex,omitempty" xml:"CategoryIndex,omitempty"`
}

func (s DeleteDiversifyResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiversifyResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *DeleteDiversifyResponseBodyResultParameter) SetWindow(v int32) *DeleteDiversifyResponseBodyResultParameter {
	s.Window = &v
	return s
}

func (s *DeleteDiversifyResponseBodyResultParameter) SetCategoryIndex(v int32) *DeleteDiversifyResponseBodyResultParameter {
	s.CategoryIndex = &v
	return s
}

type DeleteDiversifyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDiversifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDiversifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiversifyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiversifyResponse) SetHeaders(v map[string]*string) *DeleteDiversifyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiversifyResponse) SetBody(v *DeleteDiversifyResponseBody) *DeleteDiversifyResponse {
	s.Body = v
	return s
}

type DeleteMixResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DeleteMixResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteMixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMixResponseBody) SetMessage(v string) *DeleteMixResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMixResponseBody) SetRequestId(v string) *DeleteMixResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMixResponseBody) SetCode(v string) *DeleteMixResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMixResponseBody) SetResult(v *DeleteMixResponseBodyResult) *DeleteMixResponseBody {
	s.Result = v
	return s
}

type DeleteMixResponseBodyResult struct {
	GmtCreate   *string                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *DeleteMixResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s DeleteMixResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteMixResponseBodyResult) SetGmtCreate(v string) *DeleteMixResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DeleteMixResponseBodyResult) SetName(v string) *DeleteMixResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DeleteMixResponseBodyResult) SetGmtModified(v string) *DeleteMixResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *DeleteMixResponseBodyResult) SetParameter(v *DeleteMixResponseBodyResultParameter) *DeleteMixResponseBodyResult {
	s.Parameter = v
	return s
}

type DeleteMixResponseBodyResultParameter struct {
	Settings []*DeleteMixResponseBodyResultParameterSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s DeleteMixResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *DeleteMixResponseBodyResultParameter) SetSettings(v []*DeleteMixResponseBodyResultParameterSettings) *DeleteMixResponseBodyResultParameter {
	s.Settings = v
	return s
}

type DeleteMixResponseBodyResultParameterSettings struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteMixResponseBodyResultParameterSettings) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixResponseBodyResultParameterSettings) GoString() string {
	return s.String()
}

func (s *DeleteMixResponseBodyResultParameterSettings) SetValue(v string) *DeleteMixResponseBodyResultParameterSettings {
	s.Value = &v
	return s
}

func (s *DeleteMixResponseBodyResultParameterSettings) SetName(v string) *DeleteMixResponseBodyResultParameterSettings {
	s.Name = &v
	return s
}

type DeleteMixResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMixResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMixResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMixResponse) GoString() string {
	return s.String()
}

func (s *DeleteMixResponse) SetHeaders(v map[string]*string) *DeleteMixResponse {
	s.Headers = v
	return s
}

func (s *DeleteMixResponse) SetBody(v *DeleteMixResponseBody) *DeleteMixResponse {
	s.Body = v
	return s
}

type DeleteSceneResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteSceneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponseBody) SetRequestId(v string) *DeleteSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSceneResponseBody) SetResult(v *DeleteSceneResponseBodyResult) *DeleteSceneResponseBody {
	s.Result = v
	return s
}

type DeleteSceneResponseBodyResult struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponseBodyResult) SetSceneId(v string) *DeleteSceneResponseBodyResult {
	s.SceneId = &v
	return s
}

type DeleteSceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponse) SetHeaders(v map[string]*string) *DeleteSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneResponse) SetBody(v *DeleteSceneResponseBody) *DeleteSceneResponse {
	s.Body = v
	return s
}

type DescribeDataSetMessageResponseBody struct {
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*DescribeDataSetMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeDataSetMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSetMessageResponseBody) SetMessage(v string) *DescribeDataSetMessageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDataSetMessageResponseBody) SetRequestId(v string) *DescribeDataSetMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSetMessageResponseBody) SetCode(v string) *DescribeDataSetMessageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDataSetMessageResponseBody) SetResult(v []*DescribeDataSetMessageResponseBodyResult) *DescribeDataSetMessageResponseBody {
	s.Result = v
	return s
}

type DescribeDataSetMessageResponseBodyResult struct {
	ErrorLevel *string `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Timestamp  *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ErrorType  *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
}

func (s DescribeDataSetMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDataSetMessageResponseBodyResult) SetErrorLevel(v string) *DescribeDataSetMessageResponseBodyResult {
	s.ErrorLevel = &v
	return s
}

func (s *DescribeDataSetMessageResponseBodyResult) SetMessage(v string) *DescribeDataSetMessageResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DescribeDataSetMessageResponseBodyResult) SetTimestamp(v string) *DescribeDataSetMessageResponseBodyResult {
	s.Timestamp = &v
	return s
}

func (s *DescribeDataSetMessageResponseBodyResult) SetErrorType(v string) *DescribeDataSetMessageResponseBodyResult {
	s.ErrorType = &v
	return s
}

type DescribeDataSetMessageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataSetMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataSetMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetMessageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSetMessageResponse) SetHeaders(v map[string]*string) *DescribeDataSetMessageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSetMessageResponse) SetBody(v *DescribeDataSetMessageResponseBody) *DescribeDataSetMessageResponse {
	s.Body = v
	return s
}

type DescribeDataSetReportResponseBody struct {
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DescribeDataSetReportResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDataSetReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSetReportResponseBody) SetMessage(v string) *DescribeDataSetReportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDataSetReportResponseBody) SetRequestId(v string) *DescribeDataSetReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSetReportResponseBody) SetCode(v string) *DescribeDataSetReportResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDataSetReportResponseBody) SetResult(v *DescribeDataSetReportResponseBodyResult) *DescribeDataSetReportResponseBody {
	s.Result = v
	return s
}

type DescribeDataSetReportResponseBodyResult struct {
	Overall *DescribeDataSetReportResponseBodyResultOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	Detail  []*DescribeDataSetReportResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeDataSetReportResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDataSetReportResponseBodyResult) SetOverall(v *DescribeDataSetReportResponseBodyResultOverall) *DescribeDataSetReportResponseBodyResult {
	s.Overall = v
	return s
}

func (s *DescribeDataSetReportResponseBodyResult) SetDetail(v []*DescribeDataSetReportResponseBodyResultDetail) *DescribeDataSetReportResponseBodyResult {
	s.Detail = v
	return s
}

type DescribeDataSetReportResponseBodyResultOverall struct {
	UserRepetitiveRate *float32 `json:"UserRepetitiveRate,omitempty" xml:"UserRepetitiveRate,omitempty"`
	UserUserCount      *int32   `json:"UserUserCount,omitempty" xml:"UserUserCount,omitempty"`
	BhvCount           *int32   `json:"BhvCount,omitempty" xml:"BhvCount,omitempty"`
	BhvLegalRate       *float32 `json:"BhvLegalRate,omitempty" xml:"BhvLegalRate,omitempty"`
	UserCompleteRate   *float32 `json:"UserCompleteRate,omitempty" xml:"UserCompleteRate,omitempty"`
	UserLoginRate      *float32 `json:"UserLoginRate,omitempty" xml:"UserLoginRate,omitempty"`
	ItemCompleteRate   *float32 `json:"ItemCompleteRate,omitempty" xml:"ItemCompleteRate,omitempty"`
	ItemRepetitiveRate *float32 `json:"ItemRepetitiveRate,omitempty" xml:"ItemRepetitiveRate,omitempty"`
	ItemItemCount      *int32   `json:"ItemItemCount,omitempty" xml:"ItemItemCount,omitempty"`
	UserLegalRate      *float32 `json:"UserLegalRate,omitempty" xml:"UserLegalRate,omitempty"`
	ItemLegalRate      *float32 `json:"ItemLegalRate,omitempty" xml:"ItemLegalRate,omitempty"`
	ItemLoginRate      *float32 `json:"ItemLoginRate,omitempty" xml:"ItemLoginRate,omitempty"`
}

func (s DescribeDataSetReportResponseBodyResultOverall) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetReportResponseBodyResultOverall) GoString() string {
	return s.String()
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetUserRepetitiveRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.UserRepetitiveRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetUserUserCount(v int32) *DescribeDataSetReportResponseBodyResultOverall {
	s.UserUserCount = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetBhvCount(v int32) *DescribeDataSetReportResponseBodyResultOverall {
	s.BhvCount = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetBhvLegalRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.BhvLegalRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetUserCompleteRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.UserCompleteRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetUserLoginRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.UserLoginRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetItemCompleteRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.ItemCompleteRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetItemRepetitiveRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.ItemRepetitiveRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetItemItemCount(v int32) *DescribeDataSetReportResponseBodyResultOverall {
	s.ItemItemCount = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetUserLegalRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.UserLegalRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetItemLegalRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.ItemLegalRate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultOverall) SetItemLoginRate(v float32) *DescribeDataSetReportResponseBodyResultOverall {
	s.ItemLoginRate = &v
	return s
}

type DescribeDataSetReportResponseBodyResultDetail struct {
	UvCtr      *float32 `json:"UvCtr,omitempty" xml:"UvCtr,omitempty"`
	BizDate    *int64   `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	PerUvClick *float32 `json:"PerUvClick,omitempty" xml:"PerUvClick,omitempty"`
	Pv         *int64   `json:"Pv,omitempty" xml:"Pv,omitempty"`
	ActiveItem *int64   `json:"ActiveItem,omitempty" xml:"ActiveItem,omitempty"`
	Ctr        *float32 `json:"Ctr,omitempty" xml:"Ctr,omitempty"`
	PerUvBhv   *float32 `json:"PerUvBhv,omitempty" xml:"PerUvBhv,omitempty"`
	Click      *int64   `json:"Click,omitempty" xml:"Click,omitempty"`
	Uv         *int64   `json:"Uv,omitempty" xml:"Uv,omitempty"`
	ClickUser  *int64   `json:"ClickUser,omitempty" xml:"ClickUser,omitempty"`
}

func (s DescribeDataSetReportResponseBodyResultDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetReportResponseBodyResultDetail) GoString() string {
	return s.String()
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetUvCtr(v float32) *DescribeDataSetReportResponseBodyResultDetail {
	s.UvCtr = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetBizDate(v int64) *DescribeDataSetReportResponseBodyResultDetail {
	s.BizDate = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetPerUvClick(v float32) *DescribeDataSetReportResponseBodyResultDetail {
	s.PerUvClick = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetPv(v int64) *DescribeDataSetReportResponseBodyResultDetail {
	s.Pv = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetActiveItem(v int64) *DescribeDataSetReportResponseBodyResultDetail {
	s.ActiveItem = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetCtr(v float32) *DescribeDataSetReportResponseBodyResultDetail {
	s.Ctr = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetPerUvBhv(v float32) *DescribeDataSetReportResponseBodyResultDetail {
	s.PerUvBhv = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetClick(v int64) *DescribeDataSetReportResponseBodyResultDetail {
	s.Click = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetUv(v int64) *DescribeDataSetReportResponseBodyResultDetail {
	s.Uv = &v
	return s
}

func (s *DescribeDataSetReportResponseBodyResultDetail) SetClickUser(v int64) *DescribeDataSetReportResponseBodyResultDetail {
	s.ClickUser = &v
	return s
}

type DescribeDataSetReportResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataSetReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataSetReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataSetReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSetReportResponse) SetHeaders(v map[string]*string) *DescribeDataSetReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSetReportResponse) SetBody(v *DescribeDataSetReportResponseBody) *DescribeDataSetReportResponse {
	s.Body = v
	return s
}

type DescribeDiversifyResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DescribeDiversifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDiversifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiversifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiversifyResponseBody) SetMessage(v string) *DescribeDiversifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDiversifyResponseBody) SetRequestId(v string) *DescribeDiversifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiversifyResponseBody) SetCode(v string) *DescribeDiversifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDiversifyResponseBody) SetResult(v *DescribeDiversifyResponseBodyResult) *DescribeDiversifyResponseBody {
	s.Result = v
	return s
}

type DescribeDiversifyResponseBodyResult struct {
	GmtCreate   *string                                       `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                                       `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *DescribeDiversifyResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s DescribeDiversifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiversifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDiversifyResponseBodyResult) SetGmtCreate(v string) *DescribeDiversifyResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDiversifyResponseBodyResult) SetName(v string) *DescribeDiversifyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDiversifyResponseBodyResult) SetGmtModified(v string) *DescribeDiversifyResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeDiversifyResponseBodyResult) SetParameter(v *DescribeDiversifyResponseBodyResultParameter) *DescribeDiversifyResponseBodyResult {
	s.Parameter = v
	return s
}

type DescribeDiversifyResponseBodyResultParameter struct {
	Window        *int32 `json:"Window,omitempty" xml:"Window,omitempty"`
	CategoryIndex *int32 `json:"CategoryIndex,omitempty" xml:"CategoryIndex,omitempty"`
}

func (s DescribeDiversifyResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiversifyResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *DescribeDiversifyResponseBodyResultParameter) SetWindow(v int32) *DescribeDiversifyResponseBodyResultParameter {
	s.Window = &v
	return s
}

func (s *DescribeDiversifyResponseBodyResultParameter) SetCategoryIndex(v int32) *DescribeDiversifyResponseBodyResultParameter {
	s.CategoryIndex = &v
	return s
}

type DescribeDiversifyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiversifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiversifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiversifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiversifyResponse) SetHeaders(v map[string]*string) *DescribeDiversifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiversifyResponse) SetBody(v *DescribeDiversifyResponseBody) *DescribeDiversifyResponse {
	s.Body = v
	return s
}

type DescribeExposureSettingsResponseBody struct {
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DescribeExposureSettingsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeExposureSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposureSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposureSettingsResponseBody) SetMessage(v string) *DescribeExposureSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExposureSettingsResponseBody) SetRequestId(v string) *DescribeExposureSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposureSettingsResponseBody) SetCode(v string) *DescribeExposureSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExposureSettingsResponseBody) SetResult(v *DescribeExposureSettingsResponseBodyResult) *DescribeExposureSettingsResponseBody {
	s.Result = v
	return s
}

type DescribeExposureSettingsResponseBodyResult struct {
	DurationSeconds *int32 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	ScenarioBased   *bool  `json:"ScenarioBased,omitempty" xml:"ScenarioBased,omitempty"`
}

func (s DescribeExposureSettingsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposureSettingsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeExposureSettingsResponseBodyResult) SetDurationSeconds(v int32) *DescribeExposureSettingsResponseBodyResult {
	s.DurationSeconds = &v
	return s
}

func (s *DescribeExposureSettingsResponseBodyResult) SetScenarioBased(v bool) *DescribeExposureSettingsResponseBodyResult {
	s.ScenarioBased = &v
	return s
}

type DescribeExposureSettingsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExposureSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExposureSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposureSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposureSettingsResponse) SetHeaders(v map[string]*string) *DescribeExposureSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposureSettingsResponse) SetBody(v *DescribeExposureSettingsResponseBody) *DescribeExposureSettingsResponse {
	s.Body = v
	return s
}

type DescribeInstanceResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DescribeInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetMessage(v string) *DescribeInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCode(v string) *DescribeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetResult(v *DescribeInstanceResponseBodyResult) *DescribeInstanceResponseBody {
	s.Result = v
	return s
}

type DescribeInstanceResponseBodyResult struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LockMode       *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DataSetVersion *string `json:"DataSetVersion,omitempty" xml:"DataSetVersion,omitempty"`
	Industry       *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s DescribeInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResult) SetStatus(v string) *DescribeInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetType(v string) *DescribeInstanceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetCommodityCode(v string) *DescribeInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetChargeType(v string) *DescribeInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetInstanceId(v string) *DescribeInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetGmtModified(v string) *DescribeInstanceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetLockMode(v string) *DescribeInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetRegionId(v string) *DescribeInstanceResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDataSetVersion(v string) *DescribeInstanceResponseBodyResult {
	s.DataSetVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetIndustry(v string) *DescribeInstanceResponseBodyResult {
	s.Industry = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetExpiredTime(v string) *DescribeInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetGmtCreate(v string) *DescribeInstanceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetName(v string) *DescribeInstanceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetScene(v string) *DescribeInstanceResponseBodyResult {
	s.Scene = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeMixResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DescribeMixResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeMixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMixResponseBody) SetMessage(v string) *DescribeMixResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMixResponseBody) SetRequestId(v string) *DescribeMixResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMixResponseBody) SetCode(v string) *DescribeMixResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMixResponseBody) SetResult(v *DescribeMixResponseBodyResult) *DescribeMixResponseBody {
	s.Result = v
	return s
}

type DescribeMixResponseBodyResult struct {
	GmtCreate   *string                                 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *DescribeMixResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s DescribeMixResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeMixResponseBodyResult) SetGmtCreate(v string) *DescribeMixResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMixResponseBodyResult) SetName(v string) *DescribeMixResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeMixResponseBodyResult) SetGmtModified(v string) *DescribeMixResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeMixResponseBodyResult) SetParameter(v *DescribeMixResponseBodyResultParameter) *DescribeMixResponseBodyResult {
	s.Parameter = v
	return s
}

type DescribeMixResponseBodyResultParameter struct {
	Settings []*DescribeMixResponseBodyResultParameterSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s DescribeMixResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *DescribeMixResponseBodyResultParameter) SetSettings(v []*DescribeMixResponseBodyResultParameterSettings) *DescribeMixResponseBodyResultParameter {
	s.Settings = v
	return s
}

type DescribeMixResponseBodyResultParameterSettings struct {
	Value *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMixResponseBodyResultParameterSettings) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixResponseBodyResultParameterSettings) GoString() string {
	return s.String()
}

func (s *DescribeMixResponseBodyResultParameterSettings) SetValue(v int32) *DescribeMixResponseBodyResultParameterSettings {
	s.Value = &v
	return s
}

func (s *DescribeMixResponseBodyResultParameterSettings) SetName(v string) *DescribeMixResponseBodyResultParameterSettings {
	s.Name = &v
	return s
}

type DescribeMixResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMixResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMixResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMixResponse) GoString() string {
	return s.String()
}

func (s *DescribeMixResponse) SetHeaders(v map[string]*string) *DescribeMixResponse {
	s.Headers = v
	return s
}

func (s *DescribeMixResponse) SetBody(v *DescribeMixResponseBody) *DescribeMixResponse {
	s.Body = v
	return s
}

type DescribeQuotaResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DescribeQuotaResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQuotaResponseBody) SetMessage(v string) *DescribeQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeQuotaResponseBody) SetRequestId(v string) *DescribeQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQuotaResponseBody) SetCode(v string) *DescribeQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeQuotaResponseBody) SetResult(v *DescribeQuotaResponseBodyResult) *DescribeQuotaResponseBody {
	s.Result = v
	return s
}

type DescribeQuotaResponseBodyResult struct {
	Qps           *int32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	ItemCount     *int64 `json:"ItemCount,omitempty" xml:"ItemCount,omitempty"`
	CurrentQps    *int32 `json:"CurrentQps,omitempty" xml:"CurrentQps,omitempty"`
	UserCount     *int64 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	UserCountUsed *int64 `json:"UserCountUsed,omitempty" xml:"UserCountUsed,omitempty"`
	ItemCountUsed *int64 `json:"ItemCountUsed,omitempty" xml:"ItemCountUsed,omitempty"`
}

func (s DescribeQuotaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeQuotaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeQuotaResponseBodyResult) SetQps(v int32) *DescribeQuotaResponseBodyResult {
	s.Qps = &v
	return s
}

func (s *DescribeQuotaResponseBodyResult) SetItemCount(v int64) *DescribeQuotaResponseBodyResult {
	s.ItemCount = &v
	return s
}

func (s *DescribeQuotaResponseBodyResult) SetCurrentQps(v int32) *DescribeQuotaResponseBodyResult {
	s.CurrentQps = &v
	return s
}

func (s *DescribeQuotaResponseBodyResult) SetUserCount(v int64) *DescribeQuotaResponseBodyResult {
	s.UserCount = &v
	return s
}

func (s *DescribeQuotaResponseBodyResult) SetUserCountUsed(v int64) *DescribeQuotaResponseBodyResult {
	s.UserCountUsed = &v
	return s
}

func (s *DescribeQuotaResponseBodyResult) SetItemCountUsed(v int64) *DescribeQuotaResponseBodyResult {
	s.ItemCountUsed = &v
	return s
}

type DescribeQuotaResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeQuotaResponse) SetHeaders(v map[string]*string) *DescribeQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeQuotaResponse) SetBody(v *DescribeQuotaResponseBody) *DescribeQuotaResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*DescribeRegionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody {
	s.Result = v
	return s
}

type DescribeRegionsResponseBodyResult struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ConsoleUrl *string `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	LocalName  *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	Endpoint   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) SetStatus(v string) *DescribeRegionsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetConsoleUrl(v string) *DescribeRegionsResponseBodyResult {
	s.ConsoleUrl = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRuleRequest struct {
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleRequest) SetSceneId(v string) *DescribeRuleRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeRuleRequest) SetRuleType(v string) *DescribeRuleRequest {
	s.RuleType = &v
	return s
}

type DescribeRuleResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleResponseBody) SetRequestId(v string) *DescribeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleResponseBody) SetResult(v *DescribeRuleResponseBodyResult) *DescribeRuleResponseBody {
	s.Result = v
	return s
}

type DescribeRuleResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRuleResponseBodyResult) SetStatus(v string) *DescribeRuleResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeRuleResponseBodyResult) SetGmtCreate(v string) *DescribeRuleResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRuleResponseBodyResult) SetGmtModified(v string) *DescribeRuleResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleResponseBodyResult) SetRuleId(v string) *DescribeRuleResponseBodyResult {
	s.RuleId = &v
	return s
}

type DescribeRuleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleResponse) SetHeaders(v map[string]*string) *DescribeRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleResponse) SetBody(v *DescribeRuleResponseBody) *DescribeRuleResponse {
	s.Body = v
	return s
}

type DescribeSceneResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeSceneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneResponseBody) SetRequestId(v string) *DescribeSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneResponseBody) SetResult(v *DescribeSceneResponseBodyResult) *DescribeSceneResponseBody {
	s.Result = v
	return s
}

type DescribeSceneResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSceneResponseBodyResult) SetStatus(v string) *DescribeSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeSceneResponseBodyResult) SetGmtCreate(v string) *DescribeSceneResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSceneResponseBodyResult) SetSceneId(v string) *DescribeSceneResponseBodyResult {
	s.SceneId = &v
	return s
}

func (s *DescribeSceneResponseBodyResult) SetGmtModified(v string) *DescribeSceneResponseBodyResult {
	s.GmtModified = &v
	return s
}

type DescribeSceneResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneResponse) SetHeaders(v map[string]*string) *DescribeSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneResponse) SetBody(v *DescribeSceneResponseBody) *DescribeSceneResponse {
	s.Body = v
	return s
}

type DescribeSceneThroughputResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeSceneThroughputResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeSceneThroughputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneThroughputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneThroughputResponseBody) SetRequestId(v string) *DescribeSceneThroughputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneThroughputResponseBody) SetResult(v *DescribeSceneThroughputResponseBodyResult) *DescribeSceneThroughputResponseBody {
	s.Result = v
	return s
}

type DescribeSceneThroughputResponseBodyResult struct {
	PvCount *int64 `json:"PvCount,omitempty" xml:"PvCount,omitempty"`
}

func (s DescribeSceneThroughputResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneThroughputResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSceneThroughputResponseBodyResult) SetPvCount(v int64) *DescribeSceneThroughputResponseBodyResult {
	s.PvCount = &v
	return s
}

type DescribeSceneThroughputResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSceneThroughputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSceneThroughputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneThroughputResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneThroughputResponse) SetHeaders(v map[string]*string) *DescribeSceneThroughputResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneThroughputResponse) SetBody(v *DescribeSceneThroughputResponseBody) *DescribeSceneThroughputResponse {
	s.Body = v
	return s
}

type DescribeSyncReportDetailRequest struct {
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	LevelType *string `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
}

func (s DescribeSyncReportDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportDetailRequest) SetStartTime(v int64) *DescribeSyncReportDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncReportDetailRequest) SetEndTime(v int64) *DescribeSyncReportDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSyncReportDetailRequest) SetType(v string) *DescribeSyncReportDetailRequest {
	s.Type = &v
	return s
}

func (s *DescribeSyncReportDetailRequest) SetLevelType(v string) *DescribeSyncReportDetailRequest {
	s.LevelType = &v
	return s
}

type DescribeSyncReportDetailResponseBody struct {
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*DescribeSyncReportDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeSyncReportDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportDetailResponseBody) SetMessage(v string) *DescribeSyncReportDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBody) SetRequestId(v string) *DescribeSyncReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBody) SetCode(v string) *DescribeSyncReportDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBody) SetResult(v []*DescribeSyncReportDetailResponseBodyResult) *DescribeSyncReportDetailResponseBody {
	s.Result = v
	return s
}

type DescribeSyncReportDetailResponseBodyResult struct {
	Type           *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	ErrorCount     *int32                                                   `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	SampleDisplay  *bool                                                    `json:"SampleDisplay,omitempty" xml:"SampleDisplay,omitempty"`
	HistoryData    []*DescribeSyncReportDetailResponseBodyResultHistoryData `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
	DefaultDisplay *bool                                                    `json:"DefaultDisplay,omitempty" xml:"DefaultDisplay,omitempty"`
	ErrorPercent   *float32                                                 `json:"ErrorPercent,omitempty" xml:"ErrorPercent,omitempty"`
}

func (s DescribeSyncReportDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportDetailResponseBodyResult) SetType(v string) *DescribeSyncReportDetailResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResult) SetErrorCount(v int32) *DescribeSyncReportDetailResponseBodyResult {
	s.ErrorCount = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResult) SetSampleDisplay(v bool) *DescribeSyncReportDetailResponseBodyResult {
	s.SampleDisplay = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResult) SetHistoryData(v []*DescribeSyncReportDetailResponseBodyResultHistoryData) *DescribeSyncReportDetailResponseBodyResult {
	s.HistoryData = v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResult) SetDefaultDisplay(v bool) *DescribeSyncReportDetailResponseBodyResult {
	s.DefaultDisplay = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResult) SetErrorPercent(v float32) *DescribeSyncReportDetailResponseBodyResult {
	s.ErrorPercent = &v
	return s
}

type DescribeSyncReportDetailResponseBodyResultHistoryData struct {
	EndTime      *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ErrorPercent *float32 `json:"ErrorPercent,omitempty" xml:"ErrorPercent,omitempty"`
}

func (s DescribeSyncReportDetailResponseBodyResultHistoryData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportDetailResponseBodyResultHistoryData) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportDetailResponseBodyResultHistoryData) SetEndTime(v int64) *DescribeSyncReportDetailResponseBodyResultHistoryData {
	s.EndTime = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResultHistoryData) SetStartTime(v int64) *DescribeSyncReportDetailResponseBodyResultHistoryData {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncReportDetailResponseBodyResultHistoryData) SetErrorPercent(v float32) *DescribeSyncReportDetailResponseBodyResultHistoryData {
	s.ErrorPercent = &v
	return s
}

type DescribeSyncReportDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSyncReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSyncReportDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportDetailResponse) SetHeaders(v map[string]*string) *DescribeSyncReportDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncReportDetailResponse) SetBody(v *DescribeSyncReportDetailResponseBody) *DescribeSyncReportDetailResponse {
	s.Body = v
	return s
}

type DescribeSyncReportOutliersRequest struct {
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LevelType *string `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
}

func (s DescribeSyncReportOutliersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportOutliersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportOutliersRequest) SetStartTime(v int64) *DescribeSyncReportOutliersRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncReportOutliersRequest) SetKey(v string) *DescribeSyncReportOutliersRequest {
	s.Key = &v
	return s
}

func (s *DescribeSyncReportOutliersRequest) SetType(v string) *DescribeSyncReportOutliersRequest {
	s.Type = &v
	return s
}

func (s *DescribeSyncReportOutliersRequest) SetEndTime(v int64) *DescribeSyncReportOutliersRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSyncReportOutliersRequest) SetLevelType(v string) *DescribeSyncReportOutliersRequest {
	s.LevelType = &v
	return s
}

type DescribeSyncReportOutliersResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeSyncReportOutliersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportOutliersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportOutliersResponseBody) SetMessage(v string) *DescribeSyncReportOutliersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSyncReportOutliersResponseBody) SetRequestId(v string) *DescribeSyncReportOutliersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncReportOutliersResponseBody) SetCode(v string) *DescribeSyncReportOutliersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSyncReportOutliersResponseBody) SetResult(v map[string]interface{}) *DescribeSyncReportOutliersResponseBody {
	s.Result = v
	return s
}

type DescribeSyncReportOutliersResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSyncReportOutliersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSyncReportOutliersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncReportOutliersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncReportOutliersResponse) SetHeaders(v map[string]*string) *DescribeSyncReportOutliersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncReportOutliersResponse) SetBody(v *DescribeSyncReportOutliersResponseBody) *DescribeSyncReportOutliersResponse {
	s.Body = v
	return s
}

type DescribeUserMetricsRequest struct {
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s DescribeUserMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserMetricsRequest) SetStartTime(v int64) *DescribeUserMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserMetricsRequest) SetEndTime(v int64) *DescribeUserMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserMetricsRequest) SetMetricType(v string) *DescribeUserMetricsRequest {
	s.MetricType = &v
	return s
}

type DescribeUserMetricsResponseBody struct {
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*DescribeUserMetricsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeUserMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserMetricsResponseBody) SetMessage(v string) *DescribeUserMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserMetricsResponseBody) SetRequestId(v string) *DescribeUserMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserMetricsResponseBody) SetCode(v string) *DescribeUserMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserMetricsResponseBody) SetResult(v []*DescribeUserMetricsResponseBodyResult) *DescribeUserMetricsResponseBody {
	s.Result = v
	return s
}

type DescribeUserMetricsResponseBodyResult struct {
	DataPoints []*DescribeUserMetricsResponseBodyResultDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	SceneId    *string                                            `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeUserMetricsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserMetricsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeUserMetricsResponseBodyResult) SetDataPoints(v []*DescribeUserMetricsResponseBodyResultDataPoints) *DescribeUserMetricsResponseBodyResult {
	s.DataPoints = v
	return s
}

func (s *DescribeUserMetricsResponseBodyResult) SetSceneId(v string) *DescribeUserMetricsResponseBodyResult {
	s.SceneId = &v
	return s
}

type DescribeUserMetricsResponseBodyResultDataPoints struct {
	EndTime   *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Val       *float32 `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s DescribeUserMetricsResponseBodyResultDataPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserMetricsResponseBodyResultDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeUserMetricsResponseBodyResultDataPoints) SetEndTime(v int64) *DescribeUserMetricsResponseBodyResultDataPoints {
	s.EndTime = &v
	return s
}

func (s *DescribeUserMetricsResponseBodyResultDataPoints) SetStartTime(v int64) *DescribeUserMetricsResponseBodyResultDataPoints {
	s.StartTime = &v
	return s
}

func (s *DescribeUserMetricsResponseBodyResultDataPoints) SetVal(v float32) *DescribeUserMetricsResponseBodyResultDataPoints {
	s.Val = &v
	return s
}

type DescribeUserMetricsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserMetricsResponse) SetHeaders(v map[string]*string) *DescribeUserMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserMetricsResponse) SetBody(v *DescribeUserMetricsResponseBody) *DescribeUserMetricsResponse {
	s.Body = v
	return s
}

type DowngradeInstanceResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *DowngradeInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DowngradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DowngradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradeInstanceResponseBody) SetMessage(v string) *DowngradeInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DowngradeInstanceResponseBody) SetRequestId(v string) *DowngradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradeInstanceResponseBody) SetCode(v string) *DowngradeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DowngradeInstanceResponseBody) SetResult(v *DowngradeInstanceResponseBodyResult) *DowngradeInstanceResponseBody {
	s.Result = v
	return s
}

type DowngradeInstanceResponseBodyResult struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DowngradeInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DowngradeInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DowngradeInstanceResponseBodyResult) SetInstanceId(v string) *DowngradeInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

type DowngradeInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DowngradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DowngradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DowngradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DowngradeInstanceResponse) SetHeaders(v map[string]*string) *DowngradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DowngradeInstanceResponse) SetBody(v *DowngradeInstanceResponseBody) *DowngradeInstanceResponse {
	s.Body = v
	return s
}

type ListDashboardRequest struct {
	StartDate *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate   *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardRequest) SetStartDate(v int64) *ListDashboardRequest {
	s.StartDate = &v
	return s
}

func (s *ListDashboardRequest) SetEndDate(v int64) *ListDashboardRequest {
	s.EndDate = &v
	return s
}

func (s *ListDashboardRequest) SetTraceId(v string) *ListDashboardRequest {
	s.TraceId = &v
	return s
}

func (s *ListDashboardRequest) SetSceneId(v string) *ListDashboardRequest {
	s.SceneId = &v
	return s
}

func (s *ListDashboardRequest) SetPage(v int32) *ListDashboardRequest {
	s.Page = &v
	return s
}

func (s *ListDashboardRequest) SetSize(v int32) *ListDashboardRequest {
	s.Size = &v
	return s
}

type ListDashboardResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ListDashboardResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardResponseBody) SetMessage(v string) *ListDashboardResponseBody {
	s.Message = &v
	return s
}

func (s *ListDashboardResponseBody) SetRequestId(v string) *ListDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardResponseBody) SetCode(v string) *ListDashboardResponseBody {
	s.Code = &v
	return s
}

func (s *ListDashboardResponseBody) SetResult(v *ListDashboardResponseBodyResult) *ListDashboardResponseBody {
	s.Result = v
	return s
}

type ListDashboardResponseBodyResult struct {
	List []*ListDashboardResponseBodyResultList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Num  *int32                                 `json:"Num,omitempty" xml:"Num,omitempty"`
}

func (s ListDashboardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardResponseBodyResult) SetList(v []*ListDashboardResponseBodyResultList) *ListDashboardResponseBodyResult {
	s.List = v
	return s
}

func (s *ListDashboardResponseBodyResult) SetNum(v int32) *ListDashboardResponseBodyResult {
	s.Num = &v
	return s
}

type ListDashboardResponseBodyResultList struct {
	UvCtr      *float32 `json:"UvCtr,omitempty" xml:"UvCtr,omitempty"`
	BizDate    *int64   `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	ActiveItem *int64   `json:"ActiveItem,omitempty" xml:"ActiveItem,omitempty"`
	PerUvBhv   *float32 `json:"PerUvBhv,omitempty" xml:"PerUvBhv,omitempty"`
	Uv         *int64   `json:"Uv,omitempty" xml:"Uv,omitempty"`
	PerUvClick *float32 `json:"PerUvClick,omitempty" xml:"PerUvClick,omitempty"`
	Pv         *int64   `json:"Pv,omitempty" xml:"Pv,omitempty"`
	Ctr        *float32 `json:"Ctr,omitempty" xml:"Ctr,omitempty"`
	SceneId    *string  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	TraceId    *string  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Click      *int64   `json:"Click,omitempty" xml:"Click,omitempty"`
	ClickUser  *int64   `json:"ClickUser,omitempty" xml:"ClickUser,omitempty"`
}

func (s ListDashboardResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *ListDashboardResponseBodyResultList) SetUvCtr(v float32) *ListDashboardResponseBodyResultList {
	s.UvCtr = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetBizDate(v int64) *ListDashboardResponseBodyResultList {
	s.BizDate = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetActiveItem(v int64) *ListDashboardResponseBodyResultList {
	s.ActiveItem = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetPerUvBhv(v float32) *ListDashboardResponseBodyResultList {
	s.PerUvBhv = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetUv(v int64) *ListDashboardResponseBodyResultList {
	s.Uv = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetPerUvClick(v float32) *ListDashboardResponseBodyResultList {
	s.PerUvClick = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetPv(v int64) *ListDashboardResponseBodyResultList {
	s.Pv = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetCtr(v float32) *ListDashboardResponseBodyResultList {
	s.Ctr = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetSceneId(v string) *ListDashboardResponseBodyResultList {
	s.SceneId = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetTraceId(v string) *ListDashboardResponseBodyResultList {
	s.TraceId = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetClick(v int64) *ListDashboardResponseBodyResultList {
	s.Click = &v
	return s
}

func (s *ListDashboardResponseBodyResultList) SetClickUser(v int64) *ListDashboardResponseBodyResultList {
	s.ClickUser = &v
	return s
}

type ListDashboardResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardResponse) SetHeaders(v map[string]*string) *ListDashboardResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardResponse) SetBody(v *ListDashboardResponseBody) *ListDashboardResponse {
	s.Body = v
	return s
}

type ListDashboardDetailsRequest struct {
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TraceIds   *string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty"`
	SceneIds   *string `json:"SceneIds,omitempty" xml:"SceneIds,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s ListDashboardDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsRequest) SetStartTime(v int64) *ListDashboardDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDashboardDetailsRequest) SetEndTime(v int64) *ListDashboardDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDashboardDetailsRequest) SetTraceIds(v string) *ListDashboardDetailsRequest {
	s.TraceIds = &v
	return s
}

func (s *ListDashboardDetailsRequest) SetSceneIds(v string) *ListDashboardDetailsRequest {
	s.SceneIds = &v
	return s
}

func (s *ListDashboardDetailsRequest) SetMetricType(v string) *ListDashboardDetailsRequest {
	s.MetricType = &v
	return s
}

type ListDashboardDetailsResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDashboardDetailsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDashboardDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsResponseBody) SetRequestId(v string) *ListDashboardDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardDetailsResponseBody) SetResult(v []*ListDashboardDetailsResponseBodyResult) *ListDashboardDetailsResponseBody {
	s.Result = v
	return s
}

type ListDashboardDetailsResponseBodyResult struct {
	MetricRes *ListDashboardDetailsResponseBodyResultMetricRes `json:"MetricRes,omitempty" xml:"MetricRes,omitempty" type:"Struct"`
	SceneId   *string                                          `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	TraceId   *string                                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListDashboardDetailsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsResponseBodyResult) SetMetricRes(v *ListDashboardDetailsResponseBodyResultMetricRes) *ListDashboardDetailsResponseBodyResult {
	s.MetricRes = v
	return s
}

func (s *ListDashboardDetailsResponseBodyResult) SetSceneId(v string) *ListDashboardDetailsResponseBodyResult {
	s.SceneId = &v
	return s
}

func (s *ListDashboardDetailsResponseBodyResult) SetTraceId(v string) *ListDashboardDetailsResponseBodyResult {
	s.TraceId = &v
	return s
}

type ListDashboardDetailsResponseBodyResultMetricRes struct {
	Total  map[string]interface{} `json:"Total,omitempty" xml:"Total,omitempty"`
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s ListDashboardDetailsResponseBodyResultMetricRes) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsResponseBodyResultMetricRes) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsResponseBodyResultMetricRes) SetTotal(v map[string]interface{}) *ListDashboardDetailsResponseBodyResultMetricRes {
	s.Total = v
	return s
}

func (s *ListDashboardDetailsResponseBodyResultMetricRes) SetDetail(v map[string]interface{}) *ListDashboardDetailsResponseBodyResultMetricRes {
	s.Detail = v
	return s
}

type ListDashboardDetailsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsResponse) SetHeaders(v map[string]*string) *ListDashboardDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardDetailsResponse) SetBody(v *ListDashboardDetailsResponseBody) *ListDashboardDetailsResponse {
	s.Body = v
	return s
}

type ListDashboardDetailsFlowsRequest struct {
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TraceIds   *string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty"`
	SceneIds   *string `json:"SceneIds,omitempty" xml:"SceneIds,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s ListDashboardDetailsFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsFlowsRequest) SetStartTime(v int64) *ListDashboardDetailsFlowsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDashboardDetailsFlowsRequest) SetEndTime(v int64) *ListDashboardDetailsFlowsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDashboardDetailsFlowsRequest) SetTraceIds(v string) *ListDashboardDetailsFlowsRequest {
	s.TraceIds = &v
	return s
}

func (s *ListDashboardDetailsFlowsRequest) SetSceneIds(v string) *ListDashboardDetailsFlowsRequest {
	s.SceneIds = &v
	return s
}

func (s *ListDashboardDetailsFlowsRequest) SetMetricType(v string) *ListDashboardDetailsFlowsRequest {
	s.MetricType = &v
	return s
}

type ListDashboardDetailsFlowsResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListDashboardDetailsFlowsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDashboardDetailsFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsFlowsResponseBody) SetRequestId(v string) *ListDashboardDetailsFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardDetailsFlowsResponseBody) SetResult(v *ListDashboardDetailsFlowsResponseBodyResult) *ListDashboardDetailsFlowsResponseBody {
	s.Result = v
	return s
}

type ListDashboardDetailsFlowsResponseBodyResult struct {
	MetricData []*ListDashboardDetailsFlowsResponseBodyResultMetricData `json:"MetricData,omitempty" xml:"MetricData,omitempty" type:"Repeated"`
	MetricType *string                                                  `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s ListDashboardDetailsFlowsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsFlowsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsFlowsResponseBodyResult) SetMetricData(v []*ListDashboardDetailsFlowsResponseBodyResultMetricData) *ListDashboardDetailsFlowsResponseBodyResult {
	s.MetricData = v
	return s
}

func (s *ListDashboardDetailsFlowsResponseBodyResult) SetMetricType(v string) *ListDashboardDetailsFlowsResponseBodyResult {
	s.MetricType = &v
	return s
}

type ListDashboardDetailsFlowsResponseBodyResultMetricData struct {
	MetricRes map[string]interface{} `json:"MetricRes,omitempty" xml:"MetricRes,omitempty"`
	SceneId   *string                `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	TraceId   *string                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListDashboardDetailsFlowsResponseBodyResultMetricData) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsFlowsResponseBodyResultMetricData) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsFlowsResponseBodyResultMetricData) SetMetricRes(v map[string]interface{}) *ListDashboardDetailsFlowsResponseBodyResultMetricData {
	s.MetricRes = v
	return s
}

func (s *ListDashboardDetailsFlowsResponseBodyResultMetricData) SetSceneId(v string) *ListDashboardDetailsFlowsResponseBodyResultMetricData {
	s.SceneId = &v
	return s
}

func (s *ListDashboardDetailsFlowsResponseBodyResultMetricData) SetTraceId(v string) *ListDashboardDetailsFlowsResponseBodyResultMetricData {
	s.TraceId = &v
	return s
}

type ListDashboardDetailsFlowsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardDetailsFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardDetailsFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardDetailsFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardDetailsFlowsResponse) SetHeaders(v map[string]*string) *ListDashboardDetailsFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardDetailsFlowsResponse) SetBody(v *ListDashboardDetailsFlowsResponseBody) *ListDashboardDetailsFlowsResponse {
	s.Body = v
	return s
}

type ListDashboardMetricsRequest struct {
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s ListDashboardMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsRequest) SetStartTime(v int64) *ListDashboardMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDashboardMetricsRequest) SetEndTime(v int64) *ListDashboardMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDashboardMetricsRequest) SetMetricType(v string) *ListDashboardMetricsRequest {
	s.MetricType = &v
	return s
}

type ListDashboardMetricsResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDashboardMetricsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDashboardMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsResponseBody) SetRequestId(v string) *ListDashboardMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardMetricsResponseBody) SetResult(v []*ListDashboardMetricsResponseBodyResult) *ListDashboardMetricsResponseBody {
	s.Result = v
	return s
}

type ListDashboardMetricsResponseBodyResult struct {
	Total  map[string]interface{}                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Detail []*ListDashboardMetricsResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s ListDashboardMetricsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsResponseBodyResult) SetTotal(v map[string]interface{}) *ListDashboardMetricsResponseBodyResult {
	s.Total = v
	return s
}

func (s *ListDashboardMetricsResponseBodyResult) SetDetail(v []*ListDashboardMetricsResponseBodyResultDetail) *ListDashboardMetricsResponseBodyResult {
	s.Detail = v
	return s
}

type ListDashboardMetricsResponseBodyResultDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Val       *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s ListDashboardMetricsResponseBodyResultDetail) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsResponseBodyResultDetail) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsResponseBodyResultDetail) SetEndTime(v string) *ListDashboardMetricsResponseBodyResultDetail {
	s.EndTime = &v
	return s
}

func (s *ListDashboardMetricsResponseBodyResultDetail) SetStartTime(v string) *ListDashboardMetricsResponseBodyResultDetail {
	s.StartTime = &v
	return s
}

func (s *ListDashboardMetricsResponseBodyResultDetail) SetVal(v string) *ListDashboardMetricsResponseBodyResultDetail {
	s.Val = &v
	return s
}

type ListDashboardMetricsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsResponse) SetHeaders(v map[string]*string) *ListDashboardMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardMetricsResponse) SetBody(v *ListDashboardMetricsResponseBody) *ListDashboardMetricsResponse {
	s.Body = v
	return s
}

type ListDashboardMetricsFlowsRequest struct {
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s ListDashboardMetricsFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsFlowsRequest) SetStartTime(v int64) *ListDashboardMetricsFlowsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDashboardMetricsFlowsRequest) SetEndTime(v int64) *ListDashboardMetricsFlowsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDashboardMetricsFlowsRequest) SetMetricType(v string) *ListDashboardMetricsFlowsRequest {
	s.MetricType = &v
	return s
}

type ListDashboardMetricsFlowsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDashboardMetricsFlowsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDashboardMetricsFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsFlowsResponseBody) SetRequestId(v string) *ListDashboardMetricsFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardMetricsFlowsResponseBody) SetResult(v []*ListDashboardMetricsFlowsResponseBodyResult) *ListDashboardMetricsFlowsResponseBody {
	s.Result = v
	return s
}

type ListDashboardMetricsFlowsResponseBodyResult struct {
	MetricData map[string]interface{} `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	MetricType *string                `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s ListDashboardMetricsFlowsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsFlowsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsFlowsResponseBodyResult) SetMetricData(v map[string]interface{}) *ListDashboardMetricsFlowsResponseBodyResult {
	s.MetricData = v
	return s
}

func (s *ListDashboardMetricsFlowsResponseBodyResult) SetMetricType(v string) *ListDashboardMetricsFlowsResponseBodyResult {
	s.MetricType = &v
	return s
}

type ListDashboardMetricsFlowsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardMetricsFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardMetricsFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardMetricsFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardMetricsFlowsResponse) SetHeaders(v map[string]*string) *ListDashboardMetricsFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardMetricsFlowsResponse) SetBody(v *ListDashboardMetricsFlowsResponseBody) *ListDashboardMetricsFlowsResponse {
	s.Body = v
	return s
}

type ListDashboardParametersResponseBody struct {
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ListDashboardParametersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDashboardParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardParametersResponseBody) SetMessage(v string) *ListDashboardParametersResponseBody {
	s.Message = &v
	return s
}

func (s *ListDashboardParametersResponseBody) SetRequestId(v string) *ListDashboardParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardParametersResponseBody) SetCode(v string) *ListDashboardParametersResponseBody {
	s.Code = &v
	return s
}

func (s *ListDashboardParametersResponseBody) SetResult(v *ListDashboardParametersResponseBodyResult) *ListDashboardParametersResponseBody {
	s.Result = v
	return s
}

type ListDashboardParametersResponseBodyResult struct {
	SceneId []*string `json:"SceneId,omitempty" xml:"SceneId,omitempty" type:"Repeated"`
	TraceId []*string `json:"TraceId,omitempty" xml:"TraceId,omitempty" type:"Repeated"`
}

func (s ListDashboardParametersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardParametersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardParametersResponseBodyResult) SetSceneId(v []*string) *ListDashboardParametersResponseBodyResult {
	s.SceneId = v
	return s
}

func (s *ListDashboardParametersResponseBodyResult) SetTraceId(v []*string) *ListDashboardParametersResponseBodyResult {
	s.TraceId = v
	return s
}

type ListDashboardParametersResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardParametersResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardParametersResponse) SetHeaders(v map[string]*string) *ListDashboardParametersResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardParametersResponse) SetBody(v *ListDashboardParametersResponseBody) *ListDashboardParametersResponse {
	s.Body = v
	return s
}

type ListDashboardUidResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ListDashboardUidResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDashboardUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardUidResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardUidResponseBody) SetMessage(v string) *ListDashboardUidResponseBody {
	s.Message = &v
	return s
}

func (s *ListDashboardUidResponseBody) SetRequestId(v string) *ListDashboardUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardUidResponseBody) SetCode(v string) *ListDashboardUidResponseBody {
	s.Code = &v
	return s
}

func (s *ListDashboardUidResponseBody) SetResult(v *ListDashboardUidResponseBodyResult) *ListDashboardUidResponseBody {
	s.Result = v
	return s
}

type ListDashboardUidResponseBodyResult struct {
	Num *int32    `json:"Num,omitempty" xml:"Num,omitempty"`
	Uid []*string `json:"Uid,omitempty" xml:"Uid,omitempty" type:"Repeated"`
}

func (s ListDashboardUidResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardUidResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDashboardUidResponseBodyResult) SetNum(v int32) *ListDashboardUidResponseBodyResult {
	s.Num = &v
	return s
}

func (s *ListDashboardUidResponseBodyResult) SetUid(v []*string) *ListDashboardUidResponseBodyResult {
	s.Uid = v
	return s
}

type ListDashboardUidResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDashboardUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardUidResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardUidResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardUidResponse) SetHeaders(v map[string]*string) *ListDashboardUidResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardUidResponse) SetBody(v *ListDashboardUidResponseBody) *ListDashboardUidResponse {
	s.Body = v
	return s
}

type ListDataSetResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListDataSetResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBody) SetMessage(v string) *ListDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataSetResponseBody) SetRequestId(v string) *ListDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSetResponseBody) SetCode(v string) *ListDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSetResponseBody) SetResult(v []*ListDataSetResponseBodyResult) *ListDataSetResponseBody {
	s.Result = v
	return s
}

type ListDataSetResponseBodyResult struct {
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s ListDataSetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSetResponseBodyResult) SetState(v string) *ListDataSetResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListDataSetResponseBodyResult) SetGmtCreate(v int64) *ListDataSetResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListDataSetResponseBodyResult) SetVersionId(v string) *ListDataSetResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *ListDataSetResponseBodyResult) SetInstanceId(v string) *ListDataSetResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListDataSetResponseBodyResult) SetGmtModified(v int64) *ListDataSetResponseBodyResult {
	s.GmtModified = &v
	return s
}

type ListDataSetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSetResponse) GoString() string {
	return s.String()
}

func (s *ListDataSetResponse) SetHeaders(v map[string]*string) *ListDataSetResponse {
	s.Headers = v
	return s
}

func (s *ListDataSetResponse) SetBody(v *ListDataSetResponseBody) *ListDataSetResponse {
	s.Body = v
	return s
}

type ListDataSourceResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListDataSourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponseBody) SetMessage(v string) *ListDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataSourceResponseBody) SetRequestId(v string) *ListDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceResponseBody) SetCode(v string) *ListDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSourceResponseBody) SetResult(v []*ListDataSourceResponseBodyResult) *ListDataSourceResponseBody {
	s.Result = v
	return s
}

type ListDataSourceResponseBodyResult struct {
	TableName   *string                               `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Meta        *ListDataSourceResponseBodyResultMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	GmtCreate   *string                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s ListDataSourceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponseBodyResult) SetTableName(v string) *ListDataSourceResponseBodyResult {
	s.TableName = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetMeta(v *ListDataSourceResponseBodyResultMeta) *ListDataSourceResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetGmtCreate(v string) *ListDataSourceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListDataSourceResponseBodyResult) SetGmtModified(v string) *ListDataSourceResponseBodyResult {
	s.GmtModified = &v
	return s
}

type ListDataSourceResponseBodyResultMeta struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	BucketName  *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Partition   *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
}

func (s ListDataSourceResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponseBodyResultMeta) SetType(v string) *ListDataSourceResponseBodyResultMeta {
	s.Type = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetTableName(v string) *ListDataSourceResponseBodyResultMeta {
	s.TableName = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetProjectName(v string) *ListDataSourceResponseBodyResultMeta {
	s.ProjectName = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetBucketName(v string) *ListDataSourceResponseBodyResultMeta {
	s.BucketName = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetPath(v string) *ListDataSourceResponseBodyResultMeta {
	s.Path = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetPartition(v string) *ListDataSourceResponseBodyResultMeta {
	s.Partition = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetTimestamp(v int64) *ListDataSourceResponseBodyResultMeta {
	s.Timestamp = &v
	return s
}

func (s *ListDataSourceResponseBodyResultMeta) SetAccessKeyId(v string) *ListDataSourceResponseBodyResultMeta {
	s.AccessKeyId = &v
	return s
}

type ListDataSourceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponse) SetHeaders(v map[string]*string) *ListDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceResponse) SetBody(v *ListDataSourceResponseBody) *ListDataSourceResponse {
	s.Body = v
	return s
}

type ListDiversifyResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListDiversifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiversifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiversifyResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiversifyResponseBody) SetMessage(v string) *ListDiversifyResponseBody {
	s.Message = &v
	return s
}

func (s *ListDiversifyResponseBody) SetRequestId(v string) *ListDiversifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiversifyResponseBody) SetCode(v string) *ListDiversifyResponseBody {
	s.Code = &v
	return s
}

func (s *ListDiversifyResponseBody) SetResult(v []*ListDiversifyResponseBodyResult) *ListDiversifyResponseBody {
	s.Result = v
	return s
}

type ListDiversifyResponseBodyResult struct {
	GmtCreate   *string                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *ListDiversifyResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s ListDiversifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDiversifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDiversifyResponseBodyResult) SetGmtCreate(v string) *ListDiversifyResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListDiversifyResponseBodyResult) SetName(v string) *ListDiversifyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDiversifyResponseBodyResult) SetGmtModified(v string) *ListDiversifyResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListDiversifyResponseBodyResult) SetParameter(v *ListDiversifyResponseBodyResultParameter) *ListDiversifyResponseBodyResult {
	s.Parameter = v
	return s
}

type ListDiversifyResponseBodyResultParameter struct {
	Window        *int32 `json:"Window,omitempty" xml:"Window,omitempty"`
	CategoryIndex *int32 `json:"CategoryIndex,omitempty" xml:"CategoryIndex,omitempty"`
}

func (s ListDiversifyResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s ListDiversifyResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *ListDiversifyResponseBodyResultParameter) SetWindow(v int32) *ListDiversifyResponseBodyResultParameter {
	s.Window = &v
	return s
}

func (s *ListDiversifyResponseBodyResultParameter) SetCategoryIndex(v int32) *ListDiversifyResponseBodyResultParameter {
	s.CategoryIndex = &v
	return s
}

type ListDiversifyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDiversifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiversifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiversifyResponse) GoString() string {
	return s.String()
}

func (s *ListDiversifyResponse) SetHeaders(v map[string]*string) *ListDiversifyResponse {
	s.Headers = v
	return s
}

func (s *ListDiversifyResponse) SetBody(v *ListDiversifyResponseBody) *ListDiversifyResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	Page        *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size        *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetPage(v int32) *ListInstanceRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceRequest) SetSize(v int32) *ListInstanceRequest {
	s.Size = &v
	return s
}

func (s *ListInstanceRequest) SetStatus(v string) *ListInstanceRequest {
	s.Status = &v
	return s
}

func (s *ListInstanceRequest) SetName(v string) *ListInstanceRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceRequest) SetExpiredTime(v string) *ListInstanceRequest {
	s.ExpiredTime = &v
	return s
}

type ListInstanceResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetMessage(v string) *ListInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetCode(v string) *ListInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceResponseBody) SetResult(v []*ListInstanceResponseBodyResult) *ListInstanceResponseBody {
	s.Result = v
	return s
}

type ListInstanceResponseBodyResult struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LockMode       *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DataSetVersion *string `json:"DataSetVersion,omitempty" xml:"DataSetVersion,omitempty"`
	Industry       *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResult) SetStatus(v string) *ListInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetType(v string) *ListInstanceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetCommodityCode(v string) *ListInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetChargeType(v string) *ListInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetInstanceId(v string) *ListInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetGmtModified(v string) *ListInstanceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetLockMode(v string) *ListInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetRegionId(v string) *ListInstanceResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetDataSetVersion(v string) *ListInstanceResponseBodyResult {
	s.DataSetVersion = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetIndustry(v string) *ListInstanceResponseBodyResult {
	s.Industry = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetExpiredTime(v string) *ListInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetGmtCreate(v string) *ListInstanceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetName(v string) *ListInstanceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetScene(v string) *ListInstanceResponseBodyResult {
	s.Scene = &v
	return s
}

type ListInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListInstanceTaskResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListInstanceTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceTaskResponseBody) SetMessage(v string) *ListInstanceTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceTaskResponseBody) SetRequestId(v string) *ListInstanceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceTaskResponseBody) SetCode(v string) *ListInstanceTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceTaskResponseBody) SetResult(v []*ListInstanceTaskResponseBodyResult) *ListInstanceTaskResponseBody {
	s.Result = v
	return s
}

type ListInstanceTaskResponseBodyResult struct {
	SubProgressInfos []*ListInstanceTaskResponseBodyResultSubProgressInfos `json:"SubProgressInfos,omitempty" xml:"SubProgressInfos,omitempty" type:"Repeated"`
	Name             *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	TotalProgress    *int32                                                `json:"TotalProgress,omitempty" xml:"TotalProgress,omitempty"`
}

func (s ListInstanceTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceTaskResponseBodyResult) SetSubProgressInfos(v []*ListInstanceTaskResponseBodyResultSubProgressInfos) *ListInstanceTaskResponseBodyResult {
	s.SubProgressInfos = v
	return s
}

func (s *ListInstanceTaskResponseBodyResult) SetName(v string) *ListInstanceTaskResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInstanceTaskResponseBodyResult) SetTotalProgress(v int32) *ListInstanceTaskResponseBodyResult {
	s.TotalProgress = &v
	return s
}

type ListInstanceTaskResponseBodyResultSubProgressInfos struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Progress    *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	FinishedNum *int32  `json:"FinishedNum,omitempty" xml:"FinishedNum,omitempty"`
	TotalNum    *int32  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	Detail      *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s ListInstanceTaskResponseBodyResultSubProgressInfos) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTaskResponseBodyResultSubProgressInfos) GoString() string {
	return s.String()
}

func (s *ListInstanceTaskResponseBodyResultSubProgressInfos) SetType(v string) *ListInstanceTaskResponseBodyResultSubProgressInfos {
	s.Type = &v
	return s
}

func (s *ListInstanceTaskResponseBodyResultSubProgressInfos) SetProgress(v int32) *ListInstanceTaskResponseBodyResultSubProgressInfos {
	s.Progress = &v
	return s
}

func (s *ListInstanceTaskResponseBodyResultSubProgressInfos) SetFinishedNum(v int32) *ListInstanceTaskResponseBodyResultSubProgressInfos {
	s.FinishedNum = &v
	return s
}

func (s *ListInstanceTaskResponseBodyResultSubProgressInfos) SetTotalNum(v int32) *ListInstanceTaskResponseBodyResultSubProgressInfos {
	s.TotalNum = &v
	return s
}

func (s *ListInstanceTaskResponseBodyResultSubProgressInfos) SetDetail(v string) *ListInstanceTaskResponseBodyResultSubProgressInfos {
	s.Detail = &v
	return s
}

type ListInstanceTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTaskResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceTaskResponse) SetHeaders(v map[string]*string) *ListInstanceTaskResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceTaskResponse) SetBody(v *ListInstanceTaskResponseBody) *ListInstanceTaskResponse {
	s.Body = v
	return s
}

type ListItemsRequest struct {
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListItemsRequest) GoString() string {
	return s.String()
}

func (s *ListItemsRequest) SetPage(v int32) *ListItemsRequest {
	s.Page = &v
	return s
}

func (s *ListItemsRequest) SetSize(v int32) *ListItemsRequest {
	s.Size = &v
	return s
}

type ListItemsResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListItemsResponseBody) SetRequestId(v string) *ListItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListItemsResponseBody) SetResult(v *ListItemsResponseBodyResult) *ListItemsResponseBody {
	s.Result = v
	return s
}

type ListItemsResponseBodyResult struct {
	Total  *ListItemsResponseBodyResultTotal    `json:"Total,omitempty" xml:"Total,omitempty" type:"Struct"`
	Detail []*ListItemsResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s ListItemsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListItemsResponseBodyResult) SetTotal(v *ListItemsResponseBodyResultTotal) *ListItemsResponseBodyResult {
	s.Total = v
	return s
}

func (s *ListItemsResponseBodyResult) SetDetail(v []*ListItemsResponseBodyResultDetail) *ListItemsResponseBodyResult {
	s.Detail = v
	return s
}

type ListItemsResponseBodyResultTotal struct {
	WeightItem            *int64 `json:"WeightItem,omitempty" xml:"WeightItem,omitempty"`
	InstanceRecommendItem *int64 `json:"InstanceRecommendItem,omitempty" xml:"InstanceRecommendItem,omitempty"`
	QueryCount            *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	TotalCount            *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SceneRecommendItem    *int64 `json:"SceneRecommendItem,omitempty" xml:"SceneRecommendItem,omitempty"`
	SceneWeightItem       *int64 `json:"SceneWeightItem,omitempty" xml:"SceneWeightItem,omitempty"`
}

func (s ListItemsResponseBodyResultTotal) String() string {
	return tea.Prettify(s)
}

func (s ListItemsResponseBodyResultTotal) GoString() string {
	return s.String()
}

func (s *ListItemsResponseBodyResultTotal) SetWeightItem(v int64) *ListItemsResponseBodyResultTotal {
	s.WeightItem = &v
	return s
}

func (s *ListItemsResponseBodyResultTotal) SetInstanceRecommendItem(v int64) *ListItemsResponseBodyResultTotal {
	s.InstanceRecommendItem = &v
	return s
}

func (s *ListItemsResponseBodyResultTotal) SetQueryCount(v int64) *ListItemsResponseBodyResultTotal {
	s.QueryCount = &v
	return s
}

func (s *ListItemsResponseBodyResultTotal) SetTotalCount(v int64) *ListItemsResponseBodyResultTotal {
	s.TotalCount = &v
	return s
}

func (s *ListItemsResponseBodyResultTotal) SetSceneRecommendItem(v int64) *ListItemsResponseBodyResultTotal {
	s.SceneRecommendItem = &v
	return s
}

func (s *ListItemsResponseBodyResultTotal) SetSceneWeightItem(v int64) *ListItemsResponseBodyResultTotal {
	s.SceneWeightItem = &v
	return s
}

type ListItemsResponseBodyResultDetail struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PubTime      *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ItemType     *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Author       *string `json:"Author,omitempty" xml:"Author,omitempty"`
	CategoryPath *string `json:"CategoryPath,omitempty" xml:"CategoryPath,omitempty"`
	ItemId       *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ShopId       *string `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	BrandId      *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	Channel      *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListItemsResponseBodyResultDetail) String() string {
	return tea.Prettify(s)
}

func (s ListItemsResponseBodyResultDetail) GoString() string {
	return s.String()
}

func (s *ListItemsResponseBodyResultDetail) SetStatus(v string) *ListItemsResponseBodyResultDetail {
	s.Status = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetPubTime(v string) *ListItemsResponseBodyResultDetail {
	s.PubTime = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetExpireTime(v string) *ListItemsResponseBodyResultDetail {
	s.ExpireTime = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetItemType(v string) *ListItemsResponseBodyResultDetail {
	s.ItemType = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetAuthor(v string) *ListItemsResponseBodyResultDetail {
	s.Author = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetCategoryPath(v string) *ListItemsResponseBodyResultDetail {
	s.CategoryPath = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetItemId(v string) *ListItemsResponseBodyResultDetail {
	s.ItemId = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetShopId(v string) *ListItemsResponseBodyResultDetail {
	s.ShopId = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetBrandId(v string) *ListItemsResponseBodyResultDetail {
	s.BrandId = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetChannel(v string) *ListItemsResponseBodyResultDetail {
	s.Channel = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetDuration(v string) *ListItemsResponseBodyResultDetail {
	s.Duration = &v
	return s
}

func (s *ListItemsResponseBodyResultDetail) SetTitle(v string) *ListItemsResponseBodyResultDetail {
	s.Title = &v
	return s
}

type ListItemsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListItemsResponse) GoString() string {
	return s.String()
}

func (s *ListItemsResponse) SetHeaders(v map[string]*string) *ListItemsResponse {
	s.Headers = v
	return s
}

func (s *ListItemsResponse) SetBody(v *ListItemsResponseBody) *ListItemsResponse {
	s.Body = v
	return s
}

type ListLogsRequest struct {
	QueryParams *string `json:"QueryParams,omitempty" xml:"QueryParams,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page        *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size        *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogsRequest) GoString() string {
	return s.String()
}

func (s *ListLogsRequest) SetQueryParams(v string) *ListLogsRequest {
	s.QueryParams = &v
	return s
}

func (s *ListLogsRequest) SetStartTime(v int32) *ListLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLogsRequest) SetEndTime(v int32) *ListLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLogsRequest) SetPage(v int32) *ListLogsRequest {
	s.Page = &v
	return s
}

func (s *ListLogsRequest) SetSize(v int32) *ListLogsRequest {
	s.Size = &v
	return s
}

type ListLogsResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Headers   *ListLogsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	Result    []map[string]interface{}     `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogsResponseBody) SetRequestId(v string) *ListLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogsResponseBody) SetHeaders(v *ListLogsResponseBodyHeaders) *ListLogsResponseBody {
	s.Headers = v
	return s
}

func (s *ListLogsResponseBody) SetResult(v []map[string]interface{}) *ListLogsResponseBody {
	s.Result = v
	return s
}

type ListLogsResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListLogsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListLogsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListLogsResponseBodyHeaders) SetXTotalCount(v int32) *ListLogsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListLogsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogsResponse) GoString() string {
	return s.String()
}

func (s *ListLogsResponse) SetHeaders(v map[string]*string) *ListLogsResponse {
	s.Headers = v
	return s
}

func (s *ListLogsResponse) SetBody(v *ListLogsResponseBody) *ListLogsResponse {
	s.Body = v
	return s
}

type ListMixResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListMixResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListMixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMixResponseBody) GoString() string {
	return s.String()
}

func (s *ListMixResponseBody) SetMessage(v string) *ListMixResponseBody {
	s.Message = &v
	return s
}

func (s *ListMixResponseBody) SetRequestId(v string) *ListMixResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMixResponseBody) SetCode(v string) *ListMixResponseBody {
	s.Code = &v
	return s
}

func (s *ListMixResponseBody) SetResult(v []*ListMixResponseBodyResult) *ListMixResponseBody {
	s.Result = v
	return s
}

type ListMixResponseBodyResult struct {
	GmtCreate   *string                             `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                             `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *ListMixResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s ListMixResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListMixResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMixResponseBodyResult) SetGmtCreate(v string) *ListMixResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListMixResponseBodyResult) SetName(v string) *ListMixResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListMixResponseBodyResult) SetGmtModified(v string) *ListMixResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListMixResponseBodyResult) SetParameter(v *ListMixResponseBodyResultParameter) *ListMixResponseBodyResult {
	s.Parameter = v
	return s
}

type ListMixResponseBodyResultParameter struct {
	Settings []*ListMixResponseBodyResultParameterSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s ListMixResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s ListMixResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *ListMixResponseBodyResultParameter) SetSettings(v []*ListMixResponseBodyResultParameterSettings) *ListMixResponseBodyResultParameter {
	s.Settings = v
	return s
}

type ListMixResponseBodyResultParameterSettings struct {
	Value *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMixResponseBodyResultParameterSettings) String() string {
	return tea.Prettify(s)
}

func (s ListMixResponseBodyResultParameterSettings) GoString() string {
	return s.String()
}

func (s *ListMixResponseBodyResultParameterSettings) SetValue(v int32) *ListMixResponseBodyResultParameterSettings {
	s.Value = &v
	return s
}

func (s *ListMixResponseBodyResultParameterSettings) SetName(v string) *ListMixResponseBodyResultParameterSettings {
	s.Name = &v
	return s
}

type ListMixResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMixResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMixResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMixResponse) GoString() string {
	return s.String()
}

func (s *ListMixResponse) SetHeaders(v map[string]*string) *ListMixResponse {
	s.Headers = v
	return s
}

func (s *ListMixResponse) SetBody(v *ListMixResponseBody) *ListMixResponse {
	s.Body = v
	return s
}

type ListRuleConditionsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListRuleConditionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRuleConditionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRuleConditionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRuleConditionsResponseBody) SetRequestId(v string) *ListRuleConditionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRuleConditionsResponseBody) SetResult(v []*ListRuleConditionsResponseBodyResult) *ListRuleConditionsResponseBody {
	s.Result = v
	return s
}

type ListRuleConditionsResponseBodyResult struct {
	SelectionOperation *string `json:"SelectionOperation,omitempty" xml:"SelectionOperation,omitempty"`
	SelectValue        *string `json:"SelectValue,omitempty" xml:"SelectValue,omitempty"`
	SelectType         *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
}

func (s ListRuleConditionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRuleConditionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRuleConditionsResponseBodyResult) SetSelectionOperation(v string) *ListRuleConditionsResponseBodyResult {
	s.SelectionOperation = &v
	return s
}

func (s *ListRuleConditionsResponseBodyResult) SetSelectValue(v string) *ListRuleConditionsResponseBodyResult {
	s.SelectValue = &v
	return s
}

func (s *ListRuleConditionsResponseBodyResult) SetSelectType(v string) *ListRuleConditionsResponseBodyResult {
	s.SelectType = &v
	return s
}

type ListRuleConditionsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRuleConditionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRuleConditionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRuleConditionsResponse) GoString() string {
	return s.String()
}

func (s *ListRuleConditionsResponse) SetHeaders(v map[string]*string) *ListRuleConditionsResponse {
	s.Headers = v
	return s
}

func (s *ListRuleConditionsResponse) SetBody(v *ListRuleConditionsResponseBody) *ListRuleConditionsResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	RuleType  *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetSceneId(v string) *ListRulesRequest {
	s.SceneId = &v
	return s
}

func (s *ListRulesRequest) SetRuleType(v string) *ListRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListRulesRequest) SetStatus(v string) *ListRulesRequest {
	s.Status = &v
	return s
}

func (s *ListRulesRequest) SetPage(v int32) *ListRulesRequest {
	s.Page = &v
	return s
}

func (s *ListRulesRequest) SetSize(v int32) *ListRulesRequest {
	s.Size = &v
	return s
}

func (s *ListRulesRequest) SetStartTime(v int64) *ListRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListRulesRequest) SetEndTime(v int64) *ListRulesRequest {
	s.EndTime = &v
	return s
}

type ListRulesResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListRulesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetResult(v []*ListRulesResponseBodyResult) *ListRulesResponseBody {
	s.Result = v
	return s
}

type ListRulesResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListRulesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyResult) SetStatus(v string) *ListRulesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListRulesResponseBodyResult) SetGmtCreate(v string) *ListRulesResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListRulesResponseBodyResult) SetGmtModified(v string) *ListRulesResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListRulesResponseBodyResult) SetRuleId(v string) *ListRulesResponseBodyResult {
	s.RuleId = &v
	return s
}

type ListRulesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListRuleTasksRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListRuleTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRuleTasksRequest) GoString() string {
	return s.String()
}

func (s *ListRuleTasksRequest) SetSceneId(v string) *ListRuleTasksRequest {
	s.SceneId = &v
	return s
}

type ListRuleTasksResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListRuleTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRuleTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRuleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListRuleTasksResponseBody) SetRequestId(v string) *ListRuleTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRuleTasksResponseBody) SetResult(v *ListRuleTasksResponseBodyResult) *ListRuleTasksResponseBody {
	s.Result = v
	return s
}

type ListRuleTasksResponseBodyResult struct {
	FinishTime *int32 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	FinishRate *int32 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
}

func (s ListRuleTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRuleTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRuleTasksResponseBodyResult) SetFinishTime(v int32) *ListRuleTasksResponseBodyResult {
	s.FinishTime = &v
	return s
}

func (s *ListRuleTasksResponseBodyResult) SetFinishRate(v int32) *ListRuleTasksResponseBodyResult {
	s.FinishRate = &v
	return s
}

type ListRuleTasksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRuleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRuleTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRuleTasksResponse) GoString() string {
	return s.String()
}

func (s *ListRuleTasksResponse) SetHeaders(v map[string]*string) *ListRuleTasksResponse {
	s.Headers = v
	return s
}

func (s *ListRuleTasksResponse) SetBody(v *ListRuleTasksResponseBody) *ListRuleTasksResponse {
	s.Body = v
	return s
}

type ListSceneItemsRequest struct {
	OperationRuleId *string `json:"OperationRuleId,omitempty" xml:"OperationRuleId,omitempty"`
	SelectionRuleId *string `json:"SelectionRuleId,omitempty" xml:"SelectionRuleId,omitempty"`
	Page            *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size            *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	PreviewType     *string `json:"PreviewType,omitempty" xml:"PreviewType,omitempty"`
	QueryCount      *int32  `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
}

func (s ListSceneItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSceneItemsRequest) GoString() string {
	return s.String()
}

func (s *ListSceneItemsRequest) SetOperationRuleId(v string) *ListSceneItemsRequest {
	s.OperationRuleId = &v
	return s
}

func (s *ListSceneItemsRequest) SetSelectionRuleId(v string) *ListSceneItemsRequest {
	s.SelectionRuleId = &v
	return s
}

func (s *ListSceneItemsRequest) SetPage(v int32) *ListSceneItemsRequest {
	s.Page = &v
	return s
}

func (s *ListSceneItemsRequest) SetSize(v int32) *ListSceneItemsRequest {
	s.Size = &v
	return s
}

func (s *ListSceneItemsRequest) SetPreviewType(v string) *ListSceneItemsRequest {
	s.PreviewType = &v
	return s
}

func (s *ListSceneItemsRequest) SetQueryCount(v int32) *ListSceneItemsRequest {
	s.QueryCount = &v
	return s
}

type ListSceneItemsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSceneItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSceneItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSceneItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneItemsResponseBody) SetRequestId(v string) *ListSceneItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSceneItemsResponseBody) SetResult(v *ListSceneItemsResponseBodyResult) *ListSceneItemsResponseBody {
	s.Result = v
	return s
}

type ListSceneItemsResponseBodyResult struct {
	Total  *ListSceneItemsResponseBodyResultTotal    `json:"Total,omitempty" xml:"Total,omitempty" type:"Struct"`
	Detail []*ListSceneItemsResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s ListSceneItemsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSceneItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSceneItemsResponseBodyResult) SetTotal(v *ListSceneItemsResponseBodyResultTotal) *ListSceneItemsResponseBodyResult {
	s.Total = v
	return s
}

func (s *ListSceneItemsResponseBodyResult) SetDetail(v []*ListSceneItemsResponseBodyResultDetail) *ListSceneItemsResponseBodyResult {
	s.Detail = v
	return s
}

type ListSceneItemsResponseBodyResultTotal struct {
	WeightItem            *int64 `json:"WeightItem,omitempty" xml:"WeightItem,omitempty"`
	InstanceRecommendItem *int64 `json:"InstanceRecommendItem,omitempty" xml:"InstanceRecommendItem,omitempty"`
	TotalCount            *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SceneRecommendItem    *int64 `json:"SceneRecommendItem,omitempty" xml:"SceneRecommendItem,omitempty"`
	SceneWeightItem       *int64 `json:"SceneWeightItem,omitempty" xml:"SceneWeightItem,omitempty"`
}

func (s ListSceneItemsResponseBodyResultTotal) String() string {
	return tea.Prettify(s)
}

func (s ListSceneItemsResponseBodyResultTotal) GoString() string {
	return s.String()
}

func (s *ListSceneItemsResponseBodyResultTotal) SetWeightItem(v int64) *ListSceneItemsResponseBodyResultTotal {
	s.WeightItem = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultTotal) SetInstanceRecommendItem(v int64) *ListSceneItemsResponseBodyResultTotal {
	s.InstanceRecommendItem = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultTotal) SetTotalCount(v int64) *ListSceneItemsResponseBodyResultTotal {
	s.TotalCount = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultTotal) SetSceneRecommendItem(v int64) *ListSceneItemsResponseBodyResultTotal {
	s.SceneRecommendItem = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultTotal) SetSceneWeightItem(v int64) *ListSceneItemsResponseBodyResultTotal {
	s.SceneWeightItem = &v
	return s
}

type ListSceneItemsResponseBodyResultDetail struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PubTime      *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ItemType     *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Author       *string `json:"Author,omitempty" xml:"Author,omitempty"`
	CategoryPath *string `json:"CategoryPath,omitempty" xml:"CategoryPath,omitempty"`
	ItemId       *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ShopId       *string `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	BrandId      *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	Channel      *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListSceneItemsResponseBodyResultDetail) String() string {
	return tea.Prettify(s)
}

func (s ListSceneItemsResponseBodyResultDetail) GoString() string {
	return s.String()
}

func (s *ListSceneItemsResponseBodyResultDetail) SetStatus(v string) *ListSceneItemsResponseBodyResultDetail {
	s.Status = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetPubTime(v string) *ListSceneItemsResponseBodyResultDetail {
	s.PubTime = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetExpireTime(v string) *ListSceneItemsResponseBodyResultDetail {
	s.ExpireTime = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetItemType(v string) *ListSceneItemsResponseBodyResultDetail {
	s.ItemType = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetAuthor(v string) *ListSceneItemsResponseBodyResultDetail {
	s.Author = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetCategoryPath(v string) *ListSceneItemsResponseBodyResultDetail {
	s.CategoryPath = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetItemId(v string) *ListSceneItemsResponseBodyResultDetail {
	s.ItemId = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetShopId(v string) *ListSceneItemsResponseBodyResultDetail {
	s.ShopId = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetBrandId(v string) *ListSceneItemsResponseBodyResultDetail {
	s.BrandId = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetChannel(v string) *ListSceneItemsResponseBodyResultDetail {
	s.Channel = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetDuration(v string) *ListSceneItemsResponseBodyResultDetail {
	s.Duration = &v
	return s
}

func (s *ListSceneItemsResponseBodyResultDetail) SetTitle(v string) *ListSceneItemsResponseBodyResultDetail {
	s.Title = &v
	return s
}

type ListSceneItemsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSceneItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSceneItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSceneItemsResponse) GoString() string {
	return s.String()
}

func (s *ListSceneItemsResponse) SetHeaders(v map[string]*string) *ListSceneItemsResponse {
	s.Headers = v
	return s
}

func (s *ListSceneItemsResponse) SetBody(v *ListSceneItemsResponseBody) *ListSceneItemsResponse {
	s.Body = v
	return s
}

type ListScenesRequest struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenesRequest) GoString() string {
	return s.String()
}

func (s *ListScenesRequest) SetStatus(v string) *ListScenesRequest {
	s.Status = &v
	return s
}

type ListScenesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListScenesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s *ListScenesResponseBody) SetResult(v []*ListScenesResponseBodyResult) *ListScenesResponseBody {
	s.Result = v
	return s
}

type ListScenesResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s ListScenesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListScenesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListScenesResponseBodyResult) SetStatus(v string) *ListScenesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListScenesResponseBodyResult) SetGmtCreate(v string) *ListScenesResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListScenesResponseBodyResult) SetSceneId(v string) *ListScenesResponseBodyResult {
	s.SceneId = &v
	return s
}

func (s *ListScenesResponseBodyResult) SetGmtModified(v string) *ListScenesResponseBodyResult {
	s.GmtModified = &v
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

type ListUmengAppkeysResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*ListUmengAppkeysResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListUmengAppkeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUmengAppkeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListUmengAppkeysResponseBody) SetMessage(v string) *ListUmengAppkeysResponseBody {
	s.Message = &v
	return s
}

func (s *ListUmengAppkeysResponseBody) SetRequestId(v string) *ListUmengAppkeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUmengAppkeysResponseBody) SetCode(v string) *ListUmengAppkeysResponseBody {
	s.Code = &v
	return s
}

func (s *ListUmengAppkeysResponseBody) SetResult(v []*ListUmengAppkeysResponseBodyResult) *ListUmengAppkeysResponseBody {
	s.Result = v
	return s
}

type ListUmengAppkeysResponseBodyResult struct {
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Appkey   *string `json:"Appkey,omitempty" xml:"Appkey,omitempty"`
}

func (s ListUmengAppkeysResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListUmengAppkeysResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUmengAppkeysResponseBodyResult) SetPlatform(v string) *ListUmengAppkeysResponseBodyResult {
	s.Platform = &v
	return s
}

func (s *ListUmengAppkeysResponseBodyResult) SetName(v string) *ListUmengAppkeysResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListUmengAppkeysResponseBodyResult) SetAppkey(v string) *ListUmengAppkeysResponseBodyResult {
	s.Appkey = &v
	return s
}

type ListUmengAppkeysResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUmengAppkeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUmengAppkeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUmengAppkeysResponse) GoString() string {
	return s.String()
}

func (s *ListUmengAppkeysResponse) SetHeaders(v map[string]*string) *ListUmengAppkeysResponse {
	s.Headers = v
	return s
}

func (s *ListUmengAppkeysResponse) SetBody(v *ListUmengAppkeysResponseBody) *ListUmengAppkeysResponse {
	s.Body = v
	return s
}

type ModifyDataSourceResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ModifyDataSourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBody) SetMessage(v string) *ModifyDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDataSourceResponseBody) SetRequestId(v string) *ModifyDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceResponseBody) SetCode(v string) *ModifyDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDataSourceResponseBody) SetResult(v *ModifyDataSourceResponseBodyResult) *ModifyDataSourceResponseBody {
	s.Result = v
	return s
}

type ModifyDataSourceResponseBodyResult struct {
	TableName   *string                                 `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Meta        *ModifyDataSourceResponseBodyResultMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	GmtCreate   *string                                 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string                                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s ModifyDataSourceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBodyResult) SetTableName(v string) *ModifyDataSourceResponseBodyResult {
	s.TableName = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResult) SetMeta(v *ModifyDataSourceResponseBodyResultMeta) *ModifyDataSourceResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ModifyDataSourceResponseBodyResult) SetGmtCreate(v string) *ModifyDataSourceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResult) SetGmtModified(v string) *ModifyDataSourceResponseBodyResult {
	s.GmtModified = &v
	return s
}

type ModifyDataSourceResponseBodyResultMeta struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	BucketName  *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Partition   *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
}

func (s ModifyDataSourceResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetType(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.Type = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetTableName(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.TableName = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetProjectName(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.ProjectName = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetBucketName(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.BucketName = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetPath(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetPartition(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.Partition = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetTimestamp(v int64) *ModifyDataSourceResponseBodyResultMeta {
	s.Timestamp = &v
	return s
}

func (s *ModifyDataSourceResponseBodyResultMeta) SetAccessKeyId(v string) *ModifyDataSourceResponseBodyResultMeta {
	s.AccessKeyId = &v
	return s
}

type ModifyDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponse) SetHeaders(v map[string]*string) *ModifyDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceResponse) SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse {
	s.Body = v
	return s
}

type ModifyDiversifyResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ModifyDiversifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyDiversifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiversifyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiversifyResponseBody) SetMessage(v string) *ModifyDiversifyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDiversifyResponseBody) SetRequestId(v string) *ModifyDiversifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiversifyResponseBody) SetCode(v string) *ModifyDiversifyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDiversifyResponseBody) SetResult(v *ModifyDiversifyResponseBodyResult) *ModifyDiversifyResponseBody {
	s.Result = v
	return s
}

type ModifyDiversifyResponseBodyResult struct {
	GmtCreate   *string                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *ModifyDiversifyResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s ModifyDiversifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiversifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyDiversifyResponseBodyResult) SetGmtCreate(v string) *ModifyDiversifyResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ModifyDiversifyResponseBodyResult) SetName(v string) *ModifyDiversifyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyDiversifyResponseBodyResult) SetGmtModified(v string) *ModifyDiversifyResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ModifyDiversifyResponseBodyResult) SetParameter(v *ModifyDiversifyResponseBodyResultParameter) *ModifyDiversifyResponseBodyResult {
	s.Parameter = v
	return s
}

type ModifyDiversifyResponseBodyResultParameter struct {
	Window        *int32 `json:"Window,omitempty" xml:"Window,omitempty"`
	CategoryIndex *int32 `json:"CategoryIndex,omitempty" xml:"CategoryIndex,omitempty"`
}

func (s ModifyDiversifyResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiversifyResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *ModifyDiversifyResponseBodyResultParameter) SetWindow(v int32) *ModifyDiversifyResponseBodyResultParameter {
	s.Window = &v
	return s
}

func (s *ModifyDiversifyResponseBodyResultParameter) SetCategoryIndex(v int32) *ModifyDiversifyResponseBodyResultParameter {
	s.CategoryIndex = &v
	return s
}

type ModifyDiversifyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDiversifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDiversifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiversifyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiversifyResponse) SetHeaders(v map[string]*string) *ModifyDiversifyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiversifyResponse) SetBody(v *ModifyDiversifyResponseBody) *ModifyDiversifyResponse {
	s.Body = v
	return s
}

type ModifyExposureSettingsResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ModifyExposureSettingsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyExposureSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExposureSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExposureSettingsResponseBody) SetMessage(v string) *ModifyExposureSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExposureSettingsResponseBody) SetRequestId(v string) *ModifyExposureSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExposureSettingsResponseBody) SetCode(v string) *ModifyExposureSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExposureSettingsResponseBody) SetResult(v *ModifyExposureSettingsResponseBodyResult) *ModifyExposureSettingsResponseBody {
	s.Result = v
	return s
}

type ModifyExposureSettingsResponseBodyResult struct {
	DurationSeconds *int32 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	ScenarioBased   *bool  `json:"ScenarioBased,omitempty" xml:"ScenarioBased,omitempty"`
}

func (s ModifyExposureSettingsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyExposureSettingsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyExposureSettingsResponseBodyResult) SetDurationSeconds(v int32) *ModifyExposureSettingsResponseBodyResult {
	s.DurationSeconds = &v
	return s
}

func (s *ModifyExposureSettingsResponseBodyResult) SetScenarioBased(v bool) *ModifyExposureSettingsResponseBodyResult {
	s.ScenarioBased = &v
	return s
}

type ModifyExposureSettingsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyExposureSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyExposureSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExposureSettingsResponse) GoString() string {
	return s.String()
}

func (s *ModifyExposureSettingsResponse) SetHeaders(v map[string]*string) *ModifyExposureSettingsResponse {
	s.Headers = v
	return s
}

func (s *ModifyExposureSettingsResponse) SetBody(v *ModifyExposureSettingsResponseBody) *ModifyExposureSettingsResponse {
	s.Body = v
	return s
}

type ModifyInstanceResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ModifyInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) SetMessage(v string) *ModifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetCode(v string) *ModifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetResult(v *ModifyInstanceResponseBodyResult) *ModifyInstanceResponseBody {
	s.Result = v
	return s
}

type ModifyInstanceResponseBodyResult struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LockMode       *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DataSetVersion *string `json:"DataSetVersion,omitempty" xml:"DataSetVersion,omitempty"`
	Industry       *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ModifyInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBodyResult) SetStatus(v string) *ModifyInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetType(v string) *ModifyInstanceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetCommodityCode(v string) *ModifyInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetChargeType(v string) *ModifyInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetInstanceId(v string) *ModifyInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetGmtModified(v string) *ModifyInstanceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetLockMode(v string) *ModifyInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetRegionId(v string) *ModifyInstanceResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetDataSetVersion(v string) *ModifyInstanceResponseBodyResult {
	s.DataSetVersion = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetIndustry(v string) *ModifyInstanceResponseBodyResult {
	s.Industry = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetExpiredTime(v string) *ModifyInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetGmtCreate(v string) *ModifyInstanceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetName(v string) *ModifyInstanceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyInstanceResponseBodyResult) SetScene(v string) *ModifyInstanceResponseBodyResult {
	s.Scene = &v
	return s
}

type ModifyInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponse) SetHeaders(v map[string]*string) *ModifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceResponse) SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse {
	s.Body = v
	return s
}

type ModifyItemsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyItemsResponseBody) SetRequestId(v string) *ModifyItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyItemsResponseBody) SetResult(v bool) *ModifyItemsResponseBody {
	s.Result = &v
	return s
}

type ModifyItemsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyItemsResponse) GoString() string {
	return s.String()
}

func (s *ModifyItemsResponse) SetHeaders(v map[string]*string) *ModifyItemsResponse {
	s.Headers = v
	return s
}

func (s *ModifyItemsResponse) SetBody(v *ModifyItemsResponseBody) *ModifyItemsResponse {
	s.Body = v
	return s
}

type ModifyMixResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *ModifyMixResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyMixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMixResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMixResponseBody) SetMessage(v string) *ModifyMixResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMixResponseBody) SetRequestId(v string) *ModifyMixResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMixResponseBody) SetCode(v string) *ModifyMixResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMixResponseBody) SetResult(v *ModifyMixResponseBodyResult) *ModifyMixResponseBody {
	s.Result = v
	return s
}

type ModifyMixResponseBodyResult struct {
	GmtCreate   *string                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Parameter   *ModifyMixResponseBodyResultParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
}

func (s ModifyMixResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyMixResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyMixResponseBodyResult) SetGmtCreate(v string) *ModifyMixResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ModifyMixResponseBodyResult) SetName(v string) *ModifyMixResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyMixResponseBodyResult) SetGmtModified(v string) *ModifyMixResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ModifyMixResponseBodyResult) SetParameter(v *ModifyMixResponseBodyResultParameter) *ModifyMixResponseBodyResult {
	s.Parameter = v
	return s
}

type ModifyMixResponseBodyResultParameter struct {
	Settings []*ModifyMixResponseBodyResultParameterSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s ModifyMixResponseBodyResultParameter) String() string {
	return tea.Prettify(s)
}

func (s ModifyMixResponseBodyResultParameter) GoString() string {
	return s.String()
}

func (s *ModifyMixResponseBodyResultParameter) SetSettings(v []*ModifyMixResponseBodyResultParameterSettings) *ModifyMixResponseBodyResultParameter {
	s.Settings = v
	return s
}

type ModifyMixResponseBodyResultParameterSettings struct {
	Value *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyMixResponseBodyResultParameterSettings) String() string {
	return tea.Prettify(s)
}

func (s ModifyMixResponseBodyResultParameterSettings) GoString() string {
	return s.String()
}

func (s *ModifyMixResponseBodyResultParameterSettings) SetValue(v int32) *ModifyMixResponseBodyResultParameterSettings {
	s.Value = &v
	return s
}

func (s *ModifyMixResponseBodyResultParameterSettings) SetName(v string) *ModifyMixResponseBodyResultParameterSettings {
	s.Name = &v
	return s
}

type ModifyMixResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMixResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMixResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMixResponse) GoString() string {
	return s.String()
}

func (s *ModifyMixResponse) SetHeaders(v map[string]*string) *ModifyMixResponse {
	s.Headers = v
	return s
}

func (s *ModifyMixResponse) SetBody(v *ModifyMixResponseBody) *ModifyMixResponse {
	s.Body = v
	return s
}

type ModifyRuleResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifyRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponseBody) SetRequestId(v string) *ModifyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRuleResponseBody) SetResult(v *ModifyRuleResponseBodyResult) *ModifyRuleResponseBody {
	s.Result = v
	return s
}

type ModifyRuleResponseBodyResult struct {
	Status      *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	RuleMeta    map[string]interface{} `json:"RuleMeta,omitempty" xml:"RuleMeta,omitempty"`
	GmtModified *string                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleId      *string                `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponseBodyResult) SetStatus(v string) *ModifyRuleResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyRuleResponseBodyResult) SetGmtCreate(v string) *ModifyRuleResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ModifyRuleResponseBodyResult) SetRuleMeta(v map[string]interface{}) *ModifyRuleResponseBodyResult {
	s.RuleMeta = v
	return s
}

func (s *ModifyRuleResponseBodyResult) SetGmtModified(v string) *ModifyRuleResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ModifyRuleResponseBodyResult) SetRuleId(v string) *ModifyRuleResponseBodyResult {
	s.RuleId = &v
	return s
}

type ModifyRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponse) SetHeaders(v map[string]*string) *ModifyRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyRuleResponse) SetBody(v *ModifyRuleResponseBody) *ModifyRuleResponse {
	s.Body = v
	return s
}

type ModifySceneResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifySceneResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifySceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySceneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySceneResponseBody) SetRequestId(v string) *ModifySceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySceneResponseBody) SetResult(v *ModifySceneResponseBodyResult) *ModifySceneResponseBody {
	s.Result = v
	return s
}

type ModifySceneResponseBodyResult struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s ModifySceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifySceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifySceneResponseBodyResult) SetStatus(v string) *ModifySceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifySceneResponseBodyResult) SetGmtCreate(v string) *ModifySceneResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ModifySceneResponseBodyResult) SetSceneId(v string) *ModifySceneResponseBodyResult {
	s.SceneId = &v
	return s
}

func (s *ModifySceneResponseBodyResult) SetGmtModified(v string) *ModifySceneResponseBodyResult {
	s.GmtModified = &v
	return s
}

type ModifySceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySceneResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySceneResponse) GoString() string {
	return s.String()
}

func (s *ModifySceneResponse) SetHeaders(v map[string]*string) *ModifySceneResponse {
	s.Headers = v
	return s
}

func (s *ModifySceneResponse) SetBody(v *ModifySceneResponseBody) *ModifySceneResponse {
	s.Body = v
	return s
}

type PublishRuleRequest struct {
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s PublishRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleRequest) GoString() string {
	return s.String()
}

func (s *PublishRuleRequest) SetRuleType(v string) *PublishRuleRequest {
	s.RuleType = &v
	return s
}

func (s *PublishRuleRequest) SetSceneId(v string) *PublishRuleRequest {
	s.SceneId = &v
	return s
}

type PublishRuleResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *PublishRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s PublishRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRuleResponseBody) SetRequestId(v string) *PublishRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishRuleResponseBody) SetResult(v *PublishRuleResponseBodyResult) *PublishRuleResponseBody {
	s.Result = v
	return s
}

type PublishRuleResponseBodyResult struct {
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s PublishRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PublishRuleResponseBodyResult) SetRuleId(v string) *PublishRuleResponseBodyResult {
	s.RuleId = &v
	return s
}

type PublishRuleResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleResponse) GoString() string {
	return s.String()
}

func (s *PublishRuleResponse) SetHeaders(v map[string]*string) *PublishRuleResponse {
	s.Headers = v
	return s
}

func (s *PublishRuleResponse) SetBody(v *PublishRuleResponseBody) *PublishRuleResponse {
	s.Body = v
	return s
}

type PushDocumentResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PushDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *PushDocumentResponseBody) SetMessage(v string) *PushDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *PushDocumentResponseBody) SetRequestId(v string) *PushDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushDocumentResponseBody) SetCode(v string) *PushDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *PushDocumentResponseBody) SetResult(v bool) *PushDocumentResponseBody {
	s.Result = &v
	return s
}

type PushDocumentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s PushDocumentResponse) GoString() string {
	return s.String()
}

func (s *PushDocumentResponse) SetHeaders(v map[string]*string) *PushDocumentResponse {
	s.Headers = v
	return s
}

func (s *PushDocumentResponse) SetBody(v *PushDocumentResponseBody) *PushDocumentResponse {
	s.Body = v
	return s
}

type PushInterventionResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PushInterventionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionResponseBody) GoString() string {
	return s.String()
}

func (s *PushInterventionResponseBody) SetMessage(v string) *PushInterventionResponseBody {
	s.Message = &v
	return s
}

func (s *PushInterventionResponseBody) SetRequestId(v string) *PushInterventionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushInterventionResponseBody) SetCode(v string) *PushInterventionResponseBody {
	s.Code = &v
	return s
}

func (s *PushInterventionResponseBody) SetResult(v bool) *PushInterventionResponseBody {
	s.Result = &v
	return s
}

type PushInterventionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushInterventionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushInterventionResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionResponse) GoString() string {
	return s.String()
}

func (s *PushInterventionResponse) SetHeaders(v map[string]*string) *PushInterventionResponse {
	s.Headers = v
	return s
}

func (s *PushInterventionResponse) SetBody(v *PushInterventionResponseBody) *PushInterventionResponse {
	s.Body = v
	return s
}

type QueryDataMessageRequest struct {
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CmdType       *string `json:"CmdType,omitempty" xml:"CmdType,omitempty"`
	ItemId        *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemType      *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	Page          *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size          *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	TraceId       *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	SceneId       *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	BhvType       *string `json:"BhvType,omitempty" xml:"BhvType,omitempty"`
	MessageSource *string `json:"MessageSource,omitempty" xml:"MessageSource,omitempty"`
}

func (s QueryDataMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDataMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryDataMessageRequest) SetStartTime(v int64) *QueryDataMessageRequest {
	s.StartTime = &v
	return s
}

func (s *QueryDataMessageRequest) SetEndTime(v int64) *QueryDataMessageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDataMessageRequest) SetCmdType(v string) *QueryDataMessageRequest {
	s.CmdType = &v
	return s
}

func (s *QueryDataMessageRequest) SetItemId(v string) *QueryDataMessageRequest {
	s.ItemId = &v
	return s
}

func (s *QueryDataMessageRequest) SetItemType(v string) *QueryDataMessageRequest {
	s.ItemType = &v
	return s
}

func (s *QueryDataMessageRequest) SetUserId(v string) *QueryDataMessageRequest {
	s.UserId = &v
	return s
}

func (s *QueryDataMessageRequest) SetUserType(v string) *QueryDataMessageRequest {
	s.UserType = &v
	return s
}

func (s *QueryDataMessageRequest) SetPage(v int32) *QueryDataMessageRequest {
	s.Page = &v
	return s
}

func (s *QueryDataMessageRequest) SetSize(v int32) *QueryDataMessageRequest {
	s.Size = &v
	return s
}

func (s *QueryDataMessageRequest) SetTraceId(v string) *QueryDataMessageRequest {
	s.TraceId = &v
	return s
}

func (s *QueryDataMessageRequest) SetSceneId(v string) *QueryDataMessageRequest {
	s.SceneId = &v
	return s
}

func (s *QueryDataMessageRequest) SetBhvType(v string) *QueryDataMessageRequest {
	s.BhvType = &v
	return s
}

func (s *QueryDataMessageRequest) SetMessageSource(v string) *QueryDataMessageRequest {
	s.MessageSource = &v
	return s
}

type QueryDataMessageResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QueryDataMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDataMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataMessageResponseBody) SetMessage(v string) *QueryDataMessageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDataMessageResponseBody) SetRequestId(v string) *QueryDataMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataMessageResponseBody) SetCode(v string) *QueryDataMessageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDataMessageResponseBody) SetResult(v map[string]interface{}) *QueryDataMessageResponseBody {
	s.Result = v
	return s
}

type QueryDataMessageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDataMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDataMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDataMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryDataMessageResponse) SetHeaders(v map[string]*string) *QueryDataMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryDataMessageResponse) SetBody(v *QueryDataMessageResponseBody) *QueryDataMessageResponse {
	s.Body = v
	return s
}

type QueryDataMessageStatisticsRequest struct {
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CmdType       *string `json:"CmdType,omitempty" xml:"CmdType,omitempty"`
	ItemId        *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemType      *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	TraceId       *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	SceneId       *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	BhvType       *string `json:"BhvType,omitempty" xml:"BhvType,omitempty"`
	MessageSource *string `json:"MessageSource,omitempty" xml:"MessageSource,omitempty"`
}

func (s QueryDataMessageStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDataMessageStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QueryDataMessageStatisticsRequest) SetStartTime(v int64) *QueryDataMessageStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetEndTime(v int64) *QueryDataMessageStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetCmdType(v string) *QueryDataMessageStatisticsRequest {
	s.CmdType = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetItemId(v string) *QueryDataMessageStatisticsRequest {
	s.ItemId = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetItemType(v string) *QueryDataMessageStatisticsRequest {
	s.ItemType = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetUserId(v string) *QueryDataMessageStatisticsRequest {
	s.UserId = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetUserType(v string) *QueryDataMessageStatisticsRequest {
	s.UserType = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetTraceId(v string) *QueryDataMessageStatisticsRequest {
	s.TraceId = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetSceneId(v string) *QueryDataMessageStatisticsRequest {
	s.SceneId = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetBhvType(v string) *QueryDataMessageStatisticsRequest {
	s.BhvType = &v
	return s
}

func (s *QueryDataMessageStatisticsRequest) SetMessageSource(v string) *QueryDataMessageStatisticsRequest {
	s.MessageSource = &v
	return s
}

type QueryDataMessageStatisticsResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QueryDataMessageStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDataMessageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataMessageStatisticsResponseBody) SetMessage(v string) *QueryDataMessageStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDataMessageStatisticsResponseBody) SetRequestId(v string) *QueryDataMessageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataMessageStatisticsResponseBody) SetCode(v string) *QueryDataMessageStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDataMessageStatisticsResponseBody) SetResult(v map[string]interface{}) *QueryDataMessageStatisticsResponseBody {
	s.Result = v
	return s
}

type QueryDataMessageStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDataMessageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDataMessageStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDataMessageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QueryDataMessageStatisticsResponse) SetHeaders(v map[string]*string) *QueryDataMessageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QueryDataMessageStatisticsResponse) SetBody(v *QueryDataMessageStatisticsResponseBody) *QueryDataMessageStatisticsResponse {
	s.Body = v
	return s
}

type QueryExceptionHistoryRequest struct {
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryExceptionHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryExceptionHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryExceptionHistoryRequest) SetStartTime(v int64) *QueryExceptionHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *QueryExceptionHistoryRequest) SetEndTime(v int64) *QueryExceptionHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *QueryExceptionHistoryRequest) SetType(v string) *QueryExceptionHistoryRequest {
	s.Type = &v
	return s
}

type QueryExceptionHistoryResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QueryExceptionHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExceptionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExceptionHistoryResponseBody) SetMessage(v string) *QueryExceptionHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *QueryExceptionHistoryResponseBody) SetRequestId(v string) *QueryExceptionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExceptionHistoryResponseBody) SetCode(v string) *QueryExceptionHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *QueryExceptionHistoryResponseBody) SetResult(v map[string]interface{}) *QueryExceptionHistoryResponseBody {
	s.Result = v
	return s
}

type QueryExceptionHistoryResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryExceptionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryExceptionHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExceptionHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryExceptionHistoryResponse) SetHeaders(v map[string]*string) *QueryExceptionHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryExceptionHistoryResponse) SetBody(v *QueryExceptionHistoryResponseBody) *QueryExceptionHistoryResponse {
	s.Body = v
	return s
}

type QueryRawDataRequest struct {
	ItemId   *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s QueryRawDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRawDataRequest) GoString() string {
	return s.String()
}

func (s *QueryRawDataRequest) SetItemId(v string) *QueryRawDataRequest {
	s.ItemId = &v
	return s
}

func (s *QueryRawDataRequest) SetItemType(v string) *QueryRawDataRequest {
	s.ItemType = &v
	return s
}

func (s *QueryRawDataRequest) SetUserId(v string) *QueryRawDataRequest {
	s.UserId = &v
	return s
}

func (s *QueryRawDataRequest) SetUserType(v string) *QueryRawDataRequest {
	s.UserType = &v
	return s
}

type QueryRawDataResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QueryRawDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRawDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRawDataResponseBody) SetMessage(v string) *QueryRawDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRawDataResponseBody) SetRequestId(v string) *QueryRawDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRawDataResponseBody) SetCode(v string) *QueryRawDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRawDataResponseBody) SetResult(v map[string]interface{}) *QueryRawDataResponseBody {
	s.Result = v
	return s
}

type QueryRawDataResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRawDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRawDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRawDataResponse) GoString() string {
	return s.String()
}

func (s *QueryRawDataResponse) SetHeaders(v map[string]*string) *QueryRawDataResponse {
	s.Headers = v
	return s
}

func (s *QueryRawDataResponse) SetBody(v *QueryRawDataResponseBody) *QueryRawDataResponse {
	s.Body = v
	return s
}

type QuerySingleAggregationReportResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QuerySingleAggregationReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleAggregationReportResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleAggregationReportResponseBody) SetMessage(v string) *QuerySingleAggregationReportResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySingleAggregationReportResponseBody) SetRequestId(v string) *QuerySingleAggregationReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySingleAggregationReportResponseBody) SetCode(v string) *QuerySingleAggregationReportResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySingleAggregationReportResponseBody) SetResult(v map[string]interface{}) *QuerySingleAggregationReportResponseBody {
	s.Result = v
	return s
}

type QuerySingleAggregationReportResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySingleAggregationReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySingleAggregationReportResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleAggregationReportResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleAggregationReportResponse) SetHeaders(v map[string]*string) *QuerySingleAggregationReportResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleAggregationReportResponse) SetBody(v *QuerySingleAggregationReportResponseBody) *QuerySingleAggregationReportResponse {
	s.Body = v
	return s
}

type QuerySingleReportRequest struct {
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s QuerySingleReportRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleReportRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleReportRequest) SetReportType(v string) *QuerySingleReportRequest {
	s.ReportType = &v
	return s
}

type QuerySingleReportResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QuerySingleReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleReportResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleReportResponseBody) SetMessage(v string) *QuerySingleReportResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySingleReportResponseBody) SetRequestId(v string) *QuerySingleReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySingleReportResponseBody) SetCode(v string) *QuerySingleReportResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySingleReportResponseBody) SetResult(v map[string]interface{}) *QuerySingleReportResponseBody {
	s.Result = v
	return s
}

type QuerySingleReportResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySingleReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySingleReportResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleReportResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleReportResponse) SetHeaders(v map[string]*string) *QuerySingleReportResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleReportResponse) SetBody(v *QuerySingleReportResponseBody) *QuerySingleReportResponse {
	s.Body = v
	return s
}

type QuerySyncReportAggregationRequest struct {
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s QuerySyncReportAggregationRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncReportAggregationRequest) GoString() string {
	return s.String()
}

func (s *QuerySyncReportAggregationRequest) SetStartTime(v int64) *QuerySyncReportAggregationRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySyncReportAggregationRequest) SetEndTime(v int64) *QuerySyncReportAggregationRequest {
	s.EndTime = &v
	return s
}

type QuerySyncReportAggregationResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QuerySyncReportAggregationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncReportAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySyncReportAggregationResponseBody) SetMessage(v string) *QuerySyncReportAggregationResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySyncReportAggregationResponseBody) SetRequestId(v string) *QuerySyncReportAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySyncReportAggregationResponseBody) SetCode(v string) *QuerySyncReportAggregationResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySyncReportAggregationResponseBody) SetResult(v map[string]interface{}) *QuerySyncReportAggregationResponseBody {
	s.Result = v
	return s
}

type QuerySyncReportAggregationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySyncReportAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySyncReportAggregationResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncReportAggregationResponse) GoString() string {
	return s.String()
}

func (s *QuerySyncReportAggregationResponse) SetHeaders(v map[string]*string) *QuerySyncReportAggregationResponse {
	s.Headers = v
	return s
}

func (s *QuerySyncReportAggregationResponse) SetBody(v *QuerySyncReportAggregationResponseBody) *QuerySyncReportAggregationResponse {
	s.Body = v
	return s
}

type RecommendHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	RegionId      *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RecommendHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecommendHeaders) GoString() string {
	return s.String()
}

func (s *RecommendHeaders) SetCommonHeaders(v map[string]*string) *RecommendHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecommendHeaders) SetRegionId(v string) *RecommendHeaders {
	s.RegionId = &v
	return s
}

type RecommendRequest struct {
	SceneId     *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Imei        *string `json:"Imei,omitempty" xml:"Imei,omitempty"`
	ReturnCount *int32  `json:"ReturnCount,omitempty" xml:"ReturnCount,omitempty"`
	Items       *string `json:"Items,omitempty" xml:"Items,omitempty"`
}

func (s RecommendRequest) String() string {
	return tea.Prettify(s)
}

func (s RecommendRequest) GoString() string {
	return s.String()
}

func (s *RecommendRequest) SetSceneId(v string) *RecommendRequest {
	s.SceneId = &v
	return s
}

func (s *RecommendRequest) SetUserId(v string) *RecommendRequest {
	s.UserId = &v
	return s
}

func (s *RecommendRequest) SetIp(v string) *RecommendRequest {
	s.Ip = &v
	return s
}

func (s *RecommendRequest) SetImei(v string) *RecommendRequest {
	s.Imei = &v
	return s
}

func (s *RecommendRequest) SetReturnCount(v int32) *RecommendRequest {
	s.ReturnCount = &v
	return s
}

func (s *RecommendRequest) SetItems(v string) *RecommendRequest {
	s.Items = &v
	return s
}

type RecommendResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    []*RecommendResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s RecommendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecommendResponseBody) GoString() string {
	return s.String()
}

func (s *RecommendResponseBody) SetMessage(v string) *RecommendResponseBody {
	s.Message = &v
	return s
}

func (s *RecommendResponseBody) SetRequestId(v string) *RecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecommendResponseBody) SetCode(v string) *RecommendResponseBody {
	s.Code = &v
	return s
}

func (s *RecommendResponseBody) SetResult(v []*RecommendResponseBodyResult) *RecommendResponseBody {
	s.Result = v
	return s
}

type RecommendResponseBodyResult struct {
	TraceInfo *string  `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty"`
	Weight    *float32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	MatchInfo *string  `json:"MatchInfo,omitempty" xml:"MatchInfo,omitempty"`
	ItemType  *string  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Position  *int32   `json:"Position,omitempty" xml:"Position,omitempty"`
	ItemId    *string  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	TraceId   *string  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RecommendResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RecommendResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RecommendResponseBodyResult) SetTraceInfo(v string) *RecommendResponseBodyResult {
	s.TraceInfo = &v
	return s
}

func (s *RecommendResponseBodyResult) SetWeight(v float32) *RecommendResponseBodyResult {
	s.Weight = &v
	return s
}

func (s *RecommendResponseBodyResult) SetMatchInfo(v string) *RecommendResponseBodyResult {
	s.MatchInfo = &v
	return s
}

func (s *RecommendResponseBodyResult) SetItemType(v string) *RecommendResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *RecommendResponseBodyResult) SetPosition(v int32) *RecommendResponseBodyResult {
	s.Position = &v
	return s
}

func (s *RecommendResponseBodyResult) SetItemId(v string) *RecommendResponseBodyResult {
	s.ItemId = &v
	return s
}

func (s *RecommendResponseBodyResult) SetTraceId(v string) *RecommendResponseBodyResult {
	s.TraceId = &v
	return s
}

type RecommendResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecommendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecommendResponse) String() string {
	return tea.Prettify(s)
}

func (s RecommendResponse) GoString() string {
	return s.String()
}

func (s *RecommendResponse) SetHeaders(v map[string]*string) *RecommendResponse {
	s.Headers = v
	return s
}

func (s *RecommendResponse) SetBody(v *RecommendResponseBody) *RecommendResponse {
	s.Body = v
	return s
}

type RunInstanceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RunInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RunInstanceResponseBody) SetMessage(v string) *RunInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RunInstanceResponseBody) SetRequestId(v string) *RunInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunInstanceResponseBody) SetCode(v string) *RunInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RunInstanceResponseBody) SetResult(v bool) *RunInstanceResponseBody {
	s.Result = &v
	return s
}

type RunInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RunInstanceResponse) GoString() string {
	return s.String()
}

func (s *RunInstanceResponse) SetHeaders(v map[string]*string) *RunInstanceResponse {
	s.Headers = v
	return s
}

func (s *RunInstanceResponse) SetBody(v *RunInstanceResponseBody) *RunInstanceResponse {
	s.Body = v
	return s
}

type StopDataSetResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *StopDataSetResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s StopDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *StopDataSetResponseBody) SetMessage(v string) *StopDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *StopDataSetResponseBody) SetRequestId(v string) *StopDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDataSetResponseBody) SetCode(v string) *StopDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *StopDataSetResponseBody) SetResult(v *StopDataSetResponseBodyResult) *StopDataSetResponseBody {
	s.Result = v
	return s
}

type StopDataSetResponseBodyResult struct {
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s StopDataSetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s StopDataSetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *StopDataSetResponseBodyResult) SetState(v string) *StopDataSetResponseBodyResult {
	s.State = &v
	return s
}

func (s *StopDataSetResponseBodyResult) SetGmtCreate(v int64) *StopDataSetResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *StopDataSetResponseBodyResult) SetVersionId(v string) *StopDataSetResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *StopDataSetResponseBodyResult) SetInstanceId(v string) *StopDataSetResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *StopDataSetResponseBodyResult) SetGmtModified(v int64) *StopDataSetResponseBodyResult {
	s.GmtModified = &v
	return s
}

type StopDataSetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDataSetResponse) GoString() string {
	return s.String()
}

func (s *StopDataSetResponse) SetHeaders(v map[string]*string) *StopDataSetResponse {
	s.Headers = v
	return s
}

func (s *StopDataSetResponse) SetBody(v *StopDataSetResponseBody) *StopDataSetResponse {
	s.Body = v
	return s
}

type UpgradeInstanceResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *UpgradeInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) SetMessage(v string) *UpgradeInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetCode(v string) *UpgradeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetResult(v *UpgradeInstanceResponseBodyResult) *UpgradeInstanceResponseBody {
	s.Result = v
	return s
}

type UpgradeInstanceResponseBodyResult struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpgradeInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBodyResult) SetInstanceId(v string) *UpgradeInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

type UpgradeInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) SetHeaders(v map[string]*string) *UpgradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceResponse) SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse {
	s.Body = v
	return s
}

type ValidateInstanceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateInstanceResponseBody) SetMessage(v string) *ValidateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateInstanceResponseBody) SetRequestId(v string) *ValidateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateInstanceResponseBody) SetCode(v string) *ValidateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateInstanceResponseBody) SetResult(v bool) *ValidateInstanceResponseBody {
	s.Result = &v
	return s
}

type ValidateInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateInstanceResponse) GoString() string {
	return s.String()
}

func (s *ValidateInstanceResponse) SetHeaders(v map[string]*string) *ValidateInstanceResponse {
	s.Headers = v
	return s
}

func (s *ValidateInstanceResponse) SetBody(v *ValidateInstanceResponseBody) *ValidateInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("airec"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AttachDataset(InstanceId *string, VersionId *string) (_result *AttachDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AttachDatasetResponse{}
	_body, _err := client.AttachDatasetWithOptions(InstanceId, VersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDatasetWithOptions(InstanceId *string, VersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AttachDatasetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &AttachDatasetResponse{}
	_body, _err := client.DoROARequest(tea.String("AttachDataset"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSets/{VersionId}/actions/current"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiversify(InstanceId *string) (_result *CreateDiversifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDiversifyResponse{}
	_body, _err := client.CreateDiversifyWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDiversifyWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDiversifyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateDiversifyResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDiversify"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/diversifies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance() (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMix(InstanceId *string) (_result *CreateMixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMixResponse{}
	_body, _err := client.CreateMixWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMixWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMixResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateMixResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateMix"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/mixes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(InstanceId *string) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateRule"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScene(InstanceId *string, request *CreateSceneRequest) (_result *CreateSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSceneResponse{}
	_body, _err := client.CreateSceneWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSceneWithOptions(InstanceId *string, request *CreateSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateSceneResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateScene"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataSet(InstanceId *string, VersionId *string) (_result *DeleteDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DeleteDataSetWithOptions(InstanceId, VersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataSetWithOptions(InstanceId *string, VersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDataSet"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSets/{VersionId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDiversify(InstanceId *string, Name *string) (_result *DeleteDiversifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDiversifyResponse{}
	_body, _err := client.DeleteDiversifyWithOptions(InstanceId, Name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDiversifyWithOptions(InstanceId *string, Name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDiversifyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteDiversifyResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDiversify"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/diversifies/{Name}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMix(InstanceId *string, Name *string) (_result *DeleteMixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMixResponse{}
	_body, _err := client.DeleteMixWithOptions(InstanceId, Name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMixWithOptions(InstanceId *string, Name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMixResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteMixResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteMix"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/mixes/{Name}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScene(InstanceId *string, SceneId *string) (_result *DeleteSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSceneResponse{}
	_body, _err := client.DeleteSceneWithOptions(InstanceId, SceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSceneWithOptions(InstanceId *string, SceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSceneResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteSceneResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteScene"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes/{SceneId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataSetMessage(InstanceId *string, VersionId *string) (_result *DescribeDataSetMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDataSetMessageResponse{}
	_body, _err := client.DescribeDataSetMessageWithOptions(InstanceId, VersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataSetMessageWithOptions(InstanceId *string, VersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDataSetMessageResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeDataSetMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeDataSetMessage"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSets/{VersionId}/messages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataSetReport(InstanceId *string, VersionId *string) (_result *DescribeDataSetReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDataSetReportResponse{}
	_body, _err := client.DescribeDataSetReportWithOptions(InstanceId, VersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataSetReportWithOptions(InstanceId *string, VersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDataSetReportResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeDataSetReportResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeDataSetReport"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSets/{VersionId}/report"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiversify(InstanceId *string, Name *string) (_result *DescribeDiversifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDiversifyResponse{}
	_body, _err := client.DescribeDiversifyWithOptions(InstanceId, Name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiversifyWithOptions(InstanceId *string, Name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDiversifyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeDiversifyResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeDiversify"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/diversifies/{Name}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExposureSettings(InstanceId *string) (_result *DescribeExposureSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeExposureSettingsResponse{}
	_body, _err := client.DescribeExposureSettingsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExposureSettingsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeExposureSettingsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeExposureSettingsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeExposureSettings"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/exposure-settings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(InstanceId *string) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMix(InstanceId *string, Name *string) (_result *DescribeMixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeMixResponse{}
	_body, _err := client.DescribeMixWithOptions(InstanceId, Name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMixWithOptions(InstanceId *string, Name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeMixResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeMixResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeMix"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/mixes/{Name}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQuota(InstanceId *string) (_result *DescribeQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQuotaResponse{}
	_body, _err := client.DescribeQuotaWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQuotaWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeQuota"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/quota"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeRegions"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/configurations/regions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRule(InstanceId *string, RuleId *string, request *DescribeRuleRequest) (_result *DescribeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRuleResponse{}
	_body, _err := client.DescribeRuleWithOptions(InstanceId, RuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleWithOptions(InstanceId *string, RuleId *string, request *DescribeRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeRule"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rules/{RuleId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScene(InstanceId *string, SceneId *string) (_result *DescribeSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSceneResponse{}
	_body, _err := client.DescribeSceneWithOptions(InstanceId, SceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSceneWithOptions(InstanceId *string, SceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSceneResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeSceneResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeScene"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes/{SceneId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSceneThroughput(InstanceId *string, SceneId *string) (_result *DescribeSceneThroughputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSceneThroughputResponse{}
	_body, _err := client.DescribeSceneThroughputWithOptions(InstanceId, SceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSceneThroughputWithOptions(InstanceId *string, SceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSceneThroughputResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DescribeSceneThroughputResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeSceneThroughput"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes/{SceneId}/throughput"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSyncReportDetail(InstanceId *string, request *DescribeSyncReportDetailRequest) (_result *DescribeSyncReportDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSyncReportDetailResponse{}
	_body, _err := client.DescribeSyncReportDetailWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSyncReportDetailWithOptions(InstanceId *string, request *DescribeSyncReportDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSyncReportDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.LevelType)) {
		query["LevelType"] = request.LevelType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeSyncReportDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeSyncReportDetail"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/sync-reports/detail"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSyncReportOutliers(InstanceId *string, request *DescribeSyncReportOutliersRequest) (_result *DescribeSyncReportOutliersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSyncReportOutliersResponse{}
	_body, _err := client.DescribeSyncReportOutliersWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSyncReportOutliersWithOptions(InstanceId *string, request *DescribeSyncReportOutliersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSyncReportOutliersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LevelType)) {
		query["LevelType"] = request.LevelType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeSyncReportOutliersResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeSyncReportOutliers"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/sync-reports/outliers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserMetrics(InstanceId *string, request *DescribeUserMetricsRequest) (_result *DescribeUserMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserMetricsResponse{}
	_body, _err := client.DescribeUserMetricsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserMetricsWithOptions(InstanceId *string, request *DescribeUserMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUserMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DescribeUserMetricsResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeUserMetrics"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/metrics"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DowngradeInstance(InstanceId *string) (_result *DowngradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DowngradeInstanceResponse{}
	_body, _err := client.DowngradeInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DowngradeInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DowngradeInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DowngradeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DowngradeInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/actions/downgrade"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboard(InstanceId *string, request *ListDashboardRequest) (_result *ListDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardResponse{}
	_body, _err := client.ListDashboardWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardWithOptions(InstanceId *string, request *ListDashboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["TraceId"] = request.TraceId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListDashboardResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboard"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/statistics"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardDetails(InstanceId *string, request *ListDashboardDetailsRequest) (_result *ListDashboardDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardDetailsResponse{}
	_body, _err := client.ListDashboardDetailsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardDetailsWithOptions(InstanceId *string, request *ListDashboardDetailsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TraceIds)) {
		query["TraceIds"] = request.TraceIds
	}

	if !tea.BoolValue(util.IsUnset(request.SceneIds)) {
		query["SceneIds"] = request.SceneIds
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListDashboardDetailsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboardDetails"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/details"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardDetailsFlows(InstanceId *string, request *ListDashboardDetailsFlowsRequest) (_result *ListDashboardDetailsFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardDetailsFlowsResponse{}
	_body, _err := client.ListDashboardDetailsFlowsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardDetailsFlowsWithOptions(InstanceId *string, request *ListDashboardDetailsFlowsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardDetailsFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TraceIds)) {
		query["TraceIds"] = request.TraceIds
	}

	if !tea.BoolValue(util.IsUnset(request.SceneIds)) {
		query["SceneIds"] = request.SceneIds
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListDashboardDetailsFlowsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboardDetailsFlows"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/details/flows"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardMetrics(InstanceId *string, request *ListDashboardMetricsRequest) (_result *ListDashboardMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardMetricsResponse{}
	_body, _err := client.ListDashboardMetricsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardMetricsWithOptions(InstanceId *string, request *ListDashboardMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListDashboardMetricsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboardMetrics"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/metrics"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardMetricsFlows(InstanceId *string, request *ListDashboardMetricsFlowsRequest) (_result *ListDashboardMetricsFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardMetricsFlowsResponse{}
	_body, _err := client.ListDashboardMetricsFlowsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardMetricsFlowsWithOptions(InstanceId *string, request *ListDashboardMetricsFlowsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardMetricsFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListDashboardMetricsFlowsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboardMetricsFlows"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/metrics/flows"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardParameters(InstanceId *string) (_result *ListDashboardParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardParametersResponse{}
	_body, _err := client.ListDashboardParametersWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardParametersWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardParametersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListDashboardParametersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboardParameters"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/parameters"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardUid(InstanceId *string) (_result *ListDashboardUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDashboardUidResponse{}
	_body, _err := client.ListDashboardUidWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardUidWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDashboardUidResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListDashboardUidResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDashboardUid"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dashboard/uid"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSet(InstanceId *string) (_result *ListDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSetResponse{}
	_body, _err := client.ListDataSetWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSetWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListDataSetResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDataSet"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSource(InstanceId *string) (_result *ListDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceResponse{}
	_body, _err := client.ListDataSourceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListDataSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDataSource"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiversify(InstanceId *string) (_result *ListDiversifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiversifyResponse{}
	_body, _err := client.ListDiversifyWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiversifyWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDiversifyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListDiversifyResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDiversify"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/diversifies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceTask(InstanceId *string) (_result *ListInstanceTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceTaskResponse{}
	_body, _err := client.ListInstanceTaskWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceTaskWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListInstanceTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstanceTask"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListItems(InstanceId *string, request *ListItemsRequest) (_result *ListItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListItemsResponse{}
	_body, _err := client.ListItemsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListItemsWithOptions(InstanceId *string, request *ListItemsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListItemsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListItems"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/items/actions/list"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogs(InstanceId *string, request *ListLogsRequest) (_result *ListLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogsResponse{}
	_body, _err := client.ListLogsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogsWithOptions(InstanceId *string, request *ListLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryParams)) {
		query["QueryParams"] = request.QueryParams
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListLogsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListLogs"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/logs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMix(InstanceId *string) (_result *ListMixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMixResponse{}
	_body, _err := client.ListMixWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMixWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMixResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListMixResponse{}
	_body, _err := client.DoROARequest(tea.String("ListMix"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/mixes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRuleConditions(InstanceId *string) (_result *ListRuleConditionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRuleConditionsResponse{}
	_body, _err := client.ListRuleConditionsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRuleConditionsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRuleConditionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListRuleConditionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRuleConditions"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rule-conditions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(InstanceId *string, request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(InstanceId *string, request *ListRulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRules"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRuleTasks(InstanceId *string, request *ListRuleTasksRequest) (_result *ListRuleTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRuleTasksResponse{}
	_body, _err := client.ListRuleTasksWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRuleTasksWithOptions(InstanceId *string, request *ListRuleTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRuleTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRuleTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRuleTasks"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rule-tasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSceneItems(InstanceId *string, SceneId *string, request *ListSceneItemsRequest) (_result *ListSceneItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSceneItemsResponse{}
	_body, _err := client.ListSceneItemsWithOptions(InstanceId, SceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSceneItemsWithOptions(InstanceId *string, SceneId *string, request *ListSceneItemsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSceneItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationRuleId)) {
		query["OperationRuleId"] = request.OperationRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SelectionRuleId)) {
		query["SelectionRuleId"] = request.SelectionRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewType)) {
		query["PreviewType"] = request.PreviewType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCount)) {
		query["QueryCount"] = request.QueryCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSceneItemsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSceneItems"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes/{SceneId}/items"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScenes(InstanceId *string, request *ListScenesRequest) (_result *ListScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScenesResponse{}
	_body, _err := client.ListScenesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScenesWithOptions(InstanceId *string, request *ListScenesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListScenesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListScenes"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUmengAppkeys() (_result *ListUmengAppkeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUmengAppkeysResponse{}
	_body, _err := client.ListUmengAppkeysWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUmengAppkeysWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUmengAppkeysResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListUmengAppkeysResponse{}
	_body, _err := client.DoROARequest(tea.String("ListUmengAppkeys"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/umeng/appkeys"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataSource(InstanceId *string, TableName *string) (_result *ModifyDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.ModifyDataSourceWithOptions(InstanceId, TableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataSourceWithOptions(InstanceId *string, TableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDataSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyDataSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyDataSource"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSources/{TableName}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDiversify(InstanceId *string, Name *string) (_result *ModifyDiversifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiversifyResponse{}
	_body, _err := client.ModifyDiversifyWithOptions(InstanceId, Name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDiversifyWithOptions(InstanceId *string, Name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDiversifyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyDiversifyResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyDiversify"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/diversifies/{Name}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyExposureSettings(InstanceId *string) (_result *ModifyExposureSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyExposureSettingsResponse{}
	_body, _err := client.ModifyExposureSettingsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyExposureSettingsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyExposureSettingsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyExposureSettingsResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyExposureSettings"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/exposure-settings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstance(InstanceId *string) (_result *ModifyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyItems(InstanceId *string) (_result *ModifyItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyItemsResponse{}
	_body, _err := client.ModifyItemsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyItemsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyItemsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyItemsResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyItems"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/items"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMix(InstanceId *string, Name *string) (_result *ModifyMixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyMixResponse{}
	_body, _err := client.ModifyMixWithOptions(InstanceId, Name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMixWithOptions(InstanceId *string, Name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyMixResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyMixResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyMix"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/mixes/{Name}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRule(InstanceId *string, RuleId *string) (_result *ModifyRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyRuleResponse{}
	_body, _err := client.ModifyRuleWithOptions(InstanceId, RuleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRuleWithOptions(InstanceId *string, RuleId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyRule"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rules/{RuleId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScene(InstanceId *string, SceneId *string) (_result *ModifySceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifySceneResponse{}
	_body, _err := client.ModifySceneWithOptions(InstanceId, SceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySceneWithOptions(InstanceId *string, SceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifySceneResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifySceneResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyScene"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/scenes/{SceneId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishRule(InstanceId *string, RuleId *string, request *PublishRuleRequest) (_result *PublishRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishRuleResponse{}
	_body, _err := client.PublishRuleWithOptions(InstanceId, RuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishRuleWithOptions(InstanceId *string, RuleId *string, request *PublishRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &PublishRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("PublishRule"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/rules/{RuleId}/actions/publish"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushDocument(InstanceId *string, TableName *string) (_result *PushDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushDocumentResponse{}
	_body, _err := client.PushDocumentWithOptions(InstanceId, TableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushDocumentWithOptions(InstanceId *string, TableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushDocumentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &PushDocumentResponse{}
	_body, _err := client.DoROARequest(tea.String("PushDocument"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/tables/{TableName}/actions/bulk"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushIntervention(InstanceId *string) (_result *PushInterventionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInterventionResponse{}
	_body, _err := client.PushInterventionWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushInterventionWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInterventionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &PushInterventionResponse{}
	_body, _err := client.DoROARequest(tea.String("PushIntervention"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/actions/intervene"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDataMessage(InstanceId *string, Table *string, request *QueryDataMessageRequest) (_result *QueryDataMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDataMessageResponse{}
	_body, _err := client.QueryDataMessageWithOptions(InstanceId, Table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDataMessageWithOptions(InstanceId *string, Table *string, request *QueryDataMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryDataMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CmdType)) {
		query["CmdType"] = request.CmdType
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemType)) {
		query["ItemType"] = request.ItemType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["TraceId"] = request.TraceId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.BhvType)) {
		query["BhvType"] = request.BhvType
	}

	if !tea.BoolValue(util.IsUnset(request.MessageSource)) {
		query["MessageSource"] = request.MessageSource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryDataMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDataMessage"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/tables/{Table}/data-message"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDataMessageStatistics(InstanceId *string, Table *string, request *QueryDataMessageStatisticsRequest) (_result *QueryDataMessageStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDataMessageStatisticsResponse{}
	_body, _err := client.QueryDataMessageStatisticsWithOptions(InstanceId, Table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDataMessageStatisticsWithOptions(InstanceId *string, Table *string, request *QueryDataMessageStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryDataMessageStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CmdType)) {
		query["CmdType"] = request.CmdType
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemType)) {
		query["ItemType"] = request.ItemType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		query["TraceId"] = request.TraceId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.BhvType)) {
		query["BhvType"] = request.BhvType
	}

	if !tea.BoolValue(util.IsUnset(request.MessageSource)) {
		query["MessageSource"] = request.MessageSource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryDataMessageStatisticsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDataMessageStatistics"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/tables/{Table}/data-message-statistics"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryExceptionHistory(InstanceId *string, request *QueryExceptionHistoryRequest) (_result *QueryExceptionHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryExceptionHistoryResponse{}
	_body, _err := client.QueryExceptionHistoryWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryExceptionHistoryWithOptions(InstanceId *string, request *QueryExceptionHistoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryExceptionHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryExceptionHistoryResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryExceptionHistory"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/sync-reports/exception-history"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRawData(InstanceId *string, Table *string, request *QueryRawDataRequest) (_result *QueryRawDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRawDataResponse{}
	_body, _err := client.QueryRawDataWithOptions(InstanceId, Table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRawDataWithOptions(InstanceId *string, Table *string, request *QueryRawDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryRawDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemType)) {
		query["ItemType"] = request.ItemType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryRawDataResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryRawData"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/tables/{Table}/raw-data"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySingleAggregationReport(InstanceId *string) (_result *QuerySingleAggregationReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySingleAggregationReportResponse{}
	_body, _err := client.QuerySingleAggregationReportWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySingleAggregationReportWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QuerySingleAggregationReportResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &QuerySingleAggregationReportResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySingleAggregationReport"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/sync-reports/single-aggregation-report"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySingleReport(InstanceId *string, request *QuerySingleReportRequest) (_result *QuerySingleReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySingleReportResponse{}
	_body, _err := client.QuerySingleReportWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySingleReportWithOptions(InstanceId *string, request *QuerySingleReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QuerySingleReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		query["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &QuerySingleReportResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySingleReport"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/sync-reports/single-report"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySyncReportAggregation(InstanceId *string, request *QuerySyncReportAggregationRequest) (_result *QuerySyncReportAggregationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySyncReportAggregationResponse{}
	_body, _err := client.QuerySyncReportAggregationWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySyncReportAggregationWithOptions(InstanceId *string, request *QuerySyncReportAggregationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QuerySyncReportAggregationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &QuerySyncReportAggregationResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySyncReportAggregation"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/sync-reports/aggregation"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Recommend(InstanceId *string, request *RecommendRequest) (_result *RecommendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecommendHeaders{}
	_result = &RecommendResponse{}
	_body, _err := client.RecommendWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecommendWithOptions(InstanceId *string, request *RecommendRequest, headers *RecommendHeaders, runtime *util.RuntimeOptions) (_result *RecommendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Imei)) {
		query["Imei"] = request.Imei
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnCount)) {
		query["ReturnCount"] = request.ReturnCount
	}

	if !tea.BoolValue(util.IsUnset(request.Items)) {
		query["Items"] = request.Items
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.RegionId)) {
		realHeaders["RegionId"] = headers.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &RecommendResponse{}
	_body, _err := client.DoROARequest(tea.String("Recommend"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/actions/recommend"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunInstance(InstanceId *string) (_result *RunInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunInstanceResponse{}
	_body, _err := client.RunInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &RunInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("RunInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/actions/import"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDataSet(InstanceId *string, VersionId *string) (_result *StopDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopDataSetResponse{}
	_body, _err := client.StopDataSetWithOptions(InstanceId, VersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDataSetWithOptions(InstanceId *string, VersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopDataSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &StopDataSetResponse{}
	_body, _err := client.DoROARequest(tea.String("StopDataSet"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/dataSets/{VersionId}/actions/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstance(InstanceId *string) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpgradeInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/actions/upgrade"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateInstance(InstanceId *string) (_result *ValidateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateInstanceResponse{}
	_body, _err := client.ValidateInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ValidateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateInstance"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/openapi/instances/"+tea.StringValue(InstanceId)+"/actions/validate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
