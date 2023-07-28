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

type AddImageRequest struct {
	Description *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageUri    *string                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels      []*AddImageRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name        *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetDescription(v string) *AddImageRequest {
	s.Description = &v
	return s
}

func (s *AddImageRequest) SetImageUri(v string) *AddImageRequest {
	s.ImageUri = &v
	return s
}

func (s *AddImageRequest) SetLabels(v []*AddImageRequestLabels) *AddImageRequest {
	s.Labels = v
	return s
}

func (s *AddImageRequest) SetName(v string) *AddImageRequest {
	s.Name = &v
	return s
}

type AddImageRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddImageRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequestLabels) GoString() string {
	return s.String()
}

func (s *AddImageRequestLabels) SetKey(v string) *AddImageRequestLabels {
	s.Key = &v
	return s
}

func (s *AddImageRequestLabels) SetValue(v string) *AddImageRequestLabels {
	s.Value = &v
	return s
}

type AddImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetImageId(v string) *AddImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type AddImageLabelsRequest struct {
	Labels []*AddImageLabelsRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s AddImageLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddImageLabelsRequest) SetLabels(v []*AddImageLabelsRequestLabels) *AddImageLabelsRequest {
	s.Labels = v
	return s
}

type AddImageLabelsRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddImageLabelsRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsRequestLabels) GoString() string {
	return s.String()
}

func (s *AddImageLabelsRequestLabels) SetKey(v string) *AddImageLabelsRequestLabels {
	s.Key = &v
	return s
}

func (s *AddImageLabelsRequestLabels) SetValue(v string) *AddImageLabelsRequestLabels {
	s.Value = &v
	return s
}

type AddImageLabelsResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageLabelsResponseBody) SetRequestId(v string) *AddImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type AddImageLabelsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddImageLabelsResponse) SetHeaders(v map[string]*string) *AddImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddImageLabelsResponse) SetStatusCode(v int32) *AddImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageLabelsResponse) SetBody(v *AddImageLabelsResponseBody) *AddImageLabelsResponse {
	s.Body = v
	return s
}

type CopyExperimentRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FolderId      *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CopyExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyExperimentRequest) GoString() string {
	return s.String()
}

func (s *CopyExperimentRequest) SetAccessibility(v string) *CopyExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CopyExperimentRequest) SetDescription(v string) *CopyExperimentRequest {
	s.Description = &v
	return s
}

func (s *CopyExperimentRequest) SetFolderId(v string) *CopyExperimentRequest {
	s.FolderId = &v
	return s
}

func (s *CopyExperimentRequest) SetName(v string) *CopyExperimentRequest {
	s.Name = &v
	return s
}

func (s *CopyExperimentRequest) SetSource(v string) *CopyExperimentRequest {
	s.Source = &v
	return s
}

func (s *CopyExperimentRequest) SetWorkspaceId(v string) *CopyExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type CopyExperimentResponseBody struct {
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CopyExperimentResponseBody) SetExperimentId(v string) *CopyExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CopyExperimentResponseBody) SetRequestId(v string) *CopyExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CopyExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CopyExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyExperimentResponse) GoString() string {
	return s.String()
}

func (s *CopyExperimentResponse) SetHeaders(v map[string]*string) *CopyExperimentResponse {
	s.Headers = v
	return s
}

func (s *CopyExperimentResponse) SetStatusCode(v int32) *CopyExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyExperimentResponse) SetBody(v *CopyExperimentResponseBody) *CopyExperimentResponse {
	s.Body = v
	return s
}

type CreateExperimentRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FolderId      *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options       *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) SetAccessibility(v string) *CreateExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateExperimentRequest) SetDescription(v string) *CreateExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentRequest) SetFolderId(v string) *CreateExperimentRequest {
	s.FolderId = &v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetOptions(v string) *CreateExperimentRequest {
	s.Options = &v
	return s
}

func (s *CreateExperimentRequest) SetSource(v string) *CreateExperimentRequest {
	s.Source = &v
	return s
}

func (s *CreateExperimentRequest) SetTemplateId(v string) *CreateExperimentRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateExperimentRequest) SetWorkspaceId(v string) *CreateExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type CreateExperimentResponseBody struct {
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponseBody) SetExperimentId(v string) *CreateExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateExperimentResponseBody) SetRequestId(v string) *CreateExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponse) SetHeaders(v map[string]*string) *CreateExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentResponse) SetStatusCode(v int32) *CreateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentResponse) SetBody(v *CreateExperimentResponseBody) *CreateExperimentResponse {
	s.Body = v
	return s
}

type CreateExperimentFolderRequest struct {
	Accessibility  *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentFolderRequest) SetAccessibility(v string) *CreateExperimentFolderRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetName(v string) *CreateExperimentFolderRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetParentFolderId(v string) *CreateExperimentFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetSource(v string) *CreateExperimentFolderRequest {
	s.Source = &v
	return s
}

func (s *CreateExperimentFolderRequest) SetWorkspaceId(v string) *CreateExperimentFolderRequest {
	s.WorkspaceId = &v
	return s
}

type CreateExperimentFolderResponseBody struct {
	FolderId  *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentFolderResponseBody) SetFolderId(v string) *CreateExperimentFolderResponseBody {
	s.FolderId = &v
	return s
}

func (s *CreateExperimentFolderResponseBody) SetRequestId(v string) *CreateExperimentFolderResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentFolderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateExperimentFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateExperimentFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentFolderResponse) SetHeaders(v map[string]*string) *CreateExperimentFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentFolderResponse) SetStatusCode(v int32) *CreateExperimentFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentFolderResponse) SetBody(v *CreateExperimentFolderResponseBody) *CreateExperimentFolderResponse {
	s.Body = v
	return s
}

type CreateExperimentMigrateValidationRequest struct {
	SourceExpId *int64 `json:"SourceExpId,omitempty" xml:"SourceExpId,omitempty"`
}

func (s CreateExperimentMigrateValidationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentMigrateValidationRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentMigrateValidationRequest) SetSourceExpId(v int64) *CreateExperimentMigrateValidationRequest {
	s.SourceExpId = &v
	return s
}

type CreateExperimentMigrateValidationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentMigrateValidationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentMigrateValidationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentMigrateValidationResponseBody) SetRequestId(v string) *CreateExperimentMigrateValidationResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentMigrateValidationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateExperimentMigrateValidationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateExperimentMigrateValidationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentMigrateValidationResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentMigrateValidationResponse) SetHeaders(v map[string]*string) *CreateExperimentMigrateValidationResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentMigrateValidationResponse) SetStatusCode(v int32) *CreateExperimentMigrateValidationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentMigrateValidationResponse) SetBody(v *CreateExperimentMigrateValidationResponseBody) *CreateExperimentMigrateValidationResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	ExecuteType  *string `json:"ExecuteType,omitempty" xml:"ExecuteType,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Options      *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetExecuteType(v string) *CreateJobRequest {
	s.ExecuteType = &v
	return s
}

func (s *CreateJobRequest) SetExperimentId(v string) *CreateJobRequest {
	s.ExperimentId = &v
	return s
}

func (s *CreateJobRequest) SetNodeId(v string) *CreateJobRequest {
	s.NodeId = &v
	return s
}

func (s *CreateJobRequest) SetOptions(v string) *CreateJobRequest {
	s.Options = &v
	return s
}

type CreateJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type DeleteExperimentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponseBody) SetRequestId(v string) *DeleteExperimentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponse) SetHeaders(v map[string]*string) *DeleteExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentResponse) SetStatusCode(v int32) *DeleteExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentResponse) SetBody(v *DeleteExperimentResponseBody) *DeleteExperimentResponse {
	s.Body = v
	return s
}

type DeleteExperimentFolderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentFolderResponseBody) SetRequestId(v string) *DeleteExperimentFolderResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentFolderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteExperimentFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExperimentFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentFolderResponse) SetHeaders(v map[string]*string) *DeleteExperimentFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentFolderResponse) SetStatusCode(v int32) *DeleteExperimentFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentFolderResponse) SetBody(v *DeleteExperimentFolderResponseBody) *DeleteExperimentFolderResponse {
	s.Body = v
	return s
}

type GetAlgoTreeRequest struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetAlgoTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgoTreeRequest) GoString() string {
	return s.String()
}

func (s *GetAlgoTreeRequest) SetSource(v string) *GetAlgoTreeRequest {
	s.Source = &v
	return s
}

type GetAlgoTreeResponseBody struct {
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlgoTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgoTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgoTreeResponseBody) SetData(v map[string]interface{}) *GetAlgoTreeResponseBody {
	s.Data = v
	return s
}

func (s *GetAlgoTreeResponseBody) SetRequestId(v string) *GetAlgoTreeResponseBody {
	s.RequestId = &v
	return s
}

type GetAlgoTreeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgoTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgoTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgoTreeResponse) GoString() string {
	return s.String()
}

func (s *GetAlgoTreeResponse) SetHeaders(v map[string]*string) *GetAlgoTreeResponse {
	s.Headers = v
	return s
}

func (s *GetAlgoTreeResponse) SetStatusCode(v int32) *GetAlgoTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgoTreeResponse) SetBody(v *GetAlgoTreeResponseBody) *GetAlgoTreeResponse {
	s.Body = v
	return s
}

type GetAlgorithmDefRequest struct {
	AlgoVersion *string `json:"AlgoVersion,omitempty" xml:"AlgoVersion,omitempty"`
	Identifier  *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Provider    *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetAlgorithmDefRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDefRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefRequest) SetAlgoVersion(v string) *GetAlgorithmDefRequest {
	s.AlgoVersion = &v
	return s
}

func (s *GetAlgorithmDefRequest) SetIdentifier(v string) *GetAlgorithmDefRequest {
	s.Identifier = &v
	return s
}

func (s *GetAlgorithmDefRequest) SetProvider(v string) *GetAlgorithmDefRequest {
	s.Provider = &v
	return s
}

func (s *GetAlgorithmDefRequest) SetSignature(v string) *GetAlgorithmDefRequest {
	s.Signature = &v
	return s
}

type GetAlgorithmDefResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spec      map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetAlgorithmDefResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDefResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefResponseBody) SetRequestId(v string) *GetAlgorithmDefResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDefResponseBody) SetSpec(v map[string]interface{}) *GetAlgorithmDefResponseBody {
	s.Spec = v
	return s
}

type GetAlgorithmDefResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgorithmDefResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmDefResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDefResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefResponse) SetHeaders(v map[string]*string) *GetAlgorithmDefResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDefResponse) SetStatusCode(v int32) *GetAlgorithmDefResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmDefResponse) SetBody(v *GetAlgorithmDefResponseBody) *GetAlgorithmDefResponse {
	s.Body = v
	return s
}

type GetAlgorithmDefsRequest struct {
	LatestTimestamp *string `json:"LatestTimestamp,omitempty" xml:"LatestTimestamp,omitempty"`
	RangeEnd        *int32  `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	RangeStart      *int32  `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
	Timestamp       *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetAlgorithmDefsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDefsRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefsRequest) SetLatestTimestamp(v string) *GetAlgorithmDefsRequest {
	s.LatestTimestamp = &v
	return s
}

func (s *GetAlgorithmDefsRequest) SetRangeEnd(v int32) *GetAlgorithmDefsRequest {
	s.RangeEnd = &v
	return s
}

func (s *GetAlgorithmDefsRequest) SetRangeStart(v int32) *GetAlgorithmDefsRequest {
	s.RangeStart = &v
	return s
}

func (s *GetAlgorithmDefsRequest) SetTimestamp(v string) *GetAlgorithmDefsRequest {
	s.Timestamp = &v
	return s
}

type GetAlgorithmDefsResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Specs     []map[string]interface{} `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s GetAlgorithmDefsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDefsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefsResponseBody) SetRequestId(v string) *GetAlgorithmDefsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDefsResponseBody) SetSpecs(v []map[string]interface{}) *GetAlgorithmDefsResponseBody {
	s.Specs = v
	return s
}

type GetAlgorithmDefsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgorithmDefsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmDefsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmDefsResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefsResponse) SetHeaders(v map[string]*string) *GetAlgorithmDefsResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDefsResponse) SetStatusCode(v int32) *GetAlgorithmDefsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmDefsResponse) SetBody(v *GetAlgorithmDefsResponseBody) *GetAlgorithmDefsResponse {
	s.Body = v
	return s
}

type GetAlgorithmTreeRequest struct {
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAlgorithmTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmTreeRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmTreeRequest) SetSource(v string) *GetAlgorithmTreeRequest {
	s.Source = &v
	return s
}

func (s *GetAlgorithmTreeRequest) SetWorkspaceId(v string) *GetAlgorithmTreeRequest {
	s.WorkspaceId = &v
	return s
}

type GetAlgorithmTreeResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp *string                  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Tree      []map[string]interface{} `json:"Tree,omitempty" xml:"Tree,omitempty" type:"Repeated"`
}

func (s GetAlgorithmTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmTreeResponseBody) SetRequestId(v string) *GetAlgorithmTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmTreeResponseBody) SetTimestamp(v string) *GetAlgorithmTreeResponseBody {
	s.Timestamp = &v
	return s
}

func (s *GetAlgorithmTreeResponseBody) SetTree(v []map[string]interface{}) *GetAlgorithmTreeResponseBody {
	s.Tree = v
	return s
}

type GetAlgorithmTreeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlgorithmTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlgorithmTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlgorithmTreeResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmTreeResponse) SetHeaders(v map[string]*string) *GetAlgorithmTreeResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmTreeResponse) SetStatusCode(v int32) *GetAlgorithmTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmTreeResponse) SetBody(v *GetAlgorithmTreeResponseBody) *GetAlgorithmTreeResponse {
	s.Body = v
	return s
}

type GetExperimentResponseBody struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Creator         *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options         *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Version         *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBody) SetAccessibility(v string) *GetExperimentResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetExperimentResponseBody) SetContent(v string) *GetExperimentResponseBody {
	s.Content = &v
	return s
}

func (s *GetExperimentResponseBody) SetCreator(v string) *GetExperimentResponseBody {
	s.Creator = &v
	return s
}

func (s *GetExperimentResponseBody) SetDescription(v string) *GetExperimentResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentResponseBody) SetExperimentId(v string) *GetExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtCreateTime(v string) *GetExperimentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetGmtModifiedTime(v string) *GetExperimentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentResponseBody) SetName(v string) *GetExperimentResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentResponseBody) SetOptions(v string) *GetExperimentResponseBody {
	s.Options = &v
	return s
}

func (s *GetExperimentResponseBody) SetRequestId(v string) *GetExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResponseBody) SetSource(v string) *GetExperimentResponseBody {
	s.Source = &v
	return s
}

func (s *GetExperimentResponseBody) SetVersion(v int64) *GetExperimentResponseBody {
	s.Version = &v
	return s
}

func (s *GetExperimentResponseBody) SetWorkspaceId(v string) *GetExperimentResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetExperimentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentResponse) SetHeaders(v map[string]*string) *GetExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentResponse) SetStatusCode(v int32) *GetExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentResponse) SetBody(v *GetExperimentResponseBody) *GetExperimentResponse {
	s.Body = v
	return s
}

type GetExperimentFolderChildrenRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	OnlyFolder    *bool   `json:"OnlyFolder,omitempty" xml:"OnlyFolder,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentFolderChildrenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentFolderChildrenRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenRequest) SetAccessibility(v string) *GetExperimentFolderChildrenRequest {
	s.Accessibility = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetOnlyFolder(v bool) *GetExperimentFolderChildrenRequest {
	s.OnlyFolder = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetSource(v string) *GetExperimentFolderChildrenRequest {
	s.Source = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetUserId(v string) *GetExperimentFolderChildrenRequest {
	s.UserId = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetWorkspaceId(v string) *GetExperimentFolderChildrenRequest {
	s.WorkspaceId = &v
	return s
}

type GetExperimentFolderChildrenResponseBody struct {
	Items      []*GetExperimentFolderChildrenResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetExperimentFolderChildrenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentFolderChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenResponseBody) SetItems(v []*GetExperimentFolderChildrenResponseBodyItems) *GetExperimentFolderChildrenResponseBody {
	s.Items = v
	return s
}

func (s *GetExperimentFolderChildrenResponseBody) SetRequestId(v string) *GetExperimentFolderChildrenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBody) SetTotalCount(v int32) *GetExperimentFolderChildrenResponseBody {
	s.TotalCount = &v
	return s
}

type GetExperimentFolderChildrenResponseBodyItems struct {
	Empty           *bool   `json:"Empty,omitempty" xml:"Empty,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Icon            *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExperimentFolderChildrenResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentFolderChildrenResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetEmpty(v bool) *GetExperimentFolderChildrenResponseBodyItems {
	s.Empty = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetGmtCreateTime(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetGmtModifiedTime(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetIcon(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Icon = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetId(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Id = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetName(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Name = &v
	return s
}

func (s *GetExperimentFolderChildrenResponseBodyItems) SetType(v string) *GetExperimentFolderChildrenResponseBodyItems {
	s.Type = &v
	return s
}

type GetExperimentFolderChildrenResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentFolderChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentFolderChildrenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentFolderChildrenResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenResponse) SetHeaders(v map[string]*string) *GetExperimentFolderChildrenResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentFolderChildrenResponse) SetStatusCode(v int32) *GetExperimentFolderChildrenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentFolderChildrenResponse) SetBody(v *GetExperimentFolderChildrenResponseBody) *GetExperimentFolderChildrenResponse {
	s.Body = v
	return s
}

type GetExperimentMetaResponseBody struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator         *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options         *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaResponseBody) SetAccessibility(v string) *GetExperimentMetaResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetCreator(v string) *GetExperimentMetaResponseBody {
	s.Creator = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetDescription(v string) *GetExperimentMetaResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetExperimentId(v string) *GetExperimentMetaResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetGmtCreateTime(v string) *GetExperimentMetaResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetGmtModifiedTime(v string) *GetExperimentMetaResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetName(v string) *GetExperimentMetaResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetOptions(v string) *GetExperimentMetaResponseBody {
	s.Options = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetRequestId(v string) *GetExperimentMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetSource(v string) *GetExperimentMetaResponseBody {
	s.Source = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetVersion(v string) *GetExperimentMetaResponseBody {
	s.Version = &v
	return s
}

func (s *GetExperimentMetaResponseBody) SetWorkspaceId(v string) *GetExperimentMetaResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetExperimentMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentMetaResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaResponse) SetHeaders(v map[string]*string) *GetExperimentMetaResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentMetaResponse) SetStatusCode(v int32) *GetExperimentMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentMetaResponse) SetBody(v *GetExperimentMetaResponseBody) *GetExperimentMetaResponse {
	s.Body = v
	return s
}

type GetExperimentStatusResponseBody struct {
	Nodes     []*GetExperimentStatusResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExperimentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentStatusResponseBody) SetNodes(v []*GetExperimentStatusResponseBodyNodes) *GetExperimentStatusResponseBody {
	s.Nodes = v
	return s
}

func (s *GetExperimentStatusResponseBody) SetRequestId(v string) *GetExperimentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentStatusResponseBody) SetStatus(v string) *GetExperimentStatusResponseBody {
	s.Status = &v
	return s
}

type GetExperimentStatusResponseBodyNodes struct {
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	RunId      *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	RunNodeId  *string `json:"RunNodeId,omitempty" xml:"RunNodeId,omitempty"`
	StartedAt  *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExperimentStatusResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentStatusResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *GetExperimentStatusResponseBodyNodes) SetFinishedAt(v string) *GetExperimentStatusResponseBodyNodes {
	s.FinishedAt = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetJobId(v string) *GetExperimentStatusResponseBodyNodes {
	s.JobId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetNodeId(v string) *GetExperimentStatusResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetRunId(v string) *GetExperimentStatusResponseBodyNodes {
	s.RunId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetRunNodeId(v string) *GetExperimentStatusResponseBodyNodes {
	s.RunNodeId = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetStartedAt(v string) *GetExperimentStatusResponseBodyNodes {
	s.StartedAt = &v
	return s
}

func (s *GetExperimentStatusResponseBodyNodes) SetStatus(v string) *GetExperimentStatusResponseBodyNodes {
	s.Status = &v
	return s
}

type GetExperimentStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentStatusResponse) SetHeaders(v map[string]*string) *GetExperimentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentStatusResponse) SetStatusCode(v int32) *GetExperimentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentStatusResponse) SetBody(v *GetExperimentStatusResponseBody) *GetExperimentStatusResponse {
	s.Body = v
	return s
}

type GetExperimentVisualizationMetaRequest struct {
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
}

func (s GetExperimentVisualizationMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentVisualizationMetaRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaRequest) SetNodeIds(v string) *GetExperimentVisualizationMetaRequest {
	s.NodeIds = &v
	return s
}

type GetExperimentVisualizationMetaResponseBody struct {
	VisualizationMeta []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta `json:"VisualizationMeta,omitempty" xml:"VisualizationMeta,omitempty" type:"Repeated"`
	RequestId         *string                                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetExperimentVisualizationMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentVisualizationMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaResponseBody) SetVisualizationMeta(v []*GetExperimentVisualizationMetaResponseBodyVisualizationMeta) *GetExperimentVisualizationMetaResponseBody {
	s.VisualizationMeta = v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBody) SetRequestId(v string) *GetExperimentVisualizationMetaResponseBody {
	s.RequestId = &v
	return s
}

type GetExperimentVisualizationMetaResponseBodyVisualizationMeta struct {
	Meta     *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetExperimentVisualizationMetaResponseBodyVisualizationMeta) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentVisualizationMetaResponseBodyVisualizationMeta) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) SetMeta(v string) *GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	s.Meta = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) SetNodeId(v string) *GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	s.NodeId = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponseBodyVisualizationMeta) SetNodeName(v string) *GetExperimentVisualizationMetaResponseBodyVisualizationMeta {
	s.NodeName = &v
	return s
}

type GetExperimentVisualizationMetaResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentVisualizationMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentVisualizationMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentVisualizationMetaResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentVisualizationMetaResponse) SetHeaders(v map[string]*string) *GetExperimentVisualizationMetaResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentVisualizationMetaResponse) SetStatusCode(v int32) *GetExperimentVisualizationMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentVisualizationMetaResponse) SetBody(v *GetExperimentVisualizationMetaResponseBody) *GetExperimentVisualizationMetaResponse {
	s.Body = v
	return s
}

type GetExperimentsStatisticsRequest struct {
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s GetExperimentsStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsRequest) SetSource(v string) *GetExperimentsStatisticsRequest {
	s.Source = &v
	return s
}

func (s *GetExperimentsStatisticsRequest) SetWorkspaceIds(v string) *GetExperimentsStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

type GetExperimentsStatisticsResponseBody struct {
	Data      []*GetExperimentsStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExperimentsStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsResponseBody) SetData(v []*GetExperimentsStatisticsResponseBodyData) *GetExperimentsStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetExperimentsStatisticsResponseBody) SetRequestId(v string) *GetExperimentsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetExperimentsStatisticsResponseBodyData struct {
	CreateCount *int64  `json:"CreateCount,omitempty" xml:"CreateCount,omitempty"`
	TotalCount  *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentsStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsResponseBodyData) SetCreateCount(v int64) *GetExperimentsStatisticsResponseBodyData {
	s.CreateCount = &v
	return s
}

func (s *GetExperimentsStatisticsResponseBodyData) SetTotalCount(v int64) *GetExperimentsStatisticsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetExperimentsStatisticsResponseBodyData) SetWorkspaceId(v string) *GetExperimentsStatisticsResponseBodyData {
	s.WorkspaceId = &v
	return s
}

type GetExperimentsStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentsStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentsStatisticsResponse) SetHeaders(v map[string]*string) *GetExperimentsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentsStatisticsResponse) SetStatusCode(v int32) *GetExperimentsStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentsStatisticsResponse) SetBody(v *GetExperimentsStatisticsResponseBody) *GetExperimentsStatisticsResponse {
	s.Body = v
	return s
}

type GetExperimentsUsersStatisticsRequest struct {
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentsUsersStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsUsersStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsRequest) SetSource(v string) *GetExperimentsUsersStatisticsRequest {
	s.Source = &v
	return s
}

func (s *GetExperimentsUsersStatisticsRequest) SetWorkspaceId(v string) *GetExperimentsUsersStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetExperimentsUsersStatisticsResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*GetExperimentsUsersStatisticsResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s GetExperimentsUsersStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsUsersStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsResponseBody) SetRequestId(v string) *GetExperimentsUsersStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentsUsersStatisticsResponseBody) SetUsers(v []*GetExperimentsUsersStatisticsResponseBodyUsers) *GetExperimentsUsersStatisticsResponseBody {
	s.Users = v
	return s
}

type GetExperimentsUsersStatisticsResponseBodyUsers struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetExperimentsUsersStatisticsResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsUsersStatisticsResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsResponseBodyUsers) SetUserId(v string) *GetExperimentsUsersStatisticsResponseBodyUsers {
	s.UserId = &v
	return s
}

type GetExperimentsUsersStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExperimentsUsersStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExperimentsUsersStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentsUsersStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentsUsersStatisticsResponse) SetHeaders(v map[string]*string) *GetExperimentsUsersStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentsUsersStatisticsResponse) SetStatusCode(v int32) *GetExperimentsUsersStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentsUsersStatisticsResponse) SetBody(v *GetExperimentsUsersStatisticsResponseBody) *GetExperimentsUsersStatisticsResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetVerbose(v bool) *GetImageRequest {
	s.Verbose = &v
	return s
}

type GetImageResponseBody struct {
	Description          *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime        *string                       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime      *string                       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageUri             *string                       `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels               []*GetImageResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name                 *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	OperatorCreate       *string                       `json:"OperatorCreate,omitempty" xml:"OperatorCreate,omitempty"`
	ParentOperatorCreate *string                       `json:"ParentOperatorCreate,omitempty" xml:"ParentOperatorCreate,omitempty"`
	RequestId            *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetDescription(v string) *GetImageResponseBody {
	s.Description = &v
	return s
}

func (s *GetImageResponseBody) SetGmtCreateTime(v string) *GetImageResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetImageResponseBody) SetGmtModifiedTime(v string) *GetImageResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageUri(v string) *GetImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageResponseBody) SetLabels(v []*GetImageResponseBodyLabels) *GetImageResponseBody {
	s.Labels = v
	return s
}

func (s *GetImageResponseBody) SetName(v string) *GetImageResponseBody {
	s.Name = &v
	return s
}

func (s *GetImageResponseBody) SetOperatorCreate(v string) *GetImageResponseBody {
	s.OperatorCreate = &v
	return s
}

func (s *GetImageResponseBody) SetParentOperatorCreate(v string) *GetImageResponseBody {
	s.ParentOperatorCreate = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

type GetImageResponseBodyLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetImageResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyLabels) SetKey(v string) *GetImageResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetImageResponseBodyLabels) SetValue(v string) *GetImageResponseBodyLabels {
	s.Value = &v
	return s
}

type GetImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetStatusCode(v int32) *GetImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetVerbose(v bool) *GetJobRequest {
	s.Verbose = &v
	return s
}

type GetJobResponseBody struct {
	Arguments     *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	ExecuteType   *string `json:"ExecuteType,omitempty" xml:"ExecuteType,omitempty"`
	ExperimentId  *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PaiflowNodeId *string `json:"PaiflowNodeId,omitempty" xml:"PaiflowNodeId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunId         *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	RunInfo       *string `json:"RunInfo,omitempty" xml:"RunInfo,omitempty"`
	Snapshot      *string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetArguments(v string) *GetJobResponseBody {
	s.Arguments = &v
	return s
}

func (s *GetJobResponseBody) SetCreator(v string) *GetJobResponseBody {
	s.Creator = &v
	return s
}

func (s *GetJobResponseBody) SetExecuteType(v string) *GetJobResponseBody {
	s.ExecuteType = &v
	return s
}

func (s *GetJobResponseBody) SetExperimentId(v string) *GetJobResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *GetJobResponseBody) SetGmtCreateTime(v string) *GetJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBody) SetJobId(v string) *GetJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBody) SetNodeId(v string) *GetJobResponseBody {
	s.NodeId = &v
	return s
}

func (s *GetJobResponseBody) SetPaiflowNodeId(v string) *GetJobResponseBody {
	s.PaiflowNodeId = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetRunId(v string) *GetJobResponseBody {
	s.RunId = &v
	return s
}

func (s *GetJobResponseBody) SetRunInfo(v string) *GetJobResponseBody {
	s.RunInfo = &v
	return s
}

func (s *GetJobResponseBody) SetSnapshot(v string) *GetJobResponseBody {
	s.Snapshot = &v
	return s
}

func (s *GetJobResponseBody) SetStatus(v string) *GetJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceId(v string) *GetJobResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetMCTableSchemaRequest struct {
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMCTableSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMCTableSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaRequest) SetWorkspaceId(v string) *GetMCTableSchemaRequest {
	s.WorkspaceId = &v
	return s
}

type GetMCTableSchemaResponseBody struct {
	Columns          []*GetMCTableSchemaResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	PartitionColumns []*string                              `json:"PartitionColumns,omitempty" xml:"PartitionColumns,omitempty" type:"Repeated"`
	RequestId        *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMCTableSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMCTableSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaResponseBody) SetColumns(v []*GetMCTableSchemaResponseBodyColumns) *GetMCTableSchemaResponseBody {
	s.Columns = v
	return s
}

func (s *GetMCTableSchemaResponseBody) SetPartitionColumns(v []*string) *GetMCTableSchemaResponseBody {
	s.PartitionColumns = v
	return s
}

func (s *GetMCTableSchemaResponseBody) SetRequestId(v string) *GetMCTableSchemaResponseBody {
	s.RequestId = &v
	return s
}

type GetMCTableSchemaResponseBodyColumns struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Preview []*string `json:"Preview,omitempty" xml:"Preview,omitempty" type:"Repeated"`
	Type    *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMCTableSchemaResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s GetMCTableSchemaResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaResponseBodyColumns) SetName(v string) *GetMCTableSchemaResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *GetMCTableSchemaResponseBodyColumns) SetPreview(v []*string) *GetMCTableSchemaResponseBodyColumns {
	s.Preview = v
	return s
}

func (s *GetMCTableSchemaResponseBodyColumns) SetType(v string) *GetMCTableSchemaResponseBodyColumns {
	s.Type = &v
	return s
}

type GetMCTableSchemaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMCTableSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMCTableSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMCTableSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaResponse) SetHeaders(v map[string]*string) *GetMCTableSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetMCTableSchemaResponse) SetStatusCode(v int32) *GetMCTableSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMCTableSchemaResponse) SetBody(v *GetMCTableSchemaResponseBody) *GetMCTableSchemaResponse {
	s.Body = v
	return s
}

type GetNodeInputSchemaRequest struct {
	InputId    *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	InputIndex *int32  `json:"InputIndex,omitempty" xml:"InputIndex,omitempty"`
}

func (s GetNodeInputSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInputSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetNodeInputSchemaRequest) SetInputId(v string) *GetNodeInputSchemaRequest {
	s.InputId = &v
	return s
}

func (s *GetNodeInputSchemaRequest) SetInputIndex(v int32) *GetNodeInputSchemaRequest {
	s.InputIndex = &v
	return s
}

type GetNodeInputSchemaResponseBody struct {
	ColNames  []*string `json:"ColNames,omitempty" xml:"ColNames,omitempty" type:"Repeated"`
	ColTypes  []*string `json:"ColTypes,omitempty" xml:"ColTypes,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeInputSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInputSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeInputSchemaResponseBody) SetColNames(v []*string) *GetNodeInputSchemaResponseBody {
	s.ColNames = v
	return s
}

func (s *GetNodeInputSchemaResponseBody) SetColTypes(v []*string) *GetNodeInputSchemaResponseBody {
	s.ColTypes = v
	return s
}

func (s *GetNodeInputSchemaResponseBody) SetRequestId(v string) *GetNodeInputSchemaResponseBody {
	s.RequestId = &v
	return s
}

type GetNodeInputSchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeInputSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeInputSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeInputSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetNodeInputSchemaResponse) SetHeaders(v map[string]*string) *GetNodeInputSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetNodeInputSchemaResponse) SetStatusCode(v int32) *GetNodeInputSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeInputSchemaResponse) SetBody(v *GetNodeInputSchemaResponseBody) *GetNodeInputSchemaResponse {
	s.Body = v
	return s
}

type GetNodeOutputRequest struct {
	OutputIndex *string `json:"OutputIndex,omitempty" xml:"OutputIndex,omitempty"`
}

func (s GetNodeOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeOutputRequest) GoString() string {
	return s.String()
}

func (s *GetNodeOutputRequest) SetOutputIndex(v string) *GetNodeOutputRequest {
	s.OutputIndex = &v
	return s
}

type GetNodeOutputResponseBody struct {
	AlgoName     *string                `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	DisplayName  *string                `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	LocationType *string                `json:"LocationType,omitempty" xml:"LocationType,omitempty"`
	NodeName     *string                `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Type         *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        map[string]interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNodeOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeOutputResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeOutputResponseBody) SetAlgoName(v string) *GetNodeOutputResponseBody {
	s.AlgoName = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetDisplayName(v string) *GetNodeOutputResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetLocationType(v string) *GetNodeOutputResponseBody {
	s.LocationType = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetNodeName(v string) *GetNodeOutputResponseBody {
	s.NodeName = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetRequestId(v string) *GetNodeOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetType(v string) *GetNodeOutputResponseBody {
	s.Type = &v
	return s
}

func (s *GetNodeOutputResponseBody) SetValue(v map[string]interface{}) *GetNodeOutputResponseBody {
	s.Value = v
	return s
}

type GetNodeOutputResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeOutputResponse) GoString() string {
	return s.String()
}

func (s *GetNodeOutputResponse) SetHeaders(v map[string]*string) *GetNodeOutputResponse {
	s.Headers = v
	return s
}

func (s *GetNodeOutputResponse) SetStatusCode(v int32) *GetNodeOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeOutputResponse) SetBody(v *GetNodeOutputResponseBody) *GetNodeOutputResponse {
	s.Body = v
	return s
}

type GetNodeVisualizationRequest struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s GetNodeVisualizationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeVisualizationRequest) GoString() string {
	return s.String()
}

func (s *GetNodeVisualizationRequest) SetConfig(v string) *GetNodeVisualizationRequest {
	s.Config = &v
	return s
}

type GetNodeVisualizationResponseBody struct {
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VisualizationType *string `json:"VisualizationType,omitempty" xml:"VisualizationType,omitempty"`
}

func (s GetNodeVisualizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeVisualizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeVisualizationResponseBody) SetContent(v string) *GetNodeVisualizationResponseBody {
	s.Content = &v
	return s
}

func (s *GetNodeVisualizationResponseBody) SetRequestId(v string) *GetNodeVisualizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeVisualizationResponseBody) SetVisualizationType(v string) *GetNodeVisualizationResponseBody {
	s.VisualizationType = &v
	return s
}

type GetNodeVisualizationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeVisualizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeVisualizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeVisualizationResponse) GoString() string {
	return s.String()
}

func (s *GetNodeVisualizationResponse) SetHeaders(v map[string]*string) *GetNodeVisualizationResponse {
	s.Headers = v
	return s
}

func (s *GetNodeVisualizationResponse) SetStatusCode(v int32) *GetNodeVisualizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeVisualizationResponse) SetBody(v *GetNodeVisualizationResponseBody) *GetNodeVisualizationResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetVerbose(v bool) *GetTemplateRequest {
	s.Verbose = &v
	return s
}

type GetTemplateResponseBody struct {
	Content      *string                  `json:"Content,omitempty" xml:"Content,omitempty"`
	Description  *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail       *string                  `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DocLink      *string                  `json:"DocLink,omitempty" xml:"DocLink,omitempty"`
	ImageLink    *string                  `json:"ImageLink,omitempty" xml:"ImageLink,omitempty"`
	Labels       []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name         *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceId     *string                  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType   *string                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TemplateId   *string                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType *string                  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetContent(v string) *GetTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBody) SetDescription(v string) *GetTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBody) SetDetail(v string) *GetTemplateResponseBody {
	s.Detail = &v
	return s
}

func (s *GetTemplateResponseBody) SetDocLink(v string) *GetTemplateResponseBody {
	s.DocLink = &v
	return s
}

func (s *GetTemplateResponseBody) SetImageLink(v string) *GetTemplateResponseBody {
	s.ImageLink = &v
	return s
}

func (s *GetTemplateResponseBody) SetLabels(v []map[string]interface{}) *GetTemplateResponseBody {
	s.Labels = v
	return s
}

func (s *GetTemplateResponseBody) SetName(v string) *GetTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSourceId(v string) *GetTemplateResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSourceType(v string) *GetTemplateResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateId(v string) *GetTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateType(v string) *GetTemplateResponseBody {
	s.TemplateType = &v
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

type ListAlgoDefsRequest struct {
	Body []*ListAlgoDefsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListAlgoDefsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlgoDefsRequest) GoString() string {
	return s.String()
}

func (s *ListAlgoDefsRequest) SetBody(v []*ListAlgoDefsRequestBody) *ListAlgoDefsRequest {
	s.Body = v
	return s
}

type ListAlgoDefsRequestBody struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Provider   *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Signature  *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAlgoDefsRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlgoDefsRequestBody) GoString() string {
	return s.String()
}

func (s *ListAlgoDefsRequestBody) SetIdentifier(v string) *ListAlgoDefsRequestBody {
	s.Identifier = &v
	return s
}

func (s *ListAlgoDefsRequestBody) SetProvider(v string) *ListAlgoDefsRequestBody {
	s.Provider = &v
	return s
}

func (s *ListAlgoDefsRequestBody) SetSignature(v string) *ListAlgoDefsRequestBody {
	s.Signature = &v
	return s
}

func (s *ListAlgoDefsRequestBody) SetVersion(v string) *ListAlgoDefsRequestBody {
	s.Version = &v
	return s
}

type ListAlgoDefsResponseBody struct {
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlgoDefsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlgoDefsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgoDefsResponseBody) SetData(v []map[string]interface{}) *ListAlgoDefsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlgoDefsResponseBody) SetRequestId(v string) *ListAlgoDefsResponseBody {
	s.RequestId = &v
	return s
}

type ListAlgoDefsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlgoDefsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlgoDefsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlgoDefsResponse) GoString() string {
	return s.String()
}

func (s *ListAlgoDefsResponse) SetHeaders(v map[string]*string) *ListAlgoDefsResponse {
	s.Headers = v
	return s
}

func (s *ListAlgoDefsResponse) SetStatusCode(v int32) *ListAlgoDefsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlgoDefsResponse) SetBody(v *ListAlgoDefsResponseBody) *ListAlgoDefsResponse {
	s.Body = v
	return s
}

type ListAuthRolesRequest struct {
	IsGenerateToken *string `json:"IsGenerateToken,omitempty" xml:"IsGenerateToken,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAuthRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthRolesRequest) GoString() string {
	return s.String()
}

func (s *ListAuthRolesRequest) SetIsGenerateToken(v string) *ListAuthRolesRequest {
	s.IsGenerateToken = &v
	return s
}

func (s *ListAuthRolesRequest) SetWorkspaceId(v string) *ListAuthRolesRequest {
	s.WorkspaceId = &v
	return s
}

type ListAuthRolesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     []*ListAuthRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s ListAuthRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponseBody) SetRequestId(v string) *ListAuthRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthRolesResponseBody) SetRoles(v []*ListAuthRolesResponseBodyRoles) *ListAuthRolesResponseBody {
	s.Roles = v
	return s
}

type ListAuthRolesResponseBodyRoles struct {
	IsEnabled *string                              `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	RoleARN   *string                              `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	RoleName  *string                              `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RoleType  *string                              `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	Token     *ListAuthRolesResponseBodyRolesToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s ListAuthRolesResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s ListAuthRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponseBodyRoles) SetIsEnabled(v string) *ListAuthRolesResponseBodyRoles {
	s.IsEnabled = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetRoleARN(v string) *ListAuthRolesResponseBodyRoles {
	s.RoleARN = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetRoleName(v string) *ListAuthRolesResponseBodyRoles {
	s.RoleName = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetRoleType(v string) *ListAuthRolesResponseBodyRoles {
	s.RoleType = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetToken(v *ListAuthRolesResponseBodyRolesToken) *ListAuthRolesResponseBodyRoles {
	s.Token = v
	return s
}

type ListAuthRolesResponseBodyRolesToken struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAuthRolesResponseBodyRolesToken) String() string {
	return tea.Prettify(s)
}

func (s ListAuthRolesResponseBodyRolesToken) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponseBodyRolesToken) SetAccessKeyId(v string) *ListAuthRolesResponseBodyRolesToken {
	s.AccessKeyId = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) SetAccessKeySecret(v string) *ListAuthRolesResponseBodyRolesToken {
	s.AccessKeySecret = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) SetExpiration(v string) *ListAuthRolesResponseBodyRolesToken {
	s.Expiration = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) SetSecurityToken(v string) *ListAuthRolesResponseBodyRolesToken {
	s.SecurityToken = &v
	return s
}

type ListAuthRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAuthRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAuthRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthRolesResponse) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponse) SetHeaders(v map[string]*string) *ListAuthRolesResponse {
	s.Headers = v
	return s
}

func (s *ListAuthRolesResponse) SetStatusCode(v int32) *ListAuthRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthRolesResponse) SetBody(v *ListAuthRolesResponseBody) *ListAuthRolesResponse {
	s.Body = v
	return s
}

type ListExperimentsRequest struct {
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentsRequest) SetCreator(v string) *ListExperimentsRequest {
	s.Creator = &v
	return s
}

func (s *ListExperimentsRequest) SetExperimentId(v string) *ListExperimentsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsRequest) SetName(v string) *ListExperimentsRequest {
	s.Name = &v
	return s
}

func (s *ListExperimentsRequest) SetOrder(v string) *ListExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentsRequest) SetPageNumber(v int32) *ListExperimentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentsRequest) SetPageSize(v int32) *ListExperimentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentsRequest) SetSortBy(v string) *ListExperimentsRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentsRequest) SetSource(v string) *ListExperimentsRequest {
	s.Source = &v
	return s
}

func (s *ListExperimentsRequest) SetWorkspaceId(v string) *ListExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

type ListExperimentsResponseBody struct {
	Experiments []*ListExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBody) SetExperiments(v []*ListExperimentsResponseBodyExperiments) *ListExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListExperimentsResponseBody) SetRequestId(v string) *ListExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentsResponseBody) SetTotalCount(v int64) *ListExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentsResponseBodyExperiments struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator         *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Version         *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentsResponseBodyExperiments) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyExperiments) SetAccessibility(v string) *ListExperimentsResponseBodyExperiments {
	s.Accessibility = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetCreator(v string) *ListExperimentsResponseBodyExperiments {
	s.Creator = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetDescription(v string) *ListExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtCreateTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtCreateTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetGmtModifiedTime(v string) *ListExperimentsResponseBodyExperiments {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetName(v string) *ListExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetSource(v string) *ListExperimentsResponseBodyExperiments {
	s.Source = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetVersion(v int64) *ListExperimentsResponseBodyExperiments {
	s.Version = &v
	return s
}

func (s *ListExperimentsResponseBodyExperiments) SetWorkspaceId(v string) *ListExperimentsResponseBodyExperiments {
	s.WorkspaceId = &v
	return s
}

type ListExperimentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponse) SetHeaders(v map[string]*string) *ListExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentsResponse) SetStatusCode(v int32) *ListExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentsResponse) SetBody(v *ListExperimentsResponseBody) *ListExperimentsResponse {
	s.Body = v
	return s
}

type ListImageLabelsRequest struct {
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LabelFilter *string `json:"LabelFilter,omitempty" xml:"LabelFilter,omitempty"`
	LabelKeys   *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s ListImageLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListImageLabelsRequest) SetImageId(v string) *ListImageLabelsRequest {
	s.ImageId = &v
	return s
}

func (s *ListImageLabelsRequest) SetLabelFilter(v string) *ListImageLabelsRequest {
	s.LabelFilter = &v
	return s
}

func (s *ListImageLabelsRequest) SetLabelKeys(v string) *ListImageLabelsRequest {
	s.LabelKeys = &v
	return s
}

type ListImageLabelsResponseBody struct {
	Labels     []*ListImageLabelsResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	TotalCount *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponseBody) SetLabels(v []*ListImageLabelsResponseBodyLabels) *ListImageLabelsResponseBody {
	s.Labels = v
	return s
}

func (s *ListImageLabelsResponseBody) SetTotalCount(v int64) *ListImageLabelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImageLabelsResponseBody) SetRequestId(v string) *ListImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type ListImageLabelsResponseBodyLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImageLabelsResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponseBodyLabels) SetKey(v string) *ListImageLabelsResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *ListImageLabelsResponseBodyLabels) SetValue(v string) *ListImageLabelsResponseBodyLabels {
	s.Value = &v
	return s
}

type ListImageLabelsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponse) SetHeaders(v map[string]*string) *ListImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListImageLabelsResponse) SetStatusCode(v int32) *ListImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageLabelsResponse) SetBody(v *ListImageLabelsResponseBody) *ListImageLabelsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Verbose    *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetLabels(v string) *ListImagesRequest {
	s.Labels = &v
	return s
}

func (s *ListImagesRequest) SetName(v string) *ListImagesRequest {
	s.Name = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int32) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetVerbose(v bool) *ListImagesRequest {
	s.Verbose = &v
	return s
}

type ListImagesResponseBody struct {
	Images     []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	Description   *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string                               `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	ImageId       *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUri      *string                               `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels        []*ListImagesResponseBodyImagesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name          *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetGmtCreateTime(v string) *ListImagesResponseBodyImages {
	s.GmtCreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageUri(v string) *ListImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetLabels(v []*ListImagesResponseBodyImagesLabels) *ListImagesResponseBodyImages {
	s.Labels = v
	return s
}

func (s *ListImagesResponseBodyImages) SetName(v string) *ListImagesResponseBodyImages {
	s.Name = &v
	return s
}

type ListImagesResponseBodyImagesLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImagesResponseBodyImagesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesLabels) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesLabels) SetKey(v string) *ListImagesResponseBodyImagesLabels {
	s.Key = &v
	return s
}

func (s *ListImagesResponseBodyImagesLabels) SetValue(v string) *ListImagesResponseBodyImagesLabels {
	s.Value = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetCreator(v string) *ListJobsRequest {
	s.Creator = &v
	return s
}

func (s *ListJobsRequest) SetExperimentId(v string) *ListJobsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListJobsRequest) SetOrder(v string) *ListJobsRequest {
	s.Order = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

type ListJobsResponseBody struct {
	Jobs      []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	ExecuteType   *string `json:"ExecuteType,omitempty" xml:"ExecuteType,omitempty"`
	ExperimentId  *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PaiflowNodeId *string `json:"PaiflowNodeId,omitempty" xml:"PaiflowNodeId,omitempty"`
	RunId         *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetCreator(v string) *ListJobsResponseBodyJobs {
	s.Creator = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExecuteType(v string) *ListJobsResponseBodyJobs {
	s.ExecuteType = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExperimentId(v string) *ListJobsResponseBodyJobs {
	s.ExperimentId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetGmtCreateTime(v string) *ListJobsResponseBodyJobs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetNodeId(v string) *ListJobsResponseBodyJobs {
	s.NodeId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetPaiflowNodeId(v string) *ListJobsResponseBodyJobs {
	s.PaiflowNodeId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetRunId(v string) *ListJobsResponseBodyJobs {
	s.RunId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v string) *ListJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetWorkspaceId(v string) *ListJobsResponseBodyJobs {
	s.WorkspaceId = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListNodeOutputsResponseBody struct {
	Outputs   []*ListNodeOutputsResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeOutputsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeOutputsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeOutputsResponseBody) SetOutputs(v []*ListNodeOutputsResponseBodyOutputs) *ListNodeOutputsResponseBody {
	s.Outputs = v
	return s
}

func (s *ListNodeOutputsResponseBody) SetRequestId(v string) *ListNodeOutputsResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeOutputsResponseBodyOutputs struct {
	AlgoName     *string                `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	DisplayName  *string                `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	LocationType *string                `json:"LocationType,omitempty" xml:"LocationType,omitempty"`
	NodeName     *string                `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OutputId     *string                `json:"OutputId,omitempty" xml:"OutputId,omitempty"`
	OutputIndex  *string                `json:"OutputIndex,omitempty" xml:"OutputIndex,omitempty"`
	Type         *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        map[string]interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeOutputsResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeOutputsResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeOutputsResponseBodyOutputs) SetAlgoName(v string) *ListNodeOutputsResponseBodyOutputs {
	s.AlgoName = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetDisplayName(v string) *ListNodeOutputsResponseBodyOutputs {
	s.DisplayName = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetLocationType(v string) *ListNodeOutputsResponseBodyOutputs {
	s.LocationType = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetNodeName(v string) *ListNodeOutputsResponseBodyOutputs {
	s.NodeName = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetOutputId(v string) *ListNodeOutputsResponseBodyOutputs {
	s.OutputId = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetOutputIndex(v string) *ListNodeOutputsResponseBodyOutputs {
	s.OutputIndex = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetType(v string) *ListNodeOutputsResponseBodyOutputs {
	s.Type = &v
	return s
}

func (s *ListNodeOutputsResponseBodyOutputs) SetValue(v map[string]interface{}) *ListNodeOutputsResponseBodyOutputs {
	s.Value = v
	return s
}

type ListNodeOutputsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeOutputsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeOutputsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeOutputsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeOutputsResponse) SetHeaders(v map[string]*string) *ListNodeOutputsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeOutputsResponse) SetStatusCode(v int32) *ListNodeOutputsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeOutputsResponse) SetBody(v *ListNodeOutputsResponseBody) *ListNodeOutputsResponse {
	s.Body = v
	return s
}

type ListRecentExperimentsRequest struct {
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListRecentExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsRequest) SetOrder(v string) *ListRecentExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetPageNumber(v int64) *ListRecentExperimentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetPageSize(v int64) *ListRecentExperimentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetSource(v string) *ListRecentExperimentsRequest {
	s.Source = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetType(v string) *ListRecentExperimentsRequest {
	s.Type = &v
	return s
}

func (s *ListRecentExperimentsRequest) SetWorkspaceId(v string) *ListRecentExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

type ListRecentExperimentsResponseBody struct {
	Experiments []*ListRecentExperimentsResponseBodyExperiments `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecentExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsResponseBody) SetExperiments(v []*ListRecentExperimentsResponseBodyExperiments) *ListRecentExperimentsResponseBody {
	s.Experiments = v
	return s
}

func (s *ListRecentExperimentsResponseBody) SetRequestId(v string) *ListRecentExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentExperimentsResponseBody) SetTotalCount(v int64) *ListRecentExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRecentExperimentsResponseBodyExperiments struct {
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExperimentId          *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ModifyCnt             *int64  `json:"ModifyCnt,omitempty" xml:"ModifyCnt,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecentGmtModifiedTime *string `json:"RecentGmtModifiedTime,omitempty" xml:"RecentGmtModifiedTime,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	WorkspaceId           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName         *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListRecentExperimentsResponseBodyExperiments) String() string {
	return tea.Prettify(s)
}

func (s ListRecentExperimentsResponseBodyExperiments) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetDescription(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.Description = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetExperimentId(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.ExperimentId = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetModifyCnt(v int64) *ListRecentExperimentsResponseBodyExperiments {
	s.ModifyCnt = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetName(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.Name = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetRecentGmtModifiedTime(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.RecentGmtModifiedTime = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetSource(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.Source = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetWorkspaceId(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.WorkspaceId = &v
	return s
}

func (s *ListRecentExperimentsResponseBodyExperiments) SetWorkspaceName(v string) *ListRecentExperimentsResponseBodyExperiments {
	s.WorkspaceName = &v
	return s
}

type ListRecentExperimentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecentExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecentExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentExperimentsResponse) SetHeaders(v map[string]*string) *ListRecentExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentExperimentsResponse) SetStatusCode(v int32) *ListRecentExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentExperimentsResponse) SetBody(v *ListRecentExperimentsResponseBody) *ListRecentExperimentsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	Label        *string `json:"Label,omitempty" xml:"Label,omitempty"`
	List         *string `json:"List,omitempty" xml:"List,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TagId        *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TypeId       *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
	Verbose      *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetLabel(v string) *ListTemplatesRequest {
	s.Label = &v
	return s
}

func (s *ListTemplatesRequest) SetList(v string) *ListTemplatesRequest {
	s.List = &v
	return s
}

func (s *ListTemplatesRequest) SetName(v string) *ListTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequest) SetOrder(v string) *ListTemplatesRequest {
	s.Order = &v
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

func (s *ListTemplatesRequest) SetSortBy(v string) *ListTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListTemplatesRequest) SetSource(v string) *ListTemplatesRequest {
	s.Source = &v
	return s
}

func (s *ListTemplatesRequest) SetTagId(v string) *ListTemplatesRequest {
	s.TagId = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesRequest) SetTypeId(v string) *ListTemplatesRequest {
	s.TypeId = &v
	return s
}

func (s *ListTemplatesRequest) SetVerbose(v bool) *ListTemplatesRequest {
	s.Verbose = &v
	return s
}

func (s *ListTemplatesRequest) SetWorkspaceId(v string) *ListTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

type ListTemplatesResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateData []*ListTemplatesResponseBodyTemplateData `json:"TemplateData,omitempty" xml:"TemplateData,omitempty" type:"Repeated"`
	TotalCount   *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplateData(v []*ListTemplatesResponseBodyTemplateData) *ListTemplatesResponseBody {
	s.TemplateData = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyTemplateData struct {
	Template     *ListTemplatesResponseBodyTemplateDataTemplate     `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	TemplateTag  *ListTemplatesResponseBodyTemplateDataTemplateTag  `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Struct"`
	TemplateType *ListTemplatesResponseBodyTemplateDataTemplateType `json:"TemplateType,omitempty" xml:"TemplateType,omitempty" type:"Struct"`
}

func (s ListTemplatesResponseBodyTemplateData) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateData) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateData) SetTemplate(v *ListTemplatesResponseBodyTemplateDataTemplate) *ListTemplatesResponseBodyTemplateData {
	s.Template = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateData) SetTemplateTag(v *ListTemplatesResponseBodyTemplateDataTemplateTag) *ListTemplatesResponseBodyTemplateData {
	s.TemplateTag = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateData) SetTemplateType(v *ListTemplatesResponseBodyTemplateDataTemplateType) *ListTemplatesResponseBodyTemplateData {
	s.TemplateType = v
	return s
}

type ListTemplatesResponseBodyTemplateDataTemplate struct {
	Content         *string                  `json:"Content,omitempty" xml:"Content,omitempty"`
	Creator         *string                  `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail          *string                  `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DocLink         *string                  `json:"DocLink,omitempty" xml:"DocLink,omitempty"`
	GmtCreateTime   *string                  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageLink       *string                  `json:"ImageLink,omitempty" xml:"ImageLink,omitempty"`
	Labels          []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId      *string                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplateDataTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateDataTemplate) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetContent(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Content = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetCreator(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Creator = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetDescription(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetDetail(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Detail = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetDocLink(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.DocLink = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetGmtCreateTime(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetGmtModifiedTime(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetImageLink(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.ImageLink = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetLabels(v []map[string]interface{}) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Labels = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetName(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetTemplateId(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.TemplateId = &v
	return s
}

type ListTemplatesResponseBodyTemplateDataTemplateTag struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagId  *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TypeId *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplateDataTemplateTag) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateDataTemplateTag) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) SetName(v string) *ListTemplatesResponseBodyTemplateDataTemplateTag {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) SetTagId(v string) *ListTemplatesResponseBodyTemplateDataTemplateTag {
	s.TagId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) SetTypeId(v string) *ListTemplatesResponseBodyTemplateDataTemplateTag {
	s.TypeId = &v
	return s
}

type ListTemplatesResponseBodyTemplateDataTemplateType struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TypeId *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplateDataTemplateType) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateDataTemplateType) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) SetName(v string) *ListTemplatesResponseBodyTemplateDataTemplateType {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) SetTypeId(v string) *ListTemplatesResponseBodyTemplateDataTemplateType {
	s.TypeId = &v
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

type MigrateExperimentFoldersRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	IsOwner       *bool   `json:"IsOwner,omitempty" xml:"IsOwner,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MigrateExperimentFoldersRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentFoldersRequest) GoString() string {
	return s.String()
}

func (s *MigrateExperimentFoldersRequest) SetAccessibility(v string) *MigrateExperimentFoldersRequest {
	s.Accessibility = &v
	return s
}

func (s *MigrateExperimentFoldersRequest) SetIsOwner(v bool) *MigrateExperimentFoldersRequest {
	s.IsOwner = &v
	return s
}

func (s *MigrateExperimentFoldersRequest) SetWorkspaceId(v string) *MigrateExperimentFoldersRequest {
	s.WorkspaceId = &v
	return s
}

type MigrateExperimentFoldersResponseBody struct {
	Code            *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	FolderIdMapping map[string]map[string]interface{} `json:"FolderIdMapping,omitempty" xml:"FolderIdMapping,omitempty"`
	Message         *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateExperimentFoldersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateExperimentFoldersResponseBody) SetCode(v string) *MigrateExperimentFoldersResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) SetFolderIdMapping(v map[string]map[string]interface{}) *MigrateExperimentFoldersResponseBody {
	s.FolderIdMapping = v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) SetMessage(v string) *MigrateExperimentFoldersResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) SetRequestId(v string) *MigrateExperimentFoldersResponseBody {
	s.RequestId = &v
	return s
}

type MigrateExperimentFoldersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MigrateExperimentFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateExperimentFoldersResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentFoldersResponse) GoString() string {
	return s.String()
}

func (s *MigrateExperimentFoldersResponse) SetHeaders(v map[string]*string) *MigrateExperimentFoldersResponse {
	s.Headers = v
	return s
}

func (s *MigrateExperimentFoldersResponse) SetStatusCode(v int32) *MigrateExperimentFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateExperimentFoldersResponse) SetBody(v *MigrateExperimentFoldersResponseBody) *MigrateExperimentFoldersResponse {
	s.Body = v
	return s
}

type MigrateExperimentsRequest struct {
	Accessibility  *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	DestFolderId   *string `json:"DestFolderId,omitempty" xml:"DestFolderId,omitempty"`
	IsOwner        *bool   `json:"IsOwner,omitempty" xml:"IsOwner,omitempty"`
	SourceExpId    *int64  `json:"SourceExpId,omitempty" xml:"SourceExpId,omitempty"`
	UpdateIfExists *bool   `json:"UpdateIfExists,omitempty" xml:"UpdateIfExists,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MigrateExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentsRequest) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsRequest) SetAccessibility(v string) *MigrateExperimentsRequest {
	s.Accessibility = &v
	return s
}

func (s *MigrateExperimentsRequest) SetDestFolderId(v string) *MigrateExperimentsRequest {
	s.DestFolderId = &v
	return s
}

func (s *MigrateExperimentsRequest) SetIsOwner(v bool) *MigrateExperimentsRequest {
	s.IsOwner = &v
	return s
}

func (s *MigrateExperimentsRequest) SetSourceExpId(v int64) *MigrateExperimentsRequest {
	s.SourceExpId = &v
	return s
}

func (s *MigrateExperimentsRequest) SetUpdateIfExists(v bool) *MigrateExperimentsRequest {
	s.UpdateIfExists = &v
	return s
}

func (s *MigrateExperimentsRequest) SetWorkspaceId(v string) *MigrateExperimentsRequest {
	s.WorkspaceId = &v
	return s
}

type MigrateExperimentsResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MigrateExperimentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsResponseBody) SetCode(v string) *MigrateExperimentsResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateExperimentsResponseBody) SetData(v *MigrateExperimentsResponseBodyData) *MigrateExperimentsResponseBody {
	s.Data = v
	return s
}

func (s *MigrateExperimentsResponseBody) SetMessage(v string) *MigrateExperimentsResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateExperimentsResponseBody) SetRequestId(v string) *MigrateExperimentsResponseBody {
	s.RequestId = &v
	return s
}

type MigrateExperimentsResponseBodyData struct {
	AlreadyExists *bool   `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
	ExperimentId  *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Updated       *bool   `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s MigrateExperimentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsResponseBodyData) SetAlreadyExists(v bool) *MigrateExperimentsResponseBodyData {
	s.AlreadyExists = &v
	return s
}

func (s *MigrateExperimentsResponseBodyData) SetExperimentId(v string) *MigrateExperimentsResponseBodyData {
	s.ExperimentId = &v
	return s
}

func (s *MigrateExperimentsResponseBodyData) SetUpdated(v bool) *MigrateExperimentsResponseBodyData {
	s.Updated = &v
	return s
}

type MigrateExperimentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MigrateExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateExperimentsResponse) GoString() string {
	return s.String()
}

func (s *MigrateExperimentsResponse) SetHeaders(v map[string]*string) *MigrateExperimentsResponse {
	s.Headers = v
	return s
}

func (s *MigrateExperimentsResponse) SetStatusCode(v int32) *MigrateExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateExperimentsResponse) SetBody(v *MigrateExperimentsResponseBody) *MigrateExperimentsResponse {
	s.Body = v
	return s
}

type PreviewMCTableRequest struct {
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Limit       *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Partition   *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PreviewMCTableRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewMCTableRequest) GoString() string {
	return s.String()
}

func (s *PreviewMCTableRequest) SetEndpoint(v string) *PreviewMCTableRequest {
	s.Endpoint = &v
	return s
}

func (s *PreviewMCTableRequest) SetLimit(v int32) *PreviewMCTableRequest {
	s.Limit = &v
	return s
}

func (s *PreviewMCTableRequest) SetPartition(v string) *PreviewMCTableRequest {
	s.Partition = &v
	return s
}

func (s *PreviewMCTableRequest) SetWorkspaceId(v string) *PreviewMCTableRequest {
	s.WorkspaceId = &v
	return s
}

type PreviewMCTableResponseBody struct {
	Content   [][]*string `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreviewMCTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewMCTableResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewMCTableResponseBody) SetContent(v [][]*string) *PreviewMCTableResponseBody {
	s.Content = v
	return s
}

func (s *PreviewMCTableResponseBody) SetRequestId(v string) *PreviewMCTableResponseBody {
	s.RequestId = &v
	return s
}

type PreviewMCTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PreviewMCTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviewMCTableResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewMCTableResponse) GoString() string {
	return s.String()
}

func (s *PreviewMCTableResponse) SetHeaders(v map[string]*string) *PreviewMCTableResponse {
	s.Headers = v
	return s
}

func (s *PreviewMCTableResponse) SetStatusCode(v int32) *PreviewMCTableResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewMCTableResponse) SetBody(v *PreviewMCTableResponseBody) *PreviewMCTableResponse {
	s.Body = v
	return s
}

type PublishExperimentRequest struct {
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
}

func (s PublishExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishExperimentRequest) GoString() string {
	return s.String()
}

func (s *PublishExperimentRequest) SetFolderId(v string) *PublishExperimentRequest {
	s.FolderId = &v
	return s
}

type PublishExperimentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *PublishExperimentResponseBody) SetRequestId(v string) *PublishExperimentResponseBody {
	s.RequestId = &v
	return s
}

type PublishExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishExperimentResponse) GoString() string {
	return s.String()
}

func (s *PublishExperimentResponse) SetHeaders(v map[string]*string) *PublishExperimentResponse {
	s.Headers = v
	return s
}

func (s *PublishExperimentResponse) SetStatusCode(v int32) *PublishExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishExperimentResponse) SetBody(v *PublishExperimentResponseBody) *PublishExperimentResponse {
	s.Body = v
	return s
}

type QueryExperimentVisualizationDataRequest struct {
	Body []*QueryExperimentVisualizationDataRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s QueryExperimentVisualizationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryExperimentVisualizationDataRequest) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataRequest) SetBody(v []*QueryExperimentVisualizationDataRequestBody) *QueryExperimentVisualizationDataRequest {
	s.Body = v
	return s
}

type QueryExperimentVisualizationDataRequestBody struct {
	EndTime              *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NodeId               *string   `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	StartTime            *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VisualizationDataIds []*string `json:"VisualizationDataIds,omitempty" xml:"VisualizationDataIds,omitempty" type:"Repeated"`
}

func (s QueryExperimentVisualizationDataRequestBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExperimentVisualizationDataRequestBody) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataRequestBody) SetEndTime(v string) *QueryExperimentVisualizationDataRequestBody {
	s.EndTime = &v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) SetNodeId(v string) *QueryExperimentVisualizationDataRequestBody {
	s.NodeId = &v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) SetStartTime(v string) *QueryExperimentVisualizationDataRequestBody {
	s.StartTime = &v
	return s
}

func (s *QueryExperimentVisualizationDataRequestBody) SetVisualizationDataIds(v []*string) *QueryExperimentVisualizationDataRequestBody {
	s.VisualizationDataIds = v
	return s
}

type QueryExperimentVisualizationDataResponseBody struct {
	VisualizationData []*QueryExperimentVisualizationDataResponseBodyVisualizationData `json:"VisualizationData,omitempty" xml:"VisualizationData,omitempty" type:"Repeated"`
	RequestId         *string                                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryExperimentVisualizationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExperimentVisualizationDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataResponseBody) SetVisualizationData(v []*QueryExperimentVisualizationDataResponseBodyVisualizationData) *QueryExperimentVisualizationDataResponseBody {
	s.VisualizationData = v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBody) SetRequestId(v string) *QueryExperimentVisualizationDataResponseBody {
	s.RequestId = &v
	return s
}

type QueryExperimentVisualizationDataResponseBodyVisualizationData struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	NodeId     *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s QueryExperimentVisualizationDataResponseBodyVisualizationData) String() string {
	return tea.Prettify(s)
}

func (s QueryExperimentVisualizationDataResponseBodyVisualizationData) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetCreateTime(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.CreateTime = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetData(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.Data = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetDataId(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.DataId = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponseBodyVisualizationData) SetNodeId(v string) *QueryExperimentVisualizationDataResponseBodyVisualizationData {
	s.NodeId = &v
	return s
}

type QueryExperimentVisualizationDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryExperimentVisualizationDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryExperimentVisualizationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExperimentVisualizationDataResponse) GoString() string {
	return s.String()
}

func (s *QueryExperimentVisualizationDataResponse) SetHeaders(v map[string]*string) *QueryExperimentVisualizationDataResponse {
	s.Headers = v
	return s
}

func (s *QueryExperimentVisualizationDataResponse) SetStatusCode(v int32) *QueryExperimentVisualizationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExperimentVisualizationDataResponse) SetBody(v *QueryExperimentVisualizationDataResponseBody) *QueryExperimentVisualizationDataResponse {
	s.Body = v
	return s
}

type RemoveImageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageResponseBody) SetRequestId(v string) *RemoveImageResponseBody {
	s.RequestId = &v
	return s
}

type RemoveImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageResponse) SetHeaders(v map[string]*string) *RemoveImageResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageResponse) SetStatusCode(v int32) *RemoveImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageResponse) SetBody(v *RemoveImageResponseBody) *RemoveImageResponse {
	s.Body = v
	return s
}

type RemoveImageLabelsResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageLabelsResponseBody) SetRequestId(v string) *RemoveImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveImageLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageLabelsResponse) SetHeaders(v map[string]*string) *RemoveImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageLabelsResponse) SetStatusCode(v int32) *RemoveImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageLabelsResponse) SetBody(v *RemoveImageLabelsResponseBody) *RemoveImageLabelsResponse {
	s.Body = v
	return s
}

type SearchMCTablesRequest struct {
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchMCTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchMCTablesRequest) GoString() string {
	return s.String()
}

func (s *SearchMCTablesRequest) SetKeyword(v string) *SearchMCTablesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchMCTablesRequest) SetWorkspaceId(v string) *SearchMCTablesRequest {
	s.WorkspaceId = &v
	return s
}

type SearchMCTablesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s SearchMCTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchMCTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMCTablesResponseBody) SetRequestId(v string) *SearchMCTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMCTablesResponseBody) SetTables(v []*string) *SearchMCTablesResponseBody {
	s.Tables = v
	return s
}

type SearchMCTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchMCTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchMCTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchMCTablesResponse) GoString() string {
	return s.String()
}

func (s *SearchMCTablesResponse) SetHeaders(v map[string]*string) *SearchMCTablesResponse {
	s.Headers = v
	return s
}

func (s *SearchMCTablesResponse) SetStatusCode(v int32) *SearchMCTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMCTablesResponse) SetBody(v *SearchMCTablesResponseBody) *SearchMCTablesResponse {
	s.Body = v
	return s
}

type StopExperimentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *StopExperimentResponseBody) SetRequestId(v string) *StopExperimentResponseBody {
	s.RequestId = &v
	return s
}

type StopExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentResponse) GoString() string {
	return s.String()
}

func (s *StopExperimentResponse) SetHeaders(v map[string]*string) *StopExperimentResponse {
	s.Headers = v
	return s
}

func (s *StopExperimentResponse) SetStatusCode(v int32) *StopExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopExperimentResponse) SetBody(v *StopExperimentResponseBody) *StopExperimentResponse {
	s.Body = v
	return s
}

type StopJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

type StopJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponse) GoString() string {
	return s.String()
}

func (s *StopJobResponse) SetHeaders(v map[string]*string) *StopJobResponse {
	s.Headers = v
	return s
}

func (s *StopJobResponse) SetStatusCode(v int32) *StopJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobResponse) SetBody(v *StopJobResponseBody) *StopJobResponse {
	s.Body = v
	return s
}

type UpdateExperimentContentRequest struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Version *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateExperimentContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentContentRequest) SetContent(v string) *UpdateExperimentContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateExperimentContentRequest) SetVersion(v int64) *UpdateExperimentContentRequest {
	s.Version = &v
	return s
}

type UpdateExperimentContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version   *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateExperimentContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentContentResponseBody) SetRequestId(v string) *UpdateExperimentContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentContentResponseBody) SetVersion(v int64) *UpdateExperimentContentResponseBody {
	s.Version = &v
	return s
}

type UpdateExperimentContentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateExperimentContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExperimentContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentContentResponse) SetHeaders(v map[string]*string) *UpdateExperimentContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentContentResponse) SetStatusCode(v int32) *UpdateExperimentContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentContentResponse) SetBody(v *UpdateExperimentContentResponseBody) *UpdateExperimentContentResponse {
	s.Body = v
	return s
}

type UpdateExperimentFolderRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s UpdateExperimentFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentFolderRequest) SetName(v string) *UpdateExperimentFolderRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentFolderRequest) SetParentFolderId(v string) *UpdateExperimentFolderRequest {
	s.ParentFolderId = &v
	return s
}

type UpdateExperimentFolderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentFolderResponseBody) SetRequestId(v string) *UpdateExperimentFolderResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExperimentFolderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateExperimentFolderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExperimentFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentFolderResponse) SetHeaders(v map[string]*string) *UpdateExperimentFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentFolderResponse) SetStatusCode(v int32) *UpdateExperimentFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentFolderResponse) SetBody(v *UpdateExperimentFolderResponseBody) *UpdateExperimentFolderResponse {
	s.Body = v
	return s
}

type UpdateExperimentMetaRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FolderId      *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options       *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateExperimentMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentMetaRequest) SetAccessibility(v string) *UpdateExperimentMetaRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetDescription(v string) *UpdateExperimentMetaRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetFolderId(v string) *UpdateExperimentMetaRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetName(v string) *UpdateExperimentMetaRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentMetaRequest) SetOptions(v string) *UpdateExperimentMetaRequest {
	s.Options = &v
	return s
}

type UpdateExperimentMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentMetaResponseBody) SetRequestId(v string) *UpdateExperimentMetaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExperimentMetaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateExperimentMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExperimentMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentMetaResponse) SetHeaders(v map[string]*string) *UpdateExperimentMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentMetaResponse) SetStatusCode(v int32) *UpdateExperimentMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentMetaResponse) SetBody(v *UpdateExperimentMetaResponseBody) *UpdateExperimentMetaResponse {
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
		"cn-beijing":     tea.String("pai.cn-beijing.aliyuncs.com"),
		"cn-hangzhou":    tea.String("pai.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":    tea.String("pai.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":    tea.String("pai.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":    tea.String("pai.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1": tea.String("pai.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": tea.String("pai.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3": tea.String("pai.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5": tea.String("pai.ap-southeast-5.aliyuncs.com"),
		"us-west-1":      tea.String("pai.us-west-1.aliyuncs.com"),
		"us-east-1":      tea.String("pai.us-east-1.aliyuncs.com"),
		"eu-central-1":   tea.String("pai.eu-central-1.aliyuncs.com"),
		"me-east-1":      tea.String("pai.me-east-1.aliyuncs.com"),
		"ap-south-1":     tea.String("pai.ap-south-1.aliyuncs.com"),
		"cn-qingdao":     tea.String("pai.cn-qingdao.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("pai.cn-zhangjiakou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("paistudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddImageWithOptions(request *AddImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		body["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageLabelsWithOptions(ImageId *string, request *AddImageLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImageLabels"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImageLabels(ImageId *string, request *AddImageLabelsRequest) (_result *AddImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageLabelsResponse{}
	_body, _err := client.AddImageLabelsWithOptions(ImageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyExperimentWithOptions(ExperimentId *string, request *CopyExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CopyExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		body["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyExperiment"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/%5BExperimentId%5D/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyExperiment(ExperimentId *string, request *CopyExperimentRequest) (_result *CopyExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CopyExperimentResponse{}
	_body, _err := client.CopyExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExperimentWithOptions(request *CreateExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		body["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperiment"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExperiment(request *CreateExperimentRequest) (_result *CreateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentResponse{}
	_body, _err := client.CreateExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExperimentFolderWithOptions(request *CreateExperimentFolderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		body["ParentFolderId"] = request.ParentFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperimentFolder"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentfolders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExperimentFolder(request *CreateExperimentFolderRequest) (_result *CreateExperimentFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentFolderResponse{}
	_body, _err := client.CreateExperimentFolderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExperimentMigrateValidationWithOptions(request *CreateExperimentMigrateValidationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentMigrateValidationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceExpId)) {
		query["SourceExpId"] = request.SourceExpId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperimentMigrateValidation"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/migrate/experimentvalidation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentMigrateValidationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExperimentMigrateValidation(request *CreateExperimentMigrateValidationRequest) (_result *CreateExperimentMigrateValidationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentMigrateValidationResponse{}
	_body, _err := client.CreateExperimentMigrateValidationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecuteType)) {
		body["ExecuteType"] = request.ExecuteType
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		body["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperiment"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExperiment(ExperimentId *string) (_result *DeleteExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExperimentFolderWithOptions(FolderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentFolderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperimentFolder"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentfolders/" + tea.StringValue(openapiutil.GetEncodeParam(FolderId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExperimentFolder(FolderId *string) (_result *DeleteExperimentFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentFolderResponse{}
	_body, _err := client.DeleteExperimentFolderWithOptions(FolderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgoTreeWithOptions(request *GetAlgoTreeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgoTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgoTree"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algo/tree"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlgoTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgoTree(request *GetAlgoTreeRequest) (_result *GetAlgoTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgoTreeResponse{}
	_body, _err := client.GetAlgoTreeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmDefWithOptions(request *GetAlgorithmDefRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgorithmDefResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgoVersion)) {
		query["AlgoVersion"] = request.AlgoVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgorithmDef"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithm/def"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlgorithmDefResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmDef(request *GetAlgorithmDefRequest) (_result *GetAlgorithmDefResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmDefResponse{}
	_body, _err := client.GetAlgorithmDefWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmDefsWithOptions(request *GetAlgorithmDefsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgorithmDefsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LatestTimestamp)) {
		query["LatestTimestamp"] = request.LatestTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.RangeEnd)) {
		query["RangeEnd"] = request.RangeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.RangeStart)) {
		query["RangeStart"] = request.RangeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["Timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgorithmDefs"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithm/defs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlgorithmDefsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmDefs(request *GetAlgorithmDefsRequest) (_result *GetAlgorithmDefsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmDefsResponse{}
	_body, _err := client.GetAlgorithmDefsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlgorithmTreeWithOptions(request *GetAlgorithmTreeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAlgorithmTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlgorithmTree"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algorithm/tree"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlgorithmTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlgorithmTree(request *GetAlgorithmTreeRequest) (_result *GetAlgorithmTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmTreeResponse{}
	_body, _err := client.GetAlgorithmTreeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperiment"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperiment(ExperimentId *string) (_result *GetExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentFolderChildrenWithOptions(FolderId *string, request *GetExperimentFolderChildrenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentFolderChildrenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyFolder)) {
		query["OnlyFolder"] = request.OnlyFolder
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentFolderChildren"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentfolders/" + tea.StringValue(openapiutil.GetEncodeParam(FolderId)) + "/children"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentFolderChildrenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentFolderChildren(FolderId *string, request *GetExperimentFolderChildrenRequest) (_result *GetExperimentFolderChildrenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentFolderChildrenResponse{}
	_body, _err := client.GetExperimentFolderChildrenWithOptions(FolderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentMetaWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentMetaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentMeta"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/meta"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentMeta(ExperimentId *string) (_result *GetExperimentMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentMetaResponse{}
	_body, _err := client.GetExperimentMetaWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentStatusWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentStatus"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentStatus(ExperimentId *string) (_result *GetExperimentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentStatusResponse{}
	_body, _err := client.GetExperimentStatusWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentVisualizationMetaWithOptions(ExperimentId *string, request *GetExperimentVisualizationMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentVisualizationMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		query["NodeIds"] = request.NodeIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentVisualizationMeta"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/visualizationMeta"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentVisualizationMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentVisualizationMeta(ExperimentId *string, request *GetExperimentVisualizationMetaRequest) (_result *GetExperimentVisualizationMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentVisualizationMetaResponse{}
	_body, _err := client.GetExperimentVisualizationMetaWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentsStatisticsWithOptions(request *GetExperimentsStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentsStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentsStatistics"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentsStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentsStatistics(request *GetExperimentsStatisticsRequest) (_result *GetExperimentsStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentsStatisticsResponse{}
	_body, _err := client.GetExperimentsStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExperimentsUsersStatisticsWithOptions(request *GetExperimentsUsersStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentsUsersStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentsUsersStatistics"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/experimentsusers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentsUsersStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExperimentsUsersStatistics(request *GetExperimentsUsersStatisticsRequest) (_result *GetExperimentsUsersStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentsUsersStatisticsResponse{}
	_body, _err := client.GetExperimentsUsersStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(ImageId *string, request *GetImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(ImageId *string, request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(ImageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(JobId *string, request *GetJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(JobId *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMCTableSchemaWithOptions(TableName *string, request *GetMCTableSchemaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMCTableSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMCTableSchema"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources/maxcompute/tables/" + tea.StringValue(openapiutil.GetEncodeParam(TableName)) + "/schema"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMCTableSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMCTableSchema(TableName *string, request *GetMCTableSchemaRequest) (_result *GetMCTableSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMCTableSchemaResponse{}
	_body, _err := client.GetMCTableSchemaWithOptions(TableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeInputSchemaWithOptions(ExperimentId *string, NodeId *string, request *GetNodeInputSchemaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNodeInputSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputId)) {
		query["InputId"] = request.InputId
	}

	if !tea.BoolValue(util.IsUnset(request.InputIndex)) {
		query["InputIndex"] = request.InputIndex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeInputSchema"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/schema"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeInputSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeInputSchema(ExperimentId *string, NodeId *string, request *GetNodeInputSchemaRequest) (_result *GetNodeInputSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeInputSchemaResponse{}
	_body, _err := client.GetNodeInputSchemaWithOptions(ExperimentId, NodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeOutputWithOptions(ExperimentId *string, NodeId *string, OutputId *string, request *GetNodeOutputRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNodeOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputIndex)) {
		query["OutputIndex"] = request.OutputIndex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeOutput"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/outputs/" + tea.StringValue(openapiutil.GetEncodeParam(OutputId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeOutput(ExperimentId *string, NodeId *string, OutputId *string, request *GetNodeOutputRequest) (_result *GetNodeOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeOutputResponse{}
	_body, _err := client.GetNodeOutputWithOptions(ExperimentId, NodeId, OutputId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeVisualizationWithOptions(ExperimentId *string, NodeId *string, VisualizationId *string, request *GetNodeVisualizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNodeVisualizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeVisualization"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/visualizations/" + tea.StringValue(openapiutil.GetEncodeParam(VisualizationId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeVisualizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeVisualization(ExperimentId *string, NodeId *string, VisualizationId *string, request *GetNodeVisualizationRequest) (_result *GetNodeVisualizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeVisualizationResponse{}
	_body, _err := client.GetNodeVisualizationWithOptions(ExperimentId, NodeId, VisualizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(TemplateId *string, request *GetTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/templates/" + tea.StringValue(openapiutil.GetEncodeParam(TemplateId))),
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

func (client *Client) GetTemplate(TemplateId *string, request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(TemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlgoDefsWithOptions(request *ListAlgoDefsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlgoDefsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlgoDefs"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/algo/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlgoDefsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlgoDefs(request *ListAlgoDefsRequest) (_result *ListAlgoDefsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlgoDefsResponse{}
	_body, _err := client.ListAlgoDefsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAuthRolesWithOptions(request *ListAuthRolesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAuthRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsGenerateToken)) {
		query["IsGenerateToken"] = request.IsGenerateToken
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthRoles"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/authorization/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAuthRoles(request *ListAuthRolesRequest) (_result *ListAuthRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAuthRolesResponse{}
	_body, _err := client.ListAuthRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExperimentsWithOptions(request *ListExperimentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExperimentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperiments"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExperiments(request *ListExperimentsRequest) (_result *ListExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentsResponse{}
	_body, _err := client.ListExperimentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImageLabelsWithOptions(request *ListImageLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImageLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelFilter)) {
		query["LabelFilter"] = request.LabelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImageLabels"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/image/labels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImageLabels(request *ListImageLabelsRequest) (_result *ListImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImageLabelsResponse{}
	_body, _err := client.ListImageLabelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeOutputsWithOptions(ExperimentId *string, NodeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNodeOutputsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeOutputs"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/outputs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeOutputsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeOutputs(ExperimentId *string, NodeId *string) (_result *ListNodeOutputsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodeOutputsResponse{}
	_body, _err := client.ListNodeOutputsWithOptions(ExperimentId, NodeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecentExperimentsWithOptions(request *ListRecentExperimentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRecentExperimentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecentExperiments"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/recentexperiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecentExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecentExperiments(request *ListRecentExperimentsRequest) (_result *ListRecentExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecentExperimentsResponse{}
	_body, _err := client.ListRecentExperimentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.List)) {
		query["List"] = request.List
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	if !tea.BoolValue(util.IsUnset(request.TypeId)) {
		query["TypeId"] = request.TypeId
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/templates"),
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

func (client *Client) MigrateExperimentFoldersWithOptions(request *MigrateExperimentFoldersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MigrateExperimentFoldersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.IsOwner)) {
		query["IsOwner"] = request.IsOwner
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MigrateExperimentFolders"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/migrate/folders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MigrateExperimentFoldersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateExperimentFolders(request *MigrateExperimentFoldersRequest) (_result *MigrateExperimentFoldersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateExperimentFoldersResponse{}
	_body, _err := client.MigrateExperimentFoldersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateExperimentsWithOptions(request *MigrateExperimentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MigrateExperimentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.DestFolderId)) {
		query["DestFolderId"] = request.DestFolderId
	}

	if !tea.BoolValue(util.IsUnset(request.IsOwner)) {
		query["IsOwner"] = request.IsOwner
	}

	if !tea.BoolValue(util.IsUnset(request.SourceExpId)) {
		query["SourceExpId"] = request.SourceExpId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateIfExists)) {
		query["UpdateIfExists"] = request.UpdateIfExists
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MigrateExperiments"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/migrate/experiments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MigrateExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateExperiments(request *MigrateExperimentsRequest) (_result *MigrateExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateExperimentsResponse{}
	_body, _err := client.MigrateExperimentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreviewMCTableWithOptions(TableName *string, request *PreviewMCTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PreviewMCTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		query["Partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviewMCTable"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources/maxcompute/tables/" + tea.StringValue(openapiutil.GetEncodeParam(TableName)) + "/preview"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PreviewMCTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreviewMCTable(TableName *string, request *PreviewMCTableRequest) (_result *PreviewMCTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewMCTableResponse{}
	_body, _err := client.PreviewMCTableWithOptions(TableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishExperimentWithOptions(ExperimentId *string, request *PublishExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		body["FolderId"] = request.FolderId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishExperiment"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/publish"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishExperiment(ExperimentId *string, request *PublishExperimentRequest) (_result *PublishExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishExperimentResponse{}
	_body, _err := client.PublishExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryExperimentVisualizationDataWithOptions(ExperimentId *string, request *QueryExperimentVisualizationDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryExperimentVisualizationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryExperimentVisualizationData"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/visualizationDataQuery"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryExperimentVisualizationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryExperimentVisualizationData(ExperimentId *string, request *QueryExperimentVisualizationDataRequest) (_result *QueryExperimentVisualizationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryExperimentVisualizationDataResponse{}
	_body, _err := client.QueryExperimentVisualizationDataWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageWithOptions(ImageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImage"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImage(ImageId *string) (_result *RemoveImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveImageResponse{}
	_body, _err := client.RemoveImageWithOptions(ImageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageLabelsWithOptions(ImageId *string, LabelKey *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveImageLabelsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImageLabels"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId)) + "/labels/" + tea.StringValue(openapiutil.GetEncodeParam(LabelKey))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImageLabels(ImageId *string, LabelKey *string) (_result *RemoveImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveImageLabelsResponse{}
	_body, _err := client.RemoveImageLabelsWithOptions(ImageId, LabelKey, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchMCTablesWithOptions(request *SearchMCTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchMCTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchMCTables"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources/maxcompute/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchMCTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchMCTables(request *SearchMCTablesRequest) (_result *SearchMCTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchMCTablesResponse{}
	_body, _err := client.SearchMCTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopExperiment"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopExperiment(ExperimentId *string) (_result *StopExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopExperimentResponse{}
	_body, _err := client.StopExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobWithOptions(JobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopJob"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJob(JobId *string) (_result *StopJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExperimentContentWithOptions(ExperimentId *string, request *UpdateExperimentContentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentContent"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/content"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExperimentContent(ExperimentId *string, request *UpdateExperimentContentRequest) (_result *UpdateExperimentContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentContentResponse{}
	_body, _err := client.UpdateExperimentContentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExperimentFolderWithOptions(FolderId *string, request *UpdateExperimentFolderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFolderId)) {
		body["ParentFolderId"] = request.ParentFolderId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentFolder"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experimentfolders/" + tea.StringValue(openapiutil.GetEncodeParam(FolderId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExperimentFolder(FolderId *string, request *UpdateExperimentFolderRequest) (_result *UpdateExperimentFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentFolderResponse{}
	_body, _err := client.UpdateExperimentFolderWithOptions(FolderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExperimentMetaWithOptions(ExperimentId *string, request *UpdateExperimentMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		body["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentMeta"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/meta"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExperimentMeta(ExperimentId *string, request *UpdateExperimentMetaRequest) (_result *UpdateExperimentMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentMetaResponse{}
	_body, _err := client.UpdateExperimentMetaWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
