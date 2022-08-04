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

type CreateDatasetRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetName(v string) *CreateDatasetRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequest) SetType(v string) *CreateDatasetRequest {
	s.Type = &v
	return s
}

func (s *CreateDatasetRequest) SetWorkspaceId(v int64) *CreateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

type CreateDatasetResponseBody struct {
	Data      *CreateDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) SetData(v *CreateDatasetResponseBodyData) *CreateDatasetResponseBody {
	s.Data = v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatasetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDatasetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBodyData) SetDescription(v string) *CreateDatasetResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetGmtCreate(v int64) *CreateDatasetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetId(v int64) *CreateDatasetResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateDatasetResponseBodyData) SetName(v string) *CreateDatasetResponseBodyData {
	s.Name = &v
	return s
}

type CreateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponse) SetHeaders(v map[string]*string) *CreateDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetResponse) SetStatusCode(v int32) *CreateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateLabelsetRequest struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectKey   *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	TagSettings *string `json:"TagSettings,omitempty" xml:"TagSettings,omitempty"`
	TagUserList *string `json:"TagUserList,omitempty" xml:"TagUserList,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserOssUrl  *string `json:"UserOssUrl,omitempty" xml:"UserOssUrl,omitempty"`
}

func (s CreateLabelsetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelsetRequest) GoString() string {
	return s.String()
}

func (s *CreateLabelsetRequest) SetDatasetId(v int64) *CreateLabelsetRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateLabelsetRequest) SetDescription(v string) *CreateLabelsetRequest {
	s.Description = &v
	return s
}

func (s *CreateLabelsetRequest) SetName(v string) *CreateLabelsetRequest {
	s.Name = &v
	return s
}

func (s *CreateLabelsetRequest) SetObjectKey(v string) *CreateLabelsetRequest {
	s.ObjectKey = &v
	return s
}

func (s *CreateLabelsetRequest) SetTagSettings(v string) *CreateLabelsetRequest {
	s.TagSettings = &v
	return s
}

func (s *CreateLabelsetRequest) SetTagUserList(v string) *CreateLabelsetRequest {
	s.TagUserList = &v
	return s
}

func (s *CreateLabelsetRequest) SetType(v string) *CreateLabelsetRequest {
	s.Type = &v
	return s
}

func (s *CreateLabelsetRequest) SetUserOssUrl(v string) *CreateLabelsetRequest {
	s.UserOssUrl = &v
	return s
}

type CreateLabelsetResponseBody struct {
	Data      *CreateLabelsetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLabelsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelsetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabelsetResponseBody) SetData(v *CreateLabelsetResponseBodyData) *CreateLabelsetResponseBody {
	s.Data = v
	return s
}

func (s *CreateLabelsetResponseBody) SetRequestId(v string) *CreateLabelsetResponseBody {
	s.RequestId = &v
	return s
}

type CreateLabelsetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelType   *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLabelsetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelsetResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLabelsetResponseBodyData) SetDescription(v string) *CreateLabelsetResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateLabelsetResponseBodyData) SetGmtCreate(v int64) *CreateLabelsetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateLabelsetResponseBodyData) SetId(v int64) *CreateLabelsetResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateLabelsetResponseBodyData) SetLabelType(v string) *CreateLabelsetResponseBodyData {
	s.LabelType = &v
	return s
}

func (s *CreateLabelsetResponseBodyData) SetName(v string) *CreateLabelsetResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateLabelsetResponseBodyData) SetStatus(v string) *CreateLabelsetResponseBodyData {
	s.Status = &v
	return s
}

type CreateLabelsetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLabelsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLabelsetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelsetResponse) GoString() string {
	return s.String()
}

func (s *CreateLabelsetResponse) SetHeaders(v map[string]*string) *CreateLabelsetResponse {
	s.Headers = v
	return s
}

func (s *CreateLabelsetResponse) SetStatusCode(v int32) *CreateLabelsetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLabelsetResponse) SetBody(v *CreateLabelsetResponseBody) *CreateLabelsetResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TrainTaskId *int64  `json:"TrainTaskId,omitempty" xml:"TrainTaskId,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetDescription(v string) *CreateServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateServiceRequest) SetName(v string) *CreateServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceRequest) SetTrainTaskId(v int64) *CreateServiceRequest {
	s.TrainTaskId = &v
	return s
}

type CreateServiceResponseBody struct {
	Data      *CreateServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetData(v *CreateServiceResponseBodyData) *CreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceResponseBodyData struct {
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceName        *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyData) SetGmtCreate(v int64) *CreateServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateServiceResponseBodyData) SetId(v int64) *CreateServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateServiceResponseBodyData) SetServiceDescription(v string) *CreateServiceResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *CreateServiceResponseBodyData) SetServiceName(v string) *CreateServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceResponseBodyData) SetStatus(v string) *CreateServiceResponseBodyData {
	s.Status = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateTagTaskRequest struct {
	LabelsetId *int64 `json:"LabelsetId,omitempty" xml:"LabelsetId,omitempty"`
}

func (s CreateTagTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTagTaskRequest) SetLabelsetId(v int64) *CreateTagTaskRequest {
	s.LabelsetId = &v
	return s
}

type CreateTagTaskResponseBody struct {
	Data      *CreateTagTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagTaskResponseBody) SetData(v *CreateTagTaskResponseBodyData) *CreateTagTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTagTaskResponseBody) SetRequestId(v string) *CreateTagTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateTagTaskResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelType   *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateTagTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTagTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTagTaskResponseBodyData) SetDescription(v string) *CreateTagTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTagTaskResponseBodyData) SetGmtCreate(v int64) *CreateTagTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateTagTaskResponseBodyData) SetId(v int64) *CreateTagTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTagTaskResponseBodyData) SetLabelType(v string) *CreateTagTaskResponseBodyData {
	s.LabelType = &v
	return s
}

func (s *CreateTagTaskResponseBodyData) SetName(v string) *CreateTagTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTagTaskResponseBodyData) SetStatus(v string) *CreateTagTaskResponseBodyData {
	s.Status = &v
	return s
}

type CreateTagTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTagTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTagTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTagTaskResponse) SetHeaders(v map[string]*string) *CreateTagTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTagTaskResponse) SetStatusCode(v int32) *CreateTagTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagTaskResponse) SetBody(v *CreateTagTaskResponseBody) *CreateTagTaskResponse {
	s.Body = v
	return s
}

type CreateTrainTaskRequest struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TrainMode   *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	WorkspaceId *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskRequest) SetDatasetId(v int64) *CreateTrainTaskRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateTrainTaskRequest) SetDescription(v string) *CreateTrainTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTrainTaskRequest) SetLabelId(v int64) *CreateTrainTaskRequest {
	s.LabelId = &v
	return s
}

func (s *CreateTrainTaskRequest) SetName(v string) *CreateTrainTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTrainTaskRequest) SetTrainMode(v string) *CreateTrainTaskRequest {
	s.TrainMode = &v
	return s
}

func (s *CreateTrainTaskRequest) SetWorkspaceId(v int64) *CreateTrainTaskRequest {
	s.WorkspaceId = &v
	return s
}

type CreateTrainTaskResponseBody struct {
	Data      *CreateTrainTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponseBody) SetData(v *CreateTrainTaskResponseBodyData) *CreateTrainTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTrainTaskResponseBody) SetRequestId(v string) *CreateTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateTrainTaskResponseBodyData struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName   *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	ModelEffect *string `json:"ModelEffect,omitempty" xml:"ModelEffect,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TrainMode   *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	TrainStatus *string `json:"TrainStatus,omitempty" xml:"TrainStatus,omitempty"`
}

func (s CreateTrainTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponseBodyData) SetDatasetId(v int64) *CreateTrainTaskResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetDatasetName(v string) *CreateTrainTaskResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetDescription(v string) *CreateTrainTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetGmtCreate(v int64) *CreateTrainTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetId(v int64) *CreateTrainTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetLabelId(v int64) *CreateTrainTaskResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetLabelName(v string) *CreateTrainTaskResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetModelEffect(v string) *CreateTrainTaskResponseBodyData {
	s.ModelEffect = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetModelId(v int64) *CreateTrainTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetTaskName(v string) *CreateTrainTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetTrainMode(v string) *CreateTrainTaskResponseBodyData {
	s.TrainMode = &v
	return s
}

func (s *CreateTrainTaskResponseBodyData) SetTrainStatus(v string) *CreateTrainTaskResponseBodyData {
	s.TrainStatus = &v
	return s
}

type CreateTrainTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponse) SetHeaders(v map[string]*string) *CreateTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainTaskResponse) SetStatusCode(v int32) *CreateTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainTaskResponse) SetBody(v *CreateTrainTaskResponseBody) *CreateTrainTaskResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetType(v string) *CreateWorkspaceRequest {
	s.Type = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	Data      *CreateWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetData(v *CreateWorkspaceResponseBodyData) *CreateWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateWorkspaceResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWorkspaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyData) SetDescription(v string) *CreateWorkspaceResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetGmtCreate(v int64) *CreateWorkspaceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetId(v int64) *CreateWorkspaceResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetName(v string) *CreateWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetType(v string) *CreateWorkspaceResponseBodyData {
	s.Type = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type CustomizeClassifyImageRequest struct {
	ImageUrl  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CustomizeClassifyImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeClassifyImageRequest) GoString() string {
	return s.String()
}

func (s *CustomizeClassifyImageRequest) SetImageUrl(v string) *CustomizeClassifyImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *CustomizeClassifyImageRequest) SetServiceId(v string) *CustomizeClassifyImageRequest {
	s.ServiceId = &v
	return s
}

type CustomizeClassifyImageAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	ServiceId      *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CustomizeClassifyImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeClassifyImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CustomizeClassifyImageAdvanceRequest) SetImageUrlObject(v io.Reader) *CustomizeClassifyImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *CustomizeClassifyImageAdvanceRequest) SetServiceId(v string) *CustomizeClassifyImageAdvanceRequest {
	s.ServiceId = &v
	return s
}

type CustomizeClassifyImageResponseBody struct {
	Data      *CustomizeClassifyImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomizeClassifyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeClassifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeClassifyImageResponseBody) SetData(v *CustomizeClassifyImageResponseBodyData) *CustomizeClassifyImageResponseBody {
	s.Data = v
	return s
}

func (s *CustomizeClassifyImageResponseBody) SetRequestId(v string) *CustomizeClassifyImageResponseBody {
	s.RequestId = &v
	return s
}

type CustomizeClassifyImageResponseBodyData struct {
	Category *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	Score    *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s CustomizeClassifyImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CustomizeClassifyImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomizeClassifyImageResponseBodyData) SetCategory(v string) *CustomizeClassifyImageResponseBodyData {
	s.Category = &v
	return s
}

func (s *CustomizeClassifyImageResponseBodyData) SetScore(v float32) *CustomizeClassifyImageResponseBodyData {
	s.Score = &v
	return s
}

type CustomizeClassifyImageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeClassifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeClassifyImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeClassifyImageResponse) GoString() string {
	return s.String()
}

func (s *CustomizeClassifyImageResponse) SetHeaders(v map[string]*string) *CustomizeClassifyImageResponse {
	s.Headers = v
	return s
}

func (s *CustomizeClassifyImageResponse) SetStatusCode(v int32) *CustomizeClassifyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeClassifyImageResponse) SetBody(v *CustomizeClassifyImageResponseBody) *CustomizeClassifyImageResponse {
	s.Body = v
	return s
}

type CustomizeDetectImageRequest struct {
	ImageUrl  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CustomizeDetectImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageRequest) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageRequest) SetImageUrl(v string) *CustomizeDetectImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *CustomizeDetectImageRequest) SetServiceId(v string) *CustomizeDetectImageRequest {
	s.ServiceId = &v
	return s
}

type CustomizeDetectImageAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	ServiceId      *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CustomizeDetectImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageAdvanceRequest) SetImageUrlObject(v io.Reader) *CustomizeDetectImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *CustomizeDetectImageAdvanceRequest) SetServiceId(v string) *CustomizeDetectImageAdvanceRequest {
	s.ServiceId = &v
	return s
}

type CustomizeDetectImageResponseBody struct {
	Data      *CustomizeDetectImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomizeDetectImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageResponseBody) SetData(v *CustomizeDetectImageResponseBodyData) *CustomizeDetectImageResponseBody {
	s.Data = v
	return s
}

func (s *CustomizeDetectImageResponseBody) SetRequestId(v string) *CustomizeDetectImageResponseBody {
	s.RequestId = &v
	return s
}

type CustomizeDetectImageResponseBodyData struct {
	Elements []*CustomizeDetectImageResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s CustomizeDetectImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageResponseBodyData) SetElements(v []*CustomizeDetectImageResponseBodyDataElements) *CustomizeDetectImageResponseBodyData {
	s.Elements = v
	return s
}

type CustomizeDetectImageResponseBodyDataElements struct {
	Boxes    *CustomizeDetectImageResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Struct"`
	Category *string                                            `json:"Category,omitempty" xml:"Category,omitempty"`
	Score    *float32                                           `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s CustomizeDetectImageResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageResponseBodyDataElements) SetBoxes(v *CustomizeDetectImageResponseBodyDataElementsBoxes) *CustomizeDetectImageResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *CustomizeDetectImageResponseBodyDataElements) SetCategory(v string) *CustomizeDetectImageResponseBodyDataElements {
	s.Category = &v
	return s
}

func (s *CustomizeDetectImageResponseBodyDataElements) SetScore(v float32) *CustomizeDetectImageResponseBodyDataElements {
	s.Score = &v
	return s
}

type CustomizeDetectImageResponseBodyDataElementsBoxes struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s CustomizeDetectImageResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageResponseBodyDataElementsBoxes) SetHeight(v float32) *CustomizeDetectImageResponseBodyDataElementsBoxes {
	s.Height = &v
	return s
}

func (s *CustomizeDetectImageResponseBodyDataElementsBoxes) SetWidth(v float32) *CustomizeDetectImageResponseBodyDataElementsBoxes {
	s.Width = &v
	return s
}

func (s *CustomizeDetectImageResponseBodyDataElementsBoxes) SetX(v float32) *CustomizeDetectImageResponseBodyDataElementsBoxes {
	s.X = &v
	return s
}

func (s *CustomizeDetectImageResponseBodyDataElementsBoxes) SetY(v float32) *CustomizeDetectImageResponseBodyDataElementsBoxes {
	s.Y = &v
	return s
}

type CustomizeDetectImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeDetectImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeDetectImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeDetectImageResponse) GoString() string {
	return s.String()
}

func (s *CustomizeDetectImageResponse) SetHeaders(v map[string]*string) *CustomizeDetectImageResponse {
	s.Headers = v
	return s
}

func (s *CustomizeDetectImageResponse) SetStatusCode(v int32) *CustomizeDetectImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeDetectImageResponse) SetBody(v *CustomizeDetectImageResponseBody) *CustomizeDetectImageResponse {
	s.Body = v
	return s
}

type CustomizeInstanceSegmentImageRequest struct {
	ImageUrl  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CustomizeInstanceSegmentImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageRequest) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageRequest) SetImageUrl(v string) *CustomizeInstanceSegmentImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *CustomizeInstanceSegmentImageRequest) SetServiceId(v string) *CustomizeInstanceSegmentImageRequest {
	s.ServiceId = &v
	return s
}

type CustomizeInstanceSegmentImageAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	ServiceId      *string   `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CustomizeInstanceSegmentImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageAdvanceRequest) SetImageUrlObject(v io.Reader) *CustomizeInstanceSegmentImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *CustomizeInstanceSegmentImageAdvanceRequest) SetServiceId(v string) *CustomizeInstanceSegmentImageAdvanceRequest {
	s.ServiceId = &v
	return s
}

type CustomizeInstanceSegmentImageResponseBody struct {
	Data      *CustomizeInstanceSegmentImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomizeInstanceSegmentImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponseBody) SetData(v *CustomizeInstanceSegmentImageResponseBodyData) *CustomizeInstanceSegmentImageResponseBody {
	s.Data = v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBody) SetRequestId(v string) *CustomizeInstanceSegmentImageResponseBody {
	s.RequestId = &v
	return s
}

type CustomizeInstanceSegmentImageResponseBodyData struct {
	Elements []*CustomizeInstanceSegmentImageResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s CustomizeInstanceSegmentImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponseBodyData) SetElements(v []*CustomizeInstanceSegmentImageResponseBodyDataElements) *CustomizeInstanceSegmentImageResponseBodyData {
	s.Elements = v
	return s
}

type CustomizeInstanceSegmentImageResponseBodyDataElements struct {
	Boxes    *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes      `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Struct"`
	Category *string                                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	Contours []*CustomizeInstanceSegmentImageResponseBodyDataElementsContours `json:"Contours,omitempty" xml:"Contours,omitempty" type:"Repeated"`
	Mask     *CustomizeInstanceSegmentImageResponseBodyDataElementsMask       `json:"Mask,omitempty" xml:"Mask,omitempty" type:"Struct"`
	Score    *float32                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElements) SetBoxes(v *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) *CustomizeInstanceSegmentImageResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElements) SetCategory(v string) *CustomizeInstanceSegmentImageResponseBodyDataElements {
	s.Category = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElements) SetContours(v []*CustomizeInstanceSegmentImageResponseBodyDataElementsContours) *CustomizeInstanceSegmentImageResponseBodyDataElements {
	s.Contours = v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElements) SetMask(v *CustomizeInstanceSegmentImageResponseBodyDataElementsMask) *CustomizeInstanceSegmentImageResponseBodyDataElements {
	s.Mask = v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElements) SetScore(v float32) *CustomizeInstanceSegmentImageResponseBodyDataElements {
	s.Score = &v
	return s
}

type CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) SetHeight(v int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes {
	s.Height = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) SetWidth(v int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes {
	s.Width = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) SetX(v int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes {
	s.X = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes) SetY(v int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsBoxes {
	s.Y = &v
	return s
}

type CustomizeInstanceSegmentImageResponseBodyDataElementsContours struct {
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElementsContours) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElementsContours) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsContours) SetX(v int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsContours {
	s.X = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsContours) SetY(v int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsContours {
	s.Y = &v
	return s
}

type CustomizeInstanceSegmentImageResponseBodyDataElementsMask struct {
	Counts *string  `json:"Counts,omitempty" xml:"Counts,omitempty"`
	Sizes  []*int32 `json:"Sizes,omitempty" xml:"Sizes,omitempty" type:"Repeated"`
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElementsMask) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponseBodyDataElementsMask) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsMask) SetCounts(v string) *CustomizeInstanceSegmentImageResponseBodyDataElementsMask {
	s.Counts = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponseBodyDataElementsMask) SetSizes(v []*int32) *CustomizeInstanceSegmentImageResponseBodyDataElementsMask {
	s.Sizes = v
	return s
}

type CustomizeInstanceSegmentImageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeInstanceSegmentImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeInstanceSegmentImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeInstanceSegmentImageResponse) GoString() string {
	return s.String()
}

func (s *CustomizeInstanceSegmentImageResponse) SetHeaders(v map[string]*string) *CustomizeInstanceSegmentImageResponse {
	s.Headers = v
	return s
}

func (s *CustomizeInstanceSegmentImageResponse) SetStatusCode(v int32) *CustomizeInstanceSegmentImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeInstanceSegmentImageResponse) SetBody(v *CustomizeInstanceSegmentImageResponseBody) *CustomizeInstanceSegmentImageResponse {
	s.Body = v
	return s
}

type DebugServiceRequest struct {
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s DebugServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugServiceRequest) GoString() string {
	return s.String()
}

func (s *DebugServiceRequest) SetId(v int64) *DebugServiceRequest {
	s.Id = &v
	return s
}

func (s *DebugServiceRequest) SetParam(v string) *DebugServiceRequest {
	s.Param = &v
	return s
}

type DebugServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DebugServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DebugServiceResponseBody) SetRequestId(v string) *DebugServiceResponseBody {
	s.RequestId = &v
	return s
}

type DebugServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DebugServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DebugServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugServiceResponse) GoString() string {
	return s.String()
}

func (s *DebugServiceResponse) SetHeaders(v map[string]*string) *DebugServiceResponse {
	s.Headers = v
	return s
}

func (s *DebugServiceResponse) SetStatusCode(v int32) *DebugServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugServiceResponse) SetBody(v *DebugServiceResponseBody) *DebugServiceResponse {
	s.Body = v
	return s
}

type DeleteDatasetRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) SetId(v int64) *DeleteDatasetRequest {
	s.Id = &v
	return s
}

type DeleteDatasetResponseBody struct {
	Data      *DeleteDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponseBody) SetData(v *DeleteDatasetResponseBodyData) *DeleteDatasetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDatasetResponseBody) SetRequestId(v string) *DeleteDatasetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteDatasetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponseBodyData) SetDescription(v string) *DeleteDatasetResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteDatasetResponseBodyData) SetGmtCreate(v int64) *DeleteDatasetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteDatasetResponseBodyData) SetId(v int64) *DeleteDatasetResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteDatasetResponseBodyData) SetName(v string) *DeleteDatasetResponseBodyData {
	s.Name = &v
	return s
}

type DeleteDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponse) SetHeaders(v map[string]*string) *DeleteDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetResponse) SetStatusCode(v int32) *DeleteDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetResponse) SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse {
	s.Body = v
	return s
}

type DeleteLabelsetRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteLabelsetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetRequest) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetRequest) SetId(v int64) *DeleteLabelsetRequest {
	s.Id = &v
	return s
}

type DeleteLabelsetResponseBody struct {
	Data      *DeleteLabelsetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLabelsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetResponseBody) SetData(v *DeleteLabelsetResponseBodyData) *DeleteLabelsetResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLabelsetResponseBody) SetRequestId(v string) *DeleteLabelsetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLabelsetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelType   *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteLabelsetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetResponseBodyData) SetDescription(v string) *DeleteLabelsetResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteLabelsetResponseBodyData) SetGmtCreate(v int64) *DeleteLabelsetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteLabelsetResponseBodyData) SetId(v int64) *DeleteLabelsetResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteLabelsetResponseBodyData) SetLabelType(v string) *DeleteLabelsetResponseBodyData {
	s.LabelType = &v
	return s
}

func (s *DeleteLabelsetResponseBodyData) SetName(v string) *DeleteLabelsetResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteLabelsetResponseBodyData) SetStatus(v string) *DeleteLabelsetResponseBodyData {
	s.Status = &v
	return s
}

type DeleteLabelsetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLabelsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLabelsetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetResponse) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetResponse) SetHeaders(v map[string]*string) *DeleteLabelsetResponse {
	s.Headers = v
	return s
}

func (s *DeleteLabelsetResponse) SetStatusCode(v int32) *DeleteLabelsetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLabelsetResponse) SetBody(v *DeleteLabelsetResponseBody) *DeleteLabelsetResponse {
	s.Body = v
	return s
}

type DeleteLabelsetDataRequest struct {
	Id      *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId *int64 `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
}

func (s DeleteLabelsetDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetDataRequest) SetId(v int64) *DeleteLabelsetDataRequest {
	s.Id = &v
	return s
}

func (s *DeleteLabelsetDataRequest) SetLabelId(v int64) *DeleteLabelsetDataRequest {
	s.LabelId = &v
	return s
}

type DeleteLabelsetDataResponseBody struct {
	Data      *DeleteLabelsetDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLabelsetDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetDataResponseBody) SetData(v *DeleteLabelsetDataResponseBodyData) *DeleteLabelsetDataResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLabelsetDataResponseBody) SetRequestId(v string) *DeleteLabelsetDataResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLabelsetDataResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelType   *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Total       *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DeleteLabelsetDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetDataResponseBodyData) SetDescription(v string) *DeleteLabelsetDataResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteLabelsetDataResponseBodyData) SetGmtCreate(v int64) *DeleteLabelsetDataResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteLabelsetDataResponseBodyData) SetId(v int64) *DeleteLabelsetDataResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteLabelsetDataResponseBodyData) SetLabelType(v string) *DeleteLabelsetDataResponseBodyData {
	s.LabelType = &v
	return s
}

func (s *DeleteLabelsetDataResponseBodyData) SetName(v string) *DeleteLabelsetDataResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteLabelsetDataResponseBodyData) SetStatus(v string) *DeleteLabelsetDataResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteLabelsetDataResponseBodyData) SetTotal(v int64) *DeleteLabelsetDataResponseBodyData {
	s.Total = &v
	return s
}

type DeleteLabelsetDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLabelsetDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLabelsetDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLabelsetDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteLabelsetDataResponse) SetHeaders(v map[string]*string) *DeleteLabelsetDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteLabelsetDataResponse) SetStatusCode(v int32) *DeleteLabelsetDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLabelsetDataResponse) SetBody(v *DeleteLabelsetDataResponseBody) *DeleteLabelsetDataResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetId(v int64) *DeleteServiceRequest {
	s.Id = &v
	return s
}

type DeleteServiceResponseBody struct {
	Data      *DeleteServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetData(v *DeleteServiceResponseBodyData) *DeleteServiceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponseBodyData struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBodyData) SetId(v int64) *DeleteServiceResponseBodyData {
	s.Id = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteTrainTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrainTaskRequest) SetId(v int64) *DeleteTrainTaskRequest {
	s.Id = &v
	return s
}

type DeleteTrainTaskResponseBody struct {
	Data      *DeleteTrainTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrainTaskResponseBody) SetData(v *DeleteTrainTaskResponseBodyData) *DeleteTrainTaskResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTrainTaskResponseBody) SetRequestId(v string) *DeleteTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrainTaskResponseBodyData struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName   *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	ModelEffect *string `json:"ModelEffect,omitempty" xml:"ModelEffect,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TrainMode   *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	TrainStatus *string `json:"TrainStatus,omitempty" xml:"TrainStatus,omitempty"`
}

func (s DeleteTrainTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteTrainTaskResponseBodyData) SetDatasetId(v int64) *DeleteTrainTaskResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetDatasetName(v string) *DeleteTrainTaskResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetDescription(v string) *DeleteTrainTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetGmtCreate(v int64) *DeleteTrainTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetId(v int64) *DeleteTrainTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetLabelId(v int64) *DeleteTrainTaskResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetLabelName(v string) *DeleteTrainTaskResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetModelEffect(v string) *DeleteTrainTaskResponseBodyData {
	s.ModelEffect = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetModelId(v int64) *DeleteTrainTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetTaskName(v string) *DeleteTrainTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetTrainMode(v string) *DeleteTrainTaskResponseBodyData {
	s.TrainMode = &v
	return s
}

func (s *DeleteTrainTaskResponseBodyData) SetTrainStatus(v string) *DeleteTrainTaskResponseBodyData {
	s.TrainStatus = &v
	return s
}

type DeleteTrainTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrainTaskResponse) SetHeaders(v map[string]*string) *DeleteTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrainTaskResponse) SetStatusCode(v int32) *DeleteTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrainTaskResponse) SetBody(v *DeleteTrainTaskResponseBody) *DeleteTrainTaskResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRequest) SetId(v int64) *DeleteWorkspaceRequest {
	s.Id = &v
	return s
}

type DeleteWorkspaceResponseBody struct {
	Data      *DeleteWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) SetData(v *DeleteWorkspaceResponseBodyData) *DeleteWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkspaceResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteWorkspaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBodyData) SetDescription(v string) *DeleteWorkspaceResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyData) SetGmtCreate(v int64) *DeleteWorkspaceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyData) SetId(v int64) *DeleteWorkspaceResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyData) SetName(v string) *DeleteWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyData) SetType(v string) *DeleteWorkspaceResponseBodyData {
	s.Type = &v
	return s
}

type DeleteWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResponse) SetStatusCode(v int32) *DeleteWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponse) SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse {
	s.Body = v
	return s
}

type DownloadFileNameListRequest struct {
	DatasetId *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Identity  *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
}

func (s DownloadFileNameListRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadFileNameListRequest) GoString() string {
	return s.String()
}

func (s *DownloadFileNameListRequest) SetDatasetId(v int64) *DownloadFileNameListRequest {
	s.DatasetId = &v
	return s
}

func (s *DownloadFileNameListRequest) SetIdentity(v string) *DownloadFileNameListRequest {
	s.Identity = &v
	return s
}

type DownloadFileNameListResponseBody struct {
	Data      *DownloadFileNameListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadFileNameListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadFileNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadFileNameListResponseBody) SetData(v *DownloadFileNameListResponseBodyData) *DownloadFileNameListResponseBody {
	s.Data = v
	return s
}

func (s *DownloadFileNameListResponseBody) SetRequestId(v string) *DownloadFileNameListResponseBody {
	s.RequestId = &v
	return s
}

type DownloadFileNameListResponseBodyData struct {
	OssHttpUrl *string `json:"OssHttpUrl,omitempty" xml:"OssHttpUrl,omitempty"`
}

func (s DownloadFileNameListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DownloadFileNameListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DownloadFileNameListResponseBodyData) SetOssHttpUrl(v string) *DownloadFileNameListResponseBodyData {
	s.OssHttpUrl = &v
	return s
}

type DownloadFileNameListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadFileNameListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadFileNameListResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadFileNameListResponse) GoString() string {
	return s.String()
}

func (s *DownloadFileNameListResponse) SetHeaders(v map[string]*string) *DownloadFileNameListResponse {
	s.Headers = v
	return s
}

func (s *DownloadFileNameListResponse) SetStatusCode(v int32) *DownloadFileNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadFileNameListResponse) SetBody(v *DownloadFileNameListResponseBody) *DownloadFileNameListResponse {
	s.Body = v
	return s
}

type DownloadLabelFileRequest struct {
	LabelId *int64 `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
}

func (s DownloadLabelFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadLabelFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadLabelFileRequest) SetLabelId(v int64) *DownloadLabelFileRequest {
	s.LabelId = &v
	return s
}

type DownloadLabelFileResponseBody struct {
	Data      *DownloadLabelFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadLabelFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadLabelFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadLabelFileResponseBody) SetData(v *DownloadLabelFileResponseBodyData) *DownloadLabelFileResponseBody {
	s.Data = v
	return s
}

func (s *DownloadLabelFileResponseBody) SetRequestId(v string) *DownloadLabelFileResponseBody {
	s.RequestId = &v
	return s
}

type DownloadLabelFileResponseBodyData struct {
	OssHttpUrl *string `json:"OssHttpUrl,omitempty" xml:"OssHttpUrl,omitempty"`
}

func (s DownloadLabelFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DownloadLabelFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DownloadLabelFileResponseBodyData) SetOssHttpUrl(v string) *DownloadLabelFileResponseBodyData {
	s.OssHttpUrl = &v
	return s
}

type DownloadLabelFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadLabelFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadLabelFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadLabelFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadLabelFileResponse) SetHeaders(v map[string]*string) *DownloadLabelFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadLabelFileResponse) SetStatusCode(v int32) *DownloadLabelFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadLabelFileResponse) SetBody(v *DownloadLabelFileResponseBody) *DownloadLabelFileResponse {
	s.Body = v
	return s
}

type DownloadTemplateResponseBody struct {
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadTemplateResponseBody) SetData(v map[string]interface{}) *DownloadTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DownloadTemplateResponseBody) SetRequestId(v string) *DownloadTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DownloadTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadTemplateResponse) GoString() string {
	return s.String()
}

func (s *DownloadTemplateResponse) SetHeaders(v map[string]*string) *DownloadTemplateResponse {
	s.Headers = v
	return s
}

func (s *DownloadTemplateResponse) SetStatusCode(v int32) *DownloadTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadTemplateResponse) SetBody(v *DownloadTemplateResponseBody) *DownloadTemplateResponse {
	s.Body = v
	return s
}

type GetDatasetRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) SetId(v int64) *GetDatasetRequest {
	s.Id = &v
	return s
}

type GetDatasetResponseBody struct {
	Data      *GetDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) SetData(v *GetDatasetResponseBodyData) *GetDatasetResponseBody {
	s.Data = v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

type GetDatasetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssUrl      *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	OwnerType   *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Total       *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetDatasetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBodyData) SetDescription(v string) *GetDatasetResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetGmtCreate(v int64) *GetDatasetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetId(v int64) *GetDatasetResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetName(v string) *GetDatasetResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetOssUrl(v string) *GetDatasetResponseBodyData {
	s.OssUrl = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetOwnerType(v string) *GetDatasetResponseBodyData {
	s.OwnerType = &v
	return s
}

func (s *GetDatasetResponseBodyData) SetTotal(v int64) *GetDatasetResponseBodyData {
	s.Total = &v
	return s
}

type GetDatasetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetResponse) SetHeaders(v map[string]*string) *GetDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetResponse) SetStatusCode(v int32) *GetDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

type GetDiffCountLabelsetAndDatasetRequest struct {
	LabelsetId *int64 `json:"LabelsetId,omitempty" xml:"LabelsetId,omitempty"`
}

func (s GetDiffCountLabelsetAndDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiffCountLabelsetAndDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDiffCountLabelsetAndDatasetRequest) SetLabelsetId(v int64) *GetDiffCountLabelsetAndDatasetRequest {
	s.LabelsetId = &v
	return s
}

type GetDiffCountLabelsetAndDatasetResponseBody struct {
	Data      *GetDiffCountLabelsetAndDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiffCountLabelsetAndDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiffCountLabelsetAndDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiffCountLabelsetAndDatasetResponseBody) SetData(v *GetDiffCountLabelsetAndDatasetResponseBodyData) *GetDiffCountLabelsetAndDatasetResponseBody {
	s.Data = v
	return s
}

func (s *GetDiffCountLabelsetAndDatasetResponseBody) SetRequestId(v string) *GetDiffCountLabelsetAndDatasetResponseBody {
	s.RequestId = &v
	return s
}

type GetDiffCountLabelsetAndDatasetResponseBodyData struct {
	DiffCount *int64 `json:"DiffCount,omitempty" xml:"DiffCount,omitempty"`
}

func (s GetDiffCountLabelsetAndDatasetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDiffCountLabelsetAndDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDiffCountLabelsetAndDatasetResponseBodyData) SetDiffCount(v int64) *GetDiffCountLabelsetAndDatasetResponseBodyData {
	s.DiffCount = &v
	return s
}

type GetDiffCountLabelsetAndDatasetResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiffCountLabelsetAndDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiffCountLabelsetAndDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiffCountLabelsetAndDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetDiffCountLabelsetAndDatasetResponse) SetHeaders(v map[string]*string) *GetDiffCountLabelsetAndDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetDiffCountLabelsetAndDatasetResponse) SetStatusCode(v int32) *GetDiffCountLabelsetAndDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiffCountLabelsetAndDatasetResponse) SetBody(v *GetDiffCountLabelsetAndDatasetResponseBody) *GetDiffCountLabelsetAndDatasetResponse {
	s.Body = v
	return s
}

type GetLabelDetailRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetLabelDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLabelDetailRequest) GoString() string {
	return s.String()
}

func (s *GetLabelDetailRequest) SetId(v int64) *GetLabelDetailRequest {
	s.Id = &v
	return s
}

type GetLabelDetailResponseBody struct {
	Data      *GetLabelDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLabelDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLabelDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelDetailResponseBody) SetData(v *GetLabelDetailResponseBodyData) *GetLabelDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetLabelDetailResponseBody) SetRequestId(v string) *GetLabelDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetLabelDetailResponseBodyData struct {
	LabelInfo *string `json:"LabelInfo,omitempty" xml:"LabelInfo,omitempty"`
}

func (s GetLabelDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLabelDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLabelDetailResponseBodyData) SetLabelInfo(v string) *GetLabelDetailResponseBodyData {
	s.LabelInfo = &v
	return s
}

type GetLabelDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLabelDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLabelDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLabelDetailResponse) GoString() string {
	return s.String()
}

func (s *GetLabelDetailResponse) SetHeaders(v map[string]*string) *GetLabelDetailResponse {
	s.Headers = v
	return s
}

func (s *GetLabelDetailResponse) SetStatusCode(v int32) *GetLabelDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelDetailResponse) SetBody(v *GetLabelDetailResponseBody) *GetLabelDetailResponse {
	s.Body = v
	return s
}

type GetLabelsetRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetLabelsetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLabelsetRequest) GoString() string {
	return s.String()
}

func (s *GetLabelsetRequest) SetId(v int64) *GetLabelsetRequest {
	s.Id = &v
	return s
}

type GetLabelsetResponseBody struct {
	Data      *GetLabelsetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLabelsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLabelsetResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelsetResponseBody) SetData(v *GetLabelsetResponseBodyData) *GetLabelsetResponseBody {
	s.Data = v
	return s
}

func (s *GetLabelsetResponseBody) SetRequestId(v string) *GetLabelsetResponseBody {
	s.RequestId = &v
	return s
}

type GetLabelsetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelType   *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Total       *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLabelsetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLabelsetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLabelsetResponseBodyData) SetDescription(v string) *GetLabelsetResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetLabelsetResponseBodyData) SetGmtCreate(v int64) *GetLabelsetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetLabelsetResponseBodyData) SetId(v int64) *GetLabelsetResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetLabelsetResponseBodyData) SetLabelType(v string) *GetLabelsetResponseBodyData {
	s.LabelType = &v
	return s
}

func (s *GetLabelsetResponseBodyData) SetName(v string) *GetLabelsetResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetLabelsetResponseBodyData) SetStatus(v string) *GetLabelsetResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetLabelsetResponseBodyData) SetTotal(v int64) *GetLabelsetResponseBodyData {
	s.Total = &v
	return s
}

type GetLabelsetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLabelsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLabelsetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLabelsetResponse) GoString() string {
	return s.String()
}

func (s *GetLabelsetResponse) SetHeaders(v map[string]*string) *GetLabelsetResponse {
	s.Headers = v
	return s
}

func (s *GetLabelsetResponse) SetStatusCode(v int32) *GetLabelsetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelsetResponse) SetBody(v *GetLabelsetResponseBody) *GetLabelsetResponse {
	s.Body = v
	return s
}

type GetServiceRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetId(v int64) *GetServiceRequest {
	s.Id = &v
	return s
}

type GetServiceResponseBody struct {
	Data      *GetServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetData(v *GetServiceResponseBodyData) *GetServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceResponseBodyData struct {
	CurlExample        *string `json:"CurlExample,omitempty" xml:"CurlExample,omitempty"`
	Errorcodes         *string `json:"Errorcodes,omitempty" xml:"Errorcodes,omitempty"`
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InputExample       *string `json:"InputExample,omitempty" xml:"InputExample,omitempty"`
	InputParams        *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	OutputExample      *string `json:"OutputExample,omitempty" xml:"OutputExample,omitempty"`
	OutputParams       *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceName        *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyData) SetCurlExample(v string) *GetServiceResponseBodyData {
	s.CurlExample = &v
	return s
}

func (s *GetServiceResponseBodyData) SetErrorcodes(v string) *GetServiceResponseBodyData {
	s.Errorcodes = &v
	return s
}

func (s *GetServiceResponseBodyData) SetGmtCreate(v int64) *GetServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceResponseBodyData) SetId(v int64) *GetServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetServiceResponseBodyData) SetInputExample(v string) *GetServiceResponseBodyData {
	s.InputExample = &v
	return s
}

func (s *GetServiceResponseBodyData) SetInputParams(v string) *GetServiceResponseBodyData {
	s.InputParams = &v
	return s
}

func (s *GetServiceResponseBodyData) SetOutputExample(v string) *GetServiceResponseBodyData {
	s.OutputExample = &v
	return s
}

func (s *GetServiceResponseBodyData) SetOutputParams(v string) *GetServiceResponseBodyData {
	s.OutputParams = &v
	return s
}

func (s *GetServiceResponseBodyData) SetServiceDescription(v string) *GetServiceResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *GetServiceResponseBodyData) SetServiceName(v string) *GetServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetServiceResponseBodyData) SetStatus(v string) *GetServiceResponseBodyData {
	s.Status = &v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetTrainModelRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTrainModelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrainModelRequest) GoString() string {
	return s.String()
}

func (s *GetTrainModelRequest) SetId(v int64) *GetTrainModelRequest {
	s.Id = &v
	return s
}

type GetTrainModelResponseBody struct {
	Data      *GetTrainModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainModelResponseBody) SetData(v *GetTrainModelResponseBodyData) *GetTrainModelResponseBody {
	s.Data = v
	return s
}

func (s *GetTrainModelResponseBody) SetRequestId(v string) *GetTrainModelResponseBody {
	s.RequestId = &v
	return s
}

type GetTrainModelResponseBodyData struct {
	SimpleEvaluate *int64 `json:"SimpleEvaluate,omitempty" xml:"SimpleEvaluate,omitempty"`
}

func (s GetTrainModelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrainModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrainModelResponseBodyData) SetSimpleEvaluate(v int64) *GetTrainModelResponseBodyData {
	s.SimpleEvaluate = &v
	return s
}

type GetTrainModelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrainModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrainModelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainModelResponse) GoString() string {
	return s.String()
}

func (s *GetTrainModelResponse) SetHeaders(v map[string]*string) *GetTrainModelResponse {
	s.Headers = v
	return s
}

func (s *GetTrainModelResponse) SetStatusCode(v int32) *GetTrainModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainModelResponse) SetBody(v *GetTrainModelResponseBody) *GetTrainModelResponse {
	s.Body = v
	return s
}

type GetTrainTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTrainTaskRequest) SetId(v int64) *GetTrainTaskRequest {
	s.Id = &v
	return s
}

type GetTrainTaskResponseBody struct {
	Data      *GetTrainTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainTaskResponseBody) SetData(v *GetTrainTaskResponseBodyData) *GetTrainTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTrainTaskResponseBody) SetRequestId(v string) *GetTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetTrainTaskResponseBodyData struct {
	DatasetId     *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName   *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId       *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName     *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	ModelEffect   *string `json:"ModelEffect,omitempty" xml:"ModelEffect,omitempty"`
	ModelId       *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TrainMode     *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	TrainStatus   *string `json:"TrainStatus,omitempty" xml:"TrainStatus,omitempty"`
	TrainUseTime  *int64  `json:"TrainUseTime,omitempty" xml:"TrainUseTime,omitempty"`
}

func (s GetTrainTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrainTaskResponseBodyData) SetDatasetId(v int64) *GetTrainTaskResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetDatasetName(v string) *GetTrainTaskResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetDescription(v string) *GetTrainTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetFailureReason(v string) *GetTrainTaskResponseBodyData {
	s.FailureReason = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetGmtCreate(v int64) *GetTrainTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetId(v int64) *GetTrainTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetLabelId(v int64) *GetTrainTaskResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetLabelName(v string) *GetTrainTaskResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetModelEffect(v string) *GetTrainTaskResponseBodyData {
	s.ModelEffect = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetModelId(v int64) *GetTrainTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetTaskName(v string) *GetTrainTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetTrainMode(v string) *GetTrainTaskResponseBodyData {
	s.TrainMode = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetTrainStatus(v string) *GetTrainTaskResponseBodyData {
	s.TrainStatus = &v
	return s
}

func (s *GetTrainTaskResponseBodyData) SetTrainUseTime(v int64) *GetTrainTaskResponseBodyData {
	s.TrainUseTime = &v
	return s
}

type GetTrainTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTrainTaskResponse) SetHeaders(v map[string]*string) *GetTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTrainTaskResponse) SetStatusCode(v int32) *GetTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainTaskResponse) SetBody(v *GetTrainTaskResponseBody) *GetTrainTaskResponse {
	s.Body = v
	return s
}

type GetTrainTaskEstimatedTimeRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTrainTaskEstimatedTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskEstimatedTimeRequest) GoString() string {
	return s.String()
}

func (s *GetTrainTaskEstimatedTimeRequest) SetId(v int64) *GetTrainTaskEstimatedTimeRequest {
	s.Id = &v
	return s
}

type GetTrainTaskEstimatedTimeResponseBody struct {
	Data      *GetTrainTaskEstimatedTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainTaskEstimatedTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskEstimatedTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainTaskEstimatedTimeResponseBody) SetData(v *GetTrainTaskEstimatedTimeResponseBodyData) *GetTrainTaskEstimatedTimeResponseBody {
	s.Data = v
	return s
}

func (s *GetTrainTaskEstimatedTimeResponseBody) SetRequestId(v string) *GetTrainTaskEstimatedTimeResponseBody {
	s.RequestId = &v
	return s
}

type GetTrainTaskEstimatedTimeResponseBodyData struct {
	EstimatedTime *string `json:"EstimatedTime,omitempty" xml:"EstimatedTime,omitempty"`
}

func (s GetTrainTaskEstimatedTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskEstimatedTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrainTaskEstimatedTimeResponseBodyData) SetEstimatedTime(v string) *GetTrainTaskEstimatedTimeResponseBodyData {
	s.EstimatedTime = &v
	return s
}

type GetTrainTaskEstimatedTimeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrainTaskEstimatedTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrainTaskEstimatedTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrainTaskEstimatedTimeResponse) GoString() string {
	return s.String()
}

func (s *GetTrainTaskEstimatedTimeResponse) SetHeaders(v map[string]*string) *GetTrainTaskEstimatedTimeResponse {
	s.Headers = v
	return s
}

func (s *GetTrainTaskEstimatedTimeResponse) SetStatusCode(v int32) *GetTrainTaskEstimatedTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainTaskEstimatedTimeResponse) SetBody(v *GetTrainTaskEstimatedTimeResponseBody) *GetTrainTaskEstimatedTimeResponse {
	s.Body = v
	return s
}

type GetUploadPolicyRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUploadPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyRequest) SetFileName(v string) *GetUploadPolicyRequest {
	s.FileName = &v
	return s
}

func (s *GetUploadPolicyRequest) SetId(v int64) *GetUploadPolicyRequest {
	s.Id = &v
	return s
}

func (s *GetUploadPolicyRequest) SetType(v string) *GetUploadPolicyRequest {
	s.Type = &v
	return s
}

type GetUploadPolicyResponseBody struct {
	Data      *GetUploadPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponseBody) SetData(v *GetUploadPolicyResponseBodyData) *GetUploadPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadPolicyResponseBody) SetRequestId(v string) *GetUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetUploadPolicyResponseBodyData struct {
	AccessId         *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	BucketName       *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Endpoint         *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ObjectKey        *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	OriginalFilename *string `json:"OriginalFilename,omitempty" xml:"OriginalFilename,omitempty"`
	Policy           *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature        *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	SignedHttpUrl    *string `json:"SignedHttpUrl,omitempty" xml:"SignedHttpUrl,omitempty"`
}

func (s GetUploadPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUploadPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponseBodyData) SetAccessId(v string) *GetUploadPolicyResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetBucketName(v string) *GetUploadPolicyResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetEndpoint(v string) *GetUploadPolicyResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetFileName(v string) *GetUploadPolicyResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetObjectKey(v string) *GetUploadPolicyResponseBodyData {
	s.ObjectKey = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetOriginalFilename(v string) *GetUploadPolicyResponseBodyData {
	s.OriginalFilename = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetPolicy(v string) *GetUploadPolicyResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetSignature(v string) *GetUploadPolicyResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetSignedHttpUrl(v string) *GetUploadPolicyResponseBodyData {
	s.SignedHttpUrl = &v
	return s
}

type GetUploadPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponse) SetHeaders(v map[string]*string) *GetUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetUploadPolicyResponse) SetStatusCode(v int32) *GetUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadPolicyResponse) SetBody(v *GetUploadPolicyResponseBody) *GetUploadPolicyResponse {
	s.Body = v
	return s
}

type GetWorkspaceRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) SetId(v int64) *GetWorkspaceRequest {
	s.Id = &v
	return s
}

type GetWorkspaceResponseBody struct {
	Data      *GetWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetData(v *GetWorkspaceResponseBodyData) *GetWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type GetWorkspaceResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWorkspaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyData) SetDescription(v string) *GetWorkspaceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetGmtCreate(v int64) *GetWorkspaceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetId(v int64) *GetWorkspaceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetName(v string) *GetWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetType(v string) *GetWorkspaceResponseBodyData {
	s.Type = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetStatusCode(v int32) *GetWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type ListDatasetDatasRequest struct {
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Identity    *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDatasetDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetDatasRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetDatasRequest) SetCurrentPage(v int64) *ListDatasetDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDatasetDatasRequest) SetDatasetId(v int64) *ListDatasetDatasRequest {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetDatasRequest) SetIdentity(v string) *ListDatasetDatasRequest {
	s.Identity = &v
	return s
}

func (s *ListDatasetDatasRequest) SetPageSize(v int64) *ListDatasetDatasRequest {
	s.PageSize = &v
	return s
}

type ListDatasetDatasResponseBody struct {
	Data      *ListDatasetDatasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatasetDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetDatasResponseBody) SetData(v *ListDatasetDatasResponseBodyData) *ListDatasetDatasResponseBody {
	s.Data = v
	return s
}

func (s *ListDatasetDatasResponseBody) SetRequestId(v string) *ListDatasetDatasResponseBody {
	s.RequestId = &v
	return s
}

type ListDatasetDatasResponseBodyData struct {
	CurrentPage *int64                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []map[string]interface{} `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListDatasetDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDatasetDatasResponseBodyData) SetCurrentPage(v int64) *ListDatasetDatasResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListDatasetDatasResponseBodyData) SetElements(v []map[string]interface{}) *ListDatasetDatasResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListDatasetDatasResponseBodyData) SetPageSize(v int64) *ListDatasetDatasResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDatasetDatasResponseBodyData) SetTotalCount(v int64) *ListDatasetDatasResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetDatasResponseBodyData) SetTotalPage(v int64) *ListDatasetDatasResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListDatasetDatasResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDatasetDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatasetDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetDatasResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetDatasResponse) SetHeaders(v map[string]*string) *ListDatasetDatasResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetDatasResponse) SetStatusCode(v int32) *ListDatasetDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetDatasResponse) SetBody(v *ListDatasetDatasResponseBody) *ListDatasetDatasResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) SetCurrentPage(v int64) *ListDatasetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDatasetsRequest) SetPageSize(v int64) *ListDatasetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsRequest) SetWorkspaceId(v int64) *ListDatasetsRequest {
	s.WorkspaceId = &v
	return s
}

type ListDatasetsResponseBody struct {
	Data      *ListDatasetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) SetData(v *ListDatasetsResponseBodyData) *ListDatasetsResponseBody {
	s.Data = v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

type ListDatasetsResponseBodyData struct {
	CurrentPage *int64                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []map[string]interface{} `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListDatasetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyData) SetCurrentPage(v int64) *ListDatasetsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetElements(v []map[string]interface{}) *ListDatasetsResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListDatasetsResponseBodyData) SetPageSize(v int64) *ListDatasetsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetTotalCount(v int64) *ListDatasetsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetTotalPage(v int64) *ListDatasetsResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListDatasetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatasetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponse) SetHeaders(v map[string]*string) *ListDatasetsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetsResponse) SetStatusCode(v int32) *ListDatasetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
	s.Body = v
	return s
}

type ListLabelsetDatasRequest struct {
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLabelsetDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetDatasRequest) GoString() string {
	return s.String()
}

func (s *ListLabelsetDatasRequest) SetCurrentPage(v int64) *ListLabelsetDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListLabelsetDatasRequest) SetLabelId(v int64) *ListLabelsetDatasRequest {
	s.LabelId = &v
	return s
}

func (s *ListLabelsetDatasRequest) SetName(v string) *ListLabelsetDatasRequest {
	s.Name = &v
	return s
}

func (s *ListLabelsetDatasRequest) SetOperation(v string) *ListLabelsetDatasRequest {
	s.Operation = &v
	return s
}

func (s *ListLabelsetDatasRequest) SetPageSize(v int64) *ListLabelsetDatasRequest {
	s.PageSize = &v
	return s
}

func (s *ListLabelsetDatasRequest) SetValue(v string) *ListLabelsetDatasRequest {
	s.Value = &v
	return s
}

type ListLabelsetDatasResponseBody struct {
	Data      *ListLabelsetDatasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLabelsetDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListLabelsetDatasResponseBody) SetData(v *ListLabelsetDatasResponseBodyData) *ListLabelsetDatasResponseBody {
	s.Data = v
	return s
}

func (s *ListLabelsetDatasResponseBody) SetRequestId(v string) *ListLabelsetDatasResponseBody {
	s.RequestId = &v
	return s
}

type ListLabelsetDatasResponseBodyData struct {
	CurrentPage *int64                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []map[string]interface{} `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListLabelsetDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLabelsetDatasResponseBodyData) SetCurrentPage(v int64) *ListLabelsetDatasResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListLabelsetDatasResponseBodyData) SetElements(v []map[string]interface{}) *ListLabelsetDatasResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListLabelsetDatasResponseBodyData) SetPageSize(v int64) *ListLabelsetDatasResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLabelsetDatasResponseBodyData) SetTotalCount(v int64) *ListLabelsetDatasResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLabelsetDatasResponseBodyData) SetTotalPage(v int64) *ListLabelsetDatasResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListLabelsetDatasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLabelsetDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLabelsetDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetDatasResponse) GoString() string {
	return s.String()
}

func (s *ListLabelsetDatasResponse) SetHeaders(v map[string]*string) *ListLabelsetDatasResponse {
	s.Headers = v
	return s
}

func (s *ListLabelsetDatasResponse) SetStatusCode(v int32) *ListLabelsetDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLabelsetDatasResponse) SetBody(v *ListLabelsetDatasResponseBody) *ListLabelsetDatasResponse {
	s.Body = v
	return s
}

type ListLabelsetsRequest struct {
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DatasetId   *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	PageSize    *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLabelsetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetsRequest) GoString() string {
	return s.String()
}

func (s *ListLabelsetsRequest) SetCurrentPage(v int64) *ListLabelsetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListLabelsetsRequest) SetDatasetId(v int64) *ListLabelsetsRequest {
	s.DatasetId = &v
	return s
}

func (s *ListLabelsetsRequest) SetPageSize(v int64) *ListLabelsetsRequest {
	s.PageSize = &v
	return s
}

type ListLabelsetsResponseBody struct {
	Data      *ListLabelsetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLabelsetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLabelsetsResponseBody) SetData(v *ListLabelsetsResponseBodyData) *ListLabelsetsResponseBody {
	s.Data = v
	return s
}

func (s *ListLabelsetsResponseBody) SetRequestId(v string) *ListLabelsetsResponseBody {
	s.RequestId = &v
	return s
}

type ListLabelsetsResponseBodyData struct {
	CurrentPage *int64                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []map[string]interface{} `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListLabelsetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLabelsetsResponseBodyData) SetCurrentPage(v int64) *ListLabelsetsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListLabelsetsResponseBodyData) SetElements(v []map[string]interface{}) *ListLabelsetsResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListLabelsetsResponseBodyData) SetPageSize(v int64) *ListLabelsetsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLabelsetsResponseBodyData) SetTotalCount(v int64) *ListLabelsetsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLabelsetsResponseBodyData) SetTotalPage(v int64) *ListLabelsetsResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListLabelsetsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLabelsetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLabelsetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLabelsetsResponse) GoString() string {
	return s.String()
}

func (s *ListLabelsetsResponse) SetHeaders(v map[string]*string) *ListLabelsetsResponse {
	s.Headers = v
	return s
}

func (s *ListLabelsetsResponse) SetStatusCode(v int32) *ListLabelsetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLabelsetsResponse) SetBody(v *ListLabelsetsResponseBody) *ListLabelsetsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *int64 `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize    *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetCurrentPage(v int64) *ListServicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListServicesRequest) SetId(v int64) *ListServicesRequest {
	s.Id = &v
	return s
}

func (s *ListServicesRequest) SetName(v int64) *ListServicesRequest {
	s.Name = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int64) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetWorkspaceId(v int64) *ListServicesRequest {
	s.WorkspaceId = &v
	return s
}

type ListServicesResponseBody struct {
	Data      *ListServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetData(v *ListServicesResponseBodyData) *ListServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

type ListServicesResponseBodyData struct {
	CurrentPage *int64                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []map[string]interface{} `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyData) SetCurrentPage(v int64) *ListServicesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListServicesResponseBodyData) SetElements(v []map[string]interface{}) *ListServicesResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListServicesResponseBodyData) SetPageSize(v int64) *ListServicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBodyData) SetTotalCount(v int64) *ListServicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListServicesResponseBodyData) SetTotalPage(v int64) *ListServicesResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListTrainTasksRequest struct {
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTrainTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrainTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTrainTasksRequest) SetCurrentPage(v int64) *ListTrainTasksRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTrainTasksRequest) SetPageSize(v int64) *ListTrainTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrainTasksRequest) SetWorkspaceId(v int64) *ListTrainTasksRequest {
	s.WorkspaceId = &v
	return s
}

type ListTrainTasksResponseBody struct {
	Data      *ListTrainTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTrainTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrainTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainTasksResponseBody) SetData(v *ListTrainTasksResponseBodyData) *ListTrainTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTrainTasksResponseBody) SetRequestId(v string) *ListTrainTasksResponseBody {
	s.RequestId = &v
	return s
}

type ListTrainTasksResponseBodyData struct {
	CurrentPage *int64                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []map[string]interface{} `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                   `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTrainTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTrainTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTrainTasksResponseBodyData) SetCurrentPage(v int64) *ListTrainTasksResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListTrainTasksResponseBodyData) SetElements(v []map[string]interface{}) *ListTrainTasksResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListTrainTasksResponseBodyData) SetPageSize(v int64) *ListTrainTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTrainTasksResponseBodyData) SetTotalCount(v int64) *ListTrainTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTrainTasksResponseBodyData) SetTotalPage(v int64) *ListTrainTasksResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListTrainTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTrainTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTrainTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrainTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTrainTasksResponse) SetHeaders(v map[string]*string) *ListTrainTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTrainTasksResponse) SetStatusCode(v int32) *ListTrainTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainTasksResponse) SetBody(v *ListTrainTasksResponseBody) *ListTrainTasksResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetCurrentPage(v int64) *ListWorkspacesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListWorkspacesRequest) SetName(v string) *ListWorkspacesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int64) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

type ListWorkspacesResponseBody struct {
	Data      *ListWorkspacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetData(v *ListWorkspacesResponseBodyData) *ListWorkspacesResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

type ListWorkspacesResponseBodyData struct {
	CurrentPage *int64                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Elements    []*ListWorkspacesResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	PageSize    *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage   *int64                                    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListWorkspacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyData) SetCurrentPage(v int64) *ListWorkspacesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListWorkspacesResponseBodyData) SetElements(v []*ListWorkspacesResponseBodyDataElements) *ListWorkspacesResponseBodyData {
	s.Elements = v
	return s
}

func (s *ListWorkspacesResponseBodyData) SetPageSize(v int64) *ListWorkspacesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesResponseBodyData) SetTotalCount(v int64) *ListWorkspacesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBodyData) SetTotalPage(v int64) *ListWorkspacesResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListWorkspacesResponseBodyDataElements struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkspacesResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyDataElements) SetDescription(v string) *ListWorkspacesResponseBodyDataElements {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataElements) SetGmtCreate(v int64) *ListWorkspacesResponseBodyDataElements {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataElements) SetId(v int64) *ListWorkspacesResponseBodyDataElements {
	s.Id = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataElements) SetName(v string) *ListWorkspacesResponseBodyDataElements {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataElements) SetType(v string) *ListWorkspacesResponseBodyDataElements {
	s.Type = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type SetDatasetUserOssPathRequest struct {
	DatasetId  *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	UserOssUrl *string `json:"UserOssUrl,omitempty" xml:"UserOssUrl,omitempty"`
}

func (s SetDatasetUserOssPathRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDatasetUserOssPathRequest) GoString() string {
	return s.String()
}

func (s *SetDatasetUserOssPathRequest) SetDatasetId(v int64) *SetDatasetUserOssPathRequest {
	s.DatasetId = &v
	return s
}

func (s *SetDatasetUserOssPathRequest) SetUserOssUrl(v string) *SetDatasetUserOssPathRequest {
	s.UserOssUrl = &v
	return s
}

type SetDatasetUserOssPathResponseBody struct {
	Data      *SetDatasetUserOssPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDatasetUserOssPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDatasetUserOssPathResponseBody) GoString() string {
	return s.String()
}

func (s *SetDatasetUserOssPathResponseBody) SetData(v *SetDatasetUserOssPathResponseBodyData) *SetDatasetUserOssPathResponseBody {
	s.Data = v
	return s
}

func (s *SetDatasetUserOssPathResponseBody) SetRequestId(v string) *SetDatasetUserOssPathResponseBody {
	s.RequestId = &v
	return s
}

type SetDatasetUserOssPathResponseBodyData struct {
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s SetDatasetUserOssPathResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetDatasetUserOssPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetDatasetUserOssPathResponseBodyData) SetId(v int64) *SetDatasetUserOssPathResponseBodyData {
	s.Id = &v
	return s
}

func (s *SetDatasetUserOssPathResponseBodyData) SetOssUrl(v string) *SetDatasetUserOssPathResponseBodyData {
	s.OssUrl = &v
	return s
}

type SetDatasetUserOssPathResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDatasetUserOssPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDatasetUserOssPathResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDatasetUserOssPathResponse) GoString() string {
	return s.String()
}

func (s *SetDatasetUserOssPathResponse) SetHeaders(v map[string]*string) *SetDatasetUserOssPathResponse {
	s.Headers = v
	return s
}

func (s *SetDatasetUserOssPathResponse) SetStatusCode(v int32) *SetDatasetUserOssPathResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDatasetUserOssPathResponse) SetBody(v *SetDatasetUserOssPathResponseBody) *SetDatasetUserOssPathResponse {
	s.Body = v
	return s
}

type StartServiceRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartServiceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceRequest) SetId(v int64) *StartServiceRequest {
	s.Id = &v
	return s
}

type StartServiceResponseBody struct {
	Data      *StartServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceResponseBody) SetData(v *StartServiceResponseBodyData) *StartServiceResponseBody {
	s.Data = v
	return s
}

func (s *StartServiceResponseBody) SetRequestId(v string) *StartServiceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceResponseBodyData struct {
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceName        *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StartServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartServiceResponseBodyData) SetGmtCreate(v int64) *StartServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *StartServiceResponseBodyData) SetId(v int64) *StartServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *StartServiceResponseBodyData) SetServiceDescription(v string) *StartServiceResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *StartServiceResponseBodyData) SetServiceName(v string) *StartServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *StartServiceResponseBodyData) SetStatus(v string) *StartServiceResponseBodyData {
	s.Status = &v
	return s
}

type StartServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceResponse) SetHeaders(v map[string]*string) *StartServiceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceResponse) SetStatusCode(v int32) *StartServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceResponse) SetBody(v *StartServiceResponseBody) *StartServiceResponse {
	s.Body = v
	return s
}

type StartTrainTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTrainTaskRequest) SetId(v int64) *StartTrainTaskRequest {
	s.Id = &v
	return s
}

type StartTrainTaskResponseBody struct {
	Data      *StartTrainTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTrainTaskResponseBody) SetData(v *StartTrainTaskResponseBodyData) *StartTrainTaskResponseBody {
	s.Data = v
	return s
}

func (s *StartTrainTaskResponseBody) SetRequestId(v string) *StartTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartTrainTaskResponseBodyData struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName   *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	ModelEffect *string `json:"ModelEffect,omitempty" xml:"ModelEffect,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TrainMode   *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	TrainStatus *string `json:"TrainStatus,omitempty" xml:"TrainStatus,omitempty"`
}

func (s StartTrainTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartTrainTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartTrainTaskResponseBodyData) SetDatasetId(v int64) *StartTrainTaskResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetDatasetName(v string) *StartTrainTaskResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetDescription(v string) *StartTrainTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetGmtCreate(v int64) *StartTrainTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetId(v int64) *StartTrainTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetLabelId(v int64) *StartTrainTaskResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetLabelName(v string) *StartTrainTaskResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetModelEffect(v string) *StartTrainTaskResponseBodyData {
	s.ModelEffect = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetModelId(v int64) *StartTrainTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetTaskName(v string) *StartTrainTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetTrainMode(v string) *StartTrainTaskResponseBodyData {
	s.TrainMode = &v
	return s
}

func (s *StartTrainTaskResponseBodyData) SetTrainStatus(v string) *StartTrainTaskResponseBodyData {
	s.TrainStatus = &v
	return s
}

type StartTrainTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTrainTaskResponse) SetHeaders(v map[string]*string) *StartTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTrainTaskResponse) SetStatusCode(v int32) *StartTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTrainTaskResponse) SetBody(v *StartTrainTaskResponseBody) *StartTrainTaskResponse {
	s.Body = v
	return s
}

type StopServiceRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopServiceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceRequest) SetId(v int64) *StopServiceRequest {
	s.Id = &v
	return s
}

type StopServiceResponseBody struct {
	Data      *StopServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceResponseBody) SetData(v *StopServiceResponseBodyData) *StopServiceResponseBody {
	s.Data = v
	return s
}

func (s *StopServiceResponseBody) SetRequestId(v string) *StopServiceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceResponseBodyData struct {
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceName        *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StopServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopServiceResponseBodyData) SetGmtCreate(v int64) *StopServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *StopServiceResponseBodyData) SetId(v int64) *StopServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *StopServiceResponseBodyData) SetServiceDescription(v string) *StopServiceResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *StopServiceResponseBodyData) SetServiceName(v string) *StopServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *StopServiceResponseBodyData) SetStatus(v string) *StopServiceResponseBodyData {
	s.Status = &v
	return s
}

type StopServiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceResponse) SetHeaders(v map[string]*string) *StopServiceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceResponse) SetStatusCode(v int32) *StopServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceResponse) SetBody(v *StopServiceResponseBody) *StopServiceResponse {
	s.Body = v
	return s
}

type StopTrainTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTrainTaskRequest) SetId(v int64) *StopTrainTaskRequest {
	s.Id = &v
	return s
}

type StopTrainTaskResponseBody struct {
	Data      *StopTrainTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTrainTaskResponseBody) SetData(v *StopTrainTaskResponseBodyData) *StopTrainTaskResponseBody {
	s.Data = v
	return s
}

func (s *StopTrainTaskResponseBody) SetRequestId(v string) *StopTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopTrainTaskResponseBodyData struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName   *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	ModelEffect *string `json:"ModelEffect,omitempty" xml:"ModelEffect,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TrainMode   *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	TrainStatus *string `json:"TrainStatus,omitempty" xml:"TrainStatus,omitempty"`
}

func (s StopTrainTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopTrainTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopTrainTaskResponseBodyData) SetDatasetId(v int64) *StopTrainTaskResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetDatasetName(v string) *StopTrainTaskResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetDescription(v string) *StopTrainTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetGmtCreate(v int64) *StopTrainTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetId(v int64) *StopTrainTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetLabelId(v int64) *StopTrainTaskResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetLabelName(v string) *StopTrainTaskResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetModelEffect(v string) *StopTrainTaskResponseBodyData {
	s.ModelEffect = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetModelId(v int64) *StopTrainTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetTaskName(v string) *StopTrainTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetTrainMode(v string) *StopTrainTaskResponseBodyData {
	s.TrainMode = &v
	return s
}

func (s *StopTrainTaskResponseBodyData) SetTrainStatus(v string) *StopTrainTaskResponseBodyData {
	s.TrainStatus = &v
	return s
}

type StopTrainTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTrainTaskResponse) SetHeaders(v map[string]*string) *StopTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTrainTaskResponse) SetStatusCode(v int32) *StopTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrainTaskResponse) SetBody(v *StopTrainTaskResponseBody) *StopTrainTaskResponse {
	s.Body = v
	return s
}

type UpdateDatasetRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) SetId(v int64) *UpdateDatasetRequest {
	s.Id = &v
	return s
}

func (s *UpdateDatasetRequest) SetName(v string) *UpdateDatasetRequest {
	s.Name = &v
	return s
}

type UpdateDatasetResponseBody struct {
	Data      *UpdateDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) SetData(v *UpdateDatasetResponseBodyData) *UpdateDatasetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDatasetResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDatasetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBodyData) SetDescription(v string) *UpdateDatasetResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetGmtCreate(v int64) *UpdateDatasetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetId(v int64) *UpdateDatasetResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetName(v string) *UpdateDatasetResponseBodyData {
	s.Name = &v
	return s
}

type UpdateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponse) SetHeaders(v map[string]*string) *UpdateDatasetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetResponse) SetStatusCode(v int32) *UpdateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

type UpdateLabelsetRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectKey   *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	UserOssUrl  *string `json:"UserOssUrl,omitempty" xml:"UserOssUrl,omitempty"`
}

func (s UpdateLabelsetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelsetRequest) GoString() string {
	return s.String()
}

func (s *UpdateLabelsetRequest) SetDescription(v string) *UpdateLabelsetRequest {
	s.Description = &v
	return s
}

func (s *UpdateLabelsetRequest) SetId(v int64) *UpdateLabelsetRequest {
	s.Id = &v
	return s
}

func (s *UpdateLabelsetRequest) SetName(v string) *UpdateLabelsetRequest {
	s.Name = &v
	return s
}

func (s *UpdateLabelsetRequest) SetObjectKey(v string) *UpdateLabelsetRequest {
	s.ObjectKey = &v
	return s
}

func (s *UpdateLabelsetRequest) SetUserOssUrl(v string) *UpdateLabelsetRequest {
	s.UserOssUrl = &v
	return s
}

type UpdateLabelsetResponseBody struct {
	Data      *UpdateLabelsetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLabelsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelsetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLabelsetResponseBody) SetData(v *UpdateLabelsetResponseBodyData) *UpdateLabelsetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLabelsetResponseBody) SetRequestId(v string) *UpdateLabelsetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLabelsetResponseBodyData struct {
	Description *int64 `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *int64 `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateLabelsetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelsetResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLabelsetResponseBodyData) SetDescription(v int64) *UpdateLabelsetResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateLabelsetResponseBodyData) SetGmtCreate(v int64) *UpdateLabelsetResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateLabelsetResponseBodyData) SetId(v int64) *UpdateLabelsetResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateLabelsetResponseBodyData) SetName(v int64) *UpdateLabelsetResponseBodyData {
	s.Name = &v
	return s
}

type UpdateLabelsetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLabelsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLabelsetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLabelsetResponse) GoString() string {
	return s.String()
}

func (s *UpdateLabelsetResponse) SetHeaders(v map[string]*string) *UpdateLabelsetResponse {
	s.Headers = v
	return s
}

func (s *UpdateLabelsetResponse) SetStatusCode(v int32) *UpdateLabelsetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLabelsetResponse) SetBody(v *UpdateLabelsetResponseBody) *UpdateLabelsetResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetDescription(v string) *UpdateServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateServiceRequest) SetId(v int64) *UpdateServiceRequest {
	s.Id = &v
	return s
}

func (s *UpdateServiceRequest) SetName(v string) *UpdateServiceRequest {
	s.Name = &v
	return s
}

type UpdateServiceResponseBody struct {
	Data      *UpdateServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetData(v *UpdateServiceResponseBodyData) *UpdateServiceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponseBodyData struct {
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	ServiceName        *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s UpdateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyData) SetGmtCreate(v int64) *UpdateServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateServiceResponseBodyData) SetId(v int64) *UpdateServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateServiceResponseBodyData) SetServiceDescription(v string) *UpdateServiceResponseBodyData {
	s.ServiceDescription = &v
	return s
}

func (s *UpdateServiceResponseBodyData) SetServiceName(v string) *UpdateServiceResponseBodyData {
	s.ServiceName = &v
	return s
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateTrainTaskRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTrainTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrainTaskRequest) SetDescription(v string) *UpdateTrainTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTrainTaskRequest) SetId(v int64) *UpdateTrainTaskRequest {
	s.Id = &v
	return s
}

func (s *UpdateTrainTaskRequest) SetName(v string) *UpdateTrainTaskRequest {
	s.Name = &v
	return s
}

type UpdateTrainTaskResponseBody struct {
	Data      *UpdateTrainTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrainTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrainTaskResponseBody) SetData(v *UpdateTrainTaskResponseBodyData) *UpdateTrainTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTrainTaskResponseBody) SetRequestId(v string) *UpdateTrainTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrainTaskResponseBodyData struct {
	DatasetId   *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *int64  `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LabelName   *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	ModelEffect *string `json:"ModelEffect,omitempty" xml:"ModelEffect,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TrainMode   *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
	TrainStatus *string `json:"TrainStatus,omitempty" xml:"TrainStatus,omitempty"`
}

func (s UpdateTrainTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTrainTaskResponseBodyData) SetDatasetId(v int64) *UpdateTrainTaskResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetDatasetName(v string) *UpdateTrainTaskResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetDescription(v string) *UpdateTrainTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetGmtCreate(v int64) *UpdateTrainTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetId(v int64) *UpdateTrainTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetLabelId(v int64) *UpdateTrainTaskResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetLabelName(v string) *UpdateTrainTaskResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetModelEffect(v string) *UpdateTrainTaskResponseBodyData {
	s.ModelEffect = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetModelId(v int64) *UpdateTrainTaskResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetTaskName(v string) *UpdateTrainTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetTrainMode(v string) *UpdateTrainTaskResponseBodyData {
	s.TrainMode = &v
	return s
}

func (s *UpdateTrainTaskResponseBodyData) SetTrainStatus(v string) *UpdateTrainTaskResponseBodyData {
	s.TrainStatus = &v
	return s
}

type UpdateTrainTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTrainTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrainTaskResponse) SetHeaders(v map[string]*string) *UpdateTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrainTaskResponse) SetStatusCode(v int32) *UpdateTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrainTaskResponse) SetBody(v *UpdateTrainTaskResponseBody) *UpdateTrainTaskResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetId(v int64) *UpdateWorkspaceRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetName(v string) *UpdateWorkspaceRequest {
	s.Name = &v
	return s
}

type UpdateWorkspaceResponseBody struct {
	Data      *UpdateWorkspaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) SetData(v *UpdateWorkspaceResponseBodyData) *UpdateWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkspaceResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateWorkspaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyData) SetDescription(v string) *UpdateWorkspaceResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetGmtCreate(v int64) *UpdateWorkspaceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetId(v int64) *UpdateWorkspaceResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetName(v string) *UpdateWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetType(v string) *UpdateWorkspaceResponseBodyData {
	s.Type = &v
	return s
}

type UpdateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResponse) SetStatusCode(v int32) *UpdateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResponse) SetBody(v *UpdateWorkspaceResponseBody) *UpdateWorkspaceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("viapi-regen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, runtime *util.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLabelsetWithOptions(request *CreateLabelsetRequest, runtime *util.RuntimeOptions) (_result *CreateLabelsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKey)) {
		body["ObjectKey"] = request.ObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagSettings)) {
		body["TagSettings"] = request.TagSettings
	}

	if !tea.BoolValue(util.IsUnset(request.TagUserList)) {
		body["TagUserList"] = request.TagUserList
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserOssUrl)) {
		body["UserOssUrl"] = request.UserOssUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLabelset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLabelsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLabelset(request *CreateLabelsetRequest) (_result *CreateLabelsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLabelsetResponse{}
	_body, _err := client.CreateLabelsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TrainTaskId)) {
		body["TrainTaskId"] = request.TrainTaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTagTaskWithOptions(request *CreateTagTaskRequest, runtime *util.RuntimeOptions) (_result *CreateTagTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelsetId)) {
		body["LabelsetId"] = request.LabelsetId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTagTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTagTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTagTask(request *CreateTagTaskRequest) (_result *CreateTagTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTagTaskResponse{}
	_body, _err := client.CreateTagTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrainTaskWithOptions(request *CreateTrainTaskRequest, runtime *util.RuntimeOptions) (_result *CreateTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TrainMode)) {
		body["TrainMode"] = request.TrainMode
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrainTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrainTask(request *CreateTrainTaskRequest) (_result *CreateTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrainTaskResponse{}
	_body, _err := client.CreateTrainTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspace"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeClassifyImageWithOptions(request *CustomizeClassifyImageRequest, runtime *util.RuntimeOptions) (_result *CustomizeClassifyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomizeClassifyImage"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeClassifyImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeClassifyImage(request *CustomizeClassifyImageRequest) (_result *CustomizeClassifyImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CustomizeClassifyImageResponse{}
	_body, _err := client.CustomizeClassifyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeClassifyImageAdvance(request *CustomizeClassifyImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *CustomizeClassifyImageResponse, _err error) {
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
		Product:  tea.String("viapi-regen"),
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
	customizeClassifyImageReq := &CustomizeClassifyImageRequest{}
	openapiutil.Convert(request, customizeClassifyImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		customizeClassifyImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	customizeClassifyImageResp, _err := client.CustomizeClassifyImageWithOptions(customizeClassifyImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = customizeClassifyImageResp
	return _result, _err
}

func (client *Client) CustomizeDetectImageWithOptions(request *CustomizeDetectImageRequest, runtime *util.RuntimeOptions) (_result *CustomizeDetectImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomizeDetectImage"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeDetectImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeDetectImage(request *CustomizeDetectImageRequest) (_result *CustomizeDetectImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CustomizeDetectImageResponse{}
	_body, _err := client.CustomizeDetectImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeDetectImageAdvance(request *CustomizeDetectImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *CustomizeDetectImageResponse, _err error) {
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
		Product:  tea.String("viapi-regen"),
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
	customizeDetectImageReq := &CustomizeDetectImageRequest{}
	openapiutil.Convert(request, customizeDetectImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		customizeDetectImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	customizeDetectImageResp, _err := client.CustomizeDetectImageWithOptions(customizeDetectImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = customizeDetectImageResp
	return _result, _err
}

func (client *Client) CustomizeInstanceSegmentImageWithOptions(request *CustomizeInstanceSegmentImageRequest, runtime *util.RuntimeOptions) (_result *CustomizeInstanceSegmentImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CustomizeInstanceSegmentImage"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeInstanceSegmentImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeInstanceSegmentImage(request *CustomizeInstanceSegmentImageRequest) (_result *CustomizeInstanceSegmentImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CustomizeInstanceSegmentImageResponse{}
	_body, _err := client.CustomizeInstanceSegmentImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeInstanceSegmentImageAdvance(request *CustomizeInstanceSegmentImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *CustomizeInstanceSegmentImageResponse, _err error) {
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
		Product:  tea.String("viapi-regen"),
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
	customizeInstanceSegmentImageReq := &CustomizeInstanceSegmentImageRequest{}
	openapiutil.Convert(request, customizeInstanceSegmentImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		customizeInstanceSegmentImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	customizeInstanceSegmentImageResp, _err := client.CustomizeInstanceSegmentImageWithOptions(customizeInstanceSegmentImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = customizeInstanceSegmentImageResp
	return _result, _err
}

func (client *Client) DebugServiceWithOptions(request *DebugServiceRequest, runtime *util.RuntimeOptions) (_result *DebugServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["Param"] = request.Param
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DebugServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DebugService(request *DebugServiceRequest) (_result *DebugServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DebugServiceResponse{}
	_body, _err := client.DebugServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatasetWithOptions(request *DeleteDatasetRequest, runtime *util.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataset(request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLabelsetWithOptions(request *DeleteLabelsetRequest, runtime *util.RuntimeOptions) (_result *DeleteLabelsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLabelset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLabelsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLabelset(request *DeleteLabelsetRequest) (_result *DeleteLabelsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLabelsetResponse{}
	_body, _err := client.DeleteLabelsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLabelsetDataWithOptions(request *DeleteLabelsetDataRequest, runtime *util.RuntimeOptions) (_result *DeleteLabelsetDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLabelsetData"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLabelsetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLabelsetData(request *DeleteLabelsetDataRequest) (_result *DeleteLabelsetDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLabelsetDataResponse{}
	_body, _err := client.DeleteLabelsetDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrainTaskWithOptions(request *DeleteTrainTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrainTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrainTask(request *DeleteTrainTaskRequest) (_result *DeleteTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrainTaskResponse{}
	_body, _err := client.DeleteTrainTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceWithOptions(request *DeleteWorkspaceRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspace"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadFileNameListWithOptions(request *DownloadFileNameListRequest, runtime *util.RuntimeOptions) (_result *DownloadFileNameListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["Identity"] = request.Identity
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadFileNameList"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadFileNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadFileNameList(request *DownloadFileNameListRequest) (_result *DownloadFileNameListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadFileNameListResponse{}
	_body, _err := client.DownloadFileNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadLabelFileWithOptions(request *DownloadLabelFileRequest, runtime *util.RuntimeOptions) (_result *DownloadLabelFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadLabelFile"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadLabelFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadLabelFile(request *DownloadLabelFileRequest) (_result *DownloadLabelFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadLabelFileResponse{}
	_body, _err := client.DownloadLabelFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadTemplateWithOptions(runtime *util.RuntimeOptions) (_result *DownloadTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DownloadTemplate"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadTemplate() (_result *DownloadTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadTemplateResponse{}
	_body, _err := client.DownloadTemplateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatasetWithOptions(request *GetDatasetRequest, runtime *util.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataset(request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiffCountLabelsetAndDatasetWithOptions(request *GetDiffCountLabelsetAndDatasetRequest, runtime *util.RuntimeOptions) (_result *GetDiffCountLabelsetAndDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelsetId)) {
		body["LabelsetId"] = request.LabelsetId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiffCountLabelsetAndDataset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiffCountLabelsetAndDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiffCountLabelsetAndDataset(request *GetDiffCountLabelsetAndDatasetRequest) (_result *GetDiffCountLabelsetAndDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiffCountLabelsetAndDatasetResponse{}
	_body, _err := client.GetDiffCountLabelsetAndDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLabelDetailWithOptions(request *GetLabelDetailRequest, runtime *util.RuntimeOptions) (_result *GetLabelDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLabelDetail"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLabelDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLabelDetail(request *GetLabelDetailRequest) (_result *GetLabelDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLabelDetailResponse{}
	_body, _err := client.GetLabelDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLabelsetWithOptions(request *GetLabelsetRequest, runtime *util.RuntimeOptions) (_result *GetLabelsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLabelset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLabelsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLabelset(request *GetLabelsetRequest) (_result *GetLabelsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLabelsetResponse{}
	_body, _err := client.GetLabelsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetService(request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrainModelWithOptions(request *GetTrainModelRequest, runtime *util.RuntimeOptions) (_result *GetTrainModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainModel"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrainModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrainModel(request *GetTrainModelRequest) (_result *GetTrainModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrainModelResponse{}
	_body, _err := client.GetTrainModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrainTaskWithOptions(request *GetTrainTaskRequest, runtime *util.RuntimeOptions) (_result *GetTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrainTask(request *GetTrainTaskRequest) (_result *GetTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrainTaskResponse{}
	_body, _err := client.GetTrainTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrainTaskEstimatedTimeWithOptions(request *GetTrainTaskEstimatedTimeRequest, runtime *util.RuntimeOptions) (_result *GetTrainTaskEstimatedTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrainTaskEstimatedTime"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrainTaskEstimatedTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrainTaskEstimatedTime(request *GetTrainTaskEstimatedTimeRequest) (_result *GetTrainTaskEstimatedTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrainTaskEstimatedTimeResponse{}
	_body, _err := client.GetTrainTaskEstimatedTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadPolicyWithOptions(request *GetUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *GetUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUploadPolicy"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadPolicy(request *GetUploadPolicyRequest) (_result *GetUploadPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUploadPolicyResponse{}
	_body, _err := client.GetUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(request *GetWorkspaceRequest, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasetDatasWithOptions(request *ListDatasetDatasRequest, runtime *util.RuntimeOptions) (_result *ListDatasetDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["Identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasetDatas"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasetDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatasetDatas(request *ListDatasetDatasRequest) (_result *ListDatasetDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatasetDatasResponse{}
	_body, _err := client.ListDatasetDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, runtime *util.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasets"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLabelsetDatasWithOptions(request *ListLabelsetDatasRequest, runtime *util.RuntimeOptions) (_result *ListLabelsetDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["LabelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLabelsetDatas"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLabelsetDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLabelsetDatas(request *ListLabelsetDatasRequest) (_result *ListLabelsetDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLabelsetDatasResponse{}
	_body, _err := client.ListLabelsetDatasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLabelsetsWithOptions(request *ListLabelsetsRequest, runtime *util.RuntimeOptions) (_result *ListLabelsetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLabelsets"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLabelsetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLabelsets(request *ListLabelsetsRequest) (_result *ListLabelsetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLabelsetsResponse{}
	_body, _err := client.ListLabelsetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTrainTasksWithOptions(request *ListTrainTasksRequest, runtime *util.RuntimeOptions) (_result *ListTrainTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrainTasks"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrainTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrainTasks(request *ListTrainTasksRequest) (_result *ListTrainTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrainTasksResponse{}
	_body, _err := client.ListTrainTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDatasetUserOssPathWithOptions(request *SetDatasetUserOssPathRequest, runtime *util.RuntimeOptions) (_result *SetDatasetUserOssPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.UserOssUrl)) {
		body["UserOssUrl"] = request.UserOssUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDatasetUserOssPath"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDatasetUserOssPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDatasetUserOssPath(request *SetDatasetUserOssPathRequest) (_result *SetDatasetUserOssPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDatasetUserOssPathResponse{}
	_body, _err := client.SetDatasetUserOssPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartServiceWithOptions(request *StartServiceRequest, runtime *util.RuntimeOptions) (_result *StartServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartService(request *StartServiceRequest) (_result *StartServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartServiceResponse{}
	_body, _err := client.StartServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTrainTaskWithOptions(request *StartTrainTaskRequest, runtime *util.RuntimeOptions) (_result *StartTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTrainTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTrainTask(request *StartTrainTaskRequest) (_result *StartTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTrainTaskResponse{}
	_body, _err := client.StartTrainTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopServiceWithOptions(request *StopServiceRequest, runtime *util.RuntimeOptions) (_result *StopServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopService(request *StopServiceRequest) (_result *StopServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopServiceResponse{}
	_body, _err := client.StopServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTrainTaskWithOptions(request *StopTrainTaskRequest, runtime *util.RuntimeOptions) (_result *StopTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopTrainTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTrainTask(request *StopTrainTaskRequest) (_result *StopTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopTrainTaskResponse{}
	_body, _err := client.StopTrainTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDatasetWithOptions(request *UpdateDatasetRequest, runtime *util.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataset(request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLabelsetWithOptions(request *UpdateLabelsetRequest, runtime *util.RuntimeOptions) (_result *UpdateLabelsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKey)) {
		body["ObjectKey"] = request.ObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserOssUrl)) {
		body["UserOssUrl"] = request.UserOssUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLabelset"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLabelsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLabelset(request *UpdateLabelsetRequest) (_result *UpdateLabelsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLabelsetResponse{}
	_body, _err := client.UpdateLabelsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceWithOptions(request *UpdateServiceRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateService(request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTrainTaskWithOptions(request *UpdateTrainTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateTrainTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrainTask"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTrainTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrainTask(request *UpdateTrainTaskRequest) (_result *UpdateTrainTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTrainTaskResponse{}
	_body, _err := client.UpdateTrainTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceWithOptions(request *UpdateWorkspaceRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspace"),
		Version:     tea.String("2021-11-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspace(request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
